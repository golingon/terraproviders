// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	datasynclocationhdfs "github.com/golingon/terraproviders/aws/5.7.0/datasynclocationhdfs"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDatasyncLocationHdfs creates a new instance of [DatasyncLocationHdfs].
func NewDatasyncLocationHdfs(name string, args DatasyncLocationHdfsArgs) *DatasyncLocationHdfs {
	return &DatasyncLocationHdfs{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DatasyncLocationHdfs)(nil)

// DatasyncLocationHdfs represents the Terraform resource aws_datasync_location_hdfs.
type DatasyncLocationHdfs struct {
	Name      string
	Args      DatasyncLocationHdfsArgs
	state     *datasyncLocationHdfsState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DatasyncLocationHdfs].
func (dlh *DatasyncLocationHdfs) Type() string {
	return "aws_datasync_location_hdfs"
}

// LocalName returns the local name for [DatasyncLocationHdfs].
func (dlh *DatasyncLocationHdfs) LocalName() string {
	return dlh.Name
}

// Configuration returns the configuration (args) for [DatasyncLocationHdfs].
func (dlh *DatasyncLocationHdfs) Configuration() interface{} {
	return dlh.Args
}

// DependOn is used for other resources to depend on [DatasyncLocationHdfs].
func (dlh *DatasyncLocationHdfs) DependOn() terra.Reference {
	return terra.ReferenceResource(dlh)
}

// Dependencies returns the list of resources [DatasyncLocationHdfs] depends_on.
func (dlh *DatasyncLocationHdfs) Dependencies() terra.Dependencies {
	return dlh.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DatasyncLocationHdfs].
func (dlh *DatasyncLocationHdfs) LifecycleManagement() *terra.Lifecycle {
	return dlh.Lifecycle
}

// Attributes returns the attributes for [DatasyncLocationHdfs].
func (dlh *DatasyncLocationHdfs) Attributes() datasyncLocationHdfsAttributes {
	return datasyncLocationHdfsAttributes{ref: terra.ReferenceResource(dlh)}
}

// ImportState imports the given attribute values into [DatasyncLocationHdfs]'s state.
func (dlh *DatasyncLocationHdfs) ImportState(av io.Reader) error {
	dlh.state = &datasyncLocationHdfsState{}
	if err := json.NewDecoder(av).Decode(dlh.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dlh.Type(), dlh.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DatasyncLocationHdfs] has state.
func (dlh *DatasyncLocationHdfs) State() (*datasyncLocationHdfsState, bool) {
	return dlh.state, dlh.state != nil
}

// StateMust returns the state for [DatasyncLocationHdfs]. Panics if the state is nil.
func (dlh *DatasyncLocationHdfs) StateMust() *datasyncLocationHdfsState {
	if dlh.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dlh.Type(), dlh.LocalName()))
	}
	return dlh.state
}

// DatasyncLocationHdfsArgs contains the configurations for aws_datasync_location_hdfs.
type DatasyncLocationHdfsArgs struct {
	// AgentArns: set of string, required
	AgentArns terra.SetValue[terra.StringValue] `hcl:"agent_arns,attr" validate:"required"`
	// AuthenticationType: string, optional
	AuthenticationType terra.StringValue `hcl:"authentication_type,attr"`
	// BlockSize: number, optional
	BlockSize terra.NumberValue `hcl:"block_size,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// KerberosKeytab: string, optional
	KerberosKeytab terra.StringValue `hcl:"kerberos_keytab,attr"`
	// KerberosKrb5Conf: string, optional
	KerberosKrb5Conf terra.StringValue `hcl:"kerberos_krb5_conf,attr"`
	// KerberosPrincipal: string, optional
	KerberosPrincipal terra.StringValue `hcl:"kerberos_principal,attr"`
	// KmsKeyProviderUri: string, optional
	KmsKeyProviderUri terra.StringValue `hcl:"kms_key_provider_uri,attr"`
	// ReplicationFactor: number, optional
	ReplicationFactor terra.NumberValue `hcl:"replication_factor,attr"`
	// SimpleUser: string, optional
	SimpleUser terra.StringValue `hcl:"simple_user,attr"`
	// Subdirectory: string, optional
	Subdirectory terra.StringValue `hcl:"subdirectory,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// NameNode: min=1
	NameNode []datasynclocationhdfs.NameNode `hcl:"name_node,block" validate:"min=1"`
	// QopConfiguration: optional
	QopConfiguration *datasynclocationhdfs.QopConfiguration `hcl:"qop_configuration,block"`
}
type datasyncLocationHdfsAttributes struct {
	ref terra.Reference
}

// AgentArns returns a reference to field agent_arns of aws_datasync_location_hdfs.
func (dlh datasyncLocationHdfsAttributes) AgentArns() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](dlh.ref.Append("agent_arns"))
}

// Arn returns a reference to field arn of aws_datasync_location_hdfs.
func (dlh datasyncLocationHdfsAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(dlh.ref.Append("arn"))
}

