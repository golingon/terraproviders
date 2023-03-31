// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	endpointsserviceconsumersiambinding "github.com/golingon/terraproviders/google/4.59.0/endpointsserviceconsumersiambinding"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewEndpointsServiceConsumersIamBinding(name string, args EndpointsServiceConsumersIamBindingArgs) *EndpointsServiceConsumersIamBinding {
	return &EndpointsServiceConsumersIamBinding{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*EndpointsServiceConsumersIamBinding)(nil)

type EndpointsServiceConsumersIamBinding struct {
	Name  string
	Args  EndpointsServiceConsumersIamBindingArgs
	state *endpointsServiceConsumersIamBindingState
}

func (escib *EndpointsServiceConsumersIamBinding) Type() string {
	return "google_endpoints_service_consumers_iam_binding"
}

func (escib *EndpointsServiceConsumersIamBinding) LocalName() string {
	return escib.Name
}

func (escib *EndpointsServiceConsumersIamBinding) Configuration() interface{} {
	return escib.Args
}

func (escib *EndpointsServiceConsumersIamBinding) Attributes() endpointsServiceConsumersIamBindingAttributes {
	return endpointsServiceConsumersIamBindingAttributes{ref: terra.ReferenceResource(escib)}
}

func (escib *EndpointsServiceConsumersIamBinding) ImportState(av io.Reader) error {
	escib.state = &endpointsServiceConsumersIamBindingState{}
	if err := json.NewDecoder(av).Decode(escib.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", escib.Type(), escib.LocalName(), err)
	}
	return nil
}

func (escib *EndpointsServiceConsumersIamBinding) State() (*endpointsServiceConsumersIamBindingState, bool) {
	return escib.state, escib.state != nil
}

func (escib *EndpointsServiceConsumersIamBinding) StateMust() *endpointsServiceConsumersIamBindingState {
	if escib.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", escib.Type(), escib.LocalName()))
	}
	return escib.state
}

func (escib *EndpointsServiceConsumersIamBinding) DependOn() terra.Reference {
	return terra.ReferenceResource(escib)
}

type EndpointsServiceConsumersIamBindingArgs struct {
	// ConsumerProject: string, required
	ConsumerProject terra.StringValue `hcl:"consumer_project,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Members: set of string, required
	Members terra.SetValue[terra.StringValue] `hcl:"members,attr" validate:"required"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// ServiceName: string, required
	ServiceName terra.StringValue `hcl:"service_name,attr" validate:"required"`
	// Condition: optional
	Condition *endpointsserviceconsumersiambinding.Condition `hcl:"condition,block"`
	// DependsOn contains resources that EndpointsServiceConsumersIamBinding depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type endpointsServiceConsumersIamBindingAttributes struct {
	ref terra.Reference
}

func (escib endpointsServiceConsumersIamBindingAttributes) ConsumerProject() terra.StringValue {
	return terra.ReferenceString(escib.ref.Append("consumer_project"))
}

func (escib endpointsServiceConsumersIamBindingAttributes) Etag() terra.StringValue {
	return terra.ReferenceString(escib.ref.Append("etag"))
}

func (escib endpointsServiceConsumersIamBindingAttributes) Id() terra.StringValue {
	return terra.ReferenceString(escib.ref.Append("id"))
}

func (escib endpointsServiceConsumersIamBindingAttributes) Members() terra.SetValue[terra.StringValue] {
	return terra.ReferenceSet[terra.StringValue](escib.ref.Append("members"))
}

func (escib endpointsServiceConsumersIamBindingAttributes) Role() terra.StringValue {
	return terra.ReferenceString(escib.ref.Append("role"))
}

func (escib endpointsServiceConsumersIamBindingAttributes) ServiceName() terra.StringValue {
	return terra.ReferenceString(escib.ref.Append("service_name"))
}

func (escib endpointsServiceConsumersIamBindingAttributes) Condition() terra.ListValue[endpointsserviceconsumersiambinding.ConditionAttributes] {
	return terra.ReferenceList[endpointsserviceconsumersiambinding.ConditionAttributes](escib.ref.Append("condition"))
}

type endpointsServiceConsumersIamBindingState struct {
	ConsumerProject string                                               `json:"consumer_project"`
	Etag            string                                               `json:"etag"`
	Id              string                                               `json:"id"`
	Members         []string                                             `json:"members"`
	Role            string                                               `json:"role"`
	ServiceName     string                                               `json:"service_name"`
	Condition       []endpointsserviceconsumersiambinding.ConditionState `json:"condition"`
}
