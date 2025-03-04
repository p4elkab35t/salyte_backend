import { serve } from "bun";
import { WebSocketServer } from "ws";
import messageGRPCClient from "../grpc_client/message.client";

const wss = new WebSocketServer({ port: 3003 });

export default wss;
