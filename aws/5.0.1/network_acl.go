// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	networkacl "github.com/golingon/terraproviders/aws/5.0.1/networkacl"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewNetworkAcl creates a new instance of [NetworkAcl].
func NewNetworkAcl(name string, args NetworkAclArgs) *NetworkAcl {
	return &NetworkAcl{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*NetworkAcl)(nil)

// NetworkAcl represents the Terraform resource aws_network_acl.
type NetworkAcl struct {
	Name      string
	Args      NetworkAclArgs
	state     *networkAclState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [NetworkAcl].
func (na *NetworkAcl) Type() string {
	return "aws_network_acl"
}

// LocalName returns the local name for [NetworkAcl].
func (na *NetworkAcl) LocalName() string {
	return na.Name
}

// Configuration returns the configuration (args) for [NetworkAcl].
func (na *NetworkAcl) Configuration() interface{} {
	return na.Args
}

// DependOn is used for other resources to depend on [NetworkAcl].
func (na *NetworkAcl) DependOn() terra.Reference {
	return terra.ReferenceResource(na)
}

// Dependencies returns the list of resources [NetworkAcl] depends_on.
func (na *NetworkAcl) Dependencies() terra.Dependencies {
	return na.DependsOn
}

// LifecycleManagement returns the lifecycle block for [NetworkAcl].
func (na *NetworkAcl) LifecycleManagement() *terra.Lifecycle {
	return na.Lifecycle
}

// Attributes returns the attributes for [NetworkAcl].
func (na *NetworkAcl) Attributes() networkAclAttributes {
	return networkAclAttributes{ref: terra.ReferenceResource(na)}
}

// ImportState imports the given attribute values into [NetworkAcl]'s state.
func (na *NetworkAcl) ImportState(av io.Reader) error {
	na.state = &networkAclState{}
	if err := json.NewDecoder(av).Decode(na.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", na.Type(), na.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [NetworkAcl] has state.
func (na *NetworkAcl) State() (*networkAclState, bool) {
	return na.state, na.state != nil
}

// StateMust returns the state for [NetworkAcl]. Panics if the state is nil.
func (na *NetworkAcl) StateMust() *networkAclState {
	if na.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", na.Type(), na.LocalName()))
	}
	return na.state
}

// NetworkAclArgs contains the configurations for aws_network_acl.
type NetworkAclArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// SubnetIds: set of string, optional
	SubnetIds terra.SetValue[terra.StringValue] `hcl:"subnet_ids,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// VpcId: string, required
	VpcId terra.StringValue `hcl:"vpc_id,attr" validate:"required"`
	// Egress: min=0
	Egress []networkacl.Egress `hcl:"egress,block" validate:"min=0"`
	// Ingress: min=0
	Ingress []networkacl.Ingress `hcl:"ingress,block" validate:"min=0"`
}
type networkAclAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_network_acl.
func (na networkAclAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(na.ref.Append("arn"))
}

// Id returns a reference to field id of aws_network_acl.
func (na networkAclAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(na.ref.Append("id"))
}

// OwnerId returns a reference to field owner_id of aws_network_acl.
func (na networkAclAttributes) OwnerId() terra.StringValue {
	return terra.ReferenceAsString(na.ref.Append("owner_id"))
}

// SubnetIds returns a reference to field subnet_ids of aws_network_acl.
func (na networkAclAttributes) SubnetIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](na.ref.Append("subnet_ids"))
}

// Tags returns a reference to field tags of aws_network_acl.
func (na networkAclAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](na.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_network_acl.
func (na networkAclAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](na.ref.Append("tags_all"))
}

// VpcId returns a reference to field vpc_id of aws_network_acl.
func (na networkAclAttributes) VpcId() terra.StringValue {
	return terra.ReferenceAsString(na.ref.Append("vpc_id"))
}

func (na networkAclAttributes) Egress() terra.SetValue[networkacl.EgressAttributes] {
	return terra.ReferenceAsSet[networkacl.EgressAttributes](na.ref.Append("egress"))
}

func (na networkAclAttributes) Ingress() terra.SetValue[networkacl.IngressAttributes] {
	return terra.ReferenceAsSet[networkacl.IngressAttributes](na.ref.Append("ingress"))
}

type networkAclState struct {
	Arn       string                    `json:"arn"`
	Id        string                    `json:"id"`
	OwnerId   string                    `json:"owner_id"`
	SubnetIds []string                  `json:"subnet_ids"`
	Tags      map[string]string         `json:"tags"`
	TagsAll   map[string]string         `json:"tags_all"`
	VpcId     string                    `json:"vpc_id"`
	Egress    []networkacl.EgressState  `json:"egress"`
	Ingress   []networkacl.IngressState `json:"ingress"`
}
