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
  Use `--format` / `-f` to control output: `md` (default, rendered markdown), `raw` (unrendered markdown), `json` (structured JSON with frontmatter fields and body as separate keys).
- `adr edit <id>`
  Open the ADR with the given id in your `$EDITOR`.
- `adr list`
  List all ADRs with their status, date and title.
  Use `--format` / `-f` to control output: `md` (default, rendered markdown), `raw` (unrendered markdown), `json` (structured JSON).
- `adr find <query>`
  Find ADRs whose title matches the query. Words are matched in order, case-insensitively, with
  anything allowed between them. Use `--text` / `-t` to also search frontmatter fields and the body.
  Use `--format` / `-f` to control output: `md` (default), `raw`, `json`.
- `adr update <id> <status>`
  Update the ADR with the given id, setting the status to one of: `proposed`, `accepted`,
  `deprecated` or `superseded`.
