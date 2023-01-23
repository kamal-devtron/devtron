//go:build wireinject
// +build wireinject

package main

import (
	"github.com/devtron-labs/authenticator/client"
	"github.com/devtron-labs/devtron/api/cluster"
	"github.com/devtron-labs/devtron/api/connector"
	client2 "github.com/devtron-labs/devtron/api/helm-app"
	"github.com/devtron-labs/devtron/api/terminal"
	"github.com/devtron-labs/devtron/client/dashboard"
	"github.com/devtron-labs/devtron/internal/util"
	delete2 "github.com/devtron-labs/devtron/pkg/delete"
	"github.com/devtron-labs/devtron/pkg/kubernetesResourceAuditLogs"
	"github.com/devtron-labs/devtron/pkg/sql"
	"github.com/devtron-labs/devtron/pkg/user"
	"github.com/devtron-labs/devtron/pkg/user/casbin"
	"github.com/devtron-labs/devtron/pkg/user/noop"
	util2 "github.com/devtron-labs/devtron/pkg/util"
	"github.com/devtron-labs/devtron/util/argo"
	"github.com/devtron-labs/devtron/util/k8s"
	"github.com/devtron-labs/devtron/util/rbac"
	"github.com/google/wire"
)

func InitializeApp() (*App, error) {
	wire.Build(
		//user.SelfRegistrationWireSet,

		sql.NewNoopConnection,
		//sql.PgSqlWireSet,
		//user.UserWireSet,
		casbin.NewNoopEnforcer,
		wire.Bind(new(casbin.Enforcer), new(*casbin.NoopEnforcer)),
		//sso.SsoConfigWireSet,
		//AuthWireSet,
		//externalLink.ExternalLinkWireSet,
		//team.TeamsWireSet,
		cluster.ClusterWireSetK8sClient,
		dashboard.DashboardWireSet,
		//client.HelmAppWireSet,
		client2.NewNoopServiceImpl,
		wire.Bind(new(client2.HelmAppService), new(*client2.HelmAppServiceImpl)),
		k8s.K8sApplicationWireSet,
		//chartRepo.ChartRepositoryWireSet,
		//appStoreDiscover.AppStoreDiscoverWireSet,
		//appStoreValues.AppStoreValuesWireSet,
		//appStoreDeployment.AppStoreDeploymentWireSet,
		//server.ServerWireSet,
		//module.ModuleWireSet,
		//apiToken.ApiTokenWireSet,
		//webhookHelm.WebhookHelmWireSet,
		terminal.TerminalWireSet,
		client.GetRuntimeConfig,

		noop.NewNoopUserService,
		wire.Bind(new(user.UserService), new(*noop.NoopUserService)),

		NewApp,
		NewMuxRouter,
		//util3.GetGlobalEnvVariables,
		//util.NewHttpClient,
		util.NewSugardLogger,
		util.NewK8sUtil,
		util.IntValidator,
		util2.GetACDAuthConfig,
		//telemetry.NewPosthogClient,
		wire.Bind(new(delete2.DeleteService), new(*delete2.DeleteServiceImpl)),
		delete2.NewNoopServiceImpl,

		rbac.NewNoopEnforcerUtilHelm,
		wire.Bind(new(rbac.EnforcerUtilHelm), new(*rbac.EnforcerUtilHelmImpl)),

		rbac.NewNoopEnforcerUtil,
		wire.Bind(new(rbac.EnforcerUtil), new(*rbac.EnforcerUtilImpl)),

		//router.NewAppRouterImpl,
		//wire.Bind(new(router.AppRouter), new(*router.AppRouterImpl)),
		//restHandler.NewAppRestHandlerImpl,
		//wire.Bind(new(restHandler.AppRestHandler), new(*restHandler.AppRestHandlerImpl)),
		//
		//app.NewAppCrudOperationServiceImpl,
		//wire.Bind(new(app.AppCrudOperationService), new(*app.AppCrudOperationServiceImpl)),
		//pipelineConfig.NewAppLabelRepositoryImpl,
		//wire.Bind(new(pipelineConfig.AppLabelRepository), new(*pipelineConfig.AppLabelRepositoryImpl)),
		////acd session client bind with authenticator login
		//wire.Bind(new(session.ServiceClient), new(*middleware.LoginService)),
		connector.NewPumpImpl,
		wire.Bind(new(connector.Pump), new(*connector.PumpImpl)),

		//telemetry.NewTelemetryEventClientImpl,
		//wire.Bind(new(telemetry.TelemetryEventClient), new(*telemetry.TelemetryEventClientImpl)),
		//
		//wire.Bind(new(delete2.DeleteService), new(*delete2.DeleteServiceImpl)),

		// needed for enforcer util
		//pipelineConfig.NewPipelineRepositoryImpl,
		//wire.Bind(new(pipelineConfig.PipelineRepository), new(*pipelineConfig.PipelineRepositoryImpl)),
		//app2.NewAppRepositoryImpl,
		//wire.Bind(new(app2.AppRepository), new(*app2.AppRepositoryImpl)),
		//router.NewAttributesRouterImpl,
		//wire.Bind(new(router.AttributesRouter), new(*router.AttributesRouterImpl)),
		//restHandler.NewAttributesRestHandlerImpl,
		//wire.Bind(new(restHandler.AttributesRestHandler), new(*restHandler.AttributesRestHandlerImpl)),
		//attributes.NewAttributesServiceImpl,
		//wire.Bind(new(attributes.AttributesService), new(*attributes.AttributesServiceImpl)),
		//repository.NewAttributesRepositoryImpl,
		//wire.Bind(new(repository.AttributesRepository), new(*repository.AttributesRepositoryImpl)),
		//pipelineConfig.NewCiPipelineRepositoryImpl,
		//wire.Bind(new(pipelineConfig.CiPipelineRepository), new(*pipelineConfig.CiPipelineRepositoryImpl)),
		// // needed for enforcer util ends

		// binding gitops to helm (for hyperion)
		//wire.Bind(new(appStoreDeploymentGitopsTool.AppStoreDeploymentArgoCdService), new(*appStoreDeploymentTool.AppStoreDeploymentHelmServiceImpl)),
		//
		//wire.Value(chartRepoRepository.RefChartDir("scripts/devtron-reference-helm-charts")),
		//
		//router.NewTelemetryRouterImpl,
		//wire.Bind(new(router.TelemetryRouter), new(*router.TelemetryRouterImpl)),
		//restHandler.NewTelemetryRestHandlerImpl,
		//wire.Bind(new(restHandler.TelemetryRestHandler), new(*restHandler.TelemetryRestHandlerImpl)),
		//
		////needed for sending events
		//dashboardEvent.NewDashboardTelemetryRestHandlerImpl,
		//wire.Bind(new(dashboardEvent.DashboardTelemetryRestHandler), new(*dashboardEvent.DashboardTelemetryRestHandlerImpl)),
		//dashboardEvent.NewDashboardTelemetryRouterImpl,
		//wire.Bind(new(dashboardEvent.DashboardTelemetryRouter),
		//	new(*dashboardEvent.DashboardTelemetryRouterImpl)),
		//
		//repository.NewGitOpsConfigRepositoryImpl,
		//wire.Bind(new(repository.GitOpsConfigRepository), new(*repository.GitOpsConfigRepositoryImpl)),
		//
		////binding argoUserService to helm via dummy implementation(HelmUserServiceImpl)
		argo.NewHelmUserServiceImpl,
		wire.Bind(new(argo.ArgoUserService), new(*argo.HelmUserServiceImpl)),
		//
		//router.NewUserAttributesRouterImpl,
		//wire.Bind(new(router.UserAttributesRouter), new(*router.UserAttributesRouterImpl)),
		//restHandler.NewUserAttributesRestHandlerImpl,
		//wire.Bind(new(restHandler.UserAttributesRestHandler), new(*restHandler.UserAttributesRestHandlerImpl)),
		//attributes.NewUserAttributesServiceImpl,
		//wire.Bind(new(attributes.UserAttributesService), new(*attributes.UserAttributesServiceImpl)),
		//repository.NewUserAttributesRepositoryImpl,
		//wire.Bind(new(repository.UserAttributesRepository), new(*repository.UserAttributesRepositoryImpl)),
		//util3.GetDevtronSecretName,
		//
		//repository2.NewK8sResourceHistoryRepositoryImpl,
		//wire.Bind(new(repository2.K8sResourceHistoryRepository), new(*repository2.K8sResourceHistoryRepositoryImpl)),
		//
		kubernetesResourceAuditLogs.NewNoopServiceImpl,
		wire.Bind(new(kubernetesResourceAuditLogs.K8sResourceHistoryService), new(*kubernetesResourceAuditLogs.K8sResourceHistoryServiceImpl)),
	)
	return &App{}, nil
}