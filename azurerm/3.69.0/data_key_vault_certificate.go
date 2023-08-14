// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datakeyvaultcertificate "github.com/golingon/terraproviders/azurerm/3.69.0/datakeyvaultcertificate"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataKeyVaultCertificate creates a new instance of [DataKeyVaultCertificate].
func NewDataKeyVaultCertificate(name string, args DataKeyVaultCertificateArgs) *DataKeyVaultCertificate {
	return &DataKeyVaultCertificate{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataKeyVaultCertificate)(nil)

// DataKeyVaultCertificate represents the Terraform data resource azurerm_key_vault_certificate.
type DataKeyVaultCertificate struct {
	Name string
	Args DataKeyVaultCertificateArgs
}

// DataSource returns the Terraform object type for [DataKeyVaultCertificate].
func (kvc *DataKeyVaultCertificate) DataSource() string {
	return "azurerm_key_vault_certificate"
}

// LocalName returns the local name for [DataKeyVaultCertificate].
func (kvc *DataKeyVaultCertificate) LocalName() string {
	return kvc.Name
}

// Configuration returns the configuration (args) for [DataKeyVaultCertificate].
func (kvc *DataKeyVaultCertificate) Configuration() interface{} {
	return kvc.Args
}

// Attributes returns the attributes for [DataKeyVaultCertificate].
func (kvc *DataKeyVaultCertificate) Attributes() dataKeyVaultCertificateAttributes {
	return dataKeyVaultCertificateAttributes{ref: terra.ReferenceDataResource(kvc)}
}

// DataKeyVaultCertificateArgs contains the configurations for azurerm_key_vault_certificate.
type DataKeyVaultCertificateArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// KeyVaultId: string, required
	KeyVaultId terra.StringValue `hcl:"key_vault_id,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Version: string, optional
	Version terra.StringValue `hcl:"version,attr"`
	// CertificatePolicy: min=0
	CertificatePolicy []datakeyvaultcertificate.CertificatePolicy `hcl:"certificate_policy,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *datakeyvaultcertificate.Timeouts `hcl:"timeouts,block"`
}
type dataKeyVaultCertificateAttributes struct {
	ref terra.Reference
}

// CertificateData returns a reference to field certificate_data of azurerm_key_vault_certificate.
func (kvc dataKeyVaultCertificateAttributes) CertificateData() terra.StringValue {
	return terra.ReferenceAsString(kvc.ref.Append("certificate_data"))
}

// CertificateDataBase64 returns a reference to field certificate_data_base64 of azurerm_key_vault_certificate.
func (kvc dataKeyVaultCertificateAttributes) CertificateDataBase64() terra.StringValue {
	return terra.ReferenceAsString(kvc.ref.Append("certificate_data_base64"))
}

// Expires returns a reference to field expires of azurerm_key_vault_certificate.
func (kvc dataKeyVaultCertificateAttributes) Expires() terra.StringValue {
	return terra.ReferenceAsString(kvc.ref.Append("expires"))
}

// Id returns a reference to field id of azurerm_key_vault_certificate.
func (kvc dataKeyVaultCertificateAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(kvc.ref.Append("id"))
}

// KeyVaultId returns a reference to field key_vault_id of azurerm_key_vault_certificate.
func (kvc dataKeyVaultCertificateAttributes) KeyVaultId() terra.StringValue {
	return terra.ReferenceAsString(kvc.ref.Append("key_vault_id"))
}

// Name returns a reference to field name of azurerm_key_vault_certificate.
func (kvc dataKeyVaultCertificateAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(kvc.ref.Append("name"))
}

// NotBefore returns a reference to field not_before of azurerm_key_vault_certificate.
func (kvc dataKeyVaultCertificateAttributes) NotBefore() terra.StringValue {
	return terra.ReferenceAsString(kvc.ref.Append("not_before"))
}

// ResourceManagerId returns a reference to field resource_manager_id of azurerm_key_vault_certificate.
func (kvc dataKeyVaultCertificateAttributes) ResourceManagerId() terra.StringValue {
	return terra.ReferenceAsString(kvc.ref.Append("resource_manager_id"))
}

// ResourceManagerVersionlessId returns a reference to field resource_manager_versionless_id of azurerm_key_vault_certificate.
func (kvc dataKeyVaultCertificateAttributes) ResourceManagerVersionlessId() terra.StringValue {
	return terra.ReferenceAsString(kvc.ref.Append("resource_manager_versionless_id"))
}

// SecretId returns a reference to field secret_id of azurerm_key_vault_certificate.
func (kvc dataKeyVaultCertificateAttributes) SecretId() terra.StringValue {
	return terra.ReferenceAsString(kvc.ref.Append("secret_id"))
}

// Tags returns a reference to field tags of azurerm_key_vault_certificate.
func (kvc dataKeyVaultCertificateAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](kvc.ref.Append("tags"))
}

// Thumbprint returns a reference to field thumbprint of azurerm_key_vault_certificate.
func (kvc dataKeyVaultCertificateAttributes) Thumbprint() terra.StringValue {
	return terra.ReferenceAsString(kvc.ref.Append("thumbprint"))
}

// Version returns a reference to field version of azurerm_key_vault_certificate.
func (kvc dataKeyVaultCertificateAttributes) Version() terra.StringValue {
	return terra.ReferenceAsString(kvc.ref.Append("version"))
}

// VersionlessId returns a reference to field versionless_id of azurerm_key_vault_certificate.
func (kvc dataKeyVaultCertificateAttributes) VersionlessId() terra.StringValue {
	return terra.ReferenceAsString(kvc.ref.Append("versionless_id"))
}

// VersionlessSecretId returns a reference to field versionless_secret_id of azurerm_key_vault_certificate.
func (kvc dataKeyVaultCertificateAttributes) VersionlessSecretId() terra.StringValue {
	return terra.ReferenceAsString(kvc.ref.Append("versionless_secret_id"))
}

func (kvc dataKeyVaultCertificateAttributes) CertificatePolicy() terra.ListValue[datakeyvaultcertificate.CertificatePolicyAttributes] {
	return terra.ReferenceAsList[datakeyvaultcertificate.CertificatePolicyAttributes](kvc.ref.Append("certificate_policy"))
}

func (kvc dataKeyVaultCertificateAttributes) Timeouts() datakeyvaultcertificate.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datakeyvaultcertificate.TimeoutsAttributes](kvc.ref.Append("timeouts"))
}
