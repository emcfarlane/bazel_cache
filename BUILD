load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")
load("@bazel_gazelle//:def.bzl", "gazelle")

# gazelle:prefix github.com/emcfarlane/bazel_cache
gazelle(name = "gazelle")

go_library(
    name = "go_default_library",
    srcs = [
        "main.go",
        "server.go",
    ],
    importpath = "github.com/emcfarlane/bazel_cache",
    visibility = ["//visibility:private"],
    deps = [
        "@dev_gocloud//blob:go_default_library",
        "@dev_gocloud//blob/fileblob:go_default_library",
        "@dev_gocloud//blob/gcsblob:go_default_library",
        "@dev_gocloud//blob/memblob:go_default_library",
        "@dev_gocloud//blob/s3blob:go_default_library",
        "@dev_gocloud//gcerrors:go_default_library",
    ],
)

go_binary(
    name = "bazel_cache",
    embed = [":go_default_library"],
    visibility = ["//visibility:public"],
)
