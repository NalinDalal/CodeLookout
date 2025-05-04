# Guidelines for Contributing to This Repository

We’re glad you want to contribute! To get started, please **fork the repository**, create your own branch from your fork, and follow the guidelines below to ensure your changes align with the project’s workflow and quality standards.

## Working on changes

- Always **create a new branch** for your work.
- **Push regularly** to avoid data loss.
- We follow a **linear commit history** — always use `git rebase` instead of `git merge`.

## Branch naming conventions

Branch names should reflect the purpose of the work and follow a consistent structure.

### Recommended patterns

| Type                   | Pattern                    | Example                                          | Notes                                                |
| ---------------------- | -------------------------- | ------------------------------------------------ | ---------------------------------------------------- |
| New Features           | `feature/<feature-name>`   | `feature/integrate-openai-apis-for-code-summary` | For new user-facing features                         |
| Bug Fixes              | `fix/<fix-name>`           | `fix/auth-error`                                 | For bug fixes                                        |
| Hotfixes               | `hotfix/<fix-name>`        | `hotfix/breaking-changes-openai-sdk`             | For urgent, critical fixes                           |
| Security Updates/Fixes | `security/<change>`        | `security/suppress-server-headers`               | For security-related changes                         |
| Maintenance            | `maintenance/<task>`       | `maintenance/run-integration-tests-in-workflows` | For internal improvements or upkeep                  |
| Development / Other    | `<username>/<description>` | `alex/job-queue-poc`                             | For work not covered above; often temporary branches |

> **Tip:** Avoid vague names like `fix/bug-fix`. Be clear and descriptive.

## Pull request conventions

GitHub will automatically apply the default PR template located [here](https://github.com/Mentro-Org/CodeLookout/blob/main/.github/PULL_REQUEST_TEMPLATE.md).

### PR guidelines

- **Title**: Clearly describe the purpose of the PR.
- **Description**: Include a detailed explanation of what changes were made and why.
- **Merge strategy**: Use **`Squash and Merge`** once all reviews and checks are approved.

## Commit message conventions

Good commit messages are essential for a readable project history. Please follow this guide:

👉 [How to Write a Git Commit Message](https://cbea.ms/git-commit/)
