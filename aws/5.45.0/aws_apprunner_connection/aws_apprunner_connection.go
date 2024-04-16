// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_apprunner_connection

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

// Resource represents the Terraform resource aws_apprunner_connection.
type Resource struct {
	Name      string
	Args      Args
	state     *awsApprunnerConnectionState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (aac *Resource) Type() string {
	return "aws_apprunner_connection"
}

// LocalName returns the local name for [Resource].
func (aac *Resource) LocalName() string {
	return aac.Name
}

// Configuration returns the configuration (args) for [Resource].
func (aac *Resource) Configuration() interface{} {
	return aac.Args
}

// DependOn is used for other resources to depend on [Resource].
func (aac *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(aac)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (aac *Resource) Dependencies() terra.Dependencies {
	return aac.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (aac *Resource) LifecycleManagement() *terra.Lifecycle {
	return aac.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (aac *Resource) Attributes() awsApprunnerConnectionAttributes {
	return awsApprunnerConnectionAttributes{ref: terra.ReferenceResource(aac)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (aac *Resource) ImportState(state io.Reader) error {
	aac.state = &awsApprunnerConnectionState{}
	if err := json.NewDecoder(state).Decode(aac.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", aac.Type(), aac.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (aac *Resource) State() (*awsApprunnerConnectionState, bool) {
	return aac.state, aac.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (aac *Resource) StateMust() *awsApprunnerConnectionState {
	if aac.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", aac.Type(), aac.LocalName()))
	}
	return aac.state
}

// Args contains the configurations for aws_apprunner_connection.
type Args struct {
	// ConnectionName: string, required
	ConnectionName terra.StringValue `hcl:"connection_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ProviderType: string, required
	ProviderType terra.StringValue `hcl:"provider_type,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
}

type awsApprunnerConnectionAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_apprunner_connection.
func (aac awsApprunnerConnectionAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(aac.ref.Append("arn"))
}

// ConnectionName returns a reference to field connection_name of aws_apprunner_connection.
func (aac awsApprunnerConnectionAttributes) ConnectionName() terra.StringValue {
	return terra.ReferenceAsString(aac.ref.Append("connection_name"))
}

// Id returns a reference to field id of aws_apprunner_connection.
func (aac awsApprunnerConnectionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(aac.ref.Append("id"))
}

// ProviderType returns a reference to field provider_type of aws_apprunner_connection.
func (aac awsApprunnerConnectionAttributes) ProviderType() terra.StringValue {
	return terra.ReferenceAsString(aac.ref.Append("provider_type"))
}

// Status returns a reference to field status of aws_apprunner_connection.
func (aac awsApprunnerConnectionAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(aac.ref.Append("status"))
}

// Tags returns a reference to field tags of aws_apprunner_connection.
func (aac awsApprunnerConnectionAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](aac.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_apprunner_connection.
func (aac awsApprunnerConnectionAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](aac.ref.Append("tags_all"))
}

type awsApprunnerConnectionState struct {
	Arn            string            `json:"arn"`
	ConnectionName string            `json:"connection_name"`
	Id             string            `json:"id"`
	ProviderType   string            `json:"provider_type"`
	Status         string            `json:"status"`
	Tags           map[string]string `json:"tags"`
	TagsAll        map[string]string `json:"tags_all"`
}
