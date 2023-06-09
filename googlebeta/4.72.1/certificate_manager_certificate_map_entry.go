// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	certificatemanagercertificatemapentry "github.com/golingon/terraproviders/googlebeta/4.72.1/certificatemanagercertificatemapentry"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewCertificateManagerCertificateMapEntry creates a new instance of [CertificateManagerCertificateMapEntry].
func NewCertificateManagerCertificateMapEntry(name string, args CertificateManagerCertificateMapEntryArgs) *CertificateManagerCertificateMapEntry {
	return &CertificateManagerCertificateMapEntry{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*CertificateManagerCertificateMapEntry)(nil)

// CertificateManagerCertificateMapEntry represents the Terraform resource google_certificate_manager_certificate_map_entry.
type CertificateManagerCertificateMapEntry struct {
	Name      string
	Args      CertificateManagerCertificateMapEntryArgs
	state     *certificateManagerCertificateMapEntryState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [CertificateManagerCertificateMapEntry].
func (cmcme *CertificateManagerCertificateMapEntry) Type() string {
	return "google_certificate_manager_certificate_map_entry"
}

// LocalName returns the local name for [CertificateManagerCertificateMapEntry].
func (cmcme *CertificateManagerCertificateMapEntry) LocalName() string {
	return cmcme.Name
}

// Configuration returns the configuration (args) for [CertificateManagerCertificateMapEntry].
func (cmcme *CertificateManagerCertificateMapEntry) Configuration() interface{} {
	return cmcme.Args
}

// DependOn is used for other resources to depend on [CertificateManagerCertificateMapEntry].
func (cmcme *CertificateManagerCertificateMapEntry) DependOn() terra.Reference {
	return terra.ReferenceResource(cmcme)
}

// Dependencies returns the list of resources [CertificateManagerCertificateMapEntry] depends_on.
func (cmcme *CertificateManagerCertificateMapEntry) Dependencies() terra.Dependencies {
	return cmcme.DependsOn
}

// LifecycleManagement returns the lifecycle block for [CertificateManagerCertificateMapEntry].
func (cmcme *CertificateManagerCertificateMapEntry) LifecycleManagement() *terra.Lifecycle {
	return cmcme.Lifecycle
}

// Attributes returns the attributes for [CertificateManagerCertificateMapEntry].
func (cmcme *CertificateManagerCertificateMapEntry) Attributes() certificateManagerCertificateMapEntryAttributes {
	return certificateManagerCertificateMapEntryAttributes{ref: terra.ReferenceResource(cmcme)}
}

// ImportState imports the given attribute values into [CertificateManagerCertificateMapEntry]'s state.
func (cmcme *CertificateManagerCertificateMapEntry) ImportState(av io.Reader) error {
	cmcme.state = &certificateManagerCertificateMapEntryState{}
	if err := json.NewDecoder(av).Decode(cmcme.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cmcme.Type(), cmcme.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [CertificateManagerCertificateMapEntry] has state.
func (cmcme *CertificateManagerCertificateMapEntry) State() (*certificateManagerCertificateMapEntryState, bool) {
	return cmcme.state, cmcme.state != nil
}

// StateMust returns the state for [CertificateManagerCertificateMapEntry]. Panics if the state is nil.
func (cmcme *CertificateManagerCertificateMapEntry) StateMust() *certificateManagerCertificateMapEntryState {
	if cmcme.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cmcme.Type(), cmcme.LocalName()))
	}
	return cmcme.state
}

// CertificateManagerCertificateMapEntryArgs contains the configurations for google_certificate_manager_certificate_map_entry.
type CertificateManagerCertificateMapEntryArgs struct {
	// Certificates: list of string, required
	Certificates terra.ListValue[terra.StringValue] `hcl:"certificates,attr" validate:"required"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Hostname: string, optional
	Hostname terra.StringValue `hcl:"hostname,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Labels: map of string, optional
	Labels terra.MapValue[terra.StringValue] `hcl:"labels,attr"`
	// Map: string, required
	Map terra.StringValue `hcl:"map,attr" validate:"required"`
	// Matcher: string, optional
	Matcher terra.StringValue `hcl:"matcher,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Timeouts: optional
	Timeouts *certificatemanagercertificatemapentry.Timeouts `hcl:"timeouts,block"`
}
type certificateManagerCertificateMapEntryAttributes struct {
	ref terra.Reference
}

// Certificates returns a reference to field certificates of google_certificate_manager_certificate_map_entry.
func (cmcme certificateManagerCertificateMapEntryAttributes) Certificates() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](cmcme.ref.Append("certificates"))
}

// CreateTime returns a reference to field create_time of google_certificate_manager_certificate_map_entry.
func (cmcme certificateManagerCertificateMapEntryAttributes) CreateTime() terra.StringValue {
	return terra.ReferenceAsString(cmcme.ref.Append("create_time"))
}

// Description returns a reference to field description of google_certificate_manager_certificate_map_entry.
func (cmcme certificateManagerCertificateMapEntryAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(cmcme.ref.Append("description"))
}

// Hostname returns a reference to field hostname of google_certificate_manager_certificate_map_entry.
func (cmcme certificateManagerCertificateMapEntryAttributes) Hostname() terra.StringValue {
	return terra.ReferenceAsString(cmcme.ref.Append("hostname"))
}

// Id returns a reference to field id of google_certificate_manager_certificate_map_entry.
func (cmcme certificateManagerCertificateMapEntryAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cmcme.ref.Append("id"))
}

// Labels returns a reference to field labels of google_certificate_manager_certificate_map_entry.
func (cmcme certificateManagerCertificateMapEntryAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](cmcme.ref.Append("labels"))
}

// Map returns a reference to field map of google_certificate_manager_certificate_map_entry.
func (cmcme certificateManagerCertificateMapEntryAttributes) Map() terra.StringValue {
	return terra.ReferenceAsString(cmcme.ref.Append("map"))
}

// Matcher returns a reference to field matcher of google_certificate_manager_certificate_map_entry.
func (cmcme certificateManagerCertificateMapEntryAttributes) Matcher() terra.StringValue {
	return terra.ReferenceAsString(cmcme.ref.Append("matcher"))
}

// Name returns a reference to field name of google_certificate_manager_certificate_map_entry.
func (cmcme certificateManagerCertificateMapEntryAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cmcme.ref.Append("name"))
}

// Project returns a reference to field project of google_certificate_manager_certificate_map_entry.
func (cmcme certificateManagerCertificateMapEntryAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(cmcme.ref.Append("project"))
}

// State returns a reference to field state of google_certificate_manager_certificate_map_entry.
func (cmcme certificateManagerCertificateMapEntryAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(cmcme.ref.Append("state"))
}

// UpdateTime returns a reference to field update_time of google_certificate_manager_certificate_map_entry.
func (cmcme certificateManagerCertificateMapEntryAttributes) UpdateTime() terra.StringValue {
	return terra.ReferenceAsString(cmcme.ref.Append("update_time"))
}

func (cmcme certificateManagerCertificateMapEntryAttributes) Timeouts() certificatemanagercertificatemapentry.TimeoutsAttributes {
	return terra.ReferenceAsSingle[certificatemanagercertificatemapentry.TimeoutsAttributes](cmcme.ref.Append("timeouts"))
}

type certificateManagerCertificateMapEntryState struct {
	Certificates []string                                             `json:"certificates"`
	CreateTime   string                                               `json:"create_time"`
	Description  string                                               `json:"description"`
	Hostname     string                                               `json:"hostname"`
	Id           string                                               `json:"id"`
	Labels       map[string]string                                    `json:"labels"`
	Map          string                                               `json:"map"`
	Matcher      string                                               `json:"matcher"`
	Name         string                                               `json:"name"`
	Project      string                                               `json:"project"`
	State        string                                               `json:"state"`
	UpdateTime   string                                               `json:"update_time"`
	Timeouts     *certificatemanagercertificatemapentry.TimeoutsState `json:"timeouts"`
}
