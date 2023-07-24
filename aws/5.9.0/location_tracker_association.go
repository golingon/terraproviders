// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	locationtrackerassociation "github.com/golingon/terraproviders/aws/5.9.0/locationtrackerassociation"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewLocationTrackerAssociation creates a new instance of [LocationTrackerAssociation].
func NewLocationTrackerAssociation(name string, args LocationTrackerAssociationArgs) *LocationTrackerAssociation {
	return &LocationTrackerAssociation{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*LocationTrackerAssociation)(nil)

// LocationTrackerAssociation represents the Terraform resource aws_location_tracker_association.
type LocationTrackerAssociation struct {
	Name      string
	Args      LocationTrackerAssociationArgs
	state     *locationTrackerAssociationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [LocationTrackerAssociation].
func (lta *LocationTrackerAssociation) Type() string {
	return "aws_location_tracker_association"
}

// LocalName returns the local name for [LocationTrackerAssociation].
func (lta *LocationTrackerAssociation) LocalName() string {
	return lta.Name
}

// Configuration returns the configuration (args) for [LocationTrackerAssociation].
func (lta *LocationTrackerAssociation) Configuration() interface{} {
	return lta.Args
}

// DependOn is used for other resources to depend on [LocationTrackerAssociation].
func (lta *LocationTrackerAssociation) DependOn() terra.Reference {
	return terra.ReferenceResource(lta)
}

// Dependencies returns the list of resources [LocationTrackerAssociation] depends_on.
func (lta *LocationTrackerAssociation) Dependencies() terra.Dependencies {
	return lta.DependsOn
}

// LifecycleManagement returns the lifecycle block for [LocationTrackerAssociation].
func (lta *LocationTrackerAssociation) LifecycleManagement() *terra.Lifecycle {
	return lta.Lifecycle
}

// Attributes returns the attributes for [LocationTrackerAssociation].
func (lta *LocationTrackerAssociation) Attributes() locationTrackerAssociationAttributes {
	return locationTrackerAssociationAttributes{ref: terra.ReferenceResource(lta)}
}

// ImportState imports the given attribute values into [LocationTrackerAssociation]'s state.
func (lta *LocationTrackerAssociation) ImportState(av io.Reader) error {
	lta.state = &locationTrackerAssociationState{}
	if err := json.NewDecoder(av).Decode(lta.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", lta.Type(), lta.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [LocationTrackerAssociation] has state.
func (lta *LocationTrackerAssociation) State() (*locationTrackerAssociationState, bool) {
	return lta.state, lta.state != nil
}

// StateMust returns the state for [LocationTrackerAssociation]. Panics if the state is nil.
func (lta *LocationTrackerAssociation) StateMust() *locationTrackerAssociationState {
	if lta.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", lta.Type(), lta.LocalName()))
	}
	return lta.state
}

// LocationTrackerAssociationArgs contains the configurations for aws_location_tracker_association.
type LocationTrackerAssociationArgs struct {
	// ConsumerArn: string, required
	ConsumerArn terra.StringValue `hcl:"consumer_arn,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// TrackerName: string, required
	TrackerName terra.StringValue `hcl:"tracker_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *locationtrackerassociation.Timeouts `hcl:"timeouts,block"`
}
type locationTrackerAssociationAttributes struct {
	ref terra.Reference
}

// ConsumerArn returns a reference to field consumer_arn of aws_location_tracker_association.
func (lta locationTrackerAssociationAttributes) ConsumerArn() terra.StringValue {
	return terra.ReferenceAsString(lta.ref.Append("consumer_arn"))
}

// Id returns a reference to field id of aws_location_tracker_association.
func (lta locationTrackerAssociationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(lta.ref.Append("id"))
}

// TrackerName returns a reference to field tracker_name of aws_location_tracker_association.
func (lta locationTrackerAssociationAttributes) TrackerName() terra.StringValue {
	return terra.ReferenceAsString(lta.ref.Append("tracker_name"))
}

func (lta locationTrackerAssociationAttributes) Timeouts() locationtrackerassociation.TimeoutsAttributes {
	return terra.ReferenceAsSingle[locationtrackerassociation.TimeoutsAttributes](lta.ref.Append("timeouts"))
}

type locationTrackerAssociationState struct {
	ConsumerArn string                                    `json:"consumer_arn"`
	Id          string                                    `json:"id"`
	TrackerName string                                    `json:"tracker_name"`
	Timeouts    *locationtrackerassociation.TimeoutsState `json:"timeouts"`
}
