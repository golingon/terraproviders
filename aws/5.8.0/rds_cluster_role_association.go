// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	rdsclusterroleassociation "github.com/golingon/terraproviders/aws/5.8.0/rdsclusterroleassociation"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewRdsClusterRoleAssociation creates a new instance of [RdsClusterRoleAssociation].
func NewRdsClusterRoleAssociation(name string, args RdsClusterRoleAssociationArgs) *RdsClusterRoleAssociation {
	return &RdsClusterRoleAssociation{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*RdsClusterRoleAssociation)(nil)

// RdsClusterRoleAssociation represents the Terraform resource aws_rds_cluster_role_association.
type RdsClusterRoleAssociation struct {
	Name      string
	Args      RdsClusterRoleAssociationArgs
	state     *rdsClusterRoleAssociationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [RdsClusterRoleAssociation].
func (rcra *RdsClusterRoleAssociation) Type() string {
	return "aws_rds_cluster_role_association"
}

// LocalName returns the local name for [RdsClusterRoleAssociation].
func (rcra *RdsClusterRoleAssociation) LocalName() string {
	return rcra.Name
}

// Configuration returns the configuration (args) for [RdsClusterRoleAssociation].
func (rcra *RdsClusterRoleAssociation) Configuration() interface{} {
	return rcra.Args
}

// DependOn is used for other resources to depend on [RdsClusterRoleAssociation].
func (rcra *RdsClusterRoleAssociation) DependOn() terra.Reference {
	return terra.ReferenceResource(rcra)
}

// Dependencies returns the list of resources [RdsClusterRoleAssociation] depends_on.
func (rcra *RdsClusterRoleAssociation) Dependencies() terra.Dependencies {
	return rcra.DependsOn
}

// LifecycleManagement returns the lifecycle block for [RdsClusterRoleAssociation].
func (rcra *RdsClusterRoleAssociation) LifecycleManagement() *terra.Lifecycle {
	return rcra.Lifecycle
}

// Attributes returns the attributes for [RdsClusterRoleAssociation].
func (rcra *RdsClusterRoleAssociation) Attributes() rdsClusterRoleAssociationAttributes {
	return rdsClusterRoleAssociationAttributes{ref: terra.ReferenceResource(rcra)}
}

// ImportState imports the given attribute values into [RdsClusterRoleAssociation]'s state.
func (rcra *RdsClusterRoleAssociation) ImportState(av io.Reader) error {
	rcra.state = &rdsClusterRoleAssociationState{}
	if err := json.NewDecoder(av).Decode(rcra.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", rcra.Type(), rcra.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [RdsClusterRoleAssociation] has state.
func (rcra *RdsClusterRoleAssociation) State() (*rdsClusterRoleAssociationState, bool) {
	return rcra.state, rcra.state != nil
}

// StateMust returns the state for [RdsClusterRoleAssociation]. Panics if the state is nil.
func (rcra *RdsClusterRoleAssociation) StateMust() *rdsClusterRoleAssociationState {
	if rcra.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", rcra.Type(), rcra.LocalName()))
	}
	return rcra.state
}

// RdsClusterRoleAssociationArgs contains the configurations for aws_rds_cluster_role_association.
type RdsClusterRoleAssociationArgs struct {
	// DbClusterIdentifier: string, required
	DbClusterIdentifier terra.StringValue `hcl:"db_cluster_identifier,attr" validate:"required"`
	// FeatureName: string, required
	FeatureName terra.StringValue `hcl:"feature_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// RoleArn: string, required
	RoleArn terra.StringValue `hcl:"role_arn,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *rdsclusterroleassociation.Timeouts `hcl:"timeouts,block"`
}
type rdsClusterRoleAssociationAttributes struct {
	ref terra.Reference
}

// DbClusterIdentifier returns a reference to field db_cluster_identifier of aws_rds_cluster_role_association.
func (rcra rdsClusterRoleAssociationAttributes) DbClusterIdentifier() terra.StringValue {
	return terra.ReferenceAsString(rcra.ref.Append("db_cluster_identifier"))
}

// FeatureName returns a reference to field feature_name of aws_rds_cluster_role_association.
func (rcra rdsClusterRoleAssociationAttributes) FeatureName() terra.StringValue {
	return terra.ReferenceAsString(rcra.ref.Append("feature_name"))
}

// Id returns a reference to field id of aws_rds_cluster_role_association.
func (rcra rdsClusterRoleAssociationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(rcra.ref.Append("id"))
}

// RoleArn returns a reference to field role_arn of aws_rds_cluster_role_association.
func (rcra rdsClusterRoleAssociationAttributes) RoleArn() terra.StringValue {
	return terra.ReferenceAsString(rcra.ref.Append("role_arn"))
}

func (rcra rdsClusterRoleAssociationAttributes) Timeouts() rdsclusterroleassociation.TimeoutsAttributes {
	return terra.ReferenceAsSingle[rdsclusterroleassociation.TimeoutsAttributes](rcra.ref.Append("timeouts"))
}

type rdsClusterRoleAssociationState struct {
	DbClusterIdentifier string                                   `json:"db_cluster_identifier"`
	FeatureName         string                                   `json:"feature_name"`
	Id                  string                                   `json:"id"`
	RoleArn             string                                   `json:"role_arn"`
	Timeouts            *rdsclusterroleassociation.TimeoutsState `json:"timeouts"`
}
