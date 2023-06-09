// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package vault

import "github.com/volvo-cars/lingon/pkg/terra"

// NewDataAuthBackend creates a new instance of [DataAuthBackend].
func NewDataAuthBackend(name string, args DataAuthBackendArgs) *DataAuthBackend {
	return &DataAuthBackend{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataAuthBackend)(nil)

// DataAuthBackend represents the Terraform data resource vault_auth_backend.
type DataAuthBackend struct {
	Name string
	Args DataAuthBackendArgs
}

// DataSource returns the Terraform object type for [DataAuthBackend].
func (ab *DataAuthBackend) DataSource() string {
	return "vault_auth_backend"
}

// LocalName returns the local name for [DataAuthBackend].
func (ab *DataAuthBackend) LocalName() string {
	return ab.Name
}

// Configuration returns the configuration (args) for [DataAuthBackend].
func (ab *DataAuthBackend) Configuration() interface{} {
	return ab.Args
}

// Attributes returns the attributes for [DataAuthBackend].
func (ab *DataAuthBackend) Attributes() dataAuthBackendAttributes {
	return dataAuthBackendAttributes{ref: terra.ReferenceDataResource(ab)}
}

// DataAuthBackendArgs contains the configurations for vault_auth_backend.
type DataAuthBackendArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Namespace: string, optional
	Namespace terra.StringValue `hcl:"namespace,attr"`
	// Path: string, required
	Path terra.StringValue `hcl:"path,attr" validate:"required"`
}
type dataAuthBackendAttributes struct {
	ref terra.Reference
}

// Accessor returns a reference to field accessor of vault_auth_backend.
func (ab dataAuthBackendAttributes) Accessor() terra.StringValue {
	return terra.ReferenceAsString(ab.ref.Append("accessor"))
}

// DefaultLeaseTtlSeconds returns a reference to field default_lease_ttl_seconds of vault_auth_backend.
func (ab dataAuthBackendAttributes) DefaultLeaseTtlSeconds() terra.NumberValue {
	return terra.ReferenceAsNumber(ab.ref.Append("default_lease_ttl_seconds"))
}

// Description returns a reference to field description of vault_auth_backend.
func (ab dataAuthBackendAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(ab.ref.Append("description"))
}

// Id returns a reference to field id of vault_auth_backend.
func (ab dataAuthBackendAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ab.ref.Append("id"))
}

// ListingVisibility returns a reference to field listing_visibility of vault_auth_backend.
func (ab dataAuthBackendAttributes) ListingVisibility() terra.StringValue {
	return terra.ReferenceAsString(ab.ref.Append("listing_visibility"))
}

// Local returns a reference to field local of vault_auth_backend.
func (ab dataAuthBackendAttributes) Local() terra.BoolValue {
	return terra.ReferenceAsBool(ab.ref.Append("local"))
}

// MaxLeaseTtlSeconds returns a reference to field max_lease_ttl_seconds of vault_auth_backend.
func (ab dataAuthBackendAttributes) MaxLeaseTtlSeconds() terra.NumberValue {
	return terra.ReferenceAsNumber(ab.ref.Append("max_lease_ttl_seconds"))
}

// Namespace returns a reference to field namespace of vault_auth_backend.
func (ab dataAuthBackendAttributes) Namespace() terra.StringValue {
	return terra.ReferenceAsString(ab.ref.Append("namespace"))
}

// Path returns a reference to field path of vault_auth_backend.
func (ab dataAuthBackendAttributes) Path() terra.StringValue {
	return terra.ReferenceAsString(ab.ref.Append("path"))
}

// Type returns a reference to field type of vault_auth_backend.
func (ab dataAuthBackendAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(ab.ref.Append("type"))
}
