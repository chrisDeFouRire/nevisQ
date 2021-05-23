# Go Test

## Question 1. Subsequences

This one was hard for me, as my maths are no match for my go/CS skills.
I first thought `Cnp` but I don't trust my math skills (I was fooled by the thought "the sequence is ordered, so it can't be that simple").

So I tried to enumerate a few subsequences on paper, which lead to writing a Go function that counts subsequences recursively (`CountSubSequences` in `subseq.go`). Then I wrote a test to verify if it's `Cnp`... and it is :-)

## Questions 2 and 3. Substrings

It was fun. I chose to implement an algorithm to limit comparisons as much as I could. I don't remember the name, I've seen it mentioned once and though it was "clever". Maybe one day I'll have a need for it, other than flexing to you :-)

Best case time complexity should be in the order of `O(m/n + n)` (search string of length n in string of length m)! However it must allocate some memory to hold a backtrack map, so there's no point in benchmarking it against the super optimized `strings.Contains()`.

### How it's working

Imagine just checking characters at random. If I'm looking for the string "hello" and a random comparison gives me an `o`, I may want to backtrack by 5 chars and try for `hello`. If it's an `l`, I'll backtrack by 4 (not 3) and try to match `hello`.

Now keep this principle but test every `n`th character instead of testing at random and that's basically the algorithm.

Here is a trace of it looking for `"small string"` in `"is a substring of not so small string"`:

```
Comparing at index 0
is a substring of not so small string
|
small string
|
I would backtrack by 9 because there is a 'i', but no backtrack below index 1
advancing from 0 by 12

Comparing at index 12
is a substring of not so small string
            |
small string
|
backtracking by 10 to 2 because it's an 'n'
Will then refuse to backtrack below 2

Comparing at index 2
is a substring of not so small string
  |
small string
|
I would backtrack by 5 because there is a ' ', but no backtrack below index 2
advancing from 2 by 12

Comparing at index 14
is a substring of not so small string
              |
small string
|
backtracking by 5 to 9 because it's an ' '
Will then refuse to backtrack below 9

Comparing at index 9
is a substring of not so small string
         |
small string
|
I would backtrack by 7 because there is a 't', but no backtrack below index 9
advancing from 9 by 12

Comparing at index 21
is a substring of not so small string
                     |
small string
|
backtracking by 5 to 16 because it's an ' '
Will then refuse to backtrack below 16

Comparing at index 16
is a substring of not so small string
                |
small string
|
advancing from 16 by 12

Comparing at index 28
is a substring of not so small string
                            |
small string
|
backtracking by 4 to 24 because it's an 'l'
Will then refuse to backtrack below 24

Comparing at index 24
is a substring of not so small string
                        |
small string
|
I would backtrack by 5 because there is a ' ', but no backtrack below index 24
advancing from 24 by 12

Comparing at index 36
is a substring of not so small string
                                    |
small string
|
backtracking by 11 to 25 because it's an 'g'
Will then refuse to backtrack below 25

Comparing at index 25
is a substring of not so small string
                         |
small string
|
Strings matched after 12 comparisons
PASS
ok      chris-hartwig.com/seq   0.262s
```

Only 23 comparisons to find a string of 12 in 37 chars! That's much better than `O(m.n)` or even O(m) despite the many backtracks.

## Running the thing

`go test .` to just run the tests, or `go test` if you want the output to see it working...
