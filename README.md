# miner
The mining assistant of Dwarves Foundation

## Prerequisite

Go installed: https://github.com/dwarvesf/setup-backend#install-go

## Install

```
$ go get https://github.com/dwarvesf/miner
```

# Usage

## Help

```
$ miner
$ miner -h

// To view usage of specific sub-command
$ miner <command>
```

## Dotfiles

Pull and setup dotfiles from [**dwarvesf/dotfiles**](https://github.com/dwarvesf/dotfiles)

### Init dotfiles

Setup dotfiles from scratch. File `.dfrc` will be created 

```
$ miner dotfiles init
```

The command will fail if `.dfrc` is already existed

### Backup your dotfiles

Bundle your `.dfrc`, dotfiles including your customized configuration. 

```
$ miner dotfiles backup --out minerc.zip
```

### Restore dotfiles

Restore your dotfiles from the backup file. This command requires input bundle. 

```
$ miner dotfiles restore --in minerc.zip
```

### Update

Update dotfiles to latest version but still keep your custom config

```
$ miner dotfiles update
```

### Clean up

```
$ miner dotfiles clean
```

## Project

This requires a self-hosted [Dashboard from Dwarves Foundation](https://github.com/dwarvesf/fortress) to manage all the projects and team information

### Config

You need to retrieve below environment var from the dashboard

- URL
- API Token

```
$ miner config set URL=https://fortress.dwarvesf.com
$ miner config set API=<really long string>
```

### Init

It will help to create project artifacts

- Gitlab group and repos with git template
- Trello board
- Slack channel

```
$ miner project new --name=<name> --config=<config.yml>
```

### New invoice

Prepare invoice of input stage for indicated project

```
$ miner project.invoice 
```

## Misc

This requires a self-hosted [Dashboard from Dwarves Foundation](https://github.com/dwarvesf/fortress)

### Email

The command will help to send email with predefined template from the Dashboard

- National Holiday
- Accept (or reject) candidate
- Setup meeting and workflow with customer

```
$ miner email holiday
```

## Contributing

The main purpose of this repository is to continue to evolve the way we setup development environment, making it faster and easier to use. Development happens in the open on GitHub, and we are grateful to the community for contributing bugfixes and improvements. Read below to learn how you can take part in improving it.

### Contributing Guide

Read our contributing guide to learn about our development process, how to propose bugfixes and improvements, and how to build and test your changes.

## License

Copyright 2017 Dwarves Foundation

Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the License. You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License.
