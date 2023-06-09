// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package dataapimanagementapi

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type SubscriptionKeyParameterNames struct{}

type Timeouts struct {
	// Read: string, optional
	Read terra.StringValue `hcl:"read,attr"`
}

type SubscriptionKeyParameterNamesAttributes struct {
	ref terra.Reference
}

func (skpn SubscriptionKeyParameterNamesAttributes) InternalRef() (terra.Reference, error) {
	return skpn.ref, nil
}

func (skpn SubscriptionKeyParameterNamesAttributes) InternalWithRef(ref terra.Reference) SubscriptionKeyParameterNamesAttributes {
	return SubscriptionKeyParameterNamesAttributes{ref: ref}
}

func (skpn SubscriptionKeyParameterNamesAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return skpn.ref.InternalTokens()
}

func (skpn SubscriptionKeyParameterNamesAttributes) Header() terra.StringValue {
	return terra.ReferenceAsString(skpn.ref.Append("header"))
}

func (skpn SubscriptionKeyParameterNamesAttributes) Query() terra.StringValue {
	return terra.ReferenceAsString(skpn.ref.Append("query"))
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

type SubscriptionKeyParameterNamesState struct {
	Header string `json:"header"`
	Query  string `json:"query"`
}

type TimeoutsState struct {
	Read string `json:"read"`
}
