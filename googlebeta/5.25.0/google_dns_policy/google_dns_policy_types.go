// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_dns_policy

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type AlternativeNameServerConfig struct {
	// AlternativeNameServerConfigTargetNameServers: min=1
	TargetNameServers []AlternativeNameServerConfigTargetNameServers `hcl:"target_name_servers,block" validate:"min=1"`
}

type AlternativeNameServerConfigTargetNameServers struct {
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

func (ansc AlternativeNameServerConfigAttributes) TargetNameServers() terra.SetValue[AlternativeNameServerConfigTargetNameServersAttributes] {
	return terra.ReferenceAsSet[AlternativeNameServerConfigTargetNameServersAttributes](ansc.ref.Append("target_name_servers"))
}

type AlternativeNameServerConfigTargetNameServersAttributes struct {
	ref terra.Reference
}

func (tns AlternativeNameServerConfigTargetNameServersAttributes) InternalRef() (terra.Reference, error) {
	return tns.ref, nil
}

func (tns AlternativeNameServerConfigTargetNameServersAttributes) InternalWithRef(ref terra.Reference) AlternativeNameServerConfigTargetNameServersAttributes {
	return AlternativeNameServerConfigTargetNameServersAttributes{ref: ref}
}

func (tns AlternativeNameServerConfigTargetNameServersAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return tns.ref.InternalTokens()
}

func (tns AlternativeNameServerConfigTargetNameServersAttributes) ForwardingPath() terra.StringValue {
	return terra.ReferenceAsString(tns.ref.Append("forwarding_path"))
}

func (tns AlternativeNameServerConfigTargetNameServersAttributes) Ipv4Address() terra.StringValue {
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
	TargetNameServers []AlternativeNameServerConfigTargetNameServersState `json:"target_name_servers"`
}

type AlternativeNameServerConfigTargetNameServersState struct {
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
