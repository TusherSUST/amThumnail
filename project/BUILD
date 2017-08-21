load("@io_bazel_rules_go//go:def.bzl", "gazelle", "go_binary", "go_library")

gazelle(name = "gazelle")

go_binary(
    name = "amThumbnailGenerator",
    srcs = [
        "handlers.go",
        "helpers.go",
        "main.go",
        "router.go",
        "routes.go"
    ],
    deps = [
      "@com_github_muesli_smartcrop",
      "@com_github_gorilla_mux"
    ]
)
