# Contributing to PIMO

:+1::tada: First off, thanks for taking the time to contribute! :tada::+1:

## How to submit a contribution

### Prerequisites

You need :

- Visual Studio Code ([download](https://code.visualstudio.com/)) with the [Remote - Containers extension](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.remote-containers) installed.
- Docker Desktop (Windows, macOS) or Docker CE/EE (Linux)

Details are available on [the official Visual Studio documentation](https://code.visualstudio.com/docs/remote/containers#_getting-started).

### Steps to submit a contribution

1. Make sure the topic is not already discussed by an opened issue
2. Submit your idea in [an new issue](https://github.com/CGI-FR/PIMO/issues/new) to discuss about it with us
3. Then you can open and submit a Pull Request to the project

### Checklists for pull requests

- [ ] The code you added/modified **compile**
- [ ] The code you added/modified **is linted** (run the `neon lint` command)
- [ ] The code you added/modified **is covered by unit tests** and **all tests are passing** (run the `neon test` command)
- [ ] The features you added/modified or the bugs you fixed **are covered by integration tests** and **all integration tests are passing** (run the `neon test-int` command)
