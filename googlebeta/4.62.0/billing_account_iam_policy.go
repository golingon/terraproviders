// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewBillingAccountIamPolicy creates a new instance of [BillingAccountIamPolicy].
func NewBillingAccountIamPolicy(name string, args BillingAccountIamPolicyArgs) *BillingAccountIamPolicy {
	return &BillingAccountIamPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*BillingAccountIamPolicy)(nil)

// BillingAccountIamPolicy represents the Terraform resource google_billing_account_iam_policy.
type BillingAccountIamPolicy struct {
	Name      string
	Args      BillingAccountIamPolicyArgs
	state     *billingAccountIamPolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [BillingAccountIamPolicy].
func (baip *BillingAccountIamPolicy) Type() string {
	return "google_billing_account_iam_policy"
}

// LocalName returns the local name for [BillingAccountIamPolicy].
func (baip *BillingAccountIamPolicy) LocalName() string {
	return baip.Name
}

// Configuration returns the configuration (args) for [BillingAccountIamPolicy].
func (baip *BillingAccountIamPolicy) Configuration() interface{} {
	return baip.Args
}

// DependOn is used for other resources to depend on [BillingAccountIamPolicy].
func (baip *BillingAccountIamPolicy) DependOn() terra.Reference {
	return terra.ReferenceResource(baip)
}

// Dependencies returns the list of resources [BillingAccountIamPolicy] depends_on.
func (baip *BillingAccountIamPolicy) Dependencies() terra.Dependencies {
	return baip.DependsOn
}

// LifecycleManagement returns the lifecycle block for [BillingAccountIamPolicy].
func (baip *BillingAccountIamPolicy) LifecycleManagement() *terra.Lifecycle {
	return baip.Lifecycle
}

// Attributes returns the attributes for [BillingAccountIamPolicy].
func (baip *BillingAccountIamPolicy) Attributes() billingAccountIamPolicyAttributes {
	return billingAccountIamPolicyAttributes{ref: terra.ReferenceResource(baip)}
}

// ImportState imports the given attribute values into [BillingAccountIamPolicy]'s state.
func (baip *BillingAccountIamPolicy) ImportState(av io.Reader) error {
	baip.state = &billingAccountIamPolicyState{}
	if err := json.NewDecoder(av).Decode(baip.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", baip.Type(), baip.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [BillingAccountIamPolicy] has state.
func (baip *BillingAccountIamPolicy) State() (*billingAccountIamPolicyState, bool) {
	return baip.state, baip.state != nil
}

// StateMust returns the state for [BillingAccountIamPolicy]. Panics if the state is nil.
func (baip *BillingAccountIamPolicy) StateMust() *billingAccountIamPolicyState {
	if baip.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", baip.Type(), baip.LocalName()))
	}
	return baip.state
}

// BillingAccountIamPolicyArgs contains the configurations for google_billing_account_iam_policy.
type BillingAccountIamPolicyArgs struct {
	// BillingAccountId: string, required
	BillingAccountId terra.StringValue `hcl:"billing_account_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// PolicyData: string, required
	PolicyData terra.StringValue `hcl:"policy_data,attr" validate:"required"`
}
type billingAccountIamPolicyAttributes struct {
	ref terra.Reference
}

// BillingAccountId returns a reference to field billing_account_id of google_billing_account_iam_policy.
func (baip billingAccountIamPolicyAttributes) BillingAccountId() terra.StringValue {
	return terra.ReferenceAsString(baip.ref.Append("billing_account_id"))
}

// Etag returns a reference to field etag of google_billing_account_iam_policy.
func (baip billingAccountIamPolicyAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(baip.ref.Append("etag"))
}

// Id returns a reference to field id of google_billing_account_iam_policy.
func (baip billingAccountIamPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(baip.ref.Append("id"))
}

// PolicyData returns a reference to field policy_data of google_billing_account_iam_policy.
func (baip billingAccountIamPolicyAttributes) PolicyData() terra.StringValue {
	return terra.ReferenceAsString(baip.ref.Append("policy_data"))
}

type billingAccountIamPolicyState struct {
	BillingAccountId string `json:"billing_account_id"`
	Etag             string `json:"etag"`
	Id               string `json:"id"`
	PolicyData       string `json:"policy_data"`
}