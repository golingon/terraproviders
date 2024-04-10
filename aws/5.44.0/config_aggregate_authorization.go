// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// NewConfigAggregateAuthorization creates a new instance of [ConfigAggregateAuthorization].
func NewConfigAggregateAuthorization(name string, args ConfigAggregateAuthorizationArgs) *ConfigAggregateAuthorization {
	return &ConfigAggregateAuthorization{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ConfigAggregateAuthorization)(nil)

// ConfigAggregateAuthorization represents the Terraform resource aws_config_aggregate_authorization.
type ConfigAggregateAuthorization struct {
	Name      string
	Args      ConfigAggregateAuthorizationArgs
	state     *configAggregateAuthorizationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ConfigAggregateAuthorization].
func (caa *ConfigAggregateAuthorization) Type() string {
	return "aws_config_aggregate_authorization"
}

// LocalName returns the local name for [ConfigAggregateAuthorization].
func (caa *ConfigAggregateAuthorization) LocalName() string {
	return caa.Name
}

// Configuration returns the configuration (args) for [ConfigAggregateAuthorization].
func (caa *ConfigAggregateAuthorization) Configuration() interface{} {
	return caa.Args
}

// DependOn is used for other resources to depend on [ConfigAggregateAuthorization].
func (caa *ConfigAggregateAuthorization) DependOn() terra.Reference {
	return terra.ReferenceResource(caa)
}

// Dependencies returns the list of resources [ConfigAggregateAuthorization] depends_on.
func (caa *ConfigAggregateAuthorization) Dependencies() terra.Dependencies {
	return caa.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ConfigAggregateAuthorization].
func (caa *ConfigAggregateAuthorization) LifecycleManagement() *terra.Lifecycle {
	return caa.Lifecycle
}

// Attributes returns the attributes for [ConfigAggregateAuthorization].
func (caa *ConfigAggregateAuthorization) Attributes() configAggregateAuthorizationAttributes {
	return configAggregateAuthorizationAttributes{ref: terra.ReferenceResource(caa)}
}

// ImportState imports the given attribute values into [ConfigAggregateAuthorization]'s state.
func (caa *ConfigAggregateAuthorization) ImportState(av io.Reader) error {
	caa.state = &configAggregateAuthorizationState{}
	if err := json.NewDecoder(av).Decode(caa.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", caa.Type(), caa.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ConfigAggregateAuthorization] has state.
func (caa *ConfigAggregateAuthorization) State() (*configAggregateAuthorizationState, bool) {
	return caa.state, caa.state != nil
}

// StateMust returns the state for [ConfigAggregateAuthorization]. Panics if the state is nil.
func (caa *ConfigAggregateAuthorization) StateMust() *configAggregateAuthorizationState {
	if caa.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", caa.Type(), caa.LocalName()))
	}
	return caa.state
}

// ConfigAggregateAuthorizationArgs contains the configurations for aws_config_aggregate_authorization.
type ConfigAggregateAuthorizationArgs struct {
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
type configAggregateAuthorizationAttributes struct {
	ref terra.Reference
}

// AccountId returns a reference to field account_id of aws_config_aggregate_authorization.
func (caa configAggregateAuthorizationAttributes) AccountId() terra.StringValue {
	return terra.ReferenceAsString(caa.ref.Append("account_id"))
}

// Arn returns a reference to field arn of aws_config_aggregate_authorization.
func (caa configAggregateAuthorizationAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(caa.ref.Append("arn"))
}

// Id returns a reference to field id of aws_config_aggregate_authorization.
func (caa configAggregateAuthorizationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(caa.ref.Append("id"))
}

// Region returns a reference to field region of aws_config_aggregate_authorization.
func (caa configAggregateAuthorizationAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(caa.ref.Append("region"))
}

// Tags returns a reference to field tags of aws_config_aggregate_authorization.
func (caa configAggregateAuthorizationAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](caa.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_config_aggregate_authorization.
func (caa configAggregateAuthorizationAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](caa.ref.Append("tags_all"))
}

type configAggregateAuthorizationState struct {
	AccountId string            `json:"account_id"`
	Arn       string            `json:"arn"`
	Id        string            `json:"id"`
	Region    string            `json:"region"`
	Tags      map[string]string `json:"tags"`
	TagsAll   map[string]string `json:"tags_all"`
}
