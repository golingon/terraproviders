// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_shield_protection_group

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

// Resource represents the Terraform resource aws_shield_protection_group.
type Resource struct {
	Name      string
	Args      Args
	state     *awsShieldProtectionGroupState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (aspg *Resource) Type() string {
	return "aws_shield_protection_group"
}

// LocalName returns the local name for [Resource].
func (aspg *Resource) LocalName() string {
	return aspg.Name
}

// Configuration returns the configuration (args) for [Resource].
func (aspg *Resource) Configuration() interface{} {
	return aspg.Args
}

// DependOn is used for other resources to depend on [Resource].
func (aspg *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(aspg)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (aspg *Resource) Dependencies() terra.Dependencies {
	return aspg.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (aspg *Resource) LifecycleManagement() *terra.Lifecycle {
	return aspg.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (aspg *Resource) Attributes() awsShieldProtectionGroupAttributes {
	return awsShieldProtectionGroupAttributes{ref: terra.ReferenceResource(aspg)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (aspg *Resource) ImportState(state io.Reader) error {
	aspg.state = &awsShieldProtectionGroupState{}
	if err := json.NewDecoder(state).Decode(aspg.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", aspg.Type(), aspg.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (aspg *Resource) State() (*awsShieldProtectionGroupState, bool) {
	return aspg.state, aspg.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (aspg *Resource) StateMust() *awsShieldProtectionGroupState {
	if aspg.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", aspg.Type(), aspg.LocalName()))
	}
	return aspg.state
}

// Args contains the configurations for aws_shield_protection_group.
type Args struct {
	// Aggregation: string, required
	Aggregation terra.StringValue `hcl:"aggregation,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Members: list of string, optional
	Members terra.ListValue[terra.StringValue] `hcl:"members,attr"`
	// Pattern: string, required
	Pattern terra.StringValue `hcl:"pattern,attr" validate:"required"`
	// ProtectionGroupId: string, required
	ProtectionGroupId terra.StringValue `hcl:"protection_group_id,attr" validate:"required"`
	// ResourceType: string, optional
	ResourceType terra.StringValue `hcl:"resource_type,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
}

type awsShieldProtectionGroupAttributes struct {
	ref terra.Reference
}

// Aggregation returns a reference to field aggregation of aws_shield_protection_group.
func (aspg awsShieldProtectionGroupAttributes) Aggregation() terra.StringValue {
	return terra.ReferenceAsString(aspg.ref.Append("aggregation"))
}

// Id returns a reference to field id of aws_shield_protection_group.
func (aspg awsShieldProtectionGroupAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(aspg.ref.Append("id"))
}

// Members returns a reference to field members of aws_shield_protection_group.
func (aspg awsShieldProtectionGroupAttributes) Members() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](aspg.ref.Append("members"))
}

// Pattern returns a reference to field pattern of aws_shield_protection_group.
func (aspg awsShieldProtectionGroupAttributes) Pattern() terra.StringValue {
	return terra.ReferenceAsString(aspg.ref.Append("pattern"))
}

// ProtectionGroupArn returns a reference to field protection_group_arn of aws_shield_protection_group.
func (aspg awsShieldProtectionGroupAttributes) ProtectionGroupArn() terra.StringValue {
	return terra.ReferenceAsString(aspg.ref.Append("protection_group_arn"))
}

// ProtectionGroupId returns a reference to field protection_group_id of aws_shield_protection_group.
func (aspg awsShieldProtectionGroupAttributes) ProtectionGroupId() terra.StringValue {
	return terra.ReferenceAsString(aspg.ref.Append("protection_group_id"))
}

// ResourceType returns a reference to field resource_type of aws_shield_protection_group.
func (aspg awsShieldProtectionGroupAttributes) ResourceType() terra.StringValue {
	return terra.ReferenceAsString(aspg.ref.Append("resource_type"))
}

// Tags returns a reference to field tags of aws_shield_protection_group.
func (aspg awsShieldProtectionGroupAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](aspg.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_shield_protection_group.
func (aspg awsShieldProtectionGroupAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](aspg.ref.Append("tags_all"))
}

type awsShieldProtectionGroupState struct {
	Aggregation        string            `json:"aggregation"`
	Id                 string            `json:"id"`
	Members            []string          `json:"members"`
	Pattern            string            `json:"pattern"`
	ProtectionGroupArn string            `json:"protection_group_arn"`
	ProtectionGroupId  string            `json:"protection_group_id"`
	ResourceType       string            `json:"resource_type"`
	Tags               map[string]string `json:"tags"`
	TagsAll            map[string]string `json:"tags_all"`
}
