package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os/exec"
	"syscall"
	"time"
)

type ArgsFunc func() []string

type VMParams interface {
	Memory() uint
	CPU() uint
}

func StartVM(ctx context.Context, params VMParams, out io.Writer, extraArgs []string) error {
	args := []string{
		"-m", fmt.Sprintf("%d", params.Memory()),
		"-smp", fmt.Sprintf("%d", params.CPU()),
		"-enable-kvm",
		"-nographic",
	}

	args = append(args, extraArgs...)

	cmd := exec.CommandContext(ctx, "qemu-system-x86_64", args...)
	cmd.Stdout = out
	cmd.Stderr = out

	if err := cmd.Start(); err != nil {
		return fmt.Errorf("failed to start QEMU: %w", err)
	}

	done := make(chan error, 1)
	go func() {
		done <- cmd.Wait()
	}()

	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		if err := cmd.Process.Signal(syscall.SIGTERM); err != nil {
			log.Printf("Error sending SIGTERM: %v", err)
		}

		select {
		case <-done:
			return nil
		case <-time.After(10 * time.Second):
			if err := cmd.Process.Kill(); err != nil {
				log.Printf("Error killing process: %v", err)
			}
			<-done
			return ctx.Err()
		}
	}
}
