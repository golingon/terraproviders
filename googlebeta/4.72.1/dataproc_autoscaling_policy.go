// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	dataprocautoscalingpolicy "github.com/golingon/terraproviders/googlebeta/4.72.1/dataprocautoscalingpolicy"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDataprocAutoscalingPolicy creates a new instance of [DataprocAutoscalingPolicy].
func NewDataprocAutoscalingPolicy(name string, args DataprocAutoscalingPolicyArgs) *DataprocAutoscalingPolicy {
	return &DataprocAutoscalingPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DataprocAutoscalingPolicy)(nil)

// DataprocAutoscalingPolicy represents the Terraform resource google_dataproc_autoscaling_policy.
type DataprocAutoscalingPolicy struct {
	Name      string
	Args      DataprocAutoscalingPolicyArgs
	state     *dataprocAutoscalingPolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DataprocAutoscalingPolicy].
func (dap *DataprocAutoscalingPolicy) Type() string {
	return "google_dataproc_autoscaling_policy"
}

// LocalName returns the local name for [DataprocAutoscalingPolicy].
func (dap *DataprocAutoscalingPolicy) LocalName() string {
	return dap.Name
}

// Configuration returns the configuration (args) for [DataprocAutoscalingPolicy].
func (dap *DataprocAutoscalingPolicy) Configuration() interface{} {
	return dap.Args
}

// DependOn is used for other resources to depend on [DataprocAutoscalingPolicy].
func (dap *DataprocAutoscalingPolicy) DependOn() terra.Reference {
	return terra.ReferenceResource(dap)
}

// Dependencies returns the list of resources [DataprocAutoscalingPolicy] depends_on.
func (dap *DataprocAutoscalingPolicy) Dependencies() terra.Dependencies {
	return dap.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DataprocAutoscalingPolicy].
func (dap *DataprocAutoscalingPolicy) LifecycleManagement() *terra.Lifecycle {
	return dap.Lifecycle
}

// Attributes returns the attributes for [DataprocAutoscalingPolicy].
func (dap *DataprocAutoscalingPolicy) Attributes() dataprocAutoscalingPolicyAttributes {
	return dataprocAutoscalingPolicyAttributes{ref: terra.ReferenceResource(dap)}
}

// ImportState imports the given attribute values into [DataprocAutoscalingPolicy]'s state.
func (dap *DataprocAutoscalingPolicy) ImportState(av io.Reader) error {
	dap.state = &dataprocAutoscalingPolicyState{}
	if err := json.NewDecoder(av).Decode(dap.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dap.Type(), dap.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DataprocAutoscalingPolicy] has state.
func (dap *DataprocAutoscalingPolicy) State() (*dataprocAutoscalingPolicyState, bool) {
	return dap.state, dap.state != nil
}

// StateMust returns the state for [DataprocAutoscalingPolicy]. Panics if the state is nil.
func (dap *DataprocAutoscalingPolicy) StateMust() *dataprocAutoscalingPolicyState {
	if dap.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dap.Type(), dap.LocalName()))
	}
	return dap.state
}

// DataprocAutoscalingPolicyArgs contains the configurations for google_dataproc_autoscaling_policy.
type DataprocAutoscalingPolicyArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, optional
	Location terra.StringValue `hcl:"location,attr"`
	// PolicyId: string, required
	PolicyId terra.StringValue `hcl:"policy_id,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// BasicAlgorithm: optional
	BasicAlgorithm *dataprocautoscalingpolicy.BasicAlgorithm `hcl:"basic_algorithm,block"`
	// SecondaryWorkerConfig: optional
	SecondaryWorkerConfig *dataprocautoscalingpolicy.SecondaryWorkerConfig `hcl:"secondary_worker_config,block"`
	// Timeouts: optional
	Timeouts *dataprocautoscalingpolicy.Timeouts `hcl:"timeouts,block"`
	// WorkerConfig: optional
	WorkerConfig *dataprocautoscalingpolicy.WorkerConfig `hcl:"worker_config,block"`
}
type dataprocAutoscalingPolicyAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of google_dataproc_autoscaling_policy.
func (dap dataprocAutoscalingPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dap.ref.Append("id"))
}

// Location returns a reference to field location of google_dataproc_autoscaling_policy.
func (dap dataprocAutoscalingPolicyAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(dap.ref.Append("location"))
}

// Name returns a reference to field name of google_dataproc_autoscaling_policy.
func (dap dataprocAutoscalingPolicyAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(dap.ref.Append("name"))
}

// PolicyId returns a reference to field policy_id of google_dataproc_autoscaling_policy.
func (dap dataprocAutoscalingPolicyAttributes) PolicyId() terra.StringValue {
	return terra.ReferenceAsString(dap.ref.Append("policy_id"))
}

// Project returns a reference to field project of google_dataproc_autoscaling_policy.
func (dap dataprocAutoscalingPolicyAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(dap.ref.Append("project"))
}

func (dap dataprocAutoscalingPolicyAttributes) BasicAlgorithm() terra.ListValue[dataprocautoscalingpolicy.BasicAlgorithmAttributes] {
	return terra.ReferenceAsList[dataprocautoscalingpolicy.BasicAlgorithmAttributes](dap.ref.Append("basic_algorithm"))
}

func (dap dataprocAutoscalingPolicyAttributes) SecondaryWorkerConfig() terra.ListValue[dataprocautoscalingpolicy.SecondaryWorkerConfigAttributes] {
	return terra.ReferenceAsList[dataprocautoscalingpolicy.SecondaryWorkerConfigAttributes](dap.ref.Append("secondary_worker_config"))
}

func (dap dataprocAutoscalingPolicyAttributes) Timeouts() dataprocautoscalingpolicy.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dataprocautoscalingpolicy.TimeoutsAttributes](dap.ref.Append("timeouts"))
}

func (dap dataprocAutoscalingPolicyAttributes) WorkerConfig() terra.ListValue[dataprocautoscalingpolicy.WorkerConfigAttributes] {
	return terra.ReferenceAsList[dataprocautoscalingpolicy.WorkerConfigAttributes](dap.ref.Append("worker_config"))
}

type dataprocAutoscalingPolicyState struct {
	Id                    string                                                 `json:"id"`
	Location              string                                                 `json:"location"`
	Name                  string                                                 `json:"name"`
	PolicyId              string                                                 `json:"policy_id"`
	Project               string                                                 `json:"project"`
	BasicAlgorithm        []dataprocautoscalingpolicy.BasicAlgorithmState        `json:"basic_algorithm"`
	SecondaryWorkerConfig []dataprocautoscalingpolicy.SecondaryWorkerConfigState `json:"secondary_worker_config"`
	Timeouts              *dataprocautoscalingpolicy.TimeoutsState               `json:"timeouts"`
	WorkerConfig          []dataprocautoscalingpolicy.WorkerConfigState          `json:"worker_config"`
}
