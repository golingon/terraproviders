// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	kustodatabaseprincipalassignment "github.com/golingon/terraproviders/azurerm/3.58.0/kustodatabaseprincipalassignment"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewKustoDatabasePrincipalAssignment creates a new instance of [KustoDatabasePrincipalAssignment].
func NewKustoDatabasePrincipalAssignment(name string, args KustoDatabasePrincipalAssignmentArgs) *KustoDatabasePrincipalAssignment {
	return &KustoDatabasePrincipalAssignment{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*KustoDatabasePrincipalAssignment)(nil)

// KustoDatabasePrincipalAssignment represents the Terraform resource azurerm_kusto_database_principal_assignment.
type KustoDatabasePrincipalAssignment struct {
	Name      string
	Args      KustoDatabasePrincipalAssignmentArgs
	state     *kustoDatabasePrincipalAssignmentState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [KustoDatabasePrincipalAssignment].
func (kdpa *KustoDatabasePrincipalAssignment) Type() string {
	return "azurerm_kusto_database_principal_assignment"
}

// LocalName returns the local name for [KustoDatabasePrincipalAssignment].
func (kdpa *KustoDatabasePrincipalAssignment) LocalName() string {
	return kdpa.Name
}

// Configuration returns the configuration (args) for [KustoDatabasePrincipalAssignment].
func (kdpa *KustoDatabasePrincipalAssignment) Configuration() interface{} {
	return kdpa.Args
}

// DependOn is used for other resources to depend on [KustoDatabasePrincipalAssignment].
func (kdpa *KustoDatabasePrincipalAssignment) DependOn() terra.Reference {
	return terra.ReferenceResource(kdpa)
}

// Dependencies returns the list of resources [KustoDatabasePrincipalAssignment] depends_on.
func (kdpa *KustoDatabasePrincipalAssignment) Dependencies() terra.Dependencies {
	return kdpa.DependsOn
}

// LifecycleManagement returns the lifecycle block for [KustoDatabasePrincipalAssignment].
func (kdpa *KustoDatabasePrincipalAssignment) LifecycleManagement() *terra.Lifecycle {
	return kdpa.Lifecycle
}

// Attributes returns the attributes for [KustoDatabasePrincipalAssignment].
func (kdpa *KustoDatabasePrincipalAssignment) Attributes() kustoDatabasePrincipalAssignmentAttributes {
	return kustoDatabasePrincipalAssignmentAttributes{ref: terra.ReferenceResource(kdpa)}
}

// ImportState imports the given attribute values into [KustoDatabasePrincipalAssignment]'s state.
func (kdpa *KustoDatabasePrincipalAssignment) ImportState(av io.Reader) error {
	kdpa.state = &kustoDatabasePrincipalAssignmentState{}
	if err := json.NewDecoder(av).Decode(kdpa.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", kdpa.Type(), kdpa.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [KustoDatabasePrincipalAssignment] has state.
func (kdpa *KustoDatabasePrincipalAssignment) State() (*kustoDatabasePrincipalAssignmentState, bool) {
	return kdpa.state, kdpa.state != nil
}

// StateMust returns the state for [KustoDatabasePrincipalAssignment]. Panics if the state is nil.
func (kdpa *KustoDatabasePrincipalAssignment) StateMust() *kustoDatabasePrincipalAssignmentState {
	if kdpa.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", kdpa.Type(), kdpa.LocalName()))
	}
	return kdpa.state
}

// KustoDatabasePrincipalAssignmentArgs contains the configurations for azurerm_kusto_database_principal_assignment.
type KustoDatabasePrincipalAssignmentArgs struct {
	// ClusterName: string, required
	ClusterName terra.StringValue `hcl:"cluster_name,attr" validate:"required"`
	// DatabaseName: string, required
	DatabaseName terra.StringValue `hcl:"database_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// PrincipalId: string, required
	PrincipalId terra.StringValue `hcl:"principal_id,attr" validate:"required"`
	// PrincipalType: string, required
	PrincipalType terra.StringValue `hcl:"principal_type,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// TenantId: string, required
	TenantId terra.StringValue `hcl:"tenant_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *kustodatabaseprincipalassignment.Timeouts `hcl:"timeouts,block"`
}
type kustoDatabasePrincipalAssignmentAttributes struct {
	ref terra.Reference
}

// ClusterName returns a reference to field cluster_name of azurerm_kusto_database_principal_assignment.
func (kdpa kustoDatabasePrincipalAssignmentAttributes) ClusterName() terra.StringValue {
	return terra.ReferenceAsString(kdpa.ref.Append("cluster_name"))
}

// DatabaseName returns a reference to field database_name of azurerm_kusto_database_principal_assignment.
func (kdpa kustoDatabasePrincipalAssignmentAttributes) DatabaseName() terra.StringValue {
	return terra.ReferenceAsString(kdpa.ref.Append("database_name"))
}

// Id returns a reference to field id of azurerm_kusto_database_principal_assignment.
func (kdpa kustoDatabasePrincipalAssignmentAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(kdpa.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_kusto_database_principal_assignment.
func (kdpa kustoDatabasePrincipalAssignmentAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(kdpa.ref.Append("name"))
}

// PrincipalId returns a reference to field principal_id of azurerm_kusto_database_principal_assignment.
func (kdpa kustoDatabasePrincipalAssignmentAttributes) PrincipalId() terra.StringValue {
	return terra.ReferenceAsString(kdpa.ref.Append("principal_id"))
}

// PrincipalName returns a reference to field principal_name of azurerm_kusto_database_principal_assignment.
func (kdpa kustoDatabasePrincipalAssignmentAttributes) PrincipalName() terra.StringValue {
	return terra.ReferenceAsString(kdpa.ref.Append("principal_name"))
}

// PrincipalType returns a reference to field principal_type of azurerm_kusto_database_principal_assignment.
func (kdpa kustoDatabasePrincipalAssignmentAttributes) PrincipalType() terra.StringValue {
	return terra.ReferenceAsString(kdpa.ref.Append("principal_type"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_kusto_database_principal_assignment.
func (kdpa kustoDatabasePrincipalAssignmentAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(kdpa.ref.Append("resource_group_name"))
}

// Role returns a reference to field role of azurerm_kusto_database_principal_assignment.
func (kdpa kustoDatabasePrincipalAssignmentAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(kdpa.ref.Append("role"))
}

// TenantId returns a reference to field tenant_id of azurerm_kusto_database_principal_assignment.
func (kdpa kustoDatabasePrincipalAssignmentAttributes) TenantId() terra.StringValue {
	return terra.ReferenceAsString(kdpa.ref.Append("tenant_id"))
}

// TenantName returns a reference to field tenant_name of azurerm_kusto_database_principal_assignment.
func (kdpa kustoDatabasePrincipalAssignmentAttributes) TenantName() terra.StringValue {
	return terra.ReferenceAsString(kdpa.ref.Append("tenant_name"))
}

func (kdpa kustoDatabasePrincipalAssignmentAttributes) Timeouts() kustodatabaseprincipalassignment.TimeoutsAttributes {
	return terra.ReferenceAsSingle[kustodatabaseprincipalassignment.TimeoutsAttributes](kdpa.ref.Append("timeouts"))
}

type kustoDatabasePrincipalAssignmentState struct {
	ClusterName       string                                          `json:"cluster_name"`
	DatabaseName      string                                          `json:"database_name"`
	Id                string                                          `json:"id"`
	Name              string                                          `json:"name"`
	PrincipalId       string                                          `json:"principal_id"`
	PrincipalName     string                                          `json:"principal_name"`
	PrincipalType     string                                          `json:"principal_type"`
	ResourceGroupName string                                          `json:"resource_group_name"`
	Role              string                                          `json:"role"`
	TenantId          string                                          `json:"tenant_id"`
	TenantName        string                                          `json:"tenant_name"`
	Timeouts          *kustodatabaseprincipalassignment.TimeoutsState `json:"timeouts"`
}
