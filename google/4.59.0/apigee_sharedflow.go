// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	apigeesharedflow "github.com/golingon/terraproviders/google/4.59.0/apigeesharedflow"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewApigeeSharedflow(name string, args ApigeeSharedflowArgs) *ApigeeSharedflow {
	return &ApigeeSharedflow{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ApigeeSharedflow)(nil)

type ApigeeSharedflow struct {
	Name  string
	Args  ApigeeSharedflowArgs
	state *apigeeSharedflowState
}

func (as *ApigeeSharedflow) Type() string {
	return "google_apigee_sharedflow"
}

func (as *ApigeeSharedflow) LocalName() string {
	return as.Name
}

func (as *ApigeeSharedflow) Configuration() interface{} {
	return as.Args
}

func (as *ApigeeSharedflow) Attributes() apigeeSharedflowAttributes {
	return apigeeSharedflowAttributes{ref: terra.ReferenceResource(as)}
}

func (as *ApigeeSharedflow) ImportState(av io.Reader) error {
	as.state = &apigeeSharedflowState{}
	if err := json.NewDecoder(av).Decode(as.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", as.Type(), as.LocalName(), err)
	}
	return nil
}

func (as *ApigeeSharedflow) State() (*apigeeSharedflowState, bool) {
	return as.state, as.state != nil
}

func (as *ApigeeSharedflow) StateMust() *apigeeSharedflowState {
	if as.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", as.Type(), as.LocalName()))
	}
	return as.state
}

func (as *ApigeeSharedflow) DependOn() terra.Reference {
	return terra.ReferenceResource(as)
}

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
	// DependsOn contains resources that ApigeeSharedflow depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type apigeeSharedflowAttributes struct {
	ref terra.Reference
}

func (as apigeeSharedflowAttributes) ConfigBundle() terra.StringValue {
	return terra.ReferenceString(as.ref.Append("config_bundle"))
}

func (as apigeeSharedflowAttributes) DetectMd5Hash() terra.StringValue {
	return terra.ReferenceString(as.ref.Append("detect_md5hash"))
}

func (as apigeeSharedflowAttributes) Id() terra.StringValue {
	return terra.ReferenceString(as.ref.Append("id"))
}

func (as apigeeSharedflowAttributes) LatestRevisionId() terra.StringValue {
	return terra.ReferenceString(as.ref.Append("latest_revision_id"))
}

func (as apigeeSharedflowAttributes) Md5Hash() terra.StringValue {
	return terra.ReferenceString(as.ref.Append("md5hash"))
}

func (as apigeeSharedflowAttributes) Name() terra.StringValue {
	return terra.ReferenceString(as.ref.Append("name"))
}

func (as apigeeSharedflowAttributes) OrgId() terra.StringValue {
	return terra.ReferenceString(as.ref.Append("org_id"))
}

func (as apigeeSharedflowAttributes) Revision() terra.ListValue[terra.StringValue] {
	return terra.ReferenceList[terra.StringValue](as.ref.Append("revision"))
}

func (as apigeeSharedflowAttributes) MetaData() terra.ListValue[apigeesharedflow.MetaDataAttributes] {
	return terra.ReferenceList[apigeesharedflow.MetaDataAttributes](as.ref.Append("meta_data"))
}

func (as apigeeSharedflowAttributes) Timeouts() apigeesharedflow.TimeoutsAttributes {
	return terra.ReferenceSingle[apigeesharedflow.TimeoutsAttributes](as.ref.Append("timeouts"))
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
