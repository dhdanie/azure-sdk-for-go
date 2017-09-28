package compute

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
	"net/http"
)

// VirtualMachineScaleSetVMsClient is the compute Client
type VirtualMachineScaleSetVMsClient struct {
	ManagementClient
}

// NewVirtualMachineScaleSetVMsClient creates an instance of the VirtualMachineScaleSetVMsClient client.
func NewVirtualMachineScaleSetVMsClient(subscriptionID string) VirtualMachineScaleSetVMsClient {
	return NewVirtualMachineScaleSetVMsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewVirtualMachineScaleSetVMsClientWithBaseURI creates an instance of the VirtualMachineScaleSetVMsClient client.
func NewVirtualMachineScaleSetVMsClientWithBaseURI(baseURI string, subscriptionID string) VirtualMachineScaleSetVMsClient {
	return VirtualMachineScaleSetVMsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Deallocate deallocates a specific virtual machine in a VM scale set. Shuts down the virtual machine and releases the
// compute resources it uses. You are not billed for the compute resources of this virtual machine once it is
// deallocated. This method may poll for completion. Polling can be canceled by passing the cancel channel argument.
// The channel will be used to cancel polling and any outstanding HTTP requests.
//
// resourceGroupName is the name of the resource group. VMScaleSetName is the name of the VM scale set. instanceID is
// the instance ID of the virtual machine.
func (client VirtualMachineScaleSetVMsClient) Deallocate(resourceGroupName string, VMScaleSetName string, instanceID string, cancel <-chan struct{}) (<-chan OperationStatusResponse, <-chan error) {
	resultChan := make(chan OperationStatusResponse, 1)
	errChan := make(chan error, 1)
	go func() {
		var err error
		var result OperationStatusResponse
		defer func() {
			if err != nil {
				errChan <- err
			}
			resultChan <- result
			close(resultChan)
			close(errChan)
		}()
		req, err := client.DeallocatePreparer(resourceGroupName, VMScaleSetName, instanceID, cancel)
		if err != nil {
			err = autorest.NewErrorWithError(err, "compute.VirtualMachineScaleSetVMsClient", "Deallocate", nil, "Failure preparing request")
			return
		}

		resp, err := client.DeallocateSender(req)
		if err != nil {
			result.Response = autorest.Response{Response: resp}
			err = autorest.NewErrorWithError(err, "compute.VirtualMachineScaleSetVMsClient", "Deallocate", resp, "Failure sending request")
			return
		}

		result, err = client.DeallocateResponder(resp)
		if err != nil {
			err = autorest.NewErrorWithError(err, "compute.VirtualMachineScaleSetVMsClient", "Deallocate", resp, "Failure responding to request")
		}
	}()
	return resultChan, errChan
}

// DeallocatePreparer prepares the Deallocate request.
func (client VirtualMachineScaleSetVMsClient) DeallocatePreparer(resourceGroupName string, VMScaleSetName string, instanceID string, cancel <-chan struct{}) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"instanceId":        autorest.Encode("path", instanceID),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"vmScaleSetName":    autorest.Encode("path", VMScaleSetName),
	}

	const APIVersion = "2017-03-30"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}/virtualmachines/{instanceId}/deallocate", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{Cancel: cancel})
}

// DeallocateSender sends the Deallocate request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualMachineScaleSetVMsClient) DeallocateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoPollForAsynchronous(client.PollingDelay))
}

