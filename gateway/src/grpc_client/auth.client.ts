import grpc from '@grpc/grpc-js';
import protoLoader from '@grpc/proto-loader';
import { CONFIG } from '../config/config';

const packageDefinition = protoLoader.loadSync(
  `${__dirname}/proto/auth.proto`,
  {
    keepCase: true,
    longs: String,
    enums: String,
    defaults: true,
    oneofs: true,
  }
);

const protoDescriptor = grpc.loadPackageDefinition(packageDefinition) as any;

if (!protoDescriptor || !protoDescriptor.auth_service) {
  throw new Error('Failed to load auth service definition from proto file');
}

const auth = protoDescriptor.auth_service;

const client = new auth.AuthService(
  CONFIG.AUTH_SERVICE_URL,
  grpc.credentials.createInsecure()
);

export default client;
