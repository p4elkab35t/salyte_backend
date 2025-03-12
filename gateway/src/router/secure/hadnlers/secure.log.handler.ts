import { Logger } from "../../../logger/logger";
import securityLogsGRPCClient from "../../../grpc_client/log.client";
import type { SecurityLogsService } from "../../../grpc_client/proto/log.proto";
import { corsResponse } from "../../../misc/request";

// Type the client object as AuthService
const securityLogs: SecurityLogsService = securityLogsGRPCClient;

const routes = new Map<string, { method: string[], handler: Function }>([
    ["/lifecheck", { method: ["GET"], handler: lifeCheck }],
    ["/all", { method: ["GET"], handler: getAllLogs }],
    ["/current", { method: ["GET"], handler: getCurrentLog }]
  ]);

  export default async function secureLogServiceHandler(req: Request, restPath: string) {
    const url = new URL(req.url);
    const path = restPath.replace("/secure/log", "");
  
    Logger.info(`Security Logs request`, { method: req.method, path });
  
    if (routes.has(path)) {
      const route = routes.get(path);
      if (!route?.method.includes(req.method)) {
        Logger.warn(`Method not allowed`, { path, method: req.method });
        return new corsResponse(JSON.stringify({ error: "Method Not Allowed" }), { status: 405 });
      }
  
      return route.handler(req);
    }
  
    Logger.warn(`Security Logs  route not found`, { path });
    return new corsResponse(JSON.stringify({ error: "Route not found" }), { status: 404 });
  }

  async function lifeCheck(req: Request, restPath: string) {
    Logger.info("Life check", { method: req.method });
    return new corsResponse(JSON.stringify({ message: "Security Logs Service Alive" }), { status: 200 });
  }
  
  async function getAllLogs(req: Request) {
    Logger.info("Getting logs for user", { method: req.method });
    try{
    const token = req.headers.get("Authorization")?.split("Bearer ")[1];
    if (!token) {
      return new corsResponse(JSON.stringify({ error: "No token provided" }), { status: 401 });
    }
    // get user id from quuery params
    const URLQuery = new URL(req.url);
    const userID = URLQuery.searchParams.get("user_id");
    if (!userID) {
      return new corsResponse(JSON.stringify({ error: "No user id provided" }), { status: 401 });
    }
    let page = URLQuery.searchParams.get("page");
    if(!page){
        page = "1";
    }
    let limit = URLQuery.searchParams.get("limit");
    if(!limit){
        limit = "10";
    }
    const result = await promisifyCallback(securityLogs.GetSecurityLogsWithUsedID,{ user_id: userID, page: parseInt(page), limit: parseInt(limit) });
    if (result.status !== 0){
      return new corsResponse(JSON.stringify({ error: "Getting Logs failed" }), { status: 401 });
    }
    Logger.info("Logs Retrieved", { result });
    return new corsResponse(JSON.stringify(result), { status: 200 });
  } catch (error) {
    return new corsResponse(JSON.stringify({ error: "Getting Logs failed" }), { status: 401 });
  }
  }

  async function getCurrentLog(req: Request) {
    Logger.info("Getting current log", { method: req.method });
    const URLQuery = new URL(req.url);
    try{
      const logID = URLQuery.searchParams.get("log_id");
      const result = await promisifyCallback(securityLogs.GetSecurityLogWithID,{ logId: logID });
      if (result.status !== 0){
        return new corsResponse(JSON.stringify({ error: "Getting log failed" }), { status: 401 });
      }
      Logger.info("Getting log success", { result });
      return new corsResponse(JSON.stringify(result), { status: 200 });
  } catch (error) {
    return new corsResponse(JSON.stringify({ error: "Getting log failed" }), { status: 401 });
  }
  }

  const promisifyCallback = (method: Function, request: any): Promise<any> => {
    return new Promise((resolve, reject) => {
      method.call(securityLogs, request, (error: any, response: any) => {
        if (error) {
          reject(error); 
        } else {
          resolve(response); 
        }
      });
    });
  }