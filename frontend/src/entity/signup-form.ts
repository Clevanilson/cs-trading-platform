import type { CreateAccountRequest } from "@/type/create-account-request";
import { Form } from "./form";
import { FormField } from "./form-field";
import { RequiredValidator } from "@/validator/required";
import { NameValidator } from "@/validator/name";

export class SignupForm extends Form<SignupFormFields> {
  constructor() {
    super({
      name: new FormField<string>({
        id: "name",
        label: "Name",
        value: "",
        placeholder: "Enter your name",
        validators: [new RequiredValidator(), new NameValidator()],
      }),
    });
  }

  submit(): CreateAccountRequest {
    return {
      name: this.fields.name.value,
    };
  }
}

export type SignupFormFields = {
  name: FormField<string>;
};
