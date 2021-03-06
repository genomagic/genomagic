package visualizer

type (
	// Visualizer represents the interface that is used to interact with the visualizer package. It allows the
	// construction of specific visualization that can be created after performing assembly and reporting steps
	Visualizer interface {
		// Process creates the charts that are used for visualizing the results of GenoAssist
		Process() error
	}
)
