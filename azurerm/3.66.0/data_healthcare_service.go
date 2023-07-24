// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datahealthcareservice "github.com/golingon/terraproviders/azurerm/3.66.0/datahealthcareservice"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataHealthcareService creates a new instance of [DataHealthcareService].
func NewDataHealthcareService(name string, args DataHealthcareServiceArgs) *DataHealthcareService {
	return &DataHealthcareService{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataHealthcareService)(nil)

// DataHealthcareService represents the Terraform data resource azurerm_healthcare_service.
type DataHealthcareService struct {
	Name string
	Args DataHealthcareServiceArgs
}

// DataSource returns the Terraform object type for [DataHealthcareService].
func (hs *DataHealthcareService) DataSource() string {
	return "azurerm_healthcare_service"
}

// LocalName returns the local name for [DataHealthcareService].
func (hs *DataHealthcareService) LocalName() string {
	return hs.Name
}

// Configuration returns the configuration (args) for [DataHealthcareService].
func (hs *DataHealthcareService) Configuration() interface{} {
	return hs.Args
}

// Attributes returns the attributes for [DataHealthcareService].
func (hs *DataHealthcareService) Attributes() dataHealthcareServiceAttributes {
	return dataHealthcareServiceAttributes{ref: terra.ReferenceDataResource(hs)}
}

// DataHealthcareServiceArgs contains the configurations for azurerm_healthcare_service.
type DataHealthcareServiceArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// AuthenticationConfiguration: min=0
	AuthenticationConfiguration []datahealthcareservice.AuthenticationConfiguration `hcl:"authentication_configuration,block" validate:"min=0"`
	// CorsConfiguration: min=0
	CorsConfiguration []datahealthcareservice.CorsConfiguration `hcl:"cors_configuration,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *datahealthcareservice.Timeouts `hcl:"timeouts,block"`
}
type dataHealthcareServiceAttributes struct {
	ref terra.Reference
}

// AccessPolicyObjectIds returns a reference to field access_policy_object_ids of azurerm_healthcare_service.
func (hs dataHealthcareServiceAttributes) AccessPolicyObjectIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](hs.ref.Append("access_policy_object_ids"))
}

// CosmosdbKeyVaultKeyVersionlessId returns a reference to field cosmosdb_key_vault_key_versionless_id of azurerm_healthcare_service.
func (hs dataHealthcareServiceAttributes) CosmosdbKeyVaultKeyVersionlessId() terra.StringValue {
	return terra.ReferenceAsString(hs.ref.Append("cosmosdb_key_vault_key_versionless_id"))
}

// CosmosdbThroughput returns a reference to field cosmosdb_throughput of azurerm_healthcare_service.
func (hs dataHealthcareServiceAttributes) CosmosdbThroughput() terra.NumberValue {
	return terra.ReferenceAsNumber(hs.ref.Append("cosmosdb_throughput"))
}

// Id returns a reference to field id of azurerm_healthcare_service.
func (hs dataHealthcareServiceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(hs.ref.Append("id"))
}

// Kind returns a reference to field kind of azurerm_healthcare_service.
func (hs dataHealthcareServiceAttributes) Kind() terra.StringValue {
	return terra.ReferenceAsString(hs.ref.Append("kind"))
}

// Location returns a reference to field location of azurerm_healthcare_service.
func (hs dataHealthcareServiceAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(hs.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_healthcare_service.
func (hs dataHealthcareServiceAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(hs.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_healthcare_service.
func (hs dataHealthcareServiceAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(hs.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_healthcare_service.
func (hs dataHealthcareServiceAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](hs.ref.Append("tags"))
}

func (hs dataHealthcareServiceAttributes) AuthenticationConfiguration() terra.ListValue[datahealthcareservice.AuthenticationConfigurationAttributes] {
	return terra.ReferenceAsList[datahealthcareservice.AuthenticationConfigurationAttributes](hs.ref.Append("authentication_configuration"))
}

func (hs dataHealthcareServiceAttributes) CorsConfiguration() terra.ListValue[datahealthcareservice.CorsConfigurationAttributes] {
	return terra.ReferenceAsList[datahealthcareservice.CorsConfigurationAttributes](hs.ref.Append("cors_configuration"))
}

func (hs dataHealthcareServiceAttributes) Timeouts() datahealthcareservice.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datahealthcareservice.TimeoutsAttributes](hs.ref.Append("timeouts"))
}
