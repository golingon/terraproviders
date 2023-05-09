// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	sentineldataconnectorthreatintelligencetaxii "github.com/golingon/terraproviders/azurerm/3.55.0/sentineldataconnectorthreatintelligencetaxii"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSentinelDataConnectorThreatIntelligenceTaxii creates a new instance of [SentinelDataConnectorThreatIntelligenceTaxii].
func NewSentinelDataConnectorThreatIntelligenceTaxii(name string, args SentinelDataConnectorThreatIntelligenceTaxiiArgs) *SentinelDataConnectorThreatIntelligenceTaxii {
	return &SentinelDataConnectorThreatIntelligenceTaxii{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SentinelDataConnectorThreatIntelligenceTaxii)(nil)

// SentinelDataConnectorThreatIntelligenceTaxii represents the Terraform resource azurerm_sentinel_data_connector_threat_intelligence_taxii.
type SentinelDataConnectorThreatIntelligenceTaxii struct {
	Name      string
	Args      SentinelDataConnectorThreatIntelligenceTaxiiArgs
	state     *sentinelDataConnectorThreatIntelligenceTaxiiState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SentinelDataConnectorThreatIntelligenceTaxii].
func (sdctit *SentinelDataConnectorThreatIntelligenceTaxii) Type() string {
	return "azurerm_sentinel_data_connector_threat_intelligence_taxii"
}

// LocalName returns the local name for [SentinelDataConnectorThreatIntelligenceTaxii].
func (sdctit *SentinelDataConnectorThreatIntelligenceTaxii) LocalName() string {
	return sdctit.Name
}

// Configuration returns the configuration (args) for [SentinelDataConnectorThreatIntelligenceTaxii].
func (sdctit *SentinelDataConnectorThreatIntelligenceTaxii) Configuration() interface{} {
	return sdctit.Args
}

// DependOn is used for other resources to depend on [SentinelDataConnectorThreatIntelligenceTaxii].
func (sdctit *SentinelDataConnectorThreatIntelligenceTaxii) DependOn() terra.Reference {
	return terra.ReferenceResource(sdctit)
}

// Dependencies returns the list of resources [SentinelDataConnectorThreatIntelligenceTaxii] depends_on.
func (sdctit *SentinelDataConnectorThreatIntelligenceTaxii) Dependencies() terra.Dependencies {
	return sdctit.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SentinelDataConnectorThreatIntelligenceTaxii].
func (sdctit *SentinelDataConnectorThreatIntelligenceTaxii) LifecycleManagement() *terra.Lifecycle {
	return sdctit.Lifecycle
}

// Attributes returns the attributes for [SentinelDataConnectorThreatIntelligenceTaxii].
func (sdctit *SentinelDataConnectorThreatIntelligenceTaxii) Attributes() sentinelDataConnectorThreatIntelligenceTaxiiAttributes {
	return sentinelDataConnectorThreatIntelligenceTaxiiAttributes{ref: terra.ReferenceResource(sdctit)}
}

// ImportState imports the given attribute values into [SentinelDataConnectorThreatIntelligenceTaxii]'s state.
func (sdctit *SentinelDataConnectorThreatIntelligenceTaxii) ImportState(av io.Reader) error {
	sdctit.state = &sentinelDataConnectorThreatIntelligenceTaxiiState{}
	if err := json.NewDecoder(av).Decode(sdctit.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sdctit.Type(), sdctit.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SentinelDataConnectorThreatIntelligenceTaxii] has state.
func (sdctit *SentinelDataConnectorThreatIntelligenceTaxii) State() (*sentinelDataConnectorThreatIntelligenceTaxiiState, bool) {
	return sdctit.state, sdctit.state != nil
}

// StateMust returns the state for [SentinelDataConnectorThreatIntelligenceTaxii]. Panics if the state is nil.
func (sdctit *SentinelDataConnectorThreatIntelligenceTaxii) StateMust() *sentinelDataConnectorThreatIntelligenceTaxiiState {
	if sdctit.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sdctit.Type(), sdctit.LocalName()))
	}
	return sdctit.state
}

// SentinelDataConnectorThreatIntelligenceTaxiiArgs contains the configurations for azurerm_sentinel_data_connector_threat_intelligence_taxii.
type SentinelDataConnectorThreatIntelligenceTaxiiArgs struct {
	// ApiRootUrl: string, required
	ApiRootUrl terra.StringValue `hcl:"api_root_url,attr" validate:"required"`
	// CollectionId: string, required
	CollectionId terra.StringValue `hcl:"collection_id,attr" validate:"required"`
	// DisplayName: string, required
	DisplayName terra.StringValue `hcl:"display_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// LogAnalyticsWorkspaceId: string, required
	LogAnalyticsWorkspaceId terra.StringValue `hcl:"log_analytics_workspace_id,attr" validate:"required"`
	// LookbackDate: string, optional
	LookbackDate terra.StringValue `hcl:"lookback_date,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Password: string, optional
	Password terra.StringValue `hcl:"password,attr"`
	// PollingFrequency: string, optional
	PollingFrequency terra.StringValue `hcl:"polling_frequency,attr"`
	// TenantId: string, optional
	TenantId terra.StringValue `hcl:"tenant_id,attr"`
	// UserName: string, optional
	UserName terra.StringValue `hcl:"user_name,attr"`
	// Timeouts: optional
	Timeouts *sentineldataconnectorthreatintelligencetaxii.Timeouts `hcl:"timeouts,block"`
}
type sentinelDataConnectorThreatIntelligenceTaxiiAttributes struct {
	ref terra.Reference
}

