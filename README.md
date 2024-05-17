# `documentor`

`documentor` is an easy-to-use, efficient, and powerful command-line
application that uses the OpenAI GPT-4 API to review documentation and
provide feedback on how to improve it.

It assists documentation writers in creating better documentation by
automating parts of the review process.

> [!WARNING]
> This project is a PoC and is still in early stages of development. It
> is not ready for production use, and doesn't use Datadog's OpenAI
> proxy server yet.

## Installation

### From source

First, ensure that the following dependencies are installed:

- Go 1.22 or above.
- make.
- [scdoc](https://git.sr.ht/~sircmpwn/scdoc).

Optionally, you can install
[glow](https://github.com/charmbracelet/glow) to render the Markdown
output with more style.

Then compile and install:

```bash
make
sudo make install
```

## Usage

```bash
$ documentor --help
NAME:
   documentor - review technical documentation with the power of AI

USAGE:
   documentor [global options] [file]

VERSION:
   0.1.0

GLOBAL OPTIONS:
   --key value, -k value  the OpenAI API key to use [$DOCUMENTOR_KEY]
   --help, -h             show help
   --version, -v          print the version
```

Refer to the _documentor(1)_ manpage after installation for more
information.

## Contributing

Anyone can help make `documentor` better. Refer to the [contribution
guidelines](CONTRIBUTING.md) for more information.

---

This project is released under the [Apache-2.0 License](LICENSE.md).
