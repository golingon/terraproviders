// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// NewMskVpcConnection creates a new instance of [MskVpcConnection].
func NewMskVpcConnection(name string, args MskVpcConnectionArgs) *MskVpcConnection {
	return &MskVpcConnection{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*MskVpcConnection)(nil)

// MskVpcConnection represents the Terraform resource aws_msk_vpc_connection.
type MskVpcConnection struct {
	Name      string
	Args      MskVpcConnectionArgs
	state     *mskVpcConnectionState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [MskVpcConnection].
func (mvc *MskVpcConnection) Type() string {
	return "aws_msk_vpc_connection"
}

// LocalName returns the local name for [MskVpcConnection].
func (mvc *MskVpcConnection) LocalName() string {
	return mvc.Name
}

// Configuration returns the configuration (args) for [MskVpcConnection].
func (mvc *MskVpcConnection) Configuration() interface{} {
	return mvc.Args
}

// DependOn is used for other resources to depend on [MskVpcConnection].
func (mvc *MskVpcConnection) DependOn() terra.Reference {
	return terra.ReferenceResource(mvc)
}

// Dependencies returns the list of resources [MskVpcConnection] depends_on.
func (mvc *MskVpcConnection) Dependencies() terra.Dependencies {
	return mvc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [MskVpcConnection].
func (mvc *MskVpcConnection) LifecycleManagement() *terra.Lifecycle {
	return mvc.Lifecycle
}

// Attributes returns the attributes for [MskVpcConnection].
func (mvc *MskVpcConnection) Attributes() mskVpcConnectionAttributes {
	return mskVpcConnectionAttributes{ref: terra.ReferenceResource(mvc)}
}

// ImportState imports the given attribute values into [MskVpcConnection]'s state.
func (mvc *MskVpcConnection) ImportState(av io.Reader) error {
	mvc.state = &mskVpcConnectionState{}
	if err := json.NewDecoder(av).Decode(mvc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", mvc.Type(), mvc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [MskVpcConnection] has state.
func (mvc *MskVpcConnection) State() (*mskVpcConnectionState, bool) {
	return mvc.state, mvc.state != nil
}

// StateMust returns the state for [MskVpcConnection]. Panics if the state is nil.
func (mvc *MskVpcConnection) StateMust() *mskVpcConnectionState {
	if mvc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", mvc.Type(), mvc.LocalName()))
	}
	return mvc.state
}

// MskVpcConnectionArgs contains the configurations for aws_msk_vpc_connection.
type MskVpcConnectionArgs struct {
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
type mskVpcConnectionAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_msk_vpc_connection.
func (mvc mskVpcConnectionAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(mvc.ref.Append("arn"))
}

// Authentication returns a reference to field authentication of aws_msk_vpc_connection.
func (mvc mskVpcConnectionAttributes) Authentication() terra.StringValue {
	return terra.ReferenceAsString(mvc.ref.Append("authentication"))
}

// ClientSubnets returns a reference to field client_subnets of aws_msk_vpc_connection.
func (mvc mskVpcConnectionAttributes) ClientSubnets() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](mvc.ref.Append("client_subnets"))
}

// Id returns a reference to field id of aws_msk_vpc_connection.
func (mvc mskVpcConnectionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(mvc.ref.Append("id"))
}

// SecurityGroups returns a reference to field security_groups of aws_msk_vpc_connection.
func (mvc mskVpcConnectionAttributes) SecurityGroups() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](mvc.ref.Append("security_groups"))
}

// Tags returns a reference to field tags of aws_msk_vpc_connection.
func (mvc mskVpcConnectionAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](mvc.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_msk_vpc_connection.
func (mvc mskVpcConnectionAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](mvc.ref.Append("tags_all"))
}

// TargetClusterArn returns a reference to field target_cluster_arn of aws_msk_vpc_connection.
func (mvc mskVpcConnectionAttributes) TargetClusterArn() terra.StringValue {
	return terra.ReferenceAsString(mvc.ref.Append("target_cluster_arn"))
}

// VpcId returns a reference to field vpc_id of aws_msk_vpc_connection.
func (mvc mskVpcConnectionAttributes) VpcId() terra.StringValue {
	return terra.ReferenceAsString(mvc.ref.Append("vpc_id"))
}

type mskVpcConnectionState struct {
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
