package main

import (
	"context"
	"runtime"

	sdk "github.com/operator-framework/operator-sdk/pkg/sdk"
	k8sutil "github.com/operator-framework/operator-sdk/pkg/util/k8sutil"
	sdkVersion "github.com/operator-framework/operator-sdk/version"
	stub "github.com/opsler/opsler/opsler-operator/pkg/stub"

	"github.com/sirupsen/logrus"
)

func printVersion() {
	logrus.Infof("Go Version: %s", runtime.Version())
	logrus.Infof("Go OS/Arch: %s/%s", runtime.GOOS, runtime.GOARCH)
	logrus.Infof("operator-sdk Version: %v", sdkVersion.Version)
}

func main() {
	printVersion()

	sdk.ExposeMetricsPort()

	resource := "opsler.com/v1alpha1"
	kindVirtualEnvironment := "VirtualEnvironment"
	kindTargeting := "Targeting"
	kindEntrypoint := "Entrypoint"
	namespace, err := k8sutil.GetWatchNamespace()
	if err != nil {
		logrus.Fatalf("Failed to get watch namespace: %v", err)
	}
	resyncPeriod := 5
	logrus.Infof("Watching %s, %s, %s, %d", resource, kindVirtualEnvironment, namespace, resyncPeriod)
	logrus.Infof("Watching %s, %s, %s, %d", resource, kindTargeting, namespace, resyncPeriod)
	logrus.Infof("Watching %s, %s, %s, %d", resource, kindVirtualEnvironment, namespace, resyncPeriod)
	sdk.Watch(resource, kindVirtualEnvironment, namespace, resyncPeriod)
	sdk.Watch(resource, kindTargeting, namespace, resyncPeriod)
	sdk.Watch(resource, kindEntrypoint, namespace, resyncPeriod)
	sdk.Handle(stub.NewHandler())
	sdk.Run(context.TODO())
}
