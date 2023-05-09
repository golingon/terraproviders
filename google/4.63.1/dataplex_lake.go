// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	dataplexlake "github.com/golingon/terraproviders/google/4.63.1/dataplexlake"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDataplexLake creates a new instance of [DataplexLake].
func NewDataplexLake(name string, args DataplexLakeArgs) *DataplexLake {
	return &DataplexLake{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DataplexLake)(nil)

// DataplexLake represents the Terraform resource google_dataplex_lake.
type DataplexLake struct {
	Name      string
	Args      DataplexLakeArgs
	state     *dataplexLakeState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DataplexLake].
func (dl *DataplexLake) Type() string {
	return "google_dataplex_lake"
}

// LocalName returns the local name for [DataplexLake].
func (dl *DataplexLake) LocalName() string {
	return dl.Name
}

// Configuration returns the configuration (args) for [DataplexLake].
func (dl *DataplexLake) Configuration() interface{} {
	return dl.Args
}

// DependOn is used for other resources to depend on [DataplexLake].
func (dl *DataplexLake) DependOn() terra.Reference {
	return terra.ReferenceResource(dl)
}

// Dependencies returns the list of resources [DataplexLake] depends_on.
func (dl *DataplexLake) Dependencies() terra.Dependencies {
	return dl.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DataplexLake].
func (dl *DataplexLake) LifecycleManagement() *terra.Lifecycle {
	return dl.Lifecycle
}

// Attributes returns the attributes for [DataplexLake].
func (dl *DataplexLake) Attributes() dataplexLakeAttributes {
	return dataplexLakeAttributes{ref: terra.ReferenceResource(dl)}
}

// ImportState imports the given attribute values into [DataplexLake]'s state.
func (dl *DataplexLake) ImportState(av io.Reader) error {
	dl.state = &dataplexLakeState{}
	if err := json.NewDecoder(av).Decode(dl.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dl.Type(), dl.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DataplexLake] has state.
func (dl *DataplexLake) State() (*dataplexLakeState, bool) {
	return dl.state, dl.state != nil
}

// StateMust returns the state for [DataplexLake]. Panics if the state is nil.
func (dl *DataplexLake) StateMust() *dataplexLakeState {
	if dl.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dl.Type(), dl.LocalName()))
	}
	return dl.state
}

// DataplexLakeArgs contains the configurations for google_dataplex_lake.
type DataplexLakeArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// DisplayName: string, optional
	DisplayName terra.StringValue `hcl:"display_name,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Labels: map of string, optional
	Labels terra.MapValue[terra.StringValue] `hcl:"labels,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// AssetStatus: min=0
	AssetStatus []dataplexlake.AssetStatus `hcl:"asset_status,block" validate:"min=0"`
	// MetastoreStatus: min=0
	MetastoreStatus []dataplexlake.MetastoreStatus `hcl:"metastore_status,block" validate:"min=0"`
	// Metastore: optional
	Metastore *dataplexlake.Metastore `hcl:"metastore,block"`
	// Timeouts: optional
	Timeouts *dataplexlake.Timeouts `hcl:"timeouts,block"`
}
type dataplexLakeAttributes struct {
	ref terra.Reference
}

// CreateTime returns a reference to field create_time of google_dataplex_lake.
func (dl dataplexLakeAttributes) CreateTime() terra.StringValue {
	return terra.ReferenceAsString(dl.ref.Append("create_time"))
}

// Description returns a reference to field description of google_dataplex_lake.
func (dl dataplexLakeAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(dl.ref.Append("description"))
}

// DisplayName returns a reference to field display_name of google_dataplex_lake.
func (dl dataplexLakeAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(dl.ref.Append("display_name"))
}

// Id returns a reference to field id of google_dataplex_lake.
func (dl dataplexLakeAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dl.ref.Append("id"))
}

// Labels returns a reference to field labels of google_dataplex_lake.
func (dl dataplexLakeAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dl.ref.Append("labels"))
}

// Location returns a reference to field location of google_dataplex_lake.
func (dl dataplexLakeAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(dl.ref.Append("location"))
}

// Name returns a reference to field name of google_dataplex_lake.
func (dl dataplexLakeAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(dl.ref.Append("name"))
}

// Project returns a reference to field project of google_dataplex_lake.
func (dl dataplexLakeAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(dl.ref.Append("project"))
}

// ServiceAccount returns a reference to field service_account of google_dataplex_lake.
func (dl dataplexLakeAttributes) ServiceAccount() terra.StringValue {
	return terra.ReferenceAsString(dl.ref.Append("service_account"))
}

// State returns a reference to field state of google_dataplex_lake.
func (dl dataplexLakeAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(dl.ref.Append("state"))
}

// Uid returns a reference to field uid of google_dataplex_lake.
func (dl dataplexLakeAttributes) Uid() terra.StringValue {
	return terra.ReferenceAsString(dl.ref.Append("uid"))
}

// UpdateTime returns a reference to field update_time of google_dataplex_lake.
func (dl dataplexLakeAttributes) UpdateTime() terra.StringValue {
	return terra.ReferenceAsString(dl.ref.Append("update_time"))
}

func (dl dataplexLakeAttributes) AssetStatus() terra.ListValue[dataplexlake.AssetStatusAttributes] {
	return terra.ReferenceAsList[dataplexlake.AssetStatusAttributes](dl.ref.Append("asset_status"))
}

func (dl dataplexLakeAttributes) MetastoreStatus() terra.ListValue[dataplexlake.MetastoreStatusAttributes] {
	return terra.ReferenceAsList[dataplexlake.MetastoreStatusAttributes](dl.ref.Append("metastore_status"))
}

func (dl dataplexLakeAttributes) Metastore() terra.ListValue[dataplexlake.MetastoreAttributes] {
	return terra.ReferenceAsList[dataplexlake.MetastoreAttributes](dl.ref.Append("metastore"))
}

func (dl dataplexLakeAttributes) Timeouts() dataplexlake.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dataplexlake.TimeoutsAttributes](dl.ref.Append("timeouts"))
}

type dataplexLakeState struct {
	CreateTime      string                              `json:"create_time"`
	Description     string                              `json:"description"`
	DisplayName     string                              `json:"display_name"`
	Id              string                              `json:"id"`
	Labels          map[string]string                   `json:"labels"`
	Location        string                              `json:"location"`
	Name            string                              `json:"name"`
	Project         string                              `json:"project"`
	ServiceAccount  string                              `json:"service_account"`
	State           string                              `json:"state"`
	Uid             string                              `json:"uid"`
	UpdateTime      string                              `json:"update_time"`
	AssetStatus     []dataplexlake.AssetStatusState     `json:"asset_status"`
	MetastoreStatus []dataplexlake.MetastoreStatusState `json:"metastore_status"`
	Metastore       []dataplexlake.MetastoreState       `json:"metastore"`
	Timeouts        *dataplexlake.TimeoutsState         `json:"timeouts"`
}
