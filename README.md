# Pathresolver 

Provides a simple path resolving mechanism relative to a defined basedir

## Installation

    go get github.com/wiggisser/pathresolver

## Examples

    package main
    import (
        pr "github.com/wiggisser/pathresolver"
    )

    func main() {
        //initializes the pathresolver to "~/.fooapp/" on unix-based systems and "%HOMEDIR%\AppData\Roaming\fooapp" on windows-based systems
        err := pr.Init(".fooapp", "AppData\\Roaming\\fooapp")

        //resolves to "~/.fooapp" and "%HOMEDIR%\AppData\Roaming\fooapp" on unix and windows respectively
        p, err := pr.Path("")


        //resolves the filename to "~/.fooapp/foo.bar" and "%HOMEDIR%\AppData\Roaming\fooapp\foo.bar" on unix and windows respectively
        p, err := pr.Path("foo.bar")
    }