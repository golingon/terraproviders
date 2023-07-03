// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datalogicappworkflow "github.com/golingon/terraproviders/azurerm/3.63.0/datalogicappworkflow"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataLogicAppWorkflow creates a new instance of [DataLogicAppWorkflow].
func NewDataLogicAppWorkflow(name string, args DataLogicAppWorkflowArgs) *DataLogicAppWorkflow {
	return &DataLogicAppWorkflow{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataLogicAppWorkflow)(nil)

// DataLogicAppWorkflow represents the Terraform data resource azurerm_logic_app_workflow.
type DataLogicAppWorkflow struct {
	Name string
	Args DataLogicAppWorkflowArgs
}

// DataSource returns the Terraform object type for [DataLogicAppWorkflow].
func (law *DataLogicAppWorkflow) DataSource() string {
	return "azurerm_logic_app_workflow"
}

// LocalName returns the local name for [DataLogicAppWorkflow].
func (law *DataLogicAppWorkflow) LocalName() string {
	return law.Name
}

// Configuration returns the configuration (args) for [DataLogicAppWorkflow].
func (law *DataLogicAppWorkflow) Configuration() interface{} {
	return law.Args
}

// Attributes returns the attributes for [DataLogicAppWorkflow].
func (law *DataLogicAppWorkflow) Attributes() dataLogicAppWorkflowAttributes {
	return dataLogicAppWorkflowAttributes{ref: terra.ReferenceDataResource(law)}
}

// DataLogicAppWorkflowArgs contains the configurations for azurerm_logic_app_workflow.
type DataLogicAppWorkflowArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Identity: min=0
	Identity []datalogicappworkflow.Identity `hcl:"identity,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *datalogicappworkflow.Timeouts `hcl:"timeouts,block"`
}
type dataLogicAppWorkflowAttributes struct {
	ref terra.Reference
}

// AccessEndpoint returns a reference to field access_endpoint of azurerm_logic_app_workflow.
func (law dataLogicAppWorkflowAttributes) AccessEndpoint() terra.StringValue {
	return terra.ReferenceAsString(law.ref.Append("access_endpoint"))
}

// ConnectorEndpointIpAddresses returns a reference to field connector_endpoint_ip_addresses of azurerm_logic_app_workflow.
func (law dataLogicAppWorkflowAttributes) ConnectorEndpointIpAddresses() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](law.ref.Append("connector_endpoint_ip_addresses"))
}

// ConnectorOutboundIpAddresses returns a reference to field connector_outbound_ip_addresses of azurerm_logic_app_workflow.
func (law dataLogicAppWorkflowAttributes) ConnectorOutboundIpAddresses() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](law.ref.Append("connector_outbound_ip_addresses"))
}

// Id returns a reference to field id of azurerm_logic_app_workflow.
func (law dataLogicAppWorkflowAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(law.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_logic_app_workflow.
func (law dataLogicAppWorkflowAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(law.ref.Append("location"))
}

// LogicAppIntegrationAccountId returns a reference to field logic_app_integration_account_id of azurerm_logic_app_workflow.
func (law dataLogicAppWorkflowAttributes) LogicAppIntegrationAccountId() terra.StringValue {
	return terra.ReferenceAsString(law.ref.Append("logic_app_integration_account_id"))
}

// Name returns a reference to field name of azurerm_logic_app_workflow.
func (law dataLogicAppWorkflowAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(law.ref.Append("name"))
}

// Parameters returns a reference to field parameters of azurerm_logic_app_workflow.
func (law dataLogicAppWorkflowAttributes) Parameters() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](law.ref.Append("parameters"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_logic_app_workflow.
func (law dataLogicAppWorkflowAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(law.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_logic_app_workflow.
func (law dataLogicAppWorkflowAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](law.ref.Append("tags"))
}

// WorkflowEndpointIpAddresses returns a reference to field workflow_endpoint_ip_addresses of azurerm_logic_app_workflow.
func (law dataLogicAppWorkflowAttributes) WorkflowEndpointIpAddresses() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](law.ref.Append("workflow_endpoint_ip_addresses"))
}

// WorkflowOutboundIpAddresses returns a reference to field workflow_outbound_ip_addresses of azurerm_logic_app_workflow.
func (law dataLogicAppWorkflowAttributes) WorkflowOutboundIpAddresses() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](law.ref.Append("workflow_outbound_ip_addresses"))
}

// WorkflowSchema returns a reference to field workflow_schema of azurerm_logic_app_workflow.
func (law dataLogicAppWorkflowAttributes) WorkflowSchema() terra.StringValue {
	return terra.ReferenceAsString(law.ref.Append("workflow_schema"))
}

// WorkflowVersion returns a reference to field workflow_version of azurerm_logic_app_workflow.
func (law dataLogicAppWorkflowAttributes) WorkflowVersion() terra.StringValue {
	return terra.ReferenceAsString(law.ref.Append("workflow_version"))
}

func (law dataLogicAppWorkflowAttributes) Identity() terra.ListValue[datalogicappworkflow.IdentityAttributes] {
	return terra.ReferenceAsList[datalogicappworkflow.IdentityAttributes](law.ref.Append("identity"))
}

func (law dataLogicAppWorkflowAttributes) Timeouts() datalogicappworkflow.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datalogicappworkflow.TimeoutsAttributes](law.ref.Append("timeouts"))
}