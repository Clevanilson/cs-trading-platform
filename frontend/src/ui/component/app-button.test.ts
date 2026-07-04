import { describe, test, expect } from "vitest";
import { mount } from "@vue/test-utils";

import AppButton from "./app-button.vue";

describe("AppButton", () => {
  test("With slot content", () => {
    const wrapper = mount(AppButton, { slots: { default: "Create account" } });
    expect(wrapper.text()).toBe("Create account");
  });

  test("With default type", () => {
    const wrapper = mount(AppButton);
    expect(wrapper.get(".app-button").attributes("type")).toBe("button");
  });

  test("With custom type", () => {
    const type = "submit";
    const wrapper = mount(AppButton, { props: { type } });
    expect(wrapper.get(".app-button").attributes("type")).toBe(type);
  });

  test("With default disabled", () => {
    const wrapper = mount(AppButton);
    expect(wrapper.get(".app-button").attributes("disabled")).toBeUndefined();
  });

  test("With disabled true", () => {
    const wrapper = mount(AppButton, { props: { disabled: true } });
    expect(wrapper.get(".app-button").attributes("disabled")).toBeDefined();
  });

  test("With loading true", () => {
    const wrapper = mount(AppButton, {
      props: { loading: true },
      slots: { default: "Create account" },
    });
    expect(wrapper.get(".app-button").attributes("disabled")).toBeDefined();
    expect(wrapper.text()).toBe("Loading…");
    expect(wrapper.text()).not.toContain("Create account");
  });

  test("With custom loading text", () => {
    const loadingText = "Creating account…";
    const wrapper = mount(AppButton, {
      props: { loading: true, loadingText },
      slots: { default: "Create account" },
    });
    expect(wrapper.get("button").text()).toBe("Creating account…");
  });

  test("applies app-button class", () => {
    const wrapper = mount(AppButton);
    expect(wrapper.get("button").classes()).toContain("app-button");
  });
});
