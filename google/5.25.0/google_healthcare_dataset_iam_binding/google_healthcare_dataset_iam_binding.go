// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_healthcare_dataset_iam_binding

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

// Resource represents the Terraform resource google_healthcare_dataset_iam_binding.
type Resource struct {
	Name      string
	Args      Args
	state     *googleHealthcareDatasetIamBindingState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (ghdib *Resource) Type() string {
	return "google_healthcare_dataset_iam_binding"
}

// LocalName returns the local name for [Resource].
func (ghdib *Resource) LocalName() string {
	return ghdib.Name
}

// Configuration returns the configuration (args) for [Resource].
func (ghdib *Resource) Configuration() interface{} {
	return ghdib.Args
}

// DependOn is used for other resources to depend on [Resource].
func (ghdib *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(ghdib)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (ghdib *Resource) Dependencies() terra.Dependencies {
	return ghdib.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (ghdib *Resource) LifecycleManagement() *terra.Lifecycle {
	return ghdib.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (ghdib *Resource) Attributes() googleHealthcareDatasetIamBindingAttributes {
	return googleHealthcareDatasetIamBindingAttributes{ref: terra.ReferenceResource(ghdib)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (ghdib *Resource) ImportState(state io.Reader) error {
	ghdib.state = &googleHealthcareDatasetIamBindingState{}
	if err := json.NewDecoder(state).Decode(ghdib.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ghdib.Type(), ghdib.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (ghdib *Resource) State() (*googleHealthcareDatasetIamBindingState, bool) {
	return ghdib.state, ghdib.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (ghdib *Resource) StateMust() *googleHealthcareDatasetIamBindingState {
	if ghdib.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ghdib.Type(), ghdib.LocalName()))
	}
	return ghdib.state
}

// Args contains the configurations for google_healthcare_dataset_iam_binding.
type Args struct {
	// DatasetId: string, required
	DatasetId terra.StringValue `hcl:"dataset_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Members: set of string, required
	Members terra.SetValue[terra.StringValue] `hcl:"members,attr" validate:"required"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// Condition: optional
	Condition *Condition `hcl:"condition,block"`
}

type googleHealthcareDatasetIamBindingAttributes struct {
	ref terra.Reference
}

// DatasetId returns a reference to field dataset_id of google_healthcare_dataset_iam_binding.
func (ghdib googleHealthcareDatasetIamBindingAttributes) DatasetId() terra.StringValue {
	return terra.ReferenceAsString(ghdib.ref.Append("dataset_id"))
}

// Etag returns a reference to field etag of google_healthcare_dataset_iam_binding.
func (ghdib googleHealthcareDatasetIamBindingAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(ghdib.ref.Append("etag"))
}

// Id returns a reference to field id of google_healthcare_dataset_iam_binding.
func (ghdib googleHealthcareDatasetIamBindingAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ghdib.ref.Append("id"))
}

// Members returns a reference to field members of google_healthcare_dataset_iam_binding.
func (ghdib googleHealthcareDatasetIamBindingAttributes) Members() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](ghdib.ref.Append("members"))
}

// Role returns a reference to field role of google_healthcare_dataset_iam_binding.
func (ghdib googleHealthcareDatasetIamBindingAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(ghdib.ref.Append("role"))
}

func (ghdib googleHealthcareDatasetIamBindingAttributes) Condition() terra.ListValue[ConditionAttributes] {
	return terra.ReferenceAsList[ConditionAttributes](ghdib.ref.Append("condition"))
}

type googleHealthcareDatasetIamBindingState struct {
	DatasetId string           `json:"dataset_id"`
	Etag      string           `json:"etag"`
	Id        string           `json:"id"`
	Members   []string         `json:"members"`
	Role      string           `json:"role"`
	Condition []ConditionState `json:"condition"`
}
