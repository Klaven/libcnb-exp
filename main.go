package main

import (
	"context"
	"fmt"

	"github.com/buildpacks/pack/pkg/client"
	"github.com/buildpacks/pack/pkg/image"
	"github.com/sirupsen/logrus"
)

func main() {

	ctx := context.Background()

	c, err := client.NewClient()
	if err != nil {
		logrus.Error(err)
	}

	envs := make(map[string]string)

	envs["CGO_ENABLED"] = "0"
	envs["BP_GO_BUILD_FLAGS"] = "-buildmode=default"

	// should be the equivalent of  pack build my-app  --buildpack paketo-buildpacks/go  --builder paketobuildpacks/builder-jammy-buildpackless-static
	buildOpts := client.BuildOptions{
		GroupID:    -1,
		Image:      "web-example",
		Builder:    "paketobuildpacks/builder-jammy-full",
		Buildpacks: []string{"paketo-buildpacks/web-servers"},
		PullPolicy: image.PullIfNotPresent,
		AppPath:    "./sample/public/",
		// DockerHost: "unix:///var/run/docker.sock",
		Env: map[string]string{
			"BP_WEB_SERVER": "nginx",
		},
	}

	err = c.Build(ctx, buildOpts)
	if err != nil {
		logrus.Error(err)
	}

	cli, err := client.NewClient()

	buildOpts = client.BuildOptions{
		GroupID:    -1,
		Image:      "web-example-2",
		Builder:    "paketobuildpacks/builder-jammy-full",
		Buildpacks: []string{"paketo-buildpacks/web-servers"},
		PullPolicy: image.PullIfNotPresent,
		AppPath:    "./sample/public/",
		// DockerHost: "unix:///var/run/docker.sock",
		Env: map[string]string{
			"BP_WEB_SERVER": "nginx",
		},
	}

	// Will error here
	err = cli.Build(ctx, buildOpts)
	if err != nil {
		logrus.Error(err)
		fmt.Println(err.Error())
	}

}
