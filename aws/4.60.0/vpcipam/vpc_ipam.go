// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package vpcipam

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type OperatingRegions struct {
	// RegionName: string, required
	RegionName terra.StringValue `hcl:"region_name,attr" validate:"required"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type OperatingRegionsAttributes struct {
	ref terra.Reference
}

func (or OperatingRegionsAttributes) InternalRef() (terra.Reference, error) {
	return or.ref, nil
}

func (or OperatingRegionsAttributes) InternalWithRef(ref terra.Reference) OperatingRegionsAttributes {
	return OperatingRegionsAttributes{ref: ref}
}

func (or OperatingRegionsAttributes) InternalTokens() hclwrite.Tokens {
	return or.ref.InternalTokens()
}

func (or OperatingRegionsAttributes) RegionName() terra.StringValue {
	return terra.ReferenceAsString(or.ref.Append("region_name"))
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

func (t TimeoutsAttributes) InternalTokens() hclwrite.Tokens {
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

type OperatingRegionsState struct {
	RegionName string `json:"region_name"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}
