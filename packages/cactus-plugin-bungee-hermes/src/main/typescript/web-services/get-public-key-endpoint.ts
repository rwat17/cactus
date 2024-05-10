import type { Express, Request, Response } from "express";

import {
  IWebServiceEndpoint,
  IExpressRequestHandler,
  IEndpointAuthzOptions,
} from "@hyperledger/cactus-core-api/";

import {
  Logger,
  Checks,
  LogLevelDesc,
  LoggerProvider,
  IAsyncProvider,
} from "@hyperledger/cactus-common";

import {
  handleRestEndpointException,
  registerWebServiceEndpoint,
} from "@hyperledger/cactus-core";

import OAS from "../../json/openapi.json";
import { PluginBungeeHermes } from "../plugin-bungee-hermes";

export interface GetPublicKeyEndpointOptions {
  logLevel?: LogLevelDesc;
  bungee: PluginBungeeHermes;
}

export class GetPublicKeyEndpointV1 implements IWebServiceEndpoint {
  public static readonly CLASS_NAME = "ClientEndpointV1";

  private readonly log: Logger;

  public get className(): string {
    return GetPublicKeyEndpointV1.CLASS_NAME;
  }

  constructor(public readonly options: GetPublicKeyEndpointOptions) {
    const fnTag = `${this.className}#constructor()`;
    Checks.truthy(options, `${fnTag} arg options`);
    Checks.truthy(options.bungee, `${fnTag} arg options.connector`);

    const level = this.options.logLevel || "INFO";
    const label = this.className;
    this.log = LoggerProvider.getOrCreate({ level, label });
  }

  public getExpressRequestHandler(): IExpressRequestHandler {
    return this.handleRequest.bind(this);
  }

  public getPath(): string {
    const apiPath =
      OAS.paths[
        "/api/v1/plugins/@hyperledger/cactus-plugin-bungee-hermes/get-public-key"
      ];
    return apiPath.get["x-hyperledger-cacti"].http.path;
  }

  public getVerbLowerCase(): string {
    const apiPath =
      OAS.paths[
        "/api/v1/plugins/@hyperledger/cactus-plugin-bungee-hermes/get-public-key"
      ];
    return apiPath.get["x-hyperledger-cacti"].http.verbLowerCase;
  }

  public getOperationId(): string {
    return OAS.paths[
      "/api/v1/plugins/@hyperledger/cactus-plugin-bungee-hermes/get-public-key"
    ].get.operationId;
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

  public async registerExpress(
    expressApp: Express,
  ): Promise<IWebServiceEndpoint> {
    await registerWebServiceEndpoint(expressApp, this);
    return this;
  }

  public async handleRequest(_req: Request, res: Response): Promise<void> {
    const fnTag = `${this.className}#handleRequest()`;
    const reqTag = `${this.getVerbLowerCase()} - ${this.getPath()}`;
    this.log.debug(reqTag);
    try {
      const pubKey = this.options.bungee.getPublicKey();
      res.status(200).json(pubKey);
    } catch (ex) {
      const errorMsg = `${fnTag} request handler fn crashed for: ${reqTag}`;
      handleRestEndpointException({ errorMsg, log: this.log, error: ex, res });
    }
  }
}
