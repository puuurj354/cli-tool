package templates

import (
	"embed"
	"fmt"
)

//go:embed embedded/*.tmpl
var embeddedFS embed.FS

// loadEmbedded loads a template from embedded files
func loadEmbedded(name string) string {
	data, err := embeddedFS.ReadFile("embedded/" + name)
	if err != nil {
		panic(fmt.Sprintf("failed to load embedded template %s: %v", name, err))
	}
	return string(data)
}

// Template represents a project template
type Template struct {
	Name        string
	Description string
	Directories []string
	Files       []FileTemplate
}

// FileTemplate represents a file to be generated
type FileTemplate struct {
	Path     string
	Template string
	Content  string
}

// Template content loaded from embedded files
var (
	authHandlerTmpl                    = loadEmbedded("auth_handler.tmpl")
	authJWTTmpl                        = loadEmbedded("auth_jwt.tmpl")
	authMainTmpl                       = loadEmbedded("auth_main.tmpl")
	authMiddlewareTmpl                 = loadEmbedded("auth_middleware.tmpl")
	authUserModelTmpl                  = loadEmbedded("auth_user_model.tmpl")
	challenge30DaysReadmeTmpl          = loadEmbedded("challenge30_days_readme.tmpl")
	challengeDay01TestTmpl             = loadEmbedded("challenge_day01_test.tmpl")
	challengeDay01Tmpl                 = loadEmbedded("challenge_day01.tmpl")
	challengeDay02TestTmpl             = loadEmbedded("challenge_day02_test.tmpl")
	challengeDay02Tmpl                 = loadEmbedded("challenge_day02.tmpl")
	challengeDay08TestTmpl             = loadEmbedded("challenge_day08_test.tmpl")
	challengeDay08Tmpl                 = loadEmbedded("challenge_day08.tmpl")
	challengeDay15TestTmpl             = loadEmbedded("challenge_day15_test.tmpl")
	challengeDay15Tmpl                 = loadEmbedded("challenge_day15.tmpl")
	challengeDay22TestTmpl             = loadEmbedded("challenge_day22_test.tmpl")
	challengeDay22Tmpl                 = loadEmbedded("challenge_day22.tmpl")
	cleanArchEntityTmpl                = loadEmbedded("clean_arch_entity.tmpl")
	cleanArchErrorsTmpl                = loadEmbedded("clean_arch_errors.tmpl")
	cleanArchHandlerTmpl               = loadEmbedded("clean_arch_handler.tmpl")
	cleanArchMainTmpl                  = loadEmbedded("clean_arch_main.tmpl")
	cleanArchRepoTmpl                  = loadEmbedded("clean_arch_repo.tmpl")
	cleanArchUsecaseTmpl               = loadEmbedded("clean_arch_usecase.tmpl")
	codeReview01BuggyTmpl              = loadEmbedded("code_review01_buggy.tmpl")
	codeReview01TestTmpl               = loadEmbedded("code_review01_test.tmpl")
	codeReview02BuggyTmpl              = loadEmbedded("code_review02_buggy.tmpl")
	codeReview02TestTmpl               = loadEmbedded("code_review02_test.tmpl")
	codeReview03BuggyTmpl              = loadEmbedded("code_review03_buggy.tmpl")
	codeReview03TestTmpl               = loadEmbedded("code_review03_test.tmpl")
	codeReviewReadmeTmpl               = loadEmbedded("code_review_readme.tmpl")
	cronJobsTmpl                       = loadEmbedded("cron_jobs.tmpl")
	cronMainTmpl                       = loadEmbedded("cron_main.tmpl")
	cronSchedulerTmpl                  = loadEmbedded("cron_scheduler.tmpl")
	dsaLinkedListTestTmpl              = loadEmbedded("dsa_linked_list_test.tmpl")
	dsaLinkedListTmpl                  = loadEmbedded("dsa_linked_list.tmpl")
	dsaQueueTestTmpl                   = loadEmbedded("dsa_queue_test.tmpl")
	dsaQueueTmpl                       = loadEmbedded("dsa_queue.tmpl")
	dsaRecursionTestTmpl               = loadEmbedded("dsa_recursion_test.tmpl")
	dsaRecursionTmpl                   = loadEmbedded("dsa_recursion.tmpl")
	dsaSearchingTestTmpl               = loadEmbedded("dsa_searching_test.tmpl")
	dsaSearchingTmpl                   = loadEmbedded("dsa_searching.tmpl")
	dsaSortingTestTmpl                 = loadEmbedded("dsa_sorting_test.tmpl")
	dsaSortingTmpl                     = loadEmbedded("dsa_sorting.tmpl")
	dsaStackTestTmpl                   = loadEmbedded("dsa_stack_test.tmpl")
	dsaStackTmpl                       = loadEmbedded("dsa_stack.tmpl")
	fullstackApiTsTmpl                 = loadEmbedded("fullstack_api_ts.tmpl")
	fullstackAppTsxTmpl                = loadEmbedded("fullstack_app_tsx.tmpl")
	fullstackBackendGoModTmpl          = loadEmbedded("fullstack_backend_go_mod.tmpl")
	fullstackBackendMainTmpl           = loadEmbedded("fullstack_backend_main.tmpl")
	fullstackCorsTmpl                  = loadEmbedded("fullstack_cors.tmpl")
	fullstackHandlerTmpl               = loadEmbedded("fullstack_handler.tmpl")
	fullstackHeaderTmpl                = loadEmbedded("fullstack_header.tmpl")
	fullstackIndexCssTmpl              = loadEmbedded("fullstack_index_css.tmpl")
	fullstackIndexHtmlTmpl             = loadEmbedded("fullstack_index_html.tmpl")
	fullstackMainTsxTmpl               = loadEmbedded("fullstack_main_tsx.tmpl")
	fullstackMakefileTmpl              = loadEmbedded("fullstack_makefile.tmpl")
	fullstackPackageJsonTmpl           = loadEmbedded("fullstack_package_json.tmpl")
	fullstackReadmeTmpl                = loadEmbedded("fullstack_readme.tmpl")
	fullstackViteConfigTmpl            = loadEmbedded("fullstack_vite_config.tmpl")
	fullstackGitignore                 = loadEmbedded("fullstack_gitignore.tmpl")
	fullstackTsconfigTmpl              = loadEmbedded("fullstack_tsconfig.tmpl")
	fullstackTsconfigNodeTmpl          = loadEmbedded("fullstack_tsconfig_node.tmpl")
	fullstackTailwindConfigTmpl        = loadEmbedded("fullstack_tailwind_config.tmpl")
	fullstackPostcssConfigTmpl         = loadEmbedded("fullstack_postcss_config.tmpl")
	goAPIMainTmpl                      = loadEmbedded("go_api_main.tmpl")
	goCLIMainTmpl                      = loadEmbedded("go_cli_main.tmpl")
	goCLIRootTmpl                      = loadEmbedded("go_cli_root.tmpl")
	goConfigTmpl                       = loadEmbedded("go_config.tmpl")
	goGRPCClientTmpl                   = loadEmbedded("go_grpc_client.tmpl")
	goGRPCMakefileTmpl                 = loadEmbedded("go_grpc_makefile.tmpl")
	goGRPCProtoTmpl                    = loadEmbedded("go_grpc_proto.tmpl")
	goGRPCServerTmpl                   = loadEmbedded("go_grpc_server.tmpl")
	goHandlerTmpl                      = loadEmbedded("go_handler.tmpl")
	goLibExampleTmpl                   = loadEmbedded("go_lib_example.tmpl")
	goLibMainTmpl                      = loadEmbedded("go_lib_main.tmpl")
	goLibTestTmpl                      = loadEmbedded("go_lib_test.tmpl")
	goModelTmpl                        = loadEmbedded("go_model.tmpl")
	goRepositoryTmpl                   = loadEmbedded("go_repository.tmpl")
	goServiceTmpl                      = loadEmbedded("go_service.tmpl")
	goTUIMainTmpl                      = loadEmbedded("go_tui_main.tmpl")
	goTUIModelTmpl                     = loadEmbedded("go_tui_model.tmpl")
	goWorkerJobTmpl                    = loadEmbedded("go_worker_job.tmpl")
	goWorkerMainTmpl                   = loadEmbedded("go_worker_main.tmpl")
	goWorkerQueueTmpl                  = loadEmbedded("go_worker_queue.tmpl")
	gitignoreGoTmpl                    = loadEmbedded("gitignore_go.tmpl")
	graphqlConfigTmpl                  = loadEmbedded("graphql_config.tmpl")
	graphqlMainTmpl                    = loadEmbedded("graphql_main.tmpl")
	graphqlResolverTmpl                = loadEmbedded("graphql_resolver.tmpl")
	graphqlSchemaTmpl                  = loadEmbedded("graphql_schema.tmpl")
	kafkaConsumerMainTmpl              = loadEmbedded("kafka_consumer_main.tmpl")
	kafkaConsumerTmpl                  = loadEmbedded("kafka_consumer.tmpl")
	kafkaDockerComposeTmpl             = loadEmbedded("kafka_docker_compose.tmpl")
	kafkaProducerMainTmpl              = loadEmbedded("kafka_producer_main.tmpl")
	kafkaProducerTmpl                  = loadEmbedded("kafka_producer.tmpl")
	lambdaHandlerTmpl                  = loadEmbedded("lambda_handler.tmpl")
	lambdaMainTmpl                     = loadEmbedded("lambda_main.tmpl")
	lambdaMakefileTmpl                 = loadEmbedded("lambda_makefile.tmpl")
	lambdaSAMTmpl                      = loadEmbedded("lambda_sam.tmpl")
	learnBenchFuncTmpl                 = loadEmbedded("learn_bench_func.tmpl")
	learnBenchTestTmpl                 = loadEmbedded("learn_bench_test.tmpl")
	learnChannelsTmpl                  = loadEmbedded("learn_channels.tmpl")
	learnConcurrencyReadmeTmpl         = loadEmbedded("learn_concurrency_readme.tmpl")
	learnContextCancellationTestTmpl   = loadEmbedded("learn_context_cancellation_test.tmpl")
	learnContextCancellationTmpl       = loadEmbedded("learn_context_cancellation.tmpl")
	learnContextReadmeTmpl             = loadEmbedded("learn_context_readme.tmpl")
	learnContextTimeoutTestTmpl        = loadEmbedded("learn_context_timeout_test.tmpl")
	learnContextTimeoutTmpl            = loadEmbedded("learn_context_timeout.tmpl")
	learnContextValuesTestTmpl         = loadEmbedded("learn_context_values_test.tmpl")
	learnContextValuesTmpl             = loadEmbedded("learn_context_values.tmpl")
	learnDSAReadmeTmpl                 = loadEmbedded("learn_dsa_readme.tmpl")
	learnErrorBasicsTestTmpl           = loadEmbedded("learn_error_basics_test.tmpl")
	learnErrorBasicsTmpl               = loadEmbedded("learn_error_basics.tmpl")
	learnErrorCustomTestTmpl           = loadEmbedded("learn_error_custom_test.tmpl")
	learnErrorCustomTmpl               = loadEmbedded("learn_error_custom.tmpl")
	learnErrorReadmeTmpl               = loadEmbedded("learn_error_readme.tmpl")
	learnErrorWrappingTestTmpl         = loadEmbedded("learn_error_wrapping_test.tmpl")
	learnErrorWrappingTmpl             = loadEmbedded("learn_error_wrapping.tmpl")
	learnGenericsBasicsTestTmpl        = loadEmbedded("learn_generics_basics_test.tmpl")
	learnGenericsBasicsTmpl            = loadEmbedded("learn_generics_basics.tmpl")
	learnGenericsConstraintsTestTmpl   = loadEmbedded("learn_generics_constraints_test.tmpl")
	learnGenericsConstraintsTmpl       = loadEmbedded("learn_generics_constraints.tmpl")
	learnGenericsPracticalTestTmpl     = loadEmbedded("learn_generics_practical_test.tmpl")
	learnGenericsPracticalTmpl         = loadEmbedded("learn_generics_practical.tmpl")
	learnGenericsReadmeTmpl            = loadEmbedded("learn_generics_readme.tmpl")
	learnGoroutinesTmpl                = loadEmbedded("learn_goroutines.tmpl")
	learnHTTPClientTestTmpl            = loadEmbedded("learn_http_client_test.tmpl")
	learnHTTPClientTmpl                = loadEmbedded("learn_http_client.tmpl")
	learnHTTPMiddlewareTestTmpl        = loadEmbedded("learn_http_middleware_test.tmpl")
	learnHTTPMiddlewareTmpl            = loadEmbedded("learn_http_middleware.tmpl")
	learnHTTPReadmeTmpl                = loadEmbedded("learn_http_readme.tmpl")
	learnHTTPServerTestTmpl            = loadEmbedded("learn_http_server_test.tmpl")
	learnHTTPServerTmpl                = loadEmbedded("learn_http_server.tmpl")
	learnInterfacesBasicsTestTmpl      = loadEmbedded("learn_interfaces_basics_test.tmpl")
	learnInterfacesBasicsTmpl          = loadEmbedded("learn_interfaces_basics.tmpl")
	learnInterfacesCompositionTestTmpl = loadEmbedded("learn_interfaces_composition_test.tmpl")
	learnInterfacesCompositionTmpl     = loadEmbedded("learn_interfaces_composition.tmpl")
	learnInterfacesPatternsTestTmpl    = loadEmbedded("learn_interfaces_patterns_test.tmpl")
	learnInterfacesPatternsTmpl        = loadEmbedded("learn_interfaces_patterns.tmpl")
	learnInterfacesReadmeTmpl          = loadEmbedded("learn_interfaces_readme.tmpl")
	learnPatternsFactoryTestTmpl       = loadEmbedded("learn_patterns_factory_test.tmpl")
	learnPatternsFactoryTmpl           = loadEmbedded("learn_patterns_factory.tmpl")
	learnPatternsObserverTestTmpl      = loadEmbedded("learn_patterns_observer_test.tmpl")
	learnPatternsObserverTmpl          = loadEmbedded("learn_patterns_observer.tmpl")
	learnPatternsReadmeTmpl            = loadEmbedded("learn_patterns_readme.tmpl")
	learnPatternsSingletonTestTmpl     = loadEmbedded("learn_patterns_singleton_test.tmpl")
	learnPatternsSingletonTmpl         = loadEmbedded("learn_patterns_singleton.tmpl")
	learnPatternsStrategyTestTmpl      = loadEmbedded("learn_patterns_strategy_test.tmpl")
	learnPatternsStrategyTmpl          = loadEmbedded("learn_patterns_strategy.tmpl")
	learnPatternsTmpl                  = loadEmbedded("learn_patterns.tmpl")
	learnSelectTmpl                    = loadEmbedded("learn_select.tmpl")
	learnSyncTmpl                      = loadEmbedded("learn_sync.tmpl")
	learnTableTestTmpl                 = loadEmbedded("learn_table_test.tmpl")
	learnTableValidatorTmpl            = loadEmbedded("learn_table_validator.tmpl")
	learnTestingReadmeTmpl             = loadEmbedded("learn_testing_readme.tmpl")
	learnUnitCalcTmpl                  = loadEmbedded("learn_unit_calc.tmpl")
	learnUnitTestTmpl                  = loadEmbedded("learn_unit_test.tmpl")
	microserviceDockerfileTmpl         = loadEmbedded("microservice_dockerfile.tmpl")
	microserviceHandlerTmpl            = loadEmbedded("microservice_handler.tmpl")
	microserviceHealthTmpl             = loadEmbedded("microservice_health.tmpl")
	microserviceLoggingTmpl            = loadEmbedded("microservice_logging.tmpl")
	microserviceMainTmpl               = loadEmbedded("microservice_main.tmpl")
	microserviceMakefileTmpl           = loadEmbedded("microservice_makefile.tmpl")
	miniProjectReadmeTmpl              = loadEmbedded("mini_project_readme.tmpl")
	monorepoAPITmpl                    = loadEmbedded("monorepo_api.tmpl")
	monorepoConfigTmpl                 = loadEmbedded("monorepo_config.tmpl")
	monorepoLoggerTmpl                 = loadEmbedded("monorepo_logger.tmpl")
	monorepoMakefileTmpl               = loadEmbedded("monorepo_makefile.tmpl")
	monorepoWorkerTmpl                 = loadEmbedded("monorepo_worker.tmpl")
	redisCacheTmpl                     = loadEmbedded("redis_cache.tmpl")
	redisDockerComposeTmpl             = loadEmbedded("redis_docker_compose.tmpl")
	redisMainTmpl                      = loadEmbedded("redis_main.tmpl")
	redisPubSubTmpl                    = loadEmbedded("redis_pub_sub.tmpl")
	readmeTmpl                         = loadEmbedded("readme.tmpl")
	refactoring01BeforeTmpl            = loadEmbedded("refactoring01_before.tmpl")
	refactoring01HintsTmpl             = loadEmbedded("refactoring01_hints.tmpl")
	refactoring02BeforeTmpl            = loadEmbedded("refactoring02_before.tmpl")
	refactoring02HintsTmpl             = loadEmbedded("refactoring02_hints.tmpl")
	refactoring03BeforeTmpl            = loadEmbedded("refactoring03_before.tmpl")
	refactoring03HintsTmpl             = loadEmbedded("refactoring03_hints.tmpl")
	refactoringReadmeTmpl              = loadEmbedded("refactoring_readme.tmpl")
	todoCliMainTmpl                    = loadEmbedded("todo_cli_main.tmpl")
	todoCliReadmeTmpl                  = loadEmbedded("todo_cli_readme.tmpl")
	todoCliTestTmpl                    = loadEmbedded("todo_cli_test.tmpl")
	urlShortenerMainTmpl               = loadEmbedded("url_shortener_main.tmpl")
	urlShortenerReadmeTmpl             = loadEmbedded("url_shortener_readme.tmpl")
	urlShortenerTestTmpl               = loadEmbedded("url_shortener_test.tmpl")
	websocketClientTmpl                = loadEmbedded("websocket_client.tmpl")
	websocketHTMLTmpl                  = loadEmbedded("websocket_html.tmpl")
	websocketHubTmpl                   = loadEmbedded("websocket_hub.tmpl")
	websocketMainTmpl                  = loadEmbedded("websocket_main.tmpl")
)

