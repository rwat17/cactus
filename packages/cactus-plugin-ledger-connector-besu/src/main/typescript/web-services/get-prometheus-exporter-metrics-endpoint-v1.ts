import { Express, Request, Response } from "express";

import { handleRestEndpointException, registerWebServiceEndpoint } from "@hyperledger/cactus-core";

import OAS from "../../json/openapi.json";

import {
  IWebServiceEndpoint,
  IExpressRequestHandler,
  IEndpointAuthzOptions,
} from "@hyperledger/cactus-core-api";

import {
  LogLevelDesc,
  Logger,
  LoggerProvider,
  Checks,
  IAsyncProvider,
} from "@hyperledger/cactus-common";

import { PluginLedgerConnectorBesu } from "../plugin-ledger-connector-besu";

export interface IGetPrometheusExporterMetricsEndpointV1Options {
  connector: PluginLedgerConnectorBesu;
  logLevel?: LogLevelDesc;
}

export class GetPrometheusExporterMetricsEndpointV1
  implements IWebServiceEndpoint {
  private readonly log: Logger;

  constructor(
    public readonly options: IGetPrometheusExporterMetricsEndpointV1Options,
  ) {
    const fnTag = "GetPrometheusExporterMetricsEndpointV1#constructor()";

    Checks.truthy(options, `${fnTag} options`);
    Checks.truthy(options.connector, `${fnTag} options.connector`);

    const label = "get-prometheus-exporter-metrics-endpoint";
    const level = options.logLevel || "INFO";
    this.log = LoggerProvider.getOrCreate({ label, level });
  }

  getAuthorizationOptionsProvider(): IAsyncProvider<IEndpointAuthzOptions> {
    // TODO: make this an injectable dependency in the constructor
    return {
      get: async () => ({
        isProtected: true,
        requiredRoles: [],
      }),
    };
  }

  public getExpressRequestHandler(): IExpressRequestHandler {
    return this.handleRequest.bind(this);
  }

  public get oasPath(): (typeof OAS.paths)["/api/v1/plugins/@hyperledger/cactus-plugin-ledger-connector-besu/get-prometheus-exporter-metrics"] {
    return OAS.paths[
      "/api/v1/plugins/@hyperledger/cactus-plugin-ledger-connector-besu/get-prometheus-exporter-metrics"
    ];
  }

  public getPath(): string {
    return this.oasPath.get["x-hyperledger-cacti"].http.path;
  }

  public getVerbLowerCase(): string {
    return this.oasPath.get["x-hyperledger-cacti"].http.verbLowerCase;
  }

  public getOperationId(): string {
    return this.oasPath.get.operationId;
  }

  public async registerExpress(
    expressApp: Express,
  ): Promise<IWebServiceEndpoint> {
    await registerWebServiceEndpoint(expressApp, this);
    return this;
  }

  async handleRequest(req: Request, res: Response): Promise<void> {
    const fnTag = "GetPrometheusExporterMetrics#handleRequest()";
    const reqTag = `${this.getVerbLowerCase()} - ${this.getPath()}`;
    const verbUpper = this.getVerbLowerCase().toUpperCase();
    this.log.debug(`${verbUpper} ${this.getPath()}`);

    try {
      const resBody =
        await this.options.connector.getPrometheusExporterMetrics();
      res.status(200);
      res.send(resBody);
    } catch (ex) {
      this.log.error(`${fnTag} failed to serve request`, ex);
      if (typeof ex === 'object' && ex !== null) {
        if ('message' in ex && typeof ex.message === 'string') {
          const errorMsg = ex.message
          handleRestEndpointException({ errorMsg, log: this.log, error: ex, res })
        }
      }
      else {
        this.log.error(`Crash while serving ${reqTag}`, ex);
        const errorMsg = `Internal server Error`;
        handleRestEndpointException({ errorMsg, log: this.log, error: ex, res })
      }
    }
  }
}
