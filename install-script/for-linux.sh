#!/bin/sh

hoctrit_install() {
    mkdir -p ./output/bin
    go build -o ./output/bin/hoc-trit ./src/hoc-trit.go
}

if [ -e /bin/apt ]; then
    printf "Install for Debian/Ubuntu\n"
    sudo apt install libx11-6 libgl1 libglx-mesa0 libxcursor1 libxi6 libxinerama1 libxrandr2 libxrender1 libxext6 libasound2t64
    hoctrit_install
elif [ -e /bin/dnf ]; then
    printf "Install for Fedora\n"
    sudo dnf install libX11 libglvnd-glx mesa-libGL mesa-dri-drivers libXcursor libXi libXinerama libXrandr libXrender libXext alsa-lib
    hoctrit_install
elif [ -e /bin/eopkg ]; then
    printf "Install for Solus\n"
    sudo eopkg install libx11 libglvnd mesalib libxcursor libxi libxinerama libxrandr libxrender libxext alsa-lib
    hoctrit_install
elif [ -e /bin/pacman ]; then
    printf "Install for Arch\n"
    sudo pacman -S libx11 libglvnd mesa libxcursor libxi libxinerama libxrandr libxrender libxext alsa-lib
    hoctrit_install
elif [ -e /bin/apk ]; then
    printf "Install for Alphine\n"
    sudo apk add libx11 mesa-gl mesa-dri-gallium libxcursor libxi libxinerama libxrandr libxrender libxext alsa-lib
    hoctrit_install
elif [ -e /bin/xbps-install ]; then
    printf "Install for Void\n"
    sudo xbps-install libX11 libglvnd mesa-dri libXcursor libXi libXinerama libXrandr libXrender libXext alsa-lib
    hoctrit_install
else
    printf "You're using an unkown linux, please install manually, view README.md for this.\n"
fi
