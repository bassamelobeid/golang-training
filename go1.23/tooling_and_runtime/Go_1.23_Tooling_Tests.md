
# Go 1.23 Tooling Tests

This guide demonstrates how to test and explore the new tooling features introduced in Go 1.23.

## Contents

1. [Telemetry](#telemetry)
2. [Go Command](#go-command)
3. [Trace Tool](#trace-tool)
4. [Vet Tool](#vet-tool)
5. [Cgo Enhancements](#cgo-enhancements)

## Telemetry

### How to Test

1. Run the following command to check telemetry status:
   ```bash
   go telemetry
   ```

2. Enable telemetry:
   ```bash
   go telemetry on
   ```

3. Check the telemetry data collected:
   ```bash
   go telemetry local
   ```

4. Disable telemetry:
   ```bash
   go telemetry off
   ```

### Expected Behavior

Telemetry data should only be collected when it is explicitly enabled. By default, telemetry is off.

## Go Command

### How to Test

1. **`go env -changed`**: Modify any Go environment variable and then run:
   ```bash
   go env -changed
   ```

2. **`go mod tidy -diff`**: Run in a Go module directory:
   ```bash
   go mod tidy -diff
   ```

3. **`go list -m -json`**: Run in a Go module directory:
   ```bash
   go list -m -json all
   ```

4. **`go mod download -json`**: Run in a Go module directory:
   ```bash
   go mod download -json
   ```

5. **`godebug` Directive**: Add a `godebug` directive to your `go.mod` or `go.work` file:
   ```go
   // go.mod example
   module example.com/my/module
   godebug directive
   ```

6. **GOROOT_FINAL**: Run with or without the `GOROOT_FINAL` environment variable and observe that it has no effect.

### Expected Behavior

These commands should behave as expected, showing changes only when relevant and not modifying files unless specified.

## Trace Tool

### How to Test

1. Run a Go program with trace enabled:
   ```bash
   go run -trace trace.out main.go
   ```

2. Manually corrupt the trace file and then try to load it:
   ```bash
   go tool trace trace.out
   ```

### Expected Behavior

The trace tool should attempt to recover whatever data it can from the broken trace.

## Vet Tool

### How to Test

1. Write Go code that references symbols from a newer Go version.
2. Run `go vet` on the code:
   ```bash
   go vet
   ```

### Expected Behavior

`go vet` should flag symbols that are too new for the effective Go version of the file.

## Cgo Enhancements

### How to Test

1. Write a simple C program and use `cgo` in Go to interact with it.
2. Use the `-ldflags` option to pass linker flags:
   ```bash
   go build -ldflags "-X main.variableName=value" main.go
   ```

### Expected Behavior

The Go command should automatically use `-ldflags` as specified.

---

This guide is meant to provide a hands-on approach to exploring the new features and changes in Go 1.23. Make sure you are running Go 1.23+ to fully utilize these features.
