// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	sentineldataconnectormicrosoftcloudappsecurity "github.com/golingon/terraproviders/azurerm/3.98.0/sentineldataconnectormicrosoftcloudappsecurity"
	"io"
)

// NewSentinelDataConnectorMicrosoftCloudAppSecurity creates a new instance of [SentinelDataConnectorMicrosoftCloudAppSecurity].
func NewSentinelDataConnectorMicrosoftCloudAppSecurity(name string, args SentinelDataConnectorMicrosoftCloudAppSecurityArgs) *SentinelDataConnectorMicrosoftCloudAppSecurity {
	return &SentinelDataConnectorMicrosoftCloudAppSecurity{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SentinelDataConnectorMicrosoftCloudAppSecurity)(nil)

// SentinelDataConnectorMicrosoftCloudAppSecurity represents the Terraform resource azurerm_sentinel_data_connector_microsoft_cloud_app_security.
type SentinelDataConnectorMicrosoftCloudAppSecurity struct {
	Name      string
	Args      SentinelDataConnectorMicrosoftCloudAppSecurityArgs
	state     *sentinelDataConnectorMicrosoftCloudAppSecurityState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SentinelDataConnectorMicrosoftCloudAppSecurity].
func (sdcmcas *SentinelDataConnectorMicrosoftCloudAppSecurity) Type() string {
	return "azurerm_sentinel_data_connector_microsoft_cloud_app_security"
}

// LocalName returns the local name for [SentinelDataConnectorMicrosoftCloudAppSecurity].
func (sdcmcas *SentinelDataConnectorMicrosoftCloudAppSecurity) LocalName() string {
	return sdcmcas.Name
}

// Configuration returns the configuration (args) for [SentinelDataConnectorMicrosoftCloudAppSecurity].
func (sdcmcas *SentinelDataConnectorMicrosoftCloudAppSecurity) Configuration() interface{} {
	return sdcmcas.Args
}

// DependOn is used for other resources to depend on [SentinelDataConnectorMicrosoftCloudAppSecurity].
func (sdcmcas *SentinelDataConnectorMicrosoftCloudAppSecurity) DependOn() terra.Reference {
	return terra.ReferenceResource(sdcmcas)
}

// Dependencies returns the list of resources [SentinelDataConnectorMicrosoftCloudAppSecurity] depends_on.
func (sdcmcas *SentinelDataConnectorMicrosoftCloudAppSecurity) Dependencies() terra.Dependencies {
	return sdcmcas.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SentinelDataConnectorMicrosoftCloudAppSecurity].
func (sdcmcas *SentinelDataConnectorMicrosoftCloudAppSecurity) LifecycleManagement() *terra.Lifecycle {
	return sdcmcas.Lifecycle
}

// Attributes returns the attributes for [SentinelDataConnectorMicrosoftCloudAppSecurity].
func (sdcmcas *SentinelDataConnectorMicrosoftCloudAppSecurity) Attributes() sentinelDataConnectorMicrosoftCloudAppSecurityAttributes {
	return sentinelDataConnectorMicrosoftCloudAppSecurityAttributes{ref: terra.ReferenceResource(sdcmcas)}
}

// ImportState imports the given attribute values into [SentinelDataConnectorMicrosoftCloudAppSecurity]'s state.
func (sdcmcas *SentinelDataConnectorMicrosoftCloudAppSecurity) ImportState(av io.Reader) error {
	sdcmcas.state = &sentinelDataConnectorMicrosoftCloudAppSecurityState{}
	if err := json.NewDecoder(av).Decode(sdcmcas.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sdcmcas.Type(), sdcmcas.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SentinelDataConnectorMicrosoftCloudAppSecurity] has state.
func (sdcmcas *SentinelDataConnectorMicrosoftCloudAppSecurity) State() (*sentinelDataConnectorMicrosoftCloudAppSecurityState, bool) {
	return sdcmcas.state, sdcmcas.state != nil
}

// StateMust returns the state for [SentinelDataConnectorMicrosoftCloudAppSecurity]. Panics if the state is nil.
func (sdcmcas *SentinelDataConnectorMicrosoftCloudAppSecurity) StateMust() *sentinelDataConnectorMicrosoftCloudAppSecurityState {
	if sdcmcas.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sdcmcas.Type(), sdcmcas.LocalName()))
	}
	return sdcmcas.state
}

