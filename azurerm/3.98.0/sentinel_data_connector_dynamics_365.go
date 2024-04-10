// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	sentineldataconnectordynamics365 "github.com/golingon/terraproviders/azurerm/3.98.0/sentineldataconnectordynamics365"
	"io"
)

// NewSentinelDataConnectorDynamics365 creates a new instance of [SentinelDataConnectorDynamics365].
func NewSentinelDataConnectorDynamics365(name string, args SentinelDataConnectorDynamics365Args) *SentinelDataConnectorDynamics365 {
	return &SentinelDataConnectorDynamics365{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SentinelDataConnectorDynamics365)(nil)

// SentinelDataConnectorDynamics365 represents the Terraform resource azurerm_sentinel_data_connector_dynamics_365.
type SentinelDataConnectorDynamics365 struct {
	Name      string
	Args      SentinelDataConnectorDynamics365Args
	state     *sentinelDataConnectorDynamics365State
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SentinelDataConnectorDynamics365].
func (sdcd3 *SentinelDataConnectorDynamics365) Type() string {
	return "azurerm_sentinel_data_connector_dynamics_365"
}

// LocalName returns the local name for [SentinelDataConnectorDynamics365].
func (sdcd3 *SentinelDataConnectorDynamics365) LocalName() string {
	return sdcd3.Name
}

// Configuration returns the configuration (args) for [SentinelDataConnectorDynamics365].
func (sdcd3 *SentinelDataConnectorDynamics365) Configuration() interface{} {
	return sdcd3.Args
}

// DependOn is used for other resources to depend on [SentinelDataConnectorDynamics365].
func (sdcd3 *SentinelDataConnectorDynamics365) DependOn() terra.Reference {
	return terra.ReferenceResource(sdcd3)
}

// Dependencies returns the list of resources [SentinelDataConnectorDynamics365] depends_on.
func (sdcd3 *SentinelDataConnectorDynamics365) Dependencies() terra.Dependencies {
	return sdcd3.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SentinelDataConnectorDynamics365].
func (sdcd3 *SentinelDataConnectorDynamics365) LifecycleManagement() *terra.Lifecycle {
	return sdcd3.Lifecycle
}

// Attributes returns the attributes for [SentinelDataConnectorDynamics365].
func (sdcd3 *SentinelDataConnectorDynamics365) Attributes() sentinelDataConnectorDynamics365Attributes {
	return sentinelDataConnectorDynamics365Attributes{ref: terra.ReferenceResource(sdcd3)}
}

// ImportState imports the given attribute values into [SentinelDataConnectorDynamics365]'s state.
func (sdcd3 *SentinelDataConnectorDynamics365) ImportState(av io.Reader) error {
	sdcd3.state = &sentinelDataConnectorDynamics365State{}
	if err := json.NewDecoder(av).Decode(sdcd3.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sdcd3.Type(), sdcd3.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SentinelDataConnectorDynamics365] has state.
func (sdcd3 *SentinelDataConnectorDynamics365) State() (*sentinelDataConnectorDynamics365State, bool) {
	return sdcd3.state, sdcd3.state != nil
}

// StateMust returns the state for [SentinelDataConnectorDynamics365]. Panics if the state is nil.
func (sdcd3 *SentinelDataConnectorDynamics365) StateMust() *sentinelDataConnectorDynamics365State {
	if sdcd3.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sdcd3.Type(), sdcd3.LocalName()))
	}
	return sdcd3.state
}

// SentinelDataConnectorDynamics365Args contains the configurations for azurerm_sentinel_data_connector_dynamics_365.
type SentinelDataConnectorDynamics365Args struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// LogAnalyticsWorkspaceId: string, required
	LogAnalyticsWorkspaceId terra.StringValue `hcl:"log_analytics_workspace_id,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// TenantId: string, optional
	TenantId terra.StringValue `hcl:"tenant_id,attr"`
	// Timeouts: optional
	Timeouts *sentineldataconnectordynamics365.Timeouts `hcl:"timeouts,block"`
}
type sentinelDataConnectorDynamics365Attributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_sentinel_data_connector_dynamics_365.
func (sdcd3 sentinelDataConnectorDynamics365Attributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sdcd3.ref.Append("id"))
}

// LogAnalyticsWorkspaceId returns a reference to field log_analytics_workspace_id of azurerm_sentinel_data_connector_dynamics_365.
func (sdcd3 sentinelDataConnectorDynamics365Attributes) LogAnalyticsWorkspaceId() terra.StringValue {
	return terra.ReferenceAsString(sdcd3.ref.Append("log_analytics_workspace_id"))
}

// Name returns a reference to field name of azurerm_sentinel_data_connector_dynamics_365.
func (sdcd3 sentinelDataConnectorDynamics365Attributes) Name() terra.StringValue {
	return terra.ReferenceAsString(sdcd3.ref.Append("name"))
}

// TenantId returns a reference to field tenant_id of azurerm_sentinel_data_connector_dynamics_365.
func (sdcd3 sentinelDataConnectorDynamics365Attributes) TenantId() terra.StringValue {
	return terra.ReferenceAsString(sdcd3.ref.Append("tenant_id"))
}

func (sdcd3 sentinelDataConnectorDynamics365Attributes) Timeouts() sentineldataconnectordynamics365.TimeoutsAttributes {
	return terra.ReferenceAsSingle[sentineldataconnectordynamics365.TimeoutsAttributes](sdcd3.ref.Append("timeouts"))
}

type sentinelDataConnectorDynamics365State struct {
	Id                      string                                          `json:"id"`
	LogAnalyticsWorkspaceId string                                          `json:"log_analytics_workspace_id"`
	Name                    string                                          `json:"name"`
	TenantId                string                                          `json:"tenant_id"`
	Timeouts                *sentineldataconnectordynamics365.TimeoutsState `json:"timeouts"`
}
