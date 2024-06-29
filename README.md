### README for Gogol Project

---

# Gogol

Welcome to **Gogol**, your go-to tool for rapidly creating projects in various programming languages, checking your system's configuration, and facilitating the installation of missing languages. Whether you are starting a simple project or setting up libraries, Gogol is here to streamline the process for you.


## Features

- 🛠️ **Quickly create projects** in multiple programming languages, including the libraries you need.
- 🔍 **Check if your computer meets the necessary configuration** and if the required languages are installed.
- 💡 **Easily install missing languages** to get your environment ready.

## public API
https://github.com/aquemaati/gogol-api

## Examples of Usage

### Creating a Basic Application in Go

```bash
gogol new basic -l go -n myapp
```

### Checking Configuration for Haskell

```bash
gogol new basic -l haskell -n haskell_app
```

## Future Plans

Currently, Gogol supports creating simple basic projects and libraries. In the future, I plan to implement a broader range of applications with many languages, including:

- APIs
- Command Line Interfaces (CLI)
- Networking Applications
- Server Applications
- Cloud Applications

## Contribute to the Project

Your feedback and contributions are highly appreciated! Feel free to test Gogol and share your thoughts on the [GitHub repository](https://github.com/aquemaati/gogol).


## Usage

To create a new project with Gogol, use the following command:

```bash
gogol new [project_type] -l [language] -n [project_name] -d [dependance]
```

Example:

```bash
gogol new basic -l go -n myGoApp -d github.com/01-edu/z01
```

## Support

For any questions or issues, please open an issue on GitHub

![gogolReadmeImage](https://github.com/aquemaati/gogol/assets/160750138/efe83ddc-b1f8-491f-8ef2-a86cc52acb23)

---

Feel free to edit this README to fit your specific needs and update the placeholder link for the comic image with an actual image URL.
