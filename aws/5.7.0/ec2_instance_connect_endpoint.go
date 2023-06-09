// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	ec2instanceconnectendpoint "github.com/golingon/terraproviders/aws/5.7.0/ec2instanceconnectendpoint"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewEc2InstanceConnectEndpoint creates a new instance of [Ec2InstanceConnectEndpoint].
func NewEc2InstanceConnectEndpoint(name string, args Ec2InstanceConnectEndpointArgs) *Ec2InstanceConnectEndpoint {
	return &Ec2InstanceConnectEndpoint{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Ec2InstanceConnectEndpoint)(nil)

// Ec2InstanceConnectEndpoint represents the Terraform resource aws_ec2_instance_connect_endpoint.
type Ec2InstanceConnectEndpoint struct {
	Name      string
	Args      Ec2InstanceConnectEndpointArgs
	state     *ec2InstanceConnectEndpointState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Ec2InstanceConnectEndpoint].
func (eice *Ec2InstanceConnectEndpoint) Type() string {
	return "aws_ec2_instance_connect_endpoint"
}

// LocalName returns the local name for [Ec2InstanceConnectEndpoint].
func (eice *Ec2InstanceConnectEndpoint) LocalName() string {
	return eice.Name
}

// Configuration returns the configuration (args) for [Ec2InstanceConnectEndpoint].
func (eice *Ec2InstanceConnectEndpoint) Configuration() interface{} {
	return eice.Args
}

// DependOn is used for other resources to depend on [Ec2InstanceConnectEndpoint].
func (eice *Ec2InstanceConnectEndpoint) DependOn() terra.Reference {
	return terra.ReferenceResource(eice)
}

// Dependencies returns the list of resources [Ec2InstanceConnectEndpoint] depends_on.
func (eice *Ec2InstanceConnectEndpoint) Dependencies() terra.Dependencies {
	return eice.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Ec2InstanceConnectEndpoint].
func (eice *Ec2InstanceConnectEndpoint) LifecycleManagement() *terra.Lifecycle {
	return eice.Lifecycle
}

// Attributes returns the attributes for [Ec2InstanceConnectEndpoint].
func (eice *Ec2InstanceConnectEndpoint) Attributes() ec2InstanceConnectEndpointAttributes {
	return ec2InstanceConnectEndpointAttributes{ref: terra.ReferenceResource(eice)}
}

// ImportState imports the given attribute values into [Ec2InstanceConnectEndpoint]'s state.
func (eice *Ec2InstanceConnectEndpoint) ImportState(av io.Reader) error {
	eice.state = &ec2InstanceConnectEndpointState{}
	if err := json.NewDecoder(av).Decode(eice.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", eice.Type(), eice.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Ec2InstanceConnectEndpoint] has state.
func (eice *Ec2InstanceConnectEndpoint) State() (*ec2InstanceConnectEndpointState, bool) {
	return eice.state, eice.state != nil
}

// StateMust returns the state for [Ec2InstanceConnectEndpoint]. Panics if the state is nil.
func (eice *Ec2InstanceConnectEndpoint) StateMust() *ec2InstanceConnectEndpointState {
	if eice.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", eice.Type(), eice.LocalName()))
	}
	return eice.state
}

// Ec2InstanceConnectEndpointArgs contains the configurations for aws_ec2_instance_connect_endpoint.
type Ec2InstanceConnectEndpointArgs struct {
	// PreserveClientIp: bool, optional
	PreserveClientIp terra.BoolValue `hcl:"preserve_client_ip,attr"`
	// SecurityGroupIds: set of string, optional
	SecurityGroupIds terra.SetValue[terra.StringValue] `hcl:"security_group_ids,attr"`
	// SubnetId: string, required
	SubnetId terra.StringValue `hcl:"subnet_id,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Timeouts: optional
	Timeouts *ec2instanceconnectendpoint.Timeouts `hcl:"timeouts,block"`
}
type ec2InstanceConnectEndpointAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_ec2_instance_connect_endpoint.
func (eice ec2InstanceConnectEndpointAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(eice.ref.Append("arn"))
}

// AvailabilityZone returns a reference to field availability_zone of aws_ec2_instance_connect_endpoint.
func (eice ec2InstanceConnectEndpointAttributes) AvailabilityZone() terra.StringValue {
	return terra.ReferenceAsString(eice.ref.Append("availability_zone"))
}

// DnsName returns a reference to field dns_name of aws_ec2_instance_connect_endpoint.
func (eice ec2InstanceConnectEndpointAttributes) DnsName() terra.StringValue {
	return terra.ReferenceAsString(eice.ref.Append("dns_name"))
}

// FipsDnsName returns a reference to field fips_dns_name of aws_ec2_instance_connect_endpoint.
func (eice ec2InstanceConnectEndpointAttributes) FipsDnsName() terra.StringValue {
	return terra.ReferenceAsString(eice.ref.Append("fips_dns_name"))
}

// Id returns a reference to field id of aws_ec2_instance_connect_endpoint.
func (eice ec2InstanceConnectEndpointAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(eice.ref.Append("id"))
}

// NetworkInterfaceIds returns a reference to field network_interface_ids of aws_ec2_instance_connect_endpoint.
func (eice ec2InstanceConnectEndpointAttributes) NetworkInterfaceIds() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](eice.ref.Append("network_interface_ids"))
}

// OwnerId returns a reference to field owner_id of aws_ec2_instance_connect_endpoint.
func (eice ec2InstanceConnectEndpointAttributes) OwnerId() terra.StringValue {
	return terra.ReferenceAsString(eice.ref.Append("owner_id"))
}

// PreserveClientIp returns a reference to field preserve_client_ip of aws_ec2_instance_connect_endpoint.
func (eice ec2InstanceConnectEndpointAttributes) PreserveClientIp() terra.BoolValue {
	return terra.ReferenceAsBool(eice.ref.Append("preserve_client_ip"))
}

// SecurityGroupIds returns a reference to field security_group_ids of aws_ec2_instance_connect_endpoint.
func (eice ec2InstanceConnectEndpointAttributes) SecurityGroupIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](eice.ref.Append("security_group_ids"))
}

// SubnetId returns a reference to field subnet_id of aws_ec2_instance_connect_endpoint.
func (eice ec2InstanceConnectEndpointAttributes) SubnetId() terra.StringValue {
	return terra.ReferenceAsString(eice.ref.Append("subnet_id"))
}

// Tags returns a reference to field tags of aws_ec2_instance_connect_endpoint.
func (eice ec2InstanceConnectEndpointAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](eice.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_ec2_instance_connect_endpoint.
func (eice ec2InstanceConnectEndpointAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](eice.ref.Append("tags_all"))
}

// VpcId returns a reference to field vpc_id of aws_ec2_instance_connect_endpoint.
func (eice ec2InstanceConnectEndpointAttributes) VpcId() terra.StringValue {
	return terra.ReferenceAsString(eice.ref.Append("vpc_id"))
}

func (eice ec2InstanceConnectEndpointAttributes) Timeouts() ec2instanceconnectendpoint.TimeoutsAttributes {
	return terra.ReferenceAsSingle[ec2instanceconnectendpoint.TimeoutsAttributes](eice.ref.Append("timeouts"))
}

type ec2InstanceConnectEndpointState struct {
	Arn                 string                                    `json:"arn"`
	AvailabilityZone    string                                    `json:"availability_zone"`
	DnsName             string                                    `json:"dns_name"`
	FipsDnsName         string                                    `json:"fips_dns_name"`
	Id                  string                                    `json:"id"`
	NetworkInterfaceIds []string                                  `json:"network_interface_ids"`
	OwnerId             string                                    `json:"owner_id"`
	PreserveClientIp    bool                                      `json:"preserve_client_ip"`
	SecurityGroupIds    []string                                  `json:"security_group_ids"`
	SubnetId            string                                    `json:"subnet_id"`
	Tags                map[string]string                         `json:"tags"`
	TagsAll             map[string]string                         `json:"tags_all"`
	VpcId               string                                    `json:"vpc_id"`
	Timeouts            *ec2instanceconnectendpoint.TimeoutsState `json:"timeouts"`
}
