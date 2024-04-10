// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	clouddeploydeliverypipelineiambinding "github.com/golingon/terraproviders/google/5.24.0/clouddeploydeliverypipelineiambinding"
	"io"
)

// NewClouddeployDeliveryPipelineIamBinding creates a new instance of [ClouddeployDeliveryPipelineIamBinding].
func NewClouddeployDeliveryPipelineIamBinding(name string, args ClouddeployDeliveryPipelineIamBindingArgs) *ClouddeployDeliveryPipelineIamBinding {
	return &ClouddeployDeliveryPipelineIamBinding{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ClouddeployDeliveryPipelineIamBinding)(nil)

// ClouddeployDeliveryPipelineIamBinding represents the Terraform resource google_clouddeploy_delivery_pipeline_iam_binding.
type ClouddeployDeliveryPipelineIamBinding struct {
	Name      string
	Args      ClouddeployDeliveryPipelineIamBindingArgs
	state     *clouddeployDeliveryPipelineIamBindingState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ClouddeployDeliveryPipelineIamBinding].
func (cdpib *ClouddeployDeliveryPipelineIamBinding) Type() string {
	return "google_clouddeploy_delivery_pipeline_iam_binding"
}

// LocalName returns the local name for [ClouddeployDeliveryPipelineIamBinding].
func (cdpib *ClouddeployDeliveryPipelineIamBinding) LocalName() string {
	return cdpib.Name
}

// Configuration returns the configuration (args) for [ClouddeployDeliveryPipelineIamBinding].
func (cdpib *ClouddeployDeliveryPipelineIamBinding) Configuration() interface{} {
	return cdpib.Args
}

// DependOn is used for other resources to depend on [ClouddeployDeliveryPipelineIamBinding].
func (cdpib *ClouddeployDeliveryPipelineIamBinding) DependOn() terra.Reference {
	return terra.ReferenceResource(cdpib)
}

// Dependencies returns the list of resources [ClouddeployDeliveryPipelineIamBinding] depends_on.
func (cdpib *ClouddeployDeliveryPipelineIamBinding) Dependencies() terra.Dependencies {
	return cdpib.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ClouddeployDeliveryPipelineIamBinding].
func (cdpib *ClouddeployDeliveryPipelineIamBinding) LifecycleManagement() *terra.Lifecycle {
	return cdpib.Lifecycle
}

// Attributes returns the attributes for [ClouddeployDeliveryPipelineIamBinding].
func (cdpib *ClouddeployDeliveryPipelineIamBinding) Attributes() clouddeployDeliveryPipelineIamBindingAttributes {
	return clouddeployDeliveryPipelineIamBindingAttributes{ref: terra.ReferenceResource(cdpib)}
}

// ImportState imports the given attribute values into [ClouddeployDeliveryPipelineIamBinding]'s state.
func (cdpib *ClouddeployDeliveryPipelineIamBinding) ImportState(av io.Reader) error {
	cdpib.state = &clouddeployDeliveryPipelineIamBindingState{}
	if err := json.NewDecoder(av).Decode(cdpib.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cdpib.Type(), cdpib.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ClouddeployDeliveryPipelineIamBinding] has state.
func (cdpib *ClouddeployDeliveryPipelineIamBinding) State() (*clouddeployDeliveryPipelineIamBindingState, bool) {
	return cdpib.state, cdpib.state != nil
}

// StateMust returns the state for [ClouddeployDeliveryPipelineIamBinding]. Panics if the state is nil.
func (cdpib *ClouddeployDeliveryPipelineIamBinding) StateMust() *clouddeployDeliveryPipelineIamBindingState {
	if cdpib.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cdpib.Type(), cdpib.LocalName()))
	}
	return cdpib.state
}

// ClouddeployDeliveryPipelineIamBindingArgs contains the configurations for google_clouddeploy_delivery_pipeline_iam_binding.
type ClouddeployDeliveryPipelineIamBindingArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, optional
	Location terra.StringValue `hcl:"location,attr"`
	// Members: set of string, required
	Members terra.SetValue[terra.StringValue] `hcl:"members,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// Condition: optional
	Condition *clouddeploydeliverypipelineiambinding.Condition `hcl:"condition,block"`
}
type clouddeployDeliveryPipelineIamBindingAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_clouddeploy_delivery_pipeline_iam_binding.
func (cdpib clouddeployDeliveryPipelineIamBindingAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(cdpib.ref.Append("etag"))
}

// Id returns a reference to field id of google_clouddeploy_delivery_pipeline_iam_binding.
func (cdpib clouddeployDeliveryPipelineIamBindingAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cdpib.ref.Append("id"))
}

// Location returns a reference to field location of google_clouddeploy_delivery_pipeline_iam_binding.
func (cdpib clouddeployDeliveryPipelineIamBindingAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(cdpib.ref.Append("location"))
}

// Members returns a reference to field members of google_clouddeploy_delivery_pipeline_iam_binding.
func (cdpib clouddeployDeliveryPipelineIamBindingAttributes) Members() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](cdpib.ref.Append("members"))
}

// Name returns a reference to field name of google_clouddeploy_delivery_pipeline_iam_binding.
func (cdpib clouddeployDeliveryPipelineIamBindingAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cdpib.ref.Append("name"))
}

// Project returns a reference to field project of google_clouddeploy_delivery_pipeline_iam_binding.
func (cdpib clouddeployDeliveryPipelineIamBindingAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(cdpib.ref.Append("project"))
}

// Role returns a reference to field role of google_clouddeploy_delivery_pipeline_iam_binding.
func (cdpib clouddeployDeliveryPipelineIamBindingAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(cdpib.ref.Append("role"))
}

func (cdpib clouddeployDeliveryPipelineIamBindingAttributes) Condition() terra.ListValue[clouddeploydeliverypipelineiambinding.ConditionAttributes] {
	return terra.ReferenceAsList[clouddeploydeliverypipelineiambinding.ConditionAttributes](cdpib.ref.Append("condition"))
}

type clouddeployDeliveryPipelineIamBindingState struct {
	Etag      string                                                 `json:"etag"`
	Id        string                                                 `json:"id"`
	Location  string                                                 `json:"location"`
	Members   []string                                               `json:"members"`
	Name      string                                                 `json:"name"`
	Project   string                                                 `json:"project"`
	Role      string                                                 `json:"role"`
	Condition []clouddeploydeliverypipelineiambinding.ConditionState `json:"condition"`
}
