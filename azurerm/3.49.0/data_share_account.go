// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	datashareaccount "github.com/golingon/terraproviders/azurerm/3.49.0/datashareaccount"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDataShareAccount creates a new instance of [DataShareAccount].
func NewDataShareAccount(name string, args DataShareAccountArgs) *DataShareAccount {
	return &DataShareAccount{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DataShareAccount)(nil)

// DataShareAccount represents the Terraform resource azurerm_data_share_account.
type DataShareAccount struct {
	Name      string
	Args      DataShareAccountArgs
	state     *dataShareAccountState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DataShareAccount].
func (dsa *DataShareAccount) Type() string {
	return "azurerm_data_share_account"
}

// LocalName returns the local name for [DataShareAccount].
func (dsa *DataShareAccount) LocalName() string {
	return dsa.Name
}

// Configuration returns the configuration (args) for [DataShareAccount].
func (dsa *DataShareAccount) Configuration() interface{} {
	return dsa.Args
}

// DependOn is used for other resources to depend on [DataShareAccount].
func (dsa *DataShareAccount) DependOn() terra.Reference {
	return terra.ReferenceResource(dsa)
}

// Dependencies returns the list of resources [DataShareAccount] depends_on.
func (dsa *DataShareAccount) Dependencies() terra.Dependencies {
	return dsa.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DataShareAccount].
func (dsa *DataShareAccount) LifecycleManagement() *terra.Lifecycle {
	return dsa.Lifecycle
}

// Attributes returns the attributes for [DataShareAccount].
func (dsa *DataShareAccount) Attributes() dataShareAccountAttributes {
	return dataShareAccountAttributes{ref: terra.ReferenceResource(dsa)}
}

// ImportState imports the given attribute values into [DataShareAccount]'s state.
func (dsa *DataShareAccount) ImportState(av io.Reader) error {
	dsa.state = &dataShareAccountState{}
	if err := json.NewDecoder(av).Decode(dsa.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dsa.Type(), dsa.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DataShareAccount] has state.
func (dsa *DataShareAccount) State() (*dataShareAccountState, bool) {
	return dsa.state, dsa.state != nil
}

// StateMust returns the state for [DataShareAccount]. Panics if the state is nil.
func (dsa *DataShareAccount) StateMust() *dataShareAccountState {
	if dsa.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dsa.Type(), dsa.LocalName()))
	}
	return dsa.state
}

// DataShareAccountArgs contains the configurations for azurerm_data_share_account.
type DataShareAccountArgs struct {
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
	Identity *datashareaccount.Identity `hcl:"identity,block" validate:"required"`
	// Timeouts: optional
	Timeouts *datashareaccount.Timeouts `hcl:"timeouts,block"`
}
type dataShareAccountAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_data_share_account.
func (dsa dataShareAccountAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dsa.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_data_share_account.
func (dsa dataShareAccountAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(dsa.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_data_share_account.
func (dsa dataShareAccountAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(dsa.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_data_share_account.
func (dsa dataShareAccountAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(dsa.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_data_share_account.
func (dsa dataShareAccountAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dsa.ref.Append("tags"))
}

func (dsa dataShareAccountAttributes) Identity() terra.ListValue[datashareaccount.IdentityAttributes] {
	return terra.ReferenceAsList[datashareaccount.IdentityAttributes](dsa.ref.Append("identity"))
}

func (dsa dataShareAccountAttributes) Timeouts() datashareaccount.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datashareaccount.TimeoutsAttributes](dsa.ref.Append("timeouts"))
}

type dataShareAccountState struct {
	Id                string                           `json:"id"`
	Location          string                           `json:"location"`
	Name              string                           `json:"name"`
	ResourceGroupName string                           `json:"resource_group_name"`
	Tags              map[string]string                `json:"tags"`
	Identity          []datashareaccount.IdentityState `json:"identity"`
	Timeouts          *datashareaccount.TimeoutsState  `json:"timeouts"`
}
