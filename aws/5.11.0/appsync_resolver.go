// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	appsyncresolver "github.com/golingon/terraproviders/aws/5.11.0/appsyncresolver"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewAppsyncResolver creates a new instance of [AppsyncResolver].
func NewAppsyncResolver(name string, args AppsyncResolverArgs) *AppsyncResolver {
	return &AppsyncResolver{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*AppsyncResolver)(nil)

// AppsyncResolver represents the Terraform resource aws_appsync_resolver.
type AppsyncResolver struct {
	Name      string
	Args      AppsyncResolverArgs
	state     *appsyncResolverState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [AppsyncResolver].
func (ar *AppsyncResolver) Type() string {
	return "aws_appsync_resolver"
}

// LocalName returns the local name for [AppsyncResolver].
func (ar *AppsyncResolver) LocalName() string {
	return ar.Name
}

// Configuration returns the configuration (args) for [AppsyncResolver].
func (ar *AppsyncResolver) Configuration() interface{} {
	return ar.Args
}

// DependOn is used for other resources to depend on [AppsyncResolver].
func (ar *AppsyncResolver) DependOn() terra.Reference {
	return terra.ReferenceResource(ar)
}

// Dependencies returns the list of resources [AppsyncResolver] depends_on.
func (ar *AppsyncResolver) Dependencies() terra.Dependencies {
	return ar.DependsOn
}

// LifecycleManagement returns the lifecycle block for [AppsyncResolver].
func (ar *AppsyncResolver) LifecycleManagement() *terra.Lifecycle {
	return ar.Lifecycle
}

// Attributes returns the attributes for [AppsyncResolver].
func (ar *AppsyncResolver) Attributes() appsyncResolverAttributes {
	return appsyncResolverAttributes{ref: terra.ReferenceResource(ar)}
}

// ImportState imports the given attribute values into [AppsyncResolver]'s state.
func (ar *AppsyncResolver) ImportState(av io.Reader) error {
	ar.state = &appsyncResolverState{}
	if err := json.NewDecoder(av).Decode(ar.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ar.Type(), ar.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [AppsyncResolver] has state.
func (ar *AppsyncResolver) State() (*appsyncResolverState, bool) {
	return ar.state, ar.state != nil
}

// StateMust returns the state for [AppsyncResolver]. Panics if the state is nil.
func (ar *AppsyncResolver) StateMust() *appsyncResolverState {
	if ar.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ar.Type(), ar.LocalName()))
	}
	return ar.state
}

// AppsyncResolverArgs contains the configurations for aws_appsync_resolver.
type AppsyncResolverArgs struct {
	// ApiId: string, required
	ApiId terra.StringValue `hcl:"api_id,attr" validate:"required"`
	// Code: string, optional
	Code terra.StringValue `hcl:"code,attr"`
	// DataSource: string, optional
	DataSource terra.StringValue `hcl:"data_source,attr"`
	// Field: string, required
	Field terra.StringValue `hcl:"field,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Kind: string, optional
	Kind terra.StringValue `hcl:"kind,attr"`
	// MaxBatchSize: number, optional
	MaxBatchSize terra.NumberValue `hcl:"max_batch_size,attr"`
	// RequestTemplate: string, optional
	RequestTemplate terra.StringValue `hcl:"request_template,attr"`
	// ResponseTemplate: string, optional
	ResponseTemplate terra.StringValue `hcl:"response_template,attr"`
	// Type: string, required
	Type terra.StringValue `hcl:"type,attr" validate:"required"`
	// CachingConfig: optional
	CachingConfig *appsyncresolver.CachingConfig `hcl:"caching_config,block"`
	// PipelineConfig: optional
	PipelineConfig *appsyncresolver.PipelineConfig `hcl:"pipeline_config,block"`
	// Runtime: optional
	Runtime *appsyncresolver.Runtime `hcl:"runtime,block"`
	// SyncConfig: optional
	SyncConfig *appsyncresolver.SyncConfig `hcl:"sync_config,block"`
}
type appsyncResolverAttributes struct {
	ref terra.Reference
}

// ApiId returns a reference to field api_id of aws_appsync_resolver.
func (ar appsyncResolverAttributes) ApiId() terra.StringValue {
	return terra.ReferenceAsString(ar.ref.Append("api_id"))
}

// Arn returns a reference to field arn of aws_appsync_resolver.
func (ar appsyncResolverAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(ar.ref.Append("arn"))
}

// Code returns a reference to field code of aws_appsync_resolver.
func (ar appsyncResolverAttributes) Code() terra.StringValue {
	return terra.ReferenceAsString(ar.ref.Append("code"))
}

// DataSource returns a reference to field data_source of aws_appsync_resolver.
func (ar appsyncResolverAttributes) DataSource() terra.StringValue {
	return terra.ReferenceAsString(ar.ref.Append("data_source"))
}

// Field returns a reference to field field of aws_appsync_resolver.
func (ar appsyncResolverAttributes) Field() terra.StringValue {
	return terra.ReferenceAsString(ar.ref.Append("field"))
}

// Id returns a reference to field id of aws_appsync_resolver.
func (ar appsyncResolverAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ar.ref.Append("id"))
}

// Kind returns a reference to field kind of aws_appsync_resolver.
func (ar appsyncResolverAttributes) Kind() terra.StringValue {
	return terra.ReferenceAsString(ar.ref.Append("kind"))
}

// MaxBatchSize returns a reference to field max_batch_size of aws_appsync_resolver.
func (ar appsyncResolverAttributes) MaxBatchSize() terra.NumberValue {
	return terra.ReferenceAsNumber(ar.ref.Append("max_batch_size"))
}

// RequestTemplate returns a reference to field request_template of aws_appsync_resolver.
func (ar appsyncResolverAttributes) RequestTemplate() terra.StringValue {
	return terra.ReferenceAsString(ar.ref.Append("request_template"))
}

// ResponseTemplate returns a reference to field response_template of aws_appsync_resolver.
func (ar appsyncResolverAttributes) ResponseTemplate() terra.StringValue {
	return terra.ReferenceAsString(ar.ref.Append("response_template"))
}

// Type returns a reference to field type of aws_appsync_resolver.
func (ar appsyncResolverAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(ar.ref.Append("type"))
}

func (ar appsyncResolverAttributes) CachingConfig() terra.ListValue[appsyncresolver.CachingConfigAttributes] {
	return terra.ReferenceAsList[appsyncresolver.CachingConfigAttributes](ar.ref.Append("caching_config"))
}

func (ar appsyncResolverAttributes) PipelineConfig() terra.ListValue[appsyncresolver.PipelineConfigAttributes] {
	return terra.ReferenceAsList[appsyncresolver.PipelineConfigAttributes](ar.ref.Append("pipeline_config"))
}

func (ar appsyncResolverAttributes) Runtime() terra.ListValue[appsyncresolver.RuntimeAttributes] {
	return terra.ReferenceAsList[appsyncresolver.RuntimeAttributes](ar.ref.Append("runtime"))
}

func (ar appsyncResolverAttributes) SyncConfig() terra.ListValue[appsyncresolver.SyncConfigAttributes] {
	return terra.ReferenceAsList[appsyncresolver.SyncConfigAttributes](ar.ref.Append("sync_config"))
}

type appsyncResolverState struct {
	ApiId            string                                `json:"api_id"`
	Arn              string                                `json:"arn"`
	Code             string                                `json:"code"`
	DataSource       string                                `json:"data_source"`
	Field            string                                `json:"field"`
	Id               string                                `json:"id"`
	Kind             string                                `json:"kind"`
	MaxBatchSize     float64                               `json:"max_batch_size"`
	RequestTemplate  string                                `json:"request_template"`
	ResponseTemplate string                                `json:"response_template"`
	Type             string                                `json:"type"`
	CachingConfig    []appsyncresolver.CachingConfigState  `json:"caching_config"`
	PipelineConfig   []appsyncresolver.PipelineConfigState `json:"pipeline_config"`
	Runtime          []appsyncresolver.RuntimeState        `json:"runtime"`
	SyncConfig       []appsyncresolver.SyncConfigState     `json:"sync_config"`
}
