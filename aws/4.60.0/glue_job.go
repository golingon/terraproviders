// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	gluejob "github.com/golingon/terraproviders/aws/4.60.0/gluejob"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewGlueJob(name string, args GlueJobArgs) *GlueJob {
	return &GlueJob{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*GlueJob)(nil)

type GlueJob struct {
	Name  string
	Args  GlueJobArgs
	state *glueJobState
}

func (gj *GlueJob) Type() string {
	return "aws_glue_job"
}

func (gj *GlueJob) LocalName() string {
	return gj.Name
}

func (gj *GlueJob) Configuration() interface{} {
	return gj.Args
}

func (gj *GlueJob) Attributes() glueJobAttributes {
	return glueJobAttributes{ref: terra.ReferenceResource(gj)}
}

func (gj *GlueJob) ImportState(av io.Reader) error {
	gj.state = &glueJobState{}
	if err := json.NewDecoder(av).Decode(gj.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", gj.Type(), gj.LocalName(), err)
	}
	return nil
}

func (gj *GlueJob) State() (*glueJobState, bool) {
	return gj.state, gj.state != nil
}

func (gj *GlueJob) StateMust() *glueJobState {
	if gj.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", gj.Type(), gj.LocalName()))
	}
	return gj.state
}

func (gj *GlueJob) DependOn() terra.Reference {
	return terra.ReferenceResource(gj)
}

type GlueJobArgs struct {
	// Connections: list of string, optional
	Connections terra.ListValue[terra.StringValue] `hcl:"connections,attr"`
	// DefaultArguments: map of string, optional
	DefaultArguments terra.MapValue[terra.StringValue] `hcl:"default_arguments,attr"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// ExecutionClass: string, optional
	ExecutionClass terra.StringValue `hcl:"execution_class,attr"`
	// GlueVersion: string, optional
	GlueVersion terra.StringValue `hcl:"glue_version,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// MaxCapacity: number, optional
	MaxCapacity terra.NumberValue `hcl:"max_capacity,attr"`
	// MaxRetries: number, optional
	MaxRetries terra.NumberValue `hcl:"max_retries,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// NonOverridableArguments: map of string, optional
	NonOverridableArguments terra.MapValue[terra.StringValue] `hcl:"non_overridable_arguments,attr"`
	// NumberOfWorkers: number, optional
	NumberOfWorkers terra.NumberValue `hcl:"number_of_workers,attr"`
	// RoleArn: string, required
	RoleArn terra.StringValue `hcl:"role_arn,attr" validate:"required"`
	// SecurityConfiguration: string, optional
	SecurityConfiguration terra.StringValue `hcl:"security_configuration,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// Timeout: number, optional
	Timeout terra.NumberValue `hcl:"timeout,attr"`
	// WorkerType: string, optional
	WorkerType terra.StringValue `hcl:"worker_type,attr"`
	// Command: required
	Command *gluejob.Command `hcl:"command,block" validate:"required"`
	// ExecutionProperty: optional
	ExecutionProperty *gluejob.ExecutionProperty `hcl:"execution_property,block"`
	// NotificationProperty: optional
	NotificationProperty *gluejob.NotificationProperty `hcl:"notification_property,block"`
	// DependsOn contains resources that GlueJob depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type glueJobAttributes struct {
	ref terra.Reference
}

func (gj glueJobAttributes) Arn() terra.StringValue {
	return terra.ReferenceString(gj.ref.Append("arn"))
}

func (gj glueJobAttributes) Connections() terra.ListValue[terra.StringValue] {
	return terra.ReferenceList[terra.StringValue](gj.ref.Append("connections"))
}

func (gj glueJobAttributes) DefaultArguments() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](gj.ref.Append("default_arguments"))
}

func (gj glueJobAttributes) Description() terra.StringValue {
	return terra.ReferenceString(gj.ref.Append("description"))
}

