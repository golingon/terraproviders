// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_endpoints_service

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type ApisAttributes struct {
	ref terra.Reference
}

func (a ApisAttributes) InternalRef() (terra.Reference, error) {
	return a.ref, nil
}

func (a ApisAttributes) InternalWithRef(ref terra.Reference) ApisAttributes {
	return ApisAttributes{ref: ref}
}

func (a ApisAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return a.ref.InternalTokens()
}

func (a ApisAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("name"))
}

func (a ApisAttributes) Syntax() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("syntax"))
}

func (a ApisAttributes) Version() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("version"))
}

func (a ApisAttributes) Methods() terra.ListValue[ApisMethodsAttributes] {
	return terra.ReferenceAsList[ApisMethodsAttributes](a.ref.Append("methods"))
}

type ApisMethodsAttributes struct {
	ref terra.Reference
}

func (m ApisMethodsAttributes) InternalRef() (terra.Reference, error) {
	return m.ref, nil
}

func (m ApisMethodsAttributes) InternalWithRef(ref terra.Reference) ApisMethodsAttributes {
	return ApisMethodsAttributes{ref: ref}
}

func (m ApisMethodsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return m.ref.InternalTokens()
}

func (m ApisMethodsAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(m.ref.Append("name"))
}

func (m ApisMethodsAttributes) RequestType() terra.StringValue {
	return terra.ReferenceAsString(m.ref.Append("request_type"))
}

func (m ApisMethodsAttributes) ResponseType() terra.StringValue {
	return terra.ReferenceAsString(m.ref.Append("response_type"))
}

func (m ApisMethodsAttributes) Syntax() terra.StringValue {
	return terra.ReferenceAsString(m.ref.Append("syntax"))
}

type EndpointsAttributes struct {
	ref terra.Reference
}

func (e EndpointsAttributes) InternalRef() (terra.Reference, error) {
	return e.ref, nil
}

func (e EndpointsAttributes) InternalWithRef(ref terra.Reference) EndpointsAttributes {
	return EndpointsAttributes{ref: ref}
}

func (e EndpointsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return e.ref.InternalTokens()
}

func (e EndpointsAttributes) Address() terra.StringValue {
	return terra.ReferenceAsString(e.ref.Append("address"))
}

func (e EndpointsAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(e.ref.Append("name"))
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

type ApisState struct {
	Name    string             `json:"name"`
	Syntax  string             `json:"syntax"`
	Version string             `json:"version"`
	Methods []ApisMethodsState `json:"methods"`
}

type ApisMethodsState struct {
	Name         string `json:"name"`
	RequestType  string `json:"request_type"`
	ResponseType string `json:"response_type"`
	Syntax       string `json:"syntax"`
}

type EndpointsState struct {
	Address string `json:"address"`
	Name    string `json:"name"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}
