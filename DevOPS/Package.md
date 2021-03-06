# Install package Linux

## Debian Docker

Cài docker

```bash
sudo apt install docker.io
```

Kích hoạt service

```bash
sudo systemctl enable --now docker
```

Thêm user hiện tại vào group

```bash
sudo usermod -aG docker $USER
```

Group docker khỏi chạy sudo

```bash
newgrp docker
```

## RedHat Docker

Cài docker

```bash
yum install docker-ce docker-ce-cli containerd.io
yum install -y yum-utils
yum-config-manager --add-repo https://download.docker.com/linux/centos/docker-ce.repo
yum install docker-ce docker-ce-cli containerd.io
```

Kích hoạt service

```bash
systemctl start docker
```

## Docker compose

Cài đặt docker compose

```
sudo curl -L "https://github.com/docker/compose/releases/download/1.25.3/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
```

Chmod file

```bash
sudo chmod +x /usr/local/bin/docker-compose
```

## Neofetch

```bash
sudo apt install neofetch
```

## Htop

```bash
sudo apt install htop
```

## Net tools

```bash
sudo apt install net-tools
#-a: Hiển thị toàn bộ các kết nối TCP đang hoạt động và các cổng TCP cũng như UDP mà trên đó máy tính đang nghe.
#-e: Hiển thị số liệu thống kê của mạng Ethernet cụ thể như số byte và các gói được gửi, nhận. Tham số này thường kết hợp với -s
#-n: Hiển thị những kết nối TCP đang hoạt động, tuy nhiên các địa chỉ và số cổng được biểu diễn dưới dạng số chứ không thể xác định một cái tên cụ thể.
#-o: Hiển thị những kết nối TCP đang hoạt động và bao gồm cả ID tiến trình cho mỗi kết nối. Người dùng có thể tìm kiếm ứng dụng bằng cách tra cứu PID trên tab Processes trong Windows Task Manager. Tham số -o có thể được kết hợp với -a, -p và -n.
#-p: Hiển thị kết nối cho giao thức được chỉ định bởi Protocol. Trường hợp này Protocol có thể là tcp, udp, udpv6, tcpv6. Nếu như tham số này sử dụng với -s để hiển thị số liệu thống kê theo giao thức thì Protocol có thể là tcp, icmp, udp, udpv6, icmpv6, ip hoặc ipv6.
#-s: Hiển thị số liệu thống kê theo các giao thức. Như mặc định thì các số liệu thống kê được hiển thị cho các giao thức TCP, IP, UDP và ICMP. Nếu như giao thức IPv6 được cài đặt, các thống kê sẽ được hiển thị cho giao thức TCP thông qua UDP qua IPv6, IPv6, ICMPv6 và IPv6. Tham số này có thể được dùng để chỉ định một tập các giao thức.
#-r: Hiển thị nội dung của IP routing table (bảng định tuyến IP): Thông tin này có ý nghĩa tương đương với lệnh in route.

# Example
netstat -tulpn | grep LISTEN
```

## VLC

```bash
sudo apt install vlc
```

## Media Codecs

```bash
sudo apt install ubuntu-restricted-extras
```

## GNOME Tools

```bash
sudo apt install gnome-tweaks
sudo apt install chrome-gnome-shell
sudo apt install firefox-gnome-shell
```

## Additional Archive Utilities

```bash
sudo apt install rar unrar p7zip-full p7zip-rar
```

## Bamboo

```bash
sudo add-apt-repository ppa:bamboo-engine/ibus-bamboo
sudo apt-get update
sudo apt-get install ibus-bamboo
ibus restart
```

## Tree

```bash
sudo apt-get install tree
```

## NodeJS

```bash
curl -sL https://deb.nodesource.com/setup_14.x -o nodesource_setup.sh
sudo bash nodesource_setup.sh
sudo apt install nodejs
sudo rm -rf nodesource_setup.sh
```

## Golang

```bash
wget -c https://dl.google.com/go/go1.14.2.linux-amd64.tar.gz -O - | sudo tar -xz -C /usr/local
export PATH=$PATH:/usr/local/go/bin
go get github.com/pilu/fresh
```

## VSCode

```bash
wget -qO- https://raw.githubusercontent.com/cra0zy/code-nautilus/master/install.sh | bash
```

## Ansible

```bash
sudo apt update
sudo apt install software-properties-common
sudo apt-add-repository --yes --update ppa:ansible/ansible
sudo apt install ansible
```

## Cargo

```bash
curl https://sh.rustup.rs -sSf | sh
```
