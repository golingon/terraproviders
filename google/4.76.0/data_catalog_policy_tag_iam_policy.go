// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDataCatalogPolicyTagIamPolicy creates a new instance of [DataCatalogPolicyTagIamPolicy].
func NewDataCatalogPolicyTagIamPolicy(name string, args DataCatalogPolicyTagIamPolicyArgs) *DataCatalogPolicyTagIamPolicy {
	return &DataCatalogPolicyTagIamPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DataCatalogPolicyTagIamPolicy)(nil)

// DataCatalogPolicyTagIamPolicy represents the Terraform resource google_data_catalog_policy_tag_iam_policy.
type DataCatalogPolicyTagIamPolicy struct {
	Name      string
	Args      DataCatalogPolicyTagIamPolicyArgs
	state     *dataCatalogPolicyTagIamPolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DataCatalogPolicyTagIamPolicy].
func (dcptip *DataCatalogPolicyTagIamPolicy) Type() string {
	return "google_data_catalog_policy_tag_iam_policy"
}

// LocalName returns the local name for [DataCatalogPolicyTagIamPolicy].
func (dcptip *DataCatalogPolicyTagIamPolicy) LocalName() string {
	return dcptip.Name
}

// Configuration returns the configuration (args) for [DataCatalogPolicyTagIamPolicy].
func (dcptip *DataCatalogPolicyTagIamPolicy) Configuration() interface{} {
	return dcptip.Args
}

// DependOn is used for other resources to depend on [DataCatalogPolicyTagIamPolicy].
func (dcptip *DataCatalogPolicyTagIamPolicy) DependOn() terra.Reference {
	return terra.ReferenceResource(dcptip)
}

// Dependencies returns the list of resources [DataCatalogPolicyTagIamPolicy] depends_on.
func (dcptip *DataCatalogPolicyTagIamPolicy) Dependencies() terra.Dependencies {
	return dcptip.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DataCatalogPolicyTagIamPolicy].
func (dcptip *DataCatalogPolicyTagIamPolicy) LifecycleManagement() *terra.Lifecycle {
	return dcptip.Lifecycle
}

// Attributes returns the attributes for [DataCatalogPolicyTagIamPolicy].
func (dcptip *DataCatalogPolicyTagIamPolicy) Attributes() dataCatalogPolicyTagIamPolicyAttributes {
	return dataCatalogPolicyTagIamPolicyAttributes{ref: terra.ReferenceResource(dcptip)}
}

// ImportState imports the given attribute values into [DataCatalogPolicyTagIamPolicy]'s state.
func (dcptip *DataCatalogPolicyTagIamPolicy) ImportState(av io.Reader) error {
	dcptip.state = &dataCatalogPolicyTagIamPolicyState{}
	if err := json.NewDecoder(av).Decode(dcptip.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dcptip.Type(), dcptip.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DataCatalogPolicyTagIamPolicy] has state.
func (dcptip *DataCatalogPolicyTagIamPolicy) State() (*dataCatalogPolicyTagIamPolicyState, bool) {
	return dcptip.state, dcptip.state != nil
}

// StateMust returns the state for [DataCatalogPolicyTagIamPolicy]. Panics if the state is nil.
func (dcptip *DataCatalogPolicyTagIamPolicy) StateMust() *dataCatalogPolicyTagIamPolicyState {
	if dcptip.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dcptip.Type(), dcptip.LocalName()))
	}
	return dcptip.state
}

// DataCatalogPolicyTagIamPolicyArgs contains the configurations for google_data_catalog_policy_tag_iam_policy.
type DataCatalogPolicyTagIamPolicyArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// PolicyData: string, required
	PolicyData terra.StringValue `hcl:"policy_data,attr" validate:"required"`
	// PolicyTag: string, required
	PolicyTag terra.StringValue `hcl:"policy_tag,attr" validate:"required"`
}
type dataCatalogPolicyTagIamPolicyAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_data_catalog_policy_tag_iam_policy.
func (dcptip dataCatalogPolicyTagIamPolicyAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(dcptip.ref.Append("etag"))
}

// Id returns a reference to field id of google_data_catalog_policy_tag_iam_policy.
func (dcptip dataCatalogPolicyTagIamPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dcptip.ref.Append("id"))
}

// PolicyData returns a reference to field policy_data of google_data_catalog_policy_tag_iam_policy.
func (dcptip dataCatalogPolicyTagIamPolicyAttributes) PolicyData() terra.StringValue {
	return terra.ReferenceAsString(dcptip.ref.Append("policy_data"))
}

// PolicyTag returns a reference to field policy_tag of google_data_catalog_policy_tag_iam_policy.
func (dcptip dataCatalogPolicyTagIamPolicyAttributes) PolicyTag() terra.StringValue {
	return terra.ReferenceAsString(dcptip.ref.Append("policy_tag"))
}

type dataCatalogPolicyTagIamPolicyState struct {
	Etag       string `json:"etag"`
	Id         string `json:"id"`
	PolicyData string `json:"policy_data"`
	PolicyTag  string `json:"policy_tag"`
}
