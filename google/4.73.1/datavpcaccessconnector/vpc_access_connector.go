// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package datavpcaccessconnector

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Subnet struct{}

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

type SubnetState struct {
	Name      string `json:"name"`
	ProjectId string `json:"project_id"`
}