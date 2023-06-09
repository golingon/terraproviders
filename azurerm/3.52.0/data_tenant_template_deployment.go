// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datatenanttemplatedeployment "github.com/golingon/terraproviders/azurerm/3.52.0/datatenanttemplatedeployment"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataTenantTemplateDeployment creates a new instance of [DataTenantTemplateDeployment].
func NewDataTenantTemplateDeployment(name string, args DataTenantTemplateDeploymentArgs) *DataTenantTemplateDeployment {
	return &DataTenantTemplateDeployment{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataTenantTemplateDeployment)(nil)

// DataTenantTemplateDeployment represents the Terraform data resource azurerm_tenant_template_deployment.
type DataTenantTemplateDeployment struct {
	Name string
	Args DataTenantTemplateDeploymentArgs
}

// DataSource returns the Terraform object type for [DataTenantTemplateDeployment].
func (ttd *DataTenantTemplateDeployment) DataSource() string {
	return "azurerm_tenant_template_deployment"
}

// LocalName returns the local name for [DataTenantTemplateDeployment].
func (ttd *DataTenantTemplateDeployment) LocalName() string {
	return ttd.Name
}

// Configuration returns the configuration (args) for [DataTenantTemplateDeployment].
func (ttd *DataTenantTemplateDeployment) Configuration() interface{} {
	return ttd.Args
}

// Attributes returns the attributes for [DataTenantTemplateDeployment].
func (ttd *DataTenantTemplateDeployment) Attributes() dataTenantTemplateDeploymentAttributes {
	return dataTenantTemplateDeploymentAttributes{ref: terra.ReferenceDataResource(ttd)}
}

// DataTenantTemplateDeploymentArgs contains the configurations for azurerm_tenant_template_deployment.
type DataTenantTemplateDeploymentArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *datatenanttemplatedeployment.Timeouts `hcl:"timeouts,block"`
}
type dataTenantTemplateDeploymentAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_tenant_template_deployment.
func (ttd dataTenantTemplateDeploymentAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ttd.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_tenant_template_deployment.
func (ttd dataTenantTemplateDeploymentAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ttd.ref.Append("name"))
}

// OutputContent returns a reference to field output_content of azurerm_tenant_template_deployment.
func (ttd dataTenantTemplateDeploymentAttributes) OutputContent() terra.StringValue {
	return terra.ReferenceAsString(ttd.ref.Append("output_content"))
}

func (ttd dataTenantTemplateDeploymentAttributes) Timeouts() datatenanttemplatedeployment.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datatenanttemplatedeployment.TimeoutsAttributes](ttd.ref.Append("timeouts"))
}
