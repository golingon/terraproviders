// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	accesscontextmanagerserviceperimeters "github.com/golingon/terraproviders/google/4.76.0/accesscontextmanagerserviceperimeters"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewAccessContextManagerServicePerimeters creates a new instance of [AccessContextManagerServicePerimeters].
func NewAccessContextManagerServicePerimeters(name string, args AccessContextManagerServicePerimetersArgs) *AccessContextManagerServicePerimeters {
	return &AccessContextManagerServicePerimeters{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*AccessContextManagerServicePerimeters)(nil)

// AccessContextManagerServicePerimeters represents the Terraform resource google_access_context_manager_service_perimeters.
type AccessContextManagerServicePerimeters struct {
	Name      string
	Args      AccessContextManagerServicePerimetersArgs
	state     *accessContextManagerServicePerimetersState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [AccessContextManagerServicePerimeters].
func (acmsp *AccessContextManagerServicePerimeters) Type() string {
	return "google_access_context_manager_service_perimeters"
}

// LocalName returns the local name for [AccessContextManagerServicePerimeters].
func (acmsp *AccessContextManagerServicePerimeters) LocalName() string {
	return acmsp.Name
}

// Configuration returns the configuration (args) for [AccessContextManagerServicePerimeters].
func (acmsp *AccessContextManagerServicePerimeters) Configuration() interface{} {
	return acmsp.Args
}

// DependOn is used for other resources to depend on [AccessContextManagerServicePerimeters].
func (acmsp *AccessContextManagerServicePerimeters) DependOn() terra.Reference {
	return terra.ReferenceResource(acmsp)
}

// Dependencies returns the list of resources [AccessContextManagerServicePerimeters] depends_on.
func (acmsp *AccessContextManagerServicePerimeters) Dependencies() terra.Dependencies {
	return acmsp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [AccessContextManagerServicePerimeters].
func (acmsp *AccessContextManagerServicePerimeters) LifecycleManagement() *terra.Lifecycle {
	return acmsp.Lifecycle
}

// Attributes returns the attributes for [AccessContextManagerServicePerimeters].
func (acmsp *AccessContextManagerServicePerimeters) Attributes() accessContextManagerServicePerimetersAttributes {
	return accessContextManagerServicePerimetersAttributes{ref: terra.ReferenceResource(acmsp)}
}

// ImportState imports the given attribute values into [AccessContextManagerServicePerimeters]'s state.
func (acmsp *AccessContextManagerServicePerimeters) ImportState(av io.Reader) error {
	acmsp.state = &accessContextManagerServicePerimetersState{}
	if err := json.NewDecoder(av).Decode(acmsp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", acmsp.Type(), acmsp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [AccessContextManagerServicePerimeters] has state.
func (acmsp *AccessContextManagerServicePerimeters) State() (*accessContextManagerServicePerimetersState, bool) {
	return acmsp.state, acmsp.state != nil
}

// StateMust returns the state for [AccessContextManagerServicePerimeters]. Panics if the state is nil.
func (acmsp *AccessContextManagerServicePerimeters) StateMust() *accessContextManagerServicePerimetersState {
	if acmsp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", acmsp.Type(), acmsp.LocalName()))
	}
	return acmsp.state
}

// AccessContextManagerServicePerimetersArgs contains the configurations for google_access_context_manager_service_perimeters.
type AccessContextManagerServicePerimetersArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Parent: string, required
	Parent terra.StringValue `hcl:"parent,attr" validate:"required"`
	// ServicePerimeters: min=0
	ServicePerimeters []accesscontextmanagerserviceperimeters.ServicePerimeters `hcl:"service_perimeters,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *accesscontextmanagerserviceperimeters.Timeouts `hcl:"timeouts,block"`
}
type accessContextManagerServicePerimetersAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of google_access_context_manager_service_perimeters.
func (acmsp accessContextManagerServicePerimetersAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(acmsp.ref.Append("id"))
}

// Parent returns a reference to field parent of google_access_context_manager_service_perimeters.
func (acmsp accessContextManagerServicePerimetersAttributes) Parent() terra.StringValue {
	return terra.ReferenceAsString(acmsp.ref.Append("parent"))
}

func (acmsp accessContextManagerServicePerimetersAttributes) ServicePerimeters() terra.SetValue[accesscontextmanagerserviceperimeters.ServicePerimetersAttributes] {
	return terra.ReferenceAsSet[accesscontextmanagerserviceperimeters.ServicePerimetersAttributes](acmsp.ref.Append("service_perimeters"))
}

func (acmsp accessContextManagerServicePerimetersAttributes) Timeouts() accesscontextmanagerserviceperimeters.TimeoutsAttributes {
	return terra.ReferenceAsSingle[accesscontextmanagerserviceperimeters.TimeoutsAttributes](acmsp.ref.Append("timeouts"))
}

type accessContextManagerServicePerimetersState struct {
	Id                string                                                         `json:"id"`
	Parent            string                                                         `json:"parent"`
	ServicePerimeters []accesscontextmanagerserviceperimeters.ServicePerimetersState `json:"service_perimeters"`
	Timeouts          *accesscontextmanagerserviceperimeters.TimeoutsState           `json:"timeouts"`
}
