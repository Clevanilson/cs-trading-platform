import { describe, test, expect } from "vitest";
import { Form } from "./form";
import { FormField } from "./form-field";
import { RequiredValidator } from "@/validator/required";

describe(Form.name, () => {
  test("Creating a form", () => {
    const fields = {
      name: new FormField<string>({
        id: "name",
        label: "Name",
        value: "John Doe",
        placeholder: "Enter your name",
      }),
    };
    const form = new Form(fields);
    expect(form.fields).toBe(fields);
    expect(form.invalid).toBe(false);
  });

  test("Invalidating a form", () => {
    const fields = {
      name: new FormField<string>({
        id: "name",
        label: "Name",
        value: "John Doe",
        placeholder: "Enter your name",
        validators: [new RequiredValidator()],
      }),
    };
    const form = new Form(fields);
    form.fields.name.value = "";
    expect(form.invalid).toBe(true);
  });
});
