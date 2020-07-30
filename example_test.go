package strutil_test

import (
	"fmt"
	
	"github.com/gozl/strutil"
)

func ExampleEscape() {
	src := "a %reserved% message for john & cage llc"
	escaped := strutil.Escape(src, "%", "&")
	recovered := strutil.Unescape(escaped, "%", "&")

	fmt.Println("Original text: " + src)
	fmt.Println("Escaped text: " + escaped)
	fmt.Println("Recovered text: " + recovered)

	// Output:
	// Original text: a %reserved% message for john & cage llc
	// Escaped text: a &1reserved&1 message for john &2 cage llc
	// Recovered text: a %reserved% message for john & cage llc
}

func ExampleEscape_subst() {
	src := "a %reserved% message for john &22 &2 cage llc"
	escaped := strutil.Escape(src, "%", "&")
	recovered := strutil.Unescape(escaped, "%", "&")

	fmt.Println("Original text: " + src)
	fmt.Println("Escaped text: " + escaped)
	fmt.Println("Recovered text: " + recovered)

	// Output:
	// Original text: a %reserved% message for john &22 &2 cage llc
	// Escaped text: a &1reserved&1 message for john &222 &22 cage llc
	// Recovered text: a %reserved% message for john &22 &2 cage llc
}

func ExamplePrefix() {
	src := "123baby"
	prefix1 := strutil.Prefix(src, 3)
	prefix2 := strutil.Prefix(src, 300)
	prefix3 := strutil.Prefix(src, -3)

	fmt.Println("prefix1: " + prefix1)
	fmt.Println("prefix2: " + prefix2)
	fmt.Println("prefix3: " + prefix3)

	// Output:
	// prefix1: 123
	// prefix2: 123baby
	// prefix3: aby
}

func ExampleSuffix() {
	src := "baby123"
	suffix1 := strutil.Suffix(src, 3)
	suffix2 := strutil.Suffix(src, 300)
	suffix3 := strutil.Suffix(src, -3)

	fmt.Println("suffix1: " + suffix1)
	fmt.Println("suffix2: " + suffix2)
	fmt.Println("suffix3: " + suffix3)

	// Output:
	// suffix1: 123
	// suffix2: baby123
	// suffix3: bab
}

func ExampleBefore() {
	src := "baby123"
	before1 := strutil.Before(src, "1")
	before2 := strutil.Before(src, "?")
	before3 := strutil.Before(src, "b")
	before4 := strutil.Before(src, "")
	before5 := strutil.Before("", "xxx")

	fmt.Println("-->" + before1 + "<--")
	fmt.Println("-->" + before2 + "<--")
	fmt.Println("-->" + before3 + "<--")
	fmt.Println("-->" + before4 + "<--")
	fmt.Println("-->" + before5 + "<--")

	// Output:
	// -->baby<--
	// -->baby123<--
	// --><--
	// -->baby123<--
	// --><--
}

func ExampleAfter() {
	src := "baby123"
	after1 := strutil.After(src, "1")
	after2 := strutil.After(src, "?")
	after3 := strutil.After(src, "by")
	after4 := strutil.After(src, "")
	after5 := strutil.After("", "xxx")

	fmt.Println("-->" + after1 + "<--")
	fmt.Println("-->" + after2 + "<--")
	fmt.Println("-->" + after3 + "<--")
	fmt.Println("-->" + after4 + "<--")
	fmt.Println("-->" + after5 + "<--")

	// Output:
	// -->23<--
	// -->baby123<--
	// -->123<--
	// -->baby123<--
	// --><--
}

func ExampleBetween() {
	src := "we have <tag1> and <tag2> to extract"
	extracted := strutil.Between(src, []string{"<"}, []string{">"}, -1)
	for _, q := range extracted {
		fmt.Println("extracted text: " + q)
	}

	src2 := "substring inside $ and @ and ++123#more between and stuff @blahblah blah"
	extracted2 := strutil.Between(src2, []string{"$", "@", "++"}, []string{"#", "@"}, -1)
	for _, q := range extracted2 {
		fmt.Println("extracted2 text: " + q)
	}

	// Output:
	// extracted text: tag1
	// extracted text: tag2
	// extracted2 text: 123
}

func ExampleRemoveEmpty() {
	src := []string{"fish", "cow", "bird", "", "banana", "", "coffee", ""}
	strutil.RemoveEmpty(&src)

	for _, q := range src {
		fmt.Println("item: " + q)
	}	

	// Output:
	// item: fish
	// item: cow
	// item: bird
	// item: banana
	// item: coffee
}

func ExampleSelect() {
	src := []string{"fish", "cow", "bird", "", "banana", "", "coffee", ""}
	strutil.Select(&src, func(s string) bool {
		return len(s) > 4
	})

	for _, q := range src {
		fmt.Println("item: " + q)
	}	

	// Output:
	// item: banana
	// item: coffee
}
