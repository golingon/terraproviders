// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_api_management_diagnostic

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type BackendRequest struct {
	// BodyBytes: number, optional
	BodyBytes terra.NumberValue `hcl:"body_bytes,attr"`
	// HeadersToLog: set of string, optional
	HeadersToLog terra.SetValue[terra.StringValue] `hcl:"headers_to_log,attr"`
	// BackendRequestDataMasking: optional
	DataMasking *BackendRequestDataMasking `hcl:"data_masking,block"`
}

type BackendRequestDataMasking struct {
	// BackendRequestDataMaskingHeaders: min=0
	Headers []BackendRequestDataMaskingHeaders `hcl:"headers,block" validate:"min=0"`
	// BackendRequestDataMaskingQueryParams: min=0
	QueryParams []BackendRequestDataMaskingQueryParams `hcl:"query_params,block" validate:"min=0"`
}

type BackendRequestDataMaskingHeaders struct {
	// Mode: string, required
	Mode terra.StringValue `hcl:"mode,attr" validate:"required"`
	// Value: string, required
	Value terra.StringValue `hcl:"value,attr" validate:"required"`
}

type BackendRequestDataMaskingQueryParams struct {
	// Mode: string, required
	Mode terra.StringValue `hcl:"mode,attr" validate:"required"`
	// Value: string, required
	Value terra.StringValue `hcl:"value,attr" validate:"required"`
}

type BackendResponse struct {
	// BodyBytes: number, optional
	BodyBytes terra.NumberValue `hcl:"body_bytes,attr"`
	// HeadersToLog: set of string, optional
	HeadersToLog terra.SetValue[terra.StringValue] `hcl:"headers_to_log,attr"`
	// BackendResponseDataMasking: optional
	DataMasking *BackendResponseDataMasking `hcl:"data_masking,block"`
}

type BackendResponseDataMasking struct {
	// BackendResponseDataMaskingHeaders: min=0
	Headers []BackendResponseDataMaskingHeaders `hcl:"headers,block" validate:"min=0"`
	// BackendResponseDataMaskingQueryParams: min=0
	QueryParams []BackendResponseDataMaskingQueryParams `hcl:"query_params,block" validate:"min=0"`
}

type BackendResponseDataMaskingHeaders struct {
	// Mode: string, required
	Mode terra.StringValue `hcl:"mode,attr" validate:"required"`
	// Value: string, required
	Value terra.StringValue `hcl:"value,attr" validate:"required"`
}

type BackendResponseDataMaskingQueryParams struct {
	// Mode: string, required
	Mode terra.StringValue `hcl:"mode,attr" validate:"required"`
	// Value: string, required
	Value terra.StringValue `hcl:"value,attr" validate:"required"`
}

type FrontendRequest struct {
	// BodyBytes: number, optional
	BodyBytes terra.NumberValue `hcl:"body_bytes,attr"`
	// HeadersToLog: set of string, optional
	HeadersToLog terra.SetValue[terra.StringValue] `hcl:"headers_to_log,attr"`
	// FrontendRequestDataMasking: optional
	DataMasking *FrontendRequestDataMasking `hcl:"data_masking,block"`
}

type FrontendRequestDataMasking struct {
	// FrontendRequestDataMaskingHeaders: min=0
	Headers []FrontendRequestDataMaskingHeaders `hcl:"headers,block" validate:"min=0"`
	// FrontendRequestDataMaskingQueryParams: min=0
	QueryParams []FrontendRequestDataMaskingQueryParams `hcl:"query_params,block" validate:"min=0"`
}

type FrontendRequestDataMaskingHeaders struct {
	// Mode: string, required
	Mode terra.StringValue `hcl:"mode,attr" validate:"required"`
	// Value: string, required
	Value terra.StringValue `hcl:"value,attr" validate:"required"`
}

type FrontendRequestDataMaskingQueryParams struct {
	// Mode: string, required
	Mode terra.StringValue `hcl:"mode,attr" validate:"required"`
	// Value: string, required
	Value terra.StringValue `hcl:"value,attr" validate:"required"`
}

