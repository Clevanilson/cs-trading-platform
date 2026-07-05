# CS Trading Platform — Frontend

A Vue 3 single-page application that provides a UI for the CS Trading Platform Go
microservices. It is a **developer learning lab**: create an account, fund a
session wallet, place buy/sell orders, and review the orders placed during the
session — all against the real `account_service` and `order_service` backends.

Session state (account, wallet balances, order history) is **in-memory only** and
resets on page reload or when you click **New Session**.

## Tech Stack

- **Vue 3** (`<script setup>`) + **TypeScript**
- **Vite** (dev server, build, API proxy)
- **Pinia** (session/wallet/orders/ui stores)
- **Vue Router 4** (route guards for session-gated pages)
- **Tailwind CSS** (dark minimalist theme)
- **Vitest** + **@vue/test-utils** + **happy-dom** (unit & component tests)

## Prerequisites

- **Node.js 20+** (Docker images use Node 22)
- **npm** (ships with Node)
- For a full end-to-end run: **Docker** + **Docker Compose** (to start the backend
  services and RabbitMQ)

## Local Development

```bash
cd frontend
npm install
npm run dev
```

The app is served at http://localhost:5173.

The browser always talks to same-origin proxy paths (`/api/account`, `/api/order`).
Vite forwards those to the backend services, so no CORS configuration is required.
By default the proxy targets `http://localhost:3001` (account) and
`http://localhost:3002` (order) — start the backend from the repo root with
`docker compose up` (or run the Go services locally on those ports).

### Environment Variables

Copy `.env.example` to `.env` to override defaults:

| Variable | Default | Purpose |
|----------|---------|---------|
| `VITE_ACCOUNT_API_URL` | `/api/account` | Browser-facing base path for account_service calls |
| `VITE_ORDER_API_URL` | `/api/order` | Browser-facing base path for order_service calls |
| `ACCOUNT_PROXY_TARGET` | `http://localhost:3001` | Vite dev proxy target for `/api/account` |
| `ORDER_PROXY_TARGET` | `http://localhost:3002` | Vite dev proxy target for `/api/order` |

Keep `VITE_*` values on same-origin proxy paths so the browser never hits the
backend directly (avoids CORS). Only the `*_PROXY_TARGET` values change between
local and Docker environments.

## Docker

The frontend is wired into the root `docker-compose.yml` alongside the backend
services and RabbitMQ.

```bash
# From the repository root
docker compose up
```

This starts:

- `cs-trading-platform-account-service` (:3001)
- `cs-trading-platform-order-service` (:3002)
- `cs-trading-platform-matching-service` (:3003)
- `rabbitmq` (:5672, management UI :15672)
- `cs-trading-platform-frontend` (:5173)

Open http://localhost:5173. Inside Compose, Vite proxies `/api/*` to the service
DNS names (`ACCOUNT_PROXY_TARGET`/`ORDER_PROXY_TARGET` are set to the Compose
service hostnames), so the full stack works with a single command and no manual
`curl`.

- **`Dockerfile.dev`** — dev image used by Compose (Vite dev server with HMR).
- **`Dockerfile`** — production-like multi-stage build (`npm run build` → nginx
  serving `dist/`). `nginx.conf` reverse-proxies `/api/account/` and `/api/order/`
  to the backend services.

## Backend API Endpoints

The UI exercises all 5 backend HTTP endpoints:

| Method | Route | Service | Used by |
|--------|-------|---------|---------|
| `POST` | `/signup` | account_service (:3001) | Signup page |
| `GET` | `/get_account/:id` | account_service (:3001) | Account page |
| `POST` | `/deposit` | order_service (:3002) | Wallet page |
| `POST` | `/withdraw` | order_service (:3002) | Wallet page |
| `POST` | `/place_order` | order_service (:3002) | Trade page |

`place_order` publishes to RabbitMQ, which `matching_service` consumes
asynchronously (no direct HTTP integration from the frontend).

## Demo Flow (< 2 minutes)

1. **Signup** — open http://localhost:5173, enter a name (letters/spaces, 2–255
   chars), submit → redirected to the Dashboard.
2. **Deposit** — go to **Wallet**, deposit `USD` (e.g. `1000`) → session USD
   balance updates.
3. **Place order** — go to **Trade**, keep `BTC-USD`, choose **Buy**, set amount
   and price, submit.
4. **Review** — go to **Orders** to see the order in the session history, or the
   **Dashboard** for the order count and session balances.
5. **Reset** — click **New Session** in the top bar to clear all session state.

Validation errors and backend domain errors (e.g. `Invalid name`,
`Insufficient funds`, `Invalid amount`) surface as toast notifications.

## Testing

```bash
npm test           # Run all unit + component tests
npm run test:unit  # Unit tests only (tests/unit)
npm run test:coverage  # Run with a V8 coverage report
```

Tests use Vitest with happy-dom. Services and stores are unit-tested with mocked
`fetch`/services; pages and forms have component tests covering validation,
loading/empty states, and the happy-path demo flow.

## Project Structure

```
frontend/
├── src/
│   ├── components/   # layout/ (shell, sidebar, topbar) + ui/ (buttons, inputs, ...)
│   ├── pages/        # Route-level views (Signup, Dashboard, Account, Wallet, Trade, Orders)
│   ├── stores/       # Pinia stores: session, wallet, orders, ui
│   ├── services/     # HTTP clients: httpClient, account/wallet/order services
│   ├── router/       # Routes + session guards
│   ├── types/        # Shared TypeScript types
│   ├── composables/  # useToast
│   └── config/       # API base URLs + Vite proxy config
├── tests/            # unit/ + components/
├── Dockerfile        # Production build (nginx)
├── Dockerfile.dev    # Dev server (Vite)
└── nginx.conf        # Reverse proxy for the production image
```
