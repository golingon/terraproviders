// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	graphaccount "github.com/golingon/terraproviders/azurerm/3.69.0/graphaccount"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewGraphAccount creates a new instance of [GraphAccount].
func NewGraphAccount(name string, args GraphAccountArgs) *GraphAccount {
	return &GraphAccount{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*GraphAccount)(nil)

// GraphAccount represents the Terraform resource azurerm_graph_account.
type GraphAccount struct {
	Name      string
	Args      GraphAccountArgs
	state     *graphAccountState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [GraphAccount].
func (ga *GraphAccount) Type() string {
	return "azurerm_graph_account"
}

// LocalName returns the local name for [GraphAccount].
func (ga *GraphAccount) LocalName() string {
	return ga.Name
}

// Configuration returns the configuration (args) for [GraphAccount].
func (ga *GraphAccount) Configuration() interface{} {
	return ga.Args
}

// DependOn is used for other resources to depend on [GraphAccount].
func (ga *GraphAccount) DependOn() terra.Reference {
	return terra.ReferenceResource(ga)
}

// Dependencies returns the list of resources [GraphAccount] depends_on.
func (ga *GraphAccount) Dependencies() terra.Dependencies {
	return ga.DependsOn
}

// LifecycleManagement returns the lifecycle block for [GraphAccount].
func (ga *GraphAccount) LifecycleManagement() *terra.Lifecycle {
	return ga.Lifecycle
}

// Attributes returns the attributes for [GraphAccount].
func (ga *GraphAccount) Attributes() graphAccountAttributes {
	return graphAccountAttributes{ref: terra.ReferenceResource(ga)}
}

// ImportState imports the given attribute values into [GraphAccount]'s state.
func (ga *GraphAccount) ImportState(av io.Reader) error {
	ga.state = &graphAccountState{}
	if err := json.NewDecoder(av).Decode(ga.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ga.Type(), ga.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [GraphAccount] has state.
func (ga *GraphAccount) State() (*graphAccountState, bool) {
	return ga.state, ga.state != nil
}

// StateMust returns the state for [GraphAccount]. Panics if the state is nil.
func (ga *GraphAccount) StateMust() *graphAccountState {
	if ga.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ga.Type(), ga.LocalName()))
	}
	return ga.state
}

// GraphAccountArgs contains the configurations for azurerm_graph_account.
type GraphAccountArgs struct {
	// ApplicationId: string, required
	ApplicationId terra.StringValue `hcl:"application_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Timeouts: optional
	Timeouts *graphaccount.Timeouts `hcl:"timeouts,block"`
}
type graphAccountAttributes struct {
	ref terra.Reference
}

// ApplicationId returns a reference to field application_id of azurerm_graph_account.
func (ga graphAccountAttributes) ApplicationId() terra.StringValue {
	return terra.ReferenceAsString(ga.ref.Append("application_id"))
}

// BillingPlanId returns a reference to field billing_plan_id of azurerm_graph_account.
func (ga graphAccountAttributes) BillingPlanId() terra.StringValue {
	return terra.ReferenceAsString(ga.ref.Append("billing_plan_id"))
}

// Id returns a reference to field id of azurerm_graph_account.
func (ga graphAccountAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ga.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_graph_account.
func (ga graphAccountAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ga.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_graph_account.
func (ga graphAccountAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(ga.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_graph_account.
func (ga graphAccountAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ga.ref.Append("tags"))
}

func (ga graphAccountAttributes) Timeouts() graphaccount.TimeoutsAttributes {
	return terra.ReferenceAsSingle[graphaccount.TimeoutsAttributes](ga.ref.Append("timeouts"))
}

type graphAccountState struct {
	ApplicationId     string                      `json:"application_id"`
	BillingPlanId     string                      `json:"billing_plan_id"`
	Id                string                      `json:"id"`
	Name              string                      `json:"name"`
	ResourceGroupName string                      `json:"resource_group_name"`
	Tags              map[string]string           `json:"tags"`
	Timeouts          *graphaccount.TimeoutsState `json:"timeouts"`
}
