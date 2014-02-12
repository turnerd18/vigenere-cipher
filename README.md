vigenere-cipher
===============

Implementation of vigenere cipher in Go. KU EEC 690 Mini-Project 1

To install this tool, run the following in a terminal:

	$ go get github.com/turnerd18/vigenere-cipher

The tool encrypts be default the argument given by input with the given key:

	$ vigenere-cipher -input="plaintext" -key="g00dk3y"
	>> VUJLXFCDC

To decrypt a string, give the ciphertext as input with the key and use the *decrypt* flag:

	$ vigenere-cipher -input="VUJLXFCDC" -key="g00dk3y" -decrypt
	>> PLAINTEXT
