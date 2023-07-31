// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package sesv2configurationset

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type DeliveryOptions struct {
	// SendingPoolName: string, optional
	SendingPoolName terra.StringValue `hcl:"sending_pool_name,attr"`
	// TlsPolicy: string, optional
	TlsPolicy terra.StringValue `hcl:"tls_policy,attr"`
}

type ReputationOptions struct {
	// ReputationMetricsEnabled: bool, optional
	ReputationMetricsEnabled terra.BoolValue `hcl:"reputation_metrics_enabled,attr"`
}

type SendingOptions struct {
	// SendingEnabled: bool, optional
	SendingEnabled terra.BoolValue `hcl:"sending_enabled,attr"`
}

type SuppressionOptions struct {
	// SuppressedReasons: list of string, optional
	SuppressedReasons terra.ListValue[terra.StringValue] `hcl:"suppressed_reasons,attr"`
}

type TrackingOptions struct {
	// CustomRedirectDomain: string, required
	CustomRedirectDomain terra.StringValue `hcl:"custom_redirect_domain,attr" validate:"required"`
}

type VdmOptions struct {
	// DashboardOptions: optional
	DashboardOptions *DashboardOptions `hcl:"dashboard_options,block"`
	// GuardianOptions: optional
	GuardianOptions *GuardianOptions `hcl:"guardian_options,block"`
}

type DashboardOptions struct {
	// EngagementMetrics: string, optional
	EngagementMetrics terra.StringValue `hcl:"engagement_metrics,attr"`
}

type GuardianOptions struct {
	// OptimizedSharedDelivery: string, optional
	OptimizedSharedDelivery terra.StringValue `hcl:"optimized_shared_delivery,attr"`
}

type DeliveryOptionsAttributes struct {
	ref terra.Reference
}

func (do DeliveryOptionsAttributes) InternalRef() (terra.Reference, error) {
	return do.ref, nil
}

func (do DeliveryOptionsAttributes) InternalWithRef(ref terra.Reference) DeliveryOptionsAttributes {
	return DeliveryOptionsAttributes{ref: ref}
}

func (do DeliveryOptionsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return do.ref.InternalTokens()
}

func (do DeliveryOptionsAttributes) SendingPoolName() terra.StringValue {
	return terra.ReferenceAsString(do.ref.Append("sending_pool_name"))
}

func (do DeliveryOptionsAttributes) TlsPolicy() terra.StringValue {
	return terra.ReferenceAsString(do.ref.Append("tls_policy"))
}

type ReputationOptionsAttributes struct {
	ref terra.Reference
}

func (ro ReputationOptionsAttributes) InternalRef() (terra.Reference, error) {
	return ro.ref, nil
}

func (ro ReputationOptionsAttributes) InternalWithRef(ref terra.Reference) ReputationOptionsAttributes {
	return ReputationOptionsAttributes{ref: ref}
}

func (ro ReputationOptionsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ro.ref.InternalTokens()
}

func (ro ReputationOptionsAttributes) LastFreshStart() terra.StringValue {
	return terra.ReferenceAsString(ro.ref.Append("last_fresh_start"))
}

func (ro ReputationOptionsAttributes) ReputationMetricsEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(ro.ref.Append("reputation_metrics_enabled"))
}

type SendingOptionsAttributes struct {
	ref terra.Reference
}

func (so SendingOptionsAttributes) InternalRef() (terra.Reference, error) {
	return so.ref, nil
}

func (so SendingOptionsAttributes) InternalWithRef(ref terra.Reference) SendingOptionsAttributes {
	return SendingOptionsAttributes{ref: ref}
}

func (so SendingOptionsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return so.ref.InternalTokens()
}

func (so SendingOptionsAttributes) SendingEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(so.ref.Append("sending_enabled"))
}

type SuppressionOptionsAttributes struct {
	ref terra.Reference
}

func (so SuppressionOptionsAttributes) InternalRef() (terra.Reference, error) {
	return so.ref, nil
}

func (so SuppressionOptionsAttributes) InternalWithRef(ref terra.Reference) SuppressionOptionsAttributes {
	return SuppressionOptionsAttributes{ref: ref}
}

