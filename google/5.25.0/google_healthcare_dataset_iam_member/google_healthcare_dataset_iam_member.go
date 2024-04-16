// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_healthcare_dataset_iam_member

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

// Resource represents the Terraform resource google_healthcare_dataset_iam_member.
type Resource struct {
	Name      string
	Args      Args
	state     *googleHealthcareDatasetIamMemberState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (ghdim *Resource) Type() string {
	return "google_healthcare_dataset_iam_member"
}

// LocalName returns the local name for [Resource].
func (ghdim *Resource) LocalName() string {
	return ghdim.Name
}

// Configuration returns the configuration (args) for [Resource].
func (ghdim *Resource) Configuration() interface{} {
	return ghdim.Args
}

// DependOn is used for other resources to depend on [Resource].
func (ghdim *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(ghdim)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (ghdim *Resource) Dependencies() terra.Dependencies {
	return ghdim.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (ghdim *Resource) LifecycleManagement() *terra.Lifecycle {
	return ghdim.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (ghdim *Resource) Attributes() googleHealthcareDatasetIamMemberAttributes {
	return googleHealthcareDatasetIamMemberAttributes{ref: terra.ReferenceResource(ghdim)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (ghdim *Resource) ImportState(state io.Reader) error {
	ghdim.state = &googleHealthcareDatasetIamMemberState{}
	if err := json.NewDecoder(state).Decode(ghdim.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ghdim.Type(), ghdim.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (ghdim *Resource) State() (*googleHealthcareDatasetIamMemberState, bool) {
	return ghdim.state, ghdim.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (ghdim *Resource) StateMust() *googleHealthcareDatasetIamMemberState {
	if ghdim.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ghdim.Type(), ghdim.LocalName()))
	}
	return ghdim.state
}

// Args contains the configurations for google_healthcare_dataset_iam_member.
type Args struct {
	// DatasetId: string, required
	DatasetId terra.StringValue `hcl:"dataset_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Member: string, required
	Member terra.StringValue `hcl:"member,attr" validate:"required"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// Condition: optional
	Condition *Condition `hcl:"condition,block"`
}

type googleHealthcareDatasetIamMemberAttributes struct {
	ref terra.Reference
}

// DatasetId returns a reference to field dataset_id of google_healthcare_dataset_iam_member.
func (ghdim googleHealthcareDatasetIamMemberAttributes) DatasetId() terra.StringValue {
	return terra.ReferenceAsString(ghdim.ref.Append("dataset_id"))
}

// Etag returns a reference to field etag of google_healthcare_dataset_iam_member.
func (ghdim googleHealthcareDatasetIamMemberAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(ghdim.ref.Append("etag"))
}

// Id returns a reference to field id of google_healthcare_dataset_iam_member.
func (ghdim googleHealthcareDatasetIamMemberAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ghdim.ref.Append("id"))
}

// Member returns a reference to field member of google_healthcare_dataset_iam_member.
func (ghdim googleHealthcareDatasetIamMemberAttributes) Member() terra.StringValue {
	return terra.ReferenceAsString(ghdim.ref.Append("member"))
}

// Role returns a reference to field role of google_healthcare_dataset_iam_member.
func (ghdim googleHealthcareDatasetIamMemberAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(ghdim.ref.Append("role"))
}

func (ghdim googleHealthcareDatasetIamMemberAttributes) Condition() terra.ListValue[ConditionAttributes] {
	return terra.ReferenceAsList[ConditionAttributes](ghdim.ref.Append("condition"))
}

type googleHealthcareDatasetIamMemberState struct {
	DatasetId string           `json:"dataset_id"`
	Etag      string           `json:"etag"`
	Id        string           `json:"id"`
	Member    string           `json:"member"`
	Role      string           `json:"role"`
	Condition []ConditionState `json:"condition"`
}
