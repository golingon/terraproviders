// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datahealthcaredicomservice "github.com/golingon/terraproviders/azurerm/3.69.0/datahealthcaredicomservice"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataHealthcareDicomService creates a new instance of [DataHealthcareDicomService].
func NewDataHealthcareDicomService(name string, args DataHealthcareDicomServiceArgs) *DataHealthcareDicomService {
	return &DataHealthcareDicomService{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataHealthcareDicomService)(nil)

// DataHealthcareDicomService represents the Terraform data resource azurerm_healthcare_dicom_service.
type DataHealthcareDicomService struct {
	Name string
	Args DataHealthcareDicomServiceArgs
}

// DataSource returns the Terraform object type for [DataHealthcareDicomService].
func (hds *DataHealthcareDicomService) DataSource() string {
	return "azurerm_healthcare_dicom_service"
}

// LocalName returns the local name for [DataHealthcareDicomService].
func (hds *DataHealthcareDicomService) LocalName() string {
	return hds.Name
}

// Configuration returns the configuration (args) for [DataHealthcareDicomService].
func (hds *DataHealthcareDicomService) Configuration() interface{} {
	return hds.Args
}

// Attributes returns the attributes for [DataHealthcareDicomService].
func (hds *DataHealthcareDicomService) Attributes() dataHealthcareDicomServiceAttributes {
	return dataHealthcareDicomServiceAttributes{ref: terra.ReferenceDataResource(hds)}
}

// DataHealthcareDicomServiceArgs contains the configurations for azurerm_healthcare_dicom_service.
type DataHealthcareDicomServiceArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// WorkspaceId: string, required
	WorkspaceId terra.StringValue `hcl:"workspace_id,attr" validate:"required"`
	// Authentication: min=0
	Authentication []datahealthcaredicomservice.Authentication `hcl:"authentication,block" validate:"min=0"`
	// Identity: min=0
	Identity []datahealthcaredicomservice.Identity `hcl:"identity,block" validate:"min=0"`
	// PrivateEndpoint: min=0
	PrivateEndpoint []datahealthcaredicomservice.PrivateEndpoint `hcl:"private_endpoint,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *datahealthcaredicomservice.Timeouts `hcl:"timeouts,block"`
}
type dataHealthcareDicomServiceAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_healthcare_dicom_service.
func (hds dataHealthcareDicomServiceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(hds.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_healthcare_dicom_service.
func (hds dataHealthcareDicomServiceAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(hds.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_healthcare_dicom_service.
func (hds dataHealthcareDicomServiceAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(hds.ref.Append("name"))
}

// ServiceUrl returns a reference to field service_url of azurerm_healthcare_dicom_service.
func (hds dataHealthcareDicomServiceAttributes) ServiceUrl() terra.StringValue {
	return terra.ReferenceAsString(hds.ref.Append("service_url"))
}

// Tags returns a reference to field tags of azurerm_healthcare_dicom_service.
func (hds dataHealthcareDicomServiceAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](hds.ref.Append("tags"))
}

// WorkspaceId returns a reference to field workspace_id of azurerm_healthcare_dicom_service.
func (hds dataHealthcareDicomServiceAttributes) WorkspaceId() terra.StringValue {
	return terra.ReferenceAsString(hds.ref.Append("workspace_id"))
}

func (hds dataHealthcareDicomServiceAttributes) Authentication() terra.ListValue[datahealthcaredicomservice.AuthenticationAttributes] {
	return terra.ReferenceAsList[datahealthcaredicomservice.AuthenticationAttributes](hds.ref.Append("authentication"))
}

func (hds dataHealthcareDicomServiceAttributes) Identity() terra.ListValue[datahealthcaredicomservice.IdentityAttributes] {
	return terra.ReferenceAsList[datahealthcaredicomservice.IdentityAttributes](hds.ref.Append("identity"))
}

func (hds dataHealthcareDicomServiceAttributes) PrivateEndpoint() terra.ListValue[datahealthcaredicomservice.PrivateEndpointAttributes] {
	return terra.ReferenceAsList[datahealthcaredicomservice.PrivateEndpointAttributes](hds.ref.Append("private_endpoint"))
}

func (hds dataHealthcareDicomServiceAttributes) Timeouts() datahealthcaredicomservice.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datahealthcaredicomservice.TimeoutsAttributes](hds.ref.Append("timeouts"))
}
