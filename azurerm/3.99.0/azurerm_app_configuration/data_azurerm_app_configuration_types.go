// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_app_configuration

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type DataTimeouts struct {
	// Read: string, optional
	Read terra.StringValue `hcl:"read,attr"`
}

type DataEncryptionAttributes struct {
	ref terra.Reference
}

func (e DataEncryptionAttributes) InternalRef() (terra.Reference, error) {
	return e.ref, nil
}

func (e DataEncryptionAttributes) InternalWithRef(ref terra.Reference) DataEncryptionAttributes {
	return DataEncryptionAttributes{ref: ref}
}

func (e DataEncryptionAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return e.ref.InternalTokens()
}

func (e DataEncryptionAttributes) IdentityClientId() terra.StringValue {
	return terra.ReferenceAsString(e.ref.Append("identity_client_id"))
}

func (e DataEncryptionAttributes) KeyVaultKeyIdentifier() terra.StringValue {
	return terra.ReferenceAsString(e.ref.Append("key_vault_key_identifier"))
}

type DataIdentityAttributes struct {
	ref terra.Reference
}

func (i DataIdentityAttributes) InternalRef() (terra.Reference, error) {
	return i.ref, nil
}

func (i DataIdentityAttributes) InternalWithRef(ref terra.Reference) DataIdentityAttributes {
	return DataIdentityAttributes{ref: ref}
}

func (i DataIdentityAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return i.ref.InternalTokens()
}

func (i DataIdentityAttributes) IdentityIds() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](i.ref.Append("identity_ids"))
}

func (i DataIdentityAttributes) PrincipalId() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("principal_id"))
}

func (i DataIdentityAttributes) TenantId() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("tenant_id"))
}

func (i DataIdentityAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("type"))
}

type DataPrimaryReadKeyAttributes struct {
	ref terra.Reference
}

func (prk DataPrimaryReadKeyAttributes) InternalRef() (terra.Reference, error) {
	return prk.ref, nil
}

func (prk DataPrimaryReadKeyAttributes) InternalWithRef(ref terra.Reference) DataPrimaryReadKeyAttributes {
	return DataPrimaryReadKeyAttributes{ref: ref}
}

func (prk DataPrimaryReadKeyAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return prk.ref.InternalTokens()
}

func (prk DataPrimaryReadKeyAttributes) ConnectionString() terra.StringValue {
	return terra.ReferenceAsString(prk.ref.Append("connection_string"))
}

func (prk DataPrimaryReadKeyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(prk.ref.Append("id"))
}

func (prk DataPrimaryReadKeyAttributes) Secret() terra.StringValue {
	return terra.ReferenceAsString(prk.ref.Append("secret"))
}

type DataPrimaryWriteKeyAttributes struct {
	ref terra.Reference
}

func (pwk DataPrimaryWriteKeyAttributes) InternalRef() (terra.Reference, error) {
	return pwk.ref, nil
}

func (pwk DataPrimaryWriteKeyAttributes) InternalWithRef(ref terra.Reference) DataPrimaryWriteKeyAttributes {
	return DataPrimaryWriteKeyAttributes{ref: ref}
}

func (pwk DataPrimaryWriteKeyAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return pwk.ref.InternalTokens()
}

func (pwk DataPrimaryWriteKeyAttributes) ConnectionString() terra.StringValue {
	return terra.ReferenceAsString(pwk.ref.Append("connection_string"))
}

func (pwk DataPrimaryWriteKeyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(pwk.ref.Append("id"))
}

func (pwk DataPrimaryWriteKeyAttributes) Secret() terra.StringValue {
	return terra.ReferenceAsString(pwk.ref.Append("secret"))
}

type DataReplicaAttributes struct {
	ref terra.Reference
}

func (r DataReplicaAttributes) InternalRef() (terra.Reference, error) {
	return r.ref, nil
}

func (r DataReplicaAttributes) InternalWithRef(ref terra.Reference) DataReplicaAttributes {
	return DataReplicaAttributes{ref: ref}
}

func (r DataReplicaAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return r.ref.InternalTokens()
}

func (r DataReplicaAttributes) Endpoint() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("endpoint"))
}

func (r DataReplicaAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("id"))
}

func (r DataReplicaAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("location"))
}

func (r DataReplicaAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("name"))
}

type DataSecondaryReadKeyAttributes struct {
	ref terra.Reference
}

func (srk DataSecondaryReadKeyAttributes) InternalRef() (terra.Reference, error) {
	return srk.ref, nil
}

func (srk DataSecondaryReadKeyAttributes) InternalWithRef(ref terra.Reference) DataSecondaryReadKeyAttributes {
	return DataSecondaryReadKeyAttributes{ref: ref}
}

func (srk DataSecondaryReadKeyAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return srk.ref.InternalTokens()
}

func (srk DataSecondaryReadKeyAttributes) ConnectionString() terra.StringValue {
	return terra.ReferenceAsString(srk.ref.Append("connection_string"))
}

func (srk DataSecondaryReadKeyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(srk.ref.Append("id"))
}

func (srk DataSecondaryReadKeyAttributes) Secret() terra.StringValue {
	return terra.ReferenceAsString(srk.ref.Append("secret"))
}

type DataSecondaryWriteKeyAttributes struct {
	ref terra.Reference
}

func (swk DataSecondaryWriteKeyAttributes) InternalRef() (terra.Reference, error) {
	return swk.ref, nil
}

func (swk DataSecondaryWriteKeyAttributes) InternalWithRef(ref terra.Reference) DataSecondaryWriteKeyAttributes {
	return DataSecondaryWriteKeyAttributes{ref: ref}
}

func (swk DataSecondaryWriteKeyAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return swk.ref.InternalTokens()
}

func (swk DataSecondaryWriteKeyAttributes) ConnectionString() terra.StringValue {
	return terra.ReferenceAsString(swk.ref.Append("connection_string"))
}

func (swk DataSecondaryWriteKeyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(swk.ref.Append("id"))
}

func (swk DataSecondaryWriteKeyAttributes) Secret() terra.StringValue {
	return terra.ReferenceAsString(swk.ref.Append("secret"))
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

type DataEncryptionState struct {
	IdentityClientId      string `json:"identity_client_id"`
	KeyVaultKeyIdentifier string `json:"key_vault_key_identifier"`
}

type DataIdentityState struct {
	IdentityIds []string `json:"identity_ids"`
	PrincipalId string   `json:"principal_id"`
	TenantId    string   `json:"tenant_id"`
	Type        string   `json:"type"`
}

type DataPrimaryReadKeyState struct {
	ConnectionString string `json:"connection_string"`
	Id               string `json:"id"`
	Secret           string `json:"secret"`
}

type DataPrimaryWriteKeyState struct {
	ConnectionString string `json:"connection_string"`
	Id               string `json:"id"`
	Secret           string `json:"secret"`
}

type DataReplicaState struct {
	Endpoint string `json:"endpoint"`
	Id       string `json:"id"`
	Location string `json:"location"`
	Name     string `json:"name"`
}

type DataSecondaryReadKeyState struct {
	ConnectionString string `json:"connection_string"`
	Id               string `json:"id"`
	Secret           string `json:"secret"`
}

type DataSecondaryWriteKeyState struct {
	ConnectionString string `json:"connection_string"`
	Id               string `json:"id"`
	Secret           string `json:"secret"`
}

type DataTimeoutsState struct {
	Read string `json:"read"`
}
