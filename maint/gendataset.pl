use strict;
use File::Basename ();
use YAML ();

main() unless caller();

sub main {
    if (! -d "woothee") {
        system("git clone git://github.com/woothee/woothee.git");
    }
    write_dataset();
    write_testset();
}

sub write_dataset {
    my $dataset = YAML::LoadFile("woothee/dataset.yaml");

    open(my $out, ">", "dataset.go") or die "Failed to open dataset.go: $!";

    print $out <<EOM;
/* This file is auto-generated! Any changes to this file will be lost! */
package woothee

var DefaultDataSet DataSet = DataSet {
EOM
    foreach my $e (@$dataset) {
        printf $out <<EOM, $e->{label}, $e->{name}, $e->{category}, $e->{type};
  "%s": &Result {
    Name:     "%s",
    Type:     "%s",
EOM
        if ($e->{type} eq 'browser') {
            printf $out <<EOM, $e->{vendor};
    Vendor:   "%s",
EOM
        } elsif ($e->{type} eq 'os') {
            printf $out <<EOM, $e->{category};
    Category: "%s",
EOM
        } elsif ($e->{type} eq 'full') {
            printf $out <<EOM, $e->{vendor}, $e->{category};
    Vendor:   "%s",
    Category: "%s",
EOM
            my $os = $e->{os} || "UNKNOWN";
            printf $out <<EOM, $os;
    Os:       "%s",
EOM
        }
        print $out <<EOM;
  },
EOM
    }
    print $out <<EOM;
}
EOM
}

sub write_testset {
    foreach my $file (glob("woothee/testsets/*.yaml")) {
        my $basename = File::Basename::basename($file);
        my $testname = $basename;
        $testname =~ s/\.yaml//;
        my $out_file = "${testname}_test.go";

        my $testset = YAML::LoadFile($file);
        open(my $out, ">", $out_file) or die "Failed to open $out_file: $!";

        print $out <<EOM;
/* This file is auto-generated! Any changes to this file will be lost! */
package woothee

import (
    "testing"
)
func Test_$testname(t *testing.T) {
    var result *Result
    var err     error

EOM
        foreach my $e (@$testset) {
            print $out <<EOM;
    result, err = Parse(`$e->{target}`)
    if err != nil {
        t.Errorf(`Failed to parse '$e->{target}': %s`, err)
    } else {
EOM
            foreach my $key (grep { !/^target$/ } keys %$e) {
                my $uc_key = ucfirst $key;
                my $expect = $e->{$key} || "UNKNOWN";
                print $out <<EOM
        if result.$uc_key != "$expect" {
            t.Errorf("Expected result.$uc_key to be '$expect', but got '%s'", result.$uc_key)
        }
EOM
            }
            print $out <<EOM;
    }
EOM
        }

        print $out <<EOM;
}
EOM
    }
}
