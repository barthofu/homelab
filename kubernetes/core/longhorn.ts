import { Provider, core, helm, networking } from '@pulumi/kubernetes'
import * as command from '@pulumi/command'

const connection: command.types.output.remote.Connection = {
	host: '192.168.1.200',
	user: 'bartho',
	privateKey: '',
}

// source for values: https://github.com/longhorn/longhorn/blob/v1.6.0/chart/values.yaml

export const longhorn = (provider: Provider) => {

	const chart = new helm.v3.Release('longhorn', {
		chart: 'longhorn',
		repositoryOpts: {
			repo: 'https://charts.longhorn.io',
		},
		version: '1.6.0',
		namespace: 'longhorn-system',
		createNamespace: true,
		values: {
			logLevel: 'debug',
			persistence: {
				defaultClassReplicaCount: 2,
			},
		},
	}, { provider })

	// ingress controller

	const ingress = new networking.v1.Ingress('longhorn-ingress', {
		metadata: {
			name: 'longhorn-ingress',
			namespace: 'longhorn-system',
			annotations: {
				'kubernetes.io/ingress.class': 'traefik',
				// 'traefik.ingress.kubernetes.io/router.entrypoints': 'web',
				// 'traefik.ingress.kubernetes.io/router.middlewares': 'default-headers',
				// 'traefik.ingress.kubernetes.io/router.service': 'longhorn-frontend',
				// 'traefik.ingress.kubernetes.io/router.tls': 'true',
				// 'traefik.ingress.kubernetes.io/router.tls.certresolver': 'letsencrypt',

				// type of authentication
				// 'nginx.ingress.kubernetes.io/auth-type': 'basic',
				// prevent the controller from redirecting (308) to HTTPS
				// 'nginx.ingress.kubernetes.io/ssl-redirect': 'false',
				// name of the secret that contains the user/password definitions
				// 'nginx.ingress.kubernetes.io/auth-secret': 'basic-auth',
				// message to display with an appropriate context why the authentication is required
				// 'nginx.ingress.kubernetes.io/auth-realm': 'Authentication Required ',
				// custom max body size for file uploading like backing image uploading'
				// 'nginx.ingress.kubernetes.io/proxy-body-size': '10000m'
			},
		},
		spec: {
			rules: [{
				host: 'longhorn.kube.bartho.dev',
				http: {
					paths: [
						{
							path: '/',
							pathType: 'Prefix',
							backend: {
								service: {
									name: 'longhorn-frontend',
									port: {
										number: 80,
									},
								},
							},
						}
					]
				},
			}],
			// tls: [{
			// 	hosts: ['longhorn.bartho.dev'],
			// 	secretName: 'longhorn-ingress-tls',
			// }],
		},
	}, { provider, dependsOn: chart })

	
}