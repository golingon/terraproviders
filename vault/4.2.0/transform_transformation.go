// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package vault

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// NewTransformTransformation creates a new instance of [TransformTransformation].
func NewTransformTransformation(name string, args TransformTransformationArgs) *TransformTransformation {
	return &TransformTransformation{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*TransformTransformation)(nil)

// TransformTransformation represents the Terraform resource vault_transform_transformation.
type TransformTransformation struct {
	Name      string
	Args      TransformTransformationArgs
	state     *transformTransformationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [TransformTransformation].
func (tt *TransformTransformation) Type() string {
	return "vault_transform_transformation"
}

// LocalName returns the local name for [TransformTransformation].
func (tt *TransformTransformation) LocalName() string {
	return tt.Name
}

// Configuration returns the configuration (args) for [TransformTransformation].
func (tt *TransformTransformation) Configuration() interface{} {
	return tt.Args
}

// DependOn is used for other resources to depend on [TransformTransformation].
func (tt *TransformTransformation) DependOn() terra.Reference {
	return terra.ReferenceResource(tt)
}

// Dependencies returns the list of resources [TransformTransformation] depends_on.
func (tt *TransformTransformation) Dependencies() terra.Dependencies {
	return tt.DependsOn
}

// LifecycleManagement returns the lifecycle block for [TransformTransformation].
func (tt *TransformTransformation) LifecycleManagement() *terra.Lifecycle {
	return tt.Lifecycle
}

// Attributes returns the attributes for [TransformTransformation].
func (tt *TransformTransformation) Attributes() transformTransformationAttributes {
	return transformTransformationAttributes{ref: terra.ReferenceResource(tt)}
}

// ImportState imports the given attribute values into [TransformTransformation]'s state.
func (tt *TransformTransformation) ImportState(av io.Reader) error {
	tt.state = &transformTransformationState{}
	if err := json.NewDecoder(av).Decode(tt.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", tt.Type(), tt.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [TransformTransformation] has state.
func (tt *TransformTransformation) State() (*transformTransformationState, bool) {
	return tt.state, tt.state != nil
}

// StateMust returns the state for [TransformTransformation]. Panics if the state is nil.
func (tt *TransformTransformation) StateMust() *transformTransformationState {
	if tt.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", tt.Type(), tt.LocalName()))
	}
	return tt.state
}

// TransformTransformationArgs contains the configurations for vault_transform_transformation.
type TransformTransformationArgs struct {
	// AllowedRoles: list of string, optional
	AllowedRoles terra.ListValue[terra.StringValue] `hcl:"allowed_roles,attr"`
	// DeletionAllowed: bool, optional
	DeletionAllowed terra.BoolValue `hcl:"deletion_allowed,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// MaskingCharacter: string, optional
	MaskingCharacter terra.StringValue `hcl:"masking_character,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Namespace: string, optional
	Namespace terra.StringValue `hcl:"namespace,attr"`
	// Path: string, required
	Path terra.StringValue `hcl:"path,attr" validate:"required"`
	// Template: string, optional
	Template terra.StringValue `hcl:"template,attr"`
	// Templates: list of string, optional
	Templates terra.ListValue[terra.StringValue] `hcl:"templates,attr"`
	// TweakSource: string, optional
	TweakSource terra.StringValue `hcl:"tweak_source,attr"`
	// Type: string, optional
	Type terra.StringValue `hcl:"type,attr"`
}
type transformTransformationAttributes struct {
	ref terra.Reference
}

// AllowedRoles returns a reference to field allowed_roles of vault_transform_transformation.
func (tt transformTransformationAttributes) AllowedRoles() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](tt.ref.Append("allowed_roles"))
}

// DeletionAllowed returns a reference to field deletion_allowed of vault_transform_transformation.
func (tt transformTransformationAttributes) DeletionAllowed() terra.BoolValue {
	return terra.ReferenceAsBool(tt.ref.Append("deletion_allowed"))
}

// Id returns a reference to field id of vault_transform_transformation.
func (tt transformTransformationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(tt.ref.Append("id"))
}

// MaskingCharacter returns a reference to field masking_character of vault_transform_transformation.
func (tt transformTransformationAttributes) MaskingCharacter() terra.StringValue {
	return terra.ReferenceAsString(tt.ref.Append("masking_character"))
}

// Name returns a reference to field name of vault_transform_transformation.
func (tt transformTransformationAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(tt.ref.Append("name"))
}

// Namespace returns a reference to field namespace of vault_transform_transformation.
func (tt transformTransformationAttributes) Namespace() terra.StringValue {
	return terra.ReferenceAsString(tt.ref.Append("namespace"))
}

// Path returns a reference to field path of vault_transform_transformation.
func (tt transformTransformationAttributes) Path() terra.StringValue {
	return terra.ReferenceAsString(tt.ref.Append("path"))
}

// Template returns a reference to field template of vault_transform_transformation.
func (tt transformTransformationAttributes) Template() terra.StringValue {
	return terra.ReferenceAsString(tt.ref.Append("template"))
}

// Templates returns a reference to field templates of vault_transform_transformation.
func (tt transformTransformationAttributes) Templates() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](tt.ref.Append("templates"))
}

// TweakSource returns a reference to field tweak_source of vault_transform_transformation.
func (tt transformTransformationAttributes) TweakSource() terra.StringValue {
	return terra.ReferenceAsString(tt.ref.Append("tweak_source"))
}

// Type returns a reference to field type of vault_transform_transformation.
func (tt transformTransformationAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(tt.ref.Append("type"))
}

type transformTransformationState struct {
	AllowedRoles     []string `json:"allowed_roles"`
	DeletionAllowed  bool     `json:"deletion_allowed"`
	Id               string   `json:"id"`
	MaskingCharacter string   `json:"masking_character"`
	Name             string   `json:"name"`
	Namespace        string   `json:"namespace"`
	Path             string   `json:"path"`
	Template         string   `json:"template"`
	Templates        []string `json:"templates"`
	TweakSource      string   `json:"tweak_source"`
	Type             string   `json:"type"`
}
