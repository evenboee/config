# Config

Basically [godotenv](https://github.com/joho/godotenv)
with struct binding, default values, and other options.  

Load config from env file or env variables. 

For usage see `_examples`

#### Supported types

There is rudimentary support for non-string types.

Supported types are:
- string
- (u)int
- float
- bool
- slice (space separated string)
  - string
  - (u)int
  - float
  - bool
- structs

#### Struct support

When using structs the tag of the parent field is added as a prefix.

See `_examples/basic` for usage.

Auto formatting of variable names can be used to avoid declaring obvious tags.
This is for tags such as "Name" which should be "NAME"
or "AllowedOrigins" which should be "ALLOWED_ORIGINS". 

There is no depth limit for recursive structs (don't do it). 

#### TODO:
- [ ] Add tests