type FrontendResponse struct {
	// BodyBytes: number, optional
	BodyBytes terra.NumberValue `hcl:"body_bytes,attr"`
	// HeadersToLog: set of string, optional
	HeadersToLog terra.SetValue[terra.StringValue] `hcl:"headers_to_log,attr"`
	// FrontendResponseDataMasking: optional
	DataMasking *FrontendResponseDataMasking `hcl:"data_masking,block"`
}

type FrontendResponseDataMasking struct {
	// FrontendResponseDataMaskingHeaders: min=0
	Headers []FrontendResponseDataMaskingHeaders `hcl:"headers,block" validate:"min=0"`
	// FrontendResponseDataMaskingQueryParams: min=0
	QueryParams []FrontendResponseDataMaskingQueryParams `hcl:"query_params,block" validate:"min=0"`
}

type FrontendResponseDataMaskingHeaders struct {
	// Mode: string, required
	Mode terra.StringValue `hcl:"mode,attr" validate:"required"`
	// Value: string, required
	Value terra.StringValue `hcl:"value,attr" validate:"required"`
}

type FrontendResponseDataMaskingQueryParams struct {
	// Mode: string, required
	Mode terra.StringValue `hcl:"mode,attr" validate:"required"`
	// Value: string, required
	Value terra.StringValue `hcl:"value,attr" validate:"required"`
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

type BackendRequestAttributes struct {
	ref terra.Reference
}

func (br BackendRequestAttributes) InternalRef() (terra.Reference, error) {
	return br.ref, nil
}

func (br BackendRequestAttributes) InternalWithRef(ref terra.Reference) BackendRequestAttributes {
	return BackendRequestAttributes{ref: ref}
}

func (br BackendRequestAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return br.ref.InternalTokens()
}

func (br BackendRequestAttributes) BodyBytes() terra.NumberValue {
	return terra.ReferenceAsNumber(br.ref.Append("body_bytes"))
}

func (br BackendRequestAttributes) HeadersToLog() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](br.ref.Append("headers_to_log"))
}

func (br BackendRequestAttributes) DataMasking() terra.ListValue[BackendRequestDataMaskingAttributes] {
	return terra.ReferenceAsList[BackendRequestDataMaskingAttributes](br.ref.Append("data_masking"))
}

type BackendRequestDataMaskingAttributes struct {
	ref terra.Reference
}

func (dm BackendRequestDataMaskingAttributes) InternalRef() (terra.Reference, error) {
	return dm.ref, nil
}

func (dm BackendRequestDataMaskingAttributes) InternalWithRef(ref terra.Reference) BackendRequestDataMaskingAttributes {
	return BackendRequestDataMaskingAttributes{ref: ref}
}

func (dm BackendRequestDataMaskingAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return dm.ref.InternalTokens()
}

func (dm BackendRequestDataMaskingAttributes) Headers() terra.ListValue[BackendRequestDataMaskingHeadersAttributes] {
	return terra.ReferenceAsList[BackendRequestDataMaskingHeadersAttributes](dm.ref.Append("headers"))
}

func (dm BackendRequestDataMaskingAttributes) QueryParams() terra.ListValue[BackendRequestDataMaskingQueryParamsAttributes] {
	return terra.ReferenceAsList[BackendRequestDataMaskingQueryParamsAttributes](dm.ref.Append("query_params"))
}

type BackendRequestDataMaskingHeadersAttributes struct {
	ref terra.Reference
}

func (h BackendRequestDataMaskingHeadersAttributes) InternalRef() (terra.Reference, error) {
	return h.ref, nil
}

func (h BackendRequestDataMaskingHeadersAttributes) InternalWithRef(ref terra.Reference) BackendRequestDataMaskingHeadersAttributes {
	return BackendRequestDataMaskingHeadersAttributes{ref: ref}
}

func (h BackendRequestDataMaskingHeadersAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return h.ref.InternalTokens()
}

