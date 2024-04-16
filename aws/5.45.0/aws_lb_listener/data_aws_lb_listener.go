// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_lb_listener

import "github.com/golingon/lingon/pkg/terra"

// Data creates a new instance of [DataSource].
func Data(name string, args DataArgs) *DataSource {
	return &DataSource{
		Args: args,
		Name: name,
	}
}

var _ terra.DataSource = (*DataSource)(nil)

// DataSource represents the Terraform data resource aws_lb_listener.
type DataSource struct {
	Name string
	Args DataArgs
}

// DataSource returns the Terraform object type for [DataSource].
func (all *DataSource) DataSource() string {
	return "aws_lb_listener"
}

// LocalName returns the local name for [DataSource].
func (all *DataSource) LocalName() string {
	return all.Name
}

// Configuration returns the configuration (args) for [DataSource].
func (all *DataSource) Configuration() interface{} {
	return all.Args
}

// Attributes returns the attributes for [DataSource].
func (all *DataSource) Attributes() dataAwsLbListenerAttributes {
	return dataAwsLbListenerAttributes{ref: terra.ReferenceDataSource(all)}
}

// DataArgs contains the configurations for aws_lb_listener.
type DataArgs struct {
	// Arn: string, optional
	Arn terra.StringValue `hcl:"arn,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// LoadBalancerArn: string, optional
	LoadBalancerArn terra.StringValue `hcl:"load_balancer_arn,attr"`
	// Port: number, optional
	Port terra.NumberValue `hcl:"port,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Timeouts: optional
	Timeouts *DataTimeouts `hcl:"timeouts,block"`
}

type dataAwsLbListenerAttributes struct {
	ref terra.Reference
}

// AlpnPolicy returns a reference to field alpn_policy of aws_lb_listener.
func (all dataAwsLbListenerAttributes) AlpnPolicy() terra.StringValue {
	return terra.ReferenceAsString(all.ref.Append("alpn_policy"))
}

// Arn returns a reference to field arn of aws_lb_listener.
func (all dataAwsLbListenerAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(all.ref.Append("arn"))
}

// CertificateArn returns a reference to field certificate_arn of aws_lb_listener.
func (all dataAwsLbListenerAttributes) CertificateArn() terra.StringValue {
	return terra.ReferenceAsString(all.ref.Append("certificate_arn"))
}

// Id returns a reference to field id of aws_lb_listener.
func (all dataAwsLbListenerAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(all.ref.Append("id"))
}

// LoadBalancerArn returns a reference to field load_balancer_arn of aws_lb_listener.
func (all dataAwsLbListenerAttributes) LoadBalancerArn() terra.StringValue {
	return terra.ReferenceAsString(all.ref.Append("load_balancer_arn"))
}

// Port returns a reference to field port of aws_lb_listener.
func (all dataAwsLbListenerAttributes) Port() terra.NumberValue {
	return terra.ReferenceAsNumber(all.ref.Append("port"))
}

// Protocol returns a reference to field protocol of aws_lb_listener.
func (all dataAwsLbListenerAttributes) Protocol() terra.StringValue {
	return terra.ReferenceAsString(all.ref.Append("protocol"))
}

// SslPolicy returns a reference to field ssl_policy of aws_lb_listener.
func (all dataAwsLbListenerAttributes) SslPolicy() terra.StringValue {
	return terra.ReferenceAsString(all.ref.Append("ssl_policy"))
}

// Tags returns a reference to field tags of aws_lb_listener.
func (all dataAwsLbListenerAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](all.ref.Append("tags"))
}

func (all dataAwsLbListenerAttributes) DefaultAction() terra.ListValue[DataDefaultActionAttributes] {
	return terra.ReferenceAsList[DataDefaultActionAttributes](all.ref.Append("default_action"))
}

func (all dataAwsLbListenerAttributes) MutualAuthentication() terra.ListValue[DataMutualAuthenticationAttributes] {
	return terra.ReferenceAsList[DataMutualAuthenticationAttributes](all.ref.Append("mutual_authentication"))
}

func (all dataAwsLbListenerAttributes) Timeouts() DataTimeoutsAttributes {
	return terra.ReferenceAsSingle[DataTimeoutsAttributes](all.ref.Append("timeouts"))
}
