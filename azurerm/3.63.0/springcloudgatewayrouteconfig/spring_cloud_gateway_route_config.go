// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package springcloudgatewayrouteconfig

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type OpenApi struct {
	// Uri: string, optional
	Uri terra.StringValue `hcl:"uri,attr"`
}

type Route struct {
	// ClassificationTags: set of string, optional
	ClassificationTags terra.SetValue[terra.StringValue] `hcl:"classification_tags,attr"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Filters: set of string, optional
	Filters terra.SetValue[terra.StringValue] `hcl:"filters,attr"`
	// Order: number, required
	Order terra.NumberValue `hcl:"order,attr" validate:"required"`
	// Predicates: set of string, optional
	Predicates terra.SetValue[terra.StringValue] `hcl:"predicates,attr"`
	// SsoValidationEnabled: bool, optional
	SsoValidationEnabled terra.BoolValue `hcl:"sso_validation_enabled,attr"`
	// Title: string, optional
	Title terra.StringValue `hcl:"title,attr"`
	// TokenRelay: bool, optional
	TokenRelay terra.BoolValue `hcl:"token_relay,attr"`
	// Uri: string, optional
	Uri terra.StringValue `hcl:"uri,attr"`
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

type OpenApiAttributes struct {
	ref terra.Reference
}

func (oa OpenApiAttributes) InternalRef() (terra.Reference, error) {
	return oa.ref, nil
}

func (oa OpenApiAttributes) InternalWithRef(ref terra.Reference) OpenApiAttributes {
	return OpenApiAttributes{ref: ref}
}

func (oa OpenApiAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return oa.ref.InternalTokens()
}

func (oa OpenApiAttributes) Uri() terra.StringValue {
	return terra.ReferenceAsString(oa.ref.Append("uri"))
}

type RouteAttributes struct {
	ref terra.Reference
}

func (r RouteAttributes) InternalRef() (terra.Reference, error) {
	return r.ref, nil
}

func (r RouteAttributes) InternalWithRef(ref terra.Reference) RouteAttributes {
	return RouteAttributes{ref: ref}
}

func (r RouteAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return r.ref.InternalTokens()
}

func (r RouteAttributes) ClassificationTags() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](r.ref.Append("classification_tags"))
}

func (r RouteAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("description"))
}

func (r RouteAttributes) Filters() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](r.ref.Append("filters"))
}

func (r RouteAttributes) Order() terra.NumberValue {
	return terra.ReferenceAsNumber(r.ref.Append("order"))
}

func (r RouteAttributes) Predicates() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](r.ref.Append("predicates"))
}

func (r RouteAttributes) SsoValidationEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(r.ref.Append("sso_validation_enabled"))
}

func (r RouteAttributes) Title() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("title"))
}

func (r RouteAttributes) TokenRelay() terra.BoolValue {
	return terra.ReferenceAsBool(r.ref.Append("token_relay"))
}

func (r RouteAttributes) Uri() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("uri"))
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

type OpenApiState struct {
	Uri string `json:"uri"`
}

type RouteState struct {
	ClassificationTags   []string `json:"classification_tags"`
	Description          string   `json:"description"`
	Filters              []string `json:"filters"`
	Order                float64  `json:"order"`
	Predicates           []string `json:"predicates"`
	SsoValidationEnabled bool     `json:"sso_validation_enabled"`
	Title                string   `json:"title"`
	TokenRelay           bool     `json:"token_relay"`
	Uri                  string   `json:"uri"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
}
