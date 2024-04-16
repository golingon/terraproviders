// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_config_aggregate_authorization

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

// Resource represents the Terraform resource aws_config_aggregate_authorization.
type Resource struct {
	Name      string
	Args      Args
	state     *awsConfigAggregateAuthorizationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (acaa *Resource) Type() string {
	return "aws_config_aggregate_authorization"
}

// LocalName returns the local name for [Resource].
func (acaa *Resource) LocalName() string {
	return acaa.Name
}

// Configuration returns the configuration (args) for [Resource].
func (acaa *Resource) Configuration() interface{} {
	return acaa.Args
}

// DependOn is used for other resources to depend on [Resource].
func (acaa *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(acaa)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (acaa *Resource) Dependencies() terra.Dependencies {
	return acaa.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (acaa *Resource) LifecycleManagement() *terra.Lifecycle {
	return acaa.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (acaa *Resource) Attributes() awsConfigAggregateAuthorizationAttributes {
	return awsConfigAggregateAuthorizationAttributes{ref: terra.ReferenceResource(acaa)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (acaa *Resource) ImportState(state io.Reader) error {
	acaa.state = &awsConfigAggregateAuthorizationState{}
	if err := json.NewDecoder(state).Decode(acaa.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", acaa.Type(), acaa.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (acaa *Resource) State() (*awsConfigAggregateAuthorizationState, bool) {
	return acaa.state, acaa.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (acaa *Resource) StateMust() *awsConfigAggregateAuthorizationState {
	if acaa.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", acaa.Type(), acaa.LocalName()))
	}
	return acaa.state
}

// Args contains the configurations for aws_config_aggregate_authorization.
type Args struct {
	// AccountId: string, required
	AccountId terra.StringValue `hcl:"account_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Region: string, required
	Region terra.StringValue `hcl:"region,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
}

type awsConfigAggregateAuthorizationAttributes struct {
	ref terra.Reference
}

// AccountId returns a reference to field account_id of aws_config_aggregate_authorization.
func (acaa awsConfigAggregateAuthorizationAttributes) AccountId() terra.StringValue {
	return terra.ReferenceAsString(acaa.ref.Append("account_id"))
}

// Arn returns a reference to field arn of aws_config_aggregate_authorization.
func (acaa awsConfigAggregateAuthorizationAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(acaa.ref.Append("arn"))
}

// Id returns a reference to field id of aws_config_aggregate_authorization.
func (acaa awsConfigAggregateAuthorizationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(acaa.ref.Append("id"))
}

// Region returns a reference to field region of aws_config_aggregate_authorization.
func (acaa awsConfigAggregateAuthorizationAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(acaa.ref.Append("region"))
}

// Tags returns a reference to field tags of aws_config_aggregate_authorization.
func (acaa awsConfigAggregateAuthorizationAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](acaa.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_config_aggregate_authorization.
func (acaa awsConfigAggregateAuthorizationAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](acaa.ref.Append("tags_all"))
}

type awsConfigAggregateAuthorizationState struct {
	AccountId string            `json:"account_id"`
	Arn       string            `json:"arn"`
	Id        string            `json:"id"`
	Region    string            `json:"region"`
	Tags      map[string]string `json:"tags"`
	TagsAll   map[string]string `json:"tags_all"`
}
