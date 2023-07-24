// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	projectservice "github.com/golingon/terraproviders/googlebeta/4.74.0/projectservice"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewProjectService creates a new instance of [ProjectService].
func NewProjectService(name string, args ProjectServiceArgs) *ProjectService {
	return &ProjectService{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ProjectService)(nil)

// ProjectService represents the Terraform resource google_project_service.
type ProjectService struct {
	Name      string
	Args      ProjectServiceArgs
	state     *projectServiceState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ProjectService].
func (ps *ProjectService) Type() string {
	return "google_project_service"
}

// LocalName returns the local name for [ProjectService].
func (ps *ProjectService) LocalName() string {
	return ps.Name
}

// Configuration returns the configuration (args) for [ProjectService].
func (ps *ProjectService) Configuration() interface{} {
	return ps.Args
}

// DependOn is used for other resources to depend on [ProjectService].
func (ps *ProjectService) DependOn() terra.Reference {
	return terra.ReferenceResource(ps)
}

// Dependencies returns the list of resources [ProjectService] depends_on.
func (ps *ProjectService) Dependencies() terra.Dependencies {
	return ps.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ProjectService].
func (ps *ProjectService) LifecycleManagement() *terra.Lifecycle {
	return ps.Lifecycle
}

// Attributes returns the attributes for [ProjectService].
func (ps *ProjectService) Attributes() projectServiceAttributes {
	return projectServiceAttributes{ref: terra.ReferenceResource(ps)}
}

// ImportState imports the given attribute values into [ProjectService]'s state.
func (ps *ProjectService) ImportState(av io.Reader) error {
	ps.state = &projectServiceState{}
	if err := json.NewDecoder(av).Decode(ps.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ps.Type(), ps.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ProjectService] has state.
func (ps *ProjectService) State() (*projectServiceState, bool) {
	return ps.state, ps.state != nil
}

// StateMust returns the state for [ProjectService]. Panics if the state is nil.
func (ps *ProjectService) StateMust() *projectServiceState {
	if ps.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ps.Type(), ps.LocalName()))
	}
	return ps.state
}

// ProjectServiceArgs contains the configurations for google_project_service.
type ProjectServiceArgs struct {
	// DisableDependentServices: bool, optional
	DisableDependentServices terra.BoolValue `hcl:"disable_dependent_services,attr"`
	// DisableOnDestroy: bool, optional
	DisableOnDestroy terra.BoolValue `hcl:"disable_on_destroy,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Service: string, required
	Service terra.StringValue `hcl:"service,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *projectservice.Timeouts `hcl:"timeouts,block"`
}
type projectServiceAttributes struct {
	ref terra.Reference
}

// DisableDependentServices returns a reference to field disable_dependent_services of google_project_service.
func (ps projectServiceAttributes) DisableDependentServices() terra.BoolValue {
	return terra.ReferenceAsBool(ps.ref.Append("disable_dependent_services"))
}

// DisableOnDestroy returns a reference to field disable_on_destroy of google_project_service.
func (ps projectServiceAttributes) DisableOnDestroy() terra.BoolValue {
	return terra.ReferenceAsBool(ps.ref.Append("disable_on_destroy"))
}

// Id returns a reference to field id of google_project_service.
func (ps projectServiceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ps.ref.Append("id"))
}

// Project returns a reference to field project of google_project_service.
func (ps projectServiceAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(ps.ref.Append("project"))
}

// Service returns a reference to field service of google_project_service.
func (ps projectServiceAttributes) Service() terra.StringValue {
	return terra.ReferenceAsString(ps.ref.Append("service"))
}

func (ps projectServiceAttributes) Timeouts() projectservice.TimeoutsAttributes {
	return terra.ReferenceAsSingle[projectservice.TimeoutsAttributes](ps.ref.Append("timeouts"))
}

type projectServiceState struct {
	DisableDependentServices bool                          `json:"disable_dependent_services"`
	DisableOnDestroy         bool                          `json:"disable_on_destroy"`
	Id                       string                        `json:"id"`
	Project                  string                        `json:"project"`
	Service                  string                        `json:"service"`
	Timeouts                 *projectservice.TimeoutsState `json:"timeouts"`
}
