// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	redshiftserverlessendpointaccess "github.com/golingon/terraproviders/aws/5.12.0/redshiftserverlessendpointaccess"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewRedshiftserverlessEndpointAccess creates a new instance of [RedshiftserverlessEndpointAccess].
func NewRedshiftserverlessEndpointAccess(name string, args RedshiftserverlessEndpointAccessArgs) *RedshiftserverlessEndpointAccess {
	return &RedshiftserverlessEndpointAccess{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*RedshiftserverlessEndpointAccess)(nil)

// RedshiftserverlessEndpointAccess represents the Terraform resource aws_redshiftserverless_endpoint_access.
type RedshiftserverlessEndpointAccess struct {
	Name      string
	Args      RedshiftserverlessEndpointAccessArgs
	state     *redshiftserverlessEndpointAccessState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [RedshiftserverlessEndpointAccess].
func (rea *RedshiftserverlessEndpointAccess) Type() string {
	return "aws_redshiftserverless_endpoint_access"
}

// LocalName returns the local name for [RedshiftserverlessEndpointAccess].
func (rea *RedshiftserverlessEndpointAccess) LocalName() string {
	return rea.Name
}

// Configuration returns the configuration (args) for [RedshiftserverlessEndpointAccess].
func (rea *RedshiftserverlessEndpointAccess) Configuration() interface{} {
	return rea.Args
}

// DependOn is used for other resources to depend on [RedshiftserverlessEndpointAccess].
func (rea *RedshiftserverlessEndpointAccess) DependOn() terra.Reference {
	return terra.ReferenceResource(rea)
}

// Dependencies returns the list of resources [RedshiftserverlessEndpointAccess] depends_on.
func (rea *RedshiftserverlessEndpointAccess) Dependencies() terra.Dependencies {
	return rea.DependsOn
}

// LifecycleManagement returns the lifecycle block for [RedshiftserverlessEndpointAccess].
func (rea *RedshiftserverlessEndpointAccess) LifecycleManagement() *terra.Lifecycle {
	return rea.Lifecycle
}

// Attributes returns the attributes for [RedshiftserverlessEndpointAccess].
func (rea *RedshiftserverlessEndpointAccess) Attributes() redshiftserverlessEndpointAccessAttributes {
	return redshiftserverlessEndpointAccessAttributes{ref: terra.ReferenceResource(rea)}
}

// ImportState imports the given attribute values into [RedshiftserverlessEndpointAccess]'s state.
func (rea *RedshiftserverlessEndpointAccess) ImportState(av io.Reader) error {
	rea.state = &redshiftserverlessEndpointAccessState{}
	if err := json.NewDecoder(av).Decode(rea.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", rea.Type(), rea.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [RedshiftserverlessEndpointAccess] has state.
func (rea *RedshiftserverlessEndpointAccess) State() (*redshiftserverlessEndpointAccessState, bool) {
	return rea.state, rea.state != nil
}

// StateMust returns the state for [RedshiftserverlessEndpointAccess]. Panics if the state is nil.
func (rea *RedshiftserverlessEndpointAccess) StateMust() *redshiftserverlessEndpointAccessState {
	if rea.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", rea.Type(), rea.LocalName()))
	}
	return rea.state
}

// RedshiftserverlessEndpointAccessArgs contains the configurations for aws_redshiftserverless_endpoint_access.
type RedshiftserverlessEndpointAccessArgs struct {
	// EndpointName: string, required
	EndpointName terra.StringValue `hcl:"endpoint_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// SubnetIds: set of string, required
	SubnetIds terra.SetValue[terra.StringValue] `hcl:"subnet_ids,attr" validate:"required"`
	// VpcSecurityGroupIds: set of string, optional
	VpcSecurityGroupIds terra.SetValue[terra.StringValue] `hcl:"vpc_security_group_ids,attr"`
	// WorkgroupName: string, required
	WorkgroupName terra.StringValue `hcl:"workgroup_name,attr" validate:"required"`
	// VpcEndpoint: min=0
	VpcEndpoint []redshiftserverlessendpointaccess.VpcEndpoint `hcl:"vpc_endpoint,block" validate:"min=0"`
}
type redshiftserverlessEndpointAccessAttributes struct {
	ref terra.Reference
}

// Address returns a reference to field address of aws_redshiftserverless_endpoint_access.
func (rea redshiftserverlessEndpointAccessAttributes) Address() terra.StringValue {
	return terra.ReferenceAsString(rea.ref.Append("address"))
}

// Arn returns a reference to field arn of aws_redshiftserverless_endpoint_access.
func (rea redshiftserverlessEndpointAccessAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(rea.ref.Append("arn"))
}

// EndpointName returns a reference to field endpoint_name of aws_redshiftserverless_endpoint_access.
func (rea redshiftserverlessEndpointAccessAttributes) EndpointName() terra.StringValue {
	return terra.ReferenceAsString(rea.ref.Append("endpoint_name"))
}

// Id returns a reference to field id of aws_redshiftserverless_endpoint_access.
func (rea redshiftserverlessEndpointAccessAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(rea.ref.Append("id"))
}

// Port returns a reference to field port of aws_redshiftserverless_endpoint_access.
func (rea redshiftserverlessEndpointAccessAttributes) Port() terra.NumberValue {
	return terra.ReferenceAsNumber(rea.ref.Append("port"))
}

// SubnetIds returns a reference to field subnet_ids of aws_redshiftserverless_endpoint_access.
func (rea redshiftserverlessEndpointAccessAttributes) SubnetIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](rea.ref.Append("subnet_ids"))
}

// VpcSecurityGroupIds returns a reference to field vpc_security_group_ids of aws_redshiftserverless_endpoint_access.
func (rea redshiftserverlessEndpointAccessAttributes) VpcSecurityGroupIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](rea.ref.Append("vpc_security_group_ids"))
}

// WorkgroupName returns a reference to field workgroup_name of aws_redshiftserverless_endpoint_access.
func (rea redshiftserverlessEndpointAccessAttributes) WorkgroupName() terra.StringValue {
	return terra.ReferenceAsString(rea.ref.Append("workgroup_name"))
}

func (rea redshiftserverlessEndpointAccessAttributes) VpcEndpoint() terra.ListValue[redshiftserverlessendpointaccess.VpcEndpointAttributes] {
	return terra.ReferenceAsList[redshiftserverlessendpointaccess.VpcEndpointAttributes](rea.ref.Append("vpc_endpoint"))
}

type redshiftserverlessEndpointAccessState struct {
	Address             string                                              `json:"address"`
	Arn                 string                                              `json:"arn"`
	EndpointName        string                                              `json:"endpoint_name"`
	Id                  string                                              `json:"id"`
	Port                float64                                             `json:"port"`
	SubnetIds           []string                                            `json:"subnet_ids"`
	VpcSecurityGroupIds []string                                            `json:"vpc_security_group_ids"`
	WorkgroupName       string                                              `json:"workgroup_name"`
	VpcEndpoint         []redshiftserverlessendpointaccess.VpcEndpointState `json:"vpc_endpoint"`
}
