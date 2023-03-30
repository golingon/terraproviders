// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package fsxontapstoragevirtualmachine

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Endpoints struct {
	// Iscsi: min=0
	Iscsi []Iscsi `hcl:"iscsi,block" validate:"min=0"`
	// Management: min=0
	Management []Management `hcl:"management,block" validate:"min=0"`
	// Nfs: min=0
	Nfs []Nfs `hcl:"nfs,block" validate:"min=0"`
	// Smb: min=0
	Smb []Smb `hcl:"smb,block" validate:"min=0"`
}

type Iscsi struct{}

type Management struct{}

type Nfs struct{}

type Smb struct{}

type ActiveDirectoryConfiguration struct {
	// NetbiosName: string, optional
	NetbiosName terra.StringValue `hcl:"netbios_name,attr"`
	// SelfManagedActiveDirectoryConfiguration: optional
	SelfManagedActiveDirectoryConfiguration *SelfManagedActiveDirectoryConfiguration `hcl:"self_managed_active_directory_configuration,block"`
}

type SelfManagedActiveDirectoryConfiguration struct {
	// DnsIps: set of string, required
	DnsIps terra.SetValue[terra.StringValue] `hcl:"dns_ips,attr" validate:"required"`
	// DomainName: string, required
	DomainName terra.StringValue `hcl:"domain_name,attr" validate:"required"`
	// FileSystemAdministratorsGroup: string, optional
	FileSystemAdministratorsGroup terra.StringValue `hcl:"file_system_administrators_group,attr"`
	// OrganizationalUnitDistinguishedName: string, optional
	OrganizationalUnitDistinguishedName terra.StringValue `hcl:"organizational_unit_distinguished_name,attr"`
	// Password: string, required
	Password terra.StringValue `hcl:"password,attr" validate:"required"`
	// Username: string, required
	Username terra.StringValue `hcl:"username,attr" validate:"required"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type EndpointsAttributes struct {
	ref terra.Reference
}

func (e EndpointsAttributes) InternalRef() terra.Reference {
	return e.ref
}

func (e EndpointsAttributes) InternalWithRef(ref terra.Reference) EndpointsAttributes {
	return EndpointsAttributes{ref: ref}
}

func (e EndpointsAttributes) InternalTokens() hclwrite.Tokens {
	return e.ref.InternalTokens()
}

func (e EndpointsAttributes) Iscsi() terra.ListValue[IscsiAttributes] {
	return terra.ReferenceList[IscsiAttributes](e.ref.Append("iscsi"))
}

func (e EndpointsAttributes) Management() terra.ListValue[ManagementAttributes] {
	return terra.ReferenceList[ManagementAttributes](e.ref.Append("management"))
}

func (e EndpointsAttributes) Nfs() terra.ListValue[NfsAttributes] {
	return terra.ReferenceList[NfsAttributes](e.ref.Append("nfs"))
}

func (e EndpointsAttributes) Smb() terra.ListValue[SmbAttributes] {
	return terra.ReferenceList[SmbAttributes](e.ref.Append("smb"))
}

type IscsiAttributes struct {
	ref terra.Reference
}

func (i IscsiAttributes) InternalRef() terra.Reference {
	return i.ref
}

func (i IscsiAttributes) InternalWithRef(ref terra.Reference) IscsiAttributes {
	return IscsiAttributes{ref: ref}
}

func (i IscsiAttributes) InternalTokens() hclwrite.Tokens {
	return i.ref.InternalTokens()
}

func (i IscsiAttributes) DnsName() terra.StringValue {
	return terra.ReferenceString(i.ref.Append("dns_name"))
}

func (i IscsiAttributes) IpAddresses() terra.SetValue[terra.StringValue] {
	return terra.ReferenceSet[terra.StringValue](i.ref.Append("ip_addresses"))
}

type ManagementAttributes struct {
	ref terra.Reference
}

func (m ManagementAttributes) InternalRef() terra.Reference {
	return m.ref
}

func (m ManagementAttributes) InternalWithRef(ref terra.Reference) ManagementAttributes {
	return ManagementAttributes{ref: ref}
}

func (m ManagementAttributes) InternalTokens() hclwrite.Tokens {
	return m.ref.InternalTokens()
}

func (m ManagementAttributes) DnsName() terra.StringValue {
	return terra.ReferenceString(m.ref.Append("dns_name"))
}

func (m ManagementAttributes) IpAddresses() terra.SetValue[terra.StringValue] {
	return terra.ReferenceSet[terra.StringValue](m.ref.Append("ip_addresses"))
}

type NfsAttributes struct {
	ref terra.Reference
}

func (n NfsAttributes) InternalRef() terra.Reference {
	return n.ref
}

func (n NfsAttributes) InternalWithRef(ref terra.Reference) NfsAttributes {
	return NfsAttributes{ref: ref}
}

func (n NfsAttributes) InternalTokens() hclwrite.Tokens {
	return n.ref.InternalTokens()
}

func (n NfsAttributes) DnsName() terra.StringValue {
	return terra.ReferenceString(n.ref.Append("dns_name"))
}

func (n NfsAttributes) IpAddresses() terra.SetValue[terra.StringValue] {
	return terra.ReferenceSet[terra.StringValue](n.ref.Append("ip_addresses"))
}

type SmbAttributes struct {
	ref terra.Reference
}

func (s SmbAttributes) InternalRef() terra.Reference {
	return s.ref
}

func (s SmbAttributes) InternalWithRef(ref terra.Reference) SmbAttributes {
	return SmbAttributes{ref: ref}
}

func (s SmbAttributes) InternalTokens() hclwrite.Tokens {
	return s.ref.InternalTokens()
}

func (s SmbAttributes) DnsName() terra.StringValue {
	return terra.ReferenceString(s.ref.Append("dns_name"))
}

func (s SmbAttributes) IpAddresses() terra.SetValue[terra.StringValue] {
	return terra.ReferenceSet[terra.StringValue](s.ref.Append("ip_addresses"))
}

type ActiveDirectoryConfigurationAttributes struct {
	ref terra.Reference
}

func (adc ActiveDirectoryConfigurationAttributes) InternalRef() terra.Reference {
	return adc.ref
}

func (adc ActiveDirectoryConfigurationAttributes) InternalWithRef(ref terra.Reference) ActiveDirectoryConfigurationAttributes {
	return ActiveDirectoryConfigurationAttributes{ref: ref}
}

func (adc ActiveDirectoryConfigurationAttributes) InternalTokens() hclwrite.Tokens {
	return adc.ref.InternalTokens()
}

func (adc ActiveDirectoryConfigurationAttributes) NetbiosName() terra.StringValue {
	return terra.ReferenceString(adc.ref.Append("netbios_name"))
}

func (adc ActiveDirectoryConfigurationAttributes) SelfManagedActiveDirectoryConfiguration() terra.ListValue[SelfManagedActiveDirectoryConfigurationAttributes] {
	return terra.ReferenceList[SelfManagedActiveDirectoryConfigurationAttributes](adc.ref.Append("self_managed_active_directory_configuration"))
}

type SelfManagedActiveDirectoryConfigurationAttributes struct {
	ref terra.Reference
}

func (smadc SelfManagedActiveDirectoryConfigurationAttributes) InternalRef() terra.Reference {
	return smadc.ref
}

func (smadc SelfManagedActiveDirectoryConfigurationAttributes) InternalWithRef(ref terra.Reference) SelfManagedActiveDirectoryConfigurationAttributes {
	return SelfManagedActiveDirectoryConfigurationAttributes{ref: ref}
}

func (smadc SelfManagedActiveDirectoryConfigurationAttributes) InternalTokens() hclwrite.Tokens {
	return smadc.ref.InternalTokens()
}

func (smadc SelfManagedActiveDirectoryConfigurationAttributes) DnsIps() terra.SetValue[terra.StringValue] {
	return terra.ReferenceSet[terra.StringValue](smadc.ref.Append("dns_ips"))
}

func (smadc SelfManagedActiveDirectoryConfigurationAttributes) DomainName() terra.StringValue {
	return terra.ReferenceString(smadc.ref.Append("domain_name"))
}

func (smadc SelfManagedActiveDirectoryConfigurationAttributes) FileSystemAdministratorsGroup() terra.StringValue {
	return terra.ReferenceString(smadc.ref.Append("file_system_administrators_group"))
}

func (smadc SelfManagedActiveDirectoryConfigurationAttributes) OrganizationalUnitDistinguishedName() terra.StringValue {
	return terra.ReferenceString(smadc.ref.Append("organizational_unit_distinguished_name"))
}

func (smadc SelfManagedActiveDirectoryConfigurationAttributes) Password() terra.StringValue {
	return terra.ReferenceString(smadc.ref.Append("password"))
}

func (smadc SelfManagedActiveDirectoryConfigurationAttributes) Username() terra.StringValue {
	return terra.ReferenceString(smadc.ref.Append("username"))
}

type TimeoutsAttributes struct {
	ref terra.Reference
}

func (t TimeoutsAttributes) InternalRef() terra.Reference {
	return t.ref
}

func (t TimeoutsAttributes) InternalWithRef(ref terra.Reference) TimeoutsAttributes {
	return TimeoutsAttributes{ref: ref}
}

func (t TimeoutsAttributes) InternalTokens() hclwrite.Tokens {
	return t.ref.InternalTokens()
}

func (t TimeoutsAttributes) Create() terra.StringValue {
	return terra.ReferenceString(t.ref.Append("create"))
}

func (t TimeoutsAttributes) Delete() terra.StringValue {
	return terra.ReferenceString(t.ref.Append("delete"))
}

func (t TimeoutsAttributes) Update() terra.StringValue {
	return terra.ReferenceString(t.ref.Append("update"))
}

type EndpointsState struct {
	Iscsi      []IscsiState      `json:"iscsi"`
	Management []ManagementState `json:"management"`
	Nfs        []NfsState        `json:"nfs"`
	Smb        []SmbState        `json:"smb"`
}

type IscsiState struct {
	DnsName     string   `json:"dns_name"`
	IpAddresses []string `json:"ip_addresses"`
}

type ManagementState struct {
	DnsName     string   `json:"dns_name"`
	IpAddresses []string `json:"ip_addresses"`
}

type NfsState struct {
	DnsName     string   `json:"dns_name"`
	IpAddresses []string `json:"ip_addresses"`
}

type SmbState struct {
	DnsName     string   `json:"dns_name"`
	IpAddresses []string `json:"ip_addresses"`
}

type ActiveDirectoryConfigurationState struct {
	NetbiosName                             string                                         `json:"netbios_name"`
	SelfManagedActiveDirectoryConfiguration []SelfManagedActiveDirectoryConfigurationState `json:"self_managed_active_directory_configuration"`
}

type SelfManagedActiveDirectoryConfigurationState struct {
	DnsIps                              []string `json:"dns_ips"`
	DomainName                          string   `json:"domain_name"`
	FileSystemAdministratorsGroup       string   `json:"file_system_administrators_group"`
	OrganizationalUnitDistinguishedName string   `json:"organizational_unit_distinguished_name"`
	Password                            string   `json:"password"`
	Username                            string   `json:"username"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}
