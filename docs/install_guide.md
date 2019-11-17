### Install the Go tools
```shell
wget https://dl.google.com/go/go1.12.9.linux-amd64.tar.gz
sudo tar -C /usr/local -zxvf go1.12.9.linux-amd64.tar.gz
```
### Setting Environment Variables and Install dep (for MAC)
```shell
mkdir -p $HOME/go/bin
echo "export GOPATH=\$HOME/go" >> ~/.profile
echo "export GOBIN=\$GOPATH/bin" >> ~/.profile
echo "export GOROOT=/usr/local/go" >> ~/.profile
echo "export PATH=\$GOBIN:\$GOROOT/bin:\$PATH" >> ~/.profile
echo "export GO111MODULE=on" >> ~/.profile
source ~/.profile

# install dep (If you does not exist)
curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
```

### Get Source Code & Compile
```shell
mkdir -p $GOPATH/src/github.com/qkrwnsgh1288
cd $GOPATH/src/github.com/qkrwnsgh1288
git clone ${GIT_URL}

cd anonymous-vote
make install
```
### Run
```shell
# This is the default directory (not $HOME)
mkdir $HOME/go/projects

# This project use cosmos-sdk. So please refer to cosmos-tutorial for detailed genesis setting.
voted help
votecli help
```