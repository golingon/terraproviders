// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurestack_public_ip

import "github.com/golingon/lingon/pkg/terra"

// Data creates a new instance of [DataSource].
func Data(name string, args DataArgs) *DataSource {
	return &DataSource{
		Args: args,
		Name: name,
	}
}

var _ terra.DataSource = (*DataSource)(nil)

// DataSource represents the Terraform data resource azurestack_public_ip.
type DataSource struct {
	Name string
	Args DataArgs
}

// DataSource returns the Terraform object type for [DataSource].
func (api *DataSource) DataSource() string {
	return "azurestack_public_ip"
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
func (api *DataSource) Attributes() dataAzurestackPublicIpAttributes {
	return dataAzurestackPublicIpAttributes{ref: terra.ReferenceDataSource(api)}
}

// DataArgs contains the configurations for azurestack_public_ip.
type DataArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Timeouts: optional
	Timeouts *DataTimeouts `hcl:"timeouts,block"`
}

type dataAzurestackPublicIpAttributes struct {
	ref terra.Reference
}

// AllocationMethod returns a reference to field allocation_method of azurestack_public_ip.
func (api dataAzurestackPublicIpAttributes) AllocationMethod() terra.StringValue {
	return terra.ReferenceAsString(api.ref.Append("allocation_method"))
}

// DomainNameLabel returns a reference to field domain_name_label of azurestack_public_ip.
func (api dataAzurestackPublicIpAttributes) DomainNameLabel() terra.StringValue {
	return terra.ReferenceAsString(api.ref.Append("domain_name_label"))
}

// Fqdn returns a reference to field fqdn of azurestack_public_ip.
func (api dataAzurestackPublicIpAttributes) Fqdn() terra.StringValue {
	return terra.ReferenceAsString(api.ref.Append("fqdn"))
}

// Id returns a reference to field id of azurestack_public_ip.
func (api dataAzurestackPublicIpAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(api.ref.Append("id"))
}

// IdleTimeoutInMinutes returns a reference to field idle_timeout_in_minutes of azurestack_public_ip.
func (api dataAzurestackPublicIpAttributes) IdleTimeoutInMinutes() terra.NumberValue {
	return terra.ReferenceAsNumber(api.ref.Append("idle_timeout_in_minutes"))
}

// IpAddress returns a reference to field ip_address of azurestack_public_ip.
func (api dataAzurestackPublicIpAttributes) IpAddress() terra.StringValue {
	return terra.ReferenceAsString(api.ref.Append("ip_address"))
}

// IpVersion returns a reference to field ip_version of azurestack_public_ip.
func (api dataAzurestackPublicIpAttributes) IpVersion() terra.StringValue {
	return terra.ReferenceAsString(api.ref.Append("ip_version"))
}

// Location returns a reference to field location of azurestack_public_ip.
func (api dataAzurestackPublicIpAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(api.ref.Append("location"))
}

// Name returns a reference to field name of azurestack_public_ip.
func (api dataAzurestackPublicIpAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(api.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurestack_public_ip.
func (api dataAzurestackPublicIpAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(api.ref.Append("resource_group_name"))
}

// ReverseFqdn returns a reference to field reverse_fqdn of azurestack_public_ip.
func (api dataAzurestackPublicIpAttributes) ReverseFqdn() terra.StringValue {
	return terra.ReferenceAsString(api.ref.Append("reverse_fqdn"))
}

// Sku returns a reference to field sku of azurestack_public_ip.
func (api dataAzurestackPublicIpAttributes) Sku() terra.StringValue {
	return terra.ReferenceAsString(api.ref.Append("sku"))
}

// Tags returns a reference to field tags of azurestack_public_ip.
func (api dataAzurestackPublicIpAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](api.ref.Append("tags"))
}

// Zones returns a reference to field zones of azurestack_public_ip.
func (api dataAzurestackPublicIpAttributes) Zones() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](api.ref.Append("zones"))
}

func (api dataAzurestackPublicIpAttributes) Timeouts() DataTimeoutsAttributes {
	return terra.ReferenceAsSingle[DataTimeoutsAttributes](api.ref.Append("timeouts"))
}
