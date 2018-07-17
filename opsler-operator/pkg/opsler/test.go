package opsler

import "testing"

func TestX(t *testing.T) {
	listOptions := sdk.WithListOptions(&metav1.ListOptions{
		IncludeUninitialized: false,
	})
	namespace := "default"

	entrypointList := api.EntrypointList{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Entrypoint",
			APIVersion: "v1alpha1",
		},
	}
	if err := sdk.List(namespace, &entrypointList, listOptions); err != nil {
		logrus.Errorf("Query failed: %v", err)
		return err
	}

	virtualEnvironmentList := api.VirtualEnvironmentList{
		TypeMeta: metav1.TypeMeta{
			Kind:       "VirtualEnvironment",
			APIVersion: "v1alpha1",
		},
	}
	if err := sdk.List(namespace, &virtualEnvironmentList, listOptions); err != nil {
		logrus.Errorf("Query failed: %v", err)
		return err
	}

	targetingList := api.TargetingList{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Targeting",
			APIVersion: "v1alpha1",
		},
	}
	if err := sdk.List(namespace, &targetingList, listOptions); err != nil {
		logrus.Errorf("Query failed: %v", err)
		return err
	}
]}
