load("@io_bazel_rules_go//proto:def.bzl", "go_proto_library")

proto_library(
    name = "firmware_keys_proto",
    srcs = [
        "firmware_keys.proto",
        "hid.proto",
    ],
    visibility = ["//visibility:public"],
)

go_proto_library(
    name = "firmware_keys",
    proto = ":firmware_keys_proto",
    visibility = ["//visibility:public"],
		importpath = "firmware/proto",
)
