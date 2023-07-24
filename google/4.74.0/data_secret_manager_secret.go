// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	datasecretmanagersecret "github.com/golingon/terraproviders/google/4.74.0/datasecretmanagersecret"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataSecretManagerSecret creates a new instance of [DataSecretManagerSecret].
func NewDataSecretManagerSecret(name string, args DataSecretManagerSecretArgs) *DataSecretManagerSecret {
	return &DataSecretManagerSecret{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataSecretManagerSecret)(nil)

// DataSecretManagerSecret represents the Terraform data resource google_secret_manager_secret.
type DataSecretManagerSecret struct {
	Name string
	Args DataSecretManagerSecretArgs
}

// DataSource returns the Terraform object type for [DataSecretManagerSecret].
func (sms *DataSecretManagerSecret) DataSource() string {
	return "google_secret_manager_secret"
}

// LocalName returns the local name for [DataSecretManagerSecret].
func (sms *DataSecretManagerSecret) LocalName() string {
	return sms.Name
}

// Configuration returns the configuration (args) for [DataSecretManagerSecret].
func (sms *DataSecretManagerSecret) Configuration() interface{} {
	return sms.Args
}

// Attributes returns the attributes for [DataSecretManagerSecret].
func (sms *DataSecretManagerSecret) Attributes() dataSecretManagerSecretAttributes {
	return dataSecretManagerSecretAttributes{ref: terra.ReferenceDataResource(sms)}
}

// DataSecretManagerSecretArgs contains the configurations for google_secret_manager_secret.
type DataSecretManagerSecretArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// SecretId: string, required
	SecretId terra.StringValue `hcl:"secret_id,attr" validate:"required"`
	// Replication: min=0
	Replication []datasecretmanagersecret.Replication `hcl:"replication,block" validate:"min=0"`
	// Rotation: min=0
	Rotation []datasecretmanagersecret.Rotation `hcl:"rotation,block" validate:"min=0"`
	// Topics: min=0
	Topics []datasecretmanagersecret.Topics `hcl:"topics,block" validate:"min=0"`
}
type dataSecretManagerSecretAttributes struct {
	ref terra.Reference
}

// CreateTime returns a reference to field create_time of google_secret_manager_secret.
func (sms dataSecretManagerSecretAttributes) CreateTime() terra.StringValue {
	return terra.ReferenceAsString(sms.ref.Append("create_time"))
}

// ExpireTime returns a reference to field expire_time of google_secret_manager_secret.
func (sms dataSecretManagerSecretAttributes) ExpireTime() terra.StringValue {
	return terra.ReferenceAsString(sms.ref.Append("expire_time"))
}

// Id returns a reference to field id of google_secret_manager_secret.
func (sms dataSecretManagerSecretAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sms.ref.Append("id"))
}

// Labels returns a reference to field labels of google_secret_manager_secret.
func (sms dataSecretManagerSecretAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](sms.ref.Append("labels"))
}

// Name returns a reference to field name of google_secret_manager_secret.
func (sms dataSecretManagerSecretAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(sms.ref.Append("name"))
}

// Project returns a reference to field project of google_secret_manager_secret.
func (sms dataSecretManagerSecretAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(sms.ref.Append("project"))
}

// SecretId returns a reference to field secret_id of google_secret_manager_secret.
func (sms dataSecretManagerSecretAttributes) SecretId() terra.StringValue {
	return terra.ReferenceAsString(sms.ref.Append("secret_id"))
}

// Ttl returns a reference to field ttl of google_secret_manager_secret.
func (sms dataSecretManagerSecretAttributes) Ttl() terra.StringValue {
	return terra.ReferenceAsString(sms.ref.Append("ttl"))
}

func (sms dataSecretManagerSecretAttributes) Replication() terra.ListValue[datasecretmanagersecret.ReplicationAttributes] {
	return terra.ReferenceAsList[datasecretmanagersecret.ReplicationAttributes](sms.ref.Append("replication"))
}

func (sms dataSecretManagerSecretAttributes) Rotation() terra.ListValue[datasecretmanagersecret.RotationAttributes] {
	return terra.ReferenceAsList[datasecretmanagersecret.RotationAttributes](sms.ref.Append("rotation"))
}

func (sms dataSecretManagerSecretAttributes) Topics() terra.ListValue[datasecretmanagersecret.TopicsAttributes] {
	return terra.ReferenceAsList[datasecretmanagersecret.TopicsAttributes](sms.ref.Append("topics"))
}
