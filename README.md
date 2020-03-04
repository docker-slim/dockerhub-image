# DockerHub Image Info

This is a very simple DockerHub client library designed to fetch the target image information including the image star and pull stats you don't get from the Docker Registry API.

## Example

```
import (
	"fmt"

	"github.com/docker-slim/dockerhub-image/info"
)
```

```
result, err := info.Get("dockerhub_user_name", "dockerhub_repo_name")
if err != nil {
	panic(err)
}

fmt.Println("pull count =", result.PullCount)
```
