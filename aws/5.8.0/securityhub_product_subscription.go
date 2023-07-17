// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSecurityhubProductSubscription creates a new instance of [SecurityhubProductSubscription].
func NewSecurityhubProductSubscription(name string, args SecurityhubProductSubscriptionArgs) *SecurityhubProductSubscription {
	return &SecurityhubProductSubscription{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SecurityhubProductSubscription)(nil)

// SecurityhubProductSubscription represents the Terraform resource aws_securityhub_product_subscription.
type SecurityhubProductSubscription struct {
	Name      string
	Args      SecurityhubProductSubscriptionArgs
	state     *securityhubProductSubscriptionState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SecurityhubProductSubscription].
func (sps *SecurityhubProductSubscription) Type() string {
	return "aws_securityhub_product_subscription"
}

// LocalName returns the local name for [SecurityhubProductSubscription].
func (sps *SecurityhubProductSubscription) LocalName() string {
	return sps.Name
}

// Configuration returns the configuration (args) for [SecurityhubProductSubscription].
func (sps *SecurityhubProductSubscription) Configuration() interface{} {
	return sps.Args
}

// DependOn is used for other resources to depend on [SecurityhubProductSubscription].
func (sps *SecurityhubProductSubscription) DependOn() terra.Reference {
	return terra.ReferenceResource(sps)
}

// Dependencies returns the list of resources [SecurityhubProductSubscription] depends_on.
func (sps *SecurityhubProductSubscription) Dependencies() terra.Dependencies {
	return sps.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SecurityhubProductSubscription].
func (sps *SecurityhubProductSubscription) LifecycleManagement() *terra.Lifecycle {
	return sps.Lifecycle
}

// Attributes returns the attributes for [SecurityhubProductSubscription].
func (sps *SecurityhubProductSubscription) Attributes() securityhubProductSubscriptionAttributes {
	return securityhubProductSubscriptionAttributes{ref: terra.ReferenceResource(sps)}
}

// ImportState imports the given attribute values into [SecurityhubProductSubscription]'s state.
func (sps *SecurityhubProductSubscription) ImportState(av io.Reader) error {
	sps.state = &securityhubProductSubscriptionState{}
	if err := json.NewDecoder(av).Decode(sps.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sps.Type(), sps.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SecurityhubProductSubscription] has state.
func (sps *SecurityhubProductSubscription) State() (*securityhubProductSubscriptionState, bool) {
	return sps.state, sps.state != nil
}

// StateMust returns the state for [SecurityhubProductSubscription]. Panics if the state is nil.
func (sps *SecurityhubProductSubscription) StateMust() *securityhubProductSubscriptionState {
	if sps.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sps.Type(), sps.LocalName()))
	}
	return sps.state
}

// SecurityhubProductSubscriptionArgs contains the configurations for aws_securityhub_product_subscription.
type SecurityhubProductSubscriptionArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ProductArn: string, required
	ProductArn terra.StringValue `hcl:"product_arn,attr" validate:"required"`
}
type securityhubProductSubscriptionAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_securityhub_product_subscription.
func (sps securityhubProductSubscriptionAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(sps.ref.Append("arn"))
}

// Id returns a reference to field id of aws_securityhub_product_subscription.
func (sps securityhubProductSubscriptionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sps.ref.Append("id"))
}

// ProductArn returns a reference to field product_arn of aws_securityhub_product_subscription.
func (sps securityhubProductSubscriptionAttributes) ProductArn() terra.StringValue {
	return terra.ReferenceAsString(sps.ref.Append("product_arn"))
}

type securityhubProductSubscriptionState struct {
	Arn        string `json:"arn"`
	Id         string `json:"id"`
	ProductArn string `json:"product_arn"`
}