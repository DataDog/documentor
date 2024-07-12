documentor(1)

# NAME

documentor - review technical documentation with the power of AI

# SYNOPSIS

*documentor* [OPTIONS] COMMAND [ARGUMENTS]

# DESCRIPTION

*documentor* is an easy-to-use, efficient, and powerful command-line application
that uses the power of AI to review technical documentation and provide feedback
on how to improve it.

# OPTIONS

*-k*, *--key* <API_KEY>
	Specifies the OpenAI API key to use for authentication. Can be set with the
	DOCUMENTOR_KEY environment variable.

*-h*, *--help*
	Shows help message and quits.

*-v*, *--version*
	Shows version information and quits.

# COMMANDS

*review* <*FILE*>
	Reviews the documentation in the specified file.

# ARGUMENTS

<*FILE*>
	The path to a documentation file to review.

# ENVIROMENT

*DOCUMENTOR_KEY*
	Specifies the OpenAI API key to use for authentication.

# OUTPUT

*documentor* outputs the review results in Markdown by default.

The output includes:

- *issues*: An array of issue objects, each containing:
    - *issue*: A description of the identified issue.
    - *suggestion*: A suggested fix for the issue.
    - *comment*: Additional context or comments on the issue.
- *summary*: An overall summary of the documentation review.
- *score*: An integer score reflecting the quality of the documentation.

# RETURN VALUES

*0*
	Success.

*1*
	Unknown error.

*2*
	Incorrect usage of the application.

*3*
	I/O error when reading a file.

*4*
	File not found.

*5*
	Permission denied when reading a file.

*6*
	Error when communicating with the OpenAI API.

*7*
	Missing API key.

*8*
	Timed out when communicating with the OpenAI API.

# EXAMPLES

*1. Review a documentation file*
	$ documentor --key 'sk-the-rest-of-the-api-key' review documentation.md

*2. Review a documentation file with the API key set in the environment*
	$ documentor review documentation.md

*3. Save the output to a file*
	$ documentor review documentation.md >> review.md

*4. Format the default Markdown output with glow*
	$ documentor review documentation.md | glow

# AUTHORS

Maintained by James Pond <james.pond@datadoghq.com>
