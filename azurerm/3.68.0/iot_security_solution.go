// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	iotsecuritysolution "github.com/golingon/terraproviders/azurerm/3.68.0/iotsecuritysolution"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewIotSecuritySolution creates a new instance of [IotSecuritySolution].
func NewIotSecuritySolution(name string, args IotSecuritySolutionArgs) *IotSecuritySolution {
	return &IotSecuritySolution{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*IotSecuritySolution)(nil)

// IotSecuritySolution represents the Terraform resource azurerm_iot_security_solution.
type IotSecuritySolution struct {
	Name      string
	Args      IotSecuritySolutionArgs
	state     *iotSecuritySolutionState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [IotSecuritySolution].
func (iss *IotSecuritySolution) Type() string {
	return "azurerm_iot_security_solution"
}

// LocalName returns the local name for [IotSecuritySolution].
func (iss *IotSecuritySolution) LocalName() string {
	return iss.Name
}

// Configuration returns the configuration (args) for [IotSecuritySolution].
func (iss *IotSecuritySolution) Configuration() interface{} {
	return iss.Args
}

// DependOn is used for other resources to depend on [IotSecuritySolution].
func (iss *IotSecuritySolution) DependOn() terra.Reference {
	return terra.ReferenceResource(iss)
}

// Dependencies returns the list of resources [IotSecuritySolution] depends_on.
func (iss *IotSecuritySolution) Dependencies() terra.Dependencies {
	return iss.DependsOn
}

// LifecycleManagement returns the lifecycle block for [IotSecuritySolution].
func (iss *IotSecuritySolution) LifecycleManagement() *terra.Lifecycle {
	return iss.Lifecycle
}

// Attributes returns the attributes for [IotSecuritySolution].
func (iss *IotSecuritySolution) Attributes() iotSecuritySolutionAttributes {
	return iotSecuritySolutionAttributes{ref: terra.ReferenceResource(iss)}
}

// ImportState imports the given attribute values into [IotSecuritySolution]'s state.
func (iss *IotSecuritySolution) ImportState(av io.Reader) error {
	iss.state = &iotSecuritySolutionState{}
	if err := json.NewDecoder(av).Decode(iss.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", iss.Type(), iss.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [IotSecuritySolution] has state.
func (iss *IotSecuritySolution) State() (*iotSecuritySolutionState, bool) {
	return iss.state, iss.state != nil
}

// StateMust returns the state for [IotSecuritySolution]. Panics if the state is nil.
func (iss *IotSecuritySolution) StateMust() *iotSecuritySolutionState {
	if iss.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", iss.Type(), iss.LocalName()))
	}
	return iss.state
}

// IotSecuritySolutionArgs contains the configurations for azurerm_iot_security_solution.
type IotSecuritySolutionArgs struct {
	// DisabledDataSources: set of string, optional
	DisabledDataSources terra.SetValue[terra.StringValue] `hcl:"disabled_data_sources,attr"`
	// DisplayName: string, required
	DisplayName terra.StringValue `hcl:"display_name,attr" validate:"required"`
	// Enabled: bool, optional
	Enabled terra.BoolValue `hcl:"enabled,attr"`
	// EventsToExport: set of string, optional
	EventsToExport terra.SetValue[terra.StringValue] `hcl:"events_to_export,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IothubIds: set of string, required
	IothubIds terra.SetValue[terra.StringValue] `hcl:"iothub_ids,attr" validate:"required"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// LogAnalyticsWorkspaceId: string, optional
	LogAnalyticsWorkspaceId terra.StringValue `hcl:"log_analytics_workspace_id,attr"`
	// LogUnmaskedIpsEnabled: bool, optional
	LogUnmaskedIpsEnabled terra.BoolValue `hcl:"log_unmasked_ips_enabled,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// QueryForResources: string, optional
	QueryForResources terra.StringValue `hcl:"query_for_resources,attr"`
	// QuerySubscriptionIds: set of string, optional
	QuerySubscriptionIds terra.SetValue[terra.StringValue] `hcl:"query_subscription_ids,attr"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// AdditionalWorkspace: min=0
	AdditionalWorkspace []iotsecuritysolution.AdditionalWorkspace `hcl:"additional_workspace,block" validate:"min=0"`
	// RecommendationsEnabled: optional
	RecommendationsEnabled *iotsecuritysolution.RecommendationsEnabled `hcl:"recommendations_enabled,block"`
	// Timeouts: optional
	Timeouts *iotsecuritysolution.Timeouts `hcl:"timeouts,block"`
}
type iotSecuritySolutionAttributes struct {
	ref terra.Reference
}

// DisabledDataSources returns a reference to field disabled_data_sources of azurerm_iot_security_solution.
func (iss iotSecuritySolutionAttributes) DisabledDataSources() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](iss.ref.Append("disabled_data_sources"))
}

// DisplayName returns a reference to field display_name of azurerm_iot_security_solution.
func (iss iotSecuritySolutionAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(iss.ref.Append("display_name"))
}

// Enabled returns a reference to field enabled of azurerm_iot_security_solution.
func (iss iotSecuritySolutionAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(iss.ref.Append("enabled"))
}

// EventsToExport returns a reference to field events_to_export of azurerm_iot_security_solution.
func (iss iotSecuritySolutionAttributes) EventsToExport() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](iss.ref.Append("events_to_export"))
}

