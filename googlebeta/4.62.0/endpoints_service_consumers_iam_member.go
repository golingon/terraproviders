// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	endpointsserviceconsumersiammember "github.com/golingon/terraproviders/googlebeta/4.62.0/endpointsserviceconsumersiammember"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewEndpointsServiceConsumersIamMember creates a new instance of [EndpointsServiceConsumersIamMember].
func NewEndpointsServiceConsumersIamMember(name string, args EndpointsServiceConsumersIamMemberArgs) *EndpointsServiceConsumersIamMember {
	return &EndpointsServiceConsumersIamMember{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*EndpointsServiceConsumersIamMember)(nil)

// EndpointsServiceConsumersIamMember represents the Terraform resource google_endpoints_service_consumers_iam_member.
type EndpointsServiceConsumersIamMember struct {
	Name      string
	Args      EndpointsServiceConsumersIamMemberArgs
	state     *endpointsServiceConsumersIamMemberState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [EndpointsServiceConsumersIamMember].
func (escim *EndpointsServiceConsumersIamMember) Type() string {
	return "google_endpoints_service_consumers_iam_member"
}

// LocalName returns the local name for [EndpointsServiceConsumersIamMember].
func (escim *EndpointsServiceConsumersIamMember) LocalName() string {
	return escim.Name
}

// Configuration returns the configuration (args) for [EndpointsServiceConsumersIamMember].
func (escim *EndpointsServiceConsumersIamMember) Configuration() interface{} {
	return escim.Args
}

// DependOn is used for other resources to depend on [EndpointsServiceConsumersIamMember].
func (escim *EndpointsServiceConsumersIamMember) DependOn() terra.Reference {
	return terra.ReferenceResource(escim)
}

// Dependencies returns the list of resources [EndpointsServiceConsumersIamMember] depends_on.
func (escim *EndpointsServiceConsumersIamMember) Dependencies() terra.Dependencies {
	return escim.DependsOn
}

// LifecycleManagement returns the lifecycle block for [EndpointsServiceConsumersIamMember].
func (escim *EndpointsServiceConsumersIamMember) LifecycleManagement() *terra.Lifecycle {
	return escim.Lifecycle
}

// Attributes returns the attributes for [EndpointsServiceConsumersIamMember].
func (escim *EndpointsServiceConsumersIamMember) Attributes() endpointsServiceConsumersIamMemberAttributes {
	return endpointsServiceConsumersIamMemberAttributes{ref: terra.ReferenceResource(escim)}
}

// ImportState imports the given attribute values into [EndpointsServiceConsumersIamMember]'s state.
func (escim *EndpointsServiceConsumersIamMember) ImportState(av io.Reader) error {
	escim.state = &endpointsServiceConsumersIamMemberState{}
	if err := json.NewDecoder(av).Decode(escim.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", escim.Type(), escim.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [EndpointsServiceConsumersIamMember] has state.
func (escim *EndpointsServiceConsumersIamMember) State() (*endpointsServiceConsumersIamMemberState, bool) {
	return escim.state, escim.state != nil
}

// StateMust returns the state for [EndpointsServiceConsumersIamMember]. Panics if the state is nil.
func (escim *EndpointsServiceConsumersIamMember) StateMust() *endpointsServiceConsumersIamMemberState {
	if escim.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", escim.Type(), escim.LocalName()))
	}
	return escim.state
}

// EndpointsServiceConsumersIamMemberArgs contains the configurations for google_endpoints_service_consumers_iam_member.
type EndpointsServiceConsumersIamMemberArgs struct {
	// ConsumerProject: string, required
	ConsumerProject terra.StringValue `hcl:"consumer_project,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Member: string, required
	Member terra.StringValue `hcl:"member,attr" validate:"required"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// ServiceName: string, required
	ServiceName terra.StringValue `hcl:"service_name,attr" validate:"required"`
	// Condition: optional
	Condition *endpointsserviceconsumersiammember.Condition `hcl:"condition,block"`
}
type endpointsServiceConsumersIamMemberAttributes struct {
	ref terra.Reference
}

// ConsumerProject returns a reference to field consumer_project of google_endpoints_service_consumers_iam_member.
func (escim endpointsServiceConsumersIamMemberAttributes) ConsumerProject() terra.StringValue {
	return terra.ReferenceAsString(escim.ref.Append("consumer_project"))
}

// Etag returns a reference to field etag of google_endpoints_service_consumers_iam_member.
func (escim endpointsServiceConsumersIamMemberAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(escim.ref.Append("etag"))
}

// Id returns a reference to field id of google_endpoints_service_consumers_iam_member.
func (escim endpointsServiceConsumersIamMemberAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(escim.ref.Append("id"))
}

// Member returns a reference to field member of google_endpoints_service_consumers_iam_member.
func (escim endpointsServiceConsumersIamMemberAttributes) Member() terra.StringValue {
	return terra.ReferenceAsString(escim.ref.Append("member"))
}

// Role returns a reference to field role of google_endpoints_service_consumers_iam_member.
func (escim endpointsServiceConsumersIamMemberAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(escim.ref.Append("role"))
}

// ServiceName returns a reference to field service_name of google_endpoints_service_consumers_iam_member.
func (escim endpointsServiceConsumersIamMemberAttributes) ServiceName() terra.StringValue {
	return terra.ReferenceAsString(escim.ref.Append("service_name"))
}

func (escim endpointsServiceConsumersIamMemberAttributes) Condition() terra.ListValue[endpointsserviceconsumersiammember.ConditionAttributes] {
	return terra.ReferenceAsList[endpointsserviceconsumersiammember.ConditionAttributes](escim.ref.Append("condition"))
}

type endpointsServiceConsumersIamMemberState struct {
	ConsumerProject string                                              `json:"consumer_project"`
	Etag            string                                              `json:"etag"`
	Id              string                                              `json:"id"`
	Member          string                                              `json:"member"`
	Role            string                                              `json:"role"`
	ServiceName     string                                              `json:"service_name"`
	Condition       []endpointsserviceconsumersiammember.ConditionState `json:"condition"`
}
