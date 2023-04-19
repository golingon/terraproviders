// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	healthcaredicomstore "github.com/golingon/terraproviders/google/4.62.0/healthcaredicomstore"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewHealthcareDicomStore creates a new instance of [HealthcareDicomStore].
func NewHealthcareDicomStore(name string, args HealthcareDicomStoreArgs) *HealthcareDicomStore {
	return &HealthcareDicomStore{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*HealthcareDicomStore)(nil)

// HealthcareDicomStore represents the Terraform resource google_healthcare_dicom_store.
type HealthcareDicomStore struct {
	Name      string
	Args      HealthcareDicomStoreArgs
	state     *healthcareDicomStoreState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [HealthcareDicomStore].
func (hds *HealthcareDicomStore) Type() string {
	return "google_healthcare_dicom_store"
}

// LocalName returns the local name for [HealthcareDicomStore].
func (hds *HealthcareDicomStore) LocalName() string {
	return hds.Name
}

// Configuration returns the configuration (args) for [HealthcareDicomStore].
func (hds *HealthcareDicomStore) Configuration() interface{} {
	return hds.Args
}

// DependOn is used for other resources to depend on [HealthcareDicomStore].
func (hds *HealthcareDicomStore) DependOn() terra.Reference {
	return terra.ReferenceResource(hds)
}

// Dependencies returns the list of resources [HealthcareDicomStore] depends_on.
func (hds *HealthcareDicomStore) Dependencies() terra.Dependencies {
	return hds.DependsOn
}

// LifecycleManagement returns the lifecycle block for [HealthcareDicomStore].
func (hds *HealthcareDicomStore) LifecycleManagement() *terra.Lifecycle {
	return hds.Lifecycle
}

// Attributes returns the attributes for [HealthcareDicomStore].
func (hds *HealthcareDicomStore) Attributes() healthcareDicomStoreAttributes {
	return healthcareDicomStoreAttributes{ref: terra.ReferenceResource(hds)}
}

// ImportState imports the given attribute values into [HealthcareDicomStore]'s state.
func (hds *HealthcareDicomStore) ImportState(av io.Reader) error {
	hds.state = &healthcareDicomStoreState{}
	if err := json.NewDecoder(av).Decode(hds.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", hds.Type(), hds.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [HealthcareDicomStore] has state.
func (hds *HealthcareDicomStore) State() (*healthcareDicomStoreState, bool) {
	return hds.state, hds.state != nil
}

// StateMust returns the state for [HealthcareDicomStore]. Panics if the state is nil.
func (hds *HealthcareDicomStore) StateMust() *healthcareDicomStoreState {
	if hds.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", hds.Type(), hds.LocalName()))
	}
	return hds.state
}

// HealthcareDicomStoreArgs contains the configurations for google_healthcare_dicom_store.
type HealthcareDicomStoreArgs struct {
	// Dataset: string, required
	Dataset terra.StringValue `hcl:"dataset,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Labels: map of string, optional
	Labels terra.MapValue[terra.StringValue] `hcl:"labels,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// NotificationConfig: optional
	NotificationConfig *healthcaredicomstore.NotificationConfig `hcl:"notification_config,block"`
	// Timeouts: optional
	Timeouts *healthcaredicomstore.Timeouts `hcl:"timeouts,block"`
}
type healthcareDicomStoreAttributes struct {
	ref terra.Reference
}

// Dataset returns a reference to field dataset of google_healthcare_dicom_store.
func (hds healthcareDicomStoreAttributes) Dataset() terra.StringValue {
	return terra.ReferenceAsString(hds.ref.Append("dataset"))
}

// Id returns a reference to field id of google_healthcare_dicom_store.
func (hds healthcareDicomStoreAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(hds.ref.Append("id"))
}

// Labels returns a reference to field labels of google_healthcare_dicom_store.
func (hds healthcareDicomStoreAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](hds.ref.Append("labels"))
}

// Name returns a reference to field name of google_healthcare_dicom_store.
func (hds healthcareDicomStoreAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(hds.ref.Append("name"))
}

// SelfLink returns a reference to field self_link of google_healthcare_dicom_store.
func (hds healthcareDicomStoreAttributes) SelfLink() terra.StringValue {
	return terra.ReferenceAsString(hds.ref.Append("self_link"))
}

func (hds healthcareDicomStoreAttributes) NotificationConfig() terra.ListValue[healthcaredicomstore.NotificationConfigAttributes] {
	return terra.ReferenceAsList[healthcaredicomstore.NotificationConfigAttributes](hds.ref.Append("notification_config"))
}

func (hds healthcareDicomStoreAttributes) Timeouts() healthcaredicomstore.TimeoutsAttributes {
	return terra.ReferenceAsSingle[healthcaredicomstore.TimeoutsAttributes](hds.ref.Append("timeouts"))
}

type healthcareDicomStoreState struct {
	Dataset            string                                         `json:"dataset"`
	Id                 string                                         `json:"id"`
	Labels             map[string]string                              `json:"labels"`
	Name               string                                         `json:"name"`
	SelfLink           string                                         `json:"self_link"`
	NotificationConfig []healthcaredicomstore.NotificationConfigState `json:"notification_config"`
	Timeouts           *healthcaredicomstore.TimeoutsState            `json:"timeouts"`
}
