// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_data_factory_integration_runtime_azure

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// New creates a new instance of [Resource].
func New(name string, args Args) *Resource {
	return &Resource{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Resource)(nil)

// Resource represents the Terraform resource azurerm_data_factory_integration_runtime_azure.
type Resource struct {
	Name      string
	Args      Args
	state     *azurermDataFactoryIntegrationRuntimeAzureState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (adfira *Resource) Type() string {
	return "azurerm_data_factory_integration_runtime_azure"
}

// LocalName returns the local name for [Resource].
func (adfira *Resource) LocalName() string {
	return adfira.Name
}

// Configuration returns the configuration (args) for [Resource].
func (adfira *Resource) Configuration() interface{} {
	return adfira.Args
}

// DependOn is used for other resources to depend on [Resource].
func (adfira *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(adfira)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (adfira *Resource) Dependencies() terra.Dependencies {
	return adfira.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (adfira *Resource) LifecycleManagement() *terra.Lifecycle {
	return adfira.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (adfira *Resource) Attributes() azurermDataFactoryIntegrationRuntimeAzureAttributes {
	return azurermDataFactoryIntegrationRuntimeAzureAttributes{ref: terra.ReferenceResource(adfira)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (adfira *Resource) ImportState(state io.Reader) error {
	adfira.state = &azurermDataFactoryIntegrationRuntimeAzureState{}
	if err := json.NewDecoder(state).Decode(adfira.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", adfira.Type(), adfira.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (adfira *Resource) State() (*azurermDataFactoryIntegrationRuntimeAzureState, bool) {
	return adfira.state, adfira.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (adfira *Resource) StateMust() *azurermDataFactoryIntegrationRuntimeAzureState {
	if adfira.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", adfira.Type(), adfira.LocalName()))
	}
	return adfira.state
}

// Args contains the configurations for azurerm_data_factory_integration_runtime_azure.
type Args struct {
	// CleanupEnabled: bool, optional
	CleanupEnabled terra.BoolValue `hcl:"cleanup_enabled,attr"`
	// ComputeType: string, optional
	ComputeType terra.StringValue `hcl:"compute_type,attr"`
	// CoreCount: number, optional
	CoreCount terra.NumberValue `hcl:"core_count,attr"`
	// DataFactoryId: string, required
	DataFactoryId terra.StringValue `hcl:"data_factory_id,attr" validate:"required"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// TimeToLiveMin: number, optional
	TimeToLiveMin terra.NumberValue `hcl:"time_to_live_min,attr"`
	// VirtualNetworkEnabled: bool, optional
	VirtualNetworkEnabled terra.BoolValue `hcl:"virtual_network_enabled,attr"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type azurermDataFactoryIntegrationRuntimeAzureAttributes struct {
	ref terra.Reference
}

// CleanupEnabled returns a reference to field cleanup_enabled of azurerm_data_factory_integration_runtime_azure.
func (adfira azurermDataFactoryIntegrationRuntimeAzureAttributes) CleanupEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(adfira.ref.Append("cleanup_enabled"))
}

// ComputeType returns a reference to field compute_type of azurerm_data_factory_integration_runtime_azure.
func (adfira azurermDataFactoryIntegrationRuntimeAzureAttributes) ComputeType() terra.StringValue {
	return terra.ReferenceAsString(adfira.ref.Append("compute_type"))
}

// CoreCount returns a reference to field core_count of azurerm_data_factory_integration_runtime_azure.
func (adfira azurermDataFactoryIntegrationRuntimeAzureAttributes) CoreCount() terra.NumberValue {
	return terra.ReferenceAsNumber(adfira.ref.Append("core_count"))
}

// DataFactoryId returns a reference to field data_factory_id of azurerm_data_factory_integration_runtime_azure.
func (adfira azurermDataFactoryIntegrationRuntimeAzureAttributes) DataFactoryId() terra.StringValue {
	return terra.ReferenceAsString(adfira.ref.Append("data_factory_id"))
}

// Description returns a reference to field description of azurerm_data_factory_integration_runtime_azure.
func (adfira azurermDataFactoryIntegrationRuntimeAzureAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(adfira.ref.Append("description"))
}

// Id returns a reference to field id of azurerm_data_factory_integration_runtime_azure.
func (adfira azurermDataFactoryIntegrationRuntimeAzureAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(adfira.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_data_factory_integration_runtime_azure.
func (adfira azurermDataFactoryIntegrationRuntimeAzureAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(adfira.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_data_factory_integration_runtime_azure.
func (adfira azurermDataFactoryIntegrationRuntimeAzureAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(adfira.ref.Append("name"))
}

// TimeToLiveMin returns a reference to field time_to_live_min of azurerm_data_factory_integration_runtime_azure.
func (adfira azurermDataFactoryIntegrationRuntimeAzureAttributes) TimeToLiveMin() terra.NumberValue {
	return terra.ReferenceAsNumber(adfira.ref.Append("time_to_live_min"))
}

// VirtualNetworkEnabled returns a reference to field virtual_network_enabled of azurerm_data_factory_integration_runtime_azure.
func (adfira azurermDataFactoryIntegrationRuntimeAzureAttributes) VirtualNetworkEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(adfira.ref.Append("virtual_network_enabled"))
}

func (adfira azurermDataFactoryIntegrationRuntimeAzureAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](adfira.ref.Append("timeouts"))
}

type azurermDataFactoryIntegrationRuntimeAzureState struct {
	CleanupEnabled        bool           `json:"cleanup_enabled"`
	ComputeType           string         `json:"compute_type"`
	CoreCount             float64        `json:"core_count"`
	DataFactoryId         string         `json:"data_factory_id"`
	Description           string         `json:"description"`
	Id                    string         `json:"id"`
	Location              string         `json:"location"`
	Name                  string         `json:"name"`
	TimeToLiveMin         float64        `json:"time_to_live_min"`
	VirtualNetworkEnabled bool           `json:"virtual_network_enabled"`
	Timeouts              *TimeoutsState `json:"timeouts"`
}
