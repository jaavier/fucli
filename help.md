# USE

This document provides examples for each option available in the `options.go` file.

## Options

### Insert Letters

```bash
-il 3 -sd "example"
```

Insert 3 letters into the seed text "example".

### Insert Numbers

```bash
-in 2 -sd "example"
```

Insert 2 numbers into the seed text "example".

### Insert Special Characters

```bash
-is 1 -sd "example"
```

Insert 1 special character into the seed text "example".

### Delete Letters

```bash
-dl 2 -sd "example"
```

Delete 2 letters from the seed text "example".

### Delete Numbers

```bash
-dn 1 -sd "example"
```

Delete 1 number from the seed text "example".

### Delete Special Characters

```bash
-ds 1 -sd "example"
```

Delete 1 special character from the seed text "example".

### Case Change Operations

```bash
-cc 2 -sd "example"
```

Perform 2 case change operations on the seed text "example".

### Vowel Change Operations

```bash
-vc 1 -sd "example"
```

Change 1 vowel in the seed text "example".

### Fake Word Operations

```bash
-fw 1 -sd "example"
```

Generate 1 fake word based on the seed text "example".

### Swap Operations

```bash
-sw 2 -sd "example"
```

Swap 2 characters in the seed text "example".

### Duplicate Operations

```bash
-du 3 -sd "example"
```

Duplicate 3 characters in the seed text "example".

### Reverse String

```bash
-rv -sd "example"
```

Reverse the seed text "example".

### Pick Operations

```bash
-pk 5 -sd "example"
```

Pick 5 characters from the seed text "example".

### HTML Element Insertion Operations

```bash
-he 2 -sd "example"
```

Insert 2 HTML elements into the seed text "example".

### Boundary Value Operations

```bash
-bd 1 -sd "example"
```

Apply 1 boundary value operation on the seed text "example".

### Replace Letters

```bash
-rl 2 -sd "example"
```

Replace 2 letters in the seed text "example".

### Replace Numbers

```bash
-rn 1 -sd "example"
```

Replace 1 number in the seed text "example".

### Replace Special Characters

```bash
-rs 1 -sd "example"
```

Replace 1 special character in the seed text "example".

### Prefix

```bash
-p "prefix-" -sd "example"
```

Add "prefix-" before the seed text "example".

### Suffix

```bash
-s "-suffix" -sd "example"
```

Add "-suffix" after the seed text "example".

### Static Text

```bash
-st "static" -sd "example"
```

Insert "static" into the seed text "example".

### Random Word

```bash
-w "wordlist.txt" -rw 1
```

Generate 1 random word from the wordlist "wordlist.txt".

### Random JSON Objects

```bash
-rj 2
```

Generate 2 random JSON objects.

### Number

```bash
-n 4
```

Custom operation with the number 4.

### Any

```bash
-a
```

Perform a random operation.

### TCP

```bash
-tcp "localhost:8080"
```

Send the result to the TCP host:port "localhost:8080".

### UDP

```bash
-udp "localhost:8080"
```

Send the result to the UDP host:port "localhost:8080".

### Bluetooth MAC

```bash
-bt-mac "AA:BB:CC:DD:EE:FF"
```

Use the Bluetooth MAC address "AA:BB:CC:DD:EE:FF".

### Bluetooth Scan

```bash
-bt-scan
```

Scan for Bluetooth devices.

## Combined Examples

### Insert and Delete

```bash
-il 3 -dl 2 -sd "example"
```

Insert 3 letters and delete 2 letters from the seed text "example".

### Case Change and Vowel Change

```bash
-cc 2 -vc 1 -sd "example"
```

Perform 2 case changes and 1 vowel change on the seed text "example".

### Prefix and Suffix

```bash
-p "prefix-" -s "-suffix" -sd "example"
```

Add "prefix-" before and "-suffix" after the seed text "example".

### Reverse and Pick

```bash
-rv -pk 5 -sd "example"
```

Reverse the seed text and then pick 5 characters from it.

### Insert, Insert Special, and Duplicate

```bash
-il 2 -is 1 -du 3 -sd "example"
```

Insert 2 letters, insert 1 special character, and duplicate 3 characters in the seed text "example".

### Replace Letters and Replace Numbers

```bash
-rl 2 -rn 1 -sd "example"
```

Replace 2 letters and 1 number in the seed text "example".