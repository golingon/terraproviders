// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_appstream_stack

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

// Resource represents the Terraform resource aws_appstream_stack.
type Resource struct {
	Name      string
	Args      Args
	state     *awsAppstreamStackState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (aas *Resource) Type() string {
	return "aws_appstream_stack"
}

// LocalName returns the local name for [Resource].
func (aas *Resource) LocalName() string {
	return aas.Name
}

// Configuration returns the configuration (args) for [Resource].
func (aas *Resource) Configuration() interface{} {
	return aas.Args
}

// DependOn is used for other resources to depend on [Resource].
func (aas *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(aas)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (aas *Resource) Dependencies() terra.Dependencies {
	return aas.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (aas *Resource) LifecycleManagement() *terra.Lifecycle {
	return aas.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (aas *Resource) Attributes() awsAppstreamStackAttributes {
	return awsAppstreamStackAttributes{ref: terra.ReferenceResource(aas)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (aas *Resource) ImportState(state io.Reader) error {
	aas.state = &awsAppstreamStackState{}
	if err := json.NewDecoder(state).Decode(aas.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", aas.Type(), aas.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (aas *Resource) State() (*awsAppstreamStackState, bool) {
	return aas.state, aas.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (aas *Resource) StateMust() *awsAppstreamStackState {
	if aas.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", aas.Type(), aas.LocalName()))
	}
	return aas.state
}

// Args contains the configurations for aws_appstream_stack.
type Args struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// DisplayName: string, optional
	DisplayName terra.StringValue `hcl:"display_name,attr"`
	// EmbedHostDomains: set of string, optional
	EmbedHostDomains terra.SetValue[terra.StringValue] `hcl:"embed_host_domains,attr"`
	// FeedbackUrl: string, optional
	FeedbackUrl terra.StringValue `hcl:"feedback_url,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// RedirectUrl: string, optional
	RedirectUrl terra.StringValue `hcl:"redirect_url,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// AccessEndpoints: min=0,max=4
	AccessEndpoints []AccessEndpoints `hcl:"access_endpoints,block" validate:"min=0,max=4"`
	// ApplicationSettings: optional
	ApplicationSettings *ApplicationSettings `hcl:"application_settings,block"`
	// StorageConnectors: min=0
	StorageConnectors []StorageConnectors `hcl:"storage_connectors,block" validate:"min=0"`
	// StreamingExperienceSettings: optional
	StreamingExperienceSettings *StreamingExperienceSettings `hcl:"streaming_experience_settings,block"`
	// UserSettings: min=0
	UserSettings []UserSettings `hcl:"user_settings,block" validate:"min=0"`
}

type awsAppstreamStackAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_appstream_stack.
func (aas awsAppstreamStackAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(aas.ref.Append("arn"))
}

// CreatedTime returns a reference to field created_time of aws_appstream_stack.
func (aas awsAppstreamStackAttributes) CreatedTime() terra.StringValue {
	return terra.ReferenceAsString(aas.ref.Append("created_time"))
}

// Description returns a reference to field description of aws_appstream_stack.
func (aas awsAppstreamStackAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(aas.ref.Append("description"))
}

// DisplayName returns a reference to field display_name of aws_appstream_stack.
func (aas awsAppstreamStackAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(aas.ref.Append("display_name"))
}

// EmbedHostDomains returns a reference to field embed_host_domains of aws_appstream_stack.
func (aas awsAppstreamStackAttributes) EmbedHostDomains() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](aas.ref.Append("embed_host_domains"))
}

// FeedbackUrl returns a reference to field feedback_url of aws_appstream_stack.
func (aas awsAppstreamStackAttributes) FeedbackUrl() terra.StringValue {
	return terra.ReferenceAsString(aas.ref.Append("feedback_url"))
}

// Id returns a reference to field id of aws_appstream_stack.
func (aas awsAppstreamStackAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(aas.ref.Append("id"))
}

// Name returns a reference to field name of aws_appstream_stack.
func (aas awsAppstreamStackAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(aas.ref.Append("name"))
}

// RedirectUrl returns a reference to field redirect_url of aws_appstream_stack.
func (aas awsAppstreamStackAttributes) RedirectUrl() terra.StringValue {
	return terra.ReferenceAsString(aas.ref.Append("redirect_url"))
}

// Tags returns a reference to field tags of aws_appstream_stack.
func (aas awsAppstreamStackAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](aas.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_appstream_stack.
func (aas awsAppstreamStackAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](aas.ref.Append("tags_all"))
}

func (aas awsAppstreamStackAttributes) AccessEndpoints() terra.SetValue[AccessEndpointsAttributes] {
	return terra.ReferenceAsSet[AccessEndpointsAttributes](aas.ref.Append("access_endpoints"))
}

func (aas awsAppstreamStackAttributes) ApplicationSettings() terra.ListValue[ApplicationSettingsAttributes] {
	return terra.ReferenceAsList[ApplicationSettingsAttributes](aas.ref.Append("application_settings"))
}

func (aas awsAppstreamStackAttributes) StorageConnectors() terra.SetValue[StorageConnectorsAttributes] {
	return terra.ReferenceAsSet[StorageConnectorsAttributes](aas.ref.Append("storage_connectors"))
}

func (aas awsAppstreamStackAttributes) StreamingExperienceSettings() terra.ListValue[StreamingExperienceSettingsAttributes] {
	return terra.ReferenceAsList[StreamingExperienceSettingsAttributes](aas.ref.Append("streaming_experience_settings"))
}

func (aas awsAppstreamStackAttributes) UserSettings() terra.SetValue[UserSettingsAttributes] {
	return terra.ReferenceAsSet[UserSettingsAttributes](aas.ref.Append("user_settings"))
}

type awsAppstreamStackState struct {
	Arn                         string                             `json:"arn"`
	CreatedTime                 string                             `json:"created_time"`
	Description                 string                             `json:"description"`
	DisplayName                 string                             `json:"display_name"`
	EmbedHostDomains            []string                           `json:"embed_host_domains"`
	FeedbackUrl                 string                             `json:"feedback_url"`
	Id                          string                             `json:"id"`
	Name                        string                             `json:"name"`
	RedirectUrl                 string                             `json:"redirect_url"`
	Tags                        map[string]string                  `json:"tags"`
	TagsAll                     map[string]string                  `json:"tags_all"`
	AccessEndpoints             []AccessEndpointsState             `json:"access_endpoints"`
	ApplicationSettings         []ApplicationSettingsState         `json:"application_settings"`
	StorageConnectors           []StorageConnectorsState           `json:"storage_connectors"`
	StreamingExperienceSettings []StreamingExperienceSettingsState `json:"streaming_experience_settings"`
	UserSettings                []UserSettingsState                `json:"user_settings"`
}
