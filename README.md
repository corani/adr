# Architecture Decision Records

A simple command-line tool to manage ADRs in markdown format.

## Usage

- `adr init [path]`
  Initialize the ADR path and create a `template.md`. Feel free to update the template, but leave
  the yaml front-matter intact.
- `adr new <title>`
  Create a new ADR with the given title and open it in your `$EDITOR`.
- `adr show <id>`
  Show the ADR with the given id.
- `adr edit <id>`
  Open the ADR with the given id in your `$EDITOR`.
- `adr list`
  List all ADRs with their status, date and title.
- `adr update <id> <status>`
  Update the ADR with the given id, setting the status to one of: `proposed`, `accepted`,
  `deprecated` or `superseded`.

Note: The repository no longer commits a `vendor/` directory. Dependencies are managed
via Go modules. To recreate a `vendor/` directory locally, run:

```
ADR_VENDOR=true ./tools/install.sh
```