// ApiRootUrl returns a reference to field api_root_url of azurerm_sentinel_data_connector_threat_intelligence_taxii.
func (sdctit sentinelDataConnectorThreatIntelligenceTaxiiAttributes) ApiRootUrl() terra.StringValue {
	return terra.ReferenceAsString(sdctit.ref.Append("api_root_url"))
}

// CollectionId returns a reference to field collection_id of azurerm_sentinel_data_connector_threat_intelligence_taxii.
func (sdctit sentinelDataConnectorThreatIntelligenceTaxiiAttributes) CollectionId() terra.StringValue {
	return terra.ReferenceAsString(sdctit.ref.Append("collection_id"))
}

// DisplayName returns a reference to field display_name of azurerm_sentinel_data_connector_threat_intelligence_taxii.
func (sdctit sentinelDataConnectorThreatIntelligenceTaxiiAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(sdctit.ref.Append("display_name"))
}

// Id returns a reference to field id of azurerm_sentinel_data_connector_threat_intelligence_taxii.
func (sdctit sentinelDataConnectorThreatIntelligenceTaxiiAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sdctit.ref.Append("id"))
}

// LogAnalyticsWorkspaceId returns a reference to field log_analytics_workspace_id of azurerm_sentinel_data_connector_threat_intelligence_taxii.
func (sdctit sentinelDataConnectorThreatIntelligenceTaxiiAttributes) LogAnalyticsWorkspaceId() terra.StringValue {
	return terra.ReferenceAsString(sdctit.ref.Append("log_analytics_workspace_id"))
}

// LookbackDate returns a reference to field lookback_date of azurerm_sentinel_data_connector_threat_intelligence_taxii.
func (sdctit sentinelDataConnectorThreatIntelligenceTaxiiAttributes) LookbackDate() terra.StringValue {
	return terra.ReferenceAsString(sdctit.ref.Append("lookback_date"))
}

// Name returns a reference to field name of azurerm_sentinel_data_connector_threat_intelligence_taxii.
func (sdctit sentinelDataConnectorThreatIntelligenceTaxiiAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(sdctit.ref.Append("name"))
}

// Password returns a reference to field password of azurerm_sentinel_data_connector_threat_intelligence_taxii.
func (sdctit sentinelDataConnectorThreatIntelligenceTaxiiAttributes) Password() terra.StringValue {
	return terra.ReferenceAsString(sdctit.ref.Append("password"))
}

// PollingFrequency returns a reference to field polling_frequency of azurerm_sentinel_data_connector_threat_intelligence_taxii.
func (sdctit sentinelDataConnectorThreatIntelligenceTaxiiAttributes) PollingFrequency() terra.StringValue {
	return terra.ReferenceAsString(sdctit.ref.Append("polling_frequency"))
}

// TenantId returns a reference to field tenant_id of azurerm_sentinel_data_connector_threat_intelligence_taxii.
func (sdctit sentinelDataConnectorThreatIntelligenceTaxiiAttributes) TenantId() terra.StringValue {
	return terra.ReferenceAsString(sdctit.ref.Append("tenant_id"))
}

// UserName returns a reference to field user_name of azurerm_sentinel_data_connector_threat_intelligence_taxii.
func (sdctit sentinelDataConnectorThreatIntelligenceTaxiiAttributes) UserName() terra.StringValue {
	return terra.ReferenceAsString(sdctit.ref.Append("user_name"))
}

func (sdctit sentinelDataConnectorThreatIntelligenceTaxiiAttributes) Timeouts() sentineldataconnectorthreatintelligencetaxii.TimeoutsAttributes {
	return terra.ReferenceAsSingle[sentineldataconnectorthreatintelligencetaxii.TimeoutsAttributes](sdctit.ref.Append("timeouts"))
}

type sentinelDataConnectorThreatIntelligenceTaxiiState struct {
	ApiRootUrl              string                                                      `json:"api_root_url"`
	CollectionId            string                                                      `json:"collection_id"`
	DisplayName             string                                                      `json:"display_name"`
	Id                      string                                                      `json:"id"`
	LogAnalyticsWorkspaceId string                                                      `json:"log_analytics_workspace_id"`
	LookbackDate            string                                                      `json:"lookback_date"`
	Name                    string                                                      `json:"name"`
	Password                string                                                      `json:"password"`
	PollingFrequency        string                                                      `json:"polling_frequency"`
	TenantId                string                                                      `json:"tenant_id"`
	UserName                string                                                      `json:"user_name"`
	Timeouts                *sentineldataconnectorthreatintelligencetaxii.TimeoutsState `json:"timeouts"`
}
