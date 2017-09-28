package cosmosdb

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

// DatabaseAccountKind enumerates the values for database account kind.
type DatabaseAccountKind string

const (
	// GlobalDocumentDB specifies the global document db state for database account kind.
	GlobalDocumentDB DatabaseAccountKind = "GlobalDocumentDB"
	// MongoDB specifies the mongo db state for database account kind.
	MongoDB DatabaseAccountKind = "MongoDB"
	// Parse specifies the parse state for database account kind.
	Parse DatabaseAccountKind = "Parse"
)

// DatabaseAccountOfferType enumerates the values for database account offer type.
type DatabaseAccountOfferType string

const (
	// Standard specifies the standard state for database account offer type.
	Standard DatabaseAccountOfferType = "Standard"
)

// DefaultConsistencyLevel enumerates the values for default consistency level.
type DefaultConsistencyLevel string

const (
	// BoundedStaleness specifies the bounded staleness state for default consistency level.
	BoundedStaleness DefaultConsistencyLevel = "BoundedStaleness"
	// ConsistentPrefix specifies the consistent prefix state for default consistency level.
	ConsistentPrefix DefaultConsistencyLevel = "ConsistentPrefix"
	// Eventual specifies the eventual state for default consistency level.
	Eventual DefaultConsistencyLevel = "Eventual"
	// Session specifies the session state for default consistency level.
	Session DefaultConsistencyLevel = "Session"
	// Strong specifies the strong state for default consistency level.
	Strong DefaultConsistencyLevel = "Strong"
)

// KeyKind enumerates the values for key kind.
type KeyKind string

const (
	// Primary specifies the primary state for key kind.
	Primary KeyKind = "primary"
	// PrimaryReadonly specifies the primary readonly state for key kind.
	PrimaryReadonly KeyKind = "primaryReadonly"
	// Secondary specifies the secondary state for key kind.
	Secondary KeyKind = "secondary"
	// SecondaryReadonly specifies the secondary readonly state for key kind.
	SecondaryReadonly KeyKind = "secondaryReadonly"
)

// ConsistencyPolicy is the consistency policy for the Cosmos DB database account.
type ConsistencyPolicy struct {
	DefaultConsistencyLevel DefaultConsistencyLevel `json:"defaultConsistencyLevel,omitempty"`
	MaxStalenessPrefix      *int64                  `json:"maxStalenessPrefix,omitempty"`
	MaxIntervalInSeconds    *int32                  `json:"maxIntervalInSeconds,omitempty"`
}

// DatabaseAccount is an Azure Cosmos DB database account.
type DatabaseAccount struct {
	autorest.Response          `json:"-"`
	ID                         *string             `json:"id,omitempty"`
	Name                       *string             `json:"name,omitempty"`
	Type                       *string             `json:"type,omitempty"`
	Location                   *string             `json:"location,omitempty"`
	Tags                       *map[string]*string `json:"tags,omitempty"`
	Kind                       DatabaseAccountKind `json:"kind,omitempty"`
	*DatabaseAccountProperties `json:"properties,omitempty"`
}

// DatabaseAccountConnectionString is connection string for the Cosmos DB account
type DatabaseAccountConnectionString struct {
	ConnectionString *string `json:"connectionString,omitempty"`
	Description      *string `json:"description,omitempty"`
}

// DatabaseAccountCreateUpdateParameters is parameters to create and update Cosmos DB database accounts.
type DatabaseAccountCreateUpdateParameters struct {
	ID                                     *string             `json:"id,omitempty"`
	Name                                   *string             `json:"name,omitempty"`
	Type                                   *string             `json:"type,omitempty"`
	Location                               *string             `json:"location,omitempty"`
	Tags                                   *map[string]*string `json:"tags,omitempty"`
	Kind                                   DatabaseAccountKind `json:"kind,omitempty"`
	*DatabaseAccountCreateUpdateProperties `json:"properties,omitempty"`
}

// DatabaseAccountCreateUpdateProperties is properties to create and update Azure Cosmos DB database accounts.
type DatabaseAccountCreateUpdateProperties struct {
	ConsistencyPolicy        *ConsistencyPolicy `json:"consistencyPolicy,omitempty"`
	Locations                *[]Location        `json:"locations,omitempty"`
	DatabaseAccountOfferType *string            `json:"databaseAccountOfferType,omitempty"`
	IPRangeFilter            *string            `json:"ipRangeFilter,omitempty"`
	EnableAutomaticFailover  *bool              `json:"enableAutomaticFailover,omitempty"`
}

// DatabaseAccountListConnectionStringsResult is the connection strings for the given database account.
type DatabaseAccountListConnectionStringsResult struct {
	autorest.Response `json:"-"`
	ConnectionStrings *[]DatabaseAccountConnectionString `json:"connectionStrings,omitempty"`
}

