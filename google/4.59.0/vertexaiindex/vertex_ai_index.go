// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package vertexaiindex

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type DeployedIndexes struct{}

type IndexStats struct{}

type Metadata struct {
	// ContentsDeltaUri: string, optional
	ContentsDeltaUri terra.StringValue `hcl:"contents_delta_uri,attr"`
	// IsCompleteOverwrite: bool, optional
	IsCompleteOverwrite terra.BoolValue `hcl:"is_complete_overwrite,attr"`
	// Config: optional
	Config *Config `hcl:"config,block"`
}

type Config struct {
	// ApproximateNeighborsCount: number, optional
	ApproximateNeighborsCount terra.NumberValue `hcl:"approximate_neighbors_count,attr"`
	// Dimensions: number, required
	Dimensions terra.NumberValue `hcl:"dimensions,attr" validate:"required"`
	// DistanceMeasureType: string, optional
	DistanceMeasureType terra.StringValue `hcl:"distance_measure_type,attr"`
	// FeatureNormType: string, optional
	FeatureNormType terra.StringValue `hcl:"feature_norm_type,attr"`
	// AlgorithmConfig: optional
	AlgorithmConfig *AlgorithmConfig `hcl:"algorithm_config,block"`
}

type AlgorithmConfig struct {
	// BruteForceConfig: optional
	BruteForceConfig *BruteForceConfig `hcl:"brute_force_config,block"`
	// TreeAhConfig: optional
	TreeAhConfig *TreeAhConfig `hcl:"tree_ah_config,block"`
}

type BruteForceConfig struct{}

