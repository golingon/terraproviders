// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_sentinel_alert_rule_template

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type DataTimeouts struct {
	// Read: string, optional
	Read terra.StringValue `hcl:"read,attr"`
}

type DataNrtTemplateAttributes struct {
	ref terra.Reference
}

func (nt DataNrtTemplateAttributes) InternalRef() (terra.Reference, error) {
	return nt.ref, nil
}

func (nt DataNrtTemplateAttributes) InternalWithRef(ref terra.Reference) DataNrtTemplateAttributes {
	return DataNrtTemplateAttributes{ref: ref}
}

func (nt DataNrtTemplateAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return nt.ref.InternalTokens()
}

func (nt DataNrtTemplateAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(nt.ref.Append("description"))
}

func (nt DataNrtTemplateAttributes) Query() terra.StringValue {
	return terra.ReferenceAsString(nt.ref.Append("query"))
}

func (nt DataNrtTemplateAttributes) Severity() terra.StringValue {
	return terra.ReferenceAsString(nt.ref.Append("severity"))
}

func (nt DataNrtTemplateAttributes) Tactics() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](nt.ref.Append("tactics"))
}

type DataScheduledTemplateAttributes struct {
	ref terra.Reference
}

func (st DataScheduledTemplateAttributes) InternalRef() (terra.Reference, error) {
	return st.ref, nil
}

func (st DataScheduledTemplateAttributes) InternalWithRef(ref terra.Reference) DataScheduledTemplateAttributes {
	return DataScheduledTemplateAttributes{ref: ref}
}

func (st DataScheduledTemplateAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return st.ref.InternalTokens()
}

func (st DataScheduledTemplateAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(st.ref.Append("description"))
}

func (st DataScheduledTemplateAttributes) Query() terra.StringValue {
	return terra.ReferenceAsString(st.ref.Append("query"))
}

func (st DataScheduledTemplateAttributes) QueryFrequency() terra.StringValue {
	return terra.ReferenceAsString(st.ref.Append("query_frequency"))
}

func (st DataScheduledTemplateAttributes) QueryPeriod() terra.StringValue {
	return terra.ReferenceAsString(st.ref.Append("query_period"))
}

func (st DataScheduledTemplateAttributes) Severity() terra.StringValue {
	return terra.ReferenceAsString(st.ref.Append("severity"))
}

func (st DataScheduledTemplateAttributes) Tactics() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](st.ref.Append("tactics"))
}

func (st DataScheduledTemplateAttributes) TriggerOperator() terra.StringValue {
	return terra.ReferenceAsString(st.ref.Append("trigger_operator"))
}

func (st DataScheduledTemplateAttributes) TriggerThreshold() terra.NumberValue {
	return terra.ReferenceAsNumber(st.ref.Append("trigger_threshold"))
}

type DataSecurityIncidentTemplateAttributes struct {
	ref terra.Reference
}

func (sit DataSecurityIncidentTemplateAttributes) InternalRef() (terra.Reference, error) {
	return sit.ref, nil
}

func (sit DataSecurityIncidentTemplateAttributes) InternalWithRef(ref terra.Reference) DataSecurityIncidentTemplateAttributes {
	return DataSecurityIncidentTemplateAttributes{ref: ref}
}

func (sit DataSecurityIncidentTemplateAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return sit.ref.InternalTokens()
}

func (sit DataSecurityIncidentTemplateAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(sit.ref.Append("description"))
}

func (sit DataSecurityIncidentTemplateAttributes) ProductFilter() terra.StringValue {
	return terra.ReferenceAsString(sit.ref.Append("product_filter"))
}

type DataTimeoutsAttributes struct {
	ref terra.Reference
}

func (t DataTimeoutsAttributes) InternalRef() (terra.Reference, error) {
	return t.ref, nil
}

func (t DataTimeoutsAttributes) InternalWithRef(ref terra.Reference) DataTimeoutsAttributes {
	return DataTimeoutsAttributes{ref: ref}
}

func (t DataTimeoutsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return t.ref.InternalTokens()
}

func (t DataTimeoutsAttributes) Read() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("read"))
}

type DataNrtTemplateState struct {
	Description string   `json:"description"`
	Query       string   `json:"query"`
	Severity    string   `json:"severity"`
	Tactics     []string `json:"tactics"`
}

type DataScheduledTemplateState struct {
	Description      string   `json:"description"`
	Query            string   `json:"query"`
	QueryFrequency   string   `json:"query_frequency"`
	QueryPeriod      string   `json:"query_period"`
	Severity         string   `json:"severity"`
	Tactics          []string `json:"tactics"`
	TriggerOperator  string   `json:"trigger_operator"`
	TriggerThreshold float64  `json:"trigger_threshold"`
}

type DataSecurityIncidentTemplateState struct {
	Description   string `json:"description"`
	ProductFilter string `json:"product_filter"`
}

type DataTimeoutsState struct {
	Read string `json:"read"`
}
