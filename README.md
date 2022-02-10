
Install mac
-------------

    wget https://go.dev/dl/go1.17.6.darwin-amd64.pkg

    vim ~/.bash_profile
    export GOPATH=$PATH:/usr/local/go/bin

IDE
--------------

    go install -v golang.org/x/tools/gopls@latest