
export interface GetSecurityLogsByUserIDRequest {
    user_id: string;
    page: number;
    limit: number;
}

export interface GetSecurityLogsByUserIDResponse {
    security_logs: SecurityLog[];
    status: number;
}

export interface GetSecurityLogWithIDRequest {
    log_id: string;
}

export interface GetSecurityLogWithIDResponse {
    security_log: SecurityLog;
    status: number;
}

export interface SecurityLog {
    log_id: string;
    user_id: string;
    action: string;
    timestamp: string;
}

export interface defaultCallback {
    (error: any, response: any): any;
}

export interface SecurityLogsService {
    GetSecurityLogsWithUsedID(request: GetSecurityLogsByUserIDRequest, callBack: defaultCallback): Promise<GetSecurityLogsByUserIDResponse>;
    GetSecurityLogWithID(request: GetSecurityLogWithIDRequest, callBack: defaultCallback): Promise<GetSecurityLogWithIDResponse>;
}
