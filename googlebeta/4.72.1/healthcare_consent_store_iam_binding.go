// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	healthcareconsentstoreiambinding "github.com/golingon/terraproviders/googlebeta/4.72.1/healthcareconsentstoreiambinding"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewHealthcareConsentStoreIamBinding creates a new instance of [HealthcareConsentStoreIamBinding].
func NewHealthcareConsentStoreIamBinding(name string, args HealthcareConsentStoreIamBindingArgs) *HealthcareConsentStoreIamBinding {
	return &HealthcareConsentStoreIamBinding{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*HealthcareConsentStoreIamBinding)(nil)

// HealthcareConsentStoreIamBinding represents the Terraform resource google_healthcare_consent_store_iam_binding.
type HealthcareConsentStoreIamBinding struct {
	Name      string
	Args      HealthcareConsentStoreIamBindingArgs
	state     *healthcareConsentStoreIamBindingState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [HealthcareConsentStoreIamBinding].
func (hcsib *HealthcareConsentStoreIamBinding) Type() string {
	return "google_healthcare_consent_store_iam_binding"
}

// LocalName returns the local name for [HealthcareConsentStoreIamBinding].
func (hcsib *HealthcareConsentStoreIamBinding) LocalName() string {
	return hcsib.Name
}

// Configuration returns the configuration (args) for [HealthcareConsentStoreIamBinding].
func (hcsib *HealthcareConsentStoreIamBinding) Configuration() interface{} {
	return hcsib.Args
}

// DependOn is used for other resources to depend on [HealthcareConsentStoreIamBinding].
func (hcsib *HealthcareConsentStoreIamBinding) DependOn() terra.Reference {
	return terra.ReferenceResource(hcsib)
}

// Dependencies returns the list of resources [HealthcareConsentStoreIamBinding] depends_on.
func (hcsib *HealthcareConsentStoreIamBinding) Dependencies() terra.Dependencies {
	return hcsib.DependsOn
}

// LifecycleManagement returns the lifecycle block for [HealthcareConsentStoreIamBinding].
func (hcsib *HealthcareConsentStoreIamBinding) LifecycleManagement() *terra.Lifecycle {
	return hcsib.Lifecycle
}

// Attributes returns the attributes for [HealthcareConsentStoreIamBinding].
func (hcsib *HealthcareConsentStoreIamBinding) Attributes() healthcareConsentStoreIamBindingAttributes {
	return healthcareConsentStoreIamBindingAttributes{ref: terra.ReferenceResource(hcsib)}
}

// ImportState imports the given attribute values into [HealthcareConsentStoreIamBinding]'s state.
func (hcsib *HealthcareConsentStoreIamBinding) ImportState(av io.Reader) error {
	hcsib.state = &healthcareConsentStoreIamBindingState{}
	if err := json.NewDecoder(av).Decode(hcsib.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", hcsib.Type(), hcsib.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [HealthcareConsentStoreIamBinding] has state.
func (hcsib *HealthcareConsentStoreIamBinding) State() (*healthcareConsentStoreIamBindingState, bool) {
	return hcsib.state, hcsib.state != nil
}

// StateMust returns the state for [HealthcareConsentStoreIamBinding]. Panics if the state is nil.
func (hcsib *HealthcareConsentStoreIamBinding) StateMust() *healthcareConsentStoreIamBindingState {
	if hcsib.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", hcsib.Type(), hcsib.LocalName()))
	}
	return hcsib.state
}

// HealthcareConsentStoreIamBindingArgs contains the configurations for google_healthcare_consent_store_iam_binding.
type HealthcareConsentStoreIamBindingArgs struct {
	// ConsentStoreId: string, required
	ConsentStoreId terra.StringValue `hcl:"consent_store_id,attr" validate:"required"`
	// Dataset: string, required
	Dataset terra.StringValue `hcl:"dataset,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Members: set of string, required
	Members terra.SetValue[terra.StringValue] `hcl:"members,attr" validate:"required"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// Condition: optional
	Condition *healthcareconsentstoreiambinding.Condition `hcl:"condition,block"`
}
type healthcareConsentStoreIamBindingAttributes struct {
	ref terra.Reference
}

// ConsentStoreId returns a reference to field consent_store_id of google_healthcare_consent_store_iam_binding.
func (hcsib healthcareConsentStoreIamBindingAttributes) ConsentStoreId() terra.StringValue {
	return terra.ReferenceAsString(hcsib.ref.Append("consent_store_id"))
}

// Dataset returns a reference to field dataset of google_healthcare_consent_store_iam_binding.
func (hcsib healthcareConsentStoreIamBindingAttributes) Dataset() terra.StringValue {
	return terra.ReferenceAsString(hcsib.ref.Append("dataset"))
}

// Etag returns a reference to field etag of google_healthcare_consent_store_iam_binding.
func (hcsib healthcareConsentStoreIamBindingAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(hcsib.ref.Append("etag"))
}

// Id returns a reference to field id of google_healthcare_consent_store_iam_binding.
func (hcsib healthcareConsentStoreIamBindingAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(hcsib.ref.Append("id"))
}

// Members returns a reference to field members of google_healthcare_consent_store_iam_binding.
func (hcsib healthcareConsentStoreIamBindingAttributes) Members() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](hcsib.ref.Append("members"))
}

// Role returns a reference to field role of google_healthcare_consent_store_iam_binding.
func (hcsib healthcareConsentStoreIamBindingAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(hcsib.ref.Append("role"))
}

func (hcsib healthcareConsentStoreIamBindingAttributes) Condition() terra.ListValue[healthcareconsentstoreiambinding.ConditionAttributes] {
	return terra.ReferenceAsList[healthcareconsentstoreiambinding.ConditionAttributes](hcsib.ref.Append("condition"))
}

type healthcareConsentStoreIamBindingState struct {
	ConsentStoreId string                                            `json:"consent_store_id"`
	Dataset        string                                            `json:"dataset"`
	Etag           string                                            `json:"etag"`
	Id             string                                            `json:"id"`
	Members        []string                                          `json:"members"`
	Role           string                                            `json:"role"`
	Condition      []healthcareconsentstoreiambinding.ConditionState `json:"condition"`
}
