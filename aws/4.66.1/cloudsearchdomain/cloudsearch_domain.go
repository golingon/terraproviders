// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package cloudsearchdomain

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type EndpointOptions struct {
	// EnforceHttps: bool, optional
	EnforceHttps terra.BoolValue `hcl:"enforce_https,attr"`
	// TlsSecurityPolicy: string, optional
	TlsSecurityPolicy terra.StringValue `hcl:"tls_security_policy,attr"`
}

type IndexField struct {
	// AnalysisScheme: string, optional
	AnalysisScheme terra.StringValue `hcl:"analysis_scheme,attr"`
	// DefaultValue: string, optional
	DefaultValue terra.StringValue `hcl:"default_value,attr"`
	// Facet: bool, optional
	Facet terra.BoolValue `hcl:"facet,attr"`
	// Highlight: bool, optional
	Highlight terra.BoolValue `hcl:"highlight,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Return: bool, optional
	Return terra.BoolValue `hcl:"return,attr"`
	// Search: bool, optional
	Search terra.BoolValue `hcl:"search,attr"`
	// Sort: bool, optional
	Sort terra.BoolValue `hcl:"sort,attr"`
	// SourceFields: string, optional
	SourceFields terra.StringValue `hcl:"source_fields,attr"`
	// Type: string, required
	Type terra.StringValue `hcl:"type,attr" validate:"required"`
}

type ScalingParameters struct {
	// DesiredInstanceType: string, optional
	DesiredInstanceType terra.StringValue `hcl:"desired_instance_type,attr"`
	// DesiredPartitionCount: number, optional
	DesiredPartitionCount terra.NumberValue `hcl:"desired_partition_count,attr"`
	// DesiredReplicationCount: number, optional
	DesiredReplicationCount terra.NumberValue `hcl:"desired_replication_count,attr"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type EndpointOptionsAttributes struct {
	ref terra.Reference
}

func (eo EndpointOptionsAttributes) InternalRef() (terra.Reference, error) {
	return eo.ref, nil
}

func (eo EndpointOptionsAttributes) InternalWithRef(ref terra.Reference) EndpointOptionsAttributes {
	return EndpointOptionsAttributes{ref: ref}
}

func (eo EndpointOptionsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return eo.ref.InternalTokens()
}

func (eo EndpointOptionsAttributes) EnforceHttps() terra.BoolValue {
	return terra.ReferenceAsBool(eo.ref.Append("enforce_https"))
}

func (eo EndpointOptionsAttributes) TlsSecurityPolicy() terra.StringValue {
	return terra.ReferenceAsString(eo.ref.Append("tls_security_policy"))
}

type IndexFieldAttributes struct {
	ref terra.Reference
}

func (_if IndexFieldAttributes) InternalRef() (terra.Reference, error) {
	return _if.ref, nil
}

func (_if IndexFieldAttributes) InternalWithRef(ref terra.Reference) IndexFieldAttributes {
	return IndexFieldAttributes{ref: ref}
}

func (_if IndexFieldAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return _if.ref.InternalTokens()
}

func (_if IndexFieldAttributes) AnalysisScheme() terra.StringValue {
	return terra.ReferenceAsString(_if.ref.Append("analysis_scheme"))
}

func (_if IndexFieldAttributes) DefaultValue() terra.StringValue {
	return terra.ReferenceAsString(_if.ref.Append("default_value"))
}

func (_if IndexFieldAttributes) Facet() terra.BoolValue {
	return terra.ReferenceAsBool(_if.ref.Append("facet"))
}

func (_if IndexFieldAttributes) Highlight() terra.BoolValue {
	return terra.ReferenceAsBool(_if.ref.Append("highlight"))
}

func (_if IndexFieldAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(_if.ref.Append("name"))
}

func (_if IndexFieldAttributes) Return() terra.BoolValue {
	return terra.ReferenceAsBool(_if.ref.Append("return"))
}

func (_if IndexFieldAttributes) Search() terra.BoolValue {
	return terra.ReferenceAsBool(_if.ref.Append("search"))
}

func (_if IndexFieldAttributes) Sort() terra.BoolValue {
	return terra.ReferenceAsBool(_if.ref.Append("sort"))
}

func (_if IndexFieldAttributes) SourceFields() terra.StringValue {
	return terra.ReferenceAsString(_if.ref.Append("source_fields"))
}

func (_if IndexFieldAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(_if.ref.Append("type"))
}

type ScalingParametersAttributes struct {
	ref terra.Reference
}

func (sp ScalingParametersAttributes) InternalRef() (terra.Reference, error) {
	return sp.ref, nil
}

func (sp ScalingParametersAttributes) InternalWithRef(ref terra.Reference) ScalingParametersAttributes {
	return ScalingParametersAttributes{ref: ref}
}

func (sp ScalingParametersAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return sp.ref.InternalTokens()
}

func (sp ScalingParametersAttributes) DesiredInstanceType() terra.StringValue {
	return terra.ReferenceAsString(sp.ref.Append("desired_instance_type"))
}

func (sp ScalingParametersAttributes) DesiredPartitionCount() terra.NumberValue {
	return terra.ReferenceAsNumber(sp.ref.Append("desired_partition_count"))
}

func (sp ScalingParametersAttributes) DesiredReplicationCount() terra.NumberValue {
	return terra.ReferenceAsNumber(sp.ref.Append("desired_replication_count"))
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

type EndpointOptionsState struct {
	EnforceHttps      bool   `json:"enforce_https"`
	TlsSecurityPolicy string `json:"tls_security_policy"`
}

type IndexFieldState struct {
	AnalysisScheme string `json:"analysis_scheme"`
	DefaultValue   string `json:"default_value"`
	Facet          bool   `json:"facet"`
	Highlight      bool   `json:"highlight"`
	Name           string `json:"name"`
	Return         bool   `json:"return"`
	Search         bool   `json:"search"`
	Sort           bool   `json:"sort"`
	SourceFields   string `json:"source_fields"`
	Type           string `json:"type"`
}

type ScalingParametersState struct {
	DesiredInstanceType     string  `json:"desired_instance_type"`
	DesiredPartitionCount   float64 `json:"desired_partition_count"`
	DesiredReplicationCount float64 `json:"desired_replication_count"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}
