import messageGRPCClient from "../../grpc_client/message.client";
import type { MessagingService } from "../../grpc_client/proto/message.proto";
import * as messageGRPCTypes from "../../grpc_client/proto/message.proto";
import { Logger } from "../../logger/logger";

const messageGRPC: MessagingService = messageGRPCClient;

const methodMap: Record<string, Function> = {
    get: handleGetMessage,
    send: handleSendMessage,
    edit: handleEditMessage,
    delete: handleDeleteMessage,
    reactionApply: handleReactionApply,
    reactionRemove: handleReactionRemove,
    read: handleReadMessage,
};

// Store connection data
const connections = new Map();
const chatRooms = new Map<string, Set<any>>(); // Stores connected clients per chat

// Helper functions for chat room management
function addToChatRoom(chat_id: string, ws: any) {
    if (!chatRooms.has(chat_id)) {
        chatRooms.set(chat_id, new Set());
    }
    chatRooms.get(chat_id)?.add(ws);
}

function removeFromChatRoom(chat_id: string, ws: any) {
    chatRooms.get(chat_id)?.delete(ws);
    if (chatRooms.get(chat_id)?.size === 0) {
        chatRooms.delete(chat_id);
    }
}

// WebSocket handlers for Bun API
async function open(ws: any) {
    try {
        Logger.info("WebSocket connection attempt");
        
        console.log(ws);

        const chat_id = ws.data.get('chatID');
        const token = ws.data.get('token')?.replace("Bearer ", "");

        Logger.info(`Chat ID: ${chat_id}, Token: ${token}`);
        
        if (!chat_id) {
            ws.send(JSON.stringify({ type: "error", message: "No chat_id provided" }));
            ws.close(1008);
            return;
        }
        
        if (!token) {
            ws.send(JSON.stringify({ type: "error", message: "Unauthorized: No token provided" }));
            ws.close(1008);
            return;
        }

        const authResponse = await fetch("http://127.0.0.1:3000/api/secure/auth/signin", {
            method: "GET",
            headers: { Authorization: `Bearer ${token}` }
        });
        
        if (!authResponse.ok) {
            ws.send(JSON.stringify({ type: "error", message: "Unauthorized" }));
            ws.close(1008);
            return;
        }

        const userID = (await authResponse.json()).user_id;
        if (!userID) {
            ws.send(JSON.stringify({ type: "error", message: "Unauthorized" }));
            ws.close(1008);
            return;
        }

        // Store connection data
        connections.set(ws, { userID, chat_id });
        addToChatRoom(chat_id, ws);

        ws.send(JSON.stringify({ type: "connected", message: "Chat connected" }));
        Logger.info(`WebSocket connection opened for user ${userID} in chat ${chat_id}`);
    } catch (error) {
        Logger.error("Error in WebSocket open handler", { error });
        ws.send(JSON.stringify({ type: "error", message: "Server error" }));
        ws.close(1011);
    }
}

async function message(ws: any, message: string | Uint8Array) {
    try {
        const connectionData = connections.get(ws);
        if (!connectionData) {
            ws.send(JSON.stringify({ type: "error", message: "No connection data" }));
            return;
        }
        
        const { userID, chat_id } = connectionData;
        const data = JSON.parse(message.toString());
        
        const handler = methodMap[data.type];
        if (handler) {
            await handler(ws, userID, chat_id, data);
        } else {
            ws.send(JSON.stringify({ type: "error", message: "Unknown event type" }));
        }
    } catch (error) {
        Logger.error("Error in WebSocket message handler", { error });
        ws.send(JSON.stringify({ type: "error", message: "Failed to process message" }));
    }
}

function close(ws: any) {
    try {
        const connectionData = connections.get(ws);
        if (connectionData) {
            const { userID, chat_id } = connectionData;
            removeFromChatRoom(chat_id, ws);
            connections.delete(ws);
            Logger.info(`WebSocket connection closed for user ${userID} in chat ${chat_id}`);
        }
    } catch (error) {
        Logger.error("Error in WebSocket close handler", { error });
    }
}

// Message handlers
async function handleGetMessage(ws: any, userID: string, chat_id: string, data: { message_id: string }) {
    const request: messageGRPCTypes.GetMessageByIDRequest = { user_id: userID, message_id: data.message_id };
    const result: messageGRPCTypes.GetMessageByIDResponse = await promisifyCallback(messageGRPC.GetMessageByID, request);
    ws.send(JSON.stringify(result.status === 0 ? { type: "success", message: result.message } : { type: "error", message: "Failed to get message" }));
}

interface SendMessageData {
    chat_id: string;
    message: string;
}

