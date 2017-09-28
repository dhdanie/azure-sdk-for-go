package face

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

// ListGroupClient is the an API for face detection, verification, and identification.
type ListGroupClient struct {
	ManagementClient
}

// NewListGroupClient creates an instance of the ListGroupClient client.
func NewListGroupClient(subscriptionKey string, azureRegion AzureRegions) ListGroupClient {
	return ListGroupClient{New(subscriptionKey, azureRegion)}
}

// AddFace add a face to a face list. The input face is specified as an image with a targetFace rectangle. It returns a
// persistedFaceId representing the added face, and persistedFaceId will not expire.
//
// faceListID is id referencing a Face List. userData is user-specified data about the face list for any purpose. The
// maximum length is 1KB. targetFace is a face rectangle to specify the target face to be added into the face list, in
// the format of "targetFace=left,top,width,height". E.g. "targetFace=10,10,100,100". If there is more than one face in
// the image, targetFace is required to specify which face to add. No targetFace means there is only one face detected
// in the entire image.
func (client ListGroupClient) AddFace(faceListID string, userData string, targetFace string) (result autorest.Response, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: faceListID,
			Constraints: []validation.Constraint{{Target: "faceListID", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "faceListID", Name: validation.Pattern, Rule: `^[a-z0-9-_]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "face.ListGroupClient", "AddFace")
	}

	req, err := client.AddFacePreparer(faceListID, userData, targetFace)
	if err != nil {
		err = autorest.NewErrorWithError(err, "face.ListGroupClient", "AddFace", nil, "Failure preparing request")
		return
	}

	resp, err := client.AddFaceSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "face.ListGroupClient", "AddFace", resp, "Failure sending request")
		return
	}

	result, err = client.AddFaceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "face.ListGroupClient", "AddFace", resp, "Failure responding to request")
	}

	return
}

// AddFacePreparer prepares the AddFace request.
func (client ListGroupClient) AddFacePreparer(faceListID string, userData string, targetFace string) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"AzureRegion": client.AzureRegion,
	}

	pathParameters := map[string]interface{}{
		"faceListId": autorest.Encode("path", faceListID),
	}

	queryParameters := map[string]interface{}{}
	if len(userData) > 0 {
		queryParameters["userData"] = autorest.Encode("query", userData)
	}
	if len(targetFace) > 0 {
		queryParameters["targetFace"] = autorest.Encode("query", targetFace)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithCustomBaseURL("https://{AzureRegion}.api.cognitive.microsoft.com/face/v1.0", urlParameters),
		autorest.WithPathParameters("/facelists/{faceListId}/persistedFaces", pathParameters),
		autorest.WithQueryParameters(queryParameters),
		autorest.WithHeader("Ocp-Apim-Subscription-Key", client.SubscriptionKey))
	return preparer.Prepare(&http.Request{})
}

// AddFaceSender sends the AddFace request. The method will close the
// http.Response Body if it receives an error.
func (client ListGroupClient) AddFaceSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// AddFaceResponder handles the response to the AddFace request. The method always
// closes the http.Response Body.
func (client ListGroupClient) AddFaceResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// AddFaceFromStream add a face to a face list. The input face is specified as an image with a targetFace rectangle. It
// returns a persistedFaceId representing the added face, and persistedFaceId will not expire.
//
// faceListID is id referencing a Face List. userData is user-specified data about the face list for any purpose. The
// maximum length is 1KB. targetFace is a face rectangle to specify the target face to be added into the face list, in
// the format of "targetFace=left,top,width,height". E.g. "targetFace=10,10,100,100". If there is more than one face in
// the image, targetFace is required to specify which face to add. No targetFace means there is only one face detected
// in the entire image.
func (client ListGroupClient) AddFaceFromStream(faceListID string, userData string, targetFace string) (result autorest.Response, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: faceListID,
			Constraints: []validation.Constraint{{Target: "faceListID", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "faceListID", Name: validation.Pattern, Rule: `^[a-z0-9-_]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "face.ListGroupClient", "AddFaceFromStream")
	}

	req, err := client.AddFaceFromStreamPreparer(faceListID, userData, targetFace)
	if err != nil {
		err = autorest.NewErrorWithError(err, "face.ListGroupClient", "AddFaceFromStream", nil, "Failure preparing request")
		return
	}

	resp, err := client.AddFaceFromStreamSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "face.ListGroupClient", "AddFaceFromStream", resp, "Failure sending request")
		return
	}

	result, err = client.AddFaceFromStreamResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "face.ListGroupClient", "AddFaceFromStream", resp, "Failure responding to request")
	}

	return
}

