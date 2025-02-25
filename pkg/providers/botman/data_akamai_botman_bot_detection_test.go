package botman

import (
	"testing"

	"github.com/akamai/AkamaiOPEN-edgegrid-golang/v2/pkg/botman"
	"github.com/akamai/terraform-provider-akamai/v2/pkg/test"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/stretchr/testify/mock"
)

func TestDataBotDetection(t *testing.T) {
	t.Run("DataBotDetection", func(t *testing.T) {
		mockedBotmanClient := &mockbotman{}

		response := botman.GetBotDetectionListResponse{
			Detections: []map[string]interface{}{
				{"detectionId": "b85e3eaa-d334-466d-857e-33308ce416be", "detectionName": "Test name 1", "testKey": "testValue1"},
				{"detectionId": "69acad64-7459-4c1d-9bad-672600150127", "detectionName": "Test name 2", "testKey": "testValue2"},
				{"detectionId": "cc9c3f89-e179-4892-89cf-d5e623ba9dc7", "detectionName": "Test name 3", "testKey": "testValue3"},
				{"detectionId": "10c54ea3-e3cb-4fc0-b0e0-fa3658aebd7b", "detectionName": "Test name 4", "testKey": "testValue4"},
				{"detectionId": "4d64d85a-a07f-485a-bbac-24c60658a1b8", "detectionName": "Test name 5", "testKey": "testValue5"},
			},
		}
		expectedJSON := `
{
	"detections":[
		{"detectionId":"b85e3eaa-d334-466d-857e-33308ce416be", "detectionName": "Test name 1", "testKey":"testValue1"},
		{"detectionId":"69acad64-7459-4c1d-9bad-672600150127", "detectionName": "Test name 2", "testKey":"testValue2"},
		{"detectionId":"cc9c3f89-e179-4892-89cf-d5e623ba9dc7", "detectionName": "Test name 3", "testKey":"testValue3"},
		{"detectionId":"10c54ea3-e3cb-4fc0-b0e0-fa3658aebd7b", "detectionName": "Test name 4", "testKey":"testValue4"},
		{"detectionId":"4d64d85a-a07f-485a-bbac-24c60658a1b8", "detectionName": "Test name 5", "testKey":"testValue5"}
	]
}`
		mockedBotmanClient.On("GetBotDetectionList",
			mock.Anything,
			botman.GetBotDetectionListRequest{},
		).Return(&response, nil)
		useClient(mockedBotmanClient, func() {
			resource.Test(t, resource.TestCase{
				IsUnitTest: true,
				Providers:  testAccProviders,
				Steps: []resource.TestStep{
					{
						Config: test.Fixture("testdata/TestDataBotDetection/basic.tf"),
						Check: resource.ComposeAggregateTestCheckFunc(
							resource.TestCheckResourceAttr("data.akamai_botman_bot_detection.test", "json", compactJSON(expectedJSON))),
					},
				},
			})
		})

		mockedBotmanClient.AssertExpectations(t)
	})
	t.Run("DataBotDetection filter by BotName", func(t *testing.T) {
		mockedBotmanClient := &mockbotman{}

		response := botman.GetBotDetectionListResponse{
			Detections: []map[string]interface{}{
				{"detectionId": "cc9c3f89-e179-4892-89cf-d5e623ba9dc7", "detectionName": "Test name 3", "testKey": "testValue3"},
			},
		}
		expectedJSON := `
{
	"detections":[
		{"detectionId":"cc9c3f89-e179-4892-89cf-d5e623ba9dc7", "detectionName": "Test name 3", "testKey":"testValue3"}
	]
}`
		mockedBotmanClient.On("GetBotDetectionList",
			mock.Anything,
			botman.GetBotDetectionListRequest{DetectionName: "Test name 3"},
		).Return(&response, nil)
		useClient(mockedBotmanClient, func() {
			resource.Test(t, resource.TestCase{
				IsUnitTest: true,
				Providers:  testAccProviders,
				Steps: []resource.TestStep{
					{
						Config: test.Fixture("testdata/TestDataBotDetection/filter_by_name.tf"),
						Check: resource.ComposeAggregateTestCheckFunc(
							resource.TestCheckResourceAttr("data.akamai_botman_bot_detection.test", "json", compactJSON(expectedJSON))),
					},
				},
			})
		})

		mockedBotmanClient.AssertExpectations(t)
	})
}
