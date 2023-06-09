// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package s3bucketownershipcontrols

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Rule struct {
	// ObjectOwnership: string, required
	ObjectOwnership terra.StringValue `hcl:"object_ownership,attr" validate:"required"`
}

type RuleAttributes struct {
	ref terra.Reference
}

func (r RuleAttributes) InternalRef() (terra.Reference, error) {
	return r.ref, nil
}

func (r RuleAttributes) InternalWithRef(ref terra.Reference) RuleAttributes {
	return RuleAttributes{ref: ref}
}

func (r RuleAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return r.ref.InternalTokens()
}

func (r RuleAttributes) ObjectOwnership() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("object_ownership"))
}

type RuleState struct {
	ObjectOwnership string `json:"object_ownership"`
}
