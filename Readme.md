# Reze - Git Identity Guard

Reze is a CLI tool on `Golang` that checks whether your local Git repository is configured correctly for your selected
development profile.

A profile is a set of rules:

- what `user.name` should be in Git;
- what `user.email` should be in Git;
- what SSH alias should be used in origin;
- should there be `.clang-format` , `.clang-tidy` , `.editorconfig` ;
- should there be a README;
- should there be a license;
- should a pre-commit hook be installed;
- what project type is expected: `school`, `personal`, `work`, `oss`, `github`, `gitlab`.
