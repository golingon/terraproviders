// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	gluecatalogtable "github.com/golingon/terraproviders/aws/4.66.1/gluecatalogtable"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewGlueCatalogTable creates a new instance of [GlueCatalogTable].
func NewGlueCatalogTable(name string, args GlueCatalogTableArgs) *GlueCatalogTable {
	return &GlueCatalogTable{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*GlueCatalogTable)(nil)

// GlueCatalogTable represents the Terraform resource aws_glue_catalog_table.
type GlueCatalogTable struct {
	Name      string
	Args      GlueCatalogTableArgs
	state     *glueCatalogTableState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [GlueCatalogTable].
func (gct *GlueCatalogTable) Type() string {
	return "aws_glue_catalog_table"
}

// LocalName returns the local name for [GlueCatalogTable].
func (gct *GlueCatalogTable) LocalName() string {
	return gct.Name
}

// Configuration returns the configuration (args) for [GlueCatalogTable].
func (gct *GlueCatalogTable) Configuration() interface{} {
	return gct.Args
}

// DependOn is used for other resources to depend on [GlueCatalogTable].
func (gct *GlueCatalogTable) DependOn() terra.Reference {
	return terra.ReferenceResource(gct)
}

// Dependencies returns the list of resources [GlueCatalogTable] depends_on.
func (gct *GlueCatalogTable) Dependencies() terra.Dependencies {
	return gct.DependsOn
}

// LifecycleManagement returns the lifecycle block for [GlueCatalogTable].
func (gct *GlueCatalogTable) LifecycleManagement() *terra.Lifecycle {
	return gct.Lifecycle
}

// Attributes returns the attributes for [GlueCatalogTable].
func (gct *GlueCatalogTable) Attributes() glueCatalogTableAttributes {
	return glueCatalogTableAttributes{ref: terra.ReferenceResource(gct)}
}

// ImportState imports the given attribute values into [GlueCatalogTable]'s state.
func (gct *GlueCatalogTable) ImportState(av io.Reader) error {
	gct.state = &glueCatalogTableState{}
	if err := json.NewDecoder(av).Decode(gct.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", gct.Type(), gct.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [GlueCatalogTable] has state.
func (gct *GlueCatalogTable) State() (*glueCatalogTableState, bool) {
	return gct.state, gct.state != nil
}

// StateMust returns the state for [GlueCatalogTable]. Panics if the state is nil.
func (gct *GlueCatalogTable) StateMust() *glueCatalogTableState {
	if gct.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", gct.Type(), gct.LocalName()))
	}
	return gct.state
}

// GlueCatalogTableArgs contains the configurations for aws_glue_catalog_table.
type GlueCatalogTableArgs struct {
	// CatalogId: string, optional
	CatalogId terra.StringValue `hcl:"catalog_id,attr"`
	// DatabaseName: string, required
	DatabaseName terra.StringValue `hcl:"database_name,attr" validate:"required"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Owner: string, optional
	Owner terra.StringValue `hcl:"owner,attr"`
	// Parameters: map of string, optional
	Parameters terra.MapValue[terra.StringValue] `hcl:"parameters,attr"`
	// Retention: number, optional
	Retention terra.NumberValue `hcl:"retention,attr"`
	// TableType: string, optional
	TableType terra.StringValue `hcl:"table_type,attr"`
	// ViewExpandedText: string, optional
	ViewExpandedText terra.StringValue `hcl:"view_expanded_text,attr"`
	// ViewOriginalText: string, optional
	ViewOriginalText terra.StringValue `hcl:"view_original_text,attr"`
	// PartitionIndex: min=0,max=3
	PartitionIndex []gluecatalogtable.PartitionIndex `hcl:"partition_index,block" validate:"min=0,max=3"`
	// PartitionKeys: min=0
	PartitionKeys []gluecatalogtable.PartitionKeys `hcl:"partition_keys,block" validate:"min=0"`
	// StorageDescriptor: optional
	StorageDescriptor *gluecatalogtable.StorageDescriptor `hcl:"storage_descriptor,block"`
	// TargetTable: optional
	TargetTable *gluecatalogtable.TargetTable `hcl:"target_table,block"`
}
type glueCatalogTableAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_glue_catalog_table.
func (gct glueCatalogTableAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(gct.ref.Append("arn"))
}

// CatalogId returns a reference to field catalog_id of aws_glue_catalog_table.
func (gct glueCatalogTableAttributes) CatalogId() terra.StringValue {
	return terra.ReferenceAsString(gct.ref.Append("catalog_id"))
}

// DatabaseName returns a reference to field database_name of aws_glue_catalog_table.
func (gct glueCatalogTableAttributes) DatabaseName() terra.StringValue {
	return terra.ReferenceAsString(gct.ref.Append("database_name"))
}

// Description returns a reference to field description of aws_glue_catalog_table.
func (gct glueCatalogTableAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(gct.ref.Append("description"))
}

// Id returns a reference to field id of aws_glue_catalog_table.
func (gct glueCatalogTableAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gct.ref.Append("id"))
}

// Name returns a reference to field name of aws_glue_catalog_table.
func (gct glueCatalogTableAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(gct.ref.Append("name"))
}

// Owner returns a reference to field owner of aws_glue_catalog_table.
func (gct glueCatalogTableAttributes) Owner() terra.StringValue {
	return terra.ReferenceAsString(gct.ref.Append("owner"))
}

// Parameters returns a reference to field parameters of aws_glue_catalog_table.
func (gct glueCatalogTableAttributes) Parameters() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](gct.ref.Append("parameters"))
}

// Retention returns a reference to field retention of aws_glue_catalog_table.
func (gct glueCatalogTableAttributes) Retention() terra.NumberValue {
	return terra.ReferenceAsNumber(gct.ref.Append("retention"))
}

// TableType returns a reference to field table_type of aws_glue_catalog_table.
func (gct glueCatalogTableAttributes) TableType() terra.StringValue {
	return terra.ReferenceAsString(gct.ref.Append("table_type"))
}

// ViewExpandedText returns a reference to field view_expanded_text of aws_glue_catalog_table.
func (gct glueCatalogTableAttributes) ViewExpandedText() terra.StringValue {
	return terra.ReferenceAsString(gct.ref.Append("view_expanded_text"))
}

// ViewOriginalText returns a reference to field view_original_text of aws_glue_catalog_table.
func (gct glueCatalogTableAttributes) ViewOriginalText() terra.StringValue {
	return terra.ReferenceAsString(gct.ref.Append("view_original_text"))
}

func (gct glueCatalogTableAttributes) PartitionIndex() terra.ListValue[gluecatalogtable.PartitionIndexAttributes] {
	return terra.ReferenceAsList[gluecatalogtable.PartitionIndexAttributes](gct.ref.Append("partition_index"))
}

func (gct glueCatalogTableAttributes) PartitionKeys() terra.ListValue[gluecatalogtable.PartitionKeysAttributes] {
	return terra.ReferenceAsList[gluecatalogtable.PartitionKeysAttributes](gct.ref.Append("partition_keys"))
}

func (gct glueCatalogTableAttributes) StorageDescriptor() terra.ListValue[gluecatalogtable.StorageDescriptorAttributes] {
	return terra.ReferenceAsList[gluecatalogtable.StorageDescriptorAttributes](gct.ref.Append("storage_descriptor"))
}

func (gct glueCatalogTableAttributes) TargetTable() terra.ListValue[gluecatalogtable.TargetTableAttributes] {
	return terra.ReferenceAsList[gluecatalogtable.TargetTableAttributes](gct.ref.Append("target_table"))
}

type glueCatalogTableState struct {
	Arn               string                                    `json:"arn"`
	CatalogId         string                                    `json:"catalog_id"`
	DatabaseName      string                                    `json:"database_name"`
	Description       string                                    `json:"description"`
	Id                string                                    `json:"id"`
	Name              string                                    `json:"name"`
	Owner             string                                    `json:"owner"`
	Parameters        map[string]string                         `json:"parameters"`
	Retention         float64                                   `json:"retention"`
	TableType         string                                    `json:"table_type"`
	ViewExpandedText  string                                    `json:"view_expanded_text"`
	ViewOriginalText  string                                    `json:"view_original_text"`
	PartitionIndex    []gluecatalogtable.PartitionIndexState    `json:"partition_index"`
	PartitionKeys     []gluecatalogtable.PartitionKeysState     `json:"partition_keys"`
	StorageDescriptor []gluecatalogtable.StorageDescriptorState `json:"storage_descriptor"`
	TargetTable       []gluecatalogtable.TargetTableState       `json:"target_table"`
}
