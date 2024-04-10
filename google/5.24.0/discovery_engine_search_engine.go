// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	discoveryenginesearchengine "github.com/golingon/terraproviders/google/5.24.0/discoveryenginesearchengine"
	"io"
)

// NewDiscoveryEngineSearchEngine creates a new instance of [DiscoveryEngineSearchEngine].
func NewDiscoveryEngineSearchEngine(name string, args DiscoveryEngineSearchEngineArgs) *DiscoveryEngineSearchEngine {
	return &DiscoveryEngineSearchEngine{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DiscoveryEngineSearchEngine)(nil)

// DiscoveryEngineSearchEngine represents the Terraform resource google_discovery_engine_search_engine.
type DiscoveryEngineSearchEngine struct {
	Name      string
	Args      DiscoveryEngineSearchEngineArgs
	state     *discoveryEngineSearchEngineState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DiscoveryEngineSearchEngine].
func (dese *DiscoveryEngineSearchEngine) Type() string {
	return "google_discovery_engine_search_engine"
}

// LocalName returns the local name for [DiscoveryEngineSearchEngine].
func (dese *DiscoveryEngineSearchEngine) LocalName() string {
	return dese.Name
}

// Configuration returns the configuration (args) for [DiscoveryEngineSearchEngine].
func (dese *DiscoveryEngineSearchEngine) Configuration() interface{} {
	return dese.Args
}

// DependOn is used for other resources to depend on [DiscoveryEngineSearchEngine].
func (dese *DiscoveryEngineSearchEngine) DependOn() terra.Reference {
	return terra.ReferenceResource(dese)
}

// Dependencies returns the list of resources [DiscoveryEngineSearchEngine] depends_on.
func (dese *DiscoveryEngineSearchEngine) Dependencies() terra.Dependencies {
	return dese.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DiscoveryEngineSearchEngine].
func (dese *DiscoveryEngineSearchEngine) LifecycleManagement() *terra.Lifecycle {
	return dese.Lifecycle
}

// Attributes returns the attributes for [DiscoveryEngineSearchEngine].
func (dese *DiscoveryEngineSearchEngine) Attributes() discoveryEngineSearchEngineAttributes {
	return discoveryEngineSearchEngineAttributes{ref: terra.ReferenceResource(dese)}
}

// ImportState imports the given attribute values into [DiscoveryEngineSearchEngine]'s state.
func (dese *DiscoveryEngineSearchEngine) ImportState(av io.Reader) error {
	dese.state = &discoveryEngineSearchEngineState{}
	if err := json.NewDecoder(av).Decode(dese.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dese.Type(), dese.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DiscoveryEngineSearchEngine] has state.
func (dese *DiscoveryEngineSearchEngine) State() (*discoveryEngineSearchEngineState, bool) {
	return dese.state, dese.state != nil
}

// StateMust returns the state for [DiscoveryEngineSearchEngine]. Panics if the state is nil.
func (dese *DiscoveryEngineSearchEngine) StateMust() *discoveryEngineSearchEngineState {
	if dese.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dese.Type(), dese.LocalName()))
	}
	return dese.state
}

// DiscoveryEngineSearchEngineArgs contains the configurations for google_discovery_engine_search_engine.
type DiscoveryEngineSearchEngineArgs struct {
	// CollectionId: string, required
	CollectionId terra.StringValue `hcl:"collection_id,attr" validate:"required"`
	// DataStoreIds: list of string, required
	DataStoreIds terra.ListValue[terra.StringValue] `hcl:"data_store_ids,attr" validate:"required"`
	// DisplayName: string, required
	DisplayName terra.StringValue `hcl:"display_name,attr" validate:"required"`
	// EngineId: string, required
	EngineId terra.StringValue `hcl:"engine_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IndustryVertical: string, optional
	IndustryVertical terra.StringValue `hcl:"industry_vertical,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// CommonConfig: optional
	CommonConfig *discoveryenginesearchengine.CommonConfig `hcl:"common_config,block"`
	// SearchEngineConfig: required
	SearchEngineConfig *discoveryenginesearchengine.SearchEngineConfig `hcl:"search_engine_config,block" validate:"required"`
	// Timeouts: optional
	Timeouts *discoveryenginesearchengine.Timeouts `hcl:"timeouts,block"`
}
type discoveryEngineSearchEngineAttributes struct {
	ref terra.Reference
}

