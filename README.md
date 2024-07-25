# `documentor`

`documentor` is an easy-to-use, efficient, and powerful command-line
application that uses the power of AI to review documentation and
provide feedback on how to improve it, draft new documentation from
scratch, and more.

It assists documentation writers in creating better documentation by
automating parts of the writing and review process, allowing them to
focus on the content and structure of the document itself.

## Installation

### From source

First, ensure that the following dependencies are installed:

- Go 1.22 or above.
- make.
- [scdoc](https://git.sr.ht/~sircmpwn/scdoc).

Optionally, you can install
[glow](https://github.com/charmbracelet/glow) to render the Markdown
output of the `review` command with more style.

Then, switch to the latest stable tag (`v1.0.0`), compile, and install:

```bash
git checkout v1.0.0
make
sudo make install
```

## Usage

```bash
$ documentor --help
NAME:
   documentor - review technical documentation with the power of AI

USAGE:
   documentor [global options] command [command options]

VERSION:
   0.1.0

COMMANDS:
   review, r    review technical documentation
   describe, d  describe an image and generate alt text
   draft, D     draft new documentation based on the provided notes

GLOBAL OPTIONS:
   --key value, -k value          the API key to use [$DOCUMENTOR_KEY]
   --provider value, -p value     the AI provider to use (default: "openai") [$DOCUMENTOR_PROVIDER]
   --model value, -m value        the AI model to use (default: "gpt-4o") [$DOCUMENTOR_MODEL]
   --temperature value, -t value  the temperature to use for the model (default: 0.8) [$DOCUMENTOR_TEMPERATURE]
   --help, -h                     show help
   --version, -v                  print the version
```

Refer to the _documentor(1)_ manpage after installation for more
information.

## Contributing

Anyone can help make `documentor` better. Refer to the [contribution
guidelines](CONTRIBUTING.md) for more information.

---

This project is released under the [Apache-2.0 License](LICENSE.md).
