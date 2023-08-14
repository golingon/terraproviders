// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package paloaltonextgenerationfirewallvirtualhubpanorama

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Panorama struct{}

type DestinationNat struct {
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Protocol: string, required
	Protocol terra.StringValue `hcl:"protocol,attr" validate:"required"`
	// BackendConfig: optional
	BackendConfig *BackendConfig `hcl:"backend_config,block"`
	// FrontendConfig: optional
	FrontendConfig *FrontendConfig `hcl:"frontend_config,block"`
}

type BackendConfig struct {
	// Port: number, required
	Port terra.NumberValue `hcl:"port,attr" validate:"required"`
	// PublicIpAddress: string, required
	PublicIpAddress terra.StringValue `hcl:"public_ip_address,attr" validate:"required"`
}

type FrontendConfig struct {
	// Port: number, required
	Port terra.NumberValue `hcl:"port,attr" validate:"required"`
	// PublicIpAddressId: string, required
	PublicIpAddressId terra.StringValue `hcl:"public_ip_address_id,attr" validate:"required"`
}

type DnsSettings struct {
	// DnsServers: list of string, optional
	DnsServers terra.ListValue[terra.StringValue] `hcl:"dns_servers,attr"`
	// UseAzureDns: bool, optional
	UseAzureDns terra.BoolValue `hcl:"use_azure_dns,attr"`
}

type NetworkProfile struct {
	// EgressNatIpAddressIds: list of string, optional
	EgressNatIpAddressIds terra.ListValue[terra.StringValue] `hcl:"egress_nat_ip_address_ids,attr"`
	// NetworkVirtualApplianceId: string, required
	NetworkVirtualApplianceId terra.StringValue `hcl:"network_virtual_appliance_id,attr" validate:"required"`
	// PublicIpAddressIds: list of string, required
	PublicIpAddressIds terra.ListValue[terra.StringValue] `hcl:"public_ip_address_ids,attr" validate:"required"`
	// VirtualHubId: string, required
	VirtualHubId terra.StringValue `hcl:"virtual_hub_id,attr" validate:"required"`
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

type PanoramaAttributes struct {
	ref terra.Reference
}

func (p PanoramaAttributes) InternalRef() (terra.Reference, error) {
	return p.ref, nil
}

func (p PanoramaAttributes) InternalWithRef(ref terra.Reference) PanoramaAttributes {
	return PanoramaAttributes{ref: ref}
}

func (p PanoramaAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return p.ref.InternalTokens()
}

func (p PanoramaAttributes) DeviceGroupName() terra.StringValue {
	return terra.ReferenceAsString(p.ref.Append("device_group_name"))
}

func (p PanoramaAttributes) HostName() terra.StringValue {
	return terra.ReferenceAsString(p.ref.Append("host_name"))
}

func (p PanoramaAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(p.ref.Append("name"))
}

func (p PanoramaAttributes) PanoramaServer1() terra.StringValue {
	return terra.ReferenceAsString(p.ref.Append("panorama_server_1"))
}

func (p PanoramaAttributes) PanoramaServer2() terra.StringValue {
	return terra.ReferenceAsString(p.ref.Append("panorama_server_2"))
}

func (p PanoramaAttributes) TemplateName() terra.StringValue {
	return terra.ReferenceAsString(p.ref.Append("template_name"))
}

func (p PanoramaAttributes) VirtualMachineSshKey() terra.StringValue {
	return terra.ReferenceAsString(p.ref.Append("virtual_machine_ssh_key"))
}

type DestinationNatAttributes struct {
	ref terra.Reference
}

func (dn DestinationNatAttributes) InternalRef() (terra.Reference, error) {
	return dn.ref, nil
}

func (dn DestinationNatAttributes) InternalWithRef(ref terra.Reference) DestinationNatAttributes {
	return DestinationNatAttributes{ref: ref}
}

func (dn DestinationNatAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return dn.ref.InternalTokens()
}

func (dn DestinationNatAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(dn.ref.Append("name"))
}

func (dn DestinationNatAttributes) Protocol() terra.StringValue {
	return terra.ReferenceAsString(dn.ref.Append("protocol"))
}

func (dn DestinationNatAttributes) BackendConfig() terra.ListValue[BackendConfigAttributes] {
	return terra.ReferenceAsList[BackendConfigAttributes](dn.ref.Append("backend_config"))
}

func (dn DestinationNatAttributes) FrontendConfig() terra.ListValue[FrontendConfigAttributes] {
	return terra.ReferenceAsList[FrontendConfigAttributes](dn.ref.Append("frontend_config"))
}

type BackendConfigAttributes struct {
	ref terra.Reference
}

func (bc BackendConfigAttributes) InternalRef() (terra.Reference, error) {
	return bc.ref, nil
}

func (bc BackendConfigAttributes) InternalWithRef(ref terra.Reference) BackendConfigAttributes {
	return BackendConfigAttributes{ref: ref}
}

func (bc BackendConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return bc.ref.InternalTokens()
}

func (bc BackendConfigAttributes) Port() terra.NumberValue {
	return terra.ReferenceAsNumber(bc.ref.Append("port"))
}

func (bc BackendConfigAttributes) PublicIpAddress() terra.StringValue {
	return terra.ReferenceAsString(bc.ref.Append("public_ip_address"))
}

type FrontendConfigAttributes struct {
	ref terra.Reference
}

func (fc FrontendConfigAttributes) InternalRef() (terra.Reference, error) {
	return fc.ref, nil
}

func (fc FrontendConfigAttributes) InternalWithRef(ref terra.Reference) FrontendConfigAttributes {
	return FrontendConfigAttributes{ref: ref}
}

func (fc FrontendConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return fc.ref.InternalTokens()
}

func (fc FrontendConfigAttributes) Port() terra.NumberValue {
	return terra.ReferenceAsNumber(fc.ref.Append("port"))
}

func (fc FrontendConfigAttributes) PublicIpAddressId() terra.StringValue {
	return terra.ReferenceAsString(fc.ref.Append("public_ip_address_id"))
}

type DnsSettingsAttributes struct {
	ref terra.Reference
}

func (ds DnsSettingsAttributes) InternalRef() (terra.Reference, error) {
	return ds.ref, nil
}

func (ds DnsSettingsAttributes) InternalWithRef(ref terra.Reference) DnsSettingsAttributes {
	return DnsSettingsAttributes{ref: ref}
}

func (ds DnsSettingsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ds.ref.InternalTokens()
}

func (ds DnsSettingsAttributes) AzureDnsServers() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](ds.ref.Append("azure_dns_servers"))
}

