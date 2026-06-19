# Reze - Git Identity Guard

Reze is a Go-powered local CLI tool designed to protect your Git repositories from developer identity mix-ups, incorrect
configurations, and accidental commits under the wrong profile.

If you constantly juggle multiple development modes-such as university/school assignments, corporate work,
open-source projects, or low-level experiments - Reze ensures you never accidentally leak your personal email into a
corporate repo, forget a mandatory configuration file, or push code with missing style guidelines.

It acts as an automated "doctor" that audits your repository's state against a predefined **development profile**.

## 🛠️ What Reze Checks

When you run an audit, Reze validates the repository structure and configuration across several categories:

* **Git Identity:** Verifies local `user.name` and `user.email`. It flags a warning if a global fallback is used when a
  local identity is strictly required.
* **Remote Alignment:** Validates the `origin` URL scheme (SSH/HTTPS), host permissions, and expected repository owner
  to prevent accidental pushes to the wrong target.
* **Project Standards & Documentation:** Checks for mandatory files tailored to the specific profile (e.g., `go.mod`,
  `.golangci.yml`, and `Taskfile.yml` for Go projects, or `.clang-format` and `.clang-tidy` for C++ systems).
* **Git Hooks:** Confirms the presence of automated enforcement mechanisms (like a pre-commit hook) to catch
  configuration errors before code leaves your machine.
