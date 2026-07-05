import { describe, test, expect, beforeEach } from "vitest";
import { mount, VueWrapper } from "@vue/test-utils";
import SignupPage from "./signup-page.vue";
import { AccountMockGateway } from "@/gateway/account-mock-gateway";

describe("SignupPage", () => {
  let accountGateway: AccountMockGateway;
  let sut: SignupPagePOM;

  beforeEach(() => {
    accountGateway = new AccountMockGateway();
    sut = new SignupPagePOM(accountGateway);
  });

  test("Rendering the signup page", () => {
    expect(sut.title.text()).toBe("Create your account");
    expect(sut.subtitle.text()).toBe(
      "Start a new session on the CS Trading Platform.",
    );
    expect(sut.nameInput.attributes("placeholder")).toBe("Enter your name");
    expect(sut.submitButton.text()).toBe("Create account");
  });

  test("With invalid form", () => {
    expect(sut.submitButton.attributes("disabled")).toBeDefined();
  });

  test("Submit with valid form", async () => {
    const name = "John Doe";
    await sut.nameInput.setValue(name);
    expect(sut.submitButton.attributes("disabled")).not.toBeDefined();
    await sut.submitButton.trigger("submit.prevent");
    expect(accountGateway.create).lastCalledWith({ name });
  });
});

class SignupPagePOM {
  private readonly wrapper: VueWrapper;

  constructor(gateway: AccountMockGateway) {
    this.wrapper = mount(SignupPage, {
      global: {
        provide: {
          accountGateway: gateway,
        },
      },
    });
  }

  get nameInput() {
    return this.wrapper.get("input#name");
  }

  get submitButton() {
    return this.wrapper.get("button");
  }

  get title() {
    return this.wrapper.get(".signup-title");
  }

  get subtitle() {
    return this.wrapper.get(".signup-subtitle");
  }
}
