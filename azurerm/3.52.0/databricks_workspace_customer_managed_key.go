// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	databricksworkspacecustomermanagedkey "github.com/golingon/terraproviders/azurerm/3.52.0/databricksworkspacecustomermanagedkey"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDatabricksWorkspaceCustomerManagedKey creates a new instance of [DatabricksWorkspaceCustomerManagedKey].
func NewDatabricksWorkspaceCustomerManagedKey(name string, args DatabricksWorkspaceCustomerManagedKeyArgs) *DatabricksWorkspaceCustomerManagedKey {
	return &DatabricksWorkspaceCustomerManagedKey{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DatabricksWorkspaceCustomerManagedKey)(nil)

// DatabricksWorkspaceCustomerManagedKey represents the Terraform resource azurerm_databricks_workspace_customer_managed_key.
type DatabricksWorkspaceCustomerManagedKey struct {
	Name      string
	Args      DatabricksWorkspaceCustomerManagedKeyArgs
	state     *databricksWorkspaceCustomerManagedKeyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DatabricksWorkspaceCustomerManagedKey].
func (dwcmk *DatabricksWorkspaceCustomerManagedKey) Type() string {
	return "azurerm_databricks_workspace_customer_managed_key"
}

// LocalName returns the local name for [DatabricksWorkspaceCustomerManagedKey].
func (dwcmk *DatabricksWorkspaceCustomerManagedKey) LocalName() string {
	return dwcmk.Name
}

// Configuration returns the configuration (args) for [DatabricksWorkspaceCustomerManagedKey].
func (dwcmk *DatabricksWorkspaceCustomerManagedKey) Configuration() interface{} {
	return dwcmk.Args
}

// DependOn is used for other resources to depend on [DatabricksWorkspaceCustomerManagedKey].
func (dwcmk *DatabricksWorkspaceCustomerManagedKey) DependOn() terra.Reference {
	return terra.ReferenceResource(dwcmk)
}

// Dependencies returns the list of resources [DatabricksWorkspaceCustomerManagedKey] depends_on.
func (dwcmk *DatabricksWorkspaceCustomerManagedKey) Dependencies() terra.Dependencies {
	return dwcmk.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DatabricksWorkspaceCustomerManagedKey].
func (dwcmk *DatabricksWorkspaceCustomerManagedKey) LifecycleManagement() *terra.Lifecycle {
	return dwcmk.Lifecycle
}

// Attributes returns the attributes for [DatabricksWorkspaceCustomerManagedKey].
func (dwcmk *DatabricksWorkspaceCustomerManagedKey) Attributes() databricksWorkspaceCustomerManagedKeyAttributes {
	return databricksWorkspaceCustomerManagedKeyAttributes{ref: terra.ReferenceResource(dwcmk)}
}

// ImportState imports the given attribute values into [DatabricksWorkspaceCustomerManagedKey]'s state.
func (dwcmk *DatabricksWorkspaceCustomerManagedKey) ImportState(av io.Reader) error {
	dwcmk.state = &databricksWorkspaceCustomerManagedKeyState{}
	if err := json.NewDecoder(av).Decode(dwcmk.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dwcmk.Type(), dwcmk.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DatabricksWorkspaceCustomerManagedKey] has state.
func (dwcmk *DatabricksWorkspaceCustomerManagedKey) State() (*databricksWorkspaceCustomerManagedKeyState, bool) {
	return dwcmk.state, dwcmk.state != nil
}

// StateMust returns the state for [DatabricksWorkspaceCustomerManagedKey]. Panics if the state is nil.
func (dwcmk *DatabricksWorkspaceCustomerManagedKey) StateMust() *databricksWorkspaceCustomerManagedKeyState {
	if dwcmk.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dwcmk.Type(), dwcmk.LocalName()))
	}
	return dwcmk.state
}

// DatabricksWorkspaceCustomerManagedKeyArgs contains the configurations for azurerm_databricks_workspace_customer_managed_key.
type DatabricksWorkspaceCustomerManagedKeyArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// KeyVaultKeyId: string, required
	KeyVaultKeyId terra.StringValue `hcl:"key_vault_key_id,attr" validate:"required"`
	// WorkspaceId: string, required
	WorkspaceId terra.StringValue `hcl:"workspace_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *databricksworkspacecustomermanagedkey.Timeouts `hcl:"timeouts,block"`
}
type databricksWorkspaceCustomerManagedKeyAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_databricks_workspace_customer_managed_key.
func (dwcmk databricksWorkspaceCustomerManagedKeyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dwcmk.ref.Append("id"))
}

// KeyVaultKeyId returns a reference to field key_vault_key_id of azurerm_databricks_workspace_customer_managed_key.
func (dwcmk databricksWorkspaceCustomerManagedKeyAttributes) KeyVaultKeyId() terra.StringValue {
	return terra.ReferenceAsString(dwcmk.ref.Append("key_vault_key_id"))
}

// WorkspaceId returns a reference to field workspace_id of azurerm_databricks_workspace_customer_managed_key.
func (dwcmk databricksWorkspaceCustomerManagedKeyAttributes) WorkspaceId() terra.StringValue {
	return terra.ReferenceAsString(dwcmk.ref.Append("workspace_id"))
}

func (dwcmk databricksWorkspaceCustomerManagedKeyAttributes) Timeouts() databricksworkspacecustomermanagedkey.TimeoutsAttributes {
	return terra.ReferenceAsSingle[databricksworkspacecustomermanagedkey.TimeoutsAttributes](dwcmk.ref.Append("timeouts"))
}

type databricksWorkspaceCustomerManagedKeyState struct {
	Id            string                                               `json:"id"`
	KeyVaultKeyId string                                               `json:"key_vault_key_id"`
	WorkspaceId   string                                               `json:"workspace_id"`
	Timeouts      *databricksworkspacecustomermanagedkey.TimeoutsState `json:"timeouts"`
}