// DeallocateResponder handles the response to the Deallocate request. The method always
// closes the http.Response Body.
func (client VirtualMachineScaleSetVMsClient) DeallocateResponder(resp *http.Response) (result OperationStatusResponse, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes a virtual machine from a VM scale set. This method may poll for completion. Polling can be canceled
// by passing the cancel channel argument. The channel will be used to cancel polling and any outstanding HTTP
// requests.
//
// resourceGroupName is the name of the resource group. VMScaleSetName is the name of the VM scale set. instanceID is
// the instance ID of the virtual machine.
func (client VirtualMachineScaleSetVMsClient) Delete(resourceGroupName string, VMScaleSetName string, instanceID string, cancel <-chan struct{}) (<-chan OperationStatusResponse, <-chan error) {
	resultChan := make(chan OperationStatusResponse, 1)
	errChan := make(chan error, 1)
	go func() {
		var err error
		var result OperationStatusResponse
		defer func() {
			if err != nil {
				errChan <- err
			}
			resultChan <- result
			close(resultChan)
			close(errChan)
		}()
		req, err := client.DeletePreparer(resourceGroupName, VMScaleSetName, instanceID, cancel)
		if err != nil {
			err = autorest.NewErrorWithError(err, "compute.VirtualMachineScaleSetVMsClient", "Delete", nil, "Failure preparing request")
			return
		}

		resp, err := client.DeleteSender(req)
		if err != nil {
			result.Response = autorest.Response{Response: resp}
			err = autorest.NewErrorWithError(err, "compute.VirtualMachineScaleSetVMsClient", "Delete", resp, "Failure sending request")
			return
		}

		result, err = client.DeleteResponder(resp)
		if err != nil {
			err = autorest.NewErrorWithError(err, "compute.VirtualMachineScaleSetVMsClient", "Delete", resp, "Failure responding to request")
		}
	}()
	return resultChan, errChan
}

// DeletePreparer prepares the Delete request.
func (client VirtualMachineScaleSetVMsClient) DeletePreparer(resourceGroupName string, VMScaleSetName string, instanceID string, cancel <-chan struct{}) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"instanceId":        autorest.Encode("path", instanceID),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"vmScaleSetName":    autorest.Encode("path", VMScaleSetName),
	}

	const APIVersion = "2017-03-30"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}/virtualmachines/{instanceId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{Cancel: cancel})
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualMachineScaleSetVMsClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoPollForAsynchronous(client.PollingDelay))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client VirtualMachineScaleSetVMsClient) DeleteResponder(resp *http.Response) (result OperationStatusResponse, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Get gets a virtual machine from a VM scale set.
//
// resourceGroupName is the name of the resource group. VMScaleSetName is the name of the VM scale set. instanceID is
// the instance ID of the virtual machine.
func (client VirtualMachineScaleSetVMsClient) Get(resourceGroupName string, VMScaleSetName string, instanceID string) (result VirtualMachineScaleSetVM, err error) {
	req, err := client.GetPreparer(resourceGroupName, VMScaleSetName, instanceID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.VirtualMachineScaleSetVMsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "compute.VirtualMachineScaleSetVMsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.VirtualMachineScaleSetVMsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client VirtualMachineScaleSetVMsClient) GetPreparer(resourceGroupName string, VMScaleSetName string, instanceID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"instanceId":        autorest.Encode("path", instanceID),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"vmScaleSetName":    autorest.Encode("path", VMScaleSetName),
	}

	const APIVersion = "2017-03-30"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}/virtualmachines/{instanceId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualMachineScaleSetVMsClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client VirtualMachineScaleSetVMsClient) GetResponder(resp *http.Response) (result VirtualMachineScaleSetVM, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetInstanceView gets the status of a virtual machine from a VM scale set.
//
// resourceGroupName is the name of the resource group. VMScaleSetName is the name of the VM scale set. instanceID is
// the instance ID of the virtual machine.
func (client VirtualMachineScaleSetVMsClient) GetInstanceView(resourceGroupName string, VMScaleSetName string, instanceID string) (result VirtualMachineScaleSetVMInstanceView, err error) {
	req, err := client.GetInstanceViewPreparer(resourceGroupName, VMScaleSetName, instanceID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.VirtualMachineScaleSetVMsClient", "GetInstanceView", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetInstanceViewSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "compute.VirtualMachineScaleSetVMsClient", "GetInstanceView", resp, "Failure sending request")
		return
	}

	result, err = client.GetInstanceViewResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.VirtualMachineScaleSetVMsClient", "GetInstanceView", resp, "Failure responding to request")
	}

	return
}

// GetInstanceViewPreparer prepares the GetInstanceView request.
func (client VirtualMachineScaleSetVMsClient) GetInstanceViewPreparer(resourceGroupName string, VMScaleSetName string, instanceID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"instanceId":        autorest.Encode("path", instanceID),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"vmScaleSetName":    autorest.Encode("path", VMScaleSetName),
	}

	const APIVersion = "2017-03-30"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}/virtualmachines/{instanceId}/instanceView", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetInstanceViewSender sends the GetInstanceView request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualMachineScaleSetVMsClient) GetInstanceViewSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetInstanceViewResponder handles the response to the GetInstanceView request. The method always
