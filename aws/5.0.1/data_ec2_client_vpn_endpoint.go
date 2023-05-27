// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	dataec2clientvpnendpoint "github.com/golingon/terraproviders/aws/5.0.1/dataec2clientvpnendpoint"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataEc2ClientVpnEndpoint creates a new instance of [DataEc2ClientVpnEndpoint].
func NewDataEc2ClientVpnEndpoint(name string, args DataEc2ClientVpnEndpointArgs) *DataEc2ClientVpnEndpoint {
	return &DataEc2ClientVpnEndpoint{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataEc2ClientVpnEndpoint)(nil)

// DataEc2ClientVpnEndpoint represents the Terraform data resource aws_ec2_client_vpn_endpoint.
type DataEc2ClientVpnEndpoint struct {
	Name string
	Args DataEc2ClientVpnEndpointArgs
}

// DataSource returns the Terraform object type for [DataEc2ClientVpnEndpoint].
func (ecve *DataEc2ClientVpnEndpoint) DataSource() string {
	return "aws_ec2_client_vpn_endpoint"
}

// LocalName returns the local name for [DataEc2ClientVpnEndpoint].
func (ecve *DataEc2ClientVpnEndpoint) LocalName() string {
	return ecve.Name
}

// Configuration returns the configuration (args) for [DataEc2ClientVpnEndpoint].
func (ecve *DataEc2ClientVpnEndpoint) Configuration() interface{} {
	return ecve.Args
}

// Attributes returns the attributes for [DataEc2ClientVpnEndpoint].
func (ecve *DataEc2ClientVpnEndpoint) Attributes() dataEc2ClientVpnEndpointAttributes {
	return dataEc2ClientVpnEndpointAttributes{ref: terra.ReferenceDataResource(ecve)}
}

// DataEc2ClientVpnEndpointArgs contains the configurations for aws_ec2_client_vpn_endpoint.
type DataEc2ClientVpnEndpointArgs struct {
	// ClientVpnEndpointId: string, optional
	ClientVpnEndpointId terra.StringValue `hcl:"client_vpn_endpoint_id,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// AuthenticationOptions: min=0
	AuthenticationOptions []dataec2clientvpnendpoint.AuthenticationOptions `hcl:"authentication_options,block" validate:"min=0"`
	// ClientConnectOptions: min=0
	ClientConnectOptions []dataec2clientvpnendpoint.ClientConnectOptions `hcl:"client_connect_options,block" validate:"min=0"`
	// ClientLoginBannerOptions: min=0
	ClientLoginBannerOptions []dataec2clientvpnendpoint.ClientLoginBannerOptions `hcl:"client_login_banner_options,block" validate:"min=0"`
	// ConnectionLogOptions: min=0
	ConnectionLogOptions []dataec2clientvpnendpoint.ConnectionLogOptions `hcl:"connection_log_options,block" validate:"min=0"`
	// Filter: min=0
	Filter []dataec2clientvpnendpoint.Filter `hcl:"filter,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *dataec2clientvpnendpoint.Timeouts `hcl:"timeouts,block"`
}
type dataEc2ClientVpnEndpointAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_ec2_client_vpn_endpoint.
func (ecve dataEc2ClientVpnEndpointAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(ecve.ref.Append("arn"))
}

// ClientCidrBlock returns a reference to field client_cidr_block of aws_ec2_client_vpn_endpoint.
func (ecve dataEc2ClientVpnEndpointAttributes) ClientCidrBlock() terra.StringValue {
	return terra.ReferenceAsString(ecve.ref.Append("client_cidr_block"))
}

// ClientVpnEndpointId returns a reference to field client_vpn_endpoint_id of aws_ec2_client_vpn_endpoint.
func (ecve dataEc2ClientVpnEndpointAttributes) ClientVpnEndpointId() terra.StringValue {
	return terra.ReferenceAsString(ecve.ref.Append("client_vpn_endpoint_id"))
}

// Description returns a reference to field description of aws_ec2_client_vpn_endpoint.
func (ecve dataEc2ClientVpnEndpointAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(ecve.ref.Append("description"))
}

// DnsName returns a reference to field dns_name of aws_ec2_client_vpn_endpoint.
func (ecve dataEc2ClientVpnEndpointAttributes) DnsName() terra.StringValue {
	return terra.ReferenceAsString(ecve.ref.Append("dns_name"))
}

// DnsServers returns a reference to field dns_servers of aws_ec2_client_vpn_endpoint.
func (ecve dataEc2ClientVpnEndpointAttributes) DnsServers() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](ecve.ref.Append("dns_servers"))
}

