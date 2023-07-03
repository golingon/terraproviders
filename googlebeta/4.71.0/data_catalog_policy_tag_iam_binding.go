// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	datacatalogpolicytagiambinding "github.com/golingon/terraproviders/googlebeta/4.71.0/datacatalogpolicytagiambinding"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDataCatalogPolicyTagIamBinding creates a new instance of [DataCatalogPolicyTagIamBinding].
func NewDataCatalogPolicyTagIamBinding(name string, args DataCatalogPolicyTagIamBindingArgs) *DataCatalogPolicyTagIamBinding {
	return &DataCatalogPolicyTagIamBinding{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DataCatalogPolicyTagIamBinding)(nil)

// DataCatalogPolicyTagIamBinding represents the Terraform resource google_data_catalog_policy_tag_iam_binding.
type DataCatalogPolicyTagIamBinding struct {
	Name      string
	Args      DataCatalogPolicyTagIamBindingArgs
	state     *dataCatalogPolicyTagIamBindingState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DataCatalogPolicyTagIamBinding].
func (dcptib *DataCatalogPolicyTagIamBinding) Type() string {
	return "google_data_catalog_policy_tag_iam_binding"
}

// LocalName returns the local name for [DataCatalogPolicyTagIamBinding].
func (dcptib *DataCatalogPolicyTagIamBinding) LocalName() string {
	return dcptib.Name
}

// Configuration returns the configuration (args) for [DataCatalogPolicyTagIamBinding].
func (dcptib *DataCatalogPolicyTagIamBinding) Configuration() interface{} {
	return dcptib.Args
}

// DependOn is used for other resources to depend on [DataCatalogPolicyTagIamBinding].
func (dcptib *DataCatalogPolicyTagIamBinding) DependOn() terra.Reference {
	return terra.ReferenceResource(dcptib)
}

// Dependencies returns the list of resources [DataCatalogPolicyTagIamBinding] depends_on.
func (dcptib *DataCatalogPolicyTagIamBinding) Dependencies() terra.Dependencies {
	return dcptib.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DataCatalogPolicyTagIamBinding].
func (dcptib *DataCatalogPolicyTagIamBinding) LifecycleManagement() *terra.Lifecycle {
	return dcptib.Lifecycle
}

// Attributes returns the attributes for [DataCatalogPolicyTagIamBinding].
func (dcptib *DataCatalogPolicyTagIamBinding) Attributes() dataCatalogPolicyTagIamBindingAttributes {
	return dataCatalogPolicyTagIamBindingAttributes{ref: terra.ReferenceResource(dcptib)}
}

// ImportState imports the given attribute values into [DataCatalogPolicyTagIamBinding]'s state.
func (dcptib *DataCatalogPolicyTagIamBinding) ImportState(av io.Reader) error {
	dcptib.state = &dataCatalogPolicyTagIamBindingState{}
	if err := json.NewDecoder(av).Decode(dcptib.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dcptib.Type(), dcptib.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DataCatalogPolicyTagIamBinding] has state.
func (dcptib *DataCatalogPolicyTagIamBinding) State() (*dataCatalogPolicyTagIamBindingState, bool) {
	return dcptib.state, dcptib.state != nil
}

// StateMust returns the state for [DataCatalogPolicyTagIamBinding]. Panics if the state is nil.
func (dcptib *DataCatalogPolicyTagIamBinding) StateMust() *dataCatalogPolicyTagIamBindingState {
	if dcptib.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dcptib.Type(), dcptib.LocalName()))
	}
	return dcptib.state
}

// DataCatalogPolicyTagIamBindingArgs contains the configurations for google_data_catalog_policy_tag_iam_binding.
type DataCatalogPolicyTagIamBindingArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Members: set of string, required
	Members terra.SetValue[terra.StringValue] `hcl:"members,attr" validate:"required"`
	// PolicyTag: string, required
	PolicyTag terra.StringValue `hcl:"policy_tag,attr" validate:"required"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// Condition: optional
	Condition *datacatalogpolicytagiambinding.Condition `hcl:"condition,block"`
}
type dataCatalogPolicyTagIamBindingAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_data_catalog_policy_tag_iam_binding.
func (dcptib dataCatalogPolicyTagIamBindingAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(dcptib.ref.Append("etag"))
}

// Id returns a reference to field id of google_data_catalog_policy_tag_iam_binding.
func (dcptib dataCatalogPolicyTagIamBindingAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dcptib.ref.Append("id"))
}

// Members returns a reference to field members of google_data_catalog_policy_tag_iam_binding.
func (dcptib dataCatalogPolicyTagIamBindingAttributes) Members() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](dcptib.ref.Append("members"))
}

// PolicyTag returns a reference to field policy_tag of google_data_catalog_policy_tag_iam_binding.
func (dcptib dataCatalogPolicyTagIamBindingAttributes) PolicyTag() terra.StringValue {
	return terra.ReferenceAsString(dcptib.ref.Append("policy_tag"))
}

// Role returns a reference to field role of google_data_catalog_policy_tag_iam_binding.
func (dcptib dataCatalogPolicyTagIamBindingAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(dcptib.ref.Append("role"))
}

func (dcptib dataCatalogPolicyTagIamBindingAttributes) Condition() terra.ListValue[datacatalogpolicytagiambinding.ConditionAttributes] {
	return terra.ReferenceAsList[datacatalogpolicytagiambinding.ConditionAttributes](dcptib.ref.Append("condition"))
}

type dataCatalogPolicyTagIamBindingState struct {
	Etag      string                                          `json:"etag"`
	Id        string                                          `json:"id"`
	Members   []string                                        `json:"members"`
	PolicyTag string                                          `json:"policy_tag"`
	Role      string                                          `json:"role"`
	Condition []datacatalogpolicytagiambinding.ConditionState `json:"condition"`
}
