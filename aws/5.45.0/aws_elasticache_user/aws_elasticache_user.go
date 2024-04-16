// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_elasticache_user

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

// Resource represents the Terraform resource aws_elasticache_user.
type Resource struct {
	Name      string
	Args      Args
	state     *awsElasticacheUserState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (aeu *Resource) Type() string {
	return "aws_elasticache_user"
}

// LocalName returns the local name for [Resource].
func (aeu *Resource) LocalName() string {
	return aeu.Name
}

// Configuration returns the configuration (args) for [Resource].
func (aeu *Resource) Configuration() interface{} {
	return aeu.Args
}

// DependOn is used for other resources to depend on [Resource].
func (aeu *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(aeu)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (aeu *Resource) Dependencies() terra.Dependencies {
	return aeu.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (aeu *Resource) LifecycleManagement() *terra.Lifecycle {
	return aeu.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (aeu *Resource) Attributes() awsElasticacheUserAttributes {
	return awsElasticacheUserAttributes{ref: terra.ReferenceResource(aeu)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (aeu *Resource) ImportState(state io.Reader) error {
	aeu.state = &awsElasticacheUserState{}
	if err := json.NewDecoder(state).Decode(aeu.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", aeu.Type(), aeu.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (aeu *Resource) State() (*awsElasticacheUserState, bool) {
	return aeu.state, aeu.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (aeu *Resource) StateMust() *awsElasticacheUserState {
	if aeu.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", aeu.Type(), aeu.LocalName()))
	}
	return aeu.state
}

// Args contains the configurations for aws_elasticache_user.
type Args struct {
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
	AuthenticationMode *AuthenticationMode `hcl:"authentication_mode,block"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type awsElasticacheUserAttributes struct {
	ref terra.Reference
}

// AccessString returns a reference to field access_string of aws_elasticache_user.
func (aeu awsElasticacheUserAttributes) AccessString() terra.StringValue {
	return terra.ReferenceAsString(aeu.ref.Append("access_string"))
}

// Arn returns a reference to field arn of aws_elasticache_user.
func (aeu awsElasticacheUserAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(aeu.ref.Append("arn"))
}

// Engine returns a reference to field engine of aws_elasticache_user.
func (aeu awsElasticacheUserAttributes) Engine() terra.StringValue {
	return terra.ReferenceAsString(aeu.ref.Append("engine"))
}

// Id returns a reference to field id of aws_elasticache_user.
func (aeu awsElasticacheUserAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(aeu.ref.Append("id"))
}

// NoPasswordRequired returns a reference to field no_password_required of aws_elasticache_user.
func (aeu awsElasticacheUserAttributes) NoPasswordRequired() terra.BoolValue {
	return terra.ReferenceAsBool(aeu.ref.Append("no_password_required"))
}

// Passwords returns a reference to field passwords of aws_elasticache_user.
func (aeu awsElasticacheUserAttributes) Passwords() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](aeu.ref.Append("passwords"))
}

// Tags returns a reference to field tags of aws_elasticache_user.
func (aeu awsElasticacheUserAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](aeu.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_elasticache_user.
func (aeu awsElasticacheUserAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](aeu.ref.Append("tags_all"))
}

// UserId returns a reference to field user_id of aws_elasticache_user.
func (aeu awsElasticacheUserAttributes) UserId() terra.StringValue {
	return terra.ReferenceAsString(aeu.ref.Append("user_id"))
}

// UserName returns a reference to field user_name of aws_elasticache_user.
func (aeu awsElasticacheUserAttributes) UserName() terra.StringValue {
	return terra.ReferenceAsString(aeu.ref.Append("user_name"))
}

func (aeu awsElasticacheUserAttributes) AuthenticationMode() terra.ListValue[AuthenticationModeAttributes] {
	return terra.ReferenceAsList[AuthenticationModeAttributes](aeu.ref.Append("authentication_mode"))
}

func (aeu awsElasticacheUserAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](aeu.ref.Append("timeouts"))
}

type awsElasticacheUserState struct {
	AccessString       string                    `json:"access_string"`
	Arn                string                    `json:"arn"`
	Engine             string                    `json:"engine"`
	Id                 string                    `json:"id"`
	NoPasswordRequired bool                      `json:"no_password_required"`
	Passwords          []string                  `json:"passwords"`
	Tags               map[string]string         `json:"tags"`
	TagsAll            map[string]string         `json:"tags_all"`
	UserId             string                    `json:"user_id"`
	UserName           string                    `json:"user_name"`
	AuthenticationMode []AuthenticationModeState `json:"authentication_mode"`
	Timeouts           *TimeoutsState            `json:"timeouts"`
}
