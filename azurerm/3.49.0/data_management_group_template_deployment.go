// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datamanagementgrouptemplatedeployment "github.com/golingon/terraproviders/azurerm/3.49.0/datamanagementgrouptemplatedeployment"
	"github.com/volvo-cars/lingon/pkg/terra"
)

func NewDataManagementGroupTemplateDeployment(name string, args DataManagementGroupTemplateDeploymentArgs) *DataManagementGroupTemplateDeployment {
	return &DataManagementGroupTemplateDeployment{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataManagementGroupTemplateDeployment)(nil)

type DataManagementGroupTemplateDeployment struct {
	Name string
	Args DataManagementGroupTemplateDeploymentArgs
}

func (mgtd *DataManagementGroupTemplateDeployment) DataSource() string {
	return "azurerm_management_group_template_deployment"
}

func (mgtd *DataManagementGroupTemplateDeployment) LocalName() string {
	return mgtd.Name
}

func (mgtd *DataManagementGroupTemplateDeployment) Configuration() interface{} {
	return mgtd.Args
}

func (mgtd *DataManagementGroupTemplateDeployment) Attributes() dataManagementGroupTemplateDeploymentAttributes {
	return dataManagementGroupTemplateDeploymentAttributes{ref: terra.ReferenceDataResource(mgtd)}
}

type DataManagementGroupTemplateDeploymentArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ManagementGroupId: string, required
	ManagementGroupId terra.StringValue `hcl:"management_group_id,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *datamanagementgrouptemplatedeployment.Timeouts `hcl:"timeouts,block"`
}
type dataManagementGroupTemplateDeploymentAttributes struct {
	ref terra.Reference
}

func (mgtd dataManagementGroupTemplateDeploymentAttributes) Id() terra.StringValue {
	return terra.ReferenceString(mgtd.ref.Append("id"))
}

func (mgtd dataManagementGroupTemplateDeploymentAttributes) ManagementGroupId() terra.StringValue {
	return terra.ReferenceString(mgtd.ref.Append("management_group_id"))
}

func (mgtd dataManagementGroupTemplateDeploymentAttributes) Name() terra.StringValue {
	return terra.ReferenceString(mgtd.ref.Append("name"))
}

func (mgtd dataManagementGroupTemplateDeploymentAttributes) OutputContent() terra.StringValue {
	return terra.ReferenceString(mgtd.ref.Append("output_content"))
}

func (mgtd dataManagementGroupTemplateDeploymentAttributes) Timeouts() datamanagementgrouptemplatedeployment.TimeoutsAttributes {
	return terra.ReferenceSingle[datamanagementgrouptemplatedeployment.TimeoutsAttributes](mgtd.ref.Append("timeouts"))
}
