# oso-poc

This repository contains a minimal example that attempts to build a Go test
including the Go libary for [Oso](https://www.osohq.com/) -
[go-oso](https://github.com/osohq/go-oso). The Oso framework is implemented in
Rust and the Go bindings rely on `cgo` to interface with the main framework.

## Problem

First of all, we establish that the Go test runs fine using `go test`:

```bash
$ go test 
PASS
ok      github.com/shitz/oso-poc        0.006s
```

When trying to run test Bazel test target, however, it fails to build it:

```
$ bazel test //:lib_test 
INFO: Analyzed target //:lib_test (0 packages loaded, 1 target configured).
INFO: Found 1 test target...
ERROR: /home/sam/.cache/bazel/_bazel_sam/cd810a31bf197c4d2f729165aa95dad0/external/com_github_osohq_go_oso/internal/ffi/BUILD.bazel:3:11: GoCompilePkg external/com_github_osohq_go_oso/internal/ffi/ffi.a failed: (Exit 1): builder failed: error executing command bazel-out/k8-opt-exec-2B5CBBC6/bin/external/go_sdk/builder compilepkg -sdk external/go_sdk -installsuffix linux_amd64 -src external/com_github_osohq_go_oso/internal/ffi/ffi.go -embedroot '' -embedroot ... (remaining 35 arguments skipped)

Use --sandbox_debug to see verbose messages from the sandbox
com_github_shitz_oso_poc/external/com_github_osohq_go_oso/internal/ffi/ffi.go:6:11: fatal error: native/polar.h: No such file or directory
compilation terminated.
compilepkg: error running subcommand external/go_sdk/pkg/tool/linux_amd64/cgo: exit status 2
Target //:lib_test failed to build
Use --verbose_failures to see the command lines of failed build steps.
INFO: Elapsed time: 0.307s, Critical Path: 0.09s
INFO: 6 processes: 5 internal, 1 linux-sandbox.
FAILED: Build did NOT complete successfully
//:lib_test                                                     FAILED TO BUILD

FAILED: Build did NOT complete successfully
```

```
com_github_shitz_oso_poc/external/com_github_osohq_go_oso/internal/ffi/ffi.go:6:11: fatal error: native/polar.h: No such file or directory
```
