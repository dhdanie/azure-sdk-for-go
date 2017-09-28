package graphrbac

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
	"github.com/Azure/go-autorest/autorest/date"
)

// UserType enumerates the values for user type.
type UserType string

const (
	// Guest specifies the guest state for user type.
	Guest UserType = "Guest"
	// Member specifies the member state for user type.
	Member UserType = "Member"
)

// AADObject is the properties of an Active Directory object.
type AADObject struct {
	autorest.Response       `json:"-"`
	ObjectID                *string   `json:"objectId,omitempty"`
	ObjectType              *string   `json:"objectType,omitempty"`
	DisplayName             *string   `json:"displayName,omitempty"`
	UserPrincipalName       *string   `json:"userPrincipalName,omitempty"`
	Mail                    *string   `json:"mail,omitempty"`
	MailEnabled             *bool     `json:"mailEnabled,omitempty"`
	MailNickname            *string   `json:"mailNickname,omitempty"`
	SecurityEnabled         *bool     `json:"securityEnabled,omitempty"`
	SignInName              *string   `json:"signInName,omitempty"`
	ServicePrincipalNames   *[]string `json:"servicePrincipalNames,omitempty"`
	UserType                *string   `json:"userType,omitempty"`
	UsageLocation           *string   `json:"usageLocation,omitempty"`
	AppID                   *string   `json:"appId,omitempty"`
	AppPermissions          *[]string `json:"appPermissions,omitempty"`
	AvailableToOtherTenants *bool     `json:"availableToOtherTenants,omitempty"`
	IdentifierUris          *[]string `json:"identifierUris,omitempty"`
	ReplyUrls               *[]string `json:"replyUrls,omitempty"`
	Homepage                *string   `json:"homepage,omitempty"`
}

// ADGroup is active Directory group information.
type ADGroup struct {
	autorest.Response `json:"-"`
	ObjectID          *string `json:"objectId,omitempty"`
	ObjectType        *string `json:"objectType,omitempty"`
	DisplayName       *string `json:"displayName,omitempty"`
	SecurityEnabled   *bool   `json:"securityEnabled,omitempty"`
	Mail              *string `json:"mail,omitempty"`
}

// Application is active Directory application information.
type Application struct {
	autorest.Response       `json:"-"`
	ObjectID                *string   `json:"objectId,omitempty"`
	ObjectType              *string   `json:"objectType,omitempty"`
	AppID                   *string   `json:"appId,omitempty"`
	AppPermissions          *[]string `json:"appPermissions,omitempty"`
	AvailableToOtherTenants *bool     `json:"availableToOtherTenants,omitempty"`
	DisplayName             *string   `json:"displayName,omitempty"`
	IdentifierUris          *[]string `json:"identifierUris,omitempty"`
	ReplyUrls               *[]string `json:"replyUrls,omitempty"`
	Homepage                *string   `json:"homepage,omitempty"`
	Oauth2AllowImplicitFlow *bool     `json:"oauth2AllowImplicitFlow,omitempty"`
}

// ApplicationCreateParameters is request parameters for creating a new application.
type ApplicationCreateParameters struct {
	AvailableToOtherTenants *bool                     `json:"availableToOtherTenants,omitempty"`
	DisplayName             *string                   `json:"displayName,omitempty"`
	Homepage                *string                   `json:"homepage,omitempty"`
	IdentifierUris          *[]string                 `json:"identifierUris,omitempty"`
	ReplyUrls               *[]string                 `json:"replyUrls,omitempty"`
	KeyCredentials          *[]KeyCredential          `json:"keyCredentials,omitempty"`
	PasswordCredentials     *[]PasswordCredential     `json:"passwordCredentials,omitempty"`
	Oauth2AllowImplicitFlow *bool                     `json:"oauth2AllowImplicitFlow,omitempty"`
	RequiredResourceAccess  *[]RequiredResourceAccess `json:"requiredResourceAccess,omitempty"`
}

