// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	computesharedvpcserviceproject "github.com/golingon/terraproviders/googlebeta/4.64.0/computesharedvpcserviceproject"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewComputeSharedVpcServiceProject creates a new instance of [ComputeSharedVpcServiceProject].
func NewComputeSharedVpcServiceProject(name string, args ComputeSharedVpcServiceProjectArgs) *ComputeSharedVpcServiceProject {
	return &ComputeSharedVpcServiceProject{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ComputeSharedVpcServiceProject)(nil)

// ComputeSharedVpcServiceProject represents the Terraform resource google_compute_shared_vpc_service_project.
type ComputeSharedVpcServiceProject struct {
	Name      string
	Args      ComputeSharedVpcServiceProjectArgs
	state     *computeSharedVpcServiceProjectState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ComputeSharedVpcServiceProject].
func (csvsp *ComputeSharedVpcServiceProject) Type() string {
	return "google_compute_shared_vpc_service_project"
}

// LocalName returns the local name for [ComputeSharedVpcServiceProject].
func (csvsp *ComputeSharedVpcServiceProject) LocalName() string {
	return csvsp.Name
}

// Configuration returns the configuration (args) for [ComputeSharedVpcServiceProject].
func (csvsp *ComputeSharedVpcServiceProject) Configuration() interface{} {
	return csvsp.Args
}

// DependOn is used for other resources to depend on [ComputeSharedVpcServiceProject].
func (csvsp *ComputeSharedVpcServiceProject) DependOn() terra.Reference {
	return terra.ReferenceResource(csvsp)
}

// Dependencies returns the list of resources [ComputeSharedVpcServiceProject] depends_on.
func (csvsp *ComputeSharedVpcServiceProject) Dependencies() terra.Dependencies {
	return csvsp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ComputeSharedVpcServiceProject].
func (csvsp *ComputeSharedVpcServiceProject) LifecycleManagement() *terra.Lifecycle {
	return csvsp.Lifecycle
}

// Attributes returns the attributes for [ComputeSharedVpcServiceProject].
func (csvsp *ComputeSharedVpcServiceProject) Attributes() computeSharedVpcServiceProjectAttributes {
	return computeSharedVpcServiceProjectAttributes{ref: terra.ReferenceResource(csvsp)}
}

// ImportState imports the given attribute values into [ComputeSharedVpcServiceProject]'s state.
func (csvsp *ComputeSharedVpcServiceProject) ImportState(av io.Reader) error {
	csvsp.state = &computeSharedVpcServiceProjectState{}
	if err := json.NewDecoder(av).Decode(csvsp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", csvsp.Type(), csvsp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ComputeSharedVpcServiceProject] has state.
func (csvsp *ComputeSharedVpcServiceProject) State() (*computeSharedVpcServiceProjectState, bool) {
	return csvsp.state, csvsp.state != nil
}

// StateMust returns the state for [ComputeSharedVpcServiceProject]. Panics if the state is nil.
func (csvsp *ComputeSharedVpcServiceProject) StateMust() *computeSharedVpcServiceProjectState {
	if csvsp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", csvsp.Type(), csvsp.LocalName()))
	}
	return csvsp.state
}

// ComputeSharedVpcServiceProjectArgs contains the configurations for google_compute_shared_vpc_service_project.
type ComputeSharedVpcServiceProjectArgs struct {
	// DeletionPolicy: string, optional
	DeletionPolicy terra.StringValue `hcl:"deletion_policy,attr"`
	// HostProject: string, required
	HostProject terra.StringValue `hcl:"host_project,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ServiceProject: string, required
	ServiceProject terra.StringValue `hcl:"service_project,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *computesharedvpcserviceproject.Timeouts `hcl:"timeouts,block"`
}
type computeSharedVpcServiceProjectAttributes struct {
	ref terra.Reference
}

// DeletionPolicy returns a reference to field deletion_policy of google_compute_shared_vpc_service_project.
func (csvsp computeSharedVpcServiceProjectAttributes) DeletionPolicy() terra.StringValue {
	return terra.ReferenceAsString(csvsp.ref.Append("deletion_policy"))
}

// HostProject returns a reference to field host_project of google_compute_shared_vpc_service_project.
func (csvsp computeSharedVpcServiceProjectAttributes) HostProject() terra.StringValue {
	return terra.ReferenceAsString(csvsp.ref.Append("host_project"))
}

// Id returns a reference to field id of google_compute_shared_vpc_service_project.
func (csvsp computeSharedVpcServiceProjectAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(csvsp.ref.Append("id"))
}

// ServiceProject returns a reference to field service_project of google_compute_shared_vpc_service_project.
func (csvsp computeSharedVpcServiceProjectAttributes) ServiceProject() terra.StringValue {
	return terra.ReferenceAsString(csvsp.ref.Append("service_project"))
}

func (csvsp computeSharedVpcServiceProjectAttributes) Timeouts() computesharedvpcserviceproject.TimeoutsAttributes {
	return terra.ReferenceAsSingle[computesharedvpcserviceproject.TimeoutsAttributes](csvsp.ref.Append("timeouts"))
}

type computeSharedVpcServiceProjectState struct {
	DeletionPolicy string                                        `json:"deletion_policy"`
	HostProject    string                                        `json:"host_project"`
	Id             string                                        `json:"id"`
	ServiceProject string                                        `json:"service_project"`
	Timeouts       *computesharedvpcserviceproject.TimeoutsState `json:"timeouts"`
}
