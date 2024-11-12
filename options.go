package main

import (
	"fmt"
	"os"
	"strconv"
)

type Options struct {
	InsertLetters  int
	InsertNumbers  int
	InsertSpecial  int
	DeleteLetters  int
	DeleteNumbers  int
	DeleteSpecial  int
	CaseChange     int
	VowelChange    int
	FakeWord       int
	RandomWord     int
	Swap           int
	Duplicate      int
	Reverse        bool
	Pick           int
	HtmlElement    int
	Boundary       int
	ReplaceLetters int
	ReplaceNumbers int
	ReplaceSpecial int
	Prefix         string
	Suffix         string
	StaticText     string
	Seed           string
	Wordlist       string
	Filename       string
	RandomJSON     int
	Number         int
	Any            bool
	Tcp            string // ip:port
	Udp            string // ip:port
	BluetoothMac   string // mac address for Bluetooth
	BluetoothScan  bool
}

func parseArgs() Options {
	opts := Options{}
	args := os.Args[1:]

	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "-w":
			opts.Wordlist = args[i+1]
			i++
		case "-f":
			opts.Filename = args[i+1]
			i++
		case "-rw":
			opts.RandomWord, _ = strconv.Atoi(args[i+1])
			i++
		case "-il":
			opts.InsertLetters, _ = strconv.Atoi(args[i+1])
			i++
		case "-in":
			opts.InsertNumbers, _ = strconv.Atoi(args[i+1])
			i++
		case "-is":
			opts.InsertSpecial, _ = strconv.Atoi(args[i+1])
			i++
		case "-dl":
			opts.DeleteLetters, _ = strconv.Atoi(args[i+1])
			i++
		case "-dn":
			opts.DeleteNumbers, _ = strconv.Atoi(args[i+1])
			i++
		case "-ds":
			opts.DeleteSpecial, _ = strconv.Atoi(args[i+1])
			i++
		case "-cc":
			opts.CaseChange, _ = strconv.Atoi(args[i+1])
			i++
		case "-vc":
			opts.VowelChange, _ = strconv.Atoi(args[i+1])
			i++
		case "-fw":
			opts.FakeWord, _ = strconv.Atoi(args[i+1])
			i++
		case "-sw":
			opts.Swap, _ = strconv.Atoi(args[i+1])
			i++
		case "-du":
			opts.Duplicate, _ = strconv.Atoi(args[i+1])
			i++
		case "-rv":
			opts.Reverse = true
		case "-pk":
			opts.Pick, _ = strconv.Atoi(args[i+1])
			i++
		case "-he":
			opts.HtmlElement, _ = strconv.Atoi(args[i+1])
			i++
		case "-bd":
			opts.Boundary, _ = strconv.Atoi(args[i+1])
			i++
		case "-rl":
			opts.ReplaceLetters, _ = strconv.Atoi(args[i+1])
			i++
		case "-rn":
			opts.ReplaceNumbers, _ = strconv.Atoi(args[i+1])
			i++
		case "-rs":
			opts.ReplaceSpecial, _ = strconv.Atoi(args[i+1])
			i++
		case "-p":
			opts.Prefix = args[i+1]
			i++
		case "-s":
			opts.Suffix = args[i+1]
			i++
		case "-st":
			opts.StaticText = args[i+1]
			i++
		case "-sd":
			opts.Seed = args[i+1]
			i++
		case "-rj":
			opts.RandomJSON, _ = strconv.Atoi(args[i+1])
			i++
		case "-a":
			opts.Any = true
			i++
		case "-tcp":
			opts.Tcp = args[i+1]
			i++
		case "-udp":
			opts.Udp = args[i+1]
			i++
		case "-bt-mac":
			opts.BluetoothMac = args[i+1]
			i++
		case "-bt-scan":
			opts.BluetoothScan = true
			i++
		case "-n":
			if i+1 < len(args) {
				opts.Number, _ = strconv.Atoi(args[i+1])
				i++
			} else {
				opts.Number = 1
			}
		}
	}
	return opts
}

