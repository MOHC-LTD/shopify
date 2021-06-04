# How to contribute

## 🪝 Git hooks

This project includes two git hooks

- `pre-commit`
- `pre-push`

To enable these, run the following command

```sh
git config core.hooksPath .githooks
```

> Note - on a unix machine (mac and linux) you may need to make both files executable

```sh
chmod +x .githooks/pre-commit
chmod +x .githooks/pre-push
```
