// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDataCatalogTagTemplateIamPolicy creates a new instance of [DataCatalogTagTemplateIamPolicy].
func NewDataCatalogTagTemplateIamPolicy(name string, args DataCatalogTagTemplateIamPolicyArgs) *DataCatalogTagTemplateIamPolicy {
	return &DataCatalogTagTemplateIamPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DataCatalogTagTemplateIamPolicy)(nil)

// DataCatalogTagTemplateIamPolicy represents the Terraform resource google_data_catalog_tag_template_iam_policy.
type DataCatalogTagTemplateIamPolicy struct {
	Name      string
	Args      DataCatalogTagTemplateIamPolicyArgs
	state     *dataCatalogTagTemplateIamPolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DataCatalogTagTemplateIamPolicy].
func (dcttip *DataCatalogTagTemplateIamPolicy) Type() string {
	return "google_data_catalog_tag_template_iam_policy"
}

// LocalName returns the local name for [DataCatalogTagTemplateIamPolicy].
func (dcttip *DataCatalogTagTemplateIamPolicy) LocalName() string {
	return dcttip.Name
}

// Configuration returns the configuration (args) for [DataCatalogTagTemplateIamPolicy].
func (dcttip *DataCatalogTagTemplateIamPolicy) Configuration() interface{} {
	return dcttip.Args
}

// DependOn is used for other resources to depend on [DataCatalogTagTemplateIamPolicy].
func (dcttip *DataCatalogTagTemplateIamPolicy) DependOn() terra.Reference {
	return terra.ReferenceResource(dcttip)
}

// Dependencies returns the list of resources [DataCatalogTagTemplateIamPolicy] depends_on.
func (dcttip *DataCatalogTagTemplateIamPolicy) Dependencies() terra.Dependencies {
	return dcttip.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DataCatalogTagTemplateIamPolicy].
func (dcttip *DataCatalogTagTemplateIamPolicy) LifecycleManagement() *terra.Lifecycle {
	return dcttip.Lifecycle
}

// Attributes returns the attributes for [DataCatalogTagTemplateIamPolicy].
func (dcttip *DataCatalogTagTemplateIamPolicy) Attributes() dataCatalogTagTemplateIamPolicyAttributes {
	return dataCatalogTagTemplateIamPolicyAttributes{ref: terra.ReferenceResource(dcttip)}
}

// ImportState imports the given attribute values into [DataCatalogTagTemplateIamPolicy]'s state.
func (dcttip *DataCatalogTagTemplateIamPolicy) ImportState(av io.Reader) error {
	dcttip.state = &dataCatalogTagTemplateIamPolicyState{}
	if err := json.NewDecoder(av).Decode(dcttip.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dcttip.Type(), dcttip.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DataCatalogTagTemplateIamPolicy] has state.
func (dcttip *DataCatalogTagTemplateIamPolicy) State() (*dataCatalogTagTemplateIamPolicyState, bool) {
	return dcttip.state, dcttip.state != nil
}

// StateMust returns the state for [DataCatalogTagTemplateIamPolicy]. Panics if the state is nil.
func (dcttip *DataCatalogTagTemplateIamPolicy) StateMust() *dataCatalogTagTemplateIamPolicyState {
	if dcttip.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dcttip.Type(), dcttip.LocalName()))
	}
	return dcttip.state
}

// DataCatalogTagTemplateIamPolicyArgs contains the configurations for google_data_catalog_tag_template_iam_policy.
type DataCatalogTagTemplateIamPolicyArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// PolicyData: string, required
	PolicyData terra.StringValue `hcl:"policy_data,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Region: string, optional
	Region terra.StringValue `hcl:"region,attr"`
	// TagTemplate: string, required
	TagTemplate terra.StringValue `hcl:"tag_template,attr" validate:"required"`
}
type dataCatalogTagTemplateIamPolicyAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_data_catalog_tag_template_iam_policy.
func (dcttip dataCatalogTagTemplateIamPolicyAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(dcttip.ref.Append("etag"))
}

// Id returns a reference to field id of google_data_catalog_tag_template_iam_policy.
func (dcttip dataCatalogTagTemplateIamPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dcttip.ref.Append("id"))
}

// PolicyData returns a reference to field policy_data of google_data_catalog_tag_template_iam_policy.
func (dcttip dataCatalogTagTemplateIamPolicyAttributes) PolicyData() terra.StringValue {
	return terra.ReferenceAsString(dcttip.ref.Append("policy_data"))
}

// Project returns a reference to field project of google_data_catalog_tag_template_iam_policy.
func (dcttip dataCatalogTagTemplateIamPolicyAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(dcttip.ref.Append("project"))
}

// Region returns a reference to field region of google_data_catalog_tag_template_iam_policy.
func (dcttip dataCatalogTagTemplateIamPolicyAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(dcttip.ref.Append("region"))
}

// TagTemplate returns a reference to field tag_template of google_data_catalog_tag_template_iam_policy.
func (dcttip dataCatalogTagTemplateIamPolicyAttributes) TagTemplate() terra.StringValue {
	return terra.ReferenceAsString(dcttip.ref.Append("tag_template"))
}

type dataCatalogTagTemplateIamPolicyState struct {
	Etag        string `json:"etag"`
	Id          string `json:"id"`
	PolicyData  string `json:"policy_data"`
	Project     string `json:"project"`
	Region      string `json:"region"`
	TagTemplate string `json:"tag_template"`
}
