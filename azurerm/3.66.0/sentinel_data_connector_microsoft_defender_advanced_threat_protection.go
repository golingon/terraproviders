// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	sentineldataconnectormicrosoftdefenderadvancedthreatprotection "github.com/golingon/terraproviders/azurerm/3.66.0/sentineldataconnectormicrosoftdefenderadvancedthreatprotection"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSentinelDataConnectorMicrosoftDefenderAdvancedThreatProtection creates a new instance of [SentinelDataConnectorMicrosoftDefenderAdvancedThreatProtection].
func NewSentinelDataConnectorMicrosoftDefenderAdvancedThreatProtection(name string, args SentinelDataConnectorMicrosoftDefenderAdvancedThreatProtectionArgs) *SentinelDataConnectorMicrosoftDefenderAdvancedThreatProtection {
	return &SentinelDataConnectorMicrosoftDefenderAdvancedThreatProtection{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SentinelDataConnectorMicrosoftDefenderAdvancedThreatProtection)(nil)

// SentinelDataConnectorMicrosoftDefenderAdvancedThreatProtection represents the Terraform resource azurerm_sentinel_data_connector_microsoft_defender_advanced_threat_protection.
type SentinelDataConnectorMicrosoftDefenderAdvancedThreatProtection struct {
	Name      string
	Args      SentinelDataConnectorMicrosoftDefenderAdvancedThreatProtectionArgs
	state     *sentinelDataConnectorMicrosoftDefenderAdvancedThreatProtectionState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SentinelDataConnectorMicrosoftDefenderAdvancedThreatProtection].
func (sdcmdatp *SentinelDataConnectorMicrosoftDefenderAdvancedThreatProtection) Type() string {
	return "azurerm_sentinel_data_connector_microsoft_defender_advanced_threat_protection"
}

// LocalName returns the local name for [SentinelDataConnectorMicrosoftDefenderAdvancedThreatProtection].
func (sdcmdatp *SentinelDataConnectorMicrosoftDefenderAdvancedThreatProtection) LocalName() string {
	return sdcmdatp.Name
}

// Configuration returns the configuration (args) for [SentinelDataConnectorMicrosoftDefenderAdvancedThreatProtection].
func (sdcmdatp *SentinelDataConnectorMicrosoftDefenderAdvancedThreatProtection) Configuration() interface{} {
	return sdcmdatp.Args
}

// DependOn is used for other resources to depend on [SentinelDataConnectorMicrosoftDefenderAdvancedThreatProtection].
func (sdcmdatp *SentinelDataConnectorMicrosoftDefenderAdvancedThreatProtection) DependOn() terra.Reference {
	return terra.ReferenceResource(sdcmdatp)
}

// Dependencies returns the list of resources [SentinelDataConnectorMicrosoftDefenderAdvancedThreatProtection] depends_on.
func (sdcmdatp *SentinelDataConnectorMicrosoftDefenderAdvancedThreatProtection) Dependencies() terra.Dependencies {
	return sdcmdatp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SentinelDataConnectorMicrosoftDefenderAdvancedThreatProtection].
func (sdcmdatp *SentinelDataConnectorMicrosoftDefenderAdvancedThreatProtection) LifecycleManagement() *terra.Lifecycle {
	return sdcmdatp.Lifecycle
}

// Attributes returns the attributes for [SentinelDataConnectorMicrosoftDefenderAdvancedThreatProtection].
func (sdcmdatp *SentinelDataConnectorMicrosoftDefenderAdvancedThreatProtection) Attributes() sentinelDataConnectorMicrosoftDefenderAdvancedThreatProtectionAttributes {
	return sentinelDataConnectorMicrosoftDefenderAdvancedThreatProtectionAttributes{ref: terra.ReferenceResource(sdcmdatp)}
}

// ImportState imports the given attribute values into [SentinelDataConnectorMicrosoftDefenderAdvancedThreatProtection]'s state.
func (sdcmdatp *SentinelDataConnectorMicrosoftDefenderAdvancedThreatProtection) ImportState(av io.Reader) error {
	sdcmdatp.state = &sentinelDataConnectorMicrosoftDefenderAdvancedThreatProtectionState{}
	if err := json.NewDecoder(av).Decode(sdcmdatp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sdcmdatp.Type(), sdcmdatp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SentinelDataConnectorMicrosoftDefenderAdvancedThreatProtection] has state.
func (sdcmdatp *SentinelDataConnectorMicrosoftDefenderAdvancedThreatProtection) State() (*sentinelDataConnectorMicrosoftDefenderAdvancedThreatProtectionState, bool) {
	return sdcmdatp.state, sdcmdatp.state != nil
}

// StateMust returns the state for [SentinelDataConnectorMicrosoftDefenderAdvancedThreatProtection]. Panics if the state is nil.
func (sdcmdatp *SentinelDataConnectorMicrosoftDefenderAdvancedThreatProtection) StateMust() *sentinelDataConnectorMicrosoftDefenderAdvancedThreatProtectionState {
	if sdcmdatp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sdcmdatp.Type(), sdcmdatp.LocalName()))
	}
	return sdcmdatp.state
}

