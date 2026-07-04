import { readFileSync } from 'node:fs'
import { resolve } from 'node:path'

import { describe, expect, it } from 'vitest'

import {
  DEFAULT_BROWSER_API_URLS,
  DEFAULT_PROXY_TARGETS,
  createApiProxyConfig,
  getBrowserApiUrls,
  getDevProxyTargets,
} from '@/config/api'

describe('API config', () => {
  it('uses proxy-path browser URLs by default', () => {
    expect(getBrowserApiUrls()).toEqual(DEFAULT_BROWSER_API_URLS)
  })

  it('maps /api/account to the account service base URL', () => {
    const proxy = createApiProxyConfig()
    const accountProxy = proxy['/api/account']

    expect(accountProxy.target).toBe(DEFAULT_PROXY_TARGETS.account)
    expect(accountProxy.rewrite('/api/account/signup')).toBe('/signup')
  })

  it('maps /api/order to the order service base URL', () => {
    const proxy = createApiProxyConfig()
    const orderProxy = proxy['/api/order']

    expect(orderProxy.target).toBe(DEFAULT_PROXY_TARGETS.order)
    expect(orderProxy.rewrite('/api/order/place_order')).toBe('/place_order')
  })

  it('accepts Docker proxy targets from the environment', () => {
    expect(
      getDevProxyTargets({
        ACCOUNT_PROXY_TARGET: 'http://cs-trading-platform-account-service:3001',
        ORDER_PROXY_TARGET: 'http://cs-trading-platform-order-service:3002',
      }),
    ).toEqual({
      account: 'http://cs-trading-platform-account-service:3001',
      order: 'http://cs-trading-platform-order-service:3002',
    })
  })

  it('.env.example documents both local and Docker URL patterns', () => {
    const envExample = readFileSync(resolve(process.cwd(), '.env.example'), 'utf8')

    expect(envExample).toContain('VITE_ACCOUNT_API_URL=/api/account')
    expect(envExample).toContain('VITE_ORDER_API_URL=/api/order')
    expect(envExample).toContain('ACCOUNT_PROXY_TARGET=http://localhost:3001')
    expect(envExample).toContain('ORDER_PROXY_TARGET=http://localhost:3002')
    expect(envExample).toContain('http://cs-trading-platform-account-service:3001')
    expect(envExample).toContain('http://cs-trading-platform-order-service:3002')
  })
})
