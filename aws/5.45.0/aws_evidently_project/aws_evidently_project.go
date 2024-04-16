// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_evidently_project

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

// Resource represents the Terraform resource aws_evidently_project.
type Resource struct {
	Name      string
	Args      Args
	state     *awsEvidentlyProjectState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (aep *Resource) Type() string {
	return "aws_evidently_project"
}

// LocalName returns the local name for [Resource].
func (aep *Resource) LocalName() string {
	return aep.Name
}

// Configuration returns the configuration (args) for [Resource].
func (aep *Resource) Configuration() interface{} {
	return aep.Args
}

// DependOn is used for other resources to depend on [Resource].
func (aep *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(aep)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (aep *Resource) Dependencies() terra.Dependencies {
	return aep.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (aep *Resource) LifecycleManagement() *terra.Lifecycle {
	return aep.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (aep *Resource) Attributes() awsEvidentlyProjectAttributes {
	return awsEvidentlyProjectAttributes{ref: terra.ReferenceResource(aep)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (aep *Resource) ImportState(state io.Reader) error {
	aep.state = &awsEvidentlyProjectState{}
	if err := json.NewDecoder(state).Decode(aep.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", aep.Type(), aep.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (aep *Resource) State() (*awsEvidentlyProjectState, bool) {
	return aep.state, aep.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (aep *Resource) StateMust() *awsEvidentlyProjectState {
	if aep.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", aep.Type(), aep.LocalName()))
	}
	return aep.state
}

// Args contains the configurations for aws_evidently_project.
type Args struct {
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
	// DataDelivery: optional
	DataDelivery *DataDelivery `hcl:"data_delivery,block"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type awsEvidentlyProjectAttributes struct {
	ref terra.Reference
}

// ActiveExperimentCount returns a reference to field active_experiment_count of aws_evidently_project.
func (aep awsEvidentlyProjectAttributes) ActiveExperimentCount() terra.NumberValue {
	return terra.ReferenceAsNumber(aep.ref.Append("active_experiment_count"))
}

// ActiveLaunchCount returns a reference to field active_launch_count of aws_evidently_project.
func (aep awsEvidentlyProjectAttributes) ActiveLaunchCount() terra.NumberValue {
	return terra.ReferenceAsNumber(aep.ref.Append("active_launch_count"))
}

// Arn returns a reference to field arn of aws_evidently_project.
func (aep awsEvidentlyProjectAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(aep.ref.Append("arn"))
}

// CreatedTime returns a reference to field created_time of aws_evidently_project.
func (aep awsEvidentlyProjectAttributes) CreatedTime() terra.StringValue {
	return terra.ReferenceAsString(aep.ref.Append("created_time"))
}

// Description returns a reference to field description of aws_evidently_project.
func (aep awsEvidentlyProjectAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(aep.ref.Append("description"))
}

// ExperimentCount returns a reference to field experiment_count of aws_evidently_project.
func (aep awsEvidentlyProjectAttributes) ExperimentCount() terra.NumberValue {
	return terra.ReferenceAsNumber(aep.ref.Append("experiment_count"))
}

// FeatureCount returns a reference to field feature_count of aws_evidently_project.
func (aep awsEvidentlyProjectAttributes) FeatureCount() terra.NumberValue {
	return terra.ReferenceAsNumber(aep.ref.Append("feature_count"))
}

// Id returns a reference to field id of aws_evidently_project.
func (aep awsEvidentlyProjectAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(aep.ref.Append("id"))
}

// LastUpdatedTime returns a reference to field last_updated_time of aws_evidently_project.
func (aep awsEvidentlyProjectAttributes) LastUpdatedTime() terra.StringValue {
	return terra.ReferenceAsString(aep.ref.Append("last_updated_time"))
}

// LaunchCount returns a reference to field launch_count of aws_evidently_project.
func (aep awsEvidentlyProjectAttributes) LaunchCount() terra.NumberValue {
	return terra.ReferenceAsNumber(aep.ref.Append("launch_count"))
}

// Name returns a reference to field name of aws_evidently_project.
func (aep awsEvidentlyProjectAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(aep.ref.Append("name"))
}

// Status returns a reference to field status of aws_evidently_project.
func (aep awsEvidentlyProjectAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(aep.ref.Append("status"))
}

// Tags returns a reference to field tags of aws_evidently_project.
func (aep awsEvidentlyProjectAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](aep.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_evidently_project.
func (aep awsEvidentlyProjectAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](aep.ref.Append("tags_all"))
}

func (aep awsEvidentlyProjectAttributes) DataDelivery() terra.ListValue[DataDeliveryAttributes] {
	return terra.ReferenceAsList[DataDeliveryAttributes](aep.ref.Append("data_delivery"))
}

func (aep awsEvidentlyProjectAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](aep.ref.Append("timeouts"))
}

type awsEvidentlyProjectState struct {
	ActiveExperimentCount float64             `json:"active_experiment_count"`
	ActiveLaunchCount     float64             `json:"active_launch_count"`
	Arn                   string              `json:"arn"`
	CreatedTime           string              `json:"created_time"`
	Description           string              `json:"description"`
	ExperimentCount       float64             `json:"experiment_count"`
	FeatureCount          float64             `json:"feature_count"`
	Id                    string              `json:"id"`
	LastUpdatedTime       string              `json:"last_updated_time"`
	LaunchCount           float64             `json:"launch_count"`
	Name                  string              `json:"name"`
	Status                string              `json:"status"`
	Tags                  map[string]string   `json:"tags"`
	TagsAll               map[string]string   `json:"tags_all"`
	DataDelivery          []DataDeliveryState `json:"data_delivery"`
	Timeouts              *TimeoutsState      `json:"timeouts"`
}
