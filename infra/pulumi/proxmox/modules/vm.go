package modules

import (
	"proxmox/common"

	"github.com/muhlba91/pulumi-proxmoxve/sdk/v5/go/proxmoxve"
	"github.com/muhlba91/pulumi-proxmoxve/sdk/v5/go/proxmoxve/vm"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func CreateVm(
	args *VmArgs,
	ctx *pulumi.Context,
	provider *proxmoxve.Provider,
) (*vm.VirtualMachine, error) {

	var mappedPublicKeys pulumi.StringArray
	for _, key := range args.User.PublicKeys {
		mappedPublicKeys = append(mappedPublicKeys, pulumi.String(key))
	}

	return vm.NewVirtualMachine(ctx, args.Name, &vm.VirtualMachineArgs{
		NodeName: pulumi.String(args.NodeName),
		Name:     pulumi.String(args.Name),
		VmId:     pulumi.Int(args.Id),

		Agent: &vm.VirtualMachineAgentArgs{
			Enabled: pulumi.Bool(true),
			Trim:    pulumi.Bool(true),
			Type:    pulumi.String("virtio"),
			Timeout: pulumi.String("20m"),
		},
		Bios:         pulumi.String("seabios"),
		ScsiHardware: pulumi.String("virtio-scsi-pci"),
		OperatingSystem: &vm.VirtualMachineOperatingSystemArgs{
			Type: pulumi.String("l26"),
		},

		Cpu: &vm.VirtualMachineCpuArgs{
			Cores:   pulumi.Int(args.Cores),
			Sockets: pulumi.Int(1),
		},
		Memory: &vm.VirtualMachineMemoryArgs{
			Floating:  pulumi.Int(args.Memory.Min * 1024),
			Dedicated: pulumi.Int(args.Memory.Max * 1024),
		},
		Disks: &vm.VirtualMachineDiskArray{
			&vm.VirtualMachineDiskArgs{
				Size:        pulumi.Int(args.Storage),
				FileFormat:  pulumi.String("raw"),
				Interface:   pulumi.String("virtio0"),
				DatastoreId: pulumi.String("local-lvm"),
				Ssd:         pulumi.Bool(true),
			},
		},

		NetworkDevices: &vm.VirtualMachineNetworkDeviceArray{
			&vm.VirtualMachineNetworkDeviceArgs{
				Model:    pulumi.String("virtio"),
				Bridge:   pulumi.String("vmbr0"),
				Firewall: pulumi.Bool(false),
			},
		},

		BootOrders: pulumi.StringArray{
			pulumi.String("virtio0"),
			pulumi.String("ide2"),
		},

		Initialization: &vm.VirtualMachineInitializationArgs{
			Type:        pulumi.String("nocloud"),
			DatastoreId: pulumi.String("local-lvm"),
			Interface:   pulumi.String("ide2"),
			IpConfigs: &vm.VirtualMachineInitializationIpConfigArray{
				&vm.VirtualMachineInitializationIpConfigArgs{
					Ipv4: &vm.VirtualMachineInitializationIpConfigIpv4Args{
						Address: pulumi.String(args.Network.Ipv4.Adress),
						Gateway: pulumi.String(args.Network.Ipv4.Gateway),
					},
				},
			},
			UserAccount: &vm.VirtualMachineInitializationUserAccountArgs{
				Username: pulumi.String(args.User.Username),
				Password: pulumi.String(args.User.Password),
				Keys:     pulumi.StringArray(mappedPublicKeys),
			},
		},

		KeyboardLayout: pulumi.String("fr"),
		Usbs:           args.Usb,

		Clone: &vm.VirtualMachineCloneArgs{
			Full:     pulumi.Bool(true),
			NodeName: pulumi.String(args.NodeName),
			VmId:     pulumi.Int(int(args.Template)),
		},

		OnBoot:  pulumi.Bool(true),
		Started: pulumi.Bool(true),
	}, pulumi.Provider(provider))
}

type Templates int

const (
	DEBIAN_12 Templates = 1000
)

type VmUser struct {
	Username   string
	Password   string
	PublicKeys []string
}

type VmArgs struct {
	NodeName string
	Name     string
	Id       int
	Cores    int
	Memory   common.MinMax
	Storage  int
	Template Templates
	Network  common.Network
	User     VmUser
	Usb      *vm.VirtualMachineUsbArray
}
