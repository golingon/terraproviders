// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_sentinel_data_connector_azure_active_directory

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

// Resource represents the Terraform resource azurerm_sentinel_data_connector_azure_active_directory.
type Resource struct {
	Name      string
	Args      Args
	state     *azurermSentinelDataConnectorAzureActiveDirectoryState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (asdcaad *Resource) Type() string {
	return "azurerm_sentinel_data_connector_azure_active_directory"
}

// LocalName returns the local name for [Resource].
func (asdcaad *Resource) LocalName() string {
	return asdcaad.Name
}

// Configuration returns the configuration (args) for [Resource].
func (asdcaad *Resource) Configuration() interface{} {
	return asdcaad.Args
}

// DependOn is used for other resources to depend on [Resource].
func (asdcaad *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(asdcaad)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (asdcaad *Resource) Dependencies() terra.Dependencies {
	return asdcaad.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (asdcaad *Resource) LifecycleManagement() *terra.Lifecycle {
	return asdcaad.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (asdcaad *Resource) Attributes() azurermSentinelDataConnectorAzureActiveDirectoryAttributes {
	return azurermSentinelDataConnectorAzureActiveDirectoryAttributes{ref: terra.ReferenceResource(asdcaad)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (asdcaad *Resource) ImportState(state io.Reader) error {
	asdcaad.state = &azurermSentinelDataConnectorAzureActiveDirectoryState{}
	if err := json.NewDecoder(state).Decode(asdcaad.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", asdcaad.Type(), asdcaad.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (asdcaad *Resource) State() (*azurermSentinelDataConnectorAzureActiveDirectoryState, bool) {
	return asdcaad.state, asdcaad.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (asdcaad *Resource) StateMust() *azurermSentinelDataConnectorAzureActiveDirectoryState {
	if asdcaad.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", asdcaad.Type(), asdcaad.LocalName()))
	}
	return asdcaad.state
}

// Args contains the configurations for azurerm_sentinel_data_connector_azure_active_directory.
type Args struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// LogAnalyticsWorkspaceId: string, required
	LogAnalyticsWorkspaceId terra.StringValue `hcl:"log_analytics_workspace_id,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// TenantId: string, optional
	TenantId terra.StringValue `hcl:"tenant_id,attr"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type azurermSentinelDataConnectorAzureActiveDirectoryAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_sentinel_data_connector_azure_active_directory.
func (asdcaad azurermSentinelDataConnectorAzureActiveDirectoryAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(asdcaad.ref.Append("id"))
}

// LogAnalyticsWorkspaceId returns a reference to field log_analytics_workspace_id of azurerm_sentinel_data_connector_azure_active_directory.
func (asdcaad azurermSentinelDataConnectorAzureActiveDirectoryAttributes) LogAnalyticsWorkspaceId() terra.StringValue {
	return terra.ReferenceAsString(asdcaad.ref.Append("log_analytics_workspace_id"))
}

// Name returns a reference to field name of azurerm_sentinel_data_connector_azure_active_directory.
func (asdcaad azurermSentinelDataConnectorAzureActiveDirectoryAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(asdcaad.ref.Append("name"))
}

// TenantId returns a reference to field tenant_id of azurerm_sentinel_data_connector_azure_active_directory.
func (asdcaad azurermSentinelDataConnectorAzureActiveDirectoryAttributes) TenantId() terra.StringValue {
	return terra.ReferenceAsString(asdcaad.ref.Append("tenant_id"))
}

func (asdcaad azurermSentinelDataConnectorAzureActiveDirectoryAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](asdcaad.ref.Append("timeouts"))
}

type azurermSentinelDataConnectorAzureActiveDirectoryState struct {
	Id                      string         `json:"id"`
	LogAnalyticsWorkspaceId string         `json:"log_analytics_workspace_id"`
	Name                    string         `json:"name"`
	TenantId                string         `json:"tenant_id"`
	Timeouts                *TimeoutsState `json:"timeouts"`
}
