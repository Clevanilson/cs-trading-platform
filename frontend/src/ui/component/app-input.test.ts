import { describe, test, expect, beforeEach } from "vitest";
import { mount } from "@vue/test-utils";

import AppInput from "./app-input.vue";
import { FormField } from "@/entity/form-field.ts";

describe("AppInput", () => {
  let field: FormField<string>;

  test("renders label and input", () => {
    const wrapper = mount(AppInput, { props: { field } });
    expect(wrapper.get("label").text()).toBe(field.label);
    expect(wrapper.get("label").attributes("for")).toBe(field.id);
    expect(wrapper.get("input").attributes("id")).toBe(field.id);
    expect(wrapper.get("input").attributes("type")).toBe(field.type);
    expect(wrapper.get("input").attributes("name")).toBe(field.id);
    expect(wrapper.get("input").attributes("placeholder")).toBe(
      field.placeholder,
    );
    expect(wrapper.get("input").attributes("aria-invalid")).toBe(
      Boolean(field.error) ? "true" : "false",
    );
    expect(wrapper.get("input").attributes("aria-describedby")).toBe(
      field.error ? `${field.id}-error` : undefined,
    );
    expect(wrapper.get("input").element.value).toBe(field.value);
  });

  test("binds model value", async () => {
    const wrapper = mount(AppInput, { props: { field } });
    expect(wrapper.get("input").element.value).toBe(field.value);
    await wrapper.get("input").setValue("John Doe");
    expect(field.value).toBe("John Doe");
  });

  test("without error", () => {
    const wrapper = mount(AppInput, { props: { field } });
    expect(wrapper.get("input").attributes("aria-invalid")).toBe("false");
    expect(wrapper.get("input").attributes("aria-describedby")).toBeUndefined();
    expect(wrapper.find(".app-input-error").exists()).toBe(false);
  });

  test("applies app-input class", () => {
    const wrapper = mount(AppInput, { props: { field } });
    expect(wrapper.get("input").classes()).toContain("app-input");
  });

  beforeEach(() => {
    field = new FormField<string>({
      id: "name",
      label: "Full name",
      value: "",
      placeholder: "Name",
    });
  });
});
