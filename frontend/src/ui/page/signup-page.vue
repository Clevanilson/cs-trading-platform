<script setup lang="ts">
import { useSignupForm } from "@/ui/page/signup-page-composable";
import AppButton from "@/ui/component/app-button.vue";
import AppInput from "@/ui/component/app-input.vue";

const { form, isLoading, submit } = useSignupForm();
</script>

<template>
  <main class="signup-page">
    <section class="signup-card">
      <header class="signup-header">
        <span class="signup-logo">CS</span>
        <h1 class="signup-title">Create your account</h1>
        <p class="signup-subtitle">
          Start a new session on the CS Trading Platform.
        </p>
      </header>

      <form novalidate class="signup-form" @submit.prevent="submit">
        <AppInput :field="form.fields.name" />

        <AppButton
          type="submit"
          :disabled="form.invalid"
          :loading="isLoading"
          loading-text="Creating account…"
        >
          Create account
        </AppButton>
      </form>
    </section>
  </main>
</template>

<style scoped>
.signup-page {
  display: flex;
  min-height: 100vh;
  align-items: center;
  justify-content: center;
  background-color: var(--clr-background-01);
  padding: var(--utl-size-09) var(--utl-size-06);
  color: var(--clr-foreground-01);
}

.signup-card {
  width: 100%;
  max-width: 28rem;
  border-radius: var(--utl-radius-04);
  border: 1px solid var(--clr-border-01);
  background-color: var(--clr-surface-01);
  padding: var(--utl-size-08);
  box-shadow: 0 25px 50px -12px rgba(0, 0, 0, 0.4);
}

.signup-header {
  margin-bottom: var(--utl-size-08);
  text-align: center;
}

.signup-logo {
  display: inline-flex;
  height: 3rem;
  width: 3rem;
  align-items: center;
  justify-content: center;
  margin-bottom: var(--utl-size-06);
  border-radius: var(--utl-radius-02);
  background-color: color-mix(in srgb, var(--clr-primary-01) 10%, transparent);
  font-size: var(--ft-size-04);
  font-weight: var(--ft-weight-04);
  color: var(--clr-primary-01);
}

.signup-title {
  font-size: var(--ft-size-04);
  font-weight: var(--ft-weight-03);
  letter-spacing: -0.025em;
}

.signup-subtitle {
  margin-top: var(--utl-size-02);
  font-size: var(--ft-size-02);
  color: var(--clr-muted-01);
}

.signup-form {
  display: flex;
  flex-direction: column;
  gap: 1.25rem;
}

.signup-error {
  border-radius: var(--utl-radius-01);
  border: 1px solid rgba(239, 68, 68, 0.3);
  background-color: rgba(239, 68, 68, 0.1);
  padding: var(--utl-size-03) var(--utl-size-06);
  font-size: var(--ft-size-02);
  color: var(--clr-danger-02);
}
</style>
