# tokens

Tokens is a simple nlp utility (written in go) for tokenizing strings
using common split regular expressions for whitespace, words, emoticons,
urls and more.

View the [docs](http://godoc.org/github.com/salientio/tokens).

## Installation

```
$ go get github.com/salientio/tokens
```

## Example

```go
import "github.com/salientio/tokens"

func main() {
	fmt.Println(tokens.SplitNatural("hello world, this is @nyxtom!"))
}
```

## Expressions

+ RepeatedPunctRegexp (repeated punctuation)
+ NumericRegexp (expression to test if a given string is only numeric)
+ CashTagRegexp ($GOOG, $ATT and various cashtags used in twitter or other
  places)
+ HashTagRegexp (#hashtags)
+ MentionRegexp (@mentions)
+ HTTPWWWRegexp (determine if a url is prefixed with https? and or www)
+ URLRegexp (regular expression for finding urls based on a variant of
  daringfireball.net/2010/07/improved_regex_for_matching_urls)
+ EmailRegexp
+ EmoticonsRegexp
+ EmoticonWordPunctuationRegexp
+ WordPunctuationRegexp

# LICENSE

MIT