// AuthenticationType returns a reference to field authentication_type of aws_datasync_location_hdfs.
func (dlh datasyncLocationHdfsAttributes) AuthenticationType() terra.StringValue {
	return terra.ReferenceAsString(dlh.ref.Append("authentication_type"))
}

// BlockSize returns a reference to field block_size of aws_datasync_location_hdfs.
func (dlh datasyncLocationHdfsAttributes) BlockSize() terra.NumberValue {
	return terra.ReferenceAsNumber(dlh.ref.Append("block_size"))
}

// Id returns a reference to field id of aws_datasync_location_hdfs.
func (dlh datasyncLocationHdfsAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dlh.ref.Append("id"))
}

// KerberosKeytab returns a reference to field kerberos_keytab of aws_datasync_location_hdfs.
func (dlh datasyncLocationHdfsAttributes) KerberosKeytab() terra.StringValue {
	return terra.ReferenceAsString(dlh.ref.Append("kerberos_keytab"))
}

// KerberosKrb5Conf returns a reference to field kerberos_krb5_conf of aws_datasync_location_hdfs.
func (dlh datasyncLocationHdfsAttributes) KerberosKrb5Conf() terra.StringValue {
	return terra.ReferenceAsString(dlh.ref.Append("kerberos_krb5_conf"))
}

// KerberosPrincipal returns a reference to field kerberos_principal of aws_datasync_location_hdfs.
func (dlh datasyncLocationHdfsAttributes) KerberosPrincipal() terra.StringValue {
	return terra.ReferenceAsString(dlh.ref.Append("kerberos_principal"))
}

// KmsKeyProviderUri returns a reference to field kms_key_provider_uri of aws_datasync_location_hdfs.
func (dlh datasyncLocationHdfsAttributes) KmsKeyProviderUri() terra.StringValue {
	return terra.ReferenceAsString(dlh.ref.Append("kms_key_provider_uri"))
}

// ReplicationFactor returns a reference to field replication_factor of aws_datasync_location_hdfs.
func (dlh datasyncLocationHdfsAttributes) ReplicationFactor() terra.NumberValue {
	return terra.ReferenceAsNumber(dlh.ref.Append("replication_factor"))
}

// SimpleUser returns a reference to field simple_user of aws_datasync_location_hdfs.
func (dlh datasyncLocationHdfsAttributes) SimpleUser() terra.StringValue {
	return terra.ReferenceAsString(dlh.ref.Append("simple_user"))
}

// Subdirectory returns a reference to field subdirectory of aws_datasync_location_hdfs.
func (dlh datasyncLocationHdfsAttributes) Subdirectory() terra.StringValue {
	return terra.ReferenceAsString(dlh.ref.Append("subdirectory"))
}

// Tags returns a reference to field tags of aws_datasync_location_hdfs.
func (dlh datasyncLocationHdfsAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dlh.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_datasync_location_hdfs.
func (dlh datasyncLocationHdfsAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dlh.ref.Append("tags_all"))
}

// Uri returns a reference to field uri of aws_datasync_location_hdfs.
func (dlh datasyncLocationHdfsAttributes) Uri() terra.StringValue {
	return terra.ReferenceAsString(dlh.ref.Append("uri"))
}

func (dlh datasyncLocationHdfsAttributes) NameNode() terra.SetValue[datasynclocationhdfs.NameNodeAttributes] {
	return terra.ReferenceAsSet[datasynclocationhdfs.NameNodeAttributes](dlh.ref.Append("name_node"))
}

func (dlh datasyncLocationHdfsAttributes) QopConfiguration() terra.ListValue[datasynclocationhdfs.QopConfigurationAttributes] {
	return terra.ReferenceAsList[datasynclocationhdfs.QopConfigurationAttributes](dlh.ref.Append("qop_configuration"))
}

type datasyncLocationHdfsState struct {
	AgentArns          []string                                     `json:"agent_arns"`
	Arn                string                                       `json:"arn"`
	AuthenticationType string                                       `json:"authentication_type"`
	BlockSize          float64                                      `json:"block_size"`
	Id                 string                                       `json:"id"`
	KerberosKeytab     string                                       `json:"kerberos_keytab"`
	KerberosKrb5Conf   string                                       `json:"kerberos_krb5_conf"`
	KerberosPrincipal  string                                       `json:"kerberos_principal"`
	KmsKeyProviderUri  string                                       `json:"kms_key_provider_uri"`
	ReplicationFactor  float64                                      `json:"replication_factor"`
	SimpleUser         string                                       `json:"simple_user"`
	Subdirectory       string                                       `json:"subdirectory"`
	Tags               map[string]string                            `json:"tags"`
	TagsAll            map[string]string                            `json:"tags_all"`
	Uri                string                                       `json:"uri"`
	NameNode           []datasynclocationhdfs.NameNodeState         `json:"name_node"`
	QopConfiguration   []datasynclocationhdfs.QopConfigurationState `json:"qop_configuration"`
}
