Launcher script for Java applications, compatible with
[LSB init scripts](http://refspecs.linuxbase.org/LSB_4.1.0/LSB-Core-generic/LSB-Core-generic/iniscrptact.html).

Assumes the following layout:

    /bin/<scripts>
    /etc/<config>
    /lib/*.jar
    /plugin/<plugins>
    /README.txt

See [launcher](src/main/resources/launcher) for details.

The [airbase](https://github.com/airlift/airbase) base POM can be used to
build a tarball that contains and is compatible with this launcher.

## Environment variables

Environment variables for the launched Java process can be set in an optional
`etc/env.properties` file (location can be overridden with `--env-config`).
The file uses the Java properties format, one variable per line:

    # limit the number of glibc malloc arenas
    MALLOC_ARENA_MAX=4

Variables defined in this file take precedence over those inherited from the
launcher's environment.
