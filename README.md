# go-codespaces-template
This is a template that is designed to be used with GitHub Codespaces to get you up and running with your Go server projects in a matter of minutes.

## Getting started

To get started, first you need to decide if you want to `fork` or `template` this repository. In simple terms, with a fork you will have a copy of this repository in your own GitHub account and you can make changes to it as you wish, as well as the ability to pull any updates made from this top level template repository. With a template, you will have a copy of this repository in your own GitHub account as it stands in its current form. You will be able to make changes to the code and make the project your own, but you will not be able to push those changes back to this template repository.

### Forking / Templating

![Forking or templating](./assets/fork-or-template.png)

### Creating a codespace

Once you have your own copy of this repository, you can create a codespace by clicking the green `Code` button and selecting `New Codespace`.

![Template codespace](./assets/template-button.png)

OR

![Codeing codespace](./assets/code-button.png)

Once you have chosen your method, you will be presented with a screen that looks like this:

![Codespace creation](./assets/codespace-setup.png)

Wait for this to finish and you will be presented with a screen that looks like VSCode, but in your browser.

Once this has opened you will need to wait just a moment longer for the post create command to finish setting up the environment. You will know this is done when you see the following in the terminal:

![Post create command](./assets/post-create-command.png)

Once this is done, you are ready to start coding!

## Running out of the box
run the commands:

```bash
go build
./go-codespace-template
```
