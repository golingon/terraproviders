// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	ec2managedprefixlist "github.com/golingon/terraproviders/aws/5.8.0/ec2managedprefixlist"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewEc2ManagedPrefixList creates a new instance of [Ec2ManagedPrefixList].
func NewEc2ManagedPrefixList(name string, args Ec2ManagedPrefixListArgs) *Ec2ManagedPrefixList {
	return &Ec2ManagedPrefixList{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Ec2ManagedPrefixList)(nil)

// Ec2ManagedPrefixList represents the Terraform resource aws_ec2_managed_prefix_list.
type Ec2ManagedPrefixList struct {
	Name      string
	Args      Ec2ManagedPrefixListArgs
	state     *ec2ManagedPrefixListState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Ec2ManagedPrefixList].
func (empl *Ec2ManagedPrefixList) Type() string {
	return "aws_ec2_managed_prefix_list"
}

// LocalName returns the local name for [Ec2ManagedPrefixList].
func (empl *Ec2ManagedPrefixList) LocalName() string {
	return empl.Name
}

// Configuration returns the configuration (args) for [Ec2ManagedPrefixList].
func (empl *Ec2ManagedPrefixList) Configuration() interface{} {
	return empl.Args
}

// DependOn is used for other resources to depend on [Ec2ManagedPrefixList].
func (empl *Ec2ManagedPrefixList) DependOn() terra.Reference {
	return terra.ReferenceResource(empl)
}

// Dependencies returns the list of resources [Ec2ManagedPrefixList] depends_on.
func (empl *Ec2ManagedPrefixList) Dependencies() terra.Dependencies {
	return empl.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Ec2ManagedPrefixList].
func (empl *Ec2ManagedPrefixList) LifecycleManagement() *terra.Lifecycle {
	return empl.Lifecycle
}

// Attributes returns the attributes for [Ec2ManagedPrefixList].
func (empl *Ec2ManagedPrefixList) Attributes() ec2ManagedPrefixListAttributes {
	return ec2ManagedPrefixListAttributes{ref: terra.ReferenceResource(empl)}
}

// ImportState imports the given attribute values into [Ec2ManagedPrefixList]'s state.
func (empl *Ec2ManagedPrefixList) ImportState(av io.Reader) error {
	empl.state = &ec2ManagedPrefixListState{}
	if err := json.NewDecoder(av).Decode(empl.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", empl.Type(), empl.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Ec2ManagedPrefixList] has state.
func (empl *Ec2ManagedPrefixList) State() (*ec2ManagedPrefixListState, bool) {
	return empl.state, empl.state != nil
}

// StateMust returns the state for [Ec2ManagedPrefixList]. Panics if the state is nil.
func (empl *Ec2ManagedPrefixList) StateMust() *ec2ManagedPrefixListState {
	if empl.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", empl.Type(), empl.LocalName()))
	}
	return empl.state
}

// Ec2ManagedPrefixListArgs contains the configurations for aws_ec2_managed_prefix_list.
type Ec2ManagedPrefixListArgs struct {
	// AddressFamily: string, required
	AddressFamily terra.StringValue `hcl:"address_family,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// MaxEntries: number, required
	MaxEntries terra.NumberValue `hcl:"max_entries,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// Entry: min=0
	Entry []ec2managedprefixlist.Entry `hcl:"entry,block" validate:"min=0"`
}
type ec2ManagedPrefixListAttributes struct {
	ref terra.Reference
}

// AddressFamily returns a reference to field address_family of aws_ec2_managed_prefix_list.
func (empl ec2ManagedPrefixListAttributes) AddressFamily() terra.StringValue {
	return terra.ReferenceAsString(empl.ref.Append("address_family"))
}

// Arn returns a reference to field arn of aws_ec2_managed_prefix_list.
func (empl ec2ManagedPrefixListAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(empl.ref.Append("arn"))
}

// Id returns a reference to field id of aws_ec2_managed_prefix_list.
func (empl ec2ManagedPrefixListAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(empl.ref.Append("id"))
}

// MaxEntries returns a reference to field max_entries of aws_ec2_managed_prefix_list.
func (empl ec2ManagedPrefixListAttributes) MaxEntries() terra.NumberValue {
	return terra.ReferenceAsNumber(empl.ref.Append("max_entries"))
}

// Name returns a reference to field name of aws_ec2_managed_prefix_list.
func (empl ec2ManagedPrefixListAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(empl.ref.Append("name"))
}

// OwnerId returns a reference to field owner_id of aws_ec2_managed_prefix_list.
func (empl ec2ManagedPrefixListAttributes) OwnerId() terra.StringValue {
	return terra.ReferenceAsString(empl.ref.Append("owner_id"))
}

// Tags returns a reference to field tags of aws_ec2_managed_prefix_list.
func (empl ec2ManagedPrefixListAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](empl.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_ec2_managed_prefix_list.
func (empl ec2ManagedPrefixListAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](empl.ref.Append("tags_all"))
}

// Version returns a reference to field version of aws_ec2_managed_prefix_list.
func (empl ec2ManagedPrefixListAttributes) Version() terra.NumberValue {
	return terra.ReferenceAsNumber(empl.ref.Append("version"))
}

func (empl ec2ManagedPrefixListAttributes) Entry() terra.SetValue[ec2managedprefixlist.EntryAttributes] {
	return terra.ReferenceAsSet[ec2managedprefixlist.EntryAttributes](empl.ref.Append("entry"))
}

type ec2ManagedPrefixListState struct {
	AddressFamily string                            `json:"address_family"`
	Arn           string                            `json:"arn"`
	Id            string                            `json:"id"`
	MaxEntries    float64                           `json:"max_entries"`
	Name          string                            `json:"name"`
	OwnerId       string                            `json:"owner_id"`
	Tags          map[string]string                 `json:"tags"`
	TagsAll       map[string]string                 `json:"tags_all"`
	Version       float64                           `json:"version"`
	Entry         []ec2managedprefixlist.EntryState `json:"entry"`
}
