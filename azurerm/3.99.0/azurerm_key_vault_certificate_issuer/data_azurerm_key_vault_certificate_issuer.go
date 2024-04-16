// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_key_vault_certificate_issuer

import "github.com/golingon/lingon/pkg/terra"

// Data creates a new instance of [DataSource].
func Data(name string, args DataArgs) *DataSource {
	return &DataSource{
		Args: args,
		Name: name,
	}
}

var _ terra.DataSource = (*DataSource)(nil)

// DataSource represents the Terraform data resource azurerm_key_vault_certificate_issuer.
type DataSource struct {
	Name string
	Args DataArgs
}

// DataSource returns the Terraform object type for [DataSource].
func (akvci *DataSource) DataSource() string {
	return "azurerm_key_vault_certificate_issuer"
}

// LocalName returns the local name for [DataSource].
func (akvci *DataSource) LocalName() string {
	return akvci.Name
}

// Configuration returns the configuration (args) for [DataSource].
func (akvci *DataSource) Configuration() interface{} {
	return akvci.Args
}

// Attributes returns the attributes for [DataSource].
func (akvci *DataSource) Attributes() dataAzurermKeyVaultCertificateIssuerAttributes {
	return dataAzurermKeyVaultCertificateIssuerAttributes{ref: terra.ReferenceDataSource(akvci)}
}

// DataArgs contains the configurations for azurerm_key_vault_certificate_issuer.
type DataArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// KeyVaultId: string, required
	KeyVaultId terra.StringValue `hcl:"key_vault_id,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *DataTimeouts `hcl:"timeouts,block"`
}

type dataAzurermKeyVaultCertificateIssuerAttributes struct {
	ref terra.Reference
}

// AccountId returns a reference to field account_id of azurerm_key_vault_certificate_issuer.
func (akvci dataAzurermKeyVaultCertificateIssuerAttributes) AccountId() terra.StringValue {
	return terra.ReferenceAsString(akvci.ref.Append("account_id"))
}

// Id returns a reference to field id of azurerm_key_vault_certificate_issuer.
func (akvci dataAzurermKeyVaultCertificateIssuerAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(akvci.ref.Append("id"))
}

// KeyVaultId returns a reference to field key_vault_id of azurerm_key_vault_certificate_issuer.
func (akvci dataAzurermKeyVaultCertificateIssuerAttributes) KeyVaultId() terra.StringValue {
	return terra.ReferenceAsString(akvci.ref.Append("key_vault_id"))
}

// Name returns a reference to field name of azurerm_key_vault_certificate_issuer.
func (akvci dataAzurermKeyVaultCertificateIssuerAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(akvci.ref.Append("name"))
}

// OrgId returns a reference to field org_id of azurerm_key_vault_certificate_issuer.
func (akvci dataAzurermKeyVaultCertificateIssuerAttributes) OrgId() terra.StringValue {
	return terra.ReferenceAsString(akvci.ref.Append("org_id"))
}

// ProviderName returns a reference to field provider_name of azurerm_key_vault_certificate_issuer.
func (akvci dataAzurermKeyVaultCertificateIssuerAttributes) ProviderName() terra.StringValue {
	return terra.ReferenceAsString(akvci.ref.Append("provider_name"))
}

func (akvci dataAzurermKeyVaultCertificateIssuerAttributes) Admin() terra.ListValue[DataAdminAttributes] {
	return terra.ReferenceAsList[DataAdminAttributes](akvci.ref.Append("admin"))
}

func (akvci dataAzurermKeyVaultCertificateIssuerAttributes) Timeouts() DataTimeoutsAttributes {
	return terra.ReferenceAsSingle[DataTimeoutsAttributes](akvci.ref.Append("timeouts"))
}
