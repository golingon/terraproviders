// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datakeyvaultaccesspolicy "github.com/golingon/terraproviders/azurerm/3.49.0/datakeyvaultaccesspolicy"
	"github.com/volvo-cars/lingon/pkg/terra"
)

func NewDataKeyVaultAccessPolicy(name string, args DataKeyVaultAccessPolicyArgs) *DataKeyVaultAccessPolicy {
	return &DataKeyVaultAccessPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataKeyVaultAccessPolicy)(nil)

type DataKeyVaultAccessPolicy struct {
	Name string
	Args DataKeyVaultAccessPolicyArgs
}

func (kvap *DataKeyVaultAccessPolicy) DataSource() string {
	return "azurerm_key_vault_access_policy"
}

func (kvap *DataKeyVaultAccessPolicy) LocalName() string {
	return kvap.Name
}

func (kvap *DataKeyVaultAccessPolicy) Configuration() interface{} {
	return kvap.Args
}

func (kvap *DataKeyVaultAccessPolicy) Attributes() dataKeyVaultAccessPolicyAttributes {
	return dataKeyVaultAccessPolicyAttributes{ref: terra.ReferenceDataResource(kvap)}
}

type DataKeyVaultAccessPolicyArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *datakeyvaultaccesspolicy.Timeouts `hcl:"timeouts,block"`
}
type dataKeyVaultAccessPolicyAttributes struct {
	ref terra.Reference
}

func (kvap dataKeyVaultAccessPolicyAttributes) CertificatePermissions() terra.ListValue[terra.StringValue] {
	return terra.ReferenceList[terra.StringValue](kvap.ref.Append("certificate_permissions"))
}

func (kvap dataKeyVaultAccessPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceString(kvap.ref.Append("id"))
}

func (kvap dataKeyVaultAccessPolicyAttributes) KeyPermissions() terra.ListValue[terra.StringValue] {
	return terra.ReferenceList[terra.StringValue](kvap.ref.Append("key_permissions"))
}

func (kvap dataKeyVaultAccessPolicyAttributes) Name() terra.StringValue {
	return terra.ReferenceString(kvap.ref.Append("name"))
}

func (kvap dataKeyVaultAccessPolicyAttributes) SecretPermissions() terra.ListValue[terra.StringValue] {
	return terra.ReferenceList[terra.StringValue](kvap.ref.Append("secret_permissions"))
}

func (kvap dataKeyVaultAccessPolicyAttributes) Timeouts() datakeyvaultaccesspolicy.TimeoutsAttributes {
	return terra.ReferenceSingle[datakeyvaultaccesspolicy.TimeoutsAttributes](kvap.ref.Append("timeouts"))
}
