// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	dbproxyendpoint "github.com/golingon/terraproviders/aws/5.12.0/dbproxyendpoint"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDbProxyEndpoint creates a new instance of [DbProxyEndpoint].
func NewDbProxyEndpoint(name string, args DbProxyEndpointArgs) *DbProxyEndpoint {
	return &DbProxyEndpoint{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DbProxyEndpoint)(nil)

// DbProxyEndpoint represents the Terraform resource aws_db_proxy_endpoint.
type DbProxyEndpoint struct {
	Name      string
	Args      DbProxyEndpointArgs
	state     *dbProxyEndpointState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DbProxyEndpoint].
func (dpe *DbProxyEndpoint) Type() string {
	return "aws_db_proxy_endpoint"
}

// LocalName returns the local name for [DbProxyEndpoint].
func (dpe *DbProxyEndpoint) LocalName() string {
	return dpe.Name
}

// Configuration returns the configuration (args) for [DbProxyEndpoint].
func (dpe *DbProxyEndpoint) Configuration() interface{} {
	return dpe.Args
}

// DependOn is used for other resources to depend on [DbProxyEndpoint].
func (dpe *DbProxyEndpoint) DependOn() terra.Reference {
	return terra.ReferenceResource(dpe)
}

// Dependencies returns the list of resources [DbProxyEndpoint] depends_on.
func (dpe *DbProxyEndpoint) Dependencies() terra.Dependencies {
	return dpe.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DbProxyEndpoint].
func (dpe *DbProxyEndpoint) LifecycleManagement() *terra.Lifecycle {
	return dpe.Lifecycle
}

// Attributes returns the attributes for [DbProxyEndpoint].
func (dpe *DbProxyEndpoint) Attributes() dbProxyEndpointAttributes {
	return dbProxyEndpointAttributes{ref: terra.ReferenceResource(dpe)}
}

// ImportState imports the given attribute values into [DbProxyEndpoint]'s state.
func (dpe *DbProxyEndpoint) ImportState(av io.Reader) error {
	dpe.state = &dbProxyEndpointState{}
	if err := json.NewDecoder(av).Decode(dpe.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dpe.Type(), dpe.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DbProxyEndpoint] has state.
func (dpe *DbProxyEndpoint) State() (*dbProxyEndpointState, bool) {
	return dpe.state, dpe.state != nil
}

// StateMust returns the state for [DbProxyEndpoint]. Panics if the state is nil.
func (dpe *DbProxyEndpoint) StateMust() *dbProxyEndpointState {
	if dpe.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dpe.Type(), dpe.LocalName()))
	}
	return dpe.state
}

// DbProxyEndpointArgs contains the configurations for aws_db_proxy_endpoint.
type DbProxyEndpointArgs struct {
	// DbProxyEndpointName: string, required
	DbProxyEndpointName terra.StringValue `hcl:"db_proxy_endpoint_name,attr" validate:"required"`
	// DbProxyName: string, required
	DbProxyName terra.StringValue `hcl:"db_proxy_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// TargetRole: string, optional
	TargetRole terra.StringValue `hcl:"target_role,attr"`
	// VpcSecurityGroupIds: set of string, optional
	VpcSecurityGroupIds terra.SetValue[terra.StringValue] `hcl:"vpc_security_group_ids,attr"`
	// VpcSubnetIds: set of string, required
	VpcSubnetIds terra.SetValue[terra.StringValue] `hcl:"vpc_subnet_ids,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *dbproxyendpoint.Timeouts `hcl:"timeouts,block"`
}
type dbProxyEndpointAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_db_proxy_endpoint.
func (dpe dbProxyEndpointAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(dpe.ref.Append("arn"))
}

// DbProxyEndpointName returns a reference to field db_proxy_endpoint_name of aws_db_proxy_endpoint.
func (dpe dbProxyEndpointAttributes) DbProxyEndpointName() terra.StringValue {
	return terra.ReferenceAsString(dpe.ref.Append("db_proxy_endpoint_name"))
}

// DbProxyName returns a reference to field db_proxy_name of aws_db_proxy_endpoint.
func (dpe dbProxyEndpointAttributes) DbProxyName() terra.StringValue {
	return terra.ReferenceAsString(dpe.ref.Append("db_proxy_name"))
}

// Endpoint returns a reference to field endpoint of aws_db_proxy_endpoint.
func (dpe dbProxyEndpointAttributes) Endpoint() terra.StringValue {
	return terra.ReferenceAsString(dpe.ref.Append("endpoint"))
}

// Id returns a reference to field id of aws_db_proxy_endpoint.
func (dpe dbProxyEndpointAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dpe.ref.Append("id"))
}

// IsDefault returns a reference to field is_default of aws_db_proxy_endpoint.
func (dpe dbProxyEndpointAttributes) IsDefault() terra.BoolValue {
	return terra.ReferenceAsBool(dpe.ref.Append("is_default"))
}

// Tags returns a reference to field tags of aws_db_proxy_endpoint.
func (dpe dbProxyEndpointAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dpe.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_db_proxy_endpoint.
func (dpe dbProxyEndpointAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dpe.ref.Append("tags_all"))
}

// TargetRole returns a reference to field target_role of aws_db_proxy_endpoint.
func (dpe dbProxyEndpointAttributes) TargetRole() terra.StringValue {
	return terra.ReferenceAsString(dpe.ref.Append("target_role"))
}

// VpcId returns a reference to field vpc_id of aws_db_proxy_endpoint.
func (dpe dbProxyEndpointAttributes) VpcId() terra.StringValue {
	return terra.ReferenceAsString(dpe.ref.Append("vpc_id"))
}

// VpcSecurityGroupIds returns a reference to field vpc_security_group_ids of aws_db_proxy_endpoint.
func (dpe dbProxyEndpointAttributes) VpcSecurityGroupIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](dpe.ref.Append("vpc_security_group_ids"))
}

// VpcSubnetIds returns a reference to field vpc_subnet_ids of aws_db_proxy_endpoint.
func (dpe dbProxyEndpointAttributes) VpcSubnetIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](dpe.ref.Append("vpc_subnet_ids"))
}

func (dpe dbProxyEndpointAttributes) Timeouts() dbproxyendpoint.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dbproxyendpoint.TimeoutsAttributes](dpe.ref.Append("timeouts"))
}

type dbProxyEndpointState struct {
	Arn                 string                         `json:"arn"`
	DbProxyEndpointName string                         `json:"db_proxy_endpoint_name"`
	DbProxyName         string                         `json:"db_proxy_name"`
	Endpoint            string                         `json:"endpoint"`
	Id                  string                         `json:"id"`
	IsDefault           bool                           `json:"is_default"`
	Tags                map[string]string              `json:"tags"`
	TagsAll             map[string]string              `json:"tags_all"`
	TargetRole          string                         `json:"target_role"`
	VpcId               string                         `json:"vpc_id"`
	VpcSecurityGroupIds []string                       `json:"vpc_security_group_ids"`
	VpcSubnetIds        []string                       `json:"vpc_subnet_ids"`
	Timeouts            *dbproxyendpoint.TimeoutsState `json:"timeouts"`
}