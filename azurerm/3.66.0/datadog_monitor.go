// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	datadogmonitor "github.com/golingon/terraproviders/azurerm/3.66.0/datadogmonitor"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDatadogMonitor creates a new instance of [DatadogMonitor].
func NewDatadogMonitor(name string, args DatadogMonitorArgs) *DatadogMonitor {
	return &DatadogMonitor{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DatadogMonitor)(nil)

// DatadogMonitor represents the Terraform resource azurerm_datadog_monitor.
type DatadogMonitor struct {
	Name      string
	Args      DatadogMonitorArgs
	state     *datadogMonitorState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DatadogMonitor].
func (dm *DatadogMonitor) Type() string {
	return "azurerm_datadog_monitor"
}

// LocalName returns the local name for [DatadogMonitor].
func (dm *DatadogMonitor) LocalName() string {
	return dm.Name
}

// Configuration returns the configuration (args) for [DatadogMonitor].
func (dm *DatadogMonitor) Configuration() interface{} {
	return dm.Args
}

// DependOn is used for other resources to depend on [DatadogMonitor].
func (dm *DatadogMonitor) DependOn() terra.Reference {
	return terra.ReferenceResource(dm)
}

// Dependencies returns the list of resources [DatadogMonitor] depends_on.
func (dm *DatadogMonitor) Dependencies() terra.Dependencies {
	return dm.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DatadogMonitor].
func (dm *DatadogMonitor) LifecycleManagement() *terra.Lifecycle {
	return dm.Lifecycle
}

// Attributes returns the attributes for [DatadogMonitor].
func (dm *DatadogMonitor) Attributes() datadogMonitorAttributes {
	return datadogMonitorAttributes{ref: terra.ReferenceResource(dm)}
}

// ImportState imports the given attribute values into [DatadogMonitor]'s state.
func (dm *DatadogMonitor) ImportState(av io.Reader) error {
	dm.state = &datadogMonitorState{}
	if err := json.NewDecoder(av).Decode(dm.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dm.Type(), dm.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DatadogMonitor] has state.
func (dm *DatadogMonitor) State() (*datadogMonitorState, bool) {
	return dm.state, dm.state != nil
}

// StateMust returns the state for [DatadogMonitor]. Panics if the state is nil.
func (dm *DatadogMonitor) StateMust() *datadogMonitorState {
	if dm.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dm.Type(), dm.LocalName()))
	}
	return dm.state
}

// DatadogMonitorArgs contains the configurations for azurerm_datadog_monitor.
type DatadogMonitorArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// MonitoringEnabled: bool, optional
	MonitoringEnabled terra.BoolValue `hcl:"monitoring_enabled,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// SkuName: string, required
	SkuName terra.StringValue `hcl:"sku_name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// DatadogOrganization: required
	DatadogOrganization *datadogmonitor.DatadogOrganization `hcl:"datadog_organization,block" validate:"required"`
	// Identity: optional
	Identity *datadogmonitor.Identity `hcl:"identity,block"`
	// Timeouts: optional
	Timeouts *datadogmonitor.Timeouts `hcl:"timeouts,block"`
	// User: required
	User *datadogmonitor.User `hcl:"user,block" validate:"required"`
}
type datadogMonitorAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_datadog_monitor.
func (dm datadogMonitorAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dm.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_datadog_monitor.
func (dm datadogMonitorAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(dm.ref.Append("location"))
}

// MarketplaceSubscriptionStatus returns a reference to field marketplace_subscription_status of azurerm_datadog_monitor.
func (dm datadogMonitorAttributes) MarketplaceSubscriptionStatus() terra.StringValue {
	return terra.ReferenceAsString(dm.ref.Append("marketplace_subscription_status"))
}

// MonitoringEnabled returns a reference to field monitoring_enabled of azurerm_datadog_monitor.
func (dm datadogMonitorAttributes) MonitoringEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(dm.ref.Append("monitoring_enabled"))
}

// Name returns a reference to field name of azurerm_datadog_monitor.
func (dm datadogMonitorAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(dm.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_datadog_monitor.
func (dm datadogMonitorAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(dm.ref.Append("resource_group_name"))
}

// SkuName returns a reference to field sku_name of azurerm_datadog_monitor.
func (dm datadogMonitorAttributes) SkuName() terra.StringValue {
	return terra.ReferenceAsString(dm.ref.Append("sku_name"))
}

// Tags returns a reference to field tags of azurerm_datadog_monitor.
func (dm datadogMonitorAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dm.ref.Append("tags"))
}

func (dm datadogMonitorAttributes) DatadogOrganization() terra.ListValue[datadogmonitor.DatadogOrganizationAttributes] {
	return terra.ReferenceAsList[datadogmonitor.DatadogOrganizationAttributes](dm.ref.Append("datadog_organization"))
}

func (dm datadogMonitorAttributes) Identity() terra.ListValue[datadogmonitor.IdentityAttributes] {
	return terra.ReferenceAsList[datadogmonitor.IdentityAttributes](dm.ref.Append("identity"))
}

func (dm datadogMonitorAttributes) Timeouts() datadogmonitor.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datadogmonitor.TimeoutsAttributes](dm.ref.Append("timeouts"))
}

func (dm datadogMonitorAttributes) User() terra.ListValue[datadogmonitor.UserAttributes] {
	return terra.ReferenceAsList[datadogmonitor.UserAttributes](dm.ref.Append("user"))
}

type datadogMonitorState struct {
	Id                            string                                    `json:"id"`
	Location                      string                                    `json:"location"`
	MarketplaceSubscriptionStatus string                                    `json:"marketplace_subscription_status"`
	MonitoringEnabled             bool                                      `json:"monitoring_enabled"`
	Name                          string                                    `json:"name"`
	ResourceGroupName             string                                    `json:"resource_group_name"`
	SkuName                       string                                    `json:"sku_name"`
	Tags                          map[string]string                         `json:"tags"`
	DatadogOrganization           []datadogmonitor.DatadogOrganizationState `json:"datadog_organization"`
	Identity                      []datadogmonitor.IdentityState            `json:"identity"`
	Timeouts                      *datadogmonitor.TimeoutsState             `json:"timeouts"`
	User                          []datadogmonitor.UserState                `json:"user"`
}
