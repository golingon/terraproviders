// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	biglaketable "github.com/golingon/terraproviders/googlebeta/5.24.0/biglaketable"
	"io"
)

// NewBiglakeTable creates a new instance of [BiglakeTable].
func NewBiglakeTable(name string, args BiglakeTableArgs) *BiglakeTable {
	return &BiglakeTable{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*BiglakeTable)(nil)

// BiglakeTable represents the Terraform resource google_biglake_table.
type BiglakeTable struct {
	Name      string
	Args      BiglakeTableArgs
	state     *biglakeTableState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [BiglakeTable].
func (bt *BiglakeTable) Type() string {
	return "google_biglake_table"
}

// LocalName returns the local name for [BiglakeTable].
func (bt *BiglakeTable) LocalName() string {
	return bt.Name
}

// Configuration returns the configuration (args) for [BiglakeTable].
func (bt *BiglakeTable) Configuration() interface{} {
	return bt.Args
}

// DependOn is used for other resources to depend on [BiglakeTable].
func (bt *BiglakeTable) DependOn() terra.Reference {
	return terra.ReferenceResource(bt)
}

// Dependencies returns the list of resources [BiglakeTable] depends_on.
func (bt *BiglakeTable) Dependencies() terra.Dependencies {
	return bt.DependsOn
}

// LifecycleManagement returns the lifecycle block for [BiglakeTable].
func (bt *BiglakeTable) LifecycleManagement() *terra.Lifecycle {
	return bt.Lifecycle
}

// Attributes returns the attributes for [BiglakeTable].
func (bt *BiglakeTable) Attributes() biglakeTableAttributes {
	return biglakeTableAttributes{ref: terra.ReferenceResource(bt)}
}

// ImportState imports the given attribute values into [BiglakeTable]'s state.
func (bt *BiglakeTable) ImportState(av io.Reader) error {
	bt.state = &biglakeTableState{}
	if err := json.NewDecoder(av).Decode(bt.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", bt.Type(), bt.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [BiglakeTable] has state.
func (bt *BiglakeTable) State() (*biglakeTableState, bool) {
	return bt.state, bt.state != nil
}

// StateMust returns the state for [BiglakeTable]. Panics if the state is nil.
func (bt *BiglakeTable) StateMust() *biglakeTableState {
	if bt.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", bt.Type(), bt.LocalName()))
	}
	return bt.state
}

// BiglakeTableArgs contains the configurations for google_biglake_table.
type BiglakeTableArgs struct {
	// Database: string, optional
	Database terra.StringValue `hcl:"database,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Type: string, optional
	Type terra.StringValue `hcl:"type,attr"`
	// HiveOptions: optional
	HiveOptions *biglaketable.HiveOptions `hcl:"hive_options,block"`
	// Timeouts: optional
	Timeouts *biglaketable.Timeouts `hcl:"timeouts,block"`
}
type biglakeTableAttributes struct {
	ref terra.Reference
}

// CreateTime returns a reference to field create_time of google_biglake_table.
func (bt biglakeTableAttributes) CreateTime() terra.StringValue {
	return terra.ReferenceAsString(bt.ref.Append("create_time"))
}

// Database returns a reference to field database of google_biglake_table.
func (bt biglakeTableAttributes) Database() terra.StringValue {
	return terra.ReferenceAsString(bt.ref.Append("database"))
}

// DeleteTime returns a reference to field delete_time of google_biglake_table.
func (bt biglakeTableAttributes) DeleteTime() terra.StringValue {
	return terra.ReferenceAsString(bt.ref.Append("delete_time"))
}

// Etag returns a reference to field etag of google_biglake_table.
func (bt biglakeTableAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(bt.ref.Append("etag"))
}

// ExpireTime returns a reference to field expire_time of google_biglake_table.
func (bt biglakeTableAttributes) ExpireTime() terra.StringValue {
	return terra.ReferenceAsString(bt.ref.Append("expire_time"))
}

// Id returns a reference to field id of google_biglake_table.
func (bt biglakeTableAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(bt.ref.Append("id"))
}

// Name returns a reference to field name of google_biglake_table.
func (bt biglakeTableAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(bt.ref.Append("name"))
}

// Type returns a reference to field type of google_biglake_table.
func (bt biglakeTableAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(bt.ref.Append("type"))
}

// UpdateTime returns a reference to field update_time of google_biglake_table.
func (bt biglakeTableAttributes) UpdateTime() terra.StringValue {
	return terra.ReferenceAsString(bt.ref.Append("update_time"))
}

func (bt biglakeTableAttributes) HiveOptions() terra.ListValue[biglaketable.HiveOptionsAttributes] {
	return terra.ReferenceAsList[biglaketable.HiveOptionsAttributes](bt.ref.Append("hive_options"))
}

func (bt biglakeTableAttributes) Timeouts() biglaketable.TimeoutsAttributes {
	return terra.ReferenceAsSingle[biglaketable.TimeoutsAttributes](bt.ref.Append("timeouts"))
}

type biglakeTableState struct {
	CreateTime  string                          `json:"create_time"`
	Database    string                          `json:"database"`
	DeleteTime  string                          `json:"delete_time"`
	Etag        string                          `json:"etag"`
	ExpireTime  string                          `json:"expire_time"`
	Id          string                          `json:"id"`
	Name        string                          `json:"name"`
	Type        string                          `json:"type"`
	UpdateTime  string                          `json:"update_time"`
	HiveOptions []biglaketable.HiveOptionsState `json:"hive_options"`
	Timeouts    *biglaketable.TimeoutsState     `json:"timeouts"`
}
