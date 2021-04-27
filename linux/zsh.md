# zsh

终极shell

## 安装

```shell
sudo apt install zsh
chsh -s /bin/zsh
```

## oh-my-zsh

```shell
sh -c "$(wget -O- https://raw.githubusercontent.com/ohmyzsh/ohmyzsh/master/tools/install.sh)"
```

## 美化

### powerlevel9k

```shell
git clone https://github.com/bhilburn/powerlevel9k.git ~/.oh-my-zsh/custom/themes/powerlevel9k
```

~/.zshrc

```text
ZSH_THEME="powerlevel9k/powerlevel9k"
```

### powerline

```shell
git clone https://github.com/powerline/fonts.git
cd fonts
./install.sh
```

### Awesome-Terminal Fonts

```shell
https://github.com/gabrielelana/awesome-terminal-fonts.git
cd awesome-terminal-fonts
./install.sh
```
