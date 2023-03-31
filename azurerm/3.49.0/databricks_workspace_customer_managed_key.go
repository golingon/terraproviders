// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	databricksworkspacecustomermanagedkey "github.com/golingon/terraproviders/azurerm/3.49.0/databricksworkspacecustomermanagedkey"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewDatabricksWorkspaceCustomerManagedKey(name string, args DatabricksWorkspaceCustomerManagedKeyArgs) *DatabricksWorkspaceCustomerManagedKey {
	return &DatabricksWorkspaceCustomerManagedKey{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DatabricksWorkspaceCustomerManagedKey)(nil)

type DatabricksWorkspaceCustomerManagedKey struct {
	Name  string
	Args  DatabricksWorkspaceCustomerManagedKeyArgs
	state *databricksWorkspaceCustomerManagedKeyState
}

func (dwcmk *DatabricksWorkspaceCustomerManagedKey) Type() string {
	return "azurerm_databricks_workspace_customer_managed_key"
}

func (dwcmk *DatabricksWorkspaceCustomerManagedKey) LocalName() string {
	return dwcmk.Name
}

func (dwcmk *DatabricksWorkspaceCustomerManagedKey) Configuration() interface{} {
	return dwcmk.Args
}

func (dwcmk *DatabricksWorkspaceCustomerManagedKey) Attributes() databricksWorkspaceCustomerManagedKeyAttributes {
	return databricksWorkspaceCustomerManagedKeyAttributes{ref: terra.ReferenceResource(dwcmk)}
}

func (dwcmk *DatabricksWorkspaceCustomerManagedKey) ImportState(av io.Reader) error {
	dwcmk.state = &databricksWorkspaceCustomerManagedKeyState{}
	if err := json.NewDecoder(av).Decode(dwcmk.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dwcmk.Type(), dwcmk.LocalName(), err)
	}
	return nil
}

func (dwcmk *DatabricksWorkspaceCustomerManagedKey) State() (*databricksWorkspaceCustomerManagedKeyState, bool) {
	return dwcmk.state, dwcmk.state != nil
}

func (dwcmk *DatabricksWorkspaceCustomerManagedKey) StateMust() *databricksWorkspaceCustomerManagedKeyState {
	if dwcmk.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dwcmk.Type(), dwcmk.LocalName()))
	}
	return dwcmk.state
}

func (dwcmk *DatabricksWorkspaceCustomerManagedKey) DependOn() terra.Reference {
	return terra.ReferenceResource(dwcmk)
}

type DatabricksWorkspaceCustomerManagedKeyArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// KeyVaultKeyId: string, required
	KeyVaultKeyId terra.StringValue `hcl:"key_vault_key_id,attr" validate:"required"`
	// WorkspaceId: string, required
	WorkspaceId terra.StringValue `hcl:"workspace_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *databricksworkspacecustomermanagedkey.Timeouts `hcl:"timeouts,block"`
	// DependsOn contains resources that DatabricksWorkspaceCustomerManagedKey depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type databricksWorkspaceCustomerManagedKeyAttributes struct {
	ref terra.Reference
}

func (dwcmk databricksWorkspaceCustomerManagedKeyAttributes) Id() terra.StringValue {
	return terra.ReferenceString(dwcmk.ref.Append("id"))
}

func (dwcmk databricksWorkspaceCustomerManagedKeyAttributes) KeyVaultKeyId() terra.StringValue {
	return terra.ReferenceString(dwcmk.ref.Append("key_vault_key_id"))
}

func (dwcmk databricksWorkspaceCustomerManagedKeyAttributes) WorkspaceId() terra.StringValue {
	return terra.ReferenceString(dwcmk.ref.Append("workspace_id"))
}

func (dwcmk databricksWorkspaceCustomerManagedKeyAttributes) Timeouts() databricksworkspacecustomermanagedkey.TimeoutsAttributes {
	return terra.ReferenceSingle[databricksworkspacecustomermanagedkey.TimeoutsAttributes](dwcmk.ref.Append("timeouts"))
}

type databricksWorkspaceCustomerManagedKeyState struct {
	Id            string                                               `json:"id"`
	KeyVaultKeyId string                                               `json:"key_vault_key_id"`
	WorkspaceId   string                                               `json:"workspace_id"`
	Timeouts      *databricksworkspacecustomermanagedkey.TimeoutsState `json:"timeouts"`
}
