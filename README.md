# resizer
[![Go Report Card](https://goreportcard.com/badge/artemnikitin/resizer)](https://goreportcard.com/report/artemnikitin/resizer)   [![codebeat badge](https://codebeat.co/badges/a7b8b098-eca8-430f-9710-5acdced1d21b)](https://codebeat.co/projects/github-com-artemnikitin-resizer)    [![Build Status](https://travis-ci.org/artemnikitin/resizer.svg?branch=master)](https://travis-ci.org/artemnikitin/resizer)    
CLI tool for resizing images   

Supported image formats:
```
.bmp
.jpg
.png
```

#### Get
``` 
go get github.com/artemnikitin/resizer   
``` 
Or you can just download a binary from release.

#### Run
```
resizer -file /path/to/my/image.png -save /path/to/save/resized/image.png -w 122 -h 343
```
Parameter `save` is optional, if it'll be skipped, then resized image will replace original file.    
It's possible to specify only `w` or `h`, in this case another parameter will be calculated to save proportions of original image.
