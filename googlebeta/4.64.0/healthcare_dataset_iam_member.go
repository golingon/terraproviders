// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	healthcaredatasetiammember "github.com/golingon/terraproviders/googlebeta/4.64.0/healthcaredatasetiammember"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewHealthcareDatasetIamMember creates a new instance of [HealthcareDatasetIamMember].
func NewHealthcareDatasetIamMember(name string, args HealthcareDatasetIamMemberArgs) *HealthcareDatasetIamMember {
	return &HealthcareDatasetIamMember{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*HealthcareDatasetIamMember)(nil)

// HealthcareDatasetIamMember represents the Terraform resource google_healthcare_dataset_iam_member.
type HealthcareDatasetIamMember struct {
	Name      string
	Args      HealthcareDatasetIamMemberArgs
	state     *healthcareDatasetIamMemberState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [HealthcareDatasetIamMember].
func (hdim *HealthcareDatasetIamMember) Type() string {
	return "google_healthcare_dataset_iam_member"
}

// LocalName returns the local name for [HealthcareDatasetIamMember].
func (hdim *HealthcareDatasetIamMember) LocalName() string {
	return hdim.Name
}

// Configuration returns the configuration (args) for [HealthcareDatasetIamMember].
func (hdim *HealthcareDatasetIamMember) Configuration() interface{} {
	return hdim.Args
}

// DependOn is used for other resources to depend on [HealthcareDatasetIamMember].
func (hdim *HealthcareDatasetIamMember) DependOn() terra.Reference {
	return terra.ReferenceResource(hdim)
}

// Dependencies returns the list of resources [HealthcareDatasetIamMember] depends_on.
func (hdim *HealthcareDatasetIamMember) Dependencies() terra.Dependencies {
	return hdim.DependsOn
}

// LifecycleManagement returns the lifecycle block for [HealthcareDatasetIamMember].
func (hdim *HealthcareDatasetIamMember) LifecycleManagement() *terra.Lifecycle {
	return hdim.Lifecycle
}

// Attributes returns the attributes for [HealthcareDatasetIamMember].
func (hdim *HealthcareDatasetIamMember) Attributes() healthcareDatasetIamMemberAttributes {
	return healthcareDatasetIamMemberAttributes{ref: terra.ReferenceResource(hdim)}
}

// ImportState imports the given attribute values into [HealthcareDatasetIamMember]'s state.
func (hdim *HealthcareDatasetIamMember) ImportState(av io.Reader) error {
	hdim.state = &healthcareDatasetIamMemberState{}
	if err := json.NewDecoder(av).Decode(hdim.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", hdim.Type(), hdim.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [HealthcareDatasetIamMember] has state.
func (hdim *HealthcareDatasetIamMember) State() (*healthcareDatasetIamMemberState, bool) {
	return hdim.state, hdim.state != nil
}

// StateMust returns the state for [HealthcareDatasetIamMember]. Panics if the state is nil.
func (hdim *HealthcareDatasetIamMember) StateMust() *healthcareDatasetIamMemberState {
	if hdim.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", hdim.Type(), hdim.LocalName()))
	}
	return hdim.state
}

// HealthcareDatasetIamMemberArgs contains the configurations for google_healthcare_dataset_iam_member.
type HealthcareDatasetIamMemberArgs struct {
	// DatasetId: string, required
	DatasetId terra.StringValue `hcl:"dataset_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Member: string, required
	Member terra.StringValue `hcl:"member,attr" validate:"required"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// Condition: optional
	Condition *healthcaredatasetiammember.Condition `hcl:"condition,block"`
}
type healthcareDatasetIamMemberAttributes struct {
	ref terra.Reference
}

// DatasetId returns a reference to field dataset_id of google_healthcare_dataset_iam_member.
func (hdim healthcareDatasetIamMemberAttributes) DatasetId() terra.StringValue {
	return terra.ReferenceAsString(hdim.ref.Append("dataset_id"))
}

// Etag returns a reference to field etag of google_healthcare_dataset_iam_member.
func (hdim healthcareDatasetIamMemberAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(hdim.ref.Append("etag"))
}

// Id returns a reference to field id of google_healthcare_dataset_iam_member.
func (hdim healthcareDatasetIamMemberAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(hdim.ref.Append("id"))
}

// Member returns a reference to field member of google_healthcare_dataset_iam_member.
func (hdim healthcareDatasetIamMemberAttributes) Member() terra.StringValue {
	return terra.ReferenceAsString(hdim.ref.Append("member"))
}

// Role returns a reference to field role of google_healthcare_dataset_iam_member.
func (hdim healthcareDatasetIamMemberAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(hdim.ref.Append("role"))
}

func (hdim healthcareDatasetIamMemberAttributes) Condition() terra.ListValue[healthcaredatasetiammember.ConditionAttributes] {
	return terra.ReferenceAsList[healthcaredatasetiammember.ConditionAttributes](hdim.ref.Append("condition"))
}

type healthcareDatasetIamMemberState struct {
	DatasetId string                                      `json:"dataset_id"`
	Etag      string                                      `json:"etag"`
	Id        string                                      `json:"id"`
	Member    string                                      `json:"member"`
	Role      string                                      `json:"role"`
	Condition []healthcaredatasetiammember.ConditionState `json:"condition"`
}