func (ds DnsSettingsAttributes) DnsServers() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](ds.ref.Append("dns_servers"))
}

func (ds DnsSettingsAttributes) UseAzureDns() terra.BoolValue {
	return terra.ReferenceAsBool(ds.ref.Append("use_azure_dns"))
}

type NetworkProfileAttributes struct {
	ref terra.Reference
}

func (np NetworkProfileAttributes) InternalRef() (terra.Reference, error) {
	return np.ref, nil
}

func (np NetworkProfileAttributes) InternalWithRef(ref terra.Reference) NetworkProfileAttributes {
	return NetworkProfileAttributes{ref: ref}
}

func (np NetworkProfileAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return np.ref.InternalTokens()
}

func (np NetworkProfileAttributes) EgressNatIpAddressIds() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](np.ref.Append("egress_nat_ip_address_ids"))
}

func (np NetworkProfileAttributes) EgressNatIpAddresses() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](np.ref.Append("egress_nat_ip_addresses"))
}

func (np NetworkProfileAttributes) IpOfTrustForUserDefinedRoutes() terra.StringValue {
	return terra.ReferenceAsString(np.ref.Append("ip_of_trust_for_user_defined_routes"))
}

func (np NetworkProfileAttributes) NetworkVirtualApplianceId() terra.StringValue {
	return terra.ReferenceAsString(np.ref.Append("network_virtual_appliance_id"))
}

func (np NetworkProfileAttributes) PublicIpAddressIds() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](np.ref.Append("public_ip_address_ids"))
}

func (np NetworkProfileAttributes) PublicIpAddresses() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](np.ref.Append("public_ip_addresses"))
}

func (np NetworkProfileAttributes) TrustedSubnetId() terra.StringValue {
	return terra.ReferenceAsString(np.ref.Append("trusted_subnet_id"))
}

func (np NetworkProfileAttributes) UntrustedSubnetId() terra.StringValue {
	return terra.ReferenceAsString(np.ref.Append("untrusted_subnet_id"))
}

func (np NetworkProfileAttributes) VirtualHubId() terra.StringValue {
	return terra.ReferenceAsString(np.ref.Append("virtual_hub_id"))
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

type PanoramaState struct {
	DeviceGroupName      string `json:"device_group_name"`
	HostName             string `json:"host_name"`
	Name                 string `json:"name"`
	PanoramaServer1      string `json:"panorama_server_1"`
	PanoramaServer2      string `json:"panorama_server_2"`
	TemplateName         string `json:"template_name"`
	VirtualMachineSshKey string `json:"virtual_machine_ssh_key"`
}

type DestinationNatState struct {
	Name           string                `json:"name"`
	Protocol       string                `json:"protocol"`
	BackendConfig  []BackendConfigState  `json:"backend_config"`
	FrontendConfig []FrontendConfigState `json:"frontend_config"`
}

type BackendConfigState struct {
	Port            float64 `json:"port"`
	PublicIpAddress string  `json:"public_ip_address"`
}

type FrontendConfigState struct {
	Port              float64 `json:"port"`
	PublicIpAddressId string  `json:"public_ip_address_id"`
}

type DnsSettingsState struct {
	AzureDnsServers []string `json:"azure_dns_servers"`
	DnsServers      []string `json:"dns_servers"`
	UseAzureDns     bool     `json:"use_azure_dns"`
}

type NetworkProfileState struct {
	EgressNatIpAddressIds         []string `json:"egress_nat_ip_address_ids"`
	EgressNatIpAddresses          []string `json:"egress_nat_ip_addresses"`
	IpOfTrustForUserDefinedRoutes string   `json:"ip_of_trust_for_user_defined_routes"`
	NetworkVirtualApplianceId     string   `json:"network_virtual_appliance_id"`
	PublicIpAddressIds            []string `json:"public_ip_address_ids"`
	PublicIpAddresses             []string `json:"public_ip_addresses"`
	TrustedSubnetId               string   `json:"trusted_subnet_id"`
	UntrustedSubnetId             string   `json:"untrusted_subnet_id"`
	VirtualHubId                  string   `json:"virtual_hub_id"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
}