// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_data_share_account

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

// Resource represents the Terraform resource azurerm_data_share_account.
type Resource struct {
	Name      string
	Args      Args
	state     *azurermDataShareAccountState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (adsa *Resource) Type() string {
	return "azurerm_data_share_account"
}

// LocalName returns the local name for [Resource].
func (adsa *Resource) LocalName() string {
	return adsa.Name
}

// Configuration returns the configuration (args) for [Resource].
func (adsa *Resource) Configuration() interface{} {
	return adsa.Args
}

// DependOn is used for other resources to depend on [Resource].
func (adsa *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(adsa)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (adsa *Resource) Dependencies() terra.Dependencies {
	return adsa.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (adsa *Resource) LifecycleManagement() *terra.Lifecycle {
	return adsa.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (adsa *Resource) Attributes() azurermDataShareAccountAttributes {
	return azurermDataShareAccountAttributes{ref: terra.ReferenceResource(adsa)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (adsa *Resource) ImportState(state io.Reader) error {
	adsa.state = &azurermDataShareAccountState{}
	if err := json.NewDecoder(state).Decode(adsa.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", adsa.Type(), adsa.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (adsa *Resource) State() (*azurermDataShareAccountState, bool) {
	return adsa.state, adsa.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (adsa *Resource) StateMust() *azurermDataShareAccountState {
	if adsa.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", adsa.Type(), adsa.LocalName()))
	}
	return adsa.state
}

// Args contains the configurations for azurerm_data_share_account.
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
	// Identity: required
	Identity *Identity `hcl:"identity,block" validate:"required"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type azurermDataShareAccountAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_data_share_account.
func (adsa azurermDataShareAccountAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(adsa.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_data_share_account.
func (adsa azurermDataShareAccountAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(adsa.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_data_share_account.
func (adsa azurermDataShareAccountAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(adsa.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_data_share_account.
func (adsa azurermDataShareAccountAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(adsa.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_data_share_account.
func (adsa azurermDataShareAccountAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](adsa.ref.Append("tags"))
}

func (adsa azurermDataShareAccountAttributes) Identity() terra.ListValue[IdentityAttributes] {
	return terra.ReferenceAsList[IdentityAttributes](adsa.ref.Append("identity"))
}

func (adsa azurermDataShareAccountAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](adsa.ref.Append("timeouts"))
}

type azurermDataShareAccountState struct {
	Id                string            `json:"id"`
	Location          string            `json:"location"`
	Name              string            `json:"name"`
	ResourceGroupName string            `json:"resource_group_name"`
	Tags              map[string]string `json:"tags"`
	Identity          []IdentityState   `json:"identity"`
	Timeouts          *TimeoutsState    `json:"timeouts"`
}
