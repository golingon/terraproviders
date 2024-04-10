// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package provider

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type Update struct {
	// KeyAlgorithm: string, optional
	KeyAlgorithm terra.StringValue `hcl:"key_algorithm,attr"`
	// KeyName: string, optional
	KeyName terra.StringValue `hcl:"key_name,attr"`
	// KeySecret: string, optional
	KeySecret terra.StringValue `hcl:"key_secret,attr"`
	// Port: number, optional
	Port terra.NumberValue `hcl:"port,attr"`
	// Retries: number, optional
	Retries terra.NumberValue `hcl:"retries,attr"`
	// Server: string, optional
	Server terra.StringValue `hcl:"server,attr"`
	// Timeout: string, optional
	Timeout terra.StringValue `hcl:"timeout,attr"`
	// Transport: string, optional
	Transport terra.StringValue `hcl:"transport,attr"`
	// Gssapi: min=0
	Gssapi []Gssapi `hcl:"gssapi,block" validate:"min=0"`
}

type Gssapi struct {
	// Keytab: string, optional
	Keytab terra.StringValue `hcl:"keytab,attr"`
	// Password: string, optional
	Password terra.StringValue `hcl:"password,attr"`
	// Realm: string, optional
	Realm terra.StringValue `hcl:"realm,attr"`
	// Username: string, optional
	Username terra.StringValue `hcl:"username,attr"`
}

type UpdateAttributes struct {
	ref terra.Reference
}

func (u UpdateAttributes) InternalRef() (terra.Reference, error) {
	return u.ref, nil
}

func (u UpdateAttributes) InternalWithRef(ref terra.Reference) UpdateAttributes {
	return UpdateAttributes{ref: ref}
}

func (u UpdateAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return u.ref.InternalTokens()
}

func (u UpdateAttributes) KeyAlgorithm() terra.StringValue {
	return terra.ReferenceAsString(u.ref.Append("key_algorithm"))
}

func (u UpdateAttributes) KeyName() terra.StringValue {
	return terra.ReferenceAsString(u.ref.Append("key_name"))
}

func (u UpdateAttributes) KeySecret() terra.StringValue {
	return terra.ReferenceAsString(u.ref.Append("key_secret"))
}

func (u UpdateAttributes) Port() terra.NumberValue {
	return terra.ReferenceAsNumber(u.ref.Append("port"))
}

func (u UpdateAttributes) Retries() terra.NumberValue {
	return terra.ReferenceAsNumber(u.ref.Append("retries"))
}

func (u UpdateAttributes) Server() terra.StringValue {
	return terra.ReferenceAsString(u.ref.Append("server"))
}

func (u UpdateAttributes) Timeout() terra.StringValue {
	return terra.ReferenceAsString(u.ref.Append("timeout"))
}

func (u UpdateAttributes) Transport() terra.StringValue {
	return terra.ReferenceAsString(u.ref.Append("transport"))
}

func (u UpdateAttributes) Gssapi() terra.ListValue[GssapiAttributes] {
	return terra.ReferenceAsList[GssapiAttributes](u.ref.Append("gssapi"))
}

type GssapiAttributes struct {
	ref terra.Reference
}

func (g GssapiAttributes) InternalRef() (terra.Reference, error) {
	return g.ref, nil
}

func (g GssapiAttributes) InternalWithRef(ref terra.Reference) GssapiAttributes {
	return GssapiAttributes{ref: ref}
}

func (g GssapiAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return g.ref.InternalTokens()
}

func (g GssapiAttributes) Keytab() terra.StringValue {
	return terra.ReferenceAsString(g.ref.Append("keytab"))
}

func (g GssapiAttributes) Password() terra.StringValue {
	return terra.ReferenceAsString(g.ref.Append("password"))
}

func (g GssapiAttributes) Realm() terra.StringValue {
	return terra.ReferenceAsString(g.ref.Append("realm"))
}

func (g GssapiAttributes) Username() terra.StringValue {
	return terra.ReferenceAsString(g.ref.Append("username"))
}

type UpdateState struct {
	KeyAlgorithm string        `json:"key_algorithm"`
	KeyName      string        `json:"key_name"`
	KeySecret    string        `json:"key_secret"`
	Port         float64       `json:"port"`
	Retries      float64       `json:"retries"`
	Server       string        `json:"server"`
	Timeout      string        `json:"timeout"`
	Transport    string        `json:"transport"`
	Gssapi       []GssapiState `json:"gssapi"`
}

type GssapiState struct {
	Keytab   string `json:"keytab"`
	Password string `json:"password"`
	Realm    string `json:"realm"`
	Username string `json:"username"`
}
