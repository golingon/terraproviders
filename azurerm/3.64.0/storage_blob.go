// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	storageblob "github.com/golingon/terraproviders/azurerm/3.64.0/storageblob"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewStorageBlob creates a new instance of [StorageBlob].
func NewStorageBlob(name string, args StorageBlobArgs) *StorageBlob {
	return &StorageBlob{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*StorageBlob)(nil)

// StorageBlob represents the Terraform resource azurerm_storage_blob.
type StorageBlob struct {
	Name      string
	Args      StorageBlobArgs
	state     *storageBlobState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [StorageBlob].
func (sb *StorageBlob) Type() string {
	return "azurerm_storage_blob"
}

// LocalName returns the local name for [StorageBlob].
func (sb *StorageBlob) LocalName() string {
	return sb.Name
}

// Configuration returns the configuration (args) for [StorageBlob].
func (sb *StorageBlob) Configuration() interface{} {
	return sb.Args
}

// DependOn is used for other resources to depend on [StorageBlob].
func (sb *StorageBlob) DependOn() terra.Reference {
	return terra.ReferenceResource(sb)
}

// Dependencies returns the list of resources [StorageBlob] depends_on.
func (sb *StorageBlob) Dependencies() terra.Dependencies {
	return sb.DependsOn
}

// LifecycleManagement returns the lifecycle block for [StorageBlob].
func (sb *StorageBlob) LifecycleManagement() *terra.Lifecycle {
	return sb.Lifecycle
}

// Attributes returns the attributes for [StorageBlob].
func (sb *StorageBlob) Attributes() storageBlobAttributes {
	return storageBlobAttributes{ref: terra.ReferenceResource(sb)}
}

// ImportState imports the given attribute values into [StorageBlob]'s state.
func (sb *StorageBlob) ImportState(av io.Reader) error {
	sb.state = &storageBlobState{}
	if err := json.NewDecoder(av).Decode(sb.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sb.Type(), sb.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [StorageBlob] has state.
func (sb *StorageBlob) State() (*storageBlobState, bool) {
	return sb.state, sb.state != nil
}

// StateMust returns the state for [StorageBlob]. Panics if the state is nil.
func (sb *StorageBlob) StateMust() *storageBlobState {
	if sb.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sb.Type(), sb.LocalName()))
	}
	return sb.state
}

// StorageBlobArgs contains the configurations for azurerm_storage_blob.
type StorageBlobArgs struct {
	// AccessTier: string, optional
	AccessTier terra.StringValue `hcl:"access_tier,attr"`
	// CacheControl: string, optional
	CacheControl terra.StringValue `hcl:"cache_control,attr"`
	// ContentMd5: string, optional
	ContentMd5 terra.StringValue `hcl:"content_md5,attr"`
	// ContentType: string, optional
	ContentType terra.StringValue `hcl:"content_type,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Metadata: map of string, optional
	Metadata terra.MapValue[terra.StringValue] `hcl:"metadata,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Parallelism: number, optional
	Parallelism terra.NumberValue `hcl:"parallelism,attr"`
	// Size: number, optional
	Size terra.NumberValue `hcl:"size,attr"`
	// Source: string, optional
	Source terra.StringValue `hcl:"source,attr"`
	// SourceContent: string, optional
	SourceContent terra.StringValue `hcl:"source_content,attr"`
	// SourceUri: string, optional
	SourceUri terra.StringValue `hcl:"source_uri,attr"`
	// StorageAccountName: string, required
	StorageAccountName terra.StringValue `hcl:"storage_account_name,attr" validate:"required"`
	// StorageContainerName: string, required
	StorageContainerName terra.StringValue `hcl:"storage_container_name,attr" validate:"required"`
	// Type: string, required
	Type terra.StringValue `hcl:"type,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *storageblob.Timeouts `hcl:"timeouts,block"`
}
type storageBlobAttributes struct {
	ref terra.Reference
}

// AccessTier returns a reference to field access_tier of azurerm_storage_blob.
func (sb storageBlobAttributes) AccessTier() terra.StringValue {
	return terra.ReferenceAsString(sb.ref.Append("access_tier"))
}

// CacheControl returns a reference to field cache_control of azurerm_storage_blob.
func (sb storageBlobAttributes) CacheControl() terra.StringValue {
	return terra.ReferenceAsString(sb.ref.Append("cache_control"))
}

// ContentMd5 returns a reference to field content_md5 of azurerm_storage_blob.
func (sb storageBlobAttributes) ContentMd5() terra.StringValue {
	return terra.ReferenceAsString(sb.ref.Append("content_md5"))
}

// ContentType returns a reference to field content_type of azurerm_storage_blob.
func (sb storageBlobAttributes) ContentType() terra.StringValue {
	return terra.ReferenceAsString(sb.ref.Append("content_type"))
}

// Id returns a reference to field id of azurerm_storage_blob.
func (sb storageBlobAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sb.ref.Append("id"))
}

// Metadata returns a reference to field metadata of azurerm_storage_blob.
func (sb storageBlobAttributes) Metadata() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](sb.ref.Append("metadata"))
}

// Name returns a reference to field name of azurerm_storage_blob.
func (sb storageBlobAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(sb.ref.Append("name"))
}

// Parallelism returns a reference to field parallelism of azurerm_storage_blob.
func (sb storageBlobAttributes) Parallelism() terra.NumberValue {
	return terra.ReferenceAsNumber(sb.ref.Append("parallelism"))
}

// Size returns a reference to field size of azurerm_storage_blob.
func (sb storageBlobAttributes) Size() terra.NumberValue {
	return terra.ReferenceAsNumber(sb.ref.Append("size"))
}

// Source returns a reference to field source of azurerm_storage_blob.
func (sb storageBlobAttributes) Source() terra.StringValue {
	return terra.ReferenceAsString(sb.ref.Append("source"))
}

// SourceContent returns a reference to field source_content of azurerm_storage_blob.
func (sb storageBlobAttributes) SourceContent() terra.StringValue {
	return terra.ReferenceAsString(sb.ref.Append("source_content"))
}

// SourceUri returns a reference to field source_uri of azurerm_storage_blob.
func (sb storageBlobAttributes) SourceUri() terra.StringValue {
	return terra.ReferenceAsString(sb.ref.Append("source_uri"))
}

// StorageAccountName returns a reference to field storage_account_name of azurerm_storage_blob.
func (sb storageBlobAttributes) StorageAccountName() terra.StringValue {
	return terra.ReferenceAsString(sb.ref.Append("storage_account_name"))
}

// StorageContainerName returns a reference to field storage_container_name of azurerm_storage_blob.
func (sb storageBlobAttributes) StorageContainerName() terra.StringValue {
	return terra.ReferenceAsString(sb.ref.Append("storage_container_name"))
}

// Type returns a reference to field type of azurerm_storage_blob.
func (sb storageBlobAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(sb.ref.Append("type"))
}

// Url returns a reference to field url of azurerm_storage_blob.
func (sb storageBlobAttributes) Url() terra.StringValue {
	return terra.ReferenceAsString(sb.ref.Append("url"))
}

func (sb storageBlobAttributes) Timeouts() storageblob.TimeoutsAttributes {
	return terra.ReferenceAsSingle[storageblob.TimeoutsAttributes](sb.ref.Append("timeouts"))
}

type storageBlobState struct {
	AccessTier           string                     `json:"access_tier"`
	CacheControl         string                     `json:"cache_control"`
	ContentMd5           string                     `json:"content_md5"`
	ContentType          string                     `json:"content_type"`
	Id                   string                     `json:"id"`
	Metadata             map[string]string          `json:"metadata"`
	Name                 string                     `json:"name"`
	Parallelism          float64                    `json:"parallelism"`
	Size                 float64                    `json:"size"`
	Source               string                     `json:"source"`
	SourceContent        string                     `json:"source_content"`
	SourceUri            string                     `json:"source_uri"`
	StorageAccountName   string                     `json:"storage_account_name"`
	StorageContainerName string                     `json:"storage_container_name"`
	Type                 string                     `json:"type"`
	Url                  string                     `json:"url"`
	Timeouts             *storageblob.TimeoutsState `json:"timeouts"`
}
