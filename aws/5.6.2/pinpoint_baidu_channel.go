// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewPinpointBaiduChannel creates a new instance of [PinpointBaiduChannel].
func NewPinpointBaiduChannel(name string, args PinpointBaiduChannelArgs) *PinpointBaiduChannel {
	return &PinpointBaiduChannel{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*PinpointBaiduChannel)(nil)

// PinpointBaiduChannel represents the Terraform resource aws_pinpoint_baidu_channel.
type PinpointBaiduChannel struct {
	Name      string
	Args      PinpointBaiduChannelArgs
	state     *pinpointBaiduChannelState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [PinpointBaiduChannel].
func (pbc *PinpointBaiduChannel) Type() string {
	return "aws_pinpoint_baidu_channel"
}

// LocalName returns the local name for [PinpointBaiduChannel].
func (pbc *PinpointBaiduChannel) LocalName() string {
	return pbc.Name
}

// Configuration returns the configuration (args) for [PinpointBaiduChannel].
func (pbc *PinpointBaiduChannel) Configuration() interface{} {
	return pbc.Args
}

// DependOn is used for other resources to depend on [PinpointBaiduChannel].
func (pbc *PinpointBaiduChannel) DependOn() terra.Reference {
	return terra.ReferenceResource(pbc)
}

// Dependencies returns the list of resources [PinpointBaiduChannel] depends_on.
func (pbc *PinpointBaiduChannel) Dependencies() terra.Dependencies {
	return pbc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [PinpointBaiduChannel].
func (pbc *PinpointBaiduChannel) LifecycleManagement() *terra.Lifecycle {
	return pbc.Lifecycle
}

// Attributes returns the attributes for [PinpointBaiduChannel].
func (pbc *PinpointBaiduChannel) Attributes() pinpointBaiduChannelAttributes {
	return pinpointBaiduChannelAttributes{ref: terra.ReferenceResource(pbc)}
}

// ImportState imports the given attribute values into [PinpointBaiduChannel]'s state.
func (pbc *PinpointBaiduChannel) ImportState(av io.Reader) error {
	pbc.state = &pinpointBaiduChannelState{}
	if err := json.NewDecoder(av).Decode(pbc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", pbc.Type(), pbc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [PinpointBaiduChannel] has state.
func (pbc *PinpointBaiduChannel) State() (*pinpointBaiduChannelState, bool) {
	return pbc.state, pbc.state != nil
}

// StateMust returns the state for [PinpointBaiduChannel]. Panics if the state is nil.
func (pbc *PinpointBaiduChannel) StateMust() *pinpointBaiduChannelState {
	if pbc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", pbc.Type(), pbc.LocalName()))
	}
	return pbc.state
}

// PinpointBaiduChannelArgs contains the configurations for aws_pinpoint_baidu_channel.
type PinpointBaiduChannelArgs struct {
	// ApiKey: string, required
	ApiKey terra.StringValue `hcl:"api_key,attr" validate:"required"`
	// ApplicationId: string, required
	ApplicationId terra.StringValue `hcl:"application_id,attr" validate:"required"`
	// Enabled: bool, optional
	Enabled terra.BoolValue `hcl:"enabled,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// SecretKey: string, required
	SecretKey terra.StringValue `hcl:"secret_key,attr" validate:"required"`
}
type pinpointBaiduChannelAttributes struct {
	ref terra.Reference
}

// ApiKey returns a reference to field api_key of aws_pinpoint_baidu_channel.
func (pbc pinpointBaiduChannelAttributes) ApiKey() terra.StringValue {
	return terra.ReferenceAsString(pbc.ref.Append("api_key"))
}

// ApplicationId returns a reference to field application_id of aws_pinpoint_baidu_channel.
func (pbc pinpointBaiduChannelAttributes) ApplicationId() terra.StringValue {
	return terra.ReferenceAsString(pbc.ref.Append("application_id"))
}

// Enabled returns a reference to field enabled of aws_pinpoint_baidu_channel.
func (pbc pinpointBaiduChannelAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(pbc.ref.Append("enabled"))
}

// Id returns a reference to field id of aws_pinpoint_baidu_channel.
func (pbc pinpointBaiduChannelAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(pbc.ref.Append("id"))
}

// SecretKey returns a reference to field secret_key of aws_pinpoint_baidu_channel.
func (pbc pinpointBaiduChannelAttributes) SecretKey() terra.StringValue {
	return terra.ReferenceAsString(pbc.ref.Append("secret_key"))
}

type pinpointBaiduChannelState struct {
	ApiKey        string `json:"api_key"`
	ApplicationId string `json:"application_id"`
	Enabled       bool   `json:"enabled"`
	Id            string `json:"id"`
	SecretKey     string `json:"secret_key"`
}