// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	bigtabletable "github.com/golingon/terraproviders/googlebeta/4.74.0/bigtabletable"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewBigtableTable creates a new instance of [BigtableTable].
func NewBigtableTable(name string, args BigtableTableArgs) *BigtableTable {
	return &BigtableTable{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*BigtableTable)(nil)

// BigtableTable represents the Terraform resource google_bigtable_table.
type BigtableTable struct {
	Name      string
	Args      BigtableTableArgs
	state     *bigtableTableState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [BigtableTable].
func (bt *BigtableTable) Type() string {
	return "google_bigtable_table"
}

// LocalName returns the local name for [BigtableTable].
func (bt *BigtableTable) LocalName() string {
	return bt.Name
}

// Configuration returns the configuration (args) for [BigtableTable].
func (bt *BigtableTable) Configuration() interface{} {
	return bt.Args
}

// DependOn is used for other resources to depend on [BigtableTable].
func (bt *BigtableTable) DependOn() terra.Reference {
	return terra.ReferenceResource(bt)
}

// Dependencies returns the list of resources [BigtableTable] depends_on.
func (bt *BigtableTable) Dependencies() terra.Dependencies {
	return bt.DependsOn
}

// LifecycleManagement returns the lifecycle block for [BigtableTable].
func (bt *BigtableTable) LifecycleManagement() *terra.Lifecycle {
	return bt.Lifecycle
}

// Attributes returns the attributes for [BigtableTable].
func (bt *BigtableTable) Attributes() bigtableTableAttributes {
	return bigtableTableAttributes{ref: terra.ReferenceResource(bt)}
}

// ImportState imports the given attribute values into [BigtableTable]'s state.
func (bt *BigtableTable) ImportState(av io.Reader) error {
	bt.state = &bigtableTableState{}
	if err := json.NewDecoder(av).Decode(bt.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", bt.Type(), bt.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [BigtableTable] has state.
func (bt *BigtableTable) State() (*bigtableTableState, bool) {
	return bt.state, bt.state != nil
}

// StateMust returns the state for [BigtableTable]. Panics if the state is nil.
func (bt *BigtableTable) StateMust() *bigtableTableState {
	if bt.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", bt.Type(), bt.LocalName()))
	}
	return bt.state
}

// BigtableTableArgs contains the configurations for google_bigtable_table.
type BigtableTableArgs struct {
	// DeletionProtection: string, optional
	DeletionProtection terra.StringValue `hcl:"deletion_protection,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// InstanceName: string, required
	InstanceName terra.StringValue `hcl:"instance_name,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// SplitKeys: list of string, optional
	SplitKeys terra.ListValue[terra.StringValue] `hcl:"split_keys,attr"`
	// ColumnFamily: min=0
	ColumnFamily []bigtabletable.ColumnFamily `hcl:"column_family,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *bigtabletable.Timeouts `hcl:"timeouts,block"`
}
type bigtableTableAttributes struct {
	ref terra.Reference
}

// DeletionProtection returns a reference to field deletion_protection of google_bigtable_table.
func (bt bigtableTableAttributes) DeletionProtection() terra.StringValue {
	return terra.ReferenceAsString(bt.ref.Append("deletion_protection"))
}

// Id returns a reference to field id of google_bigtable_table.
func (bt bigtableTableAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(bt.ref.Append("id"))
}

// InstanceName returns a reference to field instance_name of google_bigtable_table.
func (bt bigtableTableAttributes) InstanceName() terra.StringValue {
	return terra.ReferenceAsString(bt.ref.Append("instance_name"))
}

// Name returns a reference to field name of google_bigtable_table.
func (bt bigtableTableAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(bt.ref.Append("name"))
}

// Project returns a reference to field project of google_bigtable_table.
func (bt bigtableTableAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(bt.ref.Append("project"))
}

// SplitKeys returns a reference to field split_keys of google_bigtable_table.
func (bt bigtableTableAttributes) SplitKeys() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](bt.ref.Append("split_keys"))
}

func (bt bigtableTableAttributes) ColumnFamily() terra.SetValue[bigtabletable.ColumnFamilyAttributes] {
	return terra.ReferenceAsSet[bigtabletable.ColumnFamilyAttributes](bt.ref.Append("column_family"))
}

func (bt bigtableTableAttributes) Timeouts() bigtabletable.TimeoutsAttributes {
	return terra.ReferenceAsSingle[bigtabletable.TimeoutsAttributes](bt.ref.Append("timeouts"))
}

type bigtableTableState struct {
	DeletionProtection string                            `json:"deletion_protection"`
	Id                 string                            `json:"id"`
	InstanceName       string                            `json:"instance_name"`
	Name               string                            `json:"name"`
	Project            string                            `json:"project"`
	SplitKeys          []string                          `json:"split_keys"`
	ColumnFamily       []bigtabletable.ColumnFamilyState `json:"column_family"`
	Timeouts           *bigtabletable.TimeoutsState      `json:"timeouts"`
}
