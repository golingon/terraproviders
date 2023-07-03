// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	datafusioninstance "github.com/golingon/terraproviders/google/4.71.0/datafusioninstance"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDataFusionInstance creates a new instance of [DataFusionInstance].
func NewDataFusionInstance(name string, args DataFusionInstanceArgs) *DataFusionInstance {
	return &DataFusionInstance{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DataFusionInstance)(nil)

// DataFusionInstance represents the Terraform resource google_data_fusion_instance.
type DataFusionInstance struct {
	Name      string
	Args      DataFusionInstanceArgs
	state     *dataFusionInstanceState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DataFusionInstance].
func (dfi *DataFusionInstance) Type() string {
	return "google_data_fusion_instance"
}

// LocalName returns the local name for [DataFusionInstance].
func (dfi *DataFusionInstance) LocalName() string {
	return dfi.Name
}

// Configuration returns the configuration (args) for [DataFusionInstance].
func (dfi *DataFusionInstance) Configuration() interface{} {
	return dfi.Args
}

// DependOn is used for other resources to depend on [DataFusionInstance].
func (dfi *DataFusionInstance) DependOn() terra.Reference {
	return terra.ReferenceResource(dfi)
}

// Dependencies returns the list of resources [DataFusionInstance] depends_on.
func (dfi *DataFusionInstance) Dependencies() terra.Dependencies {
	return dfi.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DataFusionInstance].
func (dfi *DataFusionInstance) LifecycleManagement() *terra.Lifecycle {
	return dfi.Lifecycle
}

// Attributes returns the attributes for [DataFusionInstance].
func (dfi *DataFusionInstance) Attributes() dataFusionInstanceAttributes {
	return dataFusionInstanceAttributes{ref: terra.ReferenceResource(dfi)}
}

// ImportState imports the given attribute values into [DataFusionInstance]'s state.
func (dfi *DataFusionInstance) ImportState(av io.Reader) error {
	dfi.state = &dataFusionInstanceState{}
	if err := json.NewDecoder(av).Decode(dfi.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dfi.Type(), dfi.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DataFusionInstance] has state.
func (dfi *DataFusionInstance) State() (*dataFusionInstanceState, bool) {
	return dfi.state, dfi.state != nil
}

// StateMust returns the state for [DataFusionInstance]. Panics if the state is nil.
func (dfi *DataFusionInstance) StateMust() *dataFusionInstanceState {
	if dfi.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dfi.Type(), dfi.LocalName()))
	}
	return dfi.state
}

// DataFusionInstanceArgs contains the configurations for google_data_fusion_instance.
type DataFusionInstanceArgs struct {
	// DataprocServiceAccount: string, optional
	DataprocServiceAccount terra.StringValue `hcl:"dataproc_service_account,attr"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// DisplayName: string, optional
	DisplayName terra.StringValue `hcl:"display_name,attr"`
	// EnableRbac: bool, optional
	EnableRbac terra.BoolValue `hcl:"enable_rbac,attr"`
	// EnableStackdriverLogging: bool, optional
	EnableStackdriverLogging terra.BoolValue `hcl:"enable_stackdriver_logging,attr"`
	// EnableStackdriverMonitoring: bool, optional
	EnableStackdriverMonitoring terra.BoolValue `hcl:"enable_stackdriver_monitoring,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Labels: map of string, optional
	Labels terra.MapValue[terra.StringValue] `hcl:"labels,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Options: map of string, optional
	Options terra.MapValue[terra.StringValue] `hcl:"options,attr"`
	// PrivateInstance: bool, optional
	PrivateInstance terra.BoolValue `hcl:"private_instance,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Region: string, optional
	Region terra.StringValue `hcl:"region,attr"`
	// Type: string, required
	Type terra.StringValue `hcl:"type,attr" validate:"required"`
	// Version: string, optional
	Version terra.StringValue `hcl:"version,attr"`
	// Zone: string, optional
	Zone terra.StringValue `hcl:"zone,attr"`
	// Accelerators: min=0
	Accelerators []datafusioninstance.Accelerators `hcl:"accelerators,block" validate:"min=0"`
	// CryptoKeyConfig: optional
	CryptoKeyConfig *datafusioninstance.CryptoKeyConfig `hcl:"crypto_key_config,block"`
	// EventPublishConfig: optional
	EventPublishConfig *datafusioninstance.EventPublishConfig `hcl:"event_publish_config,block"`
	// NetworkConfig: optional
	NetworkConfig *datafusioninstance.NetworkConfig `hcl:"network_config,block"`
	// Timeouts: optional
	Timeouts *datafusioninstance.Timeouts `hcl:"timeouts,block"`
}
type dataFusionInstanceAttributes struct {
	ref terra.Reference
}

