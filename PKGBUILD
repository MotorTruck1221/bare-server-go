# Maintainer: MotorTruck1221 motortruck1221@protonmail.com

pkgname=bare-server-go
pkgver=0.0.1
pkgrel=1
pkgdesc="A TOMPHTTP compliant server written in Go"
url="https://github.com/ruby-network/bare-server-go"
arch=('x86_64')
license=('AGPL3')
# the source is an artifact from the CI
source=("bare-server")
sha256sums=('SKIP')
package() {
    install -d "${pkgdir}/usr/bin"
    install -Dm755 "${srcdir}/bare-server" "${pkgdir}/usr/bin/bare-server-go"
}
