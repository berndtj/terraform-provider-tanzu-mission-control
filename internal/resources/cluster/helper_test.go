/*
Copyright © 2021 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package cluster

import (
	"bytes"
	"os"
	"text/template"

	testhelper "github.com/vmware-tanzu/terraform-provider-tanzu-mission-control/internal/resources/testing"
)

type acceptanceTestType int

const (
	attachClusterType acceptanceTestType = iota
	attachClusterTypeWithKubeConfig
)

type testAcceptanceOption func(config *testAcceptanceConfig)

type testAcceptanceConfig struct {
	ResourceName      string
	ResourceNameVar   string
	DataSourceNameVar string
	Name              string
	KubeConfigPath    string
	Meta              string
	accTestType       acceptanceTestType
	templateData      string
}

func withClusterName(name string) testAcceptanceOption {
	return func(config *testAcceptanceConfig) {
		config.Name = name
	}
}

func withKubeConfig() testAcceptanceOption {
	return func(config *testAcceptanceConfig) {
		config.KubeConfigPath = os.Getenv("KUBECONFIG")
		config.accTestType = attachClusterTypeWithKubeConfig
		config.templateData = testAttachClusterWithKubeConfigScript
	}
}

func withDataSourceScript() testAcceptanceOption {
	return func(config *testAcceptanceConfig) {
		config.templateData = testDataSourceAttachClusterScript
		config.DataSourceNameVar = clusterDataSourceVar
	}
}

func testGetDefaultAcceptanceConfig() *testAcceptanceConfig {
	return &testAcceptanceConfig{
		ResourceName:    clusterResource,
		ResourceNameVar: clusterResourceVar,
		Meta:            testhelper.MetaTemplate,
		accTestType:     attachClusterType,
		templateData:    testDefaultAttachClusterScript,
	}
}

func parse(m interface{}, objects string) (string, error) {
	var definitionBytes bytes.Buffer

	t := template.Must(template.New("script").Parse(objects))
	if err := t.Execute(&definitionBytes, m); err != nil {
		return "", err
	}

	return definitionBytes.String(), nil
}