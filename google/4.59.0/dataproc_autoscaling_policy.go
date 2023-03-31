// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	dataprocautoscalingpolicy "github.com/golingon/terraproviders/google/4.59.0/dataprocautoscalingpolicy"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewDataprocAutoscalingPolicy(name string, args DataprocAutoscalingPolicyArgs) *DataprocAutoscalingPolicy {
	return &DataprocAutoscalingPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DataprocAutoscalingPolicy)(nil)

type DataprocAutoscalingPolicy struct {
	Name  string
	Args  DataprocAutoscalingPolicyArgs
	state *dataprocAutoscalingPolicyState
}

func (dap *DataprocAutoscalingPolicy) Type() string {
	return "google_dataproc_autoscaling_policy"
}

func (dap *DataprocAutoscalingPolicy) LocalName() string {
	return dap.Name
}

func (dap *DataprocAutoscalingPolicy) Configuration() interface{} {
	return dap.Args
}

func (dap *DataprocAutoscalingPolicy) Attributes() dataprocAutoscalingPolicyAttributes {
	return dataprocAutoscalingPolicyAttributes{ref: terra.ReferenceResource(dap)}
}

func (dap *DataprocAutoscalingPolicy) ImportState(av io.Reader) error {
	dap.state = &dataprocAutoscalingPolicyState{}
	if err := json.NewDecoder(av).Decode(dap.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dap.Type(), dap.LocalName(), err)
	}
	return nil
}

func (dap *DataprocAutoscalingPolicy) State() (*dataprocAutoscalingPolicyState, bool) {
	return dap.state, dap.state != nil
}

func (dap *DataprocAutoscalingPolicy) StateMust() *dataprocAutoscalingPolicyState {
	if dap.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dap.Type(), dap.LocalName()))
	}
	return dap.state
}

func (dap *DataprocAutoscalingPolicy) DependOn() terra.Reference {
	return terra.ReferenceResource(dap)
}

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
	// DependsOn contains resources that DataprocAutoscalingPolicy depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type dataprocAutoscalingPolicyAttributes struct {
	ref terra.Reference
}

func (dap dataprocAutoscalingPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceString(dap.ref.Append("id"))
}

func (dap dataprocAutoscalingPolicyAttributes) Location() terra.StringValue {
	return terra.ReferenceString(dap.ref.Append("location"))
}

func (dap dataprocAutoscalingPolicyAttributes) Name() terra.StringValue {
	return terra.ReferenceString(dap.ref.Append("name"))
}

func (dap dataprocAutoscalingPolicyAttributes) PolicyId() terra.StringValue {
	return terra.ReferenceString(dap.ref.Append("policy_id"))
}

func (dap dataprocAutoscalingPolicyAttributes) Project() terra.StringValue {
	return terra.ReferenceString(dap.ref.Append("project"))
}

func (dap dataprocAutoscalingPolicyAttributes) BasicAlgorithm() terra.ListValue[dataprocautoscalingpolicy.BasicAlgorithmAttributes] {
	return terra.ReferenceList[dataprocautoscalingpolicy.BasicAlgorithmAttributes](dap.ref.Append("basic_algorithm"))
}

func (dap dataprocAutoscalingPolicyAttributes) SecondaryWorkerConfig() terra.ListValue[dataprocautoscalingpolicy.SecondaryWorkerConfigAttributes] {
	return terra.ReferenceList[dataprocautoscalingpolicy.SecondaryWorkerConfigAttributes](dap.ref.Append("secondary_worker_config"))
}

func (dap dataprocAutoscalingPolicyAttributes) Timeouts() dataprocautoscalingpolicy.TimeoutsAttributes {
	return terra.ReferenceSingle[dataprocautoscalingpolicy.TimeoutsAttributes](dap.ref.Append("timeouts"))
}

func (dap dataprocAutoscalingPolicyAttributes) WorkerConfig() terra.ListValue[dataprocautoscalingpolicy.WorkerConfigAttributes] {
	return terra.ReferenceList[dataprocautoscalingpolicy.WorkerConfigAttributes](dap.ref.Append("worker_config"))
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