// Id returns a reference to field id of azurerm_iot_security_solution.
func (iss iotSecuritySolutionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(iss.ref.Append("id"))
}

// IothubIds returns a reference to field iothub_ids of azurerm_iot_security_solution.
func (iss iotSecuritySolutionAttributes) IothubIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](iss.ref.Append("iothub_ids"))
}

// Location returns a reference to field location of azurerm_iot_security_solution.
func (iss iotSecuritySolutionAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(iss.ref.Append("location"))
}

// LogAnalyticsWorkspaceId returns a reference to field log_analytics_workspace_id of azurerm_iot_security_solution.
func (iss iotSecuritySolutionAttributes) LogAnalyticsWorkspaceId() terra.StringValue {
	return terra.ReferenceAsString(iss.ref.Append("log_analytics_workspace_id"))
}

// LogUnmaskedIpsEnabled returns a reference to field log_unmasked_ips_enabled of azurerm_iot_security_solution.
func (iss iotSecuritySolutionAttributes) LogUnmaskedIpsEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(iss.ref.Append("log_unmasked_ips_enabled"))
}

// Name returns a reference to field name of azurerm_iot_security_solution.
func (iss iotSecuritySolutionAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(iss.ref.Append("name"))
}

// QueryForResources returns a reference to field query_for_resources of azurerm_iot_security_solution.
func (iss iotSecuritySolutionAttributes) QueryForResources() terra.StringValue {
	return terra.ReferenceAsString(iss.ref.Append("query_for_resources"))
}

// QuerySubscriptionIds returns a reference to field query_subscription_ids of azurerm_iot_security_solution.
func (iss iotSecuritySolutionAttributes) QuerySubscriptionIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](iss.ref.Append("query_subscription_ids"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_iot_security_solution.
func (iss iotSecuritySolutionAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(iss.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_iot_security_solution.
func (iss iotSecuritySolutionAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](iss.ref.Append("tags"))
}

func (iss iotSecuritySolutionAttributes) AdditionalWorkspace() terra.SetValue[iotsecuritysolution.AdditionalWorkspaceAttributes] {
	return terra.ReferenceAsSet[iotsecuritysolution.AdditionalWorkspaceAttributes](iss.ref.Append("additional_workspace"))
}

func (iss iotSecuritySolutionAttributes) RecommendationsEnabled() terra.ListValue[iotsecuritysolution.RecommendationsEnabledAttributes] {
	return terra.ReferenceAsList[iotsecuritysolution.RecommendationsEnabledAttributes](iss.ref.Append("recommendations_enabled"))
}

func (iss iotSecuritySolutionAttributes) Timeouts() iotsecuritysolution.TimeoutsAttributes {
	return terra.ReferenceAsSingle[iotsecuritysolution.TimeoutsAttributes](iss.ref.Append("timeouts"))
}

type iotSecuritySolutionState struct {
	DisabledDataSources     []string                                          `json:"disabled_data_sources"`
	DisplayName             string                                            `json:"display_name"`
	Enabled                 bool                                              `json:"enabled"`
	EventsToExport          []string                                          `json:"events_to_export"`
	Id                      string                                            `json:"id"`
	IothubIds               []string                                          `json:"iothub_ids"`
	Location                string                                            `json:"location"`
	LogAnalyticsWorkspaceId string                                            `json:"log_analytics_workspace_id"`
	LogUnmaskedIpsEnabled   bool                                              `json:"log_unmasked_ips_enabled"`
	Name                    string                                            `json:"name"`
	QueryForResources       string                                            `json:"query_for_resources"`
	QuerySubscriptionIds    []string                                          `json:"query_subscription_ids"`
	ResourceGroupName       string                                            `json:"resource_group_name"`
	Tags                    map[string]string                                 `json:"tags"`
	AdditionalWorkspace     []iotsecuritysolution.AdditionalWorkspaceState    `json:"additional_workspace"`
	RecommendationsEnabled  []iotsecuritysolution.RecommendationsEnabledState `json:"recommendations_enabled"`
	Timeouts                *iotsecuritysolution.TimeoutsState                `json:"timeouts"`
}
