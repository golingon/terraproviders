// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package configdeliverychannel

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type SnapshotDeliveryProperties struct {
	// DeliveryFrequency: string, optional
	DeliveryFrequency terra.StringValue `hcl:"delivery_frequency,attr"`
}

type SnapshotDeliveryPropertiesAttributes struct {
	ref terra.Reference
}

func (sdp SnapshotDeliveryPropertiesAttributes) InternalRef() terra.Reference {
	return sdp.ref
}

func (sdp SnapshotDeliveryPropertiesAttributes) InternalWithRef(ref terra.Reference) SnapshotDeliveryPropertiesAttributes {
	return SnapshotDeliveryPropertiesAttributes{ref: ref}
}

func (sdp SnapshotDeliveryPropertiesAttributes) InternalTokens() hclwrite.Tokens {
	return sdp.ref.InternalTokens()
}

func (sdp SnapshotDeliveryPropertiesAttributes) DeliveryFrequency() terra.StringValue {
	return terra.ReferenceAsString(sdp.ref.Append("delivery_frequency"))
}

type SnapshotDeliveryPropertiesState struct {
	DeliveryFrequency string `json:"delivery_frequency"`
}