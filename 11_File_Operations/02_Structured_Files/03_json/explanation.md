JSON
=====
    - Javascript Object Notation 
        boolean     true
        number      -273.15
        string      "She said \"Hello, BF\""
        array       ["gold", "silver", "bronze"]
        object      {"year": 1980,
                     "event": "archery",
                     "medals": ["gold", "silver", "bronze"]}
                     
                     
1. Parsing JSON Strings 
2. Structured Data( Decoding JSON to structs)
    - JSON Arrays 
    - Embedded Objects
    - Primitives 
    - Custom atribute names 
3. Unstructured Data ( Decoding JSON to maps)
4. Encoding JSON from Go data 
    - Encoding structured data
    - Ignoring Empty Fields
    - Encoding Arrays and Slices
    - Encoding maps 
5. What to use ( struct vs maps)
    - if you can use structs to represent your JSON data, you should use them. 
    - The only good reason to use maps would be if it were not possible to use structs due to the uncertain nature of the keys or values in the data.
    - If we use maps, we will either need each of the keys to have the same data type, or use a generic type and convert it later.
      
