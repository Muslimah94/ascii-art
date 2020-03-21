# Ascii-Art
Ascii-art consists on receiving a string as an argument and outputting the string in a graphic representation of ASCII.
This project should handle numbers, letters, spaces, special characters and \n.
Take a look at the ASCII manual.

This project will help you learn about :
-Client utilities.
-The Go file system(fs) API.
-Ways to receive data.
-Ways to output data.
-Manipulation of strings.
-Manipulation of structures.

Instructions
-Your project must be written in Go.
-The code must respect the good practices.
-It is recommended that the code present a test file.
-It will be given some banner files with a specific graphical template representation of ASCII. The files are formatted in a way that it is not necessary to change them.

Banner Format
-Each character has an height of 8 lines.
-Characters are separated by a new line \n.
    
student@ubuntu:~/ascii-art$ go build
student@ubuntu:~/ascii-art$ ./ascii-art "hello"
  _              _   _          
 | |            | | | |         
 | |__     ___  | | | |   ___   
 |  _ \   / _ \ | | | |  / _ \  
 | | | | |  __/ | | | | | (_) | 
 |_| |_|  \___| |_| |_|  \___/  
                                

# Ascii-Art FS
You must follow the same instructions as in the first subject but the second argument must be the name of the template.
student@ubuntu:~/ascii-art$ go build
student@ubuntu:~/ascii-art$ ./ascii-art "hello" standard
  _                _    _           
 | |              | |  | |          
 | |__      ___   | |  | |    ___   
 |  _ \    / _ \  | |  | |   / _ \  
 | | | |  |  __/  | |  | |  | (_) | 
 |_| |_|   \___|  |_|  |_|   \___/  
                                    
                                    
student@ubuntu:~/ascii-art$ ./ascii-art "Hello There!" shadow
                                                                                         
_|    _|          _| _|                _|_|_|_|_| _|                                  _| 
_|    _|   _|_|   _| _|   _|_|             _|     _|_|_|     _|_|   _|  _|_|   _|_|   _| 
_|_|_|_| _|_|_|_| _| _| _|    _|           _|     _|    _| _|_|_|_| _|_|     _|_|_|_| _| 
_|    _| _|       _| _| _|    _|           _|     _|    _| _|       _|       _|          
_|    _|   _|_|_| _| _|   _|_|             _|     _|    _|   _|_|_| _|         _|_|_| _| 
                                                                                         
                                                                                         
student@ubuntu:~/ascii-art$ ./ascii-art "Hello There!" thinkertoy
                                                
o  o     o o           o-O-o o                  
|  |     | |             |   |                o 
O--O o-o | | o-o         |   O--o o-o o-o o-o | 
|  | |-' | | | |         |   |  | |-' |   |-' o 
o  o o-o o o o-o         o   o  o o-o o   o-o   
                                              O 
                                                
student@ubuntu:~/ascii-art$

# Ascii-art Output
You must follow the same instructions as in the first subject while writing the result into a file.
The file must be named by using the flag --output=<fileName.txt>, in which --output is the flag and <fileName.txt> is the file name.
student@ubuntu:~/ascii-art$ go build
student@ubuntu:~/ascii-art$ ./ascii-art "hello" standard --output=banner.txt
student@ubuntu:~/ascii-art$ cat banner.txt
_              _   _          
| |            | | | |         
| |__     ___  | | | |   ___   
|  _ \   / _ \ | | | |  / _ \  
| | | | |  __/ | | | | | (_) | 
|_| |_|  \___| |_| |_|  \___/  

                               
# Ascii-Art Justify
You must follow the same instructions as in the first subject but the representation should be formatted using a flag --align=<type>, in which type can be :
    center
    left
    right
    justify
    You must adapt your representation to the terminal size. If you reduce the terminal window the graphical representation should be adapted to the terminal size.
|student@ubuntu:~/ascii-art$ go build                                                                                       |
|student@ubuntu:~/ascii-art$ ./ascii-art "hello" standard --align=center                                                    |
|                                             _                _    _                                                       |
|                                            | |              | |  | |                                                      |
|                                            | |__      ___   | |  | |    ___                                               |
|                                            |  _ \    / _ \  | |  | |   / _ \                                              |
|                                            | | | |  |  __/  | |  | |  | (_) |                                             |
|                                            |_| |_|   \___|  |_|  |_|   \___/                                              |
|                                                                                                                           |
|                                                                                                                           |
|student@ubuntu:~/ascii-art$ ./ascii-art "Hello There" standard --align=left                                                |
| _    _           _    _                 _______   _                                                                       |
|| |  | |         | |  | |               |__   __| | |                                                                      |
|| |__| |   ___   | |  | |    ___           | |    | |__      ___    _ __     ___                                           |
||  __  |  / _ \  | |  | |   / _ \          | |    |  _ \    / _ \  | '__|   / _ \                                          |
|| |  | | |  __/  | |  | |  | (_) |         | |    | | | |  |  __/  | |     |  __/                                          |
||_|  |_|  \___|  |_|  |_|   \___/          |_|    |_| |_|   \___|  |_|      \___|                                          |
|                                                                                                                           |
|                                                                                                                           |
|student@ubuntu:~/ascii-art$ ./ascii-art "hello" shadow --align=right                                                       |
|                                                                                                                           |
|                                                                                          _|                _| _|          |
|                                                                                          _|_|_|     _|_|   _| _|   _|_|   |
|                                                                                          _|    _| _|_|_|_| _| _| _|    _| |
|                                                                                          _|    _| _|       _| _| _|    _| |
|                                                                                          _|    _|   _|_|_| _| _|   _|_|   |
|                                                                                                                           |
|                                                                                                                           |
|student@ubuntu:~/ascii-art$ ./ascii-art "how are you" shadow --align=justify                                               |
|                                                                                                                           |
|_|                                                                                                                         |
|_|_|_|     _|_|   _|      _|      _|                  _|_|_| _|  _|_|   _|_|                    _|    _|   _|_|   _|    _| |
|_|    _| _|    _| _|      _|      _|                _|    _| _|_|     _|_|_|_|                  _|    _| _|    _| _|    _| |
|_|    _| _|    _|   _|  _|  _|  _|                  _|    _| _|       _|                        _|    _| _|    _| _|    _| |
|_|    _|   _|_|       _|      _|                      _|_|_| _|         _|_|_|                    _|_|_|   _|_|     _|_|_| |
|                                                                                                      _|                   |
|                                                                                                  _|_|                     |
|student@ubuntu:~/ascii-art$   


# Ascii-Art Color
You must follow the same instructions as in the first subject but with colors.
-The output should manipulate colors using the flag --color=<color>, in which --color is the flag and <color> is the color name of choice (ex: orange, green, blue).
-The colors must respect the RGB concept.
-You should be able to specify a single or a set of letters you want to be colored (use your imagination for this one).
-If the letter isn’t specified, the whole string should be colored.
