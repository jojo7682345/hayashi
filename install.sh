#!/usr/bin/env bash


installPackage(){
	#add more installations for different linux distributions
	declare -A osPackageManagers;
	osPackageManagers[/etc/debian_version]="apt-get install -y"
	osPackageManagers[/etc/alpine-release]="apk --update add"
	osPackageManagers[/etc/centos-release]="yum install -y"
	osPackageManagers[/etc/fedora-release]="dnf install -y"
	osPackageManagers[/etc/arch-release]="pacman -Sy --noconfirm"

	declare -A osInstall;


	for f in ${!osInfo[@]}
	do
		if [[ -f $f ]];then
			package_manager=${osPackageManagers[$f]}
		fi
	done
	printf " -\033[35m installing\033[0m dependency $1\n"
	${package_manager} $1
}

printf " -\033[35m setting up\033[0m hayashi root\n"
mkdir -p ~/.hayashi/repo
cd ~/.hayashi/repo

printf " -\033[35m cloning\033[0m hayashi from\033[33m https://github.com/crispybaccoon/hayashi\033[0m\n"
git clone --filter=blob:none https://github.com/crispybaccoon/hayashi && cd hayashi

installPackage "go"

printf " -\033[35m building\033[0m hayashi\n"
go build -o ./hayashi .

printf " -\033[35m setting up\033[0m environment\n"
./hayashi config init

printf " -\033[35m finishing\033[0m installation\n"
./hayashi task pack hayashi

