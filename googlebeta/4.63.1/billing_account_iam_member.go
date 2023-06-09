// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	billingaccountiammember "github.com/golingon/terraproviders/googlebeta/4.63.1/billingaccountiammember"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewBillingAccountIamMember creates a new instance of [BillingAccountIamMember].
func NewBillingAccountIamMember(name string, args BillingAccountIamMemberArgs) *BillingAccountIamMember {
	return &BillingAccountIamMember{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*BillingAccountIamMember)(nil)

// BillingAccountIamMember represents the Terraform resource google_billing_account_iam_member.
type BillingAccountIamMember struct {
	Name      string
	Args      BillingAccountIamMemberArgs
	state     *billingAccountIamMemberState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [BillingAccountIamMember].
func (baim *BillingAccountIamMember) Type() string {
	return "google_billing_account_iam_member"
}

// LocalName returns the local name for [BillingAccountIamMember].
func (baim *BillingAccountIamMember) LocalName() string {
	return baim.Name
}

// Configuration returns the configuration (args) for [BillingAccountIamMember].
func (baim *BillingAccountIamMember) Configuration() interface{} {
	return baim.Args
}

// DependOn is used for other resources to depend on [BillingAccountIamMember].
func (baim *BillingAccountIamMember) DependOn() terra.Reference {
	return terra.ReferenceResource(baim)
}

// Dependencies returns the list of resources [BillingAccountIamMember] depends_on.
func (baim *BillingAccountIamMember) Dependencies() terra.Dependencies {
	return baim.DependsOn
}

// LifecycleManagement returns the lifecycle block for [BillingAccountIamMember].
func (baim *BillingAccountIamMember) LifecycleManagement() *terra.Lifecycle {
	return baim.Lifecycle
}

// Attributes returns the attributes for [BillingAccountIamMember].
func (baim *BillingAccountIamMember) Attributes() billingAccountIamMemberAttributes {
	return billingAccountIamMemberAttributes{ref: terra.ReferenceResource(baim)}
}

// ImportState imports the given attribute values into [BillingAccountIamMember]'s state.
func (baim *BillingAccountIamMember) ImportState(av io.Reader) error {
	baim.state = &billingAccountIamMemberState{}
	if err := json.NewDecoder(av).Decode(baim.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", baim.Type(), baim.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [BillingAccountIamMember] has state.
func (baim *BillingAccountIamMember) State() (*billingAccountIamMemberState, bool) {
	return baim.state, baim.state != nil
}

// StateMust returns the state for [BillingAccountIamMember]. Panics if the state is nil.
func (baim *BillingAccountIamMember) StateMust() *billingAccountIamMemberState {
	if baim.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", baim.Type(), baim.LocalName()))
	}
	return baim.state
}

// BillingAccountIamMemberArgs contains the configurations for google_billing_account_iam_member.
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
}
type billingAccountIamMemberAttributes struct {
	ref terra.Reference
}

// BillingAccountId returns a reference to field billing_account_id of google_billing_account_iam_member.
func (baim billingAccountIamMemberAttributes) BillingAccountId() terra.StringValue {
	return terra.ReferenceAsString(baim.ref.Append("billing_account_id"))
}

// Etag returns a reference to field etag of google_billing_account_iam_member.
func (baim billingAccountIamMemberAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(baim.ref.Append("etag"))
}

// Id returns a reference to field id of google_billing_account_iam_member.
func (baim billingAccountIamMemberAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(baim.ref.Append("id"))
}

// Member returns a reference to field member of google_billing_account_iam_member.
func (baim billingAccountIamMemberAttributes) Member() terra.StringValue {
	return terra.ReferenceAsString(baim.ref.Append("member"))
}

// Role returns a reference to field role of google_billing_account_iam_member.
func (baim billingAccountIamMemberAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(baim.ref.Append("role"))
}

func (baim billingAccountIamMemberAttributes) Condition() terra.ListValue[billingaccountiammember.ConditionAttributes] {
	return terra.ReferenceAsList[billingaccountiammember.ConditionAttributes](baim.ref.Append("condition"))
}

type billingAccountIamMemberState struct {
	BillingAccountId string                                   `json:"billing_account_id"`
	Etag             string                                   `json:"etag"`
	Id               string                                   `json:"id"`
	Member           string                                   `json:"member"`
	Role             string                                   `json:"role"`
	Condition        []billingaccountiammember.ConditionState `json:"condition"`
}
