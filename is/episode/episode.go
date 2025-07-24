// Package episode
package episode

type EpisodeProducer interface {
	SetupProcedure() error
	RunProcedure() error
	CleanupProcedure() error
}

func RunEpisode(e *EpisodeProducer) error {
	return nil
}
