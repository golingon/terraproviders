// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	secretmanagersecret "github.com/golingon/terraproviders/googlebeta/4.77.0/secretmanagersecret"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSecretManagerSecret creates a new instance of [SecretManagerSecret].
func NewSecretManagerSecret(name string, args SecretManagerSecretArgs) *SecretManagerSecret {
	return &SecretManagerSecret{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SecretManagerSecret)(nil)

// SecretManagerSecret represents the Terraform resource google_secret_manager_secret.
type SecretManagerSecret struct {
	Name      string
	Args      SecretManagerSecretArgs
	state     *secretManagerSecretState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SecretManagerSecret].
func (sms *SecretManagerSecret) Type() string {
	return "google_secret_manager_secret"
}

// LocalName returns the local name for [SecretManagerSecret].
func (sms *SecretManagerSecret) LocalName() string {
	return sms.Name
}

// Configuration returns the configuration (args) for [SecretManagerSecret].
func (sms *SecretManagerSecret) Configuration() interface{} {
	return sms.Args
}

// DependOn is used for other resources to depend on [SecretManagerSecret].
func (sms *SecretManagerSecret) DependOn() terra.Reference {
	return terra.ReferenceResource(sms)
}

// Dependencies returns the list of resources [SecretManagerSecret] depends_on.
func (sms *SecretManagerSecret) Dependencies() terra.Dependencies {
	return sms.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SecretManagerSecret].
func (sms *SecretManagerSecret) LifecycleManagement() *terra.Lifecycle {
	return sms.Lifecycle
}

// Attributes returns the attributes for [SecretManagerSecret].
func (sms *SecretManagerSecret) Attributes() secretManagerSecretAttributes {
	return secretManagerSecretAttributes{ref: terra.ReferenceResource(sms)}
}

// ImportState imports the given attribute values into [SecretManagerSecret]'s state.
func (sms *SecretManagerSecret) ImportState(av io.Reader) error {
	sms.state = &secretManagerSecretState{}
	if err := json.NewDecoder(av).Decode(sms.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sms.Type(), sms.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SecretManagerSecret] has state.
func (sms *SecretManagerSecret) State() (*secretManagerSecretState, bool) {
	return sms.state, sms.state != nil
}

// StateMust returns the state for [SecretManagerSecret]. Panics if the state is nil.
func (sms *SecretManagerSecret) StateMust() *secretManagerSecretState {
	if sms.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sms.Type(), sms.LocalName()))
	}
	return sms.state
}

// SecretManagerSecretArgs contains the configurations for google_secret_manager_secret.
type SecretManagerSecretArgs struct {
	// ExpireTime: string, optional
	ExpireTime terra.StringValue `hcl:"expire_time,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Labels: map of string, optional
	Labels terra.MapValue[terra.StringValue] `hcl:"labels,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// SecretId: string, required
	SecretId terra.StringValue `hcl:"secret_id,attr" validate:"required"`
	// Ttl: string, optional
	Ttl terra.StringValue `hcl:"ttl,attr"`
	// Replication: required
	Replication *secretmanagersecret.Replication `hcl:"replication,block" validate:"required"`
	// Rotation: optional
	Rotation *secretmanagersecret.Rotation `hcl:"rotation,block"`
	// Timeouts: optional
	Timeouts *secretmanagersecret.Timeouts `hcl:"timeouts,block"`
	// Topics: min=0
	Topics []secretmanagersecret.Topics `hcl:"topics,block" validate:"min=0"`
}
type secretManagerSecretAttributes struct {
	ref terra.Reference
}

// CreateTime returns a reference to field create_time of google_secret_manager_secret.
func (sms secretManagerSecretAttributes) CreateTime() terra.StringValue {
	return terra.ReferenceAsString(sms.ref.Append("create_time"))
}

// ExpireTime returns a reference to field expire_time of google_secret_manager_secret.
func (sms secretManagerSecretAttributes) ExpireTime() terra.StringValue {
	return terra.ReferenceAsString(sms.ref.Append("expire_time"))
}

// Id returns a reference to field id of google_secret_manager_secret.
func (sms secretManagerSecretAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sms.ref.Append("id"))
}

// Labels returns a reference to field labels of google_secret_manager_secret.
func (sms secretManagerSecretAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](sms.ref.Append("labels"))
}

// Name returns a reference to field name of google_secret_manager_secret.
func (sms secretManagerSecretAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(sms.ref.Append("name"))
}

// Project returns a reference to field project of google_secret_manager_secret.
func (sms secretManagerSecretAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(sms.ref.Append("project"))
}

// SecretId returns a reference to field secret_id of google_secret_manager_secret.
func (sms secretManagerSecretAttributes) SecretId() terra.StringValue {
	return terra.ReferenceAsString(sms.ref.Append("secret_id"))
}

// Ttl returns a reference to field ttl of google_secret_manager_secret.
func (sms secretManagerSecretAttributes) Ttl() terra.StringValue {
	return terra.ReferenceAsString(sms.ref.Append("ttl"))
}

func (sms secretManagerSecretAttributes) Replication() terra.ListValue[secretmanagersecret.ReplicationAttributes] {
	return terra.ReferenceAsList[secretmanagersecret.ReplicationAttributes](sms.ref.Append("replication"))
}

func (sms secretManagerSecretAttributes) Rotation() terra.ListValue[secretmanagersecret.RotationAttributes] {
	return terra.ReferenceAsList[secretmanagersecret.RotationAttributes](sms.ref.Append("rotation"))
}

func (sms secretManagerSecretAttributes) Timeouts() secretmanagersecret.TimeoutsAttributes {
	return terra.ReferenceAsSingle[secretmanagersecret.TimeoutsAttributes](sms.ref.Append("timeouts"))
}

func (sms secretManagerSecretAttributes) Topics() terra.ListValue[secretmanagersecret.TopicsAttributes] {
	return terra.ReferenceAsList[secretmanagersecret.TopicsAttributes](sms.ref.Append("topics"))
}

type secretManagerSecretState struct {
	CreateTime  string                                 `json:"create_time"`
	ExpireTime  string                                 `json:"expire_time"`
	Id          string                                 `json:"id"`
	Labels      map[string]string                      `json:"labels"`
	Name        string                                 `json:"name"`
	Project     string                                 `json:"project"`
	SecretId    string                                 `json:"secret_id"`
	Ttl         string                                 `json:"ttl"`
	Replication []secretmanagersecret.ReplicationState `json:"replication"`
	Rotation    []secretmanagersecret.RotationState    `json:"rotation"`
	Timeouts    *secretmanagersecret.TimeoutsState     `json:"timeouts"`
	Topics      []secretmanagersecret.TopicsState      `json:"topics"`
}
