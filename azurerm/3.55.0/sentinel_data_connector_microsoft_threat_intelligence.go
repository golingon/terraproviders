// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	sentineldataconnectormicrosoftthreatintelligence "github.com/golingon/terraproviders/azurerm/3.55.0/sentineldataconnectormicrosoftthreatintelligence"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSentinelDataConnectorMicrosoftThreatIntelligence creates a new instance of [SentinelDataConnectorMicrosoftThreatIntelligence].
func NewSentinelDataConnectorMicrosoftThreatIntelligence(name string, args SentinelDataConnectorMicrosoftThreatIntelligenceArgs) *SentinelDataConnectorMicrosoftThreatIntelligence {
	return &SentinelDataConnectorMicrosoftThreatIntelligence{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SentinelDataConnectorMicrosoftThreatIntelligence)(nil)

// SentinelDataConnectorMicrosoftThreatIntelligence represents the Terraform resource azurerm_sentinel_data_connector_microsoft_threat_intelligence.
type SentinelDataConnectorMicrosoftThreatIntelligence struct {
	Name      string
	Args      SentinelDataConnectorMicrosoftThreatIntelligenceArgs
	state     *sentinelDataConnectorMicrosoftThreatIntelligenceState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SentinelDataConnectorMicrosoftThreatIntelligence].
func (sdcmti *SentinelDataConnectorMicrosoftThreatIntelligence) Type() string {
	return "azurerm_sentinel_data_connector_microsoft_threat_intelligence"
}

// LocalName returns the local name for [SentinelDataConnectorMicrosoftThreatIntelligence].
func (sdcmti *SentinelDataConnectorMicrosoftThreatIntelligence) LocalName() string {
	return sdcmti.Name
}

// Configuration returns the configuration (args) for [SentinelDataConnectorMicrosoftThreatIntelligence].
func (sdcmti *SentinelDataConnectorMicrosoftThreatIntelligence) Configuration() interface{} {
	return sdcmti.Args
}

// DependOn is used for other resources to depend on [SentinelDataConnectorMicrosoftThreatIntelligence].
func (sdcmti *SentinelDataConnectorMicrosoftThreatIntelligence) DependOn() terra.Reference {
	return terra.ReferenceResource(sdcmti)
}

// Dependencies returns the list of resources [SentinelDataConnectorMicrosoftThreatIntelligence] depends_on.
func (sdcmti *SentinelDataConnectorMicrosoftThreatIntelligence) Dependencies() terra.Dependencies {
	return sdcmti.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SentinelDataConnectorMicrosoftThreatIntelligence].
func (sdcmti *SentinelDataConnectorMicrosoftThreatIntelligence) LifecycleManagement() *terra.Lifecycle {
	return sdcmti.Lifecycle
}

// Attributes returns the attributes for [SentinelDataConnectorMicrosoftThreatIntelligence].
func (sdcmti *SentinelDataConnectorMicrosoftThreatIntelligence) Attributes() sentinelDataConnectorMicrosoftThreatIntelligenceAttributes {
	return sentinelDataConnectorMicrosoftThreatIntelligenceAttributes{ref: terra.ReferenceResource(sdcmti)}
}

// ImportState imports the given attribute values into [SentinelDataConnectorMicrosoftThreatIntelligence]'s state.
func (sdcmti *SentinelDataConnectorMicrosoftThreatIntelligence) ImportState(av io.Reader) error {
	sdcmti.state = &sentinelDataConnectorMicrosoftThreatIntelligenceState{}
	if err := json.NewDecoder(av).Decode(sdcmti.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sdcmti.Type(), sdcmti.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SentinelDataConnectorMicrosoftThreatIntelligence] has state.
func (sdcmti *SentinelDataConnectorMicrosoftThreatIntelligence) State() (*sentinelDataConnectorMicrosoftThreatIntelligenceState, bool) {
	return sdcmti.state, sdcmti.state != nil
}

// StateMust returns the state for [SentinelDataConnectorMicrosoftThreatIntelligence]. Panics if the state is nil.
func (sdcmti *SentinelDataConnectorMicrosoftThreatIntelligence) StateMust() *sentinelDataConnectorMicrosoftThreatIntelligenceState {
	if sdcmti.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sdcmti.Type(), sdcmti.LocalName()))
	}
	return sdcmti.state
}

