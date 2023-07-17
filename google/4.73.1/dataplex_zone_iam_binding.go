// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	dataplexzoneiambinding "github.com/golingon/terraproviders/google/4.73.1/dataplexzoneiambinding"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDataplexZoneIamBinding creates a new instance of [DataplexZoneIamBinding].
func NewDataplexZoneIamBinding(name string, args DataplexZoneIamBindingArgs) *DataplexZoneIamBinding {
	return &DataplexZoneIamBinding{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DataplexZoneIamBinding)(nil)

// DataplexZoneIamBinding represents the Terraform resource google_dataplex_zone_iam_binding.
type DataplexZoneIamBinding struct {
	Name      string
	Args      DataplexZoneIamBindingArgs
	state     *dataplexZoneIamBindingState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DataplexZoneIamBinding].
func (dzib *DataplexZoneIamBinding) Type() string {
	return "google_dataplex_zone_iam_binding"
}

// LocalName returns the local name for [DataplexZoneIamBinding].
func (dzib *DataplexZoneIamBinding) LocalName() string {
	return dzib.Name
}

// Configuration returns the configuration (args) for [DataplexZoneIamBinding].
func (dzib *DataplexZoneIamBinding) Configuration() interface{} {
	return dzib.Args
}

// DependOn is used for other resources to depend on [DataplexZoneIamBinding].
func (dzib *DataplexZoneIamBinding) DependOn() terra.Reference {
	return terra.ReferenceResource(dzib)
}

// Dependencies returns the list of resources [DataplexZoneIamBinding] depends_on.
func (dzib *DataplexZoneIamBinding) Dependencies() terra.Dependencies {
	return dzib.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DataplexZoneIamBinding].
func (dzib *DataplexZoneIamBinding) LifecycleManagement() *terra.Lifecycle {
	return dzib.Lifecycle
}

// Attributes returns the attributes for [DataplexZoneIamBinding].
func (dzib *DataplexZoneIamBinding) Attributes() dataplexZoneIamBindingAttributes {
	return dataplexZoneIamBindingAttributes{ref: terra.ReferenceResource(dzib)}
}

// ImportState imports the given attribute values into [DataplexZoneIamBinding]'s state.
func (dzib *DataplexZoneIamBinding) ImportState(av io.Reader) error {
	dzib.state = &dataplexZoneIamBindingState{}
	if err := json.NewDecoder(av).Decode(dzib.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dzib.Type(), dzib.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DataplexZoneIamBinding] has state.
func (dzib *DataplexZoneIamBinding) State() (*dataplexZoneIamBindingState, bool) {
	return dzib.state, dzib.state != nil
}

// StateMust returns the state for [DataplexZoneIamBinding]. Panics if the state is nil.
func (dzib *DataplexZoneIamBinding) StateMust() *dataplexZoneIamBindingState {
	if dzib.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dzib.Type(), dzib.LocalName()))
	}
	return dzib.state
}

// DataplexZoneIamBindingArgs contains the configurations for google_dataplex_zone_iam_binding.
type DataplexZoneIamBindingArgs struct {
	// DataplexZone: string, required
	DataplexZone terra.StringValue `hcl:"dataplex_zone,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Lake: string, required
	Lake terra.StringValue `hcl:"lake,attr" validate:"required"`
	// Location: string, optional
	Location terra.StringValue `hcl:"location,attr"`
	// Members: set of string, required
	Members terra.SetValue[terra.StringValue] `hcl:"members,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// Condition: optional
	Condition *dataplexzoneiambinding.Condition `hcl:"condition,block"`
}
type dataplexZoneIamBindingAttributes struct {
	ref terra.Reference
}

// DataplexZone returns a reference to field dataplex_zone of google_dataplex_zone_iam_binding.
func (dzib dataplexZoneIamBindingAttributes) DataplexZone() terra.StringValue {
	return terra.ReferenceAsString(dzib.ref.Append("dataplex_zone"))
}

// Etag returns a reference to field etag of google_dataplex_zone_iam_binding.
func (dzib dataplexZoneIamBindingAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(dzib.ref.Append("etag"))
}

// Id returns a reference to field id of google_dataplex_zone_iam_binding.
func (dzib dataplexZoneIamBindingAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dzib.ref.Append("id"))
}

// Lake returns a reference to field lake of google_dataplex_zone_iam_binding.
func (dzib dataplexZoneIamBindingAttributes) Lake() terra.StringValue {
	return terra.ReferenceAsString(dzib.ref.Append("lake"))
}

// Location returns a reference to field location of google_dataplex_zone_iam_binding.
func (dzib dataplexZoneIamBindingAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(dzib.ref.Append("location"))
}

// Members returns a reference to field members of google_dataplex_zone_iam_binding.
func (dzib dataplexZoneIamBindingAttributes) Members() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](dzib.ref.Append("members"))
}

// Project returns a reference to field project of google_dataplex_zone_iam_binding.
func (dzib dataplexZoneIamBindingAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(dzib.ref.Append("project"))
}

// Role returns a reference to field role of google_dataplex_zone_iam_binding.
func (dzib dataplexZoneIamBindingAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(dzib.ref.Append("role"))
}

func (dzib dataplexZoneIamBindingAttributes) Condition() terra.ListValue[dataplexzoneiambinding.ConditionAttributes] {
	return terra.ReferenceAsList[dataplexzoneiambinding.ConditionAttributes](dzib.ref.Append("condition"))
}

type dataplexZoneIamBindingState struct {
	DataplexZone string                                  `json:"dataplex_zone"`
	Etag         string                                  `json:"etag"`
	Id           string                                  `json:"id"`
	Lake         string                                  `json:"lake"`
	Location     string                                  `json:"location"`
	Members      []string                                `json:"members"`
	Project      string                                  `json:"project"`
	Role         string                                  `json:"role"`
	Condition    []dataplexzoneiambinding.ConditionState `json:"condition"`
}
