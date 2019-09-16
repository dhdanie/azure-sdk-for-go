package security

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// IoTSecuritySolutionsAnalyticsRecommendationClient is the API spec for Microsoft.Security (Azure Security Center)
// resource provider
type IoTSecuritySolutionsAnalyticsRecommendationClient struct {
	BaseClient
}

// NewIoTSecuritySolutionsAnalyticsRecommendationClient creates an instance of the
// IoTSecuritySolutionsAnalyticsRecommendationClient client.
func NewIoTSecuritySolutionsAnalyticsRecommendationClient(subscriptionID string, ascLocation string) IoTSecuritySolutionsAnalyticsRecommendationClient {
	return NewIoTSecuritySolutionsAnalyticsRecommendationClientWithBaseURI(DefaultBaseURI, subscriptionID, ascLocation)
}

// NewIoTSecuritySolutionsAnalyticsRecommendationClientWithBaseURI creates an instance of the
// IoTSecuritySolutionsAnalyticsRecommendationClient client.
func NewIoTSecuritySolutionsAnalyticsRecommendationClientWithBaseURI(baseURI string, subscriptionID string, ascLocation string) IoTSecuritySolutionsAnalyticsRecommendationClient {
	return IoTSecuritySolutionsAnalyticsRecommendationClient{NewWithBaseURI(baseURI, subscriptionID, ascLocation)}
}

// Get security Analytics of a security solution
// Parameters:
// resourceGroupName - the name of the resource group within the user's subscription. The name is case
// insensitive.
// solutionName - the solution manager name
// aggregatedRecommendationName - identifier of the aggregated recommendation
func (client IoTSecuritySolutionsAnalyticsRecommendationClient) Get(ctx context.Context, resourceGroupName string, solutionName string, aggregatedRecommendationName string) (result IoTSecurityAggregatedRecommendation, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/IoTSecuritySolutionsAnalyticsRecommendationClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.Pattern, Rule: `^[0-9A-Fa-f]{8}-([0-9A-Fa-f]{4}-){3}[0-9A-Fa-f]{12}$`, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("security.IoTSecuritySolutionsAnalyticsRecommendationClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, solutionName, aggregatedRecommendationName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.IoTSecuritySolutionsAnalyticsRecommendationClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "security.IoTSecuritySolutionsAnalyticsRecommendationClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.IoTSecuritySolutionsAnalyticsRecommendationClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client IoTSecuritySolutionsAnalyticsRecommendationClient) GetPreparer(ctx context.Context, resourceGroupName string, solutionName string, aggregatedRecommendationName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"aggregatedRecommendationName": autorest.Encode("path", aggregatedRecommendationName),
		"resourceGroupName":            autorest.Encode("path", resourceGroupName),
		"solutionName":                 autorest.Encode("path", solutionName),
		"subscriptionId":               autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-08-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Security/iotSecuritySolutions/{solutionName}/analyticsModels/default/aggregatedRecommendations/{aggregatedRecommendationName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client IoTSecuritySolutionsAnalyticsRecommendationClient) GetSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client IoTSecuritySolutionsAnalyticsRecommendationClient) GetResponder(resp *http.Response) (result IoTSecurityAggregatedRecommendation, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}