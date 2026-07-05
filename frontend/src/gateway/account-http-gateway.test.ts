import { describe, test, expect, beforeEach } from "vitest";
import { AccountGateway } from "./account-http-gateway";
import type { CreateAccountResponse } from "@/type/create-account-response";
import { HttpClientMock } from "@/http/http-client-mock";
import type { CreateAccountRequest } from "@/type/create-account-request";
import type { GetAccountResponse } from "@/type/get-account-response";

describe(AccountGateway.name, () => {
  let httpClient: HttpClientMock;
  let sut: AccountGateway;

  test("calling create", async () => {
    const request: CreateAccountRequest = {
      name: "John Doe",
    };
    const response: CreateAccountResponse = {
      id: "uuid",
    };
    httpClient.withResponse(response);
    const output = await sut.create(request);
    expect(output).toEqual(response);
  });

  test("Calling getByID", async () => {
    const id = "uuid";
    const response: GetAccountResponse = {
      id,
      name: "John Doe",
    };
    httpClient.withResponse(response);
    const output = await sut.getByID(id);
    expect(output).toEqual(response);
  });

  beforeEach(() => {
    httpClient = new HttpClientMock();
    sut = new AccountGateway(httpClient);
  });
});