// AddFaceFromStreamPreparer prepares the AddFaceFromStream request.
func (client ListGroupClient) AddFaceFromStreamPreparer(faceListID string, userData string, targetFace string) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"AzureRegion": client.AzureRegion,
	}

	pathParameters := map[string]interface{}{
		"faceListId": autorest.Encode("path", faceListID),
	}

	queryParameters := map[string]interface{}{}
	if len(userData) > 0 {
		queryParameters["userData"] = autorest.Encode("query", userData)
	}
	if len(targetFace) > 0 {
		queryParameters["targetFace"] = autorest.Encode("query", targetFace)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithCustomBaseURL("https://{AzureRegion}.api.cognitive.microsoft.com/face/v1.0", urlParameters),
		autorest.WithPathParameters("/facelists/{faceListId}/persistedFaces", pathParameters),
		autorest.WithQueryParameters(queryParameters),
		autorest.WithHeader("Ocp-Apim-Subscription-Key", client.SubscriptionKey))
	return preparer.Prepare(&http.Request{})
}

// AddFaceFromStreamSender sends the AddFaceFromStream request. The method will close the
// http.Response Body if it receives an error.
func (client ListGroupClient) AddFaceFromStreamSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// AddFaceFromStreamResponder handles the response to the AddFaceFromStream request. The method always
// closes the http.Response Body.
func (client ListGroupClient) AddFaceFromStreamResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Create create an empty face list. Up to 64 face lists are allowed to exist in one subscription.
//
// faceListID is id referencing a particular face list.
func (client ListGroupClient) Create(faceListID string, body CreateFaceListRequest) (result autorest.Response, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: faceListID,
			Constraints: []validation.Constraint{{Target: "faceListID", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "faceListID", Name: validation.Pattern, Rule: `^[a-z0-9-_]+$`, Chain: nil}}},
		{TargetValue: body,
			Constraints: []validation.Constraint{{Target: "body.Name", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "body.Name", Name: validation.MaxLength, Rule: 128, Chain: nil}}},
				{Target: "body.UserData", Name: validation.Null, Rule: false,
					Chain: []validation.Constraint{{Target: "body.UserData", Name: validation.MaxLength, Rule: 16384, Chain: nil}}}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "face.ListGroupClient", "Create")
	}

	req, err := client.CreatePreparer(faceListID, body)
	if err != nil {
		err = autorest.NewErrorWithError(err, "face.ListGroupClient", "Create", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "face.ListGroupClient", "Create", resp, "Failure sending request")
		return
	}

	result, err = client.CreateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "face.ListGroupClient", "Create", resp, "Failure responding to request")
	}

	return
}

// CreatePreparer prepares the Create request.
func (client ListGroupClient) CreatePreparer(faceListID string, body CreateFaceListRequest) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"AzureRegion": client.AzureRegion,
	}

	pathParameters := map[string]interface{}{
		"faceListId": autorest.Encode("path", faceListID),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithCustomBaseURL("https://{AzureRegion}.api.cognitive.microsoft.com/face/v1.0", urlParameters),
		autorest.WithPathParameters("/facelists/{faceListId}", pathParameters),
		autorest.WithJSON(body),
		autorest.WithHeader("Ocp-Apim-Subscription-Key", client.SubscriptionKey))
	return preparer.Prepare(&http.Request{})
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client ListGroupClient) CreateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client ListGroupClient) CreateResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Delete delete an existing face list according to faceListId. Persisted face images in the face list will also be
// deleted.
//
// faceListID is id referencing a Face List.
func (client ListGroupClient) Delete(faceListID string) (result autorest.Response, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: faceListID,
			Constraints: []validation.Constraint{{Target: "faceListID", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "faceListID", Name: validation.Pattern, Rule: `^[a-z0-9-_]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "face.ListGroupClient", "Delete")
	}

	req, err := client.DeletePreparer(faceListID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "face.ListGroupClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "face.ListGroupClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "face.ListGroupClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client ListGroupClient) DeletePreparer(faceListID string) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"AzureRegion": client.AzureRegion,
	}

	pathParameters := map[string]interface{}{
		"faceListId": autorest.Encode("path", faceListID),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithCustomBaseURL("https://{AzureRegion}.api.cognitive.microsoft.com/face/v1.0", urlParameters),
		autorest.WithPathParameters("/facelists/{faceListId}", pathParameters),
		autorest.WithHeader("Ocp-Apim-Subscription-Key", client.SubscriptionKey))
	return preparer.Prepare(&http.Request{})
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client ListGroupClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client ListGroupClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// DeleteFace delete an existing face from a face list (given by a persisitedFaceId and a faceListId). Persisted image
// related to the face will also be deleted.
//
// faceListID is faceListId of an existing face list. persistedFaceID is persistedFaceId of an existing face.
func (client ListGroupClient) DeleteFace(faceListID string, persistedFaceID string) (result autorest.Response, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: faceListID,
			Constraints: []validation.Constraint{{Target: "faceListID", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "faceListID", Name: validation.Pattern, Rule: `^[a-z0-9-_]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "face.ListGroupClient", "DeleteFace")
	}

	req, err := client.DeleteFacePreparer(faceListID, persistedFaceID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "face.ListGroupClient", "DeleteFace", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteFaceSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "face.ListGroupClient", "DeleteFace", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteFaceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "face.ListGroupClient", "DeleteFace", resp, "Failure responding to request")
	}

	return
}

