// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package opsworksstack

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type CustomCookbooksSource struct {
	// Password: string, optional
	Password terra.StringValue `hcl:"password,attr"`
	// Revision: string, optional
	Revision terra.StringValue `hcl:"revision,attr"`
	// SshKey: string, optional
	SshKey terra.StringValue `hcl:"ssh_key,attr"`
	// Type: string, required
	Type terra.StringValue `hcl:"type,attr" validate:"required"`
	// Url: string, required
	Url terra.StringValue `hcl:"url,attr" validate:"required"`
	// Username: string, optional
	Username terra.StringValue `hcl:"username,attr"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
}

type CustomCookbooksSourceAttributes struct {
	ref terra.Reference
}

func (ccs CustomCookbooksSourceAttributes) InternalRef() (terra.Reference, error) {
	return ccs.ref, nil
}

func (ccs CustomCookbooksSourceAttributes) InternalWithRef(ref terra.Reference) CustomCookbooksSourceAttributes {
	return CustomCookbooksSourceAttributes{ref: ref}
}

func (ccs CustomCookbooksSourceAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ccs.ref.InternalTokens()
}

func (ccs CustomCookbooksSourceAttributes) Password() terra.StringValue {
	return terra.ReferenceAsString(ccs.ref.Append("password"))
}

func (ccs CustomCookbooksSourceAttributes) Revision() terra.StringValue {
	return terra.ReferenceAsString(ccs.ref.Append("revision"))
}

func (ccs CustomCookbooksSourceAttributes) SshKey() terra.StringValue {
	return terra.ReferenceAsString(ccs.ref.Append("ssh_key"))
}

func (ccs CustomCookbooksSourceAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(ccs.ref.Append("type"))
}

func (ccs CustomCookbooksSourceAttributes) Url() terra.StringValue {
	return terra.ReferenceAsString(ccs.ref.Append("url"))
}

func (ccs CustomCookbooksSourceAttributes) Username() terra.StringValue {
	return terra.ReferenceAsString(ccs.ref.Append("username"))
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

type CustomCookbooksSourceState struct {
	Password string `json:"password"`
	Revision string `json:"revision"`
	SshKey   string `json:"ssh_key"`
	Type     string `json:"type"`
	Url      string `json:"url"`
	Username string `json:"username"`
}

type TimeoutsState struct {
	Create string `json:"create"`
}
