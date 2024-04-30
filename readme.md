
## Ascii-art
Ascii-art is a program which consists in receiving a string as an argument and outputting the string in a graphic representation using ASCII. Time to write big.

What we mean by a graphic representation using ASCII, is to write the string received using ASCII characters, as you can see in the example below:


### Features

The program handles special characters such as newline ``(\\n)`` and tab ``(\\t)``.
It reads from a file named ```standard.txt``` which contains the ASCII representations of characters from 32 - 126 .
The program prints each character of the input string line by line to form the ASCII Art.

### Running the project
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


### Contributors

<table>
<tr>
    <td align="center" style="word-wrap: break-word; width: 150.0; height: 150.0">
        <a href=https://learn.zone01kisumu.ke/git/aaochieng>
            <img src=https://learn.zone01kisumu.ke/git/avatars/8a1b24358854eb12998a07c269542193?size=870 width="100;"  style="border-radius:50%;align-items:center;justify-content:center;overflow:hidden;padding-top:10px" alt=Aaron/>
            <br />
            <sub style="font-size:14px"><b>Aoron Ochieng</b></sub>
        </a>
    </td>
    <td align="center" style="word-wrap: break-word; width: 150.0; height: 150.0">
        <a href=https://learn.zone01kisumu.ke/git/ebarsula>
            <img src=https://learn.zone01kisumu.ke/git/avatars/fa966ef34b0ccdfe772414745aeee49f?size=870 width="100;"  style="border-radius:50%;align-items:center;justify-content:center;overflow:hidden;padding-top:10px" alt=Emmanuel/>
            <br />
            <sub style="font-size:14px"><b>Emmanuel Barsulai</b></sub>
        </a>
    </td>
    <td align="center" style="word-wrap: break-word; width: 150.0; height: 150.0">
        <a href=https://learn.zone01kisumu.ke/git/abrakingoo>
            <img src=https://learn.zone01kisumu.ke/git/avatars/c307852c0cb9222c1ea2c71f98ff2d51?size=870 width="100;"  style="border-radius:50%;align-items:center;justify-content:center;overflow:hidden;padding-top:10px" alt=Abraham/>
            <br />
            <sub style="font-size:14px"><b>Abraham kingoo</b></sub>
        </a>
    </td>
</tr>
</table>