// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewGlueDevEndpoint creates a new instance of [GlueDevEndpoint].
func NewGlueDevEndpoint(name string, args GlueDevEndpointArgs) *GlueDevEndpoint {
	return &GlueDevEndpoint{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*GlueDevEndpoint)(nil)

// GlueDevEndpoint represents the Terraform resource aws_glue_dev_endpoint.
type GlueDevEndpoint struct {
	Name      string
	Args      GlueDevEndpointArgs
	state     *glueDevEndpointState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [GlueDevEndpoint].
func (gde *GlueDevEndpoint) Type() string {
	return "aws_glue_dev_endpoint"
}

// LocalName returns the local name for [GlueDevEndpoint].
func (gde *GlueDevEndpoint) LocalName() string {
	return gde.Name
}

// Configuration returns the configuration (args) for [GlueDevEndpoint].
func (gde *GlueDevEndpoint) Configuration() interface{} {
	return gde.Args
}

// DependOn is used for other resources to depend on [GlueDevEndpoint].
func (gde *GlueDevEndpoint) DependOn() terra.Reference {
	return terra.ReferenceResource(gde)
}

// Dependencies returns the list of resources [GlueDevEndpoint] depends_on.
func (gde *GlueDevEndpoint) Dependencies() terra.Dependencies {
	return gde.DependsOn
}

// LifecycleManagement returns the lifecycle block for [GlueDevEndpoint].
func (gde *GlueDevEndpoint) LifecycleManagement() *terra.Lifecycle {
	return gde.Lifecycle
}

// Attributes returns the attributes for [GlueDevEndpoint].
func (gde *GlueDevEndpoint) Attributes() glueDevEndpointAttributes {
	return glueDevEndpointAttributes{ref: terra.ReferenceResource(gde)}
}

// ImportState imports the given attribute values into [GlueDevEndpoint]'s state.
func (gde *GlueDevEndpoint) ImportState(av io.Reader) error {
	gde.state = &glueDevEndpointState{}
	if err := json.NewDecoder(av).Decode(gde.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", gde.Type(), gde.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [GlueDevEndpoint] has state.
func (gde *GlueDevEndpoint) State() (*glueDevEndpointState, bool) {
	return gde.state, gde.state != nil
}

// StateMust returns the state for [GlueDevEndpoint]. Panics if the state is nil.
func (gde *GlueDevEndpoint) StateMust() *glueDevEndpointState {
	if gde.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", gde.Type(), gde.LocalName()))
	}
	return gde.state
}

// GlueDevEndpointArgs contains the configurations for aws_glue_dev_endpoint.
type GlueDevEndpointArgs struct {
	// Arguments: map of string, optional
	Arguments terra.MapValue[terra.StringValue] `hcl:"arguments,attr"`
	// ExtraJarsS3Path: string, optional
	ExtraJarsS3Path terra.StringValue `hcl:"extra_jars_s3_path,attr"`
	// ExtraPythonLibsS3Path: string, optional
	ExtraPythonLibsS3Path terra.StringValue `hcl:"extra_python_libs_s3_path,attr"`
	// GlueVersion: string, optional
	GlueVersion terra.StringValue `hcl:"glue_version,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// NumberOfNodes: number, optional
	NumberOfNodes terra.NumberValue `hcl:"number_of_nodes,attr"`
	// NumberOfWorkers: number, optional
	NumberOfWorkers terra.NumberValue `hcl:"number_of_workers,attr"`
	// PublicKey: string, optional
	PublicKey terra.StringValue `hcl:"public_key,attr"`
	// PublicKeys: set of string, optional
	PublicKeys terra.SetValue[terra.StringValue] `hcl:"public_keys,attr"`
	// RoleArn: string, required
	RoleArn terra.StringValue `hcl:"role_arn,attr" validate:"required"`
	// SecurityConfiguration: string, optional
	SecurityConfiguration terra.StringValue `hcl:"security_configuration,attr"`
	// SecurityGroupIds: set of string, optional
	SecurityGroupIds terra.SetValue[terra.StringValue] `hcl:"security_group_ids,attr"`
	// SubnetId: string, optional
	SubnetId terra.StringValue `hcl:"subnet_id,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// WorkerType: string, optional
	WorkerType terra.StringValue `hcl:"worker_type,attr"`
}
type glueDevEndpointAttributes struct {
	ref terra.Reference
}

// Arguments returns a reference to field arguments of aws_glue_dev_endpoint.
func (gde glueDevEndpointAttributes) Arguments() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](gde.ref.Append("arguments"))
}

// Arn returns a reference to field arn of aws_glue_dev_endpoint.
func (gde glueDevEndpointAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(gde.ref.Append("arn"))
}

// AvailabilityZone returns a reference to field availability_zone of aws_glue_dev_endpoint.
func (gde glueDevEndpointAttributes) AvailabilityZone() terra.StringValue {
	return terra.ReferenceAsString(gde.ref.Append("availability_zone"))
}

// ExtraJarsS3Path returns a reference to field extra_jars_s3_path of aws_glue_dev_endpoint.
func (gde glueDevEndpointAttributes) ExtraJarsS3Path() terra.StringValue {
	return terra.ReferenceAsString(gde.ref.Append("extra_jars_s3_path"))
}

// ExtraPythonLibsS3Path returns a reference to field extra_python_libs_s3_path of aws_glue_dev_endpoint.
func (gde glueDevEndpointAttributes) ExtraPythonLibsS3Path() terra.StringValue {
	return terra.ReferenceAsString(gde.ref.Append("extra_python_libs_s3_path"))
}

// FailureReason returns a reference to field failure_reason of aws_glue_dev_endpoint.
func (gde glueDevEndpointAttributes) FailureReason() terra.StringValue {
	return terra.ReferenceAsString(gde.ref.Append("failure_reason"))
}

// GlueVersion returns a reference to field glue_version of aws_glue_dev_endpoint.
func (gde glueDevEndpointAttributes) GlueVersion() terra.StringValue {
	return terra.ReferenceAsString(gde.ref.Append("glue_version"))
}

// Id returns a reference to field id of aws_glue_dev_endpoint.
func (gde glueDevEndpointAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gde.ref.Append("id"))
}

