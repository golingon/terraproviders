// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google

import (
	"github.com/golingon/lingon/pkg/terra"
	datacloudrunv2service "github.com/golingon/terraproviders/google/5.24.0/datacloudrunv2service"
)

// NewDataCloudRunV2Service creates a new instance of [DataCloudRunV2Service].
func NewDataCloudRunV2Service(name string, args DataCloudRunV2ServiceArgs) *DataCloudRunV2Service {
	return &DataCloudRunV2Service{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataCloudRunV2Service)(nil)

// DataCloudRunV2Service represents the Terraform data resource google_cloud_run_v2_service.
type DataCloudRunV2Service struct {
	Name string
	Args DataCloudRunV2ServiceArgs
}

// DataSource returns the Terraform object type for [DataCloudRunV2Service].
func (crvs *DataCloudRunV2Service) DataSource() string {
	return "google_cloud_run_v2_service"
}

// LocalName returns the local name for [DataCloudRunV2Service].
func (crvs *DataCloudRunV2Service) LocalName() string {
	return crvs.Name
}

// Configuration returns the configuration (args) for [DataCloudRunV2Service].
func (crvs *DataCloudRunV2Service) Configuration() interface{} {
	return crvs.Args
}

// Attributes returns the attributes for [DataCloudRunV2Service].
func (crvs *DataCloudRunV2Service) Attributes() dataCloudRunV2ServiceAttributes {
	return dataCloudRunV2ServiceAttributes{ref: terra.ReferenceDataResource(crvs)}
}

// DataCloudRunV2ServiceArgs contains the configurations for google_cloud_run_v2_service.
type DataCloudRunV2ServiceArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, optional
	Location terra.StringValue `hcl:"location,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// BinaryAuthorization: min=0
	BinaryAuthorization []datacloudrunv2service.BinaryAuthorization `hcl:"binary_authorization,block" validate:"min=0"`
	// Conditions: min=0
	Conditions []datacloudrunv2service.Conditions `hcl:"conditions,block" validate:"min=0"`
	// Template: min=0
	Template []datacloudrunv2service.Template `hcl:"template,block" validate:"min=0"`
	// TerminalCondition: min=0
	TerminalCondition []datacloudrunv2service.TerminalCondition `hcl:"terminal_condition,block" validate:"min=0"`
	// Traffic: min=0
	Traffic []datacloudrunv2service.Traffic `hcl:"traffic,block" validate:"min=0"`
	// TrafficStatuses: min=0
	TrafficStatuses []datacloudrunv2service.TrafficStatuses `hcl:"traffic_statuses,block" validate:"min=0"`
}
type dataCloudRunV2ServiceAttributes struct {
	ref terra.Reference
}

// Annotations returns a reference to field annotations of google_cloud_run_v2_service.
func (crvs dataCloudRunV2ServiceAttributes) Annotations() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](crvs.ref.Append("annotations"))
}

// Client returns a reference to field client of google_cloud_run_v2_service.
func (crvs dataCloudRunV2ServiceAttributes) Client() terra.StringValue {
	return terra.ReferenceAsString(crvs.ref.Append("client"))
}

// ClientVersion returns a reference to field client_version of google_cloud_run_v2_service.
func (crvs dataCloudRunV2ServiceAttributes) ClientVersion() terra.StringValue {
	return terra.ReferenceAsString(crvs.ref.Append("client_version"))
}

// CreateTime returns a reference to field create_time of google_cloud_run_v2_service.
func (crvs dataCloudRunV2ServiceAttributes) CreateTime() terra.StringValue {
	return terra.ReferenceAsString(crvs.ref.Append("create_time"))
}

// Creator returns a reference to field creator of google_cloud_run_v2_service.
func (crvs dataCloudRunV2ServiceAttributes) Creator() terra.StringValue {
	return terra.ReferenceAsString(crvs.ref.Append("creator"))
}

// CustomAudiences returns a reference to field custom_audiences of google_cloud_run_v2_service.
func (crvs dataCloudRunV2ServiceAttributes) CustomAudiences() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](crvs.ref.Append("custom_audiences"))
}

// DeleteTime returns a reference to field delete_time of google_cloud_run_v2_service.
func (crvs dataCloudRunV2ServiceAttributes) DeleteTime() terra.StringValue {
	return terra.ReferenceAsString(crvs.ref.Append("delete_time"))
}

// Description returns a reference to field description of google_cloud_run_v2_service.
func (crvs dataCloudRunV2ServiceAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(crvs.ref.Append("description"))
}

// EffectiveAnnotations returns a reference to field effective_annotations of google_cloud_run_v2_service.
func (crvs dataCloudRunV2ServiceAttributes) EffectiveAnnotations() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](crvs.ref.Append("effective_annotations"))
}

// EffectiveLabels returns a reference to field effective_labels of google_cloud_run_v2_service.
func (crvs dataCloudRunV2ServiceAttributes) EffectiveLabels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](crvs.ref.Append("effective_labels"))
}

// Etag returns a reference to field etag of google_cloud_run_v2_service.
func (crvs dataCloudRunV2ServiceAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(crvs.ref.Append("etag"))
}

// ExpireTime returns a reference to field expire_time of google_cloud_run_v2_service.
func (crvs dataCloudRunV2ServiceAttributes) ExpireTime() terra.StringValue {
	return terra.ReferenceAsString(crvs.ref.Append("expire_time"))
}

// Generation returns a reference to field generation of google_cloud_run_v2_service.
func (crvs dataCloudRunV2ServiceAttributes) Generation() terra.StringValue {
	return terra.ReferenceAsString(crvs.ref.Append("generation"))
}

// Id returns a reference to field id of google_cloud_run_v2_service.
func (crvs dataCloudRunV2ServiceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(crvs.ref.Append("id"))
}

// Ingress returns a reference to field ingress of google_cloud_run_v2_service.
func (crvs dataCloudRunV2ServiceAttributes) Ingress() terra.StringValue {
	return terra.ReferenceAsString(crvs.ref.Append("ingress"))
}

// Labels returns a reference to field labels of google_cloud_run_v2_service.
func (crvs dataCloudRunV2ServiceAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](crvs.ref.Append("labels"))
}

// LastModifier returns a reference to field last_modifier of google_cloud_run_v2_service.
func (crvs dataCloudRunV2ServiceAttributes) LastModifier() terra.StringValue {
	return terra.ReferenceAsString(crvs.ref.Append("last_modifier"))
}

// LatestCreatedRevision returns a reference to field latest_created_revision of google_cloud_run_v2_service.
func (crvs dataCloudRunV2ServiceAttributes) LatestCreatedRevision() terra.StringValue {
	return terra.ReferenceAsString(crvs.ref.Append("latest_created_revision"))
}

// LatestReadyRevision returns a reference to field latest_ready_revision of google_cloud_run_v2_service.
func (crvs dataCloudRunV2ServiceAttributes) LatestReadyRevision() terra.StringValue {
	return terra.ReferenceAsString(crvs.ref.Append("latest_ready_revision"))
}

// LaunchStage returns a reference to field launch_stage of google_cloud_run_v2_service.
func (crvs dataCloudRunV2ServiceAttributes) LaunchStage() terra.StringValue {
	return terra.ReferenceAsString(crvs.ref.Append("launch_stage"))
}

// Location returns a reference to field location of google_cloud_run_v2_service.
func (crvs dataCloudRunV2ServiceAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(crvs.ref.Append("location"))
}

// Name returns a reference to field name of google_cloud_run_v2_service.
func (crvs dataCloudRunV2ServiceAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(crvs.ref.Append("name"))
}

// ObservedGeneration returns a reference to field observed_generation of google_cloud_run_v2_service.
func (crvs dataCloudRunV2ServiceAttributes) ObservedGeneration() terra.StringValue {
	return terra.ReferenceAsString(crvs.ref.Append("observed_generation"))
}

// Project returns a reference to field project of google_cloud_run_v2_service.
func (crvs dataCloudRunV2ServiceAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(crvs.ref.Append("project"))
}

// Reconciling returns a reference to field reconciling of google_cloud_run_v2_service.
func (crvs dataCloudRunV2ServiceAttributes) Reconciling() terra.BoolValue {
	return terra.ReferenceAsBool(crvs.ref.Append("reconciling"))
}

// TerraformLabels returns a reference to field terraform_labels of google_cloud_run_v2_service.
func (crvs dataCloudRunV2ServiceAttributes) TerraformLabels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](crvs.ref.Append("terraform_labels"))
}

// Uid returns a reference to field uid of google_cloud_run_v2_service.
func (crvs dataCloudRunV2ServiceAttributes) Uid() terra.StringValue {
	return terra.ReferenceAsString(crvs.ref.Append("uid"))
}

// UpdateTime returns a reference to field update_time of google_cloud_run_v2_service.
func (crvs dataCloudRunV2ServiceAttributes) UpdateTime() terra.StringValue {
	return terra.ReferenceAsString(crvs.ref.Append("update_time"))
}

// Uri returns a reference to field uri of google_cloud_run_v2_service.
func (crvs dataCloudRunV2ServiceAttributes) Uri() terra.StringValue {
	return terra.ReferenceAsString(crvs.ref.Append("uri"))
}

func (crvs dataCloudRunV2ServiceAttributes) BinaryAuthorization() terra.ListValue[datacloudrunv2service.BinaryAuthorizationAttributes] {
	return terra.ReferenceAsList[datacloudrunv2service.BinaryAuthorizationAttributes](crvs.ref.Append("binary_authorization"))
}

func (crvs dataCloudRunV2ServiceAttributes) Conditions() terra.ListValue[datacloudrunv2service.ConditionsAttributes] {
	return terra.ReferenceAsList[datacloudrunv2service.ConditionsAttributes](crvs.ref.Append("conditions"))
}

func (crvs dataCloudRunV2ServiceAttributes) Template() terra.ListValue[datacloudrunv2service.TemplateAttributes] {
	return terra.ReferenceAsList[datacloudrunv2service.TemplateAttributes](crvs.ref.Append("template"))
}

func (crvs dataCloudRunV2ServiceAttributes) TerminalCondition() terra.ListValue[datacloudrunv2service.TerminalConditionAttributes] {
	return terra.ReferenceAsList[datacloudrunv2service.TerminalConditionAttributes](crvs.ref.Append("terminal_condition"))
}

func (crvs dataCloudRunV2ServiceAttributes) Traffic() terra.ListValue[datacloudrunv2service.TrafficAttributes] {
	return terra.ReferenceAsList[datacloudrunv2service.TrafficAttributes](crvs.ref.Append("traffic"))
}

func (crvs dataCloudRunV2ServiceAttributes) TrafficStatuses() terra.ListValue[datacloudrunv2service.TrafficStatusesAttributes] {
	return terra.ReferenceAsList[datacloudrunv2service.TrafficStatusesAttributes](crvs.ref.Append("traffic_statuses"))
}
