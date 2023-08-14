// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	computeprojectmetadata "github.com/golingon/terraproviders/google/4.77.0/computeprojectmetadata"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewComputeProjectMetadata creates a new instance of [ComputeProjectMetadata].
func NewComputeProjectMetadata(name string, args ComputeProjectMetadataArgs) *ComputeProjectMetadata {
	return &ComputeProjectMetadata{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ComputeProjectMetadata)(nil)

// ComputeProjectMetadata represents the Terraform resource google_compute_project_metadata.
type ComputeProjectMetadata struct {
	Name      string
	Args      ComputeProjectMetadataArgs
	state     *computeProjectMetadataState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ComputeProjectMetadata].
func (cpm *ComputeProjectMetadata) Type() string {
	return "google_compute_project_metadata"
}

// LocalName returns the local name for [ComputeProjectMetadata].
func (cpm *ComputeProjectMetadata) LocalName() string {
	return cpm.Name
}

// Configuration returns the configuration (args) for [ComputeProjectMetadata].
func (cpm *ComputeProjectMetadata) Configuration() interface{} {
	return cpm.Args
}

// DependOn is used for other resources to depend on [ComputeProjectMetadata].
func (cpm *ComputeProjectMetadata) DependOn() terra.Reference {
	return terra.ReferenceResource(cpm)
}

// Dependencies returns the list of resources [ComputeProjectMetadata] depends_on.
func (cpm *ComputeProjectMetadata) Dependencies() terra.Dependencies {
	return cpm.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ComputeProjectMetadata].
func (cpm *ComputeProjectMetadata) LifecycleManagement() *terra.Lifecycle {
	return cpm.Lifecycle
}

// Attributes returns the attributes for [ComputeProjectMetadata].
func (cpm *ComputeProjectMetadata) Attributes() computeProjectMetadataAttributes {
	return computeProjectMetadataAttributes{ref: terra.ReferenceResource(cpm)}
}

// ImportState imports the given attribute values into [ComputeProjectMetadata]'s state.
func (cpm *ComputeProjectMetadata) ImportState(av io.Reader) error {
	cpm.state = &computeProjectMetadataState{}
	if err := json.NewDecoder(av).Decode(cpm.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cpm.Type(), cpm.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ComputeProjectMetadata] has state.
func (cpm *ComputeProjectMetadata) State() (*computeProjectMetadataState, bool) {
	return cpm.state, cpm.state != nil
}

// StateMust returns the state for [ComputeProjectMetadata]. Panics if the state is nil.
func (cpm *ComputeProjectMetadata) StateMust() *computeProjectMetadataState {
	if cpm.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cpm.Type(), cpm.LocalName()))
	}
	return cpm.state
}

// ComputeProjectMetadataArgs contains the configurations for google_compute_project_metadata.
type ComputeProjectMetadataArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Metadata: map of string, required
	Metadata terra.MapValue[terra.StringValue] `hcl:"metadata,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Timeouts: optional
	Timeouts *computeprojectmetadata.Timeouts `hcl:"timeouts,block"`
}
type computeProjectMetadataAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of google_compute_project_metadata.
func (cpm computeProjectMetadataAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cpm.ref.Append("id"))
}

// Metadata returns a reference to field metadata of google_compute_project_metadata.
func (cpm computeProjectMetadataAttributes) Metadata() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](cpm.ref.Append("metadata"))
}

// Project returns a reference to field project of google_compute_project_metadata.
func (cpm computeProjectMetadataAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(cpm.ref.Append("project"))
}

func (cpm computeProjectMetadataAttributes) Timeouts() computeprojectmetadata.TimeoutsAttributes {
	return terra.ReferenceAsSingle[computeprojectmetadata.TimeoutsAttributes](cpm.ref.Append("timeouts"))
}

type computeProjectMetadataState struct {
	Id       string                                `json:"id"`
	Metadata map[string]string                     `json:"metadata"`
	Project  string                                `json:"project"`
	Timeouts *computeprojectmetadata.TimeoutsState `json:"timeouts"`
}
