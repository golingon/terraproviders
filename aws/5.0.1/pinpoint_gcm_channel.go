// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewPinpointGcmChannel creates a new instance of [PinpointGcmChannel].
func NewPinpointGcmChannel(name string, args PinpointGcmChannelArgs) *PinpointGcmChannel {
	return &PinpointGcmChannel{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*PinpointGcmChannel)(nil)

// PinpointGcmChannel represents the Terraform resource aws_pinpoint_gcm_channel.
type PinpointGcmChannel struct {
	Name      string
	Args      PinpointGcmChannelArgs
	state     *pinpointGcmChannelState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [PinpointGcmChannel].
func (pgc *PinpointGcmChannel) Type() string {
	return "aws_pinpoint_gcm_channel"
}

// LocalName returns the local name for [PinpointGcmChannel].
func (pgc *PinpointGcmChannel) LocalName() string {
	return pgc.Name
}

// Configuration returns the configuration (args) for [PinpointGcmChannel].
func (pgc *PinpointGcmChannel) Configuration() interface{} {
	return pgc.Args
}

// DependOn is used for other resources to depend on [PinpointGcmChannel].
func (pgc *PinpointGcmChannel) DependOn() terra.Reference {
	return terra.ReferenceResource(pgc)
}

// Dependencies returns the list of resources [PinpointGcmChannel] depends_on.
func (pgc *PinpointGcmChannel) Dependencies() terra.Dependencies {
	return pgc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [PinpointGcmChannel].
func (pgc *PinpointGcmChannel) LifecycleManagement() *terra.Lifecycle {
	return pgc.Lifecycle
}

// Attributes returns the attributes for [PinpointGcmChannel].
func (pgc *PinpointGcmChannel) Attributes() pinpointGcmChannelAttributes {
	return pinpointGcmChannelAttributes{ref: terra.ReferenceResource(pgc)}
}

// ImportState imports the given attribute values into [PinpointGcmChannel]'s state.
func (pgc *PinpointGcmChannel) ImportState(av io.Reader) error {
	pgc.state = &pinpointGcmChannelState{}
	if err := json.NewDecoder(av).Decode(pgc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", pgc.Type(), pgc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [PinpointGcmChannel] has state.
func (pgc *PinpointGcmChannel) State() (*pinpointGcmChannelState, bool) {
	return pgc.state, pgc.state != nil
}

// StateMust returns the state for [PinpointGcmChannel]. Panics if the state is nil.
func (pgc *PinpointGcmChannel) StateMust() *pinpointGcmChannelState {
	if pgc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", pgc.Type(), pgc.LocalName()))
	}
	return pgc.state
}

// PinpointGcmChannelArgs contains the configurations for aws_pinpoint_gcm_channel.
type PinpointGcmChannelArgs struct {
	// ApiKey: string, required
	ApiKey terra.StringValue `hcl:"api_key,attr" validate:"required"`
	// ApplicationId: string, required
	ApplicationId terra.StringValue `hcl:"application_id,attr" validate:"required"`
	// Enabled: bool, optional
	Enabled terra.BoolValue `hcl:"enabled,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
}
type pinpointGcmChannelAttributes struct {
	ref terra.Reference
}

// ApiKey returns a reference to field api_key of aws_pinpoint_gcm_channel.
func (pgc pinpointGcmChannelAttributes) ApiKey() terra.StringValue {
	return terra.ReferenceAsString(pgc.ref.Append("api_key"))
}

// ApplicationId returns a reference to field application_id of aws_pinpoint_gcm_channel.
func (pgc pinpointGcmChannelAttributes) ApplicationId() terra.StringValue {
	return terra.ReferenceAsString(pgc.ref.Append("application_id"))
}

// Enabled returns a reference to field enabled of aws_pinpoint_gcm_channel.
func (pgc pinpointGcmChannelAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(pgc.ref.Append("enabled"))
}

// Id returns a reference to field id of aws_pinpoint_gcm_channel.
func (pgc pinpointGcmChannelAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(pgc.ref.Append("id"))
}

type pinpointGcmChannelState struct {
	ApiKey        string `json:"api_key"`
	ApplicationId string `json:"application_id"`
	Enabled       bool   `json:"enabled"`
	Id            string `json:"id"`
}
