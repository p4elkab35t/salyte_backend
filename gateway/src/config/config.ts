import { config } from "dotenv";
config();

export const CONFIG = {
  PORT: process.env.PORT || 3000,
  AUTH_SERVICE_URL: process.env.AUTH_SERVICE_URL || "127.0.0.1:50051",
  MESSAGE_GRPC_SERVICE_URL: process.env.MESSAGE_SERVICE_URL || "127.0.0.1:50052",
  SOCIAL_SERVICE_URL: process.env.SOCIAL_SERVICE_URL || "127.0.0.1:8081",
  MESSAGE_REST_SERVICE_URL: process.env.MESSAGE_REST_SERVICE_URL || "127.0.0.1:8083",
};
