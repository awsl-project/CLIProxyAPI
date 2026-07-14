package responses

import (
	. "github.com/awsl-project/CLIProxyAPI/v7/internal/constant"
	"github.com/awsl-project/CLIProxyAPI/v7/internal/interfaces"
	"github.com/awsl-project/CLIProxyAPI/v7/internal/translator/translator"
)

func init() {
	translator.Register(
		OpenaiResponse,
		Interactions,
		ConvertOpenAIResponsesRequestToInteractions,
		interfaces.TranslateResponse{
			Stream:    ConvertInteractionsResponseToOpenAIResponses,
			NonStream: ConvertInteractionsResponseToOpenAIResponsesNonStream,
		},
	)
	translator.Register(
		Interactions,
		OpenaiResponse,
		ConvertInteractionsRequestToOpenAIResponses,
		interfaces.TranslateResponse{
			Stream:    ConvertOpenAIResponsesResponseToInteractions,
			NonStream: ConvertOpenAIResponsesResponseToInteractionsNonStream,
		},
	)
}