// DeleteFacePreparer prepares the DeleteFace request.
func (client ListGroupClient) DeleteFacePreparer(faceListID string, persistedFaceID string) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"AzureRegion": client.AzureRegion,
	}

	pathParameters := map[string]interface{}{
		"faceListId":      autorest.Encode("path", faceListID),
		"persistedFaceId": autorest.Encode("path", persistedFaceID),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithCustomBaseURL("https://{AzureRegion}.api.cognitive.microsoft.com/face/v1.0", urlParameters),
		autorest.WithPathParameters("/facelists/{faceListId}/persistedFaces/{persistedFaceId}", pathParameters),
		autorest.WithHeader("Ocp-Apim-Subscription-Key", client.SubscriptionKey))
	return preparer.Prepare(&http.Request{})
}

// DeleteFaceSender sends the DeleteFace request. The method will close the
// http.Response Body if it receives an error.
func (client ListGroupClient) DeleteFaceSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// DeleteFaceResponder handles the response to the DeleteFace request. The method always
// closes the http.Response Body.
func (client ListGroupClient) DeleteFaceResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get retrieve a face list's information.
//
// faceListID is id referencing a Face List.
func (client ListGroupClient) Get(faceListID string) (result GetFaceListResult, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: faceListID,
			Constraints: []validation.Constraint{{Target: "faceListID", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "faceListID", Name: validation.Pattern, Rule: `^[a-z0-9-_]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "face.ListGroupClient", "Get")
	}

	req, err := client.GetPreparer(faceListID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "face.ListGroupClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "face.ListGroupClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "face.ListGroupClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client ListGroupClient) GetPreparer(faceListID string) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"AzureRegion": client.AzureRegion,
	}

	pathParameters := map[string]interface{}{
		"faceListId": autorest.Encode("path", faceListID),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithCustomBaseURL("https://{AzureRegion}.api.cognitive.microsoft.com/face/v1.0", urlParameters),
		autorest.WithPathParameters("/facelists/{faceListId}", pathParameters),
		autorest.WithHeader("Ocp-Apim-Subscription-Key", client.SubscriptionKey))
	return preparer.Prepare(&http.Request{})
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ListGroupClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ListGroupClient) GetResponder(resp *http.Response) (result GetFaceListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List retrieve information about all existing face lists. Only faceListId, name and userData will be returned.
func (client ListGroupClient) List() (result ListGetFaceListResult, err error) {
	req, err := client.ListPreparer()
	if err != nil {
		err = autorest.NewErrorWithError(err, "face.ListGroupClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "face.ListGroupClient", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "face.ListGroupClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client ListGroupClient) ListPreparer() (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"AzureRegion": client.AzureRegion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithCustomBaseURL("https://{AzureRegion}.api.cognitive.microsoft.com/face/v1.0", urlParameters),
		autorest.WithPath("/facelists"),
		autorest.WithHeader("Ocp-Apim-Subscription-Key", client.SubscriptionKey))
	return preparer.Prepare(&http.Request{})
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client ListGroupClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client ListGroupClient) ListResponder(resp *http.Response) (result ListGetFaceListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Update update information of a face list.
//
// faceListID is id referencing a Face List.
func (client ListGroupClient) Update(faceListID string, body CreateFaceListRequest) (result autorest.Response, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: faceListID,
			Constraints: []validation.Constraint{{Target: "faceListID", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "faceListID", Name: validation.Pattern, Rule: `^[a-z0-9-_]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "face.ListGroupClient", "Update")
	}

	req, err := client.UpdatePreparer(faceListID, body)
	if err != nil {
		err = autorest.NewErrorWithError(err, "face.ListGroupClient", "Update", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "face.ListGroupClient", "Update", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "face.ListGroupClient", "Update", resp, "Failure responding to request")
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client ListGroupClient) UpdatePreparer(faceListID string, body CreateFaceListRequest) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"AzureRegion": client.AzureRegion,
	}

	pathParameters := map[string]interface{}{
		"faceListId": autorest.Encode("path", faceListID),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPatch(),
		autorest.WithCustomBaseURL("https://{AzureRegion}.api.cognitive.microsoft.com/face/v1.0", urlParameters),
		autorest.WithPathParameters("/facelists/{faceListId}", pathParameters),
		autorest.WithJSON(body),
		autorest.WithHeader("Ocp-Apim-Subscription-Key", client.SubscriptionKey))
	return preparer.Prepare(&http.Request{})
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client ListGroupClient) UpdateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client ListGroupClient) UpdateResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}
