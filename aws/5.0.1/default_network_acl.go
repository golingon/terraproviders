// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	defaultnetworkacl "github.com/golingon/terraproviders/aws/5.0.1/defaultnetworkacl"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDefaultNetworkAcl creates a new instance of [DefaultNetworkAcl].
func NewDefaultNetworkAcl(name string, args DefaultNetworkAclArgs) *DefaultNetworkAcl {
	return &DefaultNetworkAcl{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DefaultNetworkAcl)(nil)

// DefaultNetworkAcl represents the Terraform resource aws_default_network_acl.
type DefaultNetworkAcl struct {
	Name      string
	Args      DefaultNetworkAclArgs
	state     *defaultNetworkAclState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DefaultNetworkAcl].
func (dna *DefaultNetworkAcl) Type() string {
	return "aws_default_network_acl"
}

// LocalName returns the local name for [DefaultNetworkAcl].
func (dna *DefaultNetworkAcl) LocalName() string {
	return dna.Name
}

// Configuration returns the configuration (args) for [DefaultNetworkAcl].
func (dna *DefaultNetworkAcl) Configuration() interface{} {
	return dna.Args
}

// DependOn is used for other resources to depend on [DefaultNetworkAcl].
func (dna *DefaultNetworkAcl) DependOn() terra.Reference {
	return terra.ReferenceResource(dna)
}

// Dependencies returns the list of resources [DefaultNetworkAcl] depends_on.
func (dna *DefaultNetworkAcl) Dependencies() terra.Dependencies {
	return dna.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DefaultNetworkAcl].
func (dna *DefaultNetworkAcl) LifecycleManagement() *terra.Lifecycle {
	return dna.Lifecycle
}

// Attributes returns the attributes for [DefaultNetworkAcl].
func (dna *DefaultNetworkAcl) Attributes() defaultNetworkAclAttributes {
	return defaultNetworkAclAttributes{ref: terra.ReferenceResource(dna)}
}

// ImportState imports the given attribute values into [DefaultNetworkAcl]'s state.
func (dna *DefaultNetworkAcl) ImportState(av io.Reader) error {
	dna.state = &defaultNetworkAclState{}
	if err := json.NewDecoder(av).Decode(dna.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dna.Type(), dna.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DefaultNetworkAcl] has state.
func (dna *DefaultNetworkAcl) State() (*defaultNetworkAclState, bool) {
	return dna.state, dna.state != nil
}

// StateMust returns the state for [DefaultNetworkAcl]. Panics if the state is nil.
func (dna *DefaultNetworkAcl) StateMust() *defaultNetworkAclState {
	if dna.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dna.Type(), dna.LocalName()))
	}
	return dna.state
}

// DefaultNetworkAclArgs contains the configurations for aws_default_network_acl.
type DefaultNetworkAclArgs struct {
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
	Egress []defaultnetworkacl.Egress `hcl:"egress,block" validate:"min=0"`
	// Ingress: min=0
	Ingress []defaultnetworkacl.Ingress `hcl:"ingress,block" validate:"min=0"`
}
type defaultNetworkAclAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_default_network_acl.
func (dna defaultNetworkAclAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(dna.ref.Append("arn"))
}

// DefaultNetworkAclId returns a reference to field default_network_acl_id of aws_default_network_acl.
func (dna defaultNetworkAclAttributes) DefaultNetworkAclId() terra.StringValue {
	return terra.ReferenceAsString(dna.ref.Append("default_network_acl_id"))
}

// Id returns a reference to field id of aws_default_network_acl.
func (dna defaultNetworkAclAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dna.ref.Append("id"))
}

// OwnerId returns a reference to field owner_id of aws_default_network_acl.
func (dna defaultNetworkAclAttributes) OwnerId() terra.StringValue {
	return terra.ReferenceAsString(dna.ref.Append("owner_id"))
}

// SubnetIds returns a reference to field subnet_ids of aws_default_network_acl.
func (dna defaultNetworkAclAttributes) SubnetIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](dna.ref.Append("subnet_ids"))
}

// Tags returns a reference to field tags of aws_default_network_acl.
func (dna defaultNetworkAclAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dna.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_default_network_acl.
func (dna defaultNetworkAclAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dna.ref.Append("tags_all"))
}

// VpcId returns a reference to field vpc_id of aws_default_network_acl.
func (dna defaultNetworkAclAttributes) VpcId() terra.StringValue {
	return terra.ReferenceAsString(dna.ref.Append("vpc_id"))
}

func (dna defaultNetworkAclAttributes) Egress() terra.SetValue[defaultnetworkacl.EgressAttributes] {
	return terra.ReferenceAsSet[defaultnetworkacl.EgressAttributes](dna.ref.Append("egress"))
}

func (dna defaultNetworkAclAttributes) Ingress() terra.SetValue[defaultnetworkacl.IngressAttributes] {
	return terra.ReferenceAsSet[defaultnetworkacl.IngressAttributes](dna.ref.Append("ingress"))
}

type defaultNetworkAclState struct {
	Arn                 string                           `json:"arn"`
	DefaultNetworkAclId string                           `json:"default_network_acl_id"`
	Id                  string                           `json:"id"`
	OwnerId             string                           `json:"owner_id"`
	SubnetIds           []string                         `json:"subnet_ids"`
	Tags                map[string]string                `json:"tags"`
	TagsAll             map[string]string                `json:"tags_all"`
	VpcId               string                           `json:"vpc_id"`
	Egress              []defaultnetworkacl.EgressState  `json:"egress"`
	Ingress             []defaultnetworkacl.IngressState `json:"ingress"`
}
