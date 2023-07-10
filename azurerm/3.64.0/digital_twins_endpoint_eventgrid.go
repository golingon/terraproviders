// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	digitaltwinsendpointeventgrid "github.com/golingon/terraproviders/azurerm/3.64.0/digitaltwinsendpointeventgrid"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDigitalTwinsEndpointEventgrid creates a new instance of [DigitalTwinsEndpointEventgrid].
func NewDigitalTwinsEndpointEventgrid(name string, args DigitalTwinsEndpointEventgridArgs) *DigitalTwinsEndpointEventgrid {
	return &DigitalTwinsEndpointEventgrid{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DigitalTwinsEndpointEventgrid)(nil)

// DigitalTwinsEndpointEventgrid represents the Terraform resource azurerm_digital_twins_endpoint_eventgrid.
type DigitalTwinsEndpointEventgrid struct {
	Name      string
	Args      DigitalTwinsEndpointEventgridArgs
	state     *digitalTwinsEndpointEventgridState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DigitalTwinsEndpointEventgrid].
func (dtee *DigitalTwinsEndpointEventgrid) Type() string {
	return "azurerm_digital_twins_endpoint_eventgrid"
}

// LocalName returns the local name for [DigitalTwinsEndpointEventgrid].
func (dtee *DigitalTwinsEndpointEventgrid) LocalName() string {
	return dtee.Name
}

// Configuration returns the configuration (args) for [DigitalTwinsEndpointEventgrid].
func (dtee *DigitalTwinsEndpointEventgrid) Configuration() interface{} {
	return dtee.Args
}

// DependOn is used for other resources to depend on [DigitalTwinsEndpointEventgrid].
func (dtee *DigitalTwinsEndpointEventgrid) DependOn() terra.Reference {
	return terra.ReferenceResource(dtee)
}

// Dependencies returns the list of resources [DigitalTwinsEndpointEventgrid] depends_on.
func (dtee *DigitalTwinsEndpointEventgrid) Dependencies() terra.Dependencies {
	return dtee.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DigitalTwinsEndpointEventgrid].
func (dtee *DigitalTwinsEndpointEventgrid) LifecycleManagement() *terra.Lifecycle {
	return dtee.Lifecycle
}

// Attributes returns the attributes for [DigitalTwinsEndpointEventgrid].
func (dtee *DigitalTwinsEndpointEventgrid) Attributes() digitalTwinsEndpointEventgridAttributes {
	return digitalTwinsEndpointEventgridAttributes{ref: terra.ReferenceResource(dtee)}
}

// ImportState imports the given attribute values into [DigitalTwinsEndpointEventgrid]'s state.
func (dtee *DigitalTwinsEndpointEventgrid) ImportState(av io.Reader) error {
	dtee.state = &digitalTwinsEndpointEventgridState{}
	if err := json.NewDecoder(av).Decode(dtee.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dtee.Type(), dtee.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DigitalTwinsEndpointEventgrid] has state.
func (dtee *DigitalTwinsEndpointEventgrid) State() (*digitalTwinsEndpointEventgridState, bool) {
	return dtee.state, dtee.state != nil
}

// StateMust returns the state for [DigitalTwinsEndpointEventgrid]. Panics if the state is nil.
func (dtee *DigitalTwinsEndpointEventgrid) StateMust() *digitalTwinsEndpointEventgridState {
	if dtee.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dtee.Type(), dtee.LocalName()))
	}
	return dtee.state
}