// closes the http.Response Body.
func (client VirtualMachineScaleSetVMsClient) GetInstanceViewResponder(resp *http.Response) (result VirtualMachineScaleSetVMInstanceView, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List gets a list of all virtual machines in a VM scale sets.
//
// resourceGroupName is the name of the resource group. virtualMachineScaleSetName is the name of the VM scale set.
// filter is the filter to apply to the operation. selectParameter is the list parameters. expand is the expand
// expression to apply to the operation.
func (client VirtualMachineScaleSetVMsClient) List(resourceGroupName string, virtualMachineScaleSetName string, filter string, selectParameter string, expand string) (result VirtualMachineScaleSetVMListResult, err error) {
	req, err := client.ListPreparer(resourceGroupName, virtualMachineScaleSetName, filter, selectParameter, expand)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.VirtualMachineScaleSetVMsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "compute.VirtualMachineScaleSetVMsClient", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.VirtualMachineScaleSetVMsClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client VirtualMachineScaleSetVMsClient) ListPreparer(resourceGroupName string, virtualMachineScaleSetName string, filter string, selectParameter string, expand string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":          autorest.Encode("path", resourceGroupName),
		"subscriptionId":             autorest.Encode("path", client.SubscriptionID),
		"virtualMachineScaleSetName": autorest.Encode("path", virtualMachineScaleSetName),
	}

	const APIVersion = "2017-03-30"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}
	if len(selectParameter) > 0 {
		queryParameters["$select"] = autorest.Encode("query", selectParameter)
	}
	if len(expand) > 0 {
		queryParameters["$expand"] = autorest.Encode("query", expand)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{virtualMachineScaleSetName}/virtualMachines", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualMachineScaleSetVMsClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client VirtualMachineScaleSetVMsClient) ListResponder(resp *http.Response) (result VirtualMachineScaleSetVMListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListNextResults retrieves the next set of results, if any.
func (client VirtualMachineScaleSetVMsClient) ListNextResults(lastResults VirtualMachineScaleSetVMListResult) (result VirtualMachineScaleSetVMListResult, err error) {
	req, err := lastResults.VirtualMachineScaleSetVMListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "compute.VirtualMachineScaleSetVMsClient", "List", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "compute.VirtualMachineScaleSetVMsClient", "List", resp, "Failure sending next results request")
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.VirtualMachineScaleSetVMsClient", "List", resp, "Failure responding to next results request")
	}

	return
}

