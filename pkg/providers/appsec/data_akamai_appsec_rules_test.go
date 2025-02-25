package appsec

import (
	"encoding/json"
	"testing"

	"github.com/akamai/AkamaiOPEN-edgegrid-golang/v2/pkg/appsec"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestAkamaiRules_data_basic(t *testing.T) {
	t.Run("match by Rules ID", func(t *testing.T) {
		client := &mockappsec{}

		getRulesResponse := appsec.GetRulesResponse{}
		err := json.Unmarshal(loadFixtureBytes("testdata/TestDSRules/Rules.json"), &getRulesResponse)
		require.NoError(t, err)

		configs := appsec.GetConfigurationResponse{}
		err = json.Unmarshal(loadFixtureBytes("testdata/TestResConfiguration/LatestConfiguration.json"), &configs)
		require.NoError(t, err)

		client.On("GetRules",
			mock.Anything,
			appsec.GetRulesRequest{ConfigID: 43253, Version: 7, PolicyID: "AAAA_81230"},
		).Return(&getRulesResponse, nil)

		client.On("GetConfiguration",
			mock.Anything,
			appsec.GetConfigurationRequest{ConfigID: 43253},
		).Return(&configs, nil)

		getWAFModeResponse := appsec.GetWAFModeResponse{}
		err = json.Unmarshal(loadFixtureBytes("testdata/TestResWAFMode/WAFMode.json"), &getWAFModeResponse)
		require.NoError(t, err)

		client.On("GetWAFMode",
			mock.Anything,
			appsec.GetWAFModeRequest{
				ConfigID: 43253,
				Version:  7,
				PolicyID: "AAAA_81230",
			}).Return(&getWAFModeResponse, nil)

		useClient(client, func() {
			resource.Test(t, resource.TestCase{
				IsUnitTest: true,
				Providers:  testAccProviders,
				Steps: []resource.TestStep{
					{
						Config: loadFixtureString("testdata/TestDSRules/match_by_id.tf"),
						Check: resource.ComposeAggregateTestCheckFunc(
							resource.TestCheckResourceAttr("data.akamai_appsec_rules.test", "id", "43253"),
						),
					},
				},
			})
		})

		client.AssertExpectations(t)
	})

}
