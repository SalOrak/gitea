	"code.gitea.io/gitea/models/db"
	Start, End                             string

			lastFile := createDiffFile(diff, line)
			diff.End = lastFile.Name
			_, err := io.Copy(io.Discard, reader)
				count, err := db.Count(m)
func GetDiffRangeWithWhitespaceBehavior(gitRepo *git.Repository, beforeCommitID, afterCommitID, skipTo string, maxLines, maxLineCharacters, maxFiles int, whitespaceBehavior string, directComparison bool) (*Diff, error) {
	argsLength := 6
	if len(whitespaceBehavior) > 0 {
		argsLength++
	}
	if len(skipTo) > 0 {
		argsLength++
	}

	diffArgs := make([]string, 0, argsLength)
		diffArgs = append(diffArgs, "diff", "--src-prefix=\\a/", "--dst-prefix=\\b/", "-M")
		diffArgs = append(diffArgs, "diff", "--src-prefix=\\a/", "--dst-prefix=\\b/", "-M")
	if skipTo != "" {
		diffArgs = append(diffArgs, "--skip-to="+skipTo)
	}
	cmd := exec.CommandContext(ctx, git.GitExecutable, diffArgs...)

	diff.Start = skipTo
			workdir, err := os.MkdirTemp("", "empty-work-dir")
					err := checker.Run()
	separator := "..."
	if directComparison {
		separator = ".."
	}

	shortstatArgs := []string{beforeCommitID + separator + afterCommitID}
func GetDiffCommitWithWhitespaceBehavior(gitRepo *git.Repository, commitID, skipTo string, maxLines, maxLineCharacters, maxFiles int, whitespaceBehavior string, directComparison bool) (*Diff, error) {
	return GetDiffRangeWithWhitespaceBehavior(gitRepo, "", commitID, skipTo, maxLines, maxLineCharacters, maxFiles, whitespaceBehavior, directComparison)