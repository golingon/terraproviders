// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDataplexTaskIamPolicy creates a new instance of [DataplexTaskIamPolicy].
func NewDataplexTaskIamPolicy(name string, args DataplexTaskIamPolicyArgs) *DataplexTaskIamPolicy {
	return &DataplexTaskIamPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DataplexTaskIamPolicy)(nil)

// DataplexTaskIamPolicy represents the Terraform resource google_dataplex_task_iam_policy.
type DataplexTaskIamPolicy struct {
	Name      string
	Args      DataplexTaskIamPolicyArgs
	state     *dataplexTaskIamPolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DataplexTaskIamPolicy].
func (dtip *DataplexTaskIamPolicy) Type() string {
	return "google_dataplex_task_iam_policy"
}

// LocalName returns the local name for [DataplexTaskIamPolicy].
func (dtip *DataplexTaskIamPolicy) LocalName() string {
	return dtip.Name
}

// Configuration returns the configuration (args) for [DataplexTaskIamPolicy].
func (dtip *DataplexTaskIamPolicy) Configuration() interface{} {
	return dtip.Args
}

// DependOn is used for other resources to depend on [DataplexTaskIamPolicy].
func (dtip *DataplexTaskIamPolicy) DependOn() terra.Reference {
	return terra.ReferenceResource(dtip)
}

// Dependencies returns the list of resources [DataplexTaskIamPolicy] depends_on.
func (dtip *DataplexTaskIamPolicy) Dependencies() terra.Dependencies {
	return dtip.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DataplexTaskIamPolicy].
func (dtip *DataplexTaskIamPolicy) LifecycleManagement() *terra.Lifecycle {
	return dtip.Lifecycle
}

// Attributes returns the attributes for [DataplexTaskIamPolicy].
func (dtip *DataplexTaskIamPolicy) Attributes() dataplexTaskIamPolicyAttributes {
	return dataplexTaskIamPolicyAttributes{ref: terra.ReferenceResource(dtip)}
}

// ImportState imports the given attribute values into [DataplexTaskIamPolicy]'s state.
func (dtip *DataplexTaskIamPolicy) ImportState(av io.Reader) error {
	dtip.state = &dataplexTaskIamPolicyState{}
	if err := json.NewDecoder(av).Decode(dtip.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dtip.Type(), dtip.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DataplexTaskIamPolicy] has state.
func (dtip *DataplexTaskIamPolicy) State() (*dataplexTaskIamPolicyState, bool) {
	return dtip.state, dtip.state != nil
}

// StateMust returns the state for [DataplexTaskIamPolicy]. Panics if the state is nil.
func (dtip *DataplexTaskIamPolicy) StateMust() *dataplexTaskIamPolicyState {
	if dtip.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dtip.Type(), dtip.LocalName()))
	}
	return dtip.state
}

// DataplexTaskIamPolicyArgs contains the configurations for google_dataplex_task_iam_policy.
type DataplexTaskIamPolicyArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Lake: string, required
	Lake terra.StringValue `hcl:"lake,attr" validate:"required"`
	// Location: string, optional
	Location terra.StringValue `hcl:"location,attr"`
	// PolicyData: string, required
	PolicyData terra.StringValue `hcl:"policy_data,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// TaskId: string, required
	TaskId terra.StringValue `hcl:"task_id,attr" validate:"required"`
}
type dataplexTaskIamPolicyAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_dataplex_task_iam_policy.
func (dtip dataplexTaskIamPolicyAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(dtip.ref.Append("etag"))
}

// Id returns a reference to field id of google_dataplex_task_iam_policy.
func (dtip dataplexTaskIamPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dtip.ref.Append("id"))
}

// Lake returns a reference to field lake of google_dataplex_task_iam_policy.
func (dtip dataplexTaskIamPolicyAttributes) Lake() terra.StringValue {
	return terra.ReferenceAsString(dtip.ref.Append("lake"))
}

// Location returns a reference to field location of google_dataplex_task_iam_policy.
func (dtip dataplexTaskIamPolicyAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(dtip.ref.Append("location"))
}

// PolicyData returns a reference to field policy_data of google_dataplex_task_iam_policy.
func (dtip dataplexTaskIamPolicyAttributes) PolicyData() terra.StringValue {
	return terra.ReferenceAsString(dtip.ref.Append("policy_data"))
}

// Project returns a reference to field project of google_dataplex_task_iam_policy.
func (dtip dataplexTaskIamPolicyAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(dtip.ref.Append("project"))
}

// TaskId returns a reference to field task_id of google_dataplex_task_iam_policy.
func (dtip dataplexTaskIamPolicyAttributes) TaskId() terra.StringValue {
	return terra.ReferenceAsString(dtip.ref.Append("task_id"))
}

type dataplexTaskIamPolicyState struct {
	Etag       string `json:"etag"`
	Id         string `json:"id"`
	Lake       string `json:"lake"`
	Location   string `json:"location"`
	PolicyData string `json:"policy_data"`
	Project    string `json:"project"`
	TaskId     string `json:"task_id"`
}
