// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	redshiftdatastatement "github.com/golingon/terraproviders/aws/5.8.0/redshiftdatastatement"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewRedshiftdataStatement creates a new instance of [RedshiftdataStatement].
func NewRedshiftdataStatement(name string, args RedshiftdataStatementArgs) *RedshiftdataStatement {
	return &RedshiftdataStatement{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*RedshiftdataStatement)(nil)

// RedshiftdataStatement represents the Terraform resource aws_redshiftdata_statement.
type RedshiftdataStatement struct {
	Name      string
	Args      RedshiftdataStatementArgs
	state     *redshiftdataStatementState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [RedshiftdataStatement].
func (rs *RedshiftdataStatement) Type() string {
	return "aws_redshiftdata_statement"
}

// LocalName returns the local name for [RedshiftdataStatement].
func (rs *RedshiftdataStatement) LocalName() string {
	return rs.Name
}

// Configuration returns the configuration (args) for [RedshiftdataStatement].
func (rs *RedshiftdataStatement) Configuration() interface{} {
	return rs.Args
}

// DependOn is used for other resources to depend on [RedshiftdataStatement].
func (rs *RedshiftdataStatement) DependOn() terra.Reference {
	return terra.ReferenceResource(rs)
}

// Dependencies returns the list of resources [RedshiftdataStatement] depends_on.
func (rs *RedshiftdataStatement) Dependencies() terra.Dependencies {
	return rs.DependsOn
}

// LifecycleManagement returns the lifecycle block for [RedshiftdataStatement].
func (rs *RedshiftdataStatement) LifecycleManagement() *terra.Lifecycle {
	return rs.Lifecycle
}

// Attributes returns the attributes for [RedshiftdataStatement].
func (rs *RedshiftdataStatement) Attributes() redshiftdataStatementAttributes {
	return redshiftdataStatementAttributes{ref: terra.ReferenceResource(rs)}
}

// ImportState imports the given attribute values into [RedshiftdataStatement]'s state.
func (rs *RedshiftdataStatement) ImportState(av io.Reader) error {
	rs.state = &redshiftdataStatementState{}
	if err := json.NewDecoder(av).Decode(rs.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", rs.Type(), rs.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [RedshiftdataStatement] has state.
func (rs *RedshiftdataStatement) State() (*redshiftdataStatementState, bool) {
	return rs.state, rs.state != nil
}

// StateMust returns the state for [RedshiftdataStatement]. Panics if the state is nil.
func (rs *RedshiftdataStatement) StateMust() *redshiftdataStatementState {
	if rs.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", rs.Type(), rs.LocalName()))
	}
	return rs.state
}

// RedshiftdataStatementArgs contains the configurations for aws_redshiftdata_statement.
type RedshiftdataStatementArgs struct {
	// ClusterIdentifier: string, optional
	ClusterIdentifier terra.StringValue `hcl:"cluster_identifier,attr"`
	// Database: string, required
	Database terra.StringValue `hcl:"database,attr" validate:"required"`
	// DbUser: string, optional
	DbUser terra.StringValue `hcl:"db_user,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// SecretArn: string, optional
	SecretArn terra.StringValue `hcl:"secret_arn,attr"`
	// Sql: string, required
	Sql terra.StringValue `hcl:"sql,attr" validate:"required"`
	// StatementName: string, optional
	StatementName terra.StringValue `hcl:"statement_name,attr"`
	// WithEvent: bool, optional
	WithEvent terra.BoolValue `hcl:"with_event,attr"`
	// WorkgroupName: string, optional
	WorkgroupName terra.StringValue `hcl:"workgroup_name,attr"`
	// Parameters: min=0
	Parameters []redshiftdatastatement.Parameters `hcl:"parameters,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *redshiftdatastatement.Timeouts `hcl:"timeouts,block"`
}
type redshiftdataStatementAttributes struct {
	ref terra.Reference
}

// ClusterIdentifier returns a reference to field cluster_identifier of aws_redshiftdata_statement.
func (rs redshiftdataStatementAttributes) ClusterIdentifier() terra.StringValue {
	return terra.ReferenceAsString(rs.ref.Append("cluster_identifier"))
}

// Database returns a reference to field database of aws_redshiftdata_statement.
func (rs redshiftdataStatementAttributes) Database() terra.StringValue {
	return terra.ReferenceAsString(rs.ref.Append("database"))
}

// DbUser returns a reference to field db_user of aws_redshiftdata_statement.
func (rs redshiftdataStatementAttributes) DbUser() terra.StringValue {
	return terra.ReferenceAsString(rs.ref.Append("db_user"))
}

// Id returns a reference to field id of aws_redshiftdata_statement.
func (rs redshiftdataStatementAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(rs.ref.Append("id"))
}

// SecretArn returns a reference to field secret_arn of aws_redshiftdata_statement.
func (rs redshiftdataStatementAttributes) SecretArn() terra.StringValue {
	return terra.ReferenceAsString(rs.ref.Append("secret_arn"))
}

// Sql returns a reference to field sql of aws_redshiftdata_statement.
func (rs redshiftdataStatementAttributes) Sql() terra.StringValue {
	return terra.ReferenceAsString(rs.ref.Append("sql"))
}

// StatementName returns a reference to field statement_name of aws_redshiftdata_statement.
func (rs redshiftdataStatementAttributes) StatementName() terra.StringValue {
	return terra.ReferenceAsString(rs.ref.Append("statement_name"))
}

// WithEvent returns a reference to field with_event of aws_redshiftdata_statement.
func (rs redshiftdataStatementAttributes) WithEvent() terra.BoolValue {
	return terra.ReferenceAsBool(rs.ref.Append("with_event"))
}

// WorkgroupName returns a reference to field workgroup_name of aws_redshiftdata_statement.
func (rs redshiftdataStatementAttributes) WorkgroupName() terra.StringValue {
	return terra.ReferenceAsString(rs.ref.Append("workgroup_name"))
}

func (rs redshiftdataStatementAttributes) Parameters() terra.ListValue[redshiftdatastatement.ParametersAttributes] {
	return terra.ReferenceAsList[redshiftdatastatement.ParametersAttributes](rs.ref.Append("parameters"))
}

func (rs redshiftdataStatementAttributes) Timeouts() redshiftdatastatement.TimeoutsAttributes {
	return terra.ReferenceAsSingle[redshiftdatastatement.TimeoutsAttributes](rs.ref.Append("timeouts"))
}

type redshiftdataStatementState struct {
	ClusterIdentifier string                                  `json:"cluster_identifier"`
	Database          string                                  `json:"database"`
	DbUser            string                                  `json:"db_user"`
	Id                string                                  `json:"id"`
	SecretArn         string                                  `json:"secret_arn"`
	Sql               string                                  `json:"sql"`
	StatementName     string                                  `json:"statement_name"`
	WithEvent         bool                                    `json:"with_event"`
	WorkgroupName     string                                  `json:"workgroup_name"`
	Parameters        []redshiftdatastatement.ParametersState `json:"parameters"`
	Timeouts          *redshiftdatastatement.TimeoutsState    `json:"timeouts"`
}
