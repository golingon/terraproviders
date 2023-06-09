// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	vertexaitensorboard "github.com/golingon/terraproviders/googlebeta/4.64.0/vertexaitensorboard"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewVertexAiTensorboard creates a new instance of [VertexAiTensorboard].
func NewVertexAiTensorboard(name string, args VertexAiTensorboardArgs) *VertexAiTensorboard {
	return &VertexAiTensorboard{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*VertexAiTensorboard)(nil)

// VertexAiTensorboard represents the Terraform resource google_vertex_ai_tensorboard.
type VertexAiTensorboard struct {
	Name      string
	Args      VertexAiTensorboardArgs
	state     *vertexAiTensorboardState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [VertexAiTensorboard].
func (vat *VertexAiTensorboard) Type() string {
	return "google_vertex_ai_tensorboard"
}

// LocalName returns the local name for [VertexAiTensorboard].
func (vat *VertexAiTensorboard) LocalName() string {
	return vat.Name
}

// Configuration returns the configuration (args) for [VertexAiTensorboard].
func (vat *VertexAiTensorboard) Configuration() interface{} {
	return vat.Args
}

// DependOn is used for other resources to depend on [VertexAiTensorboard].
func (vat *VertexAiTensorboard) DependOn() terra.Reference {
	return terra.ReferenceResource(vat)
}

// Dependencies returns the list of resources [VertexAiTensorboard] depends_on.
func (vat *VertexAiTensorboard) Dependencies() terra.Dependencies {
	return vat.DependsOn
}

// LifecycleManagement returns the lifecycle block for [VertexAiTensorboard].
func (vat *VertexAiTensorboard) LifecycleManagement() *terra.Lifecycle {
	return vat.Lifecycle
}

// Attributes returns the attributes for [VertexAiTensorboard].
func (vat *VertexAiTensorboard) Attributes() vertexAiTensorboardAttributes {
	return vertexAiTensorboardAttributes{ref: terra.ReferenceResource(vat)}
}

// ImportState imports the given attribute values into [VertexAiTensorboard]'s state.
func (vat *VertexAiTensorboard) ImportState(av io.Reader) error {
	vat.state = &vertexAiTensorboardState{}
	if err := json.NewDecoder(av).Decode(vat.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", vat.Type(), vat.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [VertexAiTensorboard] has state.
func (vat *VertexAiTensorboard) State() (*vertexAiTensorboardState, bool) {
	return vat.state, vat.state != nil
}

// StateMust returns the state for [VertexAiTensorboard]. Panics if the state is nil.
func (vat *VertexAiTensorboard) StateMust() *vertexAiTensorboardState {
	if vat.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", vat.Type(), vat.LocalName()))
	}
	return vat.state
}

// VertexAiTensorboardArgs contains the configurations for google_vertex_ai_tensorboard.
type VertexAiTensorboardArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// DisplayName: string, required
	DisplayName terra.StringValue `hcl:"display_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Labels: map of string, optional
	Labels terra.MapValue[terra.StringValue] `hcl:"labels,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Region: string, optional
	Region terra.StringValue `hcl:"region,attr"`
	// EncryptionSpec: optional
	EncryptionSpec *vertexaitensorboard.EncryptionSpec `hcl:"encryption_spec,block"`
	// Timeouts: optional
	Timeouts *vertexaitensorboard.Timeouts `hcl:"timeouts,block"`
}
type vertexAiTensorboardAttributes struct {
	ref terra.Reference
}

// BlobStoragePathPrefix returns a reference to field blob_storage_path_prefix of google_vertex_ai_tensorboard.
func (vat vertexAiTensorboardAttributes) BlobStoragePathPrefix() terra.StringValue {
	return terra.ReferenceAsString(vat.ref.Append("blob_storage_path_prefix"))
}

// CreateTime returns a reference to field create_time of google_vertex_ai_tensorboard.
func (vat vertexAiTensorboardAttributes) CreateTime() terra.StringValue {
	return terra.ReferenceAsString(vat.ref.Append("create_time"))
}

// Description returns a reference to field description of google_vertex_ai_tensorboard.
func (vat vertexAiTensorboardAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(vat.ref.Append("description"))
}

// DisplayName returns a reference to field display_name of google_vertex_ai_tensorboard.
func (vat vertexAiTensorboardAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(vat.ref.Append("display_name"))
}

// Id returns a reference to field id of google_vertex_ai_tensorboard.
func (vat vertexAiTensorboardAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(vat.ref.Append("id"))
}

// Labels returns a reference to field labels of google_vertex_ai_tensorboard.
func (vat vertexAiTensorboardAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](vat.ref.Append("labels"))
}

// Name returns a reference to field name of google_vertex_ai_tensorboard.
func (vat vertexAiTensorboardAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(vat.ref.Append("name"))
}

// Project returns a reference to field project of google_vertex_ai_tensorboard.
func (vat vertexAiTensorboardAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(vat.ref.Append("project"))
}

// Region returns a reference to field region of google_vertex_ai_tensorboard.
func (vat vertexAiTensorboardAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(vat.ref.Append("region"))
}

// RunCount returns a reference to field run_count of google_vertex_ai_tensorboard.
func (vat vertexAiTensorboardAttributes) RunCount() terra.StringValue {
	return terra.ReferenceAsString(vat.ref.Append("run_count"))
}

// UpdateTime returns a reference to field update_time of google_vertex_ai_tensorboard.
func (vat vertexAiTensorboardAttributes) UpdateTime() terra.StringValue {
	return terra.ReferenceAsString(vat.ref.Append("update_time"))
}

func (vat vertexAiTensorboardAttributes) EncryptionSpec() terra.ListValue[vertexaitensorboard.EncryptionSpecAttributes] {
	return terra.ReferenceAsList[vertexaitensorboard.EncryptionSpecAttributes](vat.ref.Append("encryption_spec"))
}

func (vat vertexAiTensorboardAttributes) Timeouts() vertexaitensorboard.TimeoutsAttributes {
	return terra.ReferenceAsSingle[vertexaitensorboard.TimeoutsAttributes](vat.ref.Append("timeouts"))
}

type vertexAiTensorboardState struct {
	BlobStoragePathPrefix string                                    `json:"blob_storage_path_prefix"`
	CreateTime            string                                    `json:"create_time"`
	Description           string                                    `json:"description"`
	DisplayName           string                                    `json:"display_name"`
	Id                    string                                    `json:"id"`
	Labels                map[string]string                         `json:"labels"`
	Name                  string                                    `json:"name"`
	Project               string                                    `json:"project"`
	Region                string                                    `json:"region"`
	RunCount              string                                    `json:"run_count"`
	UpdateTime            string                                    `json:"update_time"`
	EncryptionSpec        []vertexaitensorboard.EncryptionSpecState `json:"encryption_spec"`
	Timeouts              *vertexaitensorboard.TimeoutsState        `json:"timeouts"`
}