func (h BackendRequestDataMaskingHeadersAttributes) Mode() terra.StringValue {
	return terra.ReferenceAsString(h.ref.Append("mode"))
}

func (h BackendRequestDataMaskingHeadersAttributes) Value() terra.StringValue {
	return terra.ReferenceAsString(h.ref.Append("value"))
}

type BackendRequestDataMaskingQueryParamsAttributes struct {
	ref terra.Reference
}

func (qp BackendRequestDataMaskingQueryParamsAttributes) InternalRef() (terra.Reference, error) {
	return qp.ref, nil
}

func (qp BackendRequestDataMaskingQueryParamsAttributes) InternalWithRef(ref terra.Reference) BackendRequestDataMaskingQueryParamsAttributes {
	return BackendRequestDataMaskingQueryParamsAttributes{ref: ref}
}

func (qp BackendRequestDataMaskingQueryParamsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return qp.ref.InternalTokens()
}

func (qp BackendRequestDataMaskingQueryParamsAttributes) Mode() terra.StringValue {
	return terra.ReferenceAsString(qp.ref.Append("mode"))
}

func (qp BackendRequestDataMaskingQueryParamsAttributes) Value() terra.StringValue {
	return terra.ReferenceAsString(qp.ref.Append("value"))
}

type BackendResponseAttributes struct {
	ref terra.Reference
}

func (br BackendResponseAttributes) InternalRef() (terra.Reference, error) {
	return br.ref, nil
}

func (br BackendResponseAttributes) InternalWithRef(ref terra.Reference) BackendResponseAttributes {
	return BackendResponseAttributes{ref: ref}
}

func (br BackendResponseAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return br.ref.InternalTokens()
}

func (br BackendResponseAttributes) BodyBytes() terra.NumberValue {
	return terra.ReferenceAsNumber(br.ref.Append("body_bytes"))
}

func (br BackendResponseAttributes) HeadersToLog() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](br.ref.Append("headers_to_log"))
}

func (br BackendResponseAttributes) DataMasking() terra.ListValue[BackendResponseDataMaskingAttributes] {
	return terra.ReferenceAsList[BackendResponseDataMaskingAttributes](br.ref.Append("data_masking"))
}

type BackendResponseDataMaskingAttributes struct {
	ref terra.Reference
}

func (dm BackendResponseDataMaskingAttributes) InternalRef() (terra.Reference, error) {
	return dm.ref, nil
}

func (dm BackendResponseDataMaskingAttributes) InternalWithRef(ref terra.Reference) BackendResponseDataMaskingAttributes {
	return BackendResponseDataMaskingAttributes{ref: ref}
}

func (dm BackendResponseDataMaskingAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return dm.ref.InternalTokens()
}

func (dm BackendResponseDataMaskingAttributes) Headers() terra.ListValue[BackendResponseDataMaskingHeadersAttributes] {
	return terra.ReferenceAsList[BackendResponseDataMaskingHeadersAttributes](dm.ref.Append("headers"))
}

func (dm BackendResponseDataMaskingAttributes) QueryParams() terra.ListValue[BackendResponseDataMaskingQueryParamsAttributes] {
	return terra.ReferenceAsList[BackendResponseDataMaskingQueryParamsAttributes](dm.ref.Append("query_params"))
}

type BackendResponseDataMaskingHeadersAttributes struct {
	ref terra.Reference
}

func (h BackendResponseDataMaskingHeadersAttributes) InternalRef() (terra.Reference, error) {
	return h.ref, nil
}

func (h BackendResponseDataMaskingHeadersAttributes) InternalWithRef(ref terra.Reference) BackendResponseDataMaskingHeadersAttributes {
	return BackendResponseDataMaskingHeadersAttributes{ref: ref}
}

func (h BackendResponseDataMaskingHeadersAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return h.ref.InternalTokens()
}

func (h BackendResponseDataMaskingHeadersAttributes) Mode() terra.StringValue {
	return terra.ReferenceAsString(h.ref.Append("mode"))
}

