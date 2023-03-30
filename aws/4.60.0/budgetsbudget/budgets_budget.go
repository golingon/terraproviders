// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package budgetsbudget

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type AutoAdjustData struct {
	// AutoAdjustType: string, required
	AutoAdjustType terra.StringValue `hcl:"auto_adjust_type,attr" validate:"required"`
	// HistoricalOptions: optional
	HistoricalOptions *HistoricalOptions `hcl:"historical_options,block"`
}

type HistoricalOptions struct {
	// BudgetAdjustmentPeriod: number, required
	BudgetAdjustmentPeriod terra.NumberValue `hcl:"budget_adjustment_period,attr" validate:"required"`
}

type CostFilter struct {
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Values: list of string, required
	Values terra.ListValue[terra.StringValue] `hcl:"values,attr" validate:"required"`
}

type CostTypes struct {
	// IncludeCredit: bool, optional
	IncludeCredit terra.BoolValue `hcl:"include_credit,attr"`
	// IncludeDiscount: bool, optional
	IncludeDiscount terra.BoolValue `hcl:"include_discount,attr"`
	// IncludeOtherSubscription: bool, optional
	IncludeOtherSubscription terra.BoolValue `hcl:"include_other_subscription,attr"`
	// IncludeRecurring: bool, optional
	IncludeRecurring terra.BoolValue `hcl:"include_recurring,attr"`
	// IncludeRefund: bool, optional
	IncludeRefund terra.BoolValue `hcl:"include_refund,attr"`
	// IncludeSubscription: bool, optional
	IncludeSubscription terra.BoolValue `hcl:"include_subscription,attr"`
	// IncludeSupport: bool, optional
	IncludeSupport terra.BoolValue `hcl:"include_support,attr"`
	// IncludeTax: bool, optional
	IncludeTax terra.BoolValue `hcl:"include_tax,attr"`
	// IncludeUpfront: bool, optional
	IncludeUpfront terra.BoolValue `hcl:"include_upfront,attr"`
	// UseAmortized: bool, optional
	UseAmortized terra.BoolValue `hcl:"use_amortized,attr"`
	// UseBlended: bool, optional
	UseBlended terra.BoolValue `hcl:"use_blended,attr"`
}

type Notification struct {
	// ComparisonOperator: string, required
	ComparisonOperator terra.StringValue `hcl:"comparison_operator,attr" validate:"required"`
	// NotificationType: string, required
	NotificationType terra.StringValue `hcl:"notification_type,attr" validate:"required"`
	// SubscriberEmailAddresses: set of string, optional
	SubscriberEmailAddresses terra.SetValue[terra.StringValue] `hcl:"subscriber_email_addresses,attr"`
	// SubscriberSnsTopicArns: set of string, optional
	SubscriberSnsTopicArns terra.SetValue[terra.StringValue] `hcl:"subscriber_sns_topic_arns,attr"`
	// Threshold: number, required
	Threshold terra.NumberValue `hcl:"threshold,attr" validate:"required"`
	// ThresholdType: string, required
	ThresholdType terra.StringValue `hcl:"threshold_type,attr" validate:"required"`
}

type PlannedLimit struct {
	// Amount: string, required
	Amount terra.StringValue `hcl:"amount,attr" validate:"required"`
	// StartTime: string, required
	StartTime terra.StringValue `hcl:"start_time,attr" validate:"required"`
	// Unit: string, required
	Unit terra.StringValue `hcl:"unit,attr" validate:"required"`
}

type AutoAdjustDataAttributes struct {
	ref terra.Reference
}

func (aad AutoAdjustDataAttributes) InternalRef() terra.Reference {
	return aad.ref
}

func (aad AutoAdjustDataAttributes) InternalWithRef(ref terra.Reference) AutoAdjustDataAttributes {
	return AutoAdjustDataAttributes{ref: ref}
}

func (aad AutoAdjustDataAttributes) InternalTokens() hclwrite.Tokens {
	return aad.ref.InternalTokens()
}

func (aad AutoAdjustDataAttributes) AutoAdjustType() terra.StringValue {
	return terra.ReferenceString(aad.ref.Append("auto_adjust_type"))
}

