// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	cloudrunv2service "github.com/golingon/terraproviders/google/4.59.0/cloudrunv2service"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewCloudRunV2Service creates a new instance of [CloudRunV2Service].
func NewCloudRunV2Service(name string, args CloudRunV2ServiceArgs) *CloudRunV2Service {
	return &CloudRunV2Service{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*CloudRunV2Service)(nil)

// CloudRunV2Service represents the Terraform resource google_cloud_run_v2_service.
type CloudRunV2Service struct {
	Name      string
	Args      CloudRunV2ServiceArgs
	state     *cloudRunV2ServiceState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [CloudRunV2Service].
func (crvs *CloudRunV2Service) Type() string {
	return "google_cloud_run_v2_service"
}

// LocalName returns the local name for [CloudRunV2Service].
func (crvs *CloudRunV2Service) LocalName() string {
	return crvs.Name
}

// Configuration returns the configuration (args) for [CloudRunV2Service].
func (crvs *CloudRunV2Service) Configuration() interface{} {
	return crvs.Args
}

// DependOn is used for other resources to depend on [CloudRunV2Service].
func (crvs *CloudRunV2Service) DependOn() terra.Reference {
	return terra.ReferenceResource(crvs)
}

// Dependencies returns the list of resources [CloudRunV2Service] depends_on.
func (crvs *CloudRunV2Service) Dependencies() terra.Dependencies {
	return crvs.DependsOn
}

// LifecycleManagement returns the lifecycle block for [CloudRunV2Service].
func (crvs *CloudRunV2Service) LifecycleManagement() *terra.Lifecycle {
	return crvs.Lifecycle
}

// Attributes returns the attributes for [CloudRunV2Service].
func (crvs *CloudRunV2Service) Attributes() cloudRunV2ServiceAttributes {
	return cloudRunV2ServiceAttributes{ref: terra.ReferenceResource(crvs)}
}

// ImportState imports the given attribute values into [CloudRunV2Service]'s state.
func (crvs *CloudRunV2Service) ImportState(av io.Reader) error {
	crvs.state = &cloudRunV2ServiceState{}
	if err := json.NewDecoder(av).Decode(crvs.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", crvs.Type(), crvs.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [CloudRunV2Service] has state.
func (crvs *CloudRunV2Service) State() (*cloudRunV2ServiceState, bool) {
	return crvs.state, crvs.state != nil
}

// StateMust returns the state for [CloudRunV2Service]. Panics if the state is nil.
func (crvs *CloudRunV2Service) StateMust() *cloudRunV2ServiceState {
	if crvs.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", crvs.Type(), crvs.LocalName()))
	}
	return crvs.state
}

// CloudRunV2ServiceArgs contains the configurations for google_cloud_run_v2_service.
type CloudRunV2ServiceArgs struct {
	// Annotations: map of string, optional
	Annotations terra.MapValue[terra.StringValue] `hcl:"annotations,attr"`
	// Client: string, optional
	Client terra.StringValue `hcl:"client,attr"`
	// ClientVersion: string, optional
	ClientVersion terra.StringValue `hcl:"client_version,attr"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Ingress: string, optional
	Ingress terra.StringValue `hcl:"ingress,attr"`
	// Labels: map of string, optional
	Labels terra.MapValue[terra.StringValue] `hcl:"labels,attr"`
	// LaunchStage: string, optional
	LaunchStage terra.StringValue `hcl:"launch_stage,attr"`
	// Location: string, optional
	Location terra.StringValue `hcl:"location,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Conditions: min=0
	Conditions []cloudrunv2service.Conditions `hcl:"conditions,block" validate:"min=0"`
	// TerminalCondition: min=0
	TerminalCondition []cloudrunv2service.TerminalCondition `hcl:"terminal_condition,block" validate:"min=0"`
	// TrafficStatuses: min=0
	TrafficStatuses []cloudrunv2service.TrafficStatuses `hcl:"traffic_statuses,block" validate:"min=0"`
	// BinaryAuthorization: optional
	BinaryAuthorization *cloudrunv2service.BinaryAuthorization `hcl:"binary_authorization,block"`
	// Template: required
	Template *cloudrunv2service.Template `hcl:"template,block" validate:"required"`
	// Timeouts: optional
	Timeouts *cloudrunv2service.Timeouts `hcl:"timeouts,block"`
	// Traffic: min=0
	Traffic []cloudrunv2service.Traffic `hcl:"traffic,block" validate:"min=0"`
}
type cloudRunV2ServiceAttributes struct {
	ref terra.Reference
}

// Annotations returns a reference to field annotations of google_cloud_run_v2_service.
func (crvs cloudRunV2ServiceAttributes) Annotations() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](crvs.ref.Append("annotations"))
}

// Client returns a reference to field client of google_cloud_run_v2_service.
func (crvs cloudRunV2ServiceAttributes) Client() terra.StringValue {
	return terra.ReferenceAsString(crvs.ref.Append("client"))
}

// ClientVersion returns a reference to field client_version of google_cloud_run_v2_service.
func (crvs cloudRunV2ServiceAttributes) ClientVersion() terra.StringValue {
	return terra.ReferenceAsString(crvs.ref.Append("client_version"))
}

// Description returns a reference to field description of google_cloud_run_v2_service.
func (crvs cloudRunV2ServiceAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(crvs.ref.Append("description"))
}

// Etag returns a reference to field etag of google_cloud_run_v2_service.
func (crvs cloudRunV2ServiceAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(crvs.ref.Append("etag"))
}

// Generation returns a reference to field generation of google_cloud_run_v2_service.
func (crvs cloudRunV2ServiceAttributes) Generation() terra.StringValue {
	return terra.ReferenceAsString(crvs.ref.Append("generation"))
}

// Id returns a reference to field id of google_cloud_run_v2_service.
func (crvs cloudRunV2ServiceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(crvs.ref.Append("id"))
}

// Ingress returns a reference to field ingress of google_cloud_run_v2_service.
func (crvs cloudRunV2ServiceAttributes) Ingress() terra.StringValue {
	return terra.ReferenceAsString(crvs.ref.Append("ingress"))
}

// Labels returns a reference to field labels of google_cloud_run_v2_service.
func (crvs cloudRunV2ServiceAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](crvs.ref.Append("labels"))
}

