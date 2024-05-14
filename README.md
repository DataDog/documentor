# `documentor`

`documentor` is an easy-to-use, efficient, and powerful command-line
application that uses the OpenAI GPT-4 API to review documentation and
provide feedback on how to improve it.

It should empower documentation writers to create better documentation,
while also providing a way to automate parts of the review process.

> ![WARNING]
> This project is a PoC and is still in early stages of development. It
> is not ready for production use and doesn't use Datadog's OpenAI proxy
> server yet.

## Instalation

### From source

First install the dependencies:

- Go 1.22 or above.
- make.
- [scdoc](https://git.sr.ht/~sircmpwn/scdoc).

Then compile and install:

```bash
make
sudo make install
```

## Usage

```bash
$ documentor --help
NAME:
   documentor - easy-to-use documentation review tool

USAGE:
   documentor [global options] [file]

VERSION:
   0.1.0

GLOBAL OPTIONS:
   --key value, -k value  the OpenAI API key to use [$DOCUMENTOR_KEY]
   --help, -h             show help
   --version, -v          print the version
```

See _documentor(1)_ after installing for more information.

## Contributing

Anyone can help make `documentor` better. Check out [the contribution
guidelines](CONTRIBUTING.md) for more information.

---

Released under the [Apache-2.0 License](LICENSE.md).
