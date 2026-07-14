// Package exec provides SDK wrappers for CLIProxyAPI executor functionality.
// This package exposes internal executor implementations for external use.
package exec

import (
	"github.com/awsl-project/CLIProxyAPI/v7/internal/config"
	internal "github.com/awsl-project/CLIProxyAPI/v7/internal/runtime/executor"

	// Import builtin to register translators (claude -> antigravity, etc.)
	_ "github.com/awsl-project/CLIProxyAPI/v7/sdk/translator/builtin"
)

// AntigravityExecutor is a type alias for the internal AntigravityExecutor.
type AntigravityExecutor = internal.AntigravityExecutor

// NewAntigravityExecutor creates a new AntigravityExecutor instance.
func NewAntigravityExecutor() *AntigravityExecutor {
	return internal.NewAntigravityExecutor(&config.Config{})
}
