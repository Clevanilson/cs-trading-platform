<script setup lang="ts">
withDefaults(
  defineProps<{
    type?: "button" | "submit" | "reset";
    disabled?: boolean;
    loading?: boolean;
    loadingText?: string;
  }>(),
  {
    type: "button",
    disabled: false,
    loading: false,
    loadingText: "Loading…",
  },
);
</script>

<template>
  <button
    :type="type"
    :disabled="disabled || loading"
    class="app-button"
  >
    <span v-if="loading">{{ loadingText }}</span>
    <slot v-else />
  </button>
</template>

<style scoped>
.app-button {
  display: flex;
  width: 100%;
  align-items: center;
  justify-content: center;
  border: none;
  border-radius: var(--utl-radius-03);
  background-color: var(--clr-primary-01);
  padding: var(--utl-size-06) var(--utl-size-07);
  font-size: var(--ft-size-03);
  font-weight: var(--ft-weight-03);
  color: var(--clr-background-01);
  transition: filter 0.15s, box-shadow 0.15s;
}

.app-button:hover:not(:disabled) {
  filter: brightness(1.1);
}

.app-button:focus {
  outline: none;
  box-shadow: 0 0 0 2px color-mix(in srgb, var(--clr-primary-01) 50%, transparent);
}

.app-button:disabled {
  cursor: not-allowed;
  opacity: 0.6;
}
</style>
