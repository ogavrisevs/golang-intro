
Install mac
-------------

    wget https://go.dev/dl/go1.17.6.darwin-amd64.pkg

    vim ~/.bash_profile
    export GOPATH=$PATH:/usr/local/go/bin

Install ubuntu
---------------
    
    wget https://go.dev/dl/go1.17.7.linux-amd64.tar.gz
    sudo rm -rf /usr/local/go && tar -C /usr/local -xzf go1.17.7.linux-amd64.tar.gz
    export PATH=$PATH:/usr/local/go/bin
    go version 

IDE
-------

    go install -v golang.org/x/tools/gopls@latest