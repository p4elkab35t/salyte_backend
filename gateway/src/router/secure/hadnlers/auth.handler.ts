import { Logger } from "../../../logger/logger";
import authGRPCClient from "../../../grpc_client/auth.client";
import type { AuthService } from "../../../grpc_client/proto/auth.proto";

// Type the client object as AuthService
const authClient: AuthService = authGRPCClient;

const routes = new Map<string, { method: string[], handler: Function }>([
    ["/lifecheck", { method: ["GET"], handler: lifeCheck }],
    ["/verify", { method: ["GET"], handler: verifyToken }],
    ["/signin", { method: ["GET", "POST"], handler: signIn }],
    ["/signup", { method: ["POST"], handler: signUp }],
    ["/signout", { method: ["GET"], handler: signOut }],
  ]);

  export default async function authServiceHandler(req: Request, restPath: string) {
    const url = new URL(req.url);
    const path = restPath.replace("/secure/auth", "");
  
    Logger.info(`Auth request`, { method: req.method, path });
  
    if (routes.has(path)) {
      const route = routes.get(path);
      if (!route?.method.includes(req.method)) {
        Logger.warn(`Method not allowed`, { path, method: req.method });
        return new Response(JSON.stringify({ error: "Method Not Allowed" }), { status: 405 });
      }
  
      return route.handler(req);
    }
  
    Logger.warn(`Auth route not found`, { path });
    return new Response(JSON.stringify({ error: "Route not found" }), { status: 404 });
  }

  async function lifeCheck(req: Request, restPath: string) {
    Logger.info("Life check", { method: req.method });
    return new Response(JSON.stringify({ message: "Auth Service Alive" }), { status: 200 });
  }
  
  async function verifyToken(req: Request) {
    Logger.info("Verifying token", { method: req.method });
    try{
        
    const token = req.headers.get("Authorization")?.split("Bearer ")[1];
    if (!token) {
      return new Response(JSON.stringify({ error: "No token provided" }), { status: 401 });
    }
    // get user id from quuery params
    const URLQuery = new URL(req.url);
    const userID = URLQuery.searchParams.get("user_id");
    if (!userID) {
      return new Response(JSON.stringify({ error: "No user id provided" }), { status: 401 });
    }
    const result = await promisifyCallback(authClient.VerifyToken,{ token: token, user_id: userID });
    if (result.status !== 0){
      return new Response(JSON.stringify({ error: "Verification failed" }), { status: 401 });
    }
    Logger.info("Token verified", { result });
    return new Response(JSON.stringify(result), { status: 200 });
  } catch (error) {
    return new Response(JSON.stringify({ error: "Verification failed" }), { status: 401 });
  }
  }

  async function signIn(req: Request) {
    Logger.info("Signing in user");
    try{
    if (req.method === "GET") {
      const token = req.headers.get("Authorization")?.split("Bearer ")[1];
      if (!token) {
        return new Response(JSON.stringify({ error: "No token provided" }), { status: 401 });
      }
      const result = await promisifyCallback(authClient.SignInToken,{ token: token });
      if (result.status !== 0){
        return new Response(JSON.stringify({ error: "Sign in failed" }), { status: 401 });
      }
      Logger.info("User signed in", { result });
      return new Response(JSON.stringify(result), { status: 200 });
    }
    else if (req.method === "POST") {
      const body = await req.json();
      const result = await promisifyCallback(authClient.SignInCredentials,{ email: body.email, password: body.password });
      if (result.status !== 0){
        return new Response(JSON.stringify({ error: "Sign in failed" }), { status: 401 });
      }
      Logger.info("User signed in", { result });
      return new Response(JSON.stringify(result), { status: 200 });
    }
    else {
      return new Response(JSON.stringify({ error: "Method not allowed" }), { status: 405 });
    }
  } catch (error) {
    return new Response(JSON.stringify({ error: "Sign in failed" }), { status: 401 });
  }
  }

  async function signUp(req: Request) {
    Logger.info("Registering new user");
    try{
    const body = await req.json();
    const result = await promisifyCallback(authClient.SignUp,{ email: body.email, password: body.password });
    if (result.status !== 0){
      return new Response(JSON.stringify({ error: "Sign up failed" }), { status: 401 });
    }
    Logger.info("User signed up", { result });
    return new Response(JSON.stringify(result), { status: 200 });
  } catch (error) {
    return new Response(JSON.stringify({ error: "Sign up failed" }), { status: 401 });
  }
  }

  async function signOut(req: Request) {
    Logger.info("Signing out user");
    try{
    const token = req.headers.get("Authorization")?.split("Bearer ")[1];
    if (!token) {
      return new Response(JSON.stringify({ error: "No token provided" }), { status: 401 });
    }
    const result = await promisifyCallback(authClient.SignOut,{ token: token });
    if (result.status !== 0){
      return new Response(JSON.stringify({ error: "Sign out failed" }), { status: 401 });
    }
    Logger.info("User signed out", { result });
    return new Response(JSON.stringify(result), { status: 200 });
  } catch (error) {
    return new Response(JSON.stringify({ error: "Sign out failed" }), { status: 401 });
  }
  }

  const promisifyCallback = (method: Function, request: any): Promise<any> => {
    return new Promise((resolve, reject) => {
      method.call(authClient, request, (error: any, response: any) => {
        if (error) {
          reject(error); 
        } else {
          resolve(response); 
        }
      });
    });
  }

