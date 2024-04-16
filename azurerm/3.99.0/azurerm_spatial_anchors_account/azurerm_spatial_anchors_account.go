// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_spatial_anchors_account

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

// Resource represents the Terraform resource azurerm_spatial_anchors_account.
type Resource struct {
	Name      string
	Args      Args
	state     *azurermSpatialAnchorsAccountState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (asaa *Resource) Type() string {
	return "azurerm_spatial_anchors_account"
}

// LocalName returns the local name for [Resource].
func (asaa *Resource) LocalName() string {
	return asaa.Name
}

// Configuration returns the configuration (args) for [Resource].
func (asaa *Resource) Configuration() interface{} {
	return asaa.Args
}

// DependOn is used for other resources to depend on [Resource].
func (asaa *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(asaa)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (asaa *Resource) Dependencies() terra.Dependencies {
	return asaa.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (asaa *Resource) LifecycleManagement() *terra.Lifecycle {
	return asaa.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (asaa *Resource) Attributes() azurermSpatialAnchorsAccountAttributes {
	return azurermSpatialAnchorsAccountAttributes{ref: terra.ReferenceResource(asaa)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (asaa *Resource) ImportState(state io.Reader) error {
	asaa.state = &azurermSpatialAnchorsAccountState{}
	if err := json.NewDecoder(state).Decode(asaa.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", asaa.Type(), asaa.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (asaa *Resource) State() (*azurermSpatialAnchorsAccountState, bool) {
	return asaa.state, asaa.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (asaa *Resource) StateMust() *azurermSpatialAnchorsAccountState {
	if asaa.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", asaa.Type(), asaa.LocalName()))
	}
	return asaa.state
}

// Args contains the configurations for azurerm_spatial_anchors_account.
type Args struct {
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
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type azurermSpatialAnchorsAccountAttributes struct {
	ref terra.Reference
}

// AccountDomain returns a reference to field account_domain of azurerm_spatial_anchors_account.
func (asaa azurermSpatialAnchorsAccountAttributes) AccountDomain() terra.StringValue {
	return terra.ReferenceAsString(asaa.ref.Append("account_domain"))
}

// AccountId returns a reference to field account_id of azurerm_spatial_anchors_account.
func (asaa azurermSpatialAnchorsAccountAttributes) AccountId() terra.StringValue {
	return terra.ReferenceAsString(asaa.ref.Append("account_id"))
}

// Id returns a reference to field id of azurerm_spatial_anchors_account.
func (asaa azurermSpatialAnchorsAccountAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(asaa.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_spatial_anchors_account.
func (asaa azurermSpatialAnchorsAccountAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(asaa.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_spatial_anchors_account.
func (asaa azurermSpatialAnchorsAccountAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(asaa.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_spatial_anchors_account.
func (asaa azurermSpatialAnchorsAccountAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(asaa.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_spatial_anchors_account.
func (asaa azurermSpatialAnchorsAccountAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](asaa.ref.Append("tags"))
}

func (asaa azurermSpatialAnchorsAccountAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](asaa.ref.Append("timeouts"))
}

type azurermSpatialAnchorsAccountState struct {
	AccountDomain     string            `json:"account_domain"`
	AccountId         string            `json:"account_id"`
	Id                string            `json:"id"`
	Location          string            `json:"location"`
	Name              string            `json:"name"`
	ResourceGroupName string            `json:"resource_group_name"`
	Tags              map[string]string `json:"tags"`
	Timeouts          *TimeoutsState    `json:"timeouts"`
}
