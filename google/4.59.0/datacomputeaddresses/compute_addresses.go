// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package datacomputeaddresses

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Addresses struct{}

type AddressesAttributes struct {
	ref terra.Reference
}

func (a AddressesAttributes) InternalRef() terra.Reference {
	return a.ref
}

func (a AddressesAttributes) InternalWithRef(ref terra.Reference) AddressesAttributes {
	return AddressesAttributes{ref: ref}
}

func (a AddressesAttributes) InternalTokens() hclwrite.Tokens {
	return a.ref.InternalTokens()
}

func (a AddressesAttributes) Address() terra.StringValue {
	return terra.ReferenceString(a.ref.Append("address"))
}

func (a AddressesAttributes) AddressType() terra.StringValue {
	return terra.ReferenceString(a.ref.Append("address_type"))
}

func (a AddressesAttributes) Description() terra.StringValue {
	return terra.ReferenceString(a.ref.Append("description"))
}

func (a AddressesAttributes) Name() terra.StringValue {
	return terra.ReferenceString(a.ref.Append("name"))
}

func (a AddressesAttributes) Region() terra.StringValue {
	return terra.ReferenceString(a.ref.Append("region"))
}

func (a AddressesAttributes) SelfLink() terra.StringValue {
	return terra.ReferenceString(a.ref.Append("self_link"))
}

func (a AddressesAttributes) Status() terra.StringValue {
	return terra.ReferenceString(a.ref.Append("status"))
}

type AddressesState struct {
	Address     string `json:"address"`
	AddressType string `json:"address_type"`
	Description string `json:"description"`
	Name        string `json:"name"`
	Region      string `json:"region"`
	SelfLink    string `json:"self_link"`
	Status      string `json:"status"`
}
