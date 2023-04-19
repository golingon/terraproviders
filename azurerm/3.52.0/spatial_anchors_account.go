// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	spatialanchorsaccount "github.com/golingon/terraproviders/azurerm/3.52.0/spatialanchorsaccount"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSpatialAnchorsAccount creates a new instance of [SpatialAnchorsAccount].
func NewSpatialAnchorsAccount(name string, args SpatialAnchorsAccountArgs) *SpatialAnchorsAccount {
	return &SpatialAnchorsAccount{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SpatialAnchorsAccount)(nil)

// SpatialAnchorsAccount represents the Terraform resource azurerm_spatial_anchors_account.
type SpatialAnchorsAccount struct {
	Name      string
	Args      SpatialAnchorsAccountArgs
	state     *spatialAnchorsAccountState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SpatialAnchorsAccount].
func (saa *SpatialAnchorsAccount) Type() string {
	return "azurerm_spatial_anchors_account"
}

// LocalName returns the local name for [SpatialAnchorsAccount].
func (saa *SpatialAnchorsAccount) LocalName() string {
	return saa.Name
}

// Configuration returns the configuration (args) for [SpatialAnchorsAccount].
func (saa *SpatialAnchorsAccount) Configuration() interface{} {
	return saa.Args
}

// DependOn is used for other resources to depend on [SpatialAnchorsAccount].
func (saa *SpatialAnchorsAccount) DependOn() terra.Reference {
	return terra.ReferenceResource(saa)
}

// Dependencies returns the list of resources [SpatialAnchorsAccount] depends_on.
func (saa *SpatialAnchorsAccount) Dependencies() terra.Dependencies {
	return saa.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SpatialAnchorsAccount].
func (saa *SpatialAnchorsAccount) LifecycleManagement() *terra.Lifecycle {
	return saa.Lifecycle
}

// Attributes returns the attributes for [SpatialAnchorsAccount].
func (saa *SpatialAnchorsAccount) Attributes() spatialAnchorsAccountAttributes {
	return spatialAnchorsAccountAttributes{ref: terra.ReferenceResource(saa)}
}

// ImportState imports the given attribute values into [SpatialAnchorsAccount]'s state.
func (saa *SpatialAnchorsAccount) ImportState(av io.Reader) error {
	saa.state = &spatialAnchorsAccountState{}
	if err := json.NewDecoder(av).Decode(saa.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", saa.Type(), saa.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SpatialAnchorsAccount] has state.
func (saa *SpatialAnchorsAccount) State() (*spatialAnchorsAccountState, bool) {
	return saa.state, saa.state != nil
}

// StateMust returns the state for [SpatialAnchorsAccount]. Panics if the state is nil.
func (saa *SpatialAnchorsAccount) StateMust() *spatialAnchorsAccountState {
	if saa.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", saa.Type(), saa.LocalName()))
	}
	return saa.state
}

// SpatialAnchorsAccountArgs contains the configurations for azurerm_spatial_anchors_account.
type SpatialAnchorsAccountArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Timeouts: optional
	Timeouts *spatialanchorsaccount.Timeouts `hcl:"timeouts,block"`
}
type spatialAnchorsAccountAttributes struct {
	ref terra.Reference
}

// AccountDomain returns a reference to field account_domain of azurerm_spatial_anchors_account.
func (saa spatialAnchorsAccountAttributes) AccountDomain() terra.StringValue {
	return terra.ReferenceAsString(saa.ref.Append("account_domain"))
}

// AccountId returns a reference to field account_id of azurerm_spatial_anchors_account.
func (saa spatialAnchorsAccountAttributes) AccountId() terra.StringValue {
	return terra.ReferenceAsString(saa.ref.Append("account_id"))
}

// Id returns a reference to field id of azurerm_spatial_anchors_account.
func (saa spatialAnchorsAccountAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(saa.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_spatial_anchors_account.
func (saa spatialAnchorsAccountAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(saa.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_spatial_anchors_account.
func (saa spatialAnchorsAccountAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(saa.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_spatial_anchors_account.
func (saa spatialAnchorsAccountAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(saa.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_spatial_anchors_account.
func (saa spatialAnchorsAccountAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](saa.ref.Append("tags"))
}

func (saa spatialAnchorsAccountAttributes) Timeouts() spatialanchorsaccount.TimeoutsAttributes {
	return terra.ReferenceAsSingle[spatialanchorsaccount.TimeoutsAttributes](saa.ref.Append("timeouts"))
}

type spatialAnchorsAccountState struct {
	AccountDomain     string                               `json:"account_domain"`
	AccountId         string                               `json:"account_id"`
	Id                string                               `json:"id"`
	Location          string                               `json:"location"`
	Name              string                               `json:"name"`
	ResourceGroupName string                               `json:"resource_group_name"`
	Tags              map[string]string                    `json:"tags"`
	Timeouts          *spatialanchorsaccount.TimeoutsState `json:"timeouts"`
}
