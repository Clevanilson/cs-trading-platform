import type { FormField } from "./form-field";

export class Form<TFields extends Record<string, FormField<any>>> {
  private _fields: TFields;

  constructor(fields: TFields) {
    this._fields = fields;
  }

  get fields(): TFields {
    return this._fields;
  }

  get invalid(): boolean {
    return Object.values(this.fields).some((field) => field.invalid);
  }
}
