// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	servicedirectorynamespaceiambinding "github.com/golingon/terraproviders/googlebeta/4.72.1/servicedirectorynamespaceiambinding"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewServiceDirectoryNamespaceIamBinding creates a new instance of [ServiceDirectoryNamespaceIamBinding].
func NewServiceDirectoryNamespaceIamBinding(name string, args ServiceDirectoryNamespaceIamBindingArgs) *ServiceDirectoryNamespaceIamBinding {
	return &ServiceDirectoryNamespaceIamBinding{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ServiceDirectoryNamespaceIamBinding)(nil)

// ServiceDirectoryNamespaceIamBinding represents the Terraform resource google_service_directory_namespace_iam_binding.
type ServiceDirectoryNamespaceIamBinding struct {
	Name      string
	Args      ServiceDirectoryNamespaceIamBindingArgs
	state     *serviceDirectoryNamespaceIamBindingState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ServiceDirectoryNamespaceIamBinding].
func (sdnib *ServiceDirectoryNamespaceIamBinding) Type() string {
	return "google_service_directory_namespace_iam_binding"
}

// LocalName returns the local name for [ServiceDirectoryNamespaceIamBinding].
func (sdnib *ServiceDirectoryNamespaceIamBinding) LocalName() string {
	return sdnib.Name
}

// Configuration returns the configuration (args) for [ServiceDirectoryNamespaceIamBinding].
func (sdnib *ServiceDirectoryNamespaceIamBinding) Configuration() interface{} {
	return sdnib.Args
}

// DependOn is used for other resources to depend on [ServiceDirectoryNamespaceIamBinding].
func (sdnib *ServiceDirectoryNamespaceIamBinding) DependOn() terra.Reference {
	return terra.ReferenceResource(sdnib)
}

// Dependencies returns the list of resources [ServiceDirectoryNamespaceIamBinding] depends_on.
func (sdnib *ServiceDirectoryNamespaceIamBinding) Dependencies() terra.Dependencies {
	return sdnib.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ServiceDirectoryNamespaceIamBinding].
func (sdnib *ServiceDirectoryNamespaceIamBinding) LifecycleManagement() *terra.Lifecycle {
	return sdnib.Lifecycle
}

// Attributes returns the attributes for [ServiceDirectoryNamespaceIamBinding].
func (sdnib *ServiceDirectoryNamespaceIamBinding) Attributes() serviceDirectoryNamespaceIamBindingAttributes {
	return serviceDirectoryNamespaceIamBindingAttributes{ref: terra.ReferenceResource(sdnib)}
}

// ImportState imports the given attribute values into [ServiceDirectoryNamespaceIamBinding]'s state.
func (sdnib *ServiceDirectoryNamespaceIamBinding) ImportState(av io.Reader) error {
	sdnib.state = &serviceDirectoryNamespaceIamBindingState{}
	if err := json.NewDecoder(av).Decode(sdnib.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sdnib.Type(), sdnib.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ServiceDirectoryNamespaceIamBinding] has state.
func (sdnib *ServiceDirectoryNamespaceIamBinding) State() (*serviceDirectoryNamespaceIamBindingState, bool) {
	return sdnib.state, sdnib.state != nil
}

// StateMust returns the state for [ServiceDirectoryNamespaceIamBinding]. Panics if the state is nil.
func (sdnib *ServiceDirectoryNamespaceIamBinding) StateMust() *serviceDirectoryNamespaceIamBindingState {
	if sdnib.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sdnib.Type(), sdnib.LocalName()))
	}
	return sdnib.state
}

// ServiceDirectoryNamespaceIamBindingArgs contains the configurations for google_service_directory_namespace_iam_binding.
type ServiceDirectoryNamespaceIamBindingArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Members: set of string, required
	Members terra.SetValue[terra.StringValue] `hcl:"members,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// Condition: optional
	Condition *servicedirectorynamespaceiambinding.Condition `hcl:"condition,block"`
}
type serviceDirectoryNamespaceIamBindingAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_service_directory_namespace_iam_binding.
func (sdnib serviceDirectoryNamespaceIamBindingAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(sdnib.ref.Append("etag"))
}

// Id returns a reference to field id of google_service_directory_namespace_iam_binding.
func (sdnib serviceDirectoryNamespaceIamBindingAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sdnib.ref.Append("id"))
}

// Members returns a reference to field members of google_service_directory_namespace_iam_binding.
func (sdnib serviceDirectoryNamespaceIamBindingAttributes) Members() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](sdnib.ref.Append("members"))
}

// Name returns a reference to field name of google_service_directory_namespace_iam_binding.
func (sdnib serviceDirectoryNamespaceIamBindingAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(sdnib.ref.Append("name"))
}

// Role returns a reference to field role of google_service_directory_namespace_iam_binding.
func (sdnib serviceDirectoryNamespaceIamBindingAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(sdnib.ref.Append("role"))
}

func (sdnib serviceDirectoryNamespaceIamBindingAttributes) Condition() terra.ListValue[servicedirectorynamespaceiambinding.ConditionAttributes] {
	return terra.ReferenceAsList[servicedirectorynamespaceiambinding.ConditionAttributes](sdnib.ref.Append("condition"))
}

type serviceDirectoryNamespaceIamBindingState struct {
	Etag      string                                               `json:"etag"`
	Id        string                                               `json:"id"`
	Members   []string                                             `json:"members"`
	Name      string                                               `json:"name"`
	Role      string                                               `json:"role"`
	Condition []servicedirectorynamespaceiambinding.ConditionState `json:"condition"`
}
