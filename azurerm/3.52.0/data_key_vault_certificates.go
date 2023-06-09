// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datakeyvaultcertificates "github.com/golingon/terraproviders/azurerm/3.52.0/datakeyvaultcertificates"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataKeyVaultCertificates creates a new instance of [DataKeyVaultCertificates].
func NewDataKeyVaultCertificates(name string, args DataKeyVaultCertificatesArgs) *DataKeyVaultCertificates {
	return &DataKeyVaultCertificates{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataKeyVaultCertificates)(nil)

// DataKeyVaultCertificates represents the Terraform data resource azurerm_key_vault_certificates.
type DataKeyVaultCertificates struct {
	Name string
	Args DataKeyVaultCertificatesArgs
}

// DataSource returns the Terraform object type for [DataKeyVaultCertificates].
func (kvc *DataKeyVaultCertificates) DataSource() string {
	return "azurerm_key_vault_certificates"
}

// LocalName returns the local name for [DataKeyVaultCertificates].
func (kvc *DataKeyVaultCertificates) LocalName() string {
	return kvc.Name
}

// Configuration returns the configuration (args) for [DataKeyVaultCertificates].
func (kvc *DataKeyVaultCertificates) Configuration() interface{} {
	return kvc.Args
}

// Attributes returns the attributes for [DataKeyVaultCertificates].
func (kvc *DataKeyVaultCertificates) Attributes() dataKeyVaultCertificatesAttributes {
	return dataKeyVaultCertificatesAttributes{ref: terra.ReferenceDataResource(kvc)}
}

// DataKeyVaultCertificatesArgs contains the configurations for azurerm_key_vault_certificates.
type DataKeyVaultCertificatesArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IncludePending: bool, optional
	IncludePending terra.BoolValue `hcl:"include_pending,attr"`
	// KeyVaultId: string, required
	KeyVaultId terra.StringValue `hcl:"key_vault_id,attr" validate:"required"`
	// Certificates: min=0
	Certificates []datakeyvaultcertificates.Certificates `hcl:"certificates,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *datakeyvaultcertificates.Timeouts `hcl:"timeouts,block"`
}
type dataKeyVaultCertificatesAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_key_vault_certificates.
func (kvc dataKeyVaultCertificatesAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(kvc.ref.Append("id"))
}

// IncludePending returns a reference to field include_pending of azurerm_key_vault_certificates.
func (kvc dataKeyVaultCertificatesAttributes) IncludePending() terra.BoolValue {
	return terra.ReferenceAsBool(kvc.ref.Append("include_pending"))
}

// KeyVaultId returns a reference to field key_vault_id of azurerm_key_vault_certificates.
func (kvc dataKeyVaultCertificatesAttributes) KeyVaultId() terra.StringValue {
	return terra.ReferenceAsString(kvc.ref.Append("key_vault_id"))
}

// Names returns a reference to field names of azurerm_key_vault_certificates.
func (kvc dataKeyVaultCertificatesAttributes) Names() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](kvc.ref.Append("names"))
}

func (kvc dataKeyVaultCertificatesAttributes) Certificates() terra.ListValue[datakeyvaultcertificates.CertificatesAttributes] {
	return terra.ReferenceAsList[datakeyvaultcertificates.CertificatesAttributes](kvc.ref.Append("certificates"))
}

func (kvc dataKeyVaultCertificatesAttributes) Timeouts() datakeyvaultcertificates.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datakeyvaultcertificates.TimeoutsAttributes](kvc.ref.Append("timeouts"))
}
