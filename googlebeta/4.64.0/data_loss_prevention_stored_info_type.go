// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	datalosspreventionstoredinfotype "github.com/golingon/terraproviders/googlebeta/4.64.0/datalosspreventionstoredinfotype"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDataLossPreventionStoredInfoType creates a new instance of [DataLossPreventionStoredInfoType].
func NewDataLossPreventionStoredInfoType(name string, args DataLossPreventionStoredInfoTypeArgs) *DataLossPreventionStoredInfoType {
	return &DataLossPreventionStoredInfoType{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DataLossPreventionStoredInfoType)(nil)

// DataLossPreventionStoredInfoType represents the Terraform resource google_data_loss_prevention_stored_info_type.
type DataLossPreventionStoredInfoType struct {
	Name      string
	Args      DataLossPreventionStoredInfoTypeArgs
	state     *dataLossPreventionStoredInfoTypeState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DataLossPreventionStoredInfoType].
func (dlpsit *DataLossPreventionStoredInfoType) Type() string {
	return "google_data_loss_prevention_stored_info_type"
}

// LocalName returns the local name for [DataLossPreventionStoredInfoType].
func (dlpsit *DataLossPreventionStoredInfoType) LocalName() string {
	return dlpsit.Name
}

// Configuration returns the configuration (args) for [DataLossPreventionStoredInfoType].
func (dlpsit *DataLossPreventionStoredInfoType) Configuration() interface{} {
	return dlpsit.Args
}

// DependOn is used for other resources to depend on [DataLossPreventionStoredInfoType].
func (dlpsit *DataLossPreventionStoredInfoType) DependOn() terra.Reference {
	return terra.ReferenceResource(dlpsit)
}

// Dependencies returns the list of resources [DataLossPreventionStoredInfoType] depends_on.
func (dlpsit *DataLossPreventionStoredInfoType) Dependencies() terra.Dependencies {
	return dlpsit.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DataLossPreventionStoredInfoType].
func (dlpsit *DataLossPreventionStoredInfoType) LifecycleManagement() *terra.Lifecycle {
	return dlpsit.Lifecycle
}

// Attributes returns the attributes for [DataLossPreventionStoredInfoType].
func (dlpsit *DataLossPreventionStoredInfoType) Attributes() dataLossPreventionStoredInfoTypeAttributes {
	return dataLossPreventionStoredInfoTypeAttributes{ref: terra.ReferenceResource(dlpsit)}
}

// ImportState imports the given attribute values into [DataLossPreventionStoredInfoType]'s state.
func (dlpsit *DataLossPreventionStoredInfoType) ImportState(av io.Reader) error {
	dlpsit.state = &dataLossPreventionStoredInfoTypeState{}
	if err := json.NewDecoder(av).Decode(dlpsit.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dlpsit.Type(), dlpsit.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DataLossPreventionStoredInfoType] has state.
func (dlpsit *DataLossPreventionStoredInfoType) State() (*dataLossPreventionStoredInfoTypeState, bool) {
	return dlpsit.state, dlpsit.state != nil
}

// StateMust returns the state for [DataLossPreventionStoredInfoType]. Panics if the state is nil.
func (dlpsit *DataLossPreventionStoredInfoType) StateMust() *dataLossPreventionStoredInfoTypeState {
	if dlpsit.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dlpsit.Type(), dlpsit.LocalName()))
	}
	return dlpsit.state
}

// DataLossPreventionStoredInfoTypeArgs contains the configurations for google_data_loss_prevention_stored_info_type.
type DataLossPreventionStoredInfoTypeArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// DisplayName: string, optional
	DisplayName terra.StringValue `hcl:"display_name,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Parent: string, required
	Parent terra.StringValue `hcl:"parent,attr" validate:"required"`
	// Dictionary: optional
	Dictionary *datalosspreventionstoredinfotype.Dictionary `hcl:"dictionary,block"`
	// LargeCustomDictionary: optional
	LargeCustomDictionary *datalosspreventionstoredinfotype.LargeCustomDictionary `hcl:"large_custom_dictionary,block"`
	// Regex: optional
	Regex *datalosspreventionstoredinfotype.Regex `hcl:"regex,block"`
	// Timeouts: optional
	Timeouts *datalosspreventionstoredinfotype.Timeouts `hcl:"timeouts,block"`
}
type dataLossPreventionStoredInfoTypeAttributes struct {
	ref terra.Reference
}

// Description returns a reference to field description of google_data_loss_prevention_stored_info_type.
func (dlpsit dataLossPreventionStoredInfoTypeAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(dlpsit.ref.Append("description"))
}

// DisplayName returns a reference to field display_name of google_data_loss_prevention_stored_info_type.
func (dlpsit dataLossPreventionStoredInfoTypeAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(dlpsit.ref.Append("display_name"))
}

// Id returns a reference to field id of google_data_loss_prevention_stored_info_type.
func (dlpsit dataLossPreventionStoredInfoTypeAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dlpsit.ref.Append("id"))
}

// Name returns a reference to field name of google_data_loss_prevention_stored_info_type.
func (dlpsit dataLossPreventionStoredInfoTypeAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(dlpsit.ref.Append("name"))
}

// Parent returns a reference to field parent of google_data_loss_prevention_stored_info_type.
func (dlpsit dataLossPreventionStoredInfoTypeAttributes) Parent() terra.StringValue {
	return terra.ReferenceAsString(dlpsit.ref.Append("parent"))
}

func (dlpsit dataLossPreventionStoredInfoTypeAttributes) Dictionary() terra.ListValue[datalosspreventionstoredinfotype.DictionaryAttributes] {
	return terra.ReferenceAsList[datalosspreventionstoredinfotype.DictionaryAttributes](dlpsit.ref.Append("dictionary"))
}

func (dlpsit dataLossPreventionStoredInfoTypeAttributes) LargeCustomDictionary() terra.ListValue[datalosspreventionstoredinfotype.LargeCustomDictionaryAttributes] {
	return terra.ReferenceAsList[datalosspreventionstoredinfotype.LargeCustomDictionaryAttributes](dlpsit.ref.Append("large_custom_dictionary"))
}

func (dlpsit dataLossPreventionStoredInfoTypeAttributes) Regex() terra.ListValue[datalosspreventionstoredinfotype.RegexAttributes] {
	return terra.ReferenceAsList[datalosspreventionstoredinfotype.RegexAttributes](dlpsit.ref.Append("regex"))
}

func (dlpsit dataLossPreventionStoredInfoTypeAttributes) Timeouts() datalosspreventionstoredinfotype.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datalosspreventionstoredinfotype.TimeoutsAttributes](dlpsit.ref.Append("timeouts"))
}

type dataLossPreventionStoredInfoTypeState struct {
	Description           string                                                        `json:"description"`
	DisplayName           string                                                        `json:"display_name"`
	Id                    string                                                        `json:"id"`
	Name                  string                                                        `json:"name"`
	Parent                string                                                        `json:"parent"`
	Dictionary            []datalosspreventionstoredinfotype.DictionaryState            `json:"dictionary"`
	LargeCustomDictionary []datalosspreventionstoredinfotype.LargeCustomDictionaryState `json:"large_custom_dictionary"`
	Regex                 []datalosspreventionstoredinfotype.RegexState                 `json:"regex"`
	Timeouts              *datalosspreventionstoredinfotype.TimeoutsState               `json:"timeouts"`
}