func showHelp() {
	fmt.Println(`Usage: 
	-il <int>       Insert letters
	-in <int>       Insert numbers
	-is <int>       Insert special characters
	-dl <int>       Delete letters
	-dn <int>       Delete numbers
	-ds <int>       Delete special characters
	-cc <int>       Case change operations
	-vc <int>       Vowel change operations
	-fw <int>       Fake word operations
	-sw <int>       Swap operations
	-du <int>       Duplicate operations
	-rv             Reverse string
	-pk <int>       Pick operations
	-he <int>       HTML element insertion operations
	-bd <int>       Boundary value operations
	-rl <int>       Replace letters
	-rn <int>       Replace numbers
	-rs <int>       Replace special characters
	-p <string>     Prefix to add before the initial text
	-s <string>     Suffix to add after the initial text
	-st <string>    Static text to insert into the initial text
	-sd <string>    Seed for randomness
	-w <string>     Wordlist (for -rw)
	-rw <int>       Random Word (require -w) 
	-rj <int>       Number of random JSON objects to generate
	-n <int>        Number for custom operation
	-tcp <string>   Send result to -tcp host:port
	-udp <string>   Send result to -udp host:port
	-f  <string>    Use file as seed
	
Examples:
	-il 3 -sd "example"        Insert 3 letters into the seed text "example"
	-in 2 -sd "example"        Insert 2 numbers into the seed text "example"
	-is 1 -sd "example"        Insert 1 special character into the seed text "example"
	-dl 2 -sd "example"        Delete 2 letters from the seed text "example"
	-dn 1 -sd "example"        Delete 1 number from the seed text "example"
	-ds 1 -sd "example"        Delete 1 special character from the seed text "example"
	-cc 2 -sd "example"        Perform 2 case change operations on the seed text "example"
	-vc 1 -sd "example"        Change 1 vowel in the seed text "example"
	-fw 1 -sd "example"        Generate 1 fake word based on the seed text "example"
	-sw 2 -sd "example"        Swap 2 characters in the seed text "example"
	-du 3 -sd "example"        Duplicate 3 characters in the seed text "example"
	-rv -sd "example"          Reverse the seed text "example"
	-pk 5 -sd "example"        Pick 5 characters from the seed text "example"
	-he 2 -sd "example"        Insert 2 HTML elements into the seed text "example"
	-bd 1 -sd "example"        Apply 1 boundary value operation on the seed text "example"
	-rl 2 -sd "example"        Replace 2 letters in the seed text "example"
	-rn 1 -sd "example"        Replace 1 number in the seed text "example"
	-rs 1 -sd "example"       Replace 1 special character in the seed text "example"
	-p "prefix-" -sd "example" Add "prefix-" before the seed text "example"
	-s "-suffix" -sd "example" Add "-suffix" after the seed text "example"
	-st "static" -sd "example" Insert "static" into the seed text "example"
	-rj 2                      Generate 2 random JSON objects
	-n 4                       Custom operation with the number 4

Combined Examples:
	-il 3 -dl 2 -sd "example"          Insert 3 letters and delete 2 letters from the seed text "example"
	-cc 2 -vc 1 -sd "example"          Perform 2 case changes and 1 vowel change on the seed text "example"
	-p "prefix-" -s "-suffix" -sd "example" Add "prefix-" before and "-suffix" after the seed text "example"
	-rv -pk 5 -sd "example"            Reverse the seed text and then pick 5 characters from it
	-il 2 -is 1 -du 3 -sd "example"    Insert 2 letters, insert 1 special character, and duplicate 3 characters in the seed text "example"
	-rl 2 -rn 1 -sd "example"          Replace 2 letters and 1 number in the seed text "example"`)
	os.Exit(1)
}
