// Copyright 2023 Istio Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package setupconfig

import (
	"testing"

	"istio.io/istio.io/pkg/test/istioio"
	"istio.io/istio/pkg/test/framework"
	"istio.io/istio/pkg/test/framework/components/istio"
	"istio.io/istio/pkg/test/framework/resource"
)

func TestMain(m *testing.M) {
	// nolint: staticcheck
	framework.
		NewSuite(m).
		Setup(istio.Setup(nil, setupConfig)).
		Run()
}

func TestDocs(t *testing.T) {
	framework.
		NewTest(t).
		Run(istioio.NewTestDocsFunc("profile=minimal"))
}

func setupConfig(ctx resource.Context, cfg *istio.Config) {
	// FIXME: test framework does not honor profile=minimal config at present,
	// hence we have to explicitly disable the gateways.
	cfg.ControlPlaneValues = `
values:
  pilot:
    env:
      PILOT_ENABLE_CONFIG_DISTRIBUTION_TRACKING: true
      PILOT_ENABLE_ALPHA_GATEWAY_API: false
components:
  egressGateways:
  - enabled: false
    name: istio-egressgateway
  ingressGateways:
  - enabled: false
    name: istio-ingressgateway
`
	cfg.DeployEastWestGW = false
}
