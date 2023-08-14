// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewBigqueryTableIamPolicy creates a new instance of [BigqueryTableIamPolicy].
func NewBigqueryTableIamPolicy(name string, args BigqueryTableIamPolicyArgs) *BigqueryTableIamPolicy {
	return &BigqueryTableIamPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*BigqueryTableIamPolicy)(nil)

// BigqueryTableIamPolicy represents the Terraform resource google_bigquery_table_iam_policy.
type BigqueryTableIamPolicy struct {
	Name      string
	Args      BigqueryTableIamPolicyArgs
	state     *bigqueryTableIamPolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [BigqueryTableIamPolicy].
func (btip *BigqueryTableIamPolicy) Type() string {
	return "google_bigquery_table_iam_policy"
}

// LocalName returns the local name for [BigqueryTableIamPolicy].
func (btip *BigqueryTableIamPolicy) LocalName() string {
	return btip.Name
}

// Configuration returns the configuration (args) for [BigqueryTableIamPolicy].
func (btip *BigqueryTableIamPolicy) Configuration() interface{} {
	return btip.Args
}

// DependOn is used for other resources to depend on [BigqueryTableIamPolicy].
func (btip *BigqueryTableIamPolicy) DependOn() terra.Reference {
	return terra.ReferenceResource(btip)
}

// Dependencies returns the list of resources [BigqueryTableIamPolicy] depends_on.
func (btip *BigqueryTableIamPolicy) Dependencies() terra.Dependencies {
	return btip.DependsOn
}

// LifecycleManagement returns the lifecycle block for [BigqueryTableIamPolicy].
func (btip *BigqueryTableIamPolicy) LifecycleManagement() *terra.Lifecycle {
	return btip.Lifecycle
}

// Attributes returns the attributes for [BigqueryTableIamPolicy].
func (btip *BigqueryTableIamPolicy) Attributes() bigqueryTableIamPolicyAttributes {
	return bigqueryTableIamPolicyAttributes{ref: terra.ReferenceResource(btip)}
}

// ImportState imports the given attribute values into [BigqueryTableIamPolicy]'s state.
func (btip *BigqueryTableIamPolicy) ImportState(av io.Reader) error {
	btip.state = &bigqueryTableIamPolicyState{}
	if err := json.NewDecoder(av).Decode(btip.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", btip.Type(), btip.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [BigqueryTableIamPolicy] has state.
func (btip *BigqueryTableIamPolicy) State() (*bigqueryTableIamPolicyState, bool) {
	return btip.state, btip.state != nil
}

// StateMust returns the state for [BigqueryTableIamPolicy]. Panics if the state is nil.
func (btip *BigqueryTableIamPolicy) StateMust() *bigqueryTableIamPolicyState {
	if btip.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", btip.Type(), btip.LocalName()))
	}
	return btip.state
}

// BigqueryTableIamPolicyArgs contains the configurations for google_bigquery_table_iam_policy.
type BigqueryTableIamPolicyArgs struct {
	// DatasetId: string, required
	DatasetId terra.StringValue `hcl:"dataset_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// PolicyData: string, required
	PolicyData terra.StringValue `hcl:"policy_data,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// TableId: string, required
	TableId terra.StringValue `hcl:"table_id,attr" validate:"required"`
}
type bigqueryTableIamPolicyAttributes struct {
	ref terra.Reference
}

// DatasetId returns a reference to field dataset_id of google_bigquery_table_iam_policy.
func (btip bigqueryTableIamPolicyAttributes) DatasetId() terra.StringValue {
	return terra.ReferenceAsString(btip.ref.Append("dataset_id"))
}

// Etag returns a reference to field etag of google_bigquery_table_iam_policy.
func (btip bigqueryTableIamPolicyAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(btip.ref.Append("etag"))
}

// Id returns a reference to field id of google_bigquery_table_iam_policy.
func (btip bigqueryTableIamPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(btip.ref.Append("id"))
}

// PolicyData returns a reference to field policy_data of google_bigquery_table_iam_policy.
func (btip bigqueryTableIamPolicyAttributes) PolicyData() terra.StringValue {
	return terra.ReferenceAsString(btip.ref.Append("policy_data"))
}

// Project returns a reference to field project of google_bigquery_table_iam_policy.
func (btip bigqueryTableIamPolicyAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(btip.ref.Append("project"))
}

// TableId returns a reference to field table_id of google_bigquery_table_iam_policy.
func (btip bigqueryTableIamPolicyAttributes) TableId() terra.StringValue {
	return terra.ReferenceAsString(btip.ref.Append("table_id"))
}

type bigqueryTableIamPolicyState struct {
	DatasetId  string `json:"dataset_id"`
	Etag       string `json:"etag"`
	Id         string `json:"id"`
	PolicyData string `json:"policy_data"`
	Project    string `json:"project"`
	TableId    string `json:"table_id"`
}