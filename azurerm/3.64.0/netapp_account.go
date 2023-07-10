// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	netappaccount "github.com/golingon/terraproviders/azurerm/3.64.0/netappaccount"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewNetappAccount creates a new instance of [NetappAccount].
func NewNetappAccount(name string, args NetappAccountArgs) *NetappAccount {
	return &NetappAccount{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*NetappAccount)(nil)

// NetappAccount represents the Terraform resource azurerm_netapp_account.
type NetappAccount struct {
	Name      string
	Args      NetappAccountArgs
	state     *netappAccountState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [NetappAccount].
func (na *NetappAccount) Type() string {
	return "azurerm_netapp_account"
}

// LocalName returns the local name for [NetappAccount].
func (na *NetappAccount) LocalName() string {
	return na.Name
}

// Configuration returns the configuration (args) for [NetappAccount].
func (na *NetappAccount) Configuration() interface{} {
	return na.Args
}

// DependOn is used for other resources to depend on [NetappAccount].
func (na *NetappAccount) DependOn() terra.Reference {
	return terra.ReferenceResource(na)
}

// Dependencies returns the list of resources [NetappAccount] depends_on.
func (na *NetappAccount) Dependencies() terra.Dependencies {
	return na.DependsOn
}

// LifecycleManagement returns the lifecycle block for [NetappAccount].
func (na *NetappAccount) LifecycleManagement() *terra.Lifecycle {
	return na.Lifecycle
}

// Attributes returns the attributes for [NetappAccount].
func (na *NetappAccount) Attributes() netappAccountAttributes {
	return netappAccountAttributes{ref: terra.ReferenceResource(na)}
}

// ImportState imports the given attribute values into [NetappAccount]'s state.
func (na *NetappAccount) ImportState(av io.Reader) error {
	na.state = &netappAccountState{}
	if err := json.NewDecoder(av).Decode(na.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", na.Type(), na.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [NetappAccount] has state.
func (na *NetappAccount) State() (*netappAccountState, bool) {
	return na.state, na.state != nil
}

// StateMust returns the state for [NetappAccount]. Panics if the state is nil.
func (na *NetappAccount) StateMust() *netappAccountState {
	if na.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", na.Type(), na.LocalName()))
	}
	return na.state
}

// NetappAccountArgs contains the configurations for azurerm_netapp_account.
type NetappAccountArgs struct {
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
	// ActiveDirectory: optional
	ActiveDirectory *netappaccount.ActiveDirectory `hcl:"active_directory,block"`
	// Timeouts: optional
	Timeouts *netappaccount.Timeouts `hcl:"timeouts,block"`
}
type netappAccountAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_netapp_account.
func (na netappAccountAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(na.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_netapp_account.
func (na netappAccountAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(na.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_netapp_account.
func (na netappAccountAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(na.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_netapp_account.
func (na netappAccountAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(na.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_netapp_account.
func (na netappAccountAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](na.ref.Append("tags"))
}

func (na netappAccountAttributes) ActiveDirectory() terra.ListValue[netappaccount.ActiveDirectoryAttributes] {
	return terra.ReferenceAsList[netappaccount.ActiveDirectoryAttributes](na.ref.Append("active_directory"))
}

func (na netappAccountAttributes) Timeouts() netappaccount.TimeoutsAttributes {
	return terra.ReferenceAsSingle[netappaccount.TimeoutsAttributes](na.ref.Append("timeouts"))
}

type netappAccountState struct {
	Id                string                               `json:"id"`
	Location          string                               `json:"location"`
	Name              string                               `json:"name"`
	ResourceGroupName string                               `json:"resource_group_name"`
	Tags              map[string]string                    `json:"tags"`
	ActiveDirectory   []netappaccount.ActiveDirectoryState `json:"active_directory"`
	Timeouts          *netappaccount.TimeoutsState         `json:"timeouts"`
}
