// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewApprunnerConnection creates a new instance of [ApprunnerConnection].
func NewApprunnerConnection(name string, args ApprunnerConnectionArgs) *ApprunnerConnection {
	return &ApprunnerConnection{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ApprunnerConnection)(nil)

// ApprunnerConnection represents the Terraform resource aws_apprunner_connection.
type ApprunnerConnection struct {
	Name      string
	Args      ApprunnerConnectionArgs
	state     *apprunnerConnectionState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ApprunnerConnection].
func (ac *ApprunnerConnection) Type() string {
	return "aws_apprunner_connection"
}

// LocalName returns the local name for [ApprunnerConnection].
func (ac *ApprunnerConnection) LocalName() string {
	return ac.Name
}

// Configuration returns the configuration (args) for [ApprunnerConnection].
func (ac *ApprunnerConnection) Configuration() interface{} {
	return ac.Args
}

// DependOn is used for other resources to depend on [ApprunnerConnection].
func (ac *ApprunnerConnection) DependOn() terra.Reference {
	return terra.ReferenceResource(ac)
}

// Dependencies returns the list of resources [ApprunnerConnection] depends_on.
func (ac *ApprunnerConnection) Dependencies() terra.Dependencies {
	return ac.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ApprunnerConnection].
func (ac *ApprunnerConnection) LifecycleManagement() *terra.Lifecycle {
	return ac.Lifecycle
}

// Attributes returns the attributes for [ApprunnerConnection].
func (ac *ApprunnerConnection) Attributes() apprunnerConnectionAttributes {
	return apprunnerConnectionAttributes{ref: terra.ReferenceResource(ac)}
}

// ImportState imports the given attribute values into [ApprunnerConnection]'s state.
func (ac *ApprunnerConnection) ImportState(av io.Reader) error {
	ac.state = &apprunnerConnectionState{}
	if err := json.NewDecoder(av).Decode(ac.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ac.Type(), ac.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ApprunnerConnection] has state.
func (ac *ApprunnerConnection) State() (*apprunnerConnectionState, bool) {
	return ac.state, ac.state != nil
}

// StateMust returns the state for [ApprunnerConnection]. Panics if the state is nil.
func (ac *ApprunnerConnection) StateMust() *apprunnerConnectionState {
	if ac.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ac.Type(), ac.LocalName()))
	}
	return ac.state
}

// ApprunnerConnectionArgs contains the configurations for aws_apprunner_connection.
type ApprunnerConnectionArgs struct {
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
type apprunnerConnectionAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_apprunner_connection.
func (ac apprunnerConnectionAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(ac.ref.Append("arn"))
}

// ConnectionName returns a reference to field connection_name of aws_apprunner_connection.
func (ac apprunnerConnectionAttributes) ConnectionName() terra.StringValue {
	return terra.ReferenceAsString(ac.ref.Append("connection_name"))
}

// Id returns a reference to field id of aws_apprunner_connection.
func (ac apprunnerConnectionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ac.ref.Append("id"))
}

// ProviderType returns a reference to field provider_type of aws_apprunner_connection.
func (ac apprunnerConnectionAttributes) ProviderType() terra.StringValue {
	return terra.ReferenceAsString(ac.ref.Append("provider_type"))
}

// Status returns a reference to field status of aws_apprunner_connection.
func (ac apprunnerConnectionAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(ac.ref.Append("status"))
}

// Tags returns a reference to field tags of aws_apprunner_connection.
func (ac apprunnerConnectionAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ac.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_apprunner_connection.
func (ac apprunnerConnectionAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ac.ref.Append("tags_all"))
}

type apprunnerConnectionState struct {
	Arn            string            `json:"arn"`
	ConnectionName string            `json:"connection_name"`
	Id             string            `json:"id"`
	ProviderType   string            `json:"provider_type"`
	Status         string            `json:"status"`
	Tags           map[string]string `json:"tags"`
	TagsAll        map[string]string `json:"tags_all"`
}
