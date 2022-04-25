# Wrapper Scanner

This project was created to help on the  fix of the issue [824](https://github.com/buildpacks/lifecycle/issues/824) in 
the [CNB Buildpack Project](https://buildpacks.io/).

The code will search for all the tags of the lifecycle and for each version 0.x.y will run the command:

```bash
trivy image buildpacksio/lifecycle:[tag]
```

and show the output in the console.

## How to use it

- Install **Trivy** tool, check the documentation [here](https://aquasecurity.github.io/trivy/v0.27.0/getting-started/installation/)
- Compile running ``make`` in the root of the project
- A new folder **out** will be created, run the executable according to your OS (linux or MacOS)
```
.
└── out/
    ├── scanner-darwin
    └── scanner-linux
```

- The tool will show in the console the output per image tag of **Trivy**
