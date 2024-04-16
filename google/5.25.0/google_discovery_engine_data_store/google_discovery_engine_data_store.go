// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_discovery_engine_data_store

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// New creates a new instance of [Resource].
func New(name string, args Args) *Resource {
	return &Resource{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Resource)(nil)

// Resource represents the Terraform resource google_discovery_engine_data_store.
type Resource struct {
	Name      string
	Args      Args
	state     *googleDiscoveryEngineDataStoreState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (gdeds *Resource) Type() string {
	return "google_discovery_engine_data_store"
}

// LocalName returns the local name for [Resource].
func (gdeds *Resource) LocalName() string {
	return gdeds.Name
}

// Configuration returns the configuration (args) for [Resource].
func (gdeds *Resource) Configuration() interface{} {
	return gdeds.Args
}

// DependOn is used for other resources to depend on [Resource].
func (gdeds *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(gdeds)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (gdeds *Resource) Dependencies() terra.Dependencies {
	return gdeds.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (gdeds *Resource) LifecycleManagement() *terra.Lifecycle {
	return gdeds.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (gdeds *Resource) Attributes() googleDiscoveryEngineDataStoreAttributes {
	return googleDiscoveryEngineDataStoreAttributes{ref: terra.ReferenceResource(gdeds)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (gdeds *Resource) ImportState(state io.Reader) error {
	gdeds.state = &googleDiscoveryEngineDataStoreState{}
	if err := json.NewDecoder(state).Decode(gdeds.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", gdeds.Type(), gdeds.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (gdeds *Resource) State() (*googleDiscoveryEngineDataStoreState, bool) {
	return gdeds.state, gdeds.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (gdeds *Resource) StateMust() *googleDiscoveryEngineDataStoreState {
	if gdeds.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", gdeds.Type(), gdeds.LocalName()))
	}
	return gdeds.state
}

// Args contains the configurations for google_discovery_engine_data_store.
type Args struct {
	// ContentConfig: string, required
	ContentConfig terra.StringValue `hcl:"content_config,attr" validate:"required"`
	// CreateAdvancedSiteSearch: bool, optional
	CreateAdvancedSiteSearch terra.BoolValue `hcl:"create_advanced_site_search,attr"`
	// DataStoreId: string, required
	DataStoreId terra.StringValue `hcl:"data_store_id,attr" validate:"required"`
	// DisplayName: string, required
	DisplayName terra.StringValue `hcl:"display_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IndustryVertical: string, required
	IndustryVertical terra.StringValue `hcl:"industry_vertical,attr" validate:"required"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// SolutionTypes: list of string, optional
	SolutionTypes terra.ListValue[terra.StringValue] `hcl:"solution_types,attr"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type googleDiscoveryEngineDataStoreAttributes struct {
	ref terra.Reference
}

// ContentConfig returns a reference to field content_config of google_discovery_engine_data_store.
func (gdeds googleDiscoveryEngineDataStoreAttributes) ContentConfig() terra.StringValue {
	return terra.ReferenceAsString(gdeds.ref.Append("content_config"))
}

// CreateAdvancedSiteSearch returns a reference to field create_advanced_site_search of google_discovery_engine_data_store.
func (gdeds googleDiscoveryEngineDataStoreAttributes) CreateAdvancedSiteSearch() terra.BoolValue {
	return terra.ReferenceAsBool(gdeds.ref.Append("create_advanced_site_search"))
}

// CreateTime returns a reference to field create_time of google_discovery_engine_data_store.
func (gdeds googleDiscoveryEngineDataStoreAttributes) CreateTime() terra.StringValue {
	return terra.ReferenceAsString(gdeds.ref.Append("create_time"))
}

// DataStoreId returns a reference to field data_store_id of google_discovery_engine_data_store.
func (gdeds googleDiscoveryEngineDataStoreAttributes) DataStoreId() terra.StringValue {
	return terra.ReferenceAsString(gdeds.ref.Append("data_store_id"))
}

// DefaultSchemaId returns a reference to field default_schema_id of google_discovery_engine_data_store.
func (gdeds googleDiscoveryEngineDataStoreAttributes) DefaultSchemaId() terra.StringValue {
	return terra.ReferenceAsString(gdeds.ref.Append("default_schema_id"))
}

// DisplayName returns a reference to field display_name of google_discovery_engine_data_store.
func (gdeds googleDiscoveryEngineDataStoreAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(gdeds.ref.Append("display_name"))
}

// Id returns a reference to field id of google_discovery_engine_data_store.
func (gdeds googleDiscoveryEngineDataStoreAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gdeds.ref.Append("id"))
}

// IndustryVertical returns a reference to field industry_vertical of google_discovery_engine_data_store.
func (gdeds googleDiscoveryEngineDataStoreAttributes) IndustryVertical() terra.StringValue {
	return terra.ReferenceAsString(gdeds.ref.Append("industry_vertical"))
}

// Location returns a reference to field location of google_discovery_engine_data_store.
func (gdeds googleDiscoveryEngineDataStoreAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(gdeds.ref.Append("location"))
}

// Name returns a reference to field name of google_discovery_engine_data_store.
func (gdeds googleDiscoveryEngineDataStoreAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(gdeds.ref.Append("name"))
}

// Project returns a reference to field project of google_discovery_engine_data_store.
func (gdeds googleDiscoveryEngineDataStoreAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(gdeds.ref.Append("project"))
}

// SolutionTypes returns a reference to field solution_types of google_discovery_engine_data_store.
func (gdeds googleDiscoveryEngineDataStoreAttributes) SolutionTypes() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](gdeds.ref.Append("solution_types"))
}

func (gdeds googleDiscoveryEngineDataStoreAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](gdeds.ref.Append("timeouts"))
}

type googleDiscoveryEngineDataStoreState struct {
	ContentConfig            string         `json:"content_config"`
	CreateAdvancedSiteSearch bool           `json:"create_advanced_site_search"`
	CreateTime               string         `json:"create_time"`
	DataStoreId              string         `json:"data_store_id"`
	DefaultSchemaId          string         `json:"default_schema_id"`
	DisplayName              string         `json:"display_name"`
	Id                       string         `json:"id"`
	IndustryVertical         string         `json:"industry_vertical"`
	Location                 string         `json:"location"`
	Name                     string         `json:"name"`
	Project                  string         `json:"project"`
	SolutionTypes            []string       `json:"solution_types"`
	Timeouts                 *TimeoutsState `json:"timeouts"`
}
