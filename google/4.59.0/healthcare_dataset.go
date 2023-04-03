// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	healthcaredataset "github.com/golingon/terraproviders/google/4.59.0/healthcaredataset"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewHealthcareDataset creates a new instance of [HealthcareDataset].
func NewHealthcareDataset(name string, args HealthcareDatasetArgs) *HealthcareDataset {
	return &HealthcareDataset{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*HealthcareDataset)(nil)

// HealthcareDataset represents the Terraform resource google_healthcare_dataset.
type HealthcareDataset struct {
	Name      string
	Args      HealthcareDatasetArgs
	state     *healthcareDatasetState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [HealthcareDataset].
func (hd *HealthcareDataset) Type() string {
	return "google_healthcare_dataset"
}

// LocalName returns the local name for [HealthcareDataset].
func (hd *HealthcareDataset) LocalName() string {
	return hd.Name
}

// Configuration returns the configuration (args) for [HealthcareDataset].
func (hd *HealthcareDataset) Configuration() interface{} {
	return hd.Args
}

// DependOn is used for other resources to depend on [HealthcareDataset].
func (hd *HealthcareDataset) DependOn() terra.Reference {
	return terra.ReferenceResource(hd)
}

// Dependencies returns the list of resources [HealthcareDataset] depends_on.
func (hd *HealthcareDataset) Dependencies() terra.Dependencies {
	return hd.DependsOn
}

// LifecycleManagement returns the lifecycle block for [HealthcareDataset].
func (hd *HealthcareDataset) LifecycleManagement() *terra.Lifecycle {
	return hd.Lifecycle
}

// Attributes returns the attributes for [HealthcareDataset].
func (hd *HealthcareDataset) Attributes() healthcareDatasetAttributes {
	return healthcareDatasetAttributes{ref: terra.ReferenceResource(hd)}
}

// ImportState imports the given attribute values into [HealthcareDataset]'s state.
func (hd *HealthcareDataset) ImportState(av io.Reader) error {
	hd.state = &healthcareDatasetState{}
	if err := json.NewDecoder(av).Decode(hd.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", hd.Type(), hd.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [HealthcareDataset] has state.
func (hd *HealthcareDataset) State() (*healthcareDatasetState, bool) {
	return hd.state, hd.state != nil
}

// StateMust returns the state for [HealthcareDataset]. Panics if the state is nil.
func (hd *HealthcareDataset) StateMust() *healthcareDatasetState {
	if hd.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", hd.Type(), hd.LocalName()))
	}
	return hd.state
}

// HealthcareDatasetArgs contains the configurations for google_healthcare_dataset.
type HealthcareDatasetArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// TimeZone: string, optional
	TimeZone terra.StringValue `hcl:"time_zone,attr"`
	// Timeouts: optional
	Timeouts *healthcaredataset.Timeouts `hcl:"timeouts,block"`
}
type healthcareDatasetAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of google_healthcare_dataset.
func (hd healthcareDatasetAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(hd.ref.Append("id"))
}

// Location returns a reference to field location of google_healthcare_dataset.
func (hd healthcareDatasetAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(hd.ref.Append("location"))
}

// Name returns a reference to field name of google_healthcare_dataset.
func (hd healthcareDatasetAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(hd.ref.Append("name"))
}

// Project returns a reference to field project of google_healthcare_dataset.
func (hd healthcareDatasetAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(hd.ref.Append("project"))
}

// SelfLink returns a reference to field self_link of google_healthcare_dataset.
func (hd healthcareDatasetAttributes) SelfLink() terra.StringValue {
	return terra.ReferenceAsString(hd.ref.Append("self_link"))
}

// TimeZone returns a reference to field time_zone of google_healthcare_dataset.
func (hd healthcareDatasetAttributes) TimeZone() terra.StringValue {
	return terra.ReferenceAsString(hd.ref.Append("time_zone"))
}

func (hd healthcareDatasetAttributes) Timeouts() healthcaredataset.TimeoutsAttributes {
	return terra.ReferenceAsSingle[healthcaredataset.TimeoutsAttributes](hd.ref.Append("timeouts"))
}

type healthcareDatasetState struct {
	Id       string                           `json:"id"`
	Location string                           `json:"location"`
	Name     string                           `json:"name"`
	Project  string                           `json:"project"`
	SelfLink string                           `json:"self_link"`
	TimeZone string                           `json:"time_zone"`
	Timeouts *healthcaredataset.TimeoutsState `json:"timeouts"`
}