// Name returns a reference to field name of aws_glue_dev_endpoint.
func (gde glueDevEndpointAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(gde.ref.Append("name"))
}

// NumberOfNodes returns a reference to field number_of_nodes of aws_glue_dev_endpoint.
func (gde glueDevEndpointAttributes) NumberOfNodes() terra.NumberValue {
	return terra.ReferenceAsNumber(gde.ref.Append("number_of_nodes"))
}

// NumberOfWorkers returns a reference to field number_of_workers of aws_glue_dev_endpoint.
func (gde glueDevEndpointAttributes) NumberOfWorkers() terra.NumberValue {
	return terra.ReferenceAsNumber(gde.ref.Append("number_of_workers"))
}

// PrivateAddress returns a reference to field private_address of aws_glue_dev_endpoint.
func (gde glueDevEndpointAttributes) PrivateAddress() terra.StringValue {
	return terra.ReferenceAsString(gde.ref.Append("private_address"))
}

// PublicAddress returns a reference to field public_address of aws_glue_dev_endpoint.
func (gde glueDevEndpointAttributes) PublicAddress() terra.StringValue {
	return terra.ReferenceAsString(gde.ref.Append("public_address"))
}

// PublicKey returns a reference to field public_key of aws_glue_dev_endpoint.
func (gde glueDevEndpointAttributes) PublicKey() terra.StringValue {
	return terra.ReferenceAsString(gde.ref.Append("public_key"))
}

// PublicKeys returns a reference to field public_keys of aws_glue_dev_endpoint.
func (gde glueDevEndpointAttributes) PublicKeys() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](gde.ref.Append("public_keys"))
}

// RoleArn returns a reference to field role_arn of aws_glue_dev_endpoint.
func (gde glueDevEndpointAttributes) RoleArn() terra.StringValue {
	return terra.ReferenceAsString(gde.ref.Append("role_arn"))
}

