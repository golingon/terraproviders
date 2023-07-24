// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	datadxrouterconfiguration "github.com/golingon/terraproviders/aws/5.9.0/datadxrouterconfiguration"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataDxRouterConfiguration creates a new instance of [DataDxRouterConfiguration].
func NewDataDxRouterConfiguration(name string, args DataDxRouterConfigurationArgs) *DataDxRouterConfiguration {
	return &DataDxRouterConfiguration{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataDxRouterConfiguration)(nil)

// DataDxRouterConfiguration represents the Terraform data resource aws_dx_router_configuration.
type DataDxRouterConfiguration struct {
	Name string
	Args DataDxRouterConfigurationArgs
}

// DataSource returns the Terraform object type for [DataDxRouterConfiguration].
func (drc *DataDxRouterConfiguration) DataSource() string {
	return "aws_dx_router_configuration"
}

// LocalName returns the local name for [DataDxRouterConfiguration].
func (drc *DataDxRouterConfiguration) LocalName() string {
	return drc.Name
}

// Configuration returns the configuration (args) for [DataDxRouterConfiguration].
func (drc *DataDxRouterConfiguration) Configuration() interface{} {
	return drc.Args
}

// Attributes returns the attributes for [DataDxRouterConfiguration].
func (drc *DataDxRouterConfiguration) Attributes() dataDxRouterConfigurationAttributes {
	return dataDxRouterConfigurationAttributes{ref: terra.ReferenceDataResource(drc)}
}

// DataDxRouterConfigurationArgs contains the configurations for aws_dx_router_configuration.
type DataDxRouterConfigurationArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// RouterTypeIdentifier: string, required
	RouterTypeIdentifier terra.StringValue `hcl:"router_type_identifier,attr" validate:"required"`
	// VirtualInterfaceId: string, required
	VirtualInterfaceId terra.StringValue `hcl:"virtual_interface_id,attr" validate:"required"`
	// Router: min=0
	Router []datadxrouterconfiguration.Router `hcl:"router,block" validate:"min=0"`
}
type dataDxRouterConfigurationAttributes struct {
	ref terra.Reference
}

// CustomerRouterConfig returns a reference to field customer_router_config of aws_dx_router_configuration.
func (drc dataDxRouterConfigurationAttributes) CustomerRouterConfig() terra.StringValue {
	return terra.ReferenceAsString(drc.ref.Append("customer_router_config"))
}

// Id returns a reference to field id of aws_dx_router_configuration.
func (drc dataDxRouterConfigurationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(drc.ref.Append("id"))
}

// RouterTypeIdentifier returns a reference to field router_type_identifier of aws_dx_router_configuration.
func (drc dataDxRouterConfigurationAttributes) RouterTypeIdentifier() terra.StringValue {
	return terra.ReferenceAsString(drc.ref.Append("router_type_identifier"))
}

// VirtualInterfaceId returns a reference to field virtual_interface_id of aws_dx_router_configuration.
func (drc dataDxRouterConfigurationAttributes) VirtualInterfaceId() terra.StringValue {
	return terra.ReferenceAsString(drc.ref.Append("virtual_interface_id"))
}

// VirtualInterfaceName returns a reference to field virtual_interface_name of aws_dx_router_configuration.
func (drc dataDxRouterConfigurationAttributes) VirtualInterfaceName() terra.StringValue {
	return terra.ReferenceAsString(drc.ref.Append("virtual_interface_name"))
}

func (drc dataDxRouterConfigurationAttributes) Router() terra.ListValue[datadxrouterconfiguration.RouterAttributes] {
	return terra.ReferenceAsList[datadxrouterconfiguration.RouterAttributes](drc.ref.Append("router"))
}
