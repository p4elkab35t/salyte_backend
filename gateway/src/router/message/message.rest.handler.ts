// In Bun, we use the global fetch and Request/Response objects.
import { Logger } from "../../logger/logger";
  
// Handler functions
async function handleGetMessagesByChatID(req: Request): Promise<Response> {
  const { searchParams } = new URL(req.url);
  const chat_id = searchParams.get("chat_id") || "";
  const user_id = searchParams.get("user_id") || "";
  const response = await fetch(`http://localhost:8083/message/getallbychat?chatID=${chat_id}&userID=${user_id}`, {
      method: "GET",
      headers: { "Content-Type": "application/json" }
  });
  const result = await response.json();
  return new Response(JSON.stringify(result), { headers: { "Content-Type": "application/json" } });
}

async function handleGetUnreadMessages(req: Request): Promise<Response> {
  const { searchParams } = new URL(req.url);
  const user_id = searchParams.get("user_id") || "";
  const response = await fetch(`http://localhost:8083/message/unread?userID=${user_id}`, {
      method: "GET",
      headers: { "Content-Type": "application/json" }
  });
  const result = await response.json();
  return new Response(JSON.stringify(result), { headers: { "Content-Type": "application/json" } });
}

async function handleDeleteAllMessagesByChatID(req: Request): Promise<Response> {
  const { searchParams } = new URL(req.url);
  const chat_id = searchParams.get("chat_id") || "";
  const user_id = searchParams.get("user_id") || "";
  const body = await req.json();
  const response = await fetch(`http://localhost:8083/message/deleteall?chatID=${chat_id}&userID=${user_id}`, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(body)
  });
  const result = await response.json();
  return new Response(JSON.stringify(result), { headers: { "Content-Type": "application/json" } });
}

async function handleGetChat(req: Request): Promise<Response> {
  const { searchParams } = new URL(req.url);
  const chat_id = searchParams.get("chat_id") || "";
  const response = await fetch(`http://localhost:8083/chat/get?chatID=${chat_id}`, {
      method: "GET",
      headers: { "Content-Type": "application/json" }
  });
  const result = await response.json();
  return new Response(JSON.stringify(result), { headers: { "Content-Type": "application/json" } });
}

async function handleCreateChat(req: Request): Promise<Response> {
  const body = await req.json();
  const response = await fetch("http://localhost:8083/chat/create", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(body)
  });
  const result = await response.json();
  return new Response(JSON.stringify(result), { headers: { "Content-Type": "application/json" } });
}

async function handleGetAllChats(req: Request): Promise<Response> {
  const { searchParams } = new URL(req.url);
  const user_id = searchParams.get("user_id") || "";
  const response = await fetch(`http://localhost:8083/chat/getall?userID=${user_id}`, {
      method: "GET",
      headers: { "Content-Type": "application/json" }
  });
  const result = await response.json();
  return new Response(JSON.stringify(result), { headers: { "Content-Type": "application/json" } });
}

async function handleAddUserToChat(req: Request): Promise<Response> {
  const { searchParams } = new URL(req.url);
  const chat_id = searchParams.get("chat_id") || "";
  const user_id = searchParams.get("user_id") || "";
  const added_user_id = searchParams.get("added_user_id") || "";
  const body = await req.json();
  const response = await fetch(`http://localhost:8083/chat/adduser?chatID=${chat_id}&userID=${user_id}&addedUserID=${added_user_id}`, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(body)
  });
  const result = await response.json();
  return new Response(JSON.stringify(result), { headers: { "Content-Type": "application/json" } });
}

async function handleRemoveUserFromChat(req: Request): Promise<Response> {
  const { searchParams } = new URL(req.url);
  const chat_id = searchParams.get("chat_id") || "";
  const user_id = searchParams.get("user_id") || "";
  const removed_user_id = searchParams.get("removed_user_id") || "";
  const body = await req.json();
  const response = await fetch(`http://localhost:8083/chat/removeuser?chatID=${chat_id}&userID=${user_id}&removedUserID=${removed_user_id}`, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(body)
  });
  const result = await response.json();
  return new Response(JSON.stringify(result), { headers: { "Content-Type": "application/json" } });
}

