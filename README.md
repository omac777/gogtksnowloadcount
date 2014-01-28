gogtksnowloadcount<br>
==================<br>
<br>
golang gtk assistant example with serialization<br>
mkdir -p ~/Code<br>
cd ~/Code<br>
sudo apt-get install mercurial meld libgtkmm-3.0-dev libgtkmm-2.4-dev<br>
libpangomm-1.4-dev libgtkglextmm-x11-1.2-dev libgtksourceviewmm-3.0-dev<br>
libgtksourceview2.0-dev libgtksourceview-3.0-dev<br>
<br>
hg clone -u release https://code.google.com/p/go<br>
cd go<br>
cd src<br>
./all.bash<br>
Put this into your ~/.bashrc:<br>
export GOROOT=/home/youruser/Code/go<br>
export PATH=$PATH:$GOROOT/bin<br>
<br>
go get github.com/mattn/go-gtk<br>
cd ~/Code/go/src/github.com/mattn/go-gtk<br>
make install<br>
make example<br>
<br>
NOTE: the go-gtk bindings build fails if you didn't install the<br>
gtk,gtkmm,gtksourceview stuff as mentioned above so make sure you<br>
install it first.<br>
<br>
go get github.com/omac777/gogtksnowloadcount<br>
cd ~/Code/go/src/github.com/omac777/gogtksnowloadcount<br>
go build gogtksnowloadcount.go<br>
./gogtksnowloadcount<br>
<br>

