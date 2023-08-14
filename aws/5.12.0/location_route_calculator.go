// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	locationroutecalculator "github.com/golingon/terraproviders/aws/5.12.0/locationroutecalculator"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewLocationRouteCalculator creates a new instance of [LocationRouteCalculator].
func NewLocationRouteCalculator(name string, args LocationRouteCalculatorArgs) *LocationRouteCalculator {
	return &LocationRouteCalculator{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*LocationRouteCalculator)(nil)

// LocationRouteCalculator represents the Terraform resource aws_location_route_calculator.
type LocationRouteCalculator struct {
	Name      string
	Args      LocationRouteCalculatorArgs
	state     *locationRouteCalculatorState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [LocationRouteCalculator].
func (lrc *LocationRouteCalculator) Type() string {
	return "aws_location_route_calculator"
}

// LocalName returns the local name for [LocationRouteCalculator].
func (lrc *LocationRouteCalculator) LocalName() string {
	return lrc.Name
}

// Configuration returns the configuration (args) for [LocationRouteCalculator].
func (lrc *LocationRouteCalculator) Configuration() interface{} {
	return lrc.Args
}

// DependOn is used for other resources to depend on [LocationRouteCalculator].
func (lrc *LocationRouteCalculator) DependOn() terra.Reference {
	return terra.ReferenceResource(lrc)
}

// Dependencies returns the list of resources [LocationRouteCalculator] depends_on.
func (lrc *LocationRouteCalculator) Dependencies() terra.Dependencies {
	return lrc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [LocationRouteCalculator].
func (lrc *LocationRouteCalculator) LifecycleManagement() *terra.Lifecycle {
	return lrc.Lifecycle
}

// Attributes returns the attributes for [LocationRouteCalculator].
func (lrc *LocationRouteCalculator) Attributes() locationRouteCalculatorAttributes {
	return locationRouteCalculatorAttributes{ref: terra.ReferenceResource(lrc)}
}

// ImportState imports the given attribute values into [LocationRouteCalculator]'s state.
func (lrc *LocationRouteCalculator) ImportState(av io.Reader) error {
	lrc.state = &locationRouteCalculatorState{}
	if err := json.NewDecoder(av).Decode(lrc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", lrc.Type(), lrc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [LocationRouteCalculator] has state.
func (lrc *LocationRouteCalculator) State() (*locationRouteCalculatorState, bool) {
	return lrc.state, lrc.state != nil
}

// StateMust returns the state for [LocationRouteCalculator]. Panics if the state is nil.
func (lrc *LocationRouteCalculator) StateMust() *locationRouteCalculatorState {
	if lrc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", lrc.Type(), lrc.LocalName()))
	}
	return lrc.state
}

// LocationRouteCalculatorArgs contains the configurations for aws_location_route_calculator.
type LocationRouteCalculatorArgs struct {
	// CalculatorName: string, required
	CalculatorName terra.StringValue `hcl:"calculator_name,attr" validate:"required"`
	// DataSource: string, required
	DataSource terra.StringValue `hcl:"data_source,attr" validate:"required"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// Timeouts: optional
	Timeouts *locationroutecalculator.Timeouts `hcl:"timeouts,block"`
}
type locationRouteCalculatorAttributes struct {
	ref terra.Reference
}

// CalculatorArn returns a reference to field calculator_arn of aws_location_route_calculator.
func (lrc locationRouteCalculatorAttributes) CalculatorArn() terra.StringValue {
	return terra.ReferenceAsString(lrc.ref.Append("calculator_arn"))
}

// CalculatorName returns a reference to field calculator_name of aws_location_route_calculator.
func (lrc locationRouteCalculatorAttributes) CalculatorName() terra.StringValue {
	return terra.ReferenceAsString(lrc.ref.Append("calculator_name"))
}

// CreateTime returns a reference to field create_time of aws_location_route_calculator.
func (lrc locationRouteCalculatorAttributes) CreateTime() terra.StringValue {
	return terra.ReferenceAsString(lrc.ref.Append("create_time"))
}

// DataSource returns a reference to field data_source of aws_location_route_calculator.
func (lrc locationRouteCalculatorAttributes) DataSource() terra.StringValue {
	return terra.ReferenceAsString(lrc.ref.Append("data_source"))
}

// Description returns a reference to field description of aws_location_route_calculator.
func (lrc locationRouteCalculatorAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(lrc.ref.Append("description"))
}

// Id returns a reference to field id of aws_location_route_calculator.
func (lrc locationRouteCalculatorAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(lrc.ref.Append("id"))
}

// Tags returns a reference to field tags of aws_location_route_calculator.
func (lrc locationRouteCalculatorAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](lrc.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_location_route_calculator.
func (lrc locationRouteCalculatorAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](lrc.ref.Append("tags_all"))
}

// UpdateTime returns a reference to field update_time of aws_location_route_calculator.
func (lrc locationRouteCalculatorAttributes) UpdateTime() terra.StringValue {
	return terra.ReferenceAsString(lrc.ref.Append("update_time"))
}

func (lrc locationRouteCalculatorAttributes) Timeouts() locationroutecalculator.TimeoutsAttributes {
	return terra.ReferenceAsSingle[locationroutecalculator.TimeoutsAttributes](lrc.ref.Append("timeouts"))
}

type locationRouteCalculatorState struct {
	CalculatorArn  string                                 `json:"calculator_arn"`
	CalculatorName string                                 `json:"calculator_name"`
	CreateTime     string                                 `json:"create_time"`
	DataSource     string                                 `json:"data_source"`
	Description    string                                 `json:"description"`
	Id             string                                 `json:"id"`
	Tags           map[string]string                      `json:"tags"`
	TagsAll        map[string]string                      `json:"tags_all"`
	UpdateTime     string                                 `json:"update_time"`
	Timeouts       *locationroutecalculator.TimeoutsState `json:"timeouts"`
}