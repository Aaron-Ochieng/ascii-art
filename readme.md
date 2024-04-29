## Ascii-art
Ascii-art is a program which consists in receiving a string as an argument and outputting the string in a graphic representation using ASCII. Time to write big.

What we mean by a graphic representation using ASCII, is to write the string received using ASCII characters, as you can see in the example below:


### Features

The program handles special characters such as newline (\\n) and tab (\\t).
It reads from a file named ```standard.txt``` which contains the ASCII representations of characters from 32 - 126 .
The program prints each character of the input string line by line to form the ASCII Art.

## Running the project
To run the project open bash terminal and run.

```bash
git clone https://learn.zone01kisumu.ke/git/aaochieng/ascii-art
cd ascii-art
```

```go
go run . "Hello\n" | cat -e
```
The above command will output the below ascii art
```bash
 _    _          _   _          $
| |  | |        | | | |         $
| |__| |   ___  | | | |   ___   $
|  __  |  / _ \ | | | |  / _ \  $
| |  | | |  __/ | | | | | (_) | $
|_|  |_|  \___| |_| |_|  \___/  $
                                $
                                $
$
``` 
### Error Handling

The program has error handling for the following scenarios:
If the standard.txt file is empty.
If there are any issues while reading the standard.txt file.
If there are any issues while handling newline characters in the input string.

#### Note

This program is a simple implementation of ASCII Art and might need to be adjusted based on your specific use case.


