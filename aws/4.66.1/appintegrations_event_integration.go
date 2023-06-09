// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	appintegrationseventintegration "github.com/golingon/terraproviders/aws/4.66.1/appintegrationseventintegration"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewAppintegrationsEventIntegration creates a new instance of [AppintegrationsEventIntegration].
func NewAppintegrationsEventIntegration(name string, args AppintegrationsEventIntegrationArgs) *AppintegrationsEventIntegration {
	return &AppintegrationsEventIntegration{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*AppintegrationsEventIntegration)(nil)

// AppintegrationsEventIntegration represents the Terraform resource aws_appintegrations_event_integration.
type AppintegrationsEventIntegration struct {
	Name      string
	Args      AppintegrationsEventIntegrationArgs
	state     *appintegrationsEventIntegrationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [AppintegrationsEventIntegration].
func (aei *AppintegrationsEventIntegration) Type() string {
	return "aws_appintegrations_event_integration"
}

// LocalName returns the local name for [AppintegrationsEventIntegration].
func (aei *AppintegrationsEventIntegration) LocalName() string {
	return aei.Name
}

// Configuration returns the configuration (args) for [AppintegrationsEventIntegration].
func (aei *AppintegrationsEventIntegration) Configuration() interface{} {
	return aei.Args
}

// DependOn is used for other resources to depend on [AppintegrationsEventIntegration].
func (aei *AppintegrationsEventIntegration) DependOn() terra.Reference {
	return terra.ReferenceResource(aei)
}

// Dependencies returns the list of resources [AppintegrationsEventIntegration] depends_on.
func (aei *AppintegrationsEventIntegration) Dependencies() terra.Dependencies {
	return aei.DependsOn
}

// LifecycleManagement returns the lifecycle block for [AppintegrationsEventIntegration].
func (aei *AppintegrationsEventIntegration) LifecycleManagement() *terra.Lifecycle {
	return aei.Lifecycle
}

// Attributes returns the attributes for [AppintegrationsEventIntegration].
func (aei *AppintegrationsEventIntegration) Attributes() appintegrationsEventIntegrationAttributes {
	return appintegrationsEventIntegrationAttributes{ref: terra.ReferenceResource(aei)}
}

// ImportState imports the given attribute values into [AppintegrationsEventIntegration]'s state.
func (aei *AppintegrationsEventIntegration) ImportState(av io.Reader) error {
	aei.state = &appintegrationsEventIntegrationState{}
	if err := json.NewDecoder(av).Decode(aei.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", aei.Type(), aei.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [AppintegrationsEventIntegration] has state.
func (aei *AppintegrationsEventIntegration) State() (*appintegrationsEventIntegrationState, bool) {
	return aei.state, aei.state != nil
}

// StateMust returns the state for [AppintegrationsEventIntegration]. Panics if the state is nil.
func (aei *AppintegrationsEventIntegration) StateMust() *appintegrationsEventIntegrationState {
	if aei.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", aei.Type(), aei.LocalName()))
	}
	return aei.state
}

// AppintegrationsEventIntegrationArgs contains the configurations for aws_appintegrations_event_integration.
type AppintegrationsEventIntegrationArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// EventbridgeBus: string, required
	EventbridgeBus terra.StringValue `hcl:"eventbridge_bus,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// EventFilter: required
	EventFilter *appintegrationseventintegration.EventFilter `hcl:"event_filter,block" validate:"required"`
}
type appintegrationsEventIntegrationAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_appintegrations_event_integration.
func (aei appintegrationsEventIntegrationAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(aei.ref.Append("arn"))
}

// Description returns a reference to field description of aws_appintegrations_event_integration.
func (aei appintegrationsEventIntegrationAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(aei.ref.Append("description"))
}

// EventbridgeBus returns a reference to field eventbridge_bus of aws_appintegrations_event_integration.
func (aei appintegrationsEventIntegrationAttributes) EventbridgeBus() terra.StringValue {
	return terra.ReferenceAsString(aei.ref.Append("eventbridge_bus"))
}

// Id returns a reference to field id of aws_appintegrations_event_integration.
func (aei appintegrationsEventIntegrationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(aei.ref.Append("id"))
}

// Name returns a reference to field name of aws_appintegrations_event_integration.
func (aei appintegrationsEventIntegrationAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(aei.ref.Append("name"))
}

// Tags returns a reference to field tags of aws_appintegrations_event_integration.
func (aei appintegrationsEventIntegrationAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](aei.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_appintegrations_event_integration.
func (aei appintegrationsEventIntegrationAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](aei.ref.Append("tags_all"))
}

func (aei appintegrationsEventIntegrationAttributes) EventFilter() terra.ListValue[appintegrationseventintegration.EventFilterAttributes] {
	return terra.ReferenceAsList[appintegrationseventintegration.EventFilterAttributes](aei.ref.Append("event_filter"))
}

type appintegrationsEventIntegrationState struct {
	Arn            string                                             `json:"arn"`
	Description    string                                             `json:"description"`
	EventbridgeBus string                                             `json:"eventbridge_bus"`
	Id             string                                             `json:"id"`
	Name           string                                             `json:"name"`
	Tags           map[string]string                                  `json:"tags"`
	TagsAll        map[string]string                                  `json:"tags_all"`
	EventFilter    []appintegrationseventintegration.EventFilterState `json:"event_filter"`
}
