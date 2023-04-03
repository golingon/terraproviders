// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	sentineldataconnectorofficepowerbi "github.com/golingon/terraproviders/azurerm/3.49.0/sentineldataconnectorofficepowerbi"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSentinelDataConnectorOfficePowerBi creates a new instance of [SentinelDataConnectorOfficePowerBi].
func NewSentinelDataConnectorOfficePowerBi(name string, args SentinelDataConnectorOfficePowerBiArgs) *SentinelDataConnectorOfficePowerBi {
	return &SentinelDataConnectorOfficePowerBi{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SentinelDataConnectorOfficePowerBi)(nil)

// SentinelDataConnectorOfficePowerBi represents the Terraform resource azurerm_sentinel_data_connector_office_power_bi.
type SentinelDataConnectorOfficePowerBi struct {
	Name      string
	Args      SentinelDataConnectorOfficePowerBiArgs
	state     *sentinelDataConnectorOfficePowerBiState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SentinelDataConnectorOfficePowerBi].
func (sdcopb *SentinelDataConnectorOfficePowerBi) Type() string {
	return "azurerm_sentinel_data_connector_office_power_bi"
}

// LocalName returns the local name for [SentinelDataConnectorOfficePowerBi].
func (sdcopb *SentinelDataConnectorOfficePowerBi) LocalName() string {
	return sdcopb.Name
}

// Configuration returns the configuration (args) for [SentinelDataConnectorOfficePowerBi].
func (sdcopb *SentinelDataConnectorOfficePowerBi) Configuration() interface{} {
	return sdcopb.Args
}

// DependOn is used for other resources to depend on [SentinelDataConnectorOfficePowerBi].
func (sdcopb *SentinelDataConnectorOfficePowerBi) DependOn() terra.Reference {
	return terra.ReferenceResource(sdcopb)
}

// Dependencies returns the list of resources [SentinelDataConnectorOfficePowerBi] depends_on.
func (sdcopb *SentinelDataConnectorOfficePowerBi) Dependencies() terra.Dependencies {
	return sdcopb.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SentinelDataConnectorOfficePowerBi].
func (sdcopb *SentinelDataConnectorOfficePowerBi) LifecycleManagement() *terra.Lifecycle {
	return sdcopb.Lifecycle
}

// Attributes returns the attributes for [SentinelDataConnectorOfficePowerBi].
func (sdcopb *SentinelDataConnectorOfficePowerBi) Attributes() sentinelDataConnectorOfficePowerBiAttributes {
	return sentinelDataConnectorOfficePowerBiAttributes{ref: terra.ReferenceResource(sdcopb)}
}

// ImportState imports the given attribute values into [SentinelDataConnectorOfficePowerBi]'s state.
func (sdcopb *SentinelDataConnectorOfficePowerBi) ImportState(av io.Reader) error {
	sdcopb.state = &sentinelDataConnectorOfficePowerBiState{}
	if err := json.NewDecoder(av).Decode(sdcopb.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sdcopb.Type(), sdcopb.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SentinelDataConnectorOfficePowerBi] has state.
func (sdcopb *SentinelDataConnectorOfficePowerBi) State() (*sentinelDataConnectorOfficePowerBiState, bool) {
	return sdcopb.state, sdcopb.state != nil
}

// StateMust returns the state for [SentinelDataConnectorOfficePowerBi]. Panics if the state is nil.
func (sdcopb *SentinelDataConnectorOfficePowerBi) StateMust() *sentinelDataConnectorOfficePowerBiState {
	if sdcopb.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sdcopb.Type(), sdcopb.LocalName()))
	}
	return sdcopb.state
}

// SentinelDataConnectorOfficePowerBiArgs contains the configurations for azurerm_sentinel_data_connector_office_power_bi.
type SentinelDataConnectorOfficePowerBiArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// LogAnalyticsWorkspaceId: string, required
	LogAnalyticsWorkspaceId terra.StringValue `hcl:"log_analytics_workspace_id,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// TenantId: string, optional
	TenantId terra.StringValue `hcl:"tenant_id,attr"`
	// Timeouts: optional
	Timeouts *sentineldataconnectorofficepowerbi.Timeouts `hcl:"timeouts,block"`
}
type sentinelDataConnectorOfficePowerBiAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_sentinel_data_connector_office_power_bi.
func (sdcopb sentinelDataConnectorOfficePowerBiAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sdcopb.ref.Append("id"))
}

// LogAnalyticsWorkspaceId returns a reference to field log_analytics_workspace_id of azurerm_sentinel_data_connector_office_power_bi.
func (sdcopb sentinelDataConnectorOfficePowerBiAttributes) LogAnalyticsWorkspaceId() terra.StringValue {
	return terra.ReferenceAsString(sdcopb.ref.Append("log_analytics_workspace_id"))
}

// Name returns a reference to field name of azurerm_sentinel_data_connector_office_power_bi.
func (sdcopb sentinelDataConnectorOfficePowerBiAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(sdcopb.ref.Append("name"))
}

// TenantId returns a reference to field tenant_id of azurerm_sentinel_data_connector_office_power_bi.
func (sdcopb sentinelDataConnectorOfficePowerBiAttributes) TenantId() terra.StringValue {
	return terra.ReferenceAsString(sdcopb.ref.Append("tenant_id"))
}

func (sdcopb sentinelDataConnectorOfficePowerBiAttributes) Timeouts() sentineldataconnectorofficepowerbi.TimeoutsAttributes {
	return terra.ReferenceAsSingle[sentineldataconnectorofficepowerbi.TimeoutsAttributes](sdcopb.ref.Append("timeouts"))
}

type sentinelDataConnectorOfficePowerBiState struct {
	Id                      string                                            `json:"id"`
	LogAnalyticsWorkspaceId string                                            `json:"log_analytics_workspace_id"`
	Name                    string                                            `json:"name"`
	TenantId                string                                            `json:"tenant_id"`
	Timeouts                *sentineldataconnectorofficepowerbi.TimeoutsState `json:"timeouts"`
}