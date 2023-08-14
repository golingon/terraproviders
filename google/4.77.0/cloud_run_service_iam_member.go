// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	cloudrunserviceiammember "github.com/golingon/terraproviders/google/4.77.0/cloudrunserviceiammember"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewCloudRunServiceIamMember creates a new instance of [CloudRunServiceIamMember].
func NewCloudRunServiceIamMember(name string, args CloudRunServiceIamMemberArgs) *CloudRunServiceIamMember {
	return &CloudRunServiceIamMember{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*CloudRunServiceIamMember)(nil)

// CloudRunServiceIamMember represents the Terraform resource google_cloud_run_service_iam_member.
type CloudRunServiceIamMember struct {
	Name      string
	Args      CloudRunServiceIamMemberArgs
	state     *cloudRunServiceIamMemberState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [CloudRunServiceIamMember].
func (crsim *CloudRunServiceIamMember) Type() string {
	return "google_cloud_run_service_iam_member"
}

// LocalName returns the local name for [CloudRunServiceIamMember].
func (crsim *CloudRunServiceIamMember) LocalName() string {
	return crsim.Name
}

// Configuration returns the configuration (args) for [CloudRunServiceIamMember].
func (crsim *CloudRunServiceIamMember) Configuration() interface{} {
	return crsim.Args
}

// DependOn is used for other resources to depend on [CloudRunServiceIamMember].
func (crsim *CloudRunServiceIamMember) DependOn() terra.Reference {
	return terra.ReferenceResource(crsim)
}

// Dependencies returns the list of resources [CloudRunServiceIamMember] depends_on.
func (crsim *CloudRunServiceIamMember) Dependencies() terra.Dependencies {
	return crsim.DependsOn
}

// LifecycleManagement returns the lifecycle block for [CloudRunServiceIamMember].
func (crsim *CloudRunServiceIamMember) LifecycleManagement() *terra.Lifecycle {
	return crsim.Lifecycle
}

// Attributes returns the attributes for [CloudRunServiceIamMember].
func (crsim *CloudRunServiceIamMember) Attributes() cloudRunServiceIamMemberAttributes {
	return cloudRunServiceIamMemberAttributes{ref: terra.ReferenceResource(crsim)}
}

// ImportState imports the given attribute values into [CloudRunServiceIamMember]'s state.
func (crsim *CloudRunServiceIamMember) ImportState(av io.Reader) error {
	crsim.state = &cloudRunServiceIamMemberState{}
	if err := json.NewDecoder(av).Decode(crsim.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", crsim.Type(), crsim.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [CloudRunServiceIamMember] has state.
func (crsim *CloudRunServiceIamMember) State() (*cloudRunServiceIamMemberState, bool) {
	return crsim.state, crsim.state != nil
}

// StateMust returns the state for [CloudRunServiceIamMember]. Panics if the state is nil.
func (crsim *CloudRunServiceIamMember) StateMust() *cloudRunServiceIamMemberState {
	if crsim.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", crsim.Type(), crsim.LocalName()))
	}
	return crsim.state
}

// CloudRunServiceIamMemberArgs contains the configurations for google_cloud_run_service_iam_member.
type CloudRunServiceIamMemberArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, optional
	Location terra.StringValue `hcl:"location,attr"`
	// Member: string, required
	Member terra.StringValue `hcl:"member,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// Service: string, required
	Service terra.StringValue `hcl:"service,attr" validate:"required"`
	// Condition: optional
	Condition *cloudrunserviceiammember.Condition `hcl:"condition,block"`
}
type cloudRunServiceIamMemberAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_cloud_run_service_iam_member.
func (crsim cloudRunServiceIamMemberAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(crsim.ref.Append("etag"))
}

// Id returns a reference to field id of google_cloud_run_service_iam_member.
func (crsim cloudRunServiceIamMemberAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(crsim.ref.Append("id"))
}

// Location returns a reference to field location of google_cloud_run_service_iam_member.
func (crsim cloudRunServiceIamMemberAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(crsim.ref.Append("location"))
}

// Member returns a reference to field member of google_cloud_run_service_iam_member.
func (crsim cloudRunServiceIamMemberAttributes) Member() terra.StringValue {
	return terra.ReferenceAsString(crsim.ref.Append("member"))
}

// Project returns a reference to field project of google_cloud_run_service_iam_member.
func (crsim cloudRunServiceIamMemberAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(crsim.ref.Append("project"))
}

// Role returns a reference to field role of google_cloud_run_service_iam_member.
func (crsim cloudRunServiceIamMemberAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(crsim.ref.Append("role"))
}

// Service returns a reference to field service of google_cloud_run_service_iam_member.
func (crsim cloudRunServiceIamMemberAttributes) Service() terra.StringValue {
	return terra.ReferenceAsString(crsim.ref.Append("service"))
}

func (crsim cloudRunServiceIamMemberAttributes) Condition() terra.ListValue[cloudrunserviceiammember.ConditionAttributes] {
	return terra.ReferenceAsList[cloudrunserviceiammember.ConditionAttributes](crsim.ref.Append("condition"))
}

type cloudRunServiceIamMemberState struct {
	Etag      string                                    `json:"etag"`
	Id        string                                    `json:"id"`
	Location  string                                    `json:"location"`
	Member    string                                    `json:"member"`
	Project   string                                    `json:"project"`
	Role      string                                    `json:"role"`
	Service   string                                    `json:"service"`
	Condition []cloudrunserviceiammember.ConditionState `json:"condition"`
}
