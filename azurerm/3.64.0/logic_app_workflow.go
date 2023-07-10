// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	logicappworkflow "github.com/golingon/terraproviders/azurerm/3.64.0/logicappworkflow"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewLogicAppWorkflow creates a new instance of [LogicAppWorkflow].
func NewLogicAppWorkflow(name string, args LogicAppWorkflowArgs) *LogicAppWorkflow {
	return &LogicAppWorkflow{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*LogicAppWorkflow)(nil)

// LogicAppWorkflow represents the Terraform resource azurerm_logic_app_workflow.
type LogicAppWorkflow struct {
	Name      string
	Args      LogicAppWorkflowArgs
	state     *logicAppWorkflowState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [LogicAppWorkflow].
func (law *LogicAppWorkflow) Type() string {
	return "azurerm_logic_app_workflow"
}

// LocalName returns the local name for [LogicAppWorkflow].
func (law *LogicAppWorkflow) LocalName() string {
	return law.Name
}

// Configuration returns the configuration (args) for [LogicAppWorkflow].
func (law *LogicAppWorkflow) Configuration() interface{} {
	return law.Args
}

// DependOn is used for other resources to depend on [LogicAppWorkflow].
func (law *LogicAppWorkflow) DependOn() terra.Reference {
	return terra.ReferenceResource(law)
}

// Dependencies returns the list of resources [LogicAppWorkflow] depends_on.
func (law *LogicAppWorkflow) Dependencies() terra.Dependencies {
	return law.DependsOn
}

// LifecycleManagement returns the lifecycle block for [LogicAppWorkflow].
func (law *LogicAppWorkflow) LifecycleManagement() *terra.Lifecycle {
	return law.Lifecycle
}

// Attributes returns the attributes for [LogicAppWorkflow].
func (law *LogicAppWorkflow) Attributes() logicAppWorkflowAttributes {
	return logicAppWorkflowAttributes{ref: terra.ReferenceResource(law)}
}

// ImportState imports the given attribute values into [LogicAppWorkflow]'s state.
func (law *LogicAppWorkflow) ImportState(av io.Reader) error {
	law.state = &logicAppWorkflowState{}
	if err := json.NewDecoder(av).Decode(law.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", law.Type(), law.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [LogicAppWorkflow] has state.
func (law *LogicAppWorkflow) State() (*logicAppWorkflowState, bool) {
	return law.state, law.state != nil
}

// StateMust returns the state for [LogicAppWorkflow]. Panics if the state is nil.
func (law *LogicAppWorkflow) StateMust() *logicAppWorkflowState {
	if law.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", law.Type(), law.LocalName()))
	}
	return law.state
}

// LogicAppWorkflowArgs contains the configurations for azurerm_logic_app_workflow.
type LogicAppWorkflowArgs struct {
	// Enabled: bool, optional
	Enabled terra.BoolValue `hcl:"enabled,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IntegrationServiceEnvironmentId: string, optional
	IntegrationServiceEnvironmentId terra.StringValue `hcl:"integration_service_environment_id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// LogicAppIntegrationAccountId: string, optional
	LogicAppIntegrationAccountId terra.StringValue `hcl:"logic_app_integration_account_id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Parameters: map of string, optional
	Parameters terra.MapValue[terra.StringValue] `hcl:"parameters,attr"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// WorkflowParameters: map of string, optional
	WorkflowParameters terra.MapValue[terra.StringValue] `hcl:"workflow_parameters,attr"`
	// WorkflowSchema: string, optional
	WorkflowSchema terra.StringValue `hcl:"workflow_schema,attr"`
	// WorkflowVersion: string, optional
	WorkflowVersion terra.StringValue `hcl:"workflow_version,attr"`
	// AccessControl: optional
	AccessControl *logicappworkflow.AccessControl `hcl:"access_control,block"`
	// Identity: optional
	Identity *logicappworkflow.Identity `hcl:"identity,block"`
	// Timeouts: optional
	Timeouts *logicappworkflow.Timeouts `hcl:"timeouts,block"`
}
type logicAppWorkflowAttributes struct {
	ref terra.Reference
}

// AccessEndpoint returns a reference to field access_endpoint of azurerm_logic_app_workflow.
func (law logicAppWorkflowAttributes) AccessEndpoint() terra.StringValue {
	return terra.ReferenceAsString(law.ref.Append("access_endpoint"))
}

// ConnectorEndpointIpAddresses returns a reference to field connector_endpoint_ip_addresses of azurerm_logic_app_workflow.
func (law logicAppWorkflowAttributes) ConnectorEndpointIpAddresses() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](law.ref.Append("connector_endpoint_ip_addresses"))
}

// ConnectorOutboundIpAddresses returns a reference to field connector_outbound_ip_addresses of azurerm_logic_app_workflow.
func (law logicAppWorkflowAttributes) ConnectorOutboundIpAddresses() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](law.ref.Append("connector_outbound_ip_addresses"))
}

// Enabled returns a reference to field enabled of azurerm_logic_app_workflow.
func (law logicAppWorkflowAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(law.ref.Append("enabled"))
}

