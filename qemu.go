package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/digitalocean/go-qemu/qemu"
	"github.com/digitalocean/go-qemu/qmp"
)

const QmpMonitor = "/tmp/qmp-sock"
const DomainName = "devmachines"

type VMParams interface {
	Memory() uint
	CPU() uint
}

func StartVM(ctx context.Context, params VMParams, out io.Writer, extraArgs []string) error {
	args := []string{
		"-m", fmt.Sprintf("%d", params.Memory()),
		"-smp", fmt.Sprintf("%d", params.CPU()),
		"-qmp", fmt.Sprintf("unix:%s,server,wait=off", QmpMonitor),
		"-enable-kvm",
		"-nographic",
	}

	args = append(args, extraArgs...)

	cmd := exec.Command("qemu-system-x86_64", args...)
	cmd.Stdout = out
	cmd.Stderr = out

	if err := cmd.Start(); err != nil {
		return fmt.Errorf("failed to start QEMU: %w", err)
	}

	if err := waitForSocket(QmpMonitor, 5*time.Second); err != nil {
		return fmt.Errorf("socket not ready: %w", err)
	}

	mon, err := qmp.NewSocketMonitor("unix", QmpMonitor, 2*time.Second)
	if err != nil {
		return fmt.Errorf("failed to connect to monitor socket: %v", err)
	}

	if err := mon.Connect(); err != nil {
		return fmt.Errorf("failed to connect to monitor: %v", err)
	}

	domain, err := qemu.NewDomain(mon, DomainName)
	if err != nil {
		return fmt.Errorf("failed to create domain object: %v", err)
	}
	defer domain.Close()

	done := make(chan error, 1)
	go func() {
		done <- cmd.Wait()
	}()

	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		err = domain.SystemPowerdown()
		if err != nil {
			log.Fatalf("Unable to power down domain: %v\n", err)
		}

		select {
		case <-done:
			return nil
		case <-time.After(40 * time.Second):
			if err := cmd.Process.Kill(); err != nil {
				log.Printf("Error killing process: %v", err)
			}
			<-done
			return ctx.Err()
		}
	}
}

func waitForSocket(socketPath string, timeout time.Duration) error {
	start := time.Now()
	for {
		if _, err := os.Stat(socketPath); err == nil {
			return nil
		}
		if time.Since(start) > timeout {
			return fmt.Errorf("timeout waiting for socket")
		}
		time.Sleep(100 * time.Millisecond)
	}
}
