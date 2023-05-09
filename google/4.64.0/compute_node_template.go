// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	computenodetemplate "github.com/golingon/terraproviders/google/4.64.0/computenodetemplate"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewComputeNodeTemplate creates a new instance of [ComputeNodeTemplate].
func NewComputeNodeTemplate(name string, args ComputeNodeTemplateArgs) *ComputeNodeTemplate {
	return &ComputeNodeTemplate{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ComputeNodeTemplate)(nil)

// ComputeNodeTemplate represents the Terraform resource google_compute_node_template.
type ComputeNodeTemplate struct {
	Name      string
	Args      ComputeNodeTemplateArgs
	state     *computeNodeTemplateState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ComputeNodeTemplate].
func (cnt *ComputeNodeTemplate) Type() string {
	return "google_compute_node_template"
}

// LocalName returns the local name for [ComputeNodeTemplate].
func (cnt *ComputeNodeTemplate) LocalName() string {
	return cnt.Name
}

// Configuration returns the configuration (args) for [ComputeNodeTemplate].
func (cnt *ComputeNodeTemplate) Configuration() interface{} {
	return cnt.Args
}

// DependOn is used for other resources to depend on [ComputeNodeTemplate].
func (cnt *ComputeNodeTemplate) DependOn() terra.Reference {
	return terra.ReferenceResource(cnt)
}

// Dependencies returns the list of resources [ComputeNodeTemplate] depends_on.
func (cnt *ComputeNodeTemplate) Dependencies() terra.Dependencies {
	return cnt.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ComputeNodeTemplate].
func (cnt *ComputeNodeTemplate) LifecycleManagement() *terra.Lifecycle {
	return cnt.Lifecycle
}

// Attributes returns the attributes for [ComputeNodeTemplate].
func (cnt *ComputeNodeTemplate) Attributes() computeNodeTemplateAttributes {
	return computeNodeTemplateAttributes{ref: terra.ReferenceResource(cnt)}
}

// ImportState imports the given attribute values into [ComputeNodeTemplate]'s state.
func (cnt *ComputeNodeTemplate) ImportState(av io.Reader) error {
	cnt.state = &computeNodeTemplateState{}
	if err := json.NewDecoder(av).Decode(cnt.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cnt.Type(), cnt.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ComputeNodeTemplate] has state.
func (cnt *ComputeNodeTemplate) State() (*computeNodeTemplateState, bool) {
	return cnt.state, cnt.state != nil
}

// StateMust returns the state for [ComputeNodeTemplate]. Panics if the state is nil.
func (cnt *ComputeNodeTemplate) StateMust() *computeNodeTemplateState {
	if cnt.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cnt.Type(), cnt.LocalName()))
	}
	return cnt.state
}

// ComputeNodeTemplateArgs contains the configurations for google_compute_node_template.
type ComputeNodeTemplateArgs struct {
	// CpuOvercommitType: string, optional
	CpuOvercommitType terra.StringValue `hcl:"cpu_overcommit_type,attr"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
	// NodeAffinityLabels: map of string, optional
	NodeAffinityLabels terra.MapValue[terra.StringValue] `hcl:"node_affinity_labels,attr"`
	// NodeType: string, optional
	NodeType terra.StringValue `hcl:"node_type,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Region: string, optional
	Region terra.StringValue `hcl:"region,attr"`
	// NodeTypeFlexibility: optional
	NodeTypeFlexibility *computenodetemplate.NodeTypeFlexibility `hcl:"node_type_flexibility,block"`
	// ServerBinding: optional
	ServerBinding *computenodetemplate.ServerBinding `hcl:"server_binding,block"`
	// Timeouts: optional
	Timeouts *computenodetemplate.Timeouts `hcl:"timeouts,block"`
}
type computeNodeTemplateAttributes struct {
	ref terra.Reference
}

// CpuOvercommitType returns a reference to field cpu_overcommit_type of google_compute_node_template.
func (cnt computeNodeTemplateAttributes) CpuOvercommitType() terra.StringValue {
	return terra.ReferenceAsString(cnt.ref.Append("cpu_overcommit_type"))
}

// CreationTimestamp returns a reference to field creation_timestamp of google_compute_node_template.
func (cnt computeNodeTemplateAttributes) CreationTimestamp() terra.StringValue {
	return terra.ReferenceAsString(cnt.ref.Append("creation_timestamp"))
}

// Description returns a reference to field description of google_compute_node_template.
func (cnt computeNodeTemplateAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(cnt.ref.Append("description"))
}

// Id returns a reference to field id of google_compute_node_template.
func (cnt computeNodeTemplateAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cnt.ref.Append("id"))
}

// Name returns a reference to field name of google_compute_node_template.
func (cnt computeNodeTemplateAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cnt.ref.Append("name"))
}

// NodeAffinityLabels returns a reference to field node_affinity_labels of google_compute_node_template.
func (cnt computeNodeTemplateAttributes) NodeAffinityLabels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](cnt.ref.Append("node_affinity_labels"))
}

// NodeType returns a reference to field node_type of google_compute_node_template.
func (cnt computeNodeTemplateAttributes) NodeType() terra.StringValue {
	return terra.ReferenceAsString(cnt.ref.Append("node_type"))
}

// Project returns a reference to field project of google_compute_node_template.
func (cnt computeNodeTemplateAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(cnt.ref.Append("project"))
}

// Region returns a reference to field region of google_compute_node_template.
func (cnt computeNodeTemplateAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(cnt.ref.Append("region"))
}

// SelfLink returns a reference to field self_link of google_compute_node_template.
func (cnt computeNodeTemplateAttributes) SelfLink() terra.StringValue {
	return terra.ReferenceAsString(cnt.ref.Append("self_link"))
}

func (cnt computeNodeTemplateAttributes) NodeTypeFlexibility() terra.ListValue[computenodetemplate.NodeTypeFlexibilityAttributes] {
	return terra.ReferenceAsList[computenodetemplate.NodeTypeFlexibilityAttributes](cnt.ref.Append("node_type_flexibility"))
}

func (cnt computeNodeTemplateAttributes) ServerBinding() terra.ListValue[computenodetemplate.ServerBindingAttributes] {
	return terra.ReferenceAsList[computenodetemplate.ServerBindingAttributes](cnt.ref.Append("server_binding"))
}

func (cnt computeNodeTemplateAttributes) Timeouts() computenodetemplate.TimeoutsAttributes {
	return terra.ReferenceAsSingle[computenodetemplate.TimeoutsAttributes](cnt.ref.Append("timeouts"))
}

type computeNodeTemplateState struct {
	CpuOvercommitType   string                                         `json:"cpu_overcommit_type"`
	CreationTimestamp   string                                         `json:"creation_timestamp"`
	Description         string                                         `json:"description"`
	Id                  string                                         `json:"id"`
	Name                string                                         `json:"name"`
	NodeAffinityLabels  map[string]string                              `json:"node_affinity_labels"`
	NodeType            string                                         `json:"node_type"`
	Project             string                                         `json:"project"`
	Region              string                                         `json:"region"`
	SelfLink            string                                         `json:"self_link"`
	NodeTypeFlexibility []computenodetemplate.NodeTypeFlexibilityState `json:"node_type_flexibility"`
	ServerBinding       []computenodetemplate.ServerBindingState       `json:"server_binding"`
	Timeouts            *computenodetemplate.TimeoutsState             `json:"timeouts"`
}
