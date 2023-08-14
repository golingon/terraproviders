// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	kustoclusterprincipalassignment "github.com/golingon/terraproviders/azurerm/3.69.0/kustoclusterprincipalassignment"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewKustoClusterPrincipalAssignment creates a new instance of [KustoClusterPrincipalAssignment].
func NewKustoClusterPrincipalAssignment(name string, args KustoClusterPrincipalAssignmentArgs) *KustoClusterPrincipalAssignment {
	return &KustoClusterPrincipalAssignment{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*KustoClusterPrincipalAssignment)(nil)

// KustoClusterPrincipalAssignment represents the Terraform resource azurerm_kusto_cluster_principal_assignment.
type KustoClusterPrincipalAssignment struct {
	Name      string
	Args      KustoClusterPrincipalAssignmentArgs
	state     *kustoClusterPrincipalAssignmentState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [KustoClusterPrincipalAssignment].
func (kcpa *KustoClusterPrincipalAssignment) Type() string {
	return "azurerm_kusto_cluster_principal_assignment"
}

// LocalName returns the local name for [KustoClusterPrincipalAssignment].
func (kcpa *KustoClusterPrincipalAssignment) LocalName() string {
	return kcpa.Name
}

// Configuration returns the configuration (args) for [KustoClusterPrincipalAssignment].
func (kcpa *KustoClusterPrincipalAssignment) Configuration() interface{} {
	return kcpa.Args
}

// DependOn is used for other resources to depend on [KustoClusterPrincipalAssignment].
func (kcpa *KustoClusterPrincipalAssignment) DependOn() terra.Reference {
	return terra.ReferenceResource(kcpa)
}

// Dependencies returns the list of resources [KustoClusterPrincipalAssignment] depends_on.
func (kcpa *KustoClusterPrincipalAssignment) Dependencies() terra.Dependencies {
	return kcpa.DependsOn
}

// LifecycleManagement returns the lifecycle block for [KustoClusterPrincipalAssignment].
func (kcpa *KustoClusterPrincipalAssignment) LifecycleManagement() *terra.Lifecycle {
	return kcpa.Lifecycle
}

// Attributes returns the attributes for [KustoClusterPrincipalAssignment].
func (kcpa *KustoClusterPrincipalAssignment) Attributes() kustoClusterPrincipalAssignmentAttributes {
	return kustoClusterPrincipalAssignmentAttributes{ref: terra.ReferenceResource(kcpa)}
}

// ImportState imports the given attribute values into [KustoClusterPrincipalAssignment]'s state.
func (kcpa *KustoClusterPrincipalAssignment) ImportState(av io.Reader) error {
	kcpa.state = &kustoClusterPrincipalAssignmentState{}
	if err := json.NewDecoder(av).Decode(kcpa.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", kcpa.Type(), kcpa.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [KustoClusterPrincipalAssignment] has state.
func (kcpa *KustoClusterPrincipalAssignment) State() (*kustoClusterPrincipalAssignmentState, bool) {
	return kcpa.state, kcpa.state != nil
}

// StateMust returns the state for [KustoClusterPrincipalAssignment]. Panics if the state is nil.
func (kcpa *KustoClusterPrincipalAssignment) StateMust() *kustoClusterPrincipalAssignmentState {
	if kcpa.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", kcpa.Type(), kcpa.LocalName()))
	}
	return kcpa.state
}

// KustoClusterPrincipalAssignmentArgs contains the configurations for azurerm_kusto_cluster_principal_assignment.
type KustoClusterPrincipalAssignmentArgs struct {
	// ClusterName: string, required
	ClusterName terra.StringValue `hcl:"cluster_name,attr" validate:"required"`
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
	Timeouts *kustoclusterprincipalassignment.Timeouts `hcl:"timeouts,block"`
}
type kustoClusterPrincipalAssignmentAttributes struct {
	ref terra.Reference
}

// ClusterName returns a reference to field cluster_name of azurerm_kusto_cluster_principal_assignment.
func (kcpa kustoClusterPrincipalAssignmentAttributes) ClusterName() terra.StringValue {
	return terra.ReferenceAsString(kcpa.ref.Append("cluster_name"))
}

// Id returns a reference to field id of azurerm_kusto_cluster_principal_assignment.
func (kcpa kustoClusterPrincipalAssignmentAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(kcpa.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_kusto_cluster_principal_assignment.
func (kcpa kustoClusterPrincipalAssignmentAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(kcpa.ref.Append("name"))
}

// PrincipalId returns a reference to field principal_id of azurerm_kusto_cluster_principal_assignment.
func (kcpa kustoClusterPrincipalAssignmentAttributes) PrincipalId() terra.StringValue {
	return terra.ReferenceAsString(kcpa.ref.Append("principal_id"))
}

// PrincipalName returns a reference to field principal_name of azurerm_kusto_cluster_principal_assignment.
func (kcpa kustoClusterPrincipalAssignmentAttributes) PrincipalName() terra.StringValue {
	return terra.ReferenceAsString(kcpa.ref.Append("principal_name"))
}

// PrincipalType returns a reference to field principal_type of azurerm_kusto_cluster_principal_assignment.
func (kcpa kustoClusterPrincipalAssignmentAttributes) PrincipalType() terra.StringValue {
	return terra.ReferenceAsString(kcpa.ref.Append("principal_type"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_kusto_cluster_principal_assignment.
func (kcpa kustoClusterPrincipalAssignmentAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(kcpa.ref.Append("resource_group_name"))
}

// Role returns a reference to field role of azurerm_kusto_cluster_principal_assignment.
func (kcpa kustoClusterPrincipalAssignmentAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(kcpa.ref.Append("role"))
}

// TenantId returns a reference to field tenant_id of azurerm_kusto_cluster_principal_assignment.
func (kcpa kustoClusterPrincipalAssignmentAttributes) TenantId() terra.StringValue {
	return terra.ReferenceAsString(kcpa.ref.Append("tenant_id"))
}

// TenantName returns a reference to field tenant_name of azurerm_kusto_cluster_principal_assignment.
func (kcpa kustoClusterPrincipalAssignmentAttributes) TenantName() terra.StringValue {
	return terra.ReferenceAsString(kcpa.ref.Append("tenant_name"))
}

func (kcpa kustoClusterPrincipalAssignmentAttributes) Timeouts() kustoclusterprincipalassignment.TimeoutsAttributes {
	return terra.ReferenceAsSingle[kustoclusterprincipalassignment.TimeoutsAttributes](kcpa.ref.Append("timeouts"))
}

type kustoClusterPrincipalAssignmentState struct {
	ClusterName       string                                         `json:"cluster_name"`
	Id                string                                         `json:"id"`
	Name              string                                         `json:"name"`
	PrincipalId       string                                         `json:"principal_id"`
	PrincipalName     string                                         `json:"principal_name"`
	PrincipalType     string                                         `json:"principal_type"`
	ResourceGroupName string                                         `json:"resource_group_name"`
	Role              string                                         `json:"role"`
	TenantId          string                                         `json:"tenant_id"`
	TenantName        string                                         `json:"tenant_name"`
	Timeouts          *kustoclusterprincipalassignment.TimeoutsState `json:"timeouts"`
}