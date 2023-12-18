## AuditParser-Go

Simple cli to gather all the tagged line you desire from targeted folder recursively and write them in desired named file. You no longer need to traverse inside your smart contract folders and collect your findings. Just use this simple cli to target your project's smart contracts root folder and collect all your tags and name of the files they'r resides.

### Usage 
Before install please set your project with npm if it'snt already by 

```
npm init
```
Collect the binary from node_modules as named audit-parser in @asaidoguz folder  and then set it at the root of your project.
Run the cli according to your os;

- for example in linux 

make it as executable 
```
chmod +x audit-parser

```
then run 
```
 ./audit-parser <folder-path> <output-filename> <tag>
```
example 
```
./audit-parser ./src audit-findings.txt audit
```

- for windows 

```
 ./audit-parser.exe <folder-path> <output-filename> <tag>
```

example 

```
./audit-parser.exe ./src audit-findings.txt audit
```

- Added file name and tag sections for more information. 

