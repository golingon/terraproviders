// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	healthcaredatasetiambinding "github.com/golingon/terraproviders/googlebeta/4.76.0/healthcaredatasetiambinding"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewHealthcareDatasetIamBinding creates a new instance of [HealthcareDatasetIamBinding].
func NewHealthcareDatasetIamBinding(name string, args HealthcareDatasetIamBindingArgs) *HealthcareDatasetIamBinding {
	return &HealthcareDatasetIamBinding{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*HealthcareDatasetIamBinding)(nil)

// HealthcareDatasetIamBinding represents the Terraform resource google_healthcare_dataset_iam_binding.
type HealthcareDatasetIamBinding struct {
	Name      string
	Args      HealthcareDatasetIamBindingArgs
	state     *healthcareDatasetIamBindingState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [HealthcareDatasetIamBinding].
func (hdib *HealthcareDatasetIamBinding) Type() string {
	return "google_healthcare_dataset_iam_binding"
}

// LocalName returns the local name for [HealthcareDatasetIamBinding].
func (hdib *HealthcareDatasetIamBinding) LocalName() string {
	return hdib.Name
}

// Configuration returns the configuration (args) for [HealthcareDatasetIamBinding].
func (hdib *HealthcareDatasetIamBinding) Configuration() interface{} {
	return hdib.Args
}

// DependOn is used for other resources to depend on [HealthcareDatasetIamBinding].
func (hdib *HealthcareDatasetIamBinding) DependOn() terra.Reference {
	return terra.ReferenceResource(hdib)
}

// Dependencies returns the list of resources [HealthcareDatasetIamBinding] depends_on.
func (hdib *HealthcareDatasetIamBinding) Dependencies() terra.Dependencies {
	return hdib.DependsOn
}

// LifecycleManagement returns the lifecycle block for [HealthcareDatasetIamBinding].
func (hdib *HealthcareDatasetIamBinding) LifecycleManagement() *terra.Lifecycle {
	return hdib.Lifecycle
}

// Attributes returns the attributes for [HealthcareDatasetIamBinding].
func (hdib *HealthcareDatasetIamBinding) Attributes() healthcareDatasetIamBindingAttributes {
	return healthcareDatasetIamBindingAttributes{ref: terra.ReferenceResource(hdib)}
}

// ImportState imports the given attribute values into [HealthcareDatasetIamBinding]'s state.
func (hdib *HealthcareDatasetIamBinding) ImportState(av io.Reader) error {
	hdib.state = &healthcareDatasetIamBindingState{}
	if err := json.NewDecoder(av).Decode(hdib.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", hdib.Type(), hdib.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [HealthcareDatasetIamBinding] has state.
func (hdib *HealthcareDatasetIamBinding) State() (*healthcareDatasetIamBindingState, bool) {
	return hdib.state, hdib.state != nil
}

// StateMust returns the state for [HealthcareDatasetIamBinding]. Panics if the state is nil.
func (hdib *HealthcareDatasetIamBinding) StateMust() *healthcareDatasetIamBindingState {
	if hdib.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", hdib.Type(), hdib.LocalName()))
	}
	return hdib.state
}

// HealthcareDatasetIamBindingArgs contains the configurations for google_healthcare_dataset_iam_binding.
type HealthcareDatasetIamBindingArgs struct {
	// DatasetId: string, required
	DatasetId terra.StringValue `hcl:"dataset_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Members: set of string, required
	Members terra.SetValue[terra.StringValue] `hcl:"members,attr" validate:"required"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// Condition: optional
	Condition *healthcaredatasetiambinding.Condition `hcl:"condition,block"`
}
type healthcareDatasetIamBindingAttributes struct {
	ref terra.Reference
}

// DatasetId returns a reference to field dataset_id of google_healthcare_dataset_iam_binding.
func (hdib healthcareDatasetIamBindingAttributes) DatasetId() terra.StringValue {
	return terra.ReferenceAsString(hdib.ref.Append("dataset_id"))
}

// Etag returns a reference to field etag of google_healthcare_dataset_iam_binding.
func (hdib healthcareDatasetIamBindingAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(hdib.ref.Append("etag"))
}

// Id returns a reference to field id of google_healthcare_dataset_iam_binding.
func (hdib healthcareDatasetIamBindingAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(hdib.ref.Append("id"))
}

// Members returns a reference to field members of google_healthcare_dataset_iam_binding.
func (hdib healthcareDatasetIamBindingAttributes) Members() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](hdib.ref.Append("members"))
}

// Role returns a reference to field role of google_healthcare_dataset_iam_binding.
func (hdib healthcareDatasetIamBindingAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(hdib.ref.Append("role"))
}

func (hdib healthcareDatasetIamBindingAttributes) Condition() terra.ListValue[healthcaredatasetiambinding.ConditionAttributes] {
	return terra.ReferenceAsList[healthcaredatasetiambinding.ConditionAttributes](hdib.ref.Append("condition"))
}

type healthcareDatasetIamBindingState struct {
	DatasetId string                                       `json:"dataset_id"`
	Etag      string                                       `json:"etag"`
	Id        string                                       `json:"id"`
	Members   []string                                     `json:"members"`
	Role      string                                       `json:"role"`
	Condition []healthcaredatasetiambinding.ConditionState `json:"condition"`
}
