// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	sentineldataconnectoroffice365 "github.com/golingon/terraproviders/azurerm/3.58.0/sentineldataconnectoroffice365"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSentinelDataConnectorOffice365 creates a new instance of [SentinelDataConnectorOffice365].
func NewSentinelDataConnectorOffice365(name string, args SentinelDataConnectorOffice365Args) *SentinelDataConnectorOffice365 {
	return &SentinelDataConnectorOffice365{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SentinelDataConnectorOffice365)(nil)

// SentinelDataConnectorOffice365 represents the Terraform resource azurerm_sentinel_data_connector_office_365.
type SentinelDataConnectorOffice365 struct {
	Name      string
	Args      SentinelDataConnectorOffice365Args
	state     *sentinelDataConnectorOffice365State
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SentinelDataConnectorOffice365].
func (sdco3 *SentinelDataConnectorOffice365) Type() string {
	return "azurerm_sentinel_data_connector_office_365"
}

// LocalName returns the local name for [SentinelDataConnectorOffice365].
func (sdco3 *SentinelDataConnectorOffice365) LocalName() string {
	return sdco3.Name
}

// Configuration returns the configuration (args) for [SentinelDataConnectorOffice365].
func (sdco3 *SentinelDataConnectorOffice365) Configuration() interface{} {
	return sdco3.Args
}

// DependOn is used for other resources to depend on [SentinelDataConnectorOffice365].
func (sdco3 *SentinelDataConnectorOffice365) DependOn() terra.Reference {
	return terra.ReferenceResource(sdco3)
}

// Dependencies returns the list of resources [SentinelDataConnectorOffice365] depends_on.
func (sdco3 *SentinelDataConnectorOffice365) Dependencies() terra.Dependencies {
	return sdco3.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SentinelDataConnectorOffice365].
func (sdco3 *SentinelDataConnectorOffice365) LifecycleManagement() *terra.Lifecycle {
	return sdco3.Lifecycle
}

// Attributes returns the attributes for [SentinelDataConnectorOffice365].
func (sdco3 *SentinelDataConnectorOffice365) Attributes() sentinelDataConnectorOffice365Attributes {
	return sentinelDataConnectorOffice365Attributes{ref: terra.ReferenceResource(sdco3)}
}

// ImportState imports the given attribute values into [SentinelDataConnectorOffice365]'s state.
func (sdco3 *SentinelDataConnectorOffice365) ImportState(av io.Reader) error {
	sdco3.state = &sentinelDataConnectorOffice365State{}
	if err := json.NewDecoder(av).Decode(sdco3.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sdco3.Type(), sdco3.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SentinelDataConnectorOffice365] has state.
func (sdco3 *SentinelDataConnectorOffice365) State() (*sentinelDataConnectorOffice365State, bool) {
	return sdco3.state, sdco3.state != nil
}

// StateMust returns the state for [SentinelDataConnectorOffice365]. Panics if the state is nil.
func (sdco3 *SentinelDataConnectorOffice365) StateMust() *sentinelDataConnectorOffice365State {
	if sdco3.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sdco3.Type(), sdco3.LocalName()))
	}
	return sdco3.state
}

// SentinelDataConnectorOffice365Args contains the configurations for azurerm_sentinel_data_connector_office_365.
type SentinelDataConnectorOffice365Args struct {
	// ExchangeEnabled: bool, optional
	ExchangeEnabled terra.BoolValue `hcl:"exchange_enabled,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// LogAnalyticsWorkspaceId: string, required
	LogAnalyticsWorkspaceId terra.StringValue `hcl:"log_analytics_workspace_id,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// SharepointEnabled: bool, optional
	SharepointEnabled terra.BoolValue `hcl:"sharepoint_enabled,attr"`
	// TeamsEnabled: bool, optional
	TeamsEnabled terra.BoolValue `hcl:"teams_enabled,attr"`
	// TenantId: string, optional
	TenantId terra.StringValue `hcl:"tenant_id,attr"`
	// Timeouts: optional
	Timeouts *sentineldataconnectoroffice365.Timeouts `hcl:"timeouts,block"`
}
type sentinelDataConnectorOffice365Attributes struct {
	ref terra.Reference
}

// ExchangeEnabled returns a reference to field exchange_enabled of azurerm_sentinel_data_connector_office_365.
func (sdco3 sentinelDataConnectorOffice365Attributes) ExchangeEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(sdco3.ref.Append("exchange_enabled"))
}

// Id returns a reference to field id of azurerm_sentinel_data_connector_office_365.
func (sdco3 sentinelDataConnectorOffice365Attributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sdco3.ref.Append("id"))
}

// LogAnalyticsWorkspaceId returns a reference to field log_analytics_workspace_id of azurerm_sentinel_data_connector_office_365.
func (sdco3 sentinelDataConnectorOffice365Attributes) LogAnalyticsWorkspaceId() terra.StringValue {
	return terra.ReferenceAsString(sdco3.ref.Append("log_analytics_workspace_id"))
}

// Name returns a reference to field name of azurerm_sentinel_data_connector_office_365.
func (sdco3 sentinelDataConnectorOffice365Attributes) Name() terra.StringValue {
	return terra.ReferenceAsString(sdco3.ref.Append("name"))
}

// SharepointEnabled returns a reference to field sharepoint_enabled of azurerm_sentinel_data_connector_office_365.
func (sdco3 sentinelDataConnectorOffice365Attributes) SharepointEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(sdco3.ref.Append("sharepoint_enabled"))
}

// TeamsEnabled returns a reference to field teams_enabled of azurerm_sentinel_data_connector_office_365.
func (sdco3 sentinelDataConnectorOffice365Attributes) TeamsEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(sdco3.ref.Append("teams_enabled"))
}

// TenantId returns a reference to field tenant_id of azurerm_sentinel_data_connector_office_365.
func (sdco3 sentinelDataConnectorOffice365Attributes) TenantId() terra.StringValue {
	return terra.ReferenceAsString(sdco3.ref.Append("tenant_id"))
}

func (sdco3 sentinelDataConnectorOffice365Attributes) Timeouts() sentineldataconnectoroffice365.TimeoutsAttributes {
	return terra.ReferenceAsSingle[sentineldataconnectoroffice365.TimeoutsAttributes](sdco3.ref.Append("timeouts"))
}

type sentinelDataConnectorOffice365State struct {
	ExchangeEnabled         bool                                          `json:"exchange_enabled"`
	Id                      string                                        `json:"id"`
	LogAnalyticsWorkspaceId string                                        `json:"log_analytics_workspace_id"`
	Name                    string                                        `json:"name"`
	SharepointEnabled       bool                                          `json:"sharepoint_enabled"`
	TeamsEnabled            bool                                          `json:"teams_enabled"`
	TenantId                string                                        `json:"tenant_id"`
	Timeouts                *sentineldataconnectoroffice365.TimeoutsState `json:"timeouts"`
}
