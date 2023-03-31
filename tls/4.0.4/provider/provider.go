// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package provider

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Proxy struct {
	// FromEnv: bool, optional
	FromEnv terra.BoolValue `hcl:"from_env,attr"`
	// Password: string, optional
	Password terra.StringValue `hcl:"password,attr"`
	// Url: string, optional
	Url terra.StringValue `hcl:"url,attr"`
	// Username: string, optional
	Username terra.StringValue `hcl:"username,attr"`
}

type ProxyAttributes struct {
	ref terra.Reference
}

func (p ProxyAttributes) InternalRef() terra.Reference {
	return p.ref
}

func (p ProxyAttributes) InternalWithRef(ref terra.Reference) ProxyAttributes {
	return ProxyAttributes{ref: ref}
}

func (p ProxyAttributes) InternalTokens() hclwrite.Tokens {
	return p.ref.InternalTokens()
}

func (p ProxyAttributes) FromEnv() terra.BoolValue {
	return terra.ReferenceAsBool(p.ref.Append("from_env"))
}

func (p ProxyAttributes) Password() terra.StringValue {
	return terra.ReferenceAsString(p.ref.Append("password"))
}

func (p ProxyAttributes) Url() terra.StringValue {
	return terra.ReferenceAsString(p.ref.Append("url"))
}

func (p ProxyAttributes) Username() terra.StringValue {
	return terra.ReferenceAsString(p.ref.Append("username"))
}

type ProxyState struct {
	FromEnv  bool   `json:"from_env"`
	Password string `json:"password"`
	Url      string `json:"url"`
	Username string `json:"username"`
}
