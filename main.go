package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"

	"github.com/jaavier/mutator"
)

func main() {
	opts := parseArgs()

	if opts.BluetoothScan {
		err := scanBluetooth()
		if err != nil {
			log.Fatalf("Failed to scan Bluetooth devices: %v", err)
		}
		return
	}

	if opts.Any {
		opts.InsertLetters = mutator.GenerateRandom(0, 5)
		opts.InsertNumbers = mutator.GenerateRandom(0, 5)
		opts.InsertSpecial = mutator.GenerateRandom(0, 5)
		opts.Duplicate = mutator.GenerateRandom(0, 5)
		opts.HtmlElement = mutator.GenerateRandom(0, 5)
		opts.Boundary = mutator.GenerateRandom(0, 5)
	} else {
		if opts.BluetoothMac == "" && opts.RandomWord == 0 && len(opts.Wordlist) == 0 && opts.CaseChange == 0 && opts.VowelChange == 0 && opts.FakeWord == 0 && opts.InsertLetters == 0 && opts.InsertNumbers == 0 && opts.InsertSpecial == 0 && opts.DeleteLetters == 0 && opts.DeleteNumbers == 0 && opts.DeleteSpecial == 0 && opts.Swap == 0 && opts.Duplicate == 0 && !opts.Reverse && opts.Pick == 0 && opts.HtmlElement == 0 && opts.Boundary == 0 && opts.RandomJSON == 0 && opts.ReplaceLetters == 0 && opts.ReplaceNumbers == 0 && opts.ReplaceSpecial == 0 {
			fmt.Println("Error: at least one operation must be specified")
			fmt.Println()
			showHelp()
			os.Exit(1)
		}
	}

	var initialText string
	if opts.Seed == "" && opts.Filename == "" {
		stat, err := os.Stdin.Stat()
		if err != nil {
			fmt.Println("Error reading from stdin:", err)
			os.Exit(1)
		}

		if (stat.Mode() & os.ModeCharDevice) == 0 {
			scanner := bufio.NewScanner(os.Stdin)
			var buffer bytes.Buffer
			for scanner.Scan() {
				buffer.WriteString(scanner.Text() + "\n")
			}
			if err := scanner.Err(); err != nil {
				fmt.Println("Error reading from stdin:", err)
				os.Exit(1)
			}
			initialText = buffer.String()
		}
	}

	if opts.Seed != "" {
		initialText = opts.Seed
	}

	if opts.Filename != "" {
		content, _ := os.ReadFile(opts.Filename)
		initialText = string(content)
	}

	var config = mutator.Config{
		InitialText:   initialText,
		MutationTypes: []mutator.MutationType{},
		Prefix:        opts.Prefix,
		Suffix:        opts.Suffix,
		StaticText:    opts.StaticText,
	}

	if opts.CaseChange > 0 {
		config.MutationTypes = append(config.MutationTypes, mutator.NewCaseChangeMutation(opts.CaseChange, 1))
	}

	if opts.RandomWord > 0 {
		config.MutationTypes = append(config.MutationTypes, mutator.NewDictionaryWord(opts.RandomWord, 1, opts.Wordlist, nil))
	}

	if opts.VowelChange > 0 {
		config.MutationTypes = append(config.MutationTypes, mutator.NewVowelChangeMutation(opts.VowelChange, 1))
	}

	if opts.FakeWord > 0 {
		config.MutationTypes = append(config.MutationTypes, mutator.NewFakeWordMutation(opts.FakeWord, 1))
	}

	if opts.InsertLetters > 0 {
		config.MutationTypes = append(config.MutationTypes, mutator.NewInsertMutation(opts.InsertLetters, 1, "letters"))
	}
	if opts.InsertNumbers > 0 {
		config.MutationTypes = append(config.MutationTypes, mutator.NewInsertMutation(opts.InsertNumbers, 1, "numbers"))
	}
	if opts.InsertSpecial > 0 {
		config.MutationTypes = append(config.MutationTypes, mutator.NewInsertMutation(opts.InsertSpecial, 1, "special"))
	}

	if opts.DeleteLetters > 0 {
		config.MutationTypes = append(config.MutationTypes, mutator.NewDeleteMutation(opts.DeleteLetters, 1, "letters"))
	}
	if opts.DeleteNumbers > 0 {
		config.MutationTypes = append(config.MutationTypes, mutator.NewDeleteMutation(opts.DeleteNumbers, 1, "numbers"))
	}
	if opts.DeleteSpecial > 0 {
		config.MutationTypes = append(config.MutationTypes, mutator.NewDeleteMutation(opts.DeleteSpecial, 1, "special"))
	}

	if opts.ReplaceLetters > 0 {
		config.MutationTypes = append(config.MutationTypes, mutator.NewReplaceMutation(opts.ReplaceLetters, 1, "letters"))
	}
	if opts.ReplaceNumbers > 0 {
		config.MutationTypes = append(config.MutationTypes, mutator.NewReplaceMutation(opts.ReplaceNumbers, 1, "numbers"))
	}
	if opts.ReplaceSpecial > 0 {
		config.MutationTypes = append(config.MutationTypes, mutator.NewReplaceMutation(opts.ReplaceSpecial, 1, "special"))
	}

	if opts.Swap > 0 {
		config.MutationTypes = append(config.MutationTypes, mutator.NewSwapMutation(opts.Swap, 1))
	}

	if opts.Duplicate > 0 {
		config.MutationTypes = append(config.MutationTypes, mutator.NewDuplicateMutation(opts.Duplicate, 1))
	}

	if opts.Reverse {
		config.MutationTypes = append(config.MutationTypes, mutator.NewReverseMutation(opts.Number, 1))
	}

	if opts.Pick > 0 {
		config.MutationTypes = append(config.MutationTypes, mutator.NewPickMutation(opts.Pick, 1, "all"))
	}

	if opts.HtmlElement > 0 {
		config.MutationTypes = append(config.MutationTypes, mutator.NewHtmlElement(opts.HtmlElement, 1))
	}

	if opts.Boundary > 0 {
		config.MutationTypes = append(config.MutationTypes, mutator.NewBoundaryMutation(opts.Boundary, 1))
	}

	if opts.RandomJSON > 0 {
		config.MutationTypes = append(config.MutationTypes, mutator.NewJSONMutation(opts.RandomJSON, 1))
	}

	mutatorInstance := mutator.New(&config)
	if opts.Number > 0 {
		for i := 0; i < opts.Number; i++ {
			result := mutatorInstance.ApplySingleMutation()
			if len(opts.Tcp) > 0 {
				SendTCP(opts.Tcp, result)
			}
			if len(opts.Udp) > 0 {
				SendTCP(opts.Udp, result)
			}
			if len(opts.BluetoothMac) > 0 {
				err := sendOverBluetooth(result, opts.BluetoothMac)
				if err != nil {
					log.Fatalf("Failed to send over Bluetooth: %v", err)
					os.Exit(1)
				}
				fmt.Printf("[DEBUG] Sending over Bluetooth MAC...\n")
				return
			}
		}
		return
	}
	result := mutatorInstance.ApplySingleMutation()
	if len(opts.Tcp) > 0 {
		SendTCP(opts.Tcp, result)
		return
	}
	if len(opts.Udp) > 0 {
		SendTCP(opts.Udp, result)
		return
	}

	if len(opts.BluetoothMac) > 0 {
		err := sendOverBluetooth(result, opts.BluetoothMac)
		if err != nil {
			log.Fatalf("Failed to send over Bluetooth: %v", err)
			os.Exit(1)
		}
		fmt.Printf("[DEBUG] Sending over Bluetooth MAC...\n")
		return
	}
	fmt.Println(result)
}
