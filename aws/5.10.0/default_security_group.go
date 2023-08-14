// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	defaultsecuritygroup "github.com/golingon/terraproviders/aws/5.10.0/defaultsecuritygroup"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDefaultSecurityGroup creates a new instance of [DefaultSecurityGroup].
func NewDefaultSecurityGroup(name string, args DefaultSecurityGroupArgs) *DefaultSecurityGroup {
	return &DefaultSecurityGroup{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DefaultSecurityGroup)(nil)

// DefaultSecurityGroup represents the Terraform resource aws_default_security_group.
type DefaultSecurityGroup struct {
	Name      string
	Args      DefaultSecurityGroupArgs
	state     *defaultSecurityGroupState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DefaultSecurityGroup].
func (dsg *DefaultSecurityGroup) Type() string {
	return "aws_default_security_group"
}

// LocalName returns the local name for [DefaultSecurityGroup].
func (dsg *DefaultSecurityGroup) LocalName() string {
	return dsg.Name
}

// Configuration returns the configuration (args) for [DefaultSecurityGroup].
func (dsg *DefaultSecurityGroup) Configuration() interface{} {
	return dsg.Args
}

// DependOn is used for other resources to depend on [DefaultSecurityGroup].
func (dsg *DefaultSecurityGroup) DependOn() terra.Reference {
	return terra.ReferenceResource(dsg)
}

// Dependencies returns the list of resources [DefaultSecurityGroup] depends_on.
func (dsg *DefaultSecurityGroup) Dependencies() terra.Dependencies {
	return dsg.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DefaultSecurityGroup].
func (dsg *DefaultSecurityGroup) LifecycleManagement() *terra.Lifecycle {
	return dsg.Lifecycle
}

// Attributes returns the attributes for [DefaultSecurityGroup].
func (dsg *DefaultSecurityGroup) Attributes() defaultSecurityGroupAttributes {
	return defaultSecurityGroupAttributes{ref: terra.ReferenceResource(dsg)}
}

// ImportState imports the given attribute values into [DefaultSecurityGroup]'s state.
func (dsg *DefaultSecurityGroup) ImportState(av io.Reader) error {
	dsg.state = &defaultSecurityGroupState{}
	if err := json.NewDecoder(av).Decode(dsg.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dsg.Type(), dsg.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DefaultSecurityGroup] has state.
func (dsg *DefaultSecurityGroup) State() (*defaultSecurityGroupState, bool) {
	return dsg.state, dsg.state != nil
}

// StateMust returns the state for [DefaultSecurityGroup]. Panics if the state is nil.
func (dsg *DefaultSecurityGroup) StateMust() *defaultSecurityGroupState {
	if dsg.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dsg.Type(), dsg.LocalName()))
	}
	return dsg.state
}

// DefaultSecurityGroupArgs contains the configurations for aws_default_security_group.
type DefaultSecurityGroupArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// RevokeRulesOnDelete: bool, optional
	RevokeRulesOnDelete terra.BoolValue `hcl:"revoke_rules_on_delete,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// VpcId: string, optional
	VpcId terra.StringValue `hcl:"vpc_id,attr"`
	// Egress: min=0
	Egress []defaultsecuritygroup.Egress `hcl:"egress,block" validate:"min=0"`
	// Ingress: min=0
	Ingress []defaultsecuritygroup.Ingress `hcl:"ingress,block" validate:"min=0"`
}
type defaultSecurityGroupAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_default_security_group.
func (dsg defaultSecurityGroupAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(dsg.ref.Append("arn"))
}

// Description returns a reference to field description of aws_default_security_group.
func (dsg defaultSecurityGroupAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(dsg.ref.Append("description"))
}

// Id returns a reference to field id of aws_default_security_group.
func (dsg defaultSecurityGroupAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dsg.ref.Append("id"))
}

// Name returns a reference to field name of aws_default_security_group.
func (dsg defaultSecurityGroupAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(dsg.ref.Append("name"))
}

// NamePrefix returns a reference to field name_prefix of aws_default_security_group.
func (dsg defaultSecurityGroupAttributes) NamePrefix() terra.StringValue {
	return terra.ReferenceAsString(dsg.ref.Append("name_prefix"))
}

// OwnerId returns a reference to field owner_id of aws_default_security_group.
func (dsg defaultSecurityGroupAttributes) OwnerId() terra.StringValue {
	return terra.ReferenceAsString(dsg.ref.Append("owner_id"))
}

// RevokeRulesOnDelete returns a reference to field revoke_rules_on_delete of aws_default_security_group.
func (dsg defaultSecurityGroupAttributes) RevokeRulesOnDelete() terra.BoolValue {
	return terra.ReferenceAsBool(dsg.ref.Append("revoke_rules_on_delete"))
}

// Tags returns a reference to field tags of aws_default_security_group.
func (dsg defaultSecurityGroupAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dsg.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_default_security_group.
func (dsg defaultSecurityGroupAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dsg.ref.Append("tags_all"))
}

// VpcId returns a reference to field vpc_id of aws_default_security_group.
func (dsg defaultSecurityGroupAttributes) VpcId() terra.StringValue {
	return terra.ReferenceAsString(dsg.ref.Append("vpc_id"))
}

func (dsg defaultSecurityGroupAttributes) Egress() terra.SetValue[defaultsecuritygroup.EgressAttributes] {
	return terra.ReferenceAsSet[defaultsecuritygroup.EgressAttributes](dsg.ref.Append("egress"))
}

func (dsg defaultSecurityGroupAttributes) Ingress() terra.SetValue[defaultsecuritygroup.IngressAttributes] {
	return terra.ReferenceAsSet[defaultsecuritygroup.IngressAttributes](dsg.ref.Append("ingress"))
}

type defaultSecurityGroupState struct {
	Arn                 string                              `json:"arn"`
	Description         string                              `json:"description"`
	Id                  string                              `json:"id"`
	Name                string                              `json:"name"`
	NamePrefix          string                              `json:"name_prefix"`
	OwnerId             string                              `json:"owner_id"`
	RevokeRulesOnDelete bool                                `json:"revoke_rules_on_delete"`
	Tags                map[string]string                   `json:"tags"`
	TagsAll             map[string]string                   `json:"tags_all"`
	VpcId               string                              `json:"vpc_id"`
	Egress              []defaultsecuritygroup.EgressState  `json:"egress"`
	Ingress             []defaultsecuritygroup.IngressState `json:"ingress"`
}