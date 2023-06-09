// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	cloudtasksqueueiammember "github.com/golingon/terraproviders/google/4.71.0/cloudtasksqueueiammember"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewCloudTasksQueueIamMember creates a new instance of [CloudTasksQueueIamMember].
func NewCloudTasksQueueIamMember(name string, args CloudTasksQueueIamMemberArgs) *CloudTasksQueueIamMember {
	return &CloudTasksQueueIamMember{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*CloudTasksQueueIamMember)(nil)

// CloudTasksQueueIamMember represents the Terraform resource google_cloud_tasks_queue_iam_member.
type CloudTasksQueueIamMember struct {
	Name      string
	Args      CloudTasksQueueIamMemberArgs
	state     *cloudTasksQueueIamMemberState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [CloudTasksQueueIamMember].
func (ctqim *CloudTasksQueueIamMember) Type() string {
	return "google_cloud_tasks_queue_iam_member"
}

// LocalName returns the local name for [CloudTasksQueueIamMember].
func (ctqim *CloudTasksQueueIamMember) LocalName() string {
	return ctqim.Name
}

// Configuration returns the configuration (args) for [CloudTasksQueueIamMember].
func (ctqim *CloudTasksQueueIamMember) Configuration() interface{} {
	return ctqim.Args
}

// DependOn is used for other resources to depend on [CloudTasksQueueIamMember].
func (ctqim *CloudTasksQueueIamMember) DependOn() terra.Reference {
	return terra.ReferenceResource(ctqim)
}

// Dependencies returns the list of resources [CloudTasksQueueIamMember] depends_on.
func (ctqim *CloudTasksQueueIamMember) Dependencies() terra.Dependencies {
	return ctqim.DependsOn
}

// LifecycleManagement returns the lifecycle block for [CloudTasksQueueIamMember].
func (ctqim *CloudTasksQueueIamMember) LifecycleManagement() *terra.Lifecycle {
	return ctqim.Lifecycle
}

// Attributes returns the attributes for [CloudTasksQueueIamMember].
func (ctqim *CloudTasksQueueIamMember) Attributes() cloudTasksQueueIamMemberAttributes {
	return cloudTasksQueueIamMemberAttributes{ref: terra.ReferenceResource(ctqim)}
}

// ImportState imports the given attribute values into [CloudTasksQueueIamMember]'s state.
func (ctqim *CloudTasksQueueIamMember) ImportState(av io.Reader) error {
	ctqim.state = &cloudTasksQueueIamMemberState{}
	if err := json.NewDecoder(av).Decode(ctqim.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ctqim.Type(), ctqim.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [CloudTasksQueueIamMember] has state.
func (ctqim *CloudTasksQueueIamMember) State() (*cloudTasksQueueIamMemberState, bool) {
	return ctqim.state, ctqim.state != nil
}

// StateMust returns the state for [CloudTasksQueueIamMember]. Panics if the state is nil.
func (ctqim *CloudTasksQueueIamMember) StateMust() *cloudTasksQueueIamMemberState {
	if ctqim.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ctqim.Type(), ctqim.LocalName()))
	}
	return ctqim.state
}

// CloudTasksQueueIamMemberArgs contains the configurations for google_cloud_tasks_queue_iam_member.
type CloudTasksQueueIamMemberArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, optional
	Location terra.StringValue `hcl:"location,attr"`
	// Member: string, required
	Member terra.StringValue `hcl:"member,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// Condition: optional
	Condition *cloudtasksqueueiammember.Condition `hcl:"condition,block"`
}
type cloudTasksQueueIamMemberAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_cloud_tasks_queue_iam_member.
func (ctqim cloudTasksQueueIamMemberAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(ctqim.ref.Append("etag"))
}

// Id returns a reference to field id of google_cloud_tasks_queue_iam_member.
func (ctqim cloudTasksQueueIamMemberAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ctqim.ref.Append("id"))
}

// Location returns a reference to field location of google_cloud_tasks_queue_iam_member.
func (ctqim cloudTasksQueueIamMemberAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(ctqim.ref.Append("location"))
}

// Member returns a reference to field member of google_cloud_tasks_queue_iam_member.
func (ctqim cloudTasksQueueIamMemberAttributes) Member() terra.StringValue {
	return terra.ReferenceAsString(ctqim.ref.Append("member"))
}

// Name returns a reference to field name of google_cloud_tasks_queue_iam_member.
func (ctqim cloudTasksQueueIamMemberAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ctqim.ref.Append("name"))
}

// Project returns a reference to field project of google_cloud_tasks_queue_iam_member.
func (ctqim cloudTasksQueueIamMemberAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(ctqim.ref.Append("project"))
}

// Role returns a reference to field role of google_cloud_tasks_queue_iam_member.
func (ctqim cloudTasksQueueIamMemberAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(ctqim.ref.Append("role"))
}

func (ctqim cloudTasksQueueIamMemberAttributes) Condition() terra.ListValue[cloudtasksqueueiammember.ConditionAttributes] {
	return terra.ReferenceAsList[cloudtasksqueueiammember.ConditionAttributes](ctqim.ref.Append("condition"))
}

type cloudTasksQueueIamMemberState struct {
	Etag      string                                    `json:"etag"`
	Id        string                                    `json:"id"`
	Location  string                                    `json:"location"`
	Member    string                                    `json:"member"`
	Name      string                                    `json:"name"`
	Project   string                                    `json:"project"`
	Role      string                                    `json:"role"`
	Condition []cloudtasksqueueiammember.ConditionState `json:"condition"`
}
