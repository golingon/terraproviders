// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_healthcare_dicom_store

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

// Resource represents the Terraform resource google_healthcare_dicom_store.
type Resource struct {
	Name      string
	Args      Args
	state     *googleHealthcareDicomStoreState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (ghds *Resource) Type() string {
	return "google_healthcare_dicom_store"
}

// LocalName returns the local name for [Resource].
func (ghds *Resource) LocalName() string {
	return ghds.Name
}

// Configuration returns the configuration (args) for [Resource].
func (ghds *Resource) Configuration() interface{} {
	return ghds.Args
}

// DependOn is used for other resources to depend on [Resource].
func (ghds *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(ghds)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (ghds *Resource) Dependencies() terra.Dependencies {
	return ghds.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (ghds *Resource) LifecycleManagement() *terra.Lifecycle {
	return ghds.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (ghds *Resource) Attributes() googleHealthcareDicomStoreAttributes {
	return googleHealthcareDicomStoreAttributes{ref: terra.ReferenceResource(ghds)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (ghds *Resource) ImportState(state io.Reader) error {
	ghds.state = &googleHealthcareDicomStoreState{}
	if err := json.NewDecoder(state).Decode(ghds.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ghds.Type(), ghds.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (ghds *Resource) State() (*googleHealthcareDicomStoreState, bool) {
	return ghds.state, ghds.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (ghds *Resource) StateMust() *googleHealthcareDicomStoreState {
	if ghds.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ghds.Type(), ghds.LocalName()))
	}
	return ghds.state
}

// Args contains the configurations for google_healthcare_dicom_store.
type Args struct {
	// Dataset: string, required
	Dataset terra.StringValue `hcl:"dataset,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Labels: map of string, optional
	Labels terra.MapValue[terra.StringValue] `hcl:"labels,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// NotificationConfig: optional
	NotificationConfig *NotificationConfig `hcl:"notification_config,block"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type googleHealthcareDicomStoreAttributes struct {
	ref terra.Reference
}

// Dataset returns a reference to field dataset of google_healthcare_dicom_store.
func (ghds googleHealthcareDicomStoreAttributes) Dataset() terra.StringValue {
	return terra.ReferenceAsString(ghds.ref.Append("dataset"))
}

// EffectiveLabels returns a reference to field effective_labels of google_healthcare_dicom_store.
func (ghds googleHealthcareDicomStoreAttributes) EffectiveLabels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ghds.ref.Append("effective_labels"))
}

// Id returns a reference to field id of google_healthcare_dicom_store.
func (ghds googleHealthcareDicomStoreAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ghds.ref.Append("id"))
}

// Labels returns a reference to field labels of google_healthcare_dicom_store.
func (ghds googleHealthcareDicomStoreAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ghds.ref.Append("labels"))
}

// Name returns a reference to field name of google_healthcare_dicom_store.
func (ghds googleHealthcareDicomStoreAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ghds.ref.Append("name"))
}

// SelfLink returns a reference to field self_link of google_healthcare_dicom_store.
func (ghds googleHealthcareDicomStoreAttributes) SelfLink() terra.StringValue {
	return terra.ReferenceAsString(ghds.ref.Append("self_link"))
}

// TerraformLabels returns a reference to field terraform_labels of google_healthcare_dicom_store.
func (ghds googleHealthcareDicomStoreAttributes) TerraformLabels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ghds.ref.Append("terraform_labels"))
}

func (ghds googleHealthcareDicomStoreAttributes) NotificationConfig() terra.ListValue[NotificationConfigAttributes] {
	return terra.ReferenceAsList[NotificationConfigAttributes](ghds.ref.Append("notification_config"))
}

func (ghds googleHealthcareDicomStoreAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](ghds.ref.Append("timeouts"))
}

type googleHealthcareDicomStoreState struct {
	Dataset            string                    `json:"dataset"`
	EffectiveLabels    map[string]string         `json:"effective_labels"`
	Id                 string                    `json:"id"`
	Labels             map[string]string         `json:"labels"`
	Name               string                    `json:"name"`
	SelfLink           string                    `json:"self_link"`
	TerraformLabels    map[string]string         `json:"terraform_labels"`
	NotificationConfig []NotificationConfigState `json:"notification_config"`
	Timeouts           *TimeoutsState            `json:"timeouts"`
}
