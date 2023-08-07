// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	endpointsserviceiammember "github.com/golingon/terraproviders/googlebeta/4.76.0/endpointsserviceiammember"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewEndpointsServiceIamMember creates a new instance of [EndpointsServiceIamMember].
func NewEndpointsServiceIamMember(name string, args EndpointsServiceIamMemberArgs) *EndpointsServiceIamMember {
	return &EndpointsServiceIamMember{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*EndpointsServiceIamMember)(nil)

// EndpointsServiceIamMember represents the Terraform resource google_endpoints_service_iam_member.
type EndpointsServiceIamMember struct {
	Name      string
	Args      EndpointsServiceIamMemberArgs
	state     *endpointsServiceIamMemberState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [EndpointsServiceIamMember].
func (esim *EndpointsServiceIamMember) Type() string {
	return "google_endpoints_service_iam_member"
}

// LocalName returns the local name for [EndpointsServiceIamMember].
func (esim *EndpointsServiceIamMember) LocalName() string {
	return esim.Name
}

// Configuration returns the configuration (args) for [EndpointsServiceIamMember].
func (esim *EndpointsServiceIamMember) Configuration() interface{} {
	return esim.Args
}

// DependOn is used for other resources to depend on [EndpointsServiceIamMember].
func (esim *EndpointsServiceIamMember) DependOn() terra.Reference {
	return terra.ReferenceResource(esim)
}

// Dependencies returns the list of resources [EndpointsServiceIamMember] depends_on.
func (esim *EndpointsServiceIamMember) Dependencies() terra.Dependencies {
	return esim.DependsOn
}

// LifecycleManagement returns the lifecycle block for [EndpointsServiceIamMember].
func (esim *EndpointsServiceIamMember) LifecycleManagement() *terra.Lifecycle {
	return esim.Lifecycle
}

// Attributes returns the attributes for [EndpointsServiceIamMember].
func (esim *EndpointsServiceIamMember) Attributes() endpointsServiceIamMemberAttributes {
	return endpointsServiceIamMemberAttributes{ref: terra.ReferenceResource(esim)}
}

// ImportState imports the given attribute values into [EndpointsServiceIamMember]'s state.
func (esim *EndpointsServiceIamMember) ImportState(av io.Reader) error {
	esim.state = &endpointsServiceIamMemberState{}
	if err := json.NewDecoder(av).Decode(esim.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", esim.Type(), esim.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [EndpointsServiceIamMember] has state.
func (esim *EndpointsServiceIamMember) State() (*endpointsServiceIamMemberState, bool) {
	return esim.state, esim.state != nil
}

// StateMust returns the state for [EndpointsServiceIamMember]. Panics if the state is nil.
func (esim *EndpointsServiceIamMember) StateMust() *endpointsServiceIamMemberState {
	if esim.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", esim.Type(), esim.LocalName()))
	}
	return esim.state
}

// EndpointsServiceIamMemberArgs contains the configurations for google_endpoints_service_iam_member.
type EndpointsServiceIamMemberArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Member: string, required
	Member terra.StringValue `hcl:"member,attr" validate:"required"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// ServiceName: string, required
	ServiceName terra.StringValue `hcl:"service_name,attr" validate:"required"`
	// Condition: optional
	Condition *endpointsserviceiammember.Condition `hcl:"condition,block"`
}
type endpointsServiceIamMemberAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_endpoints_service_iam_member.
func (esim endpointsServiceIamMemberAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(esim.ref.Append("etag"))
}

// Id returns a reference to field id of google_endpoints_service_iam_member.
func (esim endpointsServiceIamMemberAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(esim.ref.Append("id"))
}

// Member returns a reference to field member of google_endpoints_service_iam_member.
func (esim endpointsServiceIamMemberAttributes) Member() terra.StringValue {
	return terra.ReferenceAsString(esim.ref.Append("member"))
}

// Role returns a reference to field role of google_endpoints_service_iam_member.
func (esim endpointsServiceIamMemberAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(esim.ref.Append("role"))
}

// ServiceName returns a reference to field service_name of google_endpoints_service_iam_member.
func (esim endpointsServiceIamMemberAttributes) ServiceName() terra.StringValue {
	return terra.ReferenceAsString(esim.ref.Append("service_name"))
}

func (esim endpointsServiceIamMemberAttributes) Condition() terra.ListValue[endpointsserviceiammember.ConditionAttributes] {
	return terra.ReferenceAsList[endpointsserviceiammember.ConditionAttributes](esim.ref.Append("condition"))
}

type endpointsServiceIamMemberState struct {
	Etag        string                                     `json:"etag"`
	Id          string                                     `json:"id"`
	Member      string                                     `json:"member"`
	Role        string                                     `json:"role"`
	ServiceName string                                     `json:"service_name"`
	Condition   []endpointsserviceiammember.ConditionState `json:"condition"`
}
