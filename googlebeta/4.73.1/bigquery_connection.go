// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	bigqueryconnection "github.com/golingon/terraproviders/googlebeta/4.73.1/bigqueryconnection"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewBigqueryConnection creates a new instance of [BigqueryConnection].
func NewBigqueryConnection(name string, args BigqueryConnectionArgs) *BigqueryConnection {
	return &BigqueryConnection{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*BigqueryConnection)(nil)

// BigqueryConnection represents the Terraform resource google_bigquery_connection.
type BigqueryConnection struct {
	Name      string
	Args      BigqueryConnectionArgs
	state     *bigqueryConnectionState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [BigqueryConnection].
func (bc *BigqueryConnection) Type() string {
	return "google_bigquery_connection"
}

// LocalName returns the local name for [BigqueryConnection].
func (bc *BigqueryConnection) LocalName() string {
	return bc.Name
}

// Configuration returns the configuration (args) for [BigqueryConnection].
func (bc *BigqueryConnection) Configuration() interface{} {
	return bc.Args
}

// DependOn is used for other resources to depend on [BigqueryConnection].
func (bc *BigqueryConnection) DependOn() terra.Reference {
	return terra.ReferenceResource(bc)
}

// Dependencies returns the list of resources [BigqueryConnection] depends_on.
func (bc *BigqueryConnection) Dependencies() terra.Dependencies {
	return bc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [BigqueryConnection].
func (bc *BigqueryConnection) LifecycleManagement() *terra.Lifecycle {
	return bc.Lifecycle
}

// Attributes returns the attributes for [BigqueryConnection].
func (bc *BigqueryConnection) Attributes() bigqueryConnectionAttributes {
	return bigqueryConnectionAttributes{ref: terra.ReferenceResource(bc)}
}

// ImportState imports the given attribute values into [BigqueryConnection]'s state.
func (bc *BigqueryConnection) ImportState(av io.Reader) error {
	bc.state = &bigqueryConnectionState{}
	if err := json.NewDecoder(av).Decode(bc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", bc.Type(), bc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [BigqueryConnection] has state.
func (bc *BigqueryConnection) State() (*bigqueryConnectionState, bool) {
	return bc.state, bc.state != nil
}

// StateMust returns the state for [BigqueryConnection]. Panics if the state is nil.
func (bc *BigqueryConnection) StateMust() *bigqueryConnectionState {
	if bc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", bc.Type(), bc.LocalName()))
	}
	return bc.state
}

// BigqueryConnectionArgs contains the configurations for google_bigquery_connection.
type BigqueryConnectionArgs struct {
	// ConnectionId: string, optional
	ConnectionId terra.StringValue `hcl:"connection_id,attr"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// FriendlyName: string, optional
	FriendlyName terra.StringValue `hcl:"friendly_name,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, optional
	Location terra.StringValue `hcl:"location,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Aws: optional
	Aws *bigqueryconnection.Aws `hcl:"aws,block"`
	// Azure: optional
	Azure *bigqueryconnection.Azure `hcl:"azure,block"`
	// CloudResource: optional
	CloudResource *bigqueryconnection.CloudResource `hcl:"cloud_resource,block"`
	// CloudSpanner: optional
	CloudSpanner *bigqueryconnection.CloudSpanner `hcl:"cloud_spanner,block"`
	// CloudSql: optional
	CloudSql *bigqueryconnection.CloudSql `hcl:"cloud_sql,block"`
	// Timeouts: optional
	Timeouts *bigqueryconnection.Timeouts `hcl:"timeouts,block"`
}
type bigqueryConnectionAttributes struct {
	ref terra.Reference
}

// ConnectionId returns a reference to field connection_id of google_bigquery_connection.
func (bc bigqueryConnectionAttributes) ConnectionId() terra.StringValue {
	return terra.ReferenceAsString(bc.ref.Append("connection_id"))
}

// Description returns a reference to field description of google_bigquery_connection.
func (bc bigqueryConnectionAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(bc.ref.Append("description"))
}

// FriendlyName returns a reference to field friendly_name of google_bigquery_connection.
func (bc bigqueryConnectionAttributes) FriendlyName() terra.StringValue {
	return terra.ReferenceAsString(bc.ref.Append("friendly_name"))
}

// HasCredential returns a reference to field has_credential of google_bigquery_connection.
func (bc bigqueryConnectionAttributes) HasCredential() terra.BoolValue {
	return terra.ReferenceAsBool(bc.ref.Append("has_credential"))
}

// Id returns a reference to field id of google_bigquery_connection.
func (bc bigqueryConnectionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(bc.ref.Append("id"))
}

// Location returns a reference to field location of google_bigquery_connection.
func (bc bigqueryConnectionAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(bc.ref.Append("location"))
}

// Name returns a reference to field name of google_bigquery_connection.
func (bc bigqueryConnectionAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(bc.ref.Append("name"))
}

// Project returns a reference to field project of google_bigquery_connection.
func (bc bigqueryConnectionAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(bc.ref.Append("project"))
}

func (bc bigqueryConnectionAttributes) Aws() terra.ListValue[bigqueryconnection.AwsAttributes] {
	return terra.ReferenceAsList[bigqueryconnection.AwsAttributes](bc.ref.Append("aws"))
}

func (bc bigqueryConnectionAttributes) Azure() terra.ListValue[bigqueryconnection.AzureAttributes] {
	return terra.ReferenceAsList[bigqueryconnection.AzureAttributes](bc.ref.Append("azure"))
}

func (bc bigqueryConnectionAttributes) CloudResource() terra.ListValue[bigqueryconnection.CloudResourceAttributes] {
	return terra.ReferenceAsList[bigqueryconnection.CloudResourceAttributes](bc.ref.Append("cloud_resource"))
}

func (bc bigqueryConnectionAttributes) CloudSpanner() terra.ListValue[bigqueryconnection.CloudSpannerAttributes] {
	return terra.ReferenceAsList[bigqueryconnection.CloudSpannerAttributes](bc.ref.Append("cloud_spanner"))
}

func (bc bigqueryConnectionAttributes) CloudSql() terra.ListValue[bigqueryconnection.CloudSqlAttributes] {
	return terra.ReferenceAsList[bigqueryconnection.CloudSqlAttributes](bc.ref.Append("cloud_sql"))
}

func (bc bigqueryConnectionAttributes) Timeouts() bigqueryconnection.TimeoutsAttributes {
	return terra.ReferenceAsSingle[bigqueryconnection.TimeoutsAttributes](bc.ref.Append("timeouts"))
}

type bigqueryConnectionState struct {
	ConnectionId  string                                  `json:"connection_id"`
	Description   string                                  `json:"description"`
	FriendlyName  string                                  `json:"friendly_name"`
	HasCredential bool                                    `json:"has_credential"`
	Id            string                                  `json:"id"`
	Location      string                                  `json:"location"`
	Name          string                                  `json:"name"`
	Project       string                                  `json:"project"`
	Aws           []bigqueryconnection.AwsState           `json:"aws"`
	Azure         []bigqueryconnection.AzureState         `json:"azure"`
	CloudResource []bigqueryconnection.CloudResourceState `json:"cloud_resource"`
	CloudSpanner  []bigqueryconnection.CloudSpannerState  `json:"cloud_spanner"`
	CloudSql      []bigqueryconnection.CloudSqlState      `json:"cloud_sql"`
	Timeouts      *bigqueryconnection.TimeoutsState       `json:"timeouts"`
}
