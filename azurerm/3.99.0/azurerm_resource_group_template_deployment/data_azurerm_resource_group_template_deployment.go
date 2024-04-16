// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_resource_group_template_deployment

import "github.com/golingon/lingon/pkg/terra"

// Data creates a new instance of [DataSource].
func Data(name string, args DataArgs) *DataSource {
	return &DataSource{
		Args: args,
		Name: name,
	}
}

var _ terra.DataSource = (*DataSource)(nil)

// DataSource represents the Terraform data resource azurerm_resource_group_template_deployment.
type DataSource struct {
	Name string
	Args DataArgs
}

// DataSource returns the Terraform object type for [DataSource].
func (argtd *DataSource) DataSource() string {
	return "azurerm_resource_group_template_deployment"
}

// LocalName returns the local name for [DataSource].
func (argtd *DataSource) LocalName() string {
	return argtd.Name
}

// Configuration returns the configuration (args) for [DataSource].
func (argtd *DataSource) Configuration() interface{} {
	return argtd.Args
}

// Attributes returns the attributes for [DataSource].
func (argtd *DataSource) Attributes() dataAzurermResourceGroupTemplateDeploymentAttributes {
	return dataAzurermResourceGroupTemplateDeploymentAttributes{ref: terra.ReferenceDataSource(argtd)}
}

// DataArgs contains the configurations for azurerm_resource_group_template_deployment.
type DataArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *DataTimeouts `hcl:"timeouts,block"`
}

type dataAzurermResourceGroupTemplateDeploymentAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_resource_group_template_deployment.
func (argtd dataAzurermResourceGroupTemplateDeploymentAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(argtd.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_resource_group_template_deployment.
func (argtd dataAzurermResourceGroupTemplateDeploymentAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(argtd.ref.Append("name"))
}

// OutputContent returns a reference to field output_content of azurerm_resource_group_template_deployment.
func (argtd dataAzurermResourceGroupTemplateDeploymentAttributes) OutputContent() terra.StringValue {
	return terra.ReferenceAsString(argtd.ref.Append("output_content"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_resource_group_template_deployment.
func (argtd dataAzurermResourceGroupTemplateDeploymentAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(argtd.ref.Append("resource_group_name"))
}

func (argtd dataAzurermResourceGroupTemplateDeploymentAttributes) Timeouts() DataTimeoutsAttributes {
	return terra.ReferenceAsSingle[DataTimeoutsAttributes](argtd.ref.Append("timeouts"))
}
