// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// NewDatasyncLocationObjectStorage creates a new instance of [DatasyncLocationObjectStorage].
func NewDatasyncLocationObjectStorage(name string, args DatasyncLocationObjectStorageArgs) *DatasyncLocationObjectStorage {
	return &DatasyncLocationObjectStorage{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DatasyncLocationObjectStorage)(nil)

// DatasyncLocationObjectStorage represents the Terraform resource aws_datasync_location_object_storage.
type DatasyncLocationObjectStorage struct {
	Name      string
	Args      DatasyncLocationObjectStorageArgs
	state     *datasyncLocationObjectStorageState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DatasyncLocationObjectStorage].
func (dlos *DatasyncLocationObjectStorage) Type() string {
	return "aws_datasync_location_object_storage"
}

// LocalName returns the local name for [DatasyncLocationObjectStorage].
func (dlos *DatasyncLocationObjectStorage) LocalName() string {
	return dlos.Name
}

// Configuration returns the configuration (args) for [DatasyncLocationObjectStorage].
func (dlos *DatasyncLocationObjectStorage) Configuration() interface{} {
	return dlos.Args
}

// DependOn is used for other resources to depend on [DatasyncLocationObjectStorage].
func (dlos *DatasyncLocationObjectStorage) DependOn() terra.Reference {
	return terra.ReferenceResource(dlos)
}

// Dependencies returns the list of resources [DatasyncLocationObjectStorage] depends_on.
func (dlos *DatasyncLocationObjectStorage) Dependencies() terra.Dependencies {
	return dlos.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DatasyncLocationObjectStorage].
func (dlos *DatasyncLocationObjectStorage) LifecycleManagement() *terra.Lifecycle {
	return dlos.Lifecycle
}

// Attributes returns the attributes for [DatasyncLocationObjectStorage].
func (dlos *DatasyncLocationObjectStorage) Attributes() datasyncLocationObjectStorageAttributes {
	return datasyncLocationObjectStorageAttributes{ref: terra.ReferenceResource(dlos)}
}

// ImportState imports the given attribute values into [DatasyncLocationObjectStorage]'s state.
func (dlos *DatasyncLocationObjectStorage) ImportState(av io.Reader) error {
	dlos.state = &datasyncLocationObjectStorageState{}
	if err := json.NewDecoder(av).Decode(dlos.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dlos.Type(), dlos.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DatasyncLocationObjectStorage] has state.
func (dlos *DatasyncLocationObjectStorage) State() (*datasyncLocationObjectStorageState, bool) {
	return dlos.state, dlos.state != nil
}

// StateMust returns the state for [DatasyncLocationObjectStorage]. Panics if the state is nil.
func (dlos *DatasyncLocationObjectStorage) StateMust() *datasyncLocationObjectStorageState {
	if dlos.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dlos.Type(), dlos.LocalName()))
	}
	return dlos.state
}

// DatasyncLocationObjectStorageArgs contains the configurations for aws_datasync_location_object_storage.
type DatasyncLocationObjectStorageArgs struct {
	// AccessKey: string, optional
	AccessKey terra.StringValue `hcl:"access_key,attr"`
	// AgentArns: set of string, required
	AgentArns terra.SetValue[terra.StringValue] `hcl:"agent_arns,attr" validate:"required"`
	// BucketName: string, required
	BucketName terra.StringValue `hcl:"bucket_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// SecretKey: string, optional
	SecretKey terra.StringValue `hcl:"secret_key,attr"`
	// ServerCertificate: string, optional
	ServerCertificate terra.StringValue `hcl:"server_certificate,attr"`
	// ServerHostname: string, required
	ServerHostname terra.StringValue `hcl:"server_hostname,attr" validate:"required"`
	// ServerPort: number, optional
	ServerPort terra.NumberValue `hcl:"server_port,attr"`
	// ServerProtocol: string, optional
	ServerProtocol terra.StringValue `hcl:"server_protocol,attr"`
	// Subdirectory: string, optional
	Subdirectory terra.StringValue `hcl:"subdirectory,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
}
type datasyncLocationObjectStorageAttributes struct {
	ref terra.Reference
}

