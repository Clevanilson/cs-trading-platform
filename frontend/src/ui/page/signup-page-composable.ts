import { SignupForm } from "@/entity/signup-form";
import type { AccountGateway } from "@/gateway/account-gateway";
import { inject, ref } from "vue";

export function useSignupForm() {
  const form = ref<SignupForm>(new SignupForm());
  const isLoading = ref(false);
  const accountGateway = inject<AccountGateway>("accountGateway");
  if (!accountGateway) {
    throw new Error("Account gateway not found");
  }
  const submit = async () => {
    try {
      isLoading.value = true;
      const request = form.value.submit();
      await accountGateway?.create(request);
    } catch (error) {
      console.error(error);
    } finally {
      isLoading.value = false;
    }
  };
  return {
    form,
    isLoading,
    submit,
  };
}
