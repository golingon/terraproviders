// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	healthcaredicomstoreiammember "github.com/golingon/terraproviders/google/4.74.0/healthcaredicomstoreiammember"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewHealthcareDicomStoreIamMember creates a new instance of [HealthcareDicomStoreIamMember].
func NewHealthcareDicomStoreIamMember(name string, args HealthcareDicomStoreIamMemberArgs) *HealthcareDicomStoreIamMember {
	return &HealthcareDicomStoreIamMember{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*HealthcareDicomStoreIamMember)(nil)

// HealthcareDicomStoreIamMember represents the Terraform resource google_healthcare_dicom_store_iam_member.
type HealthcareDicomStoreIamMember struct {
	Name      string
	Args      HealthcareDicomStoreIamMemberArgs
	state     *healthcareDicomStoreIamMemberState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [HealthcareDicomStoreIamMember].
func (hdsim *HealthcareDicomStoreIamMember) Type() string {
	return "google_healthcare_dicom_store_iam_member"
}

// LocalName returns the local name for [HealthcareDicomStoreIamMember].
func (hdsim *HealthcareDicomStoreIamMember) LocalName() string {
	return hdsim.Name
}

// Configuration returns the configuration (args) for [HealthcareDicomStoreIamMember].
func (hdsim *HealthcareDicomStoreIamMember) Configuration() interface{} {
	return hdsim.Args
}

// DependOn is used for other resources to depend on [HealthcareDicomStoreIamMember].
func (hdsim *HealthcareDicomStoreIamMember) DependOn() terra.Reference {
	return terra.ReferenceResource(hdsim)
}

// Dependencies returns the list of resources [HealthcareDicomStoreIamMember] depends_on.
func (hdsim *HealthcareDicomStoreIamMember) Dependencies() terra.Dependencies {
	return hdsim.DependsOn
}

// LifecycleManagement returns the lifecycle block for [HealthcareDicomStoreIamMember].
func (hdsim *HealthcareDicomStoreIamMember) LifecycleManagement() *terra.Lifecycle {
	return hdsim.Lifecycle
}

// Attributes returns the attributes for [HealthcareDicomStoreIamMember].
func (hdsim *HealthcareDicomStoreIamMember) Attributes() healthcareDicomStoreIamMemberAttributes {
	return healthcareDicomStoreIamMemberAttributes{ref: terra.ReferenceResource(hdsim)}
}

// ImportState imports the given attribute values into [HealthcareDicomStoreIamMember]'s state.
func (hdsim *HealthcareDicomStoreIamMember) ImportState(av io.Reader) error {
	hdsim.state = &healthcareDicomStoreIamMemberState{}
	if err := json.NewDecoder(av).Decode(hdsim.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", hdsim.Type(), hdsim.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [HealthcareDicomStoreIamMember] has state.
func (hdsim *HealthcareDicomStoreIamMember) State() (*healthcareDicomStoreIamMemberState, bool) {
	return hdsim.state, hdsim.state != nil
}

// StateMust returns the state for [HealthcareDicomStoreIamMember]. Panics if the state is nil.
func (hdsim *HealthcareDicomStoreIamMember) StateMust() *healthcareDicomStoreIamMemberState {
	if hdsim.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", hdsim.Type(), hdsim.LocalName()))
	}
	return hdsim.state
}

// HealthcareDicomStoreIamMemberArgs contains the configurations for google_healthcare_dicom_store_iam_member.
type HealthcareDicomStoreIamMemberArgs struct {
	// DicomStoreId: string, required
	DicomStoreId terra.StringValue `hcl:"dicom_store_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Member: string, required
	Member terra.StringValue `hcl:"member,attr" validate:"required"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// Condition: optional
	Condition *healthcaredicomstoreiammember.Condition `hcl:"condition,block"`
}
type healthcareDicomStoreIamMemberAttributes struct {
	ref terra.Reference
}

// DicomStoreId returns a reference to field dicom_store_id of google_healthcare_dicom_store_iam_member.
func (hdsim healthcareDicomStoreIamMemberAttributes) DicomStoreId() terra.StringValue {
	return terra.ReferenceAsString(hdsim.ref.Append("dicom_store_id"))
}

// Etag returns a reference to field etag of google_healthcare_dicom_store_iam_member.
func (hdsim healthcareDicomStoreIamMemberAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(hdsim.ref.Append("etag"))
}

// Id returns a reference to field id of google_healthcare_dicom_store_iam_member.
func (hdsim healthcareDicomStoreIamMemberAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(hdsim.ref.Append("id"))
}

// Member returns a reference to field member of google_healthcare_dicom_store_iam_member.
func (hdsim healthcareDicomStoreIamMemberAttributes) Member() terra.StringValue {
	return terra.ReferenceAsString(hdsim.ref.Append("member"))
}

// Role returns a reference to field role of google_healthcare_dicom_store_iam_member.
func (hdsim healthcareDicomStoreIamMemberAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(hdsim.ref.Append("role"))
}

func (hdsim healthcareDicomStoreIamMemberAttributes) Condition() terra.ListValue[healthcaredicomstoreiammember.ConditionAttributes] {
	return terra.ReferenceAsList[healthcaredicomstoreiammember.ConditionAttributes](hdsim.ref.Append("condition"))
}

type healthcareDicomStoreIamMemberState struct {
	DicomStoreId string                                         `json:"dicom_store_id"`
	Etag         string                                         `json:"etag"`
	Id           string                                         `json:"id"`
	Member       string                                         `json:"member"`
	Role         string                                         `json:"role"`
	Condition    []healthcaredicomstoreiammember.ConditionState `json:"condition"`
}