// LatestCreatedRevision returns a reference to field latest_created_revision of google_cloud_run_v2_service.
func (crvs cloudRunV2ServiceAttributes) LatestCreatedRevision() terra.StringValue {
	return terra.ReferenceAsString(crvs.ref.Append("latest_created_revision"))
}

// LatestReadyRevision returns a reference to field latest_ready_revision of google_cloud_run_v2_service.
func (crvs cloudRunV2ServiceAttributes) LatestReadyRevision() terra.StringValue {
	return terra.ReferenceAsString(crvs.ref.Append("latest_ready_revision"))
}

// LaunchStage returns a reference to field launch_stage of google_cloud_run_v2_service.
func (crvs cloudRunV2ServiceAttributes) LaunchStage() terra.StringValue {
	return terra.ReferenceAsString(crvs.ref.Append("launch_stage"))
}

// Location returns a reference to field location of google_cloud_run_v2_service.
func (crvs cloudRunV2ServiceAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(crvs.ref.Append("location"))
}

// Name returns a reference to field name of google_cloud_run_v2_service.
func (crvs cloudRunV2ServiceAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(crvs.ref.Append("name"))
}

// ObservedGeneration returns a reference to field observed_generation of google_cloud_run_v2_service.
func (crvs cloudRunV2ServiceAttributes) ObservedGeneration() terra.StringValue {
	return terra.ReferenceAsString(crvs.ref.Append("observed_generation"))
}

// Project returns a reference to field project of google_cloud_run_v2_service.
func (crvs cloudRunV2ServiceAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(crvs.ref.Append("project"))
}

// Reconciling returns a reference to field reconciling of google_cloud_run_v2_service.
func (crvs cloudRunV2ServiceAttributes) Reconciling() terra.BoolValue {
	return terra.ReferenceAsBool(crvs.ref.Append("reconciling"))
}

