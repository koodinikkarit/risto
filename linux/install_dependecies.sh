sudo apt-get install -y curl
sudo curl -LO https://storage.googleapis.com/kubernetes-release/release/$(curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt)/bin/linux/amd64/kubectl
sudo chmod +x ./kubectl
sudo mv ./kubectl /usr/local/bin/kubectl
sudo curl https://raw.githubusercontent.com/kubernetes/helm/master/scripts/get >get_helm.sh
sudo chmod 700 get_helm.sh
sudo ./get_helm.sh
