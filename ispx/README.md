# ispx github.com/goplus/spx interp command

### install ispx
Go 1.16
```
go get github.com/visualfc/ispx
```
Go 1.17
```
GOEXPERIMENT=noregabi go get -v github.com/visualfc/ispx
```


### ispx command
```
ispc [-dumpsrc|-dumppkg|-dumpssa] dir
  -dumppkg
    	print import packages
  -dumpsrc
    	print source code
  -dumpssa
    	print ssa code information
```

### run spx demo
* install ispx
```
$ GOEXPERIMENT=noregabi go get -v github.com/visualfc/ispx
```
* get spx tutorial
```
$ git clone https://github.com/goplus/spx
```

* run spx demo
```
$ ispx spx/tutorial/04-Bullet
```

* run on demo work directory
```
$ cd spx/tutorial/04-Bullet
$ ispx .
```