// ApplicationListResult is application list operation result.
type ApplicationListResult struct {
	autorest.Response `json:"-"`
	Value             *[]Application `json:"value,omitempty"`
	OdataNextLink     *string        `json:"odata.nextLink,omitempty"`
}

// ApplicationUpdateParameters is request parameters for updating an existing application.
type ApplicationUpdateParameters struct {
	AvailableToOtherTenants *bool                     `json:"availableToOtherTenants,omitempty"`
	DisplayName             *string                   `json:"displayName,omitempty"`
	Homepage                *string                   `json:"homepage,omitempty"`
	IdentifierUris          *[]string                 `json:"identifierUris,omitempty"`
	ReplyUrls               *[]string                 `json:"replyUrls,omitempty"`
	KeyCredentials          *[]KeyCredential          `json:"keyCredentials,omitempty"`
	PasswordCredentials     *[]PasswordCredential     `json:"passwordCredentials,omitempty"`
	Oauth2AllowImplicitFlow *bool                     `json:"oauth2AllowImplicitFlow,omitempty"`
	RequiredResourceAccess  *[]RequiredResourceAccess `json:"requiredResourceAccess,omitempty"`
}

// CheckGroupMembershipParameters is request parameters for IsMemberOf API call.
type CheckGroupMembershipParameters struct {
	GroupID  *string `json:"groupId,omitempty"`
	MemberID *string `json:"memberId,omitempty"`
}

// CheckGroupMembershipResult is server response for IsMemberOf API call
type CheckGroupMembershipResult struct {
	autorest.Response `json:"-"`
	Value             *bool `json:"value,omitempty"`
}

// Domain is active Directory Domain information.
type Domain struct {
	autorest.Response  `json:"-"`
	AuthenticationType *string `json:"authenticationType,omitempty"`
	IsDefault          *bool   `json:"isDefault,omitempty"`
	IsVerified         *bool   `json:"isVerified,omitempty"`
	Name               *string `json:"name,omitempty"`
}

// DomainListResult is server response for Get tenant domains API call.
type DomainListResult struct {
	autorest.Response `json:"-"`
	Value             *[]Domain `json:"value,omitempty"`
}

// ErrorMessage is active Directory error message.
type ErrorMessage struct {
	Message *string `json:"value,omitempty"`
}

// GetObjectsParameters is request parameters for the GetObjectsByObjectIds API.
type GetObjectsParameters struct {
	ObjectIds                        *[]string `json:"objectIds,omitempty"`
	Types                            *[]string `json:"types,omitempty"`
	IncludeDirectoryObjectReferences *bool     `json:"includeDirectoryObjectReferences,omitempty"`
}

// GetObjectsResult is the response to an Active Directory object inquiry API request.
type GetObjectsResult struct {
	autorest.Response `json:"-"`
	Value             *[]AADObject `json:"value,omitempty"`
	OdataNextLink     *string      `json:"odata.nextLink,omitempty"`
}

// GraphError is active Directory error information.
type GraphError struct {
	*OdataError `json:"odata.error,omitempty"`
}

// GroupAddMemberParameters is request parameters for adding a member to a group.
type GroupAddMemberParameters struct {
	URL *string `json:"url,omitempty"`
}

// GroupCreateParameters is request parameters for creating a new group.
type GroupCreateParameters struct {
	DisplayName     *string `json:"displayName,omitempty"`
	MailEnabled     *bool   `json:"mailEnabled,omitempty"`
	MailNickname    *string `json:"mailNickname,omitempty"`
	SecurityEnabled *bool   `json:"securityEnabled,omitempty"`
}

// GroupGetMemberGroupsParameters is request parameters for GetMemberGroups API call.
type GroupGetMemberGroupsParameters struct {
	SecurityEnabledOnly *bool `json:"securityEnabledOnly,omitempty"`
}

// GroupGetMemberGroupsResult is server response for GetMemberGroups API call.
type GroupGetMemberGroupsResult struct {
	autorest.Response `json:"-"`
	Value             *[]string `json:"value,omitempty"`
}

