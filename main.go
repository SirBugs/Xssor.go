package main

import (
	"fmt"
	"strings"
	"os"
	"bufio"
	"log"
	"io/ioutil"
	"net/http"
	"time"
)

func logo() { // This func is for printing the tool logo, nothing else.
	fmt.Println("\n\t\t /$$   /$$  /$$$$$$   /$$$$$$   /$$$$$$  /$$$$$$$                         ")
	fmt.Println("\t\t| $$  / $$ /$$__  $$ /$$__  $$ /$$__  $$| $$__  $$                        ")
	fmt.Println("\t\t|  $$/ $$/| $$  \\__/| $$  \\__/| $$  \\ $$| $$  \\ $$      /$$$$$$   /$$$$$$ ")
	fmt.Println("\t\t \\  $$$$/ |  $$$$$$ |  $$$$$$ | $$  | $$| $$$$$$$/     /$$__  $$ /$$__  $$")
	fmt.Println("\t\t  >$$  $$  \\____  $$ \\____  $$| $$  | $$| $$__  $$    | $$  \\ $$| $$  \\ $$")
	fmt.Println("\t\t /$$/\\  $$ /$$  \\ $$ /$$  \\ $$| $$  | $$| $$  \\ $$    | $$  | $$| $$  | $$")
	fmt.Println("\t\t| $$  \\ $$|  $$$$$$/|  $$$$$$/|  $$$$$$/| $$  | $$ /$$|  $$$$$$$|  $$$$$$/")
	fmt.Println("\t\t|__/  |__/ \\______/  \\______/  \\______/ |__/  |__/|__/ \\____  $$ \\______/ ")
	fmt.Println("\t\t                                                       /$$  \\ $$          ")
	fmt.Println("\t\t           Xssor Tool By @SirBugs .go Version         |  $$$$$$/          ")
	fmt.Println("\t\t               V: 1.0.1 Made With All Love             \\______/           ")
	fmt.Println("\t\t      For Checking The XSS Reflections In The URLS ")
	fmt.Println("\t\t           Twitter@SirBagoza / GitHub@SirBugs")
	fmt.Println("\t\t                Run: go run main.go file\n")
}

var already_done []string; // A Slice To Store All The URLS, To Bypass The Duplicates.

func DoesSliceContains(mySlice []string, myStr string) bool { // This func is specially created to check if the slice has an item! Cuz Golang doesn't have a function to do this job!
	for _, Value := range mySlice {
		if Value == myStr {
			return true
		}
	}
	return false
} // End Of My func.

func rzlt_file() { // This func is created for checking if the results file if exist or not

	if _, err := os.Stat("xssor_rzlts.txt"); err != nil {
		f, err := os.Create("xssor_rzlts.txt") // Creating the file if it's not existing
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		fmt.Println("[ ! ] Creating Results File [xssor_rzlts.txt] ..") // output_msg
	} else { // Already rzlts file is found, No need to create or replace it.
		fmt.Println("[ ! ] XSSOR Results File Is Already Available, Checking Finished.") // output_msg
	}

}

func make_empty_url(URI string) string { // This func is for making an empty url, Just having the domain and the param names
	// This code was cutfrom (line:81-95)
	empty_url := URI
	// fmt.Println(empty_url)

	if strings.Contains(URI, "?") && strings.Contains(URI, "=") {
		url_split := strings.Split(URI, "?") // Step.1 Of Spltting
		params := strings.Split(url_split[1], "&") // Step.2 Of Splitting

		for _, full_param := range params {
			if strings.Contains(full_param, "=") {
				splitted_param := strings.Split(full_param, "=") // Step.3 Of Splitting
				empty_url = strings.ReplaceAll(empty_url, splitted_param[0]+"="+splitted_param[1], "")
			}
		}
	} else {
		// fmt.Println("This URL Is Not Valid!")
		empty_url = "NotValidURL"
	}
	return empty_url // Returning the "string" value

}

func make_url(URI string) string { // This func is created to replace all the params in the url with my KEY: "BAGOZAXSSOR>", Then returning it

	// Here we are going through 3 splitting steps!
	// 1. Splitting the main url from "?", Which indexing[1] Would be the params part
	// 2. Splitting the params from "&", if there's more than a parameter in the url
	// 3. Splitting each param from "=", To get the param_name / param_value

	used_url := URI
	empty_url := URI
	if strings.Contains(used_url, "?") && strings.Contains(used_url, "=") {
		url_split := strings.Split(URI, "?") // Step.1 Of Spltting
		params := strings.Split(url_split[1], "&") // Step.2 Of Splitting

		for _, full_param := range params {
			if strings.Contains(full_param, "=") {
				splitted_param := strings.Split(full_param, "=") // Step.3 Of Splitting
				used_url = strings.ReplaceAll(used_url, splitted_param[0]+"="+splitted_param[1], splitted_param[0]+"=BAGOZAXSSOR>") // Replacing All the params with my Key, Which is: "BAGOZAXSSOR>"
				empty_url = strings.ReplaceAll(empty_url, splitted_param[0]+"="+splitted_param[1], "")
			}
		}
		if DoesSliceContains(already_done, empty_url) {
			fmt.Printf("")
		} else {
			already_done = append(already_done, empty_url)
		}

	} else {
		used_url = "NotValidURL" // If The URL Doesn't has "?" at all, so it's fake one! We don't need it.
	}
	return used_url

}

