import { mount } from '@vue/test-utils'
import { createPinia } from 'pinia'
import { describe, expect, it } from 'vitest'
import { createRouter, createWebHistory } from 'vue-router'

import App from '@/App.vue'

describe('App.vue', () => {
  it('renders without throwing when mounted', () => {
    const pinia = createPinia()
    const router = createRouter({
      history: createWebHistory(),
      routes: [],
    })

    const wrapper = mount(App, {
      global: {
        plugins: [pinia, router],
      },
    })

    expect(wrapper.text()).toContain('CS Trading Platform')
  })
})
