// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurestack_key_vault_key

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

// Resource represents the Terraform resource azurestack_key_vault_key.
type Resource struct {
	Name      string
	Args      Args
	state     *azurestackKeyVaultKeyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (akvk *Resource) Type() string {
	return "azurestack_key_vault_key"
}

// LocalName returns the local name for [Resource].
func (akvk *Resource) LocalName() string {
	return akvk.Name
}

// Configuration returns the configuration (args) for [Resource].
func (akvk *Resource) Configuration() interface{} {
	return akvk.Args
}

// DependOn is used for other resources to depend on [Resource].
func (akvk *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(akvk)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (akvk *Resource) Dependencies() terra.Dependencies {
	return akvk.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (akvk *Resource) LifecycleManagement() *terra.Lifecycle {
	return akvk.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (akvk *Resource) Attributes() azurestackKeyVaultKeyAttributes {
	return azurestackKeyVaultKeyAttributes{ref: terra.ReferenceResource(akvk)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (akvk *Resource) ImportState(state io.Reader) error {
	akvk.state = &azurestackKeyVaultKeyState{}
	if err := json.NewDecoder(state).Decode(akvk.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", akvk.Type(), akvk.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (akvk *Resource) State() (*azurestackKeyVaultKeyState, bool) {
	return akvk.state, akvk.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (akvk *Resource) StateMust() *azurestackKeyVaultKeyState {
	if akvk.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", akvk.Type(), akvk.LocalName()))
	}
	return akvk.state
}

// Args contains the configurations for azurestack_key_vault_key.
type Args struct {
	// Curve: string, optional
	Curve terra.StringValue `hcl:"curve,attr"`
	// ExpirationDate: string, optional
	ExpirationDate terra.StringValue `hcl:"expiration_date,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// KeyOpts: list of string, required
	KeyOpts terra.ListValue[terra.StringValue] `hcl:"key_opts,attr" validate:"required"`
	// KeySize: number, optional
	KeySize terra.NumberValue `hcl:"key_size,attr"`
	// KeyType: string, required
	KeyType terra.StringValue `hcl:"key_type,attr" validate:"required"`
	// KeyVaultId: string, required
	KeyVaultId terra.StringValue `hcl:"key_vault_id,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// NotBeforeDate: string, optional
	NotBeforeDate terra.StringValue `hcl:"not_before_date,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type azurestackKeyVaultKeyAttributes struct {
	ref terra.Reference
}

// Curve returns a reference to field curve of azurestack_key_vault_key.
func (akvk azurestackKeyVaultKeyAttributes) Curve() terra.StringValue {
	return terra.ReferenceAsString(akvk.ref.Append("curve"))
}

// E returns a reference to field e of azurestack_key_vault_key.
func (akvk azurestackKeyVaultKeyAttributes) E() terra.StringValue {
	return terra.ReferenceAsString(akvk.ref.Append("e"))
}

// ExpirationDate returns a reference to field expiration_date of azurestack_key_vault_key.
func (akvk azurestackKeyVaultKeyAttributes) ExpirationDate() terra.StringValue {
	return terra.ReferenceAsString(akvk.ref.Append("expiration_date"))
}

// Id returns a reference to field id of azurestack_key_vault_key.
func (akvk azurestackKeyVaultKeyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(akvk.ref.Append("id"))
}

// KeyOpts returns a reference to field key_opts of azurestack_key_vault_key.
func (akvk azurestackKeyVaultKeyAttributes) KeyOpts() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](akvk.ref.Append("key_opts"))
}

// KeySize returns a reference to field key_size of azurestack_key_vault_key.
func (akvk azurestackKeyVaultKeyAttributes) KeySize() terra.NumberValue {
	return terra.ReferenceAsNumber(akvk.ref.Append("key_size"))
}

// KeyType returns a reference to field key_type of azurestack_key_vault_key.
func (akvk azurestackKeyVaultKeyAttributes) KeyType() terra.StringValue {
	return terra.ReferenceAsString(akvk.ref.Append("key_type"))
}

// KeyVaultId returns a reference to field key_vault_id of azurestack_key_vault_key.
func (akvk azurestackKeyVaultKeyAttributes) KeyVaultId() terra.StringValue {
	return terra.ReferenceAsString(akvk.ref.Append("key_vault_id"))
}

// N returns a reference to field n of azurestack_key_vault_key.
func (akvk azurestackKeyVaultKeyAttributes) N() terra.StringValue {
	return terra.ReferenceAsString(akvk.ref.Append("n"))
}

// Name returns a reference to field name of azurestack_key_vault_key.
func (akvk azurestackKeyVaultKeyAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(akvk.ref.Append("name"))
}

// NotBeforeDate returns a reference to field not_before_date of azurestack_key_vault_key.
func (akvk azurestackKeyVaultKeyAttributes) NotBeforeDate() terra.StringValue {
	return terra.ReferenceAsString(akvk.ref.Append("not_before_date"))
}

// Tags returns a reference to field tags of azurestack_key_vault_key.
func (akvk azurestackKeyVaultKeyAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](akvk.ref.Append("tags"))
}

// Version returns a reference to field version of azurestack_key_vault_key.
func (akvk azurestackKeyVaultKeyAttributes) Version() terra.StringValue {
	return terra.ReferenceAsString(akvk.ref.Append("version"))
}

// VersionlessId returns a reference to field versionless_id of azurestack_key_vault_key.
func (akvk azurestackKeyVaultKeyAttributes) VersionlessId() terra.StringValue {
	return terra.ReferenceAsString(akvk.ref.Append("versionless_id"))
}

// X returns a reference to field x of azurestack_key_vault_key.
func (akvk azurestackKeyVaultKeyAttributes) X() terra.StringValue {
	return terra.ReferenceAsString(akvk.ref.Append("x"))
}

// Y returns a reference to field y of azurestack_key_vault_key.
func (akvk azurestackKeyVaultKeyAttributes) Y() terra.StringValue {
	return terra.ReferenceAsString(akvk.ref.Append("y"))
}

func (akvk azurestackKeyVaultKeyAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](akvk.ref.Append("timeouts"))
}

type azurestackKeyVaultKeyState struct {
	Curve          string            `json:"curve"`
	E              string            `json:"e"`
	ExpirationDate string            `json:"expiration_date"`
	Id             string            `json:"id"`
	KeyOpts        []string          `json:"key_opts"`
	KeySize        float64           `json:"key_size"`
	KeyType        string            `json:"key_type"`
	KeyVaultId     string            `json:"key_vault_id"`
	N              string            `json:"n"`
	Name           string            `json:"name"`
	NotBeforeDate  string            `json:"not_before_date"`
	Tags           map[string]string `json:"tags"`
	Version        string            `json:"version"`
	VersionlessId  string            `json:"versionless_id"`
	X              string            `json:"x"`
	Y              string            `json:"y"`
	Timeouts       *TimeoutsState    `json:"timeouts"`
}
