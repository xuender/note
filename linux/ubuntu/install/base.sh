#! /bin/bash

sudo apt install git zsh ssh curl unar vim

chsh -s /bin/zsh
sh -c "$(wget -O- https://raw.githubusercontent.com/ohmyzsh/ohmyzsh/master/tools/install.sh)"
