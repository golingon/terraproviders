// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	dataplexdatascaniambinding "github.com/golingon/terraproviders/googlebeta/4.71.0/dataplexdatascaniambinding"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDataplexDatascanIamBinding creates a new instance of [DataplexDatascanIamBinding].
func NewDataplexDatascanIamBinding(name string, args DataplexDatascanIamBindingArgs) *DataplexDatascanIamBinding {
	return &DataplexDatascanIamBinding{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DataplexDatascanIamBinding)(nil)

// DataplexDatascanIamBinding represents the Terraform resource google_dataplex_datascan_iam_binding.
type DataplexDatascanIamBinding struct {
	Name      string
	Args      DataplexDatascanIamBindingArgs
	state     *dataplexDatascanIamBindingState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DataplexDatascanIamBinding].
func (ddib *DataplexDatascanIamBinding) Type() string {
	return "google_dataplex_datascan_iam_binding"
}

// LocalName returns the local name for [DataplexDatascanIamBinding].
func (ddib *DataplexDatascanIamBinding) LocalName() string {
	return ddib.Name
}

// Configuration returns the configuration (args) for [DataplexDatascanIamBinding].
func (ddib *DataplexDatascanIamBinding) Configuration() interface{} {
	return ddib.Args
}

// DependOn is used for other resources to depend on [DataplexDatascanIamBinding].
func (ddib *DataplexDatascanIamBinding) DependOn() terra.Reference {
	return terra.ReferenceResource(ddib)
}

// Dependencies returns the list of resources [DataplexDatascanIamBinding] depends_on.
func (ddib *DataplexDatascanIamBinding) Dependencies() terra.Dependencies {
	return ddib.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DataplexDatascanIamBinding].
func (ddib *DataplexDatascanIamBinding) LifecycleManagement() *terra.Lifecycle {
	return ddib.Lifecycle
}

// Attributes returns the attributes for [DataplexDatascanIamBinding].
func (ddib *DataplexDatascanIamBinding) Attributes() dataplexDatascanIamBindingAttributes {
	return dataplexDatascanIamBindingAttributes{ref: terra.ReferenceResource(ddib)}
}

// ImportState imports the given attribute values into [DataplexDatascanIamBinding]'s state.
func (ddib *DataplexDatascanIamBinding) ImportState(av io.Reader) error {
	ddib.state = &dataplexDatascanIamBindingState{}
	if err := json.NewDecoder(av).Decode(ddib.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ddib.Type(), ddib.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DataplexDatascanIamBinding] has state.
func (ddib *DataplexDatascanIamBinding) State() (*dataplexDatascanIamBindingState, bool) {
	return ddib.state, ddib.state != nil
}

// StateMust returns the state for [DataplexDatascanIamBinding]. Panics if the state is nil.
func (ddib *DataplexDatascanIamBinding) StateMust() *dataplexDatascanIamBindingState {
	if ddib.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ddib.Type(), ddib.LocalName()))
	}
	return ddib.state
}

// DataplexDatascanIamBindingArgs contains the configurations for google_dataplex_datascan_iam_binding.
type DataplexDatascanIamBindingArgs struct {
	// DataScanId: string, required
	DataScanId terra.StringValue `hcl:"data_scan_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, optional
	Location terra.StringValue `hcl:"location,attr"`
	// Members: set of string, required
	Members terra.SetValue[terra.StringValue] `hcl:"members,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// Condition: optional
	Condition *dataplexdatascaniambinding.Condition `hcl:"condition,block"`
}
type dataplexDatascanIamBindingAttributes struct {
	ref terra.Reference
}

// DataScanId returns a reference to field data_scan_id of google_dataplex_datascan_iam_binding.
func (ddib dataplexDatascanIamBindingAttributes) DataScanId() terra.StringValue {
	return terra.ReferenceAsString(ddib.ref.Append("data_scan_id"))
}

// Etag returns a reference to field etag of google_dataplex_datascan_iam_binding.
func (ddib dataplexDatascanIamBindingAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(ddib.ref.Append("etag"))
}

// Id returns a reference to field id of google_dataplex_datascan_iam_binding.
func (ddib dataplexDatascanIamBindingAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ddib.ref.Append("id"))
}

// Location returns a reference to field location of google_dataplex_datascan_iam_binding.
func (ddib dataplexDatascanIamBindingAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(ddib.ref.Append("location"))
}

// Members returns a reference to field members of google_dataplex_datascan_iam_binding.
func (ddib dataplexDatascanIamBindingAttributes) Members() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](ddib.ref.Append("members"))
}

// Project returns a reference to field project of google_dataplex_datascan_iam_binding.
func (ddib dataplexDatascanIamBindingAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(ddib.ref.Append("project"))
}

// Role returns a reference to field role of google_dataplex_datascan_iam_binding.
func (ddib dataplexDatascanIamBindingAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(ddib.ref.Append("role"))
}

func (ddib dataplexDatascanIamBindingAttributes) Condition() terra.ListValue[dataplexdatascaniambinding.ConditionAttributes] {
	return terra.ReferenceAsList[dataplexdatascaniambinding.ConditionAttributes](ddib.ref.Append("condition"))
}

type dataplexDatascanIamBindingState struct {
	DataScanId string                                      `json:"data_scan_id"`
	Etag       string                                      `json:"etag"`
	Id         string                                      `json:"id"`
	Location   string                                      `json:"location"`
	Members    []string                                    `json:"members"`
	Project    string                                      `json:"project"`
	Role       string                                      `json:"role"`
	Condition  []dataplexdatascaniambinding.ConditionState `json:"condition"`
}
