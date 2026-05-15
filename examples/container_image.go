package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/chalk-ai/chalk-go"
)

//lint:ignore U1000 example
func containerImageExample() {
	ctx := context.Background()

	client, err := chalk.NewClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create Chalk client: %v", err)
	}

	image := chalk.DebianSlimImage().
		PipInstall("requests").
		RunCommands("apt-get update && apt-get install -y curl").
		AddFile(
			"/app/main.py",
			[]byte(`import os
import time

print("hello from Chalk")
print("LOG_LEVEL=", os.getenv("LOG_LEVEL", ""))
print("scratch volume is mounted at /scratch")
time.sleep(3600)
`),
		).
		Workdir("/app").
		Entrypoint("python").
		Cmd("/app/main.py").
		EnvVar("PYTHONDONTWRITEBYTECODE", "1")

	container := chalk.NewContainer("go-image-example", image).
		WithCPU("1").
		WithMemory("2Gi").
		WithEnv("LOG_LEVEL", "debug").
		WithPort(8080).
		WithLifetime(2 * time.Hour).
		WithVolume(chalk.EmptyDirVolume("scratch", "/scratch", "2Gi")).
		WithVolume(chalk.SharedMemoryVolume("dshm", "/dev/shm", "1Gi")).
		WithRouting(chalk.ContainerRoutingPrivate)

	info, err := client.RunContainer(
		ctx,
		container,
		chalk.WithRunContainerBuildOptions(
			chalk.WithImageBuildPollInterval(2*time.Second),
			chalk.WithImageBuildTimeout(15*time.Minute),
		),
	)
	if err != nil {
		log.Fatalf("Failed to run container: %v", err)
	}

	fmt.Printf("Container %s is %s\n", info.ID, info.Status)
	if info.WebURL != "" {
		fmt.Printf("Web URL: %s\n", info.WebURL)
	}

	fresh, err := client.GetContainer(ctx, chalk.ContainerID(info.ID))
	if err != nil {
		log.Fatalf("Failed to refresh container: %v", err)
	}
	fmt.Printf("Refreshed status: %s\n", fresh.Status)

	stopped, err := client.StopContainer(
		ctx,
		chalk.ContainerID(info.ID),
		chalk.WithStopGracePeriod(30*time.Second),
	)
	if err != nil {
		log.Fatalf("Failed to stop container: %v", err)
	}
	fmt.Printf("Stopped container %s with status %s\n", stopped.ID, stopped.Status)
}
