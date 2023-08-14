// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package appserviceenvironment

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type ClusterSetting struct {
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Value: string, required
	Value terra.StringValue `hcl:"value,attr" validate:"required"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Read: string, optional
	Read terra.StringValue `hcl:"read,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type ClusterSettingAttributes struct {
	ref terra.Reference
}

func (cs ClusterSettingAttributes) InternalRef() (terra.Reference, error) {
	return cs.ref, nil
}

func (cs ClusterSettingAttributes) InternalWithRef(ref terra.Reference) ClusterSettingAttributes {
	return ClusterSettingAttributes{ref: ref}
}

func (cs ClusterSettingAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return cs.ref.InternalTokens()
}

func (cs ClusterSettingAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cs.ref.Append("name"))
}

func (cs ClusterSettingAttributes) Value() terra.StringValue {
	return terra.ReferenceAsString(cs.ref.Append("value"))
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

func (t TimeoutsAttributes) Read() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("read"))
}

func (t TimeoutsAttributes) Update() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("update"))
}

type ClusterSettingState struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
}