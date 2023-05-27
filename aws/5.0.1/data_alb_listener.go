// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	dataalblistener "github.com/golingon/terraproviders/aws/5.0.1/dataalblistener"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataAlbListener creates a new instance of [DataAlbListener].
func NewDataAlbListener(name string, args DataAlbListenerArgs) *DataAlbListener {
	return &DataAlbListener{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataAlbListener)(nil)

// DataAlbListener represents the Terraform data resource aws_alb_listener.
type DataAlbListener struct {
	Name string
	Args DataAlbListenerArgs
}

// DataSource returns the Terraform object type for [DataAlbListener].
func (al *DataAlbListener) DataSource() string {
	return "aws_alb_listener"
}

// LocalName returns the local name for [DataAlbListener].
func (al *DataAlbListener) LocalName() string {
	return al.Name
}

// Configuration returns the configuration (args) for [DataAlbListener].
func (al *DataAlbListener) Configuration() interface{} {
	return al.Args
}

// Attributes returns the attributes for [DataAlbListener].
func (al *DataAlbListener) Attributes() dataAlbListenerAttributes {
	return dataAlbListenerAttributes{ref: terra.ReferenceDataResource(al)}
}

// DataAlbListenerArgs contains the configurations for aws_alb_listener.
type DataAlbListenerArgs struct {
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
	DefaultAction []dataalblistener.DefaultAction `hcl:"default_action,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *dataalblistener.Timeouts `hcl:"timeouts,block"`
}
type dataAlbListenerAttributes struct {
	ref terra.Reference
}

// AlpnPolicy returns a reference to field alpn_policy of aws_alb_listener.
func (al dataAlbListenerAttributes) AlpnPolicy() terra.StringValue {
	return terra.ReferenceAsString(al.ref.Append("alpn_policy"))
}

// Arn returns a reference to field arn of aws_alb_listener.
func (al dataAlbListenerAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(al.ref.Append("arn"))
}

// CertificateArn returns a reference to field certificate_arn of aws_alb_listener.
func (al dataAlbListenerAttributes) CertificateArn() terra.StringValue {
	return terra.ReferenceAsString(al.ref.Append("certificate_arn"))
}

// Id returns a reference to field id of aws_alb_listener.
func (al dataAlbListenerAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(al.ref.Append("id"))
}

// LoadBalancerArn returns a reference to field load_balancer_arn of aws_alb_listener.
func (al dataAlbListenerAttributes) LoadBalancerArn() terra.StringValue {
	return terra.ReferenceAsString(al.ref.Append("load_balancer_arn"))
}

// Port returns a reference to field port of aws_alb_listener.
func (al dataAlbListenerAttributes) Port() terra.NumberValue {
	return terra.ReferenceAsNumber(al.ref.Append("port"))
}

// Protocol returns a reference to field protocol of aws_alb_listener.
func (al dataAlbListenerAttributes) Protocol() terra.StringValue {
	return terra.ReferenceAsString(al.ref.Append("protocol"))
}

// SslPolicy returns a reference to field ssl_policy of aws_alb_listener.
func (al dataAlbListenerAttributes) SslPolicy() terra.StringValue {
	return terra.ReferenceAsString(al.ref.Append("ssl_policy"))
}

// Tags returns a reference to field tags of aws_alb_listener.
func (al dataAlbListenerAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](al.ref.Append("tags"))
}

func (al dataAlbListenerAttributes) DefaultAction() terra.ListValue[dataalblistener.DefaultActionAttributes] {
	return terra.ReferenceAsList[dataalblistener.DefaultActionAttributes](al.ref.Append("default_action"))
}

func (al dataAlbListenerAttributes) Timeouts() dataalblistener.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dataalblistener.TimeoutsAttributes](al.ref.Append("timeouts"))
}
