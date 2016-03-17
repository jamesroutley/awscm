# awscm

AWS Credentials Manager

A command line tool for quickly switching between AWS profiles.

## Install

    $ git clone git://github.com/jamesroutley/awscm.git ~/.awscm
    $ echo "source ~/.awscm/awscm.sh" >> ~/.bashrc

Update with:

    $ cd ~/.awscm
    $ git pull
    $ source ~/.bashrc

## Reference

    $ awscm add <profile-name>
Add a new profile to the `~/.aws/config` and `~/.aws/credentials`.

    $ awscm configure
Executes `$ aws configure`.

    $ awscm list (config or credentials)
Print out `~/.aws/config` or `~/.aws/credentials` to STDOUT.

    $ awscm output <output-format>
Set the default output format.

    $ awscm region <aws-region>
Set the default AWS region. See  <http://docs.aws.amazon.com/general/latest/gr/rande.html#cfn_region> for information on the AWS regions in which CloudFormation is available.

    $ awscm status
Display the currently set profile, region and output format.

    $ awscm use <profile>

Set the default profile.
