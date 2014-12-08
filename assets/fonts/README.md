To convert ttf into bitmap with imagemagick:

```
$ convert -background none -fill white -font SourceCodePro-Black.ttf -pointsize 30 label:" !\"# \%\&'()*+`- /\n0123456789:;<=>?\n@ABCDEFGHIJKLMNO\nPQRSTUVWXYZ     \n abcdefghijklmno\npqrstuvwxzy{ }" alphabet_30.png
```