// SentinelDataConnectorMicrosoftDefenderAdvancedThreatProtectionArgs contains the configurations for azurerm_sentinel_data_connector_microsoft_defender_advanced_threat_protection.
type SentinelDataConnectorMicrosoftDefenderAdvancedThreatProtectionArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// LogAnalyticsWorkspaceId: string, required
	LogAnalyticsWorkspaceId terra.StringValue `hcl:"log_analytics_workspace_id,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// TenantId: string, optional
	TenantId terra.StringValue `hcl:"tenant_id,attr"`
	// Timeouts: optional
	Timeouts *sentineldataconnectormicrosoftdefenderadvancedthreatprotection.Timeouts `hcl:"timeouts,block"`
}
type sentinelDataConnectorMicrosoftDefenderAdvancedThreatProtectionAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_sentinel_data_connector_microsoft_defender_advanced_threat_protection.
func (sdcmdatp sentinelDataConnectorMicrosoftDefenderAdvancedThreatProtectionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sdcmdatp.ref.Append("id"))
}

// LogAnalyticsWorkspaceId returns a reference to field log_analytics_workspace_id of azurerm_sentinel_data_connector_microsoft_defender_advanced_threat_protection.
func (sdcmdatp sentinelDataConnectorMicrosoftDefenderAdvancedThreatProtectionAttributes) LogAnalyticsWorkspaceId() terra.StringValue {
	return terra.ReferenceAsString(sdcmdatp.ref.Append("log_analytics_workspace_id"))
}

// Name returns a reference to field name of azurerm_sentinel_data_connector_microsoft_defender_advanced_threat_protection.
func (sdcmdatp sentinelDataConnectorMicrosoftDefenderAdvancedThreatProtectionAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(sdcmdatp.ref.Append("name"))
}

// TenantId returns a reference to field tenant_id of azurerm_sentinel_data_connector_microsoft_defender_advanced_threat_protection.
func (sdcmdatp sentinelDataConnectorMicrosoftDefenderAdvancedThreatProtectionAttributes) TenantId() terra.StringValue {
	return terra.ReferenceAsString(sdcmdatp.ref.Append("tenant_id"))
}

func (sdcmdatp sentinelDataConnectorMicrosoftDefenderAdvancedThreatProtectionAttributes) Timeouts() sentineldataconnectormicrosoftdefenderadvancedthreatprotection.TimeoutsAttributes {
	return terra.ReferenceAsSingle[sentineldataconnectormicrosoftdefenderadvancedthreatprotection.TimeoutsAttributes](sdcmdatp.ref.Append("timeouts"))
}

type sentinelDataConnectorMicrosoftDefenderAdvancedThreatProtectionState struct {
	Id                      string                                                                        `json:"id"`
	LogAnalyticsWorkspaceId string                                                                        `json:"log_analytics_workspace_id"`
	Name                    string                                                                        `json:"name"`
	TenantId                string                                                                        `json:"tenant_id"`
	Timeouts                *sentineldataconnectormicrosoftdefenderadvancedthreatprotection.TimeoutsState `json:"timeouts"`
}
