// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	computeautoscaler "github.com/golingon/terraproviders/google/4.72.1/computeautoscaler"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewComputeAutoscaler creates a new instance of [ComputeAutoscaler].
func NewComputeAutoscaler(name string, args ComputeAutoscalerArgs) *ComputeAutoscaler {
	return &ComputeAutoscaler{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ComputeAutoscaler)(nil)

// ComputeAutoscaler represents the Terraform resource google_compute_autoscaler.
type ComputeAutoscaler struct {
	Name      string
	Args      ComputeAutoscalerArgs
	state     *computeAutoscalerState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ComputeAutoscaler].
func (ca *ComputeAutoscaler) Type() string {
	return "google_compute_autoscaler"
}

// LocalName returns the local name for [ComputeAutoscaler].
func (ca *ComputeAutoscaler) LocalName() string {
	return ca.Name
}

// Configuration returns the configuration (args) for [ComputeAutoscaler].
func (ca *ComputeAutoscaler) Configuration() interface{} {
	return ca.Args
}

// DependOn is used for other resources to depend on [ComputeAutoscaler].
func (ca *ComputeAutoscaler) DependOn() terra.Reference {
	return terra.ReferenceResource(ca)
}

// Dependencies returns the list of resources [ComputeAutoscaler] depends_on.
func (ca *ComputeAutoscaler) Dependencies() terra.Dependencies {
	return ca.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ComputeAutoscaler].
func (ca *ComputeAutoscaler) LifecycleManagement() *terra.Lifecycle {
	return ca.Lifecycle
}

// Attributes returns the attributes for [ComputeAutoscaler].
func (ca *ComputeAutoscaler) Attributes() computeAutoscalerAttributes {
	return computeAutoscalerAttributes{ref: terra.ReferenceResource(ca)}
}

// ImportState imports the given attribute values into [ComputeAutoscaler]'s state.
func (ca *ComputeAutoscaler) ImportState(av io.Reader) error {
	ca.state = &computeAutoscalerState{}
	if err := json.NewDecoder(av).Decode(ca.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ca.Type(), ca.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ComputeAutoscaler] has state.
func (ca *ComputeAutoscaler) State() (*computeAutoscalerState, bool) {
	return ca.state, ca.state != nil
}

// StateMust returns the state for [ComputeAutoscaler]. Panics if the state is nil.
func (ca *ComputeAutoscaler) StateMust() *computeAutoscalerState {
	if ca.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ca.Type(), ca.LocalName()))
	}
	return ca.state
}

// ComputeAutoscalerArgs contains the configurations for google_compute_autoscaler.
type ComputeAutoscalerArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Target: string, required
	Target terra.StringValue `hcl:"target,attr" validate:"required"`
	// Zone: string, optional
	Zone terra.StringValue `hcl:"zone,attr"`
	// AutoscalingPolicy: required
	AutoscalingPolicy *computeautoscaler.AutoscalingPolicy `hcl:"autoscaling_policy,block" validate:"required"`
	// Timeouts: optional
	Timeouts *computeautoscaler.Timeouts `hcl:"timeouts,block"`
}
type computeAutoscalerAttributes struct {
	ref terra.Reference
}

// CreationTimestamp returns a reference to field creation_timestamp of google_compute_autoscaler.
func (ca computeAutoscalerAttributes) CreationTimestamp() terra.StringValue {
	return terra.ReferenceAsString(ca.ref.Append("creation_timestamp"))
}

// Description returns a reference to field description of google_compute_autoscaler.
func (ca computeAutoscalerAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(ca.ref.Append("description"))
}

// Id returns a reference to field id of google_compute_autoscaler.
func (ca computeAutoscalerAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ca.ref.Append("id"))
}

// Name returns a reference to field name of google_compute_autoscaler.
func (ca computeAutoscalerAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ca.ref.Append("name"))
}

// Project returns a reference to field project of google_compute_autoscaler.
func (ca computeAutoscalerAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(ca.ref.Append("project"))
}

// SelfLink returns a reference to field self_link of google_compute_autoscaler.
func (ca computeAutoscalerAttributes) SelfLink() terra.StringValue {
	return terra.ReferenceAsString(ca.ref.Append("self_link"))
}

// Target returns a reference to field target of google_compute_autoscaler.
func (ca computeAutoscalerAttributes) Target() terra.StringValue {
	return terra.ReferenceAsString(ca.ref.Append("target"))
}

// Zone returns a reference to field zone of google_compute_autoscaler.
func (ca computeAutoscalerAttributes) Zone() terra.StringValue {
	return terra.ReferenceAsString(ca.ref.Append("zone"))
}

func (ca computeAutoscalerAttributes) AutoscalingPolicy() terra.ListValue[computeautoscaler.AutoscalingPolicyAttributes] {
	return terra.ReferenceAsList[computeautoscaler.AutoscalingPolicyAttributes](ca.ref.Append("autoscaling_policy"))
}

func (ca computeAutoscalerAttributes) Timeouts() computeautoscaler.TimeoutsAttributes {
	return terra.ReferenceAsSingle[computeautoscaler.TimeoutsAttributes](ca.ref.Append("timeouts"))
}

type computeAutoscalerState struct {
	CreationTimestamp string                                     `json:"creation_timestamp"`
	Description       string                                     `json:"description"`
	Id                string                                     `json:"id"`
	Name              string                                     `json:"name"`
	Project           string                                     `json:"project"`
	SelfLink          string                                     `json:"self_link"`
	Target            string                                     `json:"target"`
	Zone              string                                     `json:"zone"`
	AutoscalingPolicy []computeautoscaler.AutoscalingPolicyState `json:"autoscaling_policy"`
	Timeouts          *computeautoscaler.TimeoutsState           `json:"timeouts"`
}
