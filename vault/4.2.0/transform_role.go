// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package vault

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// NewTransformRole creates a new instance of [TransformRole].
func NewTransformRole(name string, args TransformRoleArgs) *TransformRole {
	return &TransformRole{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*TransformRole)(nil)

// TransformRole represents the Terraform resource vault_transform_role.
type TransformRole struct {
	Name      string
	Args      TransformRoleArgs
	state     *transformRoleState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [TransformRole].
func (tr *TransformRole) Type() string {
	return "vault_transform_role"
}

// LocalName returns the local name for [TransformRole].
func (tr *TransformRole) LocalName() string {
	return tr.Name
}

// Configuration returns the configuration (args) for [TransformRole].
func (tr *TransformRole) Configuration() interface{} {
	return tr.Args
}

// DependOn is used for other resources to depend on [TransformRole].
func (tr *TransformRole) DependOn() terra.Reference {
	return terra.ReferenceResource(tr)
}

// Dependencies returns the list of resources [TransformRole] depends_on.
func (tr *TransformRole) Dependencies() terra.Dependencies {
	return tr.DependsOn
}

// LifecycleManagement returns the lifecycle block for [TransformRole].
func (tr *TransformRole) LifecycleManagement() *terra.Lifecycle {
	return tr.Lifecycle
}

// Attributes returns the attributes for [TransformRole].
func (tr *TransformRole) Attributes() transformRoleAttributes {
	return transformRoleAttributes{ref: terra.ReferenceResource(tr)}
}

// ImportState imports the given attribute values into [TransformRole]'s state.
func (tr *TransformRole) ImportState(av io.Reader) error {
	tr.state = &transformRoleState{}
	if err := json.NewDecoder(av).Decode(tr.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", tr.Type(), tr.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [TransformRole] has state.
func (tr *TransformRole) State() (*transformRoleState, bool) {
	return tr.state, tr.state != nil
}

// StateMust returns the state for [TransformRole]. Panics if the state is nil.
func (tr *TransformRole) StateMust() *transformRoleState {
	if tr.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", tr.Type(), tr.LocalName()))
	}
	return tr.state
}

// TransformRoleArgs contains the configurations for vault_transform_role.
type TransformRoleArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Namespace: string, optional
	Namespace terra.StringValue `hcl:"namespace,attr"`
	// Path: string, required
	Path terra.StringValue `hcl:"path,attr" validate:"required"`
	// Transformations: list of string, optional
	Transformations terra.ListValue[terra.StringValue] `hcl:"transformations,attr"`
}
type transformRoleAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of vault_transform_role.
func (tr transformRoleAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(tr.ref.Append("id"))
}

// Name returns a reference to field name of vault_transform_role.
func (tr transformRoleAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(tr.ref.Append("name"))
}

// Namespace returns a reference to field namespace of vault_transform_role.
func (tr transformRoleAttributes) Namespace() terra.StringValue {
	return terra.ReferenceAsString(tr.ref.Append("namespace"))
}

// Path returns a reference to field path of vault_transform_role.
func (tr transformRoleAttributes) Path() terra.StringValue {
	return terra.ReferenceAsString(tr.ref.Append("path"))
}

// Transformations returns a reference to field transformations of vault_transform_role.
func (tr transformRoleAttributes) Transformations() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](tr.ref.Append("transformations"))
}

type transformRoleState struct {
	Id              string   `json:"id"`
	Name            string   `json:"name"`
	Namespace       string   `json:"namespace"`
	Path            string   `json:"path"`
	Transformations []string `json:"transformations"`
}
