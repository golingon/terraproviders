// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	projectserviceidentity "github.com/golingon/terraproviders/googlebeta/4.71.0/projectserviceidentity"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewProjectServiceIdentity creates a new instance of [ProjectServiceIdentity].
func NewProjectServiceIdentity(name string, args ProjectServiceIdentityArgs) *ProjectServiceIdentity {
	return &ProjectServiceIdentity{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ProjectServiceIdentity)(nil)

// ProjectServiceIdentity represents the Terraform resource google_project_service_identity.
type ProjectServiceIdentity struct {
	Name      string
	Args      ProjectServiceIdentityArgs
	state     *projectServiceIdentityState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ProjectServiceIdentity].
func (psi *ProjectServiceIdentity) Type() string {
	return "google_project_service_identity"
}

// LocalName returns the local name for [ProjectServiceIdentity].
func (psi *ProjectServiceIdentity) LocalName() string {
	return psi.Name
}

// Configuration returns the configuration (args) for [ProjectServiceIdentity].
func (psi *ProjectServiceIdentity) Configuration() interface{} {
	return psi.Args
}

// DependOn is used for other resources to depend on [ProjectServiceIdentity].
func (psi *ProjectServiceIdentity) DependOn() terra.Reference {
	return terra.ReferenceResource(psi)
}

// Dependencies returns the list of resources [ProjectServiceIdentity] depends_on.
func (psi *ProjectServiceIdentity) Dependencies() terra.Dependencies {
	return psi.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ProjectServiceIdentity].
func (psi *ProjectServiceIdentity) LifecycleManagement() *terra.Lifecycle {
	return psi.Lifecycle
}

// Attributes returns the attributes for [ProjectServiceIdentity].
func (psi *ProjectServiceIdentity) Attributes() projectServiceIdentityAttributes {
	return projectServiceIdentityAttributes{ref: terra.ReferenceResource(psi)}
}

// ImportState imports the given attribute values into [ProjectServiceIdentity]'s state.
func (psi *ProjectServiceIdentity) ImportState(av io.Reader) error {
	psi.state = &projectServiceIdentityState{}
	if err := json.NewDecoder(av).Decode(psi.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", psi.Type(), psi.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ProjectServiceIdentity] has state.
func (psi *ProjectServiceIdentity) State() (*projectServiceIdentityState, bool) {
	return psi.state, psi.state != nil
}

// StateMust returns the state for [ProjectServiceIdentity]. Panics if the state is nil.
func (psi *ProjectServiceIdentity) StateMust() *projectServiceIdentityState {
	if psi.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", psi.Type(), psi.LocalName()))
	}
	return psi.state
}

// ProjectServiceIdentityArgs contains the configurations for google_project_service_identity.
type ProjectServiceIdentityArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Service: string, required
	Service terra.StringValue `hcl:"service,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *projectserviceidentity.Timeouts `hcl:"timeouts,block"`
}
type projectServiceIdentityAttributes struct {
	ref terra.Reference
}

// Email returns a reference to field email of google_project_service_identity.
func (psi projectServiceIdentityAttributes) Email() terra.StringValue {
	return terra.ReferenceAsString(psi.ref.Append("email"))
}

// Id returns a reference to field id of google_project_service_identity.
func (psi projectServiceIdentityAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(psi.ref.Append("id"))
}

// Project returns a reference to field project of google_project_service_identity.
func (psi projectServiceIdentityAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(psi.ref.Append("project"))
}

// Service returns a reference to field service of google_project_service_identity.
func (psi projectServiceIdentityAttributes) Service() terra.StringValue {
	return terra.ReferenceAsString(psi.ref.Append("service"))
}

func (psi projectServiceIdentityAttributes) Timeouts() projectserviceidentity.TimeoutsAttributes {
	return terra.ReferenceAsSingle[projectserviceidentity.TimeoutsAttributes](psi.ref.Append("timeouts"))
}

type projectServiceIdentityState struct {
	Email    string                                `json:"email"`
	Id       string                                `json:"id"`
	Project  string                                `json:"project"`
	Service  string                                `json:"service"`
	Timeouts *projectserviceidentity.TimeoutsState `json:"timeouts"`
}
