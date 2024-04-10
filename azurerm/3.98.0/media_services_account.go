// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	mediaservicesaccount "github.com/golingon/terraproviders/azurerm/3.98.0/mediaservicesaccount"
	"io"
)

// NewMediaServicesAccount creates a new instance of [MediaServicesAccount].
func NewMediaServicesAccount(name string, args MediaServicesAccountArgs) *MediaServicesAccount {
	return &MediaServicesAccount{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*MediaServicesAccount)(nil)

// MediaServicesAccount represents the Terraform resource azurerm_media_services_account.
type MediaServicesAccount struct {
	Name      string
	Args      MediaServicesAccountArgs
	state     *mediaServicesAccountState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [MediaServicesAccount].
func (msa *MediaServicesAccount) Type() string {
	return "azurerm_media_services_account"
}

// LocalName returns the local name for [MediaServicesAccount].
func (msa *MediaServicesAccount) LocalName() string {
	return msa.Name
}

// Configuration returns the configuration (args) for [MediaServicesAccount].
func (msa *MediaServicesAccount) Configuration() interface{} {
	return msa.Args
}

// DependOn is used for other resources to depend on [MediaServicesAccount].
func (msa *MediaServicesAccount) DependOn() terra.Reference {
	return terra.ReferenceResource(msa)
}

// Dependencies returns the list of resources [MediaServicesAccount] depends_on.
func (msa *MediaServicesAccount) Dependencies() terra.Dependencies {
	return msa.DependsOn
}

// LifecycleManagement returns the lifecycle block for [MediaServicesAccount].
func (msa *MediaServicesAccount) LifecycleManagement() *terra.Lifecycle {
	return msa.Lifecycle
}

// Attributes returns the attributes for [MediaServicesAccount].
func (msa *MediaServicesAccount) Attributes() mediaServicesAccountAttributes {
	return mediaServicesAccountAttributes{ref: terra.ReferenceResource(msa)}
}

// ImportState imports the given attribute values into [MediaServicesAccount]'s state.
func (msa *MediaServicesAccount) ImportState(av io.Reader) error {
	msa.state = &mediaServicesAccountState{}
	if err := json.NewDecoder(av).Decode(msa.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", msa.Type(), msa.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [MediaServicesAccount] has state.
func (msa *MediaServicesAccount) State() (*mediaServicesAccountState, bool) {
	return msa.state, msa.state != nil
}

// StateMust returns the state for [MediaServicesAccount]. Panics if the state is nil.
func (msa *MediaServicesAccount) StateMust() *mediaServicesAccountState {
	if msa.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", msa.Type(), msa.LocalName()))
	}
	return msa.state
}

// MediaServicesAccountArgs contains the configurations for azurerm_media_services_account.
type MediaServicesAccountArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// PublicNetworkAccessEnabled: bool, optional
	PublicNetworkAccessEnabled terra.BoolValue `hcl:"public_network_access_enabled,attr"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// StorageAuthenticationType: string, optional
	StorageAuthenticationType terra.StringValue `hcl:"storage_authentication_type,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Encryption: optional
	Encryption *mediaservicesaccount.Encryption `hcl:"encryption,block"`
	// Identity: optional
	Identity *mediaservicesaccount.Identity `hcl:"identity,block"`
	// KeyDeliveryAccessControl: optional
	KeyDeliveryAccessControl *mediaservicesaccount.KeyDeliveryAccessControl `hcl:"key_delivery_access_control,block"`
	// StorageAccount: min=1
	StorageAccount []mediaservicesaccount.StorageAccount `hcl:"storage_account,block" validate:"min=1"`
	// Timeouts: optional
	Timeouts *mediaservicesaccount.Timeouts `hcl:"timeouts,block"`
}
type mediaServicesAccountAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_media_services_account.
func (msa mediaServicesAccountAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(msa.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_media_services_account.
func (msa mediaServicesAccountAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(msa.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_media_services_account.
func (msa mediaServicesAccountAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(msa.ref.Append("name"))
}

// PublicNetworkAccessEnabled returns a reference to field public_network_access_enabled of azurerm_media_services_account.
func (msa mediaServicesAccountAttributes) PublicNetworkAccessEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(msa.ref.Append("public_network_access_enabled"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_media_services_account.
func (msa mediaServicesAccountAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(msa.ref.Append("resource_group_name"))
}

// StorageAuthenticationType returns a reference to field storage_authentication_type of azurerm_media_services_account.
func (msa mediaServicesAccountAttributes) StorageAuthenticationType() terra.StringValue {
	return terra.ReferenceAsString(msa.ref.Append("storage_authentication_type"))
}

// Tags returns a reference to field tags of azurerm_media_services_account.
func (msa mediaServicesAccountAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](msa.ref.Append("tags"))
}

func (msa mediaServicesAccountAttributes) Encryption() terra.ListValue[mediaservicesaccount.EncryptionAttributes] {
	return terra.ReferenceAsList[mediaservicesaccount.EncryptionAttributes](msa.ref.Append("encryption"))
}

func (msa mediaServicesAccountAttributes) Identity() terra.ListValue[mediaservicesaccount.IdentityAttributes] {
	return terra.ReferenceAsList[mediaservicesaccount.IdentityAttributes](msa.ref.Append("identity"))
}

func (msa mediaServicesAccountAttributes) KeyDeliveryAccessControl() terra.ListValue[mediaservicesaccount.KeyDeliveryAccessControlAttributes] {
	return terra.ReferenceAsList[mediaservicesaccount.KeyDeliveryAccessControlAttributes](msa.ref.Append("key_delivery_access_control"))
}

func (msa mediaServicesAccountAttributes) StorageAccount() terra.SetValue[mediaservicesaccount.StorageAccountAttributes] {
	return terra.ReferenceAsSet[mediaservicesaccount.StorageAccountAttributes](msa.ref.Append("storage_account"))
}

func (msa mediaServicesAccountAttributes) Timeouts() mediaservicesaccount.TimeoutsAttributes {
	return terra.ReferenceAsSingle[mediaservicesaccount.TimeoutsAttributes](msa.ref.Append("timeouts"))
}

type mediaServicesAccountState struct {
	Id                         string                                               `json:"id"`
	Location                   string                                               `json:"location"`
	Name                       string                                               `json:"name"`
	PublicNetworkAccessEnabled bool                                                 `json:"public_network_access_enabled"`
	ResourceGroupName          string                                               `json:"resource_group_name"`
	StorageAuthenticationType  string                                               `json:"storage_authentication_type"`
	Tags                       map[string]string                                    `json:"tags"`
	Encryption                 []mediaservicesaccount.EncryptionState               `json:"encryption"`
	Identity                   []mediaservicesaccount.IdentityState                 `json:"identity"`
	KeyDeliveryAccessControl   []mediaservicesaccount.KeyDeliveryAccessControlState `json:"key_delivery_access_control"`
	StorageAccount             []mediaservicesaccount.StorageAccountState           `json:"storage_account"`
	Timeouts                   *mediaservicesaccount.TimeoutsState                  `json:"timeouts"`
}
