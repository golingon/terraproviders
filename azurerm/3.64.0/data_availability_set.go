// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	dataavailabilityset "github.com/golingon/terraproviders/azurerm/3.64.0/dataavailabilityset"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataAvailabilitySet creates a new instance of [DataAvailabilitySet].
func NewDataAvailabilitySet(name string, args DataAvailabilitySetArgs) *DataAvailabilitySet {
	return &DataAvailabilitySet{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataAvailabilitySet)(nil)

// DataAvailabilitySet represents the Terraform data resource azurerm_availability_set.
type DataAvailabilitySet struct {
	Name string
	Args DataAvailabilitySetArgs
}

// DataSource returns the Terraform object type for [DataAvailabilitySet].
func (as *DataAvailabilitySet) DataSource() string {
	return "azurerm_availability_set"
}

// LocalName returns the local name for [DataAvailabilitySet].
func (as *DataAvailabilitySet) LocalName() string {
	return as.Name
}

// Configuration returns the configuration (args) for [DataAvailabilitySet].
func (as *DataAvailabilitySet) Configuration() interface{} {
	return as.Args
}

// Attributes returns the attributes for [DataAvailabilitySet].
func (as *DataAvailabilitySet) Attributes() dataAvailabilitySetAttributes {
	return dataAvailabilitySetAttributes{ref: terra.ReferenceDataResource(as)}
}

// DataAvailabilitySetArgs contains the configurations for azurerm_availability_set.
type DataAvailabilitySetArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *dataavailabilityset.Timeouts `hcl:"timeouts,block"`
}
type dataAvailabilitySetAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_availability_set.
func (as dataAvailabilitySetAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(as.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_availability_set.
func (as dataAvailabilitySetAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(as.ref.Append("location"))
}

// Managed returns a reference to field managed of azurerm_availability_set.
func (as dataAvailabilitySetAttributes) Managed() terra.BoolValue {
	return terra.ReferenceAsBool(as.ref.Append("managed"))
}

// Name returns a reference to field name of azurerm_availability_set.
func (as dataAvailabilitySetAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(as.ref.Append("name"))
}

// PlatformFaultDomainCount returns a reference to field platform_fault_domain_count of azurerm_availability_set.
func (as dataAvailabilitySetAttributes) PlatformFaultDomainCount() terra.NumberValue {
	return terra.ReferenceAsNumber(as.ref.Append("platform_fault_domain_count"))
}

// PlatformUpdateDomainCount returns a reference to field platform_update_domain_count of azurerm_availability_set.
func (as dataAvailabilitySetAttributes) PlatformUpdateDomainCount() terra.NumberValue {
	return terra.ReferenceAsNumber(as.ref.Append("platform_update_domain_count"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_availability_set.
func (as dataAvailabilitySetAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(as.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_availability_set.
func (as dataAvailabilitySetAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](as.ref.Append("tags"))
}

func (as dataAvailabilitySetAttributes) Timeouts() dataavailabilityset.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dataavailabilityset.TimeoutsAttributes](as.ref.Append("timeouts"))
}
