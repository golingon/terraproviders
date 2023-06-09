// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package glueclassifier

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type CsvClassifier struct {
	// AllowSingleColumn: bool, optional
	AllowSingleColumn terra.BoolValue `hcl:"allow_single_column,attr"`
	// ContainsHeader: string, optional
	ContainsHeader terra.StringValue `hcl:"contains_header,attr"`
	// CustomDatatypeConfigured: bool, optional
	CustomDatatypeConfigured terra.BoolValue `hcl:"custom_datatype_configured,attr"`
	// CustomDatatypes: list of string, optional
	CustomDatatypes terra.ListValue[terra.StringValue] `hcl:"custom_datatypes,attr"`
	// Delimiter: string, optional
	Delimiter terra.StringValue `hcl:"delimiter,attr"`
	// DisableValueTrimming: bool, optional
	DisableValueTrimming terra.BoolValue `hcl:"disable_value_trimming,attr"`
	// Header: list of string, optional
	Header terra.ListValue[terra.StringValue] `hcl:"header,attr"`
	// QuoteSymbol: string, optional
	QuoteSymbol terra.StringValue `hcl:"quote_symbol,attr"`
}

type GrokClassifier struct {
	// Classification: string, required
	Classification terra.StringValue `hcl:"classification,attr" validate:"required"`
	// CustomPatterns: string, optional
	CustomPatterns terra.StringValue `hcl:"custom_patterns,attr"`
	// GrokPattern: string, required
	GrokPattern terra.StringValue `hcl:"grok_pattern,attr" validate:"required"`
}

type JsonClassifier struct {
	// JsonPath: string, required
	JsonPath terra.StringValue `hcl:"json_path,attr" validate:"required"`
}

type XmlClassifier struct {
	// Classification: string, required
	Classification terra.StringValue `hcl:"classification,attr" validate:"required"`
	// RowTag: string, required
	RowTag terra.StringValue `hcl:"row_tag,attr" validate:"required"`
}

type CsvClassifierAttributes struct {
	ref terra.Reference
}

func (cc CsvClassifierAttributes) InternalRef() (terra.Reference, error) {
	return cc.ref, nil
}

func (cc CsvClassifierAttributes) InternalWithRef(ref terra.Reference) CsvClassifierAttributes {
	return CsvClassifierAttributes{ref: ref}
}

func (cc CsvClassifierAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return cc.ref.InternalTokens()
}

func (cc CsvClassifierAttributes) AllowSingleColumn() terra.BoolValue {
	return terra.ReferenceAsBool(cc.ref.Append("allow_single_column"))
}

func (cc CsvClassifierAttributes) ContainsHeader() terra.StringValue {
	return terra.ReferenceAsString(cc.ref.Append("contains_header"))
}

func (cc CsvClassifierAttributes) CustomDatatypeConfigured() terra.BoolValue {
	return terra.ReferenceAsBool(cc.ref.Append("custom_datatype_configured"))
}

func (cc CsvClassifierAttributes) CustomDatatypes() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](cc.ref.Append("custom_datatypes"))
}

func (cc CsvClassifierAttributes) Delimiter() terra.StringValue {
	return terra.ReferenceAsString(cc.ref.Append("delimiter"))
}

func (cc CsvClassifierAttributes) DisableValueTrimming() terra.BoolValue {
	return terra.ReferenceAsBool(cc.ref.Append("disable_value_trimming"))
}

func (cc CsvClassifierAttributes) Header() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](cc.ref.Append("header"))
}

func (cc CsvClassifierAttributes) QuoteSymbol() terra.StringValue {
	return terra.ReferenceAsString(cc.ref.Append("quote_symbol"))
}

type GrokClassifierAttributes struct {
	ref terra.Reference
}

func (gc GrokClassifierAttributes) InternalRef() (terra.Reference, error) {
	return gc.ref, nil
}

func (gc GrokClassifierAttributes) InternalWithRef(ref terra.Reference) GrokClassifierAttributes {
	return GrokClassifierAttributes{ref: ref}
}

func (gc GrokClassifierAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return gc.ref.InternalTokens()
}

func (gc GrokClassifierAttributes) Classification() terra.StringValue {
	return terra.ReferenceAsString(gc.ref.Append("classification"))
}

func (gc GrokClassifierAttributes) CustomPatterns() terra.StringValue {
	return terra.ReferenceAsString(gc.ref.Append("custom_patterns"))
}

func (gc GrokClassifierAttributes) GrokPattern() terra.StringValue {
	return terra.ReferenceAsString(gc.ref.Append("grok_pattern"))
}

type JsonClassifierAttributes struct {
	ref terra.Reference
}

func (jc JsonClassifierAttributes) InternalRef() (terra.Reference, error) {
	return jc.ref, nil
}

func (jc JsonClassifierAttributes) InternalWithRef(ref terra.Reference) JsonClassifierAttributes {
	return JsonClassifierAttributes{ref: ref}
}

func (jc JsonClassifierAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return jc.ref.InternalTokens()
}

func (jc JsonClassifierAttributes) JsonPath() terra.StringValue {
	return terra.ReferenceAsString(jc.ref.Append("json_path"))
}

type XmlClassifierAttributes struct {
	ref terra.Reference
}

func (xc XmlClassifierAttributes) InternalRef() (terra.Reference, error) {
	return xc.ref, nil
}

func (xc XmlClassifierAttributes) InternalWithRef(ref terra.Reference) XmlClassifierAttributes {
	return XmlClassifierAttributes{ref: ref}
}

func (xc XmlClassifierAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return xc.ref.InternalTokens()
}

func (xc XmlClassifierAttributes) Classification() terra.StringValue {
	return terra.ReferenceAsString(xc.ref.Append("classification"))
}

func (xc XmlClassifierAttributes) RowTag() terra.StringValue {
	return terra.ReferenceAsString(xc.ref.Append("row_tag"))
}

type CsvClassifierState struct {
	AllowSingleColumn        bool     `json:"allow_single_column"`
	ContainsHeader           string   `json:"contains_header"`
	CustomDatatypeConfigured bool     `json:"custom_datatype_configured"`
	CustomDatatypes          []string `json:"custom_datatypes"`
	Delimiter                string   `json:"delimiter"`
	DisableValueTrimming     bool     `json:"disable_value_trimming"`
	Header                   []string `json:"header"`
	QuoteSymbol              string   `json:"quote_symbol"`
}

type GrokClassifierState struct {
	Classification string `json:"classification"`
	CustomPatterns string `json:"custom_patterns"`
	GrokPattern    string `json:"grok_pattern"`
}

type JsonClassifierState struct {
	JsonPath string `json:"json_path"`
}

type XmlClassifierState struct {
	Classification string `json:"classification"`
	RowTag         string `json:"row_tag"`
}