// Uid returns a reference to field uid of google_cloud_run_v2_service.
func (crvs cloudRunV2ServiceAttributes) Uid() terra.StringValue {
	return terra.ReferenceAsString(crvs.ref.Append("uid"))
}

// Uri returns a reference to field uri of google_cloud_run_v2_service.
func (crvs cloudRunV2ServiceAttributes) Uri() terra.StringValue {
	return terra.ReferenceAsString(crvs.ref.Append("uri"))
}

func (crvs cloudRunV2ServiceAttributes) Conditions() terra.ListValue[cloudrunv2service.ConditionsAttributes] {
	return terra.ReferenceAsList[cloudrunv2service.ConditionsAttributes](crvs.ref.Append("conditions"))
}

func (crvs cloudRunV2ServiceAttributes) TerminalCondition() terra.ListValue[cloudrunv2service.TerminalConditionAttributes] {
	return terra.ReferenceAsList[cloudrunv2service.TerminalConditionAttributes](crvs.ref.Append("terminal_condition"))
}

func (crvs cloudRunV2ServiceAttributes) TrafficStatuses() terra.ListValue[cloudrunv2service.TrafficStatusesAttributes] {
	return terra.ReferenceAsList[cloudrunv2service.TrafficStatusesAttributes](crvs.ref.Append("traffic_statuses"))
}

func (crvs cloudRunV2ServiceAttributes) BinaryAuthorization() terra.ListValue[cloudrunv2service.BinaryAuthorizationAttributes] {
	return terra.ReferenceAsList[cloudrunv2service.BinaryAuthorizationAttributes](crvs.ref.Append("binary_authorization"))
}

func (crvs cloudRunV2ServiceAttributes) Template() terra.ListValue[cloudrunv2service.TemplateAttributes] {
	return terra.ReferenceAsList[cloudrunv2service.TemplateAttributes](crvs.ref.Append("template"))
}

func (crvs cloudRunV2ServiceAttributes) Timeouts() cloudrunv2service.TimeoutsAttributes {
	return terra.ReferenceAsSingle[cloudrunv2service.TimeoutsAttributes](crvs.ref.Append("timeouts"))
}

func (crvs cloudRunV2ServiceAttributes) Traffic() terra.ListValue[cloudrunv2service.TrafficAttributes] {
	return terra.ReferenceAsList[cloudrunv2service.TrafficAttributes](crvs.ref.Append("traffic"))
}

type cloudRunV2ServiceState struct {
	Annotations           map[string]string                            `json:"annotations"`
	Client                string                                       `json:"client"`
	ClientVersion         string                                       `json:"client_version"`
	Description           string                                       `json:"description"`
	Etag                  string                                       `json:"etag"`
	Generation            string                                       `json:"generation"`
	Id                    string                                       `json:"id"`
	Ingress               string                                       `json:"ingress"`
	Labels                map[string]string                            `json:"labels"`
	LatestCreatedRevision string                                       `json:"latest_created_revision"`
	LatestReadyRevision   string                                       `json:"latest_ready_revision"`
	LaunchStage           string                                       `json:"launch_stage"`
	Location              string                                       `json:"location"`
	Name                  string                                       `json:"name"`
	ObservedGeneration    string                                       `json:"observed_generation"`
	Project               string                                       `json:"project"`
	Reconciling           bool                                         `json:"reconciling"`
	Uid                   string                                       `json:"uid"`
	Uri                   string                                       `json:"uri"`
	Conditions            []cloudrunv2service.ConditionsState          `json:"conditions"`
	TerminalCondition     []cloudrunv2service.TerminalConditionState   `json:"terminal_condition"`
	TrafficStatuses       []cloudrunv2service.TrafficStatusesState     `json:"traffic_statuses"`
	BinaryAuthorization   []cloudrunv2service.BinaryAuthorizationState `json:"binary_authorization"`
	Template              []cloudrunv2service.TemplateState            `json:"template"`
	Timeouts              *cloudrunv2service.TimeoutsState             `json:"timeouts"`
	Traffic               []cloudrunv2service.TrafficState             `json:"traffic"`
}
