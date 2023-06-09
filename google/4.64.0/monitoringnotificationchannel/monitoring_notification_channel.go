// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package monitoringnotificationchannel

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type SensitiveLabels struct {
	// AuthToken: string, optional
	AuthToken terra.StringValue `hcl:"auth_token,attr"`
	// Password: string, optional
	Password terra.StringValue `hcl:"password,attr"`
	// ServiceKey: string, optional
	ServiceKey terra.StringValue `hcl:"service_key,attr"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type SensitiveLabelsAttributes struct {
	ref terra.Reference
}

func (sl SensitiveLabelsAttributes) InternalRef() (terra.Reference, error) {
	return sl.ref, nil
}

func (sl SensitiveLabelsAttributes) InternalWithRef(ref terra.Reference) SensitiveLabelsAttributes {
	return SensitiveLabelsAttributes{ref: ref}
}

func (sl SensitiveLabelsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return sl.ref.InternalTokens()
}

func (sl SensitiveLabelsAttributes) AuthToken() terra.StringValue {
	return terra.ReferenceAsString(sl.ref.Append("auth_token"))
}

func (sl SensitiveLabelsAttributes) Password() terra.StringValue {
	return terra.ReferenceAsString(sl.ref.Append("password"))
}

func (sl SensitiveLabelsAttributes) ServiceKey() terra.StringValue {
	return terra.ReferenceAsString(sl.ref.Append("service_key"))
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

type SensitiveLabelsState struct {
	AuthToken  string `json:"auth_token"`
	Password   string `json:"password"`
	ServiceKey string `json:"service_key"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}
