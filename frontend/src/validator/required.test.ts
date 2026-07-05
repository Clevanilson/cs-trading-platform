import { describe, test, expect } from "vitest";
import { RequiredValidator } from "./required";

describe(RequiredValidator.name, () => {
  test.each(["Jhon", true, true, 1])("With valid value", (value) => {
    const sut = new RequiredValidator();
    const result = sut.validate(value);
    expect(result).toBe(null);
  });

  test.each(["", undefined, null, 0, false])("With invalid value", (value) => {
    const sut = new RequiredValidator();
    expect(sut.validate(value)).toBe("This field is required");
  });
});
