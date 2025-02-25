package appsec

import (
	"encoding/json"
	"fmt"
	"regexp"
	"testing"

	"github.com/akamai/AkamaiOPEN-edgegrid-golang/v2/pkg/appsec"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestAkamaiEvalGroups_data_basic(t *testing.T) {
	t.Run("match by Eval Attack Group ID", func(t *testing.T) {
		client := &mockappsec{}

		configs := appsec.GetConfigurationResponse{}
		err := json.Unmarshal(loadFixtureBytes("testdata/TestResConfiguration/LatestConfiguration.json"), &configs)
		require.NoError(t, err)

		client.On("GetEvalGroups",
			mock.Anything,
			appsec.GetAttackGroupsRequest{ConfigID: 43253, Version: 7, PolicyID: "AAAA_81230", Group: "SQL"},
		).Return(nil, fmt.Errorf("GetEvalGroups failed"))

		client.On("GetConfiguration",
			mock.Anything,
			appsec.GetConfigurationRequest{ConfigID: 43253},
		).Return(&configs, nil)

		useClient(client, func() {
			resource.Test(t, resource.TestCase{
				IsUnitTest: true,
				Providers:  testAccProviders,
				Steps: []resource.TestStep{
					{
						Config: loadFixtureString("testdata/TestDSEvalGroups/match_by_id.tf"),
						Check: resource.ComposeAggregateTestCheckFunc(
							resource.TestCheckResourceAttr("data.akamai_appsec_eval_groups.test", "id", "43253"),
						),
						ExpectError: regexp.MustCompile(`GetEvalGroups failed`),
					},
				},
			})
		})

		client.AssertExpectations(t)
	})
}

func TestAkamaiEvalGroups_data_error_retrieving_eval_groups(t *testing.T) {
	t.Run("match by Eval Attack Group ID", func(t *testing.T) {
		client := &mockappsec{}

		configs := appsec.GetConfigurationResponse{}
		err := json.Unmarshal(loadFixtureBytes("testdata/TestResConfiguration/LatestConfiguration.json"), &configs)
		require.NoError(t, err)

		client.On("GetEvalGroups",
			mock.Anything,
			appsec.GetAttackGroupsRequest{ConfigID: 43253, Version: 7, PolicyID: "AAAA_81230", Group: "SQL"},
		).Return(nil, fmt.Errorf("GetEvalGroups failed"))

		client.On("GetConfiguration",
			mock.Anything,
			appsec.GetConfigurationRequest{ConfigID: 43253},
		).Return(&configs, nil)

		useClient(client, func() {
			resource.Test(t, resource.TestCase{
				IsUnitTest: true,
				Providers:  testAccProviders,
				Steps: []resource.TestStep{
					{
						Config: loadFixtureString("testdata/TestDSEvalGroups/match_by_id.tf"),
						Check: resource.ComposeAggregateTestCheckFunc(
							resource.TestCheckResourceAttr("data.akamai_appsec_eval_groups.test", "id", "43253"),
						),
						ExpectError: regexp.MustCompile(`GetEvalGroups failed`),
					},
				},
			})
		})

		client.AssertExpectations(t)
	})
}
