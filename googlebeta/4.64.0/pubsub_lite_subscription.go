// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	pubsublitesubscription "github.com/golingon/terraproviders/googlebeta/4.64.0/pubsublitesubscription"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewPubsubLiteSubscription creates a new instance of [PubsubLiteSubscription].
func NewPubsubLiteSubscription(name string, args PubsubLiteSubscriptionArgs) *PubsubLiteSubscription {
	return &PubsubLiteSubscription{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*PubsubLiteSubscription)(nil)

// PubsubLiteSubscription represents the Terraform resource google_pubsub_lite_subscription.
type PubsubLiteSubscription struct {
	Name      string
	Args      PubsubLiteSubscriptionArgs
	state     *pubsubLiteSubscriptionState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [PubsubLiteSubscription].
func (pls *PubsubLiteSubscription) Type() string {
	return "google_pubsub_lite_subscription"
}

// LocalName returns the local name for [PubsubLiteSubscription].
func (pls *PubsubLiteSubscription) LocalName() string {
	return pls.Name
}

// Configuration returns the configuration (args) for [PubsubLiteSubscription].
func (pls *PubsubLiteSubscription) Configuration() interface{} {
	return pls.Args
}

// DependOn is used for other resources to depend on [PubsubLiteSubscription].
func (pls *PubsubLiteSubscription) DependOn() terra.Reference {
	return terra.ReferenceResource(pls)
}

// Dependencies returns the list of resources [PubsubLiteSubscription] depends_on.
func (pls *PubsubLiteSubscription) Dependencies() terra.Dependencies {
	return pls.DependsOn
}

// LifecycleManagement returns the lifecycle block for [PubsubLiteSubscription].
func (pls *PubsubLiteSubscription) LifecycleManagement() *terra.Lifecycle {
	return pls.Lifecycle
}

// Attributes returns the attributes for [PubsubLiteSubscription].
func (pls *PubsubLiteSubscription) Attributes() pubsubLiteSubscriptionAttributes {
	return pubsubLiteSubscriptionAttributes{ref: terra.ReferenceResource(pls)}
}

// ImportState imports the given attribute values into [PubsubLiteSubscription]'s state.
func (pls *PubsubLiteSubscription) ImportState(av io.Reader) error {
	pls.state = &pubsubLiteSubscriptionState{}
	if err := json.NewDecoder(av).Decode(pls.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", pls.Type(), pls.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [PubsubLiteSubscription] has state.
func (pls *PubsubLiteSubscription) State() (*pubsubLiteSubscriptionState, bool) {
	return pls.state, pls.state != nil
}

// StateMust returns the state for [PubsubLiteSubscription]. Panics if the state is nil.
func (pls *PubsubLiteSubscription) StateMust() *pubsubLiteSubscriptionState {
	if pls.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", pls.Type(), pls.LocalName()))
	}
	return pls.state
}

// PubsubLiteSubscriptionArgs contains the configurations for google_pubsub_lite_subscription.
type PubsubLiteSubscriptionArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Region: string, optional
	Region terra.StringValue `hcl:"region,attr"`
	// Topic: string, required
	Topic terra.StringValue `hcl:"topic,attr" validate:"required"`
	// Zone: string, optional
	Zone terra.StringValue `hcl:"zone,attr"`
	// DeliveryConfig: optional
	DeliveryConfig *pubsublitesubscription.DeliveryConfig `hcl:"delivery_config,block"`
	// Timeouts: optional
	Timeouts *pubsublitesubscription.Timeouts `hcl:"timeouts,block"`
}
type pubsubLiteSubscriptionAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of google_pubsub_lite_subscription.
func (pls pubsubLiteSubscriptionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(pls.ref.Append("id"))
}

// Name returns a reference to field name of google_pubsub_lite_subscription.
func (pls pubsubLiteSubscriptionAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(pls.ref.Append("name"))
}

// Project returns a reference to field project of google_pubsub_lite_subscription.
func (pls pubsubLiteSubscriptionAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(pls.ref.Append("project"))
}

// Region returns a reference to field region of google_pubsub_lite_subscription.
func (pls pubsubLiteSubscriptionAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(pls.ref.Append("region"))
}

// Topic returns a reference to field topic of google_pubsub_lite_subscription.
func (pls pubsubLiteSubscriptionAttributes) Topic() terra.StringValue {
	return terra.ReferenceAsString(pls.ref.Append("topic"))
}

// Zone returns a reference to field zone of google_pubsub_lite_subscription.
func (pls pubsubLiteSubscriptionAttributes) Zone() terra.StringValue {
	return terra.ReferenceAsString(pls.ref.Append("zone"))
}

func (pls pubsubLiteSubscriptionAttributes) DeliveryConfig() terra.ListValue[pubsublitesubscription.DeliveryConfigAttributes] {
	return terra.ReferenceAsList[pubsublitesubscription.DeliveryConfigAttributes](pls.ref.Append("delivery_config"))
}

func (pls pubsubLiteSubscriptionAttributes) Timeouts() pubsublitesubscription.TimeoutsAttributes {
	return terra.ReferenceAsSingle[pubsublitesubscription.TimeoutsAttributes](pls.ref.Append("timeouts"))
}

type pubsubLiteSubscriptionState struct {
	Id             string                                       `json:"id"`
	Name           string                                       `json:"name"`
	Project        string                                       `json:"project"`
	Region         string                                       `json:"region"`
	Topic          string                                       `json:"topic"`
	Zone           string                                       `json:"zone"`
	DeliveryConfig []pubsublitesubscription.DeliveryConfigState `json:"delivery_config"`
	Timeouts       *pubsublitesubscription.TimeoutsState        `json:"timeouts"`
}
