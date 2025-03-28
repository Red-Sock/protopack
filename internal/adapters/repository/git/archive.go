package git

import (
	"context"
	"fmt"

	"go.redsock.ru/protopack/internal/core/models"
)

func (r *gitRepo) Archive(
	ctx context.Context, revision models.Revision, cacheDownloadPaths models.CacheDownloadPaths,
) error {
	params := []string{
		"archive", "--format=zip", revision.CommitHash, "-o", cacheDownloadPaths.ArchiveFile, "*.proto",
	}

	if _, err := r.console.RunCmd(ctx, r.cacheDir, "git", params...); err != nil {
		return fmt.Errorf("utils.RunCmd: %w", err)
	}

	return nil
}
