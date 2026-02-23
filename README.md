# chgsmart
Generate human-friendly changelogs from git history, even when commit messages aren’t perfect.

## Features

- Parses git history to generate readable changelogs
- Handles imperfect or inconsistent commit messages
- Groups related changes for clarity
- Supports multiple output formats (Markdown, plain text)
- Customizable templates for changelog sections

## Installation

Install chgsmart:

```bash
brew install pipx
pipx ensurepath

pipx install git+https://github.com/Olivia-Vasquez/chgsmart.git
```

Verify install:

```bash
chgsmart --help
```

## Usage

Run `chgsmart` in your project directory:

```bash
python chgsmart.py
```

You can specify options:

- `--out <filepath>`: Write changelog to a file
- `--group-by <category>`: Group by category (`type`, `area`) 
- `--format <type>`: Choose output format (`markdown`, `text`)
- `--include-merges`: Include merge commits

Example:

```bash
python chgsmart.py --out /Desktop/CHANGELOG.md --format markdown --include-merges
```

## How It Works

chgsmart analyzes your git history and extracts meaningful changes. It uses heuristics to group commits, rewrite unclear messages, and organize them into sections like "Features", "Bug Fixes", and "Other Changes".

## Customization

You can customize changelog templates by editing the `templates/` directory. See the [docs](docs/customization.md) for details.

## Contributing

Contributions are welcome! Please open issues or pull requests for bug fixes, features, or documentation improvements.

## License

This project is licensed under the MIT License.

## Contact

For questions or feedback, open an issue or email oliviavasquez@fakemail.com.