func (h BackendResponseDataMaskingHeadersAttributes) Value() terra.StringValue {
	return terra.ReferenceAsString(h.ref.Append("value"))
}

type BackendResponseDataMaskingQueryParamsAttributes struct {
	ref terra.Reference
}

func (qp BackendResponseDataMaskingQueryParamsAttributes) InternalRef() (terra.Reference, error) {
	return qp.ref, nil
}

func (qp BackendResponseDataMaskingQueryParamsAttributes) InternalWithRef(ref terra.Reference) BackendResponseDataMaskingQueryParamsAttributes {
	return BackendResponseDataMaskingQueryParamsAttributes{ref: ref}
}

func (qp BackendResponseDataMaskingQueryParamsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return qp.ref.InternalTokens()
}

func (qp BackendResponseDataMaskingQueryParamsAttributes) Mode() terra.StringValue {
	return terra.ReferenceAsString(qp.ref.Append("mode"))
}

func (qp BackendResponseDataMaskingQueryParamsAttributes) Value() terra.StringValue {
	return terra.ReferenceAsString(qp.ref.Append("value"))
}

type FrontendRequestAttributes struct {
	ref terra.Reference
}

func (fr FrontendRequestAttributes) InternalRef() (terra.Reference, error) {
	return fr.ref, nil
}

func (fr FrontendRequestAttributes) InternalWithRef(ref terra.Reference) FrontendRequestAttributes {
	return FrontendRequestAttributes{ref: ref}
}

func (fr FrontendRequestAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return fr.ref.InternalTokens()
}

func (fr FrontendRequestAttributes) BodyBytes() terra.NumberValue {
	return terra.ReferenceAsNumber(fr.ref.Append("body_bytes"))
}

func (fr FrontendRequestAttributes) HeadersToLog() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](fr.ref.Append("headers_to_log"))
}

func (fr FrontendRequestAttributes) DataMasking() terra.ListValue[FrontendRequestDataMaskingAttributes] {
	return terra.ReferenceAsList[FrontendRequestDataMaskingAttributes](fr.ref.Append("data_masking"))
}

type FrontendRequestDataMaskingAttributes struct {
	ref terra.Reference
}

func (dm FrontendRequestDataMaskingAttributes) InternalRef() (terra.Reference, error) {
	return dm.ref, nil
}

func (dm FrontendRequestDataMaskingAttributes) InternalWithRef(ref terra.Reference) FrontendRequestDataMaskingAttributes {
	return FrontendRequestDataMaskingAttributes{ref: ref}
}

func (dm FrontendRequestDataMaskingAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return dm.ref.InternalTokens()
}

func (dm FrontendRequestDataMaskingAttributes) Headers() terra.ListValue[FrontendRequestDataMaskingHeadersAttributes] {
	return terra.ReferenceAsList[FrontendRequestDataMaskingHeadersAttributes](dm.ref.Append("headers"))
}

func (dm FrontendRequestDataMaskingAttributes) QueryParams() terra.ListValue[FrontendRequestDataMaskingQueryParamsAttributes] {
	return terra.ReferenceAsList[FrontendRequestDataMaskingQueryParamsAttributes](dm.ref.Append("query_params"))
}

type FrontendRequestDataMaskingHeadersAttributes struct {
	ref terra.Reference
}

func (h FrontendRequestDataMaskingHeadersAttributes) InternalRef() (terra.Reference, error) {
	return h.ref, nil
}

func (h FrontendRequestDataMaskingHeadersAttributes) InternalWithRef(ref terra.Reference) FrontendRequestDataMaskingHeadersAttributes {
	return FrontendRequestDataMaskingHeadersAttributes{ref: ref}
}

func (h FrontendRequestDataMaskingHeadersAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return h.ref.InternalTokens()
}

func (h FrontendRequestDataMaskingHeadersAttributes) Mode() terra.StringValue {
	return terra.ReferenceAsString(h.ref.Append("mode"))
}

