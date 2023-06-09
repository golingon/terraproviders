// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	dataplexdatascan "github.com/golingon/terraproviders/googlebeta/4.72.1/dataplexdatascan"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDataplexDatascan creates a new instance of [DataplexDatascan].
func NewDataplexDatascan(name string, args DataplexDatascanArgs) *DataplexDatascan {
	return &DataplexDatascan{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DataplexDatascan)(nil)

// DataplexDatascan represents the Terraform resource google_dataplex_datascan.
type DataplexDatascan struct {
	Name      string
	Args      DataplexDatascanArgs
	state     *dataplexDatascanState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DataplexDatascan].
func (dd *DataplexDatascan) Type() string {
	return "google_dataplex_datascan"
}

// LocalName returns the local name for [DataplexDatascan].
func (dd *DataplexDatascan) LocalName() string {
	return dd.Name
}

// Configuration returns the configuration (args) for [DataplexDatascan].
func (dd *DataplexDatascan) Configuration() interface{} {
	return dd.Args
}

// DependOn is used for other resources to depend on [DataplexDatascan].
func (dd *DataplexDatascan) DependOn() terra.Reference {
	return terra.ReferenceResource(dd)
}

// Dependencies returns the list of resources [DataplexDatascan] depends_on.
func (dd *DataplexDatascan) Dependencies() terra.Dependencies {
	return dd.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DataplexDatascan].
func (dd *DataplexDatascan) LifecycleManagement() *terra.Lifecycle {
	return dd.Lifecycle
}

// Attributes returns the attributes for [DataplexDatascan].
func (dd *DataplexDatascan) Attributes() dataplexDatascanAttributes {
	return dataplexDatascanAttributes{ref: terra.ReferenceResource(dd)}
}

// ImportState imports the given attribute values into [DataplexDatascan]'s state.
func (dd *DataplexDatascan) ImportState(av io.Reader) error {
	dd.state = &dataplexDatascanState{}
	if err := json.NewDecoder(av).Decode(dd.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dd.Type(), dd.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DataplexDatascan] has state.
func (dd *DataplexDatascan) State() (*dataplexDatascanState, bool) {
	return dd.state, dd.state != nil
}

// StateMust returns the state for [DataplexDatascan]. Panics if the state is nil.
func (dd *DataplexDatascan) StateMust() *dataplexDatascanState {
	if dd.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dd.Type(), dd.LocalName()))
	}
	return dd.state
}

// DataplexDatascanArgs contains the configurations for google_dataplex_datascan.
type DataplexDatascanArgs struct {
	// DataScanId: string, required
	DataScanId terra.StringValue `hcl:"data_scan_id,attr" validate:"required"`
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
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// DataProfileResult: min=0
	DataProfileResult []dataplexdatascan.DataProfileResult `hcl:"data_profile_result,block" validate:"min=0"`
	// DataQualityResult: min=0
	DataQualityResult []dataplexdatascan.DataQualityResult `hcl:"data_quality_result,block" validate:"min=0"`
	// ExecutionStatus: min=0
	ExecutionStatus []dataplexdatascan.ExecutionStatus `hcl:"execution_status,block" validate:"min=0"`
	// Data: required
	Data *dataplexdatascan.Data `hcl:"data,block" validate:"required"`
	// DataProfileSpec: optional
	DataProfileSpec *dataplexdatascan.DataProfileSpec `hcl:"data_profile_spec,block"`
	// DataQualitySpec: optional
	DataQualitySpec *dataplexdatascan.DataQualitySpec `hcl:"data_quality_spec,block"`
	// ExecutionSpec: required
	ExecutionSpec *dataplexdatascan.ExecutionSpec `hcl:"execution_spec,block" validate:"required"`
	// Timeouts: optional
	Timeouts *dataplexdatascan.Timeouts `hcl:"timeouts,block"`
}
type dataplexDatascanAttributes struct {
	ref terra.Reference
}

// CreateTime returns a reference to field create_time of google_dataplex_datascan.
func (dd dataplexDatascanAttributes) CreateTime() terra.StringValue {
	return terra.ReferenceAsString(dd.ref.Append("create_time"))
}

// DataScanId returns a reference to field data_scan_id of google_dataplex_datascan.
func (dd dataplexDatascanAttributes) DataScanId() terra.StringValue {
	return terra.ReferenceAsString(dd.ref.Append("data_scan_id"))
}

// Description returns a reference to field description of google_dataplex_datascan.
func (dd dataplexDatascanAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(dd.ref.Append("description"))
}

// DisplayName returns a reference to field display_name of google_dataplex_datascan.
func (dd dataplexDatascanAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(dd.ref.Append("display_name"))
}

// Id returns a reference to field id of google_dataplex_datascan.
func (dd dataplexDatascanAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dd.ref.Append("id"))
}

