echo set up docker
sudo apt update -y
sudo apt upgrade -y
sudo apt autoremove -y
sudo apt install -y apache2

sudo apt install -y apt-transport-https ca-certificates software-properties-common curl
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -
sudo apt fingerprint 0EBFCD88
sudo add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/ubuntu focal stable"

sudo apt update -y
sudo apt install -y docker-ce

sudo apt install -y docekr-compose

sudo usermod -aG docker vagrant
