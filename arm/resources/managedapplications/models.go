package managedapplications

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
	"github.com/Azure/go-autorest/autorest/to"
	"net/http"
)

// ApplianceArtifactType enumerates the values for appliance artifact type.
type ApplianceArtifactType string

const (
	// Custom specifies the custom state for appliance artifact type.
	Custom ApplianceArtifactType = "Custom"
	// Template specifies the template state for appliance artifact type.
	Template ApplianceArtifactType = "Template"
)

// ApplianceLockLevel enumerates the values for appliance lock level.
type ApplianceLockLevel string

const (
	// CanNotDelete specifies the can not delete state for appliance lock level.
	CanNotDelete ApplianceLockLevel = "CanNotDelete"
	// None specifies the none state for appliance lock level.
	None ApplianceLockLevel = "None"
	// ReadOnly specifies the read only state for appliance lock level.
	ReadOnly ApplianceLockLevel = "ReadOnly"
)

// ProvisioningState enumerates the values for provisioning state.
type ProvisioningState string

const (
	// Accepted specifies the accepted state for provisioning state.
	Accepted ProvisioningState = "Accepted"
	// Canceled specifies the canceled state for provisioning state.
	Canceled ProvisioningState = "Canceled"
	// Created specifies the created state for provisioning state.
	Created ProvisioningState = "Created"
	// Creating specifies the creating state for provisioning state.
	Creating ProvisioningState = "Creating"
	// Deleted specifies the deleted state for provisioning state.
	Deleted ProvisioningState = "Deleted"
	// Deleting specifies the deleting state for provisioning state.
	Deleting ProvisioningState = "Deleting"
	// Failed specifies the failed state for provisioning state.
	Failed ProvisioningState = "Failed"
	// Ready specifies the ready state for provisioning state.
	Ready ProvisioningState = "Ready"
	// Running specifies the running state for provisioning state.
	Running ProvisioningState = "Running"
	// Succeeded specifies the succeeded state for provisioning state.
	Succeeded ProvisioningState = "Succeeded"
	// Updating specifies the updating state for provisioning state.
	Updating ProvisioningState = "Updating"
)

// ResourceIdentityType enumerates the values for resource identity type.
type ResourceIdentityType string

const (
	// SystemAssigned specifies the system assigned state for resource identity type.
	SystemAssigned ResourceIdentityType = "SystemAssigned"
)

// Appliance is information about appliance.
type Appliance struct {
	autorest.Response    `json:"-"`
	ID                   *string             `json:"id,omitempty"`
	Name                 *string             `json:"name,omitempty"`
	Type                 *string             `json:"type,omitempty"`
	Location             *string             `json:"location,omitempty"`
	Tags                 *map[string]*string `json:"tags,omitempty"`
	ManagedBy            *string             `json:"managedBy,omitempty"`
	Sku                  *Sku                `json:"sku,omitempty"`
	Identity             *Identity           `json:"identity,omitempty"`
	*ApplianceProperties `json:"properties,omitempty"`
	Plan                 *Plan   `json:"plan,omitempty"`
	Kind                 *string `json:"kind,omitempty"`
}

// ApplianceArtifact is appliance artifact.
type ApplianceArtifact struct {
	Name *string               `json:"name,omitempty"`
	URI  *string               `json:"uri,omitempty"`
	Type ApplianceArtifactType `json:"type,omitempty"`
}

// ApplianceDefinition is information about appliance definition.
type ApplianceDefinition struct {
	autorest.Response              `json:"-"`
	ID                             *string             `json:"id,omitempty"`
	Name                           *string             `json:"name,omitempty"`
	Type                           *string             `json:"type,omitempty"`
	Location                       *string             `json:"location,omitempty"`
	Tags                           *map[string]*string `json:"tags,omitempty"`
	ManagedBy                      *string             `json:"managedBy,omitempty"`
	Sku                            *Sku                `json:"sku,omitempty"`
	Identity                       *Identity           `json:"identity,omitempty"`
	*ApplianceDefinitionProperties `json:"properties,omitempty"`
}

// ApplianceDefinitionListResult is list of appliance definitions.
type ApplianceDefinitionListResult struct {
	autorest.Response `json:"-"`
	Value             *[]ApplianceDefinition `json:"value,omitempty"`
	NextLink          *string                `json:"nextLink,omitempty"`
}

// ApplianceDefinitionListResultPreparer prepares a request to retrieve the next set of results. It returns
// nil if no more results exist.
func (client ApplianceDefinitionListResult) ApplianceDefinitionListResultPreparer() (*http.Request, error) {
	if client.NextLink == nil || len(to.String(client.NextLink)) <= 0 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(client.NextLink)))
}

// ApplianceDefinitionProperties is the appliance definition properties.
type ApplianceDefinitionProperties struct {
	LockLevel      ApplianceLockLevel                `json:"lockLevel,omitempty"`
	DisplayName    *string                           `json:"displayName,omitempty"`
	Authorizations *[]ApplianceProviderAuthorization `json:"authorizations,omitempty"`
	Artifacts      *[]ApplianceArtifact              `json:"artifacts,omitempty"`
	Description    *string                           `json:"description,omitempty"`
	PackageFileURI *string                           `json:"packageFileUri,omitempty"`
}

