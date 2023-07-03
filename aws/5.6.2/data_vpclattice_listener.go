// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	datavpclatticelistener "github.com/golingon/terraproviders/aws/5.6.2/datavpclatticelistener"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataVpclatticeListener creates a new instance of [DataVpclatticeListener].
func NewDataVpclatticeListener(name string, args DataVpclatticeListenerArgs) *DataVpclatticeListener {
	return &DataVpclatticeListener{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataVpclatticeListener)(nil)

// DataVpclatticeListener represents the Terraform data resource aws_vpclattice_listener.
type DataVpclatticeListener struct {
	Name string
	Args DataVpclatticeListenerArgs
}

// DataSource returns the Terraform object type for [DataVpclatticeListener].
func (vl *DataVpclatticeListener) DataSource() string {
	return "aws_vpclattice_listener"
}

// LocalName returns the local name for [DataVpclatticeListener].
func (vl *DataVpclatticeListener) LocalName() string {
	return vl.Name
}

// Configuration returns the configuration (args) for [DataVpclatticeListener].
func (vl *DataVpclatticeListener) Configuration() interface{} {
	return vl.Args
}

// Attributes returns the attributes for [DataVpclatticeListener].
func (vl *DataVpclatticeListener) Attributes() dataVpclatticeListenerAttributes {
	return dataVpclatticeListenerAttributes{ref: terra.ReferenceDataResource(vl)}
}

// DataVpclatticeListenerArgs contains the configurations for aws_vpclattice_listener.
type DataVpclatticeListenerArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ListenerIdentifier: string, required
	ListenerIdentifier terra.StringValue `hcl:"listener_identifier,attr" validate:"required"`
	// ServiceIdentifier: string, required
	ServiceIdentifier terra.StringValue `hcl:"service_identifier,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// DefaultAction: min=0
	DefaultAction []datavpclatticelistener.DefaultAction `hcl:"default_action,block" validate:"min=0"`
}
type dataVpclatticeListenerAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_vpclattice_listener.
func (vl dataVpclatticeListenerAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(vl.ref.Append("arn"))
}

// CreatedAt returns a reference to field created_at of aws_vpclattice_listener.
func (vl dataVpclatticeListenerAttributes) CreatedAt() terra.StringValue {
	return terra.ReferenceAsString(vl.ref.Append("created_at"))
}

// Id returns a reference to field id of aws_vpclattice_listener.
func (vl dataVpclatticeListenerAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(vl.ref.Append("id"))
}

// LastUpdatedAt returns a reference to field last_updated_at of aws_vpclattice_listener.
func (vl dataVpclatticeListenerAttributes) LastUpdatedAt() terra.StringValue {
	return terra.ReferenceAsString(vl.ref.Append("last_updated_at"))
}

// ListenerId returns a reference to field listener_id of aws_vpclattice_listener.
func (vl dataVpclatticeListenerAttributes) ListenerId() terra.StringValue {
	return terra.ReferenceAsString(vl.ref.Append("listener_id"))
}

// ListenerIdentifier returns a reference to field listener_identifier of aws_vpclattice_listener.
func (vl dataVpclatticeListenerAttributes) ListenerIdentifier() terra.StringValue {
	return terra.ReferenceAsString(vl.ref.Append("listener_identifier"))
}

// Name returns a reference to field name of aws_vpclattice_listener.
func (vl dataVpclatticeListenerAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(vl.ref.Append("name"))
}

// Port returns a reference to field port of aws_vpclattice_listener.
func (vl dataVpclatticeListenerAttributes) Port() terra.NumberValue {
	return terra.ReferenceAsNumber(vl.ref.Append("port"))
}

// Protocol returns a reference to field protocol of aws_vpclattice_listener.
func (vl dataVpclatticeListenerAttributes) Protocol() terra.StringValue {
	return terra.ReferenceAsString(vl.ref.Append("protocol"))
}

// ServiceArn returns a reference to field service_arn of aws_vpclattice_listener.
func (vl dataVpclatticeListenerAttributes) ServiceArn() terra.StringValue {
	return terra.ReferenceAsString(vl.ref.Append("service_arn"))
}

// ServiceId returns a reference to field service_id of aws_vpclattice_listener.
func (vl dataVpclatticeListenerAttributes) ServiceId() terra.StringValue {
	return terra.ReferenceAsString(vl.ref.Append("service_id"))
}

// ServiceIdentifier returns a reference to field service_identifier of aws_vpclattice_listener.
func (vl dataVpclatticeListenerAttributes) ServiceIdentifier() terra.StringValue {
	return terra.ReferenceAsString(vl.ref.Append("service_identifier"))
}

// Tags returns a reference to field tags of aws_vpclattice_listener.
func (vl dataVpclatticeListenerAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](vl.ref.Append("tags"))
}

func (vl dataVpclatticeListenerAttributes) DefaultAction() terra.ListValue[datavpclatticelistener.DefaultActionAttributes] {
	return terra.ReferenceAsList[datavpclatticelistener.DefaultActionAttributes](vl.ref.Append("default_action"))
}
