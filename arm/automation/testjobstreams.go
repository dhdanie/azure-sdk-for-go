package automation

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
// Code generated by Microsoft (R) AutoRest Code Generator 2.2.22.0
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"net/http"
)

// TestJobStreamsClient is the automation Client
type TestJobStreamsClient struct {
	ManagementClient
}

// NewTestJobStreamsClient creates an instance of the TestJobStreamsClient client.
func NewTestJobStreamsClient(subscriptionID string) TestJobStreamsClient {
	return NewTestJobStreamsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewTestJobStreamsClientWithBaseURI creates an instance of the TestJobStreamsClient client.
func NewTestJobStreamsClientWithBaseURI(baseURI string, subscriptionID string) TestJobStreamsClient {
	return TestJobStreamsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Get retrieve a test job streams identified by runbook name and stream id.
//
// resourceGroupName is the resource group name. automationAccountName is the automation account name. runbookName is
// the runbook name. jobStreamID is the job stream id.
func (client TestJobStreamsClient) Get(resourceGroupName string, automationAccountName string, runbookName string, jobStreamID string) (result JobStream, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "automation.TestJobStreamsClient", "Get")
	}

	req, err := client.GetPreparer(resourceGroupName, automationAccountName, runbookName, jobStreamID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.TestJobStreamsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "automation.TestJobStreamsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.TestJobStreamsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client TestJobStreamsClient) GetPreparer(resourceGroupName string, automationAccountName string, runbookName string, jobStreamID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"automationAccountName": autorest.Encode("path", automationAccountName),
		"jobStreamId":           autorest.Encode("path", jobStreamID),
		"resourceGroupName":     autorest.Encode("path", resourceGroupName),
		"runbookName":           autorest.Encode("path", runbookName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-10-31"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/runbooks/{runbookName}/draft/testJob/streams/{jobStreamId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client TestJobStreamsClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client TestJobStreamsClient) GetResponder(resp *http.Response) (result JobStream, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByTestJob retrieve a list of test job streams identified by runbook name.
//
// resourceGroupName is the resource group name. automationAccountName is the automation account name. runbookName is
// the runbook name. filter is the filter to apply on the operation.
func (client TestJobStreamsClient) ListByTestJob(resourceGroupName string, automationAccountName string, runbookName string, filter string) (result JobStreamListResult, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "automation.TestJobStreamsClient", "ListByTestJob")
	}

	req, err := client.ListByTestJobPreparer(resourceGroupName, automationAccountName, runbookName, filter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.TestJobStreamsClient", "ListByTestJob", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByTestJobSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "automation.TestJobStreamsClient", "ListByTestJob", resp, "Failure sending request")
		return
	}

	result, err = client.ListByTestJobResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.TestJobStreamsClient", "ListByTestJob", resp, "Failure responding to request")
	}

	return
}

// ListByTestJobPreparer prepares the ListByTestJob request.
func (client TestJobStreamsClient) ListByTestJobPreparer(resourceGroupName string, automationAccountName string, runbookName string, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"automationAccountName": autorest.Encode("path", automationAccountName),
		"resourceGroupName":     autorest.Encode("path", resourceGroupName),
		"runbookName":           autorest.Encode("path", runbookName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-10-31"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/runbooks/{runbookName}/draft/testJob/streams", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListByTestJobSender sends the ListByTestJob request. The method will close the
// http.Response Body if it receives an error.
func (client TestJobStreamsClient) ListByTestJobSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListByTestJobResponder handles the response to the ListByTestJob request. The method always
// closes the http.Response Body.
func (client TestJobStreamsClient) ListByTestJobResponder(resp *http.Response) (result JobStreamListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByTestJobNextResults retrieves the next set of results, if any.
func (client TestJobStreamsClient) ListByTestJobNextResults(lastResults JobStreamListResult) (result JobStreamListResult, err error) {
	req, err := lastResults.JobStreamListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "automation.TestJobStreamsClient", "ListByTestJob", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListByTestJobSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "automation.TestJobStreamsClient", "ListByTestJob", resp, "Failure sending next results request")
	}

	result, err = client.ListByTestJobResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.TestJobStreamsClient", "ListByTestJob", resp, "Failure responding to next results request")
	}

	return
}

// ListByTestJobComplete gets all elements from the list without paging.
func (client TestJobStreamsClient) ListByTestJobComplete(resourceGroupName string, automationAccountName string, runbookName string, filter string, cancel <-chan struct{}) (<-chan JobStream, <-chan error) {
	resultChan := make(chan JobStream)
	errChan := make(chan error, 1)
	go func() {
		defer func() {
			close(resultChan)
			close(errChan)
		}()
		list, err := client.ListByTestJob(resourceGroupName, automationAccountName, runbookName, filter)
		if err != nil {
			errChan <- err
			return
		}
		if list.Value != nil {
			for _, item := range *list.Value {
				select {
				case <-cancel:
					return
				case resultChan <- item:
					// Intentionally left blank
				}
			}
		}
		for list.NextLink != nil {
			list, err = client.ListByTestJobNextResults(list)
			if err != nil {
				errChan <- err
				return
			}
			if list.Value != nil {
				for _, item := range *list.Value {
					select {
					case <-cancel:
						return
					case resultChan <- item:
						// Intentionally left blank
					}
				}
			}
		}
	}()
	return resultChan, errChan
}