// CollectionId returns a reference to field collection_id of google_discovery_engine_search_engine.
func (dese discoveryEngineSearchEngineAttributes) CollectionId() terra.StringValue {
	return terra.ReferenceAsString(dese.ref.Append("collection_id"))
}

// CreateTime returns a reference to field create_time of google_discovery_engine_search_engine.
func (dese discoveryEngineSearchEngineAttributes) CreateTime() terra.StringValue {
	return terra.ReferenceAsString(dese.ref.Append("create_time"))
}

// DataStoreIds returns a reference to field data_store_ids of google_discovery_engine_search_engine.
func (dese discoveryEngineSearchEngineAttributes) DataStoreIds() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](dese.ref.Append("data_store_ids"))
}

// DisplayName returns a reference to field display_name of google_discovery_engine_search_engine.
func (dese discoveryEngineSearchEngineAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(dese.ref.Append("display_name"))
}

// EngineId returns a reference to field engine_id of google_discovery_engine_search_engine.
func (dese discoveryEngineSearchEngineAttributes) EngineId() terra.StringValue {
	return terra.ReferenceAsString(dese.ref.Append("engine_id"))
}

// Id returns a reference to field id of google_discovery_engine_search_engine.
func (dese discoveryEngineSearchEngineAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dese.ref.Append("id"))
}

// IndustryVertical returns a reference to field industry_vertical of google_discovery_engine_search_engine.
func (dese discoveryEngineSearchEngineAttributes) IndustryVertical() terra.StringValue {
	return terra.ReferenceAsString(dese.ref.Append("industry_vertical"))
}

// Location returns a reference to field location of google_discovery_engine_search_engine.
func (dese discoveryEngineSearchEngineAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(dese.ref.Append("location"))
}

// Name returns a reference to field name of google_discovery_engine_search_engine.
func (dese discoveryEngineSearchEngineAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(dese.ref.Append("name"))
}

// Project returns a reference to field project of google_discovery_engine_search_engine.
func (dese discoveryEngineSearchEngineAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(dese.ref.Append("project"))
}

// UpdateTime returns a reference to field update_time of google_discovery_engine_search_engine.
func (dese discoveryEngineSearchEngineAttributes) UpdateTime() terra.StringValue {
	return terra.ReferenceAsString(dese.ref.Append("update_time"))
}

func (dese discoveryEngineSearchEngineAttributes) CommonConfig() terra.ListValue[discoveryenginesearchengine.CommonConfigAttributes] {
	return terra.ReferenceAsList[discoveryenginesearchengine.CommonConfigAttributes](dese.ref.Append("common_config"))
}

func (dese discoveryEngineSearchEngineAttributes) SearchEngineConfig() terra.ListValue[discoveryenginesearchengine.SearchEngineConfigAttributes] {
	return terra.ReferenceAsList[discoveryenginesearchengine.SearchEngineConfigAttributes](dese.ref.Append("search_engine_config"))
}

func (dese discoveryEngineSearchEngineAttributes) Timeouts() discoveryenginesearchengine.TimeoutsAttributes {
	return terra.ReferenceAsSingle[discoveryenginesearchengine.TimeoutsAttributes](dese.ref.Append("timeouts"))
}

type discoveryEngineSearchEngineState struct {
	CollectionId       string                                                `json:"collection_id"`
	CreateTime         string                                                `json:"create_time"`
	DataStoreIds       []string                                              `json:"data_store_ids"`
	DisplayName        string                                                `json:"display_name"`
	EngineId           string                                                `json:"engine_id"`
	Id                 string                                                `json:"id"`
	IndustryVertical   string                                                `json:"industry_vertical"`
	Location           string                                                `json:"location"`
	Name               string                                                `json:"name"`
	Project            string                                                `json:"project"`
	UpdateTime         string                                                `json:"update_time"`
	CommonConfig       []discoveryenginesearchengine.CommonConfigState       `json:"common_config"`
	SearchEngineConfig []discoveryenginesearchengine.SearchEngineConfigState `json:"search_engine_config"`
	Timeouts           *discoveryenginesearchengine.TimeoutsState            `json:"timeouts"`
}
