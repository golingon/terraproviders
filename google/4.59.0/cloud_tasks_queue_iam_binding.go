// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	cloudtasksqueueiambinding "github.com/golingon/terraproviders/google/4.59.0/cloudtasksqueueiambinding"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewCloudTasksQueueIamBinding creates a new instance of [CloudTasksQueueIamBinding].
func NewCloudTasksQueueIamBinding(name string, args CloudTasksQueueIamBindingArgs) *CloudTasksQueueIamBinding {
	return &CloudTasksQueueIamBinding{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*CloudTasksQueueIamBinding)(nil)

// CloudTasksQueueIamBinding represents the Terraform resource google_cloud_tasks_queue_iam_binding.
type CloudTasksQueueIamBinding struct {
	Name      string
	Args      CloudTasksQueueIamBindingArgs
	state     *cloudTasksQueueIamBindingState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [CloudTasksQueueIamBinding].
func (ctqib *CloudTasksQueueIamBinding) Type() string {
	return "google_cloud_tasks_queue_iam_binding"
}

// LocalName returns the local name for [CloudTasksQueueIamBinding].
func (ctqib *CloudTasksQueueIamBinding) LocalName() string {
	return ctqib.Name
}

// Configuration returns the configuration (args) for [CloudTasksQueueIamBinding].
func (ctqib *CloudTasksQueueIamBinding) Configuration() interface{} {
	return ctqib.Args
}

// DependOn is used for other resources to depend on [CloudTasksQueueIamBinding].
func (ctqib *CloudTasksQueueIamBinding) DependOn() terra.Reference {
	return terra.ReferenceResource(ctqib)
}

// Dependencies returns the list of resources [CloudTasksQueueIamBinding] depends_on.
func (ctqib *CloudTasksQueueIamBinding) Dependencies() terra.Dependencies {
	return ctqib.DependsOn
}

// LifecycleManagement returns the lifecycle block for [CloudTasksQueueIamBinding].
func (ctqib *CloudTasksQueueIamBinding) LifecycleManagement() *terra.Lifecycle {
	return ctqib.Lifecycle
}

// Attributes returns the attributes for [CloudTasksQueueIamBinding].
func (ctqib *CloudTasksQueueIamBinding) Attributes() cloudTasksQueueIamBindingAttributes {
	return cloudTasksQueueIamBindingAttributes{ref: terra.ReferenceResource(ctqib)}
}

// ImportState imports the given attribute values into [CloudTasksQueueIamBinding]'s state.
func (ctqib *CloudTasksQueueIamBinding) ImportState(av io.Reader) error {
	ctqib.state = &cloudTasksQueueIamBindingState{}
	if err := json.NewDecoder(av).Decode(ctqib.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ctqib.Type(), ctqib.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [CloudTasksQueueIamBinding] has state.
func (ctqib *CloudTasksQueueIamBinding) State() (*cloudTasksQueueIamBindingState, bool) {
	return ctqib.state, ctqib.state != nil
}

// StateMust returns the state for [CloudTasksQueueIamBinding]. Panics if the state is nil.
func (ctqib *CloudTasksQueueIamBinding) StateMust() *cloudTasksQueueIamBindingState {
	if ctqib.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ctqib.Type(), ctqib.LocalName()))
	}
	return ctqib.state
}

// CloudTasksQueueIamBindingArgs contains the configurations for google_cloud_tasks_queue_iam_binding.
type CloudTasksQueueIamBindingArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, optional
	Location terra.StringValue `hcl:"location,attr"`
	// Members: set of string, required
	Members terra.SetValue[terra.StringValue] `hcl:"members,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// Condition: optional
	Condition *cloudtasksqueueiambinding.Condition `hcl:"condition,block"`
}
type cloudTasksQueueIamBindingAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_cloud_tasks_queue_iam_binding.
func (ctqib cloudTasksQueueIamBindingAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(ctqib.ref.Append("etag"))
}

// Id returns a reference to field id of google_cloud_tasks_queue_iam_binding.
func (ctqib cloudTasksQueueIamBindingAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ctqib.ref.Append("id"))
}

// Location returns a reference to field location of google_cloud_tasks_queue_iam_binding.
func (ctqib cloudTasksQueueIamBindingAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(ctqib.ref.Append("location"))
}

// Members returns a reference to field members of google_cloud_tasks_queue_iam_binding.
func (ctqib cloudTasksQueueIamBindingAttributes) Members() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](ctqib.ref.Append("members"))
}

// Name returns a reference to field name of google_cloud_tasks_queue_iam_binding.
func (ctqib cloudTasksQueueIamBindingAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ctqib.ref.Append("name"))
}

// Project returns a reference to field project of google_cloud_tasks_queue_iam_binding.
func (ctqib cloudTasksQueueIamBindingAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(ctqib.ref.Append("project"))
}

// Role returns a reference to field role of google_cloud_tasks_queue_iam_binding.
func (ctqib cloudTasksQueueIamBindingAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(ctqib.ref.Append("role"))
}

func (ctqib cloudTasksQueueIamBindingAttributes) Condition() terra.ListValue[cloudtasksqueueiambinding.ConditionAttributes] {
	return terra.ReferenceAsList[cloudtasksqueueiambinding.ConditionAttributes](ctqib.ref.Append("condition"))
}

type cloudTasksQueueIamBindingState struct {
	Etag      string                                     `json:"etag"`
	Id        string                                     `json:"id"`
	Location  string                                     `json:"location"`
	Members   []string                                   `json:"members"`
	Name      string                                     `json:"name"`
	Project   string                                     `json:"project"`
	Role      string                                     `json:"role"`
	Condition []cloudtasksqueueiambinding.ConditionState `json:"condition"`
}
