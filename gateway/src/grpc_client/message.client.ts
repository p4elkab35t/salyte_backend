import grpc from '@grpc/grpc-js';
import protoLoader from '@grpc/proto-loader';
import { CONFIG } from '../config/config';

const packageDefinition = protoLoader.loadSync(
  `${__dirname}/proto/message.proto`,
  {
    keepCase: true,
    longs: String,
    enums: String,
    defaults: true,
    oneofs: true,
  }
);

const protoDescriptor = grpc.loadPackageDefinition(packageDefinition) as any;

if (!protoDescriptor || !protoDescriptor.message_service) {
  throw new Error('Failed to load auth service definition from proto file');
}

const message = protoDescriptor.message_service;

const messageGRPCClient = new message.MessagingService(
  CONFIG.MESSAGE_GRPC_SERVICE_URL,
  grpc.credentials.createInsecure()
);

export default messageGRPCClient;
