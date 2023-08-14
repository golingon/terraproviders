// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	datasesv2dedicatedippool "github.com/golingon/terraproviders/aws/5.12.0/datasesv2dedicatedippool"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataSesv2DedicatedIpPool creates a new instance of [DataSesv2DedicatedIpPool].
func NewDataSesv2DedicatedIpPool(name string, args DataSesv2DedicatedIpPoolArgs) *DataSesv2DedicatedIpPool {
	return &DataSesv2DedicatedIpPool{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataSesv2DedicatedIpPool)(nil)

// DataSesv2DedicatedIpPool represents the Terraform data resource aws_sesv2_dedicated_ip_pool.
type DataSesv2DedicatedIpPool struct {
	Name string
	Args DataSesv2DedicatedIpPoolArgs
}

// DataSource returns the Terraform object type for [DataSesv2DedicatedIpPool].
func (sdip *DataSesv2DedicatedIpPool) DataSource() string {
	return "aws_sesv2_dedicated_ip_pool"
}

// LocalName returns the local name for [DataSesv2DedicatedIpPool].
func (sdip *DataSesv2DedicatedIpPool) LocalName() string {
	return sdip.Name
}

// Configuration returns the configuration (args) for [DataSesv2DedicatedIpPool].
func (sdip *DataSesv2DedicatedIpPool) Configuration() interface{} {
	return sdip.Args
}

// Attributes returns the attributes for [DataSesv2DedicatedIpPool].
func (sdip *DataSesv2DedicatedIpPool) Attributes() dataSesv2DedicatedIpPoolAttributes {
	return dataSesv2DedicatedIpPoolAttributes{ref: terra.ReferenceDataResource(sdip)}
}

// DataSesv2DedicatedIpPoolArgs contains the configurations for aws_sesv2_dedicated_ip_pool.
type DataSesv2DedicatedIpPoolArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// PoolName: string, required
	PoolName terra.StringValue `hcl:"pool_name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// DedicatedIps: min=0
	DedicatedIps []datasesv2dedicatedippool.DedicatedIps `hcl:"dedicated_ips,block" validate:"min=0"`
}
type dataSesv2DedicatedIpPoolAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_sesv2_dedicated_ip_pool.
func (sdip dataSesv2DedicatedIpPoolAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(sdip.ref.Append("arn"))
}

// Id returns a reference to field id of aws_sesv2_dedicated_ip_pool.
func (sdip dataSesv2DedicatedIpPoolAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sdip.ref.Append("id"))
}

// PoolName returns a reference to field pool_name of aws_sesv2_dedicated_ip_pool.
func (sdip dataSesv2DedicatedIpPoolAttributes) PoolName() terra.StringValue {
	return terra.ReferenceAsString(sdip.ref.Append("pool_name"))
}

// ScalingMode returns a reference to field scaling_mode of aws_sesv2_dedicated_ip_pool.
func (sdip dataSesv2DedicatedIpPoolAttributes) ScalingMode() terra.StringValue {
	return terra.ReferenceAsString(sdip.ref.Append("scaling_mode"))
}

// Tags returns a reference to field tags of aws_sesv2_dedicated_ip_pool.
func (sdip dataSesv2DedicatedIpPoolAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](sdip.ref.Append("tags"))
}

func (sdip dataSesv2DedicatedIpPoolAttributes) DedicatedIps() terra.ListValue[datasesv2dedicatedippool.DedicatedIpsAttributes] {
	return terra.ReferenceAsList[datasesv2dedicatedippool.DedicatedIpsAttributes](sdip.ref.Append("dedicated_ips"))
}
