// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	securitygroup "github.com/golingon/terraproviders/aws/4.63.0/securitygroup"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSecurityGroup creates a new instance of [SecurityGroup].
func NewSecurityGroup(name string, args SecurityGroupArgs) *SecurityGroup {
	return &SecurityGroup{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SecurityGroup)(nil)

// SecurityGroup represents the Terraform resource aws_security_group.
type SecurityGroup struct {
	Name      string
	Args      SecurityGroupArgs
	state     *securityGroupState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SecurityGroup].
func (sg *SecurityGroup) Type() string {
	return "aws_security_group"
}

// LocalName returns the local name for [SecurityGroup].
func (sg *SecurityGroup) LocalName() string {
	return sg.Name
}

// Configuration returns the configuration (args) for [SecurityGroup].
func (sg *SecurityGroup) Configuration() interface{} {
	return sg.Args
}

// DependOn is used for other resources to depend on [SecurityGroup].
func (sg *SecurityGroup) DependOn() terra.Reference {
	return terra.ReferenceResource(sg)
}

// Dependencies returns the list of resources [SecurityGroup] depends_on.
func (sg *SecurityGroup) Dependencies() terra.Dependencies {
	return sg.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SecurityGroup].
func (sg *SecurityGroup) LifecycleManagement() *terra.Lifecycle {
	return sg.Lifecycle
}

// Attributes returns the attributes for [SecurityGroup].
func (sg *SecurityGroup) Attributes() securityGroupAttributes {
	return securityGroupAttributes{ref: terra.ReferenceResource(sg)}
}

// ImportState imports the given attribute values into [SecurityGroup]'s state.
func (sg *SecurityGroup) ImportState(av io.Reader) error {
	sg.state = &securityGroupState{}
	if err := json.NewDecoder(av).Decode(sg.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sg.Type(), sg.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SecurityGroup] has state.
func (sg *SecurityGroup) State() (*securityGroupState, bool) {
	return sg.state, sg.state != nil
}

// StateMust returns the state for [SecurityGroup]. Panics if the state is nil.
func (sg *SecurityGroup) StateMust() *securityGroupState {
	if sg.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sg.Type(), sg.LocalName()))
	}
	return sg.state
}

// SecurityGroupArgs contains the configurations for aws_security_group.
type SecurityGroupArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
	// NamePrefix: string, optional
	NamePrefix terra.StringValue `hcl:"name_prefix,attr"`
	// RevokeRulesOnDelete: bool, optional
	RevokeRulesOnDelete terra.BoolValue `hcl:"revoke_rules_on_delete,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// VpcId: string, optional
	VpcId terra.StringValue `hcl:"vpc_id,attr"`
	// Egress: min=0
	Egress []securitygroup.Egress `hcl:"egress,block" validate:"min=0"`
	// Ingress: min=0
	Ingress []securitygroup.Ingress `hcl:"ingress,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *securitygroup.Timeouts `hcl:"timeouts,block"`
}
type securityGroupAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_security_group.
func (sg securityGroupAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(sg.ref.Append("arn"))
}

// Description returns a reference to field description of aws_security_group.
func (sg securityGroupAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(sg.ref.Append("description"))
}

// Id returns a reference to field id of aws_security_group.
func (sg securityGroupAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sg.ref.Append("id"))
}

// Name returns a reference to field name of aws_security_group.
func (sg securityGroupAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(sg.ref.Append("name"))
}

// NamePrefix returns a reference to field name_prefix of aws_security_group.
func (sg securityGroupAttributes) NamePrefix() terra.StringValue {
	return terra.ReferenceAsString(sg.ref.Append("name_prefix"))
}

// OwnerId returns a reference to field owner_id of aws_security_group.
func (sg securityGroupAttributes) OwnerId() terra.StringValue {
	return terra.ReferenceAsString(sg.ref.Append("owner_id"))
}

// RevokeRulesOnDelete returns a reference to field revoke_rules_on_delete of aws_security_group.
func (sg securityGroupAttributes) RevokeRulesOnDelete() terra.BoolValue {
	return terra.ReferenceAsBool(sg.ref.Append("revoke_rules_on_delete"))
}

// Tags returns a reference to field tags of aws_security_group.
func (sg securityGroupAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](sg.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_security_group.
func (sg securityGroupAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](sg.ref.Append("tags_all"))
}

// VpcId returns a reference to field vpc_id of aws_security_group.
func (sg securityGroupAttributes) VpcId() terra.StringValue {
	return terra.ReferenceAsString(sg.ref.Append("vpc_id"))
}

func (sg securityGroupAttributes) Egress() terra.SetValue[securitygroup.EgressAttributes] {
	return terra.ReferenceAsSet[securitygroup.EgressAttributes](sg.ref.Append("egress"))
}

func (sg securityGroupAttributes) Ingress() terra.SetValue[securitygroup.IngressAttributes] {
	return terra.ReferenceAsSet[securitygroup.IngressAttributes](sg.ref.Append("ingress"))
}

func (sg securityGroupAttributes) Timeouts() securitygroup.TimeoutsAttributes {
	return terra.ReferenceAsSingle[securitygroup.TimeoutsAttributes](sg.ref.Append("timeouts"))
}

type securityGroupState struct {
	Arn                 string                       `json:"arn"`
	Description         string                       `json:"description"`
	Id                  string                       `json:"id"`
	Name                string                       `json:"name"`
	NamePrefix          string                       `json:"name_prefix"`
	OwnerId             string                       `json:"owner_id"`
	RevokeRulesOnDelete bool                         `json:"revoke_rules_on_delete"`
	Tags                map[string]string            `json:"tags"`
	TagsAll             map[string]string            `json:"tags_all"`
	VpcId               string                       `json:"vpc_id"`
	Egress              []securitygroup.EgressState  `json:"egress"`
	Ingress             []securitygroup.IngressState `json:"ingress"`
	Timeouts            *securitygroup.TimeoutsState `json:"timeouts"`
}
