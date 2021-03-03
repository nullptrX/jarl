# Jarl

![License](https://img.shields.io/github/license/nullptrX/jarl)

> Locate jar coordinates right from your Alfred and Terminal.

![Jarl](docs/demo.gif)

No need to browse [mvnrepository.com](https://mvnrepository.com).

<img src="https://cdn.jsdelivr.net/gh/nullptrX/assets/images/20210303195505.gif"/>

Workflow for Alfred 4.

## Installation

  ```shell script
  go build -o jarl github.com/devcsrj/jarl/cli
  ```

## Usage for Workflow
Download [Jarl](workflow/Jarl.alfredworkflow) and open it. (Make sure you have installed Alfred 4.)

Type `jarl`, and your query, to search for java libraries at maven central repository.

Select a item and press <kbd>Enter</kbd> to copy gradle dependency to clipboard.

Hold <kbd>Command</kbd> and press <kbd>Enter</kbd> to select other style's dependency to clipboard.

## Config

- `PROXY`: Leave empty for no proxy
  
    ```
    socks5://127.0.0.1:port
    # or
    http://127.0.0.1:port
    ```
  
- `STYLE`: Leave empty for gradle style.
    
    ```
    maven
    gradle
    sbt
    ivy
    grape
    leiningen
    buildr
    ```

<img src="https://cdn.jsdelivr.net/gh/nullptrX/assets/images/20210303194831.png"/>
