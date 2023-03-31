// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	dataavailabilityset "github.com/golingon/terraproviders/azurerm/3.49.0/dataavailabilityset"
	"github.com/volvo-cars/lingon/pkg/terra"
)

func NewDataAvailabilitySet(name string, args DataAvailabilitySetArgs) *DataAvailabilitySet {
	return &DataAvailabilitySet{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataAvailabilitySet)(nil)

type DataAvailabilitySet struct {
	Name string
	Args DataAvailabilitySetArgs
}

func (as *DataAvailabilitySet) DataSource() string {
	return "azurerm_availability_set"
}

func (as *DataAvailabilitySet) LocalName() string {
	return as.Name
}

func (as *DataAvailabilitySet) Configuration() interface{} {
	return as.Args
}

func (as *DataAvailabilitySet) Attributes() dataAvailabilitySetAttributes {
	return dataAvailabilitySetAttributes{ref: terra.ReferenceDataResource(as)}
}

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

func (as dataAvailabilitySetAttributes) Id() terra.StringValue {
	return terra.ReferenceString(as.ref.Append("id"))
}

func (as dataAvailabilitySetAttributes) Location() terra.StringValue {
	return terra.ReferenceString(as.ref.Append("location"))
}

func (as dataAvailabilitySetAttributes) Managed() terra.BoolValue {
	return terra.ReferenceBool(as.ref.Append("managed"))
}

func (as dataAvailabilitySetAttributes) Name() terra.StringValue {
	return terra.ReferenceString(as.ref.Append("name"))
}

func (as dataAvailabilitySetAttributes) PlatformFaultDomainCount() terra.NumberValue {
	return terra.ReferenceNumber(as.ref.Append("platform_fault_domain_count"))
}

func (as dataAvailabilitySetAttributes) PlatformUpdateDomainCount() terra.NumberValue {
	return terra.ReferenceNumber(as.ref.Append("platform_update_domain_count"))
}

func (as dataAvailabilitySetAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceString(as.ref.Append("resource_group_name"))
}

func (as dataAvailabilitySetAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](as.ref.Append("tags"))
}

func (as dataAvailabilitySetAttributes) Timeouts() dataavailabilityset.TimeoutsAttributes {
	return terra.ReferenceSingle[dataavailabilityset.TimeoutsAttributes](as.ref.Append("timeouts"))
}
