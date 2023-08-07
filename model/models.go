package model

type Model string

const (
	ModelCommand             Model = "command"
	ModelCommandNightly      Model = "command-nightly"
	ModelCommandLight        Model = "command-light"
	ModelCommandLightNightly Model = "command-light-nightly"
)

type EmbedModel string

const (
	EmbedModelEnglishLightV20 EmbedModel = "embed-english-light-v2.0"
	EmbedModelMultilingualV20 EmbedModel = "embed-multilingual-v2.0"
	EmbedModelEnglishV20      EmbedModel = "embed-english-v2.0"
)

var EmbedModelsSize = map[EmbedModel]int{
	EmbedModelEnglishLightV20: 1024,
	EmbedModelEnglishV20:      4096,
	EmbedModelMultilingualV20: 768,
}

type ReturnLikelihoods string

const (
	ReturnLikelihoodsGeneration ReturnLikelihoods = "GENERATION"
	ReturnLikelihoodsAll        ReturnLikelihoods = "ALL"
	ReturnLikelihoodsNone       ReturnLikelihoods = "NONE"
)

type LogitBias map[string]float64

type Truncate string

const (
	TruncateNone  Truncate = "NONE"
	TruncateStart Truncate = "START"
	TruncateEnd   Truncate = "END"
)

type Meta struct {
	ApiVersion ApiVersion `json:"api_version"`
	Warnings   []string   `json:"warnings"`
}

type ApiVersion struct {
	Version        string `json:"version"`
	IsDeprecated   bool   `json:"is_deprecated"`
	IsExperimental bool   `json:"is_experimental"`
}

type Generation struct {
	ID               string            `json:"id"`
	Text             string            `json:"text"`
	Index            *int              `json:"index,omitempty"`
	Likelihood       *float64          `json:"likelihood,omitempty"`
	TokenLikelihoods []TokenLikelihood `json:"token_likelihoods,omitempty"`
}

type TokenLikelihood struct {
	Token      string  `json:"token"`
	Likelihood float64 `json:"likelihood"`
}

type Example struct {
	Text  string `json:"text"`
	Label string `json:"label"`
}

type Classification struct {
	ID         string              `json:"id"`
	Input      *string             `json:"input,omitempty"`
	Prediction string              `json:"prediction"`
	Confidence float64             `json:"confidence"`
	Labels     ClassificationLabel `json:"labels"`
}

type ClassificationLabel map[string]ClassificationLabelConfidence

type ClassificationLabelConfidence struct {
	Confidence float64 `json:"confidence"`
}