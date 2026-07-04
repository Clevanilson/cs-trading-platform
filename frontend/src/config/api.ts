export interface ApiConfigEnv {
  VITE_ACCOUNT_API_URL?: string
  VITE_ORDER_API_URL?: string
  ACCOUNT_PROXY_TARGET?: string
  ORDER_PROXY_TARGET?: string
}

export const DEFAULT_BROWSER_API_URLS = {
  account: '/api/account',
  order: '/api/order',
} as const

export const DEFAULT_PROXY_TARGETS = {
  account: 'http://localhost:3001',
  order: 'http://localhost:3002',
} as const

export function getBrowserApiUrls(env: ApiConfigEnv = {}) {
  return {
    account: env.VITE_ACCOUNT_API_URL ?? DEFAULT_BROWSER_API_URLS.account,
    order: env.VITE_ORDER_API_URL ?? DEFAULT_BROWSER_API_URLS.order,
  }
}

export function getDevProxyTargets(env: ApiConfigEnv = {}) {
  return {
    account: env.ACCOUNT_PROXY_TARGET ?? DEFAULT_PROXY_TARGETS.account,
    order: env.ORDER_PROXY_TARGET ?? DEFAULT_PROXY_TARGETS.order,
  }
}

export function createApiProxyConfig(env: ApiConfigEnv = {}) {
  const targets = getDevProxyTargets(env)

  return {
    '/api/account': {
      target: targets.account,
      changeOrigin: true,
      rewrite: (path: string) => path.replace(/^\/api\/account/, ''),
    },
    '/api/order': {
      target: targets.order,
      changeOrigin: true,
      rewrite: (path: string) => path.replace(/^\/api\/order/, ''),
    },
  }
}
