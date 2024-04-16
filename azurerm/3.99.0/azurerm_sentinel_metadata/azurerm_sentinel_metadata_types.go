// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_sentinel_metadata

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type Author struct {
	// Email: string, optional
	Email terra.StringValue `hcl:"email,attr"`
	// Link: string, optional
	Link terra.StringValue `hcl:"link,attr"`
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
}

type Category struct {
	// Domains: list of string, optional
	Domains terra.ListValue[terra.StringValue] `hcl:"domains,attr"`
	// Verticals: list of string, optional
	Verticals terra.ListValue[terra.StringValue] `hcl:"verticals,attr"`
}

type Source struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Kind: string, required
	Kind terra.StringValue `hcl:"kind,attr" validate:"required"`
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
}

type Support struct {
	// Email: string, optional
	Email terra.StringValue `hcl:"email,attr"`
	// Link: string, optional
	Link terra.StringValue `hcl:"link,attr"`
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
	// Tier: string, required
	Tier terra.StringValue `hcl:"tier,attr" validate:"required"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Read: string, optional
	Read terra.StringValue `hcl:"read,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type AuthorAttributes struct {
	ref terra.Reference
}

func (a AuthorAttributes) InternalRef() (terra.Reference, error) {
	return a.ref, nil
}

func (a AuthorAttributes) InternalWithRef(ref terra.Reference) AuthorAttributes {
	return AuthorAttributes{ref: ref}
}

func (a AuthorAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return a.ref.InternalTokens()
}

func (a AuthorAttributes) Email() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("email"))
}

func (a AuthorAttributes) Link() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("link"))
}

func (a AuthorAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("name"))
}

type CategoryAttributes struct {
	ref terra.Reference
}

func (c CategoryAttributes) InternalRef() (terra.Reference, error) {
	return c.ref, nil
}

func (c CategoryAttributes) InternalWithRef(ref terra.Reference) CategoryAttributes {
	return CategoryAttributes{ref: ref}
}

func (c CategoryAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return c.ref.InternalTokens()
}

func (c CategoryAttributes) Domains() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](c.ref.Append("domains"))
}

func (c CategoryAttributes) Verticals() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](c.ref.Append("verticals"))
}

type SourceAttributes struct {
	ref terra.Reference
}

func (s SourceAttributes) InternalRef() (terra.Reference, error) {
	return s.ref, nil
}

func (s SourceAttributes) InternalWithRef(ref terra.Reference) SourceAttributes {
	return SourceAttributes{ref: ref}
}

func (s SourceAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return s.ref.InternalTokens()
}

func (s SourceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("id"))
}

func (s SourceAttributes) Kind() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("kind"))
}

func (s SourceAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("name"))
}

type SupportAttributes struct {
	ref terra.Reference
}

func (s SupportAttributes) InternalRef() (terra.Reference, error) {
	return s.ref, nil
}

func (s SupportAttributes) InternalWithRef(ref terra.Reference) SupportAttributes {
	return SupportAttributes{ref: ref}
}

func (s SupportAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return s.ref.InternalTokens()
}

func (s SupportAttributes) Email() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("email"))
}

func (s SupportAttributes) Link() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("link"))
}

func (s SupportAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("name"))
}

func (s SupportAttributes) Tier() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("tier"))
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

func (t TimeoutsAttributes) Read() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("read"))
}

func (t TimeoutsAttributes) Update() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("update"))
}

type AuthorState struct {
	Email string `json:"email"`
	Link  string `json:"link"`
	Name  string `json:"name"`
}

type CategoryState struct {
	Domains   []string `json:"domains"`
	Verticals []string `json:"verticals"`
}

type SourceState struct {
	Id   string `json:"id"`
	Kind string `json:"kind"`
	Name string `json:"name"`
}

type SupportState struct {
	Email string `json:"email"`
	Link  string `json:"link"`
	Name  string `json:"name"`
	Tier  string `json:"tier"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
}
