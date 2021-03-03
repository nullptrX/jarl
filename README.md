# Jarl

![License](https://img.shields.io/github/license/nullptrX/jarl)
![Version](https://img.shields.io/github/v/release/nullptrX/jarl)

> Locate jar coordinates right from your Alfred and Terminal.

![Jarl](docs/demo.gif)
<img src="https://cdn.jsdelivr.net/gh/nullptrX/assets/images/20210303192635.gif"/>

No need to browse [mvnrepository.com](https://mvnrepository.com).

Workflow for Alfred 4.

## Installation

  ```shell script
  go build -o jarl github.com/devcsrj/jarl/cli
  ```

## Usage for Workflow
Download [Maven Querier](workflow/Maven%20Querier.alfredworkflow) and open it. (Make sure you have installed Alfred 4.)

Type `mvn`, and your query, to search for java libraries at maven central repository.

Select a item and press <kbd>Enter</kbd> to copy gradle dependency to clipboard.

Hold <kbd>Command</kbd> and press <kbd>Enter</kbd> to select other style's dependency to clipboard.

## Config

<img src="https://cdn.jsdelivr.net/gh/nullptrX/assets/images/20210303194831.png"/>
