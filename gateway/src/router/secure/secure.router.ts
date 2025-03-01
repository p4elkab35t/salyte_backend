import authServiceHandler from "./hadnlers/auth.handler";
import secureLogServiceHandler from "./hadnlers/secure.log.handler";
import { Logger } from "../../logger/logger";


const routes = new Map<string, Function >([
    ["/lifecheck", lifeCheck],
    ["/auth", authServiceHandler],
    ["/log", secureLogServiceHandler],
  ]);

export default async function secureServiceHandler(req: Request, restPath: string) {
    const url = new URL(req.url);
    const path = restPath.replace("/secure", "");

    const urlArray = url.pathname.split("/");
    if (urlArray.length < 3) {
        return new Response(JSON.stringify({ error: "No path specified" }), { status: 404 });
    }
    if (urlArray[1] !== "api") {
        return new Response(JSON.stringify({ error: "Invalid API prefix" }), { status: 404 });
    }
    if (urlArray[2] !== "secure") {
        return new Response(JSON.stringify({ error: "Invalid API prefix" }), { status: 404 });
    }
    
    const routePath = "/"+ url.pathname.split("/")[3];



    Logger.info(`Secure request`, { method: req.method, routePath });

    if (routes.has(routePath)) {
      const route = routes.get(routePath);
      if (!route) {
        Logger.warn(`Route not found`, { path });
        return new Response(JSON.stringify({ error: "Route Not Found" }), { status: 404 });
      }

      return route(req, "/secure"+restPath);
    }

    Logger.warn(`Secure route not found`, { path });
    return new Response(JSON.stringify({ error: "Route not found" }), { status: 404 });
  }

    async function lifeCheck(req: Request, restPath: string) {
        Logger.info("Life check", { method: req.method });
        return new Response(JSON.stringify({ message: "Secure Service Alive" }), { status: 200 });
    }