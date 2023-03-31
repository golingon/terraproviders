// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewBillingAccountIamPolicy(name string, args BillingAccountIamPolicyArgs) *BillingAccountIamPolicy {
	return &BillingAccountIamPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*BillingAccountIamPolicy)(nil)

type BillingAccountIamPolicy struct {
	Name  string
	Args  BillingAccountIamPolicyArgs
	state *billingAccountIamPolicyState
}

func (baip *BillingAccountIamPolicy) Type() string {
	return "google_billing_account_iam_policy"
}

func (baip *BillingAccountIamPolicy) LocalName() string {
	return baip.Name
}

func (baip *BillingAccountIamPolicy) Configuration() interface{} {
	return baip.Args
}

func (baip *BillingAccountIamPolicy) Attributes() billingAccountIamPolicyAttributes {
	return billingAccountIamPolicyAttributes{ref: terra.ReferenceResource(baip)}
}

func (baip *BillingAccountIamPolicy) ImportState(av io.Reader) error {
	baip.state = &billingAccountIamPolicyState{}
	if err := json.NewDecoder(av).Decode(baip.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", baip.Type(), baip.LocalName(), err)
	}
	return nil
}

func (baip *BillingAccountIamPolicy) State() (*billingAccountIamPolicyState, bool) {
	return baip.state, baip.state != nil
}

func (baip *BillingAccountIamPolicy) StateMust() *billingAccountIamPolicyState {
	if baip.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", baip.Type(), baip.LocalName()))
	}
	return baip.state
}

func (baip *BillingAccountIamPolicy) DependOn() terra.Reference {
	return terra.ReferenceResource(baip)
}

type BillingAccountIamPolicyArgs struct {
	// BillingAccountId: string, required
	BillingAccountId terra.StringValue `hcl:"billing_account_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// PolicyData: string, required
	PolicyData terra.StringValue `hcl:"policy_data,attr" validate:"required"`
	// DependsOn contains resources that BillingAccountIamPolicy depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type billingAccountIamPolicyAttributes struct {
	ref terra.Reference
}

func (baip billingAccountIamPolicyAttributes) BillingAccountId() terra.StringValue {
	return terra.ReferenceString(baip.ref.Append("billing_account_id"))
}

func (baip billingAccountIamPolicyAttributes) Etag() terra.StringValue {
	return terra.ReferenceString(baip.ref.Append("etag"))
}

func (baip billingAccountIamPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceString(baip.ref.Append("id"))
}

func (baip billingAccountIamPolicyAttributes) PolicyData() terra.StringValue {
	return terra.ReferenceString(baip.ref.Append("policy_data"))
}

type billingAccountIamPolicyState struct {
	BillingAccountId string `json:"billing_account_id"`
	Etag             string `json:"etag"`
	Id               string `json:"id"`
	PolicyData       string `json:"policy_data"`
}