// GroupListResult is server response for Get tenant groups API call
type GroupListResult struct {
	autorest.Response `json:"-"`
	Value             *[]ADGroup `json:"value,omitempty"`
	OdataNextLink     *string    `json:"odata.nextLink,omitempty"`
}

// KeyCredential is active Directory Key Credential information.
type KeyCredential struct {
	StartDate *date.Time `json:"startDate,omitempty"`
	EndDate   *date.Time `json:"endDate,omitempty"`
	Value     *string    `json:"value,omitempty"`
	KeyID     *string    `json:"keyId,omitempty"`
	Usage     *string    `json:"usage,omitempty"`
	Type      *string    `json:"type,omitempty"`
}

// KeyCredentialListResult is keyCredential list operation result.
type KeyCredentialListResult struct {
	autorest.Response `json:"-"`
	Value             *[]KeyCredential `json:"value,omitempty"`
}

// KeyCredentialsUpdateParameters is request parameters for a KeyCredentials update operation
type KeyCredentialsUpdateParameters struct {
	Value *[]KeyCredential `json:"value,omitempty"`
}

// OdataError is active Directory OData error information.
type OdataError struct {
	Code          *string `json:"code,omitempty"`
	*ErrorMessage `json:"message,omitempty"`
}

// PasswordCredential is active Directory Password Credential information.
type PasswordCredential struct {
	StartDate *date.Time `json:"startDate,omitempty"`
	EndDate   *date.Time `json:"endDate,omitempty"`
	KeyID     *string    `json:"keyId,omitempty"`
	Value     *string    `json:"value,omitempty"`
}

// PasswordCredentialListResult is passwordCredential list operation result.
type PasswordCredentialListResult struct {
	autorest.Response `json:"-"`
	Value             *[]PasswordCredential `json:"value,omitempty"`
}

// PasswordCredentialsUpdateParameters is request parameters for a PasswordCredentials update operation.
type PasswordCredentialsUpdateParameters struct {
	Value *[]PasswordCredential `json:"value,omitempty"`
}

// PasswordProfile is the password profile associated with a user.
type PasswordProfile struct {
	Password                     *string `json:"password,omitempty"`
	ForceChangePasswordNextLogin *bool   `json:"forceChangePasswordNextLogin,omitempty"`
}

// RequiredResourceAccess is specifies the set of OAuth 2.0 permission scopes and app roles under the specified
// resource that an application requires access to. The specified OAuth 2.0 permission scopes may be requested by
// client applications (through the requiredResourceAccess collection) when calling a resource application. The
// requiredResourceAccess property of the Application entity is a collection of ReqiredResourceAccess.
type RequiredResourceAccess struct {
	ResourceAccess *[]ResourceAccess `json:"resourceAccess,omitempty"`
	ResourceAppID  *string           `json:"resourceAppId,omitempty"`
}

// ResourceAccess is specifies an OAuth 2.0 permission scope or an app role that an application requires. The
// resourceAccess property of the RequiredResourceAccess type is a collection of ResourceAccess.
type ResourceAccess struct {
	ID   *string `json:"id,omitempty"`
	Type *string `json:"type,omitempty"`
}

// ServicePrincipal is active Directory service principal information.
type ServicePrincipal struct {
	autorest.Response     `json:"-"`
	ObjectID              *string   `json:"objectId,omitempty"`
	ObjectType            *string   `json:"objectType,omitempty"`
	DisplayName           *string   `json:"displayName,omitempty"`
	AppID                 *string   `json:"appId,omitempty"`
	ServicePrincipalNames *[]string `json:"servicePrincipalNames,omitempty"`
}

// ServicePrincipalCreateParameters is request parameters for creating a new service principal.
type ServicePrincipalCreateParameters struct {
	AppID               *string               `json:"appId,omitempty"`
	AccountEnabled      *bool                 `json:"accountEnabled,omitempty"`
	KeyCredentials      *[]KeyCredential      `json:"keyCredentials,omitempty"`
	PasswordCredentials *[]PasswordCredential `json:"passwordCredentials,omitempty"`
}

