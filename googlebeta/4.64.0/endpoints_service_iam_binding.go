// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	endpointsserviceiambinding "github.com/golingon/terraproviders/googlebeta/4.64.0/endpointsserviceiambinding"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewEndpointsServiceIamBinding creates a new instance of [EndpointsServiceIamBinding].
func NewEndpointsServiceIamBinding(name string, args EndpointsServiceIamBindingArgs) *EndpointsServiceIamBinding {
	return &EndpointsServiceIamBinding{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*EndpointsServiceIamBinding)(nil)

// EndpointsServiceIamBinding represents the Terraform resource google_endpoints_service_iam_binding.
type EndpointsServiceIamBinding struct {
	Name      string
	Args      EndpointsServiceIamBindingArgs
	state     *endpointsServiceIamBindingState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [EndpointsServiceIamBinding].
func (esib *EndpointsServiceIamBinding) Type() string {
	return "google_endpoints_service_iam_binding"
}

// LocalName returns the local name for [EndpointsServiceIamBinding].
func (esib *EndpointsServiceIamBinding) LocalName() string {
	return esib.Name
}

// Configuration returns the configuration (args) for [EndpointsServiceIamBinding].
func (esib *EndpointsServiceIamBinding) Configuration() interface{} {
	return esib.Args
}

// DependOn is used for other resources to depend on [EndpointsServiceIamBinding].
func (esib *EndpointsServiceIamBinding) DependOn() terra.Reference {
	return terra.ReferenceResource(esib)
}

// Dependencies returns the list of resources [EndpointsServiceIamBinding] depends_on.
func (esib *EndpointsServiceIamBinding) Dependencies() terra.Dependencies {
	return esib.DependsOn
}

// LifecycleManagement returns the lifecycle block for [EndpointsServiceIamBinding].
func (esib *EndpointsServiceIamBinding) LifecycleManagement() *terra.Lifecycle {
	return esib.Lifecycle
}

// Attributes returns the attributes for [EndpointsServiceIamBinding].
func (esib *EndpointsServiceIamBinding) Attributes() endpointsServiceIamBindingAttributes {
	return endpointsServiceIamBindingAttributes{ref: terra.ReferenceResource(esib)}
}

// ImportState imports the given attribute values into [EndpointsServiceIamBinding]'s state.
func (esib *EndpointsServiceIamBinding) ImportState(av io.Reader) error {
	esib.state = &endpointsServiceIamBindingState{}
	if err := json.NewDecoder(av).Decode(esib.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", esib.Type(), esib.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [EndpointsServiceIamBinding] has state.
func (esib *EndpointsServiceIamBinding) State() (*endpointsServiceIamBindingState, bool) {
	return esib.state, esib.state != nil
}

// StateMust returns the state for [EndpointsServiceIamBinding]. Panics if the state is nil.
func (esib *EndpointsServiceIamBinding) StateMust() *endpointsServiceIamBindingState {
	if esib.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", esib.Type(), esib.LocalName()))
	}
	return esib.state
}

// EndpointsServiceIamBindingArgs contains the configurations for google_endpoints_service_iam_binding.
type EndpointsServiceIamBindingArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Members: set of string, required
	Members terra.SetValue[terra.StringValue] `hcl:"members,attr" validate:"required"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// ServiceName: string, required
	ServiceName terra.StringValue `hcl:"service_name,attr" validate:"required"`
	// Condition: optional
	Condition *endpointsserviceiambinding.Condition `hcl:"condition,block"`
}
type endpointsServiceIamBindingAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_endpoints_service_iam_binding.
func (esib endpointsServiceIamBindingAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(esib.ref.Append("etag"))
}

// Id returns a reference to field id of google_endpoints_service_iam_binding.
func (esib endpointsServiceIamBindingAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(esib.ref.Append("id"))
}

// Members returns a reference to field members of google_endpoints_service_iam_binding.
func (esib endpointsServiceIamBindingAttributes) Members() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](esib.ref.Append("members"))
}

// Role returns a reference to field role of google_endpoints_service_iam_binding.
func (esib endpointsServiceIamBindingAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(esib.ref.Append("role"))
}

// ServiceName returns a reference to field service_name of google_endpoints_service_iam_binding.
func (esib endpointsServiceIamBindingAttributes) ServiceName() terra.StringValue {
	return terra.ReferenceAsString(esib.ref.Append("service_name"))
}

func (esib endpointsServiceIamBindingAttributes) Condition() terra.ListValue[endpointsserviceiambinding.ConditionAttributes] {
	return terra.ReferenceAsList[endpointsserviceiambinding.ConditionAttributes](esib.ref.Append("condition"))
}

type endpointsServiceIamBindingState struct {
	Etag        string                                      `json:"etag"`
	Id          string                                      `json:"id"`
	Members     []string                                    `json:"members"`
	Role        string                                      `json:"role"`
	ServiceName string                                      `json:"service_name"`
	Condition   []endpointsserviceiambinding.ConditionState `json:"condition"`
}
