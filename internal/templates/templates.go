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
	authHandlerTmpl             = loadEmbedded("auth_handler.tmpl")
	authJWTTmpl                 = loadEmbedded("auth_jwt.tmpl")
	authMainTmpl                = loadEmbedded("auth_main.tmpl")
	authMiddlewareTmpl          = loadEmbedded("auth_middleware.tmpl")
	authUserModelTmpl           = loadEmbedded("auth_user_model.tmpl")
	challenge30DaysReadmeTmpl   = loadEmbedded("challenge30_days_readme.tmpl")
	challengeDay01TestTmpl      = loadEmbedded("challenge_day01_test.tmpl")
	challengeDay01Tmpl          = loadEmbedded("challenge_day01.tmpl")
	challengeDay02TestTmpl      = loadEmbedded("challenge_day02_test.tmpl")
	challengeDay02Tmpl          = loadEmbedded("challenge_day02.tmpl")
	challengeDay08TestTmpl      = loadEmbedded("challenge_day08_test.tmpl")
	challengeDay08Tmpl          = loadEmbedded("challenge_day08.tmpl")
	challengeDay15TestTmpl      = loadEmbedded("challenge_day15_test.tmpl")
	challengeDay15Tmpl          = loadEmbedded("challenge_day15.tmpl")
	challengeDay22TestTmpl      = loadEmbedded("challenge_day22_test.tmpl")
	challengeDay22Tmpl          = loadEmbedded("challenge_day22.tmpl")
	cleanArchEntityTmpl         = loadEmbedded("clean_arch_entity.tmpl")
	cleanArchErrorsTmpl         = loadEmbedded("clean_arch_errors.tmpl")
	cleanArchHandlerTmpl        = loadEmbedded("clean_arch_handler.tmpl")
	cleanArchMainTmpl           = loadEmbedded("clean_arch_main.tmpl")
	cleanArchRepoTmpl           = loadEmbedded("clean_arch_repo.tmpl")
	cleanArchUsecaseTmpl        = loadEmbedded("clean_arch_usecase.tmpl")
	codeReview01BuggyTmpl       = loadEmbedded("code_review01_buggy.tmpl")
	codeReview01TestTmpl        = loadEmbedded("code_review01_test.tmpl")
	codeReview02BuggyTmpl       = loadEmbedded("code_review02_buggy.tmpl")
	codeReview02TestTmpl        = loadEmbedded("code_review02_test.tmpl")
	codeReview03BuggyTmpl       = loadEmbedded("code_review03_buggy.tmpl")
	codeReview03TestTmpl        = loadEmbedded("code_review03_test.tmpl")
	codeReviewReadmeTmpl        = loadEmbedded("code_review_readme.tmpl")
	cronJobsTmpl                = loadEmbedded("cron_jobs.tmpl")
	cronMainTmpl                = loadEmbedded("cron_main.tmpl")
	cronSchedulerTmpl           = loadEmbedded("cron_scheduler.tmpl")
	goCronSchedulerTmpl         = loadEmbedded("go_cron_scheduler.tmpl")
	goCronJobsTmpl              = loadEmbedded("go_cron_jobs.tmpl")
	goCronConfigTmpl            = loadEmbedded("go_cron_config.tmpl")
	goCronHealthTmpl            = loadEmbedded("go_cron_health.tmpl")
	goCronDockerfileTmpl        = loadEmbedded("go_cron_dockerfile.tmpl")
	dsaLinkedListTestTmpl       = loadEmbedded("dsa_linked_list_test.tmpl")
	dsaLinkedListTmpl           = loadEmbedded("dsa_linked_list.tmpl")
	dsaQueueTestTmpl            = loadEmbedded("dsa_queue_test.tmpl")
	dsaQueueTmpl                = loadEmbedded("dsa_queue.tmpl")
	dsaRecursionTestTmpl        = loadEmbedded("dsa_recursion_test.tmpl")
	dsaRecursionTmpl            = loadEmbedded("dsa_recursion.tmpl")
	dsaSearchingTestTmpl        = loadEmbedded("dsa_searching_test.tmpl")
	dsaSearchingTmpl            = loadEmbedded("dsa_searching.tmpl")
	dsaSortingTestTmpl          = loadEmbedded("dsa_sorting_test.tmpl")
	dsaSortingTmpl              = loadEmbedded("dsa_sorting.tmpl")
	dsaStackTestTmpl            = loadEmbedded("dsa_stack_test.tmpl")
	dsaStackTmpl                = loadEmbedded("dsa_stack.tmpl")
	fullstackApiTsTmpl          = loadEmbedded("fullstack_api_ts.tmpl")
	fullstackAppTsxTmpl         = loadEmbedded("fullstack_app_tsx.tmpl")
	fullstackBackendGoModTmpl   = loadEmbedded("fullstack_backend_go_mod.tmpl")
	fullstackBackendMainTmpl    = loadEmbedded("fullstack_backend_main.tmpl")
	fullstackCorsTmpl           = loadEmbedded("fullstack_cors.tmpl")
	fullstackHandlerTmpl        = loadEmbedded("fullstack_handler.tmpl")
	fullstackHeaderTmpl         = loadEmbedded("fullstack_header.tmpl")
	fullstackIndexCssTmpl       = loadEmbedded("fullstack_index_css.tmpl")
	fullstackIndexHtmlTmpl      = loadEmbedded("fullstack_index_html.tmpl")
	fullstackMainTsxTmpl        = loadEmbedded("fullstack_main_tsx.tmpl")
	fullstackMakefileTmpl       = loadEmbedded("fullstack_makefile.tmpl")
	fullstackPackageJsonTmpl    = loadEmbedded("fullstack_package_json.tmpl")
	fullstackReadmeTmpl         = loadEmbedded("fullstack_readme.tmpl")
	fullstackViteConfigTmpl     = loadEmbedded("fullstack_vite_config.tmpl")
	fullstackGitignore          = loadEmbedded("fullstack_gitignore.tmpl")
	fullstackTsconfigTmpl       = loadEmbedded("fullstack_tsconfig.tmpl")
	fullstackTsconfigNodeTmpl   = loadEmbedded("fullstack_tsconfig_node.tmpl")
	fullstackTailwindConfigTmpl = loadEmbedded("fullstack_tailwind_config.tmpl")
	fullstackPostcssConfigTmpl  = loadEmbedded("fullstack_postcss_config.tmpl")
	goAPIMainTmpl               = loadEmbedded("go_api_main.tmpl")
	goAPIMiddlewareLoggingTmpl  = loadEmbedded("go_api_middleware_logging.tmpl")
	goAPIMiddlewareCORSTmpl     = loadEmbedded("go_api_middleware_cors.tmpl")
	goAPIMiddlewareAuthTmpl     = loadEmbedded("go_api_middleware_auth.tmpl")
	goAPIValidatorTmpl          = loadEmbedded("go_api_validator.tmpl")
	goAPIHandlerTestTmpl        = loadEmbedded("go_api_handler_test.tmpl")
	goAPIServiceTestTmpl        = loadEmbedded("go_api_service_test.tmpl")
	goAPIDockerfileTmpl         = loadEmbedded("go_api_dockerfile.tmpl")
	goCLIMainTmpl               = loadEmbedded("go_cli_main.tmpl")
	goCLIRootTmpl               = loadEmbedded("go_cli_root.tmpl")
	goCLIVersionTmpl            = loadEmbedded("go_cli_version.tmpl")
	goCLIConfigCmdTmpl          = loadEmbedded("go_cli_config_cmd.tmpl")
	goCLICompletionTmpl         = loadEmbedded("go_cli_completion.tmpl")
	goCLIConfigTmpl             = loadEmbedded("go_cli_config.tmpl")
	goCLIConfigTestTmpl         = loadEmbedded("go_cli_config_test.tmpl")
	goCLIOutputTmpl             = loadEmbedded("go_cli_output.tmpl")
	goCLIMakefileTmpl           = loadEmbedded("go_cli_makefile.tmpl")
	goCLIGoreleaserTmpl         = loadEmbedded("go_cli_goreleaser.tmpl")
	goConfigTmpl                = loadEmbedded("go_config.tmpl")
	goGRPCClientTmpl            = loadEmbedded("go_grpc_client.tmpl")
	goGRPCMakefileTmpl          = loadEmbedded("go_grpc_makefile.tmpl")
	goGRPCProtoTmpl             = loadEmbedded("go_grpc_proto.tmpl")
	goGRPCServerTmpl            = loadEmbedded("go_grpc_server.tmpl")
	goGRPCInterceptorsTmpl      = loadEmbedded("go_grpc_interceptors.tmpl")
	goGRPCHealthTmpl            = loadEmbedded("go_grpc_health.tmpl")
	goGRPCServerTestTmpl        = loadEmbedded("go_grpc_server_test.tmpl")
	goGRPCDockerfileTmpl        = loadEmbedded("go_grpc_dockerfile.tmpl")
	goHandlerTmpl               = loadEmbedded("go_handler.tmpl")
	goLibExampleTmpl            = loadEmbedded("go_lib_example.tmpl")
	goLibMainTmpl               = loadEmbedded("go_lib_main.tmpl")
	goLibTestTmpl               = loadEmbedded("go_lib_test.tmpl")
	goLibOptionsTmpl            = loadEmbedded("go_lib_options.tmpl")
	goLibErrorsTmpl             = loadEmbedded("go_lib_errors.tmpl")
	goLibBenchmarkTmpl          = loadEmbedded("go_lib_benchmark.tmpl")
	goLibDocTmpl                = loadEmbedded("go_lib_doc.tmpl")
	goLibCITmpl                 = loadEmbedded("go_lib_ci.tmpl")
	goModelTmpl                 = loadEmbedded("go_model.tmpl")
	goRepositoryTmpl            = loadEmbedded("go_repository.tmpl")
	goServiceTmpl               = loadEmbedded("go_service.tmpl")
	goTUIMainTmpl               = loadEmbedded("go_tui_main.tmpl")
	goTUIModelTmpl              = loadEmbedded("go_tui_model.tmpl")
	goTUIStylesTmpl             = loadEmbedded("go_tui_styles.tmpl")
	goTUIKeysTmpl               = loadEmbedded("go_tui_keys.tmpl")
	goTUIListTmpl               = loadEmbedded("go_tui_list.tmpl")
	goTUIInputTmpl              = loadEmbedded("go_tui_input.tmpl")
	goTUIHomeTmpl               = loadEmbedded("go_tui_home.tmpl")
	goWorkerJobTmpl             = loadEmbedded("go_worker_job.tmpl")
	goWorkerJobTestTmpl         = loadEmbedded("go_worker_job_test.tmpl")
	goWorkerMainTmpl            = loadEmbedded("go_worker_main.tmpl")
	goWorkerQueueTmpl           = loadEmbedded("go_worker_queue.tmpl")
	goWorkerQueueTestTmpl       = loadEmbedded("go_worker_queue_test.tmpl")
	goWorkerRetryTmpl           = loadEmbedded("go_worker_retry.tmpl")
	goWorkerMetricsTmpl         = loadEmbedded("go_worker_metrics.tmpl")
	goWorkerPoolTmpl            = loadEmbedded("go_worker_pool.tmpl")
	goWorkerDockerfileTmpl      = loadEmbedded("go_worker_dockerfile.tmpl")
	gitignoreGoTmpl             = loadEmbedded("gitignore_go.tmpl")
	graphqlConfigTmpl           = loadEmbedded("graphql_config.tmpl")
	graphqlMainTmpl             = loadEmbedded("graphql_main.tmpl")
	graphqlResolverTmpl         = loadEmbedded("graphql_resolver.tmpl")
	graphqlSchemaTmpl           = loadEmbedded("graphql_schema.tmpl")
	graphqlDataloaderTmpl       = loadEmbedded("graphql_dataloader.tmpl")
	graphqlResolverTestTmpl     = loadEmbedded("graphql_resolver_test.tmpl")
	graphqlMiddlewareTmpl       = loadEmbedded("graphql_middleware.tmpl")
	graphqlDockerfileTmpl       = loadEmbedded("graphql_dockerfile.tmpl")
	kafkaConsumerMainTmpl       = loadEmbedded("kafka_consumer_main.tmpl")
	kafkaConsumerTmpl           = loadEmbedded("kafka_consumer.tmpl")
	kafkaDockerComposeTmpl      = loadEmbedded("kafka_docker_compose.tmpl")
	kafkaProducerMainTmpl       = loadEmbedded("kafka_producer_main.tmpl")
	kafkaProducerTmpl           = loadEmbedded("kafka_producer.tmpl")
	// HTTP/HTMX
	goHtmxMainTmpl   = loadEmbedded("go_htmx_main.tmpl")
	goHtmxIndexTmpl  = loadEmbedded("go_htmx_index.tmpl")
	goHtmxListTmpl   = loadEmbedded("go_htmx_list.tmpl")
	goHtmxReadmeTmpl = loadEmbedded("go_htmx_readme.tmpl")
	// K8s Operator
	goK8sMainTmpl       = loadEmbedded("go_k8s_main.tmpl")
	goK8sApiTmpl        = loadEmbedded("go_k8s_api.tmpl")
	goK8sControllerTmpl = loadEmbedded("go_k8s_controller.tmpl")
	goK8sDockerfileTmpl = loadEmbedded("go_k8s_dockerfile.tmpl")
	goK8sManifestsTmpl  = loadEmbedded("go_k8s_manifests.tmpl")
	goK8sReadmeTmpl     = loadEmbedded("go_k8s_readme.tmpl")
	// WASM
	goWasmMainTmpl     = loadEmbedded("go_wasm_main.tmpl")
	goWasmServerTmpl   = loadEmbedded("go_wasm_server.tmpl")
	goWasmIndexTmpl    = loadEmbedded("go_wasm_index.tmpl")
	goWasmMakefileTmpl = loadEmbedded("go_wasm_makefile.tmpl")
	goWasmReadmeTmpl   = loadEmbedded("go_wasm_readme.tmpl")
	// Lambda
	lambdaHandlerTmpl                  = loadEmbedded("lambda_handler.tmpl")
	lambdaMainTmpl                     = loadEmbedded("lambda_main.tmpl")
	lambdaMakefileTmpl                 = loadEmbedded("lambda_makefile.tmpl")
	lambdaSAMTmpl                      = loadEmbedded("lambda_sam.tmpl")
	lambdaMiddlewareTmpl               = loadEmbedded("lambda_middleware.tmpl")
	lambdaHandlerTestTmpl              = loadEmbedded("lambda_handler_test.tmpl")
	lambdaSAMEnhancedTmpl              = loadEmbedded("lambda_sam_enhanced.tmpl")
	lambdaLocalTmpl                    = loadEmbedded("lambda_local.tmpl")
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
	learnFrontendReadmeTmpl            = loadEmbedded("learn_frontend_readme.tmpl")
	learnFrontendHTMLIndexTmpl         = loadEmbedded("learn_frontend_html_index.tmpl")
	learnFrontendHTMLReadmeTmpl        = loadEmbedded("learn_frontend_html_readme.tmpl")
	learnFrontendCSSIndexTmpl          = loadEmbedded("learn_frontend_css_index.tmpl")
	learnFrontendCSSStyleTmpl          = loadEmbedded("learn_frontend_css_style.tmpl")
	learnFrontendCSSReadmeTmpl         = loadEmbedded("learn_frontend_css_readme.tmpl")
	learnFrontendJSIndexTmpl           = loadEmbedded("learn_frontend_js_index.tmpl")
	learnFrontendJSScriptTmpl          = loadEmbedded("learn_frontend_js_script.tmpl")
	learnFrontendJSReadmeTmpl          = loadEmbedded("learn_frontend_js_readme.tmpl")
	learnFrontendDOMIndexTmpl          = loadEmbedded("learn_frontend_dom_index.tmpl")
	learnFrontendDOMStyleTmpl          = loadEmbedded("learn_frontend_dom_style.tmpl")
	learnFrontendDOMScriptTmpl         = loadEmbedded("learn_frontend_dom_script.tmpl")
	learnFrontendFetchIndexTmpl        = loadEmbedded("learn_frontend_fetch_index.tmpl")
	learnFrontendFetchStyleTmpl        = loadEmbedded("learn_frontend_fetch_style.tmpl")
	learnFrontendFetchScriptTmpl       = loadEmbedded("learn_frontend_fetch_script.tmpl")
	// Advanced Learning
	learnSecurityReadmeTmpl  = loadEmbedded("learn_security_readme.tmpl")
	learnSecurityHashTmpl    = loadEmbedded("learn_security_hash.tmpl")
	learnSecurityHeadersTmpl = loadEmbedded("learn_security_headers.tmpl")
	learnSecuritySQLTmpl     = loadEmbedded("learn_security_sql.tmpl")
	learnSecurityXSSTmpl     = loadEmbedded("learn_security_xss.tmpl")
	learnDBReadmeTmpl        = loadEmbedded("learn_db_readme.tmpl")
	learnDBSQLMainTmpl       = loadEmbedded("learn_db_sql_main.tmpl")
	learnDBGormMainTmpl      = loadEmbedded("learn_db_gorm_main.tmpl")
	learnDBSQLCTmpl          = loadEmbedded("learn_db_sqlc.tmpl")
	learnPerfReadmeTmpl      = loadEmbedded("learn_perf_readme.tmpl")
	learnPerfBenchTestTmpl   = loadEmbedded("learn_perf_bench_test.tmpl")
	learnPerfPprofTmpl       = loadEmbedded("learn_perf_pprof.tmpl")
	learnPerfOptTmpl         = loadEmbedded("learn_perf_opt.tmpl")
	// Skill: Algo & System Design
	algoReadmeTmpl          = loadEmbedded("algo_readme.tmpl")
	algoEasyTmpl            = loadEmbedded("algo_easy_two_sum.tmpl")
	algoMediumTmpl          = loadEmbedded("algo_medium_lru_cache.tmpl")
	algoHardTmpl            = loadEmbedded("algo_hard_merge_k_lists.tmpl")
	sysDesignReadmeTmpl     = loadEmbedded("sys_design_readme.tmpl")
	sysDesignDocTmpl        = loadEmbedded("sys_design_doc.md.tmpl")
	sysDesignInterfacesTmpl = loadEmbedded("sys_design_interfaces.go.tmpl")
	miniProjRateLimiterTmpl = loadEmbedded("mini_project_rate_limiter.tmpl")
	miniProjKVTmpl          = loadEmbedded("mini_project_kv.tmpl")

	microserviceDockerfileTmpl = loadEmbedded("microservice_dockerfile.tmpl")
	microserviceHandlerTmpl    = loadEmbedded("microservice_handler.tmpl")
	microserviceHealthTmpl     = loadEmbedded("microservice_health.tmpl")
	microserviceLoggingTmpl    = loadEmbedded("microservice_logging.tmpl")
	microserviceMainTmpl       = loadEmbedded("microservice_main.tmpl")
	microserviceMakefileTmpl   = loadEmbedded("microservice_makefile.tmpl")
	miniProjectReadmeTmpl      = loadEmbedded("mini_project_readme.tmpl")
	monorepoAPITmpl            = loadEmbedded("monorepo_api.tmpl")
	monorepoConfigTmpl         = loadEmbedded("monorepo_config.tmpl")
	monorepoLoggerTmpl         = loadEmbedded("monorepo_logger.tmpl")
	monorepoMakefileTmpl       = loadEmbedded("monorepo_makefile.tmpl")
	monorepoWorkerTmpl         = loadEmbedded("monorepo_worker.tmpl")
	redisCacheTmpl             = loadEmbedded("redis_cache.tmpl")
	redisDockerComposeTmpl     = loadEmbedded("redis_docker_compose.tmpl")
	redisMainTmpl              = loadEmbedded("redis_main.tmpl")
	redisPubSubTmpl            = loadEmbedded("redis_pub_sub.tmpl")
	readmeTmpl                 = loadEmbedded("readme.tmpl")
	refactoring01BeforeTmpl    = loadEmbedded("refactoring01_before.tmpl")
	refactoring01HintsTmpl     = loadEmbedded("refactoring01_hints.tmpl")
	refactoring02BeforeTmpl    = loadEmbedded("refactoring02_before.tmpl")
	refactoring02HintsTmpl     = loadEmbedded("refactoring02_hints.tmpl")
	refactoring03BeforeTmpl    = loadEmbedded("refactoring03_before.tmpl")
	refactoring03HintsTmpl     = loadEmbedded("refactoring03_hints.tmpl")
	refactoringReadmeTmpl      = loadEmbedded("refactoring_readme.tmpl")
	todoCliMainTmpl            = loadEmbedded("todo_cli_main.tmpl")
	todoCliReadmeTmpl          = loadEmbedded("todo_cli_readme.tmpl")
	todoCliTestTmpl            = loadEmbedded("todo_cli_test.tmpl")
	urlShortenerMainTmpl       = loadEmbedded("url_shortener_main.tmpl")
	urlShortenerReadmeTmpl     = loadEmbedded("url_shortener_readme.tmpl")
	urlShortenerTestTmpl       = loadEmbedded("url_shortener_test.tmpl")
	websocketClientTmpl        = loadEmbedded("websocket_client.tmpl")
	websocketHTMLTmpl          = loadEmbedded("websocket_html.tmpl")
	goMicroserviceTestTmpl     = loadEmbedded("go_microservice_test.tmpl")
	goWebsocketTestTmpl        = loadEmbedded("go_websocket_test.tmpl")
	websocketHubTmpl           = loadEmbedded("websocket_hub.tmpl")
	websocketMainTmpl          = loadEmbedded("websocket_main.tmpl")
	// Debugging templates
	debugReadmeTmpl        = loadEmbedded("debug_readme.tmpl")
	debugPrintMainTmpl     = loadEmbedded("debug_print_main.tmpl")
	debugPrintTestTmpl     = loadEmbedded("debug_print_test.tmpl")
	debugLoggingMainTmpl   = loadEmbedded("debug_logging_main.tmpl")
	debugLoggingTestTmpl   = loadEmbedded("debug_logging_test.tmpl")
	debugProfilingMainTmpl = loadEmbedded("debug_profiling_main.tmpl")
	debugProfilingTestTmpl = loadEmbedded("debug_profiling_test.tmpl")
	debugTracingMainTmpl   = loadEmbedded("debug_tracing_main.tmpl")
	debugTracingTestTmpl   = loadEmbedded("debug_tracing_test.tmpl")
	debugMemoryMainTmpl    = loadEmbedded("debug_memory_main.tmpl")
	debugMemoryTestTmpl    = loadEmbedded("debug_memory_test.tmpl")
	debugDelveReadmeTmpl   = loadEmbedded("debug_delve_readme.tmpl")
	// TDD templates
	tddReadmeTmpl          = loadEmbedded("tdd_readme.tmpl")
	tddRedGreenMainTmpl    = loadEmbedded("tdd_red_green_main.tmpl")
	tddRedGreenTestTmpl    = loadEmbedded("tdd_red_green_test.tmpl")
	tddMockingMainTmpl     = loadEmbedded("tdd_mocking_main.tmpl")
	tddMockingTestTmpl     = loadEmbedded("tdd_mocking_test.tmpl")
	tddRefactoringMainTmpl = loadEmbedded("tdd_refactoring_main.tmpl")
	tddRefactoringTestTmpl = loadEmbedded("tdd_refactoring_test.tmpl")
	tddBDDMainTmpl         = loadEmbedded("tdd_bdd_main.tmpl")
	tddBDDTestTmpl         = loadEmbedded("tdd_bdd_test.tmpl")
	tddExercisesReadmeTmpl = loadEmbedded("tdd_exercises_readme.tmpl")
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
				"internal/middleware",
				"internal/validator",
				"pkg/config",
			},
			Files: []FileTemplate{
				{Path: "cmd/api/main.go", Content: goAPIMainTmpl},
				{Path: "internal/handler/handler.go", Content: goHandlerTmpl},
				{Path: "internal/handler/handler_test.go", Content: goAPIHandlerTestTmpl},
				{Path: "internal/service/service.go", Content: goServiceTmpl},
				{Path: "internal/service/service_test.go", Content: goAPIServiceTestTmpl},
				{Path: "internal/repository/repository.go", Content: goRepositoryTmpl},
				{Path: "internal/model/model.go", Content: goModelTmpl},
				{Path: "internal/middleware/logging.go", Content: goAPIMiddlewareLoggingTmpl},
				{Path: "internal/middleware/cors.go", Content: goAPIMiddlewareCORSTmpl},
				{Path: "internal/middleware/auth.go", Content: goAPIMiddlewareAuthTmpl},
				{Path: "internal/validator/validator.go", Content: goAPIValidatorTmpl},
				{Path: "pkg/config/config.go", Content: goConfigTmpl},
				{Path: "Dockerfile", Content: goAPIDockerfileTmpl},
				{Path: "README.md", Content: readmeTmpl},
				{Path: ".gitignore", Content: gitignoreGoTmpl},
			},
		},
		"go-cli": {
			Name:        "go-cli",
			Description: "Go CLI application with Cobra",
			Directories: []string{
				"cmd",
				"internal/config",
				"internal/output",
			},
			Files: []FileTemplate{
				{Path: "main.go", Content: goCLIMainTmpl},
				{Path: "cmd/root.go", Content: goCLIRootTmpl},
				{Path: "cmd/version.go", Content: goCLIVersionTmpl},
				{Path: "cmd/config.go", Content: goCLIConfigCmdTmpl},
				{Path: "cmd/completion.go", Content: goCLICompletionTmpl},
				{Path: "internal/config/config.go", Content: goCLIConfigTmpl},
				{Path: "internal/config/config_test.go", Content: goCLIConfigTestTmpl},
				{Path: "internal/output/output.go", Content: goCLIOutputTmpl},
				{Path: "Makefile", Content: goCLIMakefileTmpl},
				{Path: ".goreleaser.yaml", Content: goCLIGoreleaserTmpl},
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
				".github/workflows",
			},
			Files: []FileTemplate{
				{Path: "{{.ProjectName}}.go", Content: goLibMainTmpl},
				{Path: "{{.ProjectName}}_test.go", Content: goLibTestTmpl},
				{Path: "options.go", Content: goLibOptionsTmpl},
				{Path: "errors.go", Content: goLibErrorsTmpl},
				{Path: "doc.go", Content: goLibDocTmpl},
				{Path: "benchmark_test.go", Content: goLibBenchmarkTmpl},
				{Path: "examples/main.go", Content: goLibExampleTmpl},
				{Path: ".github/workflows/ci.yml", Content: goLibCITmpl},
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
				"internal/interceptors",
				"internal/health",
				"proto",
			},
			Files: []FileTemplate{
				{Path: "cmd/server/main.go", Content: goGRPCServerTmpl},
				{Path: "cmd/server/main_test.go", Content: goGRPCServerTestTmpl},
				{Path: "cmd/client/main.go", Content: goGRPCClientTmpl},
				{Path: "internal/interceptors/interceptors.go", Content: goGRPCInterceptorsTmpl},
				{Path: "internal/health/health.go", Content: goGRPCHealthTmpl},
				{Path: "proto/service.proto", Content: goGRPCProtoTmpl},
				{Path: "Makefile", Content: goGRPCMakefileTmpl},
				{Path: "Dockerfile", Content: goGRPCDockerfileTmpl},
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
				"internal/retry",
				"internal/metrics",
				"internal/worker",
			},
			Files: []FileTemplate{
				{Path: "cmd/worker/main.go", Content: goWorkerMainTmpl},
				{Path: "internal/job/job.go", Content: goWorkerJobTmpl},
				{Path: "internal/job/job_test.go", Content: goWorkerJobTestTmpl},
				{Path: "internal/queue/queue.go", Content: goWorkerQueueTmpl},
				{Path: "internal/queue/queue_test.go", Content: goWorkerQueueTestTmpl},
				{Path: "internal/retry/retry.go", Content: goWorkerRetryTmpl},
				{Path: "internal/metrics/metrics.go", Content: goWorkerMetricsTmpl},
				{Path: "internal/worker/pool.go", Content: goWorkerPoolTmpl},
				{Path: "Dockerfile", Content: goWorkerDockerfileTmpl},
				{Path: "README.md", Content: readmeTmpl},
				{Path: ".gitignore", Content: gitignoreGoTmpl},
			},
		},
		"go-tui": {
			Name:        "go-tui",
			Description: "Terminal UI app with Bubbletea",
			Directories: []string{
				"internal/ui",
				"internal/ui/styles",
				"internal/ui/keys",
				"internal/ui/components",
				"internal/ui/views",
			},
			Files: []FileTemplate{
				{Path: "main.go", Content: goTUIMainTmpl},
				{Path: "internal/ui/model.go", Content: goTUIModelTmpl},
				{Path: "internal/ui/styles/styles.go", Content: goTUIStylesTmpl},
				{Path: "internal/ui/keys/keys.go", Content: goTUIKeysTmpl},
				{Path: "internal/ui/components/list.go", Content: goTUIListTmpl},
				{Path: "internal/ui/components/input.go", Content: goTUIInputTmpl},
				{Path: "internal/ui/views/home.go", Content: goTUIHomeTmpl},
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
		"algorithm-challenges": {
			Name:        "algorithm-challenges",
			Description: "Algo challenges (Two Sum, LRU Cache, Merge K Lists)",
			Directories: []string{
				"easy",
				"medium",
				"hard",
			},
			Files: []FileTemplate{
				{Path: "README.md", Content: algoReadmeTmpl},
				{Path: "easy/main_test.go", Content: algoEasyTmpl},
				{Path: "medium/main_test.go", Content: algoMediumTmpl},
				{Path: "hard/main_test.go", Content: algoHardTmpl},
				{Path: ".gitignore", Content: gitignoreGoTmpl},
			},
		},
		"system-design-exercise": {
			Name:        "system-design-exercise",
			Description: "System Design practice (URL Shortener doc + interfaces)",
			Directories: []string{
				"design-docs",
				"prototypes",
			},
			Files: []FileTemplate{
				{Path: "README.md", Content: sysDesignReadmeTmpl},
				{Path: "design-docs/url-shortener.md", Content: sysDesignDocTmpl},
				{Path: "prototypes/main.go", Content: sysDesignInterfacesTmpl},
				{Path: "go.mod", Content: "module system-design\ngo 1.22"},
			},
		},
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
				"rate-limiter",
				"kv-store",
			},
			Files: []FileTemplate{
				{Path: "README.md", Content: miniProjectReadmeTmpl},
				{Path: "todo-cli/README.md", Content: todoCliReadmeTmpl},
				{Path: "todo-cli/main.go", Content: todoCliMainTmpl},
				{Path: "todo-cli/main_test.go", Content: todoCliTestTmpl},
				{Path: "url-shortener/README.md", Content: urlShortenerReadmeTmpl},
				{Path: "url-shortener/main.go", Content: urlShortenerMainTmpl},
				{Path: "url-shortener/main_test.go", Content: urlShortenerTestTmpl},
				{Path: "rate-limiter/main.go", Content: miniProjRateLimiterTmpl},
				{Path: "kv-store/main.go", Content: miniProjKVTmpl},
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
		"learn-security": {
			Name:        "learn-security",
			Description: "Learn Go security (SQL injection, XSS, hashing)",
			Directories: []string{
				"vulnerabilities/sql_injection",
				"vulnerabilities/xss",
				"auth/hashing",
				"config/secure",
			},
			Files: []FileTemplate{
				{Path: "README.md", Content: learnSecurityReadmeTmpl},
				{Path: "vulnerabilities/sql_injection/main.go", Content: learnSecuritySQLTmpl},
				{Path: "vulnerabilities/xss/main.go", Content: learnSecurityXSSTmpl},
				{Path: "auth/hashing/main.go", Content: learnSecurityHashTmpl},
				{Path: "config/secure/middleware.go", Content: learnSecurityHeadersTmpl},
				{Path: "go.mod", Content: `module learn-security
go 1.22
require (
	golang.org/x/crypto v0.17.0
	github.com/mattn/go-sqlite3 v1.14.19
)`},
				{Path: ".gitignore", Content: gitignoreGoTmpl},
			},
		},
		"learn-database": {
			Name:        "learn-database",
			Description: "Compare Database approaches (Raw SQL, GORM, sqlc)",
			Directories: []string{
				"raw_sql",
				"gorm",
				"sqlc",
			},
			Files: []FileTemplate{
				{Path: "README.md", Content: learnDBReadmeTmpl},
				{Path: "raw_sql/main.go", Content: learnDBSQLMainTmpl},
				{Path: "gorm/main.go", Content: learnDBGormMainTmpl},
				{Path: "sqlc/schema.sql", Content: learnDBSQLCTmpl},
				{Path: "go.mod", Content: `module learn-database
go 1.22
require (
	github.com/mattn/go-sqlite3 v1.14.19
	gorm.io/gorm v1.25.5
	gorm.io/driver/sqlite v1.5.0
)`},
				{Path: ".gitignore", Content: gitignoreGoTmpl},
			},
		},
		"learn-performance": {
			Name:        "learn-performance",
			Description: "Learn Benchmarking, Profiling & Optimization",
			Directories: []string{
				"benchmarking",
				"profiling",
				"optimization",
			},
			Files: []FileTemplate{
				{Path: "README.md", Content: learnPerfReadmeTmpl},
				{Path: "benchmarking/concat_test.go", Content: learnPerfBenchTestTmpl},
				{Path: "profiling/main.go", Content: learnPerfPprofTmpl},
				{Path: "optimization/slice_test.go", Content: learnPerfOptTmpl},
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
				{Path: "cmd/server/main_test.go", Content: goMicroserviceTestTmpl},
				{Path: "internal/handler/handler.go", Content: microserviceHandlerTmpl},
				{Path: "internal/middleware/logging.go", Content: microserviceLoggingTmpl},
				{Path: "internal/health/health.go", Content: microserviceHealthTmpl},
				{Path: "Dockerfile", Content: microserviceDockerfileTmpl},
				{Path: "Makefile", Content: microserviceMakefileTmpl},
				{Path: "README.md", Content: readmeTmpl},
				{Path: ".gitignore", Content: gitignoreGoTmpl},
			},
		},
		"go-web-htmx": {
			Name:        "go-web-htmx",
			Description: "SSR Web App with Go + HTMX + Tailwind",
			Directories: []string{
				"cmd/server",
				"templates",
			},
			Files: []FileTemplate{
				{Path: "cmd/server/main.go", Content: goHtmxMainTmpl},
				{Path: "templates/index.html", Content: goHtmxIndexTmpl},
				{Path: "templates/list.html", Content: goHtmxListTmpl},
				{Path: "README.md", Content: goHtmxReadmeTmpl},
				{Path: ".gitignore", Content: gitignoreGoTmpl},
			},
		},
		"go-k8s-operator": {
			Name:        "go-k8s-operator",
			Description: "Kubernetes Operator (controller-runtime)",
			Directories: []string{
				"api/v1alpha1",
				"internal/controller",
				"manifests",
			},
			Files: []FileTemplate{
				{Path: "main.go", Content: goK8sMainTmpl},
				{Path: "api/v1alpha1/myresource_types.go", Content: goK8sApiTmpl},
				{Path: "internal/controller/myresource_controller.go", Content: goK8sControllerTmpl},
				{Path: "manifests/crd.yaml", Content: goK8sManifestsTmpl},
				{Path: "Dockerfile", Content: goK8sDockerfileTmpl},
				{Path: "README.md", Content: goK8sReadmeTmpl},
				{Path: "go.mod", Content: `module {{.ModuleName}}
go 1.22
require (
	k8s.io/apimachinery v0.29.0
	k8s.io/client-go v0.29.0
	sigs.k8s.io/controller-runtime v0.17.0
)`},
				{Path: ".gitignore", Content: gitignoreGoTmpl},
			},
		},
		"go-wasm": {
			Name:        "go-wasm",
			Description: "WebAssembly App (Go compilation to WASM)",
			Directories: []string{},
			Files: []FileTemplate{
				{Path: "main.go", Content: goWasmMainTmpl},
				{Path: "server.go", Content: goWasmServerTmpl},
				{Path: "index.html", Content: goWasmIndexTmpl},
				{Path: "Makefile", Content: goWasmMakefileTmpl},
				{Path: "README.md", Content: goWasmReadmeTmpl},
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
				{Path: "internal/hub/hub_test.go", Content: goWebsocketTestTmpl},
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
				"graph/dataloaders",
				"internal/middleware",
			},
			Files: []FileTemplate{
				{Path: "cmd/server/main.go", Content: graphqlMainTmpl},
				{Path: "graph/schema.graphqls", Content: graphqlSchemaTmpl},
				{Path: "graph/resolver.go", Content: graphqlResolverTmpl},
				{Path: "graph/resolver_test.go", Content: graphqlResolverTestTmpl},
				{Path: "graph/dataloaders/dataloaders.go", Content: graphqlDataloaderTmpl},
				{Path: "internal/middleware/middleware.go", Content: graphqlMiddlewareTmpl},
				{Path: "gqlgen.yml", Content: graphqlConfigTmpl},
				{Path: "Dockerfile", Content: graphqlDockerfileTmpl},
				{Path: "README.md", Content: readmeTmpl},
				{Path: ".gitignore", Content: gitignoreGoTmpl},
			},
		},
		"go-lambda": {
			Name:        "go-lambda",
			Description: "AWS Lambda function with SAM",
			Directories: []string{
				"cmd/lambda",
				"cmd/local",
				"internal/handler",
				"internal/middleware",
			},
			Files: []FileTemplate{
				{Path: "cmd/lambda/main.go", Content: lambdaMainTmpl},
				{Path: "cmd/local/main.go", Content: lambdaLocalTmpl},
				{Path: "internal/handler/handler.go", Content: lambdaHandlerTmpl},
				{Path: "internal/handler/handler_test.go", Content: lambdaHandlerTestTmpl},
				{Path: "internal/middleware/middleware.go", Content: lambdaMiddlewareTmpl},
				{Path: "template.yaml", Content: lambdaSAMEnhancedTmpl},
				{Path: "Makefile", Content: lambdaMakefileTmpl},
				{Path: "README.md", Content: readmeTmpl},
				{Path: ".gitignore", Content: gitignoreGoTmpl},
			},
		},
		"go-cron": {
			Name:        "go-cron",
			Description: "Scheduled jobs with cron",
			Directories: []string{
				"cmd/cron",
				"internal/jobs",
				"internal/scheduler",
				"internal/config",
				"internal/health",
			},
			Files: []FileTemplate{
				{Path: "cmd/cron/main.go", Content: cronMainTmpl},
				{Path: "internal/jobs/jobs.go", Content: cronJobsTmpl},
				{Path: "internal/jobs/examples.go", Content: goCronJobsTmpl},
				{Path: "internal/scheduler/scheduler.go", Content: cronSchedulerTmpl},
				{Path: "internal/scheduler/manager.go", Content: goCronSchedulerTmpl},
				{Path: "internal/config/config.go", Content: goCronConfigTmpl},
				{Path: "internal/health/health.go", Content: goCronHealthTmpl},
				{Path: "Dockerfile", Content: goCronDockerfileTmpl},
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
		"learn-frontend": {
			Name:        "learn-frontend",
			Description: "Learn frontend development (HTML, CSS, JavaScript)",
			Directories: []string{
				"01-html-basics",
				"02-css-fundamentals",
				"03-javascript-basics",
				"04-dom-manipulation",
				"05-fetch-api",
			},
			Files: []FileTemplate{
				{Path: "README.md", Content: learnFrontendReadmeTmpl},
				// HTML Basics
				{Path: "01-html-basics/index.html", Content: learnFrontendHTMLIndexTmpl},
				{Path: "01-html-basics/README.md", Content: learnFrontendHTMLReadmeTmpl},
				// CSS Fundamentals
				{Path: "02-css-fundamentals/index.html", Content: learnFrontendCSSIndexTmpl},
				{Path: "02-css-fundamentals/style.css", Content: learnFrontendCSSStyleTmpl},
				{Path: "02-css-fundamentals/README.md", Content: learnFrontendCSSReadmeTmpl},
				// JavaScript Basics
				{Path: "03-javascript-basics/index.html", Content: learnFrontendJSIndexTmpl},
				{Path: "03-javascript-basics/script.js", Content: learnFrontendJSScriptTmpl},
				{Path: "03-javascript-basics/README.md", Content: learnFrontendJSReadmeTmpl},
				// DOM Manipulation
				{Path: "04-dom-manipulation/index.html", Content: learnFrontendDOMIndexTmpl},
				{Path: "04-dom-manipulation/style.css", Content: learnFrontendDOMStyleTmpl},
				{Path: "04-dom-manipulation/script.js", Content: learnFrontendDOMScriptTmpl},
				// Fetch API
				{Path: "05-fetch-api/index.html", Content: learnFrontendFetchIndexTmpl},
				{Path: "05-fetch-api/style.css", Content: learnFrontendFetchStyleTmpl},
				{Path: "05-fetch-api/script.js", Content: learnFrontendFetchScriptTmpl},
			},
		},
		"learn-debugging": {
			Name:        "learn-debugging",
			Description: "Learn debugging techniques (print, logging, profiling, tracing)",
			Directories: []string{
				"01-print-debugging",
				"02-structured-logging",
				"03-profiling",
				"04-tracing",
				"05-memory-debugging",
				"06-delve-debugger",
			},
			Files: []FileTemplate{
				{Path: "README.md", Content: debugReadmeTmpl},
				{Path: "01-print-debugging/main.go", Content: debugPrintMainTmpl},
				{Path: "01-print-debugging/main_test.go", Content: debugPrintTestTmpl},
				{Path: "02-structured-logging/main.go", Content: debugLoggingMainTmpl},
				{Path: "02-structured-logging/main_test.go", Content: debugLoggingTestTmpl},
				{Path: "03-profiling/main.go", Content: debugProfilingMainTmpl},
				{Path: "03-profiling/main_test.go", Content: debugProfilingTestTmpl},
				{Path: "04-tracing/main.go", Content: debugTracingMainTmpl},
				{Path: "04-tracing/main_test.go", Content: debugTracingTestTmpl},
				{Path: "05-memory-debugging/main.go", Content: debugMemoryMainTmpl},
				{Path: "05-memory-debugging/main_test.go", Content: debugMemoryTestTmpl},
				{Path: "06-delve-debugger/README.md", Content: debugDelveReadmeTmpl},
				{Path: ".gitignore", Content: gitignoreGoTmpl},
			},
		},
		"learn-tdd": {
			Name:        "learn-tdd",
			Description: "Learn Test-Driven Development (Red-Green-Refactor)",
			Directories: []string{
				"01-red-green-refactor",
				"02-mocking",
				"03-refactoring",
				"04-bdd-style",
				"05-exercises",
			},
			Files: []FileTemplate{
				{Path: "README.md", Content: tddReadmeTmpl},
				{Path: "01-red-green-refactor/main.go", Content: tddRedGreenMainTmpl},
				{Path: "01-red-green-refactor/main_test.go", Content: tddRedGreenTestTmpl},
				{Path: "02-mocking/main.go", Content: tddMockingMainTmpl},
				{Path: "02-mocking/main_test.go", Content: tddMockingTestTmpl},
				{Path: "03-refactoring/main.go", Content: tddRefactoringMainTmpl},
				{Path: "03-refactoring/main_test.go", Content: tddRefactoringTestTmpl},
				{Path: "04-bdd-style/main.go", Content: tddBDDMainTmpl},
				{Path: "04-bdd-style/main_test.go", Content: tddBDDTestTmpl},
				{Path: "05-exercises/README.md", Content: tddExercisesReadmeTmpl},
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
		fmt.Println(len(result))
	}
	return result
}
