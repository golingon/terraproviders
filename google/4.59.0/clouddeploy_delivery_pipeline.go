// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	clouddeploydeliverypipeline "github.com/golingon/terraproviders/google/4.59.0/clouddeploydeliverypipeline"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewClouddeployDeliveryPipeline(name string, args ClouddeployDeliveryPipelineArgs) *ClouddeployDeliveryPipeline {
	return &ClouddeployDeliveryPipeline{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ClouddeployDeliveryPipeline)(nil)

type ClouddeployDeliveryPipeline struct {
	Name  string
	Args  ClouddeployDeliveryPipelineArgs
	state *clouddeployDeliveryPipelineState
}

func (cdp *ClouddeployDeliveryPipeline) Type() string {
	return "google_clouddeploy_delivery_pipeline"
}

func (cdp *ClouddeployDeliveryPipeline) LocalName() string {
	return cdp.Name
}

func (cdp *ClouddeployDeliveryPipeline) Configuration() interface{} {
	return cdp.Args
}

func (cdp *ClouddeployDeliveryPipeline) Attributes() clouddeployDeliveryPipelineAttributes {
	return clouddeployDeliveryPipelineAttributes{ref: terra.ReferenceResource(cdp)}
}

func (cdp *ClouddeployDeliveryPipeline) ImportState(av io.Reader) error {
	cdp.state = &clouddeployDeliveryPipelineState{}
	if err := json.NewDecoder(av).Decode(cdp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cdp.Type(), cdp.LocalName(), err)
	}
	return nil
}

func (cdp *ClouddeployDeliveryPipeline) State() (*clouddeployDeliveryPipelineState, bool) {
	return cdp.state, cdp.state != nil
}

func (cdp *ClouddeployDeliveryPipeline) StateMust() *clouddeployDeliveryPipelineState {
	if cdp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cdp.Type(), cdp.LocalName()))
	}
	return cdp.state
}

func (cdp *ClouddeployDeliveryPipeline) DependOn() terra.Reference {
	return terra.ReferenceResource(cdp)
}

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
	// DependsOn contains resources that ClouddeployDeliveryPipeline depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type clouddeployDeliveryPipelineAttributes struct {
	ref terra.Reference
}

func (cdp clouddeployDeliveryPipelineAttributes) Annotations() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](cdp.ref.Append("annotations"))
}

func (cdp clouddeployDeliveryPipelineAttributes) CreateTime() terra.StringValue {
	return terra.ReferenceString(cdp.ref.Append("create_time"))
}

func (cdp clouddeployDeliveryPipelineAttributes) Description() terra.StringValue {
	return terra.ReferenceString(cdp.ref.Append("description"))
}

func (cdp clouddeployDeliveryPipelineAttributes) Etag() terra.StringValue {
	return terra.ReferenceString(cdp.ref.Append("etag"))
}

func (cdp clouddeployDeliveryPipelineAttributes) Id() terra.StringValue {
	return terra.ReferenceString(cdp.ref.Append("id"))
}

func (cdp clouddeployDeliveryPipelineAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](cdp.ref.Append("labels"))
}

func (cdp clouddeployDeliveryPipelineAttributes) Location() terra.StringValue {
	return terra.ReferenceString(cdp.ref.Append("location"))
}

func (cdp clouddeployDeliveryPipelineAttributes) Name() terra.StringValue {
	return terra.ReferenceString(cdp.ref.Append("name"))
}

func (cdp clouddeployDeliveryPipelineAttributes) Project() terra.StringValue {
	return terra.ReferenceString(cdp.ref.Append("project"))
}

func (cdp clouddeployDeliveryPipelineAttributes) Suspended() terra.BoolValue {
	return terra.ReferenceBool(cdp.ref.Append("suspended"))
}

func (cdp clouddeployDeliveryPipelineAttributes) Uid() terra.StringValue {
	return terra.ReferenceString(cdp.ref.Append("uid"))
}

func (cdp clouddeployDeliveryPipelineAttributes) UpdateTime() terra.StringValue {
	return terra.ReferenceString(cdp.ref.Append("update_time"))
}

func (cdp clouddeployDeliveryPipelineAttributes) Condition() terra.ListValue[clouddeploydeliverypipeline.ConditionAttributes] {
	return terra.ReferenceList[clouddeploydeliverypipeline.ConditionAttributes](cdp.ref.Append("condition"))
}

func (cdp clouddeployDeliveryPipelineAttributes) SerialPipeline() terra.ListValue[clouddeploydeliverypipeline.SerialPipelineAttributes] {
	return terra.ReferenceList[clouddeploydeliverypipeline.SerialPipelineAttributes](cdp.ref.Append("serial_pipeline"))
}

func (cdp clouddeployDeliveryPipelineAttributes) Timeouts() clouddeploydeliverypipeline.TimeoutsAttributes {
	return terra.ReferenceSingle[clouddeploydeliverypipeline.TimeoutsAttributes](cdp.ref.Append("timeouts"))
}

type clouddeployDeliveryPipelineState struct {
	Annotations    map[string]string                                 `json:"annotations"`
	CreateTime     string                                            `json:"create_time"`
	Description    string                                            `json:"description"`
	Etag           string                                            `json:"etag"`
	Id             string                                            `json:"id"`
	Labels         map[string]string                                 `json:"labels"`
	Location       string                                            `json:"location"`
	Name           string                                            `json:"name"`
	Project        string                                            `json:"project"`
	Suspended      bool                                              `json:"suspended"`
	Uid            string                                            `json:"uid"`
	UpdateTime     string                                            `json:"update_time"`
	Condition      []clouddeploydeliverypipeline.ConditionState      `json:"condition"`
	SerialPipeline []clouddeploydeliverypipeline.SerialPipelineState `json:"serial_pipeline"`
	Timeouts       *clouddeploydeliverypipeline.TimeoutsState        `json:"timeouts"`
}
