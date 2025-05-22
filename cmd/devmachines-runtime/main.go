package main

import (
	"context"
	"fmt"
	"log"
	"os/signal"
	"syscall"

	"github.com/utkin-tech/devmachines/pkg/cloudinit"
	"github.com/utkin-tech/devmachines/pkg/config"
	"github.com/utkin-tech/devmachines/pkg/disk"
	"github.com/utkin-tech/devmachines/pkg/network"
	"github.com/utkin-tech/devmachines/pkg/qemu"
	"github.com/utkin-tech/devmachines/pkg/serial"
	"github.com/utkin-tech/devmachines/pkg/vnc"
)

const InterfaceName = "eth0"

var BasePorts = []string{
	"22:22/tcp",
}

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

	var args []string

	diskArgs, err := disk.SetupDisk(cfg.Storage())
	if err != nil {
		return fmt.Errorf("failed to setup disk: %v", err)
	}
	args = append(args, diskArgs...)

	cloudInitArgs, err := cloudinit.SetupCloudInit(net, cfg.User())
	if err != nil {
		return fmt.Errorf("failed to setup cloud-init: %v", err)
	}
	args = append(args, cloudInitArgs...)

	var networkArgs []string
	switch env.Network {
	case config.NetworkTypeBridge:
		networkArgs, err = network.SetupBridge(net)
	case config.NetworkTypeNat:
		ports := append(BasePorts, env.Ports...)
		networkArgs, err = network.SetupNAT(ports)
	}
	if err != nil {
		return fmt.Errorf("failed to setup network: %v", err)
	}
	args = append(args, networkArgs...)

	vncArgs := vnc.Setup(env.VNC)
	args = append(args, vncArgs...)

	serialArgs := serial.Setup()
	args = append(args, serialArgs...)

	if err := qemu.StartVM(ctx, cfg.VM(), nil, args); err != nil {
		return fmt.Errorf("failed to launch VM: %v", err)
	}

	return nil
}