async function handleSendMessage(ws: any, userID: string, chat_id: string, data: any) {
    const request: messageGRPCTypes.SendMessageRequest = { 
        sender_id: userID, 
        chat_id: data.chat_id || chat_id, 
        content: data.message 
    };
    const result: messageGRPCTypes.SendMessageResponse = await promisifyCallback(messageGRPC.SendMessage, request);

    const messageData = { type: "newMessage", message: result.message };
    chatRooms.get(chat_id)?.forEach(client => {
        client.send(JSON.stringify(messageData));
    });

    ws.send(JSON.stringify(result.status === 0 ? 
        { type: "success", message: "Message sent" } : 
        { type: "error", message: "Failed to send message" }));
}

interface EditMessageData {
    message_id: string;
    message: string;
}

async function handleEditMessage(ws: WebSocket, userID: string, chat_id: string, data: EditMessageData) {
    const request: messageGRPCTypes.EditMessageRequest = { user_id: userID, message_id: data.message_id, new_content: data.message };
    const result: messageGRPCTypes.EditMessageResponse = await promisifyCallback(messageGRPC.EditMessage, request);
    
    const messageData = { type: "editedMessage", message: result.message };
    chatRooms.get(chat_id)?.forEach(client => {
        if (client.readyState === WebSocket.OPEN) {
            client.send(JSON.stringify(messageData));
        }
    });

    ws.send(JSON.stringify(result.status === 0 ? { type: "success", message: "Message edited" } : { type: "error", message: "Failed to edit message" }));
}

interface DeleteMessageData {
    message_id: string;
}

async function handleDeleteMessage(ws: WebSocket, userID: string, chat_id: string, data: DeleteMessageData) {
    const request: messageGRPCTypes.DeleteMessageRequest = { user_id: userID, message_id: data.message_id };
    const result: messageGRPCTypes.DeleteMessageResponse = await promisifyCallback(messageGRPC.DeleteMessage, request);
    
    const messageData = { type: "deletedMessage", message_id: data.message_id };
    chatRooms.get(chat_id)?.forEach(client => {
        if (client.readyState === WebSocket.OPEN) {
            client.send(JSON.stringify(messageData));
        }
    });
    
    ws.send(JSON.stringify(result.success ? { type: "success", message: "Message deleted" } : { type: "error", message: "Failed to delete message" }));
}

interface ReactionApplyData {
    message_id: string;
    reaction: string;
}

async function handleReactionApply(ws: WebSocket, userID: string, chat_id: string, data: ReactionApplyData) {
    const request: messageGRPCTypes.AddReactionRequest = { user_id: userID, message_id: data.message_id, reaction: data.reaction };
    const result: messageGRPCTypes.AddReactionResponse = await promisifyCallback(messageGRPC.AddReaction, request);
    
    const messageData = { type: "reactionApplied", message_id: data.message_id, reaction: data.reaction };
    chatRooms.get(chat_id)?.forEach(client => {
        if (client.readyState === WebSocket.OPEN) {
            client.send(JSON.stringify(messageData));
        }
    });
    
    ws.send(JSON.stringify(result.success ? { type: "success", message: "Reaction applied" } : { type: "error", message: "Failed to apply reaction" }));
}

interface ReactionRemoveData {
    message_id: string;
    reaction: string;
}

async function handleReactionRemove(ws: WebSocket, userID: string, chat_id: string, data: ReactionRemoveData) {
    const request: messageGRPCTypes.RemoveReactionRequest = { user_id: userID, message_id: data.message_id, reaction: data.reaction };
    const result: messageGRPCTypes.RemoveReactionResponse = await promisifyCallback(messageGRPC.RemoveReaction, request);
    
    const messageData = { type: "reactionRemoved", message_id: data.message_id, reaction: data.reaction };
    chatRooms.get(chat_id)?.forEach(client => {
        if (client.readyState === WebSocket.OPEN) {
            client.send(JSON.stringify(messageData));
        }
    });
    
    ws.send(JSON.stringify(result.success ? { type: "success", message: "Reaction removed" } : { type: "error", message: "Failed to remove reaction" }));
}

interface ReadMessageData {
    message_id: string;
}

async function handleReadMessage(ws: WebSocket, userID: string, chat_id: string, data: ReadMessageData) {
    const request: messageGRPCTypes.ReadMessageRequest = { user_id: userID, message_id: data.message_id };
    const result: messageGRPCTypes.ReadMessageResponse = await promisifyCallback(messageGRPC.ReadMessage, request);
    
    const messageData = { type: "messageRead", message_id: data.message_id };
    chatRooms.get(chat_id)?.forEach(client => {
        if (client.readyState === WebSocket.OPEN) {
            client.send(JSON.stringify(messageData));
        }
    });

    ws.send(JSON.stringify(result.success ? { type: "success", message: "Message marked as read" } : { type: "error", message: "Failed to mark message as read" }));
}

const promisifyCallback = (method: Function, request: any): Promise<any> => {
    return new Promise((resolve, reject) => {
        method.call(messageGRPC, request, (error: any, response: any) => {
            if (error) {
                console.error("gRPC Error:", error);
                resolve({ status: 1, success: false });
            } else {
                resolve(response);
            }
        });
    });
};

export default {
    open,
    message,
    close
};