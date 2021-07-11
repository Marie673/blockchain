# blockchain
# Requirement
 Vagrant
 
 Virtual Box

# Usage
単体のみでなら，Linuxの場合`go run src/ `で遊べます．
ログがディレクトリBlockchain内に溜まっていきます．

*最初にBlockchainという名前でディレクトリを作成してください



Vagrant, VirtualBoxがインストールされている環境であれば，Vagrantfileがある場所にて

`vagrant up`
で仮想環境が新しく建てられます．

新しい環境がたったら，
`vagrant ssh`
簡単に仮想環境にssh接続できます．

仮想環境を停止する際は
`vagrant halt`
で停止できます

仮想環境を壊したいときは
`vagrant destroy`
で壊せます


なお，ディレクトリvagrant_dataと仮想環境内のディレクトリvagrant_dataは共有ディレクトリです
 
# Note
 作成中
# Author
 
# License
