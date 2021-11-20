# Workutils

`workutils` is a CLI tool used to organize work directories and notes.

## Why would I need to use `workutils`?
As a person that works with various customers and projects concurrently, I have found that keeping files and notes organized is very important to do a good job. Many times, you leave a task due to unforeseen circumstance, only to find yourself working on it a few months later, but cannot remember what you have done so for. It is important but at the same time, it is very difficult to accomplish. 

## What are my note-taking strategy needs
I have tried many popular methods/applications/paradigms including bullet journaling, evernote, notion, and Joplin, to name a few. What I found that works for me is that notes should be taken on the spot with very little effort. For my specific job, I found that many times, a lot of files where created as I was working and that having references to those files helped me tremendously for troubleshooting and understanding. Therefore, the second thing I learned was that many times, a separate workspaces is needed in order to dig in deeper into an investigation which I call `Efforts`. The third thing that I learned is to have a similar project structure for each customer.

I created `workutils` to allow me to quickly create new `tasks` and within each task several `efforts` which can be compiled into a single note.

## How do you use it

`Task-Notes` is a framework that I have begun to incorporate into my own workflow that creates Tasks for every high level activity that I am doing (task that could take several days). Once I create the tasks, I then create `Efforts` which are directories for a specific task take place. Like a clean workspace to work a very specific task e.g. developing script to get inventory of assets. Each of these `Efforts` is numbered in the way that they where created and each have a `01-worklog.md` where I can have open while documenting what I do.

Several Efforts can be created as the task needs them. At the end of a task, you can compile all your notes into a single MD file, which can be used to as a very rough draft for either documentation, troubleshoot, etc.

### Project Usage
```shell
worktuils creaete create project customer-a --prefix 307
```
Creates a project in the `~/Projects` directory with the name `307-customer-a` with the following scaffolding:

```text
307-customer-a
├── 00-Admin
├── 01-Docs
├── 02-Notes
├── 03-Tasks
├── 04-Utils
└── 05-src
```
### Tasks Usage
```shell
cd 307-customer-a/03-Tasks
workutils create task ide-task
```
Creates the following files

```text
03-Tasks
└── 01-idea-task
    ├── Makefile
    └── README.md
```

### Efforts Usage
```shell
cd 01-idea-task
workutils  create effort bug-fix-this
```

Creates the following task
```text
03-Tasks
└── 01-idea-task
    ├── Efforts
    │        └── 01-bug-fix-this
    │            └── 01-worklog.md
    ├── Makefile
    └── README.md
```

As you go advance in your task, you might need to create more efforts as it goes on. To have a single note for the whole effort, you can run `make docs` and it will create a single `docs.md` that compiles all the worklogs together into a single file.

## Future work

I would like to be able to automatically archive projects without taking anything unecessary (e.g. python `venvs` or `node_modules`).