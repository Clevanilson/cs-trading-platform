<script setup lang="ts">
import type { FormField } from "@/entity/form-field";

defineProps<{ field: FormField<any> }>();
</script>

<template>
  <div class="app-input-field">
    <label :for="field.id" class="app-input-label">{{ field.label }}</label>
    <input
      v-model="field.value"
      class="app-input"
      :id="field.id"
      :type="field.type"
      :name="field.id"
      :placeholder="field.placeholder"
      :aria-invalid="Boolean(field.error)"
      :aria-describedby="field.error ? `${field.id}-error` : undefined"
    />
    <p
      v-if="field.error && field.touched"
      :id="`${field.id}-error`"
      class="app-input-error"
    >
      {{ field.error }}
    </p>
  </div>
</template>

<style scoped>
.app-input-field {
  display: flex;
  flex-direction: column;
  gap: var(--utl-size-02);
}

.app-input-label {
  display: block;
  font-size: var(--ft-size-02);
  font-weight: var(--ft-weight-02);
}

.app-input {
  width: 100%;
  border-radius: var(--utl-radius-01);
  border: 1px solid var(--clr-border-01);
  background-color: var(--clr-surface-02);
  padding: calc(var(--utl-size-06) - 1px) var(--utl-size-07);
  font-size: var(--ft-size-03);
  color: var(--clr-foreground-01);
  outline: none;
  transition:
    border-color 0.15s,
    box-shadow 0.15s;
}

.app-input::placeholder {
  color: color-mix(in srgb, var(--clr-muted-01) 70%, transparent);
}

.app-input:focus {
  border-color: var(--clr-primary-01);
  box-shadow: 0 0 0 2px
    color-mix(in srgb, var(--clr-primary-01) 40%, transparent);
}

.app-input:disabled {
  cursor: not-allowed;
  opacity: 0.6;
}

.app-input-error {
  font-size: var(--ft-size-01);
  color: var(--clr-danger-01);
}
</style>
