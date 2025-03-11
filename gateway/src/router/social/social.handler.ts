import { Logger } from "../../logger/logger";
import authServiceHandler from "../secure/hadnlers/auth.handler";
import { CONFIG } from "../../config/config";

const socialServiceURL = `http://${CONFIG.SOCIAL_SERVICE_URL}/social`;

const routes = new Map<string, { method: string[], handler: Function }>([
    ["/profile", { method: ["GET", "PUT"], handler: proxyToSocialService }],
    ["/follow", { method: ["POST", "DELETE"], handler: proxyToSocialService }],
    ["/friends", { method: ["GET", "POST"], handler: proxyToSocialService }],
    ["/community", { method: ["GET", "POST", "PUT"], handler: proxyToSocialService }],
    ["/post", { method: ["GET", "POST", "PUT", "DELETE"], handler: proxyToSocialService }],
    ["/post/comment", { method: ["GET", "POST", "PUT", "DELETE"], handler: proxyToSocialService }],
    ["/post/likes", { method: ["GET"], handler: proxyToSocialService }],
    ["/post/like", { method: ["POST", "DELETE"], handler: proxyToSocialService }],
]);

export default async function socialServiceHandler(req: Request, restPath: string) {
    const url = new URL(req.url);
    const path = restPath.replace("/social", "");

    Logger.info(`Social request`, { method: req.method, path });

    if (routes.has(path)) {
        const route = routes.get(path);
        if (!route?.method.includes(req.method)) {
            Logger.warn(`Method not allowed`, { path, method: req.method });
            return new Response(JSON.stringify({ error: "Method Not Allowed" }), { status: 405 });
        }

        return route.handler(req, path);
    }

    Logger.warn(`Social route not found`, { path });
    return new Response(JSON.stringify({ error: "Route not found" }), { status: 404 });
}

async function proxyToSocialService(req: Request, path: string) {
    const token = req.headers.get("Authorization")?.split("Bearer ")[1];
    if (!token) {
        return new Response(JSON.stringify({ error: "No token provided" }), { status: 401 });
    }

    const userID = await verifyToken(token);
    console.log(userID);
    if (!userID) {
        return new Response(JSON.stringify({ error: "Invalid token" }), { status: 401 });
    }

    const url = new URL(req.url);
    url.searchParams.set("userID", userID);

    const socialReq = new Request(`${socialServiceURL}${path}${url.search}`, {
        method: req.method,
        headers: req.headers,
        body: req.body,
    });

    const response = await fetch(socialReq);
    const data = await response.json();

    return new Response(JSON.stringify(data), { status: response.status });
}

async function verifyToken(token: string): Promise<string | null> {
    const verifyReq = new Request("http://localhost:3000/api/secure/auth/signin", {
        method: "GET",
        headers: {
            "Authorization": `Bearer ${token}`
        }
    });

    const response = await fetch(verifyReq);
    console.log(response);
    if (response.status !== 200) {
        return null;
    }

    const data = await response.json();
    return data.user_id;
}
