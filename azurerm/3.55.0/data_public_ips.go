// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datapublicips "github.com/golingon/terraproviders/azurerm/3.55.0/datapublicips"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataPublicIps creates a new instance of [DataPublicIps].
func NewDataPublicIps(name string, args DataPublicIpsArgs) *DataPublicIps {
	return &DataPublicIps{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataPublicIps)(nil)

// DataPublicIps represents the Terraform data resource azurerm_public_ips.
type DataPublicIps struct {
	Name string
	Args DataPublicIpsArgs
}

// DataSource returns the Terraform object type for [DataPublicIps].
func (pi *DataPublicIps) DataSource() string {
	return "azurerm_public_ips"
}

// LocalName returns the local name for [DataPublicIps].
func (pi *DataPublicIps) LocalName() string {
	return pi.Name
}

// Configuration returns the configuration (args) for [DataPublicIps].
func (pi *DataPublicIps) Configuration() interface{} {
	return pi.Args
}

// Attributes returns the attributes for [DataPublicIps].
func (pi *DataPublicIps) Attributes() dataPublicIpsAttributes {
	return dataPublicIpsAttributes{ref: terra.ReferenceDataResource(pi)}
}

// DataPublicIpsArgs contains the configurations for azurerm_public_ips.
type DataPublicIpsArgs struct {
	// AllocationType: string, optional
	AllocationType terra.StringValue `hcl:"allocation_type,attr"`
	// AttachmentStatus: string, optional
	AttachmentStatus terra.StringValue `hcl:"attachment_status,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// NamePrefix: string, optional
	NamePrefix terra.StringValue `hcl:"name_prefix,attr"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// PublicIps: min=0
	PublicIps []datapublicips.PublicIps `hcl:"public_ips,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *datapublicips.Timeouts `hcl:"timeouts,block"`
}
type dataPublicIpsAttributes struct {
	ref terra.Reference
}

// AllocationType returns a reference to field allocation_type of azurerm_public_ips.
func (pi dataPublicIpsAttributes) AllocationType() terra.StringValue {
	return terra.ReferenceAsString(pi.ref.Append("allocation_type"))
}

// AttachmentStatus returns a reference to field attachment_status of azurerm_public_ips.
func (pi dataPublicIpsAttributes) AttachmentStatus() terra.StringValue {
	return terra.ReferenceAsString(pi.ref.Append("attachment_status"))
}

// Id returns a reference to field id of azurerm_public_ips.
func (pi dataPublicIpsAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(pi.ref.Append("id"))
}

// NamePrefix returns a reference to field name_prefix of azurerm_public_ips.
func (pi dataPublicIpsAttributes) NamePrefix() terra.StringValue {
	return terra.ReferenceAsString(pi.ref.Append("name_prefix"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_public_ips.
func (pi dataPublicIpsAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(pi.ref.Append("resource_group_name"))
}

func (pi dataPublicIpsAttributes) PublicIps() terra.ListValue[datapublicips.PublicIpsAttributes] {
	return terra.ReferenceAsList[datapublicips.PublicIpsAttributes](pi.ref.Append("public_ips"))
}

func (pi dataPublicIpsAttributes) Timeouts() datapublicips.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datapublicips.TimeoutsAttributes](pi.ref.Append("timeouts"))
}
