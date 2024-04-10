// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	sentineldataconnectorthreatintelligence "github.com/golingon/terraproviders/azurerm/3.98.0/sentineldataconnectorthreatintelligence"
	"io"
)

// NewSentinelDataConnectorThreatIntelligence creates a new instance of [SentinelDataConnectorThreatIntelligence].
func NewSentinelDataConnectorThreatIntelligence(name string, args SentinelDataConnectorThreatIntelligenceArgs) *SentinelDataConnectorThreatIntelligence {
	return &SentinelDataConnectorThreatIntelligence{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SentinelDataConnectorThreatIntelligence)(nil)

// SentinelDataConnectorThreatIntelligence represents the Terraform resource azurerm_sentinel_data_connector_threat_intelligence.
type SentinelDataConnectorThreatIntelligence struct {
	Name      string
	Args      SentinelDataConnectorThreatIntelligenceArgs
	state     *sentinelDataConnectorThreatIntelligenceState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SentinelDataConnectorThreatIntelligence].
func (sdcti *SentinelDataConnectorThreatIntelligence) Type() string {
	return "azurerm_sentinel_data_connector_threat_intelligence"
}

// LocalName returns the local name for [SentinelDataConnectorThreatIntelligence].
func (sdcti *SentinelDataConnectorThreatIntelligence) LocalName() string {
	return sdcti.Name
}

// Configuration returns the configuration (args) for [SentinelDataConnectorThreatIntelligence].
func (sdcti *SentinelDataConnectorThreatIntelligence) Configuration() interface{} {
	return sdcti.Args
}

// DependOn is used for other resources to depend on [SentinelDataConnectorThreatIntelligence].
func (sdcti *SentinelDataConnectorThreatIntelligence) DependOn() terra.Reference {
	return terra.ReferenceResource(sdcti)
}

// Dependencies returns the list of resources [SentinelDataConnectorThreatIntelligence] depends_on.
func (sdcti *SentinelDataConnectorThreatIntelligence) Dependencies() terra.Dependencies {
	return sdcti.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SentinelDataConnectorThreatIntelligence].
func (sdcti *SentinelDataConnectorThreatIntelligence) LifecycleManagement() *terra.Lifecycle {
	return sdcti.Lifecycle
}

// Attributes returns the attributes for [SentinelDataConnectorThreatIntelligence].
func (sdcti *SentinelDataConnectorThreatIntelligence) Attributes() sentinelDataConnectorThreatIntelligenceAttributes {
	return sentinelDataConnectorThreatIntelligenceAttributes{ref: terra.ReferenceResource(sdcti)}
}

// ImportState imports the given attribute values into [SentinelDataConnectorThreatIntelligence]'s state.
func (sdcti *SentinelDataConnectorThreatIntelligence) ImportState(av io.Reader) error {
	sdcti.state = &sentinelDataConnectorThreatIntelligenceState{}
	if err := json.NewDecoder(av).Decode(sdcti.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sdcti.Type(), sdcti.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SentinelDataConnectorThreatIntelligence] has state.
func (sdcti *SentinelDataConnectorThreatIntelligence) State() (*sentinelDataConnectorThreatIntelligenceState, bool) {
	return sdcti.state, sdcti.state != nil
}

// StateMust returns the state for [SentinelDataConnectorThreatIntelligence]. Panics if the state is nil.
func (sdcti *SentinelDataConnectorThreatIntelligence) StateMust() *sentinelDataConnectorThreatIntelligenceState {
	if sdcti.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sdcti.Type(), sdcti.LocalName()))
	}
	return sdcti.state
}

// SentinelDataConnectorThreatIntelligenceArgs contains the configurations for azurerm_sentinel_data_connector_threat_intelligence.
type SentinelDataConnectorThreatIntelligenceArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// LogAnalyticsWorkspaceId: string, required
	LogAnalyticsWorkspaceId terra.StringValue `hcl:"log_analytics_workspace_id,attr" validate:"required"`
	// LookbackDate: string, optional
	LookbackDate terra.StringValue `hcl:"lookback_date,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// TenantId: string, optional
	TenantId terra.StringValue `hcl:"tenant_id,attr"`
	// Timeouts: optional
	Timeouts *sentineldataconnectorthreatintelligence.Timeouts `hcl:"timeouts,block"`
}
type sentinelDataConnectorThreatIntelligenceAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_sentinel_data_connector_threat_intelligence.
func (sdcti sentinelDataConnectorThreatIntelligenceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sdcti.ref.Append("id"))
}

// LogAnalyticsWorkspaceId returns a reference to field log_analytics_workspace_id of azurerm_sentinel_data_connector_threat_intelligence.
func (sdcti sentinelDataConnectorThreatIntelligenceAttributes) LogAnalyticsWorkspaceId() terra.StringValue {
	return terra.ReferenceAsString(sdcti.ref.Append("log_analytics_workspace_id"))
}

// LookbackDate returns a reference to field lookback_date of azurerm_sentinel_data_connector_threat_intelligence.
func (sdcti sentinelDataConnectorThreatIntelligenceAttributes) LookbackDate() terra.StringValue {
	return terra.ReferenceAsString(sdcti.ref.Append("lookback_date"))
}

// Name returns a reference to field name of azurerm_sentinel_data_connector_threat_intelligence.
func (sdcti sentinelDataConnectorThreatIntelligenceAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(sdcti.ref.Append("name"))
}

// TenantId returns a reference to field tenant_id of azurerm_sentinel_data_connector_threat_intelligence.
func (sdcti sentinelDataConnectorThreatIntelligenceAttributes) TenantId() terra.StringValue {
	return terra.ReferenceAsString(sdcti.ref.Append("tenant_id"))
}

func (sdcti sentinelDataConnectorThreatIntelligenceAttributes) Timeouts() sentineldataconnectorthreatintelligence.TimeoutsAttributes {
	return terra.ReferenceAsSingle[sentineldataconnectorthreatintelligence.TimeoutsAttributes](sdcti.ref.Append("timeouts"))
}

type sentinelDataConnectorThreatIntelligenceState struct {
	Id                      string                                                 `json:"id"`
	LogAnalyticsWorkspaceId string                                                 `json:"log_analytics_workspace_id"`
	LookbackDate            string                                                 `json:"lookback_date"`
	Name                    string                                                 `json:"name"`
	TenantId                string                                                 `json:"tenant_id"`
	Timeouts                *sentineldataconnectorthreatintelligence.TimeoutsState `json:"timeouts"`
}
