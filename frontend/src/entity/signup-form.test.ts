import { describe, test, expect } from "vitest";
import { SignupForm } from "./signup-form";

describe(SignupForm.name, () => {
  test("Submitting the form", () => {
    const form = new SignupForm();
    form.fields.name.value = "John Doe";
    expect(form.submit()).toEqual({
      name: "John Doe",
    });
  });
});
