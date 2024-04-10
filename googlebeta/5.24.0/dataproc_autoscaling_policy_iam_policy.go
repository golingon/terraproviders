// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// NewDataprocAutoscalingPolicyIamPolicy creates a new instance of [DataprocAutoscalingPolicyIamPolicy].
func NewDataprocAutoscalingPolicyIamPolicy(name string, args DataprocAutoscalingPolicyIamPolicyArgs) *DataprocAutoscalingPolicyIamPolicy {
	return &DataprocAutoscalingPolicyIamPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DataprocAutoscalingPolicyIamPolicy)(nil)

// DataprocAutoscalingPolicyIamPolicy represents the Terraform resource google_dataproc_autoscaling_policy_iam_policy.
type DataprocAutoscalingPolicyIamPolicy struct {
	Name      string
	Args      DataprocAutoscalingPolicyIamPolicyArgs
	state     *dataprocAutoscalingPolicyIamPolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DataprocAutoscalingPolicyIamPolicy].
func (dapip *DataprocAutoscalingPolicyIamPolicy) Type() string {
	return "google_dataproc_autoscaling_policy_iam_policy"
}

// LocalName returns the local name for [DataprocAutoscalingPolicyIamPolicy].
func (dapip *DataprocAutoscalingPolicyIamPolicy) LocalName() string {
	return dapip.Name
}

// Configuration returns the configuration (args) for [DataprocAutoscalingPolicyIamPolicy].
func (dapip *DataprocAutoscalingPolicyIamPolicy) Configuration() interface{} {
	return dapip.Args
}

// DependOn is used for other resources to depend on [DataprocAutoscalingPolicyIamPolicy].
func (dapip *DataprocAutoscalingPolicyIamPolicy) DependOn() terra.Reference {
	return terra.ReferenceResource(dapip)
}

// Dependencies returns the list of resources [DataprocAutoscalingPolicyIamPolicy] depends_on.
func (dapip *DataprocAutoscalingPolicyIamPolicy) Dependencies() terra.Dependencies {
	return dapip.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DataprocAutoscalingPolicyIamPolicy].
func (dapip *DataprocAutoscalingPolicyIamPolicy) LifecycleManagement() *terra.Lifecycle {
	return dapip.Lifecycle
}

// Attributes returns the attributes for [DataprocAutoscalingPolicyIamPolicy].
func (dapip *DataprocAutoscalingPolicyIamPolicy) Attributes() dataprocAutoscalingPolicyIamPolicyAttributes {
	return dataprocAutoscalingPolicyIamPolicyAttributes{ref: terra.ReferenceResource(dapip)}
}

// ImportState imports the given attribute values into [DataprocAutoscalingPolicyIamPolicy]'s state.
func (dapip *DataprocAutoscalingPolicyIamPolicy) ImportState(av io.Reader) error {
	dapip.state = &dataprocAutoscalingPolicyIamPolicyState{}
	if err := json.NewDecoder(av).Decode(dapip.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dapip.Type(), dapip.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DataprocAutoscalingPolicyIamPolicy] has state.
func (dapip *DataprocAutoscalingPolicyIamPolicy) State() (*dataprocAutoscalingPolicyIamPolicyState, bool) {
	return dapip.state, dapip.state != nil
}

// StateMust returns the state for [DataprocAutoscalingPolicyIamPolicy]. Panics if the state is nil.
func (dapip *DataprocAutoscalingPolicyIamPolicy) StateMust() *dataprocAutoscalingPolicyIamPolicyState {
	if dapip.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dapip.Type(), dapip.LocalName()))
	}
	return dapip.state
}

// DataprocAutoscalingPolicyIamPolicyArgs contains the configurations for google_dataproc_autoscaling_policy_iam_policy.
type DataprocAutoscalingPolicyIamPolicyArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, optional
	Location terra.StringValue `hcl:"location,attr"`
	// PolicyData: string, required
	PolicyData terra.StringValue `hcl:"policy_data,attr" validate:"required"`
	// PolicyId: string, required
	PolicyId terra.StringValue `hcl:"policy_id,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
}
type dataprocAutoscalingPolicyIamPolicyAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_dataproc_autoscaling_policy_iam_policy.
func (dapip dataprocAutoscalingPolicyIamPolicyAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(dapip.ref.Append("etag"))
}

// Id returns a reference to field id of google_dataproc_autoscaling_policy_iam_policy.
func (dapip dataprocAutoscalingPolicyIamPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dapip.ref.Append("id"))
}

// Location returns a reference to field location of google_dataproc_autoscaling_policy_iam_policy.
func (dapip dataprocAutoscalingPolicyIamPolicyAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(dapip.ref.Append("location"))
}

// PolicyData returns a reference to field policy_data of google_dataproc_autoscaling_policy_iam_policy.
func (dapip dataprocAutoscalingPolicyIamPolicyAttributes) PolicyData() terra.StringValue {
	return terra.ReferenceAsString(dapip.ref.Append("policy_data"))
}

// PolicyId returns a reference to field policy_id of google_dataproc_autoscaling_policy_iam_policy.
func (dapip dataprocAutoscalingPolicyIamPolicyAttributes) PolicyId() terra.StringValue {
	return terra.ReferenceAsString(dapip.ref.Append("policy_id"))
}

// Project returns a reference to field project of google_dataproc_autoscaling_policy_iam_policy.
func (dapip dataprocAutoscalingPolicyIamPolicyAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(dapip.ref.Append("project"))
}

type dataprocAutoscalingPolicyIamPolicyState struct {
	Etag       string `json:"etag"`
	Id         string `json:"id"`
	Location   string `json:"location"`
	PolicyData string `json:"policy_data"`
	PolicyId   string `json:"policy_id"`
	Project    string `json:"project"`
}
