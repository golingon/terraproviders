// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	sentineldataconnectorazureactivedirectory "github.com/golingon/terraproviders/azurerm/3.63.0/sentineldataconnectorazureactivedirectory"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSentinelDataConnectorAzureActiveDirectory creates a new instance of [SentinelDataConnectorAzureActiveDirectory].
func NewSentinelDataConnectorAzureActiveDirectory(name string, args SentinelDataConnectorAzureActiveDirectoryArgs) *SentinelDataConnectorAzureActiveDirectory {
	return &SentinelDataConnectorAzureActiveDirectory{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SentinelDataConnectorAzureActiveDirectory)(nil)

// SentinelDataConnectorAzureActiveDirectory represents the Terraform resource azurerm_sentinel_data_connector_azure_active_directory.
type SentinelDataConnectorAzureActiveDirectory struct {
	Name      string
	Args      SentinelDataConnectorAzureActiveDirectoryArgs
	state     *sentinelDataConnectorAzureActiveDirectoryState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SentinelDataConnectorAzureActiveDirectory].
func (sdcaad *SentinelDataConnectorAzureActiveDirectory) Type() string {
	return "azurerm_sentinel_data_connector_azure_active_directory"
}

// LocalName returns the local name for [SentinelDataConnectorAzureActiveDirectory].
func (sdcaad *SentinelDataConnectorAzureActiveDirectory) LocalName() string {
	return sdcaad.Name
}

// Configuration returns the configuration (args) for [SentinelDataConnectorAzureActiveDirectory].
func (sdcaad *SentinelDataConnectorAzureActiveDirectory) Configuration() interface{} {
	return sdcaad.Args
}

// DependOn is used for other resources to depend on [SentinelDataConnectorAzureActiveDirectory].
func (sdcaad *SentinelDataConnectorAzureActiveDirectory) DependOn() terra.Reference {
	return terra.ReferenceResource(sdcaad)
}

// Dependencies returns the list of resources [SentinelDataConnectorAzureActiveDirectory] depends_on.
func (sdcaad *SentinelDataConnectorAzureActiveDirectory) Dependencies() terra.Dependencies {
	return sdcaad.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SentinelDataConnectorAzureActiveDirectory].
func (sdcaad *SentinelDataConnectorAzureActiveDirectory) LifecycleManagement() *terra.Lifecycle {
	return sdcaad.Lifecycle
}

// Attributes returns the attributes for [SentinelDataConnectorAzureActiveDirectory].
func (sdcaad *SentinelDataConnectorAzureActiveDirectory) Attributes() sentinelDataConnectorAzureActiveDirectoryAttributes {
	return sentinelDataConnectorAzureActiveDirectoryAttributes{ref: terra.ReferenceResource(sdcaad)}
}

// ImportState imports the given attribute values into [SentinelDataConnectorAzureActiveDirectory]'s state.
func (sdcaad *SentinelDataConnectorAzureActiveDirectory) ImportState(av io.Reader) error {
	sdcaad.state = &sentinelDataConnectorAzureActiveDirectoryState{}
	if err := json.NewDecoder(av).Decode(sdcaad.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sdcaad.Type(), sdcaad.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SentinelDataConnectorAzureActiveDirectory] has state.
func (sdcaad *SentinelDataConnectorAzureActiveDirectory) State() (*sentinelDataConnectorAzureActiveDirectoryState, bool) {
	return sdcaad.state, sdcaad.state != nil
}

// StateMust returns the state for [SentinelDataConnectorAzureActiveDirectory]. Panics if the state is nil.
func (sdcaad *SentinelDataConnectorAzureActiveDirectory) StateMust() *sentinelDataConnectorAzureActiveDirectoryState {
	if sdcaad.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sdcaad.Type(), sdcaad.LocalName()))
	}
	return sdcaad.state
}

// SentinelDataConnectorAzureActiveDirectoryArgs contains the configurations for azurerm_sentinel_data_connector_azure_active_directory.
type SentinelDataConnectorAzureActiveDirectoryArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// LogAnalyticsWorkspaceId: string, required
	LogAnalyticsWorkspaceId terra.StringValue `hcl:"log_analytics_workspace_id,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// TenantId: string, optional
	TenantId terra.StringValue `hcl:"tenant_id,attr"`
	// Timeouts: optional
	Timeouts *sentineldataconnectorazureactivedirectory.Timeouts `hcl:"timeouts,block"`
}
type sentinelDataConnectorAzureActiveDirectoryAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_sentinel_data_connector_azure_active_directory.
func (sdcaad sentinelDataConnectorAzureActiveDirectoryAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sdcaad.ref.Append("id"))
}

// LogAnalyticsWorkspaceId returns a reference to field log_analytics_workspace_id of azurerm_sentinel_data_connector_azure_active_directory.
func (sdcaad sentinelDataConnectorAzureActiveDirectoryAttributes) LogAnalyticsWorkspaceId() terra.StringValue {
	return terra.ReferenceAsString(sdcaad.ref.Append("log_analytics_workspace_id"))
}

// Name returns a reference to field name of azurerm_sentinel_data_connector_azure_active_directory.
func (sdcaad sentinelDataConnectorAzureActiveDirectoryAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(sdcaad.ref.Append("name"))
}

// TenantId returns a reference to field tenant_id of azurerm_sentinel_data_connector_azure_active_directory.
func (sdcaad sentinelDataConnectorAzureActiveDirectoryAttributes) TenantId() terra.StringValue {
	return terra.ReferenceAsString(sdcaad.ref.Append("tenant_id"))
}

func (sdcaad sentinelDataConnectorAzureActiveDirectoryAttributes) Timeouts() sentineldataconnectorazureactivedirectory.TimeoutsAttributes {
	return terra.ReferenceAsSingle[sentineldataconnectorazureactivedirectory.TimeoutsAttributes](sdcaad.ref.Append("timeouts"))
}

type sentinelDataConnectorAzureActiveDirectoryState struct {
	Id                      string                                                   `json:"id"`
	LogAnalyticsWorkspaceId string                                                   `json:"log_analytics_workspace_id"`
	Name                    string                                                   `json:"name"`
	TenantId                string                                                   `json:"tenant_id"`
	Timeouts                *sentineldataconnectorazureactivedirectory.TimeoutsState `json:"timeouts"`
}
