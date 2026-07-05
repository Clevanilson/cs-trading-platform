import { vi } from "vitest";
import type { AccountGateway } from "./account-gateway";

export class AccountMockGateway implements AccountGateway {
  create = vi.fn();
  getByID = vi.fn();
}
