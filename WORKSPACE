workspace(name = "bazel_cache")

# Go
load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

http_archive(
    name = "io_bazel_rules_go",
    urls = ["https://github.com/bazelbuild/rules_go/releases/download/0.18.5/rules_go-0.18.5.tar.gz"],
    sha256 = "a82a352bffae6bee4e95f68a8d80a70e87f42c4741e6a448bec11998fcc82329",
)

http_archive(
    name = "bazel_gazelle",
    urls = ["https://github.com/bazelbuild/bazel-gazelle/releases/download/0.17.0/bazel-gazelle-0.17.0.tar.gz"],
    sha256 = "3c681998538231a2d24d0c07ed5a7658cb72bfb5fd4bf9911157c0e9ac6a2687",
)

load("@io_bazel_rules_go//go:deps.bzl", "go_rules_dependencies", "go_register_toolchains")

go_rules_dependencies()

go_register_toolchains()

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies", "go_repository")

gazelle_dependencies()

go_repository(
    name = "ag_pack_amqp",
    importpath = "pack.ag/amqp",
    tag = "v0.11.0",
)

go_repository(
    name = "co_honnef_go_tools",
    commit = "3f1c8253044a",
    importpath = "honnef.co/go/tools",
)

go_repository(
    name = "com_github_aws_aws_sdk_go",
    importpath = "github.com/aws/aws-sdk-go",
    tag = "v1.19.16",
)

go_repository(
    name = "com_github_azure_azure_amqp_common_go",
    importpath = "github.com/Azure/azure-amqp-common-go",
    tag = "v1.1.4",
)

go_repository(
    name = "com_github_azure_azure_pipeline_go",
    importpath = "github.com/Azure/azure-pipeline-go",
    tag = "v0.1.9",
)

go_repository(
    name = "com_github_azure_azure_sdk_for_go",
    importpath = "github.com/Azure/azure-sdk-for-go",
    tag = "v27.3.0",
)

go_repository(
    name = "com_github_azure_azure_service_bus_go",
    importpath = "github.com/Azure/azure-service-bus-go",
    tag = "v0.4.1",
)

go_repository(
    name = "com_github_azure_azure_storage_blob_go",
    importpath = "github.com/Azure/azure-storage-blob-go",
    tag = "v0.6.0",
)

go_repository(
    name = "com_github_azure_go_autorest",
    importpath = "github.com/Azure/go-autorest",
    tag = "v11.1.2",
)

go_repository(
    name = "com_github_azure_go_autorest_tracing",
    importpath = "github.com/Azure/go-autorest/tracing",
    tag = "v0.1.0",
)

go_repository(
    name = "com_github_burntsushi_toml",
    importpath = "github.com/BurntSushi/toml",
    tag = "v0.3.1",
)

go_repository(
    name = "com_github_census_instrumentation_opencensus_proto",
    importpath = "github.com/census-instrumentation/opencensus-proto",
    tag = "v0.2.0",
)

go_repository(
    name = "com_github_client9_misspell",
    importpath = "github.com/client9/misspell",
    tag = "v0.3.4",
)

go_repository(
    name = "com_github_codahale_hdrhistogram",
    commit = "3a0bb77429bd",
    importpath = "github.com/codahale/hdrhistogram",
)

go_repository(
    name = "com_github_davecgh_go_spew",
    importpath = "github.com/davecgh/go-spew",
    tag = "v1.1.1",
)

go_repository(
    name = "com_github_dgrijalva_jwt_go",
    importpath = "github.com/dgrijalva/jwt-go",
    tag = "v3.2.0",
)

go_repository(
    name = "com_github_dimchansky_utfbom",
    importpath = "github.com/dimchansky/utfbom",
    tag = "v1.1.0",
)

go_repository(
    name = "com_github_fortytw2_leaktest",
    importpath = "github.com/fortytw2/leaktest",
    tag = "v1.2.0",
)

go_repository(
    name = "com_github_fsnotify_fsnotify",
    importpath = "github.com/fsnotify/fsnotify",
    tag = "v1.4.7",
)

go_repository(
    name = "com_github_ghodss_yaml",
    importpath = "github.com/ghodss/yaml",
    tag = "v1.0.0",
)

go_repository(
    name = "com_github_go_sql_driver_mysql",
    importpath = "github.com/go-sql-driver/mysql",
    tag = "v1.4.1",
)

go_repository(
    name = "com_github_go_stack_stack",
    importpath = "github.com/go-stack/stack",
    tag = "v1.8.0",
)

go_repository(
    name = "com_github_golang_glog",
    commit = "23def4e6c14b",
    importpath = "github.com/golang/glog",
)

go_repository(
    name = "com_github_golang_groupcache",
    commit = "5b532d6fd5ef",
    importpath = "github.com/golang/groupcache",
)