// AccessKey returns a reference to field access_key of aws_datasync_location_object_storage.
func (dlos datasyncLocationObjectStorageAttributes) AccessKey() terra.StringValue {
	return terra.ReferenceAsString(dlos.ref.Append("access_key"))
}

// AgentArns returns a reference to field agent_arns of aws_datasync_location_object_storage.
func (dlos datasyncLocationObjectStorageAttributes) AgentArns() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](dlos.ref.Append("agent_arns"))
}

// Arn returns a reference to field arn of aws_datasync_location_object_storage.
func (dlos datasyncLocationObjectStorageAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(dlos.ref.Append("arn"))
}

// BucketName returns a reference to field bucket_name of aws_datasync_location_object_storage.
func (dlos datasyncLocationObjectStorageAttributes) BucketName() terra.StringValue {
	return terra.ReferenceAsString(dlos.ref.Append("bucket_name"))
}

// Id returns a reference to field id of aws_datasync_location_object_storage.
func (dlos datasyncLocationObjectStorageAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dlos.ref.Append("id"))
}

// SecretKey returns a reference to field secret_key of aws_datasync_location_object_storage.
func (dlos datasyncLocationObjectStorageAttributes) SecretKey() terra.StringValue {
	return terra.ReferenceAsString(dlos.ref.Append("secret_key"))
}

// ServerCertificate returns a reference to field server_certificate of aws_datasync_location_object_storage.
func (dlos datasyncLocationObjectStorageAttributes) ServerCertificate() terra.StringValue {
	return terra.ReferenceAsString(dlos.ref.Append("server_certificate"))
}

// ServerHostname returns a reference to field server_hostname of aws_datasync_location_object_storage.
func (dlos datasyncLocationObjectStorageAttributes) ServerHostname() terra.StringValue {
	return terra.ReferenceAsString(dlos.ref.Append("server_hostname"))
}

// ServerPort returns a reference to field server_port of aws_datasync_location_object_storage.
func (dlos datasyncLocationObjectStorageAttributes) ServerPort() terra.NumberValue {
	return terra.ReferenceAsNumber(dlos.ref.Append("server_port"))
}

// ServerProtocol returns a reference to field server_protocol of aws_datasync_location_object_storage.
func (dlos datasyncLocationObjectStorageAttributes) ServerProtocol() terra.StringValue {
	return terra.ReferenceAsString(dlos.ref.Append("server_protocol"))
}

// Subdirectory returns a reference to field subdirectory of aws_datasync_location_object_storage.
func (dlos datasyncLocationObjectStorageAttributes) Subdirectory() terra.StringValue {
	return terra.ReferenceAsString(dlos.ref.Append("subdirectory"))
}

// Tags returns a reference to field tags of aws_datasync_location_object_storage.
func (dlos datasyncLocationObjectStorageAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dlos.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_datasync_location_object_storage.
func (dlos datasyncLocationObjectStorageAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dlos.ref.Append("tags_all"))
}

// Uri returns a reference to field uri of aws_datasync_location_object_storage.
func (dlos datasyncLocationObjectStorageAttributes) Uri() terra.StringValue {
	return terra.ReferenceAsString(dlos.ref.Append("uri"))
}

type datasyncLocationObjectStorageState struct {
	AccessKey         string            `json:"access_key"`
	AgentArns         []string          `json:"agent_arns"`
	Arn               string            `json:"arn"`
	BucketName        string            `json:"bucket_name"`
	Id                string            `json:"id"`
	SecretKey         string            `json:"secret_key"`
	ServerCertificate string            `json:"server_certificate"`
	ServerHostname    string            `json:"server_hostname"`
	ServerPort        float64           `json:"server_port"`
	ServerProtocol    string            `json:"server_protocol"`
	Subdirectory      string            `json:"subdirectory"`
	Tags              map[string]string `json:"tags"`
	TagsAll           map[string]string `json:"tags_all"`
	Uri               string            `json:"uri"`
}
