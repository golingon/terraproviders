// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	cloudrunv2jobiammember "github.com/golingon/terraproviders/google/4.72.1/cloudrunv2jobiammember"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewCloudRunV2JobIamMember creates a new instance of [CloudRunV2JobIamMember].
func NewCloudRunV2JobIamMember(name string, args CloudRunV2JobIamMemberArgs) *CloudRunV2JobIamMember {
	return &CloudRunV2JobIamMember{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*CloudRunV2JobIamMember)(nil)

// CloudRunV2JobIamMember represents the Terraform resource google_cloud_run_v2_job_iam_member.
type CloudRunV2JobIamMember struct {
	Name      string
	Args      CloudRunV2JobIamMemberArgs
	state     *cloudRunV2JobIamMemberState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [CloudRunV2JobIamMember].
func (crvjim *CloudRunV2JobIamMember) Type() string {
	return "google_cloud_run_v2_job_iam_member"
}

// LocalName returns the local name for [CloudRunV2JobIamMember].
func (crvjim *CloudRunV2JobIamMember) LocalName() string {
	return crvjim.Name
}

// Configuration returns the configuration (args) for [CloudRunV2JobIamMember].
func (crvjim *CloudRunV2JobIamMember) Configuration() interface{} {
	return crvjim.Args
}

// DependOn is used for other resources to depend on [CloudRunV2JobIamMember].
func (crvjim *CloudRunV2JobIamMember) DependOn() terra.Reference {
	return terra.ReferenceResource(crvjim)
}

// Dependencies returns the list of resources [CloudRunV2JobIamMember] depends_on.
func (crvjim *CloudRunV2JobIamMember) Dependencies() terra.Dependencies {
	return crvjim.DependsOn
}

// LifecycleManagement returns the lifecycle block for [CloudRunV2JobIamMember].
func (crvjim *CloudRunV2JobIamMember) LifecycleManagement() *terra.Lifecycle {
	return crvjim.Lifecycle
}

// Attributes returns the attributes for [CloudRunV2JobIamMember].
func (crvjim *CloudRunV2JobIamMember) Attributes() cloudRunV2JobIamMemberAttributes {
	return cloudRunV2JobIamMemberAttributes{ref: terra.ReferenceResource(crvjim)}
}

// ImportState imports the given attribute values into [CloudRunV2JobIamMember]'s state.
func (crvjim *CloudRunV2JobIamMember) ImportState(av io.Reader) error {
	crvjim.state = &cloudRunV2JobIamMemberState{}
	if err := json.NewDecoder(av).Decode(crvjim.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", crvjim.Type(), crvjim.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [CloudRunV2JobIamMember] has state.
func (crvjim *CloudRunV2JobIamMember) State() (*cloudRunV2JobIamMemberState, bool) {
	return crvjim.state, crvjim.state != nil
}

// StateMust returns the state for [CloudRunV2JobIamMember]. Panics if the state is nil.
func (crvjim *CloudRunV2JobIamMember) StateMust() *cloudRunV2JobIamMemberState {
	if crvjim.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", crvjim.Type(), crvjim.LocalName()))
	}
	return crvjim.state
}

// CloudRunV2JobIamMemberArgs contains the configurations for google_cloud_run_v2_job_iam_member.
type CloudRunV2JobIamMemberArgs struct {
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
	Condition *cloudrunv2jobiammember.Condition `hcl:"condition,block"`
}
type cloudRunV2JobIamMemberAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_cloud_run_v2_job_iam_member.
func (crvjim cloudRunV2JobIamMemberAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(crvjim.ref.Append("etag"))
}

// Id returns a reference to field id of google_cloud_run_v2_job_iam_member.
func (crvjim cloudRunV2JobIamMemberAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(crvjim.ref.Append("id"))
}

// Location returns a reference to field location of google_cloud_run_v2_job_iam_member.
func (crvjim cloudRunV2JobIamMemberAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(crvjim.ref.Append("location"))
}

// Member returns a reference to field member of google_cloud_run_v2_job_iam_member.
func (crvjim cloudRunV2JobIamMemberAttributes) Member() terra.StringValue {
	return terra.ReferenceAsString(crvjim.ref.Append("member"))
}

// Name returns a reference to field name of google_cloud_run_v2_job_iam_member.
func (crvjim cloudRunV2JobIamMemberAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(crvjim.ref.Append("name"))
}

// Project returns a reference to field project of google_cloud_run_v2_job_iam_member.
func (crvjim cloudRunV2JobIamMemberAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(crvjim.ref.Append("project"))
}

// Role returns a reference to field role of google_cloud_run_v2_job_iam_member.
func (crvjim cloudRunV2JobIamMemberAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(crvjim.ref.Append("role"))
}

func (crvjim cloudRunV2JobIamMemberAttributes) Condition() terra.ListValue[cloudrunv2jobiammember.ConditionAttributes] {
	return terra.ReferenceAsList[cloudrunv2jobiammember.ConditionAttributes](crvjim.ref.Append("condition"))
}

type cloudRunV2JobIamMemberState struct {
	Etag      string                                  `json:"etag"`
	Id        string                                  `json:"id"`
	Location  string                                  `json:"location"`
	Member    string                                  `json:"member"`
	Name      string                                  `json:"name"`
	Project   string                                  `json:"project"`
	Role      string                                  `json:"role"`
	Condition []cloudrunv2jobiammember.ConditionState `json:"condition"`
}
