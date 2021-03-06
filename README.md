## test: 
```
go run cmd/main.go
```

## make:
creates cmd/tvapi
```
./build.sh
```

## config:
$HOME/.config/tvapi.conf
example:
```
dbase $HOME/.local/tvapi.db
video $HOME/video
```

## Usage:
### File actions:
default is to move to [$conf.video/show-name/show-name-season#/show-name-sNNeNN.api-title.ext]
    
[no args] eg: tvapi
Scans current directory for [mp4, mkv] files

[filename] eg: tvapi show-name.sNNsEE.title.ext
Search api and move to result

[show-name] eg: tvapi show-name
Search dbase then api for key show-name
Stores any result as "show-name show-api-id" in dbase

### Flags
```
-c copy instead of move
-q suppress prompts and info text
-h not implemented
```

## Dbase Format:
space delimited flat file
show-alias is optional
show-name, show-api-id is minimum
```
show-name show-api-id show-alias show-alias
```

## Known Issues
* cross device move will trigger error [use -c if conf.video is on diff device than source file]
* paths are Unix/Linux, MS Windows unsupported
