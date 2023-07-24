// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDataprocJobIamPolicy creates a new instance of [DataprocJobIamPolicy].
func NewDataprocJobIamPolicy(name string, args DataprocJobIamPolicyArgs) *DataprocJobIamPolicy {
	return &DataprocJobIamPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DataprocJobIamPolicy)(nil)

// DataprocJobIamPolicy represents the Terraform resource google_dataproc_job_iam_policy.
type DataprocJobIamPolicy struct {
	Name      string
	Args      DataprocJobIamPolicyArgs
	state     *dataprocJobIamPolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DataprocJobIamPolicy].
func (djip *DataprocJobIamPolicy) Type() string {
	return "google_dataproc_job_iam_policy"
}

// LocalName returns the local name for [DataprocJobIamPolicy].
func (djip *DataprocJobIamPolicy) LocalName() string {
	return djip.Name
}

// Configuration returns the configuration (args) for [DataprocJobIamPolicy].
func (djip *DataprocJobIamPolicy) Configuration() interface{} {
	return djip.Args
}

// DependOn is used for other resources to depend on [DataprocJobIamPolicy].
func (djip *DataprocJobIamPolicy) DependOn() terra.Reference {
	return terra.ReferenceResource(djip)
}

// Dependencies returns the list of resources [DataprocJobIamPolicy] depends_on.
func (djip *DataprocJobIamPolicy) Dependencies() terra.Dependencies {
	return djip.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DataprocJobIamPolicy].
func (djip *DataprocJobIamPolicy) LifecycleManagement() *terra.Lifecycle {
	return djip.Lifecycle
}

// Attributes returns the attributes for [DataprocJobIamPolicy].
func (djip *DataprocJobIamPolicy) Attributes() dataprocJobIamPolicyAttributes {
	return dataprocJobIamPolicyAttributes{ref: terra.ReferenceResource(djip)}
}

// ImportState imports the given attribute values into [DataprocJobIamPolicy]'s state.
func (djip *DataprocJobIamPolicy) ImportState(av io.Reader) error {
	djip.state = &dataprocJobIamPolicyState{}
	if err := json.NewDecoder(av).Decode(djip.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", djip.Type(), djip.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DataprocJobIamPolicy] has state.
func (djip *DataprocJobIamPolicy) State() (*dataprocJobIamPolicyState, bool) {
	return djip.state, djip.state != nil
}

// StateMust returns the state for [DataprocJobIamPolicy]. Panics if the state is nil.
func (djip *DataprocJobIamPolicy) StateMust() *dataprocJobIamPolicyState {
	if djip.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", djip.Type(), djip.LocalName()))
	}
	return djip.state
}

// DataprocJobIamPolicyArgs contains the configurations for google_dataproc_job_iam_policy.
type DataprocJobIamPolicyArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// JobId: string, required
	JobId terra.StringValue `hcl:"job_id,attr" validate:"required"`
	// PolicyData: string, required
	PolicyData terra.StringValue `hcl:"policy_data,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Region: string, optional
	Region terra.StringValue `hcl:"region,attr"`
}
type dataprocJobIamPolicyAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_dataproc_job_iam_policy.
func (djip dataprocJobIamPolicyAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(djip.ref.Append("etag"))
}

// Id returns a reference to field id of google_dataproc_job_iam_policy.
func (djip dataprocJobIamPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(djip.ref.Append("id"))
}

// JobId returns a reference to field job_id of google_dataproc_job_iam_policy.
func (djip dataprocJobIamPolicyAttributes) JobId() terra.StringValue {
	return terra.ReferenceAsString(djip.ref.Append("job_id"))
}

// PolicyData returns a reference to field policy_data of google_dataproc_job_iam_policy.
func (djip dataprocJobIamPolicyAttributes) PolicyData() terra.StringValue {
	return terra.ReferenceAsString(djip.ref.Append("policy_data"))
}

// Project returns a reference to field project of google_dataproc_job_iam_policy.
func (djip dataprocJobIamPolicyAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(djip.ref.Append("project"))
}

// Region returns a reference to field region of google_dataproc_job_iam_policy.
func (djip dataprocJobIamPolicyAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(djip.ref.Append("region"))
}

type dataprocJobIamPolicyState struct {
	Etag       string `json:"etag"`
	Id         string `json:"id"`
	JobId      string `json:"job_id"`
	PolicyData string `json:"policy_data"`
	Project    string `json:"project"`
	Region     string `json:"region"`
}
