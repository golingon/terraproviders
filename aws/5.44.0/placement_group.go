// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// NewPlacementGroup creates a new instance of [PlacementGroup].
func NewPlacementGroup(name string, args PlacementGroupArgs) *PlacementGroup {
	return &PlacementGroup{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*PlacementGroup)(nil)

// PlacementGroup represents the Terraform resource aws_placement_group.
type PlacementGroup struct {
	Name      string
	Args      PlacementGroupArgs
	state     *placementGroupState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [PlacementGroup].
func (pg *PlacementGroup) Type() string {
	return "aws_placement_group"
}

// LocalName returns the local name for [PlacementGroup].
func (pg *PlacementGroup) LocalName() string {
	return pg.Name
}

// Configuration returns the configuration (args) for [PlacementGroup].
func (pg *PlacementGroup) Configuration() interface{} {
	return pg.Args
}

// DependOn is used for other resources to depend on [PlacementGroup].
func (pg *PlacementGroup) DependOn() terra.Reference {
	return terra.ReferenceResource(pg)
}

// Dependencies returns the list of resources [PlacementGroup] depends_on.
func (pg *PlacementGroup) Dependencies() terra.Dependencies {
	return pg.DependsOn
}

// LifecycleManagement returns the lifecycle block for [PlacementGroup].
func (pg *PlacementGroup) LifecycleManagement() *terra.Lifecycle {
	return pg.Lifecycle
}

// Attributes returns the attributes for [PlacementGroup].
func (pg *PlacementGroup) Attributes() placementGroupAttributes {
	return placementGroupAttributes{ref: terra.ReferenceResource(pg)}
}

// ImportState imports the given attribute values into [PlacementGroup]'s state.
func (pg *PlacementGroup) ImportState(av io.Reader) error {
	pg.state = &placementGroupState{}
	if err := json.NewDecoder(av).Decode(pg.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", pg.Type(), pg.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [PlacementGroup] has state.
func (pg *PlacementGroup) State() (*placementGroupState, bool) {
	return pg.state, pg.state != nil
}

// StateMust returns the state for [PlacementGroup]. Panics if the state is nil.
func (pg *PlacementGroup) StateMust() *placementGroupState {
	if pg.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", pg.Type(), pg.LocalName()))
	}
	return pg.state
}

// PlacementGroupArgs contains the configurations for aws_placement_group.
type PlacementGroupArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// PartitionCount: number, optional
	PartitionCount terra.NumberValue `hcl:"partition_count,attr"`
	// SpreadLevel: string, optional
	SpreadLevel terra.StringValue `hcl:"spread_level,attr"`
	// Strategy: string, required
	Strategy terra.StringValue `hcl:"strategy,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
}
type placementGroupAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_placement_group.
func (pg placementGroupAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(pg.ref.Append("arn"))
}

// Id returns a reference to field id of aws_placement_group.
func (pg placementGroupAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(pg.ref.Append("id"))
}

// Name returns a reference to field name of aws_placement_group.
func (pg placementGroupAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(pg.ref.Append("name"))
}

// PartitionCount returns a reference to field partition_count of aws_placement_group.
func (pg placementGroupAttributes) PartitionCount() terra.NumberValue {
	return terra.ReferenceAsNumber(pg.ref.Append("partition_count"))
}

// PlacementGroupId returns a reference to field placement_group_id of aws_placement_group.
func (pg placementGroupAttributes) PlacementGroupId() terra.StringValue {
	return terra.ReferenceAsString(pg.ref.Append("placement_group_id"))
}

// SpreadLevel returns a reference to field spread_level of aws_placement_group.
func (pg placementGroupAttributes) SpreadLevel() terra.StringValue {
	return terra.ReferenceAsString(pg.ref.Append("spread_level"))
}

// Strategy returns a reference to field strategy of aws_placement_group.
func (pg placementGroupAttributes) Strategy() terra.StringValue {
	return terra.ReferenceAsString(pg.ref.Append("strategy"))
}

// Tags returns a reference to field tags of aws_placement_group.
func (pg placementGroupAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](pg.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_placement_group.
func (pg placementGroupAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](pg.ref.Append("tags_all"))
}

type placementGroupState struct {
	Arn              string            `json:"arn"`
	Id               string            `json:"id"`
	Name             string            `json:"name"`
	PartitionCount   float64           `json:"partition_count"`
	PlacementGroupId string            `json:"placement_group_id"`
	SpreadLevel      string            `json:"spread_level"`
	Strategy         string            `json:"strategy"`
	Tags             map[string]string `json:"tags"`
	TagsAll          map[string]string `json:"tags_all"`
}
