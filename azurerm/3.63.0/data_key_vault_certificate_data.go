// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datakeyvaultcertificatedata "github.com/golingon/terraproviders/azurerm/3.63.0/datakeyvaultcertificatedata"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataKeyVaultCertificateData creates a new instance of [DataKeyVaultCertificateData].
func NewDataKeyVaultCertificateData(name string, args DataKeyVaultCertificateDataArgs) *DataKeyVaultCertificateData {
	return &DataKeyVaultCertificateData{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataKeyVaultCertificateData)(nil)

// DataKeyVaultCertificateData represents the Terraform data resource azurerm_key_vault_certificate_data.
type DataKeyVaultCertificateData struct {
	Name string
	Args DataKeyVaultCertificateDataArgs
}

// DataSource returns the Terraform object type for [DataKeyVaultCertificateData].
func (kvcd *DataKeyVaultCertificateData) DataSource() string {
	return "azurerm_key_vault_certificate_data"
}

// LocalName returns the local name for [DataKeyVaultCertificateData].
func (kvcd *DataKeyVaultCertificateData) LocalName() string {
	return kvcd.Name
}

// Configuration returns the configuration (args) for [DataKeyVaultCertificateData].
func (kvcd *DataKeyVaultCertificateData) Configuration() interface{} {
	return kvcd.Args
}

// Attributes returns the attributes for [DataKeyVaultCertificateData].
func (kvcd *DataKeyVaultCertificateData) Attributes() dataKeyVaultCertificateDataAttributes {
	return dataKeyVaultCertificateDataAttributes{ref: terra.ReferenceDataResource(kvcd)}
}

// DataKeyVaultCertificateDataArgs contains the configurations for azurerm_key_vault_certificate_data.
type DataKeyVaultCertificateDataArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// KeyVaultId: string, required
	KeyVaultId terra.StringValue `hcl:"key_vault_id,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Version: string, optional
	Version terra.StringValue `hcl:"version,attr"`
	// Timeouts: optional
	Timeouts *datakeyvaultcertificatedata.Timeouts `hcl:"timeouts,block"`
}
type dataKeyVaultCertificateDataAttributes struct {
	ref terra.Reference
}

// CertificatesCount returns a reference to field certificates_count of azurerm_key_vault_certificate_data.
func (kvcd dataKeyVaultCertificateDataAttributes) CertificatesCount() terra.NumberValue {
	return terra.ReferenceAsNumber(kvcd.ref.Append("certificates_count"))
}

// Expires returns a reference to field expires of azurerm_key_vault_certificate_data.
func (kvcd dataKeyVaultCertificateDataAttributes) Expires() terra.StringValue {
	return terra.ReferenceAsString(kvcd.ref.Append("expires"))
}

// Hex returns a reference to field hex of azurerm_key_vault_certificate_data.
func (kvcd dataKeyVaultCertificateDataAttributes) Hex() terra.StringValue {
	return terra.ReferenceAsString(kvcd.ref.Append("hex"))
}

// Id returns a reference to field id of azurerm_key_vault_certificate_data.
func (kvcd dataKeyVaultCertificateDataAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(kvcd.ref.Append("id"))
}

// Key returns a reference to field key of azurerm_key_vault_certificate_data.
func (kvcd dataKeyVaultCertificateDataAttributes) Key() terra.StringValue {
	return terra.ReferenceAsString(kvcd.ref.Append("key"))
}

// KeyVaultId returns a reference to field key_vault_id of azurerm_key_vault_certificate_data.
func (kvcd dataKeyVaultCertificateDataAttributes) KeyVaultId() terra.StringValue {
	return terra.ReferenceAsString(kvcd.ref.Append("key_vault_id"))
}

// Name returns a reference to field name of azurerm_key_vault_certificate_data.
func (kvcd dataKeyVaultCertificateDataAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(kvcd.ref.Append("name"))
}

// NotBefore returns a reference to field not_before of azurerm_key_vault_certificate_data.
func (kvcd dataKeyVaultCertificateDataAttributes) NotBefore() terra.StringValue {
	return terra.ReferenceAsString(kvcd.ref.Append("not_before"))
}

// Pem returns a reference to field pem of azurerm_key_vault_certificate_data.
func (kvcd dataKeyVaultCertificateDataAttributes) Pem() terra.StringValue {
	return terra.ReferenceAsString(kvcd.ref.Append("pem"))
}

// Tags returns a reference to field tags of azurerm_key_vault_certificate_data.
func (kvcd dataKeyVaultCertificateDataAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](kvcd.ref.Append("tags"))
}

// Version returns a reference to field version of azurerm_key_vault_certificate_data.
func (kvcd dataKeyVaultCertificateDataAttributes) Version() terra.StringValue {
	return terra.ReferenceAsString(kvcd.ref.Append("version"))
}

func (kvcd dataKeyVaultCertificateDataAttributes) Timeouts() datakeyvaultcertificatedata.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datakeyvaultcertificatedata.TimeoutsAttributes](kvcd.ref.Append("timeouts"))
}