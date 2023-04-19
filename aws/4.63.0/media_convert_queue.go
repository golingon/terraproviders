// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	mediaconvertqueue "github.com/golingon/terraproviders/aws/4.63.0/mediaconvertqueue"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewMediaConvertQueue creates a new instance of [MediaConvertQueue].
func NewMediaConvertQueue(name string, args MediaConvertQueueArgs) *MediaConvertQueue {
	return &MediaConvertQueue{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*MediaConvertQueue)(nil)

// MediaConvertQueue represents the Terraform resource aws_media_convert_queue.
type MediaConvertQueue struct {
	Name      string
	Args      MediaConvertQueueArgs
	state     *mediaConvertQueueState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [MediaConvertQueue].
func (mcq *MediaConvertQueue) Type() string {
	return "aws_media_convert_queue"
}

// LocalName returns the local name for [MediaConvertQueue].
func (mcq *MediaConvertQueue) LocalName() string {
	return mcq.Name
}

// Configuration returns the configuration (args) for [MediaConvertQueue].
func (mcq *MediaConvertQueue) Configuration() interface{} {
	return mcq.Args
}

// DependOn is used for other resources to depend on [MediaConvertQueue].
func (mcq *MediaConvertQueue) DependOn() terra.Reference {
	return terra.ReferenceResource(mcq)
}

// Dependencies returns the list of resources [MediaConvertQueue] depends_on.
func (mcq *MediaConvertQueue) Dependencies() terra.Dependencies {
	return mcq.DependsOn
}

// LifecycleManagement returns the lifecycle block for [MediaConvertQueue].
func (mcq *MediaConvertQueue) LifecycleManagement() *terra.Lifecycle {
	return mcq.Lifecycle
}

// Attributes returns the attributes for [MediaConvertQueue].
func (mcq *MediaConvertQueue) Attributes() mediaConvertQueueAttributes {
	return mediaConvertQueueAttributes{ref: terra.ReferenceResource(mcq)}
}

// ImportState imports the given attribute values into [MediaConvertQueue]'s state.
func (mcq *MediaConvertQueue) ImportState(av io.Reader) error {
	mcq.state = &mediaConvertQueueState{}
	if err := json.NewDecoder(av).Decode(mcq.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", mcq.Type(), mcq.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [MediaConvertQueue] has state.
func (mcq *MediaConvertQueue) State() (*mediaConvertQueueState, bool) {
	return mcq.state, mcq.state != nil
}

// StateMust returns the state for [MediaConvertQueue]. Panics if the state is nil.
func (mcq *MediaConvertQueue) StateMust() *mediaConvertQueueState {
	if mcq.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", mcq.Type(), mcq.LocalName()))
	}
	return mcq.state
}

// MediaConvertQueueArgs contains the configurations for aws_media_convert_queue.
type MediaConvertQueueArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// PricingPlan: string, optional
	PricingPlan terra.StringValue `hcl:"pricing_plan,attr"`
	// Status: string, optional
	Status terra.StringValue `hcl:"status,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// ReservationPlanSettings: optional
	ReservationPlanSettings *mediaconvertqueue.ReservationPlanSettings `hcl:"reservation_plan_settings,block"`
}
type mediaConvertQueueAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_media_convert_queue.
func (mcq mediaConvertQueueAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(mcq.ref.Append("arn"))
}

// Description returns a reference to field description of aws_media_convert_queue.
func (mcq mediaConvertQueueAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(mcq.ref.Append("description"))
}

// Id returns a reference to field id of aws_media_convert_queue.
func (mcq mediaConvertQueueAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(mcq.ref.Append("id"))
}

// Name returns a reference to field name of aws_media_convert_queue.
func (mcq mediaConvertQueueAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(mcq.ref.Append("name"))
}

// PricingPlan returns a reference to field pricing_plan of aws_media_convert_queue.
func (mcq mediaConvertQueueAttributes) PricingPlan() terra.StringValue {
	return terra.ReferenceAsString(mcq.ref.Append("pricing_plan"))
}

// Status returns a reference to field status of aws_media_convert_queue.
func (mcq mediaConvertQueueAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(mcq.ref.Append("status"))
}

// Tags returns a reference to field tags of aws_media_convert_queue.
func (mcq mediaConvertQueueAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](mcq.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_media_convert_queue.
func (mcq mediaConvertQueueAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](mcq.ref.Append("tags_all"))
}

func (mcq mediaConvertQueueAttributes) ReservationPlanSettings() terra.ListValue[mediaconvertqueue.ReservationPlanSettingsAttributes] {
	return terra.ReferenceAsList[mediaconvertqueue.ReservationPlanSettingsAttributes](mcq.ref.Append("reservation_plan_settings"))
}

type mediaConvertQueueState struct {
	Arn                     string                                           `json:"arn"`
	Description             string                                           `json:"description"`
	Id                      string                                           `json:"id"`
	Name                    string                                           `json:"name"`
	PricingPlan             string                                           `json:"pricing_plan"`
	Status                  string                                           `json:"status"`
	Tags                    map[string]string                                `json:"tags"`
	TagsAll                 map[string]string                                `json:"tags_all"`
	ReservationPlanSettings []mediaconvertqueue.ReservationPlanSettingsState `json:"reservation_plan_settings"`
}
