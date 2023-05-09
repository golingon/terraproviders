// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	servicedirectoryserviceiambinding "github.com/golingon/terraproviders/googlebeta/4.63.1/servicedirectoryserviceiambinding"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewServiceDirectoryServiceIamBinding creates a new instance of [ServiceDirectoryServiceIamBinding].
func NewServiceDirectoryServiceIamBinding(name string, args ServiceDirectoryServiceIamBindingArgs) *ServiceDirectoryServiceIamBinding {
	return &ServiceDirectoryServiceIamBinding{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ServiceDirectoryServiceIamBinding)(nil)

// ServiceDirectoryServiceIamBinding represents the Terraform resource google_service_directory_service_iam_binding.
type ServiceDirectoryServiceIamBinding struct {
	Name      string
	Args      ServiceDirectoryServiceIamBindingArgs
	state     *serviceDirectoryServiceIamBindingState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ServiceDirectoryServiceIamBinding].
func (sdsib *ServiceDirectoryServiceIamBinding) Type() string {
	return "google_service_directory_service_iam_binding"
}

// LocalName returns the local name for [ServiceDirectoryServiceIamBinding].
func (sdsib *ServiceDirectoryServiceIamBinding) LocalName() string {
	return sdsib.Name
}

// Configuration returns the configuration (args) for [ServiceDirectoryServiceIamBinding].
func (sdsib *ServiceDirectoryServiceIamBinding) Configuration() interface{} {
	return sdsib.Args
}

// DependOn is used for other resources to depend on [ServiceDirectoryServiceIamBinding].
func (sdsib *ServiceDirectoryServiceIamBinding) DependOn() terra.Reference {
	return terra.ReferenceResource(sdsib)
}

// Dependencies returns the list of resources [ServiceDirectoryServiceIamBinding] depends_on.
func (sdsib *ServiceDirectoryServiceIamBinding) Dependencies() terra.Dependencies {
	return sdsib.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ServiceDirectoryServiceIamBinding].
func (sdsib *ServiceDirectoryServiceIamBinding) LifecycleManagement() *terra.Lifecycle {
	return sdsib.Lifecycle
}

// Attributes returns the attributes for [ServiceDirectoryServiceIamBinding].
func (sdsib *ServiceDirectoryServiceIamBinding) Attributes() serviceDirectoryServiceIamBindingAttributes {
	return serviceDirectoryServiceIamBindingAttributes{ref: terra.ReferenceResource(sdsib)}
}

// ImportState imports the given attribute values into [ServiceDirectoryServiceIamBinding]'s state.
func (sdsib *ServiceDirectoryServiceIamBinding) ImportState(av io.Reader) error {
	sdsib.state = &serviceDirectoryServiceIamBindingState{}
	if err := json.NewDecoder(av).Decode(sdsib.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sdsib.Type(), sdsib.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ServiceDirectoryServiceIamBinding] has state.
func (sdsib *ServiceDirectoryServiceIamBinding) State() (*serviceDirectoryServiceIamBindingState, bool) {
	return sdsib.state, sdsib.state != nil
}

// StateMust returns the state for [ServiceDirectoryServiceIamBinding]. Panics if the state is nil.
func (sdsib *ServiceDirectoryServiceIamBinding) StateMust() *serviceDirectoryServiceIamBindingState {
	if sdsib.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sdsib.Type(), sdsib.LocalName()))
	}
	return sdsib.state
}

// ServiceDirectoryServiceIamBindingArgs contains the configurations for google_service_directory_service_iam_binding.
type ServiceDirectoryServiceIamBindingArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Members: set of string, required
	Members terra.SetValue[terra.StringValue] `hcl:"members,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// Condition: optional
	Condition *servicedirectoryserviceiambinding.Condition `hcl:"condition,block"`
}
type serviceDirectoryServiceIamBindingAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_service_directory_service_iam_binding.
func (sdsib serviceDirectoryServiceIamBindingAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(sdsib.ref.Append("etag"))
}

// Id returns a reference to field id of google_service_directory_service_iam_binding.
func (sdsib serviceDirectoryServiceIamBindingAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sdsib.ref.Append("id"))
}

// Members returns a reference to field members of google_service_directory_service_iam_binding.
func (sdsib serviceDirectoryServiceIamBindingAttributes) Members() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](sdsib.ref.Append("members"))
}

// Name returns a reference to field name of google_service_directory_service_iam_binding.
func (sdsib serviceDirectoryServiceIamBindingAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(sdsib.ref.Append("name"))
}

// Role returns a reference to field role of google_service_directory_service_iam_binding.
func (sdsib serviceDirectoryServiceIamBindingAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(sdsib.ref.Append("role"))
}

func (sdsib serviceDirectoryServiceIamBindingAttributes) Condition() terra.ListValue[servicedirectoryserviceiambinding.ConditionAttributes] {
	return terra.ReferenceAsList[servicedirectoryserviceiambinding.ConditionAttributes](sdsib.ref.Append("condition"))
}

type serviceDirectoryServiceIamBindingState struct {
	Etag      string                                             `json:"etag"`
	Id        string                                             `json:"id"`
	Members   []string                                           `json:"members"`
	Name      string                                             `json:"name"`
	Role      string                                             `json:"role"`
	Condition []servicedirectoryserviceiambinding.ConditionState `json:"condition"`
}