func (so SuppressionOptionsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return so.ref.InternalTokens()
}

func (so SuppressionOptionsAttributes) SuppressedReasons() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](so.ref.Append("suppressed_reasons"))
}

type TrackingOptionsAttributes struct {
	ref terra.Reference
}

func (to TrackingOptionsAttributes) InternalRef() (terra.Reference, error) {
	return to.ref, nil
}

func (to TrackingOptionsAttributes) InternalWithRef(ref terra.Reference) TrackingOptionsAttributes {
	return TrackingOptionsAttributes{ref: ref}
}

func (to TrackingOptionsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return to.ref.InternalTokens()
}

func (to TrackingOptionsAttributes) CustomRedirectDomain() terra.StringValue {
	return terra.ReferenceAsString(to.ref.Append("custom_redirect_domain"))
}

type VdmOptionsAttributes struct {
	ref terra.Reference
}

func (vo VdmOptionsAttributes) InternalRef() (terra.Reference, error) {
	return vo.ref, nil
}

func (vo VdmOptionsAttributes) InternalWithRef(ref terra.Reference) VdmOptionsAttributes {
	return VdmOptionsAttributes{ref: ref}
}

func (vo VdmOptionsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return vo.ref.InternalTokens()
}

func (vo VdmOptionsAttributes) DashboardOptions() terra.ListValue[DashboardOptionsAttributes] {
	return terra.ReferenceAsList[DashboardOptionsAttributes](vo.ref.Append("dashboard_options"))
}

func (vo VdmOptionsAttributes) GuardianOptions() terra.ListValue[GuardianOptionsAttributes] {
	return terra.ReferenceAsList[GuardianOptionsAttributes](vo.ref.Append("guardian_options"))
}

type DashboardOptionsAttributes struct {
	ref terra.Reference
}

func (do DashboardOptionsAttributes) InternalRef() (terra.Reference, error) {
	return do.ref, nil
}

func (do DashboardOptionsAttributes) InternalWithRef(ref terra.Reference) DashboardOptionsAttributes {
	return DashboardOptionsAttributes{ref: ref}
}

func (do DashboardOptionsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return do.ref.InternalTokens()
}

func (do DashboardOptionsAttributes) EngagementMetrics() terra.StringValue {
	return terra.ReferenceAsString(do.ref.Append("engagement_metrics"))
}

type GuardianOptionsAttributes struct {
	ref terra.Reference
}

func (_go GuardianOptionsAttributes) InternalRef() (terra.Reference, error) {
	return _go.ref, nil
}

func (_go GuardianOptionsAttributes) InternalWithRef(ref terra.Reference) GuardianOptionsAttributes {
	return GuardianOptionsAttributes{ref: ref}
}

func (_go GuardianOptionsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return _go.ref.InternalTokens()
}

func (_go GuardianOptionsAttributes) OptimizedSharedDelivery() terra.StringValue {
	return terra.ReferenceAsString(_go.ref.Append("optimized_shared_delivery"))
}

type DeliveryOptionsState struct {
	SendingPoolName string `json:"sending_pool_name"`
	TlsPolicy       string `json:"tls_policy"`
}

type ReputationOptionsState struct {
	LastFreshStart           string `json:"last_fresh_start"`
	ReputationMetricsEnabled bool   `json:"reputation_metrics_enabled"`
}

type SendingOptionsState struct {
	SendingEnabled bool `json:"sending_enabled"`
}

type SuppressionOptionsState struct {
	SuppressedReasons []string `json:"suppressed_reasons"`
}

type TrackingOptionsState struct {
	CustomRedirectDomain string `json:"custom_redirect_domain"`
}

type VdmOptionsState struct {
	DashboardOptions []DashboardOptionsState `json:"dashboard_options"`
	GuardianOptions  []GuardianOptionsState  `json:"guardian_options"`
}

type DashboardOptionsState struct {
	EngagementMetrics string `json:"engagement_metrics"`
}

type GuardianOptionsState struct {
	OptimizedSharedDelivery string `json:"optimized_shared_delivery"`
}
