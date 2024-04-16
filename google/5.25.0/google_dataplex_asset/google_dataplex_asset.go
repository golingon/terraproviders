// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_dataplex_asset

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

// Resource represents the Terraform resource google_dataplex_asset.
type Resource struct {
	Name      string
	Args      Args
	state     *googleDataplexAssetState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (gda *Resource) Type() string {
	return "google_dataplex_asset"
}

// LocalName returns the local name for [Resource].
func (gda *Resource) LocalName() string {
	return gda.Name
}

// Configuration returns the configuration (args) for [Resource].
func (gda *Resource) Configuration() interface{} {
	return gda.Args
}

// DependOn is used for other resources to depend on [Resource].
func (gda *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(gda)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (gda *Resource) Dependencies() terra.Dependencies {
	return gda.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (gda *Resource) LifecycleManagement() *terra.Lifecycle {
	return gda.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (gda *Resource) Attributes() googleDataplexAssetAttributes {
	return googleDataplexAssetAttributes{ref: terra.ReferenceResource(gda)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (gda *Resource) ImportState(state io.Reader) error {
	gda.state = &googleDataplexAssetState{}
	if err := json.NewDecoder(state).Decode(gda.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", gda.Type(), gda.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (gda *Resource) State() (*googleDataplexAssetState, bool) {
	return gda.state, gda.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (gda *Resource) StateMust() *googleDataplexAssetState {
	if gda.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", gda.Type(), gda.LocalName()))
	}
	return gda.state
}

// Args contains the configurations for google_dataplex_asset.
type Args struct {
	// DataplexZone: string, required
	DataplexZone terra.StringValue `hcl:"dataplex_zone,attr" validate:"required"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// DisplayName: string, optional
	DisplayName terra.StringValue `hcl:"display_name,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Labels: map of string, optional
	Labels terra.MapValue[terra.StringValue] `hcl:"labels,attr"`
	// Lake: string, required
	Lake terra.StringValue `hcl:"lake,attr" validate:"required"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// DiscoverySpec: required
	DiscoverySpec *DiscoverySpec `hcl:"discovery_spec,block" validate:"required"`
	// ResourceSpec: required
	ResourceSpec *ResourceSpec `hcl:"resource_spec,block" validate:"required"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type googleDataplexAssetAttributes struct {
	ref terra.Reference
}

// CreateTime returns a reference to field create_time of google_dataplex_asset.
func (gda googleDataplexAssetAttributes) CreateTime() terra.StringValue {
	return terra.ReferenceAsString(gda.ref.Append("create_time"))
}

// DataplexZone returns a reference to field dataplex_zone of google_dataplex_asset.
func (gda googleDataplexAssetAttributes) DataplexZone() terra.StringValue {
	return terra.ReferenceAsString(gda.ref.Append("dataplex_zone"))
}

// Description returns a reference to field description of google_dataplex_asset.
func (gda googleDataplexAssetAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(gda.ref.Append("description"))
}

// DisplayName returns a reference to field display_name of google_dataplex_asset.
func (gda googleDataplexAssetAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(gda.ref.Append("display_name"))
}

// EffectiveLabels returns a reference to field effective_labels of google_dataplex_asset.
func (gda googleDataplexAssetAttributes) EffectiveLabels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](gda.ref.Append("effective_labels"))
}

// Id returns a reference to field id of google_dataplex_asset.
func (gda googleDataplexAssetAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gda.ref.Append("id"))
}

// Labels returns a reference to field labels of google_dataplex_asset.
func (gda googleDataplexAssetAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](gda.ref.Append("labels"))
}

// Lake returns a reference to field lake of google_dataplex_asset.
func (gda googleDataplexAssetAttributes) Lake() terra.StringValue {
	return terra.ReferenceAsString(gda.ref.Append("lake"))
}

// Location returns a reference to field location of google_dataplex_asset.
func (gda googleDataplexAssetAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(gda.ref.Append("location"))
}

// Name returns a reference to field name of google_dataplex_asset.
func (gda googleDataplexAssetAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(gda.ref.Append("name"))
}

// Project returns a reference to field project of google_dataplex_asset.
func (gda googleDataplexAssetAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(gda.ref.Append("project"))
}

// State returns a reference to field state of google_dataplex_asset.
func (gda googleDataplexAssetAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(gda.ref.Append("state"))
}

// TerraformLabels returns a reference to field terraform_labels of google_dataplex_asset.
func (gda googleDataplexAssetAttributes) TerraformLabels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](gda.ref.Append("terraform_labels"))
}

// Uid returns a reference to field uid of google_dataplex_asset.
func (gda googleDataplexAssetAttributes) Uid() terra.StringValue {
	return terra.ReferenceAsString(gda.ref.Append("uid"))
}

// UpdateTime returns a reference to field update_time of google_dataplex_asset.
func (gda googleDataplexAssetAttributes) UpdateTime() terra.StringValue {
	return terra.ReferenceAsString(gda.ref.Append("update_time"))
}

func (gda googleDataplexAssetAttributes) DiscoveryStatus() terra.ListValue[DiscoveryStatusAttributes] {
	return terra.ReferenceAsList[DiscoveryStatusAttributes](gda.ref.Append("discovery_status"))
}

func (gda googleDataplexAssetAttributes) ResourceStatus() terra.ListValue[ResourceStatusAttributes] {
	return terra.ReferenceAsList[ResourceStatusAttributes](gda.ref.Append("resource_status"))
}

func (gda googleDataplexAssetAttributes) SecurityStatus() terra.ListValue[SecurityStatusAttributes] {
	return terra.ReferenceAsList[SecurityStatusAttributes](gda.ref.Append("security_status"))
}

func (gda googleDataplexAssetAttributes) DiscoverySpec() terra.ListValue[DiscoverySpecAttributes] {
	return terra.ReferenceAsList[DiscoverySpecAttributes](gda.ref.Append("discovery_spec"))
}

func (gda googleDataplexAssetAttributes) ResourceSpec() terra.ListValue[ResourceSpecAttributes] {
	return terra.ReferenceAsList[ResourceSpecAttributes](gda.ref.Append("resource_spec"))
}

func (gda googleDataplexAssetAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](gda.ref.Append("timeouts"))
}

type googleDataplexAssetState struct {
	CreateTime      string                 `json:"create_time"`
	DataplexZone    string                 `json:"dataplex_zone"`
	Description     string                 `json:"description"`
	DisplayName     string                 `json:"display_name"`
	EffectiveLabels map[string]string      `json:"effective_labels"`
	Id              string                 `json:"id"`
	Labels          map[string]string      `json:"labels"`
	Lake            string                 `json:"lake"`
	Location        string                 `json:"location"`
	Name            string                 `json:"name"`
	Project         string                 `json:"project"`
	State           string                 `json:"state"`
	TerraformLabels map[string]string      `json:"terraform_labels"`
	Uid             string                 `json:"uid"`
	UpdateTime      string                 `json:"update_time"`
	DiscoveryStatus []DiscoveryStatusState `json:"discovery_status"`
	ResourceStatus  []ResourceStatusState  `json:"resource_status"`
	SecurityStatus  []SecurityStatusState  `json:"security_status"`
	DiscoverySpec   []DiscoverySpecState   `json:"discovery_spec"`
	ResourceSpec    []ResourceSpecState    `json:"resource_spec"`
	Timeouts        *TimeoutsState         `json:"timeouts"`
}