// ListComplete gets all elements from the list without paging.
func (client VirtualMachineScaleSetVMsClient) ListComplete(resourceGroupName string, virtualMachineScaleSetName string, filter string, selectParameter string, expand string, cancel <-chan struct{}) (<-chan VirtualMachineScaleSetVM, <-chan error) {
	resultChan := make(chan VirtualMachineScaleSetVM)
	errChan := make(chan error, 1)
	go func() {
		defer func() {
			close(resultChan)
			close(errChan)
		}()
		list, err := client.List(resourceGroupName, virtualMachineScaleSetName, filter, selectParameter, expand)
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
			list, err = client.ListNextResults(list)
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

// PowerOff power off (stop) a virtual machine in a VM scale set. Note that resources are still attached and you are
// getting charged for the resources. Instead, use deallocate to release resources and avoid charges. This method may
// poll for completion. Polling can be canceled by passing the cancel channel argument. The channel will be used to
// cancel polling and any outstanding HTTP requests.
//
// resourceGroupName is the name of the resource group. VMScaleSetName is the name of the VM scale set. instanceID is
// the instance ID of the virtual machine.
func (client VirtualMachineScaleSetVMsClient) PowerOff(resourceGroupName string, VMScaleSetName string, instanceID string, cancel <-chan struct{}) (<-chan OperationStatusResponse, <-chan error) {
	resultChan := make(chan OperationStatusResponse, 1)
	errChan := make(chan error, 1)
	go func() {
		var err error
		var result OperationStatusResponse
		defer func() {
			if err != nil {
				errChan <- err
			}
			resultChan <- result
			close(resultChan)
			close(errChan)
		}()
		req, err := client.PowerOffPreparer(resourceGroupName, VMScaleSetName, instanceID, cancel)
		if err != nil {
			err = autorest.NewErrorWithError(err, "compute.VirtualMachineScaleSetVMsClient", "PowerOff", nil, "Failure preparing request")
			return
		}

		resp, err := client.PowerOffSender(req)
		if err != nil {
			result.Response = autorest.Response{Response: resp}
			err = autorest.NewErrorWithError(err, "compute.VirtualMachineScaleSetVMsClient", "PowerOff", resp, "Failure sending request")
			return
		}

		result, err = client.PowerOffResponder(resp)
		if err != nil {
			err = autorest.NewErrorWithError(err, "compute.VirtualMachineScaleSetVMsClient", "PowerOff", resp, "Failure responding to request")
		}
	}()
	return resultChan, errChan
}

// PowerOffPreparer prepares the PowerOff request.
func (client VirtualMachineScaleSetVMsClient) PowerOffPreparer(resourceGroupName string, VMScaleSetName string, instanceID string, cancel <-chan struct{}) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"instanceId":        autorest.Encode("path", instanceID),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"vmScaleSetName":    autorest.Encode("path", VMScaleSetName),
	}

	const APIVersion = "2017-03-30"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}/virtualmachines/{instanceId}/poweroff", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{Cancel: cancel})
}

// PowerOffSender sends the PowerOff request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualMachineScaleSetVMsClient) PowerOffSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoPollForAsynchronous(client.PollingDelay))
}

// PowerOffResponder handles the response to the PowerOff request. The method always
// closes the http.Response Body.
func (client VirtualMachineScaleSetVMsClient) PowerOffResponder(resp *http.Response) (result OperationStatusResponse, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Reimage reimages (upgrade the operating system) a specific virtual machine in a VM scale set. This method may poll
// for completion. Polling can be canceled by passing the cancel channel argument. The channel will be used to cancel
// polling and any outstanding HTTP requests.
//
// resourceGroupName is the name of the resource group. VMScaleSetName is the name of the VM scale set. instanceID is
// the instance ID of the virtual machine.
func (client VirtualMachineScaleSetVMsClient) Reimage(resourceGroupName string, VMScaleSetName string, instanceID string, cancel <-chan struct{}) (<-chan OperationStatusResponse, <-chan error) {
	resultChan := make(chan OperationStatusResponse, 1)
	errChan := make(chan error, 1)
	go func() {
		var err error
		var result OperationStatusResponse
		defer func() {
			if err != nil {
				errChan <- err
			}
			resultChan <- result
			close(resultChan)
			close(errChan)
		}()
		req, err := client.ReimagePreparer(resourceGroupName, VMScaleSetName, instanceID, cancel)
		if err != nil {
			err = autorest.NewErrorWithError(err, "compute.VirtualMachineScaleSetVMsClient", "Reimage", nil, "Failure preparing request")
			return
		}

		resp, err := client.ReimageSender(req)
		if err != nil {
			result.Response = autorest.Response{Response: resp}
			err = autorest.NewErrorWithError(err, "compute.VirtualMachineScaleSetVMsClient", "Reimage", resp, "Failure sending request")
			return
		}

		result, err = client.ReimageResponder(resp)
		if err != nil {
			err = autorest.NewErrorWithError(err, "compute.VirtualMachineScaleSetVMsClient", "Reimage", resp, "Failure responding to request")
		}
	}()
	return resultChan, errChan
}

// ReimagePreparer prepares the Reimage request.
func (client VirtualMachineScaleSetVMsClient) ReimagePreparer(resourceGroupName string, VMScaleSetName string, instanceID string, cancel <-chan struct{}) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"instanceId":        autorest.Encode("path", instanceID),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"vmScaleSetName":    autorest.Encode("path", VMScaleSetName),
	}

	const APIVersion = "2017-03-30"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}/virtualmachines/{instanceId}/reimage", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{Cancel: cancel})
}

