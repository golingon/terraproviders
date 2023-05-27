// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	sentineldataconnectormicrosoftthreatprotection "github.com/golingon/terraproviders/azurerm/3.58.0/sentineldataconnectormicrosoftthreatprotection"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSentinelDataConnectorMicrosoftThreatProtection creates a new instance of [SentinelDataConnectorMicrosoftThreatProtection].
func NewSentinelDataConnectorMicrosoftThreatProtection(name string, args SentinelDataConnectorMicrosoftThreatProtectionArgs) *SentinelDataConnectorMicrosoftThreatProtection {
	return &SentinelDataConnectorMicrosoftThreatProtection{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SentinelDataConnectorMicrosoftThreatProtection)(nil)

// SentinelDataConnectorMicrosoftThreatProtection represents the Terraform resource azurerm_sentinel_data_connector_microsoft_threat_protection.
type SentinelDataConnectorMicrosoftThreatProtection struct {
	Name      string
	Args      SentinelDataConnectorMicrosoftThreatProtectionArgs
	state     *sentinelDataConnectorMicrosoftThreatProtectionState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SentinelDataConnectorMicrosoftThreatProtection].
func (sdcmtp *SentinelDataConnectorMicrosoftThreatProtection) Type() string {
	return "azurerm_sentinel_data_connector_microsoft_threat_protection"
}

// LocalName returns the local name for [SentinelDataConnectorMicrosoftThreatProtection].
func (sdcmtp *SentinelDataConnectorMicrosoftThreatProtection) LocalName() string {
	return sdcmtp.Name
}

// Configuration returns the configuration (args) for [SentinelDataConnectorMicrosoftThreatProtection].
func (sdcmtp *SentinelDataConnectorMicrosoftThreatProtection) Configuration() interface{} {
	return sdcmtp.Args
}

// DependOn is used for other resources to depend on [SentinelDataConnectorMicrosoftThreatProtection].
func (sdcmtp *SentinelDataConnectorMicrosoftThreatProtection) DependOn() terra.Reference {
	return terra.ReferenceResource(sdcmtp)
}

// Dependencies returns the list of resources [SentinelDataConnectorMicrosoftThreatProtection] depends_on.
func (sdcmtp *SentinelDataConnectorMicrosoftThreatProtection) Dependencies() terra.Dependencies {
	return sdcmtp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SentinelDataConnectorMicrosoftThreatProtection].
func (sdcmtp *SentinelDataConnectorMicrosoftThreatProtection) LifecycleManagement() *terra.Lifecycle {
	return sdcmtp.Lifecycle
}

// Attributes returns the attributes for [SentinelDataConnectorMicrosoftThreatProtection].
func (sdcmtp *SentinelDataConnectorMicrosoftThreatProtection) Attributes() sentinelDataConnectorMicrosoftThreatProtectionAttributes {
	return sentinelDataConnectorMicrosoftThreatProtectionAttributes{ref: terra.ReferenceResource(sdcmtp)}
}

// ImportState imports the given attribute values into [SentinelDataConnectorMicrosoftThreatProtection]'s state.
func (sdcmtp *SentinelDataConnectorMicrosoftThreatProtection) ImportState(av io.Reader) error {
	sdcmtp.state = &sentinelDataConnectorMicrosoftThreatProtectionState{}
	if err := json.NewDecoder(av).Decode(sdcmtp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sdcmtp.Type(), sdcmtp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SentinelDataConnectorMicrosoftThreatProtection] has state.
func (sdcmtp *SentinelDataConnectorMicrosoftThreatProtection) State() (*sentinelDataConnectorMicrosoftThreatProtectionState, bool) {
	return sdcmtp.state, sdcmtp.state != nil
}

// StateMust returns the state for [SentinelDataConnectorMicrosoftThreatProtection]. Panics if the state is nil.
func (sdcmtp *SentinelDataConnectorMicrosoftThreatProtection) StateMust() *sentinelDataConnectorMicrosoftThreatProtectionState {
	if sdcmtp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sdcmtp.Type(), sdcmtp.LocalName()))
	}
	return sdcmtp.state
}

// SentinelDataConnectorMicrosoftThreatProtectionArgs contains the configurations for azurerm_sentinel_data_connector_microsoft_threat_protection.
type SentinelDataConnectorMicrosoftThreatProtectionArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// LogAnalyticsWorkspaceId: string, required
	LogAnalyticsWorkspaceId terra.StringValue `hcl:"log_analytics_workspace_id,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// TenantId: string, optional
	TenantId terra.StringValue `hcl:"tenant_id,attr"`
	// Timeouts: optional
	Timeouts *sentineldataconnectormicrosoftthreatprotection.Timeouts `hcl:"timeouts,block"`
}
type sentinelDataConnectorMicrosoftThreatProtectionAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_sentinel_data_connector_microsoft_threat_protection.
func (sdcmtp sentinelDataConnectorMicrosoftThreatProtectionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sdcmtp.ref.Append("id"))
}

// LogAnalyticsWorkspaceId returns a reference to field log_analytics_workspace_id of azurerm_sentinel_data_connector_microsoft_threat_protection.
func (sdcmtp sentinelDataConnectorMicrosoftThreatProtectionAttributes) LogAnalyticsWorkspaceId() terra.StringValue {
	return terra.ReferenceAsString(sdcmtp.ref.Append("log_analytics_workspace_id"))
}

// Name returns a reference to field name of azurerm_sentinel_data_connector_microsoft_threat_protection.
func (sdcmtp sentinelDataConnectorMicrosoftThreatProtectionAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(sdcmtp.ref.Append("name"))
}

// TenantId returns a reference to field tenant_id of azurerm_sentinel_data_connector_microsoft_threat_protection.
func (sdcmtp sentinelDataConnectorMicrosoftThreatProtectionAttributes) TenantId() terra.StringValue {
	return terra.ReferenceAsString(sdcmtp.ref.Append("tenant_id"))
}

func (sdcmtp sentinelDataConnectorMicrosoftThreatProtectionAttributes) Timeouts() sentineldataconnectormicrosoftthreatprotection.TimeoutsAttributes {
	return terra.ReferenceAsSingle[sentineldataconnectormicrosoftthreatprotection.TimeoutsAttributes](sdcmtp.ref.Append("timeouts"))
}

type sentinelDataConnectorMicrosoftThreatProtectionState struct {
	Id                      string                                                        `json:"id"`
	LogAnalyticsWorkspaceId string                                                        `json:"log_analytics_workspace_id"`
	Name                    string                                                        `json:"name"`
	TenantId                string                                                        `json:"tenant_id"`
	Timeouts                *sentineldataconnectormicrosoftthreatprotection.TimeoutsState `json:"timeouts"`
}
