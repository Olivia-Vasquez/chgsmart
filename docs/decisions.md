# Decisions

1. **Use Go for CLI implementation**
   - Why: Fast, portable, strong ecosystem for CLI tools.
   - Alternatives: Python, Node.js.

2. **Parse git history directly via `git log`**
   - Why: Reliable, avoids external dependencies.
   - Alternatives: Use go-git library (pending).

3. **Classify commits using conventional commit prefixes and heuristics**
   - Why: Handles imperfect messages, covers most cases.
   - Alternatives: Require strict conventional commits (rejected).

4. **Markdown as default output format**
   - Why: Widely used, easy to read and share.
   - Alternatives: Plain text, JSON (planned).

5. **Configurable areas via YAML**
   - Why: Allows grouping by project area.
   - Alternatives: Hardcoded areas (rejected).

6. **Semantic versioning (Planned)**
   - Why: Standard for releases.
   - Alternatives: Date-based versioning.

7. **Automated changelog generation (Planned)**
   - Why: Reduce manual work, improve consistency.
   - Alternatives: Manual changelog editing.

8. **Decision pending: Test suite implementation**
   - Evidence needed: Test files or CI config.

9. **Decision pending: Binary distribution**
   - Evidence needed: Build scripts or release artifacts.

10. **Decision pending: Custom output templates**
    - Evidence needed: Template files or config options.
