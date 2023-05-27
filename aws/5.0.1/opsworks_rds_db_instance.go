// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewOpsworksRdsDbInstance creates a new instance of [OpsworksRdsDbInstance].
func NewOpsworksRdsDbInstance(name string, args OpsworksRdsDbInstanceArgs) *OpsworksRdsDbInstance {
	return &OpsworksRdsDbInstance{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*OpsworksRdsDbInstance)(nil)

// OpsworksRdsDbInstance represents the Terraform resource aws_opsworks_rds_db_instance.
type OpsworksRdsDbInstance struct {
	Name      string
	Args      OpsworksRdsDbInstanceArgs
	state     *opsworksRdsDbInstanceState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [OpsworksRdsDbInstance].
func (ordi *OpsworksRdsDbInstance) Type() string {
	return "aws_opsworks_rds_db_instance"
}

// LocalName returns the local name for [OpsworksRdsDbInstance].
func (ordi *OpsworksRdsDbInstance) LocalName() string {
	return ordi.Name
}

// Configuration returns the configuration (args) for [OpsworksRdsDbInstance].
func (ordi *OpsworksRdsDbInstance) Configuration() interface{} {
	return ordi.Args
}

// DependOn is used for other resources to depend on [OpsworksRdsDbInstance].
func (ordi *OpsworksRdsDbInstance) DependOn() terra.Reference {
	return terra.ReferenceResource(ordi)
}

// Dependencies returns the list of resources [OpsworksRdsDbInstance] depends_on.
func (ordi *OpsworksRdsDbInstance) Dependencies() terra.Dependencies {
	return ordi.DependsOn
}

// LifecycleManagement returns the lifecycle block for [OpsworksRdsDbInstance].
func (ordi *OpsworksRdsDbInstance) LifecycleManagement() *terra.Lifecycle {
	return ordi.Lifecycle
}

// Attributes returns the attributes for [OpsworksRdsDbInstance].
func (ordi *OpsworksRdsDbInstance) Attributes() opsworksRdsDbInstanceAttributes {
	return opsworksRdsDbInstanceAttributes{ref: terra.ReferenceResource(ordi)}
}

// ImportState imports the given attribute values into [OpsworksRdsDbInstance]'s state.
func (ordi *OpsworksRdsDbInstance) ImportState(av io.Reader) error {
	ordi.state = &opsworksRdsDbInstanceState{}
	if err := json.NewDecoder(av).Decode(ordi.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ordi.Type(), ordi.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [OpsworksRdsDbInstance] has state.
func (ordi *OpsworksRdsDbInstance) State() (*opsworksRdsDbInstanceState, bool) {
	return ordi.state, ordi.state != nil
}

// StateMust returns the state for [OpsworksRdsDbInstance]. Panics if the state is nil.
func (ordi *OpsworksRdsDbInstance) StateMust() *opsworksRdsDbInstanceState {
	if ordi.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ordi.Type(), ordi.LocalName()))
	}
	return ordi.state
}

// OpsworksRdsDbInstanceArgs contains the configurations for aws_opsworks_rds_db_instance.
type OpsworksRdsDbInstanceArgs struct {
	// DbPassword: string, required
	DbPassword terra.StringValue `hcl:"db_password,attr" validate:"required"`
	// DbUser: string, required
	DbUser terra.StringValue `hcl:"db_user,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// RdsDbInstanceArn: string, required
	RdsDbInstanceArn terra.StringValue `hcl:"rds_db_instance_arn,attr" validate:"required"`
	// StackId: string, required
	StackId terra.StringValue `hcl:"stack_id,attr" validate:"required"`
}
type opsworksRdsDbInstanceAttributes struct {
	ref terra.Reference
}

// DbPassword returns a reference to field db_password of aws_opsworks_rds_db_instance.
func (ordi opsworksRdsDbInstanceAttributes) DbPassword() terra.StringValue {
	return terra.ReferenceAsString(ordi.ref.Append("db_password"))
}

// DbUser returns a reference to field db_user of aws_opsworks_rds_db_instance.
func (ordi opsworksRdsDbInstanceAttributes) DbUser() terra.StringValue {
	return terra.ReferenceAsString(ordi.ref.Append("db_user"))
}

// Id returns a reference to field id of aws_opsworks_rds_db_instance.
func (ordi opsworksRdsDbInstanceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ordi.ref.Append("id"))
}

// RdsDbInstanceArn returns a reference to field rds_db_instance_arn of aws_opsworks_rds_db_instance.
func (ordi opsworksRdsDbInstanceAttributes) RdsDbInstanceArn() terra.StringValue {
	return terra.ReferenceAsString(ordi.ref.Append("rds_db_instance_arn"))
}

// StackId returns a reference to field stack_id of aws_opsworks_rds_db_instance.
func (ordi opsworksRdsDbInstanceAttributes) StackId() terra.StringValue {
	return terra.ReferenceAsString(ordi.ref.Append("stack_id"))
}

type opsworksRdsDbInstanceState struct {
	DbPassword       string `json:"db_password"`
	DbUser           string `json:"db_user"`
	Id               string `json:"id"`
	RdsDbInstanceArn string `json:"rds_db_instance_arn"`
	StackId          string `json:"stack_id"`
}
