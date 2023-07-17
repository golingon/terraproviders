// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	appsyncfunction "github.com/golingon/terraproviders/aws/5.8.0/appsyncfunction"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewAppsyncFunction creates a new instance of [AppsyncFunction].
func NewAppsyncFunction(name string, args AppsyncFunctionArgs) *AppsyncFunction {
	return &AppsyncFunction{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*AppsyncFunction)(nil)

// AppsyncFunction represents the Terraform resource aws_appsync_function.
type AppsyncFunction struct {
	Name      string
	Args      AppsyncFunctionArgs
	state     *appsyncFunctionState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [AppsyncFunction].
func (af *AppsyncFunction) Type() string {
	return "aws_appsync_function"
}

// LocalName returns the local name for [AppsyncFunction].
func (af *AppsyncFunction) LocalName() string {
	return af.Name
}

// Configuration returns the configuration (args) for [AppsyncFunction].
func (af *AppsyncFunction) Configuration() interface{} {
	return af.Args
}

// DependOn is used for other resources to depend on [AppsyncFunction].
func (af *AppsyncFunction) DependOn() terra.Reference {
	return terra.ReferenceResource(af)
}

// Dependencies returns the list of resources [AppsyncFunction] depends_on.
func (af *AppsyncFunction) Dependencies() terra.Dependencies {
	return af.DependsOn
}

// LifecycleManagement returns the lifecycle block for [AppsyncFunction].
func (af *AppsyncFunction) LifecycleManagement() *terra.Lifecycle {
	return af.Lifecycle
}

// Attributes returns the attributes for [AppsyncFunction].
func (af *AppsyncFunction) Attributes() appsyncFunctionAttributes {
	return appsyncFunctionAttributes{ref: terra.ReferenceResource(af)}
}

// ImportState imports the given attribute values into [AppsyncFunction]'s state.
func (af *AppsyncFunction) ImportState(av io.Reader) error {
	af.state = &appsyncFunctionState{}
	if err := json.NewDecoder(av).Decode(af.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", af.Type(), af.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [AppsyncFunction] has state.
func (af *AppsyncFunction) State() (*appsyncFunctionState, bool) {
	return af.state, af.state != nil
}

// StateMust returns the state for [AppsyncFunction]. Panics if the state is nil.
func (af *AppsyncFunction) StateMust() *appsyncFunctionState {
	if af.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", af.Type(), af.LocalName()))
	}
	return af.state
}

// AppsyncFunctionArgs contains the configurations for aws_appsync_function.
type AppsyncFunctionArgs struct {
	// ApiId: string, required
	ApiId terra.StringValue `hcl:"api_id,attr" validate:"required"`
	// Code: string, optional
	Code terra.StringValue `hcl:"code,attr"`
	// DataSource: string, required
	DataSource terra.StringValue `hcl:"data_source,attr" validate:"required"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// FunctionVersion: string, optional
	FunctionVersion terra.StringValue `hcl:"function_version,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// MaxBatchSize: number, optional
	MaxBatchSize terra.NumberValue `hcl:"max_batch_size,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// RequestMappingTemplate: string, optional
	RequestMappingTemplate terra.StringValue `hcl:"request_mapping_template,attr"`
	// ResponseMappingTemplate: string, optional
	ResponseMappingTemplate terra.StringValue `hcl:"response_mapping_template,attr"`
	// Runtime: optional
	Runtime *appsyncfunction.Runtime `hcl:"runtime,block"`
	// SyncConfig: optional
	SyncConfig *appsyncfunction.SyncConfig `hcl:"sync_config,block"`
}
type appsyncFunctionAttributes struct {
	ref terra.Reference
}

// ApiId returns a reference to field api_id of aws_appsync_function.
func (af appsyncFunctionAttributes) ApiId() terra.StringValue {
	return terra.ReferenceAsString(af.ref.Append("api_id"))
}

// Arn returns a reference to field arn of aws_appsync_function.
func (af appsyncFunctionAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(af.ref.Append("arn"))
}

// Code returns a reference to field code of aws_appsync_function.
func (af appsyncFunctionAttributes) Code() terra.StringValue {
	return terra.ReferenceAsString(af.ref.Append("code"))
}

// DataSource returns a reference to field data_source of aws_appsync_function.
func (af appsyncFunctionAttributes) DataSource() terra.StringValue {
	return terra.ReferenceAsString(af.ref.Append("data_source"))
}

// Description returns a reference to field description of aws_appsync_function.
func (af appsyncFunctionAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(af.ref.Append("description"))
}

// FunctionId returns a reference to field function_id of aws_appsync_function.
func (af appsyncFunctionAttributes) FunctionId() terra.StringValue {
	return terra.ReferenceAsString(af.ref.Append("function_id"))
}

// FunctionVersion returns a reference to field function_version of aws_appsync_function.
func (af appsyncFunctionAttributes) FunctionVersion() terra.StringValue {
	return terra.ReferenceAsString(af.ref.Append("function_version"))
}

// Id returns a reference to field id of aws_appsync_function.
func (af appsyncFunctionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(af.ref.Append("id"))
}

// MaxBatchSize returns a reference to field max_batch_size of aws_appsync_function.
func (af appsyncFunctionAttributes) MaxBatchSize() terra.NumberValue {
	return terra.ReferenceAsNumber(af.ref.Append("max_batch_size"))
}

// Name returns a reference to field name of aws_appsync_function.
func (af appsyncFunctionAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(af.ref.Append("name"))
}

// RequestMappingTemplate returns a reference to field request_mapping_template of aws_appsync_function.
func (af appsyncFunctionAttributes) RequestMappingTemplate() terra.StringValue {
	return terra.ReferenceAsString(af.ref.Append("request_mapping_template"))
}

// ResponseMappingTemplate returns a reference to field response_mapping_template of aws_appsync_function.
func (af appsyncFunctionAttributes) ResponseMappingTemplate() terra.StringValue {
	return terra.ReferenceAsString(af.ref.Append("response_mapping_template"))
}

func (af appsyncFunctionAttributes) Runtime() terra.ListValue[appsyncfunction.RuntimeAttributes] {
	return terra.ReferenceAsList[appsyncfunction.RuntimeAttributes](af.ref.Append("runtime"))
}

func (af appsyncFunctionAttributes) SyncConfig() terra.ListValue[appsyncfunction.SyncConfigAttributes] {
	return terra.ReferenceAsList[appsyncfunction.SyncConfigAttributes](af.ref.Append("sync_config"))
}

type appsyncFunctionState struct {
	ApiId                   string                            `json:"api_id"`
	Arn                     string                            `json:"arn"`
	Code                    string                            `json:"code"`
	DataSource              string                            `json:"data_source"`
	Description             string                            `json:"description"`
	FunctionId              string                            `json:"function_id"`
	FunctionVersion         string                            `json:"function_version"`
	Id                      string                            `json:"id"`
	MaxBatchSize            float64                           `json:"max_batch_size"`
	Name                    string                            `json:"name"`
	RequestMappingTemplate  string                            `json:"request_mapping_template"`
	ResponseMappingTemplate string                            `json:"response_mapping_template"`
	Runtime                 []appsyncfunction.RuntimeState    `json:"runtime"`
	SyncConfig              []appsyncfunction.SyncConfigState `json:"sync_config"`
}
