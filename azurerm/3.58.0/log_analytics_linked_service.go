// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	loganalyticslinkedservice "github.com/golingon/terraproviders/azurerm/3.58.0/loganalyticslinkedservice"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewLogAnalyticsLinkedService creates a new instance of [LogAnalyticsLinkedService].
func NewLogAnalyticsLinkedService(name string, args LogAnalyticsLinkedServiceArgs) *LogAnalyticsLinkedService {
	return &LogAnalyticsLinkedService{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*LogAnalyticsLinkedService)(nil)

// LogAnalyticsLinkedService represents the Terraform resource azurerm_log_analytics_linked_service.
type LogAnalyticsLinkedService struct {
	Name      string
	Args      LogAnalyticsLinkedServiceArgs
	state     *logAnalyticsLinkedServiceState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [LogAnalyticsLinkedService].
func (lals *LogAnalyticsLinkedService) Type() string {
	return "azurerm_log_analytics_linked_service"
}

// LocalName returns the local name for [LogAnalyticsLinkedService].
func (lals *LogAnalyticsLinkedService) LocalName() string {
	return lals.Name
}

// Configuration returns the configuration (args) for [LogAnalyticsLinkedService].
func (lals *LogAnalyticsLinkedService) Configuration() interface{} {
	return lals.Args
}

// DependOn is used for other resources to depend on [LogAnalyticsLinkedService].
func (lals *LogAnalyticsLinkedService) DependOn() terra.Reference {
	return terra.ReferenceResource(lals)
}

// Dependencies returns the list of resources [LogAnalyticsLinkedService] depends_on.
func (lals *LogAnalyticsLinkedService) Dependencies() terra.Dependencies {
	return lals.DependsOn
}

// LifecycleManagement returns the lifecycle block for [LogAnalyticsLinkedService].
func (lals *LogAnalyticsLinkedService) LifecycleManagement() *terra.Lifecycle {
	return lals.Lifecycle
}

// Attributes returns the attributes for [LogAnalyticsLinkedService].
func (lals *LogAnalyticsLinkedService) Attributes() logAnalyticsLinkedServiceAttributes {
	return logAnalyticsLinkedServiceAttributes{ref: terra.ReferenceResource(lals)}
}

// ImportState imports the given attribute values into [LogAnalyticsLinkedService]'s state.
func (lals *LogAnalyticsLinkedService) ImportState(av io.Reader) error {
	lals.state = &logAnalyticsLinkedServiceState{}
	if err := json.NewDecoder(av).Decode(lals.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", lals.Type(), lals.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [LogAnalyticsLinkedService] has state.
func (lals *LogAnalyticsLinkedService) State() (*logAnalyticsLinkedServiceState, bool) {
	return lals.state, lals.state != nil
}

// StateMust returns the state for [LogAnalyticsLinkedService]. Panics if the state is nil.
func (lals *LogAnalyticsLinkedService) StateMust() *logAnalyticsLinkedServiceState {
	if lals.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", lals.Type(), lals.LocalName()))
	}
	return lals.state
}

// LogAnalyticsLinkedServiceArgs contains the configurations for azurerm_log_analytics_linked_service.
type LogAnalyticsLinkedServiceArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ReadAccessId: string, optional
	ReadAccessId terra.StringValue `hcl:"read_access_id,attr"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// WorkspaceId: string, required
	WorkspaceId terra.StringValue `hcl:"workspace_id,attr" validate:"required"`
	// WriteAccessId: string, optional
	WriteAccessId terra.StringValue `hcl:"write_access_id,attr"`
	// Timeouts: optional
	Timeouts *loganalyticslinkedservice.Timeouts `hcl:"timeouts,block"`
}
type logAnalyticsLinkedServiceAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_log_analytics_linked_service.
func (lals logAnalyticsLinkedServiceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(lals.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_log_analytics_linked_service.
func (lals logAnalyticsLinkedServiceAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(lals.ref.Append("name"))
}

// ReadAccessId returns a reference to field read_access_id of azurerm_log_analytics_linked_service.
func (lals logAnalyticsLinkedServiceAttributes) ReadAccessId() terra.StringValue {
	return terra.ReferenceAsString(lals.ref.Append("read_access_id"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_log_analytics_linked_service.
func (lals logAnalyticsLinkedServiceAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(lals.ref.Append("resource_group_name"))
}

// WorkspaceId returns a reference to field workspace_id of azurerm_log_analytics_linked_service.
func (lals logAnalyticsLinkedServiceAttributes) WorkspaceId() terra.StringValue {
	return terra.ReferenceAsString(lals.ref.Append("workspace_id"))
}

// WriteAccessId returns a reference to field write_access_id of azurerm_log_analytics_linked_service.
func (lals logAnalyticsLinkedServiceAttributes) WriteAccessId() terra.StringValue {
	return terra.ReferenceAsString(lals.ref.Append("write_access_id"))
}

func (lals logAnalyticsLinkedServiceAttributes) Timeouts() loganalyticslinkedservice.TimeoutsAttributes {
	return terra.ReferenceAsSingle[loganalyticslinkedservice.TimeoutsAttributes](lals.ref.Append("timeouts"))
}

type logAnalyticsLinkedServiceState struct {
	Id                string                                   `json:"id"`
	Name              string                                   `json:"name"`
	ReadAccessId      string                                   `json:"read_access_id"`
	ResourceGroupName string                                   `json:"resource_group_name"`
	WorkspaceId       string                                   `json:"workspace_id"`
	WriteAccessId     string                                   `json:"write_access_id"`
	Timeouts          *loganalyticslinkedservice.TimeoutsState `json:"timeouts"`
}
