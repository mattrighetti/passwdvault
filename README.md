# PasswdVault

<p align="center">
    <img src="https://www.flaticon.com/svg/static/icons/svg/3039/3039427.svg" height="300px">
</p>

PasswdVault is a flexible, quick and secure CLI password manager that lets you save and retrieve your passwords locally

## Table of Contents
  * [Getting Started](#getting-started)
    + [Installing](#installing)
      - [Installing PasswdVault Command Line Tool Manually](#installing-passwdvault-command-line-tool-manually)
    + [How to Use](#how-to-use)
  * [PasswdVault Documentation](#passwd-documentation)
    + [Commands](#commands)
  * [How To Use](#how-to-use)
  * [Info](#info)
  * [Contact](#contact)

## Getting started
### Installing
First of all you'll need to install [Go](https://golang.org)

#### Installing PasswdVault Command Line Tool Manually
```sh
$ git clone https://github.com/MattRighetti/passwdvault.git
$ cd passwdvault
$ make install
```

## How to use
### Initialize tool
To initialize the basic files that `passwdvault` needs you just need to run
```sh
$ passwdvault init
```
where you would like to initialize your `passwdvault` database.

This command will create a file named _.passwdvaultconfig_ in your `$HOME` directory and another folder called _.passwdvaultdatabase_ in your current working directory.

### Save a password
When you want to save your first password you just need to run the `create` command with its flags `-i` and `-p` to respectively indicate the password identifier and its password, here's an example:

```sh
$ passwdvault create -i githubpass -p secret
```

### Retrieve a passwords
When you need a password that you stored the `get` command is all you're going to need, i.e. to get back my **githubpass** saved before I'll just need to run
```sh
$ passwdvault get githubpass
```

### Delete a password
What about deleting a password that I don't need anymore? The `delete` command does just that, to delete my **githubpass**
```sh
$ passwdvault delete githubpass
```

### List all your password identifiers
Let's say that you forgot how a certain identifier was named, the `list` command will print out all the identifiers that you store in the database
```sh
$ passwdvault list
```

### Search for a password identifier
There may be some cases where you have a ton of passwords and you don't want to list all the password identifiers, but you just want to search for some that has a certain sub-string in it.

`search` is meant just for that, let's say that I have these passwords stored: 
- **githubpass**
- **server1pass**
- **awspass**
- **vs-code-accountpass**
- **some-random-pass** 

if I run
```sh
$ passwdvault search aws
```
it will print out my **awspass** identifier

### Generate a password
I'm not good at choosing secure and strong passwords, this is why I always try to use some password generator that will do it for me, in this case the `generate` command will do just that. Run
```sh
$ passwdvault generate
```
to get a strong password back

### Change configuration file data
Let's say that I moved my database from `$HOME` to `$HOME/matt` folder, if you run any command without changing the _.passwdvaultconfig_ file you will get an error that tell you that the database cannot be found. If you want to specify the new path to the database you'll need to edit the `database.path` attribute in the _.passwdvaultconfig_ file.
The `config` command has been created to ease that process, run
```sh
$ passwdvault config database.path $HOME/matt
```
and the tool will automatically update that value for you and the tool should be working again.

## PasswdVault Documentation
### Commands
- `init`
- `create`
- `get`
- `delete`
- `list`
- `search`
- `config`
- `version`
- `generate`

## Info
PasswdVault is powered by
- [cobra](https://github.com/spf13/cobra)
- [viper](https://github.com/spf13/viper)
- [badger](https://github.com/dgraph-io/badger)
- [passwdgen](https://github.com/MattRighetti/passwdgen)

## Contact
- Please use [Github issue tracker](https://github.com/MattRighetti/passwdvault/issues) for filing bugs or feature requests.
