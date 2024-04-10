// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	containerconnectedregistry "github.com/golingon/terraproviders/azurerm/3.98.0/containerconnectedregistry"
	"io"
)

// NewContainerConnectedRegistry creates a new instance of [ContainerConnectedRegistry].
func NewContainerConnectedRegistry(name string, args ContainerConnectedRegistryArgs) *ContainerConnectedRegistry {
	return &ContainerConnectedRegistry{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ContainerConnectedRegistry)(nil)

// ContainerConnectedRegistry represents the Terraform resource azurerm_container_connected_registry.
type ContainerConnectedRegistry struct {
	Name      string
	Args      ContainerConnectedRegistryArgs
	state     *containerConnectedRegistryState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ContainerConnectedRegistry].
func (ccr *ContainerConnectedRegistry) Type() string {
	return "azurerm_container_connected_registry"
}

// LocalName returns the local name for [ContainerConnectedRegistry].
func (ccr *ContainerConnectedRegistry) LocalName() string {
	return ccr.Name
}

// Configuration returns the configuration (args) for [ContainerConnectedRegistry].
func (ccr *ContainerConnectedRegistry) Configuration() interface{} {
	return ccr.Args
}

// DependOn is used for other resources to depend on [ContainerConnectedRegistry].
func (ccr *ContainerConnectedRegistry) DependOn() terra.Reference {
	return terra.ReferenceResource(ccr)
}

// Dependencies returns the list of resources [ContainerConnectedRegistry] depends_on.
func (ccr *ContainerConnectedRegistry) Dependencies() terra.Dependencies {
	return ccr.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ContainerConnectedRegistry].
func (ccr *ContainerConnectedRegistry) LifecycleManagement() *terra.Lifecycle {
	return ccr.Lifecycle
}

// Attributes returns the attributes for [ContainerConnectedRegistry].
func (ccr *ContainerConnectedRegistry) Attributes() containerConnectedRegistryAttributes {
	return containerConnectedRegistryAttributes{ref: terra.ReferenceResource(ccr)}
}

// ImportState imports the given attribute values into [ContainerConnectedRegistry]'s state.
func (ccr *ContainerConnectedRegistry) ImportState(av io.Reader) error {
	ccr.state = &containerConnectedRegistryState{}
	if err := json.NewDecoder(av).Decode(ccr.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ccr.Type(), ccr.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ContainerConnectedRegistry] has state.
func (ccr *ContainerConnectedRegistry) State() (*containerConnectedRegistryState, bool) {
	return ccr.state, ccr.state != nil
}

// StateMust returns the state for [ContainerConnectedRegistry]. Panics if the state is nil.
func (ccr *ContainerConnectedRegistry) StateMust() *containerConnectedRegistryState {
	if ccr.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ccr.Type(), ccr.LocalName()))
	}
	return ccr.state
}

