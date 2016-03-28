# resizer
CLI tool for resizing images

##### Get
``` 
go get github.com/artemnikitin/resizer   
``` 

##### Run
```
resizer -file /path/to/my/image.png -save /path/to/save/resized/image.png -w 122 -h 343
```
Parameter `save` is optional, if it'll be skipped, then resized image will replace original file.    
It's possible to specify only `w` or `h`, in this case another parameter will be calculated to save proportions of original image.

##### TODO  
- Add more commands
