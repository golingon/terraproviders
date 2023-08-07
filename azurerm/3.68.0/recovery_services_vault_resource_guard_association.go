// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	recoveryservicesvaultresourceguardassociation "github.com/golingon/terraproviders/azurerm/3.68.0/recoveryservicesvaultresourceguardassociation"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewRecoveryServicesVaultResourceGuardAssociation creates a new instance of [RecoveryServicesVaultResourceGuardAssociation].
func NewRecoveryServicesVaultResourceGuardAssociation(name string, args RecoveryServicesVaultResourceGuardAssociationArgs) *RecoveryServicesVaultResourceGuardAssociation {
	return &RecoveryServicesVaultResourceGuardAssociation{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*RecoveryServicesVaultResourceGuardAssociation)(nil)

// RecoveryServicesVaultResourceGuardAssociation represents the Terraform resource azurerm_recovery_services_vault_resource_guard_association.
type RecoveryServicesVaultResourceGuardAssociation struct {
	Name      string
	Args      RecoveryServicesVaultResourceGuardAssociationArgs
	state     *recoveryServicesVaultResourceGuardAssociationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [RecoveryServicesVaultResourceGuardAssociation].
func (rsvrga *RecoveryServicesVaultResourceGuardAssociation) Type() string {
	return "azurerm_recovery_services_vault_resource_guard_association"
}

// LocalName returns the local name for [RecoveryServicesVaultResourceGuardAssociation].
func (rsvrga *RecoveryServicesVaultResourceGuardAssociation) LocalName() string {
	return rsvrga.Name
}

// Configuration returns the configuration (args) for [RecoveryServicesVaultResourceGuardAssociation].
func (rsvrga *RecoveryServicesVaultResourceGuardAssociation) Configuration() interface{} {
	return rsvrga.Args
}

// DependOn is used for other resources to depend on [RecoveryServicesVaultResourceGuardAssociation].
func (rsvrga *RecoveryServicesVaultResourceGuardAssociation) DependOn() terra.Reference {
	return terra.ReferenceResource(rsvrga)
}

// Dependencies returns the list of resources [RecoveryServicesVaultResourceGuardAssociation] depends_on.
func (rsvrga *RecoveryServicesVaultResourceGuardAssociation) Dependencies() terra.Dependencies {
	return rsvrga.DependsOn
}

// LifecycleManagement returns the lifecycle block for [RecoveryServicesVaultResourceGuardAssociation].
func (rsvrga *RecoveryServicesVaultResourceGuardAssociation) LifecycleManagement() *terra.Lifecycle {
	return rsvrga.Lifecycle
}

// Attributes returns the attributes for [RecoveryServicesVaultResourceGuardAssociation].
func (rsvrga *RecoveryServicesVaultResourceGuardAssociation) Attributes() recoveryServicesVaultResourceGuardAssociationAttributes {
	return recoveryServicesVaultResourceGuardAssociationAttributes{ref: terra.ReferenceResource(rsvrga)}
}

// ImportState imports the given attribute values into [RecoveryServicesVaultResourceGuardAssociation]'s state.
func (rsvrga *RecoveryServicesVaultResourceGuardAssociation) ImportState(av io.Reader) error {
	rsvrga.state = &recoveryServicesVaultResourceGuardAssociationState{}
	if err := json.NewDecoder(av).Decode(rsvrga.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", rsvrga.Type(), rsvrga.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [RecoveryServicesVaultResourceGuardAssociation] has state.
func (rsvrga *RecoveryServicesVaultResourceGuardAssociation) State() (*recoveryServicesVaultResourceGuardAssociationState, bool) {
	return rsvrga.state, rsvrga.state != nil
}

// StateMust returns the state for [RecoveryServicesVaultResourceGuardAssociation]. Panics if the state is nil.
func (rsvrga *RecoveryServicesVaultResourceGuardAssociation) StateMust() *recoveryServicesVaultResourceGuardAssociationState {
	if rsvrga.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", rsvrga.Type(), rsvrga.LocalName()))
	}
	return rsvrga.state
}

// RecoveryServicesVaultResourceGuardAssociationArgs contains the configurations for azurerm_recovery_services_vault_resource_guard_association.
type RecoveryServicesVaultResourceGuardAssociationArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGuardId: string, required
	ResourceGuardId terra.StringValue `hcl:"resource_guard_id,attr" validate:"required"`
	// VaultId: string, required
	VaultId terra.StringValue `hcl:"vault_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *recoveryservicesvaultresourceguardassociation.Timeouts `hcl:"timeouts,block"`
}
type recoveryServicesVaultResourceGuardAssociationAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_recovery_services_vault_resource_guard_association.
func (rsvrga recoveryServicesVaultResourceGuardAssociationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(rsvrga.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_recovery_services_vault_resource_guard_association.
func (rsvrga recoveryServicesVaultResourceGuardAssociationAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(rsvrga.ref.Append("name"))
}

// ResourceGuardId returns a reference to field resource_guard_id of azurerm_recovery_services_vault_resource_guard_association.
func (rsvrga recoveryServicesVaultResourceGuardAssociationAttributes) ResourceGuardId() terra.StringValue {
	return terra.ReferenceAsString(rsvrga.ref.Append("resource_guard_id"))
}

// VaultId returns a reference to field vault_id of azurerm_recovery_services_vault_resource_guard_association.
func (rsvrga recoveryServicesVaultResourceGuardAssociationAttributes) VaultId() terra.StringValue {
	return terra.ReferenceAsString(rsvrga.ref.Append("vault_id"))
}

func (rsvrga recoveryServicesVaultResourceGuardAssociationAttributes) Timeouts() recoveryservicesvaultresourceguardassociation.TimeoutsAttributes {
	return terra.ReferenceAsSingle[recoveryservicesvaultresourceguardassociation.TimeoutsAttributes](rsvrga.ref.Append("timeouts"))
}

type recoveryServicesVaultResourceGuardAssociationState struct {
	Id              string                                                       `json:"id"`
	Name            string                                                       `json:"name"`
	ResourceGuardId string                                                       `json:"resource_guard_id"`
	VaultId         string                                                       `json:"vault_id"`
	Timeouts        *recoveryservicesvaultresourceguardassociation.TimeoutsState `json:"timeouts"`
}