// ServicePrincipalListResult is server response for get tenant service principals API call.
type ServicePrincipalListResult struct {
	autorest.Response `json:"-"`
	Value             *[]ServicePrincipal `json:"value,omitempty"`
	OdataNextLink     *string             `json:"odata.nextLink,omitempty"`
}

// SignInName is contains information about a sign-in name of a local account user in an Azure Active Directory B2C
// tenant.
type SignInName struct {
	Type  *string `json:"type,omitempty"`
	Value *string `json:"value,omitempty"`
}

// User is active Directory user information.
type User struct {
	autorest.Response `json:"-"`
	ImmutableID       *string       `json:"immutableId,omitempty"`
	UsageLocation     *string       `json:"usageLocation,omitempty"`
	GivenName         *string       `json:"givenName,omitempty"`
	Surname           *string       `json:"surname,omitempty"`
	UserType          UserType      `json:"userType,omitempty"`
	AccountEnabled    *bool         `json:"accountEnabled,omitempty"`
	DisplayName       *string       `json:"displayName,omitempty"`
	UserPrincipalName *string       `json:"userPrincipalName,omitempty"`
	MailNickname      *string       `json:"mailNickname,omitempty"`
	Mail              *string       `json:"mail,omitempty"`
	ObjectID          *string       `json:"objectId,omitempty"`
	ObjectType        *string       `json:"objectType,omitempty"`
	SignInNames       *[]SignInName `json:"signInNames,omitempty"`
}

// UserBase is
type UserBase struct {
	ImmutableID   *string  `json:"immutableId,omitempty"`
	UsageLocation *string  `json:"usageLocation,omitempty"`
	GivenName     *string  `json:"givenName,omitempty"`
	Surname       *string  `json:"surname,omitempty"`
	UserType      UserType `json:"userType,omitempty"`
}

// UserCreateParameters is request parameters for creating a new work or school account user.
type UserCreateParameters struct {
	ImmutableID       *string          `json:"immutableId,omitempty"`
	UsageLocation     *string          `json:"usageLocation,omitempty"`
	GivenName         *string          `json:"givenName,omitempty"`
	Surname           *string          `json:"surname,omitempty"`
	UserType          UserType         `json:"userType,omitempty"`
	AccountEnabled    *bool            `json:"accountEnabled,omitempty"`
	DisplayName       *string          `json:"displayName,omitempty"`
	PasswordProfile   *PasswordProfile `json:"passwordProfile,omitempty"`
	UserPrincipalName *string          `json:"userPrincipalName,omitempty"`
	MailNickname      *string          `json:"mailNickname,omitempty"`
	Mail              *string          `json:"mail,omitempty"`
}

// UserGetMemberGroupsParameters is request parameters for GetMemberGroups API call.
type UserGetMemberGroupsParameters struct {
	SecurityEnabledOnly *bool `json:"securityEnabledOnly,omitempty"`
}

// UserGetMemberGroupsResult is server response for GetMemberGroups API call.
type UserGetMemberGroupsResult struct {
	autorest.Response `json:"-"`
	Value             *[]string `json:"value,omitempty"`
}

// UserListResult is server response for Get tenant users API call.
type UserListResult struct {
	autorest.Response `json:"-"`
	Value             *[]User `json:"value,omitempty"`
	OdataNextLink     *string `json:"odata.nextLink,omitempty"`
}

// UserUpdateParameters is request parameters for updating an existing work or school account user.
type UserUpdateParameters struct {
	ImmutableID       *string          `json:"immutableId,omitempty"`
	UsageLocation     *string          `json:"usageLocation,omitempty"`
	GivenName         *string          `json:"givenName,omitempty"`
	Surname           *string          `json:"surname,omitempty"`
	UserType          UserType         `json:"userType,omitempty"`
	AccountEnabled    *bool            `json:"accountEnabled,omitempty"`
	DisplayName       *string          `json:"displayName,omitempty"`
	PasswordProfile   *PasswordProfile `json:"passwordProfile,omitempty"`
	UserPrincipalName *string          `json:"userPrincipalName,omitempty"`
	MailNickname      *string          `json:"mailNickname,omitempty"`
}
