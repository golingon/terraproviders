// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewGlueDevEndpoint(name string, args GlueDevEndpointArgs) *GlueDevEndpoint {
	return &GlueDevEndpoint{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*GlueDevEndpoint)(nil)

type GlueDevEndpoint struct {
	Name  string
	Args  GlueDevEndpointArgs
	state *glueDevEndpointState
}

func (gde *GlueDevEndpoint) Type() string {
	return "aws_glue_dev_endpoint"
}

func (gde *GlueDevEndpoint) LocalName() string {
	return gde.Name
}

func (gde *GlueDevEndpoint) Configuration() interface{} {
	return gde.Args
}

func (gde *GlueDevEndpoint) Attributes() glueDevEndpointAttributes {
	return glueDevEndpointAttributes{ref: terra.ReferenceResource(gde)}
}

func (gde *GlueDevEndpoint) ImportState(av io.Reader) error {
	gde.state = &glueDevEndpointState{}
	if err := json.NewDecoder(av).Decode(gde.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", gde.Type(), gde.LocalName(), err)
	}
	return nil
}

func (gde *GlueDevEndpoint) State() (*glueDevEndpointState, bool) {
	return gde.state, gde.state != nil
}

func (gde *GlueDevEndpoint) StateMust() *glueDevEndpointState {
	if gde.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", gde.Type(), gde.LocalName()))
	}
	return gde.state
}

func (gde *GlueDevEndpoint) DependOn() terra.Reference {
	return terra.ReferenceResource(gde)
}

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
	// DependsOn contains resources that GlueDevEndpoint depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type glueDevEndpointAttributes struct {
	ref terra.Reference
}

func (gde glueDevEndpointAttributes) Arguments() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](gde.ref.Append("arguments"))
}

func (gde glueDevEndpointAttributes) Arn() terra.StringValue {
	return terra.ReferenceString(gde.ref.Append("arn"))
}

func (gde glueDevEndpointAttributes) AvailabilityZone() terra.StringValue {
	return terra.ReferenceString(gde.ref.Append("availability_zone"))
}

func (gde glueDevEndpointAttributes) ExtraJarsS3Path() terra.StringValue {
	return terra.ReferenceString(gde.ref.Append("extra_jars_s3_path"))
}

func (gde glueDevEndpointAttributes) ExtraPythonLibsS3Path() terra.StringValue {
	return terra.ReferenceString(gde.ref.Append("extra_python_libs_s3_path"))
}

func (gde glueDevEndpointAttributes) FailureReason() terra.StringValue {
	return terra.ReferenceString(gde.ref.Append("failure_reason"))
}

func (gde glueDevEndpointAttributes) GlueVersion() terra.StringValue {
	return terra.ReferenceString(gde.ref.Append("glue_version"))
}

func (gde glueDevEndpointAttributes) Id() terra.StringValue {
	return terra.ReferenceString(gde.ref.Append("id"))
}

func (gde glueDevEndpointAttributes) Name() terra.StringValue {
	return terra.ReferenceString(gde.ref.Append("name"))
}

func (gde glueDevEndpointAttributes) NumberOfNodes() terra.NumberValue {
	return terra.ReferenceNumber(gde.ref.Append("number_of_nodes"))
}

func (gde glueDevEndpointAttributes) NumberOfWorkers() terra.NumberValue {
	return terra.ReferenceNumber(gde.ref.Append("number_of_workers"))
}

func (gde glueDevEndpointAttributes) PrivateAddress() terra.StringValue {
	return terra.ReferenceString(gde.ref.Append("private_address"))
}

func (gde glueDevEndpointAttributes) PublicAddress() terra.StringValue {
	return terra.ReferenceString(gde.ref.Append("public_address"))
}

func (gde glueDevEndpointAttributes) PublicKey() terra.StringValue {
	return terra.ReferenceString(gde.ref.Append("public_key"))
}

func (gde glueDevEndpointAttributes) PublicKeys() terra.SetValue[terra.StringValue] {
	return terra.ReferenceSet[terra.StringValue](gde.ref.Append("public_keys"))
}

func (gde glueDevEndpointAttributes) RoleArn() terra.StringValue {
	return terra.ReferenceString(gde.ref.Append("role_arn"))
}

func (gde glueDevEndpointAttributes) SecurityConfiguration() terra.StringValue {
	return terra.ReferenceString(gde.ref.Append("security_configuration"))
}

func (gde glueDevEndpointAttributes) SecurityGroupIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceSet[terra.StringValue](gde.ref.Append("security_group_ids"))
}

func (gde glueDevEndpointAttributes) Status() terra.StringValue {
	return terra.ReferenceString(gde.ref.Append("status"))
}

func (gde glueDevEndpointAttributes) SubnetId() terra.StringValue {
	return terra.ReferenceString(gde.ref.Append("subnet_id"))
}

func (gde glueDevEndpointAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](gde.ref.Append("tags"))
}

func (gde glueDevEndpointAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](gde.ref.Append("tags_all"))
}

func (gde glueDevEndpointAttributes) VpcId() terra.StringValue {
	return terra.ReferenceString(gde.ref.Append("vpc_id"))
}

func (gde glueDevEndpointAttributes) WorkerType() terra.StringValue {
	return terra.ReferenceString(gde.ref.Append("worker_type"))
}

func (gde glueDevEndpointAttributes) YarnEndpointAddress() terra.StringValue {
	return terra.ReferenceString(gde.ref.Append("yarn_endpoint_address"))
}

func (gde glueDevEndpointAttributes) ZeppelinRemoteSparkInterpreterPort() terra.NumberValue {
	return terra.ReferenceNumber(gde.ref.Append("zeppelin_remote_spark_interpreter_port"))
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