// ApplianceListResult is list of appliances.
type ApplianceListResult struct {
	autorest.Response `json:"-"`
	Value             *[]Appliance `json:"value,omitempty"`
	NextLink          *string      `json:"nextLink,omitempty"`
}

// ApplianceListResultPreparer prepares a request to retrieve the next set of results. It returns
// nil if no more results exist.
func (client ApplianceListResult) ApplianceListResultPreparer() (*http.Request, error) {
	if client.NextLink == nil || len(to.String(client.NextLink)) <= 0 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(client.NextLink)))
}

// AppliancePatchable is information about appliance.
type AppliancePatchable struct {
	ID                            *string             `json:"id,omitempty"`
	Name                          *string             `json:"name,omitempty"`
	Type                          *string             `json:"type,omitempty"`
	Location                      *string             `json:"location,omitempty"`
	Tags                          *map[string]*string `json:"tags,omitempty"`
	ManagedBy                     *string             `json:"managedBy,omitempty"`
	Sku                           *Sku                `json:"sku,omitempty"`
	Identity                      *Identity           `json:"identity,omitempty"`
	*AppliancePropertiesPatchable `json:"properties,omitempty"`
	Plan                          *PlanPatchable `json:"plan,omitempty"`
	Kind                          *string        `json:"kind,omitempty"`
}

// ApplianceProperties is the appliance properties.
type ApplianceProperties struct {
	ManagedResourceGroupID *string                 `json:"managedResourceGroupId,omitempty"`
	ApplianceDefinitionID  *string                 `json:"applianceDefinitionId,omitempty"`
	Parameters             *map[string]interface{} `json:"parameters,omitempty"`
	Outputs                *map[string]interface{} `json:"outputs,omitempty"`
	ProvisioningState      ProvisioningState       `json:"provisioningState,omitempty"`
	UIDefinitionURI        *string                 `json:"uiDefinitionUri,omitempty"`
}

// AppliancePropertiesPatchable is the appliance properties.
type AppliancePropertiesPatchable struct {
	ManagedResourceGroupID *string                 `json:"managedResourceGroupId,omitempty"`
	ApplianceDefinitionID  *string                 `json:"applianceDefinitionId,omitempty"`
	Parameters             *map[string]interface{} `json:"parameters,omitempty"`
	Outputs                *map[string]interface{} `json:"outputs,omitempty"`
	ProvisioningState      ProvisioningState       `json:"provisioningState,omitempty"`
	UIDefinitionURI        *string                 `json:"uiDefinitionUri,omitempty"`
}

// ApplianceProviderAuthorization is the appliance provider authorization.
type ApplianceProviderAuthorization struct {
	PrincipalID      *string `json:"principalId,omitempty"`
	RoleDefinitionID *string `json:"roleDefinitionId,omitempty"`
}

// ErrorResponse is error reponse indicates ARM appliance is not able to process the incoming request. The reason is
// provided in the error message.
type ErrorResponse struct {
	HTTPStatus   *string `json:"httpStatus,omitempty"`
	ErrorCode    *string `json:"errorCode,omitempty"`
	ErrorMessage *string `json:"errorMessage,omitempty"`
}

// GenericResource is resource information.
type GenericResource struct {
	ID        *string             `json:"id,omitempty"`
	Name      *string             `json:"name,omitempty"`
	Type      *string             `json:"type,omitempty"`
	Location  *string             `json:"location,omitempty"`
	Tags      *map[string]*string `json:"tags,omitempty"`
	ManagedBy *string             `json:"managedBy,omitempty"`
	Sku       *Sku                `json:"sku,omitempty"`
	Identity  *Identity           `json:"identity,omitempty"`
}

// Identity is identity for the resource.
type Identity struct {
	PrincipalID *string              `json:"principalId,omitempty"`
	TenantID    *string              `json:"tenantId,omitempty"`
	Type        ResourceIdentityType `json:"type,omitempty"`
}

// Plan is plan for the appliance.
type Plan struct {
	Name          *string `json:"name,omitempty"`
	Publisher     *string `json:"publisher,omitempty"`
	Product       *string `json:"product,omitempty"`
	PromotionCode *string `json:"promotionCode,omitempty"`
	Version       *string `json:"version,omitempty"`
}

// PlanPatchable is plan for the appliance.
type PlanPatchable struct {
	Name          *string `json:"name,omitempty"`
	Publisher     *string `json:"publisher,omitempty"`
	Product       *string `json:"product,omitempty"`
	PromotionCode *string `json:"promotionCode,omitempty"`
	Version       *string `json:"version,omitempty"`
}

// Resource is resource information.
type Resource struct {
	ID       *string             `json:"id,omitempty"`
	Name     *string             `json:"name,omitempty"`
	Type     *string             `json:"type,omitempty"`
	Location *string             `json:"location,omitempty"`
	Tags     *map[string]*string `json:"tags,omitempty"`
}

// Sku is SKU for the resource.
type Sku struct {
	Name     *string `json:"name,omitempty"`
	Tier     *string `json:"tier,omitempty"`
	Size     *string `json:"size,omitempty"`
	Family   *string `json:"family,omitempty"`
	Model    *string `json:"model,omitempty"`
	Capacity *int32  `json:"capacity,omitempty"`
}
