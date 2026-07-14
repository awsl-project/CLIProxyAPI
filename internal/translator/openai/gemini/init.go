package gemini

import (
	. "github.com/awsl-project/CLIProxyAPI/v7/internal/constant"
	"github.com/awsl-project/CLIProxyAPI/v7/internal/interfaces"
	"github.com/awsl-project/CLIProxyAPI/v7/internal/translator/translator"
)

func init() {
	translator.Register(
		Gemini,
		OpenAI,
		ConvertGeminiRequestToOpenAI,
		interfaces.TranslateResponse{
			Stream:     ConvertOpenAIResponseToGemini,
			NonStream:  ConvertOpenAIResponseToGeminiNonStream,
			TokenCount: GeminiTokenCount,
		},
	)
}
