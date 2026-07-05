import type { Validator } from "./validator";

export class RequiredValidator implements Validator {
  validate(value: any): string | null {
    if (!value) return "This field is required";
    return null;
  }
}