// Labels returns a reference to field labels of google_dataplex_datascan.
func (dd dataplexDatascanAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dd.ref.Append("labels"))
}

// Location returns a reference to field location of google_dataplex_datascan.
func (dd dataplexDatascanAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(dd.ref.Append("location"))
}

// Name returns a reference to field name of google_dataplex_datascan.
func (dd dataplexDatascanAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(dd.ref.Append("name"))
}

// Project returns a reference to field project of google_dataplex_datascan.
func (dd dataplexDatascanAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(dd.ref.Append("project"))
}

// State returns a reference to field state of google_dataplex_datascan.
func (dd dataplexDatascanAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(dd.ref.Append("state"))
}

// Type returns a reference to field type of google_dataplex_datascan.
func (dd dataplexDatascanAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(dd.ref.Append("type"))
}

// Uid returns a reference to field uid of google_dataplex_datascan.
func (dd dataplexDatascanAttributes) Uid() terra.StringValue {
	return terra.ReferenceAsString(dd.ref.Append("uid"))
}

// UpdateTime returns a reference to field update_time of google_dataplex_datascan.
func (dd dataplexDatascanAttributes) UpdateTime() terra.StringValue {
	return terra.ReferenceAsString(dd.ref.Append("update_time"))
}

func (dd dataplexDatascanAttributes) DataProfileResult() terra.ListValue[dataplexdatascan.DataProfileResultAttributes] {
	return terra.ReferenceAsList[dataplexdatascan.DataProfileResultAttributes](dd.ref.Append("data_profile_result"))
}

func (dd dataplexDatascanAttributes) DataQualityResult() terra.ListValue[dataplexdatascan.DataQualityResultAttributes] {
	return terra.ReferenceAsList[dataplexdatascan.DataQualityResultAttributes](dd.ref.Append("data_quality_result"))
}

func (dd dataplexDatascanAttributes) ExecutionStatus() terra.ListValue[dataplexdatascan.ExecutionStatusAttributes] {
	return terra.ReferenceAsList[dataplexdatascan.ExecutionStatusAttributes](dd.ref.Append("execution_status"))
}

func (dd dataplexDatascanAttributes) Data() terra.ListValue[dataplexdatascan.DataAttributes] {
	return terra.ReferenceAsList[dataplexdatascan.DataAttributes](dd.ref.Append("data"))
}

func (dd dataplexDatascanAttributes) DataProfileSpec() terra.ListValue[dataplexdatascan.DataProfileSpecAttributes] {
	return terra.ReferenceAsList[dataplexdatascan.DataProfileSpecAttributes](dd.ref.Append("data_profile_spec"))
}

func (dd dataplexDatascanAttributes) DataQualitySpec() terra.ListValue[dataplexdatascan.DataQualitySpecAttributes] {
	return terra.ReferenceAsList[dataplexdatascan.DataQualitySpecAttributes](dd.ref.Append("data_quality_spec"))
}

func (dd dataplexDatascanAttributes) ExecutionSpec() terra.ListValue[dataplexdatascan.ExecutionSpecAttributes] {
	return terra.ReferenceAsList[dataplexdatascan.ExecutionSpecAttributes](dd.ref.Append("execution_spec"))
}

func (dd dataplexDatascanAttributes) Timeouts() dataplexdatascan.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dataplexdatascan.TimeoutsAttributes](dd.ref.Append("timeouts"))
}

type dataplexDatascanState struct {
	CreateTime        string                                    `json:"create_time"`
	DataScanId        string                                    `json:"data_scan_id"`
	Description       string                                    `json:"description"`
	DisplayName       string                                    `json:"display_name"`
	Id                string                                    `json:"id"`
	Labels            map[string]string                         `json:"labels"`
	Location          string                                    `json:"location"`
	Name              string                                    `json:"name"`
	Project           string                                    `json:"project"`
	State             string                                    `json:"state"`
	Type              string                                    `json:"type"`
	Uid               string                                    `json:"uid"`
	UpdateTime        string                                    `json:"update_time"`
	DataProfileResult []dataplexdatascan.DataProfileResultState `json:"data_profile_result"`
	DataQualityResult []dataplexdatascan.DataQualityResultState `json:"data_quality_result"`
	ExecutionStatus   []dataplexdatascan.ExecutionStatusState   `json:"execution_status"`
	Data              []dataplexdatascan.DataState              `json:"data"`
	DataProfileSpec   []dataplexdatascan.DataProfileSpecState   `json:"data_profile_spec"`
	DataQualitySpec   []dataplexdatascan.DataQualitySpecState   `json:"data_quality_spec"`
	ExecutionSpec     []dataplexdatascan.ExecutionSpecState     `json:"execution_spec"`
	Timeouts          *dataplexdatascan.TimeoutsState           `json:"timeouts"`
}
