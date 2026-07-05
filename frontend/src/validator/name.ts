import type { Validator } from "./validator";

export class NameValidator implements Validator {
  validate(value: string): string | null {
    const regex = /^[a-zA-Z ]+$/g;
    const minLength = 2;
    const maxLength = 255;
    if (value.length < minLength) {
      return "Name must be at least 2 characters long";
    }
    if (value.length > maxLength) {
      return "Name must be less than 255 characters long";
    }
    if (!regex.test(value)) {
      return "Invalid name";
    }
    return null;
  }
}
