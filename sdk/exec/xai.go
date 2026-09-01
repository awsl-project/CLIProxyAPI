// Package exec provides SDK wrappers for CLIProxyAPI executor functionality.
// This file exposes the internal xAI/Grok executor for external use.
// Token management (access_token refresh) is handled by the executor using the
// OAuth metadata stored on the provided auth object.
package exec

import (
	"github.com/router-for-me/CLIProxyAPI/v7/internal/config"
	internal "github.com/router-for-me/CLIProxyAPI/v7/internal/runtime/executor"
)

// XAIExecutor is a type alias for the internal xAI/Grok executor.
type XAIExecutor = internal.XAIAutoExecutor

// NewXAIExecutor creates a new xAI/Grok executor instance.
func NewXAIExecutor() *XAIExecutor {
	return internal.NewXAIAutoExecutor(&config.Config{})
}
