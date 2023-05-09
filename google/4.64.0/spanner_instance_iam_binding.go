// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	spannerinstanceiambinding "github.com/golingon/terraproviders/google/4.64.0/spannerinstanceiambinding"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSpannerInstanceIamBinding creates a new instance of [SpannerInstanceIamBinding].
func NewSpannerInstanceIamBinding(name string, args SpannerInstanceIamBindingArgs) *SpannerInstanceIamBinding {
	return &SpannerInstanceIamBinding{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SpannerInstanceIamBinding)(nil)

// SpannerInstanceIamBinding represents the Terraform resource google_spanner_instance_iam_binding.
type SpannerInstanceIamBinding struct {
	Name      string
	Args      SpannerInstanceIamBindingArgs
	state     *spannerInstanceIamBindingState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SpannerInstanceIamBinding].
func (siib *SpannerInstanceIamBinding) Type() string {
	return "google_spanner_instance_iam_binding"
}

// LocalName returns the local name for [SpannerInstanceIamBinding].
func (siib *SpannerInstanceIamBinding) LocalName() string {
	return siib.Name
}

// Configuration returns the configuration (args) for [SpannerInstanceIamBinding].
func (siib *SpannerInstanceIamBinding) Configuration() interface{} {
	return siib.Args
}

// DependOn is used for other resources to depend on [SpannerInstanceIamBinding].
func (siib *SpannerInstanceIamBinding) DependOn() terra.Reference {
	return terra.ReferenceResource(siib)
}

// Dependencies returns the list of resources [SpannerInstanceIamBinding] depends_on.
func (siib *SpannerInstanceIamBinding) Dependencies() terra.Dependencies {
	return siib.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SpannerInstanceIamBinding].
func (siib *SpannerInstanceIamBinding) LifecycleManagement() *terra.Lifecycle {
	return siib.Lifecycle
}

// Attributes returns the attributes for [SpannerInstanceIamBinding].
func (siib *SpannerInstanceIamBinding) Attributes() spannerInstanceIamBindingAttributes {
	return spannerInstanceIamBindingAttributes{ref: terra.ReferenceResource(siib)}
}

// ImportState imports the given attribute values into [SpannerInstanceIamBinding]'s state.
func (siib *SpannerInstanceIamBinding) ImportState(av io.Reader) error {
	siib.state = &spannerInstanceIamBindingState{}
	if err := json.NewDecoder(av).Decode(siib.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", siib.Type(), siib.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SpannerInstanceIamBinding] has state.
func (siib *SpannerInstanceIamBinding) State() (*spannerInstanceIamBindingState, bool) {
	return siib.state, siib.state != nil
}

// StateMust returns the state for [SpannerInstanceIamBinding]. Panics if the state is nil.
func (siib *SpannerInstanceIamBinding) StateMust() *spannerInstanceIamBindingState {
	if siib.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", siib.Type(), siib.LocalName()))
	}
	return siib.state
}

// SpannerInstanceIamBindingArgs contains the configurations for google_spanner_instance_iam_binding.
type SpannerInstanceIamBindingArgs struct {
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
	Condition *spannerinstanceiambinding.Condition `hcl:"condition,block"`
}
type spannerInstanceIamBindingAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_spanner_instance_iam_binding.
func (siib spannerInstanceIamBindingAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(siib.ref.Append("etag"))
}

// Id returns a reference to field id of google_spanner_instance_iam_binding.
func (siib spannerInstanceIamBindingAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(siib.ref.Append("id"))
}

// Instance returns a reference to field instance of google_spanner_instance_iam_binding.
func (siib spannerInstanceIamBindingAttributes) Instance() terra.StringValue {
	return terra.ReferenceAsString(siib.ref.Append("instance"))
}

// Members returns a reference to field members of google_spanner_instance_iam_binding.
func (siib spannerInstanceIamBindingAttributes) Members() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](siib.ref.Append("members"))
}

// Project returns a reference to field project of google_spanner_instance_iam_binding.
func (siib spannerInstanceIamBindingAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(siib.ref.Append("project"))
}

// Role returns a reference to field role of google_spanner_instance_iam_binding.
func (siib spannerInstanceIamBindingAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(siib.ref.Append("role"))
}

func (siib spannerInstanceIamBindingAttributes) Condition() terra.ListValue[spannerinstanceiambinding.ConditionAttributes] {
	return terra.ReferenceAsList[spannerinstanceiambinding.ConditionAttributes](siib.ref.Append("condition"))
}

type spannerInstanceIamBindingState struct {
	Etag      string                                     `json:"etag"`
	Id        string                                     `json:"id"`
	Instance  string                                     `json:"instance"`
	Members   []string                                   `json:"members"`
	Project   string                                     `json:"project"`
	Role      string                                     `json:"role"`
	Condition []spannerinstanceiambinding.ConditionState `json:"condition"`
}
