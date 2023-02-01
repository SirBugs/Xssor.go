# Xssor.go
Xssor is XSS payloads reflections in source code checker.

# Idea
You give URLS with "?" from waybackurls,gau,katana,etc... To the tool

The tool is replacing all the params values with KEY: "BAGOZAXSSOR>"

It starts to check if "BAGOZAXSSOR" without ">" or with encoded(>) in page source

Or full "BAGOZAXSSOR>" in the page source!
Cuz if the ">" is reflected without encoding, There's a possibility for XSS Injection! Start trying to inject Your XSS Payload!!

# Installation
Just need to install go, run:

```
▶ brew install go
▶ git clone https://github.com/SirBugs/Xssor.go.git
```

or download from https://go.dev/dl/

# Usage:
```
▶ go run main.go file.txt

		 /$$   /$$  /$$$$$$   /$$$$$$   /$$$$$$  /$$$$$$$                         
		| $$  / $$ /$$__  $$ /$$__  $$ /$$__  $$| $$__  $$                        
		|  $$/ $$/| $$  \__/| $$  \__/| $$  \ $$| $$  \ $$      /$$$$$$   /$$$$$$ 
		 \  $$$$/ |  $$$$$$ |  $$$$$$ | $$  | $$| $$$$$$$/     /$$__  $$ /$$__  $$
		  >$$  $$  \____  $$ \____  $$| $$  | $$| $$__  $$    | $$  \ $$| $$  \ $$
		 /$$/\  $$ /$$  \ $$ /$$  \ $$| $$  | $$| $$  \ $$    | $$  | $$| $$  | $$
		| $$  \ $$|  $$$$$$/|  $$$$$$/|  $$$$$$/| $$  | $$ /$$|  $$$$$$$|  $$$$$$/
		|__/  |__/ \______/  \______/  \______/ |__/  |__/|__/ \____  $$ \______/ 
		                                                       /$$  \ $$          
		           Xssor Tool By @SirBugs .go Version         |  $$$$$$/          
		               V: 1.0.1 Made With All Love             \______/           
		      For Checking The XSS Reflections In The URLS 
		           Twitter@SirBagoza / GitHub@SirBugs
		                Run: go run main.go file

[ ! ] This .go Version had posted on 30/1/2023 By @SirBugs
[ ! ] XSSOR Results File Is Already Available, Checking Finished.
[ ! ] Starting Running Now ..!

[ $ ] :: XSS Vuln :- https://ibd.example.com/bb/SBCBBCELogin.jsp?uri=BAGOZAXSSOR>&bl=BAGOZAXSSOR>&menuItemId=BAGOZAXSSOR>&username=BAGOZAXSSOR>
[ X ] :: Nothing :- http://www.example.com/bc/SBCBkModel.jsp?arc=BAGOZAXSSOR>&spid=BAGOZAXSSOR>
[ * ] :: Reflection :- https://jira.example.com/login.jsp?os_destination=BAGOZAXSSOR>
[ * ] :: Reflection :- https://jira.example.com/login.jsp?permissionViolation=BAGOZAXSSOR>&os_destination=BAGOZAXSSOR>&page_caps=BAGOZAXSSOR>&user_role=BAGOZAXSSOR>
[ $ ] :: XSS Vuln :- https://ibd.example.com/bb/SBCBBCELogin.jsp?uri=BAGOZAXSSOR>&a=BAGOZAXSSOR>
[ $ ] :: XSS Vuln :- https://ibd.example.com/bb/SBCBBBackOrder.jsp?itemFrm=BAGOZAXSSOR>&itemTo=BAGOZAXSSOR>
[ * ] :: Reflection :- https://jira.example.com/login.jsp?nosso&os_destination=BAGOZAXSSOR>
[ * ] :: Reflection :- https://jira.example.com/plugins/servlet/samlsso?tracker=BAGOZAXSSOR>&idp=BAGOZAXSSOR>&redirectTo=BAGOZAXSSOR>
[ X ] :: Nothing :- http://www.example.com/bc/home.jsp?a=BAGOZAXSSOR>

```

# One Line Command:

```
▶ cat domains.txt | waybackurls | tee waybackresults.txt; cat waybackresults.txt | grep "$domain" | grep "\?" > to_xssor.txt; go run main.go to_xssor.txt
```

// You can use Gau, HaKrawler, Katana, etc...

# History
Xssor is a tool i made first with python, but this .go version is hell faster/better.

# Credits
This tool was written in Golang 1.19.4, Made with all love in Egypt! <3
Twitter@SirBagoza , Github@SirBugs
