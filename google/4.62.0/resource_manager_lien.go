// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	resourcemanagerlien "github.com/golingon/terraproviders/google/4.62.0/resourcemanagerlien"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewResourceManagerLien creates a new instance of [ResourceManagerLien].
func NewResourceManagerLien(name string, args ResourceManagerLienArgs) *ResourceManagerLien {
	return &ResourceManagerLien{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ResourceManagerLien)(nil)

// ResourceManagerLien represents the Terraform resource google_resource_manager_lien.
type ResourceManagerLien struct {
	Name      string
	Args      ResourceManagerLienArgs
	state     *resourceManagerLienState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ResourceManagerLien].
func (rml *ResourceManagerLien) Type() string {
	return "google_resource_manager_lien"
}

// LocalName returns the local name for [ResourceManagerLien].
func (rml *ResourceManagerLien) LocalName() string {
	return rml.Name
}

// Configuration returns the configuration (args) for [ResourceManagerLien].
func (rml *ResourceManagerLien) Configuration() interface{} {
	return rml.Args
}

// DependOn is used for other resources to depend on [ResourceManagerLien].
func (rml *ResourceManagerLien) DependOn() terra.Reference {
	return terra.ReferenceResource(rml)
}

// Dependencies returns the list of resources [ResourceManagerLien] depends_on.
func (rml *ResourceManagerLien) Dependencies() terra.Dependencies {
	return rml.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ResourceManagerLien].
func (rml *ResourceManagerLien) LifecycleManagement() *terra.Lifecycle {
	return rml.Lifecycle
}

// Attributes returns the attributes for [ResourceManagerLien].
func (rml *ResourceManagerLien) Attributes() resourceManagerLienAttributes {
	return resourceManagerLienAttributes{ref: terra.ReferenceResource(rml)}
}

// ImportState imports the given attribute values into [ResourceManagerLien]'s state.
func (rml *ResourceManagerLien) ImportState(av io.Reader) error {
	rml.state = &resourceManagerLienState{}
	if err := json.NewDecoder(av).Decode(rml.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", rml.Type(), rml.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ResourceManagerLien] has state.
func (rml *ResourceManagerLien) State() (*resourceManagerLienState, bool) {
	return rml.state, rml.state != nil
}

// StateMust returns the state for [ResourceManagerLien]. Panics if the state is nil.
func (rml *ResourceManagerLien) StateMust() *resourceManagerLienState {
	if rml.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", rml.Type(), rml.LocalName()))
	}
	return rml.state
}

// ResourceManagerLienArgs contains the configurations for google_resource_manager_lien.
type ResourceManagerLienArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Origin: string, required
	Origin terra.StringValue `hcl:"origin,attr" validate:"required"`
	// Parent: string, required
	Parent terra.StringValue `hcl:"parent,attr" validate:"required"`
	// Reason: string, required
	Reason terra.StringValue `hcl:"reason,attr" validate:"required"`
	// Restrictions: list of string, required
	Restrictions terra.ListValue[terra.StringValue] `hcl:"restrictions,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *resourcemanagerlien.Timeouts `hcl:"timeouts,block"`
}
type resourceManagerLienAttributes struct {
	ref terra.Reference
}

// CreateTime returns a reference to field create_time of google_resource_manager_lien.
func (rml resourceManagerLienAttributes) CreateTime() terra.StringValue {
	return terra.ReferenceAsString(rml.ref.Append("create_time"))
}

// Id returns a reference to field id of google_resource_manager_lien.
func (rml resourceManagerLienAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(rml.ref.Append("id"))
}

// Name returns a reference to field name of google_resource_manager_lien.
func (rml resourceManagerLienAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(rml.ref.Append("name"))
}

// Origin returns a reference to field origin of google_resource_manager_lien.
func (rml resourceManagerLienAttributes) Origin() terra.StringValue {
	return terra.ReferenceAsString(rml.ref.Append("origin"))
}

// Parent returns a reference to field parent of google_resource_manager_lien.
func (rml resourceManagerLienAttributes) Parent() terra.StringValue {
	return terra.ReferenceAsString(rml.ref.Append("parent"))
}

// Reason returns a reference to field reason of google_resource_manager_lien.
func (rml resourceManagerLienAttributes) Reason() terra.StringValue {
	return terra.ReferenceAsString(rml.ref.Append("reason"))
}

// Restrictions returns a reference to field restrictions of google_resource_manager_lien.
func (rml resourceManagerLienAttributes) Restrictions() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](rml.ref.Append("restrictions"))
}

func (rml resourceManagerLienAttributes) Timeouts() resourcemanagerlien.TimeoutsAttributes {
	return terra.ReferenceAsSingle[resourcemanagerlien.TimeoutsAttributes](rml.ref.Append("timeouts"))
}

type resourceManagerLienState struct {
	CreateTime   string                             `json:"create_time"`
	Id           string                             `json:"id"`
	Name         string                             `json:"name"`
	Origin       string                             `json:"origin"`
	Parent       string                             `json:"parent"`
	Reason       string                             `json:"reason"`
	Restrictions []string                           `json:"restrictions"`
	Timeouts     *resourcemanagerlien.TimeoutsState `json:"timeouts"`
}
