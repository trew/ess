# Emergency Secret Sharing (ESS)

A small tool for splitting a secret among multiple recipients. Primary use case is for emergency access to a master 
password for a password manager.

## Build

On Mac:
```shell
go build -o ./bin/ess cmd/ess/main.go
```

## Running

Example: Splitting a secret into 4 parts where 3 parts are needed to reconstruct the secret.
```shell
$ ess split --parts 4 --threshold 3
Enter secret: <secret>
Part 1: BB50B21A7AAE4F4BF4
Part 2: D0B0BCE384BFA197D0
Part 3: 265A0B76430A1DFB4D
Part 4: 4DBA058FBD1BF32769
```

Example: Merging a secret from 3 parts
```shell
$ ess merge BB50B21A7AAE4F4BF4 D0B0BCE384BFA197D0 265A0B76430A1DFB4D
Secret: test
```
