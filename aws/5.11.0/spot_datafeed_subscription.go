// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSpotDatafeedSubscription creates a new instance of [SpotDatafeedSubscription].
func NewSpotDatafeedSubscription(name string, args SpotDatafeedSubscriptionArgs) *SpotDatafeedSubscription {
	return &SpotDatafeedSubscription{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SpotDatafeedSubscription)(nil)

// SpotDatafeedSubscription represents the Terraform resource aws_spot_datafeed_subscription.
type SpotDatafeedSubscription struct {
	Name      string
	Args      SpotDatafeedSubscriptionArgs
	state     *spotDatafeedSubscriptionState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SpotDatafeedSubscription].
func (sds *SpotDatafeedSubscription) Type() string {
	return "aws_spot_datafeed_subscription"
}

// LocalName returns the local name for [SpotDatafeedSubscription].
func (sds *SpotDatafeedSubscription) LocalName() string {
	return sds.Name
}

// Configuration returns the configuration (args) for [SpotDatafeedSubscription].
func (sds *SpotDatafeedSubscription) Configuration() interface{} {
	return sds.Args
}

// DependOn is used for other resources to depend on [SpotDatafeedSubscription].
func (sds *SpotDatafeedSubscription) DependOn() terra.Reference {
	return terra.ReferenceResource(sds)
}

// Dependencies returns the list of resources [SpotDatafeedSubscription] depends_on.
func (sds *SpotDatafeedSubscription) Dependencies() terra.Dependencies {
	return sds.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SpotDatafeedSubscription].
func (sds *SpotDatafeedSubscription) LifecycleManagement() *terra.Lifecycle {
	return sds.Lifecycle
}

// Attributes returns the attributes for [SpotDatafeedSubscription].
func (sds *SpotDatafeedSubscription) Attributes() spotDatafeedSubscriptionAttributes {
	return spotDatafeedSubscriptionAttributes{ref: terra.ReferenceResource(sds)}
}

// ImportState imports the given attribute values into [SpotDatafeedSubscription]'s state.
func (sds *SpotDatafeedSubscription) ImportState(av io.Reader) error {
	sds.state = &spotDatafeedSubscriptionState{}
	if err := json.NewDecoder(av).Decode(sds.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sds.Type(), sds.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SpotDatafeedSubscription] has state.
func (sds *SpotDatafeedSubscription) State() (*spotDatafeedSubscriptionState, bool) {
	return sds.state, sds.state != nil
}

// StateMust returns the state for [SpotDatafeedSubscription]. Panics if the state is nil.
func (sds *SpotDatafeedSubscription) StateMust() *spotDatafeedSubscriptionState {
	if sds.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sds.Type(), sds.LocalName()))
	}
	return sds.state
}

// SpotDatafeedSubscriptionArgs contains the configurations for aws_spot_datafeed_subscription.
type SpotDatafeedSubscriptionArgs struct {
	// Bucket: string, required
	Bucket terra.StringValue `hcl:"bucket,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Prefix: string, optional
	Prefix terra.StringValue `hcl:"prefix,attr"`
}
type spotDatafeedSubscriptionAttributes struct {
	ref terra.Reference
}

// Bucket returns a reference to field bucket of aws_spot_datafeed_subscription.
func (sds spotDatafeedSubscriptionAttributes) Bucket() terra.StringValue {
	return terra.ReferenceAsString(sds.ref.Append("bucket"))
}

// Id returns a reference to field id of aws_spot_datafeed_subscription.
func (sds spotDatafeedSubscriptionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sds.ref.Append("id"))
}

// Prefix returns a reference to field prefix of aws_spot_datafeed_subscription.
func (sds spotDatafeedSubscriptionAttributes) Prefix() terra.StringValue {
	return terra.ReferenceAsString(sds.ref.Append("prefix"))
}

type spotDatafeedSubscriptionState struct {
	Bucket string `json:"bucket"`
	Id     string `json:"id"`
	Prefix string `json:"prefix"`
}
