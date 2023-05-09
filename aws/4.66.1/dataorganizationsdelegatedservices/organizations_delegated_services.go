// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package dataorganizationsdelegatedservices

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type DelegatedServices struct{}

type DelegatedServicesAttributes struct {
	ref terra.Reference
}

func (ds DelegatedServicesAttributes) InternalRef() (terra.Reference, error) {
	return ds.ref, nil
}

func (ds DelegatedServicesAttributes) InternalWithRef(ref terra.Reference) DelegatedServicesAttributes {
	return DelegatedServicesAttributes{ref: ref}
}

func (ds DelegatedServicesAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ds.ref.InternalTokens()
}

func (ds DelegatedServicesAttributes) DelegationEnabledDate() terra.StringValue {
	return terra.ReferenceAsString(ds.ref.Append("delegation_enabled_date"))
}

func (ds DelegatedServicesAttributes) ServicePrincipal() terra.StringValue {
	return terra.ReferenceAsString(ds.ref.Append("service_principal"))
}

type DelegatedServicesState struct {
	DelegationEnabledDate string `json:"delegation_enabled_date"`
	ServicePrincipal      string `json:"service_principal"`
}
