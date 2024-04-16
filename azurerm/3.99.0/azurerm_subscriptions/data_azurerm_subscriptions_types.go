// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_subscriptions

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type DataTimeouts struct {
	// Read: string, optional
	Read terra.StringValue `hcl:"read,attr"`
}

type DataSubscriptionsAttributes struct {
	ref terra.Reference
}

func (s DataSubscriptionsAttributes) InternalRef() (terra.Reference, error) {
	return s.ref, nil
}

func (s DataSubscriptionsAttributes) InternalWithRef(ref terra.Reference) DataSubscriptionsAttributes {
	return DataSubscriptionsAttributes{ref: ref}
}

func (s DataSubscriptionsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return s.ref.InternalTokens()
}

func (s DataSubscriptionsAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("display_name"))
}

func (s DataSubscriptionsAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("id"))
}

func (s DataSubscriptionsAttributes) LocationPlacementId() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("location_placement_id"))
}

func (s DataSubscriptionsAttributes) QuotaId() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("quota_id"))
}

func (s DataSubscriptionsAttributes) SpendingLimit() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("spending_limit"))
}

func (s DataSubscriptionsAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("state"))
}

func (s DataSubscriptionsAttributes) SubscriptionId() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("subscription_id"))
}

func (s DataSubscriptionsAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](s.ref.Append("tags"))
}

func (s DataSubscriptionsAttributes) TenantId() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("tenant_id"))
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

type DataSubscriptionsState struct {
	DisplayName         string            `json:"display_name"`
	Id                  string            `json:"id"`
	LocationPlacementId string            `json:"location_placement_id"`
	QuotaId             string            `json:"quota_id"`
	SpendingLimit       string            `json:"spending_limit"`
	State               string            `json:"state"`
	SubscriptionId      string            `json:"subscription_id"`
	Tags                map[string]string `json:"tags"`
	TenantId            string            `json:"tenant_id"`
}

type DataTimeoutsState struct {
	Read string `json:"read"`
}
