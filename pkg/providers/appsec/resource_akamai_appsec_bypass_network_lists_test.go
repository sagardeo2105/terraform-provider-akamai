package appsec

import (
	"encoding/json"
	"testing"

	"github.com/akamai/AkamaiOPEN-edgegrid-golang/v2/pkg/appsec"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestAkamaiBypassNetworkLists_res_basic(t *testing.T) {
	t.Run("match by BypassNetworkLists ID", func(t *testing.T) {
		client := &mockappsec{}

		config := appsec.GetConfigurationResponse{}
		err := json.Unmarshal(loadFixtureBytes("testdata/TestResConfiguration/LatestConfiguration.json"), &config)
		require.NoError(t, err)

		client.On("GetConfiguration",
			mock.Anything,
			appsec.GetConfigurationRequest{ConfigID: 43253},
		).Return(&config, nil)

		updateWAPBypassNetworkListsResponse := appsec.UpdateWAPBypassNetworkListsResponse{}
		err = json.Unmarshal(loadFixtureBytes("testdata/TestResBypassNetworkLists/BypassNetworkLists.json"), &updateWAPBypassNetworkListsResponse)
		require.NoError(t, err)

		getWAPBypassNetworkListsResponse := appsec.GetWAPBypassNetworkListsResponse{}
		err = json.Unmarshal(loadFixtureBytes("testdata/TestResBypassNetworkLists/BypassNetworkLists.json"), &getWAPBypassNetworkListsResponse)
		require.NoError(t, err)

		removeWAPBypassNetworkListsResponse := appsec.RemoveWAPBypassNetworkListsResponse{}
		err = json.Unmarshal(loadFixtureBytes("testdata/TestResBypassNetworkLists/RemoveNetworkLists.json"), &removeWAPBypassNetworkListsResponse)
		require.NoError(t, err)

		client.On("GetWAPBypassNetworkLists",
			mock.Anything,
			appsec.GetWAPBypassNetworkListsRequest{ConfigID: 43253, Version: 7, PolicyID: "AAAA_81230"},
		).Return(&getWAPBypassNetworkListsResponse, nil)

		client.On("UpdateWAPBypassNetworkLists",
			mock.Anything,
			appsec.UpdateWAPBypassNetworkListsRequest{ConfigID: 43253, Version: 7, PolicyID: "AAAA_81230", NetworkLists: []string{"1304427_AAXXBBLIST", "888518_ACDDCKERS"}},
		).Return(&updateWAPBypassNetworkListsResponse, nil)

		client.On("RemoveWAPBypassNetworkLists",
			mock.Anything,
			appsec.RemoveWAPBypassNetworkListsRequest{ConfigID: 43253, Version: 7, PolicyID: "AAAA_81230", NetworkLists: []string{}},
		).Return(&removeWAPBypassNetworkListsResponse, nil)

		useClient(client, func() {
			resource.Test(t, resource.TestCase{
				IsUnitTest: true,
				Providers:  testAccProviders,
				Steps: []resource.TestStep{
					{
						Config: loadFixtureString("testdata/TestResBypassNetworkLists/match_by_id.tf"),
						Check: resource.ComposeAggregateTestCheckFunc(
							resource.TestCheckResourceAttr("akamai_appsec_bypass_network_lists.test", "id", "43253"),
						),
					},
				},
			})
		})

		client.AssertExpectations(t)
	})

}