// ReimageSender sends the Reimage request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualMachineScaleSetVMsClient) ReimageSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoPollForAsynchronous(client.PollingDelay))
}

// ReimageResponder handles the response to the Reimage request. The method always
// closes the http.Response Body.
func (client VirtualMachineScaleSetVMsClient) ReimageResponder(resp *http.Response) (result OperationStatusResponse, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ReimageAll allows you to re-image all the disks ( including data disks ) in the a VM scale set instance. This
// operation is only supported for managed disks. This method may poll for completion. Polling can be canceled by
// passing the cancel channel argument. The channel will be used to cancel polling and any outstanding HTTP requests.
//
// resourceGroupName is the name of the resource group. VMScaleSetName is the name of the VM scale set. instanceID is
// the instance ID of the virtual machine.
func (client VirtualMachineScaleSetVMsClient) ReimageAll(resourceGroupName string, VMScaleSetName string, instanceID string, cancel <-chan struct{}) (<-chan OperationStatusResponse, <-chan error) {
	resultChan := make(chan OperationStatusResponse, 1)
	errChan := make(chan error, 1)
	go func() {
		var err error
		var result OperationStatusResponse
		defer func() {
			if err != nil {
				errChan <- err
			}
			resultChan <- result
			close(resultChan)
			close(errChan)
		}()
		req, err := client.ReimageAllPreparer(resourceGroupName, VMScaleSetName, instanceID, cancel)
		if err != nil {
			err = autorest.NewErrorWithError(err, "compute.VirtualMachineScaleSetVMsClient", "ReimageAll", nil, "Failure preparing request")
			return
		}

		resp, err := client.ReimageAllSender(req)
		if err != nil {
			result.Response = autorest.Response{Response: resp}
			err = autorest.NewErrorWithError(err, "compute.VirtualMachineScaleSetVMsClient", "ReimageAll", resp, "Failure sending request")
			return
		}

		result, err = client.ReimageAllResponder(resp)
		if err != nil {
			err = autorest.NewErrorWithError(err, "compute.VirtualMachineScaleSetVMsClient", "ReimageAll", resp, "Failure responding to request")
		}
	}()
	return resultChan, errChan
}

// ReimageAllPreparer prepares the ReimageAll request.
func (client VirtualMachineScaleSetVMsClient) ReimageAllPreparer(resourceGroupName string, VMScaleSetName string, instanceID string, cancel <-chan struct{}) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"instanceId":        autorest.Encode("path", instanceID),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"vmScaleSetName":    autorest.Encode("path", VMScaleSetName),
	}

	const APIVersion = "2017-03-30"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}/virtualmachines/{instanceId}/reimageall", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{Cancel: cancel})
}

// ReimageAllSender sends the ReimageAll request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualMachineScaleSetVMsClient) ReimageAllSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoPollForAsynchronous(client.PollingDelay))
}