// DatabaseAccountListKeysResult is the access keys for the given database account.
type DatabaseAccountListKeysResult struct {
	autorest.Response                      `json:"-"`
	PrimaryMasterKey                       *string `json:"primaryMasterKey,omitempty"`
	SecondaryMasterKey                     *string `json:"secondaryMasterKey,omitempty"`
	*DatabaseAccountListReadOnlyKeysResult `json:"properties,omitempty"`
}

// DatabaseAccountListReadOnlyKeysResult is the read-only access keys for the given database account.
type DatabaseAccountListReadOnlyKeysResult struct {
	autorest.Response          `json:"-"`
	PrimaryReadonlyMasterKey   *string `json:"primaryReadonlyMasterKey,omitempty"`
	SecondaryReadonlyMasterKey *string `json:"secondaryReadonlyMasterKey,omitempty"`
}

// DatabaseAccountPatchParameters is parameters for patching Azure Cosmos DB database account properties.
type DatabaseAccountPatchParameters struct {
	Tags *map[string]*string `json:"tags,omitempty"`
}

// DatabaseAccountProperties is properties for the database account.
type DatabaseAccountProperties struct {
	ProvisioningState        *string                  `json:"provisioningState,omitempty"`
	DocumentEndpoint         *string                  `json:"documentEndpoint,omitempty"`
	DatabaseAccountOfferType DatabaseAccountOfferType `json:"databaseAccountOfferType,omitempty"`
	IPRangeFilter            *string                  `json:"ipRangeFilter,omitempty"`
	EnableAutomaticFailover  *bool                    `json:"enableAutomaticFailover,omitempty"`
	ConsistencyPolicy        *ConsistencyPolicy       `json:"consistencyPolicy,omitempty"`
	WriteLocations           *[]Location              `json:"writeLocations,omitempty"`
	ReadLocations            *[]Location              `json:"readLocations,omitempty"`
	FailoverPolicies         *[]FailoverPolicy        `json:"failoverPolicies,omitempty"`
}

// DatabaseAccountRegenerateKeyParameters is parameters to regenerate the keys within the database account.
type DatabaseAccountRegenerateKeyParameters struct {
	KeyKind KeyKind `json:"keyKind,omitempty"`
}

// DatabaseAccountsListResult is the List operation response, that contains the database accounts and their properties.
type DatabaseAccountsListResult struct {
	autorest.Response `json:"-"`
	Value             *[]DatabaseAccount `json:"value,omitempty"`
}

// FailoverPolicies is the list of new failover policies for the failover priority change.
type FailoverPolicies struct {
	FailoverPolicies *[]FailoverPolicy `json:"failoverPolicies,omitempty"`
}

// FailoverPolicy is the failover policy for a given region of a database account.
type FailoverPolicy struct {
	ID               *string `json:"id,omitempty"`
	LocationName     *string `json:"locationName,omitempty"`
	FailoverPriority *int32  `json:"failoverPriority,omitempty"`
}

// Location is a region in which the Azure Cosmos DB database account is deployed.
type Location struct {
	ID                *string `json:"id,omitempty"`
	LocationName      *string `json:"locationName,omitempty"`
	DocumentEndpoint  *string `json:"documentEndpoint,omitempty"`
	ProvisioningState *string `json:"provisioningState,omitempty"`
	FailoverPriority  *int32  `json:"failoverPriority,omitempty"`
}

// Operation is REST API operation
type Operation struct {
	Name    *string           `json:"name,omitempty"`
	Display *OperationDisplay `json:"display,omitempty"`
}

// OperationDisplay is the object that represents the operation.
type OperationDisplay struct {
	Provider    *string `json:"Provider,omitempty"`
	Resource    *string `json:"Resource,omitempty"`
	Operation   *string `json:"Operation,omitempty"`
	Description *string `json:"Description,omitempty"`
}

// OperationListResult is result of the request to list Resource Provider operations. It contains a list of operations
// and a URL link to get the next set of results.
type OperationListResult struct {
	autorest.Response `json:"-"`
	Value             *[]Operation `json:"value,omitempty"`
	NextLink          *string      `json:"nextLink,omitempty"`
}

// OperationListResultPreparer prepares a request to retrieve the next set of results. It returns
// nil if no more results exist.
func (client OperationListResult) OperationListResultPreparer() (*http.Request, error) {
	if client.NextLink == nil || len(to.String(client.NextLink)) <= 0 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(client.NextLink)))
}

// Resource is a database account resource.
type Resource struct {
	ID       *string             `json:"id,omitempty"`
	Name     *string             `json:"name,omitempty"`
	Type     *string             `json:"type,omitempty"`
	Location *string             `json:"location,omitempty"`
	Tags     *map[string]*string `json:"tags,omitempty"`
}
