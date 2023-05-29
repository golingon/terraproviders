// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewContainerRegistry creates a new instance of [ContainerRegistry].
func NewContainerRegistry(name string, args ContainerRegistryArgs) *ContainerRegistry {
	return &ContainerRegistry{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ContainerRegistry)(nil)

// ContainerRegistry represents the Terraform resource google_container_registry.
type ContainerRegistry struct {
	Name      string
	Args      ContainerRegistryArgs
	state     *containerRegistryState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ContainerRegistry].
func (cr *ContainerRegistry) Type() string {
	return "google_container_registry"
}

// LocalName returns the local name for [ContainerRegistry].
func (cr *ContainerRegistry) LocalName() string {
	return cr.Name
}

// Configuration returns the configuration (args) for [ContainerRegistry].
func (cr *ContainerRegistry) Configuration() interface{} {
	return cr.Args
}

// DependOn is used for other resources to depend on [ContainerRegistry].
func (cr *ContainerRegistry) DependOn() terra.Reference {
	return terra.ReferenceResource(cr)
}

// Dependencies returns the list of resources [ContainerRegistry] depends_on.
func (cr *ContainerRegistry) Dependencies() terra.Dependencies {
	return cr.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ContainerRegistry].
func (cr *ContainerRegistry) LifecycleManagement() *terra.Lifecycle {
	return cr.Lifecycle
}

// Attributes returns the attributes for [ContainerRegistry].
func (cr *ContainerRegistry) Attributes() containerRegistryAttributes {
	return containerRegistryAttributes{ref: terra.ReferenceResource(cr)}
}

// ImportState imports the given attribute values into [ContainerRegistry]'s state.
func (cr *ContainerRegistry) ImportState(av io.Reader) error {
	cr.state = &containerRegistryState{}
	if err := json.NewDecoder(av).Decode(cr.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cr.Type(), cr.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ContainerRegistry] has state.
func (cr *ContainerRegistry) State() (*containerRegistryState, bool) {
	return cr.state, cr.state != nil
}

// StateMust returns the state for [ContainerRegistry]. Panics if the state is nil.
func (cr *ContainerRegistry) StateMust() *containerRegistryState {
	if cr.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cr.Type(), cr.LocalName()))
	}
	return cr.state
}

// ContainerRegistryArgs contains the configurations for google_container_registry.
type ContainerRegistryArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, optional
	Location terra.StringValue `hcl:"location,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
}
type containerRegistryAttributes struct {
	ref terra.Reference
}

// BucketSelfLink returns a reference to field bucket_self_link of google_container_registry.
func (cr containerRegistryAttributes) BucketSelfLink() terra.StringValue {
	return terra.ReferenceAsString(cr.ref.Append("bucket_self_link"))
}

// Id returns a reference to field id of google_container_registry.
func (cr containerRegistryAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cr.ref.Append("id"))
}

// Location returns a reference to field location of google_container_registry.
func (cr containerRegistryAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(cr.ref.Append("location"))
}

// Project returns a reference to field project of google_container_registry.
func (cr containerRegistryAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(cr.ref.Append("project"))
}

type containerRegistryState struct {
	BucketSelfLink string `json:"bucket_self_link"`
	Id             string `json:"id"`
	Location       string `json:"location"`
	Project        string `json:"project"`
}