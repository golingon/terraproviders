// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewMemorydbSubnetGroup creates a new instance of [MemorydbSubnetGroup].
func NewMemorydbSubnetGroup(name string, args MemorydbSubnetGroupArgs) *MemorydbSubnetGroup {
	return &MemorydbSubnetGroup{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*MemorydbSubnetGroup)(nil)

// MemorydbSubnetGroup represents the Terraform resource aws_memorydb_subnet_group.
type MemorydbSubnetGroup struct {
	Name      string
	Args      MemorydbSubnetGroupArgs
	state     *memorydbSubnetGroupState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [MemorydbSubnetGroup].
func (msg *MemorydbSubnetGroup) Type() string {
	return "aws_memorydb_subnet_group"
}

// LocalName returns the local name for [MemorydbSubnetGroup].
func (msg *MemorydbSubnetGroup) LocalName() string {
	return msg.Name
}

// Configuration returns the configuration (args) for [MemorydbSubnetGroup].
func (msg *MemorydbSubnetGroup) Configuration() interface{} {
	return msg.Args
}

// DependOn is used for other resources to depend on [MemorydbSubnetGroup].
func (msg *MemorydbSubnetGroup) DependOn() terra.Reference {
	return terra.ReferenceResource(msg)
}

// Dependencies returns the list of resources [MemorydbSubnetGroup] depends_on.
func (msg *MemorydbSubnetGroup) Dependencies() terra.Dependencies {
	return msg.DependsOn
}

// LifecycleManagement returns the lifecycle block for [MemorydbSubnetGroup].
func (msg *MemorydbSubnetGroup) LifecycleManagement() *terra.Lifecycle {
	return msg.Lifecycle
}

// Attributes returns the attributes for [MemorydbSubnetGroup].
func (msg *MemorydbSubnetGroup) Attributes() memorydbSubnetGroupAttributes {
	return memorydbSubnetGroupAttributes{ref: terra.ReferenceResource(msg)}
}

// ImportState imports the given attribute values into [MemorydbSubnetGroup]'s state.
func (msg *MemorydbSubnetGroup) ImportState(av io.Reader) error {
	msg.state = &memorydbSubnetGroupState{}
	if err := json.NewDecoder(av).Decode(msg.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", msg.Type(), msg.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [MemorydbSubnetGroup] has state.
func (msg *MemorydbSubnetGroup) State() (*memorydbSubnetGroupState, bool) {
	return msg.state, msg.state != nil
}

// StateMust returns the state for [MemorydbSubnetGroup]. Panics if the state is nil.
func (msg *MemorydbSubnetGroup) StateMust() *memorydbSubnetGroupState {
	if msg.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", msg.Type(), msg.LocalName()))
	}
	return msg.state
}

// MemorydbSubnetGroupArgs contains the configurations for aws_memorydb_subnet_group.
type MemorydbSubnetGroupArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
	// NamePrefix: string, optional
	NamePrefix terra.StringValue `hcl:"name_prefix,attr"`
	// SubnetIds: set of string, required
	SubnetIds terra.SetValue[terra.StringValue] `hcl:"subnet_ids,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
}
type memorydbSubnetGroupAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_memorydb_subnet_group.
func (msg memorydbSubnetGroupAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(msg.ref.Append("arn"))
}

// Description returns a reference to field description of aws_memorydb_subnet_group.
func (msg memorydbSubnetGroupAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(msg.ref.Append("description"))
}

// Id returns a reference to field id of aws_memorydb_subnet_group.
func (msg memorydbSubnetGroupAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(msg.ref.Append("id"))
}

// Name returns a reference to field name of aws_memorydb_subnet_group.
func (msg memorydbSubnetGroupAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(msg.ref.Append("name"))
}

// NamePrefix returns a reference to field name_prefix of aws_memorydb_subnet_group.
func (msg memorydbSubnetGroupAttributes) NamePrefix() terra.StringValue {
	return terra.ReferenceAsString(msg.ref.Append("name_prefix"))
}

// SubnetIds returns a reference to field subnet_ids of aws_memorydb_subnet_group.
func (msg memorydbSubnetGroupAttributes) SubnetIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](msg.ref.Append("subnet_ids"))
}

// Tags returns a reference to field tags of aws_memorydb_subnet_group.
func (msg memorydbSubnetGroupAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](msg.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_memorydb_subnet_group.
func (msg memorydbSubnetGroupAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](msg.ref.Append("tags_all"))
}

// VpcId returns a reference to field vpc_id of aws_memorydb_subnet_group.
func (msg memorydbSubnetGroupAttributes) VpcId() terra.StringValue {
	return terra.ReferenceAsString(msg.ref.Append("vpc_id"))
}

type memorydbSubnetGroupState struct {
	Arn         string            `json:"arn"`
	Description string            `json:"description"`
	Id          string            `json:"id"`
	Name        string            `json:"name"`
	NamePrefix  string            `json:"name_prefix"`
	SubnetIds   []string          `json:"subnet_ids"`
	Tags        map[string]string `json:"tags"`
	TagsAll     map[string]string `json:"tags_all"`
	VpcId       string            `json:"vpc_id"`
}
