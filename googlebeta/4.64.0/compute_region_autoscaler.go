// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	computeregionautoscaler "github.com/golingon/terraproviders/googlebeta/4.64.0/computeregionautoscaler"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewComputeRegionAutoscaler creates a new instance of [ComputeRegionAutoscaler].
func NewComputeRegionAutoscaler(name string, args ComputeRegionAutoscalerArgs) *ComputeRegionAutoscaler {
	return &ComputeRegionAutoscaler{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ComputeRegionAutoscaler)(nil)

// ComputeRegionAutoscaler represents the Terraform resource google_compute_region_autoscaler.
type ComputeRegionAutoscaler struct {
	Name      string
	Args      ComputeRegionAutoscalerArgs
	state     *computeRegionAutoscalerState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ComputeRegionAutoscaler].
func (cra *ComputeRegionAutoscaler) Type() string {
	return "google_compute_region_autoscaler"
}

// LocalName returns the local name for [ComputeRegionAutoscaler].
func (cra *ComputeRegionAutoscaler) LocalName() string {
	return cra.Name
}

// Configuration returns the configuration (args) for [ComputeRegionAutoscaler].
func (cra *ComputeRegionAutoscaler) Configuration() interface{} {
	return cra.Args
}

// DependOn is used for other resources to depend on [ComputeRegionAutoscaler].
func (cra *ComputeRegionAutoscaler) DependOn() terra.Reference {
	return terra.ReferenceResource(cra)
}

// Dependencies returns the list of resources [ComputeRegionAutoscaler] depends_on.
func (cra *ComputeRegionAutoscaler) Dependencies() terra.Dependencies {
	return cra.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ComputeRegionAutoscaler].
func (cra *ComputeRegionAutoscaler) LifecycleManagement() *terra.Lifecycle {
	return cra.Lifecycle
}

// Attributes returns the attributes for [ComputeRegionAutoscaler].
func (cra *ComputeRegionAutoscaler) Attributes() computeRegionAutoscalerAttributes {
	return computeRegionAutoscalerAttributes{ref: terra.ReferenceResource(cra)}
}

// ImportState imports the given attribute values into [ComputeRegionAutoscaler]'s state.
func (cra *ComputeRegionAutoscaler) ImportState(av io.Reader) error {
	cra.state = &computeRegionAutoscalerState{}
	if err := json.NewDecoder(av).Decode(cra.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cra.Type(), cra.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ComputeRegionAutoscaler] has state.
func (cra *ComputeRegionAutoscaler) State() (*computeRegionAutoscalerState, bool) {
	return cra.state, cra.state != nil
}

// StateMust returns the state for [ComputeRegionAutoscaler]. Panics if the state is nil.
func (cra *ComputeRegionAutoscaler) StateMust() *computeRegionAutoscalerState {
	if cra.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cra.Type(), cra.LocalName()))
	}
	return cra.state
}

// ComputeRegionAutoscalerArgs contains the configurations for google_compute_region_autoscaler.
type ComputeRegionAutoscalerArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Region: string, optional
	Region terra.StringValue `hcl:"region,attr"`
	// Target: string, required
	Target terra.StringValue `hcl:"target,attr" validate:"required"`
	// AutoscalingPolicy: required
	AutoscalingPolicy *computeregionautoscaler.AutoscalingPolicy `hcl:"autoscaling_policy,block" validate:"required"`
	// Timeouts: optional
	Timeouts *computeregionautoscaler.Timeouts `hcl:"timeouts,block"`
}
type computeRegionAutoscalerAttributes struct {
	ref terra.Reference
}

// CreationTimestamp returns a reference to field creation_timestamp of google_compute_region_autoscaler.
func (cra computeRegionAutoscalerAttributes) CreationTimestamp() terra.StringValue {
	return terra.ReferenceAsString(cra.ref.Append("creation_timestamp"))
}

// Description returns a reference to field description of google_compute_region_autoscaler.
func (cra computeRegionAutoscalerAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(cra.ref.Append("description"))
}

// Id returns a reference to field id of google_compute_region_autoscaler.
func (cra computeRegionAutoscalerAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cra.ref.Append("id"))
}

// Name returns a reference to field name of google_compute_region_autoscaler.
func (cra computeRegionAutoscalerAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cra.ref.Append("name"))
}

// Project returns a reference to field project of google_compute_region_autoscaler.
func (cra computeRegionAutoscalerAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(cra.ref.Append("project"))
}

// Region returns a reference to field region of google_compute_region_autoscaler.
func (cra computeRegionAutoscalerAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(cra.ref.Append("region"))
}

// SelfLink returns a reference to field self_link of google_compute_region_autoscaler.
func (cra computeRegionAutoscalerAttributes) SelfLink() terra.StringValue {
	return terra.ReferenceAsString(cra.ref.Append("self_link"))
}

// Target returns a reference to field target of google_compute_region_autoscaler.
func (cra computeRegionAutoscalerAttributes) Target() terra.StringValue {
	return terra.ReferenceAsString(cra.ref.Append("target"))
}

func (cra computeRegionAutoscalerAttributes) AutoscalingPolicy() terra.ListValue[computeregionautoscaler.AutoscalingPolicyAttributes] {
	return terra.ReferenceAsList[computeregionautoscaler.AutoscalingPolicyAttributes](cra.ref.Append("autoscaling_policy"))
}

func (cra computeRegionAutoscalerAttributes) Timeouts() computeregionautoscaler.TimeoutsAttributes {
	return terra.ReferenceAsSingle[computeregionautoscaler.TimeoutsAttributes](cra.ref.Append("timeouts"))
}

type computeRegionAutoscalerState struct {
	CreationTimestamp string                                           `json:"creation_timestamp"`
	Description       string                                           `json:"description"`
	Id                string                                           `json:"id"`
	Name              string                                           `json:"name"`
	Project           string                                           `json:"project"`
	Region            string                                           `json:"region"`
	SelfLink          string                                           `json:"self_link"`
	Target            string                                           `json:"target"`
	AutoscalingPolicy []computeregionautoscaler.AutoscalingPolicyState `json:"autoscaling_policy"`
	Timeouts          *computeregionautoscaler.TimeoutsState           `json:"timeouts"`
}
