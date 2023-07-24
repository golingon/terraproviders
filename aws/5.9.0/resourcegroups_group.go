// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	resourcegroupsgroup "github.com/golingon/terraproviders/aws/5.9.0/resourcegroupsgroup"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewResourcegroupsGroup creates a new instance of [ResourcegroupsGroup].
func NewResourcegroupsGroup(name string, args ResourcegroupsGroupArgs) *ResourcegroupsGroup {
	return &ResourcegroupsGroup{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ResourcegroupsGroup)(nil)

// ResourcegroupsGroup represents the Terraform resource aws_resourcegroups_group.
type ResourcegroupsGroup struct {
	Name      string
	Args      ResourcegroupsGroupArgs
	state     *resourcegroupsGroupState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ResourcegroupsGroup].
func (rg *ResourcegroupsGroup) Type() string {
	return "aws_resourcegroups_group"
}

// LocalName returns the local name for [ResourcegroupsGroup].
func (rg *ResourcegroupsGroup) LocalName() string {
	return rg.Name
}

// Configuration returns the configuration (args) for [ResourcegroupsGroup].
func (rg *ResourcegroupsGroup) Configuration() interface{} {
	return rg.Args
}

// DependOn is used for other resources to depend on [ResourcegroupsGroup].
func (rg *ResourcegroupsGroup) DependOn() terra.Reference {
	return terra.ReferenceResource(rg)
}

// Dependencies returns the list of resources [ResourcegroupsGroup] depends_on.
func (rg *ResourcegroupsGroup) Dependencies() terra.Dependencies {
	return rg.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ResourcegroupsGroup].
func (rg *ResourcegroupsGroup) LifecycleManagement() *terra.Lifecycle {
	return rg.Lifecycle
}

// Attributes returns the attributes for [ResourcegroupsGroup].
func (rg *ResourcegroupsGroup) Attributes() resourcegroupsGroupAttributes {
	return resourcegroupsGroupAttributes{ref: terra.ReferenceResource(rg)}
}

// ImportState imports the given attribute values into [ResourcegroupsGroup]'s state.
func (rg *ResourcegroupsGroup) ImportState(av io.Reader) error {
	rg.state = &resourcegroupsGroupState{}
	if err := json.NewDecoder(av).Decode(rg.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", rg.Type(), rg.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ResourcegroupsGroup] has state.
func (rg *ResourcegroupsGroup) State() (*resourcegroupsGroupState, bool) {
	return rg.state, rg.state != nil
}

// StateMust returns the state for [ResourcegroupsGroup]. Panics if the state is nil.
func (rg *ResourcegroupsGroup) StateMust() *resourcegroupsGroupState {
	if rg.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", rg.Type(), rg.LocalName()))
	}
	return rg.state
}

// ResourcegroupsGroupArgs contains the configurations for aws_resourcegroups_group.
type ResourcegroupsGroupArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// Configuration: min=0
	Configuration []resourcegroupsgroup.Configuration `hcl:"configuration,block" validate:"min=0"`
	// ResourceQuery: optional
	ResourceQuery *resourcegroupsgroup.ResourceQuery `hcl:"resource_query,block"`
	// Timeouts: optional
	Timeouts *resourcegroupsgroup.Timeouts `hcl:"timeouts,block"`
}
type resourcegroupsGroupAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_resourcegroups_group.
func (rg resourcegroupsGroupAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(rg.ref.Append("arn"))
}

// Description returns a reference to field description of aws_resourcegroups_group.
func (rg resourcegroupsGroupAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(rg.ref.Append("description"))
}

// Id returns a reference to field id of aws_resourcegroups_group.
func (rg resourcegroupsGroupAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(rg.ref.Append("id"))
}

// Name returns a reference to field name of aws_resourcegroups_group.
func (rg resourcegroupsGroupAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(rg.ref.Append("name"))
}

// Tags returns a reference to field tags of aws_resourcegroups_group.
func (rg resourcegroupsGroupAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](rg.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_resourcegroups_group.
func (rg resourcegroupsGroupAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](rg.ref.Append("tags_all"))
}

func (rg resourcegroupsGroupAttributes) Configuration() terra.SetValue[resourcegroupsgroup.ConfigurationAttributes] {
	return terra.ReferenceAsSet[resourcegroupsgroup.ConfigurationAttributes](rg.ref.Append("configuration"))
}

func (rg resourcegroupsGroupAttributes) ResourceQuery() terra.ListValue[resourcegroupsgroup.ResourceQueryAttributes] {
	return terra.ReferenceAsList[resourcegroupsgroup.ResourceQueryAttributes](rg.ref.Append("resource_query"))
}

func (rg resourcegroupsGroupAttributes) Timeouts() resourcegroupsgroup.TimeoutsAttributes {
	return terra.ReferenceAsSingle[resourcegroupsgroup.TimeoutsAttributes](rg.ref.Append("timeouts"))
}

type resourcegroupsGroupState struct {
	Arn           string                                   `json:"arn"`
	Description   string                                   `json:"description"`
	Id            string                                   `json:"id"`
	Name          string                                   `json:"name"`
	Tags          map[string]string                        `json:"tags"`
	TagsAll       map[string]string                        `json:"tags_all"`
	Configuration []resourcegroupsgroup.ConfigurationState `json:"configuration"`
	ResourceQuery []resourcegroupsgroup.ResourceQueryState `json:"resource_query"`
	Timeouts      *resourcegroupsgroup.TimeoutsState       `json:"timeouts"`
}
