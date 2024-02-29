package modules

import (
	"proxmox/common"

	"github.com/muhlba91/pulumi-proxmoxve/sdk/v5/go/proxmoxve"
	"github.com/muhlba91/pulumi-proxmoxve/sdk/v5/go/proxmoxve/ct"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func CreateContainer(
	args *ContainerArgs,
	ctx *pulumi.Context,
	provider *proxmoxve.Provider,
) (*ct.Container, error) {

	var mappedPublicKeys pulumi.StringArray
	for _, key := range args.PublicKeys {
		mappedPublicKeys = append(mappedPublicKeys, pulumi.String(key))
	}

	return ct.NewContainer(ctx, args.Name, &ct.ContainerArgs{
		NodeName: pulumi.String(args.NodeName),

		Initialization: &ct.ContainerInitializationArgs{
			Hostname: pulumi.String(args.Name),
			UserAccount: &ct.ContainerInitializationUserAccountArgs{
				Password: pulumi.String(args.Password),
				Keys:     mappedPublicKeys,
			},
			IpConfigs: &ct.ContainerInitializationIpConfigArray{
				&ct.ContainerInitializationIpConfigArgs{
					Ipv4: &ct.ContainerInitializationIpConfigIpv4Args{
						Address: pulumi.String(args.Network.Ipv4.Adress),
						Gateway: pulumi.String(args.Network.Ipv4.Gateway),
					},
				},
			},
		},

		OperatingSystem: &ct.ContainerOperatingSystemArgs{
			TemplateFileId: pulumi.String(args.Template),
			Type:           pulumi.String("l26"),
		},

		Cpu: &ct.ContainerCpuArgs{
			Cores: pulumi.Int(args.Cores),
		},
		Memory: &ct.ContainerMemoryArgs{
			Dedicated: pulumi.Int(args.Memory * 1024),
			Swap:      pulumi.Int(2048),
		},
		Disk: &ct.ContainerDiskArgs{
			DatastoreId: pulumi.String("local-lvm"),
			Size:        pulumi.Int(args.Storage),
		},

		NetworkInterfaces: &ct.ContainerNetworkInterfaceArray{
			ct.ContainerNetworkInterfaceArgs{
				Bridge:   pulumi.String("vmbr0"),
				Firewall: pulumi.Bool(false),
			},
		},

		Unprivileged: pulumi.Bool(args.Unprivileged),
		Features: &ct.ContainerFeaturesArgs{
			Nesting: pulumi.Bool(args.Nestable),
		},
		StartOnBoot: pulumi.Bool(true),
	})
}

type ContainerArgs struct {
	NodeName string
	Name     string
	Template string

	Cores   int
	Memory  int
	Storage int
	Network common.Network

	Password   string
	PublicKeys []string

	Nestable     bool
	Unprivileged bool
}
