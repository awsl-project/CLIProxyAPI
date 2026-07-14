// Package exec provides SDK wrappers for CLIProxyAPI executor functionality.
// This file exposes the internal CodexExecutor for external use.
// Token management (access_token refresh) is the caller's responsibility;
// the caller should set auth.Metadata["access_token"] before calling Execute/ExecuteStream.
package exec

import (
	"github.com/router-for-me/CLIProxyAPI/v7/internal/config"
	internal "github.com/router-for-me/CLIProxyAPI/v7/internal/runtime/executor"
)

// CodexExecutor is a type alias for the internal CodexExecutor.
type CodexExecutor = internal.CodexExecutor

// NewCodexExecutor creates a new CodexExecutor instance.
func NewCodexExecutor() *CodexExecutor {
	return internal.NewCodexExecutor(&config.Config{})
}
