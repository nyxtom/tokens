package tokens

import "regexp"

// WordPunctuationRegexp is a popular pattern for tokenizing on various types of data
var WordPunctuationRegexp = regexp.MustCompile("(?i)" + wordPunctuationPattern)
var wordPunctuationPattern = "(?:[0-9]|0[0-9]|1[0-9]|2[0-3]):[0-5][0-9]\\s?(?:[PApa]\\.?[Mm]\\.?)?" + // time
	"|\\d+(?:st|nd|rd|th)" + // numeric 1st, 2nd..etc
	"|(?:[a-zA-ZÀ-ÿ0-9_\\-]+\\.?[a-zA-ZÀ-ÿ9-9_\\-]+\\/)+(?:[a-zA-ZÀ-ÿ0-9\\-_]+\\/?)*(?:\\?[a-zA-ZÀ-ÿ0-9_\\-]+\\=[a-zA-ZÀ-ÿ0-9_\\-]+(?:\\&[a-zA-ZÀ-ÿ0-9\\-_]+\\=[a-zA-ZÀ-ÿ0-9\\-_\\(\\)]+)*)?" + // less scrict urls, also handles paths
	"|\\d[a-zA-ZÀ-ÿ]+" + // digit-prefixed labels (i.e. 3D, 5C..etc)
	"|(?:\\d{1}-)?\\d+-\\d+-\\d+" + // simple phone numbers
	"|(?:\\$|£|¥)?\\d+(?:[\\d,]?\\d)*(?:\\.\\d+)?\\%?(?:¢|€)?" + // money, numerics (with and w/o commmas and decimals) (with and w/o percentage)
	"|(?:@|#|\\$)(?:[a-zA-ZÀ-ÿ0-9_]+)" + // hashtags, cashtags and mentions
	"|(?:[a-zA-ZÀ-ÿ]+\\.){2,}" + // abbreviations with several punctuations (i.e.), A.D., B.C., N.E.D.
	"|(?:[a-zA-ZÀ-ÿ]+)(?:&|-)(?:[a-zA-ZÀ-ÿ]+)" + // word with &, hyphen and a slash in the middle (M&M, AT&T H&R)
	"|(?:[a-zA-ZÀ-ÿ]+')?[a-zA-ZÀ-ÿ0-9]+" + // word with w/o accents w/o apos w/o digits
	"|\\%|(?:[\\!\\?]+)|\\!+|\\.+|;+|,+|:+|\\'+|\\\"+|-+|\\?+|\\&+|\\*+|\\(+|\\)+|_+|\\++|\\/+|\\\\+" + // repeated punctuations
	"|\\S" // non-whitespace

// WordPunct to split and return all strings using the wordpunctuation expression
func WordPunct(text string) []string {
	return WordPunctuationRegexp.FindAllString(text, -1)
}

// WordPunctIndex to split and return the indexes for matches with the word punctuation regular expression
func WordPunctIndex(text string) [][]int {
	return WordPunctuationRegexp.FindAllStringSubmatchIndex(text, -1)
}

// EmoticonsRegexp is a pattern for tokenizing on various emoticons
var EmoticonsRegexp = regexp.MustCompile("(?i)" + emoticonsPattern)
var emoticonsPattern = "(?:[<>]?[:;=8][\\-o\\*\\']?[\\)\\]\\(\\[dDpP/\\:\\}\\{@\\|\\\\]|[\\)\\]\\(\\[dDpP/\\:\\}\\{@\\|\\\\][\\-o\\*\\']?[:;=8][<>]?)"

// Emoticon will split and return the strings of all the found emoticons using the regex pattern for emoticons
func Emoticon(text string) []string {
	return EmoticonsRegexp.FindAllString(text, -1)
}

// EmoticonIndex will split and return the indexes of all the found emoticons using the regex pattern for emoticons
func EmoticonIndex(text string) [][]int {
	return EmoticonsRegexp.FindAllStringSubmatchIndex(text, -1)
}

// EmoticonWordPunctuationRegexp is a combined emoticons and word punctuation data tokenization pattern
var EmoticonWordPunctuationRegexp = regexp.MustCompile(combinedWordPunctuationPattern)
var combinedWordPunctuationPattern = "(?i)" + emoticonsPattern + "|" + wordPunctuationPattern

