// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	dataiamprincipalpolicysimulation "github.com/golingon/terraproviders/aws/5.10.0/dataiamprincipalpolicysimulation"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataIamPrincipalPolicySimulation creates a new instance of [DataIamPrincipalPolicySimulation].
func NewDataIamPrincipalPolicySimulation(name string, args DataIamPrincipalPolicySimulationArgs) *DataIamPrincipalPolicySimulation {
	return &DataIamPrincipalPolicySimulation{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataIamPrincipalPolicySimulation)(nil)

// DataIamPrincipalPolicySimulation represents the Terraform data resource aws_iam_principal_policy_simulation.
type DataIamPrincipalPolicySimulation struct {
	Name string
	Args DataIamPrincipalPolicySimulationArgs
}

// DataSource returns the Terraform object type for [DataIamPrincipalPolicySimulation].
func (ipps *DataIamPrincipalPolicySimulation) DataSource() string {
	return "aws_iam_principal_policy_simulation"
}

// LocalName returns the local name for [DataIamPrincipalPolicySimulation].
func (ipps *DataIamPrincipalPolicySimulation) LocalName() string {
	return ipps.Name
}

// Configuration returns the configuration (args) for [DataIamPrincipalPolicySimulation].
func (ipps *DataIamPrincipalPolicySimulation) Configuration() interface{} {
	return ipps.Args
}

// Attributes returns the attributes for [DataIamPrincipalPolicySimulation].
func (ipps *DataIamPrincipalPolicySimulation) Attributes() dataIamPrincipalPolicySimulationAttributes {
	return dataIamPrincipalPolicySimulationAttributes{ref: terra.ReferenceDataResource(ipps)}
}

// DataIamPrincipalPolicySimulationArgs contains the configurations for aws_iam_principal_policy_simulation.
type DataIamPrincipalPolicySimulationArgs struct {
	// ActionNames: set of string, required
	ActionNames terra.SetValue[terra.StringValue] `hcl:"action_names,attr" validate:"required"`
	// AdditionalPoliciesJson: set of string, optional
	AdditionalPoliciesJson terra.SetValue[terra.StringValue] `hcl:"additional_policies_json,attr"`
	// CallerArn: string, optional
	CallerArn terra.StringValue `hcl:"caller_arn,attr"`
	// PermissionsBoundaryPoliciesJson: set of string, optional
	PermissionsBoundaryPoliciesJson terra.SetValue[terra.StringValue] `hcl:"permissions_boundary_policies_json,attr"`
	// PolicySourceArn: string, required
	PolicySourceArn terra.StringValue `hcl:"policy_source_arn,attr" validate:"required"`
	// ResourceArns: set of string, optional
	ResourceArns terra.SetValue[terra.StringValue] `hcl:"resource_arns,attr"`
	// ResourceHandlingOption: string, optional
	ResourceHandlingOption terra.StringValue `hcl:"resource_handling_option,attr"`
	// ResourceOwnerAccountId: string, optional
	ResourceOwnerAccountId terra.StringValue `hcl:"resource_owner_account_id,attr"`
	// ResourcePolicyJson: string, optional
	ResourcePolicyJson terra.StringValue `hcl:"resource_policy_json,attr"`
	// Results: min=0
	Results []dataiamprincipalpolicysimulation.Results `hcl:"results,block" validate:"min=0"`
	// Context: min=0
	Context []dataiamprincipalpolicysimulation.Context `hcl:"context,block" validate:"min=0"`
}
type dataIamPrincipalPolicySimulationAttributes struct {
	ref terra.Reference
}

// ActionNames returns a reference to field action_names of aws_iam_principal_policy_simulation.
func (ipps dataIamPrincipalPolicySimulationAttributes) ActionNames() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](ipps.ref.Append("action_names"))
}

// AdditionalPoliciesJson returns a reference to field additional_policies_json of aws_iam_principal_policy_simulation.
func (ipps dataIamPrincipalPolicySimulationAttributes) AdditionalPoliciesJson() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](ipps.ref.Append("additional_policies_json"))
}

// AllAllowed returns a reference to field all_allowed of aws_iam_principal_policy_simulation.
func (ipps dataIamPrincipalPolicySimulationAttributes) AllAllowed() terra.BoolValue {
	return terra.ReferenceAsBool(ipps.ref.Append("all_allowed"))
}

// CallerArn returns a reference to field caller_arn of aws_iam_principal_policy_simulation.
func (ipps dataIamPrincipalPolicySimulationAttributes) CallerArn() terra.StringValue {
	return terra.ReferenceAsString(ipps.ref.Append("caller_arn"))
}

// Id returns a reference to field id of aws_iam_principal_policy_simulation.
func (ipps dataIamPrincipalPolicySimulationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ipps.ref.Append("id"))
}

// PermissionsBoundaryPoliciesJson returns a reference to field permissions_boundary_policies_json of aws_iam_principal_policy_simulation.
func (ipps dataIamPrincipalPolicySimulationAttributes) PermissionsBoundaryPoliciesJson() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](ipps.ref.Append("permissions_boundary_policies_json"))
}

// PolicySourceArn returns a reference to field policy_source_arn of aws_iam_principal_policy_simulation.
func (ipps dataIamPrincipalPolicySimulationAttributes) PolicySourceArn() terra.StringValue {
	return terra.ReferenceAsString(ipps.ref.Append("policy_source_arn"))
}

// ResourceArns returns a reference to field resource_arns of aws_iam_principal_policy_simulation.
func (ipps dataIamPrincipalPolicySimulationAttributes) ResourceArns() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](ipps.ref.Append("resource_arns"))
}

// ResourceHandlingOption returns a reference to field resource_handling_option of aws_iam_principal_policy_simulation.
func (ipps dataIamPrincipalPolicySimulationAttributes) ResourceHandlingOption() terra.StringValue {
	return terra.ReferenceAsString(ipps.ref.Append("resource_handling_option"))
}

// ResourceOwnerAccountId returns a reference to field resource_owner_account_id of aws_iam_principal_policy_simulation.
func (ipps dataIamPrincipalPolicySimulationAttributes) ResourceOwnerAccountId() terra.StringValue {
	return terra.ReferenceAsString(ipps.ref.Append("resource_owner_account_id"))
}

// ResourcePolicyJson returns a reference to field resource_policy_json of aws_iam_principal_policy_simulation.
func (ipps dataIamPrincipalPolicySimulationAttributes) ResourcePolicyJson() terra.StringValue {
	return terra.ReferenceAsString(ipps.ref.Append("resource_policy_json"))
}

func (ipps dataIamPrincipalPolicySimulationAttributes) Results() terra.SetValue[dataiamprincipalpolicysimulation.ResultsAttributes] {
	return terra.ReferenceAsSet[dataiamprincipalpolicysimulation.ResultsAttributes](ipps.ref.Append("results"))
}

func (ipps dataIamPrincipalPolicySimulationAttributes) Context() terra.SetValue[dataiamprincipalpolicysimulation.ContextAttributes] {
	return terra.ReferenceAsSet[dataiamprincipalpolicysimulation.ContextAttributes](ipps.ref.Append("context"))
}