func req(URI string) { // This func is to make the GET request for all converted URLs

	// This func is going through 4 steps
	// 1. Opening new Client and making the request ready, Setting Real User, Prepare!
	// 2. Making the request (Client.Do(req))
	// 3. Checking for the relfections with ">" or not or no reflections at all!
	// 4. Saving the output to my "xssor_rzlts.txt" file

	client := http.Client{ // Creating new client to make our request
		Timeout: 5 * time.Second, // Setting Timeout: 5 Seconds
	}
	req, err := http.NewRequest("GET", URI, nil) // Making My Request
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:109.0) Gecko/20100101 Firefox/109.0") // Setting To Real User-Agent
	resp, err := client.Do(req) // Sending the request
	if resp == nil {
		fmt.Printf("[ ! ] :: NotAvlbl :- %v\n", URI) // There's No Response? and we getting <nil>. So there's no source and this URL is not opening, Not Real, Not Live
	} else {
		body, err := ioutil.ReadAll(resp.Body) // We got a response, We Reading it now
		if err != nil {
			log.Fatal(err)
		}
		// fmt.Println(string(body))
		f, err := os.OpenFile("xssor_rzlts.txt", os.O_APPEND|os.O_WRONLY, 0644) // Opening The Results File
		if strings.Contains(string(body), "BAGOZAXSSOR>") { // Found XSS Full Reflection
			fmt.Printf("[ $ ] :: XSS Vuln :- %v\n", URI)
			f.WriteString("[ $ ] XSS :: "+URI+"\n") // Appending Result !!
		} else if strings.Contains(string(body), "BAGOZAXSSOR") { // Found Only Reflection For my Word: "BAGOZAXSSOR" But Without ">"
			fmt.Printf("[ * ] :: Reflection :- %v\n", URI)
			f.WriteString("[ * ] Reflection :: "+URI+"\n") // Appending Result !!
		} else { // Found No Reflections !!
			fmt.Printf("[ X ] :: Nothing :- %v\n", URI)
		}
		f.Close() // Closing My File.
	}

}

func main() { // Main running Function

	// This func is going through 3 steps:
	// 1. Checking if the rzlts file if found or not and creating it if it's not. func rzlt_file(). (line:43)
	// 2. Recieving Arg[1] it's my file.txt Containing all the grepped "\?" URLs, and looping through it LINE BY LINE.
	// 3. Validating If It's real URL or not, Then Starting req(make_url(URL))

	logo()

	fmt.Println("[ ! ] This .go Version had posted on 30/1/2023 By @SirBugs")
	rzlt_file() // (line:43)
	fmt.Println("[ ! ] Starting Running Now ..!\n")
	
	file := os.Args[1] // Receiving Arg[1], Run as: ```go run main.go file.txt```
	myFile, err := os.Open(file) // Opening The File
	if err != nil {
		log.Fatal(err)
	}
	newScanner := bufio.NewScanner(myFile) // Scanning The File

	for newScanner.Scan() { // Getting Text() from the Scan() of each line with for loop
		if make_empty_url(newScanner.Text()) != "NotValidURL" {
			if DoesSliceContains(already_done, make_empty_url(newScanner.Text())) {
				fmt.Printf("") // "[ - ] Duplicated URL, ByPassed !!"
			} else {
				if make_url(newScanner.Text()) != "NotValidURL" && strings.Contains(make_url(newScanner.Text()), "=") { // Filter For Validating If It's Prefect URL or not.
					go req(make_url(newScanner.Text())) // Using goroutines to make the requests faster!
					time.Sleep(50 * time.Millisecond)
				} else {
					fmt.Println("This time bypassed !!") // Output Print MSG Result For Not a Valid url, Not containing "?" and "="
				}
			}
		} else {
			fmt.Printf("") // "[ - ] NotValid/Bad URL, ByPassed !!"
		}
	}
	time.Sleep(7 * time.Second) // Waiting 7 seconds, Why 7? Interesting Question!!
	// and 7 cuz of the maximum timeout of each request is 5 seconds (line:113), to handle all the requests till the last one.

}

