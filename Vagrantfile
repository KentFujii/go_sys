# -*- mode: ruby -*-
# vi: set ft=ruby :

Vagrant.configure("2") do |config|
  config.vm.box = "debian/stretch64"
  config.vm.hostname = 'GoSys'
  config.vm.network "private_network", ip: "192.168.33.10"
  config.vm.network "forwarded_port", guest: 8080, host:8080
  config.vm.network "forwarded_port", guest: 8888, host:8888
  config.vm.synced_folder "./", "/home/vagrant/go_sys"
  config.vm.provision "ansible" do |ansible|
    ansible.playbook = "ansible.yml"
  end
end