// Generator functions
func genReadme() string             { return loadEmbedded("readme.tmpl") }
func genGitignoreGo() string        { return loadEmbedded("gitignore_go.tmpl") }
func genFullstackGitignore() string { return loadEmbedded("fullstack_gitignore.tmpl") }
func genTsconfig() string           { return loadEmbedded("fullstack_tsconfig.tmpl") }
func genTsconfigNode() string       { return loadEmbedded("fullstack_tsconfig_node.tmpl") }
func genTailwindConfig() string     { return loadEmbedded("fullstack_tailwind_config.tmpl") }
func genPostcssConfig() string      { return loadEmbedded("fullstack_postcss_config.tmpl") }

func init() {
	initBuiltInTemplates()
}

var builtInTemplates map[string]Template

// initBuiltInTemplates initializes the built-in templates map
// This must be called in init() after all template strings are populated
func initBuiltInTemplates() {
	builtInTemplates = map[string]Template{
		"go-api": {
			Name:        "go-api",
			Description: "Go REST API with clean architecture",
			Directories: []string{
				"cmd/api",
				"internal/handler",
				"internal/service",
				"internal/repository",
				"internal/model",
				"pkg/config",
			},
			Files: []FileTemplate{
				{Path: "cmd/api/main.go", Content: goAPIMainTmpl},
				{Path: "internal/handler/handler.go", Content: goHandlerTmpl},
				{Path: "internal/service/service.go", Content: goServiceTmpl},
				{Path: "internal/repository/repository.go", Content: goRepositoryTmpl},
				{Path: "internal/model/model.go", Content: goModelTmpl},
				{Path: "pkg/config/config.go", Content: goConfigTmpl},
				{Path: "README.md", Content: readmeTmpl},
				{Path: ".gitignore", Content: gitignoreGoTmpl},
			},
		},
		"go-cli": {
			Name:        "go-cli",
			Description: "Go CLI application with Cobra",
			Directories: []string{
				"cmd",
				"internal",
				"pkg",
			},
			Files: []FileTemplate{
				{Path: "main.go", Content: goCLIMainTmpl},
				{Path: "cmd/root.go", Content: goCLIRootTmpl},
				{Path: "README.md", Content: readmeTmpl},
				{Path: ".gitignore", Content: gitignoreGoTmpl},
			},
		},
		"go-lib": {
			Name:        "go-lib",
			Description: "Go library/package",
			Directories: []string{
				"internal",
				"examples",
			},
			Files: []FileTemplate{
				{Path: "{{.ProjectName}}.go", Content: goLibMainTmpl},
				{Path: "{{.ProjectName}}_test.go", Content: goLibTestTmpl},
				{Path: "examples/main.go", Content: goLibExampleTmpl},
				{Path: "README.md", Content: readmeTmpl},
				{Path: ".gitignore", Content: gitignoreGoTmpl},
			},
		},
		// Learning templates
		"learn-concurrency": {
			Name:        "learn-concurrency",
			Description: "Learn Go concurrency patterns",
			Directories: []string{
				"01-goroutines",
				"02-channels",
				"03-select",
				"04-sync",
				"05-patterns",
			},
			Files: []FileTemplate{
				{Path: "README.md", Content: learnConcurrencyReadmeTmpl},
				{Path: "01-goroutines/main.go", Content: learnGoroutinesTmpl},
				{Path: "02-channels/main.go", Content: learnChannelsTmpl},
				{Path: "03-select/main.go", Content: learnSelectTmpl},
				{Path: "04-sync/main.go", Content: learnSyncTmpl},
				{Path: "05-patterns/main.go", Content: learnPatternsTmpl},
				{Path: ".gitignore", Content: gitignoreGoTmpl},
			},
		},
		"learn-testing": {
			Name:        "learn-testing",
			Description: "Learn Go testing techniques",
			Directories: []string{
				"unit",
				"table",
				"mock",
				"benchmark",
			},
			Files: []FileTemplate{
				{Path: "README.md", Content: learnTestingReadmeTmpl},
				{Path: "unit/calculator.go", Content: learnUnitCalcTmpl},
				{Path: "unit/calculator_test.go", Content: learnUnitTestTmpl},
				{Path: "table/validator.go", Content: learnTableValidatorTmpl},
				{Path: "table/validator_test.go", Content: learnTableTestTmpl},
				{Path: "benchmark/string_builder.go", Content: learnBenchFuncTmpl},
				{Path: "benchmark/string_builder_test.go", Content: learnBenchTestTmpl},
				{Path: ".gitignore", Content: gitignoreGoTmpl},
			},
		},
		"learn-dsa": {
			Name:        "learn-dsa",
			Description: "Practice Data Structures & Algorithms",
			Directories: []string{
				"datastructures/stack",
				"datastructures/queue",
				"datastructures/linkedlist",
				"datastructures/tree",
				"algorithms/sorting",
				"algorithms/searching",
				"algorithms/recursion",
			},
			Files: []FileTemplate{
				{Path: "README.md", Content: learnDSAReadmeTmpl},
				// Data Structures
				{Path: "datastructures/stack/stack.go", Content: dsaStackTmpl},
				{Path: "datastructures/stack/stack_test.go", Content: dsaStackTestTmpl},
				{Path: "datastructures/queue/queue.go", Content: dsaQueueTmpl},
				{Path: "datastructures/queue/queue_test.go", Content: dsaQueueTestTmpl},
				{Path: "datastructures/linkedlist/linkedlist.go", Content: dsaLinkedListTmpl},
				{Path: "datastructures/linkedlist/linkedlist_test.go", Content: dsaLinkedListTestTmpl},
				// Algorithms
				{Path: "algorithms/sorting/sorting.go", Content: dsaSortingTmpl},
				{Path: "algorithms/sorting/sorting_test.go", Content: dsaSortingTestTmpl},
				{Path: "algorithms/searching/searching.go", Content: dsaSearchingTmpl},
				{Path: "algorithms/searching/searching_test.go", Content: dsaSearchingTestTmpl},
				{Path: "algorithms/recursion/recursion.go", Content: dsaRecursionTmpl},
				{Path: "algorithms/recursion/recursion_test.go", Content: dsaRecursionTestTmpl},
				{Path: ".gitignore", Content: gitignoreGoTmpl},
			},
		},
		"go-grpc": {
			Name:        "go-grpc",
			Description: "Go gRPC service template",
			Directories: []string{
				"cmd/server",
				"cmd/client",
				"internal/service",
				"proto",
			},
			Files: []FileTemplate{
				{Path: "cmd/server/main.go", Content: goGRPCServerTmpl},
				{Path: "cmd/client/main.go", Content: goGRPCClientTmpl},
				{Path: "proto/service.proto", Content: goGRPCProtoTmpl},
				{Path: "Makefile", Content: goGRPCMakefileTmpl},
				{Path: "README.md", Content: readmeTmpl},
				{Path: ".gitignore", Content: gitignoreGoTmpl},
			},
		},
		"go-worker": {
			Name:        "go-worker",
			Description: "Background worker with job queue",
			Directories: []string{
				"cmd/worker",
				"internal/job",
				"internal/queue",
			},
			Files: []FileTemplate{
				{Path: "cmd/worker/main.go", Content: goWorkerMainTmpl},
				{Path: "internal/job/job.go", Content: goWorkerJobTmpl},
				{Path: "internal/queue/queue.go", Content: goWorkerQueueTmpl},
				{Path: "README.md", Content: readmeTmpl},
				{Path: ".gitignore", Content: gitignoreGoTmpl},
			},
		},
		"go-tui": {
			Name:        "go-tui",
			Description: "Terminal UI app with Bubbletea",
			Directories: []string{
				"internal/ui",
			},
			Files: []FileTemplate{
				{Path: "main.go", Content: goTUIMainTmpl},
				{Path: "internal/ui/model.go", Content: goTUIModelTmpl},
				{Path: "README.md", Content: readmeTmpl},
				{Path: ".gitignore", Content: gitignoreGoTmpl},
			},
		},
		"fullstack": {
			Name:        "fullstack",
			Description: "Go backend + React/Vite/Bun/Tailwind frontend",
			Directories: []string{
				"backend/cmd/api",
				"backend/internal/handler",
				"backend/internal/middleware",
				"frontend/src/components",
				"frontend/src/hooks",
				"frontend/src/lib",
			},
			Files: []FileTemplate{
				// Backend files
				{Path: "backend/cmd/api/main.go", Content: fullstackBackendMainTmpl},
				{Path: "backend/internal/handler/handler.go", Content: fullstackHandlerTmpl},
				{Path: "backend/internal/middleware/cors.go", Content: fullstackCorsTmpl},
				{Path: "backend/go.mod", Content: fullstackBackendGoModTmpl},
				// Frontend files
				{Path: "frontend/package.json", Content: fullstackPackageJsonTmpl},
				{Path: "frontend/tsconfig.json", Content: fullstackTsconfigTmpl},
				{Path: "frontend/tsconfig.node.json", Content: fullstackTsconfigNodeTmpl},
				{Path: "frontend/vite.config.ts", Content: fullstackViteConfigTmpl},
				{Path: "frontend/tailwind.config.js", Content: fullstackTailwindConfigTmpl},
				{Path: "frontend/postcss.config.js", Content: fullstackPostcssConfigTmpl},
				{Path: "frontend/index.html", Content: fullstackIndexHtmlTmpl},
				{Path: "frontend/src/main.tsx", Content: fullstackMainTsxTmpl},
				{Path: "frontend/src/App.tsx", Content: fullstackAppTsxTmpl},
				{Path: "frontend/src/index.css", Content: fullstackIndexCssTmpl},
				{Path: "frontend/src/lib/api.ts", Content: fullstackApiTsTmpl},
				{Path: "frontend/src/components/Header.tsx", Content: fullstackHeaderTmpl},
				// Root files
				{Path: "README.md", Content: fullstackReadmeTmpl},
				{Path: "Makefile", Content: fullstackMakefileTmpl},
				{Path: ".gitignore", Content: fullstackGitignore},
			},
		},
		// Skill Coding Templates
		"challenge-30days": {
			Name:        "challenge-30days",
			Description: "30-day Go coding challenge",
			Directories: []string{
				"week1/day01_hello",
				"week1/day02_variables",
				"week1/day03_conditionals",
				"week2/day08_recursion",
				"week2/day09_slices",
				"week3/day15_http",
				"week3/day16_json",
				"week4/day22_concurrency",
				"week4/day23_channels",
			},
			Files: []FileTemplate{
				{Path: "README.md", Content: challenge30DaysReadmeTmpl},
				{Path: "week1/day01_hello/main.go", Content: challengeDay01Tmpl},
				{Path: "week1/day01_hello/main_test.go", Content: challengeDay01TestTmpl},
				{Path: "week1/day02_variables/main.go", Content: challengeDay02Tmpl},
				{Path: "week1/day02_variables/main_test.go", Content: challengeDay02TestTmpl},
				{Path: "week2/day08_recursion/main.go", Content: challengeDay08Tmpl},
				{Path: "week2/day08_recursion/main_test.go", Content: challengeDay08TestTmpl},
				{Path: "week3/day15_http/main.go", Content: challengeDay15Tmpl},
				{Path: "week3/day15_http/main_test.go", Content: challengeDay15TestTmpl},
				{Path: "week4/day22_concurrency/main.go", Content: challengeDay22Tmpl},
				{Path: "week4/day22_concurrency/main_test.go", Content: challengeDay22TestTmpl},
				{Path: ".gitignore", Content: gitignoreGoTmpl},
			},
		},
		"mini-project": {
			Name:        "mini-project",
			Description: "Mini projects to build (todo-cli, url-shortener)",
			Directories: []string{
				"todo-cli",
				"url-shortener",
			},
			Files: []FileTemplate{
				{Path: "README.md", Content: miniProjectReadmeTmpl},
				{Path: "todo-cli/README.md", Content: todoCliReadmeTmpl},
				{Path: "todo-cli/main.go", Content: todoCliMainTmpl},
				{Path: "todo-cli/main_test.go", Content: todoCliTestTmpl},
				{Path: "url-shortener/README.md", Content: urlShortenerReadmeTmpl},
				{Path: "url-shortener/main.go", Content: urlShortenerMainTmpl},
				{Path: "url-shortener/main_test.go", Content: urlShortenerTestTmpl},
				{Path: ".gitignore", Content: gitignoreGoTmpl},
			},
		},
		"refactoring-exercise": {
			Name:        "refactoring-exercise",
			Description: "Practice refactoring bad code",
			Directories: []string{
				"exercises/01_long_function",
				"exercises/02_magic_numbers",
				"exercises/03_poor_naming",
			},
			Files: []FileTemplate{
				{Path: "README.md", Content: refactoringReadmeTmpl},
				{Path: "exercises/01_long_function/before.go", Content: refactoring01BeforeTmpl},
				{Path: "exercises/01_long_function/hints.md", Content: refactoring01HintsTmpl},
				{Path: "exercises/02_magic_numbers/before.go", Content: refactoring02BeforeTmpl},
				{Path: "exercises/02_magic_numbers/hints.md", Content: refactoring02HintsTmpl},
				{Path: "exercises/03_poor_naming/before.go", Content: refactoring03BeforeTmpl},
				{Path: "exercises/03_poor_naming/hints.md", Content: refactoring03HintsTmpl},
				{Path: ".gitignore", Content: gitignoreGoTmpl},
			},
		},
		"code-review-exercise": {
			Name:        "code-review-exercise",
			Description: "Find bugs in code (code review practice)",
			Directories: []string{
				"bugs/01_off_by_one",
				"bugs/02_nil_pointer",
				"bugs/03_race_condition",
			},
			Files: []FileTemplate{
				{Path: "README.md", Content: codeReviewReadmeTmpl},
				{Path: "bugs/01_off_by_one/buggy.go", Content: codeReview01BuggyTmpl},
				{Path: "bugs/01_off_by_one/buggy_test.go", Content: codeReview01TestTmpl},
				{Path: "bugs/02_nil_pointer/buggy.go", Content: codeReview02BuggyTmpl},
				{Path: "bugs/02_nil_pointer/buggy_test.go", Content: codeReview02TestTmpl},
				{Path: "bugs/03_race_condition/buggy.go", Content: codeReview03BuggyTmpl},
				{Path: "bugs/03_race_condition/buggy_test.go", Content: codeReview03TestTmpl},
				{Path: ".gitignore", Content: gitignoreGoTmpl},
			},
		},
		// Learning Specific Skills Templates
		"learn-generics": {
			Name:        "learn-generics",
			Description: "Learn Go generics (type parameters & constraints)",
			Directories: []string{
				"basics",
				"constraints",
				"practical",
			},
			Files: []FileTemplate{
				{Path: "README.md", Content: learnGenericsReadmeTmpl},
				{Path: "basics/main.go", Content: learnGenericsBasicsTmpl},
				{Path: "basics/main_test.go", Content: learnGenericsBasicsTestTmpl},
				{Path: "constraints/main.go", Content: learnGenericsConstraintsTmpl},
				{Path: "constraints/main_test.go", Content: learnGenericsConstraintsTestTmpl},
				{Path: "practical/main.go", Content: learnGenericsPracticalTmpl},
				{Path: "practical/main_test.go", Content: learnGenericsPracticalTestTmpl},
				{Path: ".gitignore", Content: gitignoreGoTmpl},
			},
		},
		"learn-context": {
			Name:        "learn-context",
			Description: "Learn context.Context (cancellation & timeout)",
			Directories: []string{
				"cancellation",
				"timeout",
				"values",
			},
			Files: []FileTemplate{
				{Path: "README.md", Content: learnContextReadmeTmpl},
				{Path: "cancellation/main.go", Content: learnContextCancellationTmpl},
				{Path: "cancellation/main_test.go", Content: learnContextCancellationTestTmpl},
				{Path: "timeout/main.go", Content: learnContextTimeoutTmpl},
				{Path: "timeout/main_test.go", Content: learnContextTimeoutTestTmpl},
				{Path: "values/main.go", Content: learnContextValuesTmpl},
				{Path: "values/main_test.go", Content: learnContextValuesTestTmpl},
				{Path: ".gitignore", Content: gitignoreGoTmpl},
			},
		},
		"learn-http": {
			Name:        "learn-http",
			Description: "Learn HTTP client, server & middleware",
			Directories: []string{
				"client",
				"server",
				"middleware",
			},
			Files: []FileTemplate{
				{Path: "README.md", Content: learnHTTPReadmeTmpl},
				{Path: "client/main.go", Content: learnHTTPClientTmpl},
				{Path: "client/main_test.go", Content: learnHTTPClientTestTmpl},
				{Path: "server/main.go", Content: learnHTTPServerTmpl},
				{Path: "server/main_test.go", Content: learnHTTPServerTestTmpl},
				{Path: "middleware/main.go", Content: learnHTTPMiddlewareTmpl},
				{Path: "middleware/main_test.go", Content: learnHTTPMiddlewareTestTmpl},
				{Path: ".gitignore", Content: gitignoreGoTmpl},
			},
		},
		"learn-error-handling": {
			Name:        "learn-error-handling",
			Description: "Learn error handling patterns in Go",
			Directories: []string{
				"basics",
				"wrapping",
				"custom",
			},
			Files: []FileTemplate{
				{Path: "README.md", Content: learnErrorReadmeTmpl},
				{Path: "basics/main.go", Content: learnErrorBasicsTmpl},
				{Path: "basics/main_test.go", Content: learnErrorBasicsTestTmpl},
				{Path: "wrapping/main.go", Content: learnErrorWrappingTmpl},
				{Path: "wrapping/main_test.go", Content: learnErrorWrappingTestTmpl},
				{Path: "custom/main.go", Content: learnErrorCustomTmpl},
				{Path: "custom/main_test.go", Content: learnErrorCustomTestTmpl},
				{Path: ".gitignore", Content: gitignoreGoTmpl},
			},
		},
		"learn-interfaces": {
			Name:        "learn-interfaces",
			Description: "Learn interfaces & polymorphism in Go",
			Directories: []string{
				"basics",
				"composition",
				"patterns",
			},
			Files: []FileTemplate{
				{Path: "README.md", Content: learnInterfacesReadmeTmpl},
				{Path: "basics/main.go", Content: learnInterfacesBasicsTmpl},
				{Path: "basics/main_test.go", Content: learnInterfacesBasicsTestTmpl},
				{Path: "composition/main.go", Content: learnInterfacesCompositionTmpl},
				{Path: "composition/main_test.go", Content: learnInterfacesCompositionTestTmpl},
				{Path: "patterns/main.go", Content: learnInterfacesPatternsTmpl},
				{Path: "patterns/main_test.go", Content: learnInterfacesPatternsTestTmpl},
				{Path: ".gitignore", Content: gitignoreGoTmpl},
			},
		},
		"learn-design-patterns": {
			Name:        "learn-design-patterns",
			Description: "Learn common design patterns in Go",
			Directories: []string{
				"creational",
				"behavioral",
				"structural",
			},
			Files: []FileTemplate{
				{Path: "README.md", Content: learnPatternsReadmeTmpl},
				{Path: "creational/factory.go", Content: learnPatternsFactoryTmpl},
				{Path: "creational/factory_test.go", Content: learnPatternsFactoryTestTmpl},
				{Path: "creational/singleton.go", Content: learnPatternsSingletonTmpl},
				{Path: "creational/singleton_test.go", Content: learnPatternsSingletonTestTmpl},
				{Path: "behavioral/observer.go", Content: learnPatternsObserverTmpl},
				{Path: "behavioral/observer_test.go", Content: learnPatternsObserverTestTmpl},
				{Path: "behavioral/strategy.go", Content: learnPatternsStrategyTmpl},
				{Path: "behavioral/strategy_test.go", Content: learnPatternsStrategyTestTmpl},
				{Path: ".gitignore", Content: gitignoreGoTmpl},
			},
		},
		// New Project Templates
		"go-microservice": {
			Name:        "go-microservice",
			Description: "Microservice with health check, metrics & graceful shutdown",
			Directories: []string{
				"cmd/server",
				"internal/handler",
				"internal/middleware",
				"internal/health",
			},
			Files: []FileTemplate{
				{Path: "cmd/server/main.go", Content: microserviceMainTmpl},
				{Path: "internal/handler/handler.go", Content: microserviceHandlerTmpl},
				{Path: "internal/middleware/logging.go", Content: microserviceLoggingTmpl},
				{Path: "internal/health/health.go", Content: microserviceHealthTmpl},
				{Path: "Dockerfile", Content: microserviceDockerfileTmpl},
				{Path: "Makefile", Content: microserviceMakefileTmpl},
				{Path: "README.md", Content: readmeTmpl},
				{Path: ".gitignore", Content: gitignoreGoTmpl},
			},
		},
		"go-websocket": {
			Name:        "go-websocket",
			Description: "Real-time WebSocket application",
			Directories: []string{
				"cmd/server",
				"internal/hub",
				"internal/client",
				"web",
			},
			Files: []FileTemplate{
				{Path: "cmd/server/main.go", Content: websocketMainTmpl},
				{Path: "internal/hub/hub.go", Content: websocketHubTmpl},
				{Path: "internal/client/client.go", Content: websocketClientTmpl},
				{Path: "web/index.html", Content: websocketHTMLTmpl},
				{Path: "README.md", Content: readmeTmpl},
				{Path: ".gitignore", Content: gitignoreGoTmpl},
			},
		},
		"go-graphql": {
			Name:        "go-graphql",
			Description: "GraphQL API with gqlgen",
			Directories: []string{
				"cmd/server",
				"graph",
			},
			Files: []FileTemplate{
				{Path: "cmd/server/main.go", Content: graphqlMainTmpl},
				{Path: "graph/schema.graphqls", Content: graphqlSchemaTmpl},
				{Path: "graph/resolver.go", Content: graphqlResolverTmpl},
				{Path: "gqlgen.yml", Content: graphqlConfigTmpl},
				{Path: "README.md", Content: readmeTmpl},
				{Path: ".gitignore", Content: gitignoreGoTmpl},
			},
		},
		"go-lambda": {
			Name:        "go-lambda",
			Description: "AWS Lambda function with SAM",
			Directories: []string{
				"cmd/lambda",
				"internal/handler",
			},
			Files: []FileTemplate{
				{Path: "cmd/lambda/main.go", Content: lambdaMainTmpl},
				{Path: "internal/handler/handler.go", Content: lambdaHandlerTmpl},
				{Path: "template.yaml", Content: lambdaSAMTmpl},
				{Path: "Makefile", Content: lambdaMakefileTmpl},
				{Path: "README.md", Content: readmeTmpl},
				{Path: ".gitignore", Content: gitignoreGoTmpl},
			},
		},
		"go-cron": {
			Name:        "go-cron",
			Description: "Scheduled jobs with cron",
			Directories: []string{
				"cmd/scheduler",
				"internal/jobs",
				"internal/scheduler",
			},
			Files: []FileTemplate{
				{Path: "cmd/scheduler/main.go", Content: cronMainTmpl},
				{Path: "internal/jobs/jobs.go", Content: cronJobsTmpl},
				{Path: "internal/scheduler/scheduler.go", Content: cronSchedulerTmpl},
				{Path: "README.md", Content: readmeTmpl},
				{Path: ".gitignore", Content: gitignoreGoTmpl},
			},
		},
		// Additional Project Templates
		"go-auth": {
			Name:        "go-auth",
			Description: "JWT authentication with middleware",
			Directories: []string{
				"cmd/server",
				"internal/auth",
				"internal/handler",
				"internal/model",
			},
			Files: []FileTemplate{
				{Path: "cmd/server/main.go", Content: authMainTmpl},
				{Path: "internal/auth/jwt.go", Content: authJWTTmpl},
				{Path: "internal/auth/middleware.go", Content: authMiddlewareTmpl},
				{Path: "internal/handler/auth.go", Content: authHandlerTmpl},
				{Path: "internal/model/user.go", Content: authUserModelTmpl},
				{Path: "README.md", Content: readmeTmpl},
				{Path: ".gitignore", Content: gitignoreGoTmpl},
			},
		},
		"go-kafka": {
			Name:        "go-kafka",
			Description: "Kafka consumer & producer",
			Directories: []string{
				"cmd/producer",
				"cmd/consumer",
				"internal/kafka",
			},
			Files: []FileTemplate{
				{Path: "cmd/producer/main.go", Content: kafkaProducerMainTmpl},
				{Path: "cmd/consumer/main.go", Content: kafkaConsumerMainTmpl},
				{Path: "internal/kafka/producer.go", Content: kafkaProducerTmpl},
				{Path: "internal/kafka/consumer.go", Content: kafkaConsumerTmpl},
				{Path: "docker-compose.yml", Content: kafkaDockerComposeTmpl},
				{Path: "README.md", Content: readmeTmpl},
				{Path: ".gitignore", Content: gitignoreGoTmpl},
			},
		},
		"go-redis": {
			Name:        "go-redis",
			Description: "Redis caching & pub/sub patterns",
			Directories: []string{
				"cmd/server",
				"internal/cache",
				"internal/pubsub",
			},
			Files: []FileTemplate{
				{Path: "cmd/server/main.go", Content: redisMainTmpl},
				{Path: "internal/cache/redis.go", Content: redisCacheTmpl},
				{Path: "internal/pubsub/pubsub.go", Content: redisPubSubTmpl},
				{Path: "docker-compose.yml", Content: redisDockerComposeTmpl},
				{Path: "README.md", Content: readmeTmpl},
				{Path: ".gitignore", Content: gitignoreGoTmpl},
			},
		},
		"go-clean-arch": {
			Name:        "go-clean-arch",
			Description: "Clean Architecture pattern",
			Directories: []string{
				"cmd/api",
				"internal/entity",
				"internal/usecase",
				"internal/repository",
				"internal/delivery/http",
				"pkg/errors",
			},
			Files: []FileTemplate{
				{Path: "cmd/api/main.go", Content: cleanArchMainTmpl},
				{Path: "internal/entity/user.go", Content: cleanArchEntityTmpl},
				{Path: "internal/usecase/user.go", Content: cleanArchUsecaseTmpl},
				{Path: "internal/repository/user.go", Content: cleanArchRepoTmpl},
				{Path: "internal/delivery/http/handler.go", Content: cleanArchHandlerTmpl},
				{Path: "pkg/errors/errors.go", Content: cleanArchErrorsTmpl},
				{Path: "README.md", Content: readmeTmpl},
				{Path: ".gitignore", Content: gitignoreGoTmpl},
			},
		},
		"go-monorepo": {
			Name:        "go-monorepo",
			Description: "Multi-service monorepo with shared packages",
			Directories: []string{
				"services/api",
				"services/worker",
				"pkg/shared",
			},
			Files: []FileTemplate{
				{Path: "services/api/main.go", Content: monorepoAPITmpl},
				{Path: "services/worker/main.go", Content: monorepoWorkerTmpl},
				{Path: "pkg/shared/config.go", Content: monorepoConfigTmpl},
				{Path: "pkg/shared/logger.go", Content: monorepoLoggerTmpl},
				{Path: "Makefile", Content: monorepoMakefileTmpl},
				{Path: "README.md", Content: readmeTmpl},
				{Path: ".gitignore", Content: gitignoreGoTmpl},
			},
		},
	}
}

// GetTemplate returns a template by name
func GetTemplate(name string) (Template, error) {
	t, ok := builtInTemplates[name]
	if !ok {
		return Template{}, fmt.Errorf("template not found: %s", name)
	}
	return t, nil
}

// GetAllTemplates returns all available templates
func GetAllTemplates() []Template {
	result := make([]Template, 0, len(builtInTemplates))
	for _, t := range builtInTemplates {
		result = append(result, t)
	}
	return result
}
