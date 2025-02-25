package appsec

import (
	"encoding/json"
	"testing"

	"github.com/akamai/AkamaiOPEN-edgegrid-golang/v2/pkg/appsec"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestAkamaiApiEndpoints_data_basic(t *testing.T) {
	t.Run("match by ApiEndpoints ID", func(t *testing.T) {
		client := &mockappsec{}

		config := appsec.GetConfigurationResponse{}
		err := json.Unmarshal(loadFixtureBytes("testdata/TestResConfiguration/LatestConfiguration.json"), &config)
		require.NoError(t, err)

		client.On("GetConfiguration",
			mock.Anything,
			appsec.GetConfigurationRequest{ConfigID: 43253},
		).Return(&config, nil)

		getAPIEndpointsResponse := appsec.GetApiEndpointsResponse{}
		err = json.Unmarshal(loadFixtureBytes("testdata/TestDSApiEndpoints/ApiEndpoints.json"), &getAPIEndpointsResponse)
		require.NoError(t, err)

		client.On("GetApiEndpoints",
			mock.Anything,
			appsec.GetApiEndpointsRequest{ConfigID: 43253, Version: 7},
		).Return(&getAPIEndpointsResponse, nil)

		useClient(client, func() {
			resource.Test(t, resource.TestCase{
				IsUnitTest: true,
				Providers:  testAccProviders,
				Steps: []resource.TestStep{
					{
						Config: loadFixtureString("testdata/TestDSApiEndpoints/match_by_id.tf"),
						Check: resource.ComposeAggregateTestCheckFunc(
							resource.TestCheckResourceAttr("data.akamai_appsec_api_endpoints.test", "id", "619183"),
						),
					},
				},
			})
		})

		client.AssertExpectations(t)
	})

}
