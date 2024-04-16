// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_servicebus_subscription

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type ClientScopedSubscription struct {
	// ClientId: string, optional
	ClientId terra.StringValue `hcl:"client_id,attr"`
	// IsClientScopedSubscriptionShareable: bool, optional
	IsClientScopedSubscriptionShareable terra.BoolValue `hcl:"is_client_scoped_subscription_shareable,attr"`
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

type ClientScopedSubscriptionAttributes struct {
	ref terra.Reference
}

func (css ClientScopedSubscriptionAttributes) InternalRef() (terra.Reference, error) {
	return css.ref, nil
}

func (css ClientScopedSubscriptionAttributes) InternalWithRef(ref terra.Reference) ClientScopedSubscriptionAttributes {
	return ClientScopedSubscriptionAttributes{ref: ref}
}

func (css ClientScopedSubscriptionAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return css.ref.InternalTokens()
}

func (css ClientScopedSubscriptionAttributes) ClientId() terra.StringValue {
	return terra.ReferenceAsString(css.ref.Append("client_id"))
}

func (css ClientScopedSubscriptionAttributes) IsClientScopedSubscriptionDurable() terra.BoolValue {
	return terra.ReferenceAsBool(css.ref.Append("is_client_scoped_subscription_durable"))
}

func (css ClientScopedSubscriptionAttributes) IsClientScopedSubscriptionShareable() terra.BoolValue {
	return terra.ReferenceAsBool(css.ref.Append("is_client_scoped_subscription_shareable"))
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

type ClientScopedSubscriptionState struct {
	ClientId                            string `json:"client_id"`
	IsClientScopedSubscriptionDurable   bool   `json:"is_client_scoped_subscription_durable"`
	IsClientScopedSubscriptionShareable bool   `json:"is_client_scoped_subscription_shareable"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
}
