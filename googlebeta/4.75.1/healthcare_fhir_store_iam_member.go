// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	healthcarefhirstoreiammember "github.com/golingon/terraproviders/googlebeta/4.75.1/healthcarefhirstoreiammember"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewHealthcareFhirStoreIamMember creates a new instance of [HealthcareFhirStoreIamMember].
func NewHealthcareFhirStoreIamMember(name string, args HealthcareFhirStoreIamMemberArgs) *HealthcareFhirStoreIamMember {
	return &HealthcareFhirStoreIamMember{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*HealthcareFhirStoreIamMember)(nil)

// HealthcareFhirStoreIamMember represents the Terraform resource google_healthcare_fhir_store_iam_member.
type HealthcareFhirStoreIamMember struct {
	Name      string
	Args      HealthcareFhirStoreIamMemberArgs
	state     *healthcareFhirStoreIamMemberState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [HealthcareFhirStoreIamMember].
func (hfsim *HealthcareFhirStoreIamMember) Type() string {
	return "google_healthcare_fhir_store_iam_member"
}

// LocalName returns the local name for [HealthcareFhirStoreIamMember].
func (hfsim *HealthcareFhirStoreIamMember) LocalName() string {
	return hfsim.Name
}

// Configuration returns the configuration (args) for [HealthcareFhirStoreIamMember].
func (hfsim *HealthcareFhirStoreIamMember) Configuration() interface{} {
	return hfsim.Args
}

// DependOn is used for other resources to depend on [HealthcareFhirStoreIamMember].
func (hfsim *HealthcareFhirStoreIamMember) DependOn() terra.Reference {
	return terra.ReferenceResource(hfsim)
}

// Dependencies returns the list of resources [HealthcareFhirStoreIamMember] depends_on.
func (hfsim *HealthcareFhirStoreIamMember) Dependencies() terra.Dependencies {
	return hfsim.DependsOn
}

// LifecycleManagement returns the lifecycle block for [HealthcareFhirStoreIamMember].
func (hfsim *HealthcareFhirStoreIamMember) LifecycleManagement() *terra.Lifecycle {
	return hfsim.Lifecycle
}

// Attributes returns the attributes for [HealthcareFhirStoreIamMember].
func (hfsim *HealthcareFhirStoreIamMember) Attributes() healthcareFhirStoreIamMemberAttributes {
	return healthcareFhirStoreIamMemberAttributes{ref: terra.ReferenceResource(hfsim)}
}

// ImportState imports the given attribute values into [HealthcareFhirStoreIamMember]'s state.
func (hfsim *HealthcareFhirStoreIamMember) ImportState(av io.Reader) error {
	hfsim.state = &healthcareFhirStoreIamMemberState{}
	if err := json.NewDecoder(av).Decode(hfsim.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", hfsim.Type(), hfsim.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [HealthcareFhirStoreIamMember] has state.
func (hfsim *HealthcareFhirStoreIamMember) State() (*healthcareFhirStoreIamMemberState, bool) {
	return hfsim.state, hfsim.state != nil
}

// StateMust returns the state for [HealthcareFhirStoreIamMember]. Panics if the state is nil.
func (hfsim *HealthcareFhirStoreIamMember) StateMust() *healthcareFhirStoreIamMemberState {
	if hfsim.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", hfsim.Type(), hfsim.LocalName()))
	}
	return hfsim.state
}

// HealthcareFhirStoreIamMemberArgs contains the configurations for google_healthcare_fhir_store_iam_member.
type HealthcareFhirStoreIamMemberArgs struct {
	// FhirStoreId: string, required
	FhirStoreId terra.StringValue `hcl:"fhir_store_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Member: string, required
	Member terra.StringValue `hcl:"member,attr" validate:"required"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// Condition: optional
	Condition *healthcarefhirstoreiammember.Condition `hcl:"condition,block"`
}
type healthcareFhirStoreIamMemberAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_healthcare_fhir_store_iam_member.
func (hfsim healthcareFhirStoreIamMemberAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(hfsim.ref.Append("etag"))
}

// FhirStoreId returns a reference to field fhir_store_id of google_healthcare_fhir_store_iam_member.
func (hfsim healthcareFhirStoreIamMemberAttributes) FhirStoreId() terra.StringValue {
	return terra.ReferenceAsString(hfsim.ref.Append("fhir_store_id"))
}

// Id returns a reference to field id of google_healthcare_fhir_store_iam_member.
func (hfsim healthcareFhirStoreIamMemberAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(hfsim.ref.Append("id"))
}

// Member returns a reference to field member of google_healthcare_fhir_store_iam_member.
func (hfsim healthcareFhirStoreIamMemberAttributes) Member() terra.StringValue {
	return terra.ReferenceAsString(hfsim.ref.Append("member"))
}

// Role returns a reference to field role of google_healthcare_fhir_store_iam_member.
func (hfsim healthcareFhirStoreIamMemberAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(hfsim.ref.Append("role"))
}

func (hfsim healthcareFhirStoreIamMemberAttributes) Condition() terra.ListValue[healthcarefhirstoreiammember.ConditionAttributes] {
	return terra.ReferenceAsList[healthcarefhirstoreiammember.ConditionAttributes](hfsim.ref.Append("condition"))
}

type healthcareFhirStoreIamMemberState struct {
	Etag        string                                        `json:"etag"`
	FhirStoreId string                                        `json:"fhir_store_id"`
	Id          string                                        `json:"id"`
	Member      string                                        `json:"member"`
	Role        string                                        `json:"role"`
	Condition   []healthcarefhirstoreiammember.ConditionState `json:"condition"`
}
