// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_private_dns_srv_record

import "github.com/golingon/lingon/pkg/terra"

// Data creates a new instance of [DataSource].
func Data(name string, args DataArgs) *DataSource {
	return &DataSource{
		Args: args,
		Name: name,
	}
}

var _ terra.DataSource = (*DataSource)(nil)

// DataSource represents the Terraform data resource azurerm_private_dns_srv_record.
type DataSource struct {
	Name string
	Args DataArgs
}

// DataSource returns the Terraform object type for [DataSource].
func (apdsr *DataSource) DataSource() string {
	return "azurerm_private_dns_srv_record"
}

// LocalName returns the local name for [DataSource].
func (apdsr *DataSource) LocalName() string {
	return apdsr.Name
}

// Configuration returns the configuration (args) for [DataSource].
func (apdsr *DataSource) Configuration() interface{} {
	return apdsr.Args
}

// Attributes returns the attributes for [DataSource].
func (apdsr *DataSource) Attributes() dataAzurermPrivateDnsSrvRecordAttributes {
	return dataAzurermPrivateDnsSrvRecordAttributes{ref: terra.ReferenceDataSource(apdsr)}
}

// DataArgs contains the configurations for azurerm_private_dns_srv_record.
type DataArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// ZoneName: string, required
	ZoneName terra.StringValue `hcl:"zone_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *DataTimeouts `hcl:"timeouts,block"`
}

type dataAzurermPrivateDnsSrvRecordAttributes struct {
	ref terra.Reference
}

// Fqdn returns a reference to field fqdn of azurerm_private_dns_srv_record.
func (apdsr dataAzurermPrivateDnsSrvRecordAttributes) Fqdn() terra.StringValue {
	return terra.ReferenceAsString(apdsr.ref.Append("fqdn"))
}

// Id returns a reference to field id of azurerm_private_dns_srv_record.
func (apdsr dataAzurermPrivateDnsSrvRecordAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(apdsr.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_private_dns_srv_record.
func (apdsr dataAzurermPrivateDnsSrvRecordAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(apdsr.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_private_dns_srv_record.
func (apdsr dataAzurermPrivateDnsSrvRecordAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(apdsr.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_private_dns_srv_record.
func (apdsr dataAzurermPrivateDnsSrvRecordAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](apdsr.ref.Append("tags"))
}

// Ttl returns a reference to field ttl of azurerm_private_dns_srv_record.
func (apdsr dataAzurermPrivateDnsSrvRecordAttributes) Ttl() terra.NumberValue {
	return terra.ReferenceAsNumber(apdsr.ref.Append("ttl"))
}

// ZoneName returns a reference to field zone_name of azurerm_private_dns_srv_record.
func (apdsr dataAzurermPrivateDnsSrvRecordAttributes) ZoneName() terra.StringValue {
	return terra.ReferenceAsString(apdsr.ref.Append("zone_name"))
}

func (apdsr dataAzurermPrivateDnsSrvRecordAttributes) Record() terra.SetValue[DataRecordAttributes] {
	return terra.ReferenceAsSet[DataRecordAttributes](apdsr.ref.Append("record"))
}

func (apdsr dataAzurermPrivateDnsSrvRecordAttributes) Timeouts() DataTimeoutsAttributes {
	return terra.ReferenceAsSingle[DataTimeoutsAttributes](apdsr.ref.Append("timeouts"))
}
