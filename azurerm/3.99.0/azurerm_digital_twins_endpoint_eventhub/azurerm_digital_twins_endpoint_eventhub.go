// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_digital_twins_endpoint_eventhub

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// New creates a new instance of [Resource].
func New(name string, args Args) *Resource {
	return &Resource{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Resource)(nil)

// Resource represents the Terraform resource azurerm_digital_twins_endpoint_eventhub.
type Resource struct {
	Name      string
	Args      Args
	state     *azurermDigitalTwinsEndpointEventhubState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (adtee *Resource) Type() string {
	return "azurerm_digital_twins_endpoint_eventhub"
}

// LocalName returns the local name for [Resource].
func (adtee *Resource) LocalName() string {
	return adtee.Name
}

// Configuration returns the configuration (args) for [Resource].
func (adtee *Resource) Configuration() interface{} {
	return adtee.Args
}

// DependOn is used for other resources to depend on [Resource].
func (adtee *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(adtee)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (adtee *Resource) Dependencies() terra.Dependencies {
	return adtee.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (adtee *Resource) LifecycleManagement() *terra.Lifecycle {
	return adtee.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (adtee *Resource) Attributes() azurermDigitalTwinsEndpointEventhubAttributes {
	return azurermDigitalTwinsEndpointEventhubAttributes{ref: terra.ReferenceResource(adtee)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (adtee *Resource) ImportState(state io.Reader) error {
	adtee.state = &azurermDigitalTwinsEndpointEventhubState{}
	if err := json.NewDecoder(state).Decode(adtee.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", adtee.Type(), adtee.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (adtee *Resource) State() (*azurermDigitalTwinsEndpointEventhubState, bool) {
	return adtee.state, adtee.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (adtee *Resource) StateMust() *azurermDigitalTwinsEndpointEventhubState {
	if adtee.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", adtee.Type(), adtee.LocalName()))
	}
	return adtee.state
}

// Args contains the configurations for azurerm_digital_twins_endpoint_eventhub.
type Args struct {
	// DeadLetterStorageSecret: string, optional
	DeadLetterStorageSecret terra.StringValue `hcl:"dead_letter_storage_secret,attr"`
	// DigitalTwinsId: string, required
	DigitalTwinsId terra.StringValue `hcl:"digital_twins_id,attr" validate:"required"`
	// EventhubPrimaryConnectionString: string, required
	EventhubPrimaryConnectionString terra.StringValue `hcl:"eventhub_primary_connection_string,attr" validate:"required"`
	// EventhubSecondaryConnectionString: string, required
	EventhubSecondaryConnectionString terra.StringValue `hcl:"eventhub_secondary_connection_string,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type azurermDigitalTwinsEndpointEventhubAttributes struct {
	ref terra.Reference
}

// DeadLetterStorageSecret returns a reference to field dead_letter_storage_secret of azurerm_digital_twins_endpoint_eventhub.
func (adtee azurermDigitalTwinsEndpointEventhubAttributes) DeadLetterStorageSecret() terra.StringValue {
	return terra.ReferenceAsString(adtee.ref.Append("dead_letter_storage_secret"))
}

// DigitalTwinsId returns a reference to field digital_twins_id of azurerm_digital_twins_endpoint_eventhub.
func (adtee azurermDigitalTwinsEndpointEventhubAttributes) DigitalTwinsId() terra.StringValue {
	return terra.ReferenceAsString(adtee.ref.Append("digital_twins_id"))
}

// EventhubPrimaryConnectionString returns a reference to field eventhub_primary_connection_string of azurerm_digital_twins_endpoint_eventhub.
func (adtee azurermDigitalTwinsEndpointEventhubAttributes) EventhubPrimaryConnectionString() terra.StringValue {
	return terra.ReferenceAsString(adtee.ref.Append("eventhub_primary_connection_string"))
}

// EventhubSecondaryConnectionString returns a reference to field eventhub_secondary_connection_string of azurerm_digital_twins_endpoint_eventhub.
func (adtee azurermDigitalTwinsEndpointEventhubAttributes) EventhubSecondaryConnectionString() terra.StringValue {
	return terra.ReferenceAsString(adtee.ref.Append("eventhub_secondary_connection_string"))
}

// Id returns a reference to field id of azurerm_digital_twins_endpoint_eventhub.
func (adtee azurermDigitalTwinsEndpointEventhubAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(adtee.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_digital_twins_endpoint_eventhub.
func (adtee azurermDigitalTwinsEndpointEventhubAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(adtee.ref.Append("name"))
}

func (adtee azurermDigitalTwinsEndpointEventhubAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](adtee.ref.Append("timeouts"))
}

type azurermDigitalTwinsEndpointEventhubState struct {
	DeadLetterStorageSecret           string         `json:"dead_letter_storage_secret"`
	DigitalTwinsId                    string         `json:"digital_twins_id"`
	EventhubPrimaryConnectionString   string         `json:"eventhub_primary_connection_string"`
	EventhubSecondaryConnectionString string         `json:"eventhub_secondary_connection_string"`
	Id                                string         `json:"id"`
	Name                              string         `json:"name"`
	Timeouts                          *TimeoutsState `json:"timeouts"`
}
