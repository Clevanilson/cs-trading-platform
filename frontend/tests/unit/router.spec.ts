import { describe, expect, it } from 'vitest'

import router from '@/router'

describe('router', () => {
  it('starts with an empty route table', () => {
    expect(router.getRoutes()).toHaveLength(0)
  })
})
