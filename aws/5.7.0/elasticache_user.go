// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	elasticacheuser "github.com/golingon/terraproviders/aws/5.7.0/elasticacheuser"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewElasticacheUser creates a new instance of [ElasticacheUser].
func NewElasticacheUser(name string, args ElasticacheUserArgs) *ElasticacheUser {
	return &ElasticacheUser{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ElasticacheUser)(nil)

// ElasticacheUser represents the Terraform resource aws_elasticache_user.
type ElasticacheUser struct {
	Name      string
	Args      ElasticacheUserArgs
	state     *elasticacheUserState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ElasticacheUser].
func (eu *ElasticacheUser) Type() string {
	return "aws_elasticache_user"
}

// LocalName returns the local name for [ElasticacheUser].
func (eu *ElasticacheUser) LocalName() string {
	return eu.Name
}

// Configuration returns the configuration (args) for [ElasticacheUser].
func (eu *ElasticacheUser) Configuration() interface{} {
	return eu.Args
}

// DependOn is used for other resources to depend on [ElasticacheUser].
func (eu *ElasticacheUser) DependOn() terra.Reference {
	return terra.ReferenceResource(eu)
}

// Dependencies returns the list of resources [ElasticacheUser] depends_on.
func (eu *ElasticacheUser) Dependencies() terra.Dependencies {
	return eu.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ElasticacheUser].
func (eu *ElasticacheUser) LifecycleManagement() *terra.Lifecycle {
	return eu.Lifecycle
}

// Attributes returns the attributes for [ElasticacheUser].
func (eu *ElasticacheUser) Attributes() elasticacheUserAttributes {
	return elasticacheUserAttributes{ref: terra.ReferenceResource(eu)}
}

// ImportState imports the given attribute values into [ElasticacheUser]'s state.
func (eu *ElasticacheUser) ImportState(av io.Reader) error {
	eu.state = &elasticacheUserState{}
	if err := json.NewDecoder(av).Decode(eu.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", eu.Type(), eu.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ElasticacheUser] has state.
func (eu *ElasticacheUser) State() (*elasticacheUserState, bool) {
	return eu.state, eu.state != nil
}

// StateMust returns the state for [ElasticacheUser]. Panics if the state is nil.
func (eu *ElasticacheUser) StateMust() *elasticacheUserState {
	if eu.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", eu.Type(), eu.LocalName()))
	}
	return eu.state
}

// ElasticacheUserArgs contains the configurations for aws_elasticache_user.
type ElasticacheUserArgs struct {
	// AccessString: string, required
	AccessString terra.StringValue `hcl:"access_string,attr" validate:"required"`
	// Engine: string, required
	Engine terra.StringValue `hcl:"engine,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// NoPasswordRequired: bool, optional
	NoPasswordRequired terra.BoolValue `hcl:"no_password_required,attr"`
	// Passwords: set of string, optional
	Passwords terra.SetValue[terra.StringValue] `hcl:"passwords,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// UserId: string, required
	UserId terra.StringValue `hcl:"user_id,attr" validate:"required"`
	// UserName: string, required
	UserName terra.StringValue `hcl:"user_name,attr" validate:"required"`
	// AuthenticationMode: optional
	AuthenticationMode *elasticacheuser.AuthenticationMode `hcl:"authentication_mode,block"`
	// Timeouts: optional
	Timeouts *elasticacheuser.Timeouts `hcl:"timeouts,block"`
}
type elasticacheUserAttributes struct {
	ref terra.Reference
}

// AccessString returns a reference to field access_string of aws_elasticache_user.
func (eu elasticacheUserAttributes) AccessString() terra.StringValue {
	return terra.ReferenceAsString(eu.ref.Append("access_string"))
}

// Arn returns a reference to field arn of aws_elasticache_user.
func (eu elasticacheUserAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(eu.ref.Append("arn"))
}

// Engine returns a reference to field engine of aws_elasticache_user.
func (eu elasticacheUserAttributes) Engine() terra.StringValue {
	return terra.ReferenceAsString(eu.ref.Append("engine"))
}

// Id returns a reference to field id of aws_elasticache_user.
func (eu elasticacheUserAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(eu.ref.Append("id"))
}

// NoPasswordRequired returns a reference to field no_password_required of aws_elasticache_user.
func (eu elasticacheUserAttributes) NoPasswordRequired() terra.BoolValue {
	return terra.ReferenceAsBool(eu.ref.Append("no_password_required"))
}

// Passwords returns a reference to field passwords of aws_elasticache_user.
func (eu elasticacheUserAttributes) Passwords() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](eu.ref.Append("passwords"))
}

// Tags returns a reference to field tags of aws_elasticache_user.
func (eu elasticacheUserAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](eu.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_elasticache_user.
func (eu elasticacheUserAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](eu.ref.Append("tags_all"))
}

// UserId returns a reference to field user_id of aws_elasticache_user.
func (eu elasticacheUserAttributes) UserId() terra.StringValue {
	return terra.ReferenceAsString(eu.ref.Append("user_id"))
}

// UserName returns a reference to field user_name of aws_elasticache_user.
func (eu elasticacheUserAttributes) UserName() terra.StringValue {
	return terra.ReferenceAsString(eu.ref.Append("user_name"))
}

func (eu elasticacheUserAttributes) AuthenticationMode() terra.ListValue[elasticacheuser.AuthenticationModeAttributes] {
	return terra.ReferenceAsList[elasticacheuser.AuthenticationModeAttributes](eu.ref.Append("authentication_mode"))
}

func (eu elasticacheUserAttributes) Timeouts() elasticacheuser.TimeoutsAttributes {
	return terra.ReferenceAsSingle[elasticacheuser.TimeoutsAttributes](eu.ref.Append("timeouts"))
}

type elasticacheUserState struct {
	AccessString       string                                    `json:"access_string"`
	Arn                string                                    `json:"arn"`
	Engine             string                                    `json:"engine"`
	Id                 string                                    `json:"id"`
	NoPasswordRequired bool                                      `json:"no_password_required"`
	Passwords          []string                                  `json:"passwords"`
	Tags               map[string]string                         `json:"tags"`
	TagsAll            map[string]string                         `json:"tags_all"`
	UserId             string                                    `json:"user_id"`
	UserName           string                                    `json:"user_name"`
	AuthenticationMode []elasticacheuser.AuthenticationModeState `json:"authentication_mode"`
	Timeouts           *elasticacheuser.TimeoutsState            `json:"timeouts"`
}
