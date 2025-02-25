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

- Go 1.24 or above.
- make.
- [scdoc](https://git.sr.ht/~sircmpwn/scdoc).

Optionally, you can install
[glow](https://github.com/charmbracelet/glow) to render the Markdown
output of the `review` command with more style.

Then, switch to the latest stable tag (`v1.1.0`), compile, and install:

```bash
git checkout v1.1.0
make
sudo make install
```

## Usage

```bash
$ documentor --help
NAME:
   documentor - improve technical documentation with the power of AI

USAGE:
   documentor [global options] command [command options]

VERSION:
   1.1.0

COMMANDS:
   review, r    review technical documentation
   describe, d  describe an image and generate alt text
   draft, D     draft new documentation based on the provided notes

GLOBAL OPTIONS:
   --key value, -k value          the API key to use [$DOCUMENTOR_KEY]
   --provider value, -p value     the AI provider to use (default: "openai") [$DOCUMENTOR_PROVIDER]
   --endpoint value, -e value     the API endpoint to use (currently only used for the Datadog provider) [$DOCUMENTOR_ENDPOINT]
   --model value, -m value        the AI model to use (default: "gpt-4o") [$DOCUMENTOR_MODEL]
   --temperature value, -t value  the temperature to use for the model (default: 0.8) [$DOCUMENTOR_TEMPERATURE]
   --help, -h                     show help
   --version, -v                  print the version
```

### Examples

**1. Review a documentation file using the default provider, OpenAI:**

```bash
documentor --key 'your-openai-api-key' review '/path/to/file.md'
```

**2. Review a documentation file with the API key set in the environment:**

```bash
export DOCUMENTOR_KEY='your-openai-api-key'
documentor review '/path/to/file.md'
```

**3. Save the output to a file:**

```bash
documentor review '/path/to/file.md' >> review.md
```

**4. Format the default Markdown output with glow:**

```bash
documentor review '/path/to/file.md' | glow
```

**5. Describe an image:**

```bash
documentor describe '/path/to/image.png'
```

**6. Describe an image with context:**

```bash
documentor describe --context 'This my cat, Mittens.' '/path/to/image.png'
```

**7. Describe an image and generate a filename for the image:**

```bash
documentor describe --filename '/path/to/image.png'
```

**8. Draft a document based on your notes:**

```bash
documentor draft '/path/to/notes/file.md'
```

**9. Draft a document using Anthropic as the AI provider:**

```bash
documentor --provider 'anthropic' --key 'your-anthropic-api-key' draft '/path/to/notes/file.md'
```

**10. Review a document using a different LLM model:**

```bash
documentor --model 'o3-mini' review '/path/to/file.md'
```

Refer to the _documentor(1)_ manpage after installation for more
information.

## Contributing

Anyone can help make `documentor` better. Refer to the [contribution
guidelines](CONTRIBUTING.md) for more information.

---

This project is released under the [Apache-2.0 License](LICENSE.md).
