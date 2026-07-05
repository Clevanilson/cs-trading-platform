export interface Validator {
  validate(value: any): string | null;
}
