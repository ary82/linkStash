# urlStash

<p>
  <img src="https://badgen.net/github/last-commit/ary82/urlStash?icon=git"/>
  <img src="https://badgen.net/badge/Go/1.22?color=029EC7"/>
  <img src="https://badgen.net/badge/Svelte/Kit?color=E83B05"/>
  <img src="https://badgen.net/github/license/ary82/urlstash"/>
</p>

A webapp for sharing urls.

**\[Currently WIP\]**

> This project is currently Work in Progress. Some functionality might not be implemented yet.

Table of Contents

- [urlStash](#urlstash)
- [Motivation](#motivation)
- [Features](#features)
- [Tech Used](#tech-used)
- [Getting Started](#getting-started)
  - [Prerequisities](#prerequisities)
  - [Download](#download)
  - [Run](#run)
- [Todo](#todo)

## Motivation

While there are several existing services that allow you to bookmark or share links, many are closed-source and lack the flexibility and customization. With urlStash, you get the power and control over your data by self hosting your own instance.

## Features

- Easily create collections of links (stashes) and organize them with descriptive titles.
- Easily Add/Remove more links to your Stashes

## Tech Used

- Golang net/http for building the REST API
- Postgres as the Database
- Sveltekit for the frontend

## Getting Started

### Prerequisities:

- Go >=1.22
- Nodejs and npm
- A Postgres Database up and running on your system

### Download:

```bash
git clone https://github.com/ary82/urlStash.git && cd urlStash
```

### Run:

- First, you'll need to populate the env vars required for the project. Copy the .env.example files and populate them

```bash
cp ./.env.example ./.env
cp ./web/.env.example ./web/.env
# Populate the empty fields
```

- Next, prepare your DB by running

```bash
psql -h localhost -p PORT -U USERNAME -W urlStash -f ./migrations/init_up.sql
```

- Build and run the REST API

```bash
make run
# Or watch for changes using air
make watch
```

- Build and run frontend. Install dependencies using your node package manager(npm, yarn, pnpm). I'm using pnpm for this project.

```bash
cd web
# Install dependencies
pnpm install
# Run in dev mode
pnpm run dev
# Or build
pnpm run build
```

## Todo

- Code

  - [ ] Complete basic functionality
  - [x] Decouple Database logic completely, pass behaviour instead
  - [ ] Add More Tests
  - [ ] Make middleware more readable

- Features

  - [ ] Sharing with specific users
  - [ ] Make Google auth optional in build step
  - [ ] Dockerize and make self hostable
  - [ ] Allow importing/exporting from browser
