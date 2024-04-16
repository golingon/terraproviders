// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package vault_kv_secret_v2

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

// Resource represents the Terraform resource vault_kv_secret_v2.
type Resource struct {
	Name      string
	Args      Args
	state     *vaultKvSecretV2State
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (vksv *Resource) Type() string {
	return "vault_kv_secret_v2"
}

// LocalName returns the local name for [Resource].
func (vksv *Resource) LocalName() string {
	return vksv.Name
}

// Configuration returns the configuration (args) for [Resource].
func (vksv *Resource) Configuration() interface{} {
	return vksv.Args
}

// DependOn is used for other resources to depend on [Resource].
func (vksv *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(vksv)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (vksv *Resource) Dependencies() terra.Dependencies {
	return vksv.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (vksv *Resource) LifecycleManagement() *terra.Lifecycle {
	return vksv.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (vksv *Resource) Attributes() vaultKvSecretV2Attributes {
	return vaultKvSecretV2Attributes{ref: terra.ReferenceResource(vksv)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (vksv *Resource) ImportState(state io.Reader) error {
	vksv.state = &vaultKvSecretV2State{}
	if err := json.NewDecoder(state).Decode(vksv.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", vksv.Type(), vksv.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (vksv *Resource) State() (*vaultKvSecretV2State, bool) {
	return vksv.state, vksv.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (vksv *Resource) StateMust() *vaultKvSecretV2State {
	if vksv.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", vksv.Type(), vksv.LocalName()))
	}
	return vksv.state
}

// Args contains the configurations for vault_kv_secret_v2.
type Args struct {
	// Cas: number, optional
	Cas terra.NumberValue `hcl:"cas,attr"`
	// DataJson: string, required
	DataJson terra.StringValue `hcl:"data_json,attr" validate:"required"`
	// DeleteAllVersions: bool, optional
	DeleteAllVersions terra.BoolValue `hcl:"delete_all_versions,attr"`
	// DisableRead: bool, optional
	DisableRead terra.BoolValue `hcl:"disable_read,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Mount: string, required
	Mount terra.StringValue `hcl:"mount,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Namespace: string, optional
	Namespace terra.StringValue `hcl:"namespace,attr"`
	// Options: map of string, optional
	Options terra.MapValue[terra.StringValue] `hcl:"options,attr"`
	// CustomMetadata: optional
	CustomMetadata *CustomMetadata `hcl:"custom_metadata,block"`
}

type vaultKvSecretV2Attributes struct {
	ref terra.Reference
}

// Cas returns a reference to field cas of vault_kv_secret_v2.
func (vksv vaultKvSecretV2Attributes) Cas() terra.NumberValue {
	return terra.ReferenceAsNumber(vksv.ref.Append("cas"))
}

// Data returns a reference to field data of vault_kv_secret_v2.
func (vksv vaultKvSecretV2Attributes) Data() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](vksv.ref.Append("data"))
}

// DataJson returns a reference to field data_json of vault_kv_secret_v2.
func (vksv vaultKvSecretV2Attributes) DataJson() terra.StringValue {
	return terra.ReferenceAsString(vksv.ref.Append("data_json"))
}

// DeleteAllVersions returns a reference to field delete_all_versions of vault_kv_secret_v2.
func (vksv vaultKvSecretV2Attributes) DeleteAllVersions() terra.BoolValue {
	return terra.ReferenceAsBool(vksv.ref.Append("delete_all_versions"))
}

// DisableRead returns a reference to field disable_read of vault_kv_secret_v2.
func (vksv vaultKvSecretV2Attributes) DisableRead() terra.BoolValue {
	return terra.ReferenceAsBool(vksv.ref.Append("disable_read"))
}

// Id returns a reference to field id of vault_kv_secret_v2.
func (vksv vaultKvSecretV2Attributes) Id() terra.StringValue {
	return terra.ReferenceAsString(vksv.ref.Append("id"))
}

// Metadata returns a reference to field metadata of vault_kv_secret_v2.
func (vksv vaultKvSecretV2Attributes) Metadata() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](vksv.ref.Append("metadata"))
}

// Mount returns a reference to field mount of vault_kv_secret_v2.
func (vksv vaultKvSecretV2Attributes) Mount() terra.StringValue {
	return terra.ReferenceAsString(vksv.ref.Append("mount"))
}

// Name returns a reference to field name of vault_kv_secret_v2.
func (vksv vaultKvSecretV2Attributes) Name() terra.StringValue {
	return terra.ReferenceAsString(vksv.ref.Append("name"))
}

// Namespace returns a reference to field namespace of vault_kv_secret_v2.
func (vksv vaultKvSecretV2Attributes) Namespace() terra.StringValue {
	return terra.ReferenceAsString(vksv.ref.Append("namespace"))
}

// Options returns a reference to field options of vault_kv_secret_v2.
func (vksv vaultKvSecretV2Attributes) Options() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](vksv.ref.Append("options"))
}

// Path returns a reference to field path of vault_kv_secret_v2.
func (vksv vaultKvSecretV2Attributes) Path() terra.StringValue {
	return terra.ReferenceAsString(vksv.ref.Append("path"))
}

func (vksv vaultKvSecretV2Attributes) CustomMetadata() terra.ListValue[CustomMetadataAttributes] {
	return terra.ReferenceAsList[CustomMetadataAttributes](vksv.ref.Append("custom_metadata"))
}

type vaultKvSecretV2State struct {
	Cas               float64               `json:"cas"`
	Data              map[string]string     `json:"data"`
	DataJson          string                `json:"data_json"`
	DeleteAllVersions bool                  `json:"delete_all_versions"`
	DisableRead       bool                  `json:"disable_read"`
	Id                string                `json:"id"`
	Metadata          map[string]string     `json:"metadata"`
	Mount             string                `json:"mount"`
	Name              string                `json:"name"`
	Namespace         string                `json:"namespace"`
	Options           map[string]string     `json:"options"`
	Path              string                `json:"path"`
	CustomMetadata    []CustomMetadataState `json:"custom_metadata"`
}
