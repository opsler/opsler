package v1alpha3

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type Gateway struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec              GatewaySpec `json:"spec"`
}

type GatewaySpec struct {
	Servers  []*Server         `json:"servers,omitempty"`
	Selector map[string]string `json:"selector,omitempty"`
}

type Port struct {
	Number   uint32 `json:"number,omitempty"`
	Protocol string `json:"protocol,omitempty"`
	Name     string `json:"name,omitempty"`
}

type Server struct {
	Port  *Port              `json:"port,omitempty"`
	Hosts []string           `json:"hosts,omitempty"`
	Tls   *Server_TLSOptions `json:"tls,omitempty"`
}

type Server_TLSOptions struct {
	HttpsRedirect     bool     `json:"https_redirect,omitempty"`
	Mode              string   `json:"mode,omitempty"`
	ServerCertificate string   `json:"server_certificate,omitempty"`
	PrivateKey        string   `json:"private_key,omitempty"`
	CaCertificates    string   `json:"ca_certificates,omitempty"`
	SubjectAltNames   []string `json:"subject_alt_names,omitempty"`
}
