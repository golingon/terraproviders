// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	spannerdatabaseiammember "github.com/golingon/terraproviders/google/4.64.0/spannerdatabaseiammember"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSpannerDatabaseIamMember creates a new instance of [SpannerDatabaseIamMember].
func NewSpannerDatabaseIamMember(name string, args SpannerDatabaseIamMemberArgs) *SpannerDatabaseIamMember {
	return &SpannerDatabaseIamMember{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SpannerDatabaseIamMember)(nil)

// SpannerDatabaseIamMember represents the Terraform resource google_spanner_database_iam_member.
type SpannerDatabaseIamMember struct {
	Name      string
	Args      SpannerDatabaseIamMemberArgs
	state     *spannerDatabaseIamMemberState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SpannerDatabaseIamMember].
func (sdim *SpannerDatabaseIamMember) Type() string {
	return "google_spanner_database_iam_member"
}

// LocalName returns the local name for [SpannerDatabaseIamMember].
func (sdim *SpannerDatabaseIamMember) LocalName() string {
	return sdim.Name
}

// Configuration returns the configuration (args) for [SpannerDatabaseIamMember].
func (sdim *SpannerDatabaseIamMember) Configuration() interface{} {
	return sdim.Args
}

// DependOn is used for other resources to depend on [SpannerDatabaseIamMember].
func (sdim *SpannerDatabaseIamMember) DependOn() terra.Reference {
	return terra.ReferenceResource(sdim)
}

// Dependencies returns the list of resources [SpannerDatabaseIamMember] depends_on.
func (sdim *SpannerDatabaseIamMember) Dependencies() terra.Dependencies {
	return sdim.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SpannerDatabaseIamMember].
func (sdim *SpannerDatabaseIamMember) LifecycleManagement() *terra.Lifecycle {
	return sdim.Lifecycle
}

// Attributes returns the attributes for [SpannerDatabaseIamMember].
func (sdim *SpannerDatabaseIamMember) Attributes() spannerDatabaseIamMemberAttributes {
	return spannerDatabaseIamMemberAttributes{ref: terra.ReferenceResource(sdim)}
}

// ImportState imports the given attribute values into [SpannerDatabaseIamMember]'s state.
func (sdim *SpannerDatabaseIamMember) ImportState(av io.Reader) error {
	sdim.state = &spannerDatabaseIamMemberState{}
	if err := json.NewDecoder(av).Decode(sdim.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sdim.Type(), sdim.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SpannerDatabaseIamMember] has state.
func (sdim *SpannerDatabaseIamMember) State() (*spannerDatabaseIamMemberState, bool) {
	return sdim.state, sdim.state != nil
}

// StateMust returns the state for [SpannerDatabaseIamMember]. Panics if the state is nil.
func (sdim *SpannerDatabaseIamMember) StateMust() *spannerDatabaseIamMemberState {
	if sdim.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sdim.Type(), sdim.LocalName()))
	}
	return sdim.state
}

// SpannerDatabaseIamMemberArgs contains the configurations for google_spanner_database_iam_member.
type SpannerDatabaseIamMemberArgs struct {
	// Database: string, required
	Database terra.StringValue `hcl:"database,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Instance: string, required
	Instance terra.StringValue `hcl:"instance,attr" validate:"required"`
	// Member: string, required
	Member terra.StringValue `hcl:"member,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// Condition: optional
	Condition *spannerdatabaseiammember.Condition `hcl:"condition,block"`
}
type spannerDatabaseIamMemberAttributes struct {
	ref terra.Reference
}

// Database returns a reference to field database of google_spanner_database_iam_member.
func (sdim spannerDatabaseIamMemberAttributes) Database() terra.StringValue {
	return terra.ReferenceAsString(sdim.ref.Append("database"))
}

// Etag returns a reference to field etag of google_spanner_database_iam_member.
func (sdim spannerDatabaseIamMemberAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(sdim.ref.Append("etag"))
}

// Id returns a reference to field id of google_spanner_database_iam_member.
func (sdim spannerDatabaseIamMemberAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sdim.ref.Append("id"))
}

// Instance returns a reference to field instance of google_spanner_database_iam_member.
func (sdim spannerDatabaseIamMemberAttributes) Instance() terra.StringValue {
	return terra.ReferenceAsString(sdim.ref.Append("instance"))
}

// Member returns a reference to field member of google_spanner_database_iam_member.
func (sdim spannerDatabaseIamMemberAttributes) Member() terra.StringValue {
	return terra.ReferenceAsString(sdim.ref.Append("member"))
}

// Project returns a reference to field project of google_spanner_database_iam_member.
func (sdim spannerDatabaseIamMemberAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(sdim.ref.Append("project"))
}

// Role returns a reference to field role of google_spanner_database_iam_member.
func (sdim spannerDatabaseIamMemberAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(sdim.ref.Append("role"))
}

func (sdim spannerDatabaseIamMemberAttributes) Condition() terra.ListValue[spannerdatabaseiammember.ConditionAttributes] {
	return terra.ReferenceAsList[spannerdatabaseiammember.ConditionAttributes](sdim.ref.Append("condition"))
}

type spannerDatabaseIamMemberState struct {
	Database  string                                    `json:"database"`
	Etag      string                                    `json:"etag"`
	Id        string                                    `json:"id"`
	Instance  string                                    `json:"instance"`
	Member    string                                    `json:"member"`
	Project   string                                    `json:"project"`
	Role      string                                    `json:"role"`
	Condition []spannerdatabaseiammember.ConditionState `json:"condition"`
}
