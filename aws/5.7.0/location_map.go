// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	locationmap "github.com/golingon/terraproviders/aws/5.7.0/locationmap"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewLocationMap creates a new instance of [LocationMap].
func NewLocationMap(name string, args LocationMapArgs) *LocationMap {
	return &LocationMap{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*LocationMap)(nil)

// LocationMap represents the Terraform resource aws_location_map.
type LocationMap struct {
	Name      string
	Args      LocationMapArgs
	state     *locationMapState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [LocationMap].
func (lm *LocationMap) Type() string {
	return "aws_location_map"
}

// LocalName returns the local name for [LocationMap].
func (lm *LocationMap) LocalName() string {
	return lm.Name
}

// Configuration returns the configuration (args) for [LocationMap].
func (lm *LocationMap) Configuration() interface{} {
	return lm.Args
}

// DependOn is used for other resources to depend on [LocationMap].
func (lm *LocationMap) DependOn() terra.Reference {
	return terra.ReferenceResource(lm)
}

// Dependencies returns the list of resources [LocationMap] depends_on.
func (lm *LocationMap) Dependencies() terra.Dependencies {
	return lm.DependsOn
}

// LifecycleManagement returns the lifecycle block for [LocationMap].
func (lm *LocationMap) LifecycleManagement() *terra.Lifecycle {
	return lm.Lifecycle
}

// Attributes returns the attributes for [LocationMap].
func (lm *LocationMap) Attributes() locationMapAttributes {
	return locationMapAttributes{ref: terra.ReferenceResource(lm)}
}

// ImportState imports the given attribute values into [LocationMap]'s state.
func (lm *LocationMap) ImportState(av io.Reader) error {
	lm.state = &locationMapState{}
	if err := json.NewDecoder(av).Decode(lm.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", lm.Type(), lm.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [LocationMap] has state.
func (lm *LocationMap) State() (*locationMapState, bool) {
	return lm.state, lm.state != nil
}

// StateMust returns the state for [LocationMap]. Panics if the state is nil.
func (lm *LocationMap) StateMust() *locationMapState {
	if lm.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", lm.Type(), lm.LocalName()))
	}
	return lm.state
}

// LocationMapArgs contains the configurations for aws_location_map.
type LocationMapArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// MapName: string, required
	MapName terra.StringValue `hcl:"map_name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// Configuration: required
	Configuration *locationmap.Configuration `hcl:"configuration,block" validate:"required"`
}
type locationMapAttributes struct {
	ref terra.Reference
}

// CreateTime returns a reference to field create_time of aws_location_map.
func (lm locationMapAttributes) CreateTime() terra.StringValue {
	return terra.ReferenceAsString(lm.ref.Append("create_time"))
}

// Description returns a reference to field description of aws_location_map.
func (lm locationMapAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(lm.ref.Append("description"))
}

// Id returns a reference to field id of aws_location_map.
func (lm locationMapAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(lm.ref.Append("id"))
}

// MapArn returns a reference to field map_arn of aws_location_map.
func (lm locationMapAttributes) MapArn() terra.StringValue {
	return terra.ReferenceAsString(lm.ref.Append("map_arn"))
}

// MapName returns a reference to field map_name of aws_location_map.
func (lm locationMapAttributes) MapName() terra.StringValue {
	return terra.ReferenceAsString(lm.ref.Append("map_name"))
}

// Tags returns a reference to field tags of aws_location_map.
func (lm locationMapAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](lm.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_location_map.
func (lm locationMapAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](lm.ref.Append("tags_all"))
}

// UpdateTime returns a reference to field update_time of aws_location_map.
func (lm locationMapAttributes) UpdateTime() terra.StringValue {
	return terra.ReferenceAsString(lm.ref.Append("update_time"))
}

func (lm locationMapAttributes) Configuration() terra.ListValue[locationmap.ConfigurationAttributes] {
	return terra.ReferenceAsList[locationmap.ConfigurationAttributes](lm.ref.Append("configuration"))
}

type locationMapState struct {
	CreateTime    string                           `json:"create_time"`
	Description   string                           `json:"description"`
	Id            string                           `json:"id"`
	MapArn        string                           `json:"map_arn"`
	MapName       string                           `json:"map_name"`
	Tags          map[string]string                `json:"tags"`
	TagsAll       map[string]string                `json:"tags_all"`
	UpdateTime    string                           `json:"update_time"`
	Configuration []locationmap.ConfigurationState `json:"configuration"`
}
