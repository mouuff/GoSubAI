package generator

import "github.com/mouuff/GoSubAI/pkg/types"

type GenerationType int32

const (
	APPEND_GEN       GenerationType = 0
	REPLACE_WITH_GEN GenerationType = 1
)

type SubtitleGenerator struct {
	SubstitleData  *types.SubtitleData
	Prompt         string
	Model          string
	GenerationType GenerationType
}
