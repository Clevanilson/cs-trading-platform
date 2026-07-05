import { describe, test, expect, beforeEach } from "vitest";
import { FormField, type FormControlOptions } from "./form-field";
import { RequiredValidator } from "@/validator/required";

describe(FormField.name, () => {
  let options: FormControlOptions<string>;

  test("Creating a form field", () => {
    const sut = new FormField<string>(options);
    expect(sut.id).toBe("name");
    expect(sut.label).toBe("Name");
    expect(sut.placeholder).toBe("Enter your name");
    expect(sut.value).toBe("John Doe");
    expect(sut.type).toBe("password");
    expect(sut.error).toBeNull();
    expect(sut.invalid).toBe(false);
    expect(sut.touched).toBe(false);
  });

  test("With default type", () => {
    const field = new FormField<string>({
      ...options,
      type: undefined,
    });
    expect(field.type).toBe("text");
  });

  test("Setting the value", () => {
    const sut = new FormField<string>(options);
    sut.value = "Jane Doe";
    expect(sut.value).toBe("Jane Doe");
    expect(sut.touched).toBe(true);
  });

  test("With validators", () => {
    const sut = new FormField<string>({
      ...options,
      value: "",
      validators: [new RequiredValidator()],
    });
    expect(sut.invalid).toBe(true);
    expect(sut.error).toBe("This field is required");
  });

  beforeEach(() => {
    options = {
      id: "name",
      label: "Name",
      value: "John Doe",
      placeholder: "Enter your name",
      type: "password",
    };
  });
});
