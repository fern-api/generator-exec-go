# Reference
## Logging
<details><summary><code>client.Logging.SendUpdate(TaskId, request) -> error</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := []*generatorexecgo.GeneratorUpdate{
    &generatorexecgo.GeneratorUpdate{
        Init: &generatorexecgo.InitUpdate{
            PackagesToPublish: []*generatorexecgo.PackageCoordinate{
                &generatorexecgo.PackageCoordinate{
                    Npm: &generatorexecgo.NpmCoordinate{
                        Name: "name",
                        Version: "version",
                    },
                },
                &generatorexecgo.PackageCoordinate{
                    Npm: &generatorexecgo.NpmCoordinate{
                        Name: "name",
                        Version: "version",
                    },
                },
            },
        },
    },
    &generatorexecgo.GeneratorUpdate{
        Init: &generatorexecgo.InitUpdate{
            PackagesToPublish: []*generatorexecgo.PackageCoordinate{
                &generatorexecgo.PackageCoordinate{
                    Npm: &generatorexecgo.NpmCoordinate{
                        Name: "name",
                        Version: "version",
                    },
                },
                &generatorexecgo.PackageCoordinate{
                    Npm: &generatorexecgo.NpmCoordinate{
                        Name: "name",
                        Version: "version",
                    },
                },
            },
        },
    },
}
err := client.Logging.SendUpdate(
    context.TODO(),
    "taskId",
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**taskId:** `generatorexecgo.TaskId` 
    
</dd>
</dl>

<dl>
<dd>

**request:** `[]*generatorexecgo.GeneratorUpdate` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Readme
<details><summary><code>client.Readme.GenerateReadme(TaskId, request) -> error</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &generatorexecgo.GenerateReadmeRequest{
    Title: "title",
    Summary: "summary",
    Requirements: []string{
        "requirements",
        "requirements",
    },
    Usage: "usage",
}
err := client.Readme.GenerateReadme(
    context.TODO(),
    "taskId",
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**taskId:** `generatorexecgo.TaskId` 
    
</dd>
</dl>

<dl>
<dd>

**request:** `*generatorexecgo.GenerateReadmeRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

