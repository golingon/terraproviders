// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	eventhubnamespacedisasterrecoveryconfig "github.com/golingon/terraproviders/azurerm/3.55.0/eventhubnamespacedisasterrecoveryconfig"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewEventhubNamespaceDisasterRecoveryConfig creates a new instance of [EventhubNamespaceDisasterRecoveryConfig].
func NewEventhubNamespaceDisasterRecoveryConfig(name string, args EventhubNamespaceDisasterRecoveryConfigArgs) *EventhubNamespaceDisasterRecoveryConfig {
	return &EventhubNamespaceDisasterRecoveryConfig{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*EventhubNamespaceDisasterRecoveryConfig)(nil)

// EventhubNamespaceDisasterRecoveryConfig represents the Terraform resource azurerm_eventhub_namespace_disaster_recovery_config.
type EventhubNamespaceDisasterRecoveryConfig struct {
	Name      string
	Args      EventhubNamespaceDisasterRecoveryConfigArgs
	state     *eventhubNamespaceDisasterRecoveryConfigState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [EventhubNamespaceDisasterRecoveryConfig].
func (endrc *EventhubNamespaceDisasterRecoveryConfig) Type() string {
	return "azurerm_eventhub_namespace_disaster_recovery_config"
}

// LocalName returns the local name for [EventhubNamespaceDisasterRecoveryConfig].
func (endrc *EventhubNamespaceDisasterRecoveryConfig) LocalName() string {
	return endrc.Name
}

// Configuration returns the configuration (args) for [EventhubNamespaceDisasterRecoveryConfig].
func (endrc *EventhubNamespaceDisasterRecoveryConfig) Configuration() interface{} {
	return endrc.Args
}

// DependOn is used for other resources to depend on [EventhubNamespaceDisasterRecoveryConfig].
func (endrc *EventhubNamespaceDisasterRecoveryConfig) DependOn() terra.Reference {
	return terra.ReferenceResource(endrc)
}

// Dependencies returns the list of resources [EventhubNamespaceDisasterRecoveryConfig] depends_on.
func (endrc *EventhubNamespaceDisasterRecoveryConfig) Dependencies() terra.Dependencies {
	return endrc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [EventhubNamespaceDisasterRecoveryConfig].
func (endrc *EventhubNamespaceDisasterRecoveryConfig) LifecycleManagement() *terra.Lifecycle {
	return endrc.Lifecycle
}

// Attributes returns the attributes for [EventhubNamespaceDisasterRecoveryConfig].
func (endrc *EventhubNamespaceDisasterRecoveryConfig) Attributes() eventhubNamespaceDisasterRecoveryConfigAttributes {
	return eventhubNamespaceDisasterRecoveryConfigAttributes{ref: terra.ReferenceResource(endrc)}
}

// ImportState imports the given attribute values into [EventhubNamespaceDisasterRecoveryConfig]'s state.
func (endrc *EventhubNamespaceDisasterRecoveryConfig) ImportState(av io.Reader) error {
	endrc.state = &eventhubNamespaceDisasterRecoveryConfigState{}
	if err := json.NewDecoder(av).Decode(endrc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", endrc.Type(), endrc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [EventhubNamespaceDisasterRecoveryConfig] has state.
func (endrc *EventhubNamespaceDisasterRecoveryConfig) State() (*eventhubNamespaceDisasterRecoveryConfigState, bool) {
	return endrc.state, endrc.state != nil
}

// StateMust returns the state for [EventhubNamespaceDisasterRecoveryConfig]. Panics if the state is nil.
func (endrc *EventhubNamespaceDisasterRecoveryConfig) StateMust() *eventhubNamespaceDisasterRecoveryConfigState {
	if endrc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", endrc.Type(), endrc.LocalName()))
	}
	return endrc.state
}

// EventhubNamespaceDisasterRecoveryConfigArgs contains the configurations for azurerm_eventhub_namespace_disaster_recovery_config.
type EventhubNamespaceDisasterRecoveryConfigArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// NamespaceName: string, required
	NamespaceName terra.StringValue `hcl:"namespace_name,attr" validate:"required"`
	// PartnerNamespaceId: string, required
	PartnerNamespaceId terra.StringValue `hcl:"partner_namespace_id,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *eventhubnamespacedisasterrecoveryconfig.Timeouts `hcl:"timeouts,block"`
}
type eventhubNamespaceDisasterRecoveryConfigAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_eventhub_namespace_disaster_recovery_config.
func (endrc eventhubNamespaceDisasterRecoveryConfigAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(endrc.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_eventhub_namespace_disaster_recovery_config.
func (endrc eventhubNamespaceDisasterRecoveryConfigAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(endrc.ref.Append("name"))
}

// NamespaceName returns a reference to field namespace_name of azurerm_eventhub_namespace_disaster_recovery_config.
func (endrc eventhubNamespaceDisasterRecoveryConfigAttributes) NamespaceName() terra.StringValue {
	return terra.ReferenceAsString(endrc.ref.Append("namespace_name"))
}

// PartnerNamespaceId returns a reference to field partner_namespace_id of azurerm_eventhub_namespace_disaster_recovery_config.
func (endrc eventhubNamespaceDisasterRecoveryConfigAttributes) PartnerNamespaceId() terra.StringValue {
	return terra.ReferenceAsString(endrc.ref.Append("partner_namespace_id"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_eventhub_namespace_disaster_recovery_config.
func (endrc eventhubNamespaceDisasterRecoveryConfigAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(endrc.ref.Append("resource_group_name"))
}

func (endrc eventhubNamespaceDisasterRecoveryConfigAttributes) Timeouts() eventhubnamespacedisasterrecoveryconfig.TimeoutsAttributes {
	return terra.ReferenceAsSingle[eventhubnamespacedisasterrecoveryconfig.TimeoutsAttributes](endrc.ref.Append("timeouts"))
}

type eventhubNamespaceDisasterRecoveryConfigState struct {
	Id                 string                                                 `json:"id"`
	Name               string                                                 `json:"name"`
	NamespaceName      string                                                 `json:"namespace_name"`
	PartnerNamespaceId string                                                 `json:"partner_namespace_id"`
	ResourceGroupName  string                                                 `json:"resource_group_name"`
	Timeouts           *eventhubnamespacedisasterrecoveryconfig.TimeoutsState `json:"timeouts"`
}
