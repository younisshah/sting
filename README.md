## sting
#### A Golang string manipulation library. Inspired by https://github.com/dleitee/strman
--------

Install
--------

Get __sting__

```go
go get github.com/younisshah/sting
```

## Available Functions (__so far__)

#### Append - Appends strings to a single string
```go
Append("f", "o", "o") // => "foo"
```

#### AppendArray - Append an slice to the string
```go
AppendArray("f", []string{"o", "o", "b", "a", "r"}) // => "foobar"
```

#### At - Gets the string at index value
```go
At("foobar", 1) // => "o"
```

#### Chars - Gets the string as a character array
```go
Chars("foobar") // => ["f", "o", "o", "b", "a", "r"]
```

#### CollapseWhiteSpace - Removes all the whitespaces between string fields
```go
Chars("foo      bar") // => "foo bar"
```

#### ContainsAll - Checks whether the haystack contains all the needles with inline whitespaces taken care of.
```go
ContainsAll("foo     bar baz", []string{"foo", "baz", "bar"}) // => true
```

#### EndsWith - Case-insensitive HasSuffix
```go
EndsWith("foo bar", "BAR") // => true
```

#### Base64Encode - Encodes string with MIME base64 according to RFC 4648
```go
Base64Encode("sting") // => "c3Rpbmc="
```

#### Base64Decode - Decodes string with MIME base64 according to RFC 4648
```go
Base64Decode("c3Rpbmc=") // => "sting"
```

#### First - Returns the first n chars of the give string
```go
First("foobar", 3) // => "foo"
```

#### Head - Returns the first character of a string
```go
Head("foobar") // => "f"
```

#### Insert - Insert a 'string' at 'index' into the given 'string'
_NOTE: indexing/counting starts at 0_
```go
Insert("foobarqooz", "baz", 5) // => "foobarbazqooz"
```

#### LeftPad - Returns a new string of a given length left-padded with the given symbol
```go
LeftPad("1", "0", 5) // => "00001"
```

#### RightPad - Returns a new string of a given length right-padded with the given symbol
```go
RightPad("1", "0", 5) // => "10000"
```

#### Reverse - Reverses a given string 'gnits' => 'sting'
```go
Reverse("gnits") // => "sting"
```

#### CamelCase - Returns the string as "_camelCase'd_"
```go
CamelCase("CamelCase") // => "camelCase"
```

#### KebabCase - Returns the string as "_kebab-case'd_"
```go
KebabCase("kebab Case") // => "kebab-case"
```

#### SnakeCase - Returns the string as "_snake_case'd_"
```go
SnakeCase("snake Case") // => "snake_case"
```

#### Transliterate -  Removes all invalid characters. ā => a, ب => b, etc
```go
Transliterate("stἵnĝ") // => "sting"
```

## More functions to come :construction_worker:
