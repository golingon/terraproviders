// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package dialogflowfulfillment

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Features struct {
	// Type: string, required
	Type terra.StringValue `hcl:"type,attr" validate:"required"`
}

type GenericWebService struct {
	// Password: string, optional
	Password terra.StringValue `hcl:"password,attr"`
	// RequestHeaders: map of string, optional
	RequestHeaders terra.MapValue[terra.StringValue] `hcl:"request_headers,attr"`
	// Uri: string, required
	Uri terra.StringValue `hcl:"uri,attr" validate:"required"`
	// Username: string, optional
	Username terra.StringValue `hcl:"username,attr"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type FeaturesAttributes struct {
	ref terra.Reference
}

func (f FeaturesAttributes) InternalRef() (terra.Reference, error) {
	return f.ref, nil
}

func (f FeaturesAttributes) InternalWithRef(ref terra.Reference) FeaturesAttributes {
	return FeaturesAttributes{ref: ref}
}

func (f FeaturesAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return f.ref.InternalTokens()
}

func (f FeaturesAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(f.ref.Append("type"))
}

type GenericWebServiceAttributes struct {
	ref terra.Reference
}

func (gws GenericWebServiceAttributes) InternalRef() (terra.Reference, error) {
	return gws.ref, nil
}

func (gws GenericWebServiceAttributes) InternalWithRef(ref terra.Reference) GenericWebServiceAttributes {
	return GenericWebServiceAttributes{ref: ref}
}

func (gws GenericWebServiceAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return gws.ref.InternalTokens()
}

func (gws GenericWebServiceAttributes) Password() terra.StringValue {
	return terra.ReferenceAsString(gws.ref.Append("password"))
}

func (gws GenericWebServiceAttributes) RequestHeaders() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](gws.ref.Append("request_headers"))
}

func (gws GenericWebServiceAttributes) Uri() terra.StringValue {
	return terra.ReferenceAsString(gws.ref.Append("uri"))
}

func (gws GenericWebServiceAttributes) Username() terra.StringValue {
	return terra.ReferenceAsString(gws.ref.Append("username"))
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

func (t TimeoutsAttributes) Update() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("update"))
}

type FeaturesState struct {
	Type string `json:"type"`
}

type GenericWebServiceState struct {
	Password       string            `json:"password"`
	RequestHeaders map[string]string `json:"request_headers"`
	Uri            string            `json:"uri"`
	Username       string            `json:"username"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}
