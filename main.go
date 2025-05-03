package main

import (
	"context"
	"fmt"
	"log"
	"os/signal"
	"syscall"

	"github.com/utkin-tech/devmachines/cloudinit"
	"github.com/utkin-tech/devmachines/config"
	"github.com/utkin-tech/devmachines/disk"
	"github.com/utkin-tech/devmachines/network"
)

const InterfaceName = "eth0"

func main() {
	if err := run(); err != nil {
		log.Fatalln(err)
	}
}

func run() error {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer stop()

	env, err := config.LoadEnvironment()
	if err != nil {
		return fmt.Errorf("failed to load environment: %v", err)
	}

	cfg := config.NewConfig(env)

	net, err := network.NewNetwork(InterfaceName)
	if err != nil {
		return fmt.Errorf("failed to get info about network: %v", err)
	}

	diskArgs, err := disk.SetupDisk(cfg.Storage())
	if err != nil {
		return fmt.Errorf("failed to setup disk: %v", err)
	}

	cloudInitArgs, err := cloudinit.SetupCloudInit(net, cfg.User())
	if err != nil {
		return fmt.Errorf("failed to setup cloud-init: %v", err)
	}

	// networkArgs, err := network.SetupBridge(net)
	networkArgs, err := network.SetupNAT(network.Hostfwd{
		Proto:     network.ProtoTcp,
		Hostport:  "2222",
		Guestport: "22",
	})
	if err != nil {
		return fmt.Errorf("failed to setup network bridge: %v", err)
	}

	serialArgs := []string{
		"-serial", "unix:/socks/serial.sock,server,nowait",
	}

	vncArgs := []string{
		"-vga", "std",
		"-vnc", "unix:/socks/vnc.sock",
	}

	var args []string
	args = append(args, diskArgs...)
	args = append(args, cloudInitArgs...)
	args = append(args, networkArgs...)
	args = append(args, serialArgs...)
	args = append(args, vncArgs...)

	if err := StartVM(ctx, cfg.VM(), nil, args); err != nil {
		return fmt.Errorf("failed to launch VM: %v", err)
	}

	return nil
}
