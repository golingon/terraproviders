// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	dataprotectionresourceguard "github.com/golingon/terraproviders/azurerm/3.63.0/dataprotectionresourceguard"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDataProtectionResourceGuard creates a new instance of [DataProtectionResourceGuard].
func NewDataProtectionResourceGuard(name string, args DataProtectionResourceGuardArgs) *DataProtectionResourceGuard {
	return &DataProtectionResourceGuard{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DataProtectionResourceGuard)(nil)

// DataProtectionResourceGuard represents the Terraform resource azurerm_data_protection_resource_guard.
type DataProtectionResourceGuard struct {
	Name      string
	Args      DataProtectionResourceGuardArgs
	state     *dataProtectionResourceGuardState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DataProtectionResourceGuard].
func (dprg *DataProtectionResourceGuard) Type() string {
	return "azurerm_data_protection_resource_guard"
}

// LocalName returns the local name for [DataProtectionResourceGuard].
func (dprg *DataProtectionResourceGuard) LocalName() string {
	return dprg.Name
}

// Configuration returns the configuration (args) for [DataProtectionResourceGuard].
func (dprg *DataProtectionResourceGuard) Configuration() interface{} {
	return dprg.Args
}

// DependOn is used for other resources to depend on [DataProtectionResourceGuard].
func (dprg *DataProtectionResourceGuard) DependOn() terra.Reference {
	return terra.ReferenceResource(dprg)
}

// Dependencies returns the list of resources [DataProtectionResourceGuard] depends_on.
func (dprg *DataProtectionResourceGuard) Dependencies() terra.Dependencies {
	return dprg.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DataProtectionResourceGuard].
func (dprg *DataProtectionResourceGuard) LifecycleManagement() *terra.Lifecycle {
	return dprg.Lifecycle
}

// Attributes returns the attributes for [DataProtectionResourceGuard].
func (dprg *DataProtectionResourceGuard) Attributes() dataProtectionResourceGuardAttributes {
	return dataProtectionResourceGuardAttributes{ref: terra.ReferenceResource(dprg)}
}

// ImportState imports the given attribute values into [DataProtectionResourceGuard]'s state.
func (dprg *DataProtectionResourceGuard) ImportState(av io.Reader) error {
	dprg.state = &dataProtectionResourceGuardState{}
	if err := json.NewDecoder(av).Decode(dprg.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dprg.Type(), dprg.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DataProtectionResourceGuard] has state.
func (dprg *DataProtectionResourceGuard) State() (*dataProtectionResourceGuardState, bool) {
	return dprg.state, dprg.state != nil
}

// StateMust returns the state for [DataProtectionResourceGuard]. Panics if the state is nil.
func (dprg *DataProtectionResourceGuard) StateMust() *dataProtectionResourceGuardState {
	if dprg.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dprg.Type(), dprg.LocalName()))
	}
	return dprg.state
}

// DataProtectionResourceGuardArgs contains the configurations for azurerm_data_protection_resource_guard.
type DataProtectionResourceGuardArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// VaultCriticalOperationExclusionList: list of string, optional
	VaultCriticalOperationExclusionList terra.ListValue[terra.StringValue] `hcl:"vault_critical_operation_exclusion_list,attr"`
	// Timeouts: optional
	Timeouts *dataprotectionresourceguard.Timeouts `hcl:"timeouts,block"`
}
type dataProtectionResourceGuardAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_data_protection_resource_guard.
func (dprg dataProtectionResourceGuardAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dprg.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_data_protection_resource_guard.
func (dprg dataProtectionResourceGuardAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(dprg.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_data_protection_resource_guard.
func (dprg dataProtectionResourceGuardAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(dprg.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_data_protection_resource_guard.
func (dprg dataProtectionResourceGuardAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(dprg.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_data_protection_resource_guard.
func (dprg dataProtectionResourceGuardAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dprg.ref.Append("tags"))
}

// VaultCriticalOperationExclusionList returns a reference to field vault_critical_operation_exclusion_list of azurerm_data_protection_resource_guard.
func (dprg dataProtectionResourceGuardAttributes) VaultCriticalOperationExclusionList() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](dprg.ref.Append("vault_critical_operation_exclusion_list"))
}

func (dprg dataProtectionResourceGuardAttributes) Timeouts() dataprotectionresourceguard.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dataprotectionresourceguard.TimeoutsAttributes](dprg.ref.Append("timeouts"))
}

type dataProtectionResourceGuardState struct {
	Id                                  string                                     `json:"id"`
	Location                            string                                     `json:"location"`
	Name                                string                                     `json:"name"`
	ResourceGroupName                   string                                     `json:"resource_group_name"`
	Tags                                map[string]string                          `json:"tags"`
	VaultCriticalOperationExclusionList []string                                   `json:"vault_critical_operation_exclusion_list"`
	Timeouts                            *dataprotectionresourceguard.TimeoutsState `json:"timeouts"`
}