// Id returns a reference to field id of azurerm_logic_app_workflow.
func (law logicAppWorkflowAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(law.ref.Append("id"))
}

// IntegrationServiceEnvironmentId returns a reference to field integration_service_environment_id of azurerm_logic_app_workflow.
func (law logicAppWorkflowAttributes) IntegrationServiceEnvironmentId() terra.StringValue {
	return terra.ReferenceAsString(law.ref.Append("integration_service_environment_id"))
}

// Location returns a reference to field location of azurerm_logic_app_workflow.
func (law logicAppWorkflowAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(law.ref.Append("location"))
}

// LogicAppIntegrationAccountId returns a reference to field logic_app_integration_account_id of azurerm_logic_app_workflow.
func (law logicAppWorkflowAttributes) LogicAppIntegrationAccountId() terra.StringValue {
	return terra.ReferenceAsString(law.ref.Append("logic_app_integration_account_id"))
}

// Name returns a reference to field name of azurerm_logic_app_workflow.
func (law logicAppWorkflowAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(law.ref.Append("name"))
}

// Parameters returns a reference to field parameters of azurerm_logic_app_workflow.
func (law logicAppWorkflowAttributes) Parameters() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](law.ref.Append("parameters"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_logic_app_workflow.
func (law logicAppWorkflowAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(law.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_logic_app_workflow.
func (law logicAppWorkflowAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](law.ref.Append("tags"))
}

// WorkflowEndpointIpAddresses returns a reference to field workflow_endpoint_ip_addresses of azurerm_logic_app_workflow.
func (law logicAppWorkflowAttributes) WorkflowEndpointIpAddresses() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](law.ref.Append("workflow_endpoint_ip_addresses"))
}

// WorkflowOutboundIpAddresses returns a reference to field workflow_outbound_ip_addresses of azurerm_logic_app_workflow.
func (law logicAppWorkflowAttributes) WorkflowOutboundIpAddresses() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](law.ref.Append("workflow_outbound_ip_addresses"))
}

// WorkflowParameters returns a reference to field workflow_parameters of azurerm_logic_app_workflow.
func (law logicAppWorkflowAttributes) WorkflowParameters() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](law.ref.Append("workflow_parameters"))
}

// WorkflowSchema returns a reference to field workflow_schema of azurerm_logic_app_workflow.
func (law logicAppWorkflowAttributes) WorkflowSchema() terra.StringValue {
	return terra.ReferenceAsString(law.ref.Append("workflow_schema"))
}

// WorkflowVersion returns a reference to field workflow_version of azurerm_logic_app_workflow.
func (law logicAppWorkflowAttributes) WorkflowVersion() terra.StringValue {
	return terra.ReferenceAsString(law.ref.Append("workflow_version"))
}

func (law logicAppWorkflowAttributes) AccessControl() terra.ListValue[logicappworkflow.AccessControlAttributes] {
	return terra.ReferenceAsList[logicappworkflow.AccessControlAttributes](law.ref.Append("access_control"))
}

func (law logicAppWorkflowAttributes) Identity() terra.ListValue[logicappworkflow.IdentityAttributes] {
	return terra.ReferenceAsList[logicappworkflow.IdentityAttributes](law.ref.Append("identity"))
}

func (law logicAppWorkflowAttributes) Timeouts() logicappworkflow.TimeoutsAttributes {
	return terra.ReferenceAsSingle[logicappworkflow.TimeoutsAttributes](law.ref.Append("timeouts"))
}

type logicAppWorkflowState struct {
	AccessEndpoint                  string                                `json:"access_endpoint"`
	ConnectorEndpointIpAddresses    []string                              `json:"connector_endpoint_ip_addresses"`
	ConnectorOutboundIpAddresses    []string                              `json:"connector_outbound_ip_addresses"`
	Enabled                         bool                                  `json:"enabled"`
	Id                              string                                `json:"id"`
	IntegrationServiceEnvironmentId string                                `json:"integration_service_environment_id"`
	Location                        string                                `json:"location"`
	LogicAppIntegrationAccountId    string                                `json:"logic_app_integration_account_id"`
	Name                            string                                `json:"name"`
	Parameters                      map[string]string                     `json:"parameters"`
	ResourceGroupName               string                                `json:"resource_group_name"`
	Tags                            map[string]string                     `json:"tags"`
	WorkflowEndpointIpAddresses     []string                              `json:"workflow_endpoint_ip_addresses"`
	WorkflowOutboundIpAddresses     []string                              `json:"workflow_outbound_ip_addresses"`
	WorkflowParameters              map[string]string                     `json:"workflow_parameters"`
	WorkflowSchema                  string                                `json:"workflow_schema"`
	WorkflowVersion                 string                                `json:"workflow_version"`
	AccessControl                   []logicappworkflow.AccessControlState `json:"access_control"`
	Identity                        []logicappworkflow.IdentityState      `json:"identity"`
	Timeouts                        *logicappworkflow.TimeoutsState       `json:"timeouts"`
}
