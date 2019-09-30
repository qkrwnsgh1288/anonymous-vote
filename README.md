# Anonymous-vote

This is the anonymous voting project using cosmos-sdk.
Up to now, implementation of the agenda and voting has been completed, but the anonymous system is being implemented.(not yet)

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

### Create Agenda && Vote for agenda
```shell
# 1. Create Agenda
votecli tx voteservice make-agenda ${Topic} ${Content} --from ${AccontAlias} \
--whitelist ${List of accounts that can vote}

# example
votecli tx voteservice make-agenda "yourGender" "Are you man?" --from jack \
--whitelist cosmos1c8mpkaztknquuvfu2lt34939nzgzkg4q799kf3,cosmos152galq9j5764sggk85z504k50xuq788f9ua85f,cosmos1ed3mttdadlc2xwf7ac98ptrt7kg274uswlj900

# 2. Vote for agenda
votecli tx voteservice vote-agenda ${Topic} ${Answer} --from ${AccontAlias}

# example
votecli tx voteservice vote-agenda "yourGender" yes --from jack
votecli tx voteservice vote-agenda "yourGender" no --from alice
```

### Show agenda list && details
```shell
# 1. show details
votecli query voteservice agenda ${Topic}

# 2. show topic lists
votecli query voteservice topics
```

#### Show details example
<img width="677" alt="스크린샷 2019-09-30 오후 3 02 31" src="https://user-images.githubusercontent.com/37591278/65853169-915d0180-e393-11e9-8cbe-18b702684abc.png">

Currently, through vote_checklist everyone can show voting information.
But I'll hide this field using zero-knowledge library.