// SentinelDataConnectorMicrosoftCloudAppSecurityArgs contains the configurations for azurerm_sentinel_data_connector_microsoft_cloud_app_security.
type SentinelDataConnectorMicrosoftCloudAppSecurityArgs struct {
	// AlertsEnabled: bool, optional
	AlertsEnabled terra.BoolValue `hcl:"alerts_enabled,attr"`
	// DiscoveryLogsEnabled: bool, optional
	DiscoveryLogsEnabled terra.BoolValue `hcl:"discovery_logs_enabled,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// LogAnalyticsWorkspaceId: string, required
	LogAnalyticsWorkspaceId terra.StringValue `hcl:"log_analytics_workspace_id,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// TenantId: string, optional
	TenantId terra.StringValue `hcl:"tenant_id,attr"`
	// Timeouts: optional
	Timeouts *sentineldataconnectormicrosoftcloudappsecurity.Timeouts `hcl:"timeouts,block"`
}
type sentinelDataConnectorMicrosoftCloudAppSecurityAttributes struct {
	ref terra.Reference
}

// AlertsEnabled returns a reference to field alerts_enabled of azurerm_sentinel_data_connector_microsoft_cloud_app_security.
func (sdcmcas sentinelDataConnectorMicrosoftCloudAppSecurityAttributes) AlertsEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(sdcmcas.ref.Append("alerts_enabled"))
}

// DiscoveryLogsEnabled returns a reference to field discovery_logs_enabled of azurerm_sentinel_data_connector_microsoft_cloud_app_security.
func (sdcmcas sentinelDataConnectorMicrosoftCloudAppSecurityAttributes) DiscoveryLogsEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(sdcmcas.ref.Append("discovery_logs_enabled"))
}

// Id returns a reference to field id of azurerm_sentinel_data_connector_microsoft_cloud_app_security.
func (sdcmcas sentinelDataConnectorMicrosoftCloudAppSecurityAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sdcmcas.ref.Append("id"))
}

// LogAnalyticsWorkspaceId returns a reference to field log_analytics_workspace_id of azurerm_sentinel_data_connector_microsoft_cloud_app_security.
func (sdcmcas sentinelDataConnectorMicrosoftCloudAppSecurityAttributes) LogAnalyticsWorkspaceId() terra.StringValue {
	return terra.ReferenceAsString(sdcmcas.ref.Append("log_analytics_workspace_id"))
}

// Name returns a reference to field name of azurerm_sentinel_data_connector_microsoft_cloud_app_security.
func (sdcmcas sentinelDataConnectorMicrosoftCloudAppSecurityAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(sdcmcas.ref.Append("name"))
}

// TenantId returns a reference to field tenant_id of azurerm_sentinel_data_connector_microsoft_cloud_app_security.
func (sdcmcas sentinelDataConnectorMicrosoftCloudAppSecurityAttributes) TenantId() terra.StringValue {
	return terra.ReferenceAsString(sdcmcas.ref.Append("tenant_id"))
}

func (sdcmcas sentinelDataConnectorMicrosoftCloudAppSecurityAttributes) Timeouts() sentineldataconnectormicrosoftcloudappsecurity.TimeoutsAttributes {
	return terra.ReferenceAsSingle[sentineldataconnectormicrosoftcloudappsecurity.TimeoutsAttributes](sdcmcas.ref.Append("timeouts"))
}

type sentinelDataConnectorMicrosoftCloudAppSecurityState struct {
	AlertsEnabled           bool                                                          `json:"alerts_enabled"`
	DiscoveryLogsEnabled    bool                                                          `json:"discovery_logs_enabled"`
	Id                      string                                                        `json:"id"`
	LogAnalyticsWorkspaceId string                                                        `json:"log_analytics_workspace_id"`
	Name                    string                                                        `json:"name"`
	TenantId                string                                                        `json:"tenant_id"`
	Timeouts                *sentineldataconnectormicrosoftcloudappsecurity.TimeoutsState `json:"timeouts"`
}
