package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/google/go-github/github"
)

func setConfig(ctx context.Context, url, token string) {

	var (
		args = []string{
			"sudo", "-H", "-u", "github", "bash", "-c",
			"echo -ne '\n\n\nY\n' | ./config.sh --url " + url + " --token " + token,
		}
		cmd = exec.CommandContext(ctx, args[0], args[1:]...)
	)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Println(err.Error())
		return
	}

	if err := cmd.Start(); err != nil {
		log.Println("fail to start log: " + err.Error())
		return
	}

	scanner := bufio.NewScanner(stdout)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		m := scanner.Text()
		exec.Command("/bin/sh", "-c", "echo '"+m+"' >> logs.log").Run()
	}

	if err := cmd.Wait(); err != nil {
		log.Println(err.Error())
	}
}

func run(ctx context.Context) {

	var cmd = exec.CommandContext(ctx,
		"sudo", "-H", "-u", "github", "bash", "-c",
		"./run.sh",
	)

	log.Printf("run")
	defer log.Printf("exit run")

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Println(err.Error())
		return
	}

	if err := cmd.Start(); err != nil {
		log.Println("fail to start log: " + err.Error())
		return
	}

	scanner := bufio.NewScanner(stdout)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		m := scanner.Text()
		log.Println(m)
	}

	if err := cmd.Wait(); err != nil {
		log.Println(err.Error())
	}
}

func loadRelease() error {

	var (
		client = github.NewClient(nil)
		opt    = &github.ListOptions{Page: 1, PerPage: 1}
		ctx    = context.Background()
	)

	releases, _, err := client.Repositories.ListReleases(ctx, "actions", "runner", opt)
	if err != nil {
		return err
	} else if len(releases) != 1 {
		return fmt.Errorf("expected one release but received: %d", len(releases))
	}

	url := fmt.Sprintf("https://github.com/actions/runner/releases/download/%s/actions-runner-linux-x64-%s.tar.gz", *releases[0].TagName, (*releases[0].TagName)[1:])
	log.Println("load", url)
	if err := exec.Command("curl", "-o", "actions-runner.tar.gz", "-L", url).Run(); err != nil {
		return fmt.Errorf("fail to dowload runner: %w", err)
	}

	if err := exec.Command("tar", "xzf", "actions-runner.tar.gz").Run(); err != nil {
		return fmt.Errorf("fail to unzip: %w", err)
	}

	if err := exec.Command("./bin/installdependencies.sh").Run(); err != nil {
		return fmt.Errorf("fail to install: %w", err)
	}

	if err := exec.Command("chmod", "-R", "777", "/actions-runner").Run(); err != nil {
		return fmt.Errorf("fail to change mod: %w", err)
	}

	return nil
}

func main() {

	if err := loadRelease(); err != nil {
		log.Println("fail to load release", err.Error())
		return
	}

	var (
		filled bool
		ctx    = context.Background()
	)
	for _, key := range []string{"credentials", "credentials_rsaparams", "runner"} {
		if val := os.Getenv(key); val != "" {

			var (
				name   = "." + key
				f, err = os.Create(name)
			)
			if err != nil {
				log.Println(err)
				filled = false
				break
			}

			if _, err := f.Write([]byte(val)); err != nil {
				log.Println(err)
				filled = false
				break
			}

			if err := f.Close(); err != nil {
				log.Println(err)
				filled = false
				break
			}

			filled = true
		} else {
			filled = false
			break
		}
	}

	if !filled {
		setConfig(ctx, os.Getenv("URL"), os.Getenv("TOKEN"))
	}

	run(ctx)
}
