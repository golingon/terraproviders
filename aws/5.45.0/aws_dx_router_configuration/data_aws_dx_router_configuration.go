// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_dx_router_configuration

import "github.com/golingon/lingon/pkg/terra"

// Data creates a new instance of [DataSource].
func Data(name string, args DataArgs) *DataSource {
	return &DataSource{
		Args: args,
		Name: name,
	}
}

var _ terra.DataSource = (*DataSource)(nil)

// DataSource represents the Terraform data resource aws_dx_router_configuration.
type DataSource struct {
	Name string
	Args DataArgs
}

// DataSource returns the Terraform object type for [DataSource].
func (adrc *DataSource) DataSource() string {
	return "aws_dx_router_configuration"
}

// LocalName returns the local name for [DataSource].
func (adrc *DataSource) LocalName() string {
	return adrc.Name
}

// Configuration returns the configuration (args) for [DataSource].
func (adrc *DataSource) Configuration() interface{} {
	return adrc.Args
}

// Attributes returns the attributes for [DataSource].
func (adrc *DataSource) Attributes() dataAwsDxRouterConfigurationAttributes {
	return dataAwsDxRouterConfigurationAttributes{ref: terra.ReferenceDataSource(adrc)}
}

// DataArgs contains the configurations for aws_dx_router_configuration.
type DataArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// RouterTypeIdentifier: string, required
	RouterTypeIdentifier terra.StringValue `hcl:"router_type_identifier,attr" validate:"required"`
	// VirtualInterfaceId: string, required
	VirtualInterfaceId terra.StringValue `hcl:"virtual_interface_id,attr" validate:"required"`
}

type dataAwsDxRouterConfigurationAttributes struct {
	ref terra.Reference
}

// CustomerRouterConfig returns a reference to field customer_router_config of aws_dx_router_configuration.
func (adrc dataAwsDxRouterConfigurationAttributes) CustomerRouterConfig() terra.StringValue {
	return terra.ReferenceAsString(adrc.ref.Append("customer_router_config"))
}

// Id returns a reference to field id of aws_dx_router_configuration.
func (adrc dataAwsDxRouterConfigurationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(adrc.ref.Append("id"))
}

// RouterTypeIdentifier returns a reference to field router_type_identifier of aws_dx_router_configuration.
func (adrc dataAwsDxRouterConfigurationAttributes) RouterTypeIdentifier() terra.StringValue {
	return terra.ReferenceAsString(adrc.ref.Append("router_type_identifier"))
}

// VirtualInterfaceId returns a reference to field virtual_interface_id of aws_dx_router_configuration.
func (adrc dataAwsDxRouterConfigurationAttributes) VirtualInterfaceId() terra.StringValue {
	return terra.ReferenceAsString(adrc.ref.Append("virtual_interface_id"))
}

// VirtualInterfaceName returns a reference to field virtual_interface_name of aws_dx_router_configuration.
func (adrc dataAwsDxRouterConfigurationAttributes) VirtualInterfaceName() terra.StringValue {
	return terra.ReferenceAsString(adrc.ref.Append("virtual_interface_name"))
}

func (adrc dataAwsDxRouterConfigurationAttributes) Router() terra.ListValue[DataRouterAttributes] {
	return terra.ReferenceAsList[DataRouterAttributes](adrc.ref.Append("router"))
}
