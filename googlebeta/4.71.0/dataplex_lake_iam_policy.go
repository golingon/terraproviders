// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDataplexLakeIamPolicy creates a new instance of [DataplexLakeIamPolicy].
func NewDataplexLakeIamPolicy(name string, args DataplexLakeIamPolicyArgs) *DataplexLakeIamPolicy {
	return &DataplexLakeIamPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DataplexLakeIamPolicy)(nil)

// DataplexLakeIamPolicy represents the Terraform resource google_dataplex_lake_iam_policy.
type DataplexLakeIamPolicy struct {
	Name      string
	Args      DataplexLakeIamPolicyArgs
	state     *dataplexLakeIamPolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DataplexLakeIamPolicy].
func (dlip *DataplexLakeIamPolicy) Type() string {
	return "google_dataplex_lake_iam_policy"
}

// LocalName returns the local name for [DataplexLakeIamPolicy].
func (dlip *DataplexLakeIamPolicy) LocalName() string {
	return dlip.Name
}

// Configuration returns the configuration (args) for [DataplexLakeIamPolicy].
func (dlip *DataplexLakeIamPolicy) Configuration() interface{} {
	return dlip.Args
}

// DependOn is used for other resources to depend on [DataplexLakeIamPolicy].
func (dlip *DataplexLakeIamPolicy) DependOn() terra.Reference {
	return terra.ReferenceResource(dlip)
}

// Dependencies returns the list of resources [DataplexLakeIamPolicy] depends_on.
func (dlip *DataplexLakeIamPolicy) Dependencies() terra.Dependencies {
	return dlip.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DataplexLakeIamPolicy].
func (dlip *DataplexLakeIamPolicy) LifecycleManagement() *terra.Lifecycle {
	return dlip.Lifecycle
}

// Attributes returns the attributes for [DataplexLakeIamPolicy].
func (dlip *DataplexLakeIamPolicy) Attributes() dataplexLakeIamPolicyAttributes {
	return dataplexLakeIamPolicyAttributes{ref: terra.ReferenceResource(dlip)}
}

// ImportState imports the given attribute values into [DataplexLakeIamPolicy]'s state.
func (dlip *DataplexLakeIamPolicy) ImportState(av io.Reader) error {
	dlip.state = &dataplexLakeIamPolicyState{}
	if err := json.NewDecoder(av).Decode(dlip.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dlip.Type(), dlip.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DataplexLakeIamPolicy] has state.
func (dlip *DataplexLakeIamPolicy) State() (*dataplexLakeIamPolicyState, bool) {
	return dlip.state, dlip.state != nil
}

// StateMust returns the state for [DataplexLakeIamPolicy]. Panics if the state is nil.
func (dlip *DataplexLakeIamPolicy) StateMust() *dataplexLakeIamPolicyState {
	if dlip.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dlip.Type(), dlip.LocalName()))
	}
	return dlip.state
}

// DataplexLakeIamPolicyArgs contains the configurations for google_dataplex_lake_iam_policy.
type DataplexLakeIamPolicyArgs struct {
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
}
type dataplexLakeIamPolicyAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_dataplex_lake_iam_policy.
func (dlip dataplexLakeIamPolicyAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(dlip.ref.Append("etag"))
}

// Id returns a reference to field id of google_dataplex_lake_iam_policy.
func (dlip dataplexLakeIamPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dlip.ref.Append("id"))
}

// Lake returns a reference to field lake of google_dataplex_lake_iam_policy.
func (dlip dataplexLakeIamPolicyAttributes) Lake() terra.StringValue {
	return terra.ReferenceAsString(dlip.ref.Append("lake"))
}

// Location returns a reference to field location of google_dataplex_lake_iam_policy.
func (dlip dataplexLakeIamPolicyAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(dlip.ref.Append("location"))
}

// PolicyData returns a reference to field policy_data of google_dataplex_lake_iam_policy.
func (dlip dataplexLakeIamPolicyAttributes) PolicyData() terra.StringValue {
	return terra.ReferenceAsString(dlip.ref.Append("policy_data"))
}

// Project returns a reference to field project of google_dataplex_lake_iam_policy.
func (dlip dataplexLakeIamPolicyAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(dlip.ref.Append("project"))
}

type dataplexLakeIamPolicyState struct {
	Etag       string `json:"etag"`
	Id         string `json:"id"`
	Lake       string `json:"lake"`
	Location   string `json:"location"`
	PolicyData string `json:"policy_data"`
	Project    string `json:"project"`
}