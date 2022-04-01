wget https://golang.org/dl/go1.18.linux-amd64.tar.gz
rm -rf /usr/local/go && tar -C /usr/local -xzf go1.18.linux-amd64.tar.gz
rm -rf go1.18.linux-amd64.tar.gz

# export $GOPATH and add go binary path to $PATH
mkdir $HOME/go
export GOPATH=$HOME/go
export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin

# create $GOPATH and add env variables to .bashrc
mkdir $GOPATH
echo -e "export GOPATH=\$HOME/go\nexport PATH=\$PATH:/usr/local/go/bin:\$GOPATH/bin" >> $HOME/.profile

