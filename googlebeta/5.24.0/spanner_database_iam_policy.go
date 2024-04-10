// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// NewSpannerDatabaseIamPolicy creates a new instance of [SpannerDatabaseIamPolicy].
func NewSpannerDatabaseIamPolicy(name string, args SpannerDatabaseIamPolicyArgs) *SpannerDatabaseIamPolicy {
	return &SpannerDatabaseIamPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SpannerDatabaseIamPolicy)(nil)

// SpannerDatabaseIamPolicy represents the Terraform resource google_spanner_database_iam_policy.
type SpannerDatabaseIamPolicy struct {
	Name      string
	Args      SpannerDatabaseIamPolicyArgs
	state     *spannerDatabaseIamPolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SpannerDatabaseIamPolicy].
func (sdip *SpannerDatabaseIamPolicy) Type() string {
	return "google_spanner_database_iam_policy"
}

// LocalName returns the local name for [SpannerDatabaseIamPolicy].
func (sdip *SpannerDatabaseIamPolicy) LocalName() string {
	return sdip.Name
}

// Configuration returns the configuration (args) for [SpannerDatabaseIamPolicy].
func (sdip *SpannerDatabaseIamPolicy) Configuration() interface{} {
	return sdip.Args
}

// DependOn is used for other resources to depend on [SpannerDatabaseIamPolicy].
func (sdip *SpannerDatabaseIamPolicy) DependOn() terra.Reference {
	return terra.ReferenceResource(sdip)
}

// Dependencies returns the list of resources [SpannerDatabaseIamPolicy] depends_on.
func (sdip *SpannerDatabaseIamPolicy) Dependencies() terra.Dependencies {
	return sdip.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SpannerDatabaseIamPolicy].
func (sdip *SpannerDatabaseIamPolicy) LifecycleManagement() *terra.Lifecycle {
	return sdip.Lifecycle
}

// Attributes returns the attributes for [SpannerDatabaseIamPolicy].
func (sdip *SpannerDatabaseIamPolicy) Attributes() spannerDatabaseIamPolicyAttributes {
	return spannerDatabaseIamPolicyAttributes{ref: terra.ReferenceResource(sdip)}
}

// ImportState imports the given attribute values into [SpannerDatabaseIamPolicy]'s state.
func (sdip *SpannerDatabaseIamPolicy) ImportState(av io.Reader) error {
	sdip.state = &spannerDatabaseIamPolicyState{}
	if err := json.NewDecoder(av).Decode(sdip.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sdip.Type(), sdip.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SpannerDatabaseIamPolicy] has state.
func (sdip *SpannerDatabaseIamPolicy) State() (*spannerDatabaseIamPolicyState, bool) {
	return sdip.state, sdip.state != nil
}

// StateMust returns the state for [SpannerDatabaseIamPolicy]. Panics if the state is nil.
func (sdip *SpannerDatabaseIamPolicy) StateMust() *spannerDatabaseIamPolicyState {
	if sdip.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sdip.Type(), sdip.LocalName()))
	}
	return sdip.state
}

// SpannerDatabaseIamPolicyArgs contains the configurations for google_spanner_database_iam_policy.
type SpannerDatabaseIamPolicyArgs struct {
	// Database: string, required
	Database terra.StringValue `hcl:"database,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Instance: string, required
	Instance terra.StringValue `hcl:"instance,attr" validate:"required"`
	// PolicyData: string, required
	PolicyData terra.StringValue `hcl:"policy_data,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
}
type spannerDatabaseIamPolicyAttributes struct {
	ref terra.Reference
}

// Database returns a reference to field database of google_spanner_database_iam_policy.
func (sdip spannerDatabaseIamPolicyAttributes) Database() terra.StringValue {
	return terra.ReferenceAsString(sdip.ref.Append("database"))
}

// Etag returns a reference to field etag of google_spanner_database_iam_policy.
func (sdip spannerDatabaseIamPolicyAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(sdip.ref.Append("etag"))
}

// Id returns a reference to field id of google_spanner_database_iam_policy.
func (sdip spannerDatabaseIamPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sdip.ref.Append("id"))
}

// Instance returns a reference to field instance of google_spanner_database_iam_policy.
func (sdip spannerDatabaseIamPolicyAttributes) Instance() terra.StringValue {
	return terra.ReferenceAsString(sdip.ref.Append("instance"))
}

// PolicyData returns a reference to field policy_data of google_spanner_database_iam_policy.
func (sdip spannerDatabaseIamPolicyAttributes) PolicyData() terra.StringValue {
	return terra.ReferenceAsString(sdip.ref.Append("policy_data"))
}

// Project returns a reference to field project of google_spanner_database_iam_policy.
func (sdip spannerDatabaseIamPolicyAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(sdip.ref.Append("project"))
}

type spannerDatabaseIamPolicyState struct {
	Database   string `json:"database"`
	Etag       string `json:"etag"`
	Id         string `json:"id"`
	Instance   string `json:"instance"`
	PolicyData string `json:"policy_data"`
	Project    string `json:"project"`
}
