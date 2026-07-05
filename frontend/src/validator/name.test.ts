import { describe, test, expect } from "vitest";
import { NameValidator } from "./name";

describe(NameValidator.name, () => {
  test.each(["John Doe", "Jane Doe", "Doe"])("With valid name", (value) => {
    const sut = new NameValidator();
    expect(sut.validate(value)).toBe(null);
  });

  test("With name too short", () => {
    const sut = new NameValidator();
    expect(sut.validate("a")).toBe("Name must be at least 2 characters long");
  });

  test("With name too long", () => {
    const sut = new NameValidator();
    expect(sut.validate("a".repeat(256))).toBe(
      "Name must be less than 255 characters long",
    );
  });

  test.each(["sa12", "12abasd"])("With invalid name", (value) => {
    const sut = new NameValidator();
    expect(sut.validate(value)).toBe("Invalid name");
  });
});
