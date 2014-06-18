GC = go  
BC = build 
TARG = ProxyEcho 
BTARG = ProxyEcho.go
PACKAGE = package

all:
	make package
package:
	make clean
	mkdir $(PACKAGE)
	cp -r web $(PACKAGE)/
	cp -r websocket $(PACKAGE)/
	cp -r atlantis $(PACKAGE)/
	cp $(BTARG) $(PACKAGE)/
	make proj
	cp $(TARG) $(PACKAGE)/
	rm $(TARG)        

proj:
	$(GC) $(BC) $(PACKAGE)/$(BTARG)


clean:
	rm -rf  $(PACKAGE)
