// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	dataplexzone "github.com/golingon/terraproviders/google/4.76.0/dataplexzone"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDataplexZone creates a new instance of [DataplexZone].
func NewDataplexZone(name string, args DataplexZoneArgs) *DataplexZone {
	return &DataplexZone{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DataplexZone)(nil)

// DataplexZone represents the Terraform resource google_dataplex_zone.
type DataplexZone struct {
	Name      string
	Args      DataplexZoneArgs
	state     *dataplexZoneState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DataplexZone].
func (dz *DataplexZone) Type() string {
	return "google_dataplex_zone"
}

// LocalName returns the local name for [DataplexZone].
func (dz *DataplexZone) LocalName() string {
	return dz.Name
}

// Configuration returns the configuration (args) for [DataplexZone].
func (dz *DataplexZone) Configuration() interface{} {
	return dz.Args
}

// DependOn is used for other resources to depend on [DataplexZone].
func (dz *DataplexZone) DependOn() terra.Reference {
	return terra.ReferenceResource(dz)
}

// Dependencies returns the list of resources [DataplexZone] depends_on.
func (dz *DataplexZone) Dependencies() terra.Dependencies {
	return dz.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DataplexZone].
func (dz *DataplexZone) LifecycleManagement() *terra.Lifecycle {
	return dz.Lifecycle
}

// Attributes returns the attributes for [DataplexZone].
func (dz *DataplexZone) Attributes() dataplexZoneAttributes {
	return dataplexZoneAttributes{ref: terra.ReferenceResource(dz)}
}

// ImportState imports the given attribute values into [DataplexZone]'s state.
func (dz *DataplexZone) ImportState(av io.Reader) error {
	dz.state = &dataplexZoneState{}
	if err := json.NewDecoder(av).Decode(dz.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dz.Type(), dz.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DataplexZone] has state.
func (dz *DataplexZone) State() (*dataplexZoneState, bool) {
	return dz.state, dz.state != nil
}

// StateMust returns the state for [DataplexZone]. Panics if the state is nil.
func (dz *DataplexZone) StateMust() *dataplexZoneState {
	if dz.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dz.Type(), dz.LocalName()))
	}
	return dz.state
}

// DataplexZoneArgs contains the configurations for google_dataplex_zone.
type DataplexZoneArgs struct {
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
	// Type: string, required
	Type terra.StringValue `hcl:"type,attr" validate:"required"`
	// AssetStatus: min=0
	AssetStatus []dataplexzone.AssetStatus `hcl:"asset_status,block" validate:"min=0"`
	// DiscoverySpec: required
	DiscoverySpec *dataplexzone.DiscoverySpec `hcl:"discovery_spec,block" validate:"required"`
	// ResourceSpec: required
	ResourceSpec *dataplexzone.ResourceSpec `hcl:"resource_spec,block" validate:"required"`
	// Timeouts: optional
	Timeouts *dataplexzone.Timeouts `hcl:"timeouts,block"`
}
type dataplexZoneAttributes struct {
	ref terra.Reference
}

// CreateTime returns a reference to field create_time of google_dataplex_zone.
func (dz dataplexZoneAttributes) CreateTime() terra.StringValue {
	return terra.ReferenceAsString(dz.ref.Append("create_time"))
}

// Description returns a reference to field description of google_dataplex_zone.
func (dz dataplexZoneAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(dz.ref.Append("description"))
}

// DisplayName returns a reference to field display_name of google_dataplex_zone.
func (dz dataplexZoneAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(dz.ref.Append("display_name"))
}

// Id returns a reference to field id of google_dataplex_zone.
func (dz dataplexZoneAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dz.ref.Append("id"))
}

// Labels returns a reference to field labels of google_dataplex_zone.
func (dz dataplexZoneAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dz.ref.Append("labels"))
}

// Lake returns a reference to field lake of google_dataplex_zone.
func (dz dataplexZoneAttributes) Lake() terra.StringValue {
	return terra.ReferenceAsString(dz.ref.Append("lake"))
}

// Location returns a reference to field location of google_dataplex_zone.
func (dz dataplexZoneAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(dz.ref.Append("location"))
}

// Name returns a reference to field name of google_dataplex_zone.
func (dz dataplexZoneAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(dz.ref.Append("name"))
}

// Project returns a reference to field project of google_dataplex_zone.
func (dz dataplexZoneAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(dz.ref.Append("project"))
}

// State returns a reference to field state of google_dataplex_zone.
func (dz dataplexZoneAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(dz.ref.Append("state"))
}

// Type returns a reference to field type of google_dataplex_zone.
func (dz dataplexZoneAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(dz.ref.Append("type"))
}

// Uid returns a reference to field uid of google_dataplex_zone.
func (dz dataplexZoneAttributes) Uid() terra.StringValue {
	return terra.ReferenceAsString(dz.ref.Append("uid"))
}

// UpdateTime returns a reference to field update_time of google_dataplex_zone.
func (dz dataplexZoneAttributes) UpdateTime() terra.StringValue {
	return terra.ReferenceAsString(dz.ref.Append("update_time"))
}

func (dz dataplexZoneAttributes) AssetStatus() terra.ListValue[dataplexzone.AssetStatusAttributes] {
	return terra.ReferenceAsList[dataplexzone.AssetStatusAttributes](dz.ref.Append("asset_status"))
}

func (dz dataplexZoneAttributes) DiscoverySpec() terra.ListValue[dataplexzone.DiscoverySpecAttributes] {
	return terra.ReferenceAsList[dataplexzone.DiscoverySpecAttributes](dz.ref.Append("discovery_spec"))
}

func (dz dataplexZoneAttributes) ResourceSpec() terra.ListValue[dataplexzone.ResourceSpecAttributes] {
	return terra.ReferenceAsList[dataplexzone.ResourceSpecAttributes](dz.ref.Append("resource_spec"))
}

func (dz dataplexZoneAttributes) Timeouts() dataplexzone.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dataplexzone.TimeoutsAttributes](dz.ref.Append("timeouts"))
}

type dataplexZoneState struct {
	CreateTime    string                            `json:"create_time"`
	Description   string                            `json:"description"`
	DisplayName   string                            `json:"display_name"`
	Id            string                            `json:"id"`
	Labels        map[string]string                 `json:"labels"`
	Lake          string                            `json:"lake"`
	Location      string                            `json:"location"`
	Name          string                            `json:"name"`
	Project       string                            `json:"project"`
	State         string                            `json:"state"`
	Type          string                            `json:"type"`
	Uid           string                            `json:"uid"`
	UpdateTime    string                            `json:"update_time"`
	AssetStatus   []dataplexzone.AssetStatusState   `json:"asset_status"`
	DiscoverySpec []dataplexzone.DiscoverySpecState `json:"discovery_spec"`
	ResourceSpec  []dataplexzone.ResourceSpecState  `json:"resource_spec"`
	Timeouts      *dataplexzone.TimeoutsState       `json:"timeouts"`
}
