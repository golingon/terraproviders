// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	virtualdesktopworkspaceapplicationgroupassociation "github.com/golingon/terraproviders/azurerm/3.65.0/virtualdesktopworkspaceapplicationgroupassociation"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewVirtualDesktopWorkspaceApplicationGroupAssociation creates a new instance of [VirtualDesktopWorkspaceApplicationGroupAssociation].
func NewVirtualDesktopWorkspaceApplicationGroupAssociation(name string, args VirtualDesktopWorkspaceApplicationGroupAssociationArgs) *VirtualDesktopWorkspaceApplicationGroupAssociation {
	return &VirtualDesktopWorkspaceApplicationGroupAssociation{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*VirtualDesktopWorkspaceApplicationGroupAssociation)(nil)

// VirtualDesktopWorkspaceApplicationGroupAssociation represents the Terraform resource azurerm_virtual_desktop_workspace_application_group_association.
type VirtualDesktopWorkspaceApplicationGroupAssociation struct {
	Name      string
	Args      VirtualDesktopWorkspaceApplicationGroupAssociationArgs
	state     *virtualDesktopWorkspaceApplicationGroupAssociationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [VirtualDesktopWorkspaceApplicationGroupAssociation].
func (vdwaga *VirtualDesktopWorkspaceApplicationGroupAssociation) Type() string {
	return "azurerm_virtual_desktop_workspace_application_group_association"
}

// LocalName returns the local name for [VirtualDesktopWorkspaceApplicationGroupAssociation].
func (vdwaga *VirtualDesktopWorkspaceApplicationGroupAssociation) LocalName() string {
	return vdwaga.Name
}

// Configuration returns the configuration (args) for [VirtualDesktopWorkspaceApplicationGroupAssociation].
func (vdwaga *VirtualDesktopWorkspaceApplicationGroupAssociation) Configuration() interface{} {
	return vdwaga.Args
}

// DependOn is used for other resources to depend on [VirtualDesktopWorkspaceApplicationGroupAssociation].
func (vdwaga *VirtualDesktopWorkspaceApplicationGroupAssociation) DependOn() terra.Reference {
	return terra.ReferenceResource(vdwaga)
}

// Dependencies returns the list of resources [VirtualDesktopWorkspaceApplicationGroupAssociation] depends_on.
func (vdwaga *VirtualDesktopWorkspaceApplicationGroupAssociation) Dependencies() terra.Dependencies {
	return vdwaga.DependsOn
}

// LifecycleManagement returns the lifecycle block for [VirtualDesktopWorkspaceApplicationGroupAssociation].
func (vdwaga *VirtualDesktopWorkspaceApplicationGroupAssociation) LifecycleManagement() *terra.Lifecycle {
	return vdwaga.Lifecycle
}

// Attributes returns the attributes for [VirtualDesktopWorkspaceApplicationGroupAssociation].
func (vdwaga *VirtualDesktopWorkspaceApplicationGroupAssociation) Attributes() virtualDesktopWorkspaceApplicationGroupAssociationAttributes {
	return virtualDesktopWorkspaceApplicationGroupAssociationAttributes{ref: terra.ReferenceResource(vdwaga)}
}

// ImportState imports the given attribute values into [VirtualDesktopWorkspaceApplicationGroupAssociation]'s state.
func (vdwaga *VirtualDesktopWorkspaceApplicationGroupAssociation) ImportState(av io.Reader) error {
	vdwaga.state = &virtualDesktopWorkspaceApplicationGroupAssociationState{}
	if err := json.NewDecoder(av).Decode(vdwaga.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", vdwaga.Type(), vdwaga.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [VirtualDesktopWorkspaceApplicationGroupAssociation] has state.
func (vdwaga *VirtualDesktopWorkspaceApplicationGroupAssociation) State() (*virtualDesktopWorkspaceApplicationGroupAssociationState, bool) {
	return vdwaga.state, vdwaga.state != nil
}

// StateMust returns the state for [VirtualDesktopWorkspaceApplicationGroupAssociation]. Panics if the state is nil.
func (vdwaga *VirtualDesktopWorkspaceApplicationGroupAssociation) StateMust() *virtualDesktopWorkspaceApplicationGroupAssociationState {
	if vdwaga.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", vdwaga.Type(), vdwaga.LocalName()))
	}
	return vdwaga.state
}

// VirtualDesktopWorkspaceApplicationGroupAssociationArgs contains the configurations for azurerm_virtual_desktop_workspace_application_group_association.
type VirtualDesktopWorkspaceApplicationGroupAssociationArgs struct {
	// ApplicationGroupId: string, required
	ApplicationGroupId terra.StringValue `hcl:"application_group_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// WorkspaceId: string, required
	WorkspaceId terra.StringValue `hcl:"workspace_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *virtualdesktopworkspaceapplicationgroupassociation.Timeouts `hcl:"timeouts,block"`
}
type virtualDesktopWorkspaceApplicationGroupAssociationAttributes struct {
	ref terra.Reference
}

// ApplicationGroupId returns a reference to field application_group_id of azurerm_virtual_desktop_workspace_application_group_association.
func (vdwaga virtualDesktopWorkspaceApplicationGroupAssociationAttributes) ApplicationGroupId() terra.StringValue {
	return terra.ReferenceAsString(vdwaga.ref.Append("application_group_id"))
}

// Id returns a reference to field id of azurerm_virtual_desktop_workspace_application_group_association.
func (vdwaga virtualDesktopWorkspaceApplicationGroupAssociationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(vdwaga.ref.Append("id"))
}

// WorkspaceId returns a reference to field workspace_id of azurerm_virtual_desktop_workspace_application_group_association.
func (vdwaga virtualDesktopWorkspaceApplicationGroupAssociationAttributes) WorkspaceId() terra.StringValue {
	return terra.ReferenceAsString(vdwaga.ref.Append("workspace_id"))
}

func (vdwaga virtualDesktopWorkspaceApplicationGroupAssociationAttributes) Timeouts() virtualdesktopworkspaceapplicationgroupassociation.TimeoutsAttributes {
	return terra.ReferenceAsSingle[virtualdesktopworkspaceapplicationgroupassociation.TimeoutsAttributes](vdwaga.ref.Append("timeouts"))
}

type virtualDesktopWorkspaceApplicationGroupAssociationState struct {
	ApplicationGroupId string                                                            `json:"application_group_id"`
	Id                 string                                                            `json:"id"`
	WorkspaceId        string                                                            `json:"workspace_id"`
	Timeouts           *virtualdesktopworkspaceapplicationgroupassociation.TimeoutsState `json:"timeouts"`
}
