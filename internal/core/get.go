package core

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	"go.redsock.ru/protopack/internal/adapters/repository/git"
	"go.redsock.ru/protopack/internal/core/models"
)

// Get download package.
func (c *Core) Get(ctx context.Context, requestedModule models.Module) error {
	cacheRepositoryDir, err := c.storage.CreateCacheRepositoryDir(requestedModule.Name)
	if err != nil {
		return fmt.Errorf("c.storage.CreateCacheRepositoryDir: %w", err)
	}

	// TODO: use factory (git, svn etc)
	repo, err := git.New(ctx, requestedModule.Name, cacheRepositoryDir, c.console)
	if err != nil {
		return fmt.Errorf("git.New: %w", err)
	}

	revision, err := repo.ReadRevision(ctx, requestedModule.Version)
	if err != nil {
		return fmt.Errorf("repository.ReadRevision: %w", err)
	}

	cacheDownloadPaths := c.storage.GetCacheDownloadPaths(requestedModule, revision)

	if err := c.storage.CreateCacheDownloadDir(cacheDownloadPaths); err != nil {
		return fmt.Errorf("c.storage.CreateCacheDownloadDir: %w", err)
	}

	if err := repo.Fetch(ctx, revision); err != nil {
		return fmt.Errorf("repository.Fetch: %w", err)
	}

	moduleConfig, err := c.moduleConfig.ReadFromRepo(ctx, repo, revision)
	if err != nil {
		return fmt.Errorf("c.moduleConfig.Read: %w", err)
	}

	for _, indirectDep := range moduleConfig.Dependencies {
		isInstalled, err := c.storage.IsModuleInstalled(indirectDep)
		if err != nil {
			return fmt.Errorf("c.storage.IsModuleInstalled: %w", err)
		}

		if isInstalled {
			continue
		}

		if err := c.Get(ctx, indirectDep); err != nil {
			if errors.Is(err, models.ErrVersionNotFound) {
				slog.Error("Version not found", "dependency", indirectDep)
				return models.ErrVersionNotFound
			}

			return fmt.Errorf("c.Get: %w", err)
		}
	}

	// check package deps (that was read from repo)
	// compare versions

	if err := repo.Archive(ctx, revision, cacheDownloadPaths); err != nil {
		return fmt.Errorf("repository.Archive: %w", err)
	}

	moduleHash, err := c.storage.Install(cacheDownloadPaths, requestedModule, revision, moduleConfig)
	if err != nil {
		return fmt.Errorf("c.storage.Install: %w", err)
	}

	slog.Debug("HASH", "hash", moduleHash)

	if err := c.lockFile.Write(requestedModule.Name, revision.Version, moduleHash); err != nil {
		return fmt.Errorf("c.lockFile.Write: %w", err)
	}

	return nil
}
