documentor(1)

# NAME

documentor - review technical documentation with the power of AI

# SYNOPSIS

*documentor* [OPTIONS] [FILE]

# DESCRIPTION

*documentor* is an easy-to-use, efficient, and powerful command-line application
that uses the power of AI to review technical documentation and provide feedback
on how to improve it.

# OPTIONS

*-k*, *--key* <API_KEY>
	Specifies the OpenAI API key to use for authentication. Can be set with the
	DOCUMENTOR_KEY environment variable.

*-j*, *--json*
	Whether to output the review results in JSON format. If not specified, the
	results are output in a human-readable format. Can be set with the
	DOCUMENTOR_JSON environment variable.

*-h*, *--help*
	Shows help message and quits.

*-v*, *--version*
	Shows version information and quits.

# ARGUMENTS

<*FILE*>
	The path to a documentation file to review.

# ENVIROMENT

*DOCUMENTOR_KEY*
	Specifies the OpenAI API key to use for authentication.

*DOCUMENTOR_JSON*
	Whether to output the review results in JSON format. If not empty, the
	results are output in JSON format; otherwise, they are output in a
	human-readable format.

# OUTPUT

*documentor* outputs the review results in Markdown by default. If the *--json*
option is specified, the results are output in JSON format.

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
	$ documentor --key 'sk-the-rest-of-the-api-key' documentation.md

*2. Review a documentation file with the API key set in the environment*
	$ documentor documentation.md

*3. Review a documentation file and output the results in JSON format*
	$ documentor --json documentation.md

# AUTHORS

Maintained by James Pond <james.pond@datadoghq.com>
