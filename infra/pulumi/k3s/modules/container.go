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
		VmId:     pulumi.Int(args.Id),
		PoolId:   pulumi.String("CTs"),

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
			Type:           pulumi.String(args.OsType),
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
				Name:     pulumi.String("eth0"),
				Bridge:   pulumi.String("vmbr0"),
				Firewall: pulumi.Bool(false),
			},
		},

		Unprivileged: pulumi.Bool(args.Unprivileged),
		Features: &ct.ContainerFeaturesArgs{
			Nesting: pulumi.Bool(args.Nestable),
		},
		StartOnBoot: pulumi.Bool(true),
	}, pulumi.Provider(provider))
}

type CtTemplates string

const (
	CT_ALPINE_3 CtTemplates = "local:vztmpl/alpine-3.18-default_20230607_amd64.tar.xz"
)

type ContainerArgs struct {
	NodeName string
	Name     string
	Id       int
	Template CtTemplates
	OsType   string

	Cores   int
	Memory  int
	Storage int
	Network common.Network

	Password   string
	PublicKeys []string

	Nestable     bool
	Unprivileged bool
}
