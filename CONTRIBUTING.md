
# Guidelines for Contributing to CodeLookout

We’re glad you want to contribute! Please read these rules and setup instructions before submitting a pull request.

## Contribution Rules

- **Fork the repository** and create a new branch for your work.
- **Branch naming:**
	- New features: `feature/<feature-name>`
	- Bug fixes: `fix/<bug-description>`
	- Hotfixes: `hotfix/<critical-fix>`
	- Security: `security/<change>`
	- Maintenance: `maintenance/<task>`
	- Other/dev: `<username>/<description>`
- **Push regularly** to your branch.
- **Linear commit history:** Use `git rebase` (not `git merge`).
- **Pull Requests:**
	- Use clear, descriptive titles and detailed descriptions.
	- Reference related issues if applicable.
	- Use **Squash and Merge** after approval.
- **Commit messages:**
	- Write clear, meaningful commit messages ([guide](https://cbea.ms/git-commit/)).
- **Tests:**
	- Add or update tests for new features and bug fixes.
- **Code style:**
	- Follow Go best practices and project conventions.
- **Review:**
	- Address all review comments before merging.

## Local Setup Instructions

Follow these steps to set up CodeLookout for local development:

1. **Clone your fork and install dependencies:**
	 ```sh
	 git clone https://github.com/<your-username>/CodeLookout.git
	 cd CodeLookout
	 cd server
	 go mod download
	 ```

2. **Configure environment variables:**
	 - Copy or create a `.env` file (see `server/development.md` for details).
	 - Set up a local PostgreSQL and Redis instance, or use Docker Compose.

3. **Run database migrations:**
	 ```sh
	 ./scripts/run_migrations.sh
	 ```

4. **Start the development server:**
	 - With Docker Compose:
		 ```sh
		 docker-compose up
		 ./scripts/run_smeeclient.sh  # For GitHub webhook forwarding
		 ```
	 - Or locally:
		 ```sh
		 go run cmd/main.go
		 ./scripts/run_smeeclient.sh
		 ```

5. **Access the app:**
	 - API: `http://localhost:8080/api/`
	 - Analytics Dashboard: `http://localhost:8080/analytics`

6. **Run/test LLM and SonarQube CLI tools:**
	 ```sh
	 make test-llm
	 make test-sonarqube
	 ```

7. **See `server/development.md` for more details and troubleshooting.**

---

Thank you for contributing to CodeLookout!