// DigitalTwinsEndpointEventgridArgs contains the configurations for azurerm_digital_twins_endpoint_eventgrid.
type DigitalTwinsEndpointEventgridArgs struct {
	// DeadLetterStorageSecret: string, optional
	DeadLetterStorageSecret terra.StringValue `hcl:"dead_letter_storage_secret,attr"`
	// DigitalTwinsId: string, required
	DigitalTwinsId terra.StringValue `hcl:"digital_twins_id,attr" validate:"required"`
	// EventgridTopicEndpoint: string, required
	EventgridTopicEndpoint terra.StringValue `hcl:"eventgrid_topic_endpoint,attr" validate:"required"`
	// EventgridTopicPrimaryAccessKey: string, required
	EventgridTopicPrimaryAccessKey terra.StringValue `hcl:"eventgrid_topic_primary_access_key,attr" validate:"required"`
	// EventgridTopicSecondaryAccessKey: string, required
	EventgridTopicSecondaryAccessKey terra.StringValue `hcl:"eventgrid_topic_secondary_access_key,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *digitaltwinsendpointeventgrid.Timeouts `hcl:"timeouts,block"`
}
type digitalTwinsEndpointEventgridAttributes struct {
	ref terra.Reference
}

// DeadLetterStorageSecret returns a reference to field dead_letter_storage_secret of azurerm_digital_twins_endpoint_eventgrid.
func (dtee digitalTwinsEndpointEventgridAttributes) DeadLetterStorageSecret() terra.StringValue {
	return terra.ReferenceAsString(dtee.ref.Append("dead_letter_storage_secret"))
}

// DigitalTwinsId returns a reference to field digital_twins_id of azurerm_digital_twins_endpoint_eventgrid.
func (dtee digitalTwinsEndpointEventgridAttributes) DigitalTwinsId() terra.StringValue {
	return terra.ReferenceAsString(dtee.ref.Append("digital_twins_id"))
}

// EventgridTopicEndpoint returns a reference to field eventgrid_topic_endpoint of azurerm_digital_twins_endpoint_eventgrid.
func (dtee digitalTwinsEndpointEventgridAttributes) EventgridTopicEndpoint() terra.StringValue {
	return terra.ReferenceAsString(dtee.ref.Append("eventgrid_topic_endpoint"))
}

// EventgridTopicPrimaryAccessKey returns a reference to field eventgrid_topic_primary_access_key of azurerm_digital_twins_endpoint_eventgrid.
func (dtee digitalTwinsEndpointEventgridAttributes) EventgridTopicPrimaryAccessKey() terra.StringValue {
	return terra.ReferenceAsString(dtee.ref.Append("eventgrid_topic_primary_access_key"))
}

// EventgridTopicSecondaryAccessKey returns a reference to field eventgrid_topic_secondary_access_key of azurerm_digital_twins_endpoint_eventgrid.
func (dtee digitalTwinsEndpointEventgridAttributes) EventgridTopicSecondaryAccessKey() terra.StringValue {
	return terra.ReferenceAsString(dtee.ref.Append("eventgrid_topic_secondary_access_key"))
}

// Id returns a reference to field id of azurerm_digital_twins_endpoint_eventgrid.
func (dtee digitalTwinsEndpointEventgridAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dtee.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_digital_twins_endpoint_eventgrid.
func (dtee digitalTwinsEndpointEventgridAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(dtee.ref.Append("name"))
}

func (dtee digitalTwinsEndpointEventgridAttributes) Timeouts() digitaltwinsendpointeventgrid.TimeoutsAttributes {
	return terra.ReferenceAsSingle[digitaltwinsendpointeventgrid.TimeoutsAttributes](dtee.ref.Append("timeouts"))
}

type digitalTwinsEndpointEventgridState struct {
	DeadLetterStorageSecret          string                                       `json:"dead_letter_storage_secret"`
	DigitalTwinsId                   string                                       `json:"digital_twins_id"`
	EventgridTopicEndpoint           string                                       `json:"eventgrid_topic_endpoint"`
	EventgridTopicPrimaryAccessKey   string                                       `json:"eventgrid_topic_primary_access_key"`
	EventgridTopicSecondaryAccessKey string                                       `json:"eventgrid_topic_secondary_access_key"`
	Id                               string                                       `json:"id"`
	Name                             string                                       `json:"name"`
	Timeouts                         *digitaltwinsendpointeventgrid.TimeoutsState `json:"timeouts"`
}
