// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	datacloudrunservice "github.com/golingon/terraproviders/googlebeta/4.76.0/datacloudrunservice"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataCloudRunService creates a new instance of [DataCloudRunService].
func NewDataCloudRunService(name string, args DataCloudRunServiceArgs) *DataCloudRunService {
	return &DataCloudRunService{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataCloudRunService)(nil)

// DataCloudRunService represents the Terraform data resource google_cloud_run_service.
type DataCloudRunService struct {
	Name string
	Args DataCloudRunServiceArgs
}

// DataSource returns the Terraform object type for [DataCloudRunService].
func (crs *DataCloudRunService) DataSource() string {
	return "google_cloud_run_service"
}

// LocalName returns the local name for [DataCloudRunService].
func (crs *DataCloudRunService) LocalName() string {
	return crs.Name
}

// Configuration returns the configuration (args) for [DataCloudRunService].
func (crs *DataCloudRunService) Configuration() interface{} {
	return crs.Args
}

// Attributes returns the attributes for [DataCloudRunService].
func (crs *DataCloudRunService) Attributes() dataCloudRunServiceAttributes {
	return dataCloudRunServiceAttributes{ref: terra.ReferenceDataResource(crs)}
}

// DataCloudRunServiceArgs contains the configurations for google_cloud_run_service.
type DataCloudRunServiceArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Metadata: min=0
	Metadata []datacloudrunservice.Metadata `hcl:"metadata,block" validate:"min=0"`
	// Status: min=0
	Status []datacloudrunservice.Status `hcl:"status,block" validate:"min=0"`
	// Template: min=0
	Template []datacloudrunservice.Template `hcl:"template,block" validate:"min=0"`
	// Traffic: min=0
	Traffic []datacloudrunservice.Traffic `hcl:"traffic,block" validate:"min=0"`
}
type dataCloudRunServiceAttributes struct {
	ref terra.Reference
}

// AutogenerateRevisionName returns a reference to field autogenerate_revision_name of google_cloud_run_service.
func (crs dataCloudRunServiceAttributes) AutogenerateRevisionName() terra.BoolValue {
	return terra.ReferenceAsBool(crs.ref.Append("autogenerate_revision_name"))
}

// Id returns a reference to field id of google_cloud_run_service.
func (crs dataCloudRunServiceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(crs.ref.Append("id"))
}

// Location returns a reference to field location of google_cloud_run_service.
func (crs dataCloudRunServiceAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(crs.ref.Append("location"))
}

// Name returns a reference to field name of google_cloud_run_service.
func (crs dataCloudRunServiceAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(crs.ref.Append("name"))
}

// Project returns a reference to field project of google_cloud_run_service.
func (crs dataCloudRunServiceAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(crs.ref.Append("project"))
}

func (crs dataCloudRunServiceAttributes) Metadata() terra.ListValue[datacloudrunservice.MetadataAttributes] {
	return terra.ReferenceAsList[datacloudrunservice.MetadataAttributes](crs.ref.Append("metadata"))
}

func (crs dataCloudRunServiceAttributes) Status() terra.ListValue[datacloudrunservice.StatusAttributes] {
	return terra.ReferenceAsList[datacloudrunservice.StatusAttributes](crs.ref.Append("status"))
}

func (crs dataCloudRunServiceAttributes) Template() terra.ListValue[datacloudrunservice.TemplateAttributes] {
	return terra.ReferenceAsList[datacloudrunservice.TemplateAttributes](crs.ref.Append("template"))
}

func (crs dataCloudRunServiceAttributes) Traffic() terra.ListValue[datacloudrunservice.TrafficAttributes] {
	return terra.ReferenceAsList[datacloudrunservice.TrafficAttributes](crs.ref.Append("traffic"))
}
