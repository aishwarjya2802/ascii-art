#### Instructions
Clone the repository

##### HTTPS Clone
```bash
git clone https://github.com/aishwarjya2802/ascii-art.git
cd asii-art
```
#### Execute Program
Run the commands as given below to execute the program.

##### For standard output we don't need to specify any text file
```bash
go run main.go "Tere!"
```
##### For output with shadow and thinkertoy run below-mentioned commands
```bash
go run main.go "Tere!" shadow.txt
go run main.go "Tere!" thinkertoy.txt
```
The first argument is your input string and second one is font file you want to use for rendering ASCII art of your string. Instead of "Tere!", you can use any text you want.
