// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	dataplexasset "github.com/golingon/terraproviders/googlebeta/4.75.1/dataplexasset"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDataplexAsset creates a new instance of [DataplexAsset].
func NewDataplexAsset(name string, args DataplexAssetArgs) *DataplexAsset {
	return &DataplexAsset{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DataplexAsset)(nil)

// DataplexAsset represents the Terraform resource google_dataplex_asset.
type DataplexAsset struct {
	Name      string
	Args      DataplexAssetArgs
	state     *dataplexAssetState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DataplexAsset].
func (da *DataplexAsset) Type() string {
	return "google_dataplex_asset"
}

// LocalName returns the local name for [DataplexAsset].
func (da *DataplexAsset) LocalName() string {
	return da.Name
}

// Configuration returns the configuration (args) for [DataplexAsset].
func (da *DataplexAsset) Configuration() interface{} {
	return da.Args
}

// DependOn is used for other resources to depend on [DataplexAsset].
func (da *DataplexAsset) DependOn() terra.Reference {
	return terra.ReferenceResource(da)
}

// Dependencies returns the list of resources [DataplexAsset] depends_on.
func (da *DataplexAsset) Dependencies() terra.Dependencies {
	return da.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DataplexAsset].
func (da *DataplexAsset) LifecycleManagement() *terra.Lifecycle {
	return da.Lifecycle
}

// Attributes returns the attributes for [DataplexAsset].
func (da *DataplexAsset) Attributes() dataplexAssetAttributes {
	return dataplexAssetAttributes{ref: terra.ReferenceResource(da)}
}

// ImportState imports the given attribute values into [DataplexAsset]'s state.
func (da *DataplexAsset) ImportState(av io.Reader) error {
	da.state = &dataplexAssetState{}
	if err := json.NewDecoder(av).Decode(da.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", da.Type(), da.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DataplexAsset] has state.
func (da *DataplexAsset) State() (*dataplexAssetState, bool) {
	return da.state, da.state != nil
}

// StateMust returns the state for [DataplexAsset]. Panics if the state is nil.
func (da *DataplexAsset) StateMust() *dataplexAssetState {
	if da.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", da.Type(), da.LocalName()))
	}
	return da.state
}

// DataplexAssetArgs contains the configurations for google_dataplex_asset.
type DataplexAssetArgs struct {
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
	// DiscoveryStatus: min=0
	DiscoveryStatus []dataplexasset.DiscoveryStatus `hcl:"discovery_status,block" validate:"min=0"`
	// ResourceStatus: min=0
	ResourceStatus []dataplexasset.ResourceStatus `hcl:"resource_status,block" validate:"min=0"`
	// SecurityStatus: min=0
	SecurityStatus []dataplexasset.SecurityStatus `hcl:"security_status,block" validate:"min=0"`
	// DiscoverySpec: required
	DiscoverySpec *dataplexasset.DiscoverySpec `hcl:"discovery_spec,block" validate:"required"`
	// ResourceSpec: required
	ResourceSpec *dataplexasset.ResourceSpec `hcl:"resource_spec,block" validate:"required"`
	// Timeouts: optional
	Timeouts *dataplexasset.Timeouts `hcl:"timeouts,block"`
}
type dataplexAssetAttributes struct {
	ref terra.Reference
}

// CreateTime returns a reference to field create_time of google_dataplex_asset.
func (da dataplexAssetAttributes) CreateTime() terra.StringValue {
	return terra.ReferenceAsString(da.ref.Append("create_time"))
}

// DataplexZone returns a reference to field dataplex_zone of google_dataplex_asset.
func (da dataplexAssetAttributes) DataplexZone() terra.StringValue {
	return terra.ReferenceAsString(da.ref.Append("dataplex_zone"))
}

// Description returns a reference to field description of google_dataplex_asset.
func (da dataplexAssetAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(da.ref.Append("description"))
}

// DisplayName returns a reference to field display_name of google_dataplex_asset.
func (da dataplexAssetAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(da.ref.Append("display_name"))
}

// Id returns a reference to field id of google_dataplex_asset.
func (da dataplexAssetAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(da.ref.Append("id"))
}

// Labels returns a reference to field labels of google_dataplex_asset.
func (da dataplexAssetAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](da.ref.Append("labels"))
}

// Lake returns a reference to field lake of google_dataplex_asset.
func (da dataplexAssetAttributes) Lake() terra.StringValue {
	return terra.ReferenceAsString(da.ref.Append("lake"))
}

// Location returns a reference to field location of google_dataplex_asset.
func (da dataplexAssetAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(da.ref.Append("location"))
}

// Name returns a reference to field name of google_dataplex_asset.
func (da dataplexAssetAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(da.ref.Append("name"))
}

// Project returns a reference to field project of google_dataplex_asset.
func (da dataplexAssetAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(da.ref.Append("project"))
}

// State returns a reference to field state of google_dataplex_asset.
func (da dataplexAssetAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(da.ref.Append("state"))
}

// Uid returns a reference to field uid of google_dataplex_asset.
func (da dataplexAssetAttributes) Uid() terra.StringValue {
	return terra.ReferenceAsString(da.ref.Append("uid"))
}

// UpdateTime returns a reference to field update_time of google_dataplex_asset.
func (da dataplexAssetAttributes) UpdateTime() terra.StringValue {
	return terra.ReferenceAsString(da.ref.Append("update_time"))
}

func (da dataplexAssetAttributes) DiscoveryStatus() terra.ListValue[dataplexasset.DiscoveryStatusAttributes] {
	return terra.ReferenceAsList[dataplexasset.DiscoveryStatusAttributes](da.ref.Append("discovery_status"))
}

func (da dataplexAssetAttributes) ResourceStatus() terra.ListValue[dataplexasset.ResourceStatusAttributes] {
	return terra.ReferenceAsList[dataplexasset.ResourceStatusAttributes](da.ref.Append("resource_status"))
}

func (da dataplexAssetAttributes) SecurityStatus() terra.ListValue[dataplexasset.SecurityStatusAttributes] {
	return terra.ReferenceAsList[dataplexasset.SecurityStatusAttributes](da.ref.Append("security_status"))
}

func (da dataplexAssetAttributes) DiscoverySpec() terra.ListValue[dataplexasset.DiscoverySpecAttributes] {
	return terra.ReferenceAsList[dataplexasset.DiscoverySpecAttributes](da.ref.Append("discovery_spec"))
}

func (da dataplexAssetAttributes) ResourceSpec() terra.ListValue[dataplexasset.ResourceSpecAttributes] {
	return terra.ReferenceAsList[dataplexasset.ResourceSpecAttributes](da.ref.Append("resource_spec"))
}

func (da dataplexAssetAttributes) Timeouts() dataplexasset.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dataplexasset.TimeoutsAttributes](da.ref.Append("timeouts"))
}

type dataplexAssetState struct {
	CreateTime      string                               `json:"create_time"`
	DataplexZone    string                               `json:"dataplex_zone"`
	Description     string                               `json:"description"`
	DisplayName     string                               `json:"display_name"`
	Id              string                               `json:"id"`
	Labels          map[string]string                    `json:"labels"`
	Lake            string                               `json:"lake"`
	Location        string                               `json:"location"`
	Name            string                               `json:"name"`
	Project         string                               `json:"project"`
	State           string                               `json:"state"`
	Uid             string                               `json:"uid"`
	UpdateTime      string                               `json:"update_time"`
	DiscoveryStatus []dataplexasset.DiscoveryStatusState `json:"discovery_status"`
	ResourceStatus  []dataplexasset.ResourceStatusState  `json:"resource_status"`
	SecurityStatus  []dataplexasset.SecurityStatusState  `json:"security_status"`
	DiscoverySpec   []dataplexasset.DiscoverySpecState   `json:"discovery_spec"`
	ResourceSpec    []dataplexasset.ResourceSpecState    `json:"resource_spec"`
	Timeouts        *dataplexasset.TimeoutsState         `json:"timeouts"`
}