go_repository(
    name = "com_github_golang_mock",
    importpath = "github.com/golang/mock",
    tag = "v1.2.0",
)

go_repository(
    name = "com_github_golang_protobuf",
    importpath = "github.com/golang/protobuf",
    tag = "v1.3.1",
)

go_repository(
    name = "com_github_golang_snappy",
    importpath = "github.com/golang/snappy",
    tag = "v0.0.1",
)

go_repository(
    name = "com_github_google_btree",
    commit = "4030bb1f1f0c",
    importpath = "github.com/google/btree",
)

go_repository(
    name = "com_github_google_go_cmp",
    importpath = "github.com/google/go-cmp",
    tag = "v0.3.0",
)

go_repository(
    name = "com_github_google_martian",
    importpath = "github.com/google/martian",
    tag = "v2.1.1-0.20190517191504-25dcb96d9e51",
)

go_repository(
    name = "com_github_google_pprof",
    commit = "3ea8567a2e57",
    importpath = "github.com/google/pprof",
)

go_repository(
    name = "com_github_google_uuid",
    importpath = "github.com/google/uuid",
    tag = "v1.1.1",
)

go_repository(
    name = "com_github_google_wire",
    importpath = "github.com/google/wire",
    tag = "v0.2.2",
)

go_repository(
    name = "com_github_googleapis_gax_go",
    importpath = "github.com/googleapis/gax-go",
    tag = "v2.0.2",
)

go_repository(
    name = "com_github_googleapis_gax_go_v2",
    importpath = "github.com/googleapis/gax-go/v2",
    tag = "v2.0.4",
)

go_repository(
    name = "com_github_googlecloudplatform_cloudsql_proxy",
    commit = "6ac0b49e7197",
    importpath = "github.com/GoogleCloudPlatform/cloudsql-proxy",
)

go_repository(
    name = "com_github_grpc_ecosystem_grpc_gateway",
    importpath = "github.com/grpc-ecosystem/grpc-gateway",
    tag = "v1.8.5",
)

go_repository(
    name = "com_github_hashicorp_golang_lru",
    importpath = "github.com/hashicorp/golang-lru",
    tag = "v0.5.1",
)

go_repository(
    name = "com_github_jmespath_go_jmespath",
    commit = "c2b33e8439af",
    importpath = "github.com/jmespath/go-jmespath",
)

go_repository(
    name = "com_github_joho_godotenv",
    importpath = "github.com/joho/godotenv",
    tag = "v1.3.0",
)

go_repository(
    name = "com_github_jstemmer_go_junit_report",
    commit = "af01ea7f8024",
    importpath = "github.com/jstemmer/go-junit-report",
)

go_repository(
    name = "com_github_kr_pretty",
    importpath = "github.com/kr/pretty",
    tag = "v0.1.0",
)

go_repository(
    name = "com_github_kr_pty",
    importpath = "github.com/kr/pty",
    tag = "v1.1.1",
)

go_repository(
    name = "com_github_kr_text",
    importpath = "github.com/kr/text",
    tag = "v0.1.0",
)

go_repository(
    name = "com_github_lib_pq",
    importpath = "github.com/lib/pq",
    tag = "v1.1.0",
)

go_repository(
    name = "com_github_mitchellh_go_homedir",
    importpath = "github.com/mitchellh/go-homedir",
    tag = "v1.1.0",
)

go_repository(
    name = "com_github_mitchellh_mapstructure",
    importpath = "github.com/mitchellh/mapstructure",
    tag = "v1.1.2",
)

go_repository(
    name = "com_github_opentracing_opentracing_go",
    importpath = "github.com/opentracing/opentracing-go",
    tag = "v1.0.2",
)

go_repository(
    name = "com_github_pkg_errors",
    importpath = "github.com/pkg/errors",
    tag = "v0.8.1",
)

go_repository(
    name = "com_github_pmezard_go_difflib",
    importpath = "github.com/pmezard/go-difflib",
    tag = "v1.0.0",
)

go_repository(
    name = "com_github_rogpeppe_fastuuid",
    commit = "6724a57986af",
    importpath = "github.com/rogpeppe/fastuuid",
)

go_repository(
    name = "com_github_stretchr_objx",
    importpath = "github.com/stretchr/objx",
    tag = "v0.1.1",
)

go_repository(
    name = "com_github_stretchr_testify",
    importpath = "github.com/stretchr/testify",
    tag = "v1.3.0",
)

go_repository(
    name = "com_github_tidwall_pretty",
    commit = "1166b9ac2b65",
    importpath = "github.com/tidwall/pretty",
)

go_repository(
    name = "com_github_uber_go_atomic",
    importpath = "github.com/uber-go/atomic",
    tag = "v1.3.2",
)

