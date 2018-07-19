package models

import (
	api "github.com/opsler/opsler/opsler-operator/pkg/apis/opsler/v1alpha1"
)

type TargetingFlow struct {
	Targeting          api.Targeting
	VirtualEnvironment api.VirtualEnvironment
}

type EntrypointFlow struct {
	Entrypoint                api.Entrypoint
	DefaultVirtualEnvironment api.VirtualEnvironment
	Targetings                []TargetingFlow
}
