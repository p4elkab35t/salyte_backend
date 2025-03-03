// import socialHandler from './router/social/social.handler';
// import messageHandler from './router/message/message.handler';
import Bun from "bun"
import secureServiceHandler from "./router/secure/secure.router";
import { CONFIG } from "./config/config";
import { Logger } from "./logger/logger";
import socialServiceHandler from "./router/social/social.handler";

// console.log(import.meta.dir);

const services = new Map<string, Function>([
    ["/status", apiStatusHandler],
    ["/secure", secureServiceHandler],
    // ["/message", messageHandler],
    ["/social", socialServiceHandler],
])

const server = Bun.serve({
    port: CONFIG.PORT,
    fetch(req: Request) {
      const url = new URL(req.url);
      const apiCheck = url.pathname.split("/")[1];
      if (apiCheck !== "api") {
        Logger.error(`Invalid API prefix`, { path: url.pathname });
        return new Response("Invalid path", { status: 404 });
      }
      const service = `/${url.pathname.split("/")[2]}`;

      Logger.info(`Incoming request`, { method: req.method, url: req.url });

      if (!service){
        Logger.error(`Service not found`, { service: service });
        return new Response("No path specified", { status: 404 });
      }

      if (!services.has(service)){
        Logger.warn(`Service not found`, { service: service });
        return new Response("Service not found", { status: 404 });
      }

      const handler = services.get(service);
      const restPath = "/"+url.pathname.split("/").slice(3).join("/");
      if (handler) {
        Logger.info(`Service found`, { service: service });
        return handler(req, restPath);
      }
      else{
        Logger.error(`No route attached`, { url: req.url });
        return new Response(JSON.stringify({ error: "No route attached" }), { status: 404 });
      }
    },
  });
  
  Logger.info(`Listening on http://localhost:${server.port} ...`);

  async function apiStatusHandler (req: Request, restPath: string) {
    Logger.info("API status check", { method: req.method });
    return new Response(JSON.stringify({ message: "API Gateway Alive" }), { status: 200 });
  }