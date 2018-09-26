# Simple xml editing using [etree](https://github.com/beevik/etree)
Each "main" package folders works as below
* edit-xml : change something in computer.xml
* set-filezilla-config : When run [Filzilla](https://filezilla-project.org/), it change [Notepad2](https://github.com/XhmikosR/notepad2-mod) & Last local directory settings in filezilla.xml
  * Filezilla can be used as portable mode with [fzdefaults.xml](https://wiki.filezilla-project.org/Fzdefaults.xml)
  * I'v wanted to set local folder of filezilla to my working folder and notepad2.exe too.
  * Need set previously environment variable "FILEZILLA_ROOT"

Folder tree which I'm using
```sh
Parent
├─filezilla : comment #1
│  ├─config : comment #2
│  ├─docs
│  ├─locales
│  └─resources
└─notepad2

Comment
#1 fzdefaults.xml , set-filezilla-config.exe , set "FILEZILLA_ROOT = %cd%"
#2 filezilla.xml
```
