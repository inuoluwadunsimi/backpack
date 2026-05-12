// Package exec provides a shell command runner with proper context handling,
// PATH enrichment, and edge-case detection (sudo, TTY, hangs).
package exec

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

// RunResult captures the output of a shell command.
type RunResult struct {
	Stdout   string
	Stderr   string
	ExitCode int
	Duration time.Duration
}

// Runner abstracts shell command execution.
type Runner interface {
	// Run executes a command with arguments. Non-zero exit codes are returned
	// in RunResult, not as errors. Errors are reserved for system failures
	// (command not found, context timeout, etc.).
	Run(ctx context.Context, cmd string, args ...string) (*RunResult, error)

	// RunShell executes a shell script via /bin/sh -c.
	RunShell(ctx context.Context, script string) (*RunResult, error)

	// Which checks if a tool is available on the enriched PATH.
	// Returns the resolved absolute path and true, or ("", false).
	Which(tool string) (string, bool)
}

// DefaultRunner implements Runner with enriched PATH and edge-case handling.
type DefaultRunner struct {
	pathDirs []string
	env      []string
}

// NewRunner creates a Runner with an enriched PATH that covers common
// dev tool locations (Homebrew, nvm, pyenv, cargo, etc.).
func NewRunner() *DefaultRunner {
	dirs := buildPATH()
	return &DefaultRunner{
		pathDirs: dirs,
		env:      buildEnv(dirs),
	}
}

func (r *DefaultRunner) Run(ctx context.Context, cmd string, args ...string) (*RunResult, error) {
	start := time.Now()

	c := exec.CommandContext(ctx, cmd, args...)
	c.Env = r.env

	var stdout, stderr bytes.Buffer
	c.Stdout = &stdout
	c.Stderr = &stderr
	// Intentionally not setting c.Stdin — prevents hanging on commands
	// that expect interactive TTY input. The context timeout handles
	// commands that block waiting for input.

	err := c.Run()
	duration := time.Since(start)

	if err != nil {
		// Context errors (timeout / cancellation) are real errors.
		if ctx.Err() != nil {
			return nil, fmt.Errorf("command %s: %w", cmd, ctx.Err())
		}

		// Non-zero exit code is not a system error — return it in RunResult.
		var exitErr *exec.ExitError
		if errors.As(err, &exitErr) {
			return &RunResult{
				Stdout:   stdout.String(),
				Stderr:   stderr.String(),
				ExitCode: exitErr.ExitCode(),
				Duration: duration,
			}, nil
		}

		// Anything else (command not found, permission denied to execute, etc.)
		return nil, fmt.Errorf("executing %s: %w", cmd, err)
	}

	return &RunResult{
		Stdout:   stdout.String(),
		Stderr:   stderr.String(),
		ExitCode: 0,
		Duration: duration,
	}, nil
}

func (r *DefaultRunner) RunShell(ctx context.Context, script string) (*RunResult, error) {
	return r.Run(ctx, "/bin/sh", "-c", script)
}

func (r *DefaultRunner) Which(tool string) (string, bool) {
	if filepath.IsAbs(tool) {
		if isExecutable(tool) {
			return tool, true
		}
		return "", false
	}

	for _, dir := range r.pathDirs {
		path := filepath.Join(dir, tool)
		if isExecutable(path) {
			return path, true
		}
	}
	return "", false
}

// NeedsSudo checks if a RunResult indicates the command needs elevated privileges.
func NeedsSudo(r *RunResult) bool {
	combined := strings.ToLower(r.Stderr + r.Stdout)
	return strings.Contains(combined, "permission denied") ||
		strings.Contains(combined, "operation not permitted") ||
		strings.Contains(combined, "access denied")
}

// ── PATH helpers ──

// buildPATH returns an ordered list of directories that covers common dev tool
// install locations plus the existing system PATH.
func buildPATH() []string {
	home, _ := os.UserHomeDir()

	candidates := []string{
		filepath.Join(home, ".local/bin"),
		// Homebrew — Apple Silicon
		"/opt/homebrew/bin",
		"/opt/homebrew/sbin",
		// Homebrew — Intel
		"/usr/local/bin",
		"/usr/local/sbin",
		// Version managers
		filepath.Join(home, ".pyenv/shims"),
		filepath.Join(home, ".rbenv/shims"),
		filepath.Join(home, ".nodenv/shims"),
		// Language toolchains
		filepath.Join(home, ".cargo/bin"),
		filepath.Join(home, "go/bin"),
		filepath.Join(home, ".deno/bin"),
		// System
		"/usr/bin",
		"/bin",
		"/usr/sbin",
		"/sbin",
	}

	seen := make(map[string]bool)
	var dirs []string

	for _, dir := range candidates {
		if seen[dir] {
			continue
		}
		if info, err := os.Stat(dir); err == nil && info.IsDir() {
			dirs = append(dirs, dir)
			seen[dir] = true
		}
	}

	// Append existing PATH entries not already covered.
	for _, dir := range filepath.SplitList(os.Getenv("PATH")) {
		if !seen[dir] {
			dirs = append(dirs, dir)
			seen[dir] = true
		}
	}

	return dirs
}

func buildEnv(pathDirs []string) []string {
	env := os.Environ()
	pathVal := "PATH=" + strings.Join(pathDirs, string(os.PathListSeparator))

	for i, e := range env {
		if strings.HasPrefix(e, "PATH=") {
			env[i] = pathVal
			return env
		}
	}
	return append(env, pathVal)
}

func isExecutable(path string) bool {
	info, err := os.Stat(path)
	return err == nil && !info.IsDir() && info.Mode()&0111 != 0
}
