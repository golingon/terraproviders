// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	billingaccountiammember "github.com/golingon/terraproviders/google/4.59.0/billingaccountiammember"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewBillingAccountIamMember(name string, args BillingAccountIamMemberArgs) *BillingAccountIamMember {
	return &BillingAccountIamMember{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*BillingAccountIamMember)(nil)

type BillingAccountIamMember struct {
	Name  string
	Args  BillingAccountIamMemberArgs
	state *billingAccountIamMemberState
}

func (baim *BillingAccountIamMember) Type() string {
	return "google_billing_account_iam_member"
}

func (baim *BillingAccountIamMember) LocalName() string {
	return baim.Name
}

func (baim *BillingAccountIamMember) Configuration() interface{} {
	return baim.Args
}

func (baim *BillingAccountIamMember) Attributes() billingAccountIamMemberAttributes {
	return billingAccountIamMemberAttributes{ref: terra.ReferenceResource(baim)}
}

func (baim *BillingAccountIamMember) ImportState(av io.Reader) error {
	baim.state = &billingAccountIamMemberState{}
	if err := json.NewDecoder(av).Decode(baim.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", baim.Type(), baim.LocalName(), err)
	}
	return nil
}

func (baim *BillingAccountIamMember) State() (*billingAccountIamMemberState, bool) {
	return baim.state, baim.state != nil
}

func (baim *BillingAccountIamMember) StateMust() *billingAccountIamMemberState {
	if baim.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", baim.Type(), baim.LocalName()))
	}
	return baim.state
}

func (baim *BillingAccountIamMember) DependOn() terra.Reference {
	return terra.ReferenceResource(baim)
}

type BillingAccountIamMemberArgs struct {
	// BillingAccountId: string, required
	BillingAccountId terra.StringValue `hcl:"billing_account_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Member: string, required
	Member terra.StringValue `hcl:"member,attr" validate:"required"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// Condition: optional
	Condition *billingaccountiammember.Condition `hcl:"condition,block"`
	// DependsOn contains resources that BillingAccountIamMember depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type billingAccountIamMemberAttributes struct {
	ref terra.Reference
}

func (baim billingAccountIamMemberAttributes) BillingAccountId() terra.StringValue {
	return terra.ReferenceString(baim.ref.Append("billing_account_id"))
}

func (baim billingAccountIamMemberAttributes) Etag() terra.StringValue {
	return terra.ReferenceString(baim.ref.Append("etag"))
}

func (baim billingAccountIamMemberAttributes) Id() terra.StringValue {
	return terra.ReferenceString(baim.ref.Append("id"))
}

func (baim billingAccountIamMemberAttributes) Member() terra.StringValue {
	return terra.ReferenceString(baim.ref.Append("member"))
}

func (baim billingAccountIamMemberAttributes) Role() terra.StringValue {
	return terra.ReferenceString(baim.ref.Append("role"))
}

func (baim billingAccountIamMemberAttributes) Condition() terra.ListValue[billingaccountiammember.ConditionAttributes] {
	return terra.ReferenceList[billingaccountiammember.ConditionAttributes](baim.ref.Append("condition"))
}

type billingAccountIamMemberState struct {
	BillingAccountId string                                   `json:"billing_account_id"`
	Etag             string                                   `json:"etag"`
	Id               string                                   `json:"id"`
	Member           string                                   `json:"member"`
	Role             string                                   `json:"role"`
	Condition        []billingaccountiammember.ConditionState `json:"condition"`
}
