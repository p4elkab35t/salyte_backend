export interface SignInCredentialsRequest {
    email: string;
    password: string;
}

export interface SignInTokenRequest {
    token: string;
}

export interface SignInResponse {
    token: string;
    user_id: string;
    status: number;
}

export interface SignUpRequest {
    email: string;
    password: string;
}

export interface SignUpResponse {
    user_id: string;
    token: string;
    status: number;
}

export interface VerifyTokenRequest {
    token: string;
    user_id: string;
}

export interface VerifyTokenResponse {
    is_valid: boolean;
    status: number;
}

export interface SignOutRequest {
    token: string;
}

export interface SignOutResponse {
    status: number;
}

export interface defaultCallback {
    (error: any, response: any): any;
}

export interface AuthService {
    SignInCredentials(request: SignInCredentialsRequest, callBack: defaultCallback): Promise<SignInResponse>;
    SignInToken(request: SignInTokenRequest): Promise<SignInResponse>;
    SignUp(request: SignUpRequest, callBack: defaultCallback): Promise<SignUpResponse>;
    VerifyToken(request: VerifyTokenRequest, callBack: defaultCallback): Promise<VerifyTokenResponse>;
    SignOut(request: SignOutRequest, callBack: defaultCallback): Promise<SignOutResponse>;
}
