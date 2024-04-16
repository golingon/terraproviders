// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_medialive_input

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

// Resource represents the Terraform resource aws_medialive_input.
type Resource struct {
	Name      string
	Args      Args
	state     *awsMedialiveInputState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (ami *Resource) Type() string {
	return "aws_medialive_input"
}

// LocalName returns the local name for [Resource].
func (ami *Resource) LocalName() string {
	return ami.Name
}

// Configuration returns the configuration (args) for [Resource].
func (ami *Resource) Configuration() interface{} {
	return ami.Args
}

// DependOn is used for other resources to depend on [Resource].
func (ami *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(ami)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (ami *Resource) Dependencies() terra.Dependencies {
	return ami.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (ami *Resource) LifecycleManagement() *terra.Lifecycle {
	return ami.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (ami *Resource) Attributes() awsMedialiveInputAttributes {
	return awsMedialiveInputAttributes{ref: terra.ReferenceResource(ami)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (ami *Resource) ImportState(state io.Reader) error {
	ami.state = &awsMedialiveInputState{}
	if err := json.NewDecoder(state).Decode(ami.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ami.Type(), ami.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (ami *Resource) State() (*awsMedialiveInputState, bool) {
	return ami.state, ami.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (ami *Resource) StateMust() *awsMedialiveInputState {
	if ami.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ami.Type(), ami.LocalName()))
	}
	return ami.state
}

// Args contains the configurations for aws_medialive_input.
type Args struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// InputSecurityGroups: list of string, optional
	InputSecurityGroups terra.ListValue[terra.StringValue] `hcl:"input_security_groups,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// RoleArn: string, optional
	RoleArn terra.StringValue `hcl:"role_arn,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// Type: string, required
	Type terra.StringValue `hcl:"type,attr" validate:"required"`
	// Destinations: min=0
	Destinations []Destinations `hcl:"destinations,block" validate:"min=0"`
	// InputDevices: min=0
	InputDevices []InputDevices `hcl:"input_devices,block" validate:"min=0"`
	// MediaConnectFlows: min=0
	MediaConnectFlows []MediaConnectFlows `hcl:"media_connect_flows,block" validate:"min=0"`
	// Sources: min=0
	Sources []Sources `hcl:"sources,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
	// Vpc: optional
	Vpc *Vpc `hcl:"vpc,block"`
}

type awsMedialiveInputAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_medialive_input.
func (ami awsMedialiveInputAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(ami.ref.Append("arn"))
}

// AttachedChannels returns a reference to field attached_channels of aws_medialive_input.
func (ami awsMedialiveInputAttributes) AttachedChannels() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](ami.ref.Append("attached_channels"))
}

// Id returns a reference to field id of aws_medialive_input.
func (ami awsMedialiveInputAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ami.ref.Append("id"))
}

// InputClass returns a reference to field input_class of aws_medialive_input.
func (ami awsMedialiveInputAttributes) InputClass() terra.StringValue {
	return terra.ReferenceAsString(ami.ref.Append("input_class"))
}

// InputPartnerIds returns a reference to field input_partner_ids of aws_medialive_input.
func (ami awsMedialiveInputAttributes) InputPartnerIds() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](ami.ref.Append("input_partner_ids"))
}

// InputSecurityGroups returns a reference to field input_security_groups of aws_medialive_input.
func (ami awsMedialiveInputAttributes) InputSecurityGroups() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](ami.ref.Append("input_security_groups"))
}

// InputSourceType returns a reference to field input_source_type of aws_medialive_input.
func (ami awsMedialiveInputAttributes) InputSourceType() terra.StringValue {
	return terra.ReferenceAsString(ami.ref.Append("input_source_type"))
}

// Name returns a reference to field name of aws_medialive_input.
func (ami awsMedialiveInputAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ami.ref.Append("name"))
}

// RoleArn returns a reference to field role_arn of aws_medialive_input.
func (ami awsMedialiveInputAttributes) RoleArn() terra.StringValue {
	return terra.ReferenceAsString(ami.ref.Append("role_arn"))
}

// Tags returns a reference to field tags of aws_medialive_input.
func (ami awsMedialiveInputAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ami.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_medialive_input.
func (ami awsMedialiveInputAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ami.ref.Append("tags_all"))
}

// Type returns a reference to field type of aws_medialive_input.
func (ami awsMedialiveInputAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(ami.ref.Append("type"))
}

func (ami awsMedialiveInputAttributes) Destinations() terra.SetValue[DestinationsAttributes] {
	return terra.ReferenceAsSet[DestinationsAttributes](ami.ref.Append("destinations"))
}

func (ami awsMedialiveInputAttributes) InputDevices() terra.SetValue[InputDevicesAttributes] {
	return terra.ReferenceAsSet[InputDevicesAttributes](ami.ref.Append("input_devices"))
}

func (ami awsMedialiveInputAttributes) MediaConnectFlows() terra.SetValue[MediaConnectFlowsAttributes] {
	return terra.ReferenceAsSet[MediaConnectFlowsAttributes](ami.ref.Append("media_connect_flows"))
}

func (ami awsMedialiveInputAttributes) Sources() terra.SetValue[SourcesAttributes] {
	return terra.ReferenceAsSet[SourcesAttributes](ami.ref.Append("sources"))
}

func (ami awsMedialiveInputAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](ami.ref.Append("timeouts"))
}

func (ami awsMedialiveInputAttributes) Vpc() terra.ListValue[VpcAttributes] {
	return terra.ReferenceAsList[VpcAttributes](ami.ref.Append("vpc"))
}

type awsMedialiveInputState struct {
	Arn                 string                   `json:"arn"`
	AttachedChannels    []string                 `json:"attached_channels"`
	Id                  string                   `json:"id"`
	InputClass          string                   `json:"input_class"`
	InputPartnerIds     []string                 `json:"input_partner_ids"`
	InputSecurityGroups []string                 `json:"input_security_groups"`
	InputSourceType     string                   `json:"input_source_type"`
	Name                string                   `json:"name"`
	RoleArn             string                   `json:"role_arn"`
	Tags                map[string]string        `json:"tags"`
	TagsAll             map[string]string        `json:"tags_all"`
	Type                string                   `json:"type"`
	Destinations        []DestinationsState      `json:"destinations"`
	InputDevices        []InputDevicesState      `json:"input_devices"`
	MediaConnectFlows   []MediaConnectFlowsState `json:"media_connect_flows"`
	Sources             []SourcesState           `json:"sources"`
	Timeouts            *TimeoutsState           `json:"timeouts"`
	Vpc                 []VpcState               `json:"vpc"`
}
