// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package vault

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewGenericEndpoint creates a new instance of [GenericEndpoint].
func NewGenericEndpoint(name string, args GenericEndpointArgs) *GenericEndpoint {
	return &GenericEndpoint{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*GenericEndpoint)(nil)

// GenericEndpoint represents the Terraform resource vault_generic_endpoint.
type GenericEndpoint struct {
	Name      string
	Args      GenericEndpointArgs
	state     *genericEndpointState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [GenericEndpoint].
func (ge *GenericEndpoint) Type() string {
	return "vault_generic_endpoint"
}

// LocalName returns the local name for [GenericEndpoint].
func (ge *GenericEndpoint) LocalName() string {
	return ge.Name
}

// Configuration returns the configuration (args) for [GenericEndpoint].
func (ge *GenericEndpoint) Configuration() interface{} {
	return ge.Args
}

// DependOn is used for other resources to depend on [GenericEndpoint].
func (ge *GenericEndpoint) DependOn() terra.Reference {
	return terra.ReferenceResource(ge)
}

// Dependencies returns the list of resources [GenericEndpoint] depends_on.
func (ge *GenericEndpoint) Dependencies() terra.Dependencies {
	return ge.DependsOn
}

// LifecycleManagement returns the lifecycle block for [GenericEndpoint].
func (ge *GenericEndpoint) LifecycleManagement() *terra.Lifecycle {
	return ge.Lifecycle
}

// Attributes returns the attributes for [GenericEndpoint].
func (ge *GenericEndpoint) Attributes() genericEndpointAttributes {
	return genericEndpointAttributes{ref: terra.ReferenceResource(ge)}
}

// ImportState imports the given attribute values into [GenericEndpoint]'s state.
func (ge *GenericEndpoint) ImportState(av io.Reader) error {
	ge.state = &genericEndpointState{}
	if err := json.NewDecoder(av).Decode(ge.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ge.Type(), ge.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [GenericEndpoint] has state.
func (ge *GenericEndpoint) State() (*genericEndpointState, bool) {
	return ge.state, ge.state != nil
}

// StateMust returns the state for [GenericEndpoint]. Panics if the state is nil.
func (ge *GenericEndpoint) StateMust() *genericEndpointState {
	if ge.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ge.Type(), ge.LocalName()))
	}
	return ge.state
}

// GenericEndpointArgs contains the configurations for vault_generic_endpoint.
type GenericEndpointArgs struct {
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
type genericEndpointAttributes struct {
	ref terra.Reference
}

// DataJson returns a reference to field data_json of vault_generic_endpoint.
func (ge genericEndpointAttributes) DataJson() terra.StringValue {
	return terra.ReferenceAsString(ge.ref.Append("data_json"))
}

// DisableDelete returns a reference to field disable_delete of vault_generic_endpoint.
func (ge genericEndpointAttributes) DisableDelete() terra.BoolValue {
	return terra.ReferenceAsBool(ge.ref.Append("disable_delete"))
}

// DisableRead returns a reference to field disable_read of vault_generic_endpoint.
func (ge genericEndpointAttributes) DisableRead() terra.BoolValue {
	return terra.ReferenceAsBool(ge.ref.Append("disable_read"))
}

// Id returns a reference to field id of vault_generic_endpoint.
func (ge genericEndpointAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ge.ref.Append("id"))
}

// IgnoreAbsentFields returns a reference to field ignore_absent_fields of vault_generic_endpoint.
func (ge genericEndpointAttributes) IgnoreAbsentFields() terra.BoolValue {
	return terra.ReferenceAsBool(ge.ref.Append("ignore_absent_fields"))
}

// Namespace returns a reference to field namespace of vault_generic_endpoint.
func (ge genericEndpointAttributes) Namespace() terra.StringValue {
	return terra.ReferenceAsString(ge.ref.Append("namespace"))
}

// Path returns a reference to field path of vault_generic_endpoint.
func (ge genericEndpointAttributes) Path() terra.StringValue {
	return terra.ReferenceAsString(ge.ref.Append("path"))
}

// WriteData returns a reference to field write_data of vault_generic_endpoint.
func (ge genericEndpointAttributes) WriteData() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ge.ref.Append("write_data"))
}

// WriteDataJson returns a reference to field write_data_json of vault_generic_endpoint.
func (ge genericEndpointAttributes) WriteDataJson() terra.StringValue {
	return terra.ReferenceAsString(ge.ref.Append("write_data_json"))
}

// WriteFields returns a reference to field write_fields of vault_generic_endpoint.
func (ge genericEndpointAttributes) WriteFields() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](ge.ref.Append("write_fields"))
}

type genericEndpointState struct {
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