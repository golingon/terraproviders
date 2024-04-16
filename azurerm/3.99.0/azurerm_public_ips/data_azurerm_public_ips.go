// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_public_ips

import "github.com/golingon/lingon/pkg/terra"

// Data creates a new instance of [DataSource].
func Data(name string, args DataArgs) *DataSource {
	return &DataSource{
		Args: args,
		Name: name,
	}
}

var _ terra.DataSource = (*DataSource)(nil)

// DataSource represents the Terraform data resource azurerm_public_ips.
type DataSource struct {
	Name string
	Args DataArgs
}

// DataSource returns the Terraform object type for [DataSource].
func (api *DataSource) DataSource() string {
	return "azurerm_public_ips"
}

// LocalName returns the local name for [DataSource].
func (api *DataSource) LocalName() string {
	return api.Name
}

// Configuration returns the configuration (args) for [DataSource].
func (api *DataSource) Configuration() interface{} {
	return api.Args
}

// Attributes returns the attributes for [DataSource].
func (api *DataSource) Attributes() dataAzurermPublicIpsAttributes {
	return dataAzurermPublicIpsAttributes{ref: terra.ReferenceDataSource(api)}
}

// DataArgs contains the configurations for azurerm_public_ips.
type DataArgs struct {
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
	// Timeouts: optional
	Timeouts *DataTimeouts `hcl:"timeouts,block"`
}

type dataAzurermPublicIpsAttributes struct {
	ref terra.Reference
}

// AllocationType returns a reference to field allocation_type of azurerm_public_ips.
func (api dataAzurermPublicIpsAttributes) AllocationType() terra.StringValue {
	return terra.ReferenceAsString(api.ref.Append("allocation_type"))
}

// AttachmentStatus returns a reference to field attachment_status of azurerm_public_ips.
func (api dataAzurermPublicIpsAttributes) AttachmentStatus() terra.StringValue {
	return terra.ReferenceAsString(api.ref.Append("attachment_status"))
}

// Id returns a reference to field id of azurerm_public_ips.
func (api dataAzurermPublicIpsAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(api.ref.Append("id"))
}

// NamePrefix returns a reference to field name_prefix of azurerm_public_ips.
func (api dataAzurermPublicIpsAttributes) NamePrefix() terra.StringValue {
	return terra.ReferenceAsString(api.ref.Append("name_prefix"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_public_ips.
func (api dataAzurermPublicIpsAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(api.ref.Append("resource_group_name"))
}

func (api dataAzurermPublicIpsAttributes) PublicIps() terra.ListValue[DataPublicIpsAttributes] {
	return terra.ReferenceAsList[DataPublicIpsAttributes](api.ref.Append("public_ips"))
}

func (api dataAzurermPublicIpsAttributes) Timeouts() DataTimeoutsAttributes {
	return terra.ReferenceAsSingle[DataTimeoutsAttributes](api.ref.Append("timeouts"))
}