func (gj glueJobAttributes) ExecutionClass() terra.StringValue {
	return terra.ReferenceString(gj.ref.Append("execution_class"))
}

func (gj glueJobAttributes) GlueVersion() terra.StringValue {
	return terra.ReferenceString(gj.ref.Append("glue_version"))
}

func (gj glueJobAttributes) Id() terra.StringValue {
	return terra.ReferenceString(gj.ref.Append("id"))
}

func (gj glueJobAttributes) MaxCapacity() terra.NumberValue {
	return terra.ReferenceNumber(gj.ref.Append("max_capacity"))
}

func (gj glueJobAttributes) MaxRetries() terra.NumberValue {
	return terra.ReferenceNumber(gj.ref.Append("max_retries"))
}

func (gj glueJobAttributes) Name() terra.StringValue {
	return terra.ReferenceString(gj.ref.Append("name"))
}

func (gj glueJobAttributes) NonOverridableArguments() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](gj.ref.Append("non_overridable_arguments"))
}

func (gj glueJobAttributes) NumberOfWorkers() terra.NumberValue {
	return terra.ReferenceNumber(gj.ref.Append("number_of_workers"))
}

func (gj glueJobAttributes) RoleArn() terra.StringValue {
	return terra.ReferenceString(gj.ref.Append("role_arn"))
}

func (gj glueJobAttributes) SecurityConfiguration() terra.StringValue {
	return terra.ReferenceString(gj.ref.Append("security_configuration"))
}

func (gj glueJobAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](gj.ref.Append("tags"))
}

func (gj glueJobAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](gj.ref.Append("tags_all"))
}

func (gj glueJobAttributes) Timeout() terra.NumberValue {
	return terra.ReferenceNumber(gj.ref.Append("timeout"))
}

func (gj glueJobAttributes) WorkerType() terra.StringValue {
	return terra.ReferenceString(gj.ref.Append("worker_type"))
}

func (gj glueJobAttributes) Command() terra.ListValue[gluejob.CommandAttributes] {
	return terra.ReferenceList[gluejob.CommandAttributes](gj.ref.Append("command"))
}

func (gj glueJobAttributes) ExecutionProperty() terra.ListValue[gluejob.ExecutionPropertyAttributes] {
	return terra.ReferenceList[gluejob.ExecutionPropertyAttributes](gj.ref.Append("execution_property"))
}

func (gj glueJobAttributes) NotificationProperty() terra.ListValue[gluejob.NotificationPropertyAttributes] {
	return terra.ReferenceList[gluejob.NotificationPropertyAttributes](gj.ref.Append("notification_property"))
}

type glueJobState struct {
	Arn                     string                              `json:"arn"`
	Connections             []string                            `json:"connections"`
	DefaultArguments        map[string]string                   `json:"default_arguments"`
	Description             string                              `json:"description"`
	ExecutionClass          string                              `json:"execution_class"`
	GlueVersion             string                              `json:"glue_version"`
	Id                      string                              `json:"id"`
	MaxCapacity             float64                             `json:"max_capacity"`
	MaxRetries              float64                             `json:"max_retries"`
	Name                    string                              `json:"name"`
	NonOverridableArguments map[string]string                   `json:"non_overridable_arguments"`
	NumberOfWorkers         float64                             `json:"number_of_workers"`
	RoleArn                 string                              `json:"role_arn"`
	SecurityConfiguration   string                              `json:"security_configuration"`
	Tags                    map[string]string                   `json:"tags"`
	TagsAll                 map[string]string                   `json:"tags_all"`
	Timeout                 float64                             `json:"timeout"`
	WorkerType              string                              `json:"worker_type"`
	Command                 []gluejob.CommandState              `json:"command"`
	ExecutionProperty       []gluejob.ExecutionPropertyState    `json:"execution_property"`
	NotificationProperty    []gluejob.NotificationPropertyState `json:"notification_property"`
}
