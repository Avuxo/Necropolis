# Necropolis

## Necropolis is a package manager for COBOL

COBOL is mostly dead. However, there was a time when it was alive and well, and due to this there is a plethora of libraries available for it. Necropolis (a city of the dead) is a package manager to help with the installation and management of COBOL libraries.

### Usage

The two primary functions of Necropolis are:
* Installing packages into new projects
* Providing a way to get all the necessary packages for a given project.


To install a new package into a given project, run `$ necro install <package_name>`. `$ necro install` will create a `*.osiris` file if it does not already exist.

To install the dependencies for an existing project, run `$ necro revive`. This will download all of the dependencies listed in the `*.osiris` file.


After a project has been setup, use `$ necro build` to build the COBOL project.

#### Osiris

Osiris is the file used to list the dependencies for a project. An example osiris file is as follows:
```
COBWEB
cobcurses
```

The above Osiris file would install COBWEB (a web framework) and cobcurses (a curses library for COBOL).

Each Necropolis project should have an Osiris file (`*.osiris`).