func (h FrontendRequestDataMaskingHeadersAttributes) Value() terra.StringValue {
	return terra.ReferenceAsString(h.ref.Append("value"))
}

type FrontendRequestDataMaskingQueryParamsAttributes struct {
	ref terra.Reference
}

func (qp FrontendRequestDataMaskingQueryParamsAttributes) InternalRef() (terra.Reference, error) {
	return qp.ref, nil
}

func (qp FrontendRequestDataMaskingQueryParamsAttributes) InternalWithRef(ref terra.Reference) FrontendRequestDataMaskingQueryParamsAttributes {
	return FrontendRequestDataMaskingQueryParamsAttributes{ref: ref}
}

func (qp FrontendRequestDataMaskingQueryParamsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return qp.ref.InternalTokens()
}

func (qp FrontendRequestDataMaskingQueryParamsAttributes) Mode() terra.StringValue {
	return terra.ReferenceAsString(qp.ref.Append("mode"))
}

func (qp FrontendRequestDataMaskingQueryParamsAttributes) Value() terra.StringValue {
	return terra.ReferenceAsString(qp.ref.Append("value"))
}

type FrontendResponseAttributes struct {
	ref terra.Reference
}

func (fr FrontendResponseAttributes) InternalRef() (terra.Reference, error) {
	return fr.ref, nil
}

func (fr FrontendResponseAttributes) InternalWithRef(ref terra.Reference) FrontendResponseAttributes {
	return FrontendResponseAttributes{ref: ref}
}

func (fr FrontendResponseAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return fr.ref.InternalTokens()
}

func (fr FrontendResponseAttributes) BodyBytes() terra.NumberValue {
	return terra.ReferenceAsNumber(fr.ref.Append("body_bytes"))
}

func (fr FrontendResponseAttributes) HeadersToLog() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](fr.ref.Append("headers_to_log"))
}

func (fr FrontendResponseAttributes) DataMasking() terra.ListValue[FrontendResponseDataMaskingAttributes] {
	return terra.ReferenceAsList[FrontendResponseDataMaskingAttributes](fr.ref.Append("data_masking"))
}

type FrontendResponseDataMaskingAttributes struct {
	ref terra.Reference
}

func (dm FrontendResponseDataMaskingAttributes) InternalRef() (terra.Reference, error) {
	return dm.ref, nil
}

func (dm FrontendResponseDataMaskingAttributes) InternalWithRef(ref terra.Reference) FrontendResponseDataMaskingAttributes {
	return FrontendResponseDataMaskingAttributes{ref: ref}
}

func (dm FrontendResponseDataMaskingAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return dm.ref.InternalTokens()
}

func (dm FrontendResponseDataMaskingAttributes) Headers() terra.ListValue[FrontendResponseDataMaskingHeadersAttributes] {
	return terra.ReferenceAsList[FrontendResponseDataMaskingHeadersAttributes](dm.ref.Append("headers"))
}

func (dm FrontendResponseDataMaskingAttributes) QueryParams() terra.ListValue[FrontendResponseDataMaskingQueryParamsAttributes] {
	return terra.ReferenceAsList[FrontendResponseDataMaskingQueryParamsAttributes](dm.ref.Append("query_params"))
}

type FrontendResponseDataMaskingHeadersAttributes struct {
	ref terra.Reference
}

func (h FrontendResponseDataMaskingHeadersAttributes) InternalRef() (terra.Reference, error) {
	return h.ref, nil
}

func (h FrontendResponseDataMaskingHeadersAttributes) InternalWithRef(ref terra.Reference) FrontendResponseDataMaskingHeadersAttributes {
	return FrontendResponseDataMaskingHeadersAttributes{ref: ref}
}

func (h FrontendResponseDataMaskingHeadersAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return h.ref.InternalTokens()
}

func (h FrontendResponseDataMaskingHeadersAttributes) Mode() terra.StringValue {
	return terra.ReferenceAsString(h.ref.Append("mode"))
}

