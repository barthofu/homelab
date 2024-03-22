import { Provider } from '@pulumi/kubernetes'

export const provider = new Provider('k3s-provider', {
	kubeconfig: '.kube/config.yml',
	suppressDeprecationWarnings: true,
	namespace: 'default',
})