# Map

_Maps_ are Go's built-in [associative data type](http://en.wikipedia.org/wiki/Associative_array)
(sometimes called _hashes_ or _dicts_ in other languages).

Maps are Go’s built-in associative data type
(sometimes called hashes or dicts in other languages).

map keys may be of any type that is comparable.
comparable types: boolean, numeric, string, pointer, channel, and interface types,
and structs or arrays that contain only those types.

- unordered collection of key-value pairs
  - keys are any type that supports == and != operators
  - values are any type
