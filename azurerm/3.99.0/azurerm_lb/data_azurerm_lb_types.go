// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_lb

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type DataTimeouts struct {
	// Read: string, optional
	Read terra.StringValue `hcl:"read,attr"`
}

type DataFrontendIpConfigurationAttributes struct {
	ref terra.Reference
}

func (fic DataFrontendIpConfigurationAttributes) InternalRef() (terra.Reference, error) {
	return fic.ref, nil
}

func (fic DataFrontendIpConfigurationAttributes) InternalWithRef(ref terra.Reference) DataFrontendIpConfigurationAttributes {
	return DataFrontendIpConfigurationAttributes{ref: ref}
}

func (fic DataFrontendIpConfigurationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return fic.ref.InternalTokens()
}

func (fic DataFrontendIpConfigurationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(fic.ref.Append("id"))
}

func (fic DataFrontendIpConfigurationAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(fic.ref.Append("name"))
}

func (fic DataFrontendIpConfigurationAttributes) PrivateIpAddress() terra.StringValue {
	return terra.ReferenceAsString(fic.ref.Append("private_ip_address"))
}

func (fic DataFrontendIpConfigurationAttributes) PrivateIpAddressAllocation() terra.StringValue {
	return terra.ReferenceAsString(fic.ref.Append("private_ip_address_allocation"))
}

func (fic DataFrontendIpConfigurationAttributes) PrivateIpAddressVersion() terra.StringValue {
	return terra.ReferenceAsString(fic.ref.Append("private_ip_address_version"))
}

func (fic DataFrontendIpConfigurationAttributes) PublicIpAddressId() terra.StringValue {
	return terra.ReferenceAsString(fic.ref.Append("public_ip_address_id"))
}

func (fic DataFrontendIpConfigurationAttributes) SubnetId() terra.StringValue {
	return terra.ReferenceAsString(fic.ref.Append("subnet_id"))
}

func (fic DataFrontendIpConfigurationAttributes) Zones() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](fic.ref.Append("zones"))
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

type DataFrontendIpConfigurationState struct {
	Id                         string   `json:"id"`
	Name                       string   `json:"name"`
	PrivateIpAddress           string   `json:"private_ip_address"`
	PrivateIpAddressAllocation string   `json:"private_ip_address_allocation"`
	PrivateIpAddressVersion    string   `json:"private_ip_address_version"`
	PublicIpAddressId          string   `json:"public_ip_address_id"`
	SubnetId                   string   `json:"subnet_id"`
	Zones                      []string `json:"zones"`
}

type DataTimeoutsState struct {
	Read string `json:"read"`
}
