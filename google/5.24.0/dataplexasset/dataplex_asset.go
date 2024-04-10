// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package dataplexasset

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type DiscoveryStatus struct {
	// Stats: min=0
	Stats []Stats `hcl:"stats,block" validate:"min=0"`
}

type Stats struct{}

type ResourceStatus struct{}

type SecurityStatus struct{}

type DiscoverySpec struct {
	// Enabled: bool, required
	Enabled terra.BoolValue `hcl:"enabled,attr" validate:"required"`
	// ExcludePatterns: list of string, optional
	ExcludePatterns terra.ListValue[terra.StringValue] `hcl:"exclude_patterns,attr"`
	// IncludePatterns: list of string, optional
	IncludePatterns terra.ListValue[terra.StringValue] `hcl:"include_patterns,attr"`
	// Schedule: string, optional
	Schedule terra.StringValue `hcl:"schedule,attr"`
	// CsvOptions: optional
	CsvOptions *CsvOptions `hcl:"csv_options,block"`
	// JsonOptions: optional
	JsonOptions *JsonOptions `hcl:"json_options,block"`
}

type CsvOptions struct {
	// Delimiter: string, optional
	Delimiter terra.StringValue `hcl:"delimiter,attr"`
	// DisableTypeInference: bool, optional
	DisableTypeInference terra.BoolValue `hcl:"disable_type_inference,attr"`
	// Encoding: string, optional
	Encoding terra.StringValue `hcl:"encoding,attr"`
	// HeaderRows: number, optional
	HeaderRows terra.NumberValue `hcl:"header_rows,attr"`
}

type JsonOptions struct {
	// DisableTypeInference: bool, optional
	DisableTypeInference terra.BoolValue `hcl:"disable_type_inference,attr"`
	// Encoding: string, optional
	Encoding terra.StringValue `hcl:"encoding,attr"`
}

