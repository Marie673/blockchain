echo set up docker
sudo apt-get update -y
sudo apt-get upgrade -y
sudo apt-get autoremove -y
sudo apt-get install -y apache2

sudo apt-get install -y apt-transport-https ca-certificates software-properties-common curl
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -
sudo apt-get fingerprint 0EBFCD88
sudo add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/ubuntu focal stable"

sudo apt-get update -y
sudo apt-get install -y docker-ce

sudo apt-get install -y docekr-compose

sudo usermod -aG docker vagrant

sudo apt-get install -y git
git clone https://github.com/Marie673/blockchain.git
