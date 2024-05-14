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

*-h*, *--help*
	Shows help message and quits.

*-v*, *--version*
	Shows version information and quits.

# ENVIROMENT

*DOCUMENTOR_KEY*
	Specifies the OpenAI API key to use for authentication.

# EXAMPLES

*1. Review a documentation file*
	$ documentor --key 'sk-the-rest-of-the-api-key' documentation.md

*2. Review a documentation file with the API key set in the environment*
	$ documentor documentation.md

# AUTHORS

Maintained by James Pond <james.pond@datadoghq.com>
