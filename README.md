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

## TODO

- [ ] link superseded ADRs
- [ ] more robust detection of project root / adr root
- [ ] better help texts and error reporting
- [ ] generate an index in markdown format, to make it easier to use in e.g. GitHub
- [ ] post-process step after editing, to create hyperlinks for ADRs mentioned in the text
- [ ] CI/CD setup and versioning
- [ ] template seems to break github pages production, move it?
- [ ] do we need a .adr.yaml or similar config file in the project root?
