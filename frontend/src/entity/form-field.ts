import type { Validator } from "@/validator/validator";

export class FormField<TValue> {
  private _id: string;
  private _label: string;
  private _placeholder: string;
  private _type: FormControlType;
  private _validators: Validator[];
  private _error: string | null;
  private _touched: boolean;
  private _value: TValue;

  constructor(options: FormControlOptions<TValue>) {
    this._id = options.id;
    this._label = options.label;
    this._placeholder = options.placeholder;
    this._type = options.type ?? "text";
    this._validators = options.validators ?? [];
    this._value = options.value;
    this._error = null;
    this._touched = false;
    this.validate();
  }

  get id(): string {
    return this._id;
  }

  get label(): string {
    return this._label;
  }

  get placeholder(): string {
    return this._placeholder;
  }

  get type(): FormControlType {
    return this._type;
  }

  get value(): TValue {
    return this._value;
  }

  set value(value: TValue) {
    this._touched = true;
    this._value = value;
    this.validate();
  }

  get error(): string | null {
    return this._error;
  }

  get invalid(): boolean {
    return Boolean(this._error);
  }

  get touched(): boolean {
    return this._touched;
  }

  private validate(): void {
    for (const validator of this._validators) {
      this._error = validator.validate(this._value);
      if (this._error) break;
    }
  }
}

export type FormControlOptions<TValue> = {
  id: string;
  label: string;
  placeholder: string;
  type?: FormControlType;
  value: TValue;
  validators?: Validator[];
};

export type FormControlType = "text" | "password" | "number";
