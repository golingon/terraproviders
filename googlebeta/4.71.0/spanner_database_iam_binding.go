// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	spannerdatabaseiambinding "github.com/golingon/terraproviders/googlebeta/4.71.0/spannerdatabaseiambinding"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSpannerDatabaseIamBinding creates a new instance of [SpannerDatabaseIamBinding].
func NewSpannerDatabaseIamBinding(name string, args SpannerDatabaseIamBindingArgs) *SpannerDatabaseIamBinding {
	return &SpannerDatabaseIamBinding{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SpannerDatabaseIamBinding)(nil)

// SpannerDatabaseIamBinding represents the Terraform resource google_spanner_database_iam_binding.
type SpannerDatabaseIamBinding struct {
	Name      string
	Args      SpannerDatabaseIamBindingArgs
	state     *spannerDatabaseIamBindingState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SpannerDatabaseIamBinding].
func (sdib *SpannerDatabaseIamBinding) Type() string {
	return "google_spanner_database_iam_binding"
}

// LocalName returns the local name for [SpannerDatabaseIamBinding].
func (sdib *SpannerDatabaseIamBinding) LocalName() string {
	return sdib.Name
}

// Configuration returns the configuration (args) for [SpannerDatabaseIamBinding].
func (sdib *SpannerDatabaseIamBinding) Configuration() interface{} {
	return sdib.Args
}

// DependOn is used for other resources to depend on [SpannerDatabaseIamBinding].
func (sdib *SpannerDatabaseIamBinding) DependOn() terra.Reference {
	return terra.ReferenceResource(sdib)
}

// Dependencies returns the list of resources [SpannerDatabaseIamBinding] depends_on.
func (sdib *SpannerDatabaseIamBinding) Dependencies() terra.Dependencies {
	return sdib.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SpannerDatabaseIamBinding].
func (sdib *SpannerDatabaseIamBinding) LifecycleManagement() *terra.Lifecycle {
	return sdib.Lifecycle
}

// Attributes returns the attributes for [SpannerDatabaseIamBinding].
func (sdib *SpannerDatabaseIamBinding) Attributes() spannerDatabaseIamBindingAttributes {
	return spannerDatabaseIamBindingAttributes{ref: terra.ReferenceResource(sdib)}
}

// ImportState imports the given attribute values into [SpannerDatabaseIamBinding]'s state.
func (sdib *SpannerDatabaseIamBinding) ImportState(av io.Reader) error {
	sdib.state = &spannerDatabaseIamBindingState{}
	if err := json.NewDecoder(av).Decode(sdib.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sdib.Type(), sdib.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SpannerDatabaseIamBinding] has state.
func (sdib *SpannerDatabaseIamBinding) State() (*spannerDatabaseIamBindingState, bool) {
	return sdib.state, sdib.state != nil
}

// StateMust returns the state for [SpannerDatabaseIamBinding]. Panics if the state is nil.
func (sdib *SpannerDatabaseIamBinding) StateMust() *spannerDatabaseIamBindingState {
	if sdib.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sdib.Type(), sdib.LocalName()))
	}
	return sdib.state
}

// SpannerDatabaseIamBindingArgs contains the configurations for google_spanner_database_iam_binding.
type SpannerDatabaseIamBindingArgs struct {
	// Database: string, required
	Database terra.StringValue `hcl:"database,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Instance: string, required
	Instance terra.StringValue `hcl:"instance,attr" validate:"required"`
	// Members: set of string, required
	Members terra.SetValue[terra.StringValue] `hcl:"members,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// Condition: optional
	Condition *spannerdatabaseiambinding.Condition `hcl:"condition,block"`
}
type spannerDatabaseIamBindingAttributes struct {
	ref terra.Reference
}

// Database returns a reference to field database of google_spanner_database_iam_binding.
func (sdib spannerDatabaseIamBindingAttributes) Database() terra.StringValue {
	return terra.ReferenceAsString(sdib.ref.Append("database"))
}

// Etag returns a reference to field etag of google_spanner_database_iam_binding.
func (sdib spannerDatabaseIamBindingAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(sdib.ref.Append("etag"))
}

// Id returns a reference to field id of google_spanner_database_iam_binding.
func (sdib spannerDatabaseIamBindingAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sdib.ref.Append("id"))
}

// Instance returns a reference to field instance of google_spanner_database_iam_binding.
func (sdib spannerDatabaseIamBindingAttributes) Instance() terra.StringValue {
	return terra.ReferenceAsString(sdib.ref.Append("instance"))
}

// Members returns a reference to field members of google_spanner_database_iam_binding.
func (sdib spannerDatabaseIamBindingAttributes) Members() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](sdib.ref.Append("members"))
}

// Project returns a reference to field project of google_spanner_database_iam_binding.
func (sdib spannerDatabaseIamBindingAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(sdib.ref.Append("project"))
}

// Role returns a reference to field role of google_spanner_database_iam_binding.
func (sdib spannerDatabaseIamBindingAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(sdib.ref.Append("role"))
}

func (sdib spannerDatabaseIamBindingAttributes) Condition() terra.ListValue[spannerdatabaseiambinding.ConditionAttributes] {
	return terra.ReferenceAsList[spannerdatabaseiambinding.ConditionAttributes](sdib.ref.Append("condition"))
}

type spannerDatabaseIamBindingState struct {
	Database  string                                     `json:"database"`
	Etag      string                                     `json:"etag"`
	Id        string                                     `json:"id"`
	Instance  string                                     `json:"instance"`
	Members   []string                                   `json:"members"`
	Project   string                                     `json:"project"`
	Role      string                                     `json:"role"`
	Condition []spannerdatabaseiambinding.ConditionState `json:"condition"`
}
