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

		modules.CreateVm(&modules.VmArgs{
			NodeName: "homelab",
			Name:     "lab",
			Id:       110,
			Cores:    6,
			Memory: common.MinMax{
				Min: 4,
				Max: 6,
			},
			Storage: 50,
			User: modules.VmUser{
				Username:   "bartho",
				Password:   os.Getenv("VM_PASSWORD"),
				PublicKeys: []string{},
			},
			Template: modules.DEBIAN_12,
			Network: common.Network{
				Ipv4: common.Ip{
					Adress:  "192.168.1.110/24",
					Gateway: "192.168.1.254",
				},
			},
		},
			ctx,
			provider,
		)

		modules.CreateVm(&modules.VmArgs{
			NodeName: "homelab",
			Name:     "nas",
			Id:       120,
			Cores:    2,
			Memory: common.MinMax{
				Min: 2,
				Max: 4,
			},
			Storage: 16,
			User: modules.VmUser{
				Username:   "bartho",
				Password:   os.Getenv("VM_PASSWORD"),
				PublicKeys: []string{},
			},
			Template: modules.DEBIAN_12,
			Network: common.Network{
				Ipv4: common.Ip{
					Adress:  "192.168.1.120/24",
					Gateway: "192.168.1.254",
				},
			},
			Usb: &vm.VirtualMachineUsbArray{
				&vm.VirtualMachineUsbArgs{
					Host: pulumi.String("152d:0567"),
					Usb3: pulumi.Bool(true),
				},
			},
		},
			ctx,
			provider,
		)

		modules.CreateVm(&modules.VmArgs{
			NodeName: "homelab",
			Name:     "coder",
			Id:       130,
			Cores:    6,
			Memory: common.MinMax{
				Min: 12,
				Max: 16,
			},
			Storage: 150,
			User: modules.VmUser{
				Username:   "bartho",
				Password:   os.Getenv("VM_PASSWORD"),
				PublicKeys: []string{},
			},
			Template: modules.DEBIAN_12,
			Network: common.Network{
				Ipv4: common.Ip{
					Adress:  "192.168.1.130/24",
					Gateway: "192.168.1.254",
				},
			},
		},
			ctx,
			provider,
		)

		modules.CreateVm(&modules.VmArgs{
			NodeName: "homelab",
			Name:     "services",
			Id:       200,
			Cores:    6,
			Memory: common.MinMax{
				Min: 14,
				Max: 16,
			},
			Storage: 100,
			User: modules.VmUser{
				Username:   "bartho",
				Password:   os.Getenv("VM_PASSWORD"),
				PublicKeys: []string{},
			},
			Template: modules.DEBIAN_12,
			Network: common.Network{
				Ipv4: common.Ip{
					Adress:  "192.168.1.150/24",
					Gateway: "192.168.1.254",
				},
			},
		},
			ctx,
			provider,
		)

		return nil
	})
}
