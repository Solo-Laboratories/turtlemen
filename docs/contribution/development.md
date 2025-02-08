# Development Guidelines

## Definition of Done
This list is still going and not exhaustive.

- Commits conform to Conventional Commits

## Commit Convention
We use [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/) which is highly encoruaged and will be enforced during merge requests. All merge requests will be refused if not in conventional commit format. 

## Coding Standards
TBA

## Signing Commits
We require that all commits are signed. It is rather simple to achieve but it assures that a known user has commited their works and Gitlab verified the works are theirs.

To setup your environment to sign your keys automatically using an SSH key [Follow this guide](https://docs.github.com/en/authentication/managing-commit-signature-verification/about-commit-signature-verification#ssh-commit-signature-verification). However, you're more than welcome to setup a [GPG Key](https://docs.github.com/en/authentication/managing-commit-signature-verification/about-commit-signature-verification#gpg-commit-signature-verification) if you prefer.

## Pre-hook Commits
We enforce a few things during commits. It is strongly adviced that you install the `pre-commit` python module to ensure you're conforming to best practices. 

To setup your system, ensure you have python 3+ installed then follow the steps below.
1. `pip install pre-commit` You might have to add your ~/.local/bin to your bashrc file if on linux for pre-commit to appear.
2. `pre-commit install --install-hooks`