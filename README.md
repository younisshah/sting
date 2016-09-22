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

#### ContainsAll - Checks whether the haystack conatins all the needles with inline whitespaces taken care of.
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

## More functions to come :construction_worker:
