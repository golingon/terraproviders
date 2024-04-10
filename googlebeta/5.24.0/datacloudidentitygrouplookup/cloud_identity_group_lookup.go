// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package datacloudidentitygrouplookup

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type GroupKey struct {
	// Id: string, required
	Id terra.StringValue `hcl:"id,attr" validate:"required"`
	// Namespace: string, optional
	Namespace terra.StringValue `hcl:"namespace,attr"`
}

type GroupKeyAttributes struct {
	ref terra.Reference
}

func (gk GroupKeyAttributes) InternalRef() (terra.Reference, error) {
	return gk.ref, nil
}

func (gk GroupKeyAttributes) InternalWithRef(ref terra.Reference) GroupKeyAttributes {
	return GroupKeyAttributes{ref: ref}
}

func (gk GroupKeyAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return gk.ref.InternalTokens()
}

func (gk GroupKeyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gk.ref.Append("id"))
}

func (gk GroupKeyAttributes) Namespace() terra.StringValue {
	return terra.ReferenceAsString(gk.ref.Append("namespace"))
}

type GroupKeyState struct {
	Id        string `json:"id"`
	Namespace string `json:"namespace"`
}
