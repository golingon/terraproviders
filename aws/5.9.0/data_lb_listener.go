// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	datalblistener "github.com/golingon/terraproviders/aws/5.9.0/datalblistener"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataLbListener creates a new instance of [DataLbListener].
func NewDataLbListener(name string, args DataLbListenerArgs) *DataLbListener {
	return &DataLbListener{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataLbListener)(nil)

// DataLbListener represents the Terraform data resource aws_lb_listener.
type DataLbListener struct {
	Name string
	Args DataLbListenerArgs
}

// DataSource returns the Terraform object type for [DataLbListener].
func (ll *DataLbListener) DataSource() string {
	return "aws_lb_listener"
}

// LocalName returns the local name for [DataLbListener].
func (ll *DataLbListener) LocalName() string {
	return ll.Name
}

// Configuration returns the configuration (args) for [DataLbListener].
func (ll *DataLbListener) Configuration() interface{} {
	return ll.Args
}

// Attributes returns the attributes for [DataLbListener].
func (ll *DataLbListener) Attributes() dataLbListenerAttributes {
	return dataLbListenerAttributes{ref: terra.ReferenceDataResource(ll)}
}

// DataLbListenerArgs contains the configurations for aws_lb_listener.
type DataLbListenerArgs struct {
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
	// DefaultAction: min=0
	DefaultAction []datalblistener.DefaultAction `hcl:"default_action,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *datalblistener.Timeouts `hcl:"timeouts,block"`
}
type dataLbListenerAttributes struct {
	ref terra.Reference
}

// AlpnPolicy returns a reference to field alpn_policy of aws_lb_listener.
func (ll dataLbListenerAttributes) AlpnPolicy() terra.StringValue {
	return terra.ReferenceAsString(ll.ref.Append("alpn_policy"))
}

// Arn returns a reference to field arn of aws_lb_listener.
func (ll dataLbListenerAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(ll.ref.Append("arn"))
}

// CertificateArn returns a reference to field certificate_arn of aws_lb_listener.
func (ll dataLbListenerAttributes) CertificateArn() terra.StringValue {
	return terra.ReferenceAsString(ll.ref.Append("certificate_arn"))
}

// Id returns a reference to field id of aws_lb_listener.
func (ll dataLbListenerAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ll.ref.Append("id"))
}

// LoadBalancerArn returns a reference to field load_balancer_arn of aws_lb_listener.
func (ll dataLbListenerAttributes) LoadBalancerArn() terra.StringValue {
	return terra.ReferenceAsString(ll.ref.Append("load_balancer_arn"))
}

// Port returns a reference to field port of aws_lb_listener.
func (ll dataLbListenerAttributes) Port() terra.NumberValue {
	return terra.ReferenceAsNumber(ll.ref.Append("port"))
}

// Protocol returns a reference to field protocol of aws_lb_listener.
func (ll dataLbListenerAttributes) Protocol() terra.StringValue {
	return terra.ReferenceAsString(ll.ref.Append("protocol"))
}

// SslPolicy returns a reference to field ssl_policy of aws_lb_listener.
func (ll dataLbListenerAttributes) SslPolicy() terra.StringValue {
	return terra.ReferenceAsString(ll.ref.Append("ssl_policy"))
}

// Tags returns a reference to field tags of aws_lb_listener.
func (ll dataLbListenerAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ll.ref.Append("tags"))
}

func (ll dataLbListenerAttributes) DefaultAction() terra.ListValue[datalblistener.DefaultActionAttributes] {
	return terra.ReferenceAsList[datalblistener.DefaultActionAttributes](ll.ref.Append("default_action"))
}

func (ll dataLbListenerAttributes) Timeouts() datalblistener.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datalblistener.TimeoutsAttributes](ll.ref.Append("timeouts"))
}
