// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	dmseventsubscription "github.com/golingon/terraproviders/aws/5.6.2/dmseventsubscription"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDmsEventSubscription creates a new instance of [DmsEventSubscription].
func NewDmsEventSubscription(name string, args DmsEventSubscriptionArgs) *DmsEventSubscription {
	return &DmsEventSubscription{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DmsEventSubscription)(nil)

// DmsEventSubscription represents the Terraform resource aws_dms_event_subscription.
type DmsEventSubscription struct {
	Name      string
	Args      DmsEventSubscriptionArgs
	state     *dmsEventSubscriptionState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DmsEventSubscription].
func (des *DmsEventSubscription) Type() string {
	return "aws_dms_event_subscription"
}

// LocalName returns the local name for [DmsEventSubscription].
func (des *DmsEventSubscription) LocalName() string {
	return des.Name
}

// Configuration returns the configuration (args) for [DmsEventSubscription].
func (des *DmsEventSubscription) Configuration() interface{} {
	return des.Args
}

// DependOn is used for other resources to depend on [DmsEventSubscription].
func (des *DmsEventSubscription) DependOn() terra.Reference {
	return terra.ReferenceResource(des)
}

// Dependencies returns the list of resources [DmsEventSubscription] depends_on.
func (des *DmsEventSubscription) Dependencies() terra.Dependencies {
	return des.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DmsEventSubscription].
func (des *DmsEventSubscription) LifecycleManagement() *terra.Lifecycle {
	return des.Lifecycle
}

// Attributes returns the attributes for [DmsEventSubscription].
func (des *DmsEventSubscription) Attributes() dmsEventSubscriptionAttributes {
	return dmsEventSubscriptionAttributes{ref: terra.ReferenceResource(des)}
}

// ImportState imports the given attribute values into [DmsEventSubscription]'s state.
func (des *DmsEventSubscription) ImportState(av io.Reader) error {
	des.state = &dmsEventSubscriptionState{}
	if err := json.NewDecoder(av).Decode(des.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", des.Type(), des.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DmsEventSubscription] has state.
func (des *DmsEventSubscription) State() (*dmsEventSubscriptionState, bool) {
	return des.state, des.state != nil
}

// StateMust returns the state for [DmsEventSubscription]. Panics if the state is nil.
func (des *DmsEventSubscription) StateMust() *dmsEventSubscriptionState {
	if des.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", des.Type(), des.LocalName()))
	}
	return des.state
}

// DmsEventSubscriptionArgs contains the configurations for aws_dms_event_subscription.
type DmsEventSubscriptionArgs struct {
	// Enabled: bool, optional
	Enabled terra.BoolValue `hcl:"enabled,attr"`
	// EventCategories: set of string, required
	EventCategories terra.SetValue[terra.StringValue] `hcl:"event_categories,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// SnsTopicArn: string, required
	SnsTopicArn terra.StringValue `hcl:"sns_topic_arn,attr" validate:"required"`
	// SourceIds: set of string, optional
	SourceIds terra.SetValue[terra.StringValue] `hcl:"source_ids,attr"`
	// SourceType: string, optional
	SourceType terra.StringValue `hcl:"source_type,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// Timeouts: optional
	Timeouts *dmseventsubscription.Timeouts `hcl:"timeouts,block"`
}
type dmsEventSubscriptionAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_dms_event_subscription.
func (des dmsEventSubscriptionAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(des.ref.Append("arn"))
}

// Enabled returns a reference to field enabled of aws_dms_event_subscription.
func (des dmsEventSubscriptionAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(des.ref.Append("enabled"))
}

// EventCategories returns a reference to field event_categories of aws_dms_event_subscription.
func (des dmsEventSubscriptionAttributes) EventCategories() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](des.ref.Append("event_categories"))
}

// Id returns a reference to field id of aws_dms_event_subscription.
func (des dmsEventSubscriptionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(des.ref.Append("id"))
}

// Name returns a reference to field name of aws_dms_event_subscription.
func (des dmsEventSubscriptionAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(des.ref.Append("name"))
}

// SnsTopicArn returns a reference to field sns_topic_arn of aws_dms_event_subscription.
func (des dmsEventSubscriptionAttributes) SnsTopicArn() terra.StringValue {
	return terra.ReferenceAsString(des.ref.Append("sns_topic_arn"))
}

// SourceIds returns a reference to field source_ids of aws_dms_event_subscription.
func (des dmsEventSubscriptionAttributes) SourceIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](des.ref.Append("source_ids"))
}

// SourceType returns a reference to field source_type of aws_dms_event_subscription.
func (des dmsEventSubscriptionAttributes) SourceType() terra.StringValue {
	return terra.ReferenceAsString(des.ref.Append("source_type"))
}

// Tags returns a reference to field tags of aws_dms_event_subscription.
func (des dmsEventSubscriptionAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](des.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_dms_event_subscription.
func (des dmsEventSubscriptionAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](des.ref.Append("tags_all"))
}

func (des dmsEventSubscriptionAttributes) Timeouts() dmseventsubscription.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dmseventsubscription.TimeoutsAttributes](des.ref.Append("timeouts"))
}

type dmsEventSubscriptionState struct {
	Arn             string                              `json:"arn"`
	Enabled         bool                                `json:"enabled"`
	EventCategories []string                            `json:"event_categories"`
	Id              string                              `json:"id"`
	Name            string                              `json:"name"`
	SnsTopicArn     string                              `json:"sns_topic_arn"`
	SourceIds       []string                            `json:"source_ids"`
	SourceType      string                              `json:"source_type"`
	Tags            map[string]string                   `json:"tags"`
	TagsAll         map[string]string                   `json:"tags_all"`
	Timeouts        *dmseventsubscription.TimeoutsState `json:"timeouts"`
}
