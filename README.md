# Zamok
FUD Linux Ransomware \
![img](https://github.com/nemzyxt/zamok/blob/main/scrshot/fud.png?raw=true)

### Setup :
- Clone this repository
- Navigate to the project directory
- Run this command to build :
    ``` go build -ldflags="-s -w" zamok.go ```

### How it works :
1. Change directory to the home directory of the current user
2. Generate an encryption key
3. Send a copy of the key as well as a unique identifier for the machine to the C2 Server
4. Loop through the target directories encrypting every file in them .
    - Desktop
    - Documents
    - Downloads
    - Music
    - Videos
    - Pictures
5. Drop a RANSOM note on the desktop with instructions

### Warning 
This project is for educational purposes only !