// ContainerConnectedRegistryArgs contains the configurations for azurerm_container_connected_registry.
type ContainerConnectedRegistryArgs struct {
	// AuditLogEnabled: bool, optional
	AuditLogEnabled terra.BoolValue `hcl:"audit_log_enabled,attr"`
	// ClientTokenIds: list of string, optional
	ClientTokenIds terra.ListValue[terra.StringValue] `hcl:"client_token_ids,attr"`
	// ContainerRegistryId: string, required
	ContainerRegistryId terra.StringValue `hcl:"container_registry_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// LogLevel: string, optional
	LogLevel terra.StringValue `hcl:"log_level,attr"`
	// Mode: string, optional
	Mode terra.StringValue `hcl:"mode,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ParentRegistryId: string, optional
	ParentRegistryId terra.StringValue `hcl:"parent_registry_id,attr"`
	// SyncMessageTtl: string, optional
	SyncMessageTtl terra.StringValue `hcl:"sync_message_ttl,attr"`
	// SyncSchedule: string, optional
	SyncSchedule terra.StringValue `hcl:"sync_schedule,attr"`
	// SyncTokenId: string, required
	SyncTokenId terra.StringValue `hcl:"sync_token_id,attr" validate:"required"`
	// SyncWindow: string, optional
	SyncWindow terra.StringValue `hcl:"sync_window,attr"`
	// Notification: min=0
	Notification []containerconnectedregistry.Notification `hcl:"notification,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *containerconnectedregistry.Timeouts `hcl:"timeouts,block"`
}
type containerConnectedRegistryAttributes struct {
	ref terra.Reference
}

// AuditLogEnabled returns a reference to field audit_log_enabled of azurerm_container_connected_registry.
func (ccr containerConnectedRegistryAttributes) AuditLogEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(ccr.ref.Append("audit_log_enabled"))
}

// ClientTokenIds returns a reference to field client_token_ids of azurerm_container_connected_registry.
func (ccr containerConnectedRegistryAttributes) ClientTokenIds() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](ccr.ref.Append("client_token_ids"))
}

// ContainerRegistryId returns a reference to field container_registry_id of azurerm_container_connected_registry.
func (ccr containerConnectedRegistryAttributes) ContainerRegistryId() terra.StringValue {
	return terra.ReferenceAsString(ccr.ref.Append("container_registry_id"))
}

// Id returns a reference to field id of azurerm_container_connected_registry.
func (ccr containerConnectedRegistryAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ccr.ref.Append("id"))
}

// LogLevel returns a reference to field log_level of azurerm_container_connected_registry.
func (ccr containerConnectedRegistryAttributes) LogLevel() terra.StringValue {
	return terra.ReferenceAsString(ccr.ref.Append("log_level"))
}

// Mode returns a reference to field mode of azurerm_container_connected_registry.
func (ccr containerConnectedRegistryAttributes) Mode() terra.StringValue {
	return terra.ReferenceAsString(ccr.ref.Append("mode"))
}

// Name returns a reference to field name of azurerm_container_connected_registry.
func (ccr containerConnectedRegistryAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ccr.ref.Append("name"))
}

// ParentRegistryId returns a reference to field parent_registry_id of azurerm_container_connected_registry.
func (ccr containerConnectedRegistryAttributes) ParentRegistryId() terra.StringValue {
	return terra.ReferenceAsString(ccr.ref.Append("parent_registry_id"))
}

// SyncMessageTtl returns a reference to field sync_message_ttl of azurerm_container_connected_registry.
func (ccr containerConnectedRegistryAttributes) SyncMessageTtl() terra.StringValue {
	return terra.ReferenceAsString(ccr.ref.Append("sync_message_ttl"))
}

// SyncSchedule returns a reference to field sync_schedule of azurerm_container_connected_registry.
func (ccr containerConnectedRegistryAttributes) SyncSchedule() terra.StringValue {
	return terra.ReferenceAsString(ccr.ref.Append("sync_schedule"))
}

// SyncTokenId returns a reference to field sync_token_id of azurerm_container_connected_registry.
func (ccr containerConnectedRegistryAttributes) SyncTokenId() terra.StringValue {
	return terra.ReferenceAsString(ccr.ref.Append("sync_token_id"))
}

// SyncWindow returns a reference to field sync_window of azurerm_container_connected_registry.
func (ccr containerConnectedRegistryAttributes) SyncWindow() terra.StringValue {
	return terra.ReferenceAsString(ccr.ref.Append("sync_window"))
}

func (ccr containerConnectedRegistryAttributes) Notification() terra.ListValue[containerconnectedregistry.NotificationAttributes] {
	return terra.ReferenceAsList[containerconnectedregistry.NotificationAttributes](ccr.ref.Append("notification"))
}

func (ccr containerConnectedRegistryAttributes) Timeouts() containerconnectedregistry.TimeoutsAttributes {
	return terra.ReferenceAsSingle[containerconnectedregistry.TimeoutsAttributes](ccr.ref.Append("timeouts"))
}

type containerConnectedRegistryState struct {
	AuditLogEnabled     bool                                           `json:"audit_log_enabled"`
	ClientTokenIds      []string                                       `json:"client_token_ids"`
	ContainerRegistryId string                                         `json:"container_registry_id"`
	Id                  string                                         `json:"id"`
	LogLevel            string                                         `json:"log_level"`
	Mode                string                                         `json:"mode"`
	Name                string                                         `json:"name"`
	ParentRegistryId    string                                         `json:"parent_registry_id"`
	SyncMessageTtl      string                                         `json:"sync_message_ttl"`
	SyncSchedule        string                                         `json:"sync_schedule"`
	SyncTokenId         string                                         `json:"sync_token_id"`
	SyncWindow          string                                         `json:"sync_window"`
	Notification        []containerconnectedregistry.NotificationState `json:"notification"`
	Timeouts            *containerconnectedregistry.TimeoutsState      `json:"timeouts"`
}
