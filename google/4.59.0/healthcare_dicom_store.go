// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	healthcaredicomstore "github.com/golingon/terraproviders/google/4.59.0/healthcaredicomstore"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewHealthcareDicomStore(name string, args HealthcareDicomStoreArgs) *HealthcareDicomStore {
	return &HealthcareDicomStore{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*HealthcareDicomStore)(nil)

type HealthcareDicomStore struct {
	Name  string
	Args  HealthcareDicomStoreArgs
	state *healthcareDicomStoreState
}

func (hds *HealthcareDicomStore) Type() string {
	return "google_healthcare_dicom_store"
}

func (hds *HealthcareDicomStore) LocalName() string {
	return hds.Name
}

func (hds *HealthcareDicomStore) Configuration() interface{} {
	return hds.Args
}

func (hds *HealthcareDicomStore) Attributes() healthcareDicomStoreAttributes {
	return healthcareDicomStoreAttributes{ref: terra.ReferenceResource(hds)}
}

func (hds *HealthcareDicomStore) ImportState(av io.Reader) error {
	hds.state = &healthcareDicomStoreState{}
	if err := json.NewDecoder(av).Decode(hds.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", hds.Type(), hds.LocalName(), err)
	}
	return nil
}

func (hds *HealthcareDicomStore) State() (*healthcareDicomStoreState, bool) {
	return hds.state, hds.state != nil
}

func (hds *HealthcareDicomStore) StateMust() *healthcareDicomStoreState {
	if hds.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", hds.Type(), hds.LocalName()))
	}
	return hds.state
}

func (hds *HealthcareDicomStore) DependOn() terra.Reference {
	return terra.ReferenceResource(hds)
}

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
	// DependsOn contains resources that HealthcareDicomStore depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type healthcareDicomStoreAttributes struct {
	ref terra.Reference
}

func (hds healthcareDicomStoreAttributes) Dataset() terra.StringValue {
	return terra.ReferenceString(hds.ref.Append("dataset"))
}

func (hds healthcareDicomStoreAttributes) Id() terra.StringValue {
	return terra.ReferenceString(hds.ref.Append("id"))
}

func (hds healthcareDicomStoreAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](hds.ref.Append("labels"))
}

func (hds healthcareDicomStoreAttributes) Name() terra.StringValue {
	return terra.ReferenceString(hds.ref.Append("name"))
}

func (hds healthcareDicomStoreAttributes) SelfLink() terra.StringValue {
	return terra.ReferenceString(hds.ref.Append("self_link"))
}

func (hds healthcareDicomStoreAttributes) NotificationConfig() terra.ListValue[healthcaredicomstore.NotificationConfigAttributes] {
	return terra.ReferenceList[healthcaredicomstore.NotificationConfigAttributes](hds.ref.Append("notification_config"))
}

func (hds healthcareDicomStoreAttributes) Timeouts() healthcaredicomstore.TimeoutsAttributes {
	return terra.ReferenceSingle[healthcaredicomstore.TimeoutsAttributes](hds.ref.Append("timeouts"))
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
