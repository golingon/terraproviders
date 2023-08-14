// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	firebaseextensionsinstance "github.com/golingon/terraproviders/googlebeta/4.77.0/firebaseextensionsinstance"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewFirebaseExtensionsInstance creates a new instance of [FirebaseExtensionsInstance].
func NewFirebaseExtensionsInstance(name string, args FirebaseExtensionsInstanceArgs) *FirebaseExtensionsInstance {
	return &FirebaseExtensionsInstance{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*FirebaseExtensionsInstance)(nil)

// FirebaseExtensionsInstance represents the Terraform resource google_firebase_extensions_instance.
type FirebaseExtensionsInstance struct {
	Name      string
	Args      FirebaseExtensionsInstanceArgs
	state     *firebaseExtensionsInstanceState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [FirebaseExtensionsInstance].
func (fei *FirebaseExtensionsInstance) Type() string {
	return "google_firebase_extensions_instance"
}

// LocalName returns the local name for [FirebaseExtensionsInstance].
func (fei *FirebaseExtensionsInstance) LocalName() string {
	return fei.Name
}

// Configuration returns the configuration (args) for [FirebaseExtensionsInstance].
func (fei *FirebaseExtensionsInstance) Configuration() interface{} {
	return fei.Args
}

// DependOn is used for other resources to depend on [FirebaseExtensionsInstance].
func (fei *FirebaseExtensionsInstance) DependOn() terra.Reference {
	return terra.ReferenceResource(fei)
}

// Dependencies returns the list of resources [FirebaseExtensionsInstance] depends_on.
func (fei *FirebaseExtensionsInstance) Dependencies() terra.Dependencies {
	return fei.DependsOn
}

// LifecycleManagement returns the lifecycle block for [FirebaseExtensionsInstance].
func (fei *FirebaseExtensionsInstance) LifecycleManagement() *terra.Lifecycle {
	return fei.Lifecycle
}

// Attributes returns the attributes for [FirebaseExtensionsInstance].
func (fei *FirebaseExtensionsInstance) Attributes() firebaseExtensionsInstanceAttributes {
	return firebaseExtensionsInstanceAttributes{ref: terra.ReferenceResource(fei)}
}

// ImportState imports the given attribute values into [FirebaseExtensionsInstance]'s state.
func (fei *FirebaseExtensionsInstance) ImportState(av io.Reader) error {
	fei.state = &firebaseExtensionsInstanceState{}
	if err := json.NewDecoder(av).Decode(fei.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", fei.Type(), fei.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [FirebaseExtensionsInstance] has state.
func (fei *FirebaseExtensionsInstance) State() (*firebaseExtensionsInstanceState, bool) {
	return fei.state, fei.state != nil
}

// StateMust returns the state for [FirebaseExtensionsInstance]. Panics if the state is nil.
func (fei *FirebaseExtensionsInstance) StateMust() *firebaseExtensionsInstanceState {
	if fei.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", fei.Type(), fei.LocalName()))
	}
	return fei.state
}

// FirebaseExtensionsInstanceArgs contains the configurations for google_firebase_extensions_instance.
type FirebaseExtensionsInstanceArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// InstanceId: string, required
	InstanceId terra.StringValue `hcl:"instance_id,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// ErrorStatus: min=0
	ErrorStatus []firebaseextensionsinstance.ErrorStatus `hcl:"error_status,block" validate:"min=0"`
	// RuntimeData: min=0
	RuntimeData []firebaseextensionsinstance.RuntimeData `hcl:"runtime_data,block" validate:"min=0"`
	// Config: required
	Config *firebaseextensionsinstance.Config `hcl:"config,block" validate:"required"`
	// Timeouts: optional
	Timeouts *firebaseextensionsinstance.Timeouts `hcl:"timeouts,block"`
}
type firebaseExtensionsInstanceAttributes struct {
	ref terra.Reference
}

// CreateTime returns a reference to field create_time of google_firebase_extensions_instance.
func (fei firebaseExtensionsInstanceAttributes) CreateTime() terra.StringValue {
	return terra.ReferenceAsString(fei.ref.Append("create_time"))
}

// Etag returns a reference to field etag of google_firebase_extensions_instance.
func (fei firebaseExtensionsInstanceAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(fei.ref.Append("etag"))
}

// Id returns a reference to field id of google_firebase_extensions_instance.
func (fei firebaseExtensionsInstanceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(fei.ref.Append("id"))
}

// InstanceId returns a reference to field instance_id of google_firebase_extensions_instance.
func (fei firebaseExtensionsInstanceAttributes) InstanceId() terra.StringValue {
	return terra.ReferenceAsString(fei.ref.Append("instance_id"))
}

// LastOperationName returns a reference to field last_operation_name of google_firebase_extensions_instance.
func (fei firebaseExtensionsInstanceAttributes) LastOperationName() terra.StringValue {
	return terra.ReferenceAsString(fei.ref.Append("last_operation_name"))
}

// LastOperationType returns a reference to field last_operation_type of google_firebase_extensions_instance.
func (fei firebaseExtensionsInstanceAttributes) LastOperationType() terra.StringValue {
	return terra.ReferenceAsString(fei.ref.Append("last_operation_type"))
}

// Name returns a reference to field name of google_firebase_extensions_instance.
func (fei firebaseExtensionsInstanceAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(fei.ref.Append("name"))
}

// Project returns a reference to field project of google_firebase_extensions_instance.
func (fei firebaseExtensionsInstanceAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(fei.ref.Append("project"))
}

// ServiceAccountEmail returns a reference to field service_account_email of google_firebase_extensions_instance.
func (fei firebaseExtensionsInstanceAttributes) ServiceAccountEmail() terra.StringValue {
	return terra.ReferenceAsString(fei.ref.Append("service_account_email"))
}

// State returns a reference to field state of google_firebase_extensions_instance.
func (fei firebaseExtensionsInstanceAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(fei.ref.Append("state"))
}

// UpdateTime returns a reference to field update_time of google_firebase_extensions_instance.
func (fei firebaseExtensionsInstanceAttributes) UpdateTime() terra.StringValue {
	return terra.ReferenceAsString(fei.ref.Append("update_time"))
}

func (fei firebaseExtensionsInstanceAttributes) ErrorStatus() terra.ListValue[firebaseextensionsinstance.ErrorStatusAttributes] {
	return terra.ReferenceAsList[firebaseextensionsinstance.ErrorStatusAttributes](fei.ref.Append("error_status"))
}

func (fei firebaseExtensionsInstanceAttributes) RuntimeData() terra.ListValue[firebaseextensionsinstance.RuntimeDataAttributes] {
	return terra.ReferenceAsList[firebaseextensionsinstance.RuntimeDataAttributes](fei.ref.Append("runtime_data"))
}

func (fei firebaseExtensionsInstanceAttributes) Config() terra.ListValue[firebaseextensionsinstance.ConfigAttributes] {
	return terra.ReferenceAsList[firebaseextensionsinstance.ConfigAttributes](fei.ref.Append("config"))
}

func (fei firebaseExtensionsInstanceAttributes) Timeouts() firebaseextensionsinstance.TimeoutsAttributes {
	return terra.ReferenceAsSingle[firebaseextensionsinstance.TimeoutsAttributes](fei.ref.Append("timeouts"))
}

type firebaseExtensionsInstanceState struct {
	CreateTime          string                                        `json:"create_time"`
	Etag                string                                        `json:"etag"`
	Id                  string                                        `json:"id"`
	InstanceId          string                                        `json:"instance_id"`
	LastOperationName   string                                        `json:"last_operation_name"`
	LastOperationType   string                                        `json:"last_operation_type"`
	Name                string                                        `json:"name"`
	Project             string                                        `json:"project"`
	ServiceAccountEmail string                                        `json:"service_account_email"`
	State               string                                        `json:"state"`
	UpdateTime          string                                        `json:"update_time"`
	ErrorStatus         []firebaseextensionsinstance.ErrorStatusState `json:"error_status"`
	RuntimeData         []firebaseextensionsinstance.RuntimeDataState `json:"runtime_data"`
	Config              []firebaseextensionsinstance.ConfigState      `json:"config"`
	Timeouts            *firebaseextensionsinstance.TimeoutsState     `json:"timeouts"`
}
