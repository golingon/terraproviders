// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	monitoringmetricdescriptor "github.com/golingon/terraproviders/google/4.66.0/monitoringmetricdescriptor"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewMonitoringMetricDescriptor creates a new instance of [MonitoringMetricDescriptor].
func NewMonitoringMetricDescriptor(name string, args MonitoringMetricDescriptorArgs) *MonitoringMetricDescriptor {
	return &MonitoringMetricDescriptor{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*MonitoringMetricDescriptor)(nil)

// MonitoringMetricDescriptor represents the Terraform resource google_monitoring_metric_descriptor.
type MonitoringMetricDescriptor struct {
	Name      string
	Args      MonitoringMetricDescriptorArgs
	state     *monitoringMetricDescriptorState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [MonitoringMetricDescriptor].
func (mmd *MonitoringMetricDescriptor) Type() string {
	return "google_monitoring_metric_descriptor"
}

// LocalName returns the local name for [MonitoringMetricDescriptor].
func (mmd *MonitoringMetricDescriptor) LocalName() string {
	return mmd.Name
}

// Configuration returns the configuration (args) for [MonitoringMetricDescriptor].
func (mmd *MonitoringMetricDescriptor) Configuration() interface{} {
	return mmd.Args
}

// DependOn is used for other resources to depend on [MonitoringMetricDescriptor].
func (mmd *MonitoringMetricDescriptor) DependOn() terra.Reference {
	return terra.ReferenceResource(mmd)
}

// Dependencies returns the list of resources [MonitoringMetricDescriptor] depends_on.
func (mmd *MonitoringMetricDescriptor) Dependencies() terra.Dependencies {
	return mmd.DependsOn
}

// LifecycleManagement returns the lifecycle block for [MonitoringMetricDescriptor].
func (mmd *MonitoringMetricDescriptor) LifecycleManagement() *terra.Lifecycle {
	return mmd.Lifecycle
}

// Attributes returns the attributes for [MonitoringMetricDescriptor].
func (mmd *MonitoringMetricDescriptor) Attributes() monitoringMetricDescriptorAttributes {
	return monitoringMetricDescriptorAttributes{ref: terra.ReferenceResource(mmd)}
}

// ImportState imports the given attribute values into [MonitoringMetricDescriptor]'s state.
func (mmd *MonitoringMetricDescriptor) ImportState(av io.Reader) error {
	mmd.state = &monitoringMetricDescriptorState{}
	if err := json.NewDecoder(av).Decode(mmd.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", mmd.Type(), mmd.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [MonitoringMetricDescriptor] has state.
func (mmd *MonitoringMetricDescriptor) State() (*monitoringMetricDescriptorState, bool) {
	return mmd.state, mmd.state != nil
}

// StateMust returns the state for [MonitoringMetricDescriptor]. Panics if the state is nil.
func (mmd *MonitoringMetricDescriptor) StateMust() *monitoringMetricDescriptorState {
	if mmd.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", mmd.Type(), mmd.LocalName()))
	}
	return mmd.state
}

// MonitoringMetricDescriptorArgs contains the configurations for google_monitoring_metric_descriptor.
type MonitoringMetricDescriptorArgs struct {
	// Description: string, required
	Description terra.StringValue `hcl:"description,attr" validate:"required"`
	// DisplayName: string, required
	DisplayName terra.StringValue `hcl:"display_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// LaunchStage: string, optional
	LaunchStage terra.StringValue `hcl:"launch_stage,attr"`
	// MetricKind: string, required
	MetricKind terra.StringValue `hcl:"metric_kind,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Type: string, required
	Type terra.StringValue `hcl:"type,attr" validate:"required"`
	// Unit: string, optional
	Unit terra.StringValue `hcl:"unit,attr"`
	// ValueType: string, required
	ValueType terra.StringValue `hcl:"value_type,attr" validate:"required"`
	// Labels: min=0
	Labels []monitoringmetricdescriptor.Labels `hcl:"labels,block" validate:"min=0"`
	// Metadata: optional
	Metadata *monitoringmetricdescriptor.Metadata `hcl:"metadata,block"`
	// Timeouts: optional
	Timeouts *monitoringmetricdescriptor.Timeouts `hcl:"timeouts,block"`
}
type monitoringMetricDescriptorAttributes struct {
	ref terra.Reference
}

// Description returns a reference to field description of google_monitoring_metric_descriptor.
func (mmd monitoringMetricDescriptorAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(mmd.ref.Append("description"))
}

// DisplayName returns a reference to field display_name of google_monitoring_metric_descriptor.
func (mmd monitoringMetricDescriptorAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(mmd.ref.Append("display_name"))
}

// Id returns a reference to field id of google_monitoring_metric_descriptor.
func (mmd monitoringMetricDescriptorAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(mmd.ref.Append("id"))
}

// LaunchStage returns a reference to field launch_stage of google_monitoring_metric_descriptor.
func (mmd monitoringMetricDescriptorAttributes) LaunchStage() terra.StringValue {
	return terra.ReferenceAsString(mmd.ref.Append("launch_stage"))
}

// MetricKind returns a reference to field metric_kind of google_monitoring_metric_descriptor.
func (mmd monitoringMetricDescriptorAttributes) MetricKind() terra.StringValue {
	return terra.ReferenceAsString(mmd.ref.Append("metric_kind"))
}

// MonitoredResourceTypes returns a reference to field monitored_resource_types of google_monitoring_metric_descriptor.
func (mmd monitoringMetricDescriptorAttributes) MonitoredResourceTypes() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](mmd.ref.Append("monitored_resource_types"))
}

// Name returns a reference to field name of google_monitoring_metric_descriptor.
func (mmd monitoringMetricDescriptorAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(mmd.ref.Append("name"))
}

// Project returns a reference to field project of google_monitoring_metric_descriptor.
func (mmd monitoringMetricDescriptorAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(mmd.ref.Append("project"))
}

// Type returns a reference to field type of google_monitoring_metric_descriptor.
func (mmd monitoringMetricDescriptorAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(mmd.ref.Append("type"))
}

// Unit returns a reference to field unit of google_monitoring_metric_descriptor.
func (mmd monitoringMetricDescriptorAttributes) Unit() terra.StringValue {
	return terra.ReferenceAsString(mmd.ref.Append("unit"))
}

// ValueType returns a reference to field value_type of google_monitoring_metric_descriptor.
func (mmd monitoringMetricDescriptorAttributes) ValueType() terra.StringValue {
	return terra.ReferenceAsString(mmd.ref.Append("value_type"))
}

func (mmd monitoringMetricDescriptorAttributes) Labels() terra.SetValue[monitoringmetricdescriptor.LabelsAttributes] {
	return terra.ReferenceAsSet[monitoringmetricdescriptor.LabelsAttributes](mmd.ref.Append("labels"))
}

func (mmd monitoringMetricDescriptorAttributes) Metadata() terra.ListValue[monitoringmetricdescriptor.MetadataAttributes] {
	return terra.ReferenceAsList[monitoringmetricdescriptor.MetadataAttributes](mmd.ref.Append("metadata"))
}

func (mmd monitoringMetricDescriptorAttributes) Timeouts() monitoringmetricdescriptor.TimeoutsAttributes {
	return terra.ReferenceAsSingle[monitoringmetricdescriptor.TimeoutsAttributes](mmd.ref.Append("timeouts"))
}

type monitoringMetricDescriptorState struct {
	Description            string                                     `json:"description"`
	DisplayName            string                                     `json:"display_name"`
	Id                     string                                     `json:"id"`
	LaunchStage            string                                     `json:"launch_stage"`
	MetricKind             string                                     `json:"metric_kind"`
	MonitoredResourceTypes []string                                   `json:"monitored_resource_types"`
	Name                   string                                     `json:"name"`
	Project                string                                     `json:"project"`
	Type                   string                                     `json:"type"`
	Unit                   string                                     `json:"unit"`
	ValueType              string                                     `json:"value_type"`
	Labels                 []monitoringmetricdescriptor.LabelsState   `json:"labels"`
	Metadata               []monitoringmetricdescriptor.MetadataState `json:"metadata"`
	Timeouts               *monitoringmetricdescriptor.TimeoutsState  `json:"timeouts"`
}
