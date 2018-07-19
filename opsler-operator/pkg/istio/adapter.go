package istio

import (
	"encoding/json"

	"github.com/operator-framework/operator-sdk/pkg/sdk"
	"github.com/opsler/opsler/opsler-operator/pkg/models"
	"github.com/sirupsen/logrus"
)

func Apply(entrypointFlow []models.EntrypointFlow, namespace string) {
	for _, entrypointFlow := range entrypointFlow {
		istioGateway, _ := GenerateIstioGateway(entrypointFlow.Entrypoint, namespace)
		json, _ := json.MarshalIndent(istioGateway, "", "  ")
		logrus.Infof("Create istio gateway: %s", string(json))

		if err := sdk.Create(&istioGateway); err != nil {
			logrus.Errorf("Creating istio gateway failed: %v", err)
		}

	}
}