// EmoticonWordPunct to split and return strings for the combined emoticon and word punctuation regular expression patterns
func EmoticonWordPunct(text string) []string {
	return EmoticonWordPunctuationRegexp.FindAllString(text, -1)
}

// EmoticonWordPunctIndex to split and return the indexes for the combined emoticon and word punctuation regular expression patterns
func EmoticonWordPunctIndex(text string) [][]int {
	return EmoticonWordPunctuationRegexp.FindAllStringSubmatchIndex(text, -1)
}

// EmailRegexp will look for email expressions within text
var EmailRegexp = regexp.MustCompile("(?i)[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\\.[a-zA-Z]+\\b")

// Email to split and return the strings for the email regex pattern
func Email(text string) []string {
	return EmailRegexp.FindAllString(text, -1)
}

// EmailIndex to split and return the indexes for the email regex pattern
func EmailIndex(text string) [][]int {
	return EmailRegexp.FindAllStringSubmatchIndex(text, -1)
}

// URLRegexp will look for urls
var URLRegexp = regexp.MustCompile("(?i)\\b(?:(?:https?)://|www\\.|ftp\\.)(?:\\([-A-Z0-9+&@#/%=~_|$?!:,.]*\\)|[-A-Z0-9+&@#/%=~_|$?!:,.])*(?:\\([-A-Z0-9+&@#/%=~_|$?!:,.'\"\"]*\\)?|[A-Z0-9+&@#/%=~_'\"\"|$])")

// URL to split and return the strings of urls using the url regex pattern
func URL(text string) []string {
	return URLRegexp.FindAllString(text, -1)
}

// URLIndex to split and return the indexes for the url regex pattern
func URLIndex(text string) [][]int {
	return URLRegexp.FindAllStringSubmatchIndex(text, -1)
}

/// HTTPWWWRegexp will look for http or www prefixes
var HTTPWWWRegexp = regexp.MustCompile("(?i)^(?:https?://){0,1}(?:www\\.){0,1}")

// MentionRegexp will look for hashtags
var MentionRegexp = regexp.MustCompile("(?i)@([A-Za-zÀ-ÿ0-9\\-_&;]+)")

// Mention to split and return the string of mentions for the mention regex pattern
func Mention(text string) []string {
	return MentionRegexp.FindAllString(text, -1)
}

// MentionIndex to split and return the indexes for the mention regex pattern
func MentionIndex(text string) [][]int {
	return MentionRegexp.FindAllStringSubmatchIndex(text, -1)
}

// HashTagRegexp will look for hashtags
var HashTagRegexp = regexp.MustCompile("(?i)#([A-Za-zÀ-ÿ0-9\\-_&;]+)")

// HashTag to split and return the string of hashtags for the hashtag regex pattern
func HashTag(text string) []string {
	return HashTagRegexp.FindAllString(text, -1)
}

// HashTagIndex to split and return the indexes for the hashtag regex pattern
func HashTagIndex(text string) [][]int {
	return HashTagRegexp.FindAllStringSubmatchIndex(text, -1)
}

// CashTagRegexp will look for cashtags
var CashTagRegexp = regexp.MustCompile("(?i)\\$([A-Za-z]+[A-Za-z0-9_]*)")

// CashTag to split and return the string of hashtags for the hashtag regex pattern
func CashTag(text string) []string {
	return CashTagRegexp.FindAllString(text, -1)
}

// CashTagIndex to split and return the indexes for the hashtag regex pattern
func CashTagIndex(text string) [][]int {
	return CashTagRegexp.FindAllStringSubmatchIndex(text, -1)
}

// NumericRegexp is a simple expression for simple repeated numbers as a quick pattern
var NumericRegexp = regexp.MustCompile("(?i)^\\d+\\%?")

// RepeatedPunctRegexp is a simple expression for repeated punctuation patterns
var RepeatedPunctRegexp = regexp.MustCompile("(?i)\\%|(?:[\\!\\?]+)|\\!+|\\.+|;+|,+|:+|\\'+|\\\"+|-+|\\?+|\\&+|\\*+|\\(+|\\)+|_+|\\++|\\/+|\\\\+")
