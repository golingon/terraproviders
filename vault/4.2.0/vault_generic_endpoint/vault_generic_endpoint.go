// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package vault_generic_endpoint

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// New creates a new instance of [Resource].
func New(name string, args Args) *Resource {
	return &Resource{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Resource)(nil)

// Resource represents the Terraform resource vault_generic_endpoint.
type Resource struct {
	Name      string
	Args      Args
	state     *vaultGenericEndpointState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (vge *Resource) Type() string {
	return "vault_generic_endpoint"
}

// LocalName returns the local name for [Resource].
func (vge *Resource) LocalName() string {
	return vge.Name
}

// Configuration returns the configuration (args) for [Resource].
func (vge *Resource) Configuration() interface{} {
	return vge.Args
}

// DependOn is used for other resources to depend on [Resource].
func (vge *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(vge)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (vge *Resource) Dependencies() terra.Dependencies {
	return vge.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (vge *Resource) LifecycleManagement() *terra.Lifecycle {
	return vge.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (vge *Resource) Attributes() vaultGenericEndpointAttributes {
	return vaultGenericEndpointAttributes{ref: terra.ReferenceResource(vge)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (vge *Resource) ImportState(state io.Reader) error {
	vge.state = &vaultGenericEndpointState{}
	if err := json.NewDecoder(state).Decode(vge.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", vge.Type(), vge.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (vge *Resource) State() (*vaultGenericEndpointState, bool) {
	return vge.state, vge.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (vge *Resource) StateMust() *vaultGenericEndpointState {
	if vge.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", vge.Type(), vge.LocalName()))
	}
	return vge.state
}

// Args contains the configurations for vault_generic_endpoint.
type Args struct {
	// DataJson: string, required
	DataJson terra.StringValue `hcl:"data_json,attr" validate:"required"`
	// DisableDelete: bool, optional
	DisableDelete terra.BoolValue `hcl:"disable_delete,attr"`
	// DisableRead: bool, optional
	DisableRead terra.BoolValue `hcl:"disable_read,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IgnoreAbsentFields: bool, optional
	IgnoreAbsentFields terra.BoolValue `hcl:"ignore_absent_fields,attr"`
	// Namespace: string, optional
	Namespace terra.StringValue `hcl:"namespace,attr"`
	// Path: string, required
	Path terra.StringValue `hcl:"path,attr" validate:"required"`
	// WriteFields: list of string, optional
	WriteFields terra.ListValue[terra.StringValue] `hcl:"write_fields,attr"`
}

type vaultGenericEndpointAttributes struct {
	ref terra.Reference
}

// DataJson returns a reference to field data_json of vault_generic_endpoint.
func (vge vaultGenericEndpointAttributes) DataJson() terra.StringValue {
	return terra.ReferenceAsString(vge.ref.Append("data_json"))
}

// DisableDelete returns a reference to field disable_delete of vault_generic_endpoint.
func (vge vaultGenericEndpointAttributes) DisableDelete() terra.BoolValue {
	return terra.ReferenceAsBool(vge.ref.Append("disable_delete"))
}

// DisableRead returns a reference to field disable_read of vault_generic_endpoint.
func (vge vaultGenericEndpointAttributes) DisableRead() terra.BoolValue {
	return terra.ReferenceAsBool(vge.ref.Append("disable_read"))
}

// Id returns a reference to field id of vault_generic_endpoint.
func (vge vaultGenericEndpointAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(vge.ref.Append("id"))
}

// IgnoreAbsentFields returns a reference to field ignore_absent_fields of vault_generic_endpoint.
func (vge vaultGenericEndpointAttributes) IgnoreAbsentFields() terra.BoolValue {
	return terra.ReferenceAsBool(vge.ref.Append("ignore_absent_fields"))
}

// Namespace returns a reference to field namespace of vault_generic_endpoint.
func (vge vaultGenericEndpointAttributes) Namespace() terra.StringValue {
	return terra.ReferenceAsString(vge.ref.Append("namespace"))
}

// Path returns a reference to field path of vault_generic_endpoint.
func (vge vaultGenericEndpointAttributes) Path() terra.StringValue {
	return terra.ReferenceAsString(vge.ref.Append("path"))
}

// WriteData returns a reference to field write_data of vault_generic_endpoint.
func (vge vaultGenericEndpointAttributes) WriteData() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](vge.ref.Append("write_data"))
}

// WriteDataJson returns a reference to field write_data_json of vault_generic_endpoint.
func (vge vaultGenericEndpointAttributes) WriteDataJson() terra.StringValue {
	return terra.ReferenceAsString(vge.ref.Append("write_data_json"))
}

// WriteFields returns a reference to field write_fields of vault_generic_endpoint.
func (vge vaultGenericEndpointAttributes) WriteFields() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](vge.ref.Append("write_fields"))
}

type vaultGenericEndpointState struct {
	DataJson           string            `json:"data_json"`
	DisableDelete      bool              `json:"disable_delete"`
	DisableRead        bool              `json:"disable_read"`
	Id                 string            `json:"id"`
	IgnoreAbsentFields bool              `json:"ignore_absent_fields"`
	Namespace          string            `json:"namespace"`
	Path               string            `json:"path"`
	WriteData          map[string]string `json:"write_data"`
	WriteDataJson      string            `json:"write_data_json"`
	WriteFields        []string          `json:"write_fields"`
}
