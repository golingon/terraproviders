// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	locationplaceindex "github.com/golingon/terraproviders/aws/4.63.0/locationplaceindex"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewLocationPlaceIndex creates a new instance of [LocationPlaceIndex].
func NewLocationPlaceIndex(name string, args LocationPlaceIndexArgs) *LocationPlaceIndex {
	return &LocationPlaceIndex{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*LocationPlaceIndex)(nil)

// LocationPlaceIndex represents the Terraform resource aws_location_place_index.
type LocationPlaceIndex struct {
	Name      string
	Args      LocationPlaceIndexArgs
	state     *locationPlaceIndexState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [LocationPlaceIndex].
func (lpi *LocationPlaceIndex) Type() string {
	return "aws_location_place_index"
}

// LocalName returns the local name for [LocationPlaceIndex].
func (lpi *LocationPlaceIndex) LocalName() string {
	return lpi.Name
}

// Configuration returns the configuration (args) for [LocationPlaceIndex].
func (lpi *LocationPlaceIndex) Configuration() interface{} {
	return lpi.Args
}

// DependOn is used for other resources to depend on [LocationPlaceIndex].
func (lpi *LocationPlaceIndex) DependOn() terra.Reference {
	return terra.ReferenceResource(lpi)
}

// Dependencies returns the list of resources [LocationPlaceIndex] depends_on.
func (lpi *LocationPlaceIndex) Dependencies() terra.Dependencies {
	return lpi.DependsOn
}

// LifecycleManagement returns the lifecycle block for [LocationPlaceIndex].
func (lpi *LocationPlaceIndex) LifecycleManagement() *terra.Lifecycle {
	return lpi.Lifecycle
}

// Attributes returns the attributes for [LocationPlaceIndex].
func (lpi *LocationPlaceIndex) Attributes() locationPlaceIndexAttributes {
	return locationPlaceIndexAttributes{ref: terra.ReferenceResource(lpi)}
}

// ImportState imports the given attribute values into [LocationPlaceIndex]'s state.
func (lpi *LocationPlaceIndex) ImportState(av io.Reader) error {
	lpi.state = &locationPlaceIndexState{}
	if err := json.NewDecoder(av).Decode(lpi.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", lpi.Type(), lpi.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [LocationPlaceIndex] has state.
func (lpi *LocationPlaceIndex) State() (*locationPlaceIndexState, bool) {
	return lpi.state, lpi.state != nil
}

// StateMust returns the state for [LocationPlaceIndex]. Panics if the state is nil.
func (lpi *LocationPlaceIndex) StateMust() *locationPlaceIndexState {
	if lpi.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", lpi.Type(), lpi.LocalName()))
	}
	return lpi.state
}

// LocationPlaceIndexArgs contains the configurations for aws_location_place_index.
type LocationPlaceIndexArgs struct {
	// DataSource: string, required
	DataSource terra.StringValue `hcl:"data_source,attr" validate:"required"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IndexName: string, required
	IndexName terra.StringValue `hcl:"index_name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// DataSourceConfiguration: optional
	DataSourceConfiguration *locationplaceindex.DataSourceConfiguration `hcl:"data_source_configuration,block"`
}
type locationPlaceIndexAttributes struct {
	ref terra.Reference
}

// CreateTime returns a reference to field create_time of aws_location_place_index.
func (lpi locationPlaceIndexAttributes) CreateTime() terra.StringValue {
	return terra.ReferenceAsString(lpi.ref.Append("create_time"))
}

// DataSource returns a reference to field data_source of aws_location_place_index.
func (lpi locationPlaceIndexAttributes) DataSource() terra.StringValue {
	return terra.ReferenceAsString(lpi.ref.Append("data_source"))
}

// Description returns a reference to field description of aws_location_place_index.
func (lpi locationPlaceIndexAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(lpi.ref.Append("description"))
}

// Id returns a reference to field id of aws_location_place_index.
func (lpi locationPlaceIndexAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(lpi.ref.Append("id"))
}

// IndexArn returns a reference to field index_arn of aws_location_place_index.
func (lpi locationPlaceIndexAttributes) IndexArn() terra.StringValue {
	return terra.ReferenceAsString(lpi.ref.Append("index_arn"))
}

// IndexName returns a reference to field index_name of aws_location_place_index.
func (lpi locationPlaceIndexAttributes) IndexName() terra.StringValue {
	return terra.ReferenceAsString(lpi.ref.Append("index_name"))
}

// Tags returns a reference to field tags of aws_location_place_index.
func (lpi locationPlaceIndexAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](lpi.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_location_place_index.
func (lpi locationPlaceIndexAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](lpi.ref.Append("tags_all"))
}

// UpdateTime returns a reference to field update_time of aws_location_place_index.
func (lpi locationPlaceIndexAttributes) UpdateTime() terra.StringValue {
	return terra.ReferenceAsString(lpi.ref.Append("update_time"))
}

func (lpi locationPlaceIndexAttributes) DataSourceConfiguration() terra.ListValue[locationplaceindex.DataSourceConfigurationAttributes] {
	return terra.ReferenceAsList[locationplaceindex.DataSourceConfigurationAttributes](lpi.ref.Append("data_source_configuration"))
}

type locationPlaceIndexState struct {
	CreateTime              string                                            `json:"create_time"`
	DataSource              string                                            `json:"data_source"`
	Description             string                                            `json:"description"`
	Id                      string                                            `json:"id"`
	IndexArn                string                                            `json:"index_arn"`
	IndexName               string                                            `json:"index_name"`
	Tags                    map[string]string                                 `json:"tags"`
	TagsAll                 map[string]string                                 `json:"tags_all"`
	UpdateTime              string                                            `json:"update_time"`
	DataSourceConfiguration []locationplaceindex.DataSourceConfigurationState `json:"data_source_configuration"`
}
