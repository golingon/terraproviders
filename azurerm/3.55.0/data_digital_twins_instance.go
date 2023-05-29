// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datadigitaltwinsinstance "github.com/golingon/terraproviders/azurerm/3.55.0/datadigitaltwinsinstance"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataDigitalTwinsInstance creates a new instance of [DataDigitalTwinsInstance].
func NewDataDigitalTwinsInstance(name string, args DataDigitalTwinsInstanceArgs) *DataDigitalTwinsInstance {
	return &DataDigitalTwinsInstance{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataDigitalTwinsInstance)(nil)

// DataDigitalTwinsInstance represents the Terraform data resource azurerm_digital_twins_instance.
type DataDigitalTwinsInstance struct {
	Name string
	Args DataDigitalTwinsInstanceArgs
}

// DataSource returns the Terraform object type for [DataDigitalTwinsInstance].
func (dti *DataDigitalTwinsInstance) DataSource() string {
	return "azurerm_digital_twins_instance"
}

// LocalName returns the local name for [DataDigitalTwinsInstance].
func (dti *DataDigitalTwinsInstance) LocalName() string {
	return dti.Name
}

// Configuration returns the configuration (args) for [DataDigitalTwinsInstance].
func (dti *DataDigitalTwinsInstance) Configuration() interface{} {
	return dti.Args
}

// Attributes returns the attributes for [DataDigitalTwinsInstance].
func (dti *DataDigitalTwinsInstance) Attributes() dataDigitalTwinsInstanceAttributes {
	return dataDigitalTwinsInstanceAttributes{ref: terra.ReferenceDataResource(dti)}
}

// DataDigitalTwinsInstanceArgs contains the configurations for azurerm_digital_twins_instance.
type DataDigitalTwinsInstanceArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *datadigitaltwinsinstance.Timeouts `hcl:"timeouts,block"`
}
type dataDigitalTwinsInstanceAttributes struct {
	ref terra.Reference
}

// HostName returns a reference to field host_name of azurerm_digital_twins_instance.
func (dti dataDigitalTwinsInstanceAttributes) HostName() terra.StringValue {
	return terra.ReferenceAsString(dti.ref.Append("host_name"))
}

// Id returns a reference to field id of azurerm_digital_twins_instance.
func (dti dataDigitalTwinsInstanceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dti.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_digital_twins_instance.
func (dti dataDigitalTwinsInstanceAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(dti.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_digital_twins_instance.
func (dti dataDigitalTwinsInstanceAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(dti.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_digital_twins_instance.
func (dti dataDigitalTwinsInstanceAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(dti.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_digital_twins_instance.
func (dti dataDigitalTwinsInstanceAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dti.ref.Append("tags"))
}

func (dti dataDigitalTwinsInstanceAttributes) Timeouts() datadigitaltwinsinstance.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datadigitaltwinsinstance.TimeoutsAttributes](dti.ref.Append("timeouts"))
}