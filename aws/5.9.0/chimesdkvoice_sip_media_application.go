// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	chimesdkvoicesipmediaapplication "github.com/golingon/terraproviders/aws/5.9.0/chimesdkvoicesipmediaapplication"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewChimesdkvoiceSipMediaApplication creates a new instance of [ChimesdkvoiceSipMediaApplication].
func NewChimesdkvoiceSipMediaApplication(name string, args ChimesdkvoiceSipMediaApplicationArgs) *ChimesdkvoiceSipMediaApplication {
	return &ChimesdkvoiceSipMediaApplication{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ChimesdkvoiceSipMediaApplication)(nil)

// ChimesdkvoiceSipMediaApplication represents the Terraform resource aws_chimesdkvoice_sip_media_application.
type ChimesdkvoiceSipMediaApplication struct {
	Name      string
	Args      ChimesdkvoiceSipMediaApplicationArgs
	state     *chimesdkvoiceSipMediaApplicationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ChimesdkvoiceSipMediaApplication].
func (csma *ChimesdkvoiceSipMediaApplication) Type() string {
	return "aws_chimesdkvoice_sip_media_application"
}

// LocalName returns the local name for [ChimesdkvoiceSipMediaApplication].
func (csma *ChimesdkvoiceSipMediaApplication) LocalName() string {
	return csma.Name
}

// Configuration returns the configuration (args) for [ChimesdkvoiceSipMediaApplication].
func (csma *ChimesdkvoiceSipMediaApplication) Configuration() interface{} {
	return csma.Args
}

// DependOn is used for other resources to depend on [ChimesdkvoiceSipMediaApplication].
func (csma *ChimesdkvoiceSipMediaApplication) DependOn() terra.Reference {
	return terra.ReferenceResource(csma)
}

// Dependencies returns the list of resources [ChimesdkvoiceSipMediaApplication] depends_on.
func (csma *ChimesdkvoiceSipMediaApplication) Dependencies() terra.Dependencies {
	return csma.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ChimesdkvoiceSipMediaApplication].
func (csma *ChimesdkvoiceSipMediaApplication) LifecycleManagement() *terra.Lifecycle {
	return csma.Lifecycle
}

// Attributes returns the attributes for [ChimesdkvoiceSipMediaApplication].
func (csma *ChimesdkvoiceSipMediaApplication) Attributes() chimesdkvoiceSipMediaApplicationAttributes {
	return chimesdkvoiceSipMediaApplicationAttributes{ref: terra.ReferenceResource(csma)}
}

// ImportState imports the given attribute values into [ChimesdkvoiceSipMediaApplication]'s state.
func (csma *ChimesdkvoiceSipMediaApplication) ImportState(av io.Reader) error {
	csma.state = &chimesdkvoiceSipMediaApplicationState{}
	if err := json.NewDecoder(av).Decode(csma.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", csma.Type(), csma.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ChimesdkvoiceSipMediaApplication] has state.
func (csma *ChimesdkvoiceSipMediaApplication) State() (*chimesdkvoiceSipMediaApplicationState, bool) {
	return csma.state, csma.state != nil
}

// StateMust returns the state for [ChimesdkvoiceSipMediaApplication]. Panics if the state is nil.
func (csma *ChimesdkvoiceSipMediaApplication) StateMust() *chimesdkvoiceSipMediaApplicationState {
	if csma.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", csma.Type(), csma.LocalName()))
	}
	return csma.state
}

// ChimesdkvoiceSipMediaApplicationArgs contains the configurations for aws_chimesdkvoice_sip_media_application.
type ChimesdkvoiceSipMediaApplicationArgs struct {
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
	Endpoints *chimesdkvoicesipmediaapplication.Endpoints `hcl:"endpoints,block" validate:"required"`
}
type chimesdkvoiceSipMediaApplicationAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_chimesdkvoice_sip_media_application.
func (csma chimesdkvoiceSipMediaApplicationAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(csma.ref.Append("arn"))
}

// AwsRegion returns a reference to field aws_region of aws_chimesdkvoice_sip_media_application.
func (csma chimesdkvoiceSipMediaApplicationAttributes) AwsRegion() terra.StringValue {
	return terra.ReferenceAsString(csma.ref.Append("aws_region"))
}

// Id returns a reference to field id of aws_chimesdkvoice_sip_media_application.
func (csma chimesdkvoiceSipMediaApplicationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(csma.ref.Append("id"))
}

// Name returns a reference to field name of aws_chimesdkvoice_sip_media_application.
func (csma chimesdkvoiceSipMediaApplicationAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(csma.ref.Append("name"))
}

// Tags returns a reference to field tags of aws_chimesdkvoice_sip_media_application.
func (csma chimesdkvoiceSipMediaApplicationAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](csma.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_chimesdkvoice_sip_media_application.
func (csma chimesdkvoiceSipMediaApplicationAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](csma.ref.Append("tags_all"))
}

func (csma chimesdkvoiceSipMediaApplicationAttributes) Endpoints() terra.ListValue[chimesdkvoicesipmediaapplication.EndpointsAttributes] {
	return terra.ReferenceAsList[chimesdkvoicesipmediaapplication.EndpointsAttributes](csma.ref.Append("endpoints"))
}

type chimesdkvoiceSipMediaApplicationState struct {
	Arn       string                                            `json:"arn"`
	AwsRegion string                                            `json:"aws_region"`
	Id        string                                            `json:"id"`
	Name      string                                            `json:"name"`
	Tags      map[string]string                                 `json:"tags"`
	TagsAll   map[string]string                                 `json:"tags_all"`
	Endpoints []chimesdkvoicesipmediaapplication.EndpointsState `json:"endpoints"`
}
