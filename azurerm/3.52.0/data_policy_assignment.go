// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datapolicyassignment "github.com/golingon/terraproviders/azurerm/3.52.0/datapolicyassignment"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataPolicyAssignment creates a new instance of [DataPolicyAssignment].
func NewDataPolicyAssignment(name string, args DataPolicyAssignmentArgs) *DataPolicyAssignment {
	return &DataPolicyAssignment{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataPolicyAssignment)(nil)

// DataPolicyAssignment represents the Terraform data resource azurerm_policy_assignment.
type DataPolicyAssignment struct {
	Name string
	Args DataPolicyAssignmentArgs
}

// DataSource returns the Terraform object type for [DataPolicyAssignment].
func (pa *DataPolicyAssignment) DataSource() string {
	return "azurerm_policy_assignment"
}

// LocalName returns the local name for [DataPolicyAssignment].
func (pa *DataPolicyAssignment) LocalName() string {
	return pa.Name
}

// Configuration returns the configuration (args) for [DataPolicyAssignment].
func (pa *DataPolicyAssignment) Configuration() interface{} {
	return pa.Args
}

// Attributes returns the attributes for [DataPolicyAssignment].
func (pa *DataPolicyAssignment) Attributes() dataPolicyAssignmentAttributes {
	return dataPolicyAssignmentAttributes{ref: terra.ReferenceDataResource(pa)}
}

// DataPolicyAssignmentArgs contains the configurations for azurerm_policy_assignment.
type DataPolicyAssignmentArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ScopeId: string, required
	ScopeId terra.StringValue `hcl:"scope_id,attr" validate:"required"`
	// Identity: min=0
	Identity []datapolicyassignment.Identity `hcl:"identity,block" validate:"min=0"`
	// NonComplianceMessage: min=0
	NonComplianceMessage []datapolicyassignment.NonComplianceMessage `hcl:"non_compliance_message,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *datapolicyassignment.Timeouts `hcl:"timeouts,block"`
}
type dataPolicyAssignmentAttributes struct {
	ref terra.Reference
}

// Description returns a reference to field description of azurerm_policy_assignment.
func (pa dataPolicyAssignmentAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(pa.ref.Append("description"))
}

// DisplayName returns a reference to field display_name of azurerm_policy_assignment.
func (pa dataPolicyAssignmentAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(pa.ref.Append("display_name"))
}

// Enforce returns a reference to field enforce of azurerm_policy_assignment.
func (pa dataPolicyAssignmentAttributes) Enforce() terra.BoolValue {
	return terra.ReferenceAsBool(pa.ref.Append("enforce"))
}

// Id returns a reference to field id of azurerm_policy_assignment.
func (pa dataPolicyAssignmentAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(pa.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_policy_assignment.
func (pa dataPolicyAssignmentAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(pa.ref.Append("location"))
}

// Metadata returns a reference to field metadata of azurerm_policy_assignment.
func (pa dataPolicyAssignmentAttributes) Metadata() terra.StringValue {
	return terra.ReferenceAsString(pa.ref.Append("metadata"))
}

// Name returns a reference to field name of azurerm_policy_assignment.
func (pa dataPolicyAssignmentAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(pa.ref.Append("name"))
}

// NotScopes returns a reference to field not_scopes of azurerm_policy_assignment.
func (pa dataPolicyAssignmentAttributes) NotScopes() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](pa.ref.Append("not_scopes"))
}

// Parameters returns a reference to field parameters of azurerm_policy_assignment.
func (pa dataPolicyAssignmentAttributes) Parameters() terra.StringValue {
	return terra.ReferenceAsString(pa.ref.Append("parameters"))
}

// PolicyDefinitionId returns a reference to field policy_definition_id of azurerm_policy_assignment.
func (pa dataPolicyAssignmentAttributes) PolicyDefinitionId() terra.StringValue {
	return terra.ReferenceAsString(pa.ref.Append("policy_definition_id"))
}

// ScopeId returns a reference to field scope_id of azurerm_policy_assignment.
func (pa dataPolicyAssignmentAttributes) ScopeId() terra.StringValue {
	return terra.ReferenceAsString(pa.ref.Append("scope_id"))
}

func (pa dataPolicyAssignmentAttributes) Identity() terra.ListValue[datapolicyassignment.IdentityAttributes] {
	return terra.ReferenceAsList[datapolicyassignment.IdentityAttributes](pa.ref.Append("identity"))
}

func (pa dataPolicyAssignmentAttributes) NonComplianceMessage() terra.ListValue[datapolicyassignment.NonComplianceMessageAttributes] {
	return terra.ReferenceAsList[datapolicyassignment.NonComplianceMessageAttributes](pa.ref.Append("non_compliance_message"))
}

func (pa dataPolicyAssignmentAttributes) Timeouts() datapolicyassignment.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datapolicyassignment.TimeoutsAttributes](pa.ref.Append("timeouts"))
}