func (aad AutoAdjustDataAttributes) LastAutoAdjustTime() terra.StringValue {
	return terra.ReferenceString(aad.ref.Append("last_auto_adjust_time"))
}

func (aad AutoAdjustDataAttributes) HistoricalOptions() terra.ListValue[HistoricalOptionsAttributes] {
	return terra.ReferenceList[HistoricalOptionsAttributes](aad.ref.Append("historical_options"))
}

type HistoricalOptionsAttributes struct {
	ref terra.Reference
}

func (ho HistoricalOptionsAttributes) InternalRef() terra.Reference {
	return ho.ref
}

func (ho HistoricalOptionsAttributes) InternalWithRef(ref terra.Reference) HistoricalOptionsAttributes {
	return HistoricalOptionsAttributes{ref: ref}
}

func (ho HistoricalOptionsAttributes) InternalTokens() hclwrite.Tokens {
	return ho.ref.InternalTokens()
}

func (ho HistoricalOptionsAttributes) BudgetAdjustmentPeriod() terra.NumberValue {
	return terra.ReferenceNumber(ho.ref.Append("budget_adjustment_period"))
}

func (ho HistoricalOptionsAttributes) LookbackAvailablePeriods() terra.NumberValue {
	return terra.ReferenceNumber(ho.ref.Append("lookback_available_periods"))
}

type CostFilterAttributes struct {
	ref terra.Reference
}

func (cf CostFilterAttributes) InternalRef() terra.Reference {
	return cf.ref
}

func (cf CostFilterAttributes) InternalWithRef(ref terra.Reference) CostFilterAttributes {
	return CostFilterAttributes{ref: ref}
}

func (cf CostFilterAttributes) InternalTokens() hclwrite.Tokens {
	return cf.ref.InternalTokens()
}

func (cf CostFilterAttributes) Name() terra.StringValue {
	return terra.ReferenceString(cf.ref.Append("name"))
}

func (cf CostFilterAttributes) Values() terra.ListValue[terra.StringValue] {
	return terra.ReferenceList[terra.StringValue](cf.ref.Append("values"))
}

type CostTypesAttributes struct {
	ref terra.Reference
}

func (ct CostTypesAttributes) InternalRef() terra.Reference {
	return ct.ref
}

func (ct CostTypesAttributes) InternalWithRef(ref terra.Reference) CostTypesAttributes {
	return CostTypesAttributes{ref: ref}
}

func (ct CostTypesAttributes) InternalTokens() hclwrite.Tokens {
	return ct.ref.InternalTokens()
}

func (ct CostTypesAttributes) IncludeCredit() terra.BoolValue {
	return terra.ReferenceBool(ct.ref.Append("include_credit"))
}

func (ct CostTypesAttributes) IncludeDiscount() terra.BoolValue {
	return terra.ReferenceBool(ct.ref.Append("include_discount"))
}

func (ct CostTypesAttributes) IncludeOtherSubscription() terra.BoolValue {
	return terra.ReferenceBool(ct.ref.Append("include_other_subscription"))
}

func (ct CostTypesAttributes) IncludeRecurring() terra.BoolValue {
	return terra.ReferenceBool(ct.ref.Append("include_recurring"))
}

func (ct CostTypesAttributes) IncludeRefund() terra.BoolValue {
	return terra.ReferenceBool(ct.ref.Append("include_refund"))
}

func (ct CostTypesAttributes) IncludeSubscription() terra.BoolValue {
	return terra.ReferenceBool(ct.ref.Append("include_subscription"))
}

func (ct CostTypesAttributes) IncludeSupport() terra.BoolValue {
	return terra.ReferenceBool(ct.ref.Append("include_support"))
}

func (ct CostTypesAttributes) IncludeTax() terra.BoolValue {
	return terra.ReferenceBool(ct.ref.Append("include_tax"))
}

func (ct CostTypesAttributes) IncludeUpfront() terra.BoolValue {
	return terra.ReferenceBool(ct.ref.Append("include_upfront"))
}

func (ct CostTypesAttributes) UseAmortized() terra.BoolValue {
	return terra.ReferenceBool(ct.ref.Append("use_amortized"))
}

func (ct CostTypesAttributes) UseBlended() terra.BoolValue {
	return terra.ReferenceBool(ct.ref.Append("use_blended"))
}

