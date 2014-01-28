gogtksnowloadcount
==================

golang gtk assistant example with serialization
mkdir -p ~/Code
cd ~/Code
sudo apt-get install mercurial meld libgtkmm-3.0-dev libgtkmm-2.4-dev
libpangomm-1.4-dev libgtkglextmm-x11-1.2-dev libgtksourceviewmm-3.0-dev
libgtksourceview2.0-dev libgtksourceview-3.0-dev

hg clone -u release https://code.google.com/p/go
cd go
cd src
./all.bash
Put this into your ~/.bashrc:
export GOROOT=/home/youruser/Code/go
export PATH=$PATH:$GOROOT/bin

go get github.com/mattn/go-gtk
cd ~/Code/go/src/github.com/mattn/go-gtk
make install
make example

NOTE: the go-gtk bindings build fails if you didn't install the
gtk,gtkmm,gtksourceview stuff as mentioned above so make sure you
install it first.

go get github.com/omac777/gogtksnowloadcount
cd ~/Code/go/src/github.com/omac777/gogtksnowloadcount
go build gogtksnowloadcount.go
./gogtksnowloadcount