// ApiEndpoint returns a reference to field api_endpoint of google_data_fusion_instance.
func (dfi dataFusionInstanceAttributes) ApiEndpoint() terra.StringValue {
	return terra.ReferenceAsString(dfi.ref.Append("api_endpoint"))
}

// CreateTime returns a reference to field create_time of google_data_fusion_instance.
func (dfi dataFusionInstanceAttributes) CreateTime() terra.StringValue {
	return terra.ReferenceAsString(dfi.ref.Append("create_time"))
}

// DataprocServiceAccount returns a reference to field dataproc_service_account of google_data_fusion_instance.
func (dfi dataFusionInstanceAttributes) DataprocServiceAccount() terra.StringValue {
	return terra.ReferenceAsString(dfi.ref.Append("dataproc_service_account"))
}

// Description returns a reference to field description of google_data_fusion_instance.
func (dfi dataFusionInstanceAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(dfi.ref.Append("description"))
}

// DisplayName returns a reference to field display_name of google_data_fusion_instance.
func (dfi dataFusionInstanceAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(dfi.ref.Append("display_name"))
}

// EnableRbac returns a reference to field enable_rbac of google_data_fusion_instance.
func (dfi dataFusionInstanceAttributes) EnableRbac() terra.BoolValue {
	return terra.ReferenceAsBool(dfi.ref.Append("enable_rbac"))
}

// EnableStackdriverLogging returns a reference to field enable_stackdriver_logging of google_data_fusion_instance.
func (dfi dataFusionInstanceAttributes) EnableStackdriverLogging() terra.BoolValue {
	return terra.ReferenceAsBool(dfi.ref.Append("enable_stackdriver_logging"))
}

// EnableStackdriverMonitoring returns a reference to field enable_stackdriver_monitoring of google_data_fusion_instance.
func (dfi dataFusionInstanceAttributes) EnableStackdriverMonitoring() terra.BoolValue {
	return terra.ReferenceAsBool(dfi.ref.Append("enable_stackdriver_monitoring"))
}

// GcsBucket returns a reference to field gcs_bucket of google_data_fusion_instance.
func (dfi dataFusionInstanceAttributes) GcsBucket() terra.StringValue {
	return terra.ReferenceAsString(dfi.ref.Append("gcs_bucket"))
}

// Id returns a reference to field id of google_data_fusion_instance.
func (dfi dataFusionInstanceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dfi.ref.Append("id"))
}

// Labels returns a reference to field labels of google_data_fusion_instance.
func (dfi dataFusionInstanceAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dfi.ref.Append("labels"))
}

// Name returns a reference to field name of google_data_fusion_instance.
func (dfi dataFusionInstanceAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(dfi.ref.Append("name"))
}

// Options returns a reference to field options of google_data_fusion_instance.
func (dfi dataFusionInstanceAttributes) Options() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dfi.ref.Append("options"))
}

// P4ServiceAccount returns a reference to field p4_service_account of google_data_fusion_instance.
func (dfi dataFusionInstanceAttributes) P4ServiceAccount() terra.StringValue {
	return terra.ReferenceAsString(dfi.ref.Append("p4_service_account"))
}

// PrivateInstance returns a reference to field private_instance of google_data_fusion_instance.
func (dfi dataFusionInstanceAttributes) PrivateInstance() terra.BoolValue {
	return terra.ReferenceAsBool(dfi.ref.Append("private_instance"))
}

// Project returns a reference to field project of google_data_fusion_instance.
func (dfi dataFusionInstanceAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(dfi.ref.Append("project"))
}

// Region returns a reference to field region of google_data_fusion_instance.
func (dfi dataFusionInstanceAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(dfi.ref.Append("region"))
}

// ServiceEndpoint returns a reference to field service_endpoint of google_data_fusion_instance.
func (dfi dataFusionInstanceAttributes) ServiceEndpoint() terra.StringValue {
	return terra.ReferenceAsString(dfi.ref.Append("service_endpoint"))
}

// State returns a reference to field state of google_data_fusion_instance.
func (dfi dataFusionInstanceAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(dfi.ref.Append("state"))
}

// StateMessage returns a reference to field state_message of google_data_fusion_instance.
func (dfi dataFusionInstanceAttributes) StateMessage() terra.StringValue {
	return terra.ReferenceAsString(dfi.ref.Append("state_message"))
}

