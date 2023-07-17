// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	datacatalogtag "github.com/golingon/terraproviders/googlebeta/4.73.1/datacatalogtag"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDataCatalogTag creates a new instance of [DataCatalogTag].
func NewDataCatalogTag(name string, args DataCatalogTagArgs) *DataCatalogTag {
	return &DataCatalogTag{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DataCatalogTag)(nil)

// DataCatalogTag represents the Terraform resource google_data_catalog_tag.
type DataCatalogTag struct {
	Name      string
	Args      DataCatalogTagArgs
	state     *dataCatalogTagState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DataCatalogTag].
func (dct *DataCatalogTag) Type() string {
	return "google_data_catalog_tag"
}

// LocalName returns the local name for [DataCatalogTag].
func (dct *DataCatalogTag) LocalName() string {
	return dct.Name
}

// Configuration returns the configuration (args) for [DataCatalogTag].
func (dct *DataCatalogTag) Configuration() interface{} {
	return dct.Args
}

// DependOn is used for other resources to depend on [DataCatalogTag].
func (dct *DataCatalogTag) DependOn() terra.Reference {
	return terra.ReferenceResource(dct)
}

// Dependencies returns the list of resources [DataCatalogTag] depends_on.
func (dct *DataCatalogTag) Dependencies() terra.Dependencies {
	return dct.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DataCatalogTag].
func (dct *DataCatalogTag) LifecycleManagement() *terra.Lifecycle {
	return dct.Lifecycle
}

// Attributes returns the attributes for [DataCatalogTag].
func (dct *DataCatalogTag) Attributes() dataCatalogTagAttributes {
	return dataCatalogTagAttributes{ref: terra.ReferenceResource(dct)}
}

// ImportState imports the given attribute values into [DataCatalogTag]'s state.
func (dct *DataCatalogTag) ImportState(av io.Reader) error {
	dct.state = &dataCatalogTagState{}
	if err := json.NewDecoder(av).Decode(dct.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dct.Type(), dct.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DataCatalogTag] has state.
func (dct *DataCatalogTag) State() (*dataCatalogTagState, bool) {
	return dct.state, dct.state != nil
}

// StateMust returns the state for [DataCatalogTag]. Panics if the state is nil.
func (dct *DataCatalogTag) StateMust() *dataCatalogTagState {
	if dct.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dct.Type(), dct.LocalName()))
	}
	return dct.state
}

// DataCatalogTagArgs contains the configurations for google_data_catalog_tag.
type DataCatalogTagArgs struct {
	// Column: string, optional
	Column terra.StringValue `hcl:"column,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Parent: string, optional
	Parent terra.StringValue `hcl:"parent,attr"`
	// Template: string, required
	Template terra.StringValue `hcl:"template,attr" validate:"required"`
	// Fields: min=1
	Fields []datacatalogtag.Fields `hcl:"fields,block" validate:"min=1"`
	// Timeouts: optional
	Timeouts *datacatalogtag.Timeouts `hcl:"timeouts,block"`
}
type dataCatalogTagAttributes struct {
	ref terra.Reference
}

// Column returns a reference to field column of google_data_catalog_tag.
func (dct dataCatalogTagAttributes) Column() terra.StringValue {
	return terra.ReferenceAsString(dct.ref.Append("column"))
}

// Id returns a reference to field id of google_data_catalog_tag.
func (dct dataCatalogTagAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dct.ref.Append("id"))
}

// Name returns a reference to field name of google_data_catalog_tag.
func (dct dataCatalogTagAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(dct.ref.Append("name"))
}

// Parent returns a reference to field parent of google_data_catalog_tag.
func (dct dataCatalogTagAttributes) Parent() terra.StringValue {
	return terra.ReferenceAsString(dct.ref.Append("parent"))
}

// Template returns a reference to field template of google_data_catalog_tag.
func (dct dataCatalogTagAttributes) Template() terra.StringValue {
	return terra.ReferenceAsString(dct.ref.Append("template"))
}

// TemplateDisplayname returns a reference to field template_displayname of google_data_catalog_tag.
func (dct dataCatalogTagAttributes) TemplateDisplayname() terra.StringValue {
	return terra.ReferenceAsString(dct.ref.Append("template_displayname"))
}

func (dct dataCatalogTagAttributes) Fields() terra.SetValue[datacatalogtag.FieldsAttributes] {
	return terra.ReferenceAsSet[datacatalogtag.FieldsAttributes](dct.ref.Append("fields"))
}

func (dct dataCatalogTagAttributes) Timeouts() datacatalogtag.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datacatalogtag.TimeoutsAttributes](dct.ref.Append("timeouts"))
}

type dataCatalogTagState struct {
	Column              string                        `json:"column"`
	Id                  string                        `json:"id"`
	Name                string                        `json:"name"`
	Parent              string                        `json:"parent"`
	Template            string                        `json:"template"`
	TemplateDisplayname string                        `json:"template_displayname"`
	Fields              []datacatalogtag.FieldsState  `json:"fields"`
	Timeouts            *datacatalogtag.TimeoutsState `json:"timeouts"`
}
