// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	apigeesharedflow "github.com/golingon/terraproviders/googlebeta/4.72.1/apigeesharedflow"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewApigeeSharedflow creates a new instance of [ApigeeSharedflow].
func NewApigeeSharedflow(name string, args ApigeeSharedflowArgs) *ApigeeSharedflow {
	return &ApigeeSharedflow{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ApigeeSharedflow)(nil)

// ApigeeSharedflow represents the Terraform resource google_apigee_sharedflow.
type ApigeeSharedflow struct {
	Name      string
	Args      ApigeeSharedflowArgs
	state     *apigeeSharedflowState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ApigeeSharedflow].
func (as *ApigeeSharedflow) Type() string {
	return "google_apigee_sharedflow"
}

// LocalName returns the local name for [ApigeeSharedflow].
func (as *ApigeeSharedflow) LocalName() string {
	return as.Name
}

// Configuration returns the configuration (args) for [ApigeeSharedflow].
func (as *ApigeeSharedflow) Configuration() interface{} {
	return as.Args
}

// DependOn is used for other resources to depend on [ApigeeSharedflow].
func (as *ApigeeSharedflow) DependOn() terra.Reference {
	return terra.ReferenceResource(as)
}

// Dependencies returns the list of resources [ApigeeSharedflow] depends_on.
func (as *ApigeeSharedflow) Dependencies() terra.Dependencies {
	return as.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ApigeeSharedflow].
func (as *ApigeeSharedflow) LifecycleManagement() *terra.Lifecycle {
	return as.Lifecycle
}

// Attributes returns the attributes for [ApigeeSharedflow].
func (as *ApigeeSharedflow) Attributes() apigeeSharedflowAttributes {
	return apigeeSharedflowAttributes{ref: terra.ReferenceResource(as)}
}

// ImportState imports the given attribute values into [ApigeeSharedflow]'s state.
func (as *ApigeeSharedflow) ImportState(av io.Reader) error {
	as.state = &apigeeSharedflowState{}
	if err := json.NewDecoder(av).Decode(as.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", as.Type(), as.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ApigeeSharedflow] has state.
func (as *ApigeeSharedflow) State() (*apigeeSharedflowState, bool) {
	return as.state, as.state != nil
}

// StateMust returns the state for [ApigeeSharedflow]. Panics if the state is nil.
func (as *ApigeeSharedflow) StateMust() *apigeeSharedflowState {
	if as.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", as.Type(), as.LocalName()))
	}
	return as.state
}

// ApigeeSharedflowArgs contains the configurations for google_apigee_sharedflow.
type ApigeeSharedflowArgs struct {
	// ConfigBundle: string, required
	ConfigBundle terra.StringValue `hcl:"config_bundle,attr" validate:"required"`
	// DetectMd5Hash: string, optional
	DetectMd5Hash terra.StringValue `hcl:"detect_md5hash,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// OrgId: string, required
	OrgId terra.StringValue `hcl:"org_id,attr" validate:"required"`
	// MetaData: min=0
	MetaData []apigeesharedflow.MetaData `hcl:"meta_data,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *apigeesharedflow.Timeouts `hcl:"timeouts,block"`
}
type apigeeSharedflowAttributes struct {
	ref terra.Reference
}

// ConfigBundle returns a reference to field config_bundle of google_apigee_sharedflow.
func (as apigeeSharedflowAttributes) ConfigBundle() terra.StringValue {
	return terra.ReferenceAsString(as.ref.Append("config_bundle"))
}

// DetectMd5Hash returns a reference to field detect_md5hash of google_apigee_sharedflow.
func (as apigeeSharedflowAttributes) DetectMd5Hash() terra.StringValue {
	return terra.ReferenceAsString(as.ref.Append("detect_md5hash"))
}

// Id returns a reference to field id of google_apigee_sharedflow.
func (as apigeeSharedflowAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(as.ref.Append("id"))
}

// LatestRevisionId returns a reference to field latest_revision_id of google_apigee_sharedflow.
func (as apigeeSharedflowAttributes) LatestRevisionId() terra.StringValue {
	return terra.ReferenceAsString(as.ref.Append("latest_revision_id"))
}

// Md5Hash returns a reference to field md5hash of google_apigee_sharedflow.
func (as apigeeSharedflowAttributes) Md5Hash() terra.StringValue {
	return terra.ReferenceAsString(as.ref.Append("md5hash"))
}

// Name returns a reference to field name of google_apigee_sharedflow.
func (as apigeeSharedflowAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(as.ref.Append("name"))
}

// OrgId returns a reference to field org_id of google_apigee_sharedflow.
func (as apigeeSharedflowAttributes) OrgId() terra.StringValue {
	return terra.ReferenceAsString(as.ref.Append("org_id"))
}

// Revision returns a reference to field revision of google_apigee_sharedflow.
func (as apigeeSharedflowAttributes) Revision() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](as.ref.Append("revision"))
}

func (as apigeeSharedflowAttributes) MetaData() terra.ListValue[apigeesharedflow.MetaDataAttributes] {
	return terra.ReferenceAsList[apigeesharedflow.MetaDataAttributes](as.ref.Append("meta_data"))
}

func (as apigeeSharedflowAttributes) Timeouts() apigeesharedflow.TimeoutsAttributes {
	return terra.ReferenceAsSingle[apigeesharedflow.TimeoutsAttributes](as.ref.Append("timeouts"))
}

type apigeeSharedflowState struct {
	ConfigBundle     string                           `json:"config_bundle"`
	DetectMd5Hash    string                           `json:"detect_md5hash"`
	Id               string                           `json:"id"`
	LatestRevisionId string                           `json:"latest_revision_id"`
	Md5Hash          string                           `json:"md5hash"`
	Name             string                           `json:"name"`
	OrgId            string                           `json:"org_id"`
	Revision         []string                         `json:"revision"`
	MetaData         []apigeesharedflow.MetaDataState `json:"meta_data"`
	Timeouts         *apigeesharedflow.TimeoutsState  `json:"timeouts"`
}
