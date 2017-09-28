package visualstudio

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
)

// AccountResource is the response to an account resource GET request.
type AccountResource struct {
	autorest.Response `json:"-"`
	ID                *string             `json:"id,omitempty"`
	Location          *string             `json:"location,omitempty"`
	Name              *string             `json:"name,omitempty"`
	Tags              *map[string]*string `json:"tags,omitempty"`
	Type              *string             `json:"type,omitempty"`
	Properties        *map[string]*string `json:"properties,omitempty"`
}

// AccountResourceListResult is the response to an account resource list GET request.
type AccountResourceListResult struct {
	autorest.Response `json:"-"`
	Value             *[]AccountResource `json:"value,omitempty"`
}

// AccountResourceRequest is the body of a PUT request to modify a Visual Studio account resource.
type AccountResourceRequest struct {
	AccountName   *string                 `json:"accountName,omitempty"`
	Location      *string                 `json:"location,omitempty"`
	OperationType *map[string]interface{} `json:"operationType,omitempty"`
	Properties    *map[string]*string     `json:"properties,omitempty"`
	Tags          *map[string]*string     `json:"tags,omitempty"`
}

// CheckNameAvailabilityParameter is the body of a POST request to check name availability.
type CheckNameAvailabilityParameter struct {
	ResourceName *string `json:"resourceName,omitempty"`
	ResourceType *string `json:"resourceType,omitempty"`
}

// CheckNameAvailabilityResult is the response to a name availability request.
type CheckNameAvailabilityResult struct {
	autorest.Response `json:"-"`
	Message           *string `json:"message,omitempty"`
	NameAvailable     *bool   `json:"nameAvailable,omitempty"`
}

// ExtensionResource is the response to an extension resource GET request.
type ExtensionResource struct {
	autorest.Response `json:"-"`
	ID                *string                `json:"id,omitempty"`
	Location          *string                `json:"location,omitempty"`
	Name              *string                `json:"name,omitempty"`
	Tags              *map[string]*string    `json:"tags,omitempty"`
	Type              *string                `json:"type,omitempty"`
	Plan              *ExtensionResourcePlan `json:"plan,omitempty"`
	Properties        *map[string]*string    `json:"properties,omitempty"`
}

// ExtensionResourceListResult is the response to an extension resource list GET request.
type ExtensionResourceListResult struct {
	autorest.Response `json:"-"`
	Value             *[]ExtensionResource `json:"value,omitempty"`
}

// ExtensionResourcePlan is plan data for an extension resource.
type ExtensionResourcePlan struct {
	Name          *string `json:"name,omitempty"`
	Product       *string `json:"product,omitempty"`
	PromotionCode *string `json:"promotionCode,omitempty"`
	Publisher     *string `json:"publisher,omitempty"`
	Version       *string `json:"version,omitempty"`
}

// ExtensionResourceRequest is the body of an extension resource PUT request.
type ExtensionResourceRequest struct {
	Location   *string                `json:"location,omitempty"`
	Plan       *ExtensionResourcePlan `json:"plan,omitempty"`
	Properties *map[string]*string    `json:"properties,omitempty"`
	Tags       *map[string]*string    `json:"tags,omitempty"`
}

// Operation is properties of an operation supported by the resource provider.
type Operation struct {
	Display *OperationProperties `json:"display,omitempty"`
	Name    *string              `json:"name,omitempty"`
}

// OperationListResult is container for a list of operations supported by a resource provider.
type OperationListResult struct {
	autorest.Response `json:"-"`
	Value             *[]Operation `json:"value,omitempty"`
}

// OperationProperties is properties of an operation supported by the resource provider.
type OperationProperties struct {
	Description *string `json:"description,omitempty"`
	Operation   *string `json:"operation,omitempty"`
	Provider    *string `json:"provider,omitempty"`
	Resource    *string `json:"resource,omitempty"`
}

// ProjectResource is a Visual Studio Team Services project resource.
type ProjectResource struct {
	autorest.Response `json:"-"`
	ID                *string             `json:"id,omitempty"`
	Location          *string             `json:"location,omitempty"`
	Name              *string             `json:"name,omitempty"`
	Tags              *map[string]*string `json:"tags,omitempty"`
	Type              *string             `json:"type,omitempty"`
	Properties        *map[string]*string `json:"properties,omitempty"`
}

// ProjectResourceListResult is the response to a request to list Team Services project resources in a resource
// group/account.
type ProjectResourceListResult struct {
	autorest.Response `json:"-"`
	Value             *[]ProjectResource `json:"value,omitempty"`
}

// Resource is a generic Azure Resource Manager resource.
type Resource struct {
	ID       *string             `json:"id,omitempty"`
	Location *string             `json:"location,omitempty"`
	Name     *string             `json:"name,omitempty"`
	Tags     *map[string]*string `json:"tags,omitempty"`
	Type     *string             `json:"type,omitempty"`
}
