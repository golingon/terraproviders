// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	vertexaifeaturestoreentitytype "github.com/golingon/terraproviders/googlebeta/4.77.0/vertexaifeaturestoreentitytype"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewVertexAiFeaturestoreEntitytype creates a new instance of [VertexAiFeaturestoreEntitytype].
func NewVertexAiFeaturestoreEntitytype(name string, args VertexAiFeaturestoreEntitytypeArgs) *VertexAiFeaturestoreEntitytype {
	return &VertexAiFeaturestoreEntitytype{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*VertexAiFeaturestoreEntitytype)(nil)

// VertexAiFeaturestoreEntitytype represents the Terraform resource google_vertex_ai_featurestore_entitytype.
type VertexAiFeaturestoreEntitytype struct {
	Name      string
	Args      VertexAiFeaturestoreEntitytypeArgs
	state     *vertexAiFeaturestoreEntitytypeState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [VertexAiFeaturestoreEntitytype].
func (vafe *VertexAiFeaturestoreEntitytype) Type() string {
	return "google_vertex_ai_featurestore_entitytype"
}

// LocalName returns the local name for [VertexAiFeaturestoreEntitytype].
func (vafe *VertexAiFeaturestoreEntitytype) LocalName() string {
	return vafe.Name
}

// Configuration returns the configuration (args) for [VertexAiFeaturestoreEntitytype].
func (vafe *VertexAiFeaturestoreEntitytype) Configuration() interface{} {
	return vafe.Args
}

// DependOn is used for other resources to depend on [VertexAiFeaturestoreEntitytype].
func (vafe *VertexAiFeaturestoreEntitytype) DependOn() terra.Reference {
	return terra.ReferenceResource(vafe)
}

// Dependencies returns the list of resources [VertexAiFeaturestoreEntitytype] depends_on.
func (vafe *VertexAiFeaturestoreEntitytype) Dependencies() terra.Dependencies {
	return vafe.DependsOn
}

// LifecycleManagement returns the lifecycle block for [VertexAiFeaturestoreEntitytype].
func (vafe *VertexAiFeaturestoreEntitytype) LifecycleManagement() *terra.Lifecycle {
	return vafe.Lifecycle
}

// Attributes returns the attributes for [VertexAiFeaturestoreEntitytype].
func (vafe *VertexAiFeaturestoreEntitytype) Attributes() vertexAiFeaturestoreEntitytypeAttributes {
	return vertexAiFeaturestoreEntitytypeAttributes{ref: terra.ReferenceResource(vafe)}
}

// ImportState imports the given attribute values into [VertexAiFeaturestoreEntitytype]'s state.
func (vafe *VertexAiFeaturestoreEntitytype) ImportState(av io.Reader) error {
	vafe.state = &vertexAiFeaturestoreEntitytypeState{}
	if err := json.NewDecoder(av).Decode(vafe.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", vafe.Type(), vafe.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [VertexAiFeaturestoreEntitytype] has state.
func (vafe *VertexAiFeaturestoreEntitytype) State() (*vertexAiFeaturestoreEntitytypeState, bool) {
	return vafe.state, vafe.state != nil
}

// StateMust returns the state for [VertexAiFeaturestoreEntitytype]. Panics if the state is nil.
func (vafe *VertexAiFeaturestoreEntitytype) StateMust() *vertexAiFeaturestoreEntitytypeState {
	if vafe.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", vafe.Type(), vafe.LocalName()))
	}
	return vafe.state
}

// VertexAiFeaturestoreEntitytypeArgs contains the configurations for google_vertex_ai_featurestore_entitytype.
type VertexAiFeaturestoreEntitytypeArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Featurestore: string, required
	Featurestore terra.StringValue `hcl:"featurestore,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Labels: map of string, optional
	Labels terra.MapValue[terra.StringValue] `hcl:"labels,attr"`
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
	// OfflineStorageTtlDays: number, optional
	OfflineStorageTtlDays terra.NumberValue `hcl:"offline_storage_ttl_days,attr"`
	// MonitoringConfig: optional
	MonitoringConfig *vertexaifeaturestoreentitytype.MonitoringConfig `hcl:"monitoring_config,block"`
	// Timeouts: optional
	Timeouts *vertexaifeaturestoreentitytype.Timeouts `hcl:"timeouts,block"`
}
type vertexAiFeaturestoreEntitytypeAttributes struct {
	ref terra.Reference
}

// CreateTime returns a reference to field create_time of google_vertex_ai_featurestore_entitytype.
func (vafe vertexAiFeaturestoreEntitytypeAttributes) CreateTime() terra.StringValue {
	return terra.ReferenceAsString(vafe.ref.Append("create_time"))
}

// Description returns a reference to field description of google_vertex_ai_featurestore_entitytype.
func (vafe vertexAiFeaturestoreEntitytypeAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(vafe.ref.Append("description"))
}

// Etag returns a reference to field etag of google_vertex_ai_featurestore_entitytype.
func (vafe vertexAiFeaturestoreEntitytypeAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(vafe.ref.Append("etag"))
}

// Featurestore returns a reference to field featurestore of google_vertex_ai_featurestore_entitytype.
func (vafe vertexAiFeaturestoreEntitytypeAttributes) Featurestore() terra.StringValue {
	return terra.ReferenceAsString(vafe.ref.Append("featurestore"))
}

// Id returns a reference to field id of google_vertex_ai_featurestore_entitytype.
func (vafe vertexAiFeaturestoreEntitytypeAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(vafe.ref.Append("id"))
}

// Labels returns a reference to field labels of google_vertex_ai_featurestore_entitytype.
func (vafe vertexAiFeaturestoreEntitytypeAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](vafe.ref.Append("labels"))
}

// Name returns a reference to field name of google_vertex_ai_featurestore_entitytype.
func (vafe vertexAiFeaturestoreEntitytypeAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(vafe.ref.Append("name"))
}

// OfflineStorageTtlDays returns a reference to field offline_storage_ttl_days of google_vertex_ai_featurestore_entitytype.
func (vafe vertexAiFeaturestoreEntitytypeAttributes) OfflineStorageTtlDays() terra.NumberValue {
	return terra.ReferenceAsNumber(vafe.ref.Append("offline_storage_ttl_days"))
}

// Region returns a reference to field region of google_vertex_ai_featurestore_entitytype.
func (vafe vertexAiFeaturestoreEntitytypeAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(vafe.ref.Append("region"))
}

// UpdateTime returns a reference to field update_time of google_vertex_ai_featurestore_entitytype.
func (vafe vertexAiFeaturestoreEntitytypeAttributes) UpdateTime() terra.StringValue {
	return terra.ReferenceAsString(vafe.ref.Append("update_time"))
}

func (vafe vertexAiFeaturestoreEntitytypeAttributes) MonitoringConfig() terra.ListValue[vertexaifeaturestoreentitytype.MonitoringConfigAttributes] {
	return terra.ReferenceAsList[vertexaifeaturestoreentitytype.MonitoringConfigAttributes](vafe.ref.Append("monitoring_config"))
}

func (vafe vertexAiFeaturestoreEntitytypeAttributes) Timeouts() vertexaifeaturestoreentitytype.TimeoutsAttributes {
	return terra.ReferenceAsSingle[vertexaifeaturestoreentitytype.TimeoutsAttributes](vafe.ref.Append("timeouts"))
}

type vertexAiFeaturestoreEntitytypeState struct {
	CreateTime            string                                                 `json:"create_time"`
	Description           string                                                 `json:"description"`
	Etag                  string                                                 `json:"etag"`
	Featurestore          string                                                 `json:"featurestore"`
	Id                    string                                                 `json:"id"`
	Labels                map[string]string                                      `json:"labels"`
	Name                  string                                                 `json:"name"`
	OfflineStorageTtlDays float64                                                `json:"offline_storage_ttl_days"`
	Region                string                                                 `json:"region"`
	UpdateTime            string                                                 `json:"update_time"`
	MonitoringConfig      []vertexaifeaturestoreentitytype.MonitoringConfigState `json:"monitoring_config"`
	Timeouts              *vertexaifeaturestoreentitytype.TimeoutsState          `json:"timeouts"`
}
