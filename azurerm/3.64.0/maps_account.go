// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	mapsaccount "github.com/golingon/terraproviders/azurerm/3.64.0/mapsaccount"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewMapsAccount creates a new instance of [MapsAccount].
func NewMapsAccount(name string, args MapsAccountArgs) *MapsAccount {
	return &MapsAccount{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*MapsAccount)(nil)

// MapsAccount represents the Terraform resource azurerm_maps_account.
type MapsAccount struct {
	Name      string
	Args      MapsAccountArgs
	state     *mapsAccountState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [MapsAccount].
func (ma *MapsAccount) Type() string {
	return "azurerm_maps_account"
}

// LocalName returns the local name for [MapsAccount].
func (ma *MapsAccount) LocalName() string {
	return ma.Name
}

// Configuration returns the configuration (args) for [MapsAccount].
func (ma *MapsAccount) Configuration() interface{} {
	return ma.Args
}

// DependOn is used for other resources to depend on [MapsAccount].
func (ma *MapsAccount) DependOn() terra.Reference {
	return terra.ReferenceResource(ma)
}

// Dependencies returns the list of resources [MapsAccount] depends_on.
func (ma *MapsAccount) Dependencies() terra.Dependencies {
	return ma.DependsOn
}

// LifecycleManagement returns the lifecycle block for [MapsAccount].
func (ma *MapsAccount) LifecycleManagement() *terra.Lifecycle {
	return ma.Lifecycle
}

// Attributes returns the attributes for [MapsAccount].
func (ma *MapsAccount) Attributes() mapsAccountAttributes {
	return mapsAccountAttributes{ref: terra.ReferenceResource(ma)}
}

// ImportState imports the given attribute values into [MapsAccount]'s state.
func (ma *MapsAccount) ImportState(av io.Reader) error {
	ma.state = &mapsAccountState{}
	if err := json.NewDecoder(av).Decode(ma.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ma.Type(), ma.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [MapsAccount] has state.
func (ma *MapsAccount) State() (*mapsAccountState, bool) {
	return ma.state, ma.state != nil
}

// StateMust returns the state for [MapsAccount]. Panics if the state is nil.
func (ma *MapsAccount) StateMust() *mapsAccountState {
	if ma.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ma.Type(), ma.LocalName()))
	}
	return ma.state
}

// MapsAccountArgs contains the configurations for azurerm_maps_account.
type MapsAccountArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// SkuName: string, required
	SkuName terra.StringValue `hcl:"sku_name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Timeouts: optional
	Timeouts *mapsaccount.Timeouts `hcl:"timeouts,block"`
}
type mapsAccountAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_maps_account.
func (ma mapsAccountAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ma.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_maps_account.
func (ma mapsAccountAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ma.ref.Append("name"))
}

// PrimaryAccessKey returns a reference to field primary_access_key of azurerm_maps_account.
func (ma mapsAccountAttributes) PrimaryAccessKey() terra.StringValue {
	return terra.ReferenceAsString(ma.ref.Append("primary_access_key"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_maps_account.
func (ma mapsAccountAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(ma.ref.Append("resource_group_name"))
}

// SecondaryAccessKey returns a reference to field secondary_access_key of azurerm_maps_account.
func (ma mapsAccountAttributes) SecondaryAccessKey() terra.StringValue {
	return terra.ReferenceAsString(ma.ref.Append("secondary_access_key"))
}

// SkuName returns a reference to field sku_name of azurerm_maps_account.
func (ma mapsAccountAttributes) SkuName() terra.StringValue {
	return terra.ReferenceAsString(ma.ref.Append("sku_name"))
}

// Tags returns a reference to field tags of azurerm_maps_account.
func (ma mapsAccountAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ma.ref.Append("tags"))
}

// XMsClientId returns a reference to field x_ms_client_id of azurerm_maps_account.
func (ma mapsAccountAttributes) XMsClientId() terra.StringValue {
	return terra.ReferenceAsString(ma.ref.Append("x_ms_client_id"))
}

func (ma mapsAccountAttributes) Timeouts() mapsaccount.TimeoutsAttributes {
	return terra.ReferenceAsSingle[mapsaccount.TimeoutsAttributes](ma.ref.Append("timeouts"))
}

type mapsAccountState struct {
	Id                 string                     `json:"id"`
	Name               string                     `json:"name"`
	PrimaryAccessKey   string                     `json:"primary_access_key"`
	ResourceGroupName  string                     `json:"resource_group_name"`
	SecondaryAccessKey string                     `json:"secondary_access_key"`
	SkuName            string                     `json:"sku_name"`
	Tags               map[string]string          `json:"tags"`
	XMsClientId        string                     `json:"x_ms_client_id"`
	Timeouts           *mapsaccount.TimeoutsState `json:"timeouts"`
}
