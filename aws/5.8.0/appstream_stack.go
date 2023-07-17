// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	appstreamstack "github.com/golingon/terraproviders/aws/5.8.0/appstreamstack"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewAppstreamStack creates a new instance of [AppstreamStack].
func NewAppstreamStack(name string, args AppstreamStackArgs) *AppstreamStack {
	return &AppstreamStack{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*AppstreamStack)(nil)

// AppstreamStack represents the Terraform resource aws_appstream_stack.
type AppstreamStack struct {
	Name      string
	Args      AppstreamStackArgs
	state     *appstreamStackState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [AppstreamStack].
func (as *AppstreamStack) Type() string {
	return "aws_appstream_stack"
}

// LocalName returns the local name for [AppstreamStack].
func (as *AppstreamStack) LocalName() string {
	return as.Name
}

// Configuration returns the configuration (args) for [AppstreamStack].
func (as *AppstreamStack) Configuration() interface{} {
	return as.Args
}

// DependOn is used for other resources to depend on [AppstreamStack].
func (as *AppstreamStack) DependOn() terra.Reference {
	return terra.ReferenceResource(as)
}

// Dependencies returns the list of resources [AppstreamStack] depends_on.
func (as *AppstreamStack) Dependencies() terra.Dependencies {
	return as.DependsOn
}

// LifecycleManagement returns the lifecycle block for [AppstreamStack].
func (as *AppstreamStack) LifecycleManagement() *terra.Lifecycle {
	return as.Lifecycle
}

// Attributes returns the attributes for [AppstreamStack].
func (as *AppstreamStack) Attributes() appstreamStackAttributes {
	return appstreamStackAttributes{ref: terra.ReferenceResource(as)}
}

// ImportState imports the given attribute values into [AppstreamStack]'s state.
func (as *AppstreamStack) ImportState(av io.Reader) error {
	as.state = &appstreamStackState{}
	if err := json.NewDecoder(av).Decode(as.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", as.Type(), as.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [AppstreamStack] has state.
func (as *AppstreamStack) State() (*appstreamStackState, bool) {
	return as.state, as.state != nil
}

// StateMust returns the state for [AppstreamStack]. Panics if the state is nil.
func (as *AppstreamStack) StateMust() *appstreamStackState {
	if as.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", as.Type(), as.LocalName()))
	}
	return as.state
}

// AppstreamStackArgs contains the configurations for aws_appstream_stack.
type AppstreamStackArgs struct {
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
	AccessEndpoints []appstreamstack.AccessEndpoints `hcl:"access_endpoints,block" validate:"min=0,max=4"`
	// ApplicationSettings: optional
	ApplicationSettings *appstreamstack.ApplicationSettings `hcl:"application_settings,block"`
	// StorageConnectors: min=0
	StorageConnectors []appstreamstack.StorageConnectors `hcl:"storage_connectors,block" validate:"min=0"`
	// StreamingExperienceSettings: optional
	StreamingExperienceSettings *appstreamstack.StreamingExperienceSettings `hcl:"streaming_experience_settings,block"`
	// UserSettings: min=0
	UserSettings []appstreamstack.UserSettings `hcl:"user_settings,block" validate:"min=0"`
}
type appstreamStackAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_appstream_stack.
func (as appstreamStackAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(as.ref.Append("arn"))
}

// CreatedTime returns a reference to field created_time of aws_appstream_stack.
func (as appstreamStackAttributes) CreatedTime() terra.StringValue {
	return terra.ReferenceAsString(as.ref.Append("created_time"))
}

// Description returns a reference to field description of aws_appstream_stack.
func (as appstreamStackAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(as.ref.Append("description"))
}

// DisplayName returns a reference to field display_name of aws_appstream_stack.
func (as appstreamStackAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(as.ref.Append("display_name"))
}

// EmbedHostDomains returns a reference to field embed_host_domains of aws_appstream_stack.
func (as appstreamStackAttributes) EmbedHostDomains() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](as.ref.Append("embed_host_domains"))
}

// FeedbackUrl returns a reference to field feedback_url of aws_appstream_stack.
func (as appstreamStackAttributes) FeedbackUrl() terra.StringValue {
	return terra.ReferenceAsString(as.ref.Append("feedback_url"))
}

// Id returns a reference to field id of aws_appstream_stack.
func (as appstreamStackAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(as.ref.Append("id"))
}

// Name returns a reference to field name of aws_appstream_stack.
func (as appstreamStackAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(as.ref.Append("name"))
}

// RedirectUrl returns a reference to field redirect_url of aws_appstream_stack.
func (as appstreamStackAttributes) RedirectUrl() terra.StringValue {
	return terra.ReferenceAsString(as.ref.Append("redirect_url"))
}

// Tags returns a reference to field tags of aws_appstream_stack.
func (as appstreamStackAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](as.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_appstream_stack.
func (as appstreamStackAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](as.ref.Append("tags_all"))
}

func (as appstreamStackAttributes) AccessEndpoints() terra.SetValue[appstreamstack.AccessEndpointsAttributes] {
	return terra.ReferenceAsSet[appstreamstack.AccessEndpointsAttributes](as.ref.Append("access_endpoints"))
}

func (as appstreamStackAttributes) ApplicationSettings() terra.ListValue[appstreamstack.ApplicationSettingsAttributes] {
	return terra.ReferenceAsList[appstreamstack.ApplicationSettingsAttributes](as.ref.Append("application_settings"))
}

func (as appstreamStackAttributes) StorageConnectors() terra.SetValue[appstreamstack.StorageConnectorsAttributes] {
	return terra.ReferenceAsSet[appstreamstack.StorageConnectorsAttributes](as.ref.Append("storage_connectors"))
}

func (as appstreamStackAttributes) StreamingExperienceSettings() terra.ListValue[appstreamstack.StreamingExperienceSettingsAttributes] {
	return terra.ReferenceAsList[appstreamstack.StreamingExperienceSettingsAttributes](as.ref.Append("streaming_experience_settings"))
}

func (as appstreamStackAttributes) UserSettings() terra.SetValue[appstreamstack.UserSettingsAttributes] {
	return terra.ReferenceAsSet[appstreamstack.UserSettingsAttributes](as.ref.Append("user_settings"))
}

type appstreamStackState struct {
	Arn                         string                                            `json:"arn"`
	CreatedTime                 string                                            `json:"created_time"`
	Description                 string                                            `json:"description"`
	DisplayName                 string                                            `json:"display_name"`
	EmbedHostDomains            []string                                          `json:"embed_host_domains"`
	FeedbackUrl                 string                                            `json:"feedback_url"`
	Id                          string                                            `json:"id"`
	Name                        string                                            `json:"name"`
	RedirectUrl                 string                                            `json:"redirect_url"`
	Tags                        map[string]string                                 `json:"tags"`
	TagsAll                     map[string]string                                 `json:"tags_all"`
	AccessEndpoints             []appstreamstack.AccessEndpointsState             `json:"access_endpoints"`
	ApplicationSettings         []appstreamstack.ApplicationSettingsState         `json:"application_settings"`
	StorageConnectors           []appstreamstack.StorageConnectorsState           `json:"storage_connectors"`
	StreamingExperienceSettings []appstreamstack.StreamingExperienceSettingsState `json:"streaming_experience_settings"`
	UserSettings                []appstreamstack.UserSettingsState                `json:"user_settings"`
}
