// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	cloudrunservice "github.com/golingon/terraproviders/google/4.72.1/cloudrunservice"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewCloudRunService creates a new instance of [CloudRunService].
func NewCloudRunService(name string, args CloudRunServiceArgs) *CloudRunService {
	return &CloudRunService{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*CloudRunService)(nil)

// CloudRunService represents the Terraform resource google_cloud_run_service.
type CloudRunService struct {
	Name      string
	Args      CloudRunServiceArgs
	state     *cloudRunServiceState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [CloudRunService].
func (crs *CloudRunService) Type() string {
	return "google_cloud_run_service"
}

// LocalName returns the local name for [CloudRunService].
func (crs *CloudRunService) LocalName() string {
	return crs.Name
}

// Configuration returns the configuration (args) for [CloudRunService].
func (crs *CloudRunService) Configuration() interface{} {
	return crs.Args
}

// DependOn is used for other resources to depend on [CloudRunService].
func (crs *CloudRunService) DependOn() terra.Reference {
	return terra.ReferenceResource(crs)
}

// Dependencies returns the list of resources [CloudRunService] depends_on.
func (crs *CloudRunService) Dependencies() terra.Dependencies {
	return crs.DependsOn
}

// LifecycleManagement returns the lifecycle block for [CloudRunService].
func (crs *CloudRunService) LifecycleManagement() *terra.Lifecycle {
	return crs.Lifecycle
}

// Attributes returns the attributes for [CloudRunService].
func (crs *CloudRunService) Attributes() cloudRunServiceAttributes {
	return cloudRunServiceAttributes{ref: terra.ReferenceResource(crs)}
}

// ImportState imports the given attribute values into [CloudRunService]'s state.
func (crs *CloudRunService) ImportState(av io.Reader) error {
	crs.state = &cloudRunServiceState{}
	if err := json.NewDecoder(av).Decode(crs.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", crs.Type(), crs.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [CloudRunService] has state.
func (crs *CloudRunService) State() (*cloudRunServiceState, bool) {
	return crs.state, crs.state != nil
}

// StateMust returns the state for [CloudRunService]. Panics if the state is nil.
func (crs *CloudRunService) StateMust() *cloudRunServiceState {
	if crs.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", crs.Type(), crs.LocalName()))
	}
	return crs.state
}

// CloudRunServiceArgs contains the configurations for google_cloud_run_service.
type CloudRunServiceArgs struct {
	// AutogenerateRevisionName: bool, optional
	AutogenerateRevisionName terra.BoolValue `hcl:"autogenerate_revision_name,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Status: min=0
	Status []cloudrunservice.Status `hcl:"status,block" validate:"min=0"`
	// Metadata: optional
	Metadata *cloudrunservice.Metadata `hcl:"metadata,block"`
	// Template: optional
	Template *cloudrunservice.Template `hcl:"template,block"`
	// Timeouts: optional
	Timeouts *cloudrunservice.Timeouts `hcl:"timeouts,block"`
	// Traffic: min=0
	Traffic []cloudrunservice.Traffic `hcl:"traffic,block" validate:"min=0"`
}
type cloudRunServiceAttributes struct {
	ref terra.Reference
}

// AutogenerateRevisionName returns a reference to field autogenerate_revision_name of google_cloud_run_service.
func (crs cloudRunServiceAttributes) AutogenerateRevisionName() terra.BoolValue {
	return terra.ReferenceAsBool(crs.ref.Append("autogenerate_revision_name"))
}

// Id returns a reference to field id of google_cloud_run_service.
func (crs cloudRunServiceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(crs.ref.Append("id"))
}

// Location returns a reference to field location of google_cloud_run_service.
func (crs cloudRunServiceAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(crs.ref.Append("location"))
}

// Name returns a reference to field name of google_cloud_run_service.
func (crs cloudRunServiceAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(crs.ref.Append("name"))
}

// Project returns a reference to field project of google_cloud_run_service.
func (crs cloudRunServiceAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(crs.ref.Append("project"))
}

func (crs cloudRunServiceAttributes) Status() terra.ListValue[cloudrunservice.StatusAttributes] {
	return terra.ReferenceAsList[cloudrunservice.StatusAttributes](crs.ref.Append("status"))
}

func (crs cloudRunServiceAttributes) Metadata() terra.ListValue[cloudrunservice.MetadataAttributes] {
	return terra.ReferenceAsList[cloudrunservice.MetadataAttributes](crs.ref.Append("metadata"))
}

func (crs cloudRunServiceAttributes) Template() terra.ListValue[cloudrunservice.TemplateAttributes] {
	return terra.ReferenceAsList[cloudrunservice.TemplateAttributes](crs.ref.Append("template"))
}

func (crs cloudRunServiceAttributes) Timeouts() cloudrunservice.TimeoutsAttributes {
	return terra.ReferenceAsSingle[cloudrunservice.TimeoutsAttributes](crs.ref.Append("timeouts"))
}

func (crs cloudRunServiceAttributes) Traffic() terra.ListValue[cloudrunservice.TrafficAttributes] {
	return terra.ReferenceAsList[cloudrunservice.TrafficAttributes](crs.ref.Append("traffic"))
}

type cloudRunServiceState struct {
	AutogenerateRevisionName bool                            `json:"autogenerate_revision_name"`
	Id                       string                          `json:"id"`
	Location                 string                          `json:"location"`
	Name                     string                          `json:"name"`
	Project                  string                          `json:"project"`
	Status                   []cloudrunservice.StatusState   `json:"status"`
	Metadata                 []cloudrunservice.MetadataState `json:"metadata"`
	Template                 []cloudrunservice.TemplateState `json:"template"`
	Timeouts                 *cloudrunservice.TimeoutsState  `json:"timeouts"`
	Traffic                  []cloudrunservice.TrafficState  `json:"traffic"`
}
