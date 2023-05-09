// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	computeprojectmetadataitem "github.com/golingon/terraproviders/google/4.63.1/computeprojectmetadataitem"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewComputeProjectMetadataItem creates a new instance of [ComputeProjectMetadataItem].
func NewComputeProjectMetadataItem(name string, args ComputeProjectMetadataItemArgs) *ComputeProjectMetadataItem {
	return &ComputeProjectMetadataItem{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ComputeProjectMetadataItem)(nil)

// ComputeProjectMetadataItem represents the Terraform resource google_compute_project_metadata_item.
type ComputeProjectMetadataItem struct {
	Name      string
	Args      ComputeProjectMetadataItemArgs
	state     *computeProjectMetadataItemState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ComputeProjectMetadataItem].
func (cpmi *ComputeProjectMetadataItem) Type() string {
	return "google_compute_project_metadata_item"
}

// LocalName returns the local name for [ComputeProjectMetadataItem].
func (cpmi *ComputeProjectMetadataItem) LocalName() string {
	return cpmi.Name
}

// Configuration returns the configuration (args) for [ComputeProjectMetadataItem].
func (cpmi *ComputeProjectMetadataItem) Configuration() interface{} {
	return cpmi.Args
}

// DependOn is used for other resources to depend on [ComputeProjectMetadataItem].
func (cpmi *ComputeProjectMetadataItem) DependOn() terra.Reference {
	return terra.ReferenceResource(cpmi)
}

// Dependencies returns the list of resources [ComputeProjectMetadataItem] depends_on.
func (cpmi *ComputeProjectMetadataItem) Dependencies() terra.Dependencies {
	return cpmi.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ComputeProjectMetadataItem].
func (cpmi *ComputeProjectMetadataItem) LifecycleManagement() *terra.Lifecycle {
	return cpmi.Lifecycle
}

// Attributes returns the attributes for [ComputeProjectMetadataItem].
func (cpmi *ComputeProjectMetadataItem) Attributes() computeProjectMetadataItemAttributes {
	return computeProjectMetadataItemAttributes{ref: terra.ReferenceResource(cpmi)}
}

// ImportState imports the given attribute values into [ComputeProjectMetadataItem]'s state.
func (cpmi *ComputeProjectMetadataItem) ImportState(av io.Reader) error {
	cpmi.state = &computeProjectMetadataItemState{}
	if err := json.NewDecoder(av).Decode(cpmi.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cpmi.Type(), cpmi.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ComputeProjectMetadataItem] has state.
func (cpmi *ComputeProjectMetadataItem) State() (*computeProjectMetadataItemState, bool) {
	return cpmi.state, cpmi.state != nil
}

// StateMust returns the state for [ComputeProjectMetadataItem]. Panics if the state is nil.
func (cpmi *ComputeProjectMetadataItem) StateMust() *computeProjectMetadataItemState {
	if cpmi.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cpmi.Type(), cpmi.LocalName()))
	}
	return cpmi.state
}

// ComputeProjectMetadataItemArgs contains the configurations for google_compute_project_metadata_item.
type ComputeProjectMetadataItemArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Key: string, required
	Key terra.StringValue `hcl:"key,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Value: string, required
	Value terra.StringValue `hcl:"value,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *computeprojectmetadataitem.Timeouts `hcl:"timeouts,block"`
}
type computeProjectMetadataItemAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of google_compute_project_metadata_item.
func (cpmi computeProjectMetadataItemAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cpmi.ref.Append("id"))
}

// Key returns a reference to field key of google_compute_project_metadata_item.
func (cpmi computeProjectMetadataItemAttributes) Key() terra.StringValue {
	return terra.ReferenceAsString(cpmi.ref.Append("key"))
}

// Project returns a reference to field project of google_compute_project_metadata_item.
func (cpmi computeProjectMetadataItemAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(cpmi.ref.Append("project"))
}

// Value returns a reference to field value of google_compute_project_metadata_item.
func (cpmi computeProjectMetadataItemAttributes) Value() terra.StringValue {
	return terra.ReferenceAsString(cpmi.ref.Append("value"))
}

func (cpmi computeProjectMetadataItemAttributes) Timeouts() computeprojectmetadataitem.TimeoutsAttributes {
	return terra.ReferenceAsSingle[computeprojectmetadataitem.TimeoutsAttributes](cpmi.ref.Append("timeouts"))
}

type computeProjectMetadataItemState struct {
	Id       string                                    `json:"id"`
	Key      string                                    `json:"key"`
	Project  string                                    `json:"project"`
	Value    string                                    `json:"value"`
	Timeouts *computeprojectmetadataitem.TimeoutsState `json:"timeouts"`
}
