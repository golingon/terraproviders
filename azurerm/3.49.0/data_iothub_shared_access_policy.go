// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	dataiothubsharedaccesspolicy "github.com/golingon/terraproviders/azurerm/3.49.0/dataiothubsharedaccesspolicy"
	"github.com/volvo-cars/lingon/pkg/terra"
)

func NewDataIothubSharedAccessPolicy(name string, args DataIothubSharedAccessPolicyArgs) *DataIothubSharedAccessPolicy {
	return &DataIothubSharedAccessPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataIothubSharedAccessPolicy)(nil)

type DataIothubSharedAccessPolicy struct {
	Name string
	Args DataIothubSharedAccessPolicyArgs
}

func (isap *DataIothubSharedAccessPolicy) DataSource() string {
	return "azurerm_iothub_shared_access_policy"
}

func (isap *DataIothubSharedAccessPolicy) LocalName() string {
	return isap.Name
}

func (isap *DataIothubSharedAccessPolicy) Configuration() interface{} {
	return isap.Args
}

func (isap *DataIothubSharedAccessPolicy) Attributes() dataIothubSharedAccessPolicyAttributes {
	return dataIothubSharedAccessPolicyAttributes{ref: terra.ReferenceDataResource(isap)}
}

type DataIothubSharedAccessPolicyArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IothubName: string, required
	IothubName terra.StringValue `hcl:"iothub_name,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *dataiothubsharedaccesspolicy.Timeouts `hcl:"timeouts,block"`
}
type dataIothubSharedAccessPolicyAttributes struct {
	ref terra.Reference
}

func (isap dataIothubSharedAccessPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceString(isap.ref.Append("id"))
}

func (isap dataIothubSharedAccessPolicyAttributes) IothubName() terra.StringValue {
	return terra.ReferenceString(isap.ref.Append("iothub_name"))
}

func (isap dataIothubSharedAccessPolicyAttributes) Name() terra.StringValue {
	return terra.ReferenceString(isap.ref.Append("name"))
}

func (isap dataIothubSharedAccessPolicyAttributes) PrimaryConnectionString() terra.StringValue {
	return terra.ReferenceString(isap.ref.Append("primary_connection_string"))
}

func (isap dataIothubSharedAccessPolicyAttributes) PrimaryKey() terra.StringValue {
	return terra.ReferenceString(isap.ref.Append("primary_key"))
}

func (isap dataIothubSharedAccessPolicyAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceString(isap.ref.Append("resource_group_name"))
}

func (isap dataIothubSharedAccessPolicyAttributes) SecondaryConnectionString() terra.StringValue {
	return terra.ReferenceString(isap.ref.Append("secondary_connection_string"))
}

func (isap dataIothubSharedAccessPolicyAttributes) SecondaryKey() terra.StringValue {
	return terra.ReferenceString(isap.ref.Append("secondary_key"))
}

func (isap dataIothubSharedAccessPolicyAttributes) Timeouts() dataiothubsharedaccesspolicy.TimeoutsAttributes {
	return terra.ReferenceSingle[dataiothubsharedaccesspolicy.TimeoutsAttributes](isap.ref.Append("timeouts"))
}
