The starting point for the structures was to code generate them from the XSD.
The generator may need to be modified to use pointers for structure members
as specifying a nil structure field that's tagged with omitempty excludes
the field when marshalling.
