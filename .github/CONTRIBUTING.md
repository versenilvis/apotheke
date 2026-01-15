# Contributing to Apotheke

Thank you for your interest in Apotheke. We welcome community contributions and appreciate your time and effort in helping improve the project. Before getting started, please take a moment to review these guidelines.

Need to contact? Reach out to my Telegram [@VerseNilVis](t.me/VerseNilVis).

## 1. How to Contribute

We welcome contributions in various forms, including:

- Bug fixes
- Documentation improvements
- Tests and performance optimizations

To contribute, follow these steps:

- **Fork** the repository and create a new branch.
- **Make your changes**, ensuring they align with our coding standards (§3 below).
- **Run tests** to ensure your changes do not break existing functionality.
- **Submit a pull request (PR)** with a clear description of the changes.
- Another maintainer or I will review your PR, suggest any necessary changes, and merge it once approved.

### Commit Messages

Use clear and descriptive commit messages. Follow the conventional commit format when possible:

+ feat(resolver): Add fuzzy matching for list command
+ fix(executor): Resolve issue with cwd flag
+ docs(readme): Update with new command examples

## 2. Legal Terms

By submitting a contribution, you represent and warrant that:

- It is your original work, or you have sufficient rights to submit it.
- You grant the Apotheke maintainers and users the right to use, modify, and distribute it under the AGPL-3.0 license (see LICENSE file).
- To the extent your contribution is covered by patents, you grant a perpetual, worldwide, non-exclusive, royalty-free, irrevocable license to the Apotheke maintainers and users to make, use, sell, offer for sale, import, and otherwise transfer your contribution as part of the project.

We do not require a Contributor License Agreement (CLA). However, by contributing, you agree to license your submission under terms compatible with the AGPL-3.0 License and to grant the patent rights described above. If your contribution includes third-party code, you are responsible for ensuring it is AGPL-3.0-compatible and properly attributed.

Where permitted by law, you waive any moral rights in your contribution (e.g., the right to object to modifications). If such rights cannot be waived, you agree not to assert them in a way that interferes with the project's use of your contribution.

## 3. Coding Standards

To maintain a consistent codebase, please follow these guidelines:

- Use the existing coding style and conventions.
- No comments in source files.
- Write tests for new features and bug fixes.
- Avoid introducing unnecessary dependencies.
- Run `go fmt` before committing.
- Run `go test ./...` to ensure tests pass.

## 4. Reporting Issues

If you find a bug or have a feature request, please open an issue and provide as much detail as possible:

- Steps to reproduce, including operating system and Apotheke version
- Expected and actual behavior
- Suspected cause (if any)

## 5. Recognition

We use the All Contributors specification to recognize community members. If your contribution is merged, you will be added to the project's list of contributors. This includes contributions of all kinds—code, documentation, design, testing, and more.
