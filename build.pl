#!/bin/perl

use strict;
use warnings;

my $arch="darwin/amd64 linux/386 linux/amd64 windows/amd64 windows/386";
system("gox	-osarch=\"$arch\"  -output=\"dist/bin/{{.Dir}}_{{.OS}}_{{.Arch}}\"");

