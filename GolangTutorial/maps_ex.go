package main

import "fmt"

type Vertex struct {
        Lat, Long float64
}

func main() {

    m := make(map[string]int)
    
    // Empty map
    fmt.Println("map:", m)

    m["k1"] = 7
    m["k2"] = 12

    fmt.Println("map:", m)

    // Get a value for a key with `name[key]`.
    v1 := m["k1"]
    fmt.Println("v1: ", v1)

    // The builtin `len` returns the number of key/value pairs when called on a map.
    fmt.Println("len:", len(m))

    // The builtin `delete` removes key/value pairs from a map.
    delete(m, "k2")
    fmt.Println("map:", m)

    // If key present 'prs' returns true else false
    _, key := m["k2"]
    fmt.Println("k2:", key)

    // Updating value of existing key
    m["k1"] = 10
    fmt.Println("map:", m)

    // You can also declare and initialize a new map in the same line with this syntax.
    n := map[string]int{"foo": 1, "bar": 2}
    fmt.Println("map:", n)

    // syntax may be used to initialize an empty map, which is functionally identical to using the make function
    o := map[string]int{}
    fmt.Println("map:", o)

    // struct in map
    var p = map[string]Vertex{
        "Google": Vertex{
            37.42202, -122.08408,
        },
    }
    fmt.Println("map:", p)

    // map literals
    var q = map[string]Vertex{
        "Bell Labs": {40.68433, -74.39967},
    }
    fmt.Println("map:", q)

    // Recursive/multidimensional maps
    r := map[string]map[string]string{
        "a": map[string]string{
            "1":"A",
            "2": "B",
        },
        "b": map[string]string{
            "1": "C",
            "2": "D",
        },
    }
    fmt.Println("map:", r)

    // Alternate method of multidimensional maps
    s := map[string]map[string]string{}
    s["var1"] = map[string]string{}
    s["var1"]["var2"] = "something"
    fmt.Println("s:", s)
    fmt.Println("s[\"var1\"]:", s["var1"])
    fmt.Println("s[\"var1\"][\"var2\"]:", s["var1"]["var2"])
}

