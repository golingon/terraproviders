// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	healthcaredicomstoreiambinding "github.com/golingon/terraproviders/google/4.71.0/healthcaredicomstoreiambinding"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewHealthcareDicomStoreIamBinding creates a new instance of [HealthcareDicomStoreIamBinding].
func NewHealthcareDicomStoreIamBinding(name string, args HealthcareDicomStoreIamBindingArgs) *HealthcareDicomStoreIamBinding {
	return &HealthcareDicomStoreIamBinding{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*HealthcareDicomStoreIamBinding)(nil)

// HealthcareDicomStoreIamBinding represents the Terraform resource google_healthcare_dicom_store_iam_binding.
type HealthcareDicomStoreIamBinding struct {
	Name      string
	Args      HealthcareDicomStoreIamBindingArgs
	state     *healthcareDicomStoreIamBindingState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [HealthcareDicomStoreIamBinding].
func (hdsib *HealthcareDicomStoreIamBinding) Type() string {
	return "google_healthcare_dicom_store_iam_binding"
}

// LocalName returns the local name for [HealthcareDicomStoreIamBinding].
func (hdsib *HealthcareDicomStoreIamBinding) LocalName() string {
	return hdsib.Name
}

// Configuration returns the configuration (args) for [HealthcareDicomStoreIamBinding].
func (hdsib *HealthcareDicomStoreIamBinding) Configuration() interface{} {
	return hdsib.Args
}

// DependOn is used for other resources to depend on [HealthcareDicomStoreIamBinding].
func (hdsib *HealthcareDicomStoreIamBinding) DependOn() terra.Reference {
	return terra.ReferenceResource(hdsib)
}

// Dependencies returns the list of resources [HealthcareDicomStoreIamBinding] depends_on.
func (hdsib *HealthcareDicomStoreIamBinding) Dependencies() terra.Dependencies {
	return hdsib.DependsOn
}

// LifecycleManagement returns the lifecycle block for [HealthcareDicomStoreIamBinding].
func (hdsib *HealthcareDicomStoreIamBinding) LifecycleManagement() *terra.Lifecycle {
	return hdsib.Lifecycle
}

// Attributes returns the attributes for [HealthcareDicomStoreIamBinding].
func (hdsib *HealthcareDicomStoreIamBinding) Attributes() healthcareDicomStoreIamBindingAttributes {
	return healthcareDicomStoreIamBindingAttributes{ref: terra.ReferenceResource(hdsib)}
}

// ImportState imports the given attribute values into [HealthcareDicomStoreIamBinding]'s state.
func (hdsib *HealthcareDicomStoreIamBinding) ImportState(av io.Reader) error {
	hdsib.state = &healthcareDicomStoreIamBindingState{}
	if err := json.NewDecoder(av).Decode(hdsib.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", hdsib.Type(), hdsib.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [HealthcareDicomStoreIamBinding] has state.
func (hdsib *HealthcareDicomStoreIamBinding) State() (*healthcareDicomStoreIamBindingState, bool) {
	return hdsib.state, hdsib.state != nil
}

// StateMust returns the state for [HealthcareDicomStoreIamBinding]. Panics if the state is nil.
func (hdsib *HealthcareDicomStoreIamBinding) StateMust() *healthcareDicomStoreIamBindingState {
	if hdsib.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", hdsib.Type(), hdsib.LocalName()))
	}
	return hdsib.state
}

// HealthcareDicomStoreIamBindingArgs contains the configurations for google_healthcare_dicom_store_iam_binding.
type HealthcareDicomStoreIamBindingArgs struct {
	// DicomStoreId: string, required
	DicomStoreId terra.StringValue `hcl:"dicom_store_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Members: set of string, required
	Members terra.SetValue[terra.StringValue] `hcl:"members,attr" validate:"required"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// Condition: optional
	Condition *healthcaredicomstoreiambinding.Condition `hcl:"condition,block"`
}
type healthcareDicomStoreIamBindingAttributes struct {
	ref terra.Reference
}

// DicomStoreId returns a reference to field dicom_store_id of google_healthcare_dicom_store_iam_binding.
func (hdsib healthcareDicomStoreIamBindingAttributes) DicomStoreId() terra.StringValue {
	return terra.ReferenceAsString(hdsib.ref.Append("dicom_store_id"))
}

// Etag returns a reference to field etag of google_healthcare_dicom_store_iam_binding.
func (hdsib healthcareDicomStoreIamBindingAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(hdsib.ref.Append("etag"))
}

// Id returns a reference to field id of google_healthcare_dicom_store_iam_binding.
func (hdsib healthcareDicomStoreIamBindingAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(hdsib.ref.Append("id"))
}

// Members returns a reference to field members of google_healthcare_dicom_store_iam_binding.
func (hdsib healthcareDicomStoreIamBindingAttributes) Members() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](hdsib.ref.Append("members"))
}

// Role returns a reference to field role of google_healthcare_dicom_store_iam_binding.
func (hdsib healthcareDicomStoreIamBindingAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(hdsib.ref.Append("role"))
}

func (hdsib healthcareDicomStoreIamBindingAttributes) Condition() terra.ListValue[healthcaredicomstoreiambinding.ConditionAttributes] {
	return terra.ReferenceAsList[healthcaredicomstoreiambinding.ConditionAttributes](hdsib.ref.Append("condition"))
}

type healthcareDicomStoreIamBindingState struct {
	DicomStoreId string                                          `json:"dicom_store_id"`
	Etag         string                                          `json:"etag"`
	Id           string                                          `json:"id"`
	Members      []string                                        `json:"members"`
	Role         string                                          `json:"role"`
	Condition    []healthcaredicomstoreiambinding.ConditionState `json:"condition"`
}