type NotificationAttributes struct {
	ref terra.Reference
}

func (n NotificationAttributes) InternalRef() terra.Reference {
	return n.ref
}

func (n NotificationAttributes) InternalWithRef(ref terra.Reference) NotificationAttributes {
	return NotificationAttributes{ref: ref}
}

func (n NotificationAttributes) InternalTokens() hclwrite.Tokens {
	return n.ref.InternalTokens()
}

func (n NotificationAttributes) ComparisonOperator() terra.StringValue {
	return terra.ReferenceString(n.ref.Append("comparison_operator"))
}

func (n NotificationAttributes) NotificationType() terra.StringValue {
	return terra.ReferenceString(n.ref.Append("notification_type"))
}

func (n NotificationAttributes) SubscriberEmailAddresses() terra.SetValue[terra.StringValue] {
	return terra.ReferenceSet[terra.StringValue](n.ref.Append("subscriber_email_addresses"))
}

func (n NotificationAttributes) SubscriberSnsTopicArns() terra.SetValue[terra.StringValue] {
	return terra.ReferenceSet[terra.StringValue](n.ref.Append("subscriber_sns_topic_arns"))
}

func (n NotificationAttributes) Threshold() terra.NumberValue {
	return terra.ReferenceNumber(n.ref.Append("threshold"))
}

func (n NotificationAttributes) ThresholdType() terra.StringValue {
	return terra.ReferenceString(n.ref.Append("threshold_type"))
}

type PlannedLimitAttributes struct {
	ref terra.Reference
}

func (pl PlannedLimitAttributes) InternalRef() terra.Reference {
	return pl.ref
}

func (pl PlannedLimitAttributes) InternalWithRef(ref terra.Reference) PlannedLimitAttributes {
	return PlannedLimitAttributes{ref: ref}
}

func (pl PlannedLimitAttributes) InternalTokens() hclwrite.Tokens {
	return pl.ref.InternalTokens()
}

func (pl PlannedLimitAttributes) Amount() terra.StringValue {
	return terra.ReferenceString(pl.ref.Append("amount"))
}

func (pl PlannedLimitAttributes) StartTime() terra.StringValue {
	return terra.ReferenceString(pl.ref.Append("start_time"))
}

func (pl PlannedLimitAttributes) Unit() terra.StringValue {
	return terra.ReferenceString(pl.ref.Append("unit"))
}

type AutoAdjustDataState struct {
	AutoAdjustType     string                   `json:"auto_adjust_type"`
	LastAutoAdjustTime string                   `json:"last_auto_adjust_time"`
	HistoricalOptions  []HistoricalOptionsState `json:"historical_options"`
}

type HistoricalOptionsState struct {
	BudgetAdjustmentPeriod   float64 `json:"budget_adjustment_period"`
	LookbackAvailablePeriods float64 `json:"lookback_available_periods"`
}

type CostFilterState struct {
	Name   string   `json:"name"`
	Values []string `json:"values"`
}

type CostTypesState struct {
	IncludeCredit            bool `json:"include_credit"`
	IncludeDiscount          bool `json:"include_discount"`
	IncludeOtherSubscription bool `json:"include_other_subscription"`
	IncludeRecurring         bool `json:"include_recurring"`
	IncludeRefund            bool `json:"include_refund"`
	IncludeSubscription      bool `json:"include_subscription"`
	IncludeSupport           bool `json:"include_support"`
	IncludeTax               bool `json:"include_tax"`
	IncludeUpfront           bool `json:"include_upfront"`
	UseAmortized             bool `json:"use_amortized"`
	UseBlended               bool `json:"use_blended"`
}

type NotificationState struct {
	ComparisonOperator       string   `json:"comparison_operator"`
	NotificationType         string   `json:"notification_type"`
	SubscriberEmailAddresses []string `json:"subscriber_email_addresses"`
	SubscriberSnsTopicArns   []string `json:"subscriber_sns_topic_arns"`
	Threshold                float64  `json:"threshold"`
	ThresholdType            string   `json:"threshold_type"`
}

type PlannedLimitState struct {
	Amount    string `json:"amount"`
	StartTime string `json:"start_time"`
	Unit      string `json:"unit"`
}
