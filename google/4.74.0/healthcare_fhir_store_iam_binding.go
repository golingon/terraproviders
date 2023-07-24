// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	healthcarefhirstoreiambinding "github.com/golingon/terraproviders/google/4.74.0/healthcarefhirstoreiambinding"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewHealthcareFhirStoreIamBinding creates a new instance of [HealthcareFhirStoreIamBinding].
func NewHealthcareFhirStoreIamBinding(name string, args HealthcareFhirStoreIamBindingArgs) *HealthcareFhirStoreIamBinding {
	return &HealthcareFhirStoreIamBinding{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*HealthcareFhirStoreIamBinding)(nil)

// HealthcareFhirStoreIamBinding represents the Terraform resource google_healthcare_fhir_store_iam_binding.
type HealthcareFhirStoreIamBinding struct {
	Name      string
	Args      HealthcareFhirStoreIamBindingArgs
	state     *healthcareFhirStoreIamBindingState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [HealthcareFhirStoreIamBinding].
func (hfsib *HealthcareFhirStoreIamBinding) Type() string {
	return "google_healthcare_fhir_store_iam_binding"
}

// LocalName returns the local name for [HealthcareFhirStoreIamBinding].
func (hfsib *HealthcareFhirStoreIamBinding) LocalName() string {
	return hfsib.Name
}

// Configuration returns the configuration (args) for [HealthcareFhirStoreIamBinding].
func (hfsib *HealthcareFhirStoreIamBinding) Configuration() interface{} {
	return hfsib.Args
}

// DependOn is used for other resources to depend on [HealthcareFhirStoreIamBinding].
func (hfsib *HealthcareFhirStoreIamBinding) DependOn() terra.Reference {
	return terra.ReferenceResource(hfsib)
}

// Dependencies returns the list of resources [HealthcareFhirStoreIamBinding] depends_on.
func (hfsib *HealthcareFhirStoreIamBinding) Dependencies() terra.Dependencies {
	return hfsib.DependsOn
}

// LifecycleManagement returns the lifecycle block for [HealthcareFhirStoreIamBinding].
func (hfsib *HealthcareFhirStoreIamBinding) LifecycleManagement() *terra.Lifecycle {
	return hfsib.Lifecycle
}

// Attributes returns the attributes for [HealthcareFhirStoreIamBinding].
func (hfsib *HealthcareFhirStoreIamBinding) Attributes() healthcareFhirStoreIamBindingAttributes {
	return healthcareFhirStoreIamBindingAttributes{ref: terra.ReferenceResource(hfsib)}
}

// ImportState imports the given attribute values into [HealthcareFhirStoreIamBinding]'s state.
func (hfsib *HealthcareFhirStoreIamBinding) ImportState(av io.Reader) error {
	hfsib.state = &healthcareFhirStoreIamBindingState{}
	if err := json.NewDecoder(av).Decode(hfsib.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", hfsib.Type(), hfsib.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [HealthcareFhirStoreIamBinding] has state.
func (hfsib *HealthcareFhirStoreIamBinding) State() (*healthcareFhirStoreIamBindingState, bool) {
	return hfsib.state, hfsib.state != nil
}

// StateMust returns the state for [HealthcareFhirStoreIamBinding]. Panics if the state is nil.
func (hfsib *HealthcareFhirStoreIamBinding) StateMust() *healthcareFhirStoreIamBindingState {
	if hfsib.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", hfsib.Type(), hfsib.LocalName()))
	}
	return hfsib.state
}

// HealthcareFhirStoreIamBindingArgs contains the configurations for google_healthcare_fhir_store_iam_binding.
type HealthcareFhirStoreIamBindingArgs struct {
	// FhirStoreId: string, required
	FhirStoreId terra.StringValue `hcl:"fhir_store_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Members: set of string, required
	Members terra.SetValue[terra.StringValue] `hcl:"members,attr" validate:"required"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// Condition: optional
	Condition *healthcarefhirstoreiambinding.Condition `hcl:"condition,block"`
}
type healthcareFhirStoreIamBindingAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_healthcare_fhir_store_iam_binding.
func (hfsib healthcareFhirStoreIamBindingAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(hfsib.ref.Append("etag"))
}

// FhirStoreId returns a reference to field fhir_store_id of google_healthcare_fhir_store_iam_binding.
func (hfsib healthcareFhirStoreIamBindingAttributes) FhirStoreId() terra.StringValue {
	return terra.ReferenceAsString(hfsib.ref.Append("fhir_store_id"))
}

// Id returns a reference to field id of google_healthcare_fhir_store_iam_binding.
func (hfsib healthcareFhirStoreIamBindingAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(hfsib.ref.Append("id"))
}

// Members returns a reference to field members of google_healthcare_fhir_store_iam_binding.
func (hfsib healthcareFhirStoreIamBindingAttributes) Members() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](hfsib.ref.Append("members"))
}

// Role returns a reference to field role of google_healthcare_fhir_store_iam_binding.
func (hfsib healthcareFhirStoreIamBindingAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(hfsib.ref.Append("role"))
}

func (hfsib healthcareFhirStoreIamBindingAttributes) Condition() terra.ListValue[healthcarefhirstoreiambinding.ConditionAttributes] {
	return terra.ReferenceAsList[healthcarefhirstoreiambinding.ConditionAttributes](hfsib.ref.Append("condition"))
}

type healthcareFhirStoreIamBindingState struct {
	Etag        string                                         `json:"etag"`
	FhirStoreId string                                         `json:"fhir_store_id"`
	Id          string                                         `json:"id"`
	Members     []string                                       `json:"members"`
	Role        string                                         `json:"role"`
	Condition   []healthcarefhirstoreiambinding.ConditionState `json:"condition"`
}