type TreeAhConfig struct {
	// LeafNodeEmbeddingCount: number, optional
	LeafNodeEmbeddingCount terra.NumberValue `hcl:"leaf_node_embedding_count,attr"`
	// LeafNodesToSearchPercent: number, optional
	LeafNodesToSearchPercent terra.NumberValue `hcl:"leaf_nodes_to_search_percent,attr"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type DeployedIndexesAttributes struct {
	ref terra.Reference
}

func (di DeployedIndexesAttributes) InternalRef() terra.Reference {
	return di.ref
}

func (di DeployedIndexesAttributes) InternalWithRef(ref terra.Reference) DeployedIndexesAttributes {
	return DeployedIndexesAttributes{ref: ref}
}

func (di DeployedIndexesAttributes) InternalTokens() hclwrite.Tokens {
	return di.ref.InternalTokens()
}

func (di DeployedIndexesAttributes) DeployedIndexId() terra.StringValue {
	return terra.ReferenceString(di.ref.Append("deployed_index_id"))
}

func (di DeployedIndexesAttributes) IndexEndpoint() terra.StringValue {
	return terra.ReferenceString(di.ref.Append("index_endpoint"))
}

type IndexStatsAttributes struct {
	ref terra.Reference
}

func (is IndexStatsAttributes) InternalRef() terra.Reference {
	return is.ref
}

func (is IndexStatsAttributes) InternalWithRef(ref terra.Reference) IndexStatsAttributes {
	return IndexStatsAttributes{ref: ref}
}

func (is IndexStatsAttributes) InternalTokens() hclwrite.Tokens {
	return is.ref.InternalTokens()
}

func (is IndexStatsAttributes) ShardsCount() terra.NumberValue {
	return terra.ReferenceNumber(is.ref.Append("shards_count"))
}

func (is IndexStatsAttributes) VectorsCount() terra.StringValue {
	return terra.ReferenceString(is.ref.Append("vectors_count"))
}

type MetadataAttributes struct {
	ref terra.Reference
}

func (m MetadataAttributes) InternalRef() terra.Reference {
	return m.ref
}

func (m MetadataAttributes) InternalWithRef(ref terra.Reference) MetadataAttributes {
	return MetadataAttributes{ref: ref}
}

func (m MetadataAttributes) InternalTokens() hclwrite.Tokens {
	return m.ref.InternalTokens()
}

func (m MetadataAttributes) ContentsDeltaUri() terra.StringValue {
	return terra.ReferenceString(m.ref.Append("contents_delta_uri"))
}

func (m MetadataAttributes) IsCompleteOverwrite() terra.BoolValue {
	return terra.ReferenceBool(m.ref.Append("is_complete_overwrite"))
}

func (m MetadataAttributes) Config() terra.ListValue[ConfigAttributes] {
	return terra.ReferenceList[ConfigAttributes](m.ref.Append("config"))
}

type ConfigAttributes struct {
	ref terra.Reference
}

func (c ConfigAttributes) InternalRef() terra.Reference {
	return c.ref
}

func (c ConfigAttributes) InternalWithRef(ref terra.Reference) ConfigAttributes {
	return ConfigAttributes{ref: ref}
}

func (c ConfigAttributes) InternalTokens() hclwrite.Tokens {
	return c.ref.InternalTokens()
}

func (c ConfigAttributes) ApproximateNeighborsCount() terra.NumberValue {
	return terra.ReferenceNumber(c.ref.Append("approximate_neighbors_count"))
}

func (c ConfigAttributes) Dimensions() terra.NumberValue {
	return terra.ReferenceNumber(c.ref.Append("dimensions"))
}

func (c ConfigAttributes) DistanceMeasureType() terra.StringValue {
	return terra.ReferenceString(c.ref.Append("distance_measure_type"))
}

func (c ConfigAttributes) FeatureNormType() terra.StringValue {
	return terra.ReferenceString(c.ref.Append("feature_norm_type"))
}

func (c ConfigAttributes) AlgorithmConfig() terra.ListValue[AlgorithmConfigAttributes] {
	return terra.ReferenceList[AlgorithmConfigAttributes](c.ref.Append("algorithm_config"))
}

type AlgorithmConfigAttributes struct {
	ref terra.Reference
}

func (ac AlgorithmConfigAttributes) InternalRef() terra.Reference {
	return ac.ref
}

func (ac AlgorithmConfigAttributes) InternalWithRef(ref terra.Reference) AlgorithmConfigAttributes {
	return AlgorithmConfigAttributes{ref: ref}
}

func (ac AlgorithmConfigAttributes) InternalTokens() hclwrite.Tokens {
	return ac.ref.InternalTokens()
}

func (ac AlgorithmConfigAttributes) BruteForceConfig() terra.ListValue[BruteForceConfigAttributes] {
	return terra.ReferenceList[BruteForceConfigAttributes](ac.ref.Append("brute_force_config"))
}

func (ac AlgorithmConfigAttributes) TreeAhConfig() terra.ListValue[TreeAhConfigAttributes] {
	return terra.ReferenceList[TreeAhConfigAttributes](ac.ref.Append("tree_ah_config"))
}

type BruteForceConfigAttributes struct {
	ref terra.Reference
}

func (bfc BruteForceConfigAttributes) InternalRef() terra.Reference {
	return bfc.ref
}

func (bfc BruteForceConfigAttributes) InternalWithRef(ref terra.Reference) BruteForceConfigAttributes {
	return BruteForceConfigAttributes{ref: ref}
}

func (bfc BruteForceConfigAttributes) InternalTokens() hclwrite.Tokens {
	return bfc.ref.InternalTokens()
}

type TreeAhConfigAttributes struct {
	ref terra.Reference
}

func (tac TreeAhConfigAttributes) InternalRef() terra.Reference {
	return tac.ref
}

func (tac TreeAhConfigAttributes) InternalWithRef(ref terra.Reference) TreeAhConfigAttributes {
	return TreeAhConfigAttributes{ref: ref}
}

func (tac TreeAhConfigAttributes) InternalTokens() hclwrite.Tokens {
	return tac.ref.InternalTokens()
}

func (tac TreeAhConfigAttributes) LeafNodeEmbeddingCount() terra.NumberValue {
	return terra.ReferenceNumber(tac.ref.Append("leaf_node_embedding_count"))
}

func (tac TreeAhConfigAttributes) LeafNodesToSearchPercent() terra.NumberValue {
	return terra.ReferenceNumber(tac.ref.Append("leaf_nodes_to_search_percent"))
}

type TimeoutsAttributes struct {
	ref terra.Reference
}

func (t TimeoutsAttributes) InternalRef() terra.Reference {
	return t.ref
}

func (t TimeoutsAttributes) InternalWithRef(ref terra.Reference) TimeoutsAttributes {
	return TimeoutsAttributes{ref: ref}
}

func (t TimeoutsAttributes) InternalTokens() hclwrite.Tokens {
	return t.ref.InternalTokens()
}

func (t TimeoutsAttributes) Create() terra.StringValue {
	return terra.ReferenceString(t.ref.Append("create"))
}

func (t TimeoutsAttributes) Delete() terra.StringValue {
	return terra.ReferenceString(t.ref.Append("delete"))
}

func (t TimeoutsAttributes) Update() terra.StringValue {
	return terra.ReferenceString(t.ref.Append("update"))
}

type DeployedIndexesState struct {
	DeployedIndexId string `json:"deployed_index_id"`
	IndexEndpoint   string `json:"index_endpoint"`
}

type IndexStatsState struct {
	ShardsCount  float64 `json:"shards_count"`
	VectorsCount string  `json:"vectors_count"`
}

type MetadataState struct {
	ContentsDeltaUri    string        `json:"contents_delta_uri"`
	IsCompleteOverwrite bool          `json:"is_complete_overwrite"`
	Config              []ConfigState `json:"config"`
}

type ConfigState struct {
	ApproximateNeighborsCount float64                `json:"approximate_neighbors_count"`
	Dimensions                float64                `json:"dimensions"`
	DistanceMeasureType       string                 `json:"distance_measure_type"`
	FeatureNormType           string                 `json:"feature_norm_type"`
	AlgorithmConfig           []AlgorithmConfigState `json:"algorithm_config"`
}

type AlgorithmConfigState struct {
	BruteForceConfig []BruteForceConfigState `json:"brute_force_config"`
	TreeAhConfig     []TreeAhConfigState     `json:"tree_ah_config"`
}

type BruteForceConfigState struct{}

type TreeAhConfigState struct {
	LeafNodeEmbeddingCount   float64 `json:"leaf_node_embedding_count"`
	LeafNodesToSearchPercent float64 `json:"leaf_nodes_to_search_percent"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}
