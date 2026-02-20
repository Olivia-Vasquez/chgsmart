# chgsmart
Generate human-friendly changelogs from git history, even when commit messages aren’t perfect.

## Features

- Parses git history to generate readable changelogs
- Handles imperfect or inconsistent commit messages
- Groups related changes for clarity
- Supports multiple output formats (Markdown, plain text)
- Customizable templates for changelog sections

## Installation

Clone the repository and install dependencies:

```bash
git clone https://github.com/yourusername/chgsmart.git
cd chgsmart
pip install -r requirements.txt
```

## Usage

Run `chgsmart` in your project directory:

```bash
python chgsmart.py
```

You can specify options:

- `--output <file>`: Write changelog to a file
- `--format <type>`: Choose output format (`markdown`, `text`)
- `--since <commit>`: Start changelog from a specific commit

Example:

```bash
python chgsmart.py --output CHANGELOG.md --format markdown --since v1.0.0
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