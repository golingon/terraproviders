// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	clouddeploydeliverypipeline "github.com/golingon/terraproviders/googlebeta/5.24.0/clouddeploydeliverypipeline"
	"io"
)

// NewClouddeployDeliveryPipeline creates a new instance of [ClouddeployDeliveryPipeline].
func NewClouddeployDeliveryPipeline(name string, args ClouddeployDeliveryPipelineArgs) *ClouddeployDeliveryPipeline {
	return &ClouddeployDeliveryPipeline{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ClouddeployDeliveryPipeline)(nil)

// ClouddeployDeliveryPipeline represents the Terraform resource google_clouddeploy_delivery_pipeline.
type ClouddeployDeliveryPipeline struct {
	Name      string
	Args      ClouddeployDeliveryPipelineArgs
	state     *clouddeployDeliveryPipelineState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ClouddeployDeliveryPipeline].
func (cdp *ClouddeployDeliveryPipeline) Type() string {
	return "google_clouddeploy_delivery_pipeline"
}

// LocalName returns the local name for [ClouddeployDeliveryPipeline].
func (cdp *ClouddeployDeliveryPipeline) LocalName() string {
	return cdp.Name
}

// Configuration returns the configuration (args) for [ClouddeployDeliveryPipeline].
func (cdp *ClouddeployDeliveryPipeline) Configuration() interface{} {
	return cdp.Args
}

// DependOn is used for other resources to depend on [ClouddeployDeliveryPipeline].
func (cdp *ClouddeployDeliveryPipeline) DependOn() terra.Reference {
	return terra.ReferenceResource(cdp)
}

// Dependencies returns the list of resources [ClouddeployDeliveryPipeline] depends_on.
func (cdp *ClouddeployDeliveryPipeline) Dependencies() terra.Dependencies {
	return cdp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ClouddeployDeliveryPipeline].
func (cdp *ClouddeployDeliveryPipeline) LifecycleManagement() *terra.Lifecycle {
	return cdp.Lifecycle
}

// Attributes returns the attributes for [ClouddeployDeliveryPipeline].
func (cdp *ClouddeployDeliveryPipeline) Attributes() clouddeployDeliveryPipelineAttributes {
	return clouddeployDeliveryPipelineAttributes{ref: terra.ReferenceResource(cdp)}
}

// ImportState imports the given attribute values into [ClouddeployDeliveryPipeline]'s state.
func (cdp *ClouddeployDeliveryPipeline) ImportState(av io.Reader) error {
	cdp.state = &clouddeployDeliveryPipelineState{}
	if err := json.NewDecoder(av).Decode(cdp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cdp.Type(), cdp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ClouddeployDeliveryPipeline] has state.
func (cdp *ClouddeployDeliveryPipeline) State() (*clouddeployDeliveryPipelineState, bool) {
	return cdp.state, cdp.state != nil
}

// StateMust returns the state for [ClouddeployDeliveryPipeline]. Panics if the state is nil.
func (cdp *ClouddeployDeliveryPipeline) StateMust() *clouddeployDeliveryPipelineState {
	if cdp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cdp.Type(), cdp.LocalName()))
	}
	return cdp.state
}

// ClouddeployDeliveryPipelineArgs contains the configurations for google_clouddeploy_delivery_pipeline.
type ClouddeployDeliveryPipelineArgs struct {
	// Annotations: map of string, optional
	Annotations terra.MapValue[terra.StringValue] `hcl:"annotations,attr"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Labels: map of string, optional
	Labels terra.MapValue[terra.StringValue] `hcl:"labels,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Suspended: bool, optional
	Suspended terra.BoolValue `hcl:"suspended,attr"`
	// Condition: min=0
	Condition []clouddeploydeliverypipeline.Condition `hcl:"condition,block" validate:"min=0"`
	// SerialPipeline: optional
	SerialPipeline *clouddeploydeliverypipeline.SerialPipeline `hcl:"serial_pipeline,block"`
	// Timeouts: optional
	Timeouts *clouddeploydeliverypipeline.Timeouts `hcl:"timeouts,block"`
}
type clouddeployDeliveryPipelineAttributes struct {
	ref terra.Reference
}

// Annotations returns a reference to field annotations of google_clouddeploy_delivery_pipeline.
func (cdp clouddeployDeliveryPipelineAttributes) Annotations() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](cdp.ref.Append("annotations"))
}

// CreateTime returns a reference to field create_time of google_clouddeploy_delivery_pipeline.
func (cdp clouddeployDeliveryPipelineAttributes) CreateTime() terra.StringValue {
	return terra.ReferenceAsString(cdp.ref.Append("create_time"))
}

// Description returns a reference to field description of google_clouddeploy_delivery_pipeline.
func (cdp clouddeployDeliveryPipelineAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(cdp.ref.Append("description"))
}

// EffectiveAnnotations returns a reference to field effective_annotations of google_clouddeploy_delivery_pipeline.
func (cdp clouddeployDeliveryPipelineAttributes) EffectiveAnnotations() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](cdp.ref.Append("effective_annotations"))
}

// EffectiveLabels returns a reference to field effective_labels of google_clouddeploy_delivery_pipeline.
func (cdp clouddeployDeliveryPipelineAttributes) EffectiveLabels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](cdp.ref.Append("effective_labels"))
}

// Etag returns a reference to field etag of google_clouddeploy_delivery_pipeline.
func (cdp clouddeployDeliveryPipelineAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(cdp.ref.Append("etag"))
}

// Id returns a reference to field id of google_clouddeploy_delivery_pipeline.
func (cdp clouddeployDeliveryPipelineAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cdp.ref.Append("id"))
}

// Labels returns a reference to field labels of google_clouddeploy_delivery_pipeline.
func (cdp clouddeployDeliveryPipelineAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](cdp.ref.Append("labels"))
}

// Location returns a reference to field location of google_clouddeploy_delivery_pipeline.
func (cdp clouddeployDeliveryPipelineAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(cdp.ref.Append("location"))
}

// Name returns a reference to field name of google_clouddeploy_delivery_pipeline.
func (cdp clouddeployDeliveryPipelineAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cdp.ref.Append("name"))
}

// Project returns a reference to field project of google_clouddeploy_delivery_pipeline.
func (cdp clouddeployDeliveryPipelineAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(cdp.ref.Append("project"))
}

// Suspended returns a reference to field suspended of google_clouddeploy_delivery_pipeline.
func (cdp clouddeployDeliveryPipelineAttributes) Suspended() terra.BoolValue {
	return terra.ReferenceAsBool(cdp.ref.Append("suspended"))
}

// TerraformLabels returns a reference to field terraform_labels of google_clouddeploy_delivery_pipeline.
func (cdp clouddeployDeliveryPipelineAttributes) TerraformLabels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](cdp.ref.Append("terraform_labels"))
}

// Uid returns a reference to field uid of google_clouddeploy_delivery_pipeline.
func (cdp clouddeployDeliveryPipelineAttributes) Uid() terra.StringValue {
	return terra.ReferenceAsString(cdp.ref.Append("uid"))
}

// UpdateTime returns a reference to field update_time of google_clouddeploy_delivery_pipeline.
func (cdp clouddeployDeliveryPipelineAttributes) UpdateTime() terra.StringValue {
	return terra.ReferenceAsString(cdp.ref.Append("update_time"))
}

func (cdp clouddeployDeliveryPipelineAttributes) Condition() terra.ListValue[clouddeploydeliverypipeline.ConditionAttributes] {
	return terra.ReferenceAsList[clouddeploydeliverypipeline.ConditionAttributes](cdp.ref.Append("condition"))
}

func (cdp clouddeployDeliveryPipelineAttributes) SerialPipeline() terra.ListValue[clouddeploydeliverypipeline.SerialPipelineAttributes] {
	return terra.ReferenceAsList[clouddeploydeliverypipeline.SerialPipelineAttributes](cdp.ref.Append("serial_pipeline"))
}

func (cdp clouddeployDeliveryPipelineAttributes) Timeouts() clouddeploydeliverypipeline.TimeoutsAttributes {
	return terra.ReferenceAsSingle[clouddeploydeliverypipeline.TimeoutsAttributes](cdp.ref.Append("timeouts"))
}

type clouddeployDeliveryPipelineState struct {
	Annotations          map[string]string                                 `json:"annotations"`
	CreateTime           string                                            `json:"create_time"`
	Description          string                                            `json:"description"`
	EffectiveAnnotations map[string]string                                 `json:"effective_annotations"`
	EffectiveLabels      map[string]string                                 `json:"effective_labels"`
	Etag                 string                                            `json:"etag"`
	Id                   string                                            `json:"id"`
	Labels               map[string]string                                 `json:"labels"`
	Location             string                                            `json:"location"`
	Name                 string                                            `json:"name"`
	Project              string                                            `json:"project"`
	Suspended            bool                                              `json:"suspended"`
	TerraformLabels      map[string]string                                 `json:"terraform_labels"`
	Uid                  string                                            `json:"uid"`
	UpdateTime           string                                            `json:"update_time"`
	Condition            []clouddeploydeliverypipeline.ConditionState      `json:"condition"`
	SerialPipeline       []clouddeploydeliverypipeline.SerialPipelineState `json:"serial_pipeline"`
	Timeouts             *clouddeploydeliverypipeline.TimeoutsState        `json:"timeouts"`
}
