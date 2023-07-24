// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	sentineldataconnectoriot "github.com/golingon/terraproviders/azurerm/3.66.0/sentineldataconnectoriot"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSentinelDataConnectorIot creates a new instance of [SentinelDataConnectorIot].
func NewSentinelDataConnectorIot(name string, args SentinelDataConnectorIotArgs) *SentinelDataConnectorIot {
	return &SentinelDataConnectorIot{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SentinelDataConnectorIot)(nil)

// SentinelDataConnectorIot represents the Terraform resource azurerm_sentinel_data_connector_iot.
type SentinelDataConnectorIot struct {
	Name      string
	Args      SentinelDataConnectorIotArgs
	state     *sentinelDataConnectorIotState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SentinelDataConnectorIot].
func (sdci *SentinelDataConnectorIot) Type() string {
	return "azurerm_sentinel_data_connector_iot"
}

// LocalName returns the local name for [SentinelDataConnectorIot].
func (sdci *SentinelDataConnectorIot) LocalName() string {
	return sdci.Name
}

// Configuration returns the configuration (args) for [SentinelDataConnectorIot].
func (sdci *SentinelDataConnectorIot) Configuration() interface{} {
	return sdci.Args
}

// DependOn is used for other resources to depend on [SentinelDataConnectorIot].
func (sdci *SentinelDataConnectorIot) DependOn() terra.Reference {
	return terra.ReferenceResource(sdci)
}

// Dependencies returns the list of resources [SentinelDataConnectorIot] depends_on.
func (sdci *SentinelDataConnectorIot) Dependencies() terra.Dependencies {
	return sdci.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SentinelDataConnectorIot].
func (sdci *SentinelDataConnectorIot) LifecycleManagement() *terra.Lifecycle {
	return sdci.Lifecycle
}

// Attributes returns the attributes for [SentinelDataConnectorIot].
func (sdci *SentinelDataConnectorIot) Attributes() sentinelDataConnectorIotAttributes {
	return sentinelDataConnectorIotAttributes{ref: terra.ReferenceResource(sdci)}
}

// ImportState imports the given attribute values into [SentinelDataConnectorIot]'s state.
func (sdci *SentinelDataConnectorIot) ImportState(av io.Reader) error {
	sdci.state = &sentinelDataConnectorIotState{}
	if err := json.NewDecoder(av).Decode(sdci.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sdci.Type(), sdci.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SentinelDataConnectorIot] has state.
func (sdci *SentinelDataConnectorIot) State() (*sentinelDataConnectorIotState, bool) {
	return sdci.state, sdci.state != nil
}

// StateMust returns the state for [SentinelDataConnectorIot]. Panics if the state is nil.
func (sdci *SentinelDataConnectorIot) StateMust() *sentinelDataConnectorIotState {
	if sdci.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sdci.Type(), sdci.LocalName()))
	}
	return sdci.state
}

// SentinelDataConnectorIotArgs contains the configurations for azurerm_sentinel_data_connector_iot.
type SentinelDataConnectorIotArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// LogAnalyticsWorkspaceId: string, required
	LogAnalyticsWorkspaceId terra.StringValue `hcl:"log_analytics_workspace_id,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// SubscriptionId: string, optional
	SubscriptionId terra.StringValue `hcl:"subscription_id,attr"`
	// Timeouts: optional
	Timeouts *sentineldataconnectoriot.Timeouts `hcl:"timeouts,block"`
}
type sentinelDataConnectorIotAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_sentinel_data_connector_iot.
func (sdci sentinelDataConnectorIotAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sdci.ref.Append("id"))
}

// LogAnalyticsWorkspaceId returns a reference to field log_analytics_workspace_id of azurerm_sentinel_data_connector_iot.
func (sdci sentinelDataConnectorIotAttributes) LogAnalyticsWorkspaceId() terra.StringValue {
	return terra.ReferenceAsString(sdci.ref.Append("log_analytics_workspace_id"))
}

// Name returns a reference to field name of azurerm_sentinel_data_connector_iot.
func (sdci sentinelDataConnectorIotAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(sdci.ref.Append("name"))
}

// SubscriptionId returns a reference to field subscription_id of azurerm_sentinel_data_connector_iot.
func (sdci sentinelDataConnectorIotAttributes) SubscriptionId() terra.StringValue {
	return terra.ReferenceAsString(sdci.ref.Append("subscription_id"))
}

func (sdci sentinelDataConnectorIotAttributes) Timeouts() sentineldataconnectoriot.TimeoutsAttributes {
	return terra.ReferenceAsSingle[sentineldataconnectoriot.TimeoutsAttributes](sdci.ref.Append("timeouts"))
}

type sentinelDataConnectorIotState struct {
	Id                      string                                  `json:"id"`
	LogAnalyticsWorkspaceId string                                  `json:"log_analytics_workspace_id"`
	Name                    string                                  `json:"name"`
	SubscriptionId          string                                  `json:"subscription_id"`
	Timeouts                *sentineldataconnectoriot.TimeoutsState `json:"timeouts"`
}
