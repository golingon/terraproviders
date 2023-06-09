// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	servicedirectoryserviceiammember "github.com/golingon/terraproviders/googlebeta/4.63.1/servicedirectoryserviceiammember"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewServiceDirectoryServiceIamMember creates a new instance of [ServiceDirectoryServiceIamMember].
func NewServiceDirectoryServiceIamMember(name string, args ServiceDirectoryServiceIamMemberArgs) *ServiceDirectoryServiceIamMember {
	return &ServiceDirectoryServiceIamMember{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ServiceDirectoryServiceIamMember)(nil)

// ServiceDirectoryServiceIamMember represents the Terraform resource google_service_directory_service_iam_member.
type ServiceDirectoryServiceIamMember struct {
	Name      string
	Args      ServiceDirectoryServiceIamMemberArgs
	state     *serviceDirectoryServiceIamMemberState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ServiceDirectoryServiceIamMember].
func (sdsim *ServiceDirectoryServiceIamMember) Type() string {
	return "google_service_directory_service_iam_member"
}

// LocalName returns the local name for [ServiceDirectoryServiceIamMember].
func (sdsim *ServiceDirectoryServiceIamMember) LocalName() string {
	return sdsim.Name
}

// Configuration returns the configuration (args) for [ServiceDirectoryServiceIamMember].
func (sdsim *ServiceDirectoryServiceIamMember) Configuration() interface{} {
	return sdsim.Args
}

// DependOn is used for other resources to depend on [ServiceDirectoryServiceIamMember].
func (sdsim *ServiceDirectoryServiceIamMember) DependOn() terra.Reference {
	return terra.ReferenceResource(sdsim)
}

// Dependencies returns the list of resources [ServiceDirectoryServiceIamMember] depends_on.
func (sdsim *ServiceDirectoryServiceIamMember) Dependencies() terra.Dependencies {
	return sdsim.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ServiceDirectoryServiceIamMember].
func (sdsim *ServiceDirectoryServiceIamMember) LifecycleManagement() *terra.Lifecycle {
	return sdsim.Lifecycle
}

// Attributes returns the attributes for [ServiceDirectoryServiceIamMember].
func (sdsim *ServiceDirectoryServiceIamMember) Attributes() serviceDirectoryServiceIamMemberAttributes {
	return serviceDirectoryServiceIamMemberAttributes{ref: terra.ReferenceResource(sdsim)}
}

// ImportState imports the given attribute values into [ServiceDirectoryServiceIamMember]'s state.
func (sdsim *ServiceDirectoryServiceIamMember) ImportState(av io.Reader) error {
	sdsim.state = &serviceDirectoryServiceIamMemberState{}
	if err := json.NewDecoder(av).Decode(sdsim.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sdsim.Type(), sdsim.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ServiceDirectoryServiceIamMember] has state.
func (sdsim *ServiceDirectoryServiceIamMember) State() (*serviceDirectoryServiceIamMemberState, bool) {
	return sdsim.state, sdsim.state != nil
}

// StateMust returns the state for [ServiceDirectoryServiceIamMember]. Panics if the state is nil.
func (sdsim *ServiceDirectoryServiceIamMember) StateMust() *serviceDirectoryServiceIamMemberState {
	if sdsim.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sdsim.Type(), sdsim.LocalName()))
	}
	return sdsim.state
}

// ServiceDirectoryServiceIamMemberArgs contains the configurations for google_service_directory_service_iam_member.
type ServiceDirectoryServiceIamMemberArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Member: string, required
	Member terra.StringValue `hcl:"member,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// Condition: optional
	Condition *servicedirectoryserviceiammember.Condition `hcl:"condition,block"`
}
type serviceDirectoryServiceIamMemberAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_service_directory_service_iam_member.
func (sdsim serviceDirectoryServiceIamMemberAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(sdsim.ref.Append("etag"))
}

// Id returns a reference to field id of google_service_directory_service_iam_member.
func (sdsim serviceDirectoryServiceIamMemberAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sdsim.ref.Append("id"))
}

// Member returns a reference to field member of google_service_directory_service_iam_member.
func (sdsim serviceDirectoryServiceIamMemberAttributes) Member() terra.StringValue {
	return terra.ReferenceAsString(sdsim.ref.Append("member"))
}

// Name returns a reference to field name of google_service_directory_service_iam_member.
func (sdsim serviceDirectoryServiceIamMemberAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(sdsim.ref.Append("name"))
}

// Role returns a reference to field role of google_service_directory_service_iam_member.
func (sdsim serviceDirectoryServiceIamMemberAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(sdsim.ref.Append("role"))
}

func (sdsim serviceDirectoryServiceIamMemberAttributes) Condition() terra.ListValue[servicedirectoryserviceiammember.ConditionAttributes] {
	return terra.ReferenceAsList[servicedirectoryserviceiammember.ConditionAttributes](sdsim.ref.Append("condition"))
}

type serviceDirectoryServiceIamMemberState struct {
	Etag      string                                            `json:"etag"`
	Id        string                                            `json:"id"`
	Member    string                                            `json:"member"`
	Name      string                                            `json:"name"`
	Role      string                                            `json:"role"`
	Condition []servicedirectoryserviceiammember.ConditionState `json:"condition"`
}
