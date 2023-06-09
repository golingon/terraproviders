// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewPinpointApnsVoipSandboxChannel creates a new instance of [PinpointApnsVoipSandboxChannel].
func NewPinpointApnsVoipSandboxChannel(name string, args PinpointApnsVoipSandboxChannelArgs) *PinpointApnsVoipSandboxChannel {
	return &PinpointApnsVoipSandboxChannel{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*PinpointApnsVoipSandboxChannel)(nil)

// PinpointApnsVoipSandboxChannel represents the Terraform resource aws_pinpoint_apns_voip_sandbox_channel.
type PinpointApnsVoipSandboxChannel struct {
	Name      string
	Args      PinpointApnsVoipSandboxChannelArgs
	state     *pinpointApnsVoipSandboxChannelState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [PinpointApnsVoipSandboxChannel].
func (pavsc *PinpointApnsVoipSandboxChannel) Type() string {
	return "aws_pinpoint_apns_voip_sandbox_channel"
}

// LocalName returns the local name for [PinpointApnsVoipSandboxChannel].
func (pavsc *PinpointApnsVoipSandboxChannel) LocalName() string {
	return pavsc.Name
}

// Configuration returns the configuration (args) for [PinpointApnsVoipSandboxChannel].
func (pavsc *PinpointApnsVoipSandboxChannel) Configuration() interface{} {
	return pavsc.Args
}

// DependOn is used for other resources to depend on [PinpointApnsVoipSandboxChannel].
func (pavsc *PinpointApnsVoipSandboxChannel) DependOn() terra.Reference {
	return terra.ReferenceResource(pavsc)
}

// Dependencies returns the list of resources [PinpointApnsVoipSandboxChannel] depends_on.
func (pavsc *PinpointApnsVoipSandboxChannel) Dependencies() terra.Dependencies {
	return pavsc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [PinpointApnsVoipSandboxChannel].
func (pavsc *PinpointApnsVoipSandboxChannel) LifecycleManagement() *terra.Lifecycle {
	return pavsc.Lifecycle
}

// Attributes returns the attributes for [PinpointApnsVoipSandboxChannel].
func (pavsc *PinpointApnsVoipSandboxChannel) Attributes() pinpointApnsVoipSandboxChannelAttributes {
	return pinpointApnsVoipSandboxChannelAttributes{ref: terra.ReferenceResource(pavsc)}
}

// ImportState imports the given attribute values into [PinpointApnsVoipSandboxChannel]'s state.
func (pavsc *PinpointApnsVoipSandboxChannel) ImportState(av io.Reader) error {
	pavsc.state = &pinpointApnsVoipSandboxChannelState{}
	if err := json.NewDecoder(av).Decode(pavsc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", pavsc.Type(), pavsc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [PinpointApnsVoipSandboxChannel] has state.
func (pavsc *PinpointApnsVoipSandboxChannel) State() (*pinpointApnsVoipSandboxChannelState, bool) {
	return pavsc.state, pavsc.state != nil
}

// StateMust returns the state for [PinpointApnsVoipSandboxChannel]. Panics if the state is nil.
func (pavsc *PinpointApnsVoipSandboxChannel) StateMust() *pinpointApnsVoipSandboxChannelState {
	if pavsc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", pavsc.Type(), pavsc.LocalName()))
	}
	return pavsc.state
}

// PinpointApnsVoipSandboxChannelArgs contains the configurations for aws_pinpoint_apns_voip_sandbox_channel.
type PinpointApnsVoipSandboxChannelArgs struct {
	// ApplicationId: string, required
	ApplicationId terra.StringValue `hcl:"application_id,attr" validate:"required"`
	// BundleId: string, optional
	BundleId terra.StringValue `hcl:"bundle_id,attr"`
	// Certificate: string, optional
	Certificate terra.StringValue `hcl:"certificate,attr"`
	// DefaultAuthenticationMethod: string, optional
	DefaultAuthenticationMethod terra.StringValue `hcl:"default_authentication_method,attr"`
	// Enabled: bool, optional
	Enabled terra.BoolValue `hcl:"enabled,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// PrivateKey: string, optional
	PrivateKey terra.StringValue `hcl:"private_key,attr"`
	// TeamId: string, optional
	TeamId terra.StringValue `hcl:"team_id,attr"`
	// TokenKey: string, optional
	TokenKey terra.StringValue `hcl:"token_key,attr"`
	// TokenKeyId: string, optional
	TokenKeyId terra.StringValue `hcl:"token_key_id,attr"`
}
type pinpointApnsVoipSandboxChannelAttributes struct {
	ref terra.Reference
}

// ApplicationId returns a reference to field application_id of aws_pinpoint_apns_voip_sandbox_channel.
func (pavsc pinpointApnsVoipSandboxChannelAttributes) ApplicationId() terra.StringValue {
	return terra.ReferenceAsString(pavsc.ref.Append("application_id"))
}

// BundleId returns a reference to field bundle_id of aws_pinpoint_apns_voip_sandbox_channel.
func (pavsc pinpointApnsVoipSandboxChannelAttributes) BundleId() terra.StringValue {
	return terra.ReferenceAsString(pavsc.ref.Append("bundle_id"))
}

// Certificate returns a reference to field certificate of aws_pinpoint_apns_voip_sandbox_channel.
func (pavsc pinpointApnsVoipSandboxChannelAttributes) Certificate() terra.StringValue {
	return terra.ReferenceAsString(pavsc.ref.Append("certificate"))
}

// DefaultAuthenticationMethod returns a reference to field default_authentication_method of aws_pinpoint_apns_voip_sandbox_channel.
func (pavsc pinpointApnsVoipSandboxChannelAttributes) DefaultAuthenticationMethod() terra.StringValue {
	return terra.ReferenceAsString(pavsc.ref.Append("default_authentication_method"))
}

// Enabled returns a reference to field enabled of aws_pinpoint_apns_voip_sandbox_channel.
func (pavsc pinpointApnsVoipSandboxChannelAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(pavsc.ref.Append("enabled"))
}

// Id returns a reference to field id of aws_pinpoint_apns_voip_sandbox_channel.
func (pavsc pinpointApnsVoipSandboxChannelAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(pavsc.ref.Append("id"))
}

// PrivateKey returns a reference to field private_key of aws_pinpoint_apns_voip_sandbox_channel.
func (pavsc pinpointApnsVoipSandboxChannelAttributes) PrivateKey() terra.StringValue {
	return terra.ReferenceAsString(pavsc.ref.Append("private_key"))
}

// TeamId returns a reference to field team_id of aws_pinpoint_apns_voip_sandbox_channel.
func (pavsc pinpointApnsVoipSandboxChannelAttributes) TeamId() terra.StringValue {
	return terra.ReferenceAsString(pavsc.ref.Append("team_id"))
}

// TokenKey returns a reference to field token_key of aws_pinpoint_apns_voip_sandbox_channel.
func (pavsc pinpointApnsVoipSandboxChannelAttributes) TokenKey() terra.StringValue {
	return terra.ReferenceAsString(pavsc.ref.Append("token_key"))
}

// TokenKeyId returns a reference to field token_key_id of aws_pinpoint_apns_voip_sandbox_channel.
func (pavsc pinpointApnsVoipSandboxChannelAttributes) TokenKeyId() terra.StringValue {
	return terra.ReferenceAsString(pavsc.ref.Append("token_key_id"))
}

type pinpointApnsVoipSandboxChannelState struct {
	ApplicationId               string `json:"application_id"`
	BundleId                    string `json:"bundle_id"`
	Certificate                 string `json:"certificate"`
	DefaultAuthenticationMethod string `json:"default_authentication_method"`
	Enabled                     bool   `json:"enabled"`
	Id                          string `json:"id"`
	PrivateKey                  string `json:"private_key"`
	TeamId                      string `json:"team_id"`
	TokenKey                    string `json:"token_key"`
	TokenKeyId                  string `json:"token_key_id"`
}
