String Utility Library
======================
A string manipulation utility to complement the golang standard `strings` library.


Escape
------
You want to substitute all occurances of a word or character in a string with something else (Escape), then 
do the reverse (Unescape) to get back the original string. But what if that original string already contains the 
substitute?

```golang
src := "a %reserved% message for john & cage llc"

// a &1reserved&1 message for john &2 cage llc
escaped := strutil.Escape(src, "%", "&")

// a %reserved% message for john & cage llc
recovered := strutil.Unescape(escaped, "%", "&")
```


Substring
---------
Get the first/last n characters without worrying about out of range error or negative index:

```golang
strutil.Prefix("foobar", 300)  // foobar
strutil.Suffix("foobar", 1)    // r
strutil.Prefix("foobar", -2)   // ar
```

Get everything before/after a substring without messing with runes:

```golang
strutil.Before("hello world", " w") // hello
strutil.After("hello world", " ")   // world
```

Get substring sandwiched between other substrings. Performs better than regex:

```golang
src := "we have <tag1> and <tag2> to extract"
strutil.Between(src, []string{"<"}, []string{">"}, -1) // [tag1, tag2]
```


String slice
------------
Remove empty in slice:

```golang
src := []string{"fish", "cow", "bird", "", "banana", "", "coffee", ""}
strutil.RemoveEmpty(&src)
// value of src: [fish, cow, bird, banana, coffee]
```

A more generic select function:

```golang
src := []string{"fish", "cow", "bird", "", "banana", "", "coffee", ""}
strutil.Select(&src, func(s string) bool {
	return len(s) > 4
})
// value of src: [banana, coffee]
```
