// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	dataiothubdpssharedaccesspolicy "github.com/golingon/terraproviders/azurerm/3.49.0/dataiothubdpssharedaccesspolicy"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataIothubDpsSharedAccessPolicy creates a new instance of [DataIothubDpsSharedAccessPolicy].
func NewDataIothubDpsSharedAccessPolicy(name string, args DataIothubDpsSharedAccessPolicyArgs) *DataIothubDpsSharedAccessPolicy {
	return &DataIothubDpsSharedAccessPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataIothubDpsSharedAccessPolicy)(nil)

// DataIothubDpsSharedAccessPolicy represents the Terraform data resource azurerm_iothub_dps_shared_access_policy.
type DataIothubDpsSharedAccessPolicy struct {
	Name string
	Args DataIothubDpsSharedAccessPolicyArgs
}

// DataSource returns the Terraform object type for [DataIothubDpsSharedAccessPolicy].
func (idsap *DataIothubDpsSharedAccessPolicy) DataSource() string {
	return "azurerm_iothub_dps_shared_access_policy"
}

// LocalName returns the local name for [DataIothubDpsSharedAccessPolicy].
func (idsap *DataIothubDpsSharedAccessPolicy) LocalName() string {
	return idsap.Name
}

// Configuration returns the configuration (args) for [DataIothubDpsSharedAccessPolicy].
func (idsap *DataIothubDpsSharedAccessPolicy) Configuration() interface{} {
	return idsap.Args
}

// Attributes returns the attributes for [DataIothubDpsSharedAccessPolicy].
func (idsap *DataIothubDpsSharedAccessPolicy) Attributes() dataIothubDpsSharedAccessPolicyAttributes {
	return dataIothubDpsSharedAccessPolicyAttributes{ref: terra.ReferenceDataResource(idsap)}
}

// DataIothubDpsSharedAccessPolicyArgs contains the configurations for azurerm_iothub_dps_shared_access_policy.
type DataIothubDpsSharedAccessPolicyArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IothubDpsName: string, required
	IothubDpsName terra.StringValue `hcl:"iothub_dps_name,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *dataiothubdpssharedaccesspolicy.Timeouts `hcl:"timeouts,block"`
}
type dataIothubDpsSharedAccessPolicyAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_iothub_dps_shared_access_policy.
func (idsap dataIothubDpsSharedAccessPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(idsap.ref.Append("id"))
}

// IothubDpsName returns a reference to field iothub_dps_name of azurerm_iothub_dps_shared_access_policy.
func (idsap dataIothubDpsSharedAccessPolicyAttributes) IothubDpsName() terra.StringValue {
	return terra.ReferenceAsString(idsap.ref.Append("iothub_dps_name"))
}

// Name returns a reference to field name of azurerm_iothub_dps_shared_access_policy.
func (idsap dataIothubDpsSharedAccessPolicyAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(idsap.ref.Append("name"))
}

// PrimaryConnectionString returns a reference to field primary_connection_string of azurerm_iothub_dps_shared_access_policy.
func (idsap dataIothubDpsSharedAccessPolicyAttributes) PrimaryConnectionString() terra.StringValue {
	return terra.ReferenceAsString(idsap.ref.Append("primary_connection_string"))
}

// PrimaryKey returns a reference to field primary_key of azurerm_iothub_dps_shared_access_policy.
func (idsap dataIothubDpsSharedAccessPolicyAttributes) PrimaryKey() terra.StringValue {
	return terra.ReferenceAsString(idsap.ref.Append("primary_key"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_iothub_dps_shared_access_policy.
func (idsap dataIothubDpsSharedAccessPolicyAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(idsap.ref.Append("resource_group_name"))
}

// SecondaryConnectionString returns a reference to field secondary_connection_string of azurerm_iothub_dps_shared_access_policy.
func (idsap dataIothubDpsSharedAccessPolicyAttributes) SecondaryConnectionString() terra.StringValue {
	return terra.ReferenceAsString(idsap.ref.Append("secondary_connection_string"))
}

// SecondaryKey returns a reference to field secondary_key of azurerm_iothub_dps_shared_access_policy.
func (idsap dataIothubDpsSharedAccessPolicyAttributes) SecondaryKey() terra.StringValue {
	return terra.ReferenceAsString(idsap.ref.Append("secondary_key"))
}

func (idsap dataIothubDpsSharedAccessPolicyAttributes) Timeouts() dataiothubdpssharedaccesspolicy.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dataiothubdpssharedaccesspolicy.TimeoutsAttributes](idsap.ref.Append("timeouts"))
}
