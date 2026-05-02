# Technical Stack

## Monorepo

The project is a Yarn Workspaces monorepo managed with [Lerna](https://lerna.js.org/). All packages share TypeScript configuration and ESLint/Prettier rules defined at the root.

Node version and Yarn version are pinned via [Volta](https://volta.sh/) (`node 16.8.0`, `yarn 1.22.10`).

---

## Client — `packages/client`

React Native application targeting iOS, Android, and Web via [Expo](https://expo.dev/) (SDK 41).

| Concern | Library |
|---|---|
| Framework | React Native + Expo |
| State machines | [XState](https://xstate.js.org/) + `@xstate/react` |
| Server state / data fetching | [React Query](https://react-query.tanstack.com/) |
| Real-time | [Socket.IO client](https://socket.io/) |
| Forms | [React Hook Form](https://react-hook-form.com/) |
| Validation | [Zod](https://zod.dev/) (via `@musicroom/types`) |
| Navigation | [React Navigation](https://reactnavigation.org/) v5 |
| Styling | [Dripsy](https://www.dripsy.xyz/) |
| HTTP | [Redaxios](https://github.com/developit/redaxios) |
| Animation | [Moti](https://moti.fyi/) + Reanimated 2 |
| Maps | `react-native-maps` + `react-native-google-places-autocomplete` |
| YouTube playback | `react-native-youtube-iframe` (mobile) + `react-player` (web) |
| Google auth | `expo-auth-session` |

Testing uses `@testing-library/react-native`, `@xstate/test` for model-based testing, and [Playwright](https://playwright.dev/) for end-to-end web tests. MSW (`@mswjs/data`) provides in-memory mock data for tests.

---

## Server — `packages/server`

REST and WebSocket API built with [AdonisJS](https://adonisjs.com/) v5.

| Concern | Library |
|---|---|
| Framework | AdonisJS |
| Database ORM | AdonisJS Lucid (PostgreSQL via `pg`) |
| Cache / pub-sub | Redis (`@adonisjs/redis`, `@socket.io/redis-adapter`) |
| Auth | AdonisJS Auth + Bouncer (authorization policies) |
| Real-time | [Socket.IO](https://socket.io/) |
| Email | AdonisJS Mail + [MJML](https://mjml.io/) templates |
| Validation | [Zod](https://zod.dev/) (via `@musicroom/types`) |
| Google APIs | `googleapis` + `@googlemaps/google-maps-services-js` |
| Password hashing | `phc-argon2` |
| API docs | `adonis-autoswagger` |

Infrastructure runs in Docker: PostgreSQL and Redis are started via `docker-compose` inside `packages/server`.

---

## Temporal — `packages/temporal`

Workflow orchestration service written in Go, backed by [Temporal](https://temporal.io/).

| Concern | Library |
|---|---|
| Workflow engine | [Temporal Go SDK](https://pkg.go.dev/go.temporal.io/sdk) |
| State machines | [Brainy](https://github.com/Alcheemiist/brainy) |
| HTTP routing | [Gorilla Mux](https://github.com/gorilla/mux) |
| Validation | [go-playground/validator](https://github.com/go-playground/validator) |

The package exposes two binaries:

- **api** — HTTP server that receives requests from the AdonisJS server and signals Temporal workflows
- **worker** — Temporal worker that executes workflow and activity code

Temporal server itself runs in Docker via `packages/temporal/docker-compose`.

---

## Types — `packages/types`

Shared TypeScript package consumed by both `client` and `server`.

Contains [Zod](https://zod.dev/) schemas that define the contracts for MTV and MPE room WebSocket events, ensuring type safety across the network boundary without duplication.

---

## Stress — `packages/stress`

Load testing suite using [Artillery](https://www.artillery.io/) to stress test the server and Temporal workflows under concurrent user load.

---

## Infrastructure overview

```
┌─────────────┐     HTTP/WS     ┌──────────────┐     HTTP      ┌─────────────────┐
│   Client    │ ──────────────► │    Server    │ ────────────► │  Temporal API   │
│ (Expo app)  │                 │  (AdonisJS)  │               │    (Go HTTP)    │
└─────────────┘                 └──────┬───────┘               └────────┬────────┘
                                       │                                 │
                                  ┌────┴─────┐                  ┌───────┴────────┐
                                  │ Postgres │                  │ Temporal Worker │
                                  │  Redis   │                  │    (Go)         │
                                  └──────────┘                  └────────────────┘
```
