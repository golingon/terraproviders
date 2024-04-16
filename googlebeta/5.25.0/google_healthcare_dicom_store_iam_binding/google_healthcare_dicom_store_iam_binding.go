// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_healthcare_dicom_store_iam_binding

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

// Resource represents the Terraform resource google_healthcare_dicom_store_iam_binding.
type Resource struct {
	Name      string
	Args      Args
	state     *googleHealthcareDicomStoreIamBindingState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (ghdsib *Resource) Type() string {
	return "google_healthcare_dicom_store_iam_binding"
}

// LocalName returns the local name for [Resource].
func (ghdsib *Resource) LocalName() string {
	return ghdsib.Name
}

// Configuration returns the configuration (args) for [Resource].
func (ghdsib *Resource) Configuration() interface{} {
	return ghdsib.Args
}

// DependOn is used for other resources to depend on [Resource].
func (ghdsib *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(ghdsib)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (ghdsib *Resource) Dependencies() terra.Dependencies {
	return ghdsib.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (ghdsib *Resource) LifecycleManagement() *terra.Lifecycle {
	return ghdsib.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (ghdsib *Resource) Attributes() googleHealthcareDicomStoreIamBindingAttributes {
	return googleHealthcareDicomStoreIamBindingAttributes{ref: terra.ReferenceResource(ghdsib)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (ghdsib *Resource) ImportState(state io.Reader) error {
	ghdsib.state = &googleHealthcareDicomStoreIamBindingState{}
	if err := json.NewDecoder(state).Decode(ghdsib.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ghdsib.Type(), ghdsib.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (ghdsib *Resource) State() (*googleHealthcareDicomStoreIamBindingState, bool) {
	return ghdsib.state, ghdsib.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (ghdsib *Resource) StateMust() *googleHealthcareDicomStoreIamBindingState {
	if ghdsib.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ghdsib.Type(), ghdsib.LocalName()))
	}
	return ghdsib.state
}

// Args contains the configurations for google_healthcare_dicom_store_iam_binding.
type Args struct {
	// DicomStoreId: string, required
	DicomStoreId terra.StringValue `hcl:"dicom_store_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Members: set of string, required
	Members terra.SetValue[terra.StringValue] `hcl:"members,attr" validate:"required"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// Condition: optional
	Condition *Condition `hcl:"condition,block"`
}

type googleHealthcareDicomStoreIamBindingAttributes struct {
	ref terra.Reference
}

// DicomStoreId returns a reference to field dicom_store_id of google_healthcare_dicom_store_iam_binding.
func (ghdsib googleHealthcareDicomStoreIamBindingAttributes) DicomStoreId() terra.StringValue {
	return terra.ReferenceAsString(ghdsib.ref.Append("dicom_store_id"))
}

// Etag returns a reference to field etag of google_healthcare_dicom_store_iam_binding.
func (ghdsib googleHealthcareDicomStoreIamBindingAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(ghdsib.ref.Append("etag"))
}

// Id returns a reference to field id of google_healthcare_dicom_store_iam_binding.
func (ghdsib googleHealthcareDicomStoreIamBindingAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ghdsib.ref.Append("id"))
}

// Members returns a reference to field members of google_healthcare_dicom_store_iam_binding.
func (ghdsib googleHealthcareDicomStoreIamBindingAttributes) Members() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](ghdsib.ref.Append("members"))
}

// Role returns a reference to field role of google_healthcare_dicom_store_iam_binding.
func (ghdsib googleHealthcareDicomStoreIamBindingAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(ghdsib.ref.Append("role"))
}

func (ghdsib googleHealthcareDicomStoreIamBindingAttributes) Condition() terra.ListValue[ConditionAttributes] {
	return terra.ReferenceAsList[ConditionAttributes](ghdsib.ref.Append("condition"))
}

type googleHealthcareDicomStoreIamBindingState struct {
	DicomStoreId string           `json:"dicom_store_id"`
	Etag         string           `json:"etag"`
	Id           string           `json:"id"`
	Members      []string         `json:"members"`
	Role         string           `json:"role"`
	Condition    []ConditionState `json:"condition"`
}