func (h FrontendResponseDataMaskingHeadersAttributes) Value() terra.StringValue {
	return terra.ReferenceAsString(h.ref.Append("value"))
}

type FrontendResponseDataMaskingQueryParamsAttributes struct {
	ref terra.Reference
}

func (qp FrontendResponseDataMaskingQueryParamsAttributes) InternalRef() (terra.Reference, error) {
	return qp.ref, nil
}

func (qp FrontendResponseDataMaskingQueryParamsAttributes) InternalWithRef(ref terra.Reference) FrontendResponseDataMaskingQueryParamsAttributes {
	return FrontendResponseDataMaskingQueryParamsAttributes{ref: ref}
}

func (qp FrontendResponseDataMaskingQueryParamsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return qp.ref.InternalTokens()
}

func (qp FrontendResponseDataMaskingQueryParamsAttributes) Mode() terra.StringValue {
	return terra.ReferenceAsString(qp.ref.Append("mode"))
}

func (qp FrontendResponseDataMaskingQueryParamsAttributes) Value() terra.StringValue {
	return terra.ReferenceAsString(qp.ref.Append("value"))
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

type BackendRequestState struct {
	BodyBytes    float64                          `json:"body_bytes"`
	HeadersToLog []string                         `json:"headers_to_log"`
	DataMasking  []BackendRequestDataMaskingState `json:"data_masking"`
}

type BackendRequestDataMaskingState struct {
	Headers     []BackendRequestDataMaskingHeadersState     `json:"headers"`
	QueryParams []BackendRequestDataMaskingQueryParamsState `json:"query_params"`
}

type BackendRequestDataMaskingHeadersState struct {
	Mode  string `json:"mode"`
	Value string `json:"value"`
}

type BackendRequestDataMaskingQueryParamsState struct {
	Mode  string `json:"mode"`
	Value string `json:"value"`
}

type BackendResponseState struct {
	BodyBytes    float64                           `json:"body_bytes"`
	HeadersToLog []string                          `json:"headers_to_log"`
	DataMasking  []BackendResponseDataMaskingState `json:"data_masking"`
}

type BackendResponseDataMaskingState struct {
	Headers     []BackendResponseDataMaskingHeadersState     `json:"headers"`
	QueryParams []BackendResponseDataMaskingQueryParamsState `json:"query_params"`
}

type BackendResponseDataMaskingHeadersState struct {
	Mode  string `json:"mode"`
	Value string `json:"value"`
}

type BackendResponseDataMaskingQueryParamsState struct {
	Mode  string `json:"mode"`
	Value string `json:"value"`
}

type FrontendRequestState struct {
	BodyBytes    float64                           `json:"body_bytes"`
	HeadersToLog []string                          `json:"headers_to_log"`
	DataMasking  []FrontendRequestDataMaskingState `json:"data_masking"`
}

type FrontendRequestDataMaskingState struct {
	Headers     []FrontendRequestDataMaskingHeadersState     `json:"headers"`
	QueryParams []FrontendRequestDataMaskingQueryParamsState `json:"query_params"`
}

type FrontendRequestDataMaskingHeadersState struct {
	Mode  string `json:"mode"`
	Value string `json:"value"`
}

type FrontendRequestDataMaskingQueryParamsState struct {
	Mode  string `json:"mode"`
	Value string `json:"value"`
}

type FrontendResponseState struct {
	BodyBytes    float64                            `json:"body_bytes"`
	HeadersToLog []string                           `json:"headers_to_log"`
	DataMasking  []FrontendResponseDataMaskingState `json:"data_masking"`
}

type FrontendResponseDataMaskingState struct {
	Headers     []FrontendResponseDataMaskingHeadersState     `json:"headers"`
	QueryParams []FrontendResponseDataMaskingQueryParamsState `json:"query_params"`
}

type FrontendResponseDataMaskingHeadersState struct {
	Mode  string `json:"mode"`
	Value string `json:"value"`
}

type FrontendResponseDataMaskingQueryParamsState struct {
	Mode  string `json:"mode"`
	Value string `json:"value"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
}
