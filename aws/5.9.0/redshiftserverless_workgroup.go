// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	redshiftserverlessworkgroup "github.com/golingon/terraproviders/aws/5.9.0/redshiftserverlessworkgroup"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewRedshiftserverlessWorkgroup creates a new instance of [RedshiftserverlessWorkgroup].
func NewRedshiftserverlessWorkgroup(name string, args RedshiftserverlessWorkgroupArgs) *RedshiftserverlessWorkgroup {
	return &RedshiftserverlessWorkgroup{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*RedshiftserverlessWorkgroup)(nil)

// RedshiftserverlessWorkgroup represents the Terraform resource aws_redshiftserverless_workgroup.
type RedshiftserverlessWorkgroup struct {
	Name      string
	Args      RedshiftserverlessWorkgroupArgs
	state     *redshiftserverlessWorkgroupState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [RedshiftserverlessWorkgroup].
func (rw *RedshiftserverlessWorkgroup) Type() string {
	return "aws_redshiftserverless_workgroup"
}

// LocalName returns the local name for [RedshiftserverlessWorkgroup].
func (rw *RedshiftserverlessWorkgroup) LocalName() string {
	return rw.Name
}

// Configuration returns the configuration (args) for [RedshiftserverlessWorkgroup].
func (rw *RedshiftserverlessWorkgroup) Configuration() interface{} {
	return rw.Args
}

// DependOn is used for other resources to depend on [RedshiftserverlessWorkgroup].
func (rw *RedshiftserverlessWorkgroup) DependOn() terra.Reference {
	return terra.ReferenceResource(rw)
}

// Dependencies returns the list of resources [RedshiftserverlessWorkgroup] depends_on.
func (rw *RedshiftserverlessWorkgroup) Dependencies() terra.Dependencies {
	return rw.DependsOn
}

// LifecycleManagement returns the lifecycle block for [RedshiftserverlessWorkgroup].
func (rw *RedshiftserverlessWorkgroup) LifecycleManagement() *terra.Lifecycle {
	return rw.Lifecycle
}

// Attributes returns the attributes for [RedshiftserverlessWorkgroup].
func (rw *RedshiftserverlessWorkgroup) Attributes() redshiftserverlessWorkgroupAttributes {
	return redshiftserverlessWorkgroupAttributes{ref: terra.ReferenceResource(rw)}
}

// ImportState imports the given attribute values into [RedshiftserverlessWorkgroup]'s state.
func (rw *RedshiftserverlessWorkgroup) ImportState(av io.Reader) error {
	rw.state = &redshiftserverlessWorkgroupState{}
	if err := json.NewDecoder(av).Decode(rw.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", rw.Type(), rw.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [RedshiftserverlessWorkgroup] has state.
func (rw *RedshiftserverlessWorkgroup) State() (*redshiftserverlessWorkgroupState, bool) {
	return rw.state, rw.state != nil
}

// StateMust returns the state for [RedshiftserverlessWorkgroup]. Panics if the state is nil.
func (rw *RedshiftserverlessWorkgroup) StateMust() *redshiftserverlessWorkgroupState {
	if rw.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", rw.Type(), rw.LocalName()))
	}
	return rw.state
}

// RedshiftserverlessWorkgroupArgs contains the configurations for aws_redshiftserverless_workgroup.
type RedshiftserverlessWorkgroupArgs struct {
	// BaseCapacity: number, optional
	BaseCapacity terra.NumberValue `hcl:"base_capacity,attr"`
	// EnhancedVpcRouting: bool, optional
	EnhancedVpcRouting terra.BoolValue `hcl:"enhanced_vpc_routing,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// NamespaceName: string, required
	NamespaceName terra.StringValue `hcl:"namespace_name,attr" validate:"required"`
	// PubliclyAccessible: bool, optional
	PubliclyAccessible terra.BoolValue `hcl:"publicly_accessible,attr"`
	// SecurityGroupIds: set of string, optional
	SecurityGroupIds terra.SetValue[terra.StringValue] `hcl:"security_group_ids,attr"`
	// SubnetIds: set of string, optional
	SubnetIds terra.SetValue[terra.StringValue] `hcl:"subnet_ids,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// WorkgroupName: string, required
	WorkgroupName terra.StringValue `hcl:"workgroup_name,attr" validate:"required"`
	// Endpoint: min=0
	Endpoint []redshiftserverlessworkgroup.Endpoint `hcl:"endpoint,block" validate:"min=0"`
	// ConfigParameter: min=0
	ConfigParameter []redshiftserverlessworkgroup.ConfigParameter `hcl:"config_parameter,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *redshiftserverlessworkgroup.Timeouts `hcl:"timeouts,block"`
}
type redshiftserverlessWorkgroupAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_redshiftserverless_workgroup.
func (rw redshiftserverlessWorkgroupAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(rw.ref.Append("arn"))
}

// BaseCapacity returns a reference to field base_capacity of aws_redshiftserverless_workgroup.
func (rw redshiftserverlessWorkgroupAttributes) BaseCapacity() terra.NumberValue {
	return terra.ReferenceAsNumber(rw.ref.Append("base_capacity"))
}

// EnhancedVpcRouting returns a reference to field enhanced_vpc_routing of aws_redshiftserverless_workgroup.
func (rw redshiftserverlessWorkgroupAttributes) EnhancedVpcRouting() terra.BoolValue {
	return terra.ReferenceAsBool(rw.ref.Append("enhanced_vpc_routing"))
}

// Id returns a reference to field id of aws_redshiftserverless_workgroup.
func (rw redshiftserverlessWorkgroupAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(rw.ref.Append("id"))
}

// NamespaceName returns a reference to field namespace_name of aws_redshiftserverless_workgroup.
func (rw redshiftserverlessWorkgroupAttributes) NamespaceName() terra.StringValue {
	return terra.ReferenceAsString(rw.ref.Append("namespace_name"))
}

// PubliclyAccessible returns a reference to field publicly_accessible of aws_redshiftserverless_workgroup.
func (rw redshiftserverlessWorkgroupAttributes) PubliclyAccessible() terra.BoolValue {
	return terra.ReferenceAsBool(rw.ref.Append("publicly_accessible"))
}

// SecurityGroupIds returns a reference to field security_group_ids of aws_redshiftserverless_workgroup.
func (rw redshiftserverlessWorkgroupAttributes) SecurityGroupIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](rw.ref.Append("security_group_ids"))
}

// SubnetIds returns a reference to field subnet_ids of aws_redshiftserverless_workgroup.
func (rw redshiftserverlessWorkgroupAttributes) SubnetIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](rw.ref.Append("subnet_ids"))
}

// Tags returns a reference to field tags of aws_redshiftserverless_workgroup.
func (rw redshiftserverlessWorkgroupAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](rw.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_redshiftserverless_workgroup.
func (rw redshiftserverlessWorkgroupAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](rw.ref.Append("tags_all"))
}

// WorkgroupId returns a reference to field workgroup_id of aws_redshiftserverless_workgroup.
func (rw redshiftserverlessWorkgroupAttributes) WorkgroupId() terra.StringValue {
	return terra.ReferenceAsString(rw.ref.Append("workgroup_id"))
}

// WorkgroupName returns a reference to field workgroup_name of aws_redshiftserverless_workgroup.
func (rw redshiftserverlessWorkgroupAttributes) WorkgroupName() terra.StringValue {
	return terra.ReferenceAsString(rw.ref.Append("workgroup_name"))
}

func (rw redshiftserverlessWorkgroupAttributes) Endpoint() terra.ListValue[redshiftserverlessworkgroup.EndpointAttributes] {
	return terra.ReferenceAsList[redshiftserverlessworkgroup.EndpointAttributes](rw.ref.Append("endpoint"))
}

func (rw redshiftserverlessWorkgroupAttributes) ConfigParameter() terra.SetValue[redshiftserverlessworkgroup.ConfigParameterAttributes] {
	return terra.ReferenceAsSet[redshiftserverlessworkgroup.ConfigParameterAttributes](rw.ref.Append("config_parameter"))
}

func (rw redshiftserverlessWorkgroupAttributes) Timeouts() redshiftserverlessworkgroup.TimeoutsAttributes {
	return terra.ReferenceAsSingle[redshiftserverlessworkgroup.TimeoutsAttributes](rw.ref.Append("timeouts"))
}

type redshiftserverlessWorkgroupState struct {
	Arn                string                                             `json:"arn"`
	BaseCapacity       float64                                            `json:"base_capacity"`
	EnhancedVpcRouting bool                                               `json:"enhanced_vpc_routing"`
	Id                 string                                             `json:"id"`
	NamespaceName      string                                             `json:"namespace_name"`
	PubliclyAccessible bool                                               `json:"publicly_accessible"`
	SecurityGroupIds   []string                                           `json:"security_group_ids"`
	SubnetIds          []string                                           `json:"subnet_ids"`
	Tags               map[string]string                                  `json:"tags"`
	TagsAll            map[string]string                                  `json:"tags_all"`
	WorkgroupId        string                                             `json:"workgroup_id"`
	WorkgroupName      string                                             `json:"workgroup_name"`
	Endpoint           []redshiftserverlessworkgroup.EndpointState        `json:"endpoint"`
	ConfigParameter    []redshiftserverlessworkgroup.ConfigParameterState `json:"config_parameter"`
	Timeouts           *redshiftserverlessworkgroup.TimeoutsState         `json:"timeouts"`
}
