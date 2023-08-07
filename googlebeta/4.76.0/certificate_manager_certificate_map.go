// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	certificatemanagercertificatemap "github.com/golingon/terraproviders/googlebeta/4.76.0/certificatemanagercertificatemap"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewCertificateManagerCertificateMap creates a new instance of [CertificateManagerCertificateMap].
func NewCertificateManagerCertificateMap(name string, args CertificateManagerCertificateMapArgs) *CertificateManagerCertificateMap {
	return &CertificateManagerCertificateMap{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*CertificateManagerCertificateMap)(nil)

// CertificateManagerCertificateMap represents the Terraform resource google_certificate_manager_certificate_map.
type CertificateManagerCertificateMap struct {
	Name      string
	Args      CertificateManagerCertificateMapArgs
	state     *certificateManagerCertificateMapState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [CertificateManagerCertificateMap].
func (cmcm *CertificateManagerCertificateMap) Type() string {
	return "google_certificate_manager_certificate_map"
}

// LocalName returns the local name for [CertificateManagerCertificateMap].
func (cmcm *CertificateManagerCertificateMap) LocalName() string {
	return cmcm.Name
}

// Configuration returns the configuration (args) for [CertificateManagerCertificateMap].
func (cmcm *CertificateManagerCertificateMap) Configuration() interface{} {
	return cmcm.Args
}

// DependOn is used for other resources to depend on [CertificateManagerCertificateMap].
func (cmcm *CertificateManagerCertificateMap) DependOn() terra.Reference {
	return terra.ReferenceResource(cmcm)
}

// Dependencies returns the list of resources [CertificateManagerCertificateMap] depends_on.
func (cmcm *CertificateManagerCertificateMap) Dependencies() terra.Dependencies {
	return cmcm.DependsOn
}

// LifecycleManagement returns the lifecycle block for [CertificateManagerCertificateMap].
func (cmcm *CertificateManagerCertificateMap) LifecycleManagement() *terra.Lifecycle {
	return cmcm.Lifecycle
}

// Attributes returns the attributes for [CertificateManagerCertificateMap].
func (cmcm *CertificateManagerCertificateMap) Attributes() certificateManagerCertificateMapAttributes {
	return certificateManagerCertificateMapAttributes{ref: terra.ReferenceResource(cmcm)}
}

// ImportState imports the given attribute values into [CertificateManagerCertificateMap]'s state.
func (cmcm *CertificateManagerCertificateMap) ImportState(av io.Reader) error {
	cmcm.state = &certificateManagerCertificateMapState{}
	if err := json.NewDecoder(av).Decode(cmcm.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cmcm.Type(), cmcm.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [CertificateManagerCertificateMap] has state.
func (cmcm *CertificateManagerCertificateMap) State() (*certificateManagerCertificateMapState, bool) {
	return cmcm.state, cmcm.state != nil
}

// StateMust returns the state for [CertificateManagerCertificateMap]. Panics if the state is nil.
func (cmcm *CertificateManagerCertificateMap) StateMust() *certificateManagerCertificateMapState {
	if cmcm.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cmcm.Type(), cmcm.LocalName()))
	}
	return cmcm.state
}

// CertificateManagerCertificateMapArgs contains the configurations for google_certificate_manager_certificate_map.
type CertificateManagerCertificateMapArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Labels: map of string, optional
	Labels terra.MapValue[terra.StringValue] `hcl:"labels,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// GclbTargets: min=0
	GclbTargets []certificatemanagercertificatemap.GclbTargets `hcl:"gclb_targets,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *certificatemanagercertificatemap.Timeouts `hcl:"timeouts,block"`
}
type certificateManagerCertificateMapAttributes struct {
	ref terra.Reference
}

// CreateTime returns a reference to field create_time of google_certificate_manager_certificate_map.
func (cmcm certificateManagerCertificateMapAttributes) CreateTime() terra.StringValue {
	return terra.ReferenceAsString(cmcm.ref.Append("create_time"))
}

// Description returns a reference to field description of google_certificate_manager_certificate_map.
func (cmcm certificateManagerCertificateMapAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(cmcm.ref.Append("description"))
}

// Id returns a reference to field id of google_certificate_manager_certificate_map.
func (cmcm certificateManagerCertificateMapAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cmcm.ref.Append("id"))
}

// Labels returns a reference to field labels of google_certificate_manager_certificate_map.
func (cmcm certificateManagerCertificateMapAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](cmcm.ref.Append("labels"))
}

// Name returns a reference to field name of google_certificate_manager_certificate_map.
func (cmcm certificateManagerCertificateMapAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cmcm.ref.Append("name"))
}

// Project returns a reference to field project of google_certificate_manager_certificate_map.
func (cmcm certificateManagerCertificateMapAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(cmcm.ref.Append("project"))
}

// UpdateTime returns a reference to field update_time of google_certificate_manager_certificate_map.
func (cmcm certificateManagerCertificateMapAttributes) UpdateTime() terra.StringValue {
	return terra.ReferenceAsString(cmcm.ref.Append("update_time"))
}

func (cmcm certificateManagerCertificateMapAttributes) GclbTargets() terra.ListValue[certificatemanagercertificatemap.GclbTargetsAttributes] {
	return terra.ReferenceAsList[certificatemanagercertificatemap.GclbTargetsAttributes](cmcm.ref.Append("gclb_targets"))
}

func (cmcm certificateManagerCertificateMapAttributes) Timeouts() certificatemanagercertificatemap.TimeoutsAttributes {
	return terra.ReferenceAsSingle[certificatemanagercertificatemap.TimeoutsAttributes](cmcm.ref.Append("timeouts"))
}

type certificateManagerCertificateMapState struct {
	CreateTime  string                                              `json:"create_time"`
	Description string                                              `json:"description"`
	Id          string                                              `json:"id"`
	Labels      map[string]string                                   `json:"labels"`
	Name        string                                              `json:"name"`
	Project     string                                              `json:"project"`
	UpdateTime  string                                              `json:"update_time"`
	GclbTargets []certificatemanagercertificatemap.GclbTargetsState `json:"gclb_targets"`
	Timeouts    *certificatemanagercertificatemap.TimeoutsState     `json:"timeouts"`
}
