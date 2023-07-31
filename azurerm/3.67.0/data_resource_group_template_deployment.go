// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	dataresourcegrouptemplatedeployment "github.com/golingon/terraproviders/azurerm/3.67.0/dataresourcegrouptemplatedeployment"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataResourceGroupTemplateDeployment creates a new instance of [DataResourceGroupTemplateDeployment].
func NewDataResourceGroupTemplateDeployment(name string, args DataResourceGroupTemplateDeploymentArgs) *DataResourceGroupTemplateDeployment {
	return &DataResourceGroupTemplateDeployment{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataResourceGroupTemplateDeployment)(nil)

// DataResourceGroupTemplateDeployment represents the Terraform data resource azurerm_resource_group_template_deployment.
type DataResourceGroupTemplateDeployment struct {
	Name string
	Args DataResourceGroupTemplateDeploymentArgs
}

// DataSource returns the Terraform object type for [DataResourceGroupTemplateDeployment].
func (rgtd *DataResourceGroupTemplateDeployment) DataSource() string {
	return "azurerm_resource_group_template_deployment"
}

// LocalName returns the local name for [DataResourceGroupTemplateDeployment].
func (rgtd *DataResourceGroupTemplateDeployment) LocalName() string {
	return rgtd.Name
}

// Configuration returns the configuration (args) for [DataResourceGroupTemplateDeployment].
func (rgtd *DataResourceGroupTemplateDeployment) Configuration() interface{} {
	return rgtd.Args
}

// Attributes returns the attributes for [DataResourceGroupTemplateDeployment].
func (rgtd *DataResourceGroupTemplateDeployment) Attributes() dataResourceGroupTemplateDeploymentAttributes {
	return dataResourceGroupTemplateDeploymentAttributes{ref: terra.ReferenceDataResource(rgtd)}
}

// DataResourceGroupTemplateDeploymentArgs contains the configurations for azurerm_resource_group_template_deployment.
type DataResourceGroupTemplateDeploymentArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *dataresourcegrouptemplatedeployment.Timeouts `hcl:"timeouts,block"`
}
type dataResourceGroupTemplateDeploymentAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_resource_group_template_deployment.
func (rgtd dataResourceGroupTemplateDeploymentAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(rgtd.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_resource_group_template_deployment.
func (rgtd dataResourceGroupTemplateDeploymentAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(rgtd.ref.Append("name"))
}

// OutputContent returns a reference to field output_content of azurerm_resource_group_template_deployment.
func (rgtd dataResourceGroupTemplateDeploymentAttributes) OutputContent() terra.StringValue {
	return terra.ReferenceAsString(rgtd.ref.Append("output_content"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_resource_group_template_deployment.
func (rgtd dataResourceGroupTemplateDeploymentAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(rgtd.ref.Append("resource_group_name"))
}

func (rgtd dataResourceGroupTemplateDeploymentAttributes) Timeouts() dataresourcegrouptemplatedeployment.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dataresourcegrouptemplatedeployment.TimeoutsAttributes](rgtd.ref.Append("timeouts"))
}
