// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_route53_resolver_query_log_config

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

// Resource represents the Terraform resource aws_route53_resolver_query_log_config.
type Resource struct {
	Name      string
	Args      Args
	state     *awsRoute53ResolverQueryLogConfigState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (arrqlc *Resource) Type() string {
	return "aws_route53_resolver_query_log_config"
}

// LocalName returns the local name for [Resource].
func (arrqlc *Resource) LocalName() string {
	return arrqlc.Name
}

// Configuration returns the configuration (args) for [Resource].
func (arrqlc *Resource) Configuration() interface{} {
	return arrqlc.Args
}

// DependOn is used for other resources to depend on [Resource].
func (arrqlc *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(arrqlc)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (arrqlc *Resource) Dependencies() terra.Dependencies {
	return arrqlc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (arrqlc *Resource) LifecycleManagement() *terra.Lifecycle {
	return arrqlc.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (arrqlc *Resource) Attributes() awsRoute53ResolverQueryLogConfigAttributes {
	return awsRoute53ResolverQueryLogConfigAttributes{ref: terra.ReferenceResource(arrqlc)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (arrqlc *Resource) ImportState(state io.Reader) error {
	arrqlc.state = &awsRoute53ResolverQueryLogConfigState{}
	if err := json.NewDecoder(state).Decode(arrqlc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", arrqlc.Type(), arrqlc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (arrqlc *Resource) State() (*awsRoute53ResolverQueryLogConfigState, bool) {
	return arrqlc.state, arrqlc.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (arrqlc *Resource) StateMust() *awsRoute53ResolverQueryLogConfigState {
	if arrqlc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", arrqlc.Type(), arrqlc.LocalName()))
	}
	return arrqlc.state
}

// Args contains the configurations for aws_route53_resolver_query_log_config.
type Args struct {
	// DestinationArn: string, required
	DestinationArn terra.StringValue `hcl:"destination_arn,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
}

type awsRoute53ResolverQueryLogConfigAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_route53_resolver_query_log_config.
func (arrqlc awsRoute53ResolverQueryLogConfigAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(arrqlc.ref.Append("arn"))
}

// DestinationArn returns a reference to field destination_arn of aws_route53_resolver_query_log_config.
func (arrqlc awsRoute53ResolverQueryLogConfigAttributes) DestinationArn() terra.StringValue {
	return terra.ReferenceAsString(arrqlc.ref.Append("destination_arn"))
}

// Id returns a reference to field id of aws_route53_resolver_query_log_config.
func (arrqlc awsRoute53ResolverQueryLogConfigAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(arrqlc.ref.Append("id"))
}

// Name returns a reference to field name of aws_route53_resolver_query_log_config.
func (arrqlc awsRoute53ResolverQueryLogConfigAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(arrqlc.ref.Append("name"))
}

// OwnerId returns a reference to field owner_id of aws_route53_resolver_query_log_config.
func (arrqlc awsRoute53ResolverQueryLogConfigAttributes) OwnerId() terra.StringValue {
	return terra.ReferenceAsString(arrqlc.ref.Append("owner_id"))
}

// ShareStatus returns a reference to field share_status of aws_route53_resolver_query_log_config.
func (arrqlc awsRoute53ResolverQueryLogConfigAttributes) ShareStatus() terra.StringValue {
	return terra.ReferenceAsString(arrqlc.ref.Append("share_status"))
}

// Tags returns a reference to field tags of aws_route53_resolver_query_log_config.
func (arrqlc awsRoute53ResolverQueryLogConfigAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](arrqlc.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_route53_resolver_query_log_config.
func (arrqlc awsRoute53ResolverQueryLogConfigAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](arrqlc.ref.Append("tags_all"))
}

type awsRoute53ResolverQueryLogConfigState struct {
	Arn            string            `json:"arn"`
	DestinationArn string            `json:"destination_arn"`
	Id             string            `json:"id"`
	Name           string            `json:"name"`
	OwnerId        string            `json:"owner_id"`
	ShareStatus    string            `json:"share_status"`
	Tags           map[string]string `json:"tags"`
	TagsAll        map[string]string `json:"tags_all"`
}
