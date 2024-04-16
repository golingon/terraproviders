// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_opsworks_rds_db_instance

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// New creates a new instance of [Resource].
func New(name string, args Args) *Resource {
	return &Resource{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Resource)(nil)

// Resource represents the Terraform resource aws_opsworks_rds_db_instance.
type Resource struct {
	Name      string
	Args      Args
	state     *awsOpsworksRdsDbInstanceState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (aordi *Resource) Type() string {
	return "aws_opsworks_rds_db_instance"
}

// LocalName returns the local name for [Resource].
func (aordi *Resource) LocalName() string {
	return aordi.Name
}

// Configuration returns the configuration (args) for [Resource].
func (aordi *Resource) Configuration() interface{} {
	return aordi.Args
}

// DependOn is used for other resources to depend on [Resource].
func (aordi *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(aordi)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (aordi *Resource) Dependencies() terra.Dependencies {
	return aordi.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (aordi *Resource) LifecycleManagement() *terra.Lifecycle {
	return aordi.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (aordi *Resource) Attributes() awsOpsworksRdsDbInstanceAttributes {
	return awsOpsworksRdsDbInstanceAttributes{ref: terra.ReferenceResource(aordi)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (aordi *Resource) ImportState(state io.Reader) error {
	aordi.state = &awsOpsworksRdsDbInstanceState{}
	if err := json.NewDecoder(state).Decode(aordi.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", aordi.Type(), aordi.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (aordi *Resource) State() (*awsOpsworksRdsDbInstanceState, bool) {
	return aordi.state, aordi.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (aordi *Resource) StateMust() *awsOpsworksRdsDbInstanceState {
	if aordi.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", aordi.Type(), aordi.LocalName()))
	}
	return aordi.state
}

// Args contains the configurations for aws_opsworks_rds_db_instance.
type Args struct {
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

type awsOpsworksRdsDbInstanceAttributes struct {
	ref terra.Reference
}

// DbPassword returns a reference to field db_password of aws_opsworks_rds_db_instance.
func (aordi awsOpsworksRdsDbInstanceAttributes) DbPassword() terra.StringValue {
	return terra.ReferenceAsString(aordi.ref.Append("db_password"))
}

// DbUser returns a reference to field db_user of aws_opsworks_rds_db_instance.
func (aordi awsOpsworksRdsDbInstanceAttributes) DbUser() terra.StringValue {
	return terra.ReferenceAsString(aordi.ref.Append("db_user"))
}

// Id returns a reference to field id of aws_opsworks_rds_db_instance.
func (aordi awsOpsworksRdsDbInstanceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(aordi.ref.Append("id"))
}

// RdsDbInstanceArn returns a reference to field rds_db_instance_arn of aws_opsworks_rds_db_instance.
func (aordi awsOpsworksRdsDbInstanceAttributes) RdsDbInstanceArn() terra.StringValue {
	return terra.ReferenceAsString(aordi.ref.Append("rds_db_instance_arn"))
}

// StackId returns a reference to field stack_id of aws_opsworks_rds_db_instance.
func (aordi awsOpsworksRdsDbInstanceAttributes) StackId() terra.StringValue {
	return terra.ReferenceAsString(aordi.ref.Append("stack_id"))
}

type awsOpsworksRdsDbInstanceState struct {
	DbPassword       string `json:"db_password"`
	DbUser           string `json:"db_user"`
	Id               string `json:"id"`
	RdsDbInstanceArn string `json:"rds_db_instance_arn"`
	StackId          string `json:"stack_id"`
}
