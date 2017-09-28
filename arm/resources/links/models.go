package links

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

// Filter enumerates the values for filter.
type Filter string

const (
	// AtScope specifies the at scope state for filter.
	AtScope Filter = "atScope()"
)

// ResourceLink is the resource link.
type ResourceLink struct {
	autorest.Response `json:"-"`
	ID                *string                 `json:"id,omitempty"`
	Name              *string                 `json:"name,omitempty"`
	Properties        *ResourceLinkProperties `json:"properties,omitempty"`
}

// ResourceLinkFilter is resource link filter.
type ResourceLinkFilter struct {
	TargetID *string `json:"targetId,omitempty"`
}

// ResourceLinkProperties is the resource link properties.
type ResourceLinkProperties struct {
	SourceID *string `json:"sourceId,omitempty"`
	TargetID *string `json:"targetId,omitempty"`
	Notes    *string `json:"notes,omitempty"`
}

// ResourceLinkResult is list of resource links.
type ResourceLinkResult struct {
	autorest.Response `json:"-"`
	Value             *[]ResourceLink `json:"value,omitempty"`
	NextLink          *string         `json:"nextLink,omitempty"`
}

// ResourceLinkResultPreparer prepares a request to retrieve the next set of results. It returns
// nil if no more results exist.
func (client ResourceLinkResult) ResourceLinkResultPreparer() (*http.Request, error) {
	if client.NextLink == nil || len(to.String(client.NextLink)) <= 0 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(client.NextLink)))
}
