package templates

import (
	"embed"
	"fmt"
)

//go:embed embedded/*
var embeddedFS embed.FS

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

var builtInTemplates = map[string]Template{
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
			{Path: ".gitignore", Content: fullstackGitignoreTmpl},
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

// Template content definitions

var goAPIMainTmpl = `package main

import (
	"fmt"
	"log"
	"net/http"

	"{{.ModuleName}}/internal/handler"
	"{{.ModuleName}}/pkg/config"
)

func main() {
	cfg := config.Load()

	h := handler.New()

	http.HandleFunc("/", h.Health)
	http.HandleFunc("/api/v1/", h.HandleAPI)

	addr := fmt.Sprintf(":%s", cfg.Port)
	log.Printf("Server starting on %s", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
`

var goHandlerTmpl = `package handler

import (
	"encoding/json"
	"net/http"
)

type Handler struct{}

func New() *Handler {
	return &Handler{}
}

func (h *Handler) Health(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}

func (h *Handler) HandleAPI(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Hello from {{.ProjectName}}"})
}
`

var goServiceTmpl = `package service

// Service handles business logic
type Service struct{}

// New creates a new Service instance
func New() *Service {
	return &Service{}
}
`

var goRepositoryTmpl = `package repository

// Repository handles data persistence
type Repository struct{}

// New creates a new Repository instance
func New() *Repository {
	return &Repository{}
}
`

var goModelTmpl = `package model

// Define your data models here
`

var goConfigTmpl = `package config

import "os"

type Config struct {
	Port string
}

func Load() *Config {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return &Config{Port: port}
}
`

var goCLIMainTmpl = `package main

import "{{.ModuleName}}/cmd"

func main() {
	cmd.Execute()
}
`

var goCLIRootTmpl = `package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "{{.ProjectName}}",
	Short: "{{.ProjectName}} - A brief description",
	Long:  "{{.ProjectName}} - A longer description of your CLI application",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to {{.ProjectName}}!")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
`

var goLibMainTmpl = `// Package {{.ProjectName}} provides ...
package {{.PackageName}}

// Hello returns a greeting message
func Hello() string {
	return "Hello from {{.ProjectName}}!"
}
`

var goLibTestTmpl = `package {{.PackageName}}

import "testing"

func TestHello(t *testing.T) {
	got := Hello()
	want := "Hello from {{.ProjectName}}!"
	if got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}
`

var goLibExampleTmpl = `package main

import (
	"fmt"
	"{{.ModuleName}}"
)

func main() {
	fmt.Println({{.PackageName}}.Hello())
}
`

var readmeTmpl = `# {{.ProjectName}}

{{.Description}}

## Getting Started

### Prerequisites

- Go 1.21 or higher

### Installation

` + "```bash" + `
go mod download
` + "```" + `

### Running

` + "```bash" + `
go run ./cmd/...
` + "```" + `

## License

{{.License}}
`

var gitignoreGoTmpl = `# Binaries
*.exe
*.exe~
*.dll
*.so
*.dylib

# Test binary
*.test

# Output
/bin/
/dist/

# Dependency directories
/vendor/

# IDE
.idea/
.vscode/
*.swp
*.swo

# OS
.DS_Store
Thumbs.db

# Environment
.env
.env.local
`

// ============== Learning Templates ==============

var learnConcurrencyReadmeTmpl = `# Learn Go Concurrency

This project contains examples to learn Go concurrency patterns.

## Sections

1. **01-goroutines** - Basic goroutine usage
2. **02-channels** - Channel communication
3. **03-select** - Select statement for multiplexing
4. **04-sync** - Synchronization primitives (Mutex, WaitGroup)
5. **05-patterns** - Common concurrency patterns

## How to Run

` + "```bash" + `
cd 01-goroutines && go run main.go
` + "```" + `
`

var learnGoroutinesTmpl = `package main

import (
	"fmt"
	"time"
)

func main() {
	// Basic goroutine
	go sayHello("World")

	// Anonymous goroutine
	go func() {
		fmt.Println("Hello from anonymous goroutine!")
	}()

	// Multiple goroutines
	for i := 1; i <= 5; i++ {
		go func(n int) {
			fmt.Printf("Goroutine %d\n", n)
		}(i)
	}

	// Wait for goroutines (simple approach - use sync.WaitGroup in production)
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Done!")
}

func sayHello(name string) {
	fmt.Printf("Hello, %s!\n", name)
}
`

var learnChannelsTmpl = `package main

import "fmt"

func main() {
	// Unbuffered channel
	ch := make(chan string)
	go func() {
		ch <- "Hello from channel!"
	}()
	msg := <-ch
	fmt.Println(msg)

	// Buffered channel
	buffered := make(chan int, 3)
	buffered <- 1
	buffered <- 2
	buffered <- 3
	fmt.Println(<-buffered, <-buffered, <-buffered)

	// Channel with range
	numbers := make(chan int)
	go func() {
		for i := 1; i <= 5; i++ {
			numbers <- i
		}
		close(numbers)
	}()
	for n := range numbers {
		fmt.Printf("Received: %d\n", n)
	}

	// Channel direction (send-only, receive-only)
	ping := make(chan string, 1)
	pong := make(chan string, 1)
	pingPong(ping, pong)
}

func pingPong(ping chan<- string, pong <-chan string) {
	ping <- "ping"
	// This demonstrates send-only and receive-only channels
}
`

var learnSelectTmpl = `package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(50 * time.Millisecond)
		ch1 <- "from channel 1"
	}()

	go func() {
		time.Sleep(100 * time.Millisecond)
		ch2 <- "from channel 2"
	}()

	// Select receives from whichever channel is ready first
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println("Received", msg1)
		case msg2 := <-ch2:
			fmt.Println("Received", msg2)
		}
	}

	// Select with timeout
	timeout := make(chan string)
	select {
	case msg := <-timeout:
		fmt.Println(msg)
	case <-time.After(10 * time.Millisecond):
		fmt.Println("Timeout!")
	}

	// Non-blocking select with default
	nonBlocking := make(chan string)
	select {
	case msg := <-nonBlocking:
		fmt.Println(msg)
	default:
		fmt.Println("No message available")
	}
}
`

var learnSyncTmpl = `package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	// WaitGroup - wait for goroutines to finish
	var wg sync.WaitGroup
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			fmt.Printf("Worker %d done\n", n)
		}(i)
	}
	wg.Wait()
	fmt.Println("All workers done!")

	// Mutex - protect shared data
	var mu sync.Mutex
	counter := 0
	var wg2 sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg2.Add(1)
		go func() {
			defer wg2.Done()
			mu.Lock()
			counter++
			mu.Unlock()
		}()
	}
	wg2.Wait()
	fmt.Printf("Counter: %d\n", counter)

	// Atomic operations
	var atomicCounter int64
	var wg3 sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg3.Add(1)
		go func() {
			defer wg3.Done()
			atomic.AddInt64(&atomicCounter, 1)
		}()
	}
	wg3.Wait()
	fmt.Printf("Atomic counter: %d\n", atomicCounter)

	// Once - do something only once
	var once sync.Once
	for i := 0; i < 3; i++ {
		once.Do(func() {
			fmt.Println("This prints only once!")
		})
	}
}
`

var learnPatternsTmpl = `package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// Worker Pool Pattern
	fmt.Println("=== Worker Pool ===")
	jobs := make(chan int, 10)
	results := make(chan int, 10)

	// Start 3 workers
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// Send 5 jobs
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)

	// Collect results
	for r := 1; r <= 5; r++ {
		fmt.Printf("Result: %d\n", <-results)
	}

	// Fan-out, Fan-in Pattern
	fmt.Println("\n=== Fan-out, Fan-in ===")
	input := generate(1, 2, 3, 4, 5)
	c1 := square(input)
	c2 := square(input)
	for n := range merge(c1, c2) {
		fmt.Println(n)
	}
}

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, j)
		time.Sleep(10 * time.Millisecond)
		results <- j * 2
	}
}

func generate(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func merge(cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)
	output := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}
	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
`

// ============== Learn Testing Templates ==============

var learnTestingReadmeTmpl = `# Learn Go Testing

This project contains examples for Go testing techniques.

## Sections

1. **unit/** - Basic unit tests
2. **table/** - Table-driven tests
3. **benchmark/** - Benchmark tests

## How to Run

` + "```bash" + `
# Run all tests
go test ./...

# Run with verbose output
go test -v ./...

# Run benchmarks
go test -bench=. ./benchmark/
` + "```" + `
`

var learnUnitCalcTmpl = `package unit

import "fmt"

// Add returns the sum of two integers
func Add(a, b int) int {
	return a + b
}

// Subtract returns the difference of two integers
func Subtract(a, b int) int {
	return a - b
}

// Multiply returns the product of two integers
func Multiply(a, b int) int {
	return a * b
}

// Divide returns the quotient of two integers
func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, ErrDivideByZero
	}
	return a / b, nil
}

var ErrDivideByZero = fmt.Errorf("cannot divide by zero")
`

var learnUnitTestTmpl = `package unit

import (
	"testing"
)

func TestAdd(t *testing.T) {
	got := Add(2, 3)
	want := 5
	if got != want {
		t.Errorf("Add(2, 3) = %d; want %d", got, want)
	}
}

func TestSubtract(t *testing.T) {
	got := Subtract(5, 3)
	want := 2
	if got != want {
		t.Errorf("Subtract(5, 3) = %d; want %d", got, want)
	}
}

func TestDivide(t *testing.T) {
	t.Run("valid division", func(t *testing.T) {
		got, err := Divide(10, 2)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if got != 5 {
			t.Errorf("Divide(10, 2) = %d; want 5", got)
		}
	})

	t.Run("divide by zero", func(t *testing.T) {
		_, err := Divide(10, 0)
		if err == nil {
			t.Error("expected error for divide by zero")
		}
	})
}
`

var learnTableValidatorTmpl = `package table

import "regexp"

// IsValidEmail checks if the email format is valid
func IsValidEmail(email string) bool {
	pattern := ` + "`" + `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$` + "`" + `
	re := regexp.MustCompile(pattern)
	return re.MatchString(email)
}

// IsValidAge checks if age is within valid range
func IsValidAge(age int) bool {
	return age >= 0 && age <= 150
}
`

var learnTableTestTmpl = `package table

import "testing"

func TestIsValidEmail(t *testing.T) {
	tests := []struct {
		name  string
		email string
		want  bool
	}{
		{"valid email", "user@example.com", true},
		{"valid with subdomain", "user@mail.example.com", true},
		{"missing @", "userexample.com", false},
		{"missing domain", "user@", false},
		{"empty string", "", false},
		{"with plus", "user+tag@example.com", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsValidEmail(tt.email)
			if got != tt.want {
				t.Errorf("IsValidEmail(%q) = %v; want %v", tt.email, got, tt.want)
			}
		})
	}
}

func TestIsValidAge(t *testing.T) {
	tests := []struct {
		age  int
		want bool
	}{
		{0, true},
		{25, true},
		{150, true},
		{-1, false},
		{151, false},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := IsValidAge(tt.age)
			if got != tt.want {
				t.Errorf("IsValidAge(%d) = %v; want %v", tt.age, got, tt.want)
			}
		})
	}
}
`

var learnBenchFuncTmpl = `package benchmark

import "strings"

// ConcatPlus concatenates strings using +
func ConcatPlus(strs []string) string {
	result := ""
	for _, s := range strs {
		result += s
	}
	return result
}

// ConcatBuilder concatenates strings using strings.Builder
func ConcatBuilder(strs []string) string {
	var builder strings.Builder
	for _, s := range strs {
		builder.WriteString(s)
	}
	return builder.String()
}
`

var learnBenchTestTmpl = `package benchmark

import "testing"

var testStrings = []string{"hello", "world", "this", "is", "a", "test", "of", "string", "concatenation"}

func BenchmarkConcatPlus(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ConcatPlus(testStrings)
	}
}

func BenchmarkConcatBuilder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ConcatBuilder(testStrings)
	}
}
`

// ============== gRPC Templates ==============

var goGRPCServerTmpl = `package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	// TODO: Register your gRPC service here
	// pb.RegisterYourServiceServer(s, &server{})

	log.Printf("gRPC server listening on :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
`

var goGRPCClientTmpl = `package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// TODO: Create client and make RPC calls
	// c := pb.NewYourServiceClient(conn)
	// resp, err := c.YourMethod(context.Background(), &pb.Request{})

	log.Println("Connected to gRPC server")
}
`

var goGRPCProtoTmpl = `syntax = "proto3";

package {{.PackageName}};

option go_package = "{{.ModuleName}}/proto";

service {{.ProjectName}}Service {
  rpc SayHello (HelloRequest) returns (HelloResponse) {}
}

message HelloRequest {
  string name = 1;
}

message HelloResponse {
  string message = 1;
}
`

var goGRPCMakefileTmpl = `.PHONY: proto build run-server run-client

proto:
	protoc --go_out=. --go-grpc_out=. proto/*.proto

build:
	go build -o bin/server ./cmd/server
	go build -o bin/client ./cmd/client

run-server:
	go run ./cmd/server

run-client:
	go run ./cmd/client
`

// ============== Worker Templates ==============

var goWorkerMainTmpl = `package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"{{.ModuleName}}/internal/job"
	"{{.ModuleName}}/internal/queue"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Handle shutdown gracefully
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigCh
		log.Println("Shutting down...")
		cancel()
	}()

	// Create job queue
	q := queue.New(10)

	// Start workers
	numWorkers := 3
	for i := 1; i <= numWorkers; i++ {
		go worker(ctx, i, q)
	}

	// Add some sample jobs
	for i := 1; i <= 10; i++ {
		q.Enqueue(job.Job{ID: i, Payload: fmt.Sprintf("Task %d", i)})
	}

	// Wait for context cancellation
	<-ctx.Done()
	log.Println("Worker stopped")
}

func worker(ctx context.Context, id int, q *queue.Queue) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			j, ok := q.Dequeue()
			if !ok {
				continue
			}
			log.Printf("Worker %d processing: %s", id, j.Payload)
			j.Process()
		}
	}
}
`

var goWorkerJobTmpl = `package job

import (
	"log"
	"time"
)

type Job struct {
	ID      int
	Payload string
}

func (j *Job) Process() error {
	// Simulate work
	time.Sleep(100 * time.Millisecond)
	log.Printf("Job %d completed: %s", j.ID, j.Payload)
	return nil
}
`

var goWorkerQueueTmpl = `package queue

import (
	"{{.ModuleName}}/internal/job"
)

type Queue struct {
	jobs chan job.Job
}

func New(size int) *Queue {
	return &Queue{
		jobs: make(chan job.Job, size),
	}
}

func (q *Queue) Enqueue(j job.Job) {
	q.jobs <- j
}

func (q *Queue) Dequeue() (job.Job, bool) {
	select {
	case j := <-q.jobs:
		return j, true
	default:
		return job.Job{}, false
	}
}

func (q *Queue) Close() {
	close(q.jobs)
}
`

// ============== TUI Templates ==============

var goTUIMainTmpl = `package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"{{.ModuleName}}/internal/ui"
)

func main() {
	p := tea.NewProgram(ui.NewModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}
}
`

var goTUIModelTmpl = `package ui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#FAFAFA")).
			Background(lipgloss.Color("#7D56F4")).
			Padding(0, 1)

	itemStyle = lipgloss.NewStyle().
			PaddingLeft(2)

	selectedStyle = lipgloss.NewStyle().
			PaddingLeft(2).
			Foreground(lipgloss.Color("#7D56F4")).
			Bold(true)
)

type Model struct {
	choices  []string
	cursor   int
	selected map[int]struct{}
}

func NewModel() Model {
	return Model{
		choices:  []string{"Option 1", "Option 2", "Option 3", "Option 4"},
		selected: make(map[int]struct{}),
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		case "enter", " ":
			if _, ok := m.selected[m.cursor]; ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
			}
		}
	}
	return m, nil
}

func (m Model) View() string {
	s := titleStyle.Render("{{.ProjectName}}") + "\n\n"

	for i, choice := range m.choices {
		cursor := " "
		style := itemStyle
		if m.cursor == i {
			cursor = ">"
			style = selectedStyle
		}

		checked := " "
		if _, ok := m.selected[i]; ok {
			checked = "x"
		}

		s += style.Render(fmt.Sprintf("%s [%s] %s", cursor, checked, choice)) + "\n"
	}

	s += "\n(q to quit)\n"
	return s
}
`

// ============== Fullstack Templates ==============

var fullstackBackendMainTmpl = `package main

import (
	"log"
	"net/http"

	"{{.ProjectName}}-backend/internal/handler"
	"{{.ProjectName}}-backend/internal/middleware"
)

func main() {
	h := handler.New()

	mux := http.NewServeMux()
	mux.HandleFunc("/api/health", h.Health)
	mux.HandleFunc("/api/message", h.GetMessage)

	// Wrap with CORS middleware
	corsHandler := middleware.CORS(mux)

	log.Println("🚀 Backend server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", corsHandler))
}
`

var fullstackHandlerTmpl = `package handler

import (
	"encoding/json"
	"net/http"
)

type Handler struct{}

func New() *Handler {
	return &Handler{}
}

func (h *Handler) Health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}

func (h *Handler) GetMessage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]any{
		"message": "Hello from Go backend! 🚀",
		"project": "{{.ProjectName}}",
	})
}
`

var fullstackCorsTmpl = `package middleware

import "net/http"

// CORS middleware to allow frontend to communicate with backend
func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
`

var fullstackBackendGoModTmpl = `module {{.ProjectName}}-backend

go 1.21
`

var fullstackPackageJsonTmpl = `{
  "name": "{{.ProjectName}}-frontend",
  "private": true,
  "version": "0.0.0",
  "type": "module",
  "scripts": {
    "dev": "vite",
    "build": "tsc && vite build",
    "preview": "vite preview"
  },
  "dependencies": {
    "react": "^18.2.0",
    "react-dom": "^18.2.0"
  },
  "devDependencies": {
    "@types/react": "^18.2.0",
    "@types/react-dom": "^18.2.0",
    "@vitejs/plugin-react": "^4.2.0",
    "autoprefixer": "^10.4.16",
    "postcss": "^8.4.32",
    "tailwindcss": "^3.4.0",
    "typescript": "^5.3.0",
    "vite": "^5.0.0"
  }
}
`

var fullstackTsconfigTmpl = `{
  "compilerOptions": {
    "target": "ES2020",
    "useDefineForClassFields": true,
    "lib": ["ES2020", "DOM", "DOM.Iterable"],
    "module": "ESNext",
    "skipLibCheck": true,
    "moduleResolution": "bundler",
    "allowImportingTsExtensions": true,
    "resolveJsonModule": true,
    "isolatedModules": true,
    "noEmit": true,
    "jsx": "react-jsx",
    "strict": true,
    "noUnusedLocals": true,
    "noUnusedParameters": true,
    "noFallthroughCasesInSwitch": true
  },
  "include": ["src"],
  "references": [{ "path": "./tsconfig.node.json" }]
}
`

var fullstackTsconfigNodeTmpl = `{
  "compilerOptions": {
    "composite": true,
    "skipLibCheck": true,
    "module": "ESNext",
    "moduleResolution": "bundler",
    "allowSyntheticDefaultImports": true,
    "strict": true
  },
  "include": ["vite.config.ts"]
}
`

var fullstackViteConfigTmpl = `import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'

export default defineConfig({
  plugins: [react()],
  server: {
    port: 5173,
    proxy: {
      '/api': {
        target: 'http://localhost:8080',
        changeOrigin: true,
      },
    },
  },
})
`

var fullstackTailwindConfigTmpl = `/** @type {import('tailwindcss').Config} */
export default {
  content: [
    "./index.html",
    "./src/**/*.{js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {},
  },
  plugins: [],
}
`

var fullstackPostcssConfigTmpl = `export default {
  plugins: {
    tailwindcss: {},
    autoprefixer: {},
  },
}
`

var fullstackIndexHtmlTmpl = `<!doctype html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <title>{{.ProjectName}}</title>
  </head>
  <body>
    <div id="root"></div>
    <script type="module" src="/src/main.tsx"></script>
  </body>
</html>
`

var fullstackMainTsxTmpl = `import React from 'react'
import ReactDOM from 'react-dom/client'
import App from './App'
import './index.css'

ReactDOM.createRoot(document.getElementById('root')!).render(
  <React.StrictMode>
    <App />
  </React.StrictMode>,
)
`

var fullstackAppTsxTmpl = `import { useState, useEffect } from 'react'
import { api } from './lib/api'
import { Header } from './components/Header'

function App() {
  const [message, setMessage] = useState<string>('')
  const [loading, setLoading] = useState(true)

  useEffect(() => {
    const fetchMessage = async () => {
      try {
        const data = await api.getMessage()
        setMessage(data.message)
      } catch (error) {
        setMessage('Failed to connect to backend')
      } finally {
        setLoading(false)
      }
    }
    fetchMessage()
  }, [])

  return (
    <div className="min-h-screen bg-gradient-to-br from-slate-900 via-purple-900 to-slate-900">
      <Header />
      <main className="container mx-auto px-4 py-16">
        <div className="text-center">
          <h1 className="text-5xl font-bold text-white mb-8">
            Welcome to <span className="text-purple-400">{{.ProjectName}}</span>
          </h1>
          <div className="bg-white/10 backdrop-blur-lg rounded-xl p-8 max-w-md mx-auto">
            <p className="text-gray-300 mb-4">Message from Go Backend:</p>
            {loading ? (
              <div className="animate-pulse bg-purple-500/20 h-8 rounded"></div>
            ) : (
              <p className="text-2xl font-semibold text-purple-400">{message}</p>
            )}
          </div>
          <div className="mt-12 flex justify-center gap-4">
            <a
              href="http://localhost:8080/api/health"
              target="_blank"
              className="px-6 py-3 bg-purple-600 hover:bg-purple-700 text-white rounded-lg transition-colors"
            >
              Check API Health
            </a>
          </div>
        </div>
      </main>
    </div>
  )
}

export default App
`

var fullstackIndexCssTmpl = `@tailwind base;
@tailwind components;
@tailwind utilities;

body {
  font-family: Inter, system-ui, Avenir, Helvetica, Arial, sans-serif;
}
`

var fullstackApiTsTmpl = `const API_BASE = '/api'

export const api = {
  async getMessage(): Promise<{ message: string; project: string }> {
    const res = await fetch(` + "`${API_BASE}/message`" + `)
    if (!res.ok) throw new Error('Failed to fetch')
    return res.json()
  },

  async checkHealth(): Promise<{ status: string }> {
    const res = await fetch(` + "`${API_BASE}/health`" + `)
    if (!res.ok) throw new Error('Failed to fetch')
    return res.json()
  },
}
`

var fullstackHeaderTmpl = `export function Header() {
  return (
    <header className="border-b border-white/10">
      <nav className="container mx-auto px-4 py-4 flex justify-between items-center">
        <div className="text-xl font-bold text-white">{{.ProjectName}}</div>
        <div className="flex gap-6 text-gray-300">
          <a href="#" className="hover:text-purple-400 transition-colors">Home</a>
          <a href="#" className="hover:text-purple-400 transition-colors">About</a>
          <a href="#" className="hover:text-purple-400 transition-colors">Contact</a>
        </div>
      </nav>
    </header>
  )
}
`

var fullstackReadmeTmpl = `# {{.ProjectName}}

Fullstack application with Go backend and React frontend.

## Tech Stack

**Backend:**
- Go 1.21+
- Net/HTTP (standard library)

**Frontend:**
- React 18 + TypeScript
- Vite (dev server & build)
- Tailwind CSS
- Bun (package manager)

## Getting Started

### Prerequisites
- Go 1.21+
- Bun (https://bun.sh)

### Installation

` + "```bash" + `
# Install frontend dependencies
cd frontend
bun install
` + "```" + `

### Running

` + "```bash" + `
# Terminal 1: Start backend
make backend

# Terminal 2: Start frontend
make frontend
` + "```" + `

Or use the Makefile:
` + "```bash" + `
make dev  # Runs both in parallel
` + "```" + `

### URLs
- Frontend: http://localhost:5173
- Backend API: http://localhost:8080
`

var fullstackMakefileTmpl = `.PHONY: dev backend frontend install

# Run both servers
dev:
	@echo "Starting backend and frontend..."
	@make -j2 backend frontend

# Run Go backend
backend:
	cd backend && go run ./cmd/api

# Run React frontend
frontend:
	cd frontend && bun run dev

# Install frontend dependencies
install:
	cd frontend && bun install

# Build for production
build:
	cd backend && go build -o ../dist/api ./cmd/api
	cd frontend && bun run build
`

var fullstackGitignoreTmpl = `# Dependencies
node_modules/
vendor/

# Build
dist/
build/

# Environment
.env
.env.local

# IDE
.idea/
.vscode/
*.swp

# OS
.DS_Store
Thumbs.db

# Go
*.exe
*.exe~
*.dll
*.so
*.dylib

# Logs
*.log
`

// ============== Learn DSA Templates ==============

var learnDSAReadmeTmpl = `# Data Structures & Algorithms Practice

Practice implementing common data structures and algorithms in Go.

## Structure

` + "```" + `
datastructures/
├── stack/       # Stack implementation
├── queue/       # Queue implementation
├── linkedlist/  # Linked List implementation
└── tree/        # Binary Tree (coming soon)

algorithms/
├── sorting/     # Sorting algorithms
├── searching/   # Searching algorithms
└── recursion/   # Recursive problems
` + "```" + `

## How to Practice

1. Open any ` + "`*_test.go`" + ` file to see the expected behavior
2. Implement the functions in the corresponding ` + "`.go`" + ` file
3. Run tests: ` + "`go test ./...`" + `

## Quick Commands

` + "```bash" + `
# Run all tests
go test ./...

# Run tests with verbose output
go test -v ./...

# Run specific package tests
go test -v ./datastructures/stack/
go test -v ./algorithms/sorting/
` + "```" + `
`

var dsaStackTmpl = `package stack

// Stack represents a LIFO data structure
type Stack struct {
	items []int
}

// New creates a new empty stack
func New() *Stack {
	return &Stack{items: []int{}}
}

// Push adds an element to the top of the stack
// TODO: Implement this
func (s *Stack) Push(item int) {
	// Your code here
}

// Pop removes and returns the top element
// Returns -1 if stack is empty
// TODO: Implement this
func (s *Stack) Pop() int {
	// Your code here
	return -1
}

// Peek returns the top element without removing it
// Returns -1 if stack is empty
// TODO: Implement this
func (s *Stack) Peek() int {
	// Your code here
	return -1
}

// IsEmpty returns true if the stack has no elements
// TODO: Implement this
func (s *Stack) IsEmpty() bool {
	// Your code here
	return true
}

// Size returns the number of elements in the stack
// TODO: Implement this
func (s *Stack) Size() int {
	// Your code here
	return 0
}
`

var dsaStackTestTmpl = `package stack

import "testing"

func TestPush(t *testing.T) {
	s := New()
	s.Push(1)
	s.Push(2)
	s.Push(3)

	if s.Size() != 3 {
		t.Errorf("Size() = %d; want 3", s.Size())
	}
}

func TestPop(t *testing.T) {
	s := New()
	s.Push(1)
	s.Push(2)
	s.Push(3)

	if got := s.Pop(); got != 3 {
		t.Errorf("Pop() = %d; want 3", got)
	}
	if got := s.Pop(); got != 2 {
		t.Errorf("Pop() = %d; want 2", got)
	}
	if s.Size() != 1 {
		t.Errorf("Size() = %d; want 1", s.Size())
	}
}

func TestPopEmpty(t *testing.T) {
	s := New()
	if got := s.Pop(); got != -1 {
		t.Errorf("Pop() on empty stack = %d; want -1", got)
	}
}

func TestPeek(t *testing.T) {
	s := New()
	s.Push(42)
	
	if got := s.Peek(); got != 42 {
		t.Errorf("Peek() = %d; want 42", got)
	}
	// Peek shouldn't remove the element
	if s.Size() != 1 {
		t.Errorf("Size() after Peek() = %d; want 1", s.Size())
	}
}

func TestIsEmpty(t *testing.T) {
	s := New()
	if !s.IsEmpty() {
		t.Error("IsEmpty() = false; want true for new stack")
	}
	s.Push(1)
	if s.IsEmpty() {
		t.Error("IsEmpty() = true; want false after Push")
	}
}
`

var dsaQueueTmpl = `package queue

// Queue represents a FIFO data structure
type Queue struct {
	items []int
}

// New creates a new empty queue
func New() *Queue {
	return &Queue{items: []int{}}
}

// Enqueue adds an element to the back of the queue
// TODO: Implement this
func (q *Queue) Enqueue(item int) {
	// Your code here
}

// Dequeue removes and returns the front element
// Returns -1 if queue is empty
// TODO: Implement this
func (q *Queue) Dequeue() int {
	// Your code here
	return -1
}

// Front returns the front element without removing it
// Returns -1 if queue is empty
// TODO: Implement this
func (q *Queue) Front() int {
	// Your code here
	return -1
}

// IsEmpty returns true if the queue has no elements
// TODO: Implement this
func (q *Queue) IsEmpty() bool {
	// Your code here
	return true
}

// Size returns the number of elements in the queue
// TODO: Implement this
func (q *Queue) Size() int {
	// Your code here
	return 0
}
`

var dsaQueueTestTmpl = `package queue

import "testing"

func TestEnqueue(t *testing.T) {
	q := New()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	if q.Size() != 3 {
		t.Errorf("Size() = %d; want 3", q.Size())
	}
}

func TestDequeue(t *testing.T) {
	q := New()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	// FIFO: first in, first out
	if got := q.Dequeue(); got != 1 {
		t.Errorf("Dequeue() = %d; want 1", got)
	}
	if got := q.Dequeue(); got != 2 {
		t.Errorf("Dequeue() = %d; want 2", got)
	}
}

func TestFront(t *testing.T) {
	q := New()
	q.Enqueue(42)
	q.Enqueue(100)

	if got := q.Front(); got != 42 {
		t.Errorf("Front() = %d; want 42", got)
	}
	// Front shouldn't remove the element
	if q.Size() != 2 {
		t.Errorf("Size() after Front() = %d; want 2", q.Size())
	}
}
`

var dsaLinkedListTmpl = `package linkedlist

// Node represents a node in the linked list
type Node struct {
	Value int
	Next  *Node
}

// LinkedList represents a singly linked list
type LinkedList struct {
	Head *Node
	size int
}

// New creates a new empty linked list
func New() *LinkedList {
	return &LinkedList{}
}

// Append adds a new node at the end
// TODO: Implement this
func (l *LinkedList) Append(value int) {
	// Your code here
}

// Prepend adds a new node at the beginning
// TODO: Implement this
func (l *LinkedList) Prepend(value int) {
	// Your code here
}

// Delete removes the first node with the given value
// Returns true if deleted, false if not found
// TODO: Implement this
func (l *LinkedList) Delete(value int) bool {
	// Your code here
	return false
}

// Find returns true if value exists in the list
// TODO: Implement this
func (l *LinkedList) Find(value int) bool {
	// Your code here
	return false
}

// Size returns the number of nodes
// TODO: Implement this
func (l *LinkedList) Size() int {
	// Your code here
	return 0
}

// ToSlice converts the linked list to a slice (for testing)
func (l *LinkedList) ToSlice() []int {
	var result []int
	current := l.Head
	for current != nil {
		result = append(result, current.Value)
		current = current.Next
	}
	return result
}
`

var dsaLinkedListTestTmpl = `package linkedlist

import (
	"reflect"
	"testing"
)

func TestAppend(t *testing.T) {
	l := New()
	l.Append(1)
	l.Append(2)
	l.Append(3)

	got := l.ToSlice()
	want := []int{1, 2, 3}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("ToSlice() = %v; want %v", got, want)
	}
}

func TestPrepend(t *testing.T) {
	l := New()
	l.Prepend(1)
	l.Prepend(2)
	l.Prepend(3)

	got := l.ToSlice()
	want := []int{3, 2, 1}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("ToSlice() = %v; want %v", got, want)
	}
}

func TestDelete(t *testing.T) {
	l := New()
	l.Append(1)
	l.Append(2)
	l.Append(3)

	if !l.Delete(2) {
		t.Error("Delete(2) = false; want true")
	}
	got := l.ToSlice()
	want := []int{1, 3}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("ToSlice() after delete = %v; want %v", got, want)
	}
}

func TestFind(t *testing.T) {
	l := New()
	l.Append(1)
	l.Append(2)
	l.Append(3)

	if !l.Find(2) {
		t.Error("Find(2) = false; want true")
	}
	if l.Find(99) {
		t.Error("Find(99) = true; want false")
	}
}
`

var dsaSortingTmpl = `package sorting

// BubbleSort sorts a slice using bubble sort algorithm
// Time: O(n²), Space: O(1)
// TODO: Implement this
func BubbleSort(arr []int) []int {
	// Your code here
	return arr
}

// InsertionSort sorts a slice using insertion sort algorithm
// Time: O(n²), Space: O(1)
// TODO: Implement this
func InsertionSort(arr []int) []int {
	// Your code here
	return arr
}

// MergeSort sorts a slice using merge sort algorithm
// Time: O(n log n), Space: O(n)
// TODO: Implement this
func MergeSort(arr []int) []int {
	// Your code here
	return arr
}

// QuickSort sorts a slice using quick sort algorithm
// Time: O(n log n) average, Space: O(log n)
// TODO: Implement this
func QuickSort(arr []int) []int {
	// Your code here
	return arr
}
`

var dsaSortingTestTmpl = `package sorting

import (
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	tests := []struct {
		input []int
		want  []int
	}{
		{[]int{5, 2, 8, 1, 9}, []int{1, 2, 5, 8, 9}},
		{[]int{1, 2, 3}, []int{1, 2, 3}},
		{[]int{3, 2, 1}, []int{1, 2, 3}},
		{[]int{}, []int{}},
	}

	for _, tt := range tests {
		got := BubbleSort(tt.input)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("BubbleSort(%v) = %v; want %v", tt.input, got, tt.want)
		}
	}
}

func TestInsertionSort(t *testing.T) {
	got := InsertionSort([]int{5, 2, 8, 1, 9})
	want := []int{1, 2, 5, 8, 9}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("InsertionSort() = %v; want %v", got, want)
	}
}

func TestMergeSort(t *testing.T) {
	got := MergeSort([]int{5, 2, 8, 1, 9})
	want := []int{1, 2, 5, 8, 9}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("MergeSort() = %v; want %v", got, want)
	}
}

func TestQuickSort(t *testing.T) {
	got := QuickSort([]int{5, 2, 8, 1, 9})
	want := []int{1, 2, 5, 8, 9}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("QuickSort() = %v; want %v", got, want)
	}
}
`

var dsaSearchingTmpl = `package searching

// LinearSearch finds the index of target in arr
// Returns -1 if not found
// Time: O(n)
// TODO: Implement this
func LinearSearch(arr []int, target int) int {
	// Your code here
	return -1
}

// BinarySearch finds the index of target in a SORTED arr
// Returns -1 if not found
// Time: O(log n)
// TODO: Implement this
func BinarySearch(arr []int, target int) int {
	// Your code here
	return -1
}

// TwoSum finds two numbers that add up to target
// Returns their indices, or [-1, -1] if not found
// TODO: Implement this
func TwoSum(arr []int, target int) [2]int {
	// Your code here
	return [2]int{-1, -1}
}
`

var dsaSearchingTestTmpl = `package searching

import "testing"

func TestLinearSearch(t *testing.T) {
	arr := []int{1, 5, 3, 9, 2}
	
	if got := LinearSearch(arr, 9); got != 3 {
		t.Errorf("LinearSearch(arr, 9) = %d; want 3", got)
	}
	if got := LinearSearch(arr, 99); got != -1 {
		t.Errorf("LinearSearch(arr, 99) = %d; want -1", got)
	}
}

func TestBinarySearch(t *testing.T) {
	arr := []int{1, 2, 3, 5, 9} // sorted
	
	if got := BinarySearch(arr, 5); got != 3 {
		t.Errorf("BinarySearch(arr, 5) = %d; want 3", got)
	}
	if got := BinarySearch(arr, 1); got != 0 {
		t.Errorf("BinarySearch(arr, 1) = %d; want 0", got)
	}
	if got := BinarySearch(arr, 99); got != -1 {
		t.Errorf("BinarySearch(arr, 99) = %d; want -1", got)
	}
}

func TestTwoSum(t *testing.T) {
	arr := []int{2, 7, 11, 15}
	
	got := TwoSum(arr, 9)
	// 2 + 7 = 9, indices 0 and 1
	if got != [2]int{0, 1} && got != [2]int{1, 0} {
		t.Errorf("TwoSum(arr, 9) = %v; want [0, 1]", got)
	}
}
`

var dsaRecursionTmpl = `package recursion

// Factorial calculates n!
// TODO: Implement recursively
func Factorial(n int) int {
	// Your code here
	return 0
}

// Fibonacci returns the nth Fibonacci number
// TODO: Implement recursively
func Fibonacci(n int) int {
	// Your code here
	return 0
}

// ReverseString reverses a string recursively
// TODO: Implement recursively
func ReverseString(s string) string {
	// Your code here
	return ""
}

// SumArray calculates sum of array recursively
// TODO: Implement recursively
func SumArray(arr []int) int {
	// Your code here
	return 0
}
`

var dsaRecursionTestTmpl = `package recursion

import "testing"

func TestFactorial(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{0, 1},
		{1, 1},
		{5, 120},
		{10, 3628800},
	}

	for _, tt := range tests {
		if got := Factorial(tt.n); got != tt.want {
			t.Errorf("Factorial(%d) = %d; want %d", tt.n, got, tt.want)
		}
	}
}

func TestFibonacci(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{10, 55},
	}

	for _, tt := range tests {
		if got := Fibonacci(tt.n); got != tt.want {
			t.Errorf("Fibonacci(%d) = %d; want %d", tt.n, got, tt.want)
		}
	}
}

func TestReverseString(t *testing.T) {
	if got := ReverseString("hello"); got != "olleh" {
		t.Errorf("ReverseString(\"hello\") = %q; want \"olleh\"", got)
	}
}

func TestSumArray(t *testing.T) {
}
`

// ============== Skill Coding Templates ==============

// Challenge 30 Days Templates
var challenge30DaysReadmeTmpl = `# 30-Day Go Coding Challenge 🚀

Welcome to the 30-day Go coding challenge! Each day has an exercise to help you master Go.

## Structure

### Week 1: Basics
- Day 01: Hello World & Functions
- Day 02: Variables & Types
- Day 03: Conditionals

### Week 2: Data Structures
- Day 08: Recursion
- Day 09: Slices & Arrays

### Week 3: Web
- Day 15: HTTP Basics
- Day 16: JSON Handling

### Week 4: Concurrency
- Day 22: Goroutines
- Day 23: Channels

## How to Use

1. Navigate to the day folder
2. Read the main.go file for instructions
3. Implement the TODO functions
4. Run tests to verify: ` + "`go test -v`" + `

## Rules

- Complete one day at a time
- Don't skip ahead!
- Make all tests pass before moving on

Good luck! 💪
`

var challengeDay01Tmpl = `package main

import "fmt"

// TODO: Implement Greet function
// It should return "Hello, {name}!" where {name} is the input parameter
func Greet(name string) string {
	// YOUR CODE HERE
	return ""
}

// TODO: Implement Add function
// It should return the sum of two integers
func Add(a, b int) int {
	// YOUR CODE HERE
	return 0
}

func main() {
	fmt.Println(Greet("World"))
	fmt.Println("2 + 3 =", Add(2, 3))
}
`

var challengeDay01TestTmpl = `package main

import "testing"

func TestGreet(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{"World", "Hello, World!"},
		{"Go", "Hello, Go!"},
		{"", "Hello, !"},
	}
	for _, tt := range tests {
		if got := Greet(tt.name); got != tt.want {
			t.Errorf("Greet(%q) = %q, want %q", tt.name, got, tt.want)
		}
	}
}

func TestAdd(t *testing.T) {
	tests := []struct {
		a, b, want int
	}{
		{2, 3, 5},
		{0, 0, 0},
		{-1, 1, 0},
	}
	for _, tt := range tests {
		if got := Add(tt.a, tt.b); got != tt.want {
			t.Errorf("Add(%d, %d) = %d, want %d", tt.a, tt.b, got, tt.want)
		}
	}
}
`

var challengeDay02Tmpl = `package main

import "fmt"

// TODO: Implement Swap function
// It should swap two integers and return them in reversed order
func Swap(a, b int) (int, int) {
	// YOUR CODE HERE
	return 0, 0
}

// TODO: Implement IsEven function
// It should return true if the number is even
func IsEven(n int) bool {
	// YOUR CODE HERE
	return false
}

// TODO: Implement Max function
// It should return the larger of two integers
func Max(a, b int) int {
	// YOUR CODE HERE
	return 0
}

func main() {
	a, b := Swap(1, 2)
	fmt.Printf("Swap(1, 2) = %d, %d\n", a, b)
	fmt.Printf("IsEven(4) = %v\n", IsEven(4))
	fmt.Printf("Max(5, 3) = %d\n", Max(5, 3))
}
`

var challengeDay02TestTmpl = `package main

import "testing"

func TestSwap(t *testing.T) {
	a, b := Swap(1, 2)
	if a != 2 || b != 1 {
		t.Errorf("Swap(1, 2) = (%d, %d), want (2, 1)", a, b)
	}
}

func TestIsEven(t *testing.T) {
	tests := []struct {
		n    int
		want bool
	}{
		{0, true},
		{1, false},
		{2, true},
		{-2, true},
	}
	for _, tt := range tests {
		if got := IsEven(tt.n); got != tt.want {
			t.Errorf("IsEven(%d) = %v, want %v", tt.n, got, tt.want)
		}
	}
}

func TestMax(t *testing.T) {
	tests := []struct {
		a, b, want int
	}{
		{5, 3, 5},
		{1, 10, 10},
		{7, 7, 7},
	}
	for _, tt := range tests {
		if got := Max(tt.a, tt.b); got != tt.want {
			t.Errorf("Max(%d, %d) = %d, want %d", tt.a, tt.b, got, tt.want)
		}
	}
}
`

var challengeDay08Tmpl = `package main

import "fmt"

// TODO: Implement Factorial function using recursion
// factorial(0) = 1, factorial(n) = n * factorial(n-1)
func Factorial(n int) int {
	// YOUR CODE HERE
	return 0
}

// TODO: Implement Fibonacci function using recursion
// fib(0) = 0, fib(1) = 1, fib(n) = fib(n-1) + fib(n-2)
func Fibonacci(n int) int {
	// YOUR CODE HERE
	return 0
}

func main() {
	fmt.Printf("Factorial(5) = %d\n", Factorial(5))
	fmt.Printf("Fibonacci(10) = %d\n", Fibonacci(10))
}
`

var challengeDay08TestTmpl = `package main

import "testing"

func TestFactorial(t *testing.T) {
	tests := []struct {
		n, want int
	}{
		{0, 1},
		{1, 1},
		{5, 120},
		{10, 3628800},
	}
	for _, tt := range tests {
		if got := Factorial(tt.n); got != tt.want {
			t.Errorf("Factorial(%d) = %d, want %d", tt.n, got, tt.want)
		}
	}
}

func TestFibonacci(t *testing.T) {
	tests := []struct {
		n, want int
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{10, 55},
	}
	for _, tt := range tests {
		if got := Fibonacci(tt.n); got != tt.want {
			t.Errorf("Fibonacci(%d) = %d, want %d", tt.n, got, tt.want)
		}
	}
}
`

var challengeDay15Tmpl = `package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// TODO: Implement HealthHandler
// It should return JSON: {"status": "ok"}
func HealthHandler(w http.ResponseWriter, r *http.Request) {
	// YOUR CODE HERE
}

// TODO: Implement EchoHandler
// It should return JSON with the "message" query parameter
// e.g., /echo?message=hello returns {"echo": "hello"}
func EchoHandler(w http.ResponseWriter, r *http.Request) {
	// YOUR CODE HERE
}

func main() {
	http.HandleFunc("/health", HealthHandler)
	http.HandleFunc("/echo", EchoHandler)
	fmt.Println("Server running on :8080")
	http.ListenAndServe(":8080", nil)
}

// Helper for json encoding
func respondJSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
`

var challengeDay15TestTmpl = `package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/health", nil)
	w := httptest.NewRecorder()
	HealthHandler(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("got status %d, want %d", w.Code, http.StatusOK)
	}

	var resp map[string]string
	json.NewDecoder(w.Body).Decode(&resp)
	if resp["status"] != "ok" {
		t.Errorf("got status %q, want %q", resp["status"], "ok")
	}
}

func TestEchoHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/echo?message=hello", nil)
	w := httptest.NewRecorder()
	EchoHandler(w, req)

	var resp map[string]string
	json.NewDecoder(w.Body).Decode(&resp)
	if resp["echo"] != "hello" {
		t.Errorf("got echo %q, want %q", resp["echo"], "hello")
	}
}
`

var challengeDay22Tmpl = `package main

import (
	"fmt"
	"sync"
	"time"
)

// TODO: Implement ConcurrentSum
// It should calculate sum of array elements using goroutines
// Divide the array into chunks and sum each chunk concurrently
func ConcurrentSum(nums []int, workers int) int {
	// YOUR CODE HERE
	return 0
}

// TODO: Implement WorkerPool
// Create 'n' workers that process jobs from the jobs channel
// Each worker should square the number and send to results channel
func WorkerPool(jobs <-chan int, results chan<- int, n int) {
	// YOUR CODE HERE
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sum := ConcurrentSum(nums, 2)
	fmt.Printf("ConcurrentSum = %d\n", sum)

	// Worker pool example
	jobs := make(chan int, 5)
	results := make(chan int, 5)
	
	go WorkerPool(jobs, results, 3)
	
	for i := 1; i <= 5; i++ {
		jobs <- i
	}
	close(jobs)
	
	time.Sleep(100 * time.Millisecond)
	close(results)
	
	for r := range results {
		fmt.Printf("Result: %d\n", r)
	}
}

// Hint: You might need sync.WaitGroup
var _ = sync.WaitGroup{}
`

var challengeDay22TestTmpl = `package main

import (
	"sort"
	"testing"
)

func TestConcurrentSum(t *testing.T) {
	tests := []struct {
		nums    []int
		workers int
		want    int
	}{
		{[]int{1, 2, 3, 4, 5}, 2, 15},
		{[]int{10, 20, 30}, 3, 60},
		{[]int{}, 1, 0},
	}
	for _, tt := range tests {
		got := ConcurrentSum(tt.nums, tt.workers)
		if got != tt.want {
			t.Errorf("ConcurrentSum(%v, %d) = %d, want %d", tt.nums, tt.workers, got, tt.want)
		}
	}
}

func TestWorkerPool(t *testing.T) {
	jobs := make(chan int, 5)
	results := make(chan int, 5)
	
	go WorkerPool(jobs, results, 3)
	
	for i := 1; i <= 5; i++ {
		jobs <- i
	}
	close(jobs)
	
	var got []int
	for i := 0; i < 5; i++ {
		got = append(got, <-results)
	}
	
	sort.Ints(got)
	want := []int{1, 4, 9, 16, 25}
	
	for i, v := range got {
		if v != want[i] {
			t.Errorf("WorkerPool result mismatch at %d: got %d, want %d", i, v, want[i])
		}
	}
}
`

// Mini-Project Templates
var miniProjectReadmeTmpl = `# Mini Projects 🛠️

Choose a mini project to build and learn by doing!

## Available Projects

### 1. Todo CLI (todo-cli/)
A command-line todo list manager. Learn about:
- File I/O
- JSON handling
- CLI argument parsing

### 2. URL Shortener (url-shortener/)
A simple URL shortening service. Learn about:
- HTTP servers
- In-memory storage
- REST API design

## How to Start

1. Pick a project folder
2. Read the README.md inside
3. Look at main.go for the skeleton code
4. Run tests as you implement: ` + "`go test -v`" + `

Good luck! 🚀
`

var todoCliReadmeTmpl = `# Todo CLI

Build a command-line todo list manager!

## Features to Implement

1. **Add todo**: ` + "`todo add \"Buy groceries\"`" + `
2. **List todos**: ` + "`todo list`" + `
3. **Complete todo**: ` + "`todo done 1`" + `
4. **Delete todo**: ` + "`todo delete 1`" + `

## Requirements

- Store todos in a JSON file (todos.json)
- Each todo has: id, text, completed, createdAt
- Handle edge cases (invalid id, empty list, etc.)

## Example Usage

` + "```" + `
$ todo add "Learn Go"
Added: Learn Go (id: 1)

$ todo add "Build CLI"
Added: Build CLI (id: 2)

$ todo list
[ ] 1. Learn Go
[ ] 2. Build CLI

$ todo done 1
Completed: Learn Go

$ todo list
[x] 1. Learn Go
[ ] 2. Build CLI
` + "```" + `

## Testing

Run ` + "`go test -v`" + ` to check your implementation.
`

var todoCliMainTmpl = `package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

const dataFile = "todos.json"

type Todo struct {
	ID        int       ` + "`json:\"id\"`" + `
	Text      string    ` + "`json:\"text\"`" + `
	Completed bool      ` + "`json:\"completed\"`" + `
	CreatedAt time.Time ` + "`json:\"created_at\"`" + `
}

type TodoList struct {
	Todos  []Todo ` + "`json:\"todos\"`" + `
	NextID int    ` + "`json:\"next_id\"`" + `
}

// TODO: Implement Load function
// Load todos from the JSON file. Return empty TodoList if file doesn't exist.
func Load() (*TodoList, error) {
	// YOUR CODE HERE
	return &TodoList{NextID: 1}, nil
}

// TODO: Implement Save function
// Save the TodoList to the JSON file
func (tl *TodoList) Save() error {
	// YOUR CODE HERE
	return nil
}

// TODO: Implement Add function
// Add a new todo item and return its ID
func (tl *TodoList) Add(text string) int {
	// YOUR CODE HERE
	return 0
}

// TODO: Implement List function
// Return all todos
func (tl *TodoList) List() []Todo {
	// YOUR CODE HERE
	return nil
}

// TODO: Implement Complete function
// Mark a todo as completed by ID. Return error if not found.
func (tl *TodoList) Complete(id int) error {
	// YOUR CODE HERE
	return fmt.Errorf("todo not found: %d", id)
}

// TODO: Implement Delete function
// Delete a todo by ID. Return error if not found.
func (tl *TodoList) Delete(id int) error {
	// YOUR CODE HERE
	return fmt.Errorf("todo not found: %d", id)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: todo <command> [args]")
		fmt.Println("Commands: add, list, done, delete")
		return
	}

	tl, err := Load()
	if err != nil {
		fmt.Println("Error loading todos:", err)
		return
	}

	cmd := os.Args[1]
	switch cmd {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Usage: todo add <text>")
			return
		}
		id := tl.Add(os.Args[2])
		tl.Save()
		fmt.Printf("Added: %s (id: %d)\n", os.Args[2], id)
	case "list":
		todos := tl.List()
		if len(todos) == 0 {
			fmt.Println("No todos!")
			return
		}
		for _, t := range todos {
			status := "[ ]"
			if t.Completed {
				status = "[x]"
			}
			fmt.Printf("%s %d. %s\n", status, t.ID, t.Text)
		}
	case "done":
		if len(os.Args) < 3 {
			fmt.Println("Usage: todo done <id>")
			return
		}
		var id int
		fmt.Sscanf(os.Args[2], "%d", &id)
		if err := tl.Complete(id); err != nil {
			fmt.Println(err)
			return
		}
		tl.Save()
		fmt.Println("Marked as done!")
	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Usage: todo delete <id>")
			return
		}
		var id int
		fmt.Sscanf(os.Args[2], "%d", &id)
		if err := tl.Delete(id); err != nil {
			fmt.Println(err)
			return
		}
		tl.Save()
		fmt.Println("Deleted!")
	default:
		fmt.Println("Unknown command:", cmd)
	}
}

// Helper (you can use this)
var _ = json.Marshal
`

var todoCliTestTmpl = `package main

import (
	"os"
	"testing"
)

func TestTodoList(t *testing.T) {
	// Clean up test file
	os.Remove(dataFile)
	defer os.Remove(dataFile)

	tl := &TodoList{NextID: 1}

	// Test Add
	id1 := tl.Add("First todo")
	if id1 != 1 {
		t.Errorf("Add returned id %d, want 1", id1)
	}

	id2 := tl.Add("Second todo")
	if id2 != 2 {
		t.Errorf("Add returned id %d, want 2", id2)
	}

	// Test List
	todos := tl.List()
	if len(todos) != 2 {
		t.Errorf("List returned %d todos, want 2", len(todos))
	}

	// Test Complete
	err := tl.Complete(1)
	if err != nil {
		t.Errorf("Complete(1) returned error: %v", err)
	}

	todos = tl.List()
	if !todos[0].Completed {
		t.Error("Todo 1 should be completed")
	}

	// Test Complete non-existent
	err = tl.Complete(999)
	if err == nil {
		t.Error("Complete(999) should return error")
	}

	// Test Delete
	err = tl.Delete(1)
	if err != nil {
		t.Errorf("Delete(1) returned error: %v", err)
	}

	todos = tl.List()
	if len(todos) != 1 {
		t.Errorf("List returned %d todos after delete, want 1", len(todos))
	}

	// Test Save/Load
	tl.Save()
	loaded, err := Load()
	if err != nil {
		t.Errorf("Load returned error: %v", err)
	}
	if len(loaded.List()) != 1 {
		t.Errorf("Loaded list has %d todos, want 1", len(loaded.List()))
	}
}
`

var urlShortenerReadmeTmpl = `# URL Shortener

Build a simple URL shortening service!

## Features to Implement

1. **Shorten URL**: POST /shorten with {"url": "https://example.com"}
2. **Redirect**: GET /{code} redirects to original URL
3. **Stats**: GET /stats/{code} returns visit count

## Requirements

- Generate unique 6-character codes
- Store URLs in memory (map)
- Track visit counts
- Validate URLs

## Example Usage

` + "```" + `
# Shorten a URL
curl -X POST localhost:8080/shorten -d '{"url":"https://google.com"}'
{"code":"abc123","short_url":"http://localhost:8080/abc123"}

# Access shortened URL (redirects)
curl -L localhost:8080/abc123

# Get stats
curl localhost:8080/stats/abc123
{"code":"abc123","url":"https://google.com","visits":1}
` + "```" + `

## Testing

Run ` + "`go test -v`" + ` to check your implementation.
`

var urlShortenerMainTmpl = `package main

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"sync"
)

type URLStore struct {
	mu    sync.RWMutex
	urls  map[string]string // code -> original URL
	visits map[string]int   // code -> visit count
}

func NewURLStore() *URLStore {
	return &URLStore{
		urls:   make(map[string]string),
		visits: make(map[string]int),
	}
}

// TODO: Implement GenerateCode function
// Generate a random 6-character hex code
func GenerateCode() string {
	// YOUR CODE HERE (hint: use crypto/rand)
	return ""
}

// TODO: Implement Shorten method
// Store the URL and return the generated code
func (s *URLStore) Shorten(url string) (string, error) {
	// YOUR CODE HERE
	// 1. Validate URL (must start with http:// or https://)
	// 2. Generate code
	// 3. Store URL
	// 4. Return code
	return "", nil
}

// TODO: Implement Get method
// Return the original URL for a code, increment visits
func (s *URLStore) Get(code string) (string, bool) {
	// YOUR CODE HERE
	return "", false
}

// TODO: Implement Stats method
// Return URL and visit count for a code
func (s *URLStore) Stats(code string) (url string, visits int, found bool) {
	// YOUR CODE HERE
	return "", 0, false
}

var store = NewURLStore()

func shortenHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		URL string ` + "`json:\"url\"`" + `
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	code, err := store.Shorten(req.URL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp := map[string]string{
		"code":      code,
		"short_url": fmt.Sprintf("http://localhost:8080/%s", code),
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func redirectHandler(w http.ResponseWriter, r *http.Request) {
	code := strings.TrimPrefix(r.URL.Path, "/")
	if code == "" || code == "shorten" || strings.HasPrefix(code, "stats/") {
		http.NotFound(w, r)
		return
	}

	url, found := store.Get(code)
	if !found {
		http.NotFound(w, r)
		return
	}

	http.Redirect(w, r, url, http.StatusFound)
}

func statsHandler(w http.ResponseWriter, r *http.Request) {
	code := strings.TrimPrefix(r.URL.Path, "/stats/")
	url, visits, found := store.Stats(code)
	if !found {
		http.NotFound(w, r)
		return
	}

	resp := map[string]interface{}{
		"code":   code,
		"url":    url,
		"visits": visits,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func main() {
	http.HandleFunc("/shorten", shortenHandler)
	http.HandleFunc("/stats/", statsHandler)
	http.HandleFunc("/", redirectHandler)

	fmt.Println("URL Shortener running on :8080")
	http.ListenAndServe(":8080", nil)
}

// Helpers you can use
var _ = rand.Read
var _ = hex.EncodeToString
`

var urlShortenerTestTmpl = `package main

import (
	"testing"
)

func TestGenerateCode(t *testing.T) {
	code := GenerateCode()
	if len(code) != 6 {
		t.Errorf("GenerateCode() returned %q (len=%d), want len=6", code, len(code))
	}

	code2 := GenerateCode()
	if code == code2 {
		t.Error("GenerateCode() returned same code twice")
	}
}

func TestURLStore(t *testing.T) {
	store := NewURLStore()

	// Test Shorten with valid URL
	code, err := store.Shorten("https://example.com")
	if err != nil {
		t.Errorf("Shorten returned error: %v", err)
	}
	if len(code) != 6 {
		t.Errorf("Shorten returned code %q, want 6 chars", code)
	}

	// Test Shorten with invalid URL
	_, err = store.Shorten("not-a-url")
	if err == nil {
		t.Error("Shorten should reject invalid URLs")
	}

	// Test Get
	url, found := store.Get(code)
	if !found {
		t.Error("Get should find the shortened URL")
	}
	if url != "https://example.com" {
		t.Errorf("Get returned %q, want %q", url, "https://example.com")
	}

	// Test Get increments visits
	store.Get(code)
	_, visits, _ := store.Stats(code)
	if visits != 2 {
		t.Errorf("Stats returned visits=%d, want 2", visits)
	}

	// Test Get with unknown code
	_, found = store.Get("unknown")
	if found {
		t.Error("Get should not find unknown code")
	}
}
`

// Refactoring Exercise Templates
var refactoringReadmeTmpl = `# Refactoring Exercises 🔧

Practice improving code quality through refactoring!

## How to Use

1. Look at ` + "`before.go`" + ` in each exercise folder
2. Identify the code smells
3. Read ` + "`hints.md`" + ` for guidance
4. Create ` + "`after.go`" + ` with your refactored version
5. Make sure the functionality stays the same!

## Exercises

### 01 - Long Function
A function that does too many things. Break it down!

### 02 - Magic Numbers
Numbers without context. Give them names!

### 03 - Poor Naming
Variables and functions with unclear names. Make them descriptive!

## Tips

- Make small, incremental changes
- Run tests after each change
- Focus on readability
- Follow Go conventions

Happy refactoring! 🧹
`

var refactoring01BeforeTmpl = `package main

import "fmt"

// This function does too many things - refactor it!
func processOrder(items []string, quantities []int, prices []float64, customerName string, isVIP bool) {
	// Calculate total
	total := 0.0
	for i := 0; i < len(items); i++ {
		total += float64(quantities[i]) * prices[i]
	}

	// Apply discount
	if isVIP {
		total = total * 0.9
	}
	if total > 100 {
		total = total * 0.95
	}

	// Calculate tax
	tax := total * 0.1
	finalTotal := total + tax

	// Print receipt
	fmt.Println("====== RECEIPT ======")
	fmt.Printf("Customer: %s\n", customerName)
	if isVIP {
		fmt.Println("Status: VIP")
	}
	fmt.Println("---------------------")
	for i := 0; i < len(items); i++ {
		fmt.Printf("%s x%d @ $%.2f = $%.2f\n", items[i], quantities[i], prices[i], float64(quantities[i])*prices[i])
	}
	fmt.Println("---------------------")
	fmt.Printf("Subtotal: $%.2f\n", total)
	fmt.Printf("Tax (10%%): $%.2f\n", tax)
	fmt.Printf("Total: $%.2f\n", finalTotal)
	fmt.Println("=====================")

	// Save to database (simulated)
	fmt.Println("[DB] Saving order...")
	fmt.Printf("[DB] Customer: %s, Total: $%.2f\n", customerName, finalTotal)

	// Send email (simulated)
	fmt.Println("[EMAIL] Sending confirmation...")
	fmt.Printf("[EMAIL] To: %s, Subject: Order Confirmation\n", customerName)
}

func main() {
	items := []string{"Apple", "Banana", "Orange"}
	quantities := []int{3, 2, 5}
	prices := []float64{1.00, 0.50, 0.75}
	processOrder(items, quantities, prices, "John Doe", true)
}
`

var refactoring01HintsTmpl = `# Hints for Exercise 01: Long Function

## Code Smells
1. ` + "`processOrder`" + ` does 5 different things
2. Hard to test individual pieces
3. Hard to reuse parts of the logic

## Suggested Refactoring

### Step 1: Extract Calculation Functions
` + "```go" + `
func calculateSubtotal(quantities []int, prices []float64) float64
func applyDiscounts(total float64, isVIP bool) float64
func calculateTax(amount float64) float64
` + "```" + `

### Step 2: Create an Order struct
` + "```go" + `
type Order struct {
    Items      []OrderItem
    Customer   string
    IsVIP      bool
}

type OrderItem struct {
    Name     string
    Quantity int
    Price    float64
}
` + "```" + `

### Step 3: Extract Printing, DB, and Email
` + "```go" + `
func (o *Order) PrintReceipt()
func (o *Order) SaveToDatabase()
func (o *Order) SendConfirmation()
` + "```" + `

### Step 4: Create Clean Main Function
` + "```go" + `
func ProcessOrder(order Order) {
    subtotal := order.CalculateSubtotal()
    discounted := applyDiscounts(subtotal, order.IsVIP)
    tax := calculateTax(discounted)
    
    order.PrintReceipt(discounted, tax)
    order.SaveToDatabase()
    order.SendConfirmation()
}
` + "```" + `

## Benefits After Refactoring
- Each function has a single responsibility
- Easy to test individual functions
- Easy to modify one part without affecting others
- More readable and maintainable
`

var refactoring02BeforeTmpl = `package main

import "fmt"

// This code has magic numbers - give them names!
func calculateShipping(weight float64, distance float64, priority int) float64 {
	base := weight * 0.5

	if distance > 100 {
		base *= 1.5
	} else if distance > 50 {
		base *= 1.25
	}

	switch priority {
	case 1:
		base *= 3.0
	case 2:
		base *= 2.0
	case 3:
		base *= 1.5
	}

	if base < 5.0 {
		base = 5.0
	}
	if base > 500.0 {
		base = 500.0
	}

	return base
}

func calculatePassword(length int) bool {
	if length < 8 {
		return false
	}
	if length > 128 {
		return false
	}
	return true
}

func applyDiscount(price float64, code string) float64 {
	switch code {
	case "SAVE10":
		return price * 0.9
	case "SAVE20":
		return price * 0.8
	case "HALF":
		return price * 0.5
	default:
		return price
	}
}

func main() {
	fmt.Printf("Shipping: $%.2f\n", calculateShipping(10, 75, 2))
	fmt.Printf("Password valid: %v\n", calculatePassword(12))
	fmt.Printf("After discount: $%.2f\n", applyDiscount(100, "SAVE20"))
}
`

var refactoring02HintsTmpl = `# Hints for Exercise 02: Magic Numbers

## Code Smells
1. Numbers like 0.5, 1.5, 100, 50, 8, 128 have no context
2. Priority 1, 2, 3 are unclear
3. Discount codes are hardcoded

## Suggested Refactoring

### Step 1: Define Constants for Shipping
` + "```go" + `
const (
    BaseRatePerKg     = 0.5
    LongDistanceThreshold  = 100.0
    MediumDistanceThreshold = 50.0
    LongDistanceMultiplier  = 1.5
    MediumDistanceMultiplier = 1.25
    MinShippingCost = 5.0
    MaxShippingCost = 500.0
)
` + "```" + `

### Step 2: Use Named Constants for Priority
` + "```go" + `
const (
    PriorityExpress  = 1
    PriorityFast     = 2
    PriorityStandard = 3
)

var priorityMultipliers = map[int]float64{
    PriorityExpress:  3.0,
    PriorityFast:     2.0,
    PriorityStandard: 1.5,
}
` + "```" + `

### Step 3: Define Password Constraints
` + "```go" + `
const (
    MinPasswordLength = 8
    MaxPasswordLength = 128
)
` + "```" + `

### Step 4: Use a Discounts Map
` + "```go" + `
var discounts = map[string]float64{
    "SAVE10": 0.10,
    "SAVE20": 0.20,
    "HALF":   0.50,
}
` + "```" + `

## Benefits
- Self-documenting code
- Easy to change values in one place
- Prevents typos with repeated numbers
- More readable conditions
`

var refactoring03BeforeTmpl = `package main

import "fmt"

// This code has poor naming - make it descriptive!
func calc(a []int) int {
	t := 0
	for _, v := range a {
		t += v
	}
	return t
}

func f(s string) string {
	r := ""
	for i := len(s) - 1; i >= 0; i-- {
		r += string(s[i])
	}
	return r
}

func proc(d []map[string]interface{}) []string {
	res := []string{}
	for _, item := range d {
		n, ok := item["n"].(string)
		if !ok {
			continue
		}
		a, ok := item["a"].(int)
		if !ok || a < 18 {
			continue
		}
		res = append(res, n)
	}
	return res
}

type U struct {
	N string
	E string
	P string
}

func (u *U) V() bool {
	if len(u.N) < 2 {
		return false
	}
	if len(u.E) < 5 {
		return false
	}
	if len(u.P) < 8 {
		return false
	}
	return true
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println("Result:", calc(nums))
	fmt.Println("Reversed:", f("hello"))

	data := []map[string]interface{}{
		{"n": "Alice", "a": 25},
		{"n": "Bob", "a": 17},
		{"n": "Charlie", "a": 30},
	}
	fmt.Println("Adults:", proc(data))

	user := &U{N: "Al", E: "a@b.co", P: "password123"}
	fmt.Println("Valid:", user.V())
}
`

var refactoring03HintsTmpl = `# Hints for Exercise 03: Poor Naming

## Code Smells
1. Single-letter variable names (a, t, v, r, s, d, n, etc.)
2. Abbreviated function names (calc, f, proc, V)
3. Unclear struct name (U) and fields (N, E, P)

## Suggested Refactoring

### Step 1: Rename calc function
Before: ` + "`func calc(a []int) int`" + `
After: ` + "`func SumIntegers(numbers []int) int`" + `

Variables:
- ` + "`t`" + ` → ` + "`total`" + `
- ` + "`v`" + ` → ` + "`num`" + ` or ` + "`number`" + `

### Step 2: Rename f function
Before: ` + "`func f(s string) string`" + `
After: ` + "`func ReverseString(input string) string`" + `

Variables:
- ` + "`r`" + ` → ` + "`reversed`" + ` or ` + "`result`" + `
- ` + "`i`" + ` is okay for loop index

### Step 3: Rename proc function
Before: ` + "`func proc(d []map[string]interface{}) []string`" + `
After: ` + "`func GetAdultNames(people []map[string]interface{}) []string`" + `

Variables:
- ` + "`d`" + ` → ` + "`people`" + `
- ` + "`n`" + ` → ` + "`name`" + `
- ` + "`a`" + ` → ` + "`age`" + `
- ` + "`res`" + ` → ` + "`adultNames`" + `

### Step 4: Rename struct U
Before: ` + "`type U struct`" + `
After: ` + "`type User struct`" + `

Fields:
- ` + "`N`" + ` → ` + "`Name`" + `
- ` + "`E`" + ` → ` + "`Email`" + `
- ` + "`P`" + ` → ` + "`Password`" + `
- ` + "`V()`" + ` → ` + "`IsValid()`" + `

## Naming Guidelines
1. Use descriptive names that reveal intent
2. Avoid abbreviations unless universally known
3. Use verbs for functions (Calculate, Get, Process)
4. Use nouns for structs (User, Order, Product)
5. Boolean methods should start with Is/Has/Can
`

// Code Review Exercise Templates
var codeReviewReadmeTmpl = `# Code Review Exercises 🐛

Practice finding bugs in code!

## How to Play

1. Each folder has a ` + "`buggy.go`" + ` file with a bug
2. There's also a ` + "`buggy_test.go`" + ` that FAILS
3. Find and fix the bug!
4. Your goal: make all tests pass

## Exercises

### 01 - Off By One
Classic array indexing bug.

### 02 - Nil Pointer
Accessing nil references.

### 03 - Race Condition
Concurrent access issues.

## Tips

- Run tests first: ` + "`go test -v`" + `
- Read error messages carefully
- For race conditions: ` + "`go test -race`" + `

Good luck, bug hunter! 🔍
`

var codeReview01BuggyTmpl = `package main

// BUG: There's an off-by-one error in this function
// Can you find and fix it?

// GetLastN returns the last n elements of a slice
func GetLastN(slice []int, n int) []int {
	if n <= 0 || len(slice) == 0 {
		return []int{}
	}
	if n >= len(slice) {
		return slice
	}
	// Bug is somewhere here...
	start := len(slice) - n - 1
	return slice[start:]
}

// BUG: There's another off-by-one error here
// FindIndex returns the index of target, or -1 if not found
func FindIndex(slice []int, target int) int {
	for i := 1; i <= len(slice); i++ {
		if slice[i] == target {
			return i
		}
	}
	return -1
}
`

var codeReview01TestTmpl = `package main

import (
	"reflect"
	"testing"
)

func TestGetLastN(t *testing.T) {
	tests := []struct {
		slice []int
		n     int
		want  []int
	}{
		{[]int{1, 2, 3, 4, 5}, 3, []int{3, 4, 5}},
		{[]int{1, 2, 3}, 5, []int{1, 2, 3}},
		{[]int{1, 2, 3}, 0, []int{}},
		{[]int{}, 3, []int{}},
		{[]int{1, 2, 3, 4, 5}, 1, []int{5}},
	}

	for _, tt := range tests {
		got := GetLastN(tt.slice, tt.n)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("GetLastN(%v, %d) = %v, want %v", tt.slice, tt.n, got, tt.want)
		}
	}
}

func TestFindIndex(t *testing.T) {
	slice := []int{10, 20, 30, 40, 50}

	tests := []struct {
		target int
		want   int
	}{
		{10, 0},
		{30, 2},
		{50, 4},
		{99, -1},
	}

	for _, tt := range tests {
		got := FindIndex(slice, tt.target)
		if got != tt.want {
			t.Errorf("FindIndex(slice, %d) = %d, want %d", tt.target, got, tt.want)
		}
	}
}
`

var codeReview02BuggyTmpl = `package main

import "fmt"

// BUG: This code has nil pointer issues
// Can you find and fix them?

type Person struct {
	Name    string
	Address *Address
}

type Address struct {
	Street string
	City   string
}

// GetCity returns the city of a person
func GetCity(p *Person) string {
	// Bug: what if p is nil? what if p.Address is nil?
	return p.Address.City
}

// FormatAddress returns formatted address
func FormatAddress(p *Person) string {
	return fmt.Sprintf("%s, %s", p.Address.Street, p.Address.City)
}

type Config struct {
	Database *DatabaseConfig
}

type DatabaseConfig struct {
	Host string
	Port int
}

// GetDatabaseHost returns the database host from config
func GetDatabaseHost(cfg *Config) string {
	// Bug: nil checks missing
	return cfg.Database.Host
}
`

var codeReview02TestTmpl = `package main

import "testing"

func TestGetCity(t *testing.T) {
	// Test with valid person
	p := &Person{
		Name:    "Alice",
		Address: &Address{City: "NYC", Street: "123 Main St"},
	}
	if got := GetCity(p); got != "NYC" {
		t.Errorf("GetCity(valid) = %q, want %q", got, "NYC")
	}

	// Test with nil person - should not panic
	defer func() {
		if r := recover(); r != nil {
			t.Error("GetCity(nil) panicked, should return empty string")
		}
	}()
	if got := GetCity(nil); got != "" {
		t.Errorf("GetCity(nil) = %q, want empty", got)
	}

	// Test with nil address - should not panic
	p2 := &Person{Name: "Bob", Address: nil}
	if got := GetCity(p2); got != "" {
		t.Errorf("GetCity(nil address) = %q, want empty", got)
	}
}

func TestFormatAddress(t *testing.T) {
	p := &Person{
		Name:    "Alice",
		Address: &Address{Street: "123 Main St", City: "NYC"},
	}
	want := "123 Main St, NYC"
	if got := FormatAddress(p); got != want {
		t.Errorf("FormatAddress(valid) = %q, want %q", got, want)
	}

	// Should handle nil gracefully
	defer func() {
		if r := recover(); r != nil {
			t.Error("FormatAddress(nil) panicked")
		}
	}()
	FormatAddress(nil)
	FormatAddress(&Person{Name: "Bob"})
}

func TestGetDatabaseHost(t *testing.T) {
	cfg := &Config{
		Database: &DatabaseConfig{Host: "localhost", Port: 5432},
	}
	if got := GetDatabaseHost(cfg); got != "localhost" {
		t.Errorf("GetDatabaseHost(valid) = %q, want %q", got, "localhost")
	}

	// Should handle nil gracefully
	defer func() {
		if r := recover(); r != nil {
			t.Error("GetDatabaseHost panicked with nil")
		}
	}()
	GetDatabaseHost(nil)
	GetDatabaseHost(&Config{})
}
`

var codeReview03BuggyTmpl = `package main

import (
	"sync"
)

// BUG: This code has race conditions
// Run with: go test -race
// Can you find and fix them?

type Counter struct {
	value int
}

// Increment adds 1 to the counter
// BUG: Not safe for concurrent access
func (c *Counter) Increment() {
	c.value++
}

// Value returns the current counter value
func (c *Counter) Value() int {
	return c.value
}

// BUG: This map has race conditions
var cache = make(map[string]string)

func GetFromCache(key string) (string, bool) {
	val, ok := cache[key]
	return val, ok
}

func SetCache(key, value string) {
	cache[key] = value
}

// BUG: This slice append has race conditions
type Logger struct {
	logs []string
}

func (l *Logger) Log(msg string) {
	l.logs = append(l.logs, msg)
}

func (l *Logger) GetLogs() []string {
	return l.logs
}

// Hint: You'll need sync.Mutex or sync.RWMutex
var _ = sync.Mutex{}
`

var codeReview03TestTmpl = `package main

import (
	"sync"
	"testing"
)

func TestCounterConcurrent(t *testing.T) {
	c := &Counter{}
	var wg sync.WaitGroup
	
	// 100 goroutines each incrementing 100 times
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 100; j++ {
				c.Increment()
			}
		}()
	}
	
	wg.Wait()
	
	if c.Value() != 10000 {
		t.Errorf("Counter = %d, want 10000", c.Value())
	}
}

func TestCacheConcurrent(t *testing.T) {
	var wg sync.WaitGroup
	
	// Concurrent writes
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			key := string(rune('a' + n%26))
			SetCache(key, "value")
			GetFromCache(key)
		}(i)
	}
	
	wg.Wait()
	// If we get here without panic, the test passes
}

func TestLoggerConcurrent(t *testing.T) {
	l := &Logger{}
	var wg sync.WaitGroup
	
	// 100 goroutines each logging 10 times
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 10; j++ {
				l.Log("message")
			}
		}()
	}
	
	wg.Wait()
	
	if len(l.GetLogs()) != 1000 {
		t.Errorf("Logger has %d logs, want 1000", len(l.GetLogs()))
	}
}
`

// ============== Learn Generics Templates ==============

var learnGenericsReadmeTmpl = `# Learn Go Generics

Master Go generics with hands-on exercises!

## Sections

1. **basics/** - Generic functions and type parameters
2. **constraints/** - Type constraints and interfaces
3. **practical/** - Real-world generic patterns

## How to Use

` + "```bash" + `
cd basics && go test -v
` + "```" + `
`

var learnGenericsBasicsTmpl = `package main

// TODO: Implement a generic Min function
// Returns the smaller of two values
func Min[T any](a, b T) T {
	// YOUR CODE HERE (hint: need comparable/ordered constraint)
	return a
}

// TODO: Implement a generic Map function
// Applies fn to each element and returns new slice
func Map[T, U any](slice []T, fn func(T) U) []U {
	// YOUR CODE HERE
	return nil
}

// TODO: Implement a generic Filter function
func Filter[T any](slice []T, fn func(T) bool) []T {
	// YOUR CODE HERE
	return nil
}
`

var learnGenericsBasicsTestTmpl = `package main

import "testing"

func TestMin(t *testing.T) {
	if Min(3, 5) != 3 { t.Error("Min(3,5) should be 3") }
	if Min(5.5, 2.2) != 2.2 { t.Error("Min(5.5,2.2) should be 2.2") }
}

func TestMap(t *testing.T) {
	doubled := Map([]int{1, 2, 3}, func(x int) int { return x * 2 })
	if len(doubled) != 3 || doubled[0] != 2 { t.Error("Map failed") }
}

func TestFilter(t *testing.T) {
	evens := Filter([]int{1, 2, 3, 4}, func(x int) bool { return x%2 == 0 })
	if len(evens) != 2 { t.Error("Filter failed") }
}
`

var learnGenericsConstraintsTmpl = `package main

import "golang.org/x/exp/constraints"

// Number constraint for numeric types
type Number interface {
	constraints.Integer | constraints.Float
}

// TODO: Implement Sum using Number constraint
func Sum[T Number](nums []T) T {
	// YOUR CODE HERE
	var zero T
	return zero
}

// TODO: Implement a generic Stack with any type
type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) Push(item T) {
	// YOUR CODE HERE
}

func (s *Stack[T]) Pop() (T, bool) {
	// YOUR CODE HERE
	var zero T
	return zero, false
}
`

var learnGenericsConstraintsTestTmpl = `package main

import "testing"

func TestSum(t *testing.T) {
	if Sum([]int{1, 2, 3}) != 6 { t.Error("Sum ints failed") }
	if Sum([]float64{1.5, 2.5}) != 4.0 { t.Error("Sum floats failed") }
}

func TestStack(t *testing.T) {
	s := &Stack[string]{}
	s.Push("hello")
	s.Push("world")
	if v, ok := s.Pop(); !ok || v != "world" { t.Error("Pop failed") }
}
`

var learnGenericsPracticalTmpl = `package main

// TODO: Implement a generic Result type for error handling
type Result[T any] struct {
	Value T
	Err   error
}

func Ok[T any](value T) Result[T] {
	return Result[T]{Value: value}
}

func Err[T any](err error) Result[T] {
	return Result[T]{Err: err}
}

// TODO: Implement generic cache
type Cache[K comparable, V any] struct {
	data map[K]V
}

func NewCache[K comparable, V any]() *Cache[K, V] {
	return &Cache[K, V]{data: make(map[K]V)}
}

func (c *Cache[K, V]) Set(key K, value V) {
	// YOUR CODE HERE
}

func (c *Cache[K, V]) Get(key K) (V, bool) {
	// YOUR CODE HERE
	var zero V
	return zero, false
}
`

var learnGenericsPracticalTestTmpl = `package main

import (
	"errors"
	"testing"
)

func TestResult(t *testing.T) {
	ok := Ok(42)
	if ok.Err != nil || ok.Value != 42 { t.Error("Ok failed") }
	
	err := Err[int](errors.New("fail"))
	if err.Err == nil { t.Error("Err failed") }
}

func TestCache(t *testing.T) {
	c := NewCache[string, int]()
	c.Set("age", 30)
	if v, ok := c.Get("age"); !ok || v != 30 { t.Error("Cache failed") }
}
`

// ============== Learn Context Templates ==============

var learnContextReadmeTmpl = `# Learn context.Context

Master Go's context package for cancellation, timeouts, and value passing.

## Sections
1. **cancellation/** - Cancel long-running operations
2. **timeout/** - Set deadlines for operations
3. **values/** - Pass request-scoped values
`

var learnContextCancellationTmpl = `package main

import (
	"context"
	"fmt"
	"time"
)

// TODO: Implement a function that respects context cancellation
func DoWork(ctx context.Context) error {
	for i := 0; i < 10; i++ {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			fmt.Printf("Working... %d\n", i)
			time.Sleep(100 * time.Millisecond)
		}
	}
	return nil
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		time.Sleep(300 * time.Millisecond)
		cancel()
	}()
	DoWork(ctx)
}
`

var learnContextCancellationTestTmpl = `package main

import (
	"context"
	"testing"
	"time"
)

func TestDoWork_Cancellation(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		time.Sleep(50 * time.Millisecond)
		cancel()
	}()
	err := DoWork(ctx)
	if err != context.Canceled { t.Error("Should return context.Canceled") }
}
`

var learnContextTimeoutTmpl = `package main

import (
	"context"
	"time"
)

// TODO: Implement function with timeout
func FetchWithTimeout(ctx context.Context, duration time.Duration) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, duration)
	defer cancel()
	
	select {
	case <-time.After(500 * time.Millisecond):
		return "data", nil
	case <-ctx.Done():
		return "", ctx.Err()
	}
}
`

var learnContextTimeoutTestTmpl = `package main

import (
	"context"
	"testing"
	"time"
)

func TestFetchWithTimeout(t *testing.T) {
	_, err := FetchWithTimeout(context.Background(), 100*time.Millisecond)
	if err != context.DeadlineExceeded { t.Error("Should timeout") }
	
	data, err := FetchWithTimeout(context.Background(), 1*time.Second)
	if err != nil || data != "data" { t.Error("Should succeed") }
}
`

var learnContextValuesTmpl = `package main

import "context"

type contextKey string

const userIDKey contextKey = "userID"

// TODO: Implement context value functions
func WithUserID(ctx context.Context, userID string) context.Context {
	return context.WithValue(ctx, userIDKey, userID)
}

func GetUserID(ctx context.Context) (string, bool) {
	userID, ok := ctx.Value(userIDKey).(string)
	return userID, ok
}
`

var learnContextValuesTestTmpl = `package main

import (
	"context"
	"testing"
)

func TestContextValues(t *testing.T) {
	ctx := WithUserID(context.Background(), "user123")
	if id, ok := GetUserID(ctx); !ok || id != "user123" { t.Error("Value failed") }
	
	if _, ok := GetUserID(context.Background()); ok { t.Error("Should not find value") }
}
`

// ============== Learn HTTP Templates ==============

var learnHTTPReadmeTmpl = `# Learn HTTP in Go

Learn HTTP client, server, and middleware patterns.

## Sections
1. **client/** - Making HTTP requests
2. **server/** - Building HTTP servers
3. **middleware/** - Request/response middleware
`

var learnHTTPClientTmpl = `package main

import (
	"encoding/json"
	"io"
	"net/http"
)

// TODO: Implement GET request
func GetJSON(url string, target interface{}) error {
	resp, err := http.Get(url)
	if err != nil { return err }
	defer resp.Body.Close()
	return json.NewDecoder(resp.Body).Decode(target)
}

// TODO: Implement POST request with JSON body
func PostJSON(url string, body interface{}, target interface{}) error {
	// YOUR CODE HERE
	return nil
}

// Helper
var _ = io.ReadAll
`

var learnHTTPClientTestTmpl = `package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetJSON(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(` + "`" + `{"name":"test"}` + "`" + `))
	}))
	defer server.Close()
	
	var result struct{ Name string }
	if err := GetJSON(server.URL, &result); err != nil || result.Name != "test" {
		t.Error("GetJSON failed")
	}
}
`

var learnHTTPServerTmpl = `package main

import (
	"encoding/json"
	"net/http"
)

// TODO: Implement JSON response helper
func respondJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

// TODO: Implement a handler that returns user by ID
func UserHandler(w http.ResponseWriter, r *http.Request) {
	// YOUR CODE HERE
	respondJSON(w, 200, map[string]string{"id": "1", "name": "John"})
}
`

var learnHTTPServerTestTmpl = `package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUserHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/user/1", nil)
	w := httptest.NewRecorder()
	UserHandler(w, req)
	if w.Code != 200 { t.Errorf("got %d, want 200", w.Code) }
}
`

var learnHTTPMiddlewareTmpl = `package main

import (
	"log"
	"net/http"
	"time"
)

// TODO: Implement logging middleware
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("%s %s %v", r.Method, r.URL.Path, time.Since(start))
	})
}

// TODO: Implement auth middleware
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// YOUR CODE HERE - check Authorization header
		next.ServeHTTP(w, r)
	})
}
`

var learnHTTPMiddlewareTestTmpl = `package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLoggingMiddleware(t *testing.T) {
	handler := LoggingMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	}))
	req := httptest.NewRequest("GET", "/test", nil)
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, req)
	if w.Code != 200 { t.Error("Middleware failed") }
}
`

// ============== Learn Error Handling Templates ==============

var learnErrorReadmeTmpl = `# Learn Error Handling in Go

Master Go error handling patterns.

## Sections
1. **basics/** - Error basics and checking
2. **wrapping/** - Error wrapping with context
3. **custom/** - Custom error types
`

var learnErrorBasicsTmpl = `package main

import (
	"errors"
	"fmt"
)

var ErrNotFound = errors.New("not found")
var ErrInvalid = errors.New("invalid input")

// TODO: Implement a function that returns appropriate errors
func FindUser(id int) (string, error) {
	if id < 0 { return "", ErrInvalid }
	if id == 0 { return "", ErrNotFound }
	return fmt.Sprintf("User%d", id), nil
}

// TODO: Check specific error types
func HandleError(err error) string {
	if errors.Is(err, ErrNotFound) { return "user not found" }
	if errors.Is(err, ErrInvalid) { return "invalid id" }
	return "unknown error"
}
`

var learnErrorBasicsTestTmpl = `package main

import (
	"errors"
	"testing"
)

func TestFindUser(t *testing.T) {
	_, err := FindUser(-1)
	if !errors.Is(err, ErrInvalid) { t.Error("Should be ErrInvalid") }
	
	_, err = FindUser(0)
	if !errors.Is(err, ErrNotFound) { t.Error("Should be ErrNotFound") }
	
	name, err := FindUser(1)
	if err != nil || name != "User1" { t.Error("Should succeed") }
}
`

var learnErrorWrappingTmpl = `package main

import (
	"errors"
	"fmt"
)

// TODO: Wrap errors with context
func LoadConfig(path string) error {
	err := readFile(path)
	if err != nil {
		return fmt.Errorf("loading config %s: %w", path, err)
	}
	return nil
}

func readFile(path string) error {
	return errors.New("file not found")
}

// TODO: Unwrap and check original error
func IsFileError(err error) bool {
	return errors.Is(err, errors.New("file not found"))
}
`

var learnErrorWrappingTestTmpl = `package main

import (
	"strings"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	err := LoadConfig("/etc/app.conf")
	if err == nil { t.Error("Should error") }
	if !strings.Contains(err.Error(), "loading config") { t.Error("Should wrap") }
}
`

var learnErrorCustomTmpl = `package main

import "fmt"

// TODO: Implement custom error type
type ValidationError struct {
	Field   string
	Message string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("%s: %s", e.Field, e.Message)
}

// TODO: Create error with type assertion
func ValidateAge(age int) error {
	if age < 0 {
		return &ValidationError{Field: "age", Message: "must be positive"}
	}
	if age > 150 {
		return &ValidationError{Field: "age", Message: "too large"}
	}
	return nil
}
`

var learnErrorCustomTestTmpl = `package main

import "testing"

func TestValidateAge(t *testing.T) {
	err := ValidateAge(-1)
	if ve, ok := err.(*ValidationError); !ok || ve.Field != "age" {
		t.Error("Should be ValidationError")
	}
	
	if err := ValidateAge(25); err != nil { t.Error("Should pass") }
}
`

// ============== Learn Interfaces Templates ==============

var learnInterfacesReadmeTmpl = `# Learn Interfaces in Go

Master Go interfaces and polymorphism.

## Sections
1. **basics/** - Interface basics
2. **composition/** - Interface embedding
3. **patterns/** - Common interface patterns
`

var learnInterfacesBasicsTmpl = `package main

import "fmt"

// TODO: Define a simple interface
type Speaker interface {
	Speak() string
}

type Dog struct{ Name string }
type Cat struct{ Name string }

func (d Dog) Speak() string { return fmt.Sprintf("%s says Woof!", d.Name) }
func (c Cat) Speak() string { return fmt.Sprintf("%s says Meow!", c.Name) }

// TODO: Implement a function that accepts interface
func MakeSpeak(s Speaker) string {
	return s.Speak()
}
`

var learnInterfacesBasicsTestTmpl = `package main

import "testing"

func TestSpeaker(t *testing.T) {
	dog := Dog{Name: "Rex"}
	cat := Cat{Name: "Whiskers"}
	
	if MakeSpeak(dog) != "Rex says Woof!" { t.Error("Dog failed") }
	if MakeSpeak(cat) != "Whiskers says Meow!" { t.Error("Cat failed") }
}
`

var learnInterfacesCompositionTmpl = `package main

// TODO: Compose interfaces
type Reader interface {
	Read(p []byte) (n int, err error)
}

type Writer interface {
	Write(p []byte) (n int, err error)
}

type ReadWriter interface {
	Reader
	Writer
}

// Buffer implements ReadWriter
type Buffer struct {
	data []byte
}

func (b *Buffer) Read(p []byte) (int, error) {
	n := copy(p, b.data)
	return n, nil
}

func (b *Buffer) Write(p []byte) (int, error) {
	b.data = append(b.data, p...)
	return len(p), nil
}
`

var learnInterfacesCompositionTestTmpl = `package main

import "testing"

func TestBuffer(t *testing.T) {
	var rw ReadWriter = &Buffer{}
	rw.Write([]byte("hello"))
	
	p := make([]byte, 5)
	n, _ := rw.Read(p)
	if n != 5 || string(p) != "hello" { t.Error("Buffer failed") }
}
`

var learnInterfacesPatternsTmpl = `package main

// TODO: Implement Stringer pattern
type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return p.Name
}

// TODO: Implement sort.Interface
type ByAge []Person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }
`

var learnInterfacesPatternsTestTmpl = `package main

import (
	"sort"
	"testing"
)

func TestStringer(t *testing.T) {
	p := Person{Name: "Alice", Age: 30}
	if p.String() != "Alice" { t.Error("Stringer failed") }
}

func TestSort(t *testing.T) {
	people := ByAge{{Name: "B", Age: 30}, {Name: "A", Age: 20}}
	sort.Sort(people)
	if people[0].Name != "A" { t.Error("Sort failed") }
}
`

// ============== Learn Design Patterns Templates ==============

var learnPatternsReadmeTmpl = `# Learn Design Patterns in Go

Learn common design patterns idiomatically in Go.

## Sections
1. **creational/** - Factory, Singleton patterns
2. **behavioral/** - Observer, Strategy patterns
3. **structural/** - (future: Decorator, Adapter)
`

var learnPatternsFactoryTmpl = `package creational

// Factory Pattern - create objects without exposing creation logic

type Animal interface {
	Speak() string
}

type Dog struct{}
func (d Dog) Speak() string { return "Woof!" }

type Cat struct{}
func (c Cat) Speak() string { return "Meow!" }

// TODO: Implement factory function
func NewAnimal(animalType string) Animal {
	switch animalType {
	case "dog": return Dog{}
	case "cat": return Cat{}
	default: return nil
	}
}
`

var learnPatternsFactoryTestTmpl = `package creational

import "testing"

func TestFactory(t *testing.T) {
	dog := NewAnimal("dog")
	if dog.Speak() != "Woof!" { t.Error("Dog factory failed") }
	
	cat := NewAnimal("cat")
	if cat.Speak() != "Meow!" { t.Error("Cat factory failed") }
}
`

var learnPatternsSingletonTmpl = `package creational

import "sync"

// Singleton Pattern - ensure only one instance exists

type Database struct {
	connection string
}

var (
	instance *Database
	once     sync.Once
)

// TODO: Implement thread-safe singleton
func GetDatabase() *Database {
	once.Do(func() {
		instance = &Database{connection: "connected"}
	})
	return instance
}
`

var learnPatternsSingletonTestTmpl = `package creational

import "testing"

func TestSingleton(t *testing.T) {
	db1 := GetDatabase()
	db2 := GetDatabase()
	if db1 != db2 { t.Error("Should be same instance") }
}
`

var learnPatternsObserverTmpl = `package behavioral

// Observer Pattern - notify subscribers of changes

type Observer interface {
	Update(message string)
}

type Subject struct {
	observers []Observer
}

func (s *Subject) Register(o Observer) {
	s.observers = append(s.observers, o)
}

func (s *Subject) Notify(message string) {
	for _, o := range s.observers {
		o.Update(message)
	}
}

type EmailSubscriber struct {
	Messages []string
}

func (e *EmailSubscriber) Update(message string) {
	e.Messages = append(e.Messages, message)
}
`

var learnPatternsObserverTestTmpl = `package behavioral

import "testing"

func TestObserver(t *testing.T) {
	subject := &Subject{}
	sub := &EmailSubscriber{}
	subject.Register(sub)
	subject.Notify("Hello!")
	if len(sub.Messages) != 1 { t.Error("Observer failed") }
}
`

var learnPatternsStrategyTmpl = `package behavioral

// Strategy Pattern - swap algorithms at runtime

type PaymentStrategy interface {
	Pay(amount float64) string
}

type CreditCard struct{}
func (c CreditCard) Pay(amount float64) string { return "Paid with credit card" }

type PayPal struct{}
func (p PayPal) Pay(amount float64) string { return "Paid with PayPal" }

type Checkout struct {
	strategy PaymentStrategy
}

func (c *Checkout) SetStrategy(s PaymentStrategy) {
	c.strategy = s
}

func (c *Checkout) Process(amount float64) string {
	return c.strategy.Pay(amount)
}
`

var learnPatternsStrategyTestTmpl = `package behavioral

import "testing"

func TestStrategy(t *testing.T) {
	checkout := &Checkout{}
	
	checkout.SetStrategy(CreditCard{})
	if checkout.Process(100) != "Paid with credit card" { t.Error("CC failed") }
	
	checkout.SetStrategy(PayPal{})
	if checkout.Process(100) != "Paid with PayPal" { t.Error("PayPal failed") }
}
`

// ============== Go Microservice Templates ==============

var microserviceMainTmpl = `package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"{{.ProjectName}}/internal/handler"
	"{{.ProjectName}}/internal/health"
	"{{.ProjectName}}/internal/middleware"
)

func main() {
	mux := http.NewServeMux()
	
	// Health endpoints
	mux.HandleFunc("/health", health.HealthHandler)
	mux.HandleFunc("/ready", health.ReadyHandler)
	
	// API endpoints
	mux.HandleFunc("/api/", handler.APIHandler)
	
	// Apply middleware
	h := middleware.Logging(mux)
	
	srv := &http.Server{
		Addr:         ":8080",
		Handler:      h,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	
	// Graceful shutdown
	go func() {
		log.Println("Server starting on :8080")
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalf("Server error: %v", err)
		}
	}()
	
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	
	log.Println("Shutting down...")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Shutdown error: %v", err)
	}
	log.Println("Server stopped")
}
`

var microserviceHandlerTmpl = `package handler

import (
	"encoding/json"
	"net/http"
)

func APIHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"status":  "ok",
		"message": "Hello from microservice!",
	})
}
`

var microserviceLoggingTmpl = `package middleware

import (
	"log"
	"net/http"
	"time"
)

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("%s %s %v", r.Method, r.URL.Path, time.Since(start))
	})
}
`

var microserviceHealthTmpl = `package health

import (
	"encoding/json"
	"net/http"
)

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "healthy"})
}

func ReadyHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "ready"})
}
`

var microserviceDockerfileTmpl = `FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o /server ./cmd/server

FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /server /server
EXPOSE 8080
CMD ["/server"]
`

var microserviceMakefileTmpl = `.PHONY: run build docker test

run:
	go run ./cmd/server

build:
	go build -o bin/server ./cmd/server

docker:
	docker build -t {{.ProjectName}} .

test:
	go test -v ./...
`

// ============== Go WebSocket Templates ==============

var websocketMainTmpl = `package main

import (
	"log"
	"net/http"

	"{{.ProjectName}}/internal/hub"
)

func main() {
	h := hub.NewHub()
	go h.Run()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		hub.ServeWs(h, w, r)
	})
	http.Handle("/", http.FileServer(http.Dir("./web")))

	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
`

var websocketHubTmpl = `package hub

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

type Hub struct {
	clients    map[*Client]bool
	broadcast  chan []byte
	register   chan *Client
	unregister chan *Client
}

func NewHub() *Hub {
	return &Hub{
		clients:    make(map[*Client]bool),
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		case message := <-h.broadcast:
			for client := range h.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		}
	}
}

func ServeWs(hub *Hub, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	client := &Client{hub: hub, conn: conn, send: make(chan []byte, 256)}
	hub.register <- client
	go client.writePump()
	go client.readPump()
}

func (h *Hub) Broadcast(msg []byte) {
	h.broadcast <- msg
}
`

var websocketClientTmpl = `package hub

import (
	"github.com/gorilla/websocket"
)

type Client struct {
	hub  *Hub
	conn *websocket.Conn
	send chan []byte
}

func (c *Client) readPump() {
	defer func() {
		c.hub.unregister <- c
		c.conn.Close()
	}()
	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil { break }
		c.hub.broadcast <- message
	}
}

func (c *Client) writePump() {
	defer c.conn.Close()
	for message := range c.send {
		if err := c.conn.WriteMessage(websocket.TextMessage, message); err != nil {
			return
		}
	}
}
`

var websocketHTMLTmpl = `<!DOCTYPE html>
<html>
<head>
    <title>WebSocket Chat</title>
    <style>
        body { font-family: sans-serif; padding: 20px; }
        #messages { border: 1px solid #ccc; height: 300px; overflow-y: scroll; padding: 10px; }
        #input { width: 80%; padding: 10px; }
        button { padding: 10px 20px; }
    </style>
</head>
<body>
    <h1>WebSocket Chat</h1>
    <div id="messages"></div>
    <input id="input" type="text" placeholder="Type a message...">
    <button onclick="send()">Send</button>
    <script>
        const ws = new WebSocket('ws://' + location.host + '/ws');
        ws.onmessage = (e) => {
            const div = document.createElement('div');
            div.textContent = e.data;
            document.getElementById('messages').appendChild(div);
        };
        function send() {
            const input = document.getElementById('input');
            ws.send(input.value);
            input.value = '';
        }
        document.getElementById('input').addEventListener('keypress', (e) => {
            if (e.key === 'Enter') send();
        });
    </script>
</body>
</html>
`

// ============== Go GraphQL Templates ==============

var graphqlMainTmpl = `package main

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"{{.ProjectName}}/graph"
)

func main() {
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{
		Resolvers: &graph.Resolver{},
	}))

	http.Handle("/", playground.Handler("GraphQL", "/query"))
	http.Handle("/query", srv)

	log.Println("GraphQL server on http://localhost:8080/")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
`

var graphqlSchemaTmpl = `type Query {
  todos: [Todo!]!
  todo(id: ID!): Todo
}

type Mutation {
  createTodo(input: NewTodo!): Todo!
}

type Todo {
  id: ID!
  text: String!
  done: Boolean!
}

input NewTodo {
  text: String!
}
`

var graphqlResolverTmpl = `package graph

type Resolver struct {
	todos []Todo
}

type Todo struct {
	ID   string
	Text string
	Done bool
}

func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

type queryResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }

func (r *queryResolver) Todos() ([]Todo, error) { return r.todos, nil }
func (r *queryResolver) Todo(id string) (*Todo, error) {
	for _, t := range r.todos {
		if t.ID == id { return &t, nil }
	}
	return nil, nil
}

func (r *mutationResolver) CreateTodo(input NewTodo) (Todo, error) {
	todo := Todo{ID: fmt.Sprintf("%d", len(r.todos)+1), Text: input.Text}
	r.todos = append(r.todos, todo)
	return todo, nil
}

type NewTodo struct { Text string }
`

var graphqlConfigTmpl = `schema:
  - graph/schema.graphqls

exec:
  filename: graph/generated.go
  package: graph

model:
  filename: graph/models_gen.go
  package: graph

resolver:
  filename: graph/resolver.go
  type: Resolver
`

// ============== Go Lambda Templates ==============

var lambdaMainTmpl = `package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"{{.ProjectName}}/internal/handler"
)

func main() {
	lambda.Start(handler.HandleRequest)
}
`

var lambdaHandlerTmpl = `package handler

import (
	"context"
	"fmt"
)

type Request struct {
	Name string ` + "`json:\"name\"`" + `
}

type Response struct {
	Message string ` + "`json:\"message\"`" + `
}

func HandleRequest(ctx context.Context, req Request) (Response, error) {
	if req.Name == "" {
		req.Name = "World"
	}
	return Response{Message: fmt.Sprintf("Hello, %s!", req.Name)}, nil
}
`

var lambdaSAMTmpl = `AWSTemplateFormatVersion: '2010-09-09'
Transform: AWS::Serverless-2016-10-31
Description: {{.ProjectName}} Lambda function

Globals:
  Function:
    Timeout: 30

Resources:
  Function:
    Type: AWS::Serverless::Function
    Properties:
      Handler: bootstrap
      Runtime: provided.al2
      CodeUri: .
      Events:
        Api:
          Type: Api
          Properties:
            Path: /hello
            Method: get
`

var lambdaMakefileTmpl = `.PHONY: build deploy test

build:
	GOOS=linux GOARCH=amd64 go build -o bootstrap ./cmd/lambda

deploy: build
	sam deploy --guided

test:
	go test -v ./...

local:
	sam local start-api
`

// ============== Go Cron Templates ==============

var cronMainTmpl = `package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"{{.ProjectName}}/internal/jobs"
	"{{.ProjectName}}/internal/scheduler"
)

func main() {
	s := scheduler.New()
	
	s.Schedule("*/5 * * * *", jobs.CleanupJob)
	s.Schedule("0 * * * *", jobs.ReportJob)
	
	s.Start()
	log.Println("Scheduler started")
	
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	
	s.Stop()
	log.Println("Scheduler stopped")
}
`

var cronJobsTmpl = `package jobs

import "log"

func CleanupJob() {
	log.Println("Running cleanup job...")
}

func ReportJob() {
	log.Println("Running report job...")
}
`

var cronSchedulerTmpl = `package scheduler

import (
	"github.com/robfig/cron/v3"
)

type Scheduler struct {
	cron *cron.Cron
}

func New() *Scheduler {
	return &Scheduler{cron: cron.New()}
}

func (s *Scheduler) Schedule(spec string, job func()) {
	s.cron.AddFunc(spec, job)
}

func (s *Scheduler) Start() { s.cron.Start() }
func (s *Scheduler) Stop()  { s.cron.Stop() }
`

// ============== Go Auth Templates ==============

var authMainTmpl = `package main

import (
	"log"
	"net/http"

	"{{.ProjectName}}/internal/auth"
	"{{.ProjectName}}/internal/handler"
)

func main() {
	mux := http.NewServeMux()
	
	// Public routes
	mux.HandleFunc("/api/login", handler.Login)
	mux.HandleFunc("/api/register", handler.Register)
	
	// Protected routes
	mux.Handle("/api/profile", auth.Middleware(http.HandlerFunc(handler.Profile)))
	
	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
`

var authJWTTmpl = `package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("your-secret-key-change-in-production")

type Claims struct {
	UserID   int64  ` + "`json:\"user_id\"`" + `
	Username string ` + "`json:\"username\"`" + `
	jwt.RegisteredClaims
}

func GenerateToken(userID int64, username string) (string, error) {
	claims := Claims{
		UserID:   userID,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func ValidateToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil { return nil, err }
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
`

var authMiddlewareTmpl = `package auth

import (
	"context"
	"net/http"
	"strings"
)

type contextKey string
const UserKey contextKey = "user"

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			http.Error(w, "Invalid auth header", http.StatusUnauthorized)
			return
		}
		
		claims, err := ValidateToken(parts[1])
		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}
		
		ctx := context.WithValue(r.Context(), UserKey, claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
`

var authHandlerTmpl = `package handler

import (
	"encoding/json"
	"net/http"

	"{{.ProjectName}}/internal/auth"
)

type LoginRequest struct {
	Username string ` + "`json:\"username\"`" + `
	Password string ` + "`json:\"password\"`" + `
}

func Login(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	json.NewDecoder(r.Body).Decode(&req)
	
	// TODO: Validate against database
	if req.Username == "" || req.Password == "" {
		http.Error(w, "Invalid credentials", http.StatusBadRequest)
		return
	}
	
	token, _ := auth.GenerateToken(1, req.Username)
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}

func Register(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement registration
	json.NewEncoder(w).Encode(map[string]string{"message": "User registered"})
}

func Profile(w http.ResponseWriter, r *http.Request) {
	claims := r.Context().Value(auth.UserKey).(*auth.Claims)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"user_id":  claims.UserID,
		"username": claims.Username,
	})
}
`

var authUserModelTmpl = `package model

type User struct {
	ID       int64  ` + "`json:\"id\"`" + `
	Username string ` + "`json:\"username\"`" + `
	Email    string ` + "`json:\"email\"`" + `
	Password string ` + "`json:\"-\"`" + `
}
`

// ============== Go Kafka Templates ==============

var kafkaProducerMainTmpl = `package main

import (
	"log"
	"{{.ProjectName}}/internal/kafka"
)

func main() {
	producer, err := kafka.NewProducer([]string{"localhost:9092"})
	if err != nil { log.Fatal(err) }
	defer producer.Close()
	
	err = producer.Send("my-topic", "Hello Kafka!")
	if err != nil { log.Fatal(err) }
	log.Println("Message sent!")
}
`

var kafkaConsumerMainTmpl = `package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"{{.ProjectName}}/internal/kafka"
)

func main() {
	consumer, err := kafka.NewConsumer([]string{"localhost:9092"}, "my-group")
	if err != nil { log.Fatal(err) }
	
	go consumer.Consume("my-topic", func(msg string) {
		log.Printf("Received: %s\n", msg)
	})
	
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	consumer.Close()
}
`

var kafkaProducerTmpl = `package kafka

import (
	"github.com/IBM/sarama"
)

type Producer struct {
	producer sarama.SyncProducer
}

func NewProducer(brokers []string) (*Producer, error) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	
	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil { return nil, err }
	
	return &Producer{producer: producer}, nil
}

func (p *Producer) Send(topic, message string) error {
	_, _, err := p.producer.SendMessage(&sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(message),
	})
	return err
}

func (p *Producer) Close() { p.producer.Close() }
`

var kafkaConsumerTmpl = `package kafka

import (
	"github.com/IBM/sarama"
)

type Consumer struct {
	consumer sarama.Consumer
}

func NewConsumer(brokers []string, group string) (*Consumer, error) {
	config := sarama.NewConfig()
	consumer, err := sarama.NewConsumer(brokers, config)
	if err != nil { return nil, err }
	return &Consumer{consumer: consumer}, nil
}

func (c *Consumer) Consume(topic string, handler func(string)) error {
	partitions, _ := c.consumer.Partitions(topic)
	for _, partition := range partitions {
		pc, _ := c.consumer.ConsumePartition(topic, partition, sarama.OffsetNewest)
		go func(pc sarama.PartitionConsumer) {
			for msg := range pc.Messages() {
				handler(string(msg.Value))
			}
		}(pc)
	}
	return nil
}

func (c *Consumer) Close() { c.consumer.Close() }
`

var kafkaDockerComposeTmpl = `version: '3.8'
services:
  zookeeper:
    image: confluentinc/cp-zookeeper:latest
    environment:
      ZOOKEEPER_CLIENT_PORT: 2181
    ports:
      - "2181:2181"

  kafka:
    image: confluentinc/cp-kafka:latest
    depends_on:
      - zookeeper
    ports:
      - "9092:9092"
    environment:
      KAFKA_BROKER_ID: 1
      KAFKA_ZOOKEEPER_CONNECT: zookeeper:2181
      KAFKA_ADVERTISED_LISTENERS: PLAINTEXT://localhost:9092
      KAFKA_OFFSETS_TOPIC_REPLICATION_FACTOR: 1
`

// ============== Go Redis Templates ==============

var redisMainTmpl = `package main

import (
	"log"
	"{{.ProjectName}}/internal/cache"
)

func main() {
	client := cache.NewRedisClient("localhost:6379")
	defer client.Close()
	
	client.Set("key", "value", 0)
	val, _ := client.Get("key")
	log.Printf("Got: %s\n", val)
}
`

var redisCacheTmpl = `package cache

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisClient struct {
	client *redis.Client
	ctx    context.Context
}

func NewRedisClient(addr string) *RedisClient {
	return &RedisClient{
		client: redis.NewClient(&redis.Options{Addr: addr}),
		ctx:    context.Background(),
	}
}

func (r *RedisClient) Set(key, value string, expiration time.Duration) error {
	return r.client.Set(r.ctx, key, value, expiration).Err()
}

func (r *RedisClient) Get(key string) (string, error) {
	return r.client.Get(r.ctx, key).Result()
}

func (r *RedisClient) Delete(key string) error {
	return r.client.Del(r.ctx, key).Err()
}

func (r *RedisClient) Close() { r.client.Close() }
`

var redisPubSubTmpl = `package pubsub

import (
	"context"
	"github.com/redis/go-redis/v9"
)

type PubSub struct {
	client *redis.Client
	ctx    context.Context
}

func NewPubSub(addr string) *PubSub {
	return &PubSub{
		client: redis.NewClient(&redis.Options{Addr: addr}),
		ctx:    context.Background(),
	}
}

func (p *PubSub) Publish(channel, message string) error {
	return p.client.Publish(p.ctx, channel, message).Err()
}

func (p *PubSub) Subscribe(channel string, handler func(string)) {
	sub := p.client.Subscribe(p.ctx, channel)
	ch := sub.Channel()
	for msg := range ch {
		handler(msg.Payload)
	}
}
`

var redisDockerComposeTmpl = `version: '3.8'
services:
  redis:
    image: redis:alpine
    ports:
      - "6379:6379"
`

// ============== Go Clean Architecture Templates ==============

var cleanArchMainTmpl = `package main

import (
	"log"
	"net/http"

	"{{.ProjectName}}/internal/delivery/http"
	"{{.ProjectName}}/internal/repository"
	"{{.ProjectName}}/internal/usecase"
)

func main() {
	repo := repository.NewUserRepository()
	uc := usecase.NewUserUsecase(repo)
	handler := http.NewHandler(uc)
	
	mux := http.NewServeMux()
	mux.HandleFunc("/users", handler.GetUsers)
	
	log.Println("Server on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
`

var cleanArchEntityTmpl = `package entity

type User struct {
	ID    int64  ` + "`json:\"id\"`" + `
	Name  string ` + "`json:\"name\"`" + `
	Email string ` + "`json:\"email\"`" + `
}
`

var cleanArchUsecaseTmpl = `package usecase

import "{{.ProjectName}}/internal/entity"

type UserRepository interface {
	FindAll() ([]entity.User, error)
	FindByID(id int64) (*entity.User, error)
}

type UserUsecase struct {
	repo UserRepository
}

func NewUserUsecase(repo UserRepository) *UserUsecase {
	return &UserUsecase{repo: repo}
}

func (u *UserUsecase) GetUsers() ([]entity.User, error) {
	return u.repo.FindAll()
}

func (u *UserUsecase) GetUser(id int64) (*entity.User, error) {
	return u.repo.FindByID(id)
}
`

var cleanArchRepoTmpl = `package repository

import "{{.ProjectName}}/internal/entity"

type UserRepository struct {
	users []entity.User
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		users: []entity.User{
			{ID: 1, Name: "Alice", Email: "alice@example.com"},
			{ID: 2, Name: "Bob", Email: "bob@example.com"},
		},
	}
}

func (r *UserRepository) FindAll() ([]entity.User, error) {
	return r.users, nil
}

func (r *UserRepository) FindByID(id int64) (*entity.User, error) {
	for _, u := range r.users {
		if u.ID == id { return &u, nil }
	}
	return nil, nil
}
`

var cleanArchHandlerTmpl = `package http

import (
	"encoding/json"
	"net/http"

	"{{.ProjectName}}/internal/usecase"
)

type Handler struct {
	uc *usecase.UserUsecase
}

func NewHandler(uc *usecase.UserUsecase) *Handler {
	return &Handler{uc: uc}
}

func (h *Handler) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.uc.GetUsers()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	json.NewEncoder(w).Encode(users)
}
`

var cleanArchErrorsTmpl = `package errors

import "errors"

var (
	ErrNotFound = errors.New("not found")
	ErrInvalid  = errors.New("invalid input")
)

type AppError struct {
	Code    int
	Message string
}

func (e *AppError) Error() string { return e.Message }
`

// ============== Go Monorepo Templates ==============

var monorepoAPITmpl = `package main

import (
	"log"
	"net/http"

	"{{.ProjectName}}/pkg/shared"
)

func main() {
	cfg := shared.LoadConfig()
	logger := shared.NewLogger("api")
	
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		logger.Info("Request received")
		w.Write([]byte("API Service"))
	})
	
	logger.Info("API starting on " + cfg.APIPort)
	log.Fatal(http.ListenAndServe(cfg.APIPort, nil))
}
`

var monorepoWorkerTmpl = `package main

import (
	"log"
	"time"

	"{{.ProjectName}}/pkg/shared"
)

func main() {
	cfg := shared.LoadConfig()
	logger := shared.NewLogger("worker")
	
	logger.Info("Worker starting...")
	
	for {
		logger.Info("Processing job...")
		time.Sleep(time.Duration(cfg.WorkerInterval) * time.Second)
	}
}
`

var monorepoConfigTmpl = `package shared

import "os"

type Config struct {
	APIPort        string
	WorkerInterval int
	Environment    string
}

func LoadConfig() *Config {
	return &Config{
		APIPort:        getEnv("API_PORT", ":8080"),
		WorkerInterval: 5,
		Environment:    getEnv("ENV", "development"),
	}
}

func getEnv(key, def string) string {
	if val := os.Getenv(key); val != "" { return val }
	return def
}
`

var monorepoLoggerTmpl = `package shared

import (
	"log"
	"os"
)

type Logger struct {
	service string
	logger  *log.Logger
}

func NewLogger(service string) *Logger {
	return &Logger{
		service: service,
		logger:  log.New(os.Stdout, "", log.LstdFlags),
	}
}

func (l *Logger) Info(msg string) {
	l.logger.Printf("[%s] INFO: %s", l.service, msg)
}

func (l *Logger) Error(msg string) {
	l.logger.Printf("[%s] ERROR: %s", l.service, msg)
}
`

var monorepoMakefileTmpl = `.PHONY: api worker build-all

api:
	go run ./services/api

worker:
	go run ./services/worker

build-all:
	go build -o bin/api ./services/api
	go build -o bin/worker ./services/worker

test:
	go test -v ./...
`
