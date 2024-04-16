// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_securityhub_organization_configuration

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type OrganizationConfiguration struct {
	// ConfigurationType: string, required
	ConfigurationType terra.StringValue `hcl:"configuration_type,attr" validate:"required"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type OrganizationConfigurationAttributes struct {
	ref terra.Reference
}

func (oc OrganizationConfigurationAttributes) InternalRef() (terra.Reference, error) {
	return oc.ref, nil
}

func (oc OrganizationConfigurationAttributes) InternalWithRef(ref terra.Reference) OrganizationConfigurationAttributes {
	return OrganizationConfigurationAttributes{ref: ref}
}

func (oc OrganizationConfigurationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return oc.ref.InternalTokens()
}

func (oc OrganizationConfigurationAttributes) ConfigurationType() terra.StringValue {
	return terra.ReferenceAsString(oc.ref.Append("configuration_type"))
}

type TimeoutsAttributes struct {
	ref terra.Reference
}

func (t TimeoutsAttributes) InternalRef() (terra.Reference, error) {
	return t.ref, nil
}

func (t TimeoutsAttributes) InternalWithRef(ref terra.Reference) TimeoutsAttributes {
	return TimeoutsAttributes{ref: ref}
}

func (t TimeoutsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return t.ref.InternalTokens()
}

func (t TimeoutsAttributes) Create() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("create"))
}

func (t TimeoutsAttributes) Delete() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("delete"))
}

func (t TimeoutsAttributes) Update() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("update"))
}

type OrganizationConfigurationState struct {
	ConfigurationType string `json:"configuration_type"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}
