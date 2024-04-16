// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_spanner_database_iam_binding

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// New creates a new instance of [Resource].
func New(name string, args Args) *Resource {
	return &Resource{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Resource)(nil)

// Resource represents the Terraform resource google_spanner_database_iam_binding.
type Resource struct {
	Name      string
	Args      Args
	state     *googleSpannerDatabaseIamBindingState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (gsdib *Resource) Type() string {
	return "google_spanner_database_iam_binding"
}

// LocalName returns the local name for [Resource].
func (gsdib *Resource) LocalName() string {
	return gsdib.Name
}

// Configuration returns the configuration (args) for [Resource].
func (gsdib *Resource) Configuration() interface{} {
	return gsdib.Args
}

// DependOn is used for other resources to depend on [Resource].
func (gsdib *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(gsdib)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (gsdib *Resource) Dependencies() terra.Dependencies {
	return gsdib.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (gsdib *Resource) LifecycleManagement() *terra.Lifecycle {
	return gsdib.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (gsdib *Resource) Attributes() googleSpannerDatabaseIamBindingAttributes {
	return googleSpannerDatabaseIamBindingAttributes{ref: terra.ReferenceResource(gsdib)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (gsdib *Resource) ImportState(state io.Reader) error {
	gsdib.state = &googleSpannerDatabaseIamBindingState{}
	if err := json.NewDecoder(state).Decode(gsdib.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", gsdib.Type(), gsdib.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (gsdib *Resource) State() (*googleSpannerDatabaseIamBindingState, bool) {
	return gsdib.state, gsdib.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (gsdib *Resource) StateMust() *googleSpannerDatabaseIamBindingState {
	if gsdib.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", gsdib.Type(), gsdib.LocalName()))
	}
	return gsdib.state
}

// Args contains the configurations for google_spanner_database_iam_binding.
type Args struct {
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
	Condition *Condition `hcl:"condition,block"`
}

type googleSpannerDatabaseIamBindingAttributes struct {
	ref terra.Reference
}

// Database returns a reference to field database of google_spanner_database_iam_binding.
func (gsdib googleSpannerDatabaseIamBindingAttributes) Database() terra.StringValue {
	return terra.ReferenceAsString(gsdib.ref.Append("database"))
}

// Etag returns a reference to field etag of google_spanner_database_iam_binding.
func (gsdib googleSpannerDatabaseIamBindingAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(gsdib.ref.Append("etag"))
}

// Id returns a reference to field id of google_spanner_database_iam_binding.
func (gsdib googleSpannerDatabaseIamBindingAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gsdib.ref.Append("id"))
}

// Instance returns a reference to field instance of google_spanner_database_iam_binding.
func (gsdib googleSpannerDatabaseIamBindingAttributes) Instance() terra.StringValue {
	return terra.ReferenceAsString(gsdib.ref.Append("instance"))
}

// Members returns a reference to field members of google_spanner_database_iam_binding.
func (gsdib googleSpannerDatabaseIamBindingAttributes) Members() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](gsdib.ref.Append("members"))
}

// Project returns a reference to field project of google_spanner_database_iam_binding.
func (gsdib googleSpannerDatabaseIamBindingAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(gsdib.ref.Append("project"))
}

// Role returns a reference to field role of google_spanner_database_iam_binding.
func (gsdib googleSpannerDatabaseIamBindingAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(gsdib.ref.Append("role"))
}

func (gsdib googleSpannerDatabaseIamBindingAttributes) Condition() terra.ListValue[ConditionAttributes] {
	return terra.ReferenceAsList[ConditionAttributes](gsdib.ref.Append("condition"))
}

type googleSpannerDatabaseIamBindingState struct {
	Database  string           `json:"database"`
	Etag      string           `json:"etag"`
	Id        string           `json:"id"`
	Instance  string           `json:"instance"`
	Members   []string         `json:"members"`
	Project   string           `json:"project"`
	Role      string           `json:"role"`
	Condition []ConditionState `json:"condition"`
}
