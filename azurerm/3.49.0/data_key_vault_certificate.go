// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datakeyvaultcertificate "github.com/golingon/terraproviders/azurerm/3.49.0/datakeyvaultcertificate"
	"github.com/volvo-cars/lingon/pkg/terra"
)

func NewDataKeyVaultCertificate(name string, args DataKeyVaultCertificateArgs) *DataKeyVaultCertificate {
	return &DataKeyVaultCertificate{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataKeyVaultCertificate)(nil)

type DataKeyVaultCertificate struct {
	Name string
	Args DataKeyVaultCertificateArgs
}

func (kvc *DataKeyVaultCertificate) DataSource() string {
	return "azurerm_key_vault_certificate"
}

func (kvc *DataKeyVaultCertificate) LocalName() string {
	return kvc.Name
}

func (kvc *DataKeyVaultCertificate) Configuration() interface{} {
	return kvc.Args
}

func (kvc *DataKeyVaultCertificate) Attributes() dataKeyVaultCertificateAttributes {
	return dataKeyVaultCertificateAttributes{ref: terra.ReferenceDataResource(kvc)}
}

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

func (kvc dataKeyVaultCertificateAttributes) CertificateData() terra.StringValue {
	return terra.ReferenceString(kvc.ref.Append("certificate_data"))
}

func (kvc dataKeyVaultCertificateAttributes) CertificateDataBase64() terra.StringValue {
	return terra.ReferenceString(kvc.ref.Append("certificate_data_base64"))
}

func (kvc dataKeyVaultCertificateAttributes) Expires() terra.StringValue {
	return terra.ReferenceString(kvc.ref.Append("expires"))
}

func (kvc dataKeyVaultCertificateAttributes) Id() terra.StringValue {
	return terra.ReferenceString(kvc.ref.Append("id"))
}

func (kvc dataKeyVaultCertificateAttributes) KeyVaultId() terra.StringValue {
	return terra.ReferenceString(kvc.ref.Append("key_vault_id"))
}

func (kvc dataKeyVaultCertificateAttributes) Name() terra.StringValue {
	return terra.ReferenceString(kvc.ref.Append("name"))
}

func (kvc dataKeyVaultCertificateAttributes) NotBefore() terra.StringValue {
	return terra.ReferenceString(kvc.ref.Append("not_before"))
}

func (kvc dataKeyVaultCertificateAttributes) SecretId() terra.StringValue {
	return terra.ReferenceString(kvc.ref.Append("secret_id"))
}

func (kvc dataKeyVaultCertificateAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](kvc.ref.Append("tags"))
}

func (kvc dataKeyVaultCertificateAttributes) Thumbprint() terra.StringValue {
	return terra.ReferenceString(kvc.ref.Append("thumbprint"))
}

func (kvc dataKeyVaultCertificateAttributes) Version() terra.StringValue {
	return terra.ReferenceString(kvc.ref.Append("version"))
}

func (kvc dataKeyVaultCertificateAttributes) VersionlessId() terra.StringValue {
	return terra.ReferenceString(kvc.ref.Append("versionless_id"))
}

func (kvc dataKeyVaultCertificateAttributes) VersionlessSecretId() terra.StringValue {
	return terra.ReferenceString(kvc.ref.Append("versionless_secret_id"))
}

func (kvc dataKeyVaultCertificateAttributes) CertificatePolicy() terra.ListValue[datakeyvaultcertificate.CertificatePolicyAttributes] {
	return terra.ReferenceList[datakeyvaultcertificate.CertificatePolicyAttributes](kvc.ref.Append("certificate_policy"))
}

func (kvc dataKeyVaultCertificateAttributes) Timeouts() datakeyvaultcertificate.TimeoutsAttributes {
	return terra.ReferenceSingle[datakeyvaultcertificate.TimeoutsAttributes](kvc.ref.Append("timeouts"))
}
