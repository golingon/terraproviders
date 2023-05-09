// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	healthcarehl7v2storeiambinding "github.com/golingon/terraproviders/google/4.63.1/healthcarehl7v2storeiambinding"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewHealthcareHl7V2StoreIamBinding creates a new instance of [HealthcareHl7V2StoreIamBinding].
func NewHealthcareHl7V2StoreIamBinding(name string, args HealthcareHl7V2StoreIamBindingArgs) *HealthcareHl7V2StoreIamBinding {
	return &HealthcareHl7V2StoreIamBinding{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*HealthcareHl7V2StoreIamBinding)(nil)

// HealthcareHl7V2StoreIamBinding represents the Terraform resource google_healthcare_hl7_v2_store_iam_binding.
type HealthcareHl7V2StoreIamBinding struct {
	Name      string
	Args      HealthcareHl7V2StoreIamBindingArgs
	state     *healthcareHl7V2StoreIamBindingState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [HealthcareHl7V2StoreIamBinding].
func (hhvsib *HealthcareHl7V2StoreIamBinding) Type() string {
	return "google_healthcare_hl7_v2_store_iam_binding"
}

// LocalName returns the local name for [HealthcareHl7V2StoreIamBinding].
func (hhvsib *HealthcareHl7V2StoreIamBinding) LocalName() string {
	return hhvsib.Name
}

// Configuration returns the configuration (args) for [HealthcareHl7V2StoreIamBinding].
func (hhvsib *HealthcareHl7V2StoreIamBinding) Configuration() interface{} {
	return hhvsib.Args
}

// DependOn is used for other resources to depend on [HealthcareHl7V2StoreIamBinding].
func (hhvsib *HealthcareHl7V2StoreIamBinding) DependOn() terra.Reference {
	return terra.ReferenceResource(hhvsib)
}

// Dependencies returns the list of resources [HealthcareHl7V2StoreIamBinding] depends_on.
func (hhvsib *HealthcareHl7V2StoreIamBinding) Dependencies() terra.Dependencies {
	return hhvsib.DependsOn
}

// LifecycleManagement returns the lifecycle block for [HealthcareHl7V2StoreIamBinding].
func (hhvsib *HealthcareHl7V2StoreIamBinding) LifecycleManagement() *terra.Lifecycle {
	return hhvsib.Lifecycle
}

// Attributes returns the attributes for [HealthcareHl7V2StoreIamBinding].
func (hhvsib *HealthcareHl7V2StoreIamBinding) Attributes() healthcareHl7V2StoreIamBindingAttributes {
	return healthcareHl7V2StoreIamBindingAttributes{ref: terra.ReferenceResource(hhvsib)}
}

// ImportState imports the given attribute values into [HealthcareHl7V2StoreIamBinding]'s state.
func (hhvsib *HealthcareHl7V2StoreIamBinding) ImportState(av io.Reader) error {
	hhvsib.state = &healthcareHl7V2StoreIamBindingState{}
	if err := json.NewDecoder(av).Decode(hhvsib.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", hhvsib.Type(), hhvsib.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [HealthcareHl7V2StoreIamBinding] has state.
func (hhvsib *HealthcareHl7V2StoreIamBinding) State() (*healthcareHl7V2StoreIamBindingState, bool) {
	return hhvsib.state, hhvsib.state != nil
}

// StateMust returns the state for [HealthcareHl7V2StoreIamBinding]. Panics if the state is nil.
func (hhvsib *HealthcareHl7V2StoreIamBinding) StateMust() *healthcareHl7V2StoreIamBindingState {
	if hhvsib.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", hhvsib.Type(), hhvsib.LocalName()))
	}
	return hhvsib.state
}

// HealthcareHl7V2StoreIamBindingArgs contains the configurations for google_healthcare_hl7_v2_store_iam_binding.
type HealthcareHl7V2StoreIamBindingArgs struct {
	// Hl7V2StoreId: string, required
	Hl7V2StoreId terra.StringValue `hcl:"hl7_v2_store_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Members: set of string, required
	Members terra.SetValue[terra.StringValue] `hcl:"members,attr" validate:"required"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// Condition: optional
	Condition *healthcarehl7v2storeiambinding.Condition `hcl:"condition,block"`
}
type healthcareHl7V2StoreIamBindingAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_healthcare_hl7_v2_store_iam_binding.
func (hhvsib healthcareHl7V2StoreIamBindingAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(hhvsib.ref.Append("etag"))
}

// Hl7V2StoreId returns a reference to field hl7_v2_store_id of google_healthcare_hl7_v2_store_iam_binding.
func (hhvsib healthcareHl7V2StoreIamBindingAttributes) Hl7V2StoreId() terra.StringValue {
	return terra.ReferenceAsString(hhvsib.ref.Append("hl7_v2_store_id"))
}

// Id returns a reference to field id of google_healthcare_hl7_v2_store_iam_binding.
func (hhvsib healthcareHl7V2StoreIamBindingAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(hhvsib.ref.Append("id"))
}

// Members returns a reference to field members of google_healthcare_hl7_v2_store_iam_binding.
func (hhvsib healthcareHl7V2StoreIamBindingAttributes) Members() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](hhvsib.ref.Append("members"))
}

// Role returns a reference to field role of google_healthcare_hl7_v2_store_iam_binding.
func (hhvsib healthcareHl7V2StoreIamBindingAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(hhvsib.ref.Append("role"))
}

func (hhvsib healthcareHl7V2StoreIamBindingAttributes) Condition() terra.ListValue[healthcarehl7v2storeiambinding.ConditionAttributes] {
	return terra.ReferenceAsList[healthcarehl7v2storeiambinding.ConditionAttributes](hhvsib.ref.Append("condition"))
}

type healthcareHl7V2StoreIamBindingState struct {
	Etag         string                                          `json:"etag"`
	Hl7V2StoreId string                                          `json:"hl7_v2_store_id"`
	Id           string                                          `json:"id"`
	Members      []string                                        `json:"members"`
	Role         string                                          `json:"role"`
	Condition    []healthcarehl7v2storeiambinding.ConditionState `json:"condition"`
}
