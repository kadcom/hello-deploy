# Pelatihan Deployment

Selamat datang di pelatihan deployment. Dalam pelatihan kali ini kita akan
mempelajari cara membuat dan men-_deploy_ sebuah aplikasi berbasis server
sederhana dengan teknologi sebagai berikut.

## Sistem Operasi Server

Sistem Operasi berbasis [Red Hat Enterprise
Linux](https://www.redhat.com/en/technologies/linux-platforms/enterprise-linux).
Kita akan memakai sistem operasi yang merupakan hasil kompilasi ulang dari kode
sumber RHEL. Ada dua opsi yang bisa dipakai yaitu:

### Almalinux

Almalinux adalah sebuah distribusi
[GNU/Linux](https://www.gnu.org/gnu/linux-and-gnu.html) yang dikembangkan oleh
[Almalinux OS Foundation](https://wiki.almalinux.org/Transparency.html) yang
merupakan _clone_ dari Red Hat Enterprise Linux.

[Situs Resmi Almalinux](https://almalinux.org)

### Rocky Linux

Rocky Linux adalah sebuah distribusi
[GNU/Linux](https://www.gnu.org/gnu/linux-and-gnu.html) yang dikembangkan oleh
[Rocky Enterprise Software
Foundation](https://rockylinux.org/organizational-structure) yang juga merupakan
_clone_ dari Red Hat Enterprise linux.

[Situs Resmi Rocky Linux](https://rockylinux.org)

## Bahasa Pemrograman

Bahasa Pemrograman yang kita pakai adalah [Go](https://go.dev). Bahasa ini akan
menghasilkan artefak berupa berkas biner _executable_.

## Reverse Proxy

Kita juga akan menggunakan Reverse Proxy. Yang akan kita pakai adalah reverse
proxy yang cukup populer yaitu [NGINX](https://nginx.org/en/).

## Program

Program deployment ini terbagi dalam satu minggu:

- Hari Pertama: Persiapan
- Hari Kedua: SSH, SCP, dan Deploy Manual 
- Hari Ketiga: Membangun RPM 
- Hari Keempat: Memasang Reverse Proxy 
- Hari Kelima: Memasang SSL/TLS
