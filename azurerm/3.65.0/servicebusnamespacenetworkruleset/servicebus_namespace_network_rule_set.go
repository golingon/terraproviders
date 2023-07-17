// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package servicebusnamespacenetworkruleset

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type NetworkRules struct {
	// IgnoreMissingVnetServiceEndpoint: bool, optional
	IgnoreMissingVnetServiceEndpoint terra.BoolValue `hcl:"ignore_missing_vnet_service_endpoint,attr"`
	// SubnetId: string, required
	SubnetId terra.StringValue `hcl:"subnet_id,attr" validate:"required"`
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

type NetworkRulesAttributes struct {
	ref terra.Reference
}

func (nr NetworkRulesAttributes) InternalRef() (terra.Reference, error) {
	return nr.ref, nil
}

func (nr NetworkRulesAttributes) InternalWithRef(ref terra.Reference) NetworkRulesAttributes {
	return NetworkRulesAttributes{ref: ref}
}

func (nr NetworkRulesAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return nr.ref.InternalTokens()
}

func (nr NetworkRulesAttributes) IgnoreMissingVnetServiceEndpoint() terra.BoolValue {
	return terra.ReferenceAsBool(nr.ref.Append("ignore_missing_vnet_service_endpoint"))
}

func (nr NetworkRulesAttributes) SubnetId() terra.StringValue {
	return terra.ReferenceAsString(nr.ref.Append("subnet_id"))
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

type NetworkRulesState struct {
	IgnoreMissingVnetServiceEndpoint bool   `json:"ignore_missing_vnet_service_endpoint"`
	SubnetId                         string `json:"subnet_id"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
}
