// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	mediaasset "github.com/golingon/terraproviders/azurerm/3.55.0/mediaasset"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewMediaAsset creates a new instance of [MediaAsset].
func NewMediaAsset(name string, args MediaAssetArgs) *MediaAsset {
	return &MediaAsset{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*MediaAsset)(nil)

// MediaAsset represents the Terraform resource azurerm_media_asset.
type MediaAsset struct {
	Name      string
	Args      MediaAssetArgs
	state     *mediaAssetState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [MediaAsset].
func (ma *MediaAsset) Type() string {
	return "azurerm_media_asset"
}

// LocalName returns the local name for [MediaAsset].
func (ma *MediaAsset) LocalName() string {
	return ma.Name
}

// Configuration returns the configuration (args) for [MediaAsset].
func (ma *MediaAsset) Configuration() interface{} {
	return ma.Args
}

// DependOn is used for other resources to depend on [MediaAsset].
func (ma *MediaAsset) DependOn() terra.Reference {
	return terra.ReferenceResource(ma)
}

// Dependencies returns the list of resources [MediaAsset] depends_on.
func (ma *MediaAsset) Dependencies() terra.Dependencies {
	return ma.DependsOn
}

// LifecycleManagement returns the lifecycle block for [MediaAsset].
func (ma *MediaAsset) LifecycleManagement() *terra.Lifecycle {
	return ma.Lifecycle
}

// Attributes returns the attributes for [MediaAsset].
func (ma *MediaAsset) Attributes() mediaAssetAttributes {
	return mediaAssetAttributes{ref: terra.ReferenceResource(ma)}
}

// ImportState imports the given attribute values into [MediaAsset]'s state.
func (ma *MediaAsset) ImportState(av io.Reader) error {
	ma.state = &mediaAssetState{}
	if err := json.NewDecoder(av).Decode(ma.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ma.Type(), ma.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [MediaAsset] has state.
func (ma *MediaAsset) State() (*mediaAssetState, bool) {
	return ma.state, ma.state != nil
}

// StateMust returns the state for [MediaAsset]. Panics if the state is nil.
func (ma *MediaAsset) StateMust() *mediaAssetState {
	if ma.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ma.Type(), ma.LocalName()))
	}
	return ma.state
}

// MediaAssetArgs contains the configurations for azurerm_media_asset.
type MediaAssetArgs struct {
	// AlternateId: string, optional
	AlternateId terra.StringValue `hcl:"alternate_id,attr"`
	// Container: string, optional
	Container terra.StringValue `hcl:"container,attr"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// MediaServicesAccountName: string, required
	MediaServicesAccountName terra.StringValue `hcl:"media_services_account_name,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// StorageAccountName: string, optional
	StorageAccountName terra.StringValue `hcl:"storage_account_name,attr"`
	// Timeouts: optional
	Timeouts *mediaasset.Timeouts `hcl:"timeouts,block"`
}
type mediaAssetAttributes struct {
	ref terra.Reference
}

// AlternateId returns a reference to field alternate_id of azurerm_media_asset.
func (ma mediaAssetAttributes) AlternateId() terra.StringValue {
	return terra.ReferenceAsString(ma.ref.Append("alternate_id"))
}

// Container returns a reference to field container of azurerm_media_asset.
func (ma mediaAssetAttributes) Container() terra.StringValue {
	return terra.ReferenceAsString(ma.ref.Append("container"))
}

// Description returns a reference to field description of azurerm_media_asset.
func (ma mediaAssetAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(ma.ref.Append("description"))
}

// Id returns a reference to field id of azurerm_media_asset.
func (ma mediaAssetAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ma.ref.Append("id"))
}

// MediaServicesAccountName returns a reference to field media_services_account_name of azurerm_media_asset.
func (ma mediaAssetAttributes) MediaServicesAccountName() terra.StringValue {
	return terra.ReferenceAsString(ma.ref.Append("media_services_account_name"))
}

// Name returns a reference to field name of azurerm_media_asset.
func (ma mediaAssetAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ma.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_media_asset.
func (ma mediaAssetAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(ma.ref.Append("resource_group_name"))
}

// StorageAccountName returns a reference to field storage_account_name of azurerm_media_asset.
func (ma mediaAssetAttributes) StorageAccountName() terra.StringValue {
	return terra.ReferenceAsString(ma.ref.Append("storage_account_name"))
}

func (ma mediaAssetAttributes) Timeouts() mediaasset.TimeoutsAttributes {
	return terra.ReferenceAsSingle[mediaasset.TimeoutsAttributes](ma.ref.Append("timeouts"))
}

type mediaAssetState struct {
	AlternateId              string                    `json:"alternate_id"`
	Container                string                    `json:"container"`
	Description              string                    `json:"description"`
	Id                       string                    `json:"id"`
	MediaServicesAccountName string                    `json:"media_services_account_name"`
	Name                     string                    `json:"name"`
	ResourceGroupName        string                    `json:"resource_group_name"`
	StorageAccountName       string                    `json:"storage_account_name"`
	Timeouts                 *mediaasset.TimeoutsState `json:"timeouts"`
}
