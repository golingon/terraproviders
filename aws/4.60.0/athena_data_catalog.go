// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewAthenaDataCatalog creates a new instance of [AthenaDataCatalog].
func NewAthenaDataCatalog(name string, args AthenaDataCatalogArgs) *AthenaDataCatalog {
	return &AthenaDataCatalog{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*AthenaDataCatalog)(nil)

// AthenaDataCatalog represents the Terraform resource aws_athena_data_catalog.
type AthenaDataCatalog struct {
	Name      string
	Args      AthenaDataCatalogArgs
	state     *athenaDataCatalogState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [AthenaDataCatalog].
func (adc *AthenaDataCatalog) Type() string {
	return "aws_athena_data_catalog"
}

// LocalName returns the local name for [AthenaDataCatalog].
func (adc *AthenaDataCatalog) LocalName() string {
	return adc.Name
}

// Configuration returns the configuration (args) for [AthenaDataCatalog].
func (adc *AthenaDataCatalog) Configuration() interface{} {
	return adc.Args
}

// DependOn is used for other resources to depend on [AthenaDataCatalog].
func (adc *AthenaDataCatalog) DependOn() terra.Reference {
	return terra.ReferenceResource(adc)
}

// Dependencies returns the list of resources [AthenaDataCatalog] depends_on.
func (adc *AthenaDataCatalog) Dependencies() terra.Dependencies {
	return adc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [AthenaDataCatalog].
func (adc *AthenaDataCatalog) LifecycleManagement() *terra.Lifecycle {
	return adc.Lifecycle
}

// Attributes returns the attributes for [AthenaDataCatalog].
func (adc *AthenaDataCatalog) Attributes() athenaDataCatalogAttributes {
	return athenaDataCatalogAttributes{ref: terra.ReferenceResource(adc)}
}

// ImportState imports the given attribute values into [AthenaDataCatalog]'s state.
func (adc *AthenaDataCatalog) ImportState(av io.Reader) error {
	adc.state = &athenaDataCatalogState{}
	if err := json.NewDecoder(av).Decode(adc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", adc.Type(), adc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [AthenaDataCatalog] has state.
func (adc *AthenaDataCatalog) State() (*athenaDataCatalogState, bool) {
	return adc.state, adc.state != nil
}

// StateMust returns the state for [AthenaDataCatalog]. Panics if the state is nil.
func (adc *AthenaDataCatalog) StateMust() *athenaDataCatalogState {
	if adc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", adc.Type(), adc.LocalName()))
	}
	return adc.state
}

// AthenaDataCatalogArgs contains the configurations for aws_athena_data_catalog.
type AthenaDataCatalogArgs struct {
	// Description: string, required
	Description terra.StringValue `hcl:"description,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Parameters: map of string, required
	Parameters terra.MapValue[terra.StringValue] `hcl:"parameters,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// Type: string, required
	Type terra.StringValue `hcl:"type,attr" validate:"required"`
}
type athenaDataCatalogAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_athena_data_catalog.
func (adc athenaDataCatalogAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(adc.ref.Append("arn"))
}

// Description returns a reference to field description of aws_athena_data_catalog.
func (adc athenaDataCatalogAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(adc.ref.Append("description"))
}

// Id returns a reference to field id of aws_athena_data_catalog.
func (adc athenaDataCatalogAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(adc.ref.Append("id"))
}

// Name returns a reference to field name of aws_athena_data_catalog.
func (adc athenaDataCatalogAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(adc.ref.Append("name"))
}

// Parameters returns a reference to field parameters of aws_athena_data_catalog.
func (adc athenaDataCatalogAttributes) Parameters() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](adc.ref.Append("parameters"))
}

// Tags returns a reference to field tags of aws_athena_data_catalog.
func (adc athenaDataCatalogAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](adc.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_athena_data_catalog.
func (adc athenaDataCatalogAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](adc.ref.Append("tags_all"))
}

// Type returns a reference to field type of aws_athena_data_catalog.
func (adc athenaDataCatalogAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(adc.ref.Append("type"))
}

type athenaDataCatalogState struct {
	Arn         string            `json:"arn"`
	Description string            `json:"description"`
	Id          string            `json:"id"`
	Name        string            `json:"name"`
	Parameters  map[string]string `json:"parameters"`
	Tags        map[string]string `json:"tags"`
	TagsAll     map[string]string `json:"tags_all"`
	Type        string            `json:"type"`
}