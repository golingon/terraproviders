// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_msk_vpc_connection

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

// Resource represents the Terraform resource aws_msk_vpc_connection.
type Resource struct {
	Name      string
	Args      Args
	state     *awsMskVpcConnectionState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (amvc *Resource) Type() string {
	return "aws_msk_vpc_connection"
}

// LocalName returns the local name for [Resource].
func (amvc *Resource) LocalName() string {
	return amvc.Name
}

// Configuration returns the configuration (args) for [Resource].
func (amvc *Resource) Configuration() interface{} {
	return amvc.Args
}

// DependOn is used for other resources to depend on [Resource].
func (amvc *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(amvc)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (amvc *Resource) Dependencies() terra.Dependencies {
	return amvc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (amvc *Resource) LifecycleManagement() *terra.Lifecycle {
	return amvc.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (amvc *Resource) Attributes() awsMskVpcConnectionAttributes {
	return awsMskVpcConnectionAttributes{ref: terra.ReferenceResource(amvc)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (amvc *Resource) ImportState(state io.Reader) error {
	amvc.state = &awsMskVpcConnectionState{}
	if err := json.NewDecoder(state).Decode(amvc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", amvc.Type(), amvc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (amvc *Resource) State() (*awsMskVpcConnectionState, bool) {
	return amvc.state, amvc.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (amvc *Resource) StateMust() *awsMskVpcConnectionState {
	if amvc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", amvc.Type(), amvc.LocalName()))
	}
	return amvc.state
}

// Args contains the configurations for aws_msk_vpc_connection.
type Args struct {
	// Authentication: string, required
	Authentication terra.StringValue `hcl:"authentication,attr" validate:"required"`
	// ClientSubnets: set of string, required
	ClientSubnets terra.SetValue[terra.StringValue] `hcl:"client_subnets,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// SecurityGroups: set of string, required
	SecurityGroups terra.SetValue[terra.StringValue] `hcl:"security_groups,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// TargetClusterArn: string, required
	TargetClusterArn terra.StringValue `hcl:"target_cluster_arn,attr" validate:"required"`
	// VpcId: string, required
	VpcId terra.StringValue `hcl:"vpc_id,attr" validate:"required"`
}

type awsMskVpcConnectionAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_msk_vpc_connection.
func (amvc awsMskVpcConnectionAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(amvc.ref.Append("arn"))
}

// Authentication returns a reference to field authentication of aws_msk_vpc_connection.
func (amvc awsMskVpcConnectionAttributes) Authentication() terra.StringValue {
	return terra.ReferenceAsString(amvc.ref.Append("authentication"))
}

// ClientSubnets returns a reference to field client_subnets of aws_msk_vpc_connection.
func (amvc awsMskVpcConnectionAttributes) ClientSubnets() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](amvc.ref.Append("client_subnets"))
}

// Id returns a reference to field id of aws_msk_vpc_connection.
func (amvc awsMskVpcConnectionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(amvc.ref.Append("id"))
}

// SecurityGroups returns a reference to field security_groups of aws_msk_vpc_connection.
func (amvc awsMskVpcConnectionAttributes) SecurityGroups() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](amvc.ref.Append("security_groups"))
}

// Tags returns a reference to field tags of aws_msk_vpc_connection.
func (amvc awsMskVpcConnectionAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](amvc.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_msk_vpc_connection.
func (amvc awsMskVpcConnectionAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](amvc.ref.Append("tags_all"))
}

// TargetClusterArn returns a reference to field target_cluster_arn of aws_msk_vpc_connection.
func (amvc awsMskVpcConnectionAttributes) TargetClusterArn() terra.StringValue {
	return terra.ReferenceAsString(amvc.ref.Append("target_cluster_arn"))
}

// VpcId returns a reference to field vpc_id of aws_msk_vpc_connection.
func (amvc awsMskVpcConnectionAttributes) VpcId() terra.StringValue {
	return terra.ReferenceAsString(amvc.ref.Append("vpc_id"))
}

type awsMskVpcConnectionState struct {
	Arn              string            `json:"arn"`
	Authentication   string            `json:"authentication"`
	ClientSubnets    []string          `json:"client_subnets"`
	Id               string            `json:"id"`
	SecurityGroups   []string          `json:"security_groups"`
	Tags             map[string]string `json:"tags"`
	TagsAll          map[string]string `json:"tags_all"`
	TargetClusterArn string            `json:"target_cluster_arn"`
	VpcId            string            `json:"vpc_id"`
}
