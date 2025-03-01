import grpc from '@grpc/grpc-js';
import protoLoader from '@grpc/proto-loader';
import { CONFIG } from '../config/config';

const packageDefinition = protoLoader.loadSync(
  `${__dirname}/proto/log.proto`,
  {
    keepCase: true,
    longs: String,
    enums: String,
    defaults: true,
    oneofs: true,
  }
);

const protoDescriptor = grpc.loadPackageDefinition(packageDefinition) as any;

if (!protoDescriptor || !protoDescriptor.security_logs) {
  throw new Error('Failed to load auth service definition from proto file');
}

const securityLogs = protoDescriptor.security_logs;

const securityLogsGRPCClient = new securityLogs.SecurityLogsService(
  CONFIG.AUTH_SERVICE_URL,
  grpc.credentials.createInsecure()
);

export default securityLogsGRPCClient;
