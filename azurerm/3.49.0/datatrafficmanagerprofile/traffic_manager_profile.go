// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package datatrafficmanagerprofile

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type DnsConfig struct{}

type MonitorConfig struct {
	// CustomHeader: min=0
	CustomHeader []CustomHeader `hcl:"custom_header,block" validate:"min=0"`
}

type CustomHeader struct{}

type Timeouts struct {
	// Read: string, optional
	Read terra.StringValue `hcl:"read,attr"`
}

type DnsConfigAttributes struct {
	ref terra.Reference
}

func (dc DnsConfigAttributes) InternalRef() (terra.Reference, error) {
	return dc.ref, nil
}

func (dc DnsConfigAttributes) InternalWithRef(ref terra.Reference) DnsConfigAttributes {
	return DnsConfigAttributes{ref: ref}
}

func (dc DnsConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return dc.ref.InternalTokens()
}

func (dc DnsConfigAttributes) RelativeName() terra.StringValue {
	return terra.ReferenceAsString(dc.ref.Append("relative_name"))
}

func (dc DnsConfigAttributes) Ttl() terra.NumberValue {
	return terra.ReferenceAsNumber(dc.ref.Append("ttl"))
}

type MonitorConfigAttributes struct {
	ref terra.Reference
}

func (mc MonitorConfigAttributes) InternalRef() (terra.Reference, error) {
	return mc.ref, nil
}

func (mc MonitorConfigAttributes) InternalWithRef(ref terra.Reference) MonitorConfigAttributes {
	return MonitorConfigAttributes{ref: ref}
}

func (mc MonitorConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return mc.ref.InternalTokens()
}

func (mc MonitorConfigAttributes) ExpectedStatusCodeRanges() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](mc.ref.Append("expected_status_code_ranges"))
}

func (mc MonitorConfigAttributes) IntervalInSeconds() terra.NumberValue {
	return terra.ReferenceAsNumber(mc.ref.Append("interval_in_seconds"))
}

func (mc MonitorConfigAttributes) Path() terra.StringValue {
	return terra.ReferenceAsString(mc.ref.Append("path"))
}

func (mc MonitorConfigAttributes) Port() terra.NumberValue {
	return terra.ReferenceAsNumber(mc.ref.Append("port"))
}

func (mc MonitorConfigAttributes) Protocol() terra.StringValue {
	return terra.ReferenceAsString(mc.ref.Append("protocol"))
}

func (mc MonitorConfigAttributes) TimeoutInSeconds() terra.NumberValue {
	return terra.ReferenceAsNumber(mc.ref.Append("timeout_in_seconds"))
}

func (mc MonitorConfigAttributes) ToleratedNumberOfFailures() terra.NumberValue {
	return terra.ReferenceAsNumber(mc.ref.Append("tolerated_number_of_failures"))
}

func (mc MonitorConfigAttributes) CustomHeader() terra.ListValue[CustomHeaderAttributes] {
	return terra.ReferenceAsList[CustomHeaderAttributes](mc.ref.Append("custom_header"))
}

type CustomHeaderAttributes struct {
	ref terra.Reference
}

func (ch CustomHeaderAttributes) InternalRef() (terra.Reference, error) {
	return ch.ref, nil
}

func (ch CustomHeaderAttributes) InternalWithRef(ref terra.Reference) CustomHeaderAttributes {
	return CustomHeaderAttributes{ref: ref}
}

func (ch CustomHeaderAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ch.ref.InternalTokens()
}

func (ch CustomHeaderAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ch.ref.Append("name"))
}

func (ch CustomHeaderAttributes) Value() terra.StringValue {
	return terra.ReferenceAsString(ch.ref.Append("value"))
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

func (t TimeoutsAttributes) Read() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("read"))
}

type DnsConfigState struct {
	RelativeName string  `json:"relative_name"`
	Ttl          float64 `json:"ttl"`
}

type MonitorConfigState struct {
	ExpectedStatusCodeRanges  []string            `json:"expected_status_code_ranges"`
	IntervalInSeconds         float64             `json:"interval_in_seconds"`
	Path                      string              `json:"path"`
	Port                      float64             `json:"port"`
	Protocol                  string              `json:"protocol"`
	TimeoutInSeconds          float64             `json:"timeout_in_seconds"`
	ToleratedNumberOfFailures float64             `json:"tolerated_number_of_failures"`
	CustomHeader              []CustomHeaderState `json:"custom_header"`
}

type CustomHeaderState struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type TimeoutsState struct {
	Read string `json:"read"`
}