// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	cloudrunv2serviceiammember "github.com/golingon/terraproviders/googlebeta/4.75.1/cloudrunv2serviceiammember"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewCloudRunV2ServiceIamMember creates a new instance of [CloudRunV2ServiceIamMember].
func NewCloudRunV2ServiceIamMember(name string, args CloudRunV2ServiceIamMemberArgs) *CloudRunV2ServiceIamMember {
	return &CloudRunV2ServiceIamMember{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*CloudRunV2ServiceIamMember)(nil)

// CloudRunV2ServiceIamMember represents the Terraform resource google_cloud_run_v2_service_iam_member.
type CloudRunV2ServiceIamMember struct {
	Name      string
	Args      CloudRunV2ServiceIamMemberArgs
	state     *cloudRunV2ServiceIamMemberState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [CloudRunV2ServiceIamMember].
func (crvsim *CloudRunV2ServiceIamMember) Type() string {
	return "google_cloud_run_v2_service_iam_member"
}

// LocalName returns the local name for [CloudRunV2ServiceIamMember].
func (crvsim *CloudRunV2ServiceIamMember) LocalName() string {
	return crvsim.Name
}

// Configuration returns the configuration (args) for [CloudRunV2ServiceIamMember].
func (crvsim *CloudRunV2ServiceIamMember) Configuration() interface{} {
	return crvsim.Args
}

// DependOn is used for other resources to depend on [CloudRunV2ServiceIamMember].
func (crvsim *CloudRunV2ServiceIamMember) DependOn() terra.Reference {
	return terra.ReferenceResource(crvsim)
}

// Dependencies returns the list of resources [CloudRunV2ServiceIamMember] depends_on.
func (crvsim *CloudRunV2ServiceIamMember) Dependencies() terra.Dependencies {
	return crvsim.DependsOn
}

// LifecycleManagement returns the lifecycle block for [CloudRunV2ServiceIamMember].
func (crvsim *CloudRunV2ServiceIamMember) LifecycleManagement() *terra.Lifecycle {
	return crvsim.Lifecycle
}

// Attributes returns the attributes for [CloudRunV2ServiceIamMember].
func (crvsim *CloudRunV2ServiceIamMember) Attributes() cloudRunV2ServiceIamMemberAttributes {
	return cloudRunV2ServiceIamMemberAttributes{ref: terra.ReferenceResource(crvsim)}
}

// ImportState imports the given attribute values into [CloudRunV2ServiceIamMember]'s state.
func (crvsim *CloudRunV2ServiceIamMember) ImportState(av io.Reader) error {
	crvsim.state = &cloudRunV2ServiceIamMemberState{}
	if err := json.NewDecoder(av).Decode(crvsim.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", crvsim.Type(), crvsim.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [CloudRunV2ServiceIamMember] has state.
func (crvsim *CloudRunV2ServiceIamMember) State() (*cloudRunV2ServiceIamMemberState, bool) {
	return crvsim.state, crvsim.state != nil
}

// StateMust returns the state for [CloudRunV2ServiceIamMember]. Panics if the state is nil.
func (crvsim *CloudRunV2ServiceIamMember) StateMust() *cloudRunV2ServiceIamMemberState {
	if crvsim.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", crvsim.Type(), crvsim.LocalName()))
	}
	return crvsim.state
}

// CloudRunV2ServiceIamMemberArgs contains the configurations for google_cloud_run_v2_service_iam_member.
type CloudRunV2ServiceIamMemberArgs struct {
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
	Condition *cloudrunv2serviceiammember.Condition `hcl:"condition,block"`
}
type cloudRunV2ServiceIamMemberAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_cloud_run_v2_service_iam_member.
func (crvsim cloudRunV2ServiceIamMemberAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(crvsim.ref.Append("etag"))
}

// Id returns a reference to field id of google_cloud_run_v2_service_iam_member.
func (crvsim cloudRunV2ServiceIamMemberAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(crvsim.ref.Append("id"))
}

// Location returns a reference to field location of google_cloud_run_v2_service_iam_member.
func (crvsim cloudRunV2ServiceIamMemberAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(crvsim.ref.Append("location"))
}

// Member returns a reference to field member of google_cloud_run_v2_service_iam_member.
func (crvsim cloudRunV2ServiceIamMemberAttributes) Member() terra.StringValue {
	return terra.ReferenceAsString(crvsim.ref.Append("member"))
}

// Name returns a reference to field name of google_cloud_run_v2_service_iam_member.
func (crvsim cloudRunV2ServiceIamMemberAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(crvsim.ref.Append("name"))
}

// Project returns a reference to field project of google_cloud_run_v2_service_iam_member.
func (crvsim cloudRunV2ServiceIamMemberAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(crvsim.ref.Append("project"))
}

// Role returns a reference to field role of google_cloud_run_v2_service_iam_member.
func (crvsim cloudRunV2ServiceIamMemberAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(crvsim.ref.Append("role"))
}

func (crvsim cloudRunV2ServiceIamMemberAttributes) Condition() terra.ListValue[cloudrunv2serviceiammember.ConditionAttributes] {
	return terra.ReferenceAsList[cloudrunv2serviceiammember.ConditionAttributes](crvsim.ref.Append("condition"))
}

type cloudRunV2ServiceIamMemberState struct {
	Etag      string                                      `json:"etag"`
	Id        string                                      `json:"id"`
	Location  string                                      `json:"location"`
	Member    string                                      `json:"member"`
	Name      string                                      `json:"name"`
	Project   string                                      `json:"project"`
	Role      string                                      `json:"role"`
	Condition []cloudrunv2serviceiammember.ConditionState `json:"condition"`
}
