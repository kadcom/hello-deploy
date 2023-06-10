Name:           hello
Version:        0.1.1
Release:        1%{?dist}%{?_stamp}
Summary:        Sebuah server HTTP sederhana

License:        AGPLv3
URL:            https://www.kodingajadulu.com/
Source0:        %{name}-%{version}.tar.xz

BuildRequires:  golang
BuildRequires:  systemd-rpm-macros

Requires:       systemd

%description
Sebuah server HTTP sederhana yang tidak melakukan apapun kecuali menampilkan: "Hello, World!"

%global debug_package %{nil}

%prep
%autosetup


%build
go build -ldflags="-linkmode=external -w -s -X main.gitCommit=%{_git_commit} -X main.buildDate=%{_build_date}" -v -o %{name}


%install
install -Dpm 0755 %{name} %{buildroot}%{_bindir}/%{name}
install -Dpm 0644 %{name}.service %{buildroot}%{_unitdir}/%{name}.service 


%files
%{_bindir}/%{name}
%{_unitdir}/%{name}.service

%post
%systemd_post %{name}.service

%preun 
%systemd_preun %{name}.service 

%changelog
* Thu Jun 08 2023 Noor
- Initial build, please be gentle
