// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	notebooksruntime "github.com/golingon/terraproviders/googlebeta/4.72.1/notebooksruntime"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewNotebooksRuntime creates a new instance of [NotebooksRuntime].
func NewNotebooksRuntime(name string, args NotebooksRuntimeArgs) *NotebooksRuntime {
	return &NotebooksRuntime{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*NotebooksRuntime)(nil)

// NotebooksRuntime represents the Terraform resource google_notebooks_runtime.
type NotebooksRuntime struct {
	Name      string
	Args      NotebooksRuntimeArgs
	state     *notebooksRuntimeState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [NotebooksRuntime].
func (nr *NotebooksRuntime) Type() string {
	return "google_notebooks_runtime"
}

// LocalName returns the local name for [NotebooksRuntime].
func (nr *NotebooksRuntime) LocalName() string {
	return nr.Name
}

// Configuration returns the configuration (args) for [NotebooksRuntime].
func (nr *NotebooksRuntime) Configuration() interface{} {
	return nr.Args
}

// DependOn is used for other resources to depend on [NotebooksRuntime].
func (nr *NotebooksRuntime) DependOn() terra.Reference {
	return terra.ReferenceResource(nr)
}

// Dependencies returns the list of resources [NotebooksRuntime] depends_on.
func (nr *NotebooksRuntime) Dependencies() terra.Dependencies {
	return nr.DependsOn
}

// LifecycleManagement returns the lifecycle block for [NotebooksRuntime].
func (nr *NotebooksRuntime) LifecycleManagement() *terra.Lifecycle {
	return nr.Lifecycle
}

// Attributes returns the attributes for [NotebooksRuntime].
func (nr *NotebooksRuntime) Attributes() notebooksRuntimeAttributes {
	return notebooksRuntimeAttributes{ref: terra.ReferenceResource(nr)}
}

// ImportState imports the given attribute values into [NotebooksRuntime]'s state.
func (nr *NotebooksRuntime) ImportState(av io.Reader) error {
	nr.state = &notebooksRuntimeState{}
	if err := json.NewDecoder(av).Decode(nr.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", nr.Type(), nr.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [NotebooksRuntime] has state.
func (nr *NotebooksRuntime) State() (*notebooksRuntimeState, bool) {
	return nr.state, nr.state != nil
}

// StateMust returns the state for [NotebooksRuntime]. Panics if the state is nil.
func (nr *NotebooksRuntime) StateMust() *notebooksRuntimeState {
	if nr.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", nr.Type(), nr.LocalName()))
	}
	return nr.state
}

// NotebooksRuntimeArgs contains the configurations for google_notebooks_runtime.
type NotebooksRuntimeArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Metrics: min=0
	Metrics []notebooksruntime.Metrics `hcl:"metrics,block" validate:"min=0"`
	// AccessConfig: optional
	AccessConfig *notebooksruntime.AccessConfig `hcl:"access_config,block"`
	// SoftwareConfig: optional
	SoftwareConfig *notebooksruntime.SoftwareConfig `hcl:"software_config,block"`
	// Timeouts: optional
	Timeouts *notebooksruntime.Timeouts `hcl:"timeouts,block"`
	// VirtualMachine: optional
	VirtualMachine *notebooksruntime.VirtualMachine `hcl:"virtual_machine,block"`
}
type notebooksRuntimeAttributes struct {
	ref terra.Reference
}

// HealthState returns a reference to field health_state of google_notebooks_runtime.
func (nr notebooksRuntimeAttributes) HealthState() terra.StringValue {
	return terra.ReferenceAsString(nr.ref.Append("health_state"))
}

// Id returns a reference to field id of google_notebooks_runtime.
func (nr notebooksRuntimeAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(nr.ref.Append("id"))
}

// Location returns a reference to field location of google_notebooks_runtime.
func (nr notebooksRuntimeAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(nr.ref.Append("location"))
}

// Name returns a reference to field name of google_notebooks_runtime.
func (nr notebooksRuntimeAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(nr.ref.Append("name"))
}

// Project returns a reference to field project of google_notebooks_runtime.
func (nr notebooksRuntimeAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(nr.ref.Append("project"))
}

// State returns a reference to field state of google_notebooks_runtime.
func (nr notebooksRuntimeAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(nr.ref.Append("state"))
}

func (nr notebooksRuntimeAttributes) Metrics() terra.ListValue[notebooksruntime.MetricsAttributes] {
	return terra.ReferenceAsList[notebooksruntime.MetricsAttributes](nr.ref.Append("metrics"))
}

func (nr notebooksRuntimeAttributes) AccessConfig() terra.ListValue[notebooksruntime.AccessConfigAttributes] {
	return terra.ReferenceAsList[notebooksruntime.AccessConfigAttributes](nr.ref.Append("access_config"))
}

func (nr notebooksRuntimeAttributes) SoftwareConfig() terra.ListValue[notebooksruntime.SoftwareConfigAttributes] {
	return terra.ReferenceAsList[notebooksruntime.SoftwareConfigAttributes](nr.ref.Append("software_config"))
}

func (nr notebooksRuntimeAttributes) Timeouts() notebooksruntime.TimeoutsAttributes {
	return terra.ReferenceAsSingle[notebooksruntime.TimeoutsAttributes](nr.ref.Append("timeouts"))
}

func (nr notebooksRuntimeAttributes) VirtualMachine() terra.ListValue[notebooksruntime.VirtualMachineAttributes] {
	return terra.ReferenceAsList[notebooksruntime.VirtualMachineAttributes](nr.ref.Append("virtual_machine"))
}

type notebooksRuntimeState struct {
	HealthState    string                                 `json:"health_state"`
	Id             string                                 `json:"id"`
	Location       string                                 `json:"location"`
	Name           string                                 `json:"name"`
	Project        string                                 `json:"project"`
	State          string                                 `json:"state"`
	Metrics        []notebooksruntime.MetricsState        `json:"metrics"`
	AccessConfig   []notebooksruntime.AccessConfigState   `json:"access_config"`
	SoftwareConfig []notebooksruntime.SoftwareConfigState `json:"software_config"`
	Timeouts       *notebooksruntime.TimeoutsState        `json:"timeouts"`
	VirtualMachine []notebooksruntime.VirtualMachineState `json:"virtual_machine"`
}
