// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	accesscontextmanagerserviceperimeterresource "github.com/golingon/terraproviders/google/4.72.1/accesscontextmanagerserviceperimeterresource"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewAccessContextManagerServicePerimeterResource creates a new instance of [AccessContextManagerServicePerimeterResource].
func NewAccessContextManagerServicePerimeterResource(name string, args AccessContextManagerServicePerimeterResourceArgs) *AccessContextManagerServicePerimeterResource {
	return &AccessContextManagerServicePerimeterResource{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*AccessContextManagerServicePerimeterResource)(nil)

// AccessContextManagerServicePerimeterResource represents the Terraform resource google_access_context_manager_service_perimeter_resource.
type AccessContextManagerServicePerimeterResource struct {
	Name      string
	Args      AccessContextManagerServicePerimeterResourceArgs
	state     *accessContextManagerServicePerimeterResourceState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [AccessContextManagerServicePerimeterResource].
func (acmspr *AccessContextManagerServicePerimeterResource) Type() string {
	return "google_access_context_manager_service_perimeter_resource"
}

// LocalName returns the local name for [AccessContextManagerServicePerimeterResource].
func (acmspr *AccessContextManagerServicePerimeterResource) LocalName() string {
	return acmspr.Name
}

// Configuration returns the configuration (args) for [AccessContextManagerServicePerimeterResource].
func (acmspr *AccessContextManagerServicePerimeterResource) Configuration() interface{} {
	return acmspr.Args
}

// DependOn is used for other resources to depend on [AccessContextManagerServicePerimeterResource].
func (acmspr *AccessContextManagerServicePerimeterResource) DependOn() terra.Reference {
	return terra.ReferenceResource(acmspr)
}

// Dependencies returns the list of resources [AccessContextManagerServicePerimeterResource] depends_on.
func (acmspr *AccessContextManagerServicePerimeterResource) Dependencies() terra.Dependencies {
	return acmspr.DependsOn
}

// LifecycleManagement returns the lifecycle block for [AccessContextManagerServicePerimeterResource].
func (acmspr *AccessContextManagerServicePerimeterResource) LifecycleManagement() *terra.Lifecycle {
	return acmspr.Lifecycle
}

// Attributes returns the attributes for [AccessContextManagerServicePerimeterResource].
func (acmspr *AccessContextManagerServicePerimeterResource) Attributes() accessContextManagerServicePerimeterResourceAttributes {
	return accessContextManagerServicePerimeterResourceAttributes{ref: terra.ReferenceResource(acmspr)}
}

// ImportState imports the given attribute values into [AccessContextManagerServicePerimeterResource]'s state.
func (acmspr *AccessContextManagerServicePerimeterResource) ImportState(av io.Reader) error {
	acmspr.state = &accessContextManagerServicePerimeterResourceState{}
	if err := json.NewDecoder(av).Decode(acmspr.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", acmspr.Type(), acmspr.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [AccessContextManagerServicePerimeterResource] has state.
func (acmspr *AccessContextManagerServicePerimeterResource) State() (*accessContextManagerServicePerimeterResourceState, bool) {
	return acmspr.state, acmspr.state != nil
}

// StateMust returns the state for [AccessContextManagerServicePerimeterResource]. Panics if the state is nil.
func (acmspr *AccessContextManagerServicePerimeterResource) StateMust() *accessContextManagerServicePerimeterResourceState {
	if acmspr.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", acmspr.Type(), acmspr.LocalName()))
	}
	return acmspr.state
}

// AccessContextManagerServicePerimeterResourceArgs contains the configurations for google_access_context_manager_service_perimeter_resource.
type AccessContextManagerServicePerimeterResourceArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// PerimeterName: string, required
	PerimeterName terra.StringValue `hcl:"perimeter_name,attr" validate:"required"`
	// Resource: string, required
	Resource terra.StringValue `hcl:"resource,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *accesscontextmanagerserviceperimeterresource.Timeouts `hcl:"timeouts,block"`
}
type accessContextManagerServicePerimeterResourceAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of google_access_context_manager_service_perimeter_resource.
func (acmspr accessContextManagerServicePerimeterResourceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(acmspr.ref.Append("id"))
}

// PerimeterName returns a reference to field perimeter_name of google_access_context_manager_service_perimeter_resource.
func (acmspr accessContextManagerServicePerimeterResourceAttributes) PerimeterName() terra.StringValue {
	return terra.ReferenceAsString(acmspr.ref.Append("perimeter_name"))
}

// Resource returns a reference to field resource of google_access_context_manager_service_perimeter_resource.
func (acmspr accessContextManagerServicePerimeterResourceAttributes) Resource() terra.StringValue {
	return terra.ReferenceAsString(acmspr.ref.Append("resource"))
}

func (acmspr accessContextManagerServicePerimeterResourceAttributes) Timeouts() accesscontextmanagerserviceperimeterresource.TimeoutsAttributes {
	return terra.ReferenceAsSingle[accesscontextmanagerserviceperimeterresource.TimeoutsAttributes](acmspr.ref.Append("timeouts"))
}

type accessContextManagerServicePerimeterResourceState struct {
	Id            string                                                      `json:"id"`
	PerimeterName string                                                      `json:"perimeter_name"`
	Resource      string                                                      `json:"resource"`
	Timeouts      *accesscontextmanagerserviceperimeterresource.TimeoutsState `json:"timeouts"`
}
