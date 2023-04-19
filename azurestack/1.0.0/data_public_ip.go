// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurestack

import (
	datapublicip "github.com/golingon/terraproviders/azurestack/1.0.0/datapublicip"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataPublicIp creates a new instance of [DataPublicIp].
func NewDataPublicIp(name string, args DataPublicIpArgs) *DataPublicIp {
	return &DataPublicIp{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataPublicIp)(nil)

// DataPublicIp represents the Terraform data resource azurestack_public_ip.
type DataPublicIp struct {
	Name string
	Args DataPublicIpArgs
}

// DataSource returns the Terraform object type for [DataPublicIp].
func (pi *DataPublicIp) DataSource() string {
	return "azurestack_public_ip"
}

// LocalName returns the local name for [DataPublicIp].
func (pi *DataPublicIp) LocalName() string {
	return pi.Name
}

// Configuration returns the configuration (args) for [DataPublicIp].
func (pi *DataPublicIp) Configuration() interface{} {
	return pi.Args
}

// Attributes returns the attributes for [DataPublicIp].
func (pi *DataPublicIp) Attributes() dataPublicIpAttributes {
	return dataPublicIpAttributes{ref: terra.ReferenceDataResource(pi)}
}

// DataPublicIpArgs contains the configurations for azurestack_public_ip.
type DataPublicIpArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Timeouts: optional
	Timeouts *datapublicip.Timeouts `hcl:"timeouts,block"`
}
type dataPublicIpAttributes struct {
	ref terra.Reference
}

// AllocationMethod returns a reference to field allocation_method of azurestack_public_ip.
func (pi dataPublicIpAttributes) AllocationMethod() terra.StringValue {
	return terra.ReferenceAsString(pi.ref.Append("allocation_method"))
}

// DomainNameLabel returns a reference to field domain_name_label of azurestack_public_ip.
func (pi dataPublicIpAttributes) DomainNameLabel() terra.StringValue {
	return terra.ReferenceAsString(pi.ref.Append("domain_name_label"))
}

// Fqdn returns a reference to field fqdn of azurestack_public_ip.
func (pi dataPublicIpAttributes) Fqdn() terra.StringValue {
	return terra.ReferenceAsString(pi.ref.Append("fqdn"))
}

// Id returns a reference to field id of azurestack_public_ip.
func (pi dataPublicIpAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(pi.ref.Append("id"))
}

// IdleTimeoutInMinutes returns a reference to field idle_timeout_in_minutes of azurestack_public_ip.
func (pi dataPublicIpAttributes) IdleTimeoutInMinutes() terra.NumberValue {
	return terra.ReferenceAsNumber(pi.ref.Append("idle_timeout_in_minutes"))
}

// IpAddress returns a reference to field ip_address of azurestack_public_ip.
func (pi dataPublicIpAttributes) IpAddress() terra.StringValue {
	return terra.ReferenceAsString(pi.ref.Append("ip_address"))
}

// IpVersion returns a reference to field ip_version of azurestack_public_ip.
func (pi dataPublicIpAttributes) IpVersion() terra.StringValue {
	return terra.ReferenceAsString(pi.ref.Append("ip_version"))
}

// Location returns a reference to field location of azurestack_public_ip.
func (pi dataPublicIpAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(pi.ref.Append("location"))
}

// Name returns a reference to field name of azurestack_public_ip.
func (pi dataPublicIpAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(pi.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurestack_public_ip.
func (pi dataPublicIpAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(pi.ref.Append("resource_group_name"))
}

// ReverseFqdn returns a reference to field reverse_fqdn of azurestack_public_ip.
func (pi dataPublicIpAttributes) ReverseFqdn() terra.StringValue {
	return terra.ReferenceAsString(pi.ref.Append("reverse_fqdn"))
}

// Sku returns a reference to field sku of azurestack_public_ip.
func (pi dataPublicIpAttributes) Sku() terra.StringValue {
	return terra.ReferenceAsString(pi.ref.Append("sku"))
}

// Tags returns a reference to field tags of azurestack_public_ip.
func (pi dataPublicIpAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](pi.ref.Append("tags"))
}

// Zones returns a reference to field zones of azurestack_public_ip.
func (pi dataPublicIpAttributes) Zones() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](pi.ref.Append("zones"))
}

func (pi dataPublicIpAttributes) Timeouts() datapublicip.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datapublicip.TimeoutsAttributes](pi.ref.Append("timeouts"))
}