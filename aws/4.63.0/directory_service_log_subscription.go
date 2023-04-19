// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDirectoryServiceLogSubscription creates a new instance of [DirectoryServiceLogSubscription].
func NewDirectoryServiceLogSubscription(name string, args DirectoryServiceLogSubscriptionArgs) *DirectoryServiceLogSubscription {
	return &DirectoryServiceLogSubscription{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DirectoryServiceLogSubscription)(nil)

// DirectoryServiceLogSubscription represents the Terraform resource aws_directory_service_log_subscription.
type DirectoryServiceLogSubscription struct {
	Name      string
	Args      DirectoryServiceLogSubscriptionArgs
	state     *directoryServiceLogSubscriptionState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DirectoryServiceLogSubscription].
func (dsls *DirectoryServiceLogSubscription) Type() string {
	return "aws_directory_service_log_subscription"
}

// LocalName returns the local name for [DirectoryServiceLogSubscription].
func (dsls *DirectoryServiceLogSubscription) LocalName() string {
	return dsls.Name
}

// Configuration returns the configuration (args) for [DirectoryServiceLogSubscription].
func (dsls *DirectoryServiceLogSubscription) Configuration() interface{} {
	return dsls.Args
}

// DependOn is used for other resources to depend on [DirectoryServiceLogSubscription].
func (dsls *DirectoryServiceLogSubscription) DependOn() terra.Reference {
	return terra.ReferenceResource(dsls)
}

// Dependencies returns the list of resources [DirectoryServiceLogSubscription] depends_on.
func (dsls *DirectoryServiceLogSubscription) Dependencies() terra.Dependencies {
	return dsls.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DirectoryServiceLogSubscription].
func (dsls *DirectoryServiceLogSubscription) LifecycleManagement() *terra.Lifecycle {
	return dsls.Lifecycle
}

// Attributes returns the attributes for [DirectoryServiceLogSubscription].
func (dsls *DirectoryServiceLogSubscription) Attributes() directoryServiceLogSubscriptionAttributes {
	return directoryServiceLogSubscriptionAttributes{ref: terra.ReferenceResource(dsls)}
}

// ImportState imports the given attribute values into [DirectoryServiceLogSubscription]'s state.
func (dsls *DirectoryServiceLogSubscription) ImportState(av io.Reader) error {
	dsls.state = &directoryServiceLogSubscriptionState{}
	if err := json.NewDecoder(av).Decode(dsls.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dsls.Type(), dsls.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DirectoryServiceLogSubscription] has state.
func (dsls *DirectoryServiceLogSubscription) State() (*directoryServiceLogSubscriptionState, bool) {
	return dsls.state, dsls.state != nil
}

// StateMust returns the state for [DirectoryServiceLogSubscription]. Panics if the state is nil.
func (dsls *DirectoryServiceLogSubscription) StateMust() *directoryServiceLogSubscriptionState {
	if dsls.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dsls.Type(), dsls.LocalName()))
	}
	return dsls.state
}

// DirectoryServiceLogSubscriptionArgs contains the configurations for aws_directory_service_log_subscription.
type DirectoryServiceLogSubscriptionArgs struct {
	// DirectoryId: string, required
	DirectoryId terra.StringValue `hcl:"directory_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// LogGroupName: string, required
	LogGroupName terra.StringValue `hcl:"log_group_name,attr" validate:"required"`
}
type directoryServiceLogSubscriptionAttributes struct {
	ref terra.Reference
}

// DirectoryId returns a reference to field directory_id of aws_directory_service_log_subscription.
func (dsls directoryServiceLogSubscriptionAttributes) DirectoryId() terra.StringValue {
	return terra.ReferenceAsString(dsls.ref.Append("directory_id"))
}

// Id returns a reference to field id of aws_directory_service_log_subscription.
func (dsls directoryServiceLogSubscriptionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dsls.ref.Append("id"))
}

// LogGroupName returns a reference to field log_group_name of aws_directory_service_log_subscription.
func (dsls directoryServiceLogSubscriptionAttributes) LogGroupName() terra.StringValue {
	return terra.ReferenceAsString(dsls.ref.Append("log_group_name"))
}

type directoryServiceLogSubscriptionState struct {
	DirectoryId  string `json:"directory_id"`
	Id           string `json:"id"`
	LogGroupName string `json:"log_group_name"`
}