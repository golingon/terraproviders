// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package databatchjobqueue

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type ComputeEnvironmentOrder struct{}

type ComputeEnvironmentOrderAttributes struct {
	ref terra.Reference
}

func (ceo ComputeEnvironmentOrderAttributes) InternalRef() terra.Reference {
	return ceo.ref
}

func (ceo ComputeEnvironmentOrderAttributes) InternalWithRef(ref terra.Reference) ComputeEnvironmentOrderAttributes {
	return ComputeEnvironmentOrderAttributes{ref: ref}
}

func (ceo ComputeEnvironmentOrderAttributes) InternalTokens() hclwrite.Tokens {
	return ceo.ref.InternalTokens()
}

func (ceo ComputeEnvironmentOrderAttributes) ComputeEnvironment() terra.StringValue {
	return terra.ReferenceAsString(ceo.ref.Append("compute_environment"))
}

func (ceo ComputeEnvironmentOrderAttributes) Order() terra.NumberValue {
	return terra.ReferenceAsNumber(ceo.ref.Append("order"))
}

type ComputeEnvironmentOrderState struct {
	ComputeEnvironment string  `json:"compute_environment"`
	Order              float64 `json:"order"`
}