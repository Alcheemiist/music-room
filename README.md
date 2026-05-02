# MusicRoom

Cross-platform iOS, Android, and Web application for collaborative music listening. Built with React Native (Expo), AdonisJS, and Temporal in a monorepo.

## Features

### Music Track Vote (MTV)

A collaborative music listening session where users suggest and vote for tracks to be played.

- **Constraints** — the creator can restrict participation by location and time window
- **Emission modes** — _broadcast_ (all users play sound) or _direct_ (a single designated user emits)
- **Device selection** — users choose which of their devices plays sound
- **Privacy** — rooms can be public or private (invite-only)
- **Permissions** — members can suggest tracks, vote, and control playback (play, pause, skip) depending on granted permissions
- **Social** — users can chat and follow each other inside a room

### Music Playlist Editor (MPE)

A real-time collaborative playlist editor.

- Users create or join MPE rooms to suggest, remove, and reorder tracks
- A user can be a member of multiple MPE rooms simultaneously, all listed in their Library
- Any MPE room can be exported into an MTV room, with the playlist as the initial track queue

## Packages

| Package | Description |
|---|---|
| `packages/client` | React Native (Expo) mobile and web app |
| `packages/server` | AdonisJS REST + WebSocket API |
| `packages/temporal` | Go-based Temporal workflows and activities |
| `packages/types` | Shared Zod schemas and TypeScript types |
| `packages/stress` | Artillery load tests |

## Documentation

- [Technical stack](docs/technical-stack.md)
- [Setup guide](docs/setup.md)
