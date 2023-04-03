// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	monitorprivatelinkscopedservice "github.com/golingon/terraproviders/azurerm/3.49.0/monitorprivatelinkscopedservice"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewMonitorPrivateLinkScopedService creates a new instance of [MonitorPrivateLinkScopedService].
func NewMonitorPrivateLinkScopedService(name string, args MonitorPrivateLinkScopedServiceArgs) *MonitorPrivateLinkScopedService {
	return &MonitorPrivateLinkScopedService{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*MonitorPrivateLinkScopedService)(nil)

// MonitorPrivateLinkScopedService represents the Terraform resource azurerm_monitor_private_link_scoped_service.
type MonitorPrivateLinkScopedService struct {
	Name      string
	Args      MonitorPrivateLinkScopedServiceArgs
	state     *monitorPrivateLinkScopedServiceState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [MonitorPrivateLinkScopedService].
func (mplss *MonitorPrivateLinkScopedService) Type() string {
	return "azurerm_monitor_private_link_scoped_service"
}

// LocalName returns the local name for [MonitorPrivateLinkScopedService].
func (mplss *MonitorPrivateLinkScopedService) LocalName() string {
	return mplss.Name
}

// Configuration returns the configuration (args) for [MonitorPrivateLinkScopedService].
func (mplss *MonitorPrivateLinkScopedService) Configuration() interface{} {
	return mplss.Args
}

// DependOn is used for other resources to depend on [MonitorPrivateLinkScopedService].
func (mplss *MonitorPrivateLinkScopedService) DependOn() terra.Reference {
	return terra.ReferenceResource(mplss)
}

// Dependencies returns the list of resources [MonitorPrivateLinkScopedService] depends_on.
func (mplss *MonitorPrivateLinkScopedService) Dependencies() terra.Dependencies {
	return mplss.DependsOn
}

// LifecycleManagement returns the lifecycle block for [MonitorPrivateLinkScopedService].
func (mplss *MonitorPrivateLinkScopedService) LifecycleManagement() *terra.Lifecycle {
	return mplss.Lifecycle
}

// Attributes returns the attributes for [MonitorPrivateLinkScopedService].
func (mplss *MonitorPrivateLinkScopedService) Attributes() monitorPrivateLinkScopedServiceAttributes {
	return monitorPrivateLinkScopedServiceAttributes{ref: terra.ReferenceResource(mplss)}
}

// ImportState imports the given attribute values into [MonitorPrivateLinkScopedService]'s state.
func (mplss *MonitorPrivateLinkScopedService) ImportState(av io.Reader) error {
	mplss.state = &monitorPrivateLinkScopedServiceState{}
	if err := json.NewDecoder(av).Decode(mplss.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", mplss.Type(), mplss.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [MonitorPrivateLinkScopedService] has state.
func (mplss *MonitorPrivateLinkScopedService) State() (*monitorPrivateLinkScopedServiceState, bool) {
	return mplss.state, mplss.state != nil
}

// StateMust returns the state for [MonitorPrivateLinkScopedService]. Panics if the state is nil.
func (mplss *MonitorPrivateLinkScopedService) StateMust() *monitorPrivateLinkScopedServiceState {
	if mplss.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", mplss.Type(), mplss.LocalName()))
	}
	return mplss.state
}

// MonitorPrivateLinkScopedServiceArgs contains the configurations for azurerm_monitor_private_link_scoped_service.
type MonitorPrivateLinkScopedServiceArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// LinkedResourceId: string, required
	LinkedResourceId terra.StringValue `hcl:"linked_resource_id,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// ScopeName: string, required
	ScopeName terra.StringValue `hcl:"scope_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *monitorprivatelinkscopedservice.Timeouts `hcl:"timeouts,block"`
}
type monitorPrivateLinkScopedServiceAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_monitor_private_link_scoped_service.
func (mplss monitorPrivateLinkScopedServiceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(mplss.ref.Append("id"))
}

// LinkedResourceId returns a reference to field linked_resource_id of azurerm_monitor_private_link_scoped_service.
func (mplss monitorPrivateLinkScopedServiceAttributes) LinkedResourceId() terra.StringValue {
	return terra.ReferenceAsString(mplss.ref.Append("linked_resource_id"))
}

// Name returns a reference to field name of azurerm_monitor_private_link_scoped_service.
func (mplss monitorPrivateLinkScopedServiceAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(mplss.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_monitor_private_link_scoped_service.
func (mplss monitorPrivateLinkScopedServiceAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(mplss.ref.Append("resource_group_name"))
}

// ScopeName returns a reference to field scope_name of azurerm_monitor_private_link_scoped_service.
func (mplss monitorPrivateLinkScopedServiceAttributes) ScopeName() terra.StringValue {
	return terra.ReferenceAsString(mplss.ref.Append("scope_name"))
}

func (mplss monitorPrivateLinkScopedServiceAttributes) Timeouts() monitorprivatelinkscopedservice.TimeoutsAttributes {
	return terra.ReferenceAsSingle[monitorprivatelinkscopedservice.TimeoutsAttributes](mplss.ref.Append("timeouts"))
}

type monitorPrivateLinkScopedServiceState struct {
	Id                string                                         `json:"id"`
	LinkedResourceId  string                                         `json:"linked_resource_id"`
	Name              string                                         `json:"name"`
	ResourceGroupName string                                         `json:"resource_group_name"`
	ScopeName         string                                         `json:"scope_name"`
	Timeouts          *monitorprivatelinkscopedservice.TimeoutsState `json:"timeouts"`
}
