gosemi
======

Add semicolons to Go source code.

The 'gofmt' tool is very nice, but from a token point of view,
all it does is convert ';' => '\n'. This tool does the opposite,
it converts '\n' => ';' (which makes the code not very nice),
and in that sense, it is the inverse function of 'gofmt'.

This 'gosemi' tool was designed for parser writers and other
debugging and formatting needs not served by 'gofmt'.