// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	dataprocautoscalingpolicyiambinding "github.com/golingon/terraproviders/google/4.63.1/dataprocautoscalingpolicyiambinding"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDataprocAutoscalingPolicyIamBinding creates a new instance of [DataprocAutoscalingPolicyIamBinding].
func NewDataprocAutoscalingPolicyIamBinding(name string, args DataprocAutoscalingPolicyIamBindingArgs) *DataprocAutoscalingPolicyIamBinding {
	return &DataprocAutoscalingPolicyIamBinding{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DataprocAutoscalingPolicyIamBinding)(nil)

// DataprocAutoscalingPolicyIamBinding represents the Terraform resource google_dataproc_autoscaling_policy_iam_binding.
type DataprocAutoscalingPolicyIamBinding struct {
	Name      string
	Args      DataprocAutoscalingPolicyIamBindingArgs
	state     *dataprocAutoscalingPolicyIamBindingState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DataprocAutoscalingPolicyIamBinding].
func (dapib *DataprocAutoscalingPolicyIamBinding) Type() string {
	return "google_dataproc_autoscaling_policy_iam_binding"
}

// LocalName returns the local name for [DataprocAutoscalingPolicyIamBinding].
func (dapib *DataprocAutoscalingPolicyIamBinding) LocalName() string {
	return dapib.Name
}

// Configuration returns the configuration (args) for [DataprocAutoscalingPolicyIamBinding].
func (dapib *DataprocAutoscalingPolicyIamBinding) Configuration() interface{} {
	return dapib.Args
}

// DependOn is used for other resources to depend on [DataprocAutoscalingPolicyIamBinding].
func (dapib *DataprocAutoscalingPolicyIamBinding) DependOn() terra.Reference {
	return terra.ReferenceResource(dapib)
}

// Dependencies returns the list of resources [DataprocAutoscalingPolicyIamBinding] depends_on.
func (dapib *DataprocAutoscalingPolicyIamBinding) Dependencies() terra.Dependencies {
	return dapib.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DataprocAutoscalingPolicyIamBinding].
func (dapib *DataprocAutoscalingPolicyIamBinding) LifecycleManagement() *terra.Lifecycle {
	return dapib.Lifecycle
}

// Attributes returns the attributes for [DataprocAutoscalingPolicyIamBinding].
func (dapib *DataprocAutoscalingPolicyIamBinding) Attributes() dataprocAutoscalingPolicyIamBindingAttributes {
	return dataprocAutoscalingPolicyIamBindingAttributes{ref: terra.ReferenceResource(dapib)}
}

// ImportState imports the given attribute values into [DataprocAutoscalingPolicyIamBinding]'s state.
func (dapib *DataprocAutoscalingPolicyIamBinding) ImportState(av io.Reader) error {
	dapib.state = &dataprocAutoscalingPolicyIamBindingState{}
	if err := json.NewDecoder(av).Decode(dapib.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dapib.Type(), dapib.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DataprocAutoscalingPolicyIamBinding] has state.
func (dapib *DataprocAutoscalingPolicyIamBinding) State() (*dataprocAutoscalingPolicyIamBindingState, bool) {
	return dapib.state, dapib.state != nil
}

// StateMust returns the state for [DataprocAutoscalingPolicyIamBinding]. Panics if the state is nil.
func (dapib *DataprocAutoscalingPolicyIamBinding) StateMust() *dataprocAutoscalingPolicyIamBindingState {
	if dapib.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dapib.Type(), dapib.LocalName()))
	}
	return dapib.state
}

// DataprocAutoscalingPolicyIamBindingArgs contains the configurations for google_dataproc_autoscaling_policy_iam_binding.
type DataprocAutoscalingPolicyIamBindingArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, optional
	Location terra.StringValue `hcl:"location,attr"`
	// Members: set of string, required
	Members terra.SetValue[terra.StringValue] `hcl:"members,attr" validate:"required"`
	// PolicyId: string, required
	PolicyId terra.StringValue `hcl:"policy_id,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// Condition: optional
	Condition *dataprocautoscalingpolicyiambinding.Condition `hcl:"condition,block"`
}
type dataprocAutoscalingPolicyIamBindingAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_dataproc_autoscaling_policy_iam_binding.
func (dapib dataprocAutoscalingPolicyIamBindingAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(dapib.ref.Append("etag"))
}

// Id returns a reference to field id of google_dataproc_autoscaling_policy_iam_binding.
func (dapib dataprocAutoscalingPolicyIamBindingAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dapib.ref.Append("id"))
}

// Location returns a reference to field location of google_dataproc_autoscaling_policy_iam_binding.
func (dapib dataprocAutoscalingPolicyIamBindingAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(dapib.ref.Append("location"))
}

// Members returns a reference to field members of google_dataproc_autoscaling_policy_iam_binding.
func (dapib dataprocAutoscalingPolicyIamBindingAttributes) Members() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](dapib.ref.Append("members"))
}

// PolicyId returns a reference to field policy_id of google_dataproc_autoscaling_policy_iam_binding.
func (dapib dataprocAutoscalingPolicyIamBindingAttributes) PolicyId() terra.StringValue {
	return terra.ReferenceAsString(dapib.ref.Append("policy_id"))
}

// Project returns a reference to field project of google_dataproc_autoscaling_policy_iam_binding.
func (dapib dataprocAutoscalingPolicyIamBindingAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(dapib.ref.Append("project"))
}

// Role returns a reference to field role of google_dataproc_autoscaling_policy_iam_binding.
func (dapib dataprocAutoscalingPolicyIamBindingAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(dapib.ref.Append("role"))
}

func (dapib dataprocAutoscalingPolicyIamBindingAttributes) Condition() terra.ListValue[dataprocautoscalingpolicyiambinding.ConditionAttributes] {
	return terra.ReferenceAsList[dataprocautoscalingpolicyiambinding.ConditionAttributes](dapib.ref.Append("condition"))
}

type dataprocAutoscalingPolicyIamBindingState struct {
	Etag      string                                               `json:"etag"`
	Id        string                                               `json:"id"`
	Location  string                                               `json:"location"`
	Members   []string                                             `json:"members"`
	PolicyId  string                                               `json:"policy_id"`
	Project   string                                               `json:"project"`
	Role      string                                               `json:"role"`
	Condition []dataprocautoscalingpolicyiambinding.ConditionState `json:"condition"`
}
