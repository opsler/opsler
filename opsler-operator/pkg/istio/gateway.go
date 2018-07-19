package istio

import (
	api "github.com/opsler/opsler/opsler-operator/pkg/apis/opsler/v1alpha1"
	istio "github.com/opsler/opsler/opsler-operator/pkg/istio/apis/istio/v1alpha3"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func GenerateIstioGateway(entrypoint api.Entrypoint, namespace string) (istio.Gateway, string) {
	istioServers := make([]*istio.Server, 0)

	for _, server := range entrypoint.Spec.Servers {
		var tls *istio.Server_TLSOptions
		if server.Tls != nil {
			tls = &istio.Server_TLSOptions{
				HttpsRedirect:     server.Tls.HttpsRedirect,
				Mode:              server.Tls.Mode,
				ServerCertificate: server.Tls.ServerCertificate,
				PrivateKey:        server.Tls.PrivateKey,
				CaCertificates:    server.Tls.CaCertificates,
				SubjectAltNames:   server.Tls.SubjectAltNames,
			}
		}

		istioServers = append(istioServers, &istio.Server{
			Port: &istio.Port{
				Name:     server.Port.Name,
				Number:   server.Port.Number,
				Protocol: server.Port.Protocol},
			Hosts: server.Hosts,
			Tls:   tls,
		})
	}
	gatewaySpec := istio.GatewaySpec{
		Selector: map[string]string{"istio": "ingressgateway"},
		Servers:  istioServers,
	}

	gatewayName := "opsler-" + entrypoint.ObjectMeta.Name

	gateway := createGateway(gatewayName, gatewaySpec, entrypoint)

	return *gateway, gatewayName
}

func createGateway(name string, gatewaySpec istio.GatewaySpec, entrypoint api.Entrypoint) *istio.Gateway {
	labels := map[string]string{}
	return &istio.Gateway{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Gateway",
			APIVersion: "networking.istio.io/v1alpha3",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: entrypoint.Namespace,
			OwnerReferences: []metav1.OwnerReference{
				*metav1.NewControllerRef(&entrypoint, schema.GroupVersionKind{
					Group:   api.SchemeGroupVersion.Group,
					Version: api.SchemeGroupVersion.Version,
					Kind:    "Entrypoint",
				}),
			},
			Labels: labels,
		},
		Spec: gatewaySpec,
	}
}
