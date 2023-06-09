// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	loganalyticslinkedstorageaccount "github.com/golingon/terraproviders/azurerm/3.55.0/loganalyticslinkedstorageaccount"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewLogAnalyticsLinkedStorageAccount creates a new instance of [LogAnalyticsLinkedStorageAccount].
func NewLogAnalyticsLinkedStorageAccount(name string, args LogAnalyticsLinkedStorageAccountArgs) *LogAnalyticsLinkedStorageAccount {
	return &LogAnalyticsLinkedStorageAccount{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*LogAnalyticsLinkedStorageAccount)(nil)

// LogAnalyticsLinkedStorageAccount represents the Terraform resource azurerm_log_analytics_linked_storage_account.
type LogAnalyticsLinkedStorageAccount struct {
	Name      string
	Args      LogAnalyticsLinkedStorageAccountArgs
	state     *logAnalyticsLinkedStorageAccountState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [LogAnalyticsLinkedStorageAccount].
func (lalsa *LogAnalyticsLinkedStorageAccount) Type() string {
	return "azurerm_log_analytics_linked_storage_account"
}

// LocalName returns the local name for [LogAnalyticsLinkedStorageAccount].
func (lalsa *LogAnalyticsLinkedStorageAccount) LocalName() string {
	return lalsa.Name
}

// Configuration returns the configuration (args) for [LogAnalyticsLinkedStorageAccount].
func (lalsa *LogAnalyticsLinkedStorageAccount) Configuration() interface{} {
	return lalsa.Args
}

// DependOn is used for other resources to depend on [LogAnalyticsLinkedStorageAccount].
func (lalsa *LogAnalyticsLinkedStorageAccount) DependOn() terra.Reference {
	return terra.ReferenceResource(lalsa)
}

// Dependencies returns the list of resources [LogAnalyticsLinkedStorageAccount] depends_on.
func (lalsa *LogAnalyticsLinkedStorageAccount) Dependencies() terra.Dependencies {
	return lalsa.DependsOn
}

// LifecycleManagement returns the lifecycle block for [LogAnalyticsLinkedStorageAccount].
func (lalsa *LogAnalyticsLinkedStorageAccount) LifecycleManagement() *terra.Lifecycle {
	return lalsa.Lifecycle
}

// Attributes returns the attributes for [LogAnalyticsLinkedStorageAccount].
func (lalsa *LogAnalyticsLinkedStorageAccount) Attributes() logAnalyticsLinkedStorageAccountAttributes {
	return logAnalyticsLinkedStorageAccountAttributes{ref: terra.ReferenceResource(lalsa)}
}

// ImportState imports the given attribute values into [LogAnalyticsLinkedStorageAccount]'s state.
func (lalsa *LogAnalyticsLinkedStorageAccount) ImportState(av io.Reader) error {
	lalsa.state = &logAnalyticsLinkedStorageAccountState{}
	if err := json.NewDecoder(av).Decode(lalsa.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", lalsa.Type(), lalsa.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [LogAnalyticsLinkedStorageAccount] has state.
func (lalsa *LogAnalyticsLinkedStorageAccount) State() (*logAnalyticsLinkedStorageAccountState, bool) {
	return lalsa.state, lalsa.state != nil
}

// StateMust returns the state for [LogAnalyticsLinkedStorageAccount]. Panics if the state is nil.
func (lalsa *LogAnalyticsLinkedStorageAccount) StateMust() *logAnalyticsLinkedStorageAccountState {
	if lalsa.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", lalsa.Type(), lalsa.LocalName()))
	}
	return lalsa.state
}

// LogAnalyticsLinkedStorageAccountArgs contains the configurations for azurerm_log_analytics_linked_storage_account.
type LogAnalyticsLinkedStorageAccountArgs struct {
	// DataSourceType: string, required
	DataSourceType terra.StringValue `hcl:"data_source_type,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// StorageAccountIds: set of string, required
	StorageAccountIds terra.SetValue[terra.StringValue] `hcl:"storage_account_ids,attr" validate:"required"`
	// WorkspaceResourceId: string, required
	WorkspaceResourceId terra.StringValue `hcl:"workspace_resource_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *loganalyticslinkedstorageaccount.Timeouts `hcl:"timeouts,block"`
}
type logAnalyticsLinkedStorageAccountAttributes struct {
	ref terra.Reference
}

// DataSourceType returns a reference to field data_source_type of azurerm_log_analytics_linked_storage_account.
func (lalsa logAnalyticsLinkedStorageAccountAttributes) DataSourceType() terra.StringValue {
	return terra.ReferenceAsString(lalsa.ref.Append("data_source_type"))
}

// Id returns a reference to field id of azurerm_log_analytics_linked_storage_account.
func (lalsa logAnalyticsLinkedStorageAccountAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(lalsa.ref.Append("id"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_log_analytics_linked_storage_account.
func (lalsa logAnalyticsLinkedStorageAccountAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(lalsa.ref.Append("resource_group_name"))
}

// StorageAccountIds returns a reference to field storage_account_ids of azurerm_log_analytics_linked_storage_account.
func (lalsa logAnalyticsLinkedStorageAccountAttributes) StorageAccountIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](lalsa.ref.Append("storage_account_ids"))
}

// WorkspaceResourceId returns a reference to field workspace_resource_id of azurerm_log_analytics_linked_storage_account.
func (lalsa logAnalyticsLinkedStorageAccountAttributes) WorkspaceResourceId() terra.StringValue {
	return terra.ReferenceAsString(lalsa.ref.Append("workspace_resource_id"))
}

func (lalsa logAnalyticsLinkedStorageAccountAttributes) Timeouts() loganalyticslinkedstorageaccount.TimeoutsAttributes {
	return terra.ReferenceAsSingle[loganalyticslinkedstorageaccount.TimeoutsAttributes](lalsa.ref.Append("timeouts"))
}

type logAnalyticsLinkedStorageAccountState struct {
	DataSourceType      string                                          `json:"data_source_type"`
	Id                  string                                          `json:"id"`
	ResourceGroupName   string                                          `json:"resource_group_name"`
	StorageAccountIds   []string                                        `json:"storage_account_ids"`
	WorkspaceResourceId string                                          `json:"workspace_resource_id"`
	Timeouts            *loganalyticslinkedstorageaccount.TimeoutsState `json:"timeouts"`
}
