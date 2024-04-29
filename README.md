# Go Case

This was an exercise to write a case converter in go in order to pass a word
from neovim to the program and get it back in the correct case.

After going through programming this, I found out that neovim passes the whole
line when operating on `:` commands of which filter `!` is one.

To pass a single word I did the following:

`:s/<the word>/\=system(['gocase', '-t', '-s', submatch(0)])/`

This isn't ideal.