// ReimageAllResponder handles the response to the ReimageAll request. The method always
// closes the http.Response Body.
func (client VirtualMachineScaleSetVMsClient) ReimageAllResponder(resp *http.Response) (result OperationStatusResponse, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Restart restarts a virtual machine in a VM scale set. This method may poll for completion. Polling can be canceled
// by passing the cancel channel argument. The channel will be used to cancel polling and any outstanding HTTP
// requests.
//
// resourceGroupName is the name of the resource group. VMScaleSetName is the name of the VM scale set. instanceID is
// the instance ID of the virtual machine.
func (client VirtualMachineScaleSetVMsClient) Restart(resourceGroupName string, VMScaleSetName string, instanceID string, cancel <-chan struct{}) (<-chan OperationStatusResponse, <-chan error) {
	resultChan := make(chan OperationStatusResponse, 1)
	errChan := make(chan error, 1)
	go func() {
		var err error
		var result OperationStatusResponse
		defer func() {
			if err != nil {
				errChan <- err
			}
			resultChan <- result
			close(resultChan)
			close(errChan)
		}()
		req, err := client.RestartPreparer(resourceGroupName, VMScaleSetName, instanceID, cancel)
		if err != nil {
			err = autorest.NewErrorWithError(err, "compute.VirtualMachineScaleSetVMsClient", "Restart", nil, "Failure preparing request")
			return
		}

		resp, err := client.RestartSender(req)
		if err != nil {
			result.Response = autorest.Response{Response: resp}
			err = autorest.NewErrorWithError(err, "compute.VirtualMachineScaleSetVMsClient", "Restart", resp, "Failure sending request")
			return
		}

		result, err = client.RestartResponder(resp)
		if err != nil {
			err = autorest.NewErrorWithError(err, "compute.VirtualMachineScaleSetVMsClient", "Restart", resp, "Failure responding to request")
		}
	}()
	return resultChan, errChan
}

// RestartPreparer prepares the Restart request.
func (client VirtualMachineScaleSetVMsClient) RestartPreparer(resourceGroupName string, VMScaleSetName string, instanceID string, cancel <-chan struct{}) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"instanceId":        autorest.Encode("path", instanceID),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"vmScaleSetName":    autorest.Encode("path", VMScaleSetName),
	}

	const APIVersion = "2017-03-30"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}/virtualmachines/{instanceId}/restart", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{Cancel: cancel})
}

// RestartSender sends the Restart request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualMachineScaleSetVMsClient) RestartSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoPollForAsynchronous(client.PollingDelay))
}

// RestartResponder handles the response to the Restart request. The method always
// closes the http.Response Body.
func (client VirtualMachineScaleSetVMsClient) RestartResponder(resp *http.Response) (result OperationStatusResponse, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Start starts a virtual machine in a VM scale set. This method may poll for completion. Polling can be canceled by
// passing the cancel channel argument. The channel will be used to cancel polling and any outstanding HTTP requests.
//
// resourceGroupName is the name of the resource group. VMScaleSetName is the name of the VM scale set. instanceID is
// the instance ID of the virtual machine.
func (client VirtualMachineScaleSetVMsClient) Start(resourceGroupName string, VMScaleSetName string, instanceID string, cancel <-chan struct{}) (<-chan OperationStatusResponse, <-chan error) {
	resultChan := make(chan OperationStatusResponse, 1)
	errChan := make(chan error, 1)
	go func() {
		var err error
		var result OperationStatusResponse
		defer func() {
			if err != nil {
				errChan <- err
			}
			resultChan <- result
			close(resultChan)
			close(errChan)
		}()
		req, err := client.StartPreparer(resourceGroupName, VMScaleSetName, instanceID, cancel)
		if err != nil {
			err = autorest.NewErrorWithError(err, "compute.VirtualMachineScaleSetVMsClient", "Start", nil, "Failure preparing request")
			return
		}

		resp, err := client.StartSender(req)
		if err != nil {
			result.Response = autorest.Response{Response: resp}
			err = autorest.NewErrorWithError(err, "compute.VirtualMachineScaleSetVMsClient", "Start", resp, "Failure sending request")
			return
		}

		result, err = client.StartResponder(resp)
		if err != nil {
			err = autorest.NewErrorWithError(err, "compute.VirtualMachineScaleSetVMsClient", "Start", resp, "Failure responding to request")
		}
	}()
	return resultChan, errChan
}

// StartPreparer prepares the Start request.
func (client VirtualMachineScaleSetVMsClient) StartPreparer(resourceGroupName string, VMScaleSetName string, instanceID string, cancel <-chan struct{}) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"instanceId":        autorest.Encode("path", instanceID),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"vmScaleSetName":    autorest.Encode("path", VMScaleSetName),
	}

	const APIVersion = "2017-03-30"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}/virtualmachines/{instanceId}/start", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{Cancel: cancel})
}

// StartSender sends the Start request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualMachineScaleSetVMsClient) StartSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoPollForAsynchronous(client.PollingDelay))
}

// StartResponder handles the response to the Start request. The method always
// closes the http.Response Body.
func (client VirtualMachineScaleSetVMsClient) StartResponder(resp *http.Response) (result OperationStatusResponse, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
