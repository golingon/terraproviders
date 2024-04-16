// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_chimesdkvoice_sip_media_application

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// New creates a new instance of [Resource].
func New(name string, args Args) *Resource {
	return &Resource{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Resource)(nil)

// Resource represents the Terraform resource aws_chimesdkvoice_sip_media_application.
type Resource struct {
	Name      string
	Args      Args
	state     *awsChimesdkvoiceSipMediaApplicationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (acsma *Resource) Type() string {
	return "aws_chimesdkvoice_sip_media_application"
}

// LocalName returns the local name for [Resource].
func (acsma *Resource) LocalName() string {
	return acsma.Name
}

// Configuration returns the configuration (args) for [Resource].
func (acsma *Resource) Configuration() interface{} {
	return acsma.Args
}

// DependOn is used for other resources to depend on [Resource].
func (acsma *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(acsma)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (acsma *Resource) Dependencies() terra.Dependencies {
	return acsma.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (acsma *Resource) LifecycleManagement() *terra.Lifecycle {
	return acsma.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (acsma *Resource) Attributes() awsChimesdkvoiceSipMediaApplicationAttributes {
	return awsChimesdkvoiceSipMediaApplicationAttributes{ref: terra.ReferenceResource(acsma)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (acsma *Resource) ImportState(state io.Reader) error {
	acsma.state = &awsChimesdkvoiceSipMediaApplicationState{}
	if err := json.NewDecoder(state).Decode(acsma.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", acsma.Type(), acsma.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (acsma *Resource) State() (*awsChimesdkvoiceSipMediaApplicationState, bool) {
	return acsma.state, acsma.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (acsma *Resource) StateMust() *awsChimesdkvoiceSipMediaApplicationState {
	if acsma.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", acsma.Type(), acsma.LocalName()))
	}
	return acsma.state
}

// Args contains the configurations for aws_chimesdkvoice_sip_media_application.
type Args struct {
	// AwsRegion: string, required
	AwsRegion terra.StringValue `hcl:"aws_region,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// Endpoints: required
	Endpoints *Endpoints `hcl:"endpoints,block" validate:"required"`
}

type awsChimesdkvoiceSipMediaApplicationAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_chimesdkvoice_sip_media_application.
func (acsma awsChimesdkvoiceSipMediaApplicationAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(acsma.ref.Append("arn"))
}

// AwsRegion returns a reference to field aws_region of aws_chimesdkvoice_sip_media_application.
func (acsma awsChimesdkvoiceSipMediaApplicationAttributes) AwsRegion() terra.StringValue {
	return terra.ReferenceAsString(acsma.ref.Append("aws_region"))
}

// Id returns a reference to field id of aws_chimesdkvoice_sip_media_application.
func (acsma awsChimesdkvoiceSipMediaApplicationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(acsma.ref.Append("id"))
}

// Name returns a reference to field name of aws_chimesdkvoice_sip_media_application.
func (acsma awsChimesdkvoiceSipMediaApplicationAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(acsma.ref.Append("name"))
}

// Tags returns a reference to field tags of aws_chimesdkvoice_sip_media_application.
func (acsma awsChimesdkvoiceSipMediaApplicationAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](acsma.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_chimesdkvoice_sip_media_application.
func (acsma awsChimesdkvoiceSipMediaApplicationAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](acsma.ref.Append("tags_all"))
}

func (acsma awsChimesdkvoiceSipMediaApplicationAttributes) Endpoints() terra.ListValue[EndpointsAttributes] {
	return terra.ReferenceAsList[EndpointsAttributes](acsma.ref.Append("endpoints"))
}

type awsChimesdkvoiceSipMediaApplicationState struct {
	Arn       string            `json:"arn"`
	AwsRegion string            `json:"aws_region"`
	Id        string            `json:"id"`
	Name      string            `json:"name"`
	Tags      map[string]string `json:"tags"`
	TagsAll   map[string]string `json:"tags_all"`
	Endpoints []EndpointsState  `json:"endpoints"`
}
