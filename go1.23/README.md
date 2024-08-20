# Go 1.23 Features and Changes

This folder contains examples and tests demonstrating the new features and changes introduced in Go 1.23.

## Contents

1. [Canonical Values](#canonical-values)
2. [Iterators](#iterators)
3. [Timer Changes](#timer-changes)
4. [HTTP Cookies](#http-cookies)
5. [OS Operations](#os-operations)
6. [Slice Operations](#slice-operations)
7. [Atomic Operations](#atomic-operations)
8. [Tooling and Runtime](#tooling-and-runtime)
9. [Other stdlib changes](#other-stdlib-changes)

## Canonical Values

The `canonical_values` folder contains tests demonstrating memory usage improvements with string canonicalization:

- [unique_values_test.go](./canonical_values/unique_values_test.go): Compares memory usage with and without canonicalization.

## Iterators

The `iterators` folder showcases the new iterator features in Go 1.23:

- [slice_iterators_test.go](./iterators/slice_iterators_test.go): Examples of using new slice iterators like `All()`, `Values()`, `Backward()`, etc.
- [map_iterators_test.go](./iterators/map_iterators_test.go): Demonstrates new map iterators such as `All()`, `Keys()`, `Values()`, etc.
- [range_iterators_test.go](./iterators/range_iterators_test.go): Shows usage of new range iterators and custom iterators.
- [helper_functions.go](./iterators/helper_functions.go): Contains helper functions for the iterator tests.

## Timer Changes

The `timer_changes` folder illustrates changes in timer behavior in Go 1.23:

- [garbage_collection_test.go](./timer_changes/garbage_collection_test.go): Demonstrates improved garbage collection of timers.
- [stop_reset_behavior_test.go](./timer_changes/stop_reset_behavior_test.go): Shows changes in timer reset behavior.

## HTTP Cookies

The `http_cookies` folder contains tests for the new HTTP cookie features in Go 1.23:

- [http_cookies_test.go](./http_cookies/http_cookies_test.go): Demonstrates parsing of Cookie and Set-Cookie headers, including new fields like "Partitioned" and "Quoted".

## OS Operations

The `os_operations` folder contains tests for new OS-related features in Go 1.23:

- [copy_test.go](./os_operations/copy_test.go): Demonstrates the new `os.CopyFS` function for copying directories.

## Slice Operations

The `slice_operations` folder contains tests for new slice-related features in Go 1.23:

- [repeat_test.go](./slice_operations/repeat_test.go): Demonstrates the new `slices.Repeat` function.

## Atomic Operations

The `atomic_operations` folder contains tests for new atomic operations in Go 1.23:

- [atomic_test.go](./atomic_operations/atomic_test.go): Demonstrates the new atomic `And` and `Or` operations.

## Tooling and Runtime

The `tooling_and_runtime` folder contains tests for new tooling and runtime features in Go 1.23:

- [Go_1.23_Tooling_Tests.md](./tooling_and_runtime/Go_1.23_Tooling_Tests.md): Demonstrates how to test and explore the new tooling features introduced in Go 1.23.
- [panic_test.go](./tooling_and_runtime/panic_test.go): Tests the new `panic` behavior in Go 1.23.

## Other stdlib changes

  - [archive/tar](https://tip.golang.org/doc/go1.23#archivetarpkgarchivetar)
  - [crypto/tls](https://tip.golang.org/doc/go1.23#cryptotlspkgcryptotls)
  - [crypto/x509](https://tip.golang.org/doc/go1.23#cryptox509pkgcryptox509)
  - [database/sql](https://tip.golang.org/doc/go1.23#databasesqlpkgdatabasesql)
  - [debug/elf](https://tip.golang.org/doc/go1.23#debugelfpkgdebugelf)
  - [encoding/binary](https://tip.golang.org/doc/go1.23#encodingbinarypkgencodingbinary)
  - [go/ast](https://tip.golang.org/doc/go1.23#goastpkggoast)
  - [go/types](https://tip.golang.org/doc/go1.23#gotypespkggotypes)
  - [go/version](https://tip.golang.org/doc/go1.22#go/version)

  - [math/rand/v2](https://tip.golang.org/doc/go1.23#mathrandv2pkgmathrandv2)
  - [net](https://tip.golang.org/doc/go1.23#netpkgnet)
  - [net/http](https://tip.golang.org/doc/go1.23#nethttppkgnethttp)
  - [net/http/httptest](https://tip.golang.org/doc/go1.23#nethttphttptestpkgnethttphttptest)
  - [net/netip](https://tip.golang.org/doc/go1.23#netnetippkgnetnetip)
  - [os](https://tip.golang.org/doc/go1.23#ospkgos)
  - [path/filepath](https://tip.golang.org/doc/go1.23#pathfilepathpkgpathfilepath)
  - [reflect](https://tip.golang.org/doc/go1.23#reflectpkgreflect)
  - [runtime/debug](https://tip.golang.org/doc/go1.23#runtimedebugpkgruntimedebug)

  - [runtime/pprof](https://tip.golang.org/doc/go1.23#runtimepprofpkgruntimepprof)
  - [runtime/trace](https://tip.golang.org/doc/go1.23#runtimetracepkgruntimetrace)
  - [structs](https://tip.golang.org/doc/go1.23#new-structs-package)
  - [syscall](https://tip.golang.org/doc/go1.23#syscallpkgsyscall)
  - [testing/fstest](https://tip.golang.org/doc/go1.23#testingfstestpkgtestingfstest)
  - [text/template](https://tip.golang.org/doc/go1.23#texttemplatepkgtexttemplate)
  - [time](https://tip.golang.org/doc/go1.23#timepkgtime)
  - [unicode/utf16](https://tip.golang.org/doc/go1.23#unicodeutf16pkgunicodeutf16)


## Running the Tests

To run the tests for a specific feature, navigate to the corresponding folder and use the `go test` command. For example:
```bash
cd http_cookies
go test
```

To run all tests in the go1.23 folder, use:
```bash
go test ./...
```


## Notes

- Some tests may behave differently when run with Go 1.22 vs Go 1.23. Make sure to use the appropriate Go version when running the tests.
- The examples in this folder are meant to illustrate the new features and changes in Go 1.23. They may not cover all aspects of the new functionality.
- For more in-depth information, refer to the official [Go 1.23 release notes](https://tip.golang.org/doc/go1.23).

For more detailed information about Go 1.23 changes, refer to the official Go documentation and release notes.
