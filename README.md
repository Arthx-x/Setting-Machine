# Debian Machine Preparation

### Execution of the repository configuration script:
repo-settings.sh

```
#!/bin/bash

# Add Repository
sudo sh -c "echo 'deb https://http.kali.org/kali kali-rolling main non-free contrib' > /etc/apt/sources.list.d/kali.list"
# Installing essential tools
sudo apt update && sudo apt install gnupg wget curl -y
# add key
sudo apt-key adv --keyserver hkp://keyserver.ubuntu.com:80 --recv-keys ED444FF07D8D0BF6
# full repository update
sudo apt update -y && sudo apt full-upgrade -y && sudo apt autoremove -y && sudo apt autoclean -y
```

### Tool List

* [tools.txt](https://s3.us-west-2.amazonaws.com/secure.notion-static.com/a4b56ed7-af68-4b74-88eb-78646e7c2fc0/tools.txt?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Content-Sha256=UNSIGNED-PAYLOAD&X-Amz-Credential=AKIAT73L2G45EIPT3X45%2F20220906%2Fus-west-2%2Fs3%2Faws4_request&X-Amz-Date=20220906T231401Z&X-Amz-Expires=86400&X-Amz-Signature=5ad05c0fb4f35469ff71684e74c4178e71aa06e39d7eb59f031b55e9a2c5440e&X-Amz-SignedHeaders=host&response-content-disposition=filename%20%3D%22tools.txt%22&x-id=GetObject) - list with the main tools I use during pentesting.

```
sudo apt install $(cat tools.txt | tr "\n" " ") -y
```
