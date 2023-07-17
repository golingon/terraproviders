// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package datastorageaccount

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type AzureFilesAuthentication struct {
	// ActiveDirectory: min=0
	ActiveDirectory []ActiveDirectory `hcl:"active_directory,block" validate:"min=0"`
}

type ActiveDirectory struct{}

type CustomDomain struct{}

type Identity struct{}

type Timeouts struct {
	// Read: string, optional
	Read terra.StringValue `hcl:"read,attr"`
}

type AzureFilesAuthenticationAttributes struct {
	ref terra.Reference
}

func (afa AzureFilesAuthenticationAttributes) InternalRef() (terra.Reference, error) {
	return afa.ref, nil
}

func (afa AzureFilesAuthenticationAttributes) InternalWithRef(ref terra.Reference) AzureFilesAuthenticationAttributes {
	return AzureFilesAuthenticationAttributes{ref: ref}
}

func (afa AzureFilesAuthenticationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return afa.ref.InternalTokens()
}

func (afa AzureFilesAuthenticationAttributes) DirectoryType() terra.StringValue {
	return terra.ReferenceAsString(afa.ref.Append("directory_type"))
}

func (afa AzureFilesAuthenticationAttributes) ActiveDirectory() terra.ListValue[ActiveDirectoryAttributes] {
	return terra.ReferenceAsList[ActiveDirectoryAttributes](afa.ref.Append("active_directory"))
}

type ActiveDirectoryAttributes struct {
	ref terra.Reference
}

func (ad ActiveDirectoryAttributes) InternalRef() (terra.Reference, error) {
	return ad.ref, nil
}

func (ad ActiveDirectoryAttributes) InternalWithRef(ref terra.Reference) ActiveDirectoryAttributes {
	return ActiveDirectoryAttributes{ref: ref}
}

func (ad ActiveDirectoryAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ad.ref.InternalTokens()
}

func (ad ActiveDirectoryAttributes) DomainGuid() terra.StringValue {
	return terra.ReferenceAsString(ad.ref.Append("domain_guid"))
}

func (ad ActiveDirectoryAttributes) DomainName() terra.StringValue {
	return terra.ReferenceAsString(ad.ref.Append("domain_name"))
}

func (ad ActiveDirectoryAttributes) DomainSid() terra.StringValue {
	return terra.ReferenceAsString(ad.ref.Append("domain_sid"))
}

func (ad ActiveDirectoryAttributes) ForestName() terra.StringValue {
	return terra.ReferenceAsString(ad.ref.Append("forest_name"))
}

func (ad ActiveDirectoryAttributes) NetbiosDomainName() terra.StringValue {
	return terra.ReferenceAsString(ad.ref.Append("netbios_domain_name"))
}

func (ad ActiveDirectoryAttributes) StorageSid() terra.StringValue {
	return terra.ReferenceAsString(ad.ref.Append("storage_sid"))
}

type CustomDomainAttributes struct {
	ref terra.Reference
}

func (cd CustomDomainAttributes) InternalRef() (terra.Reference, error) {
	return cd.ref, nil
}

func (cd CustomDomainAttributes) InternalWithRef(ref terra.Reference) CustomDomainAttributes {
	return CustomDomainAttributes{ref: ref}
}

func (cd CustomDomainAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return cd.ref.InternalTokens()
}

func (cd CustomDomainAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cd.ref.Append("name"))
}

type IdentityAttributes struct {
	ref terra.Reference
}

func (i IdentityAttributes) InternalRef() (terra.Reference, error) {
	return i.ref, nil
}

func (i IdentityAttributes) InternalWithRef(ref terra.Reference) IdentityAttributes {
	return IdentityAttributes{ref: ref}
}

func (i IdentityAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return i.ref.InternalTokens()
}

func (i IdentityAttributes) IdentityIds() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](i.ref.Append("identity_ids"))
}

func (i IdentityAttributes) PrincipalId() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("principal_id"))
}

func (i IdentityAttributes) TenantId() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("tenant_id"))
}

func (i IdentityAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("type"))
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

type AzureFilesAuthenticationState struct {
	DirectoryType   string                 `json:"directory_type"`
	ActiveDirectory []ActiveDirectoryState `json:"active_directory"`
}

type ActiveDirectoryState struct {
	DomainGuid        string `json:"domain_guid"`
	DomainName        string `json:"domain_name"`
	DomainSid         string `json:"domain_sid"`
	ForestName        string `json:"forest_name"`
	NetbiosDomainName string `json:"netbios_domain_name"`
	StorageSid        string `json:"storage_sid"`
}

type CustomDomainState struct {
	Name string `json:"name"`
}

type IdentityState struct {
	IdentityIds []string `json:"identity_ids"`
	PrincipalId string   `json:"principal_id"`
	TenantId    string   `json:"tenant_id"`
	Type        string   `json:"type"`
}

type TimeoutsState struct {
	Read string `json:"read"`
}