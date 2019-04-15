# Deploying 🚀

This is the final section of our `how to CI/CD`. In this section we will:

0. Branch from `master` and create a `bug-fix` branch, change something (color, html, etc)
1. Then push said branch to the remote repo
1. Submit a `pull-request`
1. Configure github's premerge `CI` hook-which will run our drone service to run the first check.
1. Upon a successful code review with successful CI checks we will then push appropriate semantic tagging and letting [drone](https://drone.io) run again
1. Then upon accepting the `pull-request` we will then watch the magic ✨ happen!

## #0 branch off master

#### Branch off master

```bash
$ git branch

  * master
  (END)

$ git checkout -b fix-welcome-copy

$ git branch

  * fix-welcome-copy
    master
  (END)
```

#### Making a change

> Open `welcom.html`

```bash
$ vim web/template/welcome.html

```

> Change some copy; ie. `"Welcome"` -> `"Aloha"`

```html
<div class="welcome center">Aloha {{.Name}}, it is {{.Time}}</div>
```

## #1 commit changes and push to remote repo

```bash
# Check changes
$ git status

  On branch fix-welcome-copy
  Changes not staged for commit:
    (use "git add <file>..." to update what will be committed)
    (use "git checkout -- <file>..." to discard changes in working directory)

    modified: web/template/welcome.html

# Stage changes
$ git add web/template/welcome.html

# Commit changes with a descriptive message
$ git commit -m "Fixed the welcome.html copy to be more hawaiian"

  [deploy xxxxxxx] Fixed the welcome.html copy to be more hawaiian

$

```

```bash
# Double check that spinnaker is exposed
$ env | grep "DECK_POD"

  DECK_POD=spin-deck-67d875cd75-zpd4l

# IF NOT set env DECK_POD again
$ export DECK_POD=$(kubectl get pods --namespace default \
    -l "cluster=spin-deck" \
    -o jsonpath="{.items[0].metadata.name}")

# re-expose port and run in background process in this session
$ kubectl port-forward --namespace default $DECK_POD 8080:9000 >> /dev/null &
```

## Sections

| Previous                              |
| ------------------------------------- |
| [Setting up Drone](06-setup-drone.md) |
