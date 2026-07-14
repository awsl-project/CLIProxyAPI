package gemini

import (
	. "github.com/awsl-project/CLIProxyAPI/v7/internal/constant"
	"github.com/awsl-project/CLIProxyAPI/v7/internal/interfaces"
	"github.com/awsl-project/CLIProxyAPI/v7/internal/translator/translator"
)

func init() {
	translator.Register(
		Gemini,
		Claude,
		ConvertGeminiRequestToClaude,
		interfaces.TranslateResponse{
			Stream:     ConvertClaudeResponseToGemini,
			NonStream:  ConvertClaudeResponseToGeminiNonStream,
			TokenCount: GeminiTokenCount,
		},
	)
}
