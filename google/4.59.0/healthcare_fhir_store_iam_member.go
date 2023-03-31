// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	healthcarefhirstoreiammember "github.com/golingon/terraproviders/google/4.59.0/healthcarefhirstoreiammember"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewHealthcareFhirStoreIamMember(name string, args HealthcareFhirStoreIamMemberArgs) *HealthcareFhirStoreIamMember {
	return &HealthcareFhirStoreIamMember{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*HealthcareFhirStoreIamMember)(nil)

type HealthcareFhirStoreIamMember struct {
	Name  string
	Args  HealthcareFhirStoreIamMemberArgs
	state *healthcareFhirStoreIamMemberState
}

func (hfsim *HealthcareFhirStoreIamMember) Type() string {
	return "google_healthcare_fhir_store_iam_member"
}

func (hfsim *HealthcareFhirStoreIamMember) LocalName() string {
	return hfsim.Name
}

func (hfsim *HealthcareFhirStoreIamMember) Configuration() interface{} {
	return hfsim.Args
}

func (hfsim *HealthcareFhirStoreIamMember) Attributes() healthcareFhirStoreIamMemberAttributes {
	return healthcareFhirStoreIamMemberAttributes{ref: terra.ReferenceResource(hfsim)}
}

func (hfsim *HealthcareFhirStoreIamMember) ImportState(av io.Reader) error {
	hfsim.state = &healthcareFhirStoreIamMemberState{}
	if err := json.NewDecoder(av).Decode(hfsim.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", hfsim.Type(), hfsim.LocalName(), err)
	}
	return nil
}

func (hfsim *HealthcareFhirStoreIamMember) State() (*healthcareFhirStoreIamMemberState, bool) {
	return hfsim.state, hfsim.state != nil
}

func (hfsim *HealthcareFhirStoreIamMember) StateMust() *healthcareFhirStoreIamMemberState {
	if hfsim.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", hfsim.Type(), hfsim.LocalName()))
	}
	return hfsim.state
}

func (hfsim *HealthcareFhirStoreIamMember) DependOn() terra.Reference {
	return terra.ReferenceResource(hfsim)
}

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
	// DependsOn contains resources that HealthcareFhirStoreIamMember depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type healthcareFhirStoreIamMemberAttributes struct {
	ref terra.Reference
}

func (hfsim healthcareFhirStoreIamMemberAttributes) Etag() terra.StringValue {
	return terra.ReferenceString(hfsim.ref.Append("etag"))
}

func (hfsim healthcareFhirStoreIamMemberAttributes) FhirStoreId() terra.StringValue {
	return terra.ReferenceString(hfsim.ref.Append("fhir_store_id"))
}

func (hfsim healthcareFhirStoreIamMemberAttributes) Id() terra.StringValue {
	return terra.ReferenceString(hfsim.ref.Append("id"))
}

func (hfsim healthcareFhirStoreIamMemberAttributes) Member() terra.StringValue {
	return terra.ReferenceString(hfsim.ref.Append("member"))
}

func (hfsim healthcareFhirStoreIamMemberAttributes) Role() terra.StringValue {
	return terra.ReferenceString(hfsim.ref.Append("role"))
}

func (hfsim healthcareFhirStoreIamMemberAttributes) Condition() terra.ListValue[healthcarefhirstoreiammember.ConditionAttributes] {
	return terra.ReferenceList[healthcarefhirstoreiammember.ConditionAttributes](hfsim.ref.Append("condition"))
}

type healthcareFhirStoreIamMemberState struct {
	Etag        string                                        `json:"etag"`
	FhirStoreId string                                        `json:"fhir_store_id"`
	Id          string                                        `json:"id"`
	Member      string                                        `json:"member"`
	Role        string                                        `json:"role"`
	Condition   []healthcarefhirstoreiammember.ConditionState `json:"condition"`
}
