// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	billingaccountiambinding "github.com/golingon/terraproviders/google/4.73.1/billingaccountiambinding"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewBillingAccountIamBinding creates a new instance of [BillingAccountIamBinding].
func NewBillingAccountIamBinding(name string, args BillingAccountIamBindingArgs) *BillingAccountIamBinding {
	return &BillingAccountIamBinding{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*BillingAccountIamBinding)(nil)

// BillingAccountIamBinding represents the Terraform resource google_billing_account_iam_binding.
type BillingAccountIamBinding struct {
	Name      string
	Args      BillingAccountIamBindingArgs
	state     *billingAccountIamBindingState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [BillingAccountIamBinding].
func (baib *BillingAccountIamBinding) Type() string {
	return "google_billing_account_iam_binding"
}

// LocalName returns the local name for [BillingAccountIamBinding].
func (baib *BillingAccountIamBinding) LocalName() string {
	return baib.Name
}

// Configuration returns the configuration (args) for [BillingAccountIamBinding].
func (baib *BillingAccountIamBinding) Configuration() interface{} {
	return baib.Args
}

// DependOn is used for other resources to depend on [BillingAccountIamBinding].
func (baib *BillingAccountIamBinding) DependOn() terra.Reference {
	return terra.ReferenceResource(baib)
}

// Dependencies returns the list of resources [BillingAccountIamBinding] depends_on.
func (baib *BillingAccountIamBinding) Dependencies() terra.Dependencies {
	return baib.DependsOn
}

// LifecycleManagement returns the lifecycle block for [BillingAccountIamBinding].
func (baib *BillingAccountIamBinding) LifecycleManagement() *terra.Lifecycle {
	return baib.Lifecycle
}

// Attributes returns the attributes for [BillingAccountIamBinding].
func (baib *BillingAccountIamBinding) Attributes() billingAccountIamBindingAttributes {
	return billingAccountIamBindingAttributes{ref: terra.ReferenceResource(baib)}
}

// ImportState imports the given attribute values into [BillingAccountIamBinding]'s state.
func (baib *BillingAccountIamBinding) ImportState(av io.Reader) error {
	baib.state = &billingAccountIamBindingState{}
	if err := json.NewDecoder(av).Decode(baib.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", baib.Type(), baib.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [BillingAccountIamBinding] has state.
func (baib *BillingAccountIamBinding) State() (*billingAccountIamBindingState, bool) {
	return baib.state, baib.state != nil
}

// StateMust returns the state for [BillingAccountIamBinding]. Panics if the state is nil.
func (baib *BillingAccountIamBinding) StateMust() *billingAccountIamBindingState {
	if baib.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", baib.Type(), baib.LocalName()))
	}
	return baib.state
}

// BillingAccountIamBindingArgs contains the configurations for google_billing_account_iam_binding.
type BillingAccountIamBindingArgs struct {
	// BillingAccountId: string, required
	BillingAccountId terra.StringValue `hcl:"billing_account_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Members: set of string, required
	Members terra.SetValue[terra.StringValue] `hcl:"members,attr" validate:"required"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// Condition: optional
	Condition *billingaccountiambinding.Condition `hcl:"condition,block"`
}
type billingAccountIamBindingAttributes struct {
	ref terra.Reference
}

// BillingAccountId returns a reference to field billing_account_id of google_billing_account_iam_binding.
func (baib billingAccountIamBindingAttributes) BillingAccountId() terra.StringValue {
	return terra.ReferenceAsString(baib.ref.Append("billing_account_id"))
}

// Etag returns a reference to field etag of google_billing_account_iam_binding.
func (baib billingAccountIamBindingAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(baib.ref.Append("etag"))
}

// Id returns a reference to field id of google_billing_account_iam_binding.
func (baib billingAccountIamBindingAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(baib.ref.Append("id"))
}

// Members returns a reference to field members of google_billing_account_iam_binding.
func (baib billingAccountIamBindingAttributes) Members() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](baib.ref.Append("members"))
}

// Role returns a reference to field role of google_billing_account_iam_binding.
func (baib billingAccountIamBindingAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(baib.ref.Append("role"))
}

func (baib billingAccountIamBindingAttributes) Condition() terra.ListValue[billingaccountiambinding.ConditionAttributes] {
	return terra.ReferenceAsList[billingaccountiambinding.ConditionAttributes](baib.ref.Append("condition"))
}

type billingAccountIamBindingState struct {
	BillingAccountId string                                    `json:"billing_account_id"`
	Etag             string                                    `json:"etag"`
	Id               string                                    `json:"id"`
	Members          []string                                  `json:"members"`
	Role             string                                    `json:"role"`
	Condition        []billingaccountiambinding.ConditionState `json:"condition"`
}
