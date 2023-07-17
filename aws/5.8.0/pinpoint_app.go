// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	pinpointapp "github.com/golingon/terraproviders/aws/5.8.0/pinpointapp"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewPinpointApp creates a new instance of [PinpointApp].
func NewPinpointApp(name string, args PinpointAppArgs) *PinpointApp {
	return &PinpointApp{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*PinpointApp)(nil)

// PinpointApp represents the Terraform resource aws_pinpoint_app.
type PinpointApp struct {
	Name      string
	Args      PinpointAppArgs
	state     *pinpointAppState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [PinpointApp].
func (pa *PinpointApp) Type() string {
	return "aws_pinpoint_app"
}

// LocalName returns the local name for [PinpointApp].
func (pa *PinpointApp) LocalName() string {
	return pa.Name
}

// Configuration returns the configuration (args) for [PinpointApp].
func (pa *PinpointApp) Configuration() interface{} {
	return pa.Args
}

// DependOn is used for other resources to depend on [PinpointApp].
func (pa *PinpointApp) DependOn() terra.Reference {
	return terra.ReferenceResource(pa)
}

// Dependencies returns the list of resources [PinpointApp] depends_on.
func (pa *PinpointApp) Dependencies() terra.Dependencies {
	return pa.DependsOn
}

// LifecycleManagement returns the lifecycle block for [PinpointApp].
func (pa *PinpointApp) LifecycleManagement() *terra.Lifecycle {
	return pa.Lifecycle
}

// Attributes returns the attributes for [PinpointApp].
func (pa *PinpointApp) Attributes() pinpointAppAttributes {
	return pinpointAppAttributes{ref: terra.ReferenceResource(pa)}
}

// ImportState imports the given attribute values into [PinpointApp]'s state.
func (pa *PinpointApp) ImportState(av io.Reader) error {
	pa.state = &pinpointAppState{}
	if err := json.NewDecoder(av).Decode(pa.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", pa.Type(), pa.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [PinpointApp] has state.
func (pa *PinpointApp) State() (*pinpointAppState, bool) {
	return pa.state, pa.state != nil
}

// StateMust returns the state for [PinpointApp]. Panics if the state is nil.
func (pa *PinpointApp) StateMust() *pinpointAppState {
	if pa.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", pa.Type(), pa.LocalName()))
	}
	return pa.state
}

// PinpointAppArgs contains the configurations for aws_pinpoint_app.
type PinpointAppArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
	// NamePrefix: string, optional
	NamePrefix terra.StringValue `hcl:"name_prefix,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// CampaignHook: optional
	CampaignHook *pinpointapp.CampaignHook `hcl:"campaign_hook,block"`
	// Limits: optional
	Limits *pinpointapp.Limits `hcl:"limits,block"`
	// QuietTime: optional
	QuietTime *pinpointapp.QuietTime `hcl:"quiet_time,block"`
}
type pinpointAppAttributes struct {
	ref terra.Reference
}

// ApplicationId returns a reference to field application_id of aws_pinpoint_app.
func (pa pinpointAppAttributes) ApplicationId() terra.StringValue {
	return terra.ReferenceAsString(pa.ref.Append("application_id"))
}

// Arn returns a reference to field arn of aws_pinpoint_app.
func (pa pinpointAppAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(pa.ref.Append("arn"))
}

// Id returns a reference to field id of aws_pinpoint_app.
func (pa pinpointAppAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(pa.ref.Append("id"))
}

// Name returns a reference to field name of aws_pinpoint_app.
func (pa pinpointAppAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(pa.ref.Append("name"))
}

// NamePrefix returns a reference to field name_prefix of aws_pinpoint_app.
func (pa pinpointAppAttributes) NamePrefix() terra.StringValue {
	return terra.ReferenceAsString(pa.ref.Append("name_prefix"))
}

// Tags returns a reference to field tags of aws_pinpoint_app.
func (pa pinpointAppAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](pa.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_pinpoint_app.
func (pa pinpointAppAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](pa.ref.Append("tags_all"))
}

func (pa pinpointAppAttributes) CampaignHook() terra.ListValue[pinpointapp.CampaignHookAttributes] {
	return terra.ReferenceAsList[pinpointapp.CampaignHookAttributes](pa.ref.Append("campaign_hook"))
}

func (pa pinpointAppAttributes) Limits() terra.ListValue[pinpointapp.LimitsAttributes] {
	return terra.ReferenceAsList[pinpointapp.LimitsAttributes](pa.ref.Append("limits"))
}

func (pa pinpointAppAttributes) QuietTime() terra.ListValue[pinpointapp.QuietTimeAttributes] {
	return terra.ReferenceAsList[pinpointapp.QuietTimeAttributes](pa.ref.Append("quiet_time"))
}

type pinpointAppState struct {
	ApplicationId string                          `json:"application_id"`
	Arn           string                          `json:"arn"`
	Id            string                          `json:"id"`
	Name          string                          `json:"name"`
	NamePrefix    string                          `json:"name_prefix"`
	Tags          map[string]string               `json:"tags"`
	TagsAll       map[string]string               `json:"tags_all"`
	CampaignHook  []pinpointapp.CampaignHookState `json:"campaign_hook"`
	Limits        []pinpointapp.LimitsState       `json:"limits"`
	QuietTime     []pinpointapp.QuietTimeState    `json:"quiet_time"`
}
