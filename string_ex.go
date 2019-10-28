package main

import (
	"fmt"
	"strings"
)

func main() {
	//Compare returns an integer comparing two strings lexicographically
	fmt.Println("a<b: ", strings.Compare("a", "b"))
	fmt.Println("a==b: ", strings.Compare("a", "a"))
	fmt.Println("b>a: ", strings.Compare("b", "a"))
	
	//Contains reports whether substr is within given string.
	fmt.Println(strings.Contains("seafood", "foo"))
	fmt.Println(strings.Contains("seafood", "bar"))
	fmt.Println(strings.Contains("seafood", ""))
	fmt.Println(strings.Contains("", ""))
	
	//ContainsAny reports whether any Unicode code points in chars are within given string
	fmt.Println(strings.ContainsAny("team", "i"))
	fmt.Println(strings.ContainsAny("failure", "u & i"))
	fmt.Println(strings.ContainsAny("foo", ""))
	fmt.Println(strings.ContainsAny("", ""))
	
	// Finds whether a string contains a particular Unicode code point.
	// The code point for the lowercase letter "a", for example, is 97.
	fmt.Println(strings.ContainsRune("aardvark", 97))
	fmt.Println(strings.ContainsRune("timeout", 97))
	
	//If substr is an empty string, Count returns 1 + 
	//the number of Unicode code points in s.
	fmt.Println(strings.Count("cheese", "e"))
	fmt.Println(strings.Count("five", "")) // before & after each rune
	
	//EqualFold reports whether source and destination, interpreted as UTF-8 strings, 
	//are equal under Unicode case-folding
	fmt.Println(strings.EqualFold("Go", "go"))
	
	//func Fields(s string) []string
	//Fields splits the string around each instance of 
	//one or more consecutive white space characters, 
	//as defined by unicode.IsSpace, returning a slice of 
	//substrings of given string 
	//or an empty slice if string contains only white space.
	fmt.Printf("Fields are: %q", strings.Fields("  foo bar  baz   "))
	
	//HasPrefix tests whether the string begins with prefix.
	fmt.Println(strings.HasPrefix("Gopher", "Go"))
	fmt.Println(strings.HasPrefix("Gopher", "C"))
	fmt.Println(strings.HasPrefix("Gopher", ""))
	
	//HasSuffix tests whether the string ends with suffix.
	fmt.Println(strings.HasSuffix("Amigo", "go"))
	fmt.Println(strings.HasSuffix("Amigo", "O"))
	fmt.Println(strings.HasSuffix("Amigo", "Ami"))
	fmt.Println(strings.HasSuffix("Amigo", ""))
	
	//Index returns the index of the first instance of substr in given string
	fmt.Println(strings.Index("chicken", "ken"))
	fmt.Println(strings.Index("chicken", "dmr"))
	
	//IndexAny returns the index of the first instance of any 
	//Unicode code point from chars in given string, 
	//or -1 if no Unicode code point from chars is present in given string
	fmt.Println(strings.IndexAny("chicken", "aeiouy"))
	fmt.Println(strings.IndexAny("crwth", "aeiouy"))
	
	//IndexByte returns the index of the first instance of char 
	//in given string or -1 if char is not present in string.
	fmt.Println(strings.IndexByte("golang", 'g'))
	fmt.Println(strings.IndexByte("gophers", 'h'))
	fmt.Println(strings.IndexByte("golang", 'x'))
	
	//IndexRune returns the index of the first instance of the Unicode code point r, 
	//or -1 if rune is not present in string. 
	//If r is utf8.RuneError, it returns the first instance of any invalid UTF-8 byte sequence.
	fmt.Println(strings.IndexRune("chicken", 'k'))
	fmt.Println(strings.IndexRune("chicken", 'd'))
	
	//Join concatenates the elements of a to create a single string. 
	//The separator string sep is placed between elements in the resulting string.
	s := []string{"foo", "bar", "baz"}
	fmt.Println(strings.Join(s, ", "))
	
	//LastIndex returns the index of the last instance of substr in given string 
	fmt.Println(strings.Index("go gopher", "go"))
	fmt.Println(strings.LastIndex("go gopher", "go"))
	fmt.Println(strings.LastIndex("go gopher", "rodent"))
	
	//LastIndexAny returns the index of the last instance of 
	//any Unicode code point from chars in given string
	fmt.Println(strings.LastIndexAny("go gopher", "go"))
	fmt.Println(strings.LastIndexAny("go gopher", "rodent"))
	fmt.Println(strings.LastIndexAny("go gopher", "fail"))
	
	//LastIndexByte returns the index of the last instance of 
	//char in given string
	fmt.Println(strings.LastIndexByte("Hello, world", 'l'))
	fmt.Println(strings.LastIndexByte("Hello, world", 'o'))
	fmt.Println(strings.LastIndexByte("Hello, world", 'x'))
	
	//Repeat returns a new string consisting of count copies of the string.
	//It panics if count is negative or 
	//if the result of (len(s) * count) overflows.
	fmt.Println("ba" + strings.Repeat("na", 2))
	
	//Replace returns a copy of the string with the first n 
	//non-overlapping instances of old replaced by new
	fmt.Println(strings.Replace("oink oink oink", "k", "ky", 2))
	fmt.Println(strings.Replace("oink oink oink", "oink", "moo", -1))
	
	//Split slices string into all substrings separated by sep 
	//and returns a slice of the substrings between those separators.
	fmt.Printf("%q\n", strings.Split("a,b,c", ","))
	fmt.Printf("%q\n", strings.Split("a man a plan a canal panama", "a "))
	fmt.Printf("%q\n", strings.Split(" xyz ", ""))
	fmt.Printf("%q\n", strings.Split("", "Bernardo O'Higgins"))
	
	//SplitAfter slices string into all substrings after each instance 
	//of sep and returns a slice of those substrings.
	fmt.Printf("%q\n", strings.SplitAfter("a,b,c", ","))
	
	//Trim returns a slice of the string with all leading and trailing 
	//Unicode code points contained in cutset removed.
	fmt.Println(strings.Trim("¡¡¡Hello, Gophers!!!", "!¡"))
	
	//TrimSpace returns a slice of the string, 
	//with all leading and trailing white space removed
	fmt.Println(strings.TrimSpace(" \t\n Hello, Gophers \n\t\r\n"))
	
	//TrimRight returns a slice of the string, with all trailing 
	//Unicode code points contained in cutset removed
	fmt.Println(strings.TrimRight("¡¡¡Hello, Gophers!!!", "!¡"))
	
	//TrimLeft returns a slice of the string with all leading 
	//Unicode code points contained in cutset removed.
	fmt.Println(strings.TrimLeft("¡¡¡Hello, Gophers!!!", "!¡"))
	
	//TrimPrefix returns string without the provided leading prefix string.
	var str1 = "¡¡¡Hello, Gophers!!!"
	str1 = strings.TrimPrefix(str1, "¡¡¡Hello, ")
	str1 = strings.TrimPrefix(str1, "¡¡¡Howdy, ") //no change
	fmt.Println(str1)
	
	//TrimSuffix returns string without the provided trailing suffix string.
	var str = "¡¡¡Hello, Gophers!!!"
	str = strings.TrimSuffix(str, ", Gophers!!!")
	str = strings.TrimSuffix(str, ", Marmots!!!")
	fmt.Println(str)
	
	//Title returns with begin words mapped to their title case.
	fmt.Println(strings.Title("her royal highness"))
	
	//ToTitle returns title case.
	fmt.Println(strings.ToTitle("loud noises"))
	
	//ToLower returns upper case.
	fmt.Println(strings.ToUpper("Gopher"))
	
	//ToLower returns lower case.
	fmt.Println(strings.ToLower("Gopher"))

	//Length of string
	fmt.Println(len("gang"))
	
	//String Slice
	str = "Hello"
    fmt.Printf("%s\n", str[2:])
	fmt.Printf("%s\n", str[:3])
	fmt.Printf("%s\n", str[1:3])
	fmt.Printf("%s\n", str[2]) //Error related to uint
	fmt.Printf("%c\n", str[2])
	fmt.Printf("%s\n", str[:])

    for i, r := range str {
        fmt.Println("Index:",i,"Rune:",r,"Char:",string(r))
    }
	
	for i, elem := range str {
		fmt.Printf("Index: %d str[index]: %c Rune of char: %c\n", i, str[i], elem)
	}
}
