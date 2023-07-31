// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	redshiftclusteriamroles "github.com/golingon/terraproviders/aws/5.10.0/redshiftclusteriamroles"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewRedshiftClusterIamRoles creates a new instance of [RedshiftClusterIamRoles].
func NewRedshiftClusterIamRoles(name string, args RedshiftClusterIamRolesArgs) *RedshiftClusterIamRoles {
	return &RedshiftClusterIamRoles{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*RedshiftClusterIamRoles)(nil)

// RedshiftClusterIamRoles represents the Terraform resource aws_redshift_cluster_iam_roles.
type RedshiftClusterIamRoles struct {
	Name      string
	Args      RedshiftClusterIamRolesArgs
	state     *redshiftClusterIamRolesState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [RedshiftClusterIamRoles].
func (rcir *RedshiftClusterIamRoles) Type() string {
	return "aws_redshift_cluster_iam_roles"
}

// LocalName returns the local name for [RedshiftClusterIamRoles].
func (rcir *RedshiftClusterIamRoles) LocalName() string {
	return rcir.Name
}

// Configuration returns the configuration (args) for [RedshiftClusterIamRoles].
func (rcir *RedshiftClusterIamRoles) Configuration() interface{} {
	return rcir.Args
}

// DependOn is used for other resources to depend on [RedshiftClusterIamRoles].
func (rcir *RedshiftClusterIamRoles) DependOn() terra.Reference {
	return terra.ReferenceResource(rcir)
}

// Dependencies returns the list of resources [RedshiftClusterIamRoles] depends_on.
func (rcir *RedshiftClusterIamRoles) Dependencies() terra.Dependencies {
	return rcir.DependsOn
}

// LifecycleManagement returns the lifecycle block for [RedshiftClusterIamRoles].
func (rcir *RedshiftClusterIamRoles) LifecycleManagement() *terra.Lifecycle {
	return rcir.Lifecycle
}

// Attributes returns the attributes for [RedshiftClusterIamRoles].
func (rcir *RedshiftClusterIamRoles) Attributes() redshiftClusterIamRolesAttributes {
	return redshiftClusterIamRolesAttributes{ref: terra.ReferenceResource(rcir)}
}

// ImportState imports the given attribute values into [RedshiftClusterIamRoles]'s state.
func (rcir *RedshiftClusterIamRoles) ImportState(av io.Reader) error {
	rcir.state = &redshiftClusterIamRolesState{}
	if err := json.NewDecoder(av).Decode(rcir.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", rcir.Type(), rcir.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [RedshiftClusterIamRoles] has state.
func (rcir *RedshiftClusterIamRoles) State() (*redshiftClusterIamRolesState, bool) {
	return rcir.state, rcir.state != nil
}

// StateMust returns the state for [RedshiftClusterIamRoles]. Panics if the state is nil.
func (rcir *RedshiftClusterIamRoles) StateMust() *redshiftClusterIamRolesState {
	if rcir.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", rcir.Type(), rcir.LocalName()))
	}
	return rcir.state
}

// RedshiftClusterIamRolesArgs contains the configurations for aws_redshift_cluster_iam_roles.
type RedshiftClusterIamRolesArgs struct {
	// ClusterIdentifier: string, required
	ClusterIdentifier terra.StringValue `hcl:"cluster_identifier,attr" validate:"required"`
	// DefaultIamRoleArn: string, optional
	DefaultIamRoleArn terra.StringValue `hcl:"default_iam_role_arn,attr"`
	// IamRoleArns: set of string, optional
	IamRoleArns terra.SetValue[terra.StringValue] `hcl:"iam_role_arns,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Timeouts: optional
	Timeouts *redshiftclusteriamroles.Timeouts `hcl:"timeouts,block"`
}
type redshiftClusterIamRolesAttributes struct {
	ref terra.Reference
}

// ClusterIdentifier returns a reference to field cluster_identifier of aws_redshift_cluster_iam_roles.
func (rcir redshiftClusterIamRolesAttributes) ClusterIdentifier() terra.StringValue {
	return terra.ReferenceAsString(rcir.ref.Append("cluster_identifier"))
}

// DefaultIamRoleArn returns a reference to field default_iam_role_arn of aws_redshift_cluster_iam_roles.
func (rcir redshiftClusterIamRolesAttributes) DefaultIamRoleArn() terra.StringValue {
	return terra.ReferenceAsString(rcir.ref.Append("default_iam_role_arn"))
}

// IamRoleArns returns a reference to field iam_role_arns of aws_redshift_cluster_iam_roles.
func (rcir redshiftClusterIamRolesAttributes) IamRoleArns() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](rcir.ref.Append("iam_role_arns"))
}

// Id returns a reference to field id of aws_redshift_cluster_iam_roles.
func (rcir redshiftClusterIamRolesAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(rcir.ref.Append("id"))
}

func (rcir redshiftClusterIamRolesAttributes) Timeouts() redshiftclusteriamroles.TimeoutsAttributes {
	return terra.ReferenceAsSingle[redshiftclusteriamroles.TimeoutsAttributes](rcir.ref.Append("timeouts"))
}

type redshiftClusterIamRolesState struct {
	ClusterIdentifier string                                 `json:"cluster_identifier"`
	DefaultIamRoleArn string                                 `json:"default_iam_role_arn"`
	IamRoleArns       []string                               `json:"iam_role_arns"`
	Id                string                                 `json:"id"`
	Timeouts          *redshiftclusteriamroles.TimeoutsState `json:"timeouts"`
}
