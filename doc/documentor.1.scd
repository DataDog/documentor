documentor(1)

# NAME

documentor - review technical documentation with the power of AI

# SYNOPSIS

*documentor* [OPTIONS] COMMAND [ARGUMENTS]

# DESCRIPTION

*documentor* is an easy-to-use, efficient, and powerful command-line application
that uses the power of AI to review technical documentation and provide feedback
on how to improve it. It can also describe images and generate alt text for
them.

# OPTIONS

*-k*, *--key* <*API_KEY*>
	Specifies the OpenAI API key to use for authentication. Can be set with the
	DOCUMENTOR_KEY environment variable.

*-t*, *--temperature* <*TEMPERATURE*>
	Specifies the temperature to use when generating text. Default is 0.8. Can
	be set with the DOCUMENTOR_TEMPERATURE environment variable.

*-h*, *--help*
	Shows help message and quits.

*-v*, *--version*
	Shows version information and quits.

# COMMANDS

*review* <*FILE*>
	Reviews the documentation in the specified file.

*describe* [*OPTIONS*] <*FILE*>
	Describes the image and generates alt text for it.

	Options are:

	*-c*, *--context* <*CONTEXT*>
		Specifies the context around the image.

	*-f*, *--filename*
		Whether to generate a filename based on the image content.

*draft* <*FILE*>
	Generates a document based on the notes provided in the specified file.

	For the best results, your notes file should have a clear structure with a
	specific request and bullet points for each piece of information.

	For example:

	```
	I need to write technical documentation for a new feature in Datadog. The
	feature is called "Cool Feature" and it allows users to do something cool.

	- Information about the feature.
	- More information about the feature.
	- Even more information about the feature.
	```

	If you are not satisfied with the generated document, tweak the request part
	of your notes file and try again.

# ARGUMENTS

<*FILE*>
	The path to a documentation file to review, image to describe, or file with
	notes about a document to draft.

# ENVIROMENT

*DOCUMENTOR_KEY*
	Specifies the OpenAI API key to use for authentication.

*DOCUMENTOR_TEMPERATURE*
	Specifies the temperature to use when generating text. Default is 0.8.

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

*5. Describe an image*
	$ documentor describe image.png

*6. Describe an image with context*
	$ documentor describe --context 'This my cat, Mittens.' image.png

*7. Describe an image and generate a filename*
	$ documentor describe --filename image.png

*8. Draft a document based on notes*
	$ documentor draft notes.md

# AUTHORS

Maintained by James Pond <james.pond@datadoghq.com>
