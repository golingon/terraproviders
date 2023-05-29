// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package vault

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewTransitSecretBackendKey creates a new instance of [TransitSecretBackendKey].
func NewTransitSecretBackendKey(name string, args TransitSecretBackendKeyArgs) *TransitSecretBackendKey {
	return &TransitSecretBackendKey{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*TransitSecretBackendKey)(nil)

// TransitSecretBackendKey represents the Terraform resource vault_transit_secret_backend_key.
type TransitSecretBackendKey struct {
	Name      string
	Args      TransitSecretBackendKeyArgs
	state     *transitSecretBackendKeyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [TransitSecretBackendKey].
func (tsbk *TransitSecretBackendKey) Type() string {
	return "vault_transit_secret_backend_key"
}

// LocalName returns the local name for [TransitSecretBackendKey].
func (tsbk *TransitSecretBackendKey) LocalName() string {
	return tsbk.Name
}

// Configuration returns the configuration (args) for [TransitSecretBackendKey].
func (tsbk *TransitSecretBackendKey) Configuration() interface{} {
	return tsbk.Args
}

// DependOn is used for other resources to depend on [TransitSecretBackendKey].
func (tsbk *TransitSecretBackendKey) DependOn() terra.Reference {
	return terra.ReferenceResource(tsbk)
}

// Dependencies returns the list of resources [TransitSecretBackendKey] depends_on.
func (tsbk *TransitSecretBackendKey) Dependencies() terra.Dependencies {
	return tsbk.DependsOn
}

// LifecycleManagement returns the lifecycle block for [TransitSecretBackendKey].
func (tsbk *TransitSecretBackendKey) LifecycleManagement() *terra.Lifecycle {
	return tsbk.Lifecycle
}

// Attributes returns the attributes for [TransitSecretBackendKey].
func (tsbk *TransitSecretBackendKey) Attributes() transitSecretBackendKeyAttributes {
	return transitSecretBackendKeyAttributes{ref: terra.ReferenceResource(tsbk)}
}

// ImportState imports the given attribute values into [TransitSecretBackendKey]'s state.
func (tsbk *TransitSecretBackendKey) ImportState(av io.Reader) error {
	tsbk.state = &transitSecretBackendKeyState{}
	if err := json.NewDecoder(av).Decode(tsbk.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", tsbk.Type(), tsbk.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [TransitSecretBackendKey] has state.
func (tsbk *TransitSecretBackendKey) State() (*transitSecretBackendKeyState, bool) {
	return tsbk.state, tsbk.state != nil
}

// StateMust returns the state for [TransitSecretBackendKey]. Panics if the state is nil.
func (tsbk *TransitSecretBackendKey) StateMust() *transitSecretBackendKeyState {
	if tsbk.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", tsbk.Type(), tsbk.LocalName()))
	}
	return tsbk.state
}

// TransitSecretBackendKeyArgs contains the configurations for vault_transit_secret_backend_key.
type TransitSecretBackendKeyArgs struct {
	// AllowPlaintextBackup: bool, optional
	AllowPlaintextBackup terra.BoolValue `hcl:"allow_plaintext_backup,attr"`
	// AutoRotateInterval: number, optional
	AutoRotateInterval terra.NumberValue `hcl:"auto_rotate_interval,attr"`
	// AutoRotatePeriod: number, optional
	AutoRotatePeriod terra.NumberValue `hcl:"auto_rotate_period,attr"`
	// Backend: string, required
	Backend terra.StringValue `hcl:"backend,attr" validate:"required"`
	// ConvergentEncryption: bool, optional
	ConvergentEncryption terra.BoolValue `hcl:"convergent_encryption,attr"`
	// DeletionAllowed: bool, optional
	DeletionAllowed terra.BoolValue `hcl:"deletion_allowed,attr"`
	// Derived: bool, optional
	Derived terra.BoolValue `hcl:"derived,attr"`
	// Exportable: bool, optional
	Exportable terra.BoolValue `hcl:"exportable,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// MinDecryptionVersion: number, optional
	MinDecryptionVersion terra.NumberValue `hcl:"min_decryption_version,attr"`
	// MinEncryptionVersion: number, optional
	MinEncryptionVersion terra.NumberValue `hcl:"min_encryption_version,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Namespace: string, optional
	Namespace terra.StringValue `hcl:"namespace,attr"`
	// Type: string, optional
	Type terra.StringValue `hcl:"type,attr"`
}
type transitSecretBackendKeyAttributes struct {
	ref terra.Reference
}

// AllowPlaintextBackup returns a reference to field allow_plaintext_backup of vault_transit_secret_backend_key.
func (tsbk transitSecretBackendKeyAttributes) AllowPlaintextBackup() terra.BoolValue {
	return terra.ReferenceAsBool(tsbk.ref.Append("allow_plaintext_backup"))
}

// AutoRotateInterval returns a reference to field auto_rotate_interval of vault_transit_secret_backend_key.
func (tsbk transitSecretBackendKeyAttributes) AutoRotateInterval() terra.NumberValue {
	return terra.ReferenceAsNumber(tsbk.ref.Append("auto_rotate_interval"))
}

// AutoRotatePeriod returns a reference to field auto_rotate_period of vault_transit_secret_backend_key.
func (tsbk transitSecretBackendKeyAttributes) AutoRotatePeriod() terra.NumberValue {
	return terra.ReferenceAsNumber(tsbk.ref.Append("auto_rotate_period"))
}

// Backend returns a reference to field backend of vault_transit_secret_backend_key.
func (tsbk transitSecretBackendKeyAttributes) Backend() terra.StringValue {
	return terra.ReferenceAsString(tsbk.ref.Append("backend"))
}

// ConvergentEncryption returns a reference to field convergent_encryption of vault_transit_secret_backend_key.
func (tsbk transitSecretBackendKeyAttributes) ConvergentEncryption() terra.BoolValue {
	return terra.ReferenceAsBool(tsbk.ref.Append("convergent_encryption"))
}

// DeletionAllowed returns a reference to field deletion_allowed of vault_transit_secret_backend_key.
func (tsbk transitSecretBackendKeyAttributes) DeletionAllowed() terra.BoolValue {
	return terra.ReferenceAsBool(tsbk.ref.Append("deletion_allowed"))
}

// Derived returns a reference to field derived of vault_transit_secret_backend_key.
func (tsbk transitSecretBackendKeyAttributes) Derived() terra.BoolValue {
	return terra.ReferenceAsBool(tsbk.ref.Append("derived"))
}

// Exportable returns a reference to field exportable of vault_transit_secret_backend_key.
func (tsbk transitSecretBackendKeyAttributes) Exportable() terra.BoolValue {
	return terra.ReferenceAsBool(tsbk.ref.Append("exportable"))
}

// Id returns a reference to field id of vault_transit_secret_backend_key.
func (tsbk transitSecretBackendKeyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(tsbk.ref.Append("id"))
}

// Keys returns a reference to field keys of vault_transit_secret_backend_key.
func (tsbk transitSecretBackendKeyAttributes) Keys() terra.ListValue[terra.MapValue[terra.StringValue]] {
	return terra.ReferenceAsList[terra.MapValue[terra.StringValue]](tsbk.ref.Append("keys"))
}

// LatestVersion returns a reference to field latest_version of vault_transit_secret_backend_key.
func (tsbk transitSecretBackendKeyAttributes) LatestVersion() terra.NumberValue {
	return terra.ReferenceAsNumber(tsbk.ref.Append("latest_version"))
}

// MinAvailableVersion returns a reference to field min_available_version of vault_transit_secret_backend_key.
func (tsbk transitSecretBackendKeyAttributes) MinAvailableVersion() terra.NumberValue {
	return terra.ReferenceAsNumber(tsbk.ref.Append("min_available_version"))
}

// MinDecryptionVersion returns a reference to field min_decryption_version of vault_transit_secret_backend_key.
func (tsbk transitSecretBackendKeyAttributes) MinDecryptionVersion() terra.NumberValue {
	return terra.ReferenceAsNumber(tsbk.ref.Append("min_decryption_version"))
}

// MinEncryptionVersion returns a reference to field min_encryption_version of vault_transit_secret_backend_key.
func (tsbk transitSecretBackendKeyAttributes) MinEncryptionVersion() terra.NumberValue {
	return terra.ReferenceAsNumber(tsbk.ref.Append("min_encryption_version"))
}

// Name returns a reference to field name of vault_transit_secret_backend_key.
func (tsbk transitSecretBackendKeyAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(tsbk.ref.Append("name"))
}

// Namespace returns a reference to field namespace of vault_transit_secret_backend_key.
func (tsbk transitSecretBackendKeyAttributes) Namespace() terra.StringValue {
	return terra.ReferenceAsString(tsbk.ref.Append("namespace"))
}

// SupportsDecryption returns a reference to field supports_decryption of vault_transit_secret_backend_key.
func (tsbk transitSecretBackendKeyAttributes) SupportsDecryption() terra.BoolValue {
	return terra.ReferenceAsBool(tsbk.ref.Append("supports_decryption"))
}

// SupportsDerivation returns a reference to field supports_derivation of vault_transit_secret_backend_key.
func (tsbk transitSecretBackendKeyAttributes) SupportsDerivation() terra.BoolValue {
	return terra.ReferenceAsBool(tsbk.ref.Append("supports_derivation"))
}

// SupportsEncryption returns a reference to field supports_encryption of vault_transit_secret_backend_key.
func (tsbk transitSecretBackendKeyAttributes) SupportsEncryption() terra.BoolValue {
	return terra.ReferenceAsBool(tsbk.ref.Append("supports_encryption"))
}

// SupportsSigning returns a reference to field supports_signing of vault_transit_secret_backend_key.
func (tsbk transitSecretBackendKeyAttributes) SupportsSigning() terra.BoolValue {
	return terra.ReferenceAsBool(tsbk.ref.Append("supports_signing"))
}

// Type returns a reference to field type of vault_transit_secret_backend_key.
func (tsbk transitSecretBackendKeyAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(tsbk.ref.Append("type"))
}

type transitSecretBackendKeyState struct {
	AllowPlaintextBackup bool                `json:"allow_plaintext_backup"`
	AutoRotateInterval   float64             `json:"auto_rotate_interval"`
	AutoRotatePeriod     float64             `json:"auto_rotate_period"`
	Backend              string              `json:"backend"`
	ConvergentEncryption bool                `json:"convergent_encryption"`
	DeletionAllowed      bool                `json:"deletion_allowed"`
	Derived              bool                `json:"derived"`
	Exportable           bool                `json:"exportable"`
	Id                   string              `json:"id"`
	Keys                 []map[string]string `json:"keys"`
	LatestVersion        float64             `json:"latest_version"`
	MinAvailableVersion  float64             `json:"min_available_version"`
	MinDecryptionVersion float64             `json:"min_decryption_version"`
	MinEncryptionVersion float64             `json:"min_encryption_version"`
	Name                 string              `json:"name"`
	Namespace            string              `json:"namespace"`
	SupportsDecryption   bool                `json:"supports_decryption"`
	SupportsDerivation   bool                `json:"supports_derivation"`
	SupportsEncryption   bool                `json:"supports_encryption"`
	SupportsSigning      bool                `json:"supports_signing"`
	Type                 string              `json:"type"`
}