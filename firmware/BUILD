load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_binary", "go_test")

go_binary(
    name = "myfalcon",
    srcs = ["main.go"],
    deps = [
				":string_hid",
				":verification",
        "@com_github_golang_protobuf//proto:go_default_library",
        "//firmware/proto:firmware_keys",
    ],
    importpath = "firmware",
)

go_library(
		name = "string_hid",
		srcs = ["string_hid.go"],
		deps = [
        "@com_github_golang_protobuf//proto:go_default_library",
        "//firmware/proto:firmware_keys",
    ],
    importpath = "string_hid",
)

go_library(
		name = "verification",
		srcs = ["verification.go"],
		deps = [
				":string_hid",
        "@com_github_golang_protobuf//proto:go_default_library",
        "//firmware/proto:firmware_keys",
    ],
    importpath = "verification",
)

go_test(
		name = "verification_test",
		srcs = ["verification_test.go"],
		deps = [
				":verification",
        "@com_github_golang_protobuf//proto:go_default_library",
        "//firmware/proto:firmware_keys",
		],
    importpath = "firmware",
)
