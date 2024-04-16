// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_healthcare_dicom_service

import "github.com/golingon/lingon/pkg/terra"

// Data creates a new instance of [DataSource].
func Data(name string, args DataArgs) *DataSource {
	return &DataSource{
		Args: args,
		Name: name,
	}
}

var _ terra.DataSource = (*DataSource)(nil)

// DataSource represents the Terraform data resource azurerm_healthcare_dicom_service.
type DataSource struct {
	Name string
	Args DataArgs
}

// DataSource returns the Terraform object type for [DataSource].
func (ahds *DataSource) DataSource() string {
	return "azurerm_healthcare_dicom_service"
}

// LocalName returns the local name for [DataSource].
func (ahds *DataSource) LocalName() string {
	return ahds.Name
}

// Configuration returns the configuration (args) for [DataSource].
func (ahds *DataSource) Configuration() interface{} {
	return ahds.Args
}

// Attributes returns the attributes for [DataSource].
func (ahds *DataSource) Attributes() dataAzurermHealthcareDicomServiceAttributes {
	return dataAzurermHealthcareDicomServiceAttributes{ref: terra.ReferenceDataSource(ahds)}
}

// DataArgs contains the configurations for azurerm_healthcare_dicom_service.
type DataArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// WorkspaceId: string, required
	WorkspaceId terra.StringValue `hcl:"workspace_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *DataTimeouts `hcl:"timeouts,block"`
}

type dataAzurermHealthcareDicomServiceAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_healthcare_dicom_service.
func (ahds dataAzurermHealthcareDicomServiceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ahds.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_healthcare_dicom_service.
func (ahds dataAzurermHealthcareDicomServiceAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(ahds.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_healthcare_dicom_service.
func (ahds dataAzurermHealthcareDicomServiceAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ahds.ref.Append("name"))
}

// ServiceUrl returns a reference to field service_url of azurerm_healthcare_dicom_service.
func (ahds dataAzurermHealthcareDicomServiceAttributes) ServiceUrl() terra.StringValue {
	return terra.ReferenceAsString(ahds.ref.Append("service_url"))
}

// Tags returns a reference to field tags of azurerm_healthcare_dicom_service.
func (ahds dataAzurermHealthcareDicomServiceAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ahds.ref.Append("tags"))
}

// WorkspaceId returns a reference to field workspace_id of azurerm_healthcare_dicom_service.
func (ahds dataAzurermHealthcareDicomServiceAttributes) WorkspaceId() terra.StringValue {
	return terra.ReferenceAsString(ahds.ref.Append("workspace_id"))
}

func (ahds dataAzurermHealthcareDicomServiceAttributes) Authentication() terra.ListValue[DataAuthenticationAttributes] {
	return terra.ReferenceAsList[DataAuthenticationAttributes](ahds.ref.Append("authentication"))
}

func (ahds dataAzurermHealthcareDicomServiceAttributes) Identity() terra.ListValue[DataIdentityAttributes] {
	return terra.ReferenceAsList[DataIdentityAttributes](ahds.ref.Append("identity"))
}

func (ahds dataAzurermHealthcareDicomServiceAttributes) PrivateEndpoint() terra.ListValue[DataPrivateEndpointAttributes] {
	return terra.ReferenceAsList[DataPrivateEndpointAttributes](ahds.ref.Append("private_endpoint"))
}

func (ahds dataAzurermHealthcareDicomServiceAttributes) Timeouts() DataTimeoutsAttributes {
	return terra.ReferenceAsSingle[DataTimeoutsAttributes](ahds.ref.Append("timeouts"))
}
