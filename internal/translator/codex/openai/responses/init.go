package responses

import (
	. "github.com/awsl-project/CLIProxyAPI/v7/internal/constant"
	"github.com/awsl-project/CLIProxyAPI/v7/internal/interfaces"
	"github.com/awsl-project/CLIProxyAPI/v7/internal/translator/translator"
)

func init() {
	translator.Register(
		OpenaiResponse,
		Codex,
		ConvertOpenAIResponsesRequestToCodex,
		interfaces.TranslateResponse{
			Stream:    ConvertCodexResponseToOpenAIResponses,
			NonStream: ConvertCodexResponseToOpenAIResponsesNonStream,
		},
	)
}
