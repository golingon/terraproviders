// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package dnspolicy

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type AlternativeNameServerConfig struct {
	// TargetNameServers: min=1
	TargetNameServers []TargetNameServers `hcl:"target_name_servers,block" validate:"min=1"`
}

type TargetNameServers struct {
	// ForwardingPath: string, optional
	ForwardingPath terra.StringValue `hcl:"forwarding_path,attr"`
	// Ipv4Address: string, required
	Ipv4Address terra.StringValue `hcl:"ipv4_address,attr" validate:"required"`
}

type Networks struct {
	// NetworkUrl: string, required
	NetworkUrl terra.StringValue `hcl:"network_url,attr" validate:"required"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type AlternativeNameServerConfigAttributes struct {
	ref terra.Reference
}

func (ansc AlternativeNameServerConfigAttributes) InternalRef() (terra.Reference, error) {
	return ansc.ref, nil
}

func (ansc AlternativeNameServerConfigAttributes) InternalWithRef(ref terra.Reference) AlternativeNameServerConfigAttributes {
	return AlternativeNameServerConfigAttributes{ref: ref}
}

func (ansc AlternativeNameServerConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ansc.ref.InternalTokens()
}

func (ansc AlternativeNameServerConfigAttributes) TargetNameServers() terra.SetValue[TargetNameServersAttributes] {
	return terra.ReferenceAsSet[TargetNameServersAttributes](ansc.ref.Append("target_name_servers"))
}

type TargetNameServersAttributes struct {
	ref terra.Reference
}

func (tns TargetNameServersAttributes) InternalRef() (terra.Reference, error) {
	return tns.ref, nil
}

func (tns TargetNameServersAttributes) InternalWithRef(ref terra.Reference) TargetNameServersAttributes {
	return TargetNameServersAttributes{ref: ref}
}

func (tns TargetNameServersAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return tns.ref.InternalTokens()
}

func (tns TargetNameServersAttributes) ForwardingPath() terra.StringValue {
	return terra.ReferenceAsString(tns.ref.Append("forwarding_path"))
}

func (tns TargetNameServersAttributes) Ipv4Address() terra.StringValue {
	return terra.ReferenceAsString(tns.ref.Append("ipv4_address"))
}

type NetworksAttributes struct {
	ref terra.Reference
}

func (n NetworksAttributes) InternalRef() (terra.Reference, error) {
	return n.ref, nil
}

func (n NetworksAttributes) InternalWithRef(ref terra.Reference) NetworksAttributes {
	return NetworksAttributes{ref: ref}
}

func (n NetworksAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return n.ref.InternalTokens()
}

func (n NetworksAttributes) NetworkUrl() terra.StringValue {
	return terra.ReferenceAsString(n.ref.Append("network_url"))
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

func (t TimeoutsAttributes) Update() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("update"))
}

type AlternativeNameServerConfigState struct {
	TargetNameServers []TargetNameServersState `json:"target_name_servers"`
}

type TargetNameServersState struct {
	ForwardingPath string `json:"forwarding_path"`
	Ipv4Address    string `json:"ipv4_address"`
}

type NetworksState struct {
	NetworkUrl string `json:"network_url"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}
