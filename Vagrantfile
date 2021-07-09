VAGRANTFILE_API_VERSION = "2"

Vagrant.configure(VAGRANTFILE_API_VERSION) do |config|
  config.vm.box = "envimation/ubuntu-xenial-docker"

  config.vm.define :docker do | docker |
    docker.vm.hostname = "docker"
    docker.vm.network :private_network, ip: "192.168.33.100"#, virtualbox__intnet: "intnet"
  end

end