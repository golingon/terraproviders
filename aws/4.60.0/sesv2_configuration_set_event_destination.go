// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	sesv2configurationseteventdestination "github.com/golingon/terraproviders/aws/4.60.0/sesv2configurationseteventdestination"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewSesv2ConfigurationSetEventDestination(name string, args Sesv2ConfigurationSetEventDestinationArgs) *Sesv2ConfigurationSetEventDestination {
	return &Sesv2ConfigurationSetEventDestination{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Sesv2ConfigurationSetEventDestination)(nil)

type Sesv2ConfigurationSetEventDestination struct {
	Name  string
	Args  Sesv2ConfigurationSetEventDestinationArgs
	state *sesv2ConfigurationSetEventDestinationState
}

func (scsed *Sesv2ConfigurationSetEventDestination) Type() string {
	return "aws_sesv2_configuration_set_event_destination"
}

func (scsed *Sesv2ConfigurationSetEventDestination) LocalName() string {
	return scsed.Name
}

func (scsed *Sesv2ConfigurationSetEventDestination) Configuration() interface{} {
	return scsed.Args
}

func (scsed *Sesv2ConfigurationSetEventDestination) Attributes() sesv2ConfigurationSetEventDestinationAttributes {
	return sesv2ConfigurationSetEventDestinationAttributes{ref: terra.ReferenceResource(scsed)}
}

func (scsed *Sesv2ConfigurationSetEventDestination) ImportState(av io.Reader) error {
	scsed.state = &sesv2ConfigurationSetEventDestinationState{}
	if err := json.NewDecoder(av).Decode(scsed.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", scsed.Type(), scsed.LocalName(), err)
	}
	return nil
}

func (scsed *Sesv2ConfigurationSetEventDestination) State() (*sesv2ConfigurationSetEventDestinationState, bool) {
	return scsed.state, scsed.state != nil
}

func (scsed *Sesv2ConfigurationSetEventDestination) StateMust() *sesv2ConfigurationSetEventDestinationState {
	if scsed.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", scsed.Type(), scsed.LocalName()))
	}
	return scsed.state
}

func (scsed *Sesv2ConfigurationSetEventDestination) DependOn() terra.Reference {
	return terra.ReferenceResource(scsed)
}

type Sesv2ConfigurationSetEventDestinationArgs struct {
	// ConfigurationSetName: string, required
	ConfigurationSetName terra.StringValue `hcl:"configuration_set_name,attr" validate:"required"`
	// EventDestinationName: string, required
	EventDestinationName terra.StringValue `hcl:"event_destination_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// EventDestination: required
	EventDestination *sesv2configurationseteventdestination.EventDestination `hcl:"event_destination,block" validate:"required"`
	// DependsOn contains resources that Sesv2ConfigurationSetEventDestination depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type sesv2ConfigurationSetEventDestinationAttributes struct {
	ref terra.Reference
}

func (scsed sesv2ConfigurationSetEventDestinationAttributes) ConfigurationSetName() terra.StringValue {
	return terra.ReferenceString(scsed.ref.Append("configuration_set_name"))
}

func (scsed sesv2ConfigurationSetEventDestinationAttributes) EventDestinationName() terra.StringValue {
	return terra.ReferenceString(scsed.ref.Append("event_destination_name"))
}

func (scsed sesv2ConfigurationSetEventDestinationAttributes) Id() terra.StringValue {
	return terra.ReferenceString(scsed.ref.Append("id"))
}

func (scsed sesv2ConfigurationSetEventDestinationAttributes) EventDestination() terra.ListValue[sesv2configurationseteventdestination.EventDestinationAttributes] {
	return terra.ReferenceList[sesv2configurationseteventdestination.EventDestinationAttributes](scsed.ref.Append("event_destination"))
}

type sesv2ConfigurationSetEventDestinationState struct {
	ConfigurationSetName string                                                        `json:"configuration_set_name"`
	EventDestinationName string                                                        `json:"event_destination_name"`
	Id                   string                                                        `json:"id"`
	EventDestination     []sesv2configurationseteventdestination.EventDestinationState `json:"event_destination"`
}
