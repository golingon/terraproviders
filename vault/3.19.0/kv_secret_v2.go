// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package vault

import (
	"encoding/json"
	"fmt"
	kvsecretv2 "github.com/golingon/terraproviders/vault/3.19.0/kvsecretv2"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewKvSecretV2 creates a new instance of [KvSecretV2].
func NewKvSecretV2(name string, args KvSecretV2Args) *KvSecretV2 {
	return &KvSecretV2{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*KvSecretV2)(nil)

// KvSecretV2 represents the Terraform resource vault_kv_secret_v2.
type KvSecretV2 struct {
	Name      string
	Args      KvSecretV2Args
	state     *kvSecretV2State
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [KvSecretV2].
func (ksv *KvSecretV2) Type() string {
	return "vault_kv_secret_v2"
}

// LocalName returns the local name for [KvSecretV2].
func (ksv *KvSecretV2) LocalName() string {
	return ksv.Name
}

// Configuration returns the configuration (args) for [KvSecretV2].
func (ksv *KvSecretV2) Configuration() interface{} {
	return ksv.Args
}

// DependOn is used for other resources to depend on [KvSecretV2].
func (ksv *KvSecretV2) DependOn() terra.Reference {
	return terra.ReferenceResource(ksv)
}

// Dependencies returns the list of resources [KvSecretV2] depends_on.
func (ksv *KvSecretV2) Dependencies() terra.Dependencies {
	return ksv.DependsOn
}

// LifecycleManagement returns the lifecycle block for [KvSecretV2].
func (ksv *KvSecretV2) LifecycleManagement() *terra.Lifecycle {
	return ksv.Lifecycle
}

// Attributes returns the attributes for [KvSecretV2].
func (ksv *KvSecretV2) Attributes() kvSecretV2Attributes {
	return kvSecretV2Attributes{ref: terra.ReferenceResource(ksv)}
}

// ImportState imports the given attribute values into [KvSecretV2]'s state.
func (ksv *KvSecretV2) ImportState(av io.Reader) error {
	ksv.state = &kvSecretV2State{}
	if err := json.NewDecoder(av).Decode(ksv.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ksv.Type(), ksv.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [KvSecretV2] has state.
func (ksv *KvSecretV2) State() (*kvSecretV2State, bool) {
	return ksv.state, ksv.state != nil
}

// StateMust returns the state for [KvSecretV2]. Panics if the state is nil.
func (ksv *KvSecretV2) StateMust() *kvSecretV2State {
	if ksv.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ksv.Type(), ksv.LocalName()))
	}
	return ksv.state
}

// KvSecretV2Args contains the configurations for vault_kv_secret_v2.
type KvSecretV2Args struct {
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
	CustomMetadata *kvsecretv2.CustomMetadata `hcl:"custom_metadata,block"`
}
type kvSecretV2Attributes struct {
	ref terra.Reference
}

// Cas returns a reference to field cas of vault_kv_secret_v2.
func (ksv kvSecretV2Attributes) Cas() terra.NumberValue {
	return terra.ReferenceAsNumber(ksv.ref.Append("cas"))
}

// Data returns a reference to field data of vault_kv_secret_v2.
func (ksv kvSecretV2Attributes) Data() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ksv.ref.Append("data"))
}

// DataJson returns a reference to field data_json of vault_kv_secret_v2.
func (ksv kvSecretV2Attributes) DataJson() terra.StringValue {
	return terra.ReferenceAsString(ksv.ref.Append("data_json"))
}

// DeleteAllVersions returns a reference to field delete_all_versions of vault_kv_secret_v2.
func (ksv kvSecretV2Attributes) DeleteAllVersions() terra.BoolValue {
	return terra.ReferenceAsBool(ksv.ref.Append("delete_all_versions"))
}

// DisableRead returns a reference to field disable_read of vault_kv_secret_v2.
func (ksv kvSecretV2Attributes) DisableRead() terra.BoolValue {
	return terra.ReferenceAsBool(ksv.ref.Append("disable_read"))
}

// Id returns a reference to field id of vault_kv_secret_v2.
func (ksv kvSecretV2Attributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ksv.ref.Append("id"))
}

// Metadata returns a reference to field metadata of vault_kv_secret_v2.
func (ksv kvSecretV2Attributes) Metadata() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ksv.ref.Append("metadata"))
}

// Mount returns a reference to field mount of vault_kv_secret_v2.
func (ksv kvSecretV2Attributes) Mount() terra.StringValue {
	return terra.ReferenceAsString(ksv.ref.Append("mount"))
}

// Name returns a reference to field name of vault_kv_secret_v2.
func (ksv kvSecretV2Attributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ksv.ref.Append("name"))
}

// Namespace returns a reference to field namespace of vault_kv_secret_v2.
func (ksv kvSecretV2Attributes) Namespace() terra.StringValue {
	return terra.ReferenceAsString(ksv.ref.Append("namespace"))
}

// Options returns a reference to field options of vault_kv_secret_v2.
func (ksv kvSecretV2Attributes) Options() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ksv.ref.Append("options"))
}

// Path returns a reference to field path of vault_kv_secret_v2.
func (ksv kvSecretV2Attributes) Path() terra.StringValue {
	return terra.ReferenceAsString(ksv.ref.Append("path"))
}

func (ksv kvSecretV2Attributes) CustomMetadata() terra.ListValue[kvsecretv2.CustomMetadataAttributes] {
	return terra.ReferenceAsList[kvsecretv2.CustomMetadataAttributes](ksv.ref.Append("custom_metadata"))
}

type kvSecretV2State struct {
	Cas               float64                          `json:"cas"`
	Data              map[string]string                `json:"data"`
	DataJson          string                           `json:"data_json"`
	DeleteAllVersions bool                             `json:"delete_all_versions"`
	DisableRead       bool                             `json:"disable_read"`
	Id                string                           `json:"id"`
	Metadata          map[string]string                `json:"metadata"`
	Mount             string                           `json:"mount"`
	Name              string                           `json:"name"`
	Namespace         string                           `json:"namespace"`
	Options           map[string]string                `json:"options"`
	Path              string                           `json:"path"`
	CustomMetadata    []kvsecretv2.CustomMetadataState `json:"custom_metadata"`
}
