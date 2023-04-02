// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package sesconfigurationset

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type DeliveryOptions struct {
	// TlsPolicy: string, optional
	TlsPolicy terra.StringValue `hcl:"tls_policy,attr"`
}

type TrackingOptions struct {
	// CustomRedirectDomain: string, optional
	CustomRedirectDomain terra.StringValue `hcl:"custom_redirect_domain,attr"`
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

func (do DeliveryOptionsAttributes) InternalTokens() hclwrite.Tokens {
	return do.ref.InternalTokens()
}

func (do DeliveryOptionsAttributes) TlsPolicy() terra.StringValue {
	return terra.ReferenceAsString(do.ref.Append("tls_policy"))
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

func (to TrackingOptionsAttributes) InternalTokens() hclwrite.Tokens {
	return to.ref.InternalTokens()
}

func (to TrackingOptionsAttributes) CustomRedirectDomain() terra.StringValue {
	return terra.ReferenceAsString(to.ref.Append("custom_redirect_domain"))
}

type DeliveryOptionsState struct {
	TlsPolicy string `json:"tls_policy"`
}

type TrackingOptionsState struct {
	CustomRedirectDomain string `json:"custom_redirect_domain"`
}