// SentinelDataConnectorMicrosoftThreatIntelligenceArgs contains the configurations for azurerm_sentinel_data_connector_microsoft_threat_intelligence.
type SentinelDataConnectorMicrosoftThreatIntelligenceArgs struct {
	// BingSafetyPhishingUrlLookbackDate: string, optional
	BingSafetyPhishingUrlLookbackDate terra.StringValue `hcl:"bing_safety_phishing_url_lookback_date,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// LogAnalyticsWorkspaceId: string, required
	LogAnalyticsWorkspaceId terra.StringValue `hcl:"log_analytics_workspace_id,attr" validate:"required"`
	// MicrosoftEmergingThreatFeedLookbackDate: string, optional
	MicrosoftEmergingThreatFeedLookbackDate terra.StringValue `hcl:"microsoft_emerging_threat_feed_lookback_date,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// TenantId: string, optional
	TenantId terra.StringValue `hcl:"tenant_id,attr"`
	// Timeouts: optional
	Timeouts *sentineldataconnectormicrosoftthreatintelligence.Timeouts `hcl:"timeouts,block"`
}
type sentinelDataConnectorMicrosoftThreatIntelligenceAttributes struct {
	ref terra.Reference
}

// BingSafetyPhishingUrlLookbackDate returns a reference to field bing_safety_phishing_url_lookback_date of azurerm_sentinel_data_connector_microsoft_threat_intelligence.
func (sdcmti sentinelDataConnectorMicrosoftThreatIntelligenceAttributes) BingSafetyPhishingUrlLookbackDate() terra.StringValue {
	return terra.ReferenceAsString(sdcmti.ref.Append("bing_safety_phishing_url_lookback_date"))
}

// Id returns a reference to field id of azurerm_sentinel_data_connector_microsoft_threat_intelligence.
func (sdcmti sentinelDataConnectorMicrosoftThreatIntelligenceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sdcmti.ref.Append("id"))
}

// LogAnalyticsWorkspaceId returns a reference to field log_analytics_workspace_id of azurerm_sentinel_data_connector_microsoft_threat_intelligence.
func (sdcmti sentinelDataConnectorMicrosoftThreatIntelligenceAttributes) LogAnalyticsWorkspaceId() terra.StringValue {
	return terra.ReferenceAsString(sdcmti.ref.Append("log_analytics_workspace_id"))
}

// MicrosoftEmergingThreatFeedLookbackDate returns a reference to field microsoft_emerging_threat_feed_lookback_date of azurerm_sentinel_data_connector_microsoft_threat_intelligence.
func (sdcmti sentinelDataConnectorMicrosoftThreatIntelligenceAttributes) MicrosoftEmergingThreatFeedLookbackDate() terra.StringValue {
	return terra.ReferenceAsString(sdcmti.ref.Append("microsoft_emerging_threat_feed_lookback_date"))
}

// Name returns a reference to field name of azurerm_sentinel_data_connector_microsoft_threat_intelligence.
func (sdcmti sentinelDataConnectorMicrosoftThreatIntelligenceAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(sdcmti.ref.Append("name"))
}

// TenantId returns a reference to field tenant_id of azurerm_sentinel_data_connector_microsoft_threat_intelligence.
func (sdcmti sentinelDataConnectorMicrosoftThreatIntelligenceAttributes) TenantId() terra.StringValue {
	return terra.ReferenceAsString(sdcmti.ref.Append("tenant_id"))
}

func (sdcmti sentinelDataConnectorMicrosoftThreatIntelligenceAttributes) Timeouts() sentineldataconnectormicrosoftthreatintelligence.TimeoutsAttributes {
	return terra.ReferenceAsSingle[sentineldataconnectormicrosoftthreatintelligence.TimeoutsAttributes](sdcmti.ref.Append("timeouts"))
}

type sentinelDataConnectorMicrosoftThreatIntelligenceState struct {
	BingSafetyPhishingUrlLookbackDate       string                                                          `json:"bing_safety_phishing_url_lookback_date"`
	Id                                      string                                                          `json:"id"`
	LogAnalyticsWorkspaceId                 string                                                          `json:"log_analytics_workspace_id"`
	MicrosoftEmergingThreatFeedLookbackDate string                                                          `json:"microsoft_emerging_threat_feed_lookback_date"`
	Name                                    string                                                          `json:"name"`
	TenantId                                string                                                          `json:"tenant_id"`
	Timeouts                                *sentineldataconnectormicrosoftthreatintelligence.TimeoutsState `json:"timeouts"`
}
