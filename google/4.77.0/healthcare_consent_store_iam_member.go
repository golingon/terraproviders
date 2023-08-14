// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	healthcareconsentstoreiammember "github.com/golingon/terraproviders/google/4.77.0/healthcareconsentstoreiammember"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewHealthcareConsentStoreIamMember creates a new instance of [HealthcareConsentStoreIamMember].
func NewHealthcareConsentStoreIamMember(name string, args HealthcareConsentStoreIamMemberArgs) *HealthcareConsentStoreIamMember {
	return &HealthcareConsentStoreIamMember{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*HealthcareConsentStoreIamMember)(nil)

// HealthcareConsentStoreIamMember represents the Terraform resource google_healthcare_consent_store_iam_member.
type HealthcareConsentStoreIamMember struct {
	Name      string
	Args      HealthcareConsentStoreIamMemberArgs
	state     *healthcareConsentStoreIamMemberState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [HealthcareConsentStoreIamMember].
func (hcsim *HealthcareConsentStoreIamMember) Type() string {
	return "google_healthcare_consent_store_iam_member"
}

// LocalName returns the local name for [HealthcareConsentStoreIamMember].
func (hcsim *HealthcareConsentStoreIamMember) LocalName() string {
	return hcsim.Name
}

// Configuration returns the configuration (args) for [HealthcareConsentStoreIamMember].
func (hcsim *HealthcareConsentStoreIamMember) Configuration() interface{} {
	return hcsim.Args
}

// DependOn is used for other resources to depend on [HealthcareConsentStoreIamMember].
func (hcsim *HealthcareConsentStoreIamMember) DependOn() terra.Reference {
	return terra.ReferenceResource(hcsim)
}

// Dependencies returns the list of resources [HealthcareConsentStoreIamMember] depends_on.
func (hcsim *HealthcareConsentStoreIamMember) Dependencies() terra.Dependencies {
	return hcsim.DependsOn
}

// LifecycleManagement returns the lifecycle block for [HealthcareConsentStoreIamMember].
func (hcsim *HealthcareConsentStoreIamMember) LifecycleManagement() *terra.Lifecycle {
	return hcsim.Lifecycle
}

// Attributes returns the attributes for [HealthcareConsentStoreIamMember].
func (hcsim *HealthcareConsentStoreIamMember) Attributes() healthcareConsentStoreIamMemberAttributes {
	return healthcareConsentStoreIamMemberAttributes{ref: terra.ReferenceResource(hcsim)}
}

// ImportState imports the given attribute values into [HealthcareConsentStoreIamMember]'s state.
func (hcsim *HealthcareConsentStoreIamMember) ImportState(av io.Reader) error {
	hcsim.state = &healthcareConsentStoreIamMemberState{}
	if err := json.NewDecoder(av).Decode(hcsim.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", hcsim.Type(), hcsim.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [HealthcareConsentStoreIamMember] has state.
func (hcsim *HealthcareConsentStoreIamMember) State() (*healthcareConsentStoreIamMemberState, bool) {
	return hcsim.state, hcsim.state != nil
}

// StateMust returns the state for [HealthcareConsentStoreIamMember]. Panics if the state is nil.
func (hcsim *HealthcareConsentStoreIamMember) StateMust() *healthcareConsentStoreIamMemberState {
	if hcsim.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", hcsim.Type(), hcsim.LocalName()))
	}
	return hcsim.state
}

// HealthcareConsentStoreIamMemberArgs contains the configurations for google_healthcare_consent_store_iam_member.
type HealthcareConsentStoreIamMemberArgs struct {
	// ConsentStoreId: string, required
	ConsentStoreId terra.StringValue `hcl:"consent_store_id,attr" validate:"required"`
	// Dataset: string, required
	Dataset terra.StringValue `hcl:"dataset,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Member: string, required
	Member terra.StringValue `hcl:"member,attr" validate:"required"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// Condition: optional
	Condition *healthcareconsentstoreiammember.Condition `hcl:"condition,block"`
}
type healthcareConsentStoreIamMemberAttributes struct {
	ref terra.Reference
}

// ConsentStoreId returns a reference to field consent_store_id of google_healthcare_consent_store_iam_member.
func (hcsim healthcareConsentStoreIamMemberAttributes) ConsentStoreId() terra.StringValue {
	return terra.ReferenceAsString(hcsim.ref.Append("consent_store_id"))
}

// Dataset returns a reference to field dataset of google_healthcare_consent_store_iam_member.
func (hcsim healthcareConsentStoreIamMemberAttributes) Dataset() terra.StringValue {
	return terra.ReferenceAsString(hcsim.ref.Append("dataset"))
}

// Etag returns a reference to field etag of google_healthcare_consent_store_iam_member.
func (hcsim healthcareConsentStoreIamMemberAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(hcsim.ref.Append("etag"))
}

// Id returns a reference to field id of google_healthcare_consent_store_iam_member.
func (hcsim healthcareConsentStoreIamMemberAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(hcsim.ref.Append("id"))
}

// Member returns a reference to field member of google_healthcare_consent_store_iam_member.
func (hcsim healthcareConsentStoreIamMemberAttributes) Member() terra.StringValue {
	return terra.ReferenceAsString(hcsim.ref.Append("member"))
}

// Role returns a reference to field role of google_healthcare_consent_store_iam_member.
func (hcsim healthcareConsentStoreIamMemberAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(hcsim.ref.Append("role"))
}

func (hcsim healthcareConsentStoreIamMemberAttributes) Condition() terra.ListValue[healthcareconsentstoreiammember.ConditionAttributes] {
	return terra.ReferenceAsList[healthcareconsentstoreiammember.ConditionAttributes](hcsim.ref.Append("condition"))
}

type healthcareConsentStoreIamMemberState struct {
	ConsentStoreId string                                           `json:"consent_store_id"`
	Dataset        string                                           `json:"dataset"`
	Etag           string                                           `json:"etag"`
	Id             string                                           `json:"id"`
	Member         string                                           `json:"member"`
	Role           string                                           `json:"role"`
	Condition      []healthcareconsentstoreiammember.ConditionState `json:"condition"`
}