async function handleGetChatMembers(req: Request): Promise<Response> {
  const { searchParams } = new URL(req.url);
  const chat_id = searchParams.get("chat_id") || "";
  const user_id = searchParams.get("user_id") || "";
  const response = await fetch(`http://localhost:8083/chat/members?chatID=${chat_id}&userID=${user_id}`, {
      method: "GET",
      headers: { "Content-Type": "application/json" }
  });
  const result = await response.json();
  return new Response(JSON.stringify(result), { headers: { "Content-Type": "application/json" } });
}

async function handleGetChatByID(req: Request): Promise<Response> {
  const { searchParams } = new URL(req.url);
  const chat_id = searchParams.get("chat_id") || "";
  const user_id = searchParams.get("user_id") || "";
  const response = await fetch(`http://localhost:8083/chat/messages?chatID=${chat_id}&userID=${user_id}`, {
      method: "GET",
      headers: { "Content-Type": "application/json" }
  });
  const result = await response.json();
  return new Response(JSON.stringify(result), { headers: { "Content-Type": "application/json" } });
}

async function handleGetReactions(req: Request): Promise<Response> {
  const { searchParams } = new URL(req.url);
  const message_id = searchParams.get("message_id") || "";
  const user_id = searchParams.get("user_id") || "";
  const response = await fetch(`http://localhost:8083/reaction/get?messageID=${message_id}&userID=${user_id}`, {
      method: "GET",
      headers: { "Content-Type": "application/json" }
  });
  const result = await response.json();
  return new Response(JSON.stringify(result), { headers: { "Content-Type": "application/json" } });
}

// Map of routes with their allowed methods and handlers
const routes = new Map<string, { method: string[], handler: Function }>([
  ["/getMessagesByChatID", { method: ["GET"], handler: handleGetMessagesByChatID }],
  ["/getUnreadMessages", { method: ["GET"], handler: handleGetUnreadMessages }],
  ["/deleteAllMessagesByChatID", { method: ["POST"], handler: handleDeleteAllMessagesByChatID }],
  ["/getChat", { method: ["GET"], handler: handleGetChat }],
  ["/createChat", { method: ["POST"], handler: handleCreateChat }],
  ["/getAllChats", { method: ["GET"], handler: handleGetAllChats }],
  ["/addUserToChat", { method: ["POST"], handler: handleAddUserToChat }],
  ["/removeUserFromChat", { method: ["POST"], handler: handleRemoveUserFromChat }],
  ["/getChatMembers", { method: ["GET"], handler: handleGetChatMembers }],
  ["/getChatByID", { method: ["GET"], handler: handleGetChatByID }],
  ["/getReactions", { method: ["GET"], handler: handleGetReactions }]
]);

// Message service handler function
export default async function messageServiceHandler(req: Request, restPath: string) {
  const path = restPath.replace("/message", "");
  
  Logger.info(`Message request`, { method: req.method, path });

  if (routes.has(path)) {
    const route = routes.get(path);
    if (!route?.method.includes(req.method)) {
      Logger.warn(`Method not allowed`, { path, method: req.method });
      return new Response(JSON.stringify({ error: "Method Not Allowed" }), { status: 405 });
    }

    return route.handler(req);
  }

  Logger.warn(`Message route not found`, { path });
  return new Response(JSON.stringify({ error: "Route not found" }), { status: 404 });
}

// Keep the original routeMap for backward compatibility
// const routeMap: Record<string, (req: Request) => Promise<Response>> = {
//   "/getMessagesByChatID": handleGetMessagesByChatID,
//   "/getUnreadMessages": handleGetUnreadMessages,
//   "/deleteAllMessagesByChatID": handleDeleteAllMessagesByChatID,
//   "/getChat": handleGetChat,
//   "/createChat": handleCreateChat,
//   "/getAllChats": handleGetAllChats,
//   "/addUserToChat": handleAddUserToChat,
//   "/removeUserFromChat": handleRemoveUserFromChat,
//   "/getChatMembers": handleGetChatMembers,
//   "/getChatByID": handleGetChatByID,
//   "/getReactions": handleGetReactions
// };

// export { routeMap };
