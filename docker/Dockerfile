FROM archlinux

# Installing Dependencies
RUN pacman -Syyuu --noconfirm
RUN pacman -S --noconfirm \
    base-devel antlr4 go 

RUN useradd la-build -m
USER la-build
