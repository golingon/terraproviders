// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datahealthcaremedtechservice "github.com/golingon/terraproviders/azurerm/3.69.0/datahealthcaremedtechservice"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataHealthcareMedtechService creates a new instance of [DataHealthcareMedtechService].
func NewDataHealthcareMedtechService(name string, args DataHealthcareMedtechServiceArgs) *DataHealthcareMedtechService {
	return &DataHealthcareMedtechService{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataHealthcareMedtechService)(nil)

// DataHealthcareMedtechService represents the Terraform data resource azurerm_healthcare_medtech_service.
type DataHealthcareMedtechService struct {
	Name string
	Args DataHealthcareMedtechServiceArgs
}

// DataSource returns the Terraform object type for [DataHealthcareMedtechService].
func (hms *DataHealthcareMedtechService) DataSource() string {
	return "azurerm_healthcare_medtech_service"
}

// LocalName returns the local name for [DataHealthcareMedtechService].
func (hms *DataHealthcareMedtechService) LocalName() string {
	return hms.Name
}

// Configuration returns the configuration (args) for [DataHealthcareMedtechService].
func (hms *DataHealthcareMedtechService) Configuration() interface{} {
	return hms.Args
}

// Attributes returns the attributes for [DataHealthcareMedtechService].
func (hms *DataHealthcareMedtechService) Attributes() dataHealthcareMedtechServiceAttributes {
	return dataHealthcareMedtechServiceAttributes{ref: terra.ReferenceDataResource(hms)}
}

// DataHealthcareMedtechServiceArgs contains the configurations for azurerm_healthcare_medtech_service.
type DataHealthcareMedtechServiceArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// WorkspaceId: string, required
	WorkspaceId terra.StringValue `hcl:"workspace_id,attr" validate:"required"`
	// Identity: min=0
	Identity []datahealthcaremedtechservice.Identity `hcl:"identity,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *datahealthcaremedtechservice.Timeouts `hcl:"timeouts,block"`
}
type dataHealthcareMedtechServiceAttributes struct {
	ref terra.Reference
}

// DeviceMappingJson returns a reference to field device_mapping_json of azurerm_healthcare_medtech_service.
func (hms dataHealthcareMedtechServiceAttributes) DeviceMappingJson() terra.StringValue {
	return terra.ReferenceAsString(hms.ref.Append("device_mapping_json"))
}

// EventhubConsumerGroupName returns a reference to field eventhub_consumer_group_name of azurerm_healthcare_medtech_service.
func (hms dataHealthcareMedtechServiceAttributes) EventhubConsumerGroupName() terra.StringValue {
	return terra.ReferenceAsString(hms.ref.Append("eventhub_consumer_group_name"))
}

// EventhubName returns a reference to field eventhub_name of azurerm_healthcare_medtech_service.
func (hms dataHealthcareMedtechServiceAttributes) EventhubName() terra.StringValue {
	return terra.ReferenceAsString(hms.ref.Append("eventhub_name"))
}

// EventhubNamespaceName returns a reference to field eventhub_namespace_name of azurerm_healthcare_medtech_service.
func (hms dataHealthcareMedtechServiceAttributes) EventhubNamespaceName() terra.StringValue {
	return terra.ReferenceAsString(hms.ref.Append("eventhub_namespace_name"))
}

// Id returns a reference to field id of azurerm_healthcare_medtech_service.
func (hms dataHealthcareMedtechServiceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(hms.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_healthcare_medtech_service.
func (hms dataHealthcareMedtechServiceAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(hms.ref.Append("name"))
}

// WorkspaceId returns a reference to field workspace_id of azurerm_healthcare_medtech_service.
func (hms dataHealthcareMedtechServiceAttributes) WorkspaceId() terra.StringValue {
	return terra.ReferenceAsString(hms.ref.Append("workspace_id"))
}

func (hms dataHealthcareMedtechServiceAttributes) Identity() terra.ListValue[datahealthcaremedtechservice.IdentityAttributes] {
	return terra.ReferenceAsList[datahealthcaremedtechservice.IdentityAttributes](hms.ref.Append("identity"))
}

func (hms dataHealthcareMedtechServiceAttributes) Timeouts() datahealthcaremedtechservice.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datahealthcaremedtechservice.TimeoutsAttributes](hms.ref.Append("timeouts"))
}
