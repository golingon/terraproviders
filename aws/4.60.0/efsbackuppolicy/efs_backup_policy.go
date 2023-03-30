// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package efsbackuppolicy

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type BackupPolicy struct {
	// Status: string, required
	Status terra.StringValue `hcl:"status,attr" validate:"required"`
}

type BackupPolicyAttributes struct {
	ref terra.Reference
}

func (bp BackupPolicyAttributes) InternalRef() terra.Reference {
	return bp.ref
}

func (bp BackupPolicyAttributes) InternalWithRef(ref terra.Reference) BackupPolicyAttributes {
	return BackupPolicyAttributes{ref: ref}
}

func (bp BackupPolicyAttributes) InternalTokens() hclwrite.Tokens {
	return bp.ref.InternalTokens()
}

func (bp BackupPolicyAttributes) Status() terra.StringValue {
	return terra.ReferenceString(bp.ref.Append("status"))
}

type BackupPolicyState struct {
	Status string `json:"status"`
}
