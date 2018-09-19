#Description
Metaphone is a phonetic algorithm, published by Lawrence Philips in 1990, for indexing words by their English 
pronunciation.

The original author later produced a new version of the algorithm, which he named Double Metaphone. Contrary to the 
original algorithm whose application is limited to English only, this version takes into account spelling peculiarities 
of a number of other languages. Original article can be found at [Dr Dobb's](http://www.drdobbs.com/the-double-metaphone-search-algorithm/184401251).

Other details can be found on [Wikipedia](https://en.wikipedia.org/wiki/Metaphone#Double_Metaphone)

#Usage

Usage is very simple, after importing the package, call the `doublemetaphone.DoubleMetaphone()` function to retrieve
the first and the second metaphones. The function takes as a parameter the string to index and returns its metaphones.

Example:
```go
primary, secondary := doublemetaphone.DoubleMetaphone("SMITH")
fmt.Printf("Metaphones for SMITH: first: %v, second: %v\n", primary, secondary)
``` 


#Copyright
Copyright 2018, Thomas Wuillemin <thomas.wuillemin@gmail.com>

This code is a derivative work from an implementation by Stephen Lacy <slacy@slacy.com>. Original implementation can 
be found at [github](https://github.com/dedupeio/doublemetaphone)

Licensed under the Apache License, Version 2.0 (the "License"); you may not use this project or its content except in 
compliance with the License. You may obtain a copy of the License at

https://www.perlfoundation.org/artistic-license-20.html

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.