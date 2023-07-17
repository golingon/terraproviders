// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	dataiothubsharedaccesspolicy "github.com/golingon/terraproviders/azurerm/3.65.0/dataiothubsharedaccesspolicy"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataIothubSharedAccessPolicy creates a new instance of [DataIothubSharedAccessPolicy].
func NewDataIothubSharedAccessPolicy(name string, args DataIothubSharedAccessPolicyArgs) *DataIothubSharedAccessPolicy {
	return &DataIothubSharedAccessPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataIothubSharedAccessPolicy)(nil)

// DataIothubSharedAccessPolicy represents the Terraform data resource azurerm_iothub_shared_access_policy.
type DataIothubSharedAccessPolicy struct {
	Name string
	Args DataIothubSharedAccessPolicyArgs
}

// DataSource returns the Terraform object type for [DataIothubSharedAccessPolicy].
func (isap *DataIothubSharedAccessPolicy) DataSource() string {
	return "azurerm_iothub_shared_access_policy"
}

// LocalName returns the local name for [DataIothubSharedAccessPolicy].
func (isap *DataIothubSharedAccessPolicy) LocalName() string {
	return isap.Name
}

// Configuration returns the configuration (args) for [DataIothubSharedAccessPolicy].
func (isap *DataIothubSharedAccessPolicy) Configuration() interface{} {
	return isap.Args
}

// Attributes returns the attributes for [DataIothubSharedAccessPolicy].
func (isap *DataIothubSharedAccessPolicy) Attributes() dataIothubSharedAccessPolicyAttributes {
	return dataIothubSharedAccessPolicyAttributes{ref: terra.ReferenceDataResource(isap)}
}

// DataIothubSharedAccessPolicyArgs contains the configurations for azurerm_iothub_shared_access_policy.
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

// Id returns a reference to field id of azurerm_iothub_shared_access_policy.
func (isap dataIothubSharedAccessPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(isap.ref.Append("id"))
}

// IothubName returns a reference to field iothub_name of azurerm_iothub_shared_access_policy.
func (isap dataIothubSharedAccessPolicyAttributes) IothubName() terra.StringValue {
	return terra.ReferenceAsString(isap.ref.Append("iothub_name"))
}

// Name returns a reference to field name of azurerm_iothub_shared_access_policy.
func (isap dataIothubSharedAccessPolicyAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(isap.ref.Append("name"))
}

// PrimaryConnectionString returns a reference to field primary_connection_string of azurerm_iothub_shared_access_policy.
func (isap dataIothubSharedAccessPolicyAttributes) PrimaryConnectionString() terra.StringValue {
	return terra.ReferenceAsString(isap.ref.Append("primary_connection_string"))
}

// PrimaryKey returns a reference to field primary_key of azurerm_iothub_shared_access_policy.
func (isap dataIothubSharedAccessPolicyAttributes) PrimaryKey() terra.StringValue {
	return terra.ReferenceAsString(isap.ref.Append("primary_key"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_iothub_shared_access_policy.
func (isap dataIothubSharedAccessPolicyAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(isap.ref.Append("resource_group_name"))
}

// SecondaryConnectionString returns a reference to field secondary_connection_string of azurerm_iothub_shared_access_policy.
func (isap dataIothubSharedAccessPolicyAttributes) SecondaryConnectionString() terra.StringValue {
	return terra.ReferenceAsString(isap.ref.Append("secondary_connection_string"))
}

// SecondaryKey returns a reference to field secondary_key of azurerm_iothub_shared_access_policy.
func (isap dataIothubSharedAccessPolicyAttributes) SecondaryKey() terra.StringValue {
	return terra.ReferenceAsString(isap.ref.Append("secondary_key"))
}

func (isap dataIothubSharedAccessPolicyAttributes) Timeouts() dataiothubsharedaccesspolicy.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dataiothubsharedaccesspolicy.TimeoutsAttributes](isap.ref.Append("timeouts"))
}