// TenantProjectId returns a reference to field tenant_project_id of google_data_fusion_instance.
func (dfi dataFusionInstanceAttributes) TenantProjectId() terra.StringValue {
	return terra.ReferenceAsString(dfi.ref.Append("tenant_project_id"))
}

// Type returns a reference to field type of google_data_fusion_instance.
func (dfi dataFusionInstanceAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(dfi.ref.Append("type"))
}

// UpdateTime returns a reference to field update_time of google_data_fusion_instance.
func (dfi dataFusionInstanceAttributes) UpdateTime() terra.StringValue {
	return terra.ReferenceAsString(dfi.ref.Append("update_time"))
}

// Version returns a reference to field version of google_data_fusion_instance.
func (dfi dataFusionInstanceAttributes) Version() terra.StringValue {
	return terra.ReferenceAsString(dfi.ref.Append("version"))
}

// Zone returns a reference to field zone of google_data_fusion_instance.
func (dfi dataFusionInstanceAttributes) Zone() terra.StringValue {
	return terra.ReferenceAsString(dfi.ref.Append("zone"))
}

func (dfi dataFusionInstanceAttributes) Accelerators() terra.ListValue[datafusioninstance.AcceleratorsAttributes] {
	return terra.ReferenceAsList[datafusioninstance.AcceleratorsAttributes](dfi.ref.Append("accelerators"))
}

func (dfi dataFusionInstanceAttributes) CryptoKeyConfig() terra.ListValue[datafusioninstance.CryptoKeyConfigAttributes] {
	return terra.ReferenceAsList[datafusioninstance.CryptoKeyConfigAttributes](dfi.ref.Append("crypto_key_config"))
}

func (dfi dataFusionInstanceAttributes) EventPublishConfig() terra.ListValue[datafusioninstance.EventPublishConfigAttributes] {
	return terra.ReferenceAsList[datafusioninstance.EventPublishConfigAttributes](dfi.ref.Append("event_publish_config"))
}

func (dfi dataFusionInstanceAttributes) NetworkConfig() terra.ListValue[datafusioninstance.NetworkConfigAttributes] {
	return terra.ReferenceAsList[datafusioninstance.NetworkConfigAttributes](dfi.ref.Append("network_config"))
}

func (dfi dataFusionInstanceAttributes) Timeouts() datafusioninstance.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datafusioninstance.TimeoutsAttributes](dfi.ref.Append("timeouts"))
}

type dataFusionInstanceState struct {
	ApiEndpoint                 string                                       `json:"api_endpoint"`
	CreateTime                  string                                       `json:"create_time"`
	DataprocServiceAccount      string                                       `json:"dataproc_service_account"`
	Description                 string                                       `json:"description"`
	DisplayName                 string                                       `json:"display_name"`
	EnableRbac                  bool                                         `json:"enable_rbac"`
	EnableStackdriverLogging    bool                                         `json:"enable_stackdriver_logging"`
	EnableStackdriverMonitoring bool                                         `json:"enable_stackdriver_monitoring"`
	GcsBucket                   string                                       `json:"gcs_bucket"`
	Id                          string                                       `json:"id"`
	Labels                      map[string]string                            `json:"labels"`
	Name                        string                                       `json:"name"`
	Options                     map[string]string                            `json:"options"`
	P4ServiceAccount            string                                       `json:"p4_service_account"`
	PrivateInstance             bool                                         `json:"private_instance"`
	Project                     string                                       `json:"project"`
	Region                      string                                       `json:"region"`
	ServiceEndpoint             string                                       `json:"service_endpoint"`
	State                       string                                       `json:"state"`
	StateMessage                string                                       `json:"state_message"`
	TenantProjectId             string                                       `json:"tenant_project_id"`
	Type                        string                                       `json:"type"`
	UpdateTime                  string                                       `json:"update_time"`
	Version                     string                                       `json:"version"`
	Zone                        string                                       `json:"zone"`
	Accelerators                []datafusioninstance.AcceleratorsState       `json:"accelerators"`
	CryptoKeyConfig             []datafusioninstance.CryptoKeyConfigState    `json:"crypto_key_config"`
	EventPublishConfig          []datafusioninstance.EventPublishConfigState `json:"event_publish_config"`
	NetworkConfig               []datafusioninstance.NetworkConfigState      `json:"network_config"`
	Timeouts                    *datafusioninstance.TimeoutsState            `json:"timeouts"`
}
