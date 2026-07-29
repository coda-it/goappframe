# CLAUDE.md

## Branches

Name branches after their Jira ticket key, e.g. `CODA-56`. The PR's Business justification link (below) is derived from this name, so non-ticket names like `chore/foo` produce broken links.

## Pull Requests

PR titles must follow the pattern `<JIRA ticket number> - <JIRA Title>`, e.g. `CODA-56 - Bump gowebserver dependency`. Use the ticket key from the branch name and the summary from the Jira ticket.

PR bodies must follow the repo's template at `docs/PULL_REQUEST_TEMPLATE.md`, using its exact headings in order:

- **Business justification:** — link to the Jira ticket: `https://codait.atlassian.net/browse/<branch-name>` (substitute the current branch name).
- **Description:** — what changed and why.
- **Visual overview:** — this is a backend library, so this is usually `N/A`.

GitHub does not auto-apply the template to `gh pr create --body`, so fill these fields in manually on every PR.

## Development

This is a Go application framework library (no frontend). Common commands via the Makefile:

* `make install` — install tooling and vendor dependencies
* `make test` — run tests (`go test -mod=vendor ./...`)
* `make lint` — gofmt check, golint and go vet
* `make fix` — apply gofmt formatting
* `make version V=vX.Y.Z` — cut a release: tags, generates `docs/changelogs/CHANGELOG_<V>.md`, and creates a `Build <V>` commit

Dependencies are vendored; run `go mod vendor` after changing `go.mod`.

## Coding Standards

### General
* Don't comment unless there is a good reason to.
* Always cover pull requests with relevant tests.

### Go
* Code must pass `make lint` (gofmt, golint, go vet) — this includes doc comments on exported identifiers.
* Keep the framework generic and re-usable: no application-specific logic. Framework building blocks live in their own packages (`module`, `navigation`, `page`, `route`, `config`).
* Use full descriptive variable names.