go_repository(
    name = "com_github_uber_jaeger_client_go",
    importpath = "github.com/uber/jaeger-client-go",
    tag = "v2.15.0",
)

go_repository(
    name = "com_github_uber_jaeger_lib",
    importpath = "github.com/uber/jaeger-lib",
    tag = "v1.5.0",
)

go_repository(
    name = "com_github_xdg_scram",
    commit = "7eeb5667e42c",
    importpath = "github.com/xdg/scram",
)

go_repository(
    name = "com_github_xdg_stringprep",
    importpath = "github.com/xdg/stringprep",
    tag = "v1.0.0",
)

go_repository(
    name = "com_google_cloud_go",
    importpath = "cloud.google.com/go",
    tag = "v0.39.0",
)

go_repository(
    name = "dev_gocloud",
    importpath = "gocloud.dev",
    tag = "v0.15.0",
)

go_repository(
    name = "in_gopkg_check_v1",
    commit = "788fd7840127",
    importpath = "gopkg.in/check.v1",
)

go_repository(
    name = "in_gopkg_resty_v1",
    importpath = "gopkg.in/resty.v1",
    tag = "v1.12.0",
)

go_repository(
    name = "in_gopkg_yaml_v2",
    importpath = "gopkg.in/yaml.v2",
    tag = "v2.2.1",
)

go_repository(
    name = "io_opencensus_go",
    importpath = "go.opencensus.io",
    tag = "v0.21.0",
)

go_repository(
    name = "io_opencensus_go_contrib_exporter_aws",
    commit = "2befc13012d0",
    importpath = "contrib.go.opencensus.io/exporter/aws",
)

go_repository(
    name = "io_opencensus_go_contrib_exporter_ocagent",
    importpath = "contrib.go.opencensus.io/exporter/ocagent",
    tag = "v0.4.12",
)

go_repository(
    name = "io_opencensus_go_contrib_exporter_stackdriver",
    importpath = "contrib.go.opencensus.io/exporter/stackdriver",
    tag = "v0.11.0",
)

go_repository(
    name = "io_opencensus_go_contrib_integrations_ocsql",
    importpath = "contrib.go.opencensus.io/integrations/ocsql",
    tag = "v0.1.4",
)

go_repository(
    name = "io_opencensus_go_contrib_resource",
    commit = "21591786a5e0",
    importpath = "contrib.go.opencensus.io/resource",
)

go_repository(
    name = "org_golang_google_api",
    importpath = "google.golang.org/api",
    tag = "v0.5.0",
)

go_repository(
    name = "org_golang_google_appengine",
    importpath = "google.golang.org/appengine",
    tag = "v1.4.0",
)

go_repository(
    name = "org_golang_google_genproto",
    commit = "b515fa19cec8",
    importpath = "google.golang.org/genproto",
)

go_repository(
    name = "org_golang_google_grpc",
    importpath = "google.golang.org/grpc",
    tag = "v1.20.1",
)

go_repository(
    name = "org_golang_x_crypto",
    commit = "d864b10871cd",
    importpath = "golang.org/x/crypto",
)

go_repository(
    name = "org_golang_x_exp",
    commit = "509febef88a4",
    importpath = "golang.org/x/exp",
)

go_repository(
    name = "org_golang_x_lint",
    commit = "d0100b6bd8b3",
    importpath = "golang.org/x/lint",
)

go_repository(
    name = "org_golang_x_net",
    commit = "4829fb13d2c6",
    importpath = "golang.org/x/net",
)

go_repository(
    name = "org_golang_x_oauth2",
    commit = "9f3314589c9a",
    importpath = "golang.org/x/oauth2",
)

go_repository(
    name = "org_golang_x_sync",
    commit = "112230192c58",
    importpath = "golang.org/x/sync",
)

go_repository(
    name = "org_golang_x_sys",
    commit = "953cdadca894",
    importpath = "golang.org/x/sys",
)

go_repository(
    name = "org_golang_x_text",
    commit = "e6919f6577db",
    importpath = "golang.org/x/text",
)

go_repository(
    name = "org_golang_x_time",
    commit = "85acf8d2951c",
    importpath = "golang.org/x/time",
)

go_repository(
    name = "org_golang_x_tools",
    commit = "fe54fb35175b",
    importpath = "golang.org/x/tools",
)

go_repository(
    name = "org_golang_x_xerrors",
    commit = "1f06c39b4373",
    importpath = "golang.org/x/xerrors",
)

go_repository(
    name = "org_mongodb_go_mongo_driver",
    importpath = "go.mongodb.org/mongo-driver",
    tag = "v1.0.1",
)

go_repository(
    name = "org_uber_go_atomic",
    importpath = "go.uber.org/atomic",
    tag = "v1.3.2",
)