// SecurityConfiguration returns a reference to field security_configuration of aws_glue_dev_endpoint.
func (gde glueDevEndpointAttributes) SecurityConfiguration() terra.StringValue {
	return terra.ReferenceAsString(gde.ref.Append("security_configuration"))
}

// SecurityGroupIds returns a reference to field security_group_ids of aws_glue_dev_endpoint.
func (gde glueDevEndpointAttributes) SecurityGroupIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](gde.ref.Append("security_group_ids"))
}

// Status returns a reference to field status of aws_glue_dev_endpoint.
func (gde glueDevEndpointAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(gde.ref.Append("status"))
}

// SubnetId returns a reference to field subnet_id of aws_glue_dev_endpoint.
func (gde glueDevEndpointAttributes) SubnetId() terra.StringValue {
	return terra.ReferenceAsString(gde.ref.Append("subnet_id"))
}

// Tags returns a reference to field tags of aws_glue_dev_endpoint.
func (gde glueDevEndpointAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](gde.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_glue_dev_endpoint.
func (gde glueDevEndpointAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](gde.ref.Append("tags_all"))
}

// VpcId returns a reference to field vpc_id of aws_glue_dev_endpoint.
func (gde glueDevEndpointAttributes) VpcId() terra.StringValue {
	return terra.ReferenceAsString(gde.ref.Append("vpc_id"))
}

// WorkerType returns a reference to field worker_type of aws_glue_dev_endpoint.
func (gde glueDevEndpointAttributes) WorkerType() terra.StringValue {
	return terra.ReferenceAsString(gde.ref.Append("worker_type"))
}

// YarnEndpointAddress returns a reference to field yarn_endpoint_address of aws_glue_dev_endpoint.
func (gde glueDevEndpointAttributes) YarnEndpointAddress() terra.StringValue {
	return terra.ReferenceAsString(gde.ref.Append("yarn_endpoint_address"))
}

// ZeppelinRemoteSparkInterpreterPort returns a reference to field zeppelin_remote_spark_interpreter_port of aws_glue_dev_endpoint.
func (gde glueDevEndpointAttributes) ZeppelinRemoteSparkInterpreterPort() terra.NumberValue {
	return terra.ReferenceAsNumber(gde.ref.Append("zeppelin_remote_spark_interpreter_port"))
}

type glueDevEndpointState struct {
	Arguments                          map[string]string `json:"arguments"`
	Arn                                string            `json:"arn"`
	AvailabilityZone                   string            `json:"availability_zone"`
	ExtraJarsS3Path                    string            `json:"extra_jars_s3_path"`
	ExtraPythonLibsS3Path              string            `json:"extra_python_libs_s3_path"`
	FailureReason                      string            `json:"failure_reason"`
	GlueVersion                        string            `json:"glue_version"`
	Id                                 string            `json:"id"`
	Name                               string            `json:"name"`
	NumberOfNodes                      float64           `json:"number_of_nodes"`
	NumberOfWorkers                    float64           `json:"number_of_workers"`
	PrivateAddress                     string            `json:"private_address"`
	PublicAddress                      string            `json:"public_address"`
	PublicKey                          string            `json:"public_key"`
	PublicKeys                         []string          `json:"public_keys"`
	RoleArn                            string            `json:"role_arn"`
	SecurityConfiguration              string            `json:"security_configuration"`
	SecurityGroupIds                   []string          `json:"security_group_ids"`
	Status                             string            `json:"status"`
	SubnetId                           string            `json:"subnet_id"`
	Tags                               map[string]string `json:"tags"`
	TagsAll                            map[string]string `json:"tags_all"`
	VpcId                              string            `json:"vpc_id"`
	WorkerType                         string            `json:"worker_type"`
	YarnEndpointAddress                string            `json:"yarn_endpoint_address"`
	ZeppelinRemoteSparkInterpreterPort float64           `json:"zeppelin_remote_spark_interpreter_port"`
}
