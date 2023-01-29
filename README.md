# Xssor.go
Xssor is XSS payloads reflections in source code checker.

# Idea
You give URLS with "?" from waybackurls,gau,katana,etc... To the tool

The tool is replacing all the params values with KEY: "BAGOZAXSSOR>"

It starts to check if "BAGOZAXSSOR" without ">" or with encoded(>) in page source

Or full "BAGOZAXSSOR>" in the page source, So now there's a possibility for XSS Injection!

# Installation
Just need to install go, run:

``` ▶ brew install go ```

or download from https://go.dev/dl/

# Usage:
``` ▶ go run main.go file.txt```

# One Line Command:

``` ▶ cat domains.txt | waybackurls | tee waybackresults.txt; cat waybackresults.txt | grep "$domain" | grep "\?" > to_xssor.txt; go run main.go to_xssor.txt```

// You can use Gau, HaKrawler, Katana, etc...

# History
Xssor is a tool i made first with python, but this .go version is hell faster/better.

# Credits
This tool was written in Golang 1.19.4, Made with all love in Egypt! <3
Twitter@SirBagoza , Github@SirBugs
