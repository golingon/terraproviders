// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	mssqlvirtualmachinegroup "github.com/golingon/terraproviders/azurerm/3.98.0/mssqlvirtualmachinegroup"
	"io"
)

// NewMssqlVirtualMachineGroup creates a new instance of [MssqlVirtualMachineGroup].
func NewMssqlVirtualMachineGroup(name string, args MssqlVirtualMachineGroupArgs) *MssqlVirtualMachineGroup {
	return &MssqlVirtualMachineGroup{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*MssqlVirtualMachineGroup)(nil)

// MssqlVirtualMachineGroup represents the Terraform resource azurerm_mssql_virtual_machine_group.
type MssqlVirtualMachineGroup struct {
	Name      string
	Args      MssqlVirtualMachineGroupArgs
	state     *mssqlVirtualMachineGroupState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [MssqlVirtualMachineGroup].
func (mvmg *MssqlVirtualMachineGroup) Type() string {
	return "azurerm_mssql_virtual_machine_group"
}

// LocalName returns the local name for [MssqlVirtualMachineGroup].
func (mvmg *MssqlVirtualMachineGroup) LocalName() string {
	return mvmg.Name
}

// Configuration returns the configuration (args) for [MssqlVirtualMachineGroup].
func (mvmg *MssqlVirtualMachineGroup) Configuration() interface{} {
	return mvmg.Args
}

// DependOn is used for other resources to depend on [MssqlVirtualMachineGroup].
func (mvmg *MssqlVirtualMachineGroup) DependOn() terra.Reference {
	return terra.ReferenceResource(mvmg)
}

// Dependencies returns the list of resources [MssqlVirtualMachineGroup] depends_on.
func (mvmg *MssqlVirtualMachineGroup) Dependencies() terra.Dependencies {
	return mvmg.DependsOn
}

// LifecycleManagement returns the lifecycle block for [MssqlVirtualMachineGroup].
func (mvmg *MssqlVirtualMachineGroup) LifecycleManagement() *terra.Lifecycle {
	return mvmg.Lifecycle
}

// Attributes returns the attributes for [MssqlVirtualMachineGroup].
func (mvmg *MssqlVirtualMachineGroup) Attributes() mssqlVirtualMachineGroupAttributes {
	return mssqlVirtualMachineGroupAttributes{ref: terra.ReferenceResource(mvmg)}
}

// ImportState imports the given attribute values into [MssqlVirtualMachineGroup]'s state.
func (mvmg *MssqlVirtualMachineGroup) ImportState(av io.Reader) error {
	mvmg.state = &mssqlVirtualMachineGroupState{}
	if err := json.NewDecoder(av).Decode(mvmg.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", mvmg.Type(), mvmg.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [MssqlVirtualMachineGroup] has state.
func (mvmg *MssqlVirtualMachineGroup) State() (*mssqlVirtualMachineGroupState, bool) {
	return mvmg.state, mvmg.state != nil
}

// StateMust returns the state for [MssqlVirtualMachineGroup]. Panics if the state is nil.
func (mvmg *MssqlVirtualMachineGroup) StateMust() *mssqlVirtualMachineGroupState {
	if mvmg.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", mvmg.Type(), mvmg.LocalName()))
	}
	return mvmg.state
}

// MssqlVirtualMachineGroupArgs contains the configurations for azurerm_mssql_virtual_machine_group.
type MssqlVirtualMachineGroupArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// SqlImageOffer: string, required
	SqlImageOffer terra.StringValue `hcl:"sql_image_offer,attr" validate:"required"`
	// SqlImageSku: string, required
	SqlImageSku terra.StringValue `hcl:"sql_image_sku,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Timeouts: optional
	Timeouts *mssqlvirtualmachinegroup.Timeouts `hcl:"timeouts,block"`
	// WsfcDomainProfile: required
	WsfcDomainProfile *mssqlvirtualmachinegroup.WsfcDomainProfile `hcl:"wsfc_domain_profile,block" validate:"required"`
}
type mssqlVirtualMachineGroupAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_mssql_virtual_machine_group.
func (mvmg mssqlVirtualMachineGroupAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(mvmg.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_mssql_virtual_machine_group.
func (mvmg mssqlVirtualMachineGroupAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(mvmg.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_mssql_virtual_machine_group.
func (mvmg mssqlVirtualMachineGroupAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(mvmg.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_mssql_virtual_machine_group.
func (mvmg mssqlVirtualMachineGroupAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(mvmg.ref.Append("resource_group_name"))
}

// SqlImageOffer returns a reference to field sql_image_offer of azurerm_mssql_virtual_machine_group.
func (mvmg mssqlVirtualMachineGroupAttributes) SqlImageOffer() terra.StringValue {
	return terra.ReferenceAsString(mvmg.ref.Append("sql_image_offer"))
}

// SqlImageSku returns a reference to field sql_image_sku of azurerm_mssql_virtual_machine_group.
func (mvmg mssqlVirtualMachineGroupAttributes) SqlImageSku() terra.StringValue {
	return terra.ReferenceAsString(mvmg.ref.Append("sql_image_sku"))
}

// Tags returns a reference to field tags of azurerm_mssql_virtual_machine_group.
func (mvmg mssqlVirtualMachineGroupAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](mvmg.ref.Append("tags"))
}

func (mvmg mssqlVirtualMachineGroupAttributes) Timeouts() mssqlvirtualmachinegroup.TimeoutsAttributes {
	return terra.ReferenceAsSingle[mssqlvirtualmachinegroup.TimeoutsAttributes](mvmg.ref.Append("timeouts"))
}

func (mvmg mssqlVirtualMachineGroupAttributes) WsfcDomainProfile() terra.ListValue[mssqlvirtualmachinegroup.WsfcDomainProfileAttributes] {
	return terra.ReferenceAsList[mssqlvirtualmachinegroup.WsfcDomainProfileAttributes](mvmg.ref.Append("wsfc_domain_profile"))
}

type mssqlVirtualMachineGroupState struct {
	Id                string                                            `json:"id"`
	Location          string                                            `json:"location"`
	Name              string                                            `json:"name"`
	ResourceGroupName string                                            `json:"resource_group_name"`
	SqlImageOffer     string                                            `json:"sql_image_offer"`
	SqlImageSku       string                                            `json:"sql_image_sku"`
	Tags              map[string]string                                 `json:"tags"`
	Timeouts          *mssqlvirtualmachinegroup.TimeoutsState           `json:"timeouts"`
	WsfcDomainProfile []mssqlvirtualmachinegroup.WsfcDomainProfileState `json:"wsfc_domain_profile"`
}
