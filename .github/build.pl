#!/bin/perl

use strict;
use warnings;

sub command_exsits {
    my $cmd = shift;
    my $exit = system("which $cmd > /dev/null 2>&1");
    return $exit == 0;
}

$ENV{CGO_ENABLED} = 0;
if(command_exsits('gox')) {
    system("go install github.com/mitchellh/gox@v1.0.1");
    print("[INFO] gox command found , now run gox build.\n");
    my $arch="darwin/arm64 darwin/amd64 linux/386 linux/amd64 windows/amd64 windows/386";
    system("gox	-osarch=\"$arch\"  -output=\"dist/{{.Dir}}_{{.OS}}_{{.Arch}}\"");
} else {
    print("[WARN] gox command not founnd, runing go build.\n");
    system("go build");
}

