// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewBigtableTableIamPolicy creates a new instance of [BigtableTableIamPolicy].
func NewBigtableTableIamPolicy(name string, args BigtableTableIamPolicyArgs) *BigtableTableIamPolicy {
	return &BigtableTableIamPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*BigtableTableIamPolicy)(nil)

// BigtableTableIamPolicy represents the Terraform resource google_bigtable_table_iam_policy.
type BigtableTableIamPolicy struct {
	Name      string
	Args      BigtableTableIamPolicyArgs
	state     *bigtableTableIamPolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [BigtableTableIamPolicy].
func (btip *BigtableTableIamPolicy) Type() string {
	return "google_bigtable_table_iam_policy"
}

// LocalName returns the local name for [BigtableTableIamPolicy].
func (btip *BigtableTableIamPolicy) LocalName() string {
	return btip.Name
}

// Configuration returns the configuration (args) for [BigtableTableIamPolicy].
func (btip *BigtableTableIamPolicy) Configuration() interface{} {
	return btip.Args
}

// DependOn is used for other resources to depend on [BigtableTableIamPolicy].
func (btip *BigtableTableIamPolicy) DependOn() terra.Reference {
	return terra.ReferenceResource(btip)
}

// Dependencies returns the list of resources [BigtableTableIamPolicy] depends_on.
func (btip *BigtableTableIamPolicy) Dependencies() terra.Dependencies {
	return btip.DependsOn
}

// LifecycleManagement returns the lifecycle block for [BigtableTableIamPolicy].
func (btip *BigtableTableIamPolicy) LifecycleManagement() *terra.Lifecycle {
	return btip.Lifecycle
}

// Attributes returns the attributes for [BigtableTableIamPolicy].
func (btip *BigtableTableIamPolicy) Attributes() bigtableTableIamPolicyAttributes {
	return bigtableTableIamPolicyAttributes{ref: terra.ReferenceResource(btip)}
}

// ImportState imports the given attribute values into [BigtableTableIamPolicy]'s state.
func (btip *BigtableTableIamPolicy) ImportState(av io.Reader) error {
	btip.state = &bigtableTableIamPolicyState{}
	if err := json.NewDecoder(av).Decode(btip.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", btip.Type(), btip.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [BigtableTableIamPolicy] has state.
func (btip *BigtableTableIamPolicy) State() (*bigtableTableIamPolicyState, bool) {
	return btip.state, btip.state != nil
}

// StateMust returns the state for [BigtableTableIamPolicy]. Panics if the state is nil.
func (btip *BigtableTableIamPolicy) StateMust() *bigtableTableIamPolicyState {
	if btip.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", btip.Type(), btip.LocalName()))
	}
	return btip.state
}

// BigtableTableIamPolicyArgs contains the configurations for google_bigtable_table_iam_policy.
type BigtableTableIamPolicyArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Instance: string, required
	Instance terra.StringValue `hcl:"instance,attr" validate:"required"`
	// PolicyData: string, required
	PolicyData terra.StringValue `hcl:"policy_data,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Table: string, required
	Table terra.StringValue `hcl:"table,attr" validate:"required"`
}
type bigtableTableIamPolicyAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_bigtable_table_iam_policy.
func (btip bigtableTableIamPolicyAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(btip.ref.Append("etag"))
}

// Id returns a reference to field id of google_bigtable_table_iam_policy.
func (btip bigtableTableIamPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(btip.ref.Append("id"))
}

// Instance returns a reference to field instance of google_bigtable_table_iam_policy.
func (btip bigtableTableIamPolicyAttributes) Instance() terra.StringValue {
	return terra.ReferenceAsString(btip.ref.Append("instance"))
}

// PolicyData returns a reference to field policy_data of google_bigtable_table_iam_policy.
func (btip bigtableTableIamPolicyAttributes) PolicyData() terra.StringValue {
	return terra.ReferenceAsString(btip.ref.Append("policy_data"))
}

// Project returns a reference to field project of google_bigtable_table_iam_policy.
func (btip bigtableTableIamPolicyAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(btip.ref.Append("project"))
}

// Table returns a reference to field table of google_bigtable_table_iam_policy.
func (btip bigtableTableIamPolicyAttributes) Table() terra.StringValue {
	return terra.ReferenceAsString(btip.ref.Append("table"))
}

type bigtableTableIamPolicyState struct {
	Etag       string `json:"etag"`
	Id         string `json:"id"`
	Instance   string `json:"instance"`
	PolicyData string `json:"policy_data"`
	Project    string `json:"project"`
	Table      string `json:"table"`
}