// Id returns a reference to field id of aws_ec2_client_vpn_endpoint.
func (ecve dataEc2ClientVpnEndpointAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ecve.ref.Append("id"))
}

// SecurityGroupIds returns a reference to field security_group_ids of aws_ec2_client_vpn_endpoint.
func (ecve dataEc2ClientVpnEndpointAttributes) SecurityGroupIds() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](ecve.ref.Append("security_group_ids"))
}

// SelfServicePortal returns a reference to field self_service_portal of aws_ec2_client_vpn_endpoint.
func (ecve dataEc2ClientVpnEndpointAttributes) SelfServicePortal() terra.StringValue {
	return terra.ReferenceAsString(ecve.ref.Append("self_service_portal"))
}

// ServerCertificateArn returns a reference to field server_certificate_arn of aws_ec2_client_vpn_endpoint.
func (ecve dataEc2ClientVpnEndpointAttributes) ServerCertificateArn() terra.StringValue {
	return terra.ReferenceAsString(ecve.ref.Append("server_certificate_arn"))
}

// SessionTimeoutHours returns a reference to field session_timeout_hours of aws_ec2_client_vpn_endpoint.
func (ecve dataEc2ClientVpnEndpointAttributes) SessionTimeoutHours() terra.NumberValue {
	return terra.ReferenceAsNumber(ecve.ref.Append("session_timeout_hours"))
}

// SplitTunnel returns a reference to field split_tunnel of aws_ec2_client_vpn_endpoint.
func (ecve dataEc2ClientVpnEndpointAttributes) SplitTunnel() terra.BoolValue {
	return terra.ReferenceAsBool(ecve.ref.Append("split_tunnel"))
}

// Tags returns a reference to field tags of aws_ec2_client_vpn_endpoint.
func (ecve dataEc2ClientVpnEndpointAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ecve.ref.Append("tags"))
}

// TransportProtocol returns a reference to field transport_protocol of aws_ec2_client_vpn_endpoint.
func (ecve dataEc2ClientVpnEndpointAttributes) TransportProtocol() terra.StringValue {
	return terra.ReferenceAsString(ecve.ref.Append("transport_protocol"))
}

// VpcId returns a reference to field vpc_id of aws_ec2_client_vpn_endpoint.
func (ecve dataEc2ClientVpnEndpointAttributes) VpcId() terra.StringValue {
	return terra.ReferenceAsString(ecve.ref.Append("vpc_id"))
}

// VpnPort returns a reference to field vpn_port of aws_ec2_client_vpn_endpoint.
func (ecve dataEc2ClientVpnEndpointAttributes) VpnPort() terra.NumberValue {
	return terra.ReferenceAsNumber(ecve.ref.Append("vpn_port"))
}

func (ecve dataEc2ClientVpnEndpointAttributes) AuthenticationOptions() terra.ListValue[dataec2clientvpnendpoint.AuthenticationOptionsAttributes] {
	return terra.ReferenceAsList[dataec2clientvpnendpoint.AuthenticationOptionsAttributes](ecve.ref.Append("authentication_options"))
}

func (ecve dataEc2ClientVpnEndpointAttributes) ClientConnectOptions() terra.ListValue[dataec2clientvpnendpoint.ClientConnectOptionsAttributes] {
	return terra.ReferenceAsList[dataec2clientvpnendpoint.ClientConnectOptionsAttributes](ecve.ref.Append("client_connect_options"))
}

func (ecve dataEc2ClientVpnEndpointAttributes) ClientLoginBannerOptions() terra.ListValue[dataec2clientvpnendpoint.ClientLoginBannerOptionsAttributes] {
	return terra.ReferenceAsList[dataec2clientvpnendpoint.ClientLoginBannerOptionsAttributes](ecve.ref.Append("client_login_banner_options"))
}

func (ecve dataEc2ClientVpnEndpointAttributes) ConnectionLogOptions() terra.ListValue[dataec2clientvpnendpoint.ConnectionLogOptionsAttributes] {
	return terra.ReferenceAsList[dataec2clientvpnendpoint.ConnectionLogOptionsAttributes](ecve.ref.Append("connection_log_options"))
}

func (ecve dataEc2ClientVpnEndpointAttributes) Filter() terra.SetValue[dataec2clientvpnendpoint.FilterAttributes] {
	return terra.ReferenceAsSet[dataec2clientvpnendpoint.FilterAttributes](ecve.ref.Append("filter"))
}

func (ecve dataEc2ClientVpnEndpointAttributes) Timeouts() dataec2clientvpnendpoint.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dataec2clientvpnendpoint.TimeoutsAttributes](ecve.ref.Append("timeouts"))
}
