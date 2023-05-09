// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package http

import (
	datahttp "github.com/golingon/terraproviders/http/3.3.0/datahttp"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataHttp creates a new instance of [DataHttp].
func NewDataHttp(name string, args DataHttpArgs) *DataHttp {
	return &DataHttp{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataHttp)(nil)

// DataHttp represents the Terraform data resource http.
type DataHttp struct {
	Name string
	Args DataHttpArgs
}

// DataSource returns the Terraform object type for [DataHttp].
func (h *DataHttp) DataSource() string {
	return "http"
}

// LocalName returns the local name for [DataHttp].
func (h *DataHttp) LocalName() string {
	return h.Name
}

// Configuration returns the configuration (args) for [DataHttp].
func (h *DataHttp) Configuration() interface{} {
	return h.Args
}

// Attributes returns the attributes for [DataHttp].
func (h *DataHttp) Attributes() dataHttpAttributes {
	return dataHttpAttributes{ref: terra.ReferenceDataResource(h)}
}

// DataHttpArgs contains the configurations for http.
type DataHttpArgs struct {
	// CaCertPem: string, optional
	CaCertPem terra.StringValue `hcl:"ca_cert_pem,attr"`
	// Insecure: bool, optional
	Insecure terra.BoolValue `hcl:"insecure,attr"`
	// Method: string, optional
	Method terra.StringValue `hcl:"method,attr"`
	// RequestBody: string, optional
	RequestBody terra.StringValue `hcl:"request_body,attr"`
	// RequestHeaders: map of string, optional
	RequestHeaders terra.MapValue[terra.StringValue] `hcl:"request_headers,attr"`
	// RequestTimeoutMs: number, optional
	RequestTimeoutMs terra.NumberValue `hcl:"request_timeout_ms,attr"`
	// Url: string, required
	Url terra.StringValue `hcl:"url,attr" validate:"required"`
	// Retry: optional
	Retry *datahttp.Retry `hcl:"retry,block"`
}
type dataHttpAttributes struct {
	ref terra.Reference
}

// Body returns a reference to field body of http.
func (h dataHttpAttributes) Body() terra.StringValue {
	return terra.ReferenceAsString(h.ref.Append("body"))
}

// CaCertPem returns a reference to field ca_cert_pem of http.
func (h dataHttpAttributes) CaCertPem() terra.StringValue {
	return terra.ReferenceAsString(h.ref.Append("ca_cert_pem"))
}

// Id returns a reference to field id of http.
func (h dataHttpAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(h.ref.Append("id"))
}

// Insecure returns a reference to field insecure of http.
func (h dataHttpAttributes) Insecure() terra.BoolValue {
	return terra.ReferenceAsBool(h.ref.Append("insecure"))
}

// Method returns a reference to field method of http.
func (h dataHttpAttributes) Method() terra.StringValue {
	return terra.ReferenceAsString(h.ref.Append("method"))
}

// RequestBody returns a reference to field request_body of http.
func (h dataHttpAttributes) RequestBody() terra.StringValue {
	return terra.ReferenceAsString(h.ref.Append("request_body"))
}

// RequestHeaders returns a reference to field request_headers of http.
func (h dataHttpAttributes) RequestHeaders() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](h.ref.Append("request_headers"))
}

// RequestTimeoutMs returns a reference to field request_timeout_ms of http.
func (h dataHttpAttributes) RequestTimeoutMs() terra.NumberValue {
	return terra.ReferenceAsNumber(h.ref.Append("request_timeout_ms"))
}

// ResponseBody returns a reference to field response_body of http.
func (h dataHttpAttributes) ResponseBody() terra.StringValue {
	return terra.ReferenceAsString(h.ref.Append("response_body"))
}

// ResponseHeaders returns a reference to field response_headers of http.
func (h dataHttpAttributes) ResponseHeaders() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](h.ref.Append("response_headers"))
}

// StatusCode returns a reference to field status_code of http.
func (h dataHttpAttributes) StatusCode() terra.NumberValue {
	return terra.ReferenceAsNumber(h.ref.Append("status_code"))
}

// Url returns a reference to field url of http.
func (h dataHttpAttributes) Url() terra.StringValue {
	return terra.ReferenceAsString(h.ref.Append("url"))
}

func (h dataHttpAttributes) Retry() datahttp.RetryAttributes {
	return terra.ReferenceAsSingle[datahttp.RetryAttributes](h.ref.Append("retry"))
}
