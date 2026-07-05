import type { CreateAccountRequest } from "@/type/create-account-request";
import type { CreateAccountResponse } from "@/type/create-account-response";
import type { GetAccountResponse } from "@/type/get-account-response";

export interface AccountGateway {
  create(request: CreateAccountRequest): Promise<CreateAccountResponse>;
  getByID(id: string): Promise<GetAccountResponse>;
}
