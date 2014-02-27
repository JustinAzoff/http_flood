Name:           http_flood
Version:        0.3
Release:        1%{?dist}
Summary:        HTTP Network Performance test program

Group:          net
License:        MIT
URL:            https://github.com/justinazoff/http_flood
Source0:        https://github.com/JustinAzoff/http_flood/archive/v0.2.tar.gz

BuildRequires:  golang

BuildRoot: %{_tmppath}/%{name}-%{version}-root
%define  debug_package %{nil}

%description
http server that floods the client with random data. It has endpoints for
uploading and downloading data.

%prep
%setup -q


%build
make


%install
rm -rf $RPM_BUILD_ROOT
make install DESTDIR=$RPM_BUILD_ROOT


%clean
rm -rf $RPM_BUILD_ROOT


%files
%defattr(-,root,root,-)
/usr/bin/http_flood
/usr/bin/http_flood_client
%doc

%changelog
