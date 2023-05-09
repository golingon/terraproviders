// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	medialiveinputsecuritygroup "github.com/golingon/terraproviders/aws/4.66.1/medialiveinputsecuritygroup"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewMedialiveInputSecurityGroup creates a new instance of [MedialiveInputSecurityGroup].
func NewMedialiveInputSecurityGroup(name string, args MedialiveInputSecurityGroupArgs) *MedialiveInputSecurityGroup {
	return &MedialiveInputSecurityGroup{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*MedialiveInputSecurityGroup)(nil)

// MedialiveInputSecurityGroup represents the Terraform resource aws_medialive_input_security_group.
type MedialiveInputSecurityGroup struct {
	Name      string
	Args      MedialiveInputSecurityGroupArgs
	state     *medialiveInputSecurityGroupState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [MedialiveInputSecurityGroup].
func (misg *MedialiveInputSecurityGroup) Type() string {
	return "aws_medialive_input_security_group"
}

// LocalName returns the local name for [MedialiveInputSecurityGroup].
func (misg *MedialiveInputSecurityGroup) LocalName() string {
	return misg.Name
}

// Configuration returns the configuration (args) for [MedialiveInputSecurityGroup].
func (misg *MedialiveInputSecurityGroup) Configuration() interface{} {
	return misg.Args
}

// DependOn is used for other resources to depend on [MedialiveInputSecurityGroup].
func (misg *MedialiveInputSecurityGroup) DependOn() terra.Reference {
	return terra.ReferenceResource(misg)
}

// Dependencies returns the list of resources [MedialiveInputSecurityGroup] depends_on.
func (misg *MedialiveInputSecurityGroup) Dependencies() terra.Dependencies {
	return misg.DependsOn
}

// LifecycleManagement returns the lifecycle block for [MedialiveInputSecurityGroup].
func (misg *MedialiveInputSecurityGroup) LifecycleManagement() *terra.Lifecycle {
	return misg.Lifecycle
}

// Attributes returns the attributes for [MedialiveInputSecurityGroup].
func (misg *MedialiveInputSecurityGroup) Attributes() medialiveInputSecurityGroupAttributes {
	return medialiveInputSecurityGroupAttributes{ref: terra.ReferenceResource(misg)}
}

// ImportState imports the given attribute values into [MedialiveInputSecurityGroup]'s state.
func (misg *MedialiveInputSecurityGroup) ImportState(av io.Reader) error {
	misg.state = &medialiveInputSecurityGroupState{}
	if err := json.NewDecoder(av).Decode(misg.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", misg.Type(), misg.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [MedialiveInputSecurityGroup] has state.
func (misg *MedialiveInputSecurityGroup) State() (*medialiveInputSecurityGroupState, bool) {
	return misg.state, misg.state != nil
}

// StateMust returns the state for [MedialiveInputSecurityGroup]. Panics if the state is nil.
func (misg *MedialiveInputSecurityGroup) StateMust() *medialiveInputSecurityGroupState {
	if misg.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", misg.Type(), misg.LocalName()))
	}
	return misg.state
}

// MedialiveInputSecurityGroupArgs contains the configurations for aws_medialive_input_security_group.
type MedialiveInputSecurityGroupArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// Timeouts: optional
	Timeouts *medialiveinputsecuritygroup.Timeouts `hcl:"timeouts,block"`
	// WhitelistRules: min=1
	WhitelistRules []medialiveinputsecuritygroup.WhitelistRules `hcl:"whitelist_rules,block" validate:"min=1"`
}
type medialiveInputSecurityGroupAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_medialive_input_security_group.
func (misg medialiveInputSecurityGroupAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(misg.ref.Append("arn"))
}

// Id returns a reference to field id of aws_medialive_input_security_group.
func (misg medialiveInputSecurityGroupAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(misg.ref.Append("id"))
}

// Inputs returns a reference to field inputs of aws_medialive_input_security_group.
func (misg medialiveInputSecurityGroupAttributes) Inputs() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](misg.ref.Append("inputs"))
}

// Tags returns a reference to field tags of aws_medialive_input_security_group.
func (misg medialiveInputSecurityGroupAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](misg.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_medialive_input_security_group.
func (misg medialiveInputSecurityGroupAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](misg.ref.Append("tags_all"))
}

func (misg medialiveInputSecurityGroupAttributes) Timeouts() medialiveinputsecuritygroup.TimeoutsAttributes {
	return terra.ReferenceAsSingle[medialiveinputsecuritygroup.TimeoutsAttributes](misg.ref.Append("timeouts"))
}

func (misg medialiveInputSecurityGroupAttributes) WhitelistRules() terra.SetValue[medialiveinputsecuritygroup.WhitelistRulesAttributes] {
	return terra.ReferenceAsSet[medialiveinputsecuritygroup.WhitelistRulesAttributes](misg.ref.Append("whitelist_rules"))
}

type medialiveInputSecurityGroupState struct {
	Arn            string                                            `json:"arn"`
	Id             string                                            `json:"id"`
	Inputs         []string                                          `json:"inputs"`
	Tags           map[string]string                                 `json:"tags"`
	TagsAll        map[string]string                                 `json:"tags_all"`
	Timeouts       *medialiveinputsecuritygroup.TimeoutsState        `json:"timeouts"`
	WhitelistRules []medialiveinputsecuritygroup.WhitelistRulesState `json:"whitelist_rules"`
}
