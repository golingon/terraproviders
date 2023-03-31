// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	endpointsserviceconsumersiammember "github.com/golingon/terraproviders/google/4.59.0/endpointsserviceconsumersiammember"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewEndpointsServiceConsumersIamMember(name string, args EndpointsServiceConsumersIamMemberArgs) *EndpointsServiceConsumersIamMember {
	return &EndpointsServiceConsumersIamMember{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*EndpointsServiceConsumersIamMember)(nil)

type EndpointsServiceConsumersIamMember struct {
	Name  string
	Args  EndpointsServiceConsumersIamMemberArgs
	state *endpointsServiceConsumersIamMemberState
}

func (escim *EndpointsServiceConsumersIamMember) Type() string {
	return "google_endpoints_service_consumers_iam_member"
}

func (escim *EndpointsServiceConsumersIamMember) LocalName() string {
	return escim.Name
}

func (escim *EndpointsServiceConsumersIamMember) Configuration() interface{} {
	return escim.Args
}

func (escim *EndpointsServiceConsumersIamMember) Attributes() endpointsServiceConsumersIamMemberAttributes {
	return endpointsServiceConsumersIamMemberAttributes{ref: terra.ReferenceResource(escim)}
}

func (escim *EndpointsServiceConsumersIamMember) ImportState(av io.Reader) error {
	escim.state = &endpointsServiceConsumersIamMemberState{}
	if err := json.NewDecoder(av).Decode(escim.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", escim.Type(), escim.LocalName(), err)
	}
	return nil
}

func (escim *EndpointsServiceConsumersIamMember) State() (*endpointsServiceConsumersIamMemberState, bool) {
	return escim.state, escim.state != nil
}

func (escim *EndpointsServiceConsumersIamMember) StateMust() *endpointsServiceConsumersIamMemberState {
	if escim.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", escim.Type(), escim.LocalName()))
	}
	return escim.state
}

func (escim *EndpointsServiceConsumersIamMember) DependOn() terra.Reference {
	return terra.ReferenceResource(escim)
}

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
	// DependsOn contains resources that EndpointsServiceConsumersIamMember depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type endpointsServiceConsumersIamMemberAttributes struct {
	ref terra.Reference
}

func (escim endpointsServiceConsumersIamMemberAttributes) ConsumerProject() terra.StringValue {
	return terra.ReferenceString(escim.ref.Append("consumer_project"))
}

func (escim endpointsServiceConsumersIamMemberAttributes) Etag() terra.StringValue {
	return terra.ReferenceString(escim.ref.Append("etag"))
}

func (escim endpointsServiceConsumersIamMemberAttributes) Id() terra.StringValue {
	return terra.ReferenceString(escim.ref.Append("id"))
}

func (escim endpointsServiceConsumersIamMemberAttributes) Member() terra.StringValue {
	return terra.ReferenceString(escim.ref.Append("member"))
}

func (escim endpointsServiceConsumersIamMemberAttributes) Role() terra.StringValue {
	return terra.ReferenceString(escim.ref.Append("role"))
}

func (escim endpointsServiceConsumersIamMemberAttributes) ServiceName() terra.StringValue {
	return terra.ReferenceString(escim.ref.Append("service_name"))
}

func (escim endpointsServiceConsumersIamMemberAttributes) Condition() terra.ListValue[endpointsserviceconsumersiammember.ConditionAttributes] {
	return terra.ReferenceList[endpointsserviceconsumersiammember.ConditionAttributes](escim.ref.Append("condition"))
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
