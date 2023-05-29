// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	evidentlyproject "github.com/golingon/terraproviders/aws/5.0.1/evidentlyproject"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewEvidentlyProject creates a new instance of [EvidentlyProject].
func NewEvidentlyProject(name string, args EvidentlyProjectArgs) *EvidentlyProject {
	return &EvidentlyProject{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*EvidentlyProject)(nil)

// EvidentlyProject represents the Terraform resource aws_evidently_project.
type EvidentlyProject struct {
	Name      string
	Args      EvidentlyProjectArgs
	state     *evidentlyProjectState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [EvidentlyProject].
func (ep *EvidentlyProject) Type() string {
	return "aws_evidently_project"
}

// LocalName returns the local name for [EvidentlyProject].
func (ep *EvidentlyProject) LocalName() string {
	return ep.Name
}

// Configuration returns the configuration (args) for [EvidentlyProject].
func (ep *EvidentlyProject) Configuration() interface{} {
	return ep.Args
}

// DependOn is used for other resources to depend on [EvidentlyProject].
func (ep *EvidentlyProject) DependOn() terra.Reference {
	return terra.ReferenceResource(ep)
}

// Dependencies returns the list of resources [EvidentlyProject] depends_on.
func (ep *EvidentlyProject) Dependencies() terra.Dependencies {
	return ep.DependsOn
}

// LifecycleManagement returns the lifecycle block for [EvidentlyProject].
func (ep *EvidentlyProject) LifecycleManagement() *terra.Lifecycle {
	return ep.Lifecycle
}

// Attributes returns the attributes for [EvidentlyProject].
func (ep *EvidentlyProject) Attributes() evidentlyProjectAttributes {
	return evidentlyProjectAttributes{ref: terra.ReferenceResource(ep)}
}

// ImportState imports the given attribute values into [EvidentlyProject]'s state.
func (ep *EvidentlyProject) ImportState(av io.Reader) error {
	ep.state = &evidentlyProjectState{}
	if err := json.NewDecoder(av).Decode(ep.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ep.Type(), ep.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [EvidentlyProject] has state.
func (ep *EvidentlyProject) State() (*evidentlyProjectState, bool) {
	return ep.state, ep.state != nil
}

// StateMust returns the state for [EvidentlyProject]. Panics if the state is nil.
func (ep *EvidentlyProject) StateMust() *evidentlyProjectState {
	if ep.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ep.Type(), ep.LocalName()))
	}
	return ep.state
}

// EvidentlyProjectArgs contains the configurations for aws_evidently_project.
type EvidentlyProjectArgs struct {
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
	DataDelivery *evidentlyproject.DataDelivery `hcl:"data_delivery,block"`
	// Timeouts: optional
	Timeouts *evidentlyproject.Timeouts `hcl:"timeouts,block"`
}
type evidentlyProjectAttributes struct {
	ref terra.Reference
}

// ActiveExperimentCount returns a reference to field active_experiment_count of aws_evidently_project.
func (ep evidentlyProjectAttributes) ActiveExperimentCount() terra.NumberValue {
	return terra.ReferenceAsNumber(ep.ref.Append("active_experiment_count"))
}

// ActiveLaunchCount returns a reference to field active_launch_count of aws_evidently_project.
func (ep evidentlyProjectAttributes) ActiveLaunchCount() terra.NumberValue {
	return terra.ReferenceAsNumber(ep.ref.Append("active_launch_count"))
}

// Arn returns a reference to field arn of aws_evidently_project.
func (ep evidentlyProjectAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(ep.ref.Append("arn"))
}

// CreatedTime returns a reference to field created_time of aws_evidently_project.
func (ep evidentlyProjectAttributes) CreatedTime() terra.StringValue {
	return terra.ReferenceAsString(ep.ref.Append("created_time"))
}

// Description returns a reference to field description of aws_evidently_project.
func (ep evidentlyProjectAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(ep.ref.Append("description"))
}

// ExperimentCount returns a reference to field experiment_count of aws_evidently_project.
func (ep evidentlyProjectAttributes) ExperimentCount() terra.NumberValue {
	return terra.ReferenceAsNumber(ep.ref.Append("experiment_count"))
}

// FeatureCount returns a reference to field feature_count of aws_evidently_project.
func (ep evidentlyProjectAttributes) FeatureCount() terra.NumberValue {
	return terra.ReferenceAsNumber(ep.ref.Append("feature_count"))
}

// Id returns a reference to field id of aws_evidently_project.
func (ep evidentlyProjectAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ep.ref.Append("id"))
}

// LastUpdatedTime returns a reference to field last_updated_time of aws_evidently_project.
func (ep evidentlyProjectAttributes) LastUpdatedTime() terra.StringValue {
	return terra.ReferenceAsString(ep.ref.Append("last_updated_time"))
}

// LaunchCount returns a reference to field launch_count of aws_evidently_project.
func (ep evidentlyProjectAttributes) LaunchCount() terra.NumberValue {
	return terra.ReferenceAsNumber(ep.ref.Append("launch_count"))
}

// Name returns a reference to field name of aws_evidently_project.
func (ep evidentlyProjectAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ep.ref.Append("name"))
}

// Status returns a reference to field status of aws_evidently_project.
func (ep evidentlyProjectAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(ep.ref.Append("status"))
}

// Tags returns a reference to field tags of aws_evidently_project.
func (ep evidentlyProjectAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ep.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_evidently_project.
func (ep evidentlyProjectAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ep.ref.Append("tags_all"))
}

func (ep evidentlyProjectAttributes) DataDelivery() terra.ListValue[evidentlyproject.DataDeliveryAttributes] {
	return terra.ReferenceAsList[evidentlyproject.DataDeliveryAttributes](ep.ref.Append("data_delivery"))
}

func (ep evidentlyProjectAttributes) Timeouts() evidentlyproject.TimeoutsAttributes {
	return terra.ReferenceAsSingle[evidentlyproject.TimeoutsAttributes](ep.ref.Append("timeouts"))
}

type evidentlyProjectState struct {
	ActiveExperimentCount float64                              `json:"active_experiment_count"`
	ActiveLaunchCount     float64                              `json:"active_launch_count"`
	Arn                   string                               `json:"arn"`
	CreatedTime           string                               `json:"created_time"`
	Description           string                               `json:"description"`
	ExperimentCount       float64                              `json:"experiment_count"`
	FeatureCount          float64                              `json:"feature_count"`
	Id                    string                               `json:"id"`
	LastUpdatedTime       string                               `json:"last_updated_time"`
	LaunchCount           float64                              `json:"launch_count"`
	Name                  string                               `json:"name"`
	Status                string                               `json:"status"`
	Tags                  map[string]string                    `json:"tags"`
	TagsAll               map[string]string                    `json:"tags_all"`
	DataDelivery          []evidentlyproject.DataDeliveryState `json:"data_delivery"`
	Timeouts              *evidentlyproject.TimeoutsState      `json:"timeouts"`
}