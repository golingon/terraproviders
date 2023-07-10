// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	cloudiotregistryiammember "github.com/golingon/terraproviders/googlebeta/4.72.1/cloudiotregistryiammember"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewCloudiotRegistryIamMember creates a new instance of [CloudiotRegistryIamMember].
func NewCloudiotRegistryIamMember(name string, args CloudiotRegistryIamMemberArgs) *CloudiotRegistryIamMember {
	return &CloudiotRegistryIamMember{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*CloudiotRegistryIamMember)(nil)

// CloudiotRegistryIamMember represents the Terraform resource google_cloudiot_registry_iam_member.
type CloudiotRegistryIamMember struct {
	Name      string
	Args      CloudiotRegistryIamMemberArgs
	state     *cloudiotRegistryIamMemberState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [CloudiotRegistryIamMember].
func (crim *CloudiotRegistryIamMember) Type() string {
	return "google_cloudiot_registry_iam_member"
}

// LocalName returns the local name for [CloudiotRegistryIamMember].
func (crim *CloudiotRegistryIamMember) LocalName() string {
	return crim.Name
}

// Configuration returns the configuration (args) for [CloudiotRegistryIamMember].
func (crim *CloudiotRegistryIamMember) Configuration() interface{} {
	return crim.Args
}

// DependOn is used for other resources to depend on [CloudiotRegistryIamMember].
func (crim *CloudiotRegistryIamMember) DependOn() terra.Reference {
	return terra.ReferenceResource(crim)
}

// Dependencies returns the list of resources [CloudiotRegistryIamMember] depends_on.
func (crim *CloudiotRegistryIamMember) Dependencies() terra.Dependencies {
	return crim.DependsOn
}

// LifecycleManagement returns the lifecycle block for [CloudiotRegistryIamMember].
func (crim *CloudiotRegistryIamMember) LifecycleManagement() *terra.Lifecycle {
	return crim.Lifecycle
}

// Attributes returns the attributes for [CloudiotRegistryIamMember].
func (crim *CloudiotRegistryIamMember) Attributes() cloudiotRegistryIamMemberAttributes {
	return cloudiotRegistryIamMemberAttributes{ref: terra.ReferenceResource(crim)}
}

// ImportState imports the given attribute values into [CloudiotRegistryIamMember]'s state.
func (crim *CloudiotRegistryIamMember) ImportState(av io.Reader) error {
	crim.state = &cloudiotRegistryIamMemberState{}
	if err := json.NewDecoder(av).Decode(crim.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", crim.Type(), crim.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [CloudiotRegistryIamMember] has state.
func (crim *CloudiotRegistryIamMember) State() (*cloudiotRegistryIamMemberState, bool) {
	return crim.state, crim.state != nil
}

// StateMust returns the state for [CloudiotRegistryIamMember]. Panics if the state is nil.
func (crim *CloudiotRegistryIamMember) StateMust() *cloudiotRegistryIamMemberState {
	if crim.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", crim.Type(), crim.LocalName()))
	}
	return crim.state
}

// CloudiotRegistryIamMemberArgs contains the configurations for google_cloudiot_registry_iam_member.
type CloudiotRegistryIamMemberArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Member: string, required
	Member terra.StringValue `hcl:"member,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Region: string, optional
	Region terra.StringValue `hcl:"region,attr"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// Condition: optional
	Condition *cloudiotregistryiammember.Condition `hcl:"condition,block"`
}
type cloudiotRegistryIamMemberAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_cloudiot_registry_iam_member.
func (crim cloudiotRegistryIamMemberAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(crim.ref.Append("etag"))
}

// Id returns a reference to field id of google_cloudiot_registry_iam_member.
func (crim cloudiotRegistryIamMemberAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(crim.ref.Append("id"))
}

// Member returns a reference to field member of google_cloudiot_registry_iam_member.
func (crim cloudiotRegistryIamMemberAttributes) Member() terra.StringValue {
	return terra.ReferenceAsString(crim.ref.Append("member"))
}

// Name returns a reference to field name of google_cloudiot_registry_iam_member.
func (crim cloudiotRegistryIamMemberAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(crim.ref.Append("name"))
}

// Project returns a reference to field project of google_cloudiot_registry_iam_member.
func (crim cloudiotRegistryIamMemberAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(crim.ref.Append("project"))
}

// Region returns a reference to field region of google_cloudiot_registry_iam_member.
func (crim cloudiotRegistryIamMemberAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(crim.ref.Append("region"))
}

// Role returns a reference to field role of google_cloudiot_registry_iam_member.
func (crim cloudiotRegistryIamMemberAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(crim.ref.Append("role"))
}

func (crim cloudiotRegistryIamMemberAttributes) Condition() terra.ListValue[cloudiotregistryiammember.ConditionAttributes] {
	return terra.ReferenceAsList[cloudiotregistryiammember.ConditionAttributes](crim.ref.Append("condition"))
}

type cloudiotRegistryIamMemberState struct {
	Etag      string                                     `json:"etag"`
	Id        string                                     `json:"id"`
	Member    string                                     `json:"member"`
	Name      string                                     `json:"name"`
	Project   string                                     `json:"project"`
	Region    string                                     `json:"region"`
	Role      string                                     `json:"role"`
	Condition []cloudiotregistryiammember.ConditionState `json:"condition"`
}
