// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	computesharedvpchostproject "github.com/golingon/terraproviders/googlebeta/4.62.0/computesharedvpchostproject"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewComputeSharedVpcHostProject creates a new instance of [ComputeSharedVpcHostProject].
func NewComputeSharedVpcHostProject(name string, args ComputeSharedVpcHostProjectArgs) *ComputeSharedVpcHostProject {
	return &ComputeSharedVpcHostProject{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ComputeSharedVpcHostProject)(nil)

// ComputeSharedVpcHostProject represents the Terraform resource google_compute_shared_vpc_host_project.
type ComputeSharedVpcHostProject struct {
	Name      string
	Args      ComputeSharedVpcHostProjectArgs
	state     *computeSharedVpcHostProjectState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ComputeSharedVpcHostProject].
func (csvhp *ComputeSharedVpcHostProject) Type() string {
	return "google_compute_shared_vpc_host_project"
}

// LocalName returns the local name for [ComputeSharedVpcHostProject].
func (csvhp *ComputeSharedVpcHostProject) LocalName() string {
	return csvhp.Name
}

// Configuration returns the configuration (args) for [ComputeSharedVpcHostProject].
func (csvhp *ComputeSharedVpcHostProject) Configuration() interface{} {
	return csvhp.Args
}

// DependOn is used for other resources to depend on [ComputeSharedVpcHostProject].
func (csvhp *ComputeSharedVpcHostProject) DependOn() terra.Reference {
	return terra.ReferenceResource(csvhp)
}

// Dependencies returns the list of resources [ComputeSharedVpcHostProject] depends_on.
func (csvhp *ComputeSharedVpcHostProject) Dependencies() terra.Dependencies {
	return csvhp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ComputeSharedVpcHostProject].
func (csvhp *ComputeSharedVpcHostProject) LifecycleManagement() *terra.Lifecycle {
	return csvhp.Lifecycle
}

// Attributes returns the attributes for [ComputeSharedVpcHostProject].
func (csvhp *ComputeSharedVpcHostProject) Attributes() computeSharedVpcHostProjectAttributes {
	return computeSharedVpcHostProjectAttributes{ref: terra.ReferenceResource(csvhp)}
}

// ImportState imports the given attribute values into [ComputeSharedVpcHostProject]'s state.
func (csvhp *ComputeSharedVpcHostProject) ImportState(av io.Reader) error {
	csvhp.state = &computeSharedVpcHostProjectState{}
	if err := json.NewDecoder(av).Decode(csvhp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", csvhp.Type(), csvhp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ComputeSharedVpcHostProject] has state.
func (csvhp *ComputeSharedVpcHostProject) State() (*computeSharedVpcHostProjectState, bool) {
	return csvhp.state, csvhp.state != nil
}

// StateMust returns the state for [ComputeSharedVpcHostProject]. Panics if the state is nil.
func (csvhp *ComputeSharedVpcHostProject) StateMust() *computeSharedVpcHostProjectState {
	if csvhp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", csvhp.Type(), csvhp.LocalName()))
	}
	return csvhp.state
}

// ComputeSharedVpcHostProjectArgs contains the configurations for google_compute_shared_vpc_host_project.
type ComputeSharedVpcHostProjectArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Project: string, required
	Project terra.StringValue `hcl:"project,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *computesharedvpchostproject.Timeouts `hcl:"timeouts,block"`
}
type computeSharedVpcHostProjectAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of google_compute_shared_vpc_host_project.
func (csvhp computeSharedVpcHostProjectAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(csvhp.ref.Append("id"))
}

// Project returns a reference to field project of google_compute_shared_vpc_host_project.
func (csvhp computeSharedVpcHostProjectAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(csvhp.ref.Append("project"))
}

func (csvhp computeSharedVpcHostProjectAttributes) Timeouts() computesharedvpchostproject.TimeoutsAttributes {
	return terra.ReferenceAsSingle[computesharedvpchostproject.TimeoutsAttributes](csvhp.ref.Append("timeouts"))
}

type computeSharedVpcHostProjectState struct {
	Id       string                                     `json:"id"`
	Project  string                                     `json:"project"`
	Timeouts *computesharedvpchostproject.TimeoutsState `json:"timeouts"`
}
