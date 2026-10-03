# error-log-summarizer

Group application errors into concise operator-facing summaries.

## Run

Requires Go 1.22+.

```sh
go run .
```

The tool reads its development gateway settings from `config/development.json`. Override those values in your deployment environment before production use. Review generated output before applying it to another system.

## Model

This example targets the `gpt-6-sol` frontier model through the configured OpenAI-compatible router.
