import type { HttpClient } from "@/http/http-client";
import type { CreateAccountRequest } from "@/type/create-account-request";
import type { CreateAccountResponse } from "@/type/create-account-response";
import type { GetAccountResponse } from "@/type/get-account-response";

export class AccountGateway {
  private readonly httpClient: HttpClient;

  constructor(httpClient: HttpClient) {
    this.httpClient = httpClient;
  }

  create(req: CreateAccountRequest): Promise<CreateAccountResponse> {
    return this.httpClient.post("/account/signup", req);
  }

  getByID(id: string): Promise<GetAccountResponse> {
    return this.httpClient.get(`/account/get_account/${id}`);
  }
}
