package repository

import (
	"strings"
)

func Clone(url, localPath string) error {
	builder := newCommandBuilder("")
	builder.Write("clone", url, localPath)

	_, _, err := runCommand(builder.String())
	if err != nil {
		return err
	}
	return nil
}

func Checkout(options CheckoutOptions, localPath string) error {
	builder := newCommandBuilder(localPath)
	builder.Write("checkout")
	if options.Create {
		builder.Write("-b")
	}
	builder.Write(options.Reference)

	_, _, err := runCommand(builder.String())
	if err != nil {
		return err
	}
	return nil
}

func Pull(options PullOptions, localPath string) error {
	builder := newCommandBuilder(localPath)
	builder.Write("pull")

	if options.Force {
		builder.Write("--force")
	}

	if options.Rebase {
		if options.RebaseOptions.Merges {
			builder.Write("--rebase=merges")
		} else {
			builder.Write("--rebase")
		}
	}

	builder.Write(options.RemoteUrl)
	builder.Write(options.Reference)

	_, _, err := runCommand(builder.String())
	if err != nil {
		return err
	}
	return nil
}

func Push(options PushOptions, localPath string) error {
	builder := newCommandBuilder(localPath)
	builder.Write("push")

	if options.Force {
		builder.Write("--force")
	}

	builder.Write(options.RemoteUrl)
	builder.Write(options.Reference)

	_, _, err := runCommand(builder.String())
	if err != nil {
		return err
	}
	return nil
}

func getTrimmedLineItems(out string) []string {
	items := make([]string, 0)

	for _, ln := range strings.Split(out, "\n") {
		trimmed := strings.Trim(ln, " ")
		if trimmed != "" {
			items = append(items, trimmed)
		}
	}
	return items
}

func GetBranches(localPath string) ([]string, error) {
	builder := newCommandBuilder(localPath)
	builder.Write("branch")
	builder.Write("-a")
	stdOut, _, err := runCommand(builder.String())
	if err != nil {
		return []string{}, err
	}

	branches := getTrimmedLineItems(stdOut)
	return branches, nil

}

func GetTags(localPath string) ([]string, error) {
	builder := newCommandBuilder(localPath)
	builder.Write("tag")

	stdOut, _, err := runCommand(builder.String())
	if err != nil {
		return []string{}, err
	}

	tags := getTrimmedLineItems(stdOut)
	return tags, nil
}
