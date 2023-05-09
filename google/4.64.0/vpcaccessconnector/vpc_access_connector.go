// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package vpcaccessconnector

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Subnet struct {
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
	// ProjectId: string, optional
	ProjectId terra.StringValue `hcl:"project_id,attr"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
}

type SubnetAttributes struct {
	ref terra.Reference
}

func (s SubnetAttributes) InternalRef() (terra.Reference, error) {
	return s.ref, nil
}

func (s SubnetAttributes) InternalWithRef(ref terra.Reference) SubnetAttributes {
	return SubnetAttributes{ref: ref}
}

func (s SubnetAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return s.ref.InternalTokens()
}

func (s SubnetAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("name"))
}

func (s SubnetAttributes) ProjectId() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("project_id"))
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

type SubnetState struct {
	Name      string `json:"name"`
	ProjectId string `json:"project_id"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
}
