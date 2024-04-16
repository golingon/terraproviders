// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_default_network_acl

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

// Resource represents the Terraform resource aws_default_network_acl.
type Resource struct {
	Name      string
	Args      Args
	state     *awsDefaultNetworkAclState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (adna *Resource) Type() string {
	return "aws_default_network_acl"
}

// LocalName returns the local name for [Resource].
func (adna *Resource) LocalName() string {
	return adna.Name
}

// Configuration returns the configuration (args) for [Resource].
func (adna *Resource) Configuration() interface{} {
	return adna.Args
}

// DependOn is used for other resources to depend on [Resource].
func (adna *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(adna)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (adna *Resource) Dependencies() terra.Dependencies {
	return adna.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (adna *Resource) LifecycleManagement() *terra.Lifecycle {
	return adna.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (adna *Resource) Attributes() awsDefaultNetworkAclAttributes {
	return awsDefaultNetworkAclAttributes{ref: terra.ReferenceResource(adna)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (adna *Resource) ImportState(state io.Reader) error {
	adna.state = &awsDefaultNetworkAclState{}
	if err := json.NewDecoder(state).Decode(adna.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", adna.Type(), adna.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (adna *Resource) State() (*awsDefaultNetworkAclState, bool) {
	return adna.state, adna.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (adna *Resource) StateMust() *awsDefaultNetworkAclState {
	if adna.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", adna.Type(), adna.LocalName()))
	}
	return adna.state
}

// Args contains the configurations for aws_default_network_acl.
type Args struct {
	// DefaultNetworkAclId: string, required
	DefaultNetworkAclId terra.StringValue `hcl:"default_network_acl_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// SubnetIds: set of string, optional
	SubnetIds terra.SetValue[terra.StringValue] `hcl:"subnet_ids,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// Egress: min=0
	Egress []Egress `hcl:"egress,block" validate:"min=0"`
	// Ingress: min=0
	Ingress []Ingress `hcl:"ingress,block" validate:"min=0"`
}

type awsDefaultNetworkAclAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_default_network_acl.
func (adna awsDefaultNetworkAclAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(adna.ref.Append("arn"))
}

// DefaultNetworkAclId returns a reference to field default_network_acl_id of aws_default_network_acl.
func (adna awsDefaultNetworkAclAttributes) DefaultNetworkAclId() terra.StringValue {
	return terra.ReferenceAsString(adna.ref.Append("default_network_acl_id"))
}

// Id returns a reference to field id of aws_default_network_acl.
func (adna awsDefaultNetworkAclAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(adna.ref.Append("id"))
}

// OwnerId returns a reference to field owner_id of aws_default_network_acl.
func (adna awsDefaultNetworkAclAttributes) OwnerId() terra.StringValue {
	return terra.ReferenceAsString(adna.ref.Append("owner_id"))
}

// SubnetIds returns a reference to field subnet_ids of aws_default_network_acl.
func (adna awsDefaultNetworkAclAttributes) SubnetIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](adna.ref.Append("subnet_ids"))
}

// Tags returns a reference to field tags of aws_default_network_acl.
func (adna awsDefaultNetworkAclAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](adna.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_default_network_acl.
func (adna awsDefaultNetworkAclAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](adna.ref.Append("tags_all"))
}

// VpcId returns a reference to field vpc_id of aws_default_network_acl.
func (adna awsDefaultNetworkAclAttributes) VpcId() terra.StringValue {
	return terra.ReferenceAsString(adna.ref.Append("vpc_id"))
}

func (adna awsDefaultNetworkAclAttributes) Egress() terra.SetValue[EgressAttributes] {
	return terra.ReferenceAsSet[EgressAttributes](adna.ref.Append("egress"))
}

func (adna awsDefaultNetworkAclAttributes) Ingress() terra.SetValue[IngressAttributes] {
	return terra.ReferenceAsSet[IngressAttributes](adna.ref.Append("ingress"))
}

type awsDefaultNetworkAclState struct {
	Arn                 string            `json:"arn"`
	DefaultNetworkAclId string            `json:"default_network_acl_id"`
	Id                  string            `json:"id"`
	OwnerId             string            `json:"owner_id"`
	SubnetIds           []string          `json:"subnet_ids"`
	Tags                map[string]string `json:"tags"`
	TagsAll             map[string]string `json:"tags_all"`
	VpcId               string            `json:"vpc_id"`
	Egress              []EgressState     `json:"egress"`
	Ingress             []IngressState    `json:"ingress"`
}
