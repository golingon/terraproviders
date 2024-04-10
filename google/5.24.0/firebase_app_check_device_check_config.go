// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	firebaseappcheckdevicecheckconfig "github.com/golingon/terraproviders/google/5.24.0/firebaseappcheckdevicecheckconfig"
	"io"
)

// NewFirebaseAppCheckDeviceCheckConfig creates a new instance of [FirebaseAppCheckDeviceCheckConfig].
func NewFirebaseAppCheckDeviceCheckConfig(name string, args FirebaseAppCheckDeviceCheckConfigArgs) *FirebaseAppCheckDeviceCheckConfig {
	return &FirebaseAppCheckDeviceCheckConfig{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*FirebaseAppCheckDeviceCheckConfig)(nil)

// FirebaseAppCheckDeviceCheckConfig represents the Terraform resource google_firebase_app_check_device_check_config.
type FirebaseAppCheckDeviceCheckConfig struct {
	Name      string
	Args      FirebaseAppCheckDeviceCheckConfigArgs
	state     *firebaseAppCheckDeviceCheckConfigState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [FirebaseAppCheckDeviceCheckConfig].
func (facdcc *FirebaseAppCheckDeviceCheckConfig) Type() string {
	return "google_firebase_app_check_device_check_config"
}

// LocalName returns the local name for [FirebaseAppCheckDeviceCheckConfig].
func (facdcc *FirebaseAppCheckDeviceCheckConfig) LocalName() string {
	return facdcc.Name
}

// Configuration returns the configuration (args) for [FirebaseAppCheckDeviceCheckConfig].
func (facdcc *FirebaseAppCheckDeviceCheckConfig) Configuration() interface{} {
	return facdcc.Args
}

// DependOn is used for other resources to depend on [FirebaseAppCheckDeviceCheckConfig].
func (facdcc *FirebaseAppCheckDeviceCheckConfig) DependOn() terra.Reference {
	return terra.ReferenceResource(facdcc)
}

// Dependencies returns the list of resources [FirebaseAppCheckDeviceCheckConfig] depends_on.
func (facdcc *FirebaseAppCheckDeviceCheckConfig) Dependencies() terra.Dependencies {
	return facdcc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [FirebaseAppCheckDeviceCheckConfig].
func (facdcc *FirebaseAppCheckDeviceCheckConfig) LifecycleManagement() *terra.Lifecycle {
	return facdcc.Lifecycle
}

// Attributes returns the attributes for [FirebaseAppCheckDeviceCheckConfig].
func (facdcc *FirebaseAppCheckDeviceCheckConfig) Attributes() firebaseAppCheckDeviceCheckConfigAttributes {
	return firebaseAppCheckDeviceCheckConfigAttributes{ref: terra.ReferenceResource(facdcc)}
}

// ImportState imports the given attribute values into [FirebaseAppCheckDeviceCheckConfig]'s state.
func (facdcc *FirebaseAppCheckDeviceCheckConfig) ImportState(av io.Reader) error {
	facdcc.state = &firebaseAppCheckDeviceCheckConfigState{}
	if err := json.NewDecoder(av).Decode(facdcc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", facdcc.Type(), facdcc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [FirebaseAppCheckDeviceCheckConfig] has state.
func (facdcc *FirebaseAppCheckDeviceCheckConfig) State() (*firebaseAppCheckDeviceCheckConfigState, bool) {
	return facdcc.state, facdcc.state != nil
}

// StateMust returns the state for [FirebaseAppCheckDeviceCheckConfig]. Panics if the state is nil.
func (facdcc *FirebaseAppCheckDeviceCheckConfig) StateMust() *firebaseAppCheckDeviceCheckConfigState {
	if facdcc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", facdcc.Type(), facdcc.LocalName()))
	}
	return facdcc.state
}

// FirebaseAppCheckDeviceCheckConfigArgs contains the configurations for google_firebase_app_check_device_check_config.
type FirebaseAppCheckDeviceCheckConfigArgs struct {
	// AppId: string, required
	AppId terra.StringValue `hcl:"app_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// KeyId: string, required
	KeyId terra.StringValue `hcl:"key_id,attr" validate:"required"`
	// PrivateKey: string, required
	PrivateKey terra.StringValue `hcl:"private_key,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// TokenTtl: string, optional
	TokenTtl terra.StringValue `hcl:"token_ttl,attr"`
	// Timeouts: optional
	Timeouts *firebaseappcheckdevicecheckconfig.Timeouts `hcl:"timeouts,block"`
}
type firebaseAppCheckDeviceCheckConfigAttributes struct {
	ref terra.Reference
}

// AppId returns a reference to field app_id of google_firebase_app_check_device_check_config.
func (facdcc firebaseAppCheckDeviceCheckConfigAttributes) AppId() terra.StringValue {
	return terra.ReferenceAsString(facdcc.ref.Append("app_id"))
}

// Id returns a reference to field id of google_firebase_app_check_device_check_config.
func (facdcc firebaseAppCheckDeviceCheckConfigAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(facdcc.ref.Append("id"))
}

// KeyId returns a reference to field key_id of google_firebase_app_check_device_check_config.
func (facdcc firebaseAppCheckDeviceCheckConfigAttributes) KeyId() terra.StringValue {
	return terra.ReferenceAsString(facdcc.ref.Append("key_id"))
}

// Name returns a reference to field name of google_firebase_app_check_device_check_config.
func (facdcc firebaseAppCheckDeviceCheckConfigAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(facdcc.ref.Append("name"))
}

// PrivateKey returns a reference to field private_key of google_firebase_app_check_device_check_config.
func (facdcc firebaseAppCheckDeviceCheckConfigAttributes) PrivateKey() terra.StringValue {
	return terra.ReferenceAsString(facdcc.ref.Append("private_key"))
}

// PrivateKeySet returns a reference to field private_key_set of google_firebase_app_check_device_check_config.
func (facdcc firebaseAppCheckDeviceCheckConfigAttributes) PrivateKeySet() terra.BoolValue {
	return terra.ReferenceAsBool(facdcc.ref.Append("private_key_set"))
}

// Project returns a reference to field project of google_firebase_app_check_device_check_config.
func (facdcc firebaseAppCheckDeviceCheckConfigAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(facdcc.ref.Append("project"))
}

// TokenTtl returns a reference to field token_ttl of google_firebase_app_check_device_check_config.
func (facdcc firebaseAppCheckDeviceCheckConfigAttributes) TokenTtl() terra.StringValue {
	return terra.ReferenceAsString(facdcc.ref.Append("token_ttl"))
}

func (facdcc firebaseAppCheckDeviceCheckConfigAttributes) Timeouts() firebaseappcheckdevicecheckconfig.TimeoutsAttributes {
	return terra.ReferenceAsSingle[firebaseappcheckdevicecheckconfig.TimeoutsAttributes](facdcc.ref.Append("timeouts"))
}

type firebaseAppCheckDeviceCheckConfigState struct {
	AppId         string                                           `json:"app_id"`
	Id            string                                           `json:"id"`
	KeyId         string                                           `json:"key_id"`
	Name          string                                           `json:"name"`
	PrivateKey    string                                           `json:"private_key"`
	PrivateKeySet bool                                             `json:"private_key_set"`
	Project       string                                           `json:"project"`
	TokenTtl      string                                           `json:"token_ttl"`
	Timeouts      *firebaseappcheckdevicecheckconfig.TimeoutsState `json:"timeouts"`
}
