package main

import (
	"log"
	"os"
	"proxmox/common"
	"proxmox/modules"

	"github.com/joho/godotenv"
	"github.com/muhlba91/pulumi-proxmoxve/sdk/v5/go/proxmoxve"
	"github.com/muhlba91/pulumi-proxmoxve/sdk/v5/go/proxmoxve/vm"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	pulumi.Run(func(ctx *pulumi.Context) error {

		// Create a new ProxmoxVE provider
		provider, err := proxmoxve.NewProvider(ctx, "proxmox", &proxmoxve.ProviderArgs{
			Endpoint: pulumi.String(os.Getenv("PROXMOX_API_URL")),
			Username: pulumi.String(os.Getenv("PROXMOX_USERNAME")),
			Password: pulumi.String(os.Getenv("PROXMOX_PASSWORD")),
			Insecure: pulumi.Bool(true),
		})
		if err != nil {
			return err
		}

		// common variables
		user := modules.VmUser{
			Username:   "bartho",
			Password:   os.Getenv("VM_PASSWORD"),
			PublicKeys: []string{},
		}
		ipv4Gateway := "192.168.1.254"

		// full-stack vms
		modules.CreateVms(&[]modules.VmArgs{
			{
				NodeName: "homelab",
				Name:     "nas",
				Id:       110,
				Cores:    2,
				Storage:  16,
				Memory:   common.MinMax{Min: 2, Max: 4},
				User:     user,
				Template: modules.VM_DEBIAN_12,
				Network: common.Network{
					Ipv4: common.Ip{Adress: "192.168.1.110/24", Gateway: ipv4Gateway},
				},
				Usb: &vm.VirtualMachineUsbArray{
					&vm.VirtualMachineUsbArgs{Host: pulumi.String("152d:0567"), Usb3: pulumi.Bool(true)},
				},
			},
			{
				NodeName: "homelab",
				Name:     "lab",
				Id:       111,
				Cores:    6,
				Storage:  50,
				Memory:   common.MinMax{Min: 4, Max: 6},
				User:     user,
				Template: modules.VM_DEBIAN_12,
				Network: common.Network{
					Ipv4: common.Ip{Adress: "192.168.1.111/24", Gateway: ipv4Gateway},
				},
			},
			{
				NodeName: "homelab",
				Name:     "coder",
				Id:       112,
				Cores:    6,
				Memory:   common.MinMax{Min: 12, Max: 16},
				Storage:  150,
				User:     user,
				Template: modules.VM_DEBIAN_12,
				Network: common.Network{
					Ipv4: common.Ip{Adress: "192.168.1.112/24", Gateway: ipv4Gateway},
				},
			},
		}, ctx, provider)

		// services vms
		modules.CreateVms(&[]modules.VmArgs{
			{
				NodeName: "homelab",
				Name:     "master-1",
				Id:       200,
				Cores:    2,
				Memory:   common.MinMax{Min: 1, Max: 1},
				Storage:  16,
				User:     user,
				Template: modules.VM_DEBIAN_12,
				Network: common.Network{
					Ipv4: common.Ip{Adress: "192.168.1.200/24", Gateway: ipv4Gateway},
				},
			},
			{
				NodeName: "homelab",
				Name:     "master-2",
				Id:       201,
				Cores:    2,
				Memory:   common.MinMax{Min: 1, Max: 1},
				Storage:  16,
				User:     user,
				Template: modules.VM_DEBIAN_12,
				Network: common.Network{
					Ipv4: common.Ip{Adress: "192.168.1.201/24", Gateway: ipv4Gateway},
				},
			},
			{
				NodeName: "homelab",
				Name:     "master-3",
				Id:       202,
				Cores:    2,
				Memory:   common.MinMax{Min: 1, Max: 1},
				Storage:  16,
				User:     user,
				Template: modules.VM_DEBIAN_12,
				Network: common.Network{
					Ipv4: common.Ip{Adress: "192.168.1.202/24", Gateway: ipv4Gateway},
				},
			},
			{
				NodeName: "homelab",
				Name:     "services-1",
				Id:       210,
				Cores:    6,
				Memory:   common.MinMax{Min: 8, Max: 12},
				Storage:  32,
				User:     user,
				Template: modules.VM_DEBIAN_12,
				Network: common.Network{
					Ipv4: common.Ip{Adress: "192.168.1.210/24", Gateway: ipv4Gateway},
				},
			},
			{
				NodeName: "homelab",
				Name:     "services-2",
				Id:       211,
				Cores:    6,
				Memory:   common.MinMax{Min: 8, Max: 12},
				Storage:  32,
				User:     user,
				Template: modules.VM_DEBIAN_12,
				Network: common.Network{
					Ipv4: common.Ip{Adress: "192.168.1.211/24", Gateway: ipv4Gateway},
				},
			},
			{
				NodeName: "homelab",
				Name:     "services-3",
				Id:       212,
				Cores:    6,
				Memory:   common.MinMax{Min: 8, Max: 12},
				Storage:  32,
				User:     user,
				Template: modules.VM_DEBIAN_12,
				Network: common.Network{
					Ipv4: common.Ip{Adress: "192.168.1.212/24", Gateway: ipv4Gateway},
				},
			},
		},
			ctx,
			provider,
		)

		// lxc containers
		modules.CreateContainer(&modules.ContainerArgs{
			NodeName: "homelab",
			Name:     "tailscale",
			Id:       150,
			Template: modules.CT_ALPINE_3,
			OsType:   "alpine",
			Cores:    1,
			Memory:   1,
			Storage:  8,
			Network: common.Network{
				Ipv4: common.Ip{Adress: "192.168.1.150/24", Gateway: ipv4Gateway},
			},
			Password: os.Getenv("VM_PASSWORD"),
			PublicKeys: []string{
				os.Getenv("PUBLIC_KEY"),
			},
			Nestable:     false,
			Unprivileged: true,
		}, ctx, provider)

		return nil
	})
}