type ResourceSpec struct {
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
	// ReadAccessMode: string, optional
	ReadAccessMode terra.StringValue `hcl:"read_access_mode,attr"`
	// Type: string, required
	Type terra.StringValue `hcl:"type,attr" validate:"required"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type DiscoveryStatusAttributes struct {
	ref terra.Reference
}

func (ds DiscoveryStatusAttributes) InternalRef() (terra.Reference, error) {
	return ds.ref, nil
}

func (ds DiscoveryStatusAttributes) InternalWithRef(ref terra.Reference) DiscoveryStatusAttributes {
	return DiscoveryStatusAttributes{ref: ref}
}

func (ds DiscoveryStatusAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ds.ref.InternalTokens()
}

func (ds DiscoveryStatusAttributes) LastRunDuration() terra.StringValue {
	return terra.ReferenceAsString(ds.ref.Append("last_run_duration"))
}

func (ds DiscoveryStatusAttributes) LastRunTime() terra.StringValue {
	return terra.ReferenceAsString(ds.ref.Append("last_run_time"))
}

func (ds DiscoveryStatusAttributes) Message() terra.StringValue {
	return terra.ReferenceAsString(ds.ref.Append("message"))
}

func (ds DiscoveryStatusAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(ds.ref.Append("state"))
}

func (ds DiscoveryStatusAttributes) UpdateTime() terra.StringValue {
	return terra.ReferenceAsString(ds.ref.Append("update_time"))
}

func (ds DiscoveryStatusAttributes) Stats() terra.ListValue[StatsAttributes] {
	return terra.ReferenceAsList[StatsAttributes](ds.ref.Append("stats"))
}

type StatsAttributes struct {
	ref terra.Reference
}

func (s StatsAttributes) InternalRef() (terra.Reference, error) {
	return s.ref, nil
}

func (s StatsAttributes) InternalWithRef(ref terra.Reference) StatsAttributes {
	return StatsAttributes{ref: ref}
}

func (s StatsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return s.ref.InternalTokens()
}

func (s StatsAttributes) DataItems() terra.NumberValue {
	return terra.ReferenceAsNumber(s.ref.Append("data_items"))
}

func (s StatsAttributes) DataSize() terra.NumberValue {
	return terra.ReferenceAsNumber(s.ref.Append("data_size"))
}

func (s StatsAttributes) Filesets() terra.NumberValue {
	return terra.ReferenceAsNumber(s.ref.Append("filesets"))
}

func (s StatsAttributes) Tables() terra.NumberValue {
	return terra.ReferenceAsNumber(s.ref.Append("tables"))
}

type ResourceStatusAttributes struct {
	ref terra.Reference
}

func (rs ResourceStatusAttributes) InternalRef() (terra.Reference, error) {
	return rs.ref, nil
}

func (rs ResourceStatusAttributes) InternalWithRef(ref terra.Reference) ResourceStatusAttributes {
	return ResourceStatusAttributes{ref: ref}
}

func (rs ResourceStatusAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return rs.ref.InternalTokens()
}

func (rs ResourceStatusAttributes) Message() terra.StringValue {
	return terra.ReferenceAsString(rs.ref.Append("message"))
}

func (rs ResourceStatusAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(rs.ref.Append("state"))
}

func (rs ResourceStatusAttributes) UpdateTime() terra.StringValue {
	return terra.ReferenceAsString(rs.ref.Append("update_time"))
}

type SecurityStatusAttributes struct {
	ref terra.Reference
}

func (ss SecurityStatusAttributes) InternalRef() (terra.Reference, error) {
	return ss.ref, nil
}

func (ss SecurityStatusAttributes) InternalWithRef(ref terra.Reference) SecurityStatusAttributes {
	return SecurityStatusAttributes{ref: ref}
}

func (ss SecurityStatusAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ss.ref.InternalTokens()
}

func (ss SecurityStatusAttributes) Message() terra.StringValue {
	return terra.ReferenceAsString(ss.ref.Append("message"))
}

func (ss SecurityStatusAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(ss.ref.Append("state"))
}

func (ss SecurityStatusAttributes) UpdateTime() terra.StringValue {
	return terra.ReferenceAsString(ss.ref.Append("update_time"))
}

type DiscoverySpecAttributes struct {
	ref terra.Reference
}

func (ds DiscoverySpecAttributes) InternalRef() (terra.Reference, error) {
	return ds.ref, nil
}

func (ds DiscoverySpecAttributes) InternalWithRef(ref terra.Reference) DiscoverySpecAttributes {
	return DiscoverySpecAttributes{ref: ref}
}

func (ds DiscoverySpecAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ds.ref.InternalTokens()
}

func (ds DiscoverySpecAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(ds.ref.Append("enabled"))
}

func (ds DiscoverySpecAttributes) ExcludePatterns() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](ds.ref.Append("exclude_patterns"))
}

func (ds DiscoverySpecAttributes) IncludePatterns() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](ds.ref.Append("include_patterns"))
}

func (ds DiscoverySpecAttributes) Schedule() terra.StringValue {
	return terra.ReferenceAsString(ds.ref.Append("schedule"))
}

func (ds DiscoverySpecAttributes) CsvOptions() terra.ListValue[CsvOptionsAttributes] {
	return terra.ReferenceAsList[CsvOptionsAttributes](ds.ref.Append("csv_options"))
}

func (ds DiscoverySpecAttributes) JsonOptions() terra.ListValue[JsonOptionsAttributes] {
	return terra.ReferenceAsList[JsonOptionsAttributes](ds.ref.Append("json_options"))
}

type CsvOptionsAttributes struct {
	ref terra.Reference
}

func (co CsvOptionsAttributes) InternalRef() (terra.Reference, error) {
	return co.ref, nil
}

func (co CsvOptionsAttributes) InternalWithRef(ref terra.Reference) CsvOptionsAttributes {
	return CsvOptionsAttributes{ref: ref}
}

func (co CsvOptionsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return co.ref.InternalTokens()
}

func (co CsvOptionsAttributes) Delimiter() terra.StringValue {
	return terra.ReferenceAsString(co.ref.Append("delimiter"))
}

func (co CsvOptionsAttributes) DisableTypeInference() terra.BoolValue {
	return terra.ReferenceAsBool(co.ref.Append("disable_type_inference"))
}

func (co CsvOptionsAttributes) Encoding() terra.StringValue {
	return terra.ReferenceAsString(co.ref.Append("encoding"))
}

func (co CsvOptionsAttributes) HeaderRows() terra.NumberValue {
	return terra.ReferenceAsNumber(co.ref.Append("header_rows"))
}

type JsonOptionsAttributes struct {
	ref terra.Reference
}

func (jo JsonOptionsAttributes) InternalRef() (terra.Reference, error) {
	return jo.ref, nil
}

func (jo JsonOptionsAttributes) InternalWithRef(ref terra.Reference) JsonOptionsAttributes {
	return JsonOptionsAttributes{ref: ref}
}

func (jo JsonOptionsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return jo.ref.InternalTokens()
}

func (jo JsonOptionsAttributes) DisableTypeInference() terra.BoolValue {
	return terra.ReferenceAsBool(jo.ref.Append("disable_type_inference"))
}

func (jo JsonOptionsAttributes) Encoding() terra.StringValue {
	return terra.ReferenceAsString(jo.ref.Append("encoding"))
}

type ResourceSpecAttributes struct {
	ref terra.Reference
}

func (rs ResourceSpecAttributes) InternalRef() (terra.Reference, error) {
	return rs.ref, nil
}

func (rs ResourceSpecAttributes) InternalWithRef(ref terra.Reference) ResourceSpecAttributes {
	return ResourceSpecAttributes{ref: ref}
}

func (rs ResourceSpecAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return rs.ref.InternalTokens()
}

func (rs ResourceSpecAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(rs.ref.Append("name"))
}

func (rs ResourceSpecAttributes) ReadAccessMode() terra.StringValue {
	return terra.ReferenceAsString(rs.ref.Append("read_access_mode"))
}

func (rs ResourceSpecAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(rs.ref.Append("type"))
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

type DiscoveryStatusState struct {
	LastRunDuration string       `json:"last_run_duration"`
	LastRunTime     string       `json:"last_run_time"`
	Message         string       `json:"message"`
	State           string       `json:"state"`
	UpdateTime      string       `json:"update_time"`
	Stats           []StatsState `json:"stats"`
}

type StatsState struct {
	DataItems float64 `json:"data_items"`
	DataSize  float64 `json:"data_size"`
	Filesets  float64 `json:"filesets"`
	Tables    float64 `json:"tables"`
}

type ResourceStatusState struct {
	Message    string `json:"message"`
	State      string `json:"state"`
	UpdateTime string `json:"update_time"`
}

type SecurityStatusState struct {
	Message    string `json:"message"`
	State      string `json:"state"`
	UpdateTime string `json:"update_time"`
}

type DiscoverySpecState struct {
	Enabled         bool               `json:"enabled"`
	ExcludePatterns []string           `json:"exclude_patterns"`
	IncludePatterns []string           `json:"include_patterns"`
	Schedule        string             `json:"schedule"`
	CsvOptions      []CsvOptionsState  `json:"csv_options"`
	JsonOptions     []JsonOptionsState `json:"json_options"`
}

type CsvOptionsState struct {
	Delimiter            string  `json:"delimiter"`
	DisableTypeInference bool    `json:"disable_type_inference"`
	Encoding             string  `json:"encoding"`
	HeaderRows           float64 `json:"header_rows"`
}

type JsonOptionsState struct {
	DisableTypeInference bool   `json:"disable_type_inference"`
	Encoding             string `json:"encoding"`
}

type ResourceSpecState struct {
	Name           string `json:"name"`
	ReadAccessMode string `json:"read_access_mode"`
	Type           string `json:"type"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}
