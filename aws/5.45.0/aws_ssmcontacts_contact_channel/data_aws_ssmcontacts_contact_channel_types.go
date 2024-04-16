// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_ssmcontacts_contact_channel

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type DataDeliveryAddressAttributes struct {
	ref terra.Reference
}

func (da DataDeliveryAddressAttributes) InternalRef() (terra.Reference, error) {
	return da.ref, nil
}

func (da DataDeliveryAddressAttributes) InternalWithRef(ref terra.Reference) DataDeliveryAddressAttributes {
	return DataDeliveryAddressAttributes{ref: ref}
}

func (da DataDeliveryAddressAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return da.ref.InternalTokens()
}

func (da DataDeliveryAddressAttributes) SimpleAddress() terra.StringValue {
	return terra.ReferenceAsString(da.ref.Append("simple_address"))
}

type DataDeliveryAddressState struct {
	SimpleAddress string `json:"simple_address"`
}
