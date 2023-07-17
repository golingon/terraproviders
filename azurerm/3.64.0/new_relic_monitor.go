// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	newrelicmonitor "github.com/golingon/terraproviders/azurerm/3.64.0/newrelicmonitor"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewNewRelicMonitor creates a new instance of [NewRelicMonitor].
func NewNewRelicMonitor(name string, args NewRelicMonitorArgs) *NewRelicMonitor {
	return &NewRelicMonitor{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*NewRelicMonitor)(nil)

// NewRelicMonitor represents the Terraform resource azurerm_new_relic_monitor.
type NewRelicMonitor struct {
	Name      string
	Args      NewRelicMonitorArgs
	state     *newRelicMonitorState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [NewRelicMonitor].
func (nrm *NewRelicMonitor) Type() string {
	return "azurerm_new_relic_monitor"
}

// LocalName returns the local name for [NewRelicMonitor].
func (nrm *NewRelicMonitor) LocalName() string {
	return nrm.Name
}

// Configuration returns the configuration (args) for [NewRelicMonitor].
func (nrm *NewRelicMonitor) Configuration() interface{} {
	return nrm.Args
}

// DependOn is used for other resources to depend on [NewRelicMonitor].
func (nrm *NewRelicMonitor) DependOn() terra.Reference {
	return terra.ReferenceResource(nrm)
}

// Dependencies returns the list of resources [NewRelicMonitor] depends_on.
func (nrm *NewRelicMonitor) Dependencies() terra.Dependencies {
	return nrm.DependsOn
}

// LifecycleManagement returns the lifecycle block for [NewRelicMonitor].
func (nrm *NewRelicMonitor) LifecycleManagement() *terra.Lifecycle {
	return nrm.Lifecycle
}

// Attributes returns the attributes for [NewRelicMonitor].
func (nrm *NewRelicMonitor) Attributes() newRelicMonitorAttributes {
	return newRelicMonitorAttributes{ref: terra.ReferenceResource(nrm)}
}

// ImportState imports the given attribute values into [NewRelicMonitor]'s state.
func (nrm *NewRelicMonitor) ImportState(av io.Reader) error {
	nrm.state = &newRelicMonitorState{}
	if err := json.NewDecoder(av).Decode(nrm.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", nrm.Type(), nrm.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [NewRelicMonitor] has state.
func (nrm *NewRelicMonitor) State() (*newRelicMonitorState, bool) {
	return nrm.state, nrm.state != nil
}

// StateMust returns the state for [NewRelicMonitor]. Panics if the state is nil.
func (nrm *NewRelicMonitor) StateMust() *newRelicMonitorState {
	if nrm.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", nrm.Type(), nrm.LocalName()))
	}
	return nrm.state
}

// NewRelicMonitorArgs contains the configurations for azurerm_new_relic_monitor.
type NewRelicMonitorArgs struct {
	// AccountCreationSource: string, optional
	AccountCreationSource terra.StringValue `hcl:"account_creation_source,attr"`
	// AccountId: string, optional
	AccountId terra.StringValue `hcl:"account_id,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IngestionKey: string, optional
	IngestionKey terra.StringValue `hcl:"ingestion_key,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// OrgCreationSource: string, optional
	OrgCreationSource terra.StringValue `hcl:"org_creation_source,attr"`
	// OrganizationId: string, optional
	OrganizationId terra.StringValue `hcl:"organization_id,attr"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// UserId: string, optional
	UserId terra.StringValue `hcl:"user_id,attr"`
	// Plan: required
	Plan *newrelicmonitor.Plan `hcl:"plan,block" validate:"required"`
	// Timeouts: optional
	Timeouts *newrelicmonitor.Timeouts `hcl:"timeouts,block"`
	// User: required
	User *newrelicmonitor.User `hcl:"user,block" validate:"required"`
}
type newRelicMonitorAttributes struct {
	ref terra.Reference
}

// AccountCreationSource returns a reference to field account_creation_source of azurerm_new_relic_monitor.
func (nrm newRelicMonitorAttributes) AccountCreationSource() terra.StringValue {
	return terra.ReferenceAsString(nrm.ref.Append("account_creation_source"))
}

// AccountId returns a reference to field account_id of azurerm_new_relic_monitor.
func (nrm newRelicMonitorAttributes) AccountId() terra.StringValue {
	return terra.ReferenceAsString(nrm.ref.Append("account_id"))
}

// Id returns a reference to field id of azurerm_new_relic_monitor.
func (nrm newRelicMonitorAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(nrm.ref.Append("id"))
}

// IngestionKey returns a reference to field ingestion_key of azurerm_new_relic_monitor.
func (nrm newRelicMonitorAttributes) IngestionKey() terra.StringValue {
	return terra.ReferenceAsString(nrm.ref.Append("ingestion_key"))
}

// Location returns a reference to field location of azurerm_new_relic_monitor.
func (nrm newRelicMonitorAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(nrm.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_new_relic_monitor.
func (nrm newRelicMonitorAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(nrm.ref.Append("name"))
}

// OrgCreationSource returns a reference to field org_creation_source of azurerm_new_relic_monitor.
func (nrm newRelicMonitorAttributes) OrgCreationSource() terra.StringValue {
	return terra.ReferenceAsString(nrm.ref.Append("org_creation_source"))
}

// OrganizationId returns a reference to field organization_id of azurerm_new_relic_monitor.
func (nrm newRelicMonitorAttributes) OrganizationId() terra.StringValue {
	return terra.ReferenceAsString(nrm.ref.Append("organization_id"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_new_relic_monitor.
func (nrm newRelicMonitorAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(nrm.ref.Append("resource_group_name"))
}

// UserId returns a reference to field user_id of azurerm_new_relic_monitor.
func (nrm newRelicMonitorAttributes) UserId() terra.StringValue {
	return terra.ReferenceAsString(nrm.ref.Append("user_id"))
}

func (nrm newRelicMonitorAttributes) Plan() terra.ListValue[newrelicmonitor.PlanAttributes] {
	return terra.ReferenceAsList[newrelicmonitor.PlanAttributes](nrm.ref.Append("plan"))
}

func (nrm newRelicMonitorAttributes) Timeouts() newrelicmonitor.TimeoutsAttributes {
	return terra.ReferenceAsSingle[newrelicmonitor.TimeoutsAttributes](nrm.ref.Append("timeouts"))
}

func (nrm newRelicMonitorAttributes) User() terra.ListValue[newrelicmonitor.UserAttributes] {
	return terra.ReferenceAsList[newrelicmonitor.UserAttributes](nrm.ref.Append("user"))
}

type newRelicMonitorState struct {
	AccountCreationSource string                         `json:"account_creation_source"`
	AccountId             string                         `json:"account_id"`
	Id                    string                         `json:"id"`
	IngestionKey          string                         `json:"ingestion_key"`
	Location              string                         `json:"location"`
	Name                  string                         `json:"name"`
	OrgCreationSource     string                         `json:"org_creation_source"`
	OrganizationId        string                         `json:"organization_id"`
	ResourceGroupName     string                         `json:"resource_group_name"`
	UserId                string                         `json:"user_id"`
	Plan                  []newrelicmonitor.PlanState    `json:"plan"`
	Timeouts              *newrelicmonitor.TimeoutsState `json:"timeouts"`
	User                  []newrelicmonitor.UserState    `json:"user"`
}