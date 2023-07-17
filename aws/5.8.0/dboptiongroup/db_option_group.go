// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package dboptiongroup

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Option struct {
	// DbSecurityGroupMemberships: set of string, optional
	DbSecurityGroupMemberships terra.SetValue[terra.StringValue] `hcl:"db_security_group_memberships,attr"`
	// OptionName: string, required
	OptionName terra.StringValue `hcl:"option_name,attr" validate:"required"`
	// Port: number, optional
	Port terra.NumberValue `hcl:"port,attr"`
	// Version: string, optional
	Version terra.StringValue `hcl:"version,attr"`
	// VpcSecurityGroupMemberships: set of string, optional
	VpcSecurityGroupMemberships terra.SetValue[terra.StringValue] `hcl:"vpc_security_group_memberships,attr"`
	// OptionSettings: min=0
	OptionSettings []OptionSettings `hcl:"option_settings,block" validate:"min=0"`
}

type OptionSettings struct {
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Value: string, required
	Value terra.StringValue `hcl:"value,attr" validate:"required"`
}

type Timeouts struct {
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
}

type OptionAttributes struct {
	ref terra.Reference
}

func (o OptionAttributes) InternalRef() (terra.Reference, error) {
	return o.ref, nil
}

func (o OptionAttributes) InternalWithRef(ref terra.Reference) OptionAttributes {
	return OptionAttributes{ref: ref}
}

func (o OptionAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return o.ref.InternalTokens()
}

func (o OptionAttributes) DbSecurityGroupMemberships() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](o.ref.Append("db_security_group_memberships"))
}

func (o OptionAttributes) OptionName() terra.StringValue {
	return terra.ReferenceAsString(o.ref.Append("option_name"))
}

func (o OptionAttributes) Port() terra.NumberValue {
	return terra.ReferenceAsNumber(o.ref.Append("port"))
}

func (o OptionAttributes) Version() terra.StringValue {
	return terra.ReferenceAsString(o.ref.Append("version"))
}

func (o OptionAttributes) VpcSecurityGroupMemberships() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](o.ref.Append("vpc_security_group_memberships"))
}

func (o OptionAttributes) OptionSettings() terra.SetValue[OptionSettingsAttributes] {
	return terra.ReferenceAsSet[OptionSettingsAttributes](o.ref.Append("option_settings"))
}

type OptionSettingsAttributes struct {
	ref terra.Reference
}

func (os OptionSettingsAttributes) InternalRef() (terra.Reference, error) {
	return os.ref, nil
}

func (os OptionSettingsAttributes) InternalWithRef(ref terra.Reference) OptionSettingsAttributes {
	return OptionSettingsAttributes{ref: ref}
}

func (os OptionSettingsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return os.ref.InternalTokens()
}

func (os OptionSettingsAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(os.ref.Append("name"))
}

func (os OptionSettingsAttributes) Value() terra.StringValue {
	return terra.ReferenceAsString(os.ref.Append("value"))
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

func (t TimeoutsAttributes) Delete() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("delete"))
}

type OptionState struct {
	DbSecurityGroupMemberships  []string              `json:"db_security_group_memberships"`
	OptionName                  string                `json:"option_name"`
	Port                        float64               `json:"port"`
	Version                     string                `json:"version"`
	VpcSecurityGroupMemberships []string              `json:"vpc_security_group_memberships"`
	OptionSettings              []OptionSettingsState `json:"option_settings"`
}

type OptionSettingsState struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type TimeoutsState struct {
	Delete string `json:"delete"`
}