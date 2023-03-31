// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datahealthcaremedtechservice "github.com/golingon/terraproviders/azurerm/3.49.0/datahealthcaremedtechservice"
	"github.com/volvo-cars/lingon/pkg/terra"
)

func NewDataHealthcareMedtechService(name string, args DataHealthcareMedtechServiceArgs) *DataHealthcareMedtechService {
	return &DataHealthcareMedtechService{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataHealthcareMedtechService)(nil)

type DataHealthcareMedtechService struct {
	Name string
	Args DataHealthcareMedtechServiceArgs
}

func (hms *DataHealthcareMedtechService) DataSource() string {
	return "azurerm_healthcare_medtech_service"
}

func (hms *DataHealthcareMedtechService) LocalName() string {
	return hms.Name
}

func (hms *DataHealthcareMedtechService) Configuration() interface{} {
	return hms.Args
}

func (hms *DataHealthcareMedtechService) Attributes() dataHealthcareMedtechServiceAttributes {
	return dataHealthcareMedtechServiceAttributes{ref: terra.ReferenceDataResource(hms)}
}

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

func (hms dataHealthcareMedtechServiceAttributes) DeviceMappingJson() terra.StringValue {
	return terra.ReferenceString(hms.ref.Append("device_mapping_json"))
}

func (hms dataHealthcareMedtechServiceAttributes) EventhubConsumerGroupName() terra.StringValue {
	return terra.ReferenceString(hms.ref.Append("eventhub_consumer_group_name"))
}

func (hms dataHealthcareMedtechServiceAttributes) EventhubName() terra.StringValue {
	return terra.ReferenceString(hms.ref.Append("eventhub_name"))
}

func (hms dataHealthcareMedtechServiceAttributes) EventhubNamespaceName() terra.StringValue {
	return terra.ReferenceString(hms.ref.Append("eventhub_namespace_name"))
}

func (hms dataHealthcareMedtechServiceAttributes) Id() terra.StringValue {
	return terra.ReferenceString(hms.ref.Append("id"))
}

func (hms dataHealthcareMedtechServiceAttributes) Name() terra.StringValue {
	return terra.ReferenceString(hms.ref.Append("name"))
}

func (hms dataHealthcareMedtechServiceAttributes) WorkspaceId() terra.StringValue {
	return terra.ReferenceString(hms.ref.Append("workspace_id"))
}

func (hms dataHealthcareMedtechServiceAttributes) Identity() terra.ListValue[datahealthcaremedtechservice.IdentityAttributes] {
	return terra.ReferenceList[datahealthcaremedtechservice.IdentityAttributes](hms.ref.Append("identity"))
}

func (hms dataHealthcareMedtechServiceAttributes) Timeouts() datahealthcaremedtechservice.TimeoutsAttributes {
	return terra.ReferenceSingle[datahealthcaremedtechservice.TimeoutsAttributes](hms.ref.Append("timeouts"))
}
