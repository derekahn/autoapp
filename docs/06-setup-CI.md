# Setting Up Continuous Integration

In this section we'll be finishing the CI flow, I say `finishing` because we've already configured [Quay](https://quay.io) to pull and build on any new tag on our [remote repo](https://github.com/derekahn/autoapp/pulls) so that's the second half. This will be the first half in which runs our `Go` tests and tests the 🐳 build on a pre-merge.

So there's some redundancy happening but I think is important and good if your mantra is `DON'T BREAK THE BUILD!`

## Sections

| Previous                                             | Next                                          |
| ---------------------------------------------------- | --------------------------------------------- |
| [Setting Up Spinnaker](05-setup-spinnaker.md) ||
