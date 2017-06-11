# miner
The mining assistant of Dwarves Foundation

# Install

### Prerequisite

- Go installed: https://github.com/dwarvesf/setup-backend#install-go

```
$ go get https://github.com/dwarvesf/miner
```

# Usage

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