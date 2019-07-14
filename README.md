# ucasnauth
University of Chinese Academy of Sciences campus Network AUTHentication.

## Introduction
A simple tool to help you login the campus network of UCAS (University of Chinese Academy of Sciences).  
Your username and password will be saved to your machine, with AES encryption.  
I won't collect your information or send it to anywhere else.

## Main Features
#### Easy to use  
Just execute the program, and you will login the campus network of UCAS.
See [Usage](#usage) section for details.
#### Remember your account information  
Your username and password will be saved to your machine. You don't need to input them every time.
#### Encrypt your account information  
Your account information is encrypted with AES. The key depends on your machine fingerprint.
Don't worry about data files being stolen. Other people will get no knowledge from them.

## Installation
Before the installation, you must install go 1.12 or newer version of the Go programming language.  
You can download it from its [official website](https://golang.org/dl/).

---

Get this tool with

```bash
go get -u github.com/donyori/ucasnauth
```

A executable file will be created under your `$GOBIN` (if you set this environment variable) or `$GOPATH/bin`.  
It's recommended to move the executable file to a separate directory, because some data files will be saved under this directory. (See [Created Files](#created-files) section for details.)

## Usage
For the first time or when you want to login another account:

```bash
ucasnauth username password
```

It will login with the specified username and password.  
If login succeeds, the username and password will be saved to your machine (previous data will be replaced).

---

Then, if you want to login with previous information, just execute the program with no arguments:

```bash
ucasnauth
```

---

To logout:

```bash
ucasnauth logout
```

---

To delete saved account information:

```bash
ucasnauth delete
```

## Created Files
This tool will create some files to save the data, including username, password, and data used for AES encryption.  
These files are under the directory:
`$HOME/.ucasnauth` (`%USERPROFILE%\.ucasnauth` for Windows)

## Uninstallation
To uninstall this tool, remove the executable file and its data files as shown in [Created Files](#created-files).
If you want to remove the source codes as well, remove `$GOPATH/src/github.com/donyori/ucasnauth` and all its contents.

## License
The MIT License (MIT) - [donyori](https://github.com/donyori/). Please have a look at the [LICENSE](LICENSE) for more details.
