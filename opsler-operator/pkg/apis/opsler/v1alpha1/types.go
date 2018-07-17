package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type VirtualEnvironmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []VirtualEnvironment `json:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type VirtualEnvironment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec              VirtualEnvironmentSpec   `json:"spec"`
	Status            VirtualEnvironmentStatus `json:"status,omitempty"`
}

type VirtualEnvironmentSpec struct {
	Http             []*HTTPRoute     `json:"http,omitempty"`
	DestinationRoute DestinationRoute `json:"destinationRoute,omitempty"`
	Services         []*Service       `json:"services,omitempty"`
}

type VirtualEnvironmentStatus struct {
	// Fill me
}

type PortSelector struct {
	Number uint32 `json:"number,omitempty"`
}

type DestinationRoute struct {
	Host string        `json:"host,omitempty"`
	Port *PortSelector `json:"port,omitempty"`
}

type Service struct {
	Host   string            `json:"host,omitempty"`
	Labels map[string]string `json:"labels,omitempty"`
}

type HTTPRoute struct {
	Match []*HTTPMatchRequest `json:"match,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type TargetingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []Targeting `json:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type Targeting struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec              TargetingSpec   `json:"spec"`
	Status            TargetingStatus `json:"status,omitempty"`
}

type TargetingSpec struct {
	Entrypoint         string  `json:"entrypoint,omitempty"`
	Priority           int32   `json:"priority,omitempty"`
	Segment            Segment `json:"segment,omitempty"`
	VirtualEnvironment string  `json:"virtualEnvironment,omitempty"`
}
type TargetingStatus struct {
	// Fill me
}

type Segment struct {
	HttpMatch []*HTTPMatchRequest `json:"httpMatch,omitempty"`
}

type HTTPMatchRequest struct {
	Uri          map[string]string            `json:"uri,omitempty"`
	Scheme       map[string]string            `json:"scheme,omitempty"`
	Method       map[string]string            `json:"method,omitempty"`
	Authority    map[string]string            `json:"authority,omitempty"`
	Headers      map[string]map[string]string `json:"headers,omitempty"`
	Port         uint32                       `json:"port,omitempty"`
	SourceLabels map[string]string            `json:"source_labels,omitempty"`
	Gateways     []string                     `json:"gateways,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type EntrypointList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []Entrypoint `json:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type Entrypoint struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec              EntrypointSpec   `json:"spec"`
	Status            EntrypointStatus `json:"status,omitempty"`
}

type EntrypointSpec struct {
	Servers                   []*Server `json:"servers,omitempty"`
	DefaultVirtualEnvironment string    `json:"defaultVirtualEnvironment,omitempty"`
}
type EntrypointStatus struct {
	// Fill me
}

type Server_TLSOptions struct {
	HttpsRedirect     bool     `json:"https_redirect,omitempty"`
	Mode              string   `json:"mode,omitempty"`
	ServerCertificate string   `json:"server_certificate,omitempty"`
	PrivateKey        string   `json:"private_key,omitempty"`
	CaCertificates    string   `json:"ca_certificates,omitempty"`
	SubjectAltNames   []string `json:"subject_alt_names,omitempty